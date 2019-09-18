package account

import (
	"context"
	"strings"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	apigatewayv1alpha1 "awsctrl.io/pkg/apis/apigateway/v1alpha1"
	cloudformationv1alpha1 "awsctrl.io/pkg/apis/cloudformation/v1alpha1"
	metav1alpha1 "awsctrl.io/pkg/apis/meta/v1alpha1"

	cfnencoder "awsctrl.io/pkg/encoding/cloudformation"

	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/types"

	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/controller"
	"sigs.k8s.io/controller-runtime/pkg/handler"
	"sigs.k8s.io/controller-runtime/pkg/manager"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"
	logf "sigs.k8s.io/controller-runtime/pkg/runtime/log"
	"sigs.k8s.io/controller-runtime/pkg/source"
)

var (
	log = logf.Log.WithName("account-controller")

	stackDeletionFinalizerName = "stack.cloudformation.awsctrl.io/deletion"
)

// Add creates a new Account Controller and adds it to the Manager with default RBAC. The Manager will set fields on the Controller
// and Start it when the Manager is Started.
func Add(mgr manager.Manager) error {
	return add(mgr, newReconciler(mgr))
}

// newReconciler returns a new reconcile.Reconciler
func newReconciler(mgr manager.Manager) reconcile.Reconciler {
	return &ReconcileAccount{Client: mgr.GetClient(), scheme: mgr.GetScheme()}
}

// add adds a new Controller to mgr with r as the reconcile.Reconciler
func add(mgr manager.Manager, r reconcile.Reconciler) error {
	// Create a new controller
	c, err := controller.New("account-controller", mgr, controller.Options{Reconciler: r})
	if err != nil {
		return err
	}

	// Watch for changes to Account
	err = c.Watch(&source.Kind{Type: &apigatewayv1alpha1.Account{}}, &handler.EnqueueRequestForObject{})
	if err != nil {
		return err
	}

	// Watch for changes to Stacks
	err = c.Watch(&source.Kind{Type: &cloudformationv1alpha1.Stack{}}, &handler.EnqueueRequestForOwner{
		IsController: true,
		OwnerType:    &apigatewayv1alpha1.Account{},
	})
	if err != nil {
		return err
	}

	return nil
}

var _ reconcile.Reconciler = &ReconcileAccount{}

// ReconcileAccount reconciles a Account object
type ReconcileAccount struct {
	client.Client
	scheme *runtime.Scheme
}

// Reconcile reads that state of the cluster for a Account object and makes changes based on the state read
// +kubebuilder:rbac:groups=cloudformation.awsctrl.io,resources=stacks,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=cloudformation.awsctrl.io,resources=stacks/status,verbs=get;update;patch
// +kubebuilder:rbac:groups=apigateway.awsctrl.io,resources=accounts,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=apigateway.awsctrl.io,resources=accounts/status,verbs=get;update;patch
func (r *ReconcileAccount) Reconcile(request reconcile.Request) (reconcile.Result, error) {
	ctx := context.Background()

	var instance apigatewayv1alpha1.Account
	if err := r.Get(ctx, request.NamespacedName, &instance); err != nil {
		if errors.IsNotFound(err) {
			return reconcile.Result{}, nil
		}
		return reconcile.Result{}, err
	}

	if instance.Spec.StackName == "" {
		if err := r.addStack(ctx, &instance); err != nil {
			return reconcile.Result{}, err
		}
		return reconcile.Result{}, nil
	}

	if instance.Status.Status == metav1alpha1.CreateCompleteStatus || instance.Status.Status == metav1alpha1.UpdateCompleteStatus {
		if err := r.updateStack(ctx, &instance); err != nil {
			return reconcile.Result{}, err
		}
		return reconcile.Result{}, nil
	}

	if instance.Status.StackID == "" {
		if err := r.updateStatus(ctx, &instance); err != nil {
			return reconcile.Result{}, err
		}
		return reconcile.Result{}, nil
	}

	return reconcile.Result{}, nil
}

func (r *ReconcileAccount) addStack(ctx context.Context, instance *apigatewayv1alpha1.Account) error {
	stack, err := r.newAccountStack(instance)
	if err != nil {
		return err
	}

	if err := r.Create(ctx, stack); err != nil {
		return err
	}

	instance.Spec.StackName = stack.Name
	return r.Update(ctx, instance)
}

func (r *ReconcileAccount) updateStack(ctx context.Context, instance *apigatewayv1alpha1.Account) error {
	stack, err := r.newAccountStack(instance)
	if err != nil {
		return err
	}

	return r.Update(ctx, stack)
}

func (r *ReconcileAccount) updateStatus(ctx context.Context, instance *apigatewayv1alpha1.Account) error {
	var stack cloudformationv1alpha1.Stack
	if err := r.Get(ctx, types.NamespacedName{Name: instance.Spec.StackName, Namespace: instance.Namespace}, &stack); err != nil {
		return err
	}

	instance.Status.StatusMeta = *stack.Status.StatusMeta.DeepCopy()
	return r.Status().Update(ctx, instance)
}

func accountStackName(name, namespace string) string {
	return strings.Join([]string{"apigateway", "account", namespace, name}, "-")
}

func (r *ReconcileAccount) newAccountStack(instance *apigatewayv1alpha1.Account) (*cloudformationv1alpha1.Stack, error) {
	params := map[string]string{}
	cfnencoder.MarshalTypes(params, instance.Spec, "Parameter")

	stack := &cloudformationv1alpha1.Stack{
		ObjectMeta: metav1.ObjectMeta{
			Name:      accountStackName(instance.Name, instance.Namespace),
			Namespace: instance.Namespace,
		},
		Spec: cloudformationv1alpha1.StackSpec{
			CloudFormationMeta: metav1alpha1.CloudFormationMeta{
				Region:                instance.Spec.Region,
				NotificationARNs:      instance.Spec.NotificationARNs,
				OnFailure:             instance.Spec.OnFailure,
				Tags:                  instance.Spec.Tags,
				TerminationProtection: instance.Spec.TerminationProtection,
			},
			Parameters:   params,
			TemplateBody: instance.GetTemplate(),
		},
	}

	if err := ctrl.SetControllerReference(instance, stack, r.scheme); err != nil {
		return nil, err
	}
	return stack, nil
}
