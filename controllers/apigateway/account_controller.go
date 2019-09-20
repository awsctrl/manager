/*

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package apigateway

import (
	"context"
	"strings"
	"time"

	"github.com/go-logr/logr"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/types"

	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	apigatewayv1alpha1 "awsctrl.io/apis/apigateway/v1alpha1"
	cloudformationv1alpha1 "awsctrl.io/apis/cloudformation/v1alpha1"
	metav1alpha1 "awsctrl.io/apis/meta/v1alpha1"
	controllerutils "awsctrl.io/controllers/utils"
	cfnencoder "awsctrl.io/encoding/cloudformation"
)

var (
	// APIGVStr returns the group version for the apigateway resource
	APIGVStr = apigatewayv1alpha1.GroupVersion.String()
)

// AccountReconciler reconciles a Account object
type AccountReconciler struct {
	client.Client
	Log    logr.Logger
	Scheme *runtime.Scheme
}

// Load the Cloudformation Stack resource
// +kubebuilder:rbac:groups=cloudformation.awsctrl.io,resources=stacks,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=cloudformation.awsctrl.io,resources=stacks/status,verbs=get;update;patch

// Load the Apigateway Account resource
// +kubebuilder:rbac:groups=apigateway.awsctrl.io,resources=accounts,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=apigateway.awsctrl.io,resources=accounts/status,verbs=get;update;patch

// Reconcile will make the desired state a reality
func (r *AccountReconciler) Reconcile(req ctrl.Request) (ctrl.Result, error) {
	ctx := context.Background()
	log := r.Log.WithValues("Account", req.NamespacedName)
	waitDuration := time.Duration(2 * time.Second)

	var instance apigatewayv1alpha1.Account
	if err := r.Get(ctx, req.NamespacedName, &instance); err != nil {
		if errors.IsNotFound(err) {
			return ctrl.Result{}, nil
		}
		return ctrl.Result{}, err
	}

	if instance.Spec.StackName == "" {
		log.Info("Create CloudFormation Stack")
		if err := r.addStack(ctx, &instance); err != nil {
			return ctrl.Result{}, err
		}

		return ctrl.Result{RequeueAfter: waitDuration}, r.updateTemplateVersion(ctx, req.NamespacedName)
	}

	if controllerutils.IsStatusComplete(instance.Status.Status) &&
		templateVersionChanged(&instance) {
		log.Info("Update CloudFormation Stack")
		if err := r.updateStack(ctx, &instance); err != nil {
			return ctrl.Result{}, err
		}

		return ctrl.Result{}, r.updateTemplateVersion(ctx, req.NamespacedName)
	}

	log.Info("Update Status")
	return reconcile.Result{}, r.updateStatus(ctx, &instance)
}

// SetupWithManager will setup the controller
func (r *AccountReconciler) SetupWithManager(mgr ctrl.Manager) error {
	if err := mgr.GetFieldIndexer().IndexField(&cloudformationv1alpha1.Stack{}, controllerutils.ControllerOwnerKey, func(rawObj runtime.Object) []string {
		stack := rawObj.(*cloudformationv1alpha1.Stack)
		owner := metav1.GetControllerOf(stack)
		if owner == nil {
			return nil
		}
		if owner.APIVersion != APIGVStr || owner.Kind != "Account" {
			return nil
		}

		return []string{owner.Name}
	}); err != nil {
		return err
	}

	return ctrl.NewControllerManagedBy(mgr).
		For(&apigatewayv1alpha1.Account{}).
		Owns(&cloudformationv1alpha1.Stack{}).
		Complete(r)
}

func (r *AccountReconciler) addStack(ctx context.Context, instance *apigatewayv1alpha1.Account) error {
	stack, err := r.newAccountStack(instance)
	if err != nil {
		return err
	}

	if err := r.Create(ctx, stack); err != nil {
		return err
	}

	instanceCopy := instance.DeepCopy()
	instanceCopy.Spec.StackName = stack.Name

	return r.Update(ctx, instanceCopy)
}

func (r *AccountReconciler) updateStack(ctx context.Context, instance *apigatewayv1alpha1.Account) error {
	params := map[string]string{}
	cfnencoder.MarshalTypes(params, instance.Spec, "Parameter")

	var stack cloudformationv1alpha1.Stack
	if err := r.Get(ctx, types.NamespacedName{Name: instance.Spec.StackName, Namespace: instance.Namespace}, &stack); err != nil {
		return err
	}

	stackCopy := stack.DeepCopy()
	stackCopy.Spec.Parameters = params
	stackCopy.Spec.TemplateBody = instance.GetTemplate()
	stackCopy.Spec.CloudFormationMeta = metav1alpha1.CloudFormationMeta{
		Region:                instance.Spec.Region,
		NotificationARNs:      instance.Spec.NotificationARNs,
		OnFailure:             instance.Spec.OnFailure,
		Tags:                  instance.Spec.Tags,
		TerminationProtection: instance.Spec.TerminationProtection,
	}

	return r.Update(ctx, stackCopy)
}

func (r *AccountReconciler) updateStatus(ctx context.Context, instance *apigatewayv1alpha1.Account) error {
	var stack cloudformationv1alpha1.Stack
	if err := r.Get(ctx, types.NamespacedName{Name: instance.Spec.StackName, Namespace: instance.Namespace}, &stack); err != nil {
		return err
	}

	instance.Status.StatusMeta = *stack.Status.StatusMeta.DeepCopy()
	return r.Status().Update(ctx, instance)
}

func (r *AccountReconciler) updateTemplateVersion(ctx context.Context, namespaceName types.NamespacedName) error {
	var instance apigatewayv1alpha1.Account
	if err := r.Get(ctx, namespaceName, &instance); err != nil {
		return err
	}
	instanceCopy := instance.DeepCopy()

	if len(instanceCopy.Labels) == 0 {
		instanceCopy.Labels = map[string]string{}
	}

	instanceCopy.Labels[controllerutils.StackTemplateVersionLabel] = controllerutils.ComputeHash(instanceCopy.Spec)

	return r.Update(ctx, instanceCopy)
}

func templateVersionChanged(instance *apigatewayv1alpha1.Account) bool {
	return instance.Labels[controllerutils.StackTemplateVersionLabel] != controllerutils.ComputeHash(instance.Spec)
}

func stackName(name, namespace string) string {
	return strings.Join([]string{"apigateway", "account", namespace, name}, "-")
}

func (r *AccountReconciler) newAccountStack(instance *apigatewayv1alpha1.Account) (*cloudformationv1alpha1.Stack, error) {
	params := map[string]string{}
	cfnencoder.MarshalTypes(params, instance.Spec, "Parameter")

	stack := &cloudformationv1alpha1.Stack{
		ObjectMeta: metav1.ObjectMeta{
			Name:      stackName(instance.Name, instance.Namespace),
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

	if err := ctrl.SetControllerReference(instance, stack, r.Scheme); err != nil {
		return nil, err
	}
	return stack, nil
}
