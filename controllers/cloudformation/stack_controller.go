/*
Copyright © 2019 AWS Controller authors

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

package cloudformation

import (
	"context"
	"time"

	"github.com/go-logr/logr"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/types"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"

	cloudformationv1alpha1 "awsctrl.io/apis/cloudformation/v1alpha1"
	metav1alpha1 "awsctrl.io/apis/meta/v1alpha1"
	selfv1alpha1 "awsctrl.io/apis/self/v1alpha1"
	"awsctrl.io/aws"
	cloudformationutils "awsctrl.io/controllers/cloudformation/utils"
	"awsctrl.io/controllers/utils"
	"awsctrl.io/token"
)

// StackReconciler reconciles a Stack object
type StackReconciler struct {
	client.Client
	Log          logr.Logger
	Scheme       *runtime.Scheme
	ConfigName   string
	PodNamespace string
	AWSClient    aws.AWS
	TokenClient  token.Token
}

// Load the Self Config for the Controller
// +kubebuilder:rbac:groups=self.awsctrl.io,resources=configs,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=self.awsctrl.io,resources=configs/status,verbs=get;update;patch

// Load the Cloudformation Stack resource
// +kubebuilder:rbac:groups=cloudformation.awsctrl.io,resources=stacks,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=cloudformation.awsctrl.io,resources=stacks/status,verbs=get;update;patch

// Reconcile will make the desired state a reality
func (r *StackReconciler) Reconcile(req ctrl.Request) (ctrl.Result, error) {
	ctx := context.Background()
	log := r.Log.WithValues("Stack", req.NamespacedName)
	waitDuration := time.Duration(2 * time.Second)

	log.Info("Starting reconcile")
	defer log.Info("Finished reconciling")

	if err := r.AWSClient.Configured(); err != nil {
		return ctrl.Result{Requeue: true}, err
	}

	var config selfv1alpha1.Config
	if err := r.Get(ctx, types.NamespacedName{Namespace: r.PodNamespace, Name: r.ConfigName}, &config); err != nil {
		return ctrl.Result{}, err
	}

	var instance cloudformationv1alpha1.Stack
	if err := r.Get(ctx, req.NamespacedName, &instance); err != nil {
		if errors.IsNotFound(err) {
			return ctrl.Result{}, nil
		}
		return ctrl.Result{}, err
	}

	if instance.Status.Status == metav1alpha1.DeleteCompleteStatus {
		return ctrl.Result{}, r.removeCFNFinalizer(ctx, &instance)
	}

	if !instance.DeletionTimestamp.IsZero() {
		if instance.Status.Status == metav1alpha1.DeleteInProgressStatus {
			return ctrl.Result{RequeueAfter: waitDuration}, r.describeCFNStackStatus(ctx, &instance)
		}

		log.Info("Deleting CloudFormation Stack")
		return ctrl.Result{RequeueAfter: waitDuration}, r.deleteCFNStack(ctx, &instance)
	}

	if ok := utils.ContainsFinalizer(instance.ObjectMeta, cloudformationutils.StackDeletionFinalizerName); !ok {
		log.Info("Adding Stack Finalizer")
		if err := r.addCFNFinalizer(ctx, &instance); err != nil {
			return ctrl.Result{}, err
		}
		return ctrl.Result{}, nil
	}

	// TODO: move to defaulting admission webhook
	if instance.Spec.ClientRequestToken == "" {
		log.Info("Adding ClientRequestToken")
		return ctrl.Result{}, r.generateClientRequestToken(ctx, &instance)
	}

	// TODO: move to defaulting admission webhook
	if !utils.SliceContains(instance.GetNotificationARNs(), config.Spec.AWS.Queue.TopicARN) {
		log.Info("Adding controller NotificationARNs")
		return ctrl.Result{}, r.addNotificationARN(ctx, &instance, &config)
	}

	if instance.Status.StackID == "" {
		log.Info("Creating CFN Stack")
		return ctrl.Result{RequeueAfter: waitDuration}, r.createCFNStack(ctx, &instance)
	}

	if utils.IsStatusComplete(instance.Status.Status) &&
		cloudformationutils.TemplateVersionChanged(&instance) {
		log.Info("Updating CFN Stack")
		return ctrl.Result{RequeueAfter: waitDuration}, r.updateCFNStack(ctx, &instance)
	}

	if instance.Status.StackID != "" {
		log.Info("Describing Stack to update Status")
		if utils.IsStatusComplete(instance.Status.Status) {
			waitDuration = time.Duration(0)
		}

		return ctrl.Result{RequeueAfter: waitDuration}, r.describeCFNStackStatus(ctx, &instance)
	}

	return ctrl.Result{}, nil

}

// SetupWithManager will setup the controller
func (r *StackReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&cloudformationv1alpha1.Stack{}).
		Complete(r)
}
