package stack

import (
	"context"
	"os"

	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/types"

	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/controller"
	"sigs.k8s.io/controller-runtime/pkg/handler"
	"sigs.k8s.io/controller-runtime/pkg/manager"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"
	"sigs.k8s.io/controller-runtime/pkg/source"

	logf "sigs.k8s.io/controller-runtime/pkg/runtime/log"

	cloudformationv1alpha1 "awsctrl.io/pkg/apis/cloudformation/v1alpha1"
	metav1alpha1 "awsctrl.io/pkg/apis/meta/v1alpha1"
	selfv1alpha1 "awsctrl.io/pkg/apis/self/v1alpha1"

	"awsctrl.io/pkg/aws"
	"awsctrl.io/pkg/controllerutils"
	"awsctrl.io/pkg/token"
)

var (
	log = logf.Log.WithName("stack-controller")

	stackDeletionFinalizerName = "stack.cloudformation.awsctrl.io/deletion"
)

// Add creates a new Stack Controller and adds it to the Manager with default RBAC. The Manager will set fields on the Controller
// and Start it when the Manager is Started.
func Add(mgr manager.Manager) error {
	return add(mgr, NewReconciler(mgr, aws.New(), token.New(), os.Getenv("POD_NAMESPACE")))
}

// NewReconciler returns a new reconcile.Reconciler
func NewReconciler(mgr manager.Manager, awsclient aws.AWS, tokenclient token.Token, podNamespace string) reconcile.Reconciler {
	return &ReconcileStack{
		Client:       mgr.GetClient(),
		scheme:       mgr.GetScheme(),
		aws:          awsclient,
		clientToken:  tokenclient,
		PodNamespace: podNamespace,
	}
}

// add adds a new Controller to mgr with r as the reconcile.Reconciler
func add(mgr manager.Manager, r reconcile.Reconciler) error {
	// Create a new controller
	c, err := controller.New("stack-controller", mgr, controller.Options{Reconciler: r})
	if err != nil {
		return err
	}

	// Watch for changes to Stack
	err = c.Watch(&source.Kind{Type: &cloudformationv1alpha1.Stack{}}, &handler.EnqueueRequestForObject{})
	if err != nil {
		return err
	}

	// Watch a Config this is what connects everything to AWS
	err = c.Watch(&source.Kind{Type: &selfv1alpha1.Config{}}, &handler.EnqueueRequestForObject{})
	if err != nil {
		return err
	}

	_ = mgr.GetFieldIndexer().IndexField(&cloudformationv1alpha1.Stack{}, "spec.status.stackID", func(o runtime.Object) []string {
		var res []string
		instance := o.(*cloudformationv1alpha1.Stack)
		res = append(res, instance.Status.StatusMeta.StackID)
		return res
	})

	return nil
}

var _ reconcile.Reconciler = &ReconcileStack{}

// ReconcileStack reconciles a Stack object
type ReconcileStack struct {
	client.Client
	scheme *runtime.Scheme

	aws         aws.AWS
	clientToken token.Token

	PodNamespace string
}

// Reconcile reads that state of the cluster for a Stack object and makes changes based on the state read
// and what is in the Stack.Spec
// Automatically generate RBAC rules to allow the Controller to read and write Deployments
// +kubebuilder:rbac:groups=self.awsctl.io,resources=config,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=cloudformation.awsctrl.io,resources=stacks,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=cloudformation.awsctrl.io,resources=stacks/status,verbs=get;update;patch
func (r *ReconcileStack) Reconcile(request reconcile.Request) (reconcile.Result, error) {
	ctx := context.Background()

	var config selfv1alpha1.Config
	if err := r.Get(ctx, types.NamespacedName{Name: "config", Namespace: r.PodNamespace}, &config); err != nil {
		return reconcile.Result{}, err
	}

	// Configure the AWS Client
	if err := r.aws.Configure(&config.Spec.AWS); err != nil {
		return reconcile.Result{}, err
	}

	var instance cloudformationv1alpha1.Stack
	if err := r.Get(ctx, request.NamespacedName, &instance); err != nil {
		if errors.IsNotFound(err) {
			return reconcile.Result{}, nil
		}
		return reconcile.Result{}, err
	}

	if !instance.DeletionTimestamp.IsZero() {
		return reconcile.Result{}, r.deleteCFNStack(ctx, &instance)
	}

	if ok := controllerutils.ContainsFinalizer(instance.ObjectMeta, stackDeletionFinalizerName); !ok {
		if err := r.addCFNFinalizer(ctx, &instance); err != nil {
			return reconcile.Result{}, err
		}

		return reconcile.Result{}, nil
	}

	if instance.Spec.ClientRequestToken == "" {
		if err := r.generateClientRequestToken(ctx, &instance); err != nil {
			return reconcile.Result{}, err
		}
		return reconcile.Result{}, nil
	}

	notifSlice := []string{}
	for _, notifArn := range instance.Spec.NotificationARNs {
		notifSlice = append(notifSlice, *notifArn)
	}

	if !controllerutils.SliceContains(notifSlice, config.Spec.AWS.Queue.TopicARN) {
		if err := r.addNotificationARN(ctx, &instance, &config); err != nil {
			return reconcile.Result{}, err
		}
		return reconcile.Result{}, nil
	}

	if instance.Status.StackID == "" {
		if err := r.createCFNStack(ctx, &instance); err != nil {
			return reconcile.Result{}, err
		}
		return reconcile.Result{}, nil
	}

	if instance.Status.Status == metav1alpha1.CreateCompleteStatus || instance.Status.Status == metav1alpha1.UpdateCompleteStatus {
		if err := r.updateCFNStack(ctx, &instance); err != nil {
			return reconcile.Result{}, err
		}
	}

	if instance.Status.StackID != "" {
		if err := r.describeCFNStackStatus(ctx, &instance); err != nil {
			return reconcile.Result{}, err
		}
	}

	return reconcile.Result{}, nil
}
