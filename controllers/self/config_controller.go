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

package self

import (
	"context"
	"time"

	"github.com/go-logr/logr"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/client-go/util/retry"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"

	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	selfv1alpha1 "go.awsctrl.io/manager/apis/self/v1alpha1"
	"go.awsctrl.io/manager/aws"
	"go.awsctrl.io/manager/controllers/utils"
)

var (
	configDeletionFinalizerName = "config.self.awsctrl.io/deletion"
	loadedAWSConfig             = "config.self.awsctrl.io/aws-config-version"
)

// ConfigReconciler reconciles a Config object
type ConfigReconciler struct {
	client.Client
	Log          logr.Logger
	Scheme       *runtime.Scheme
	ConfigName   string
	PodNamespace string
	AWSClient    aws.AWS
}

// +kubebuilder:rbac:groups=self.awsctrl.io,resources=config,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=self.awsctrl.io,resources=config/status,verbs=get;update;patch

// Reconcile will make the desired state a reality
func (r *ConfigReconciler) Reconcile(req ctrl.Request) (ctrl.Result, error) {
	ctx := context.Background()
	log := r.Log.WithValues("Config", req.NamespacedName)

	log.Info("Starting reconcile")
	defer log.Info("Finished reconciling")

	nsn := types.NamespacedName{Namespace: r.PodNamespace, Name: r.ConfigName}

	var config selfv1alpha1.Config
	if err := r.Get(ctx, nsn, &config); err != nil {
		return ctrl.Result{}, err
	}

	if len(config.Status.Conditions) == 0 {
		if err := r.updateConfigConditions(ctx, nsn, selfv1alpha1.ConfigConditionPendingAWSConfiguration, "", ""); err != nil {
			return ctrl.Result{}, err
		}
		return ctrl.Result{Requeue: true}, nil
	}

	awsConfigVersion := utils.ComputeHash(config.Spec.AWS)
	if config.Labels[loadedAWSConfig] != awsConfigVersion {
		log.Info("Configuring the AWS Client")
		if err := r.AWSClient.Configure(&config.Spec.AWS); err != nil {
			return ctrl.Result{}, err
		}

		labels := config.GetLabels()
		if len(labels) == 0 {
			labels = map[string]string{}
		}
		labels[loadedAWSConfig] = awsConfigVersion
		config.SetLabels(labels)

		if err := r.Update(ctx, &config); err != nil {
			return ctrl.Result{}, err
		}

		return ctrl.Result{Requeue: true}, r.updateConfigConditions(ctx, nsn, selfv1alpha1.ConfigConditionReady, "", "")
	}

	return ctrl.Result{}, nil
}

// SetupWithManager will setup the controller
func (r *ConfigReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&selfv1alpha1.Config{}).
		Complete(r)
}

func newConfigCondition(condtype selfv1alpha1.ConfigConditionType, message, reason string) *selfv1alpha1.ConfigStatusCondition {
	return &selfv1alpha1.ConfigStatusCondition{
		Type:               condtype,
		Status:             corev1.ConditionTrue,
		LastTransitionTime: metav1.NewTime(time.Now()),
		Message:            message,
		Reason:             reason,
	}
}

// updateConfigConditions updates the status conditions
func (r *ConfigReconciler) updateConfigConditions(
	ctx context.Context,
	nsn types.NamespacedName,
	condtype selfv1alpha1.ConfigConditionType,
	message,
	reason string) error {

	err := retry.RetryOnConflict(retry.DefaultBackoff, func() error {
		var config selfv1alpha1.Config
		if err := r.Get(ctx, nsn, &config); err != nil {
			return err
		}
		configCopy := config.DeepCopy()

		condition := newConfigCondition(condtype, message, reason)

		conditions := []selfv1alpha1.ConfigStatusCondition{}
		if len(configCopy.Status.Conditions) != 0 {
			conditions = configCopy.Status.Conditions
		}

		if i, exists := existsInConditions(conditions, condition.Type); exists {
			conditions[i] = *condition
		} else {
			conditions = append(conditions, *condition)
		}

		for i, cond := range conditions {
			if cond.Type != condition.Type {
				conditions[i].Status = corev1.ConditionFalse
			}
		}

		configCopy.Status.Conditions = conditions

		return r.Status().Update(ctx, configCopy)
	})
	if err != nil {
		r.Log.Error(err, "error updating config conditions")
	}

	return nil
}

func existsInConditions(conditions []selfv1alpha1.ConfigStatusCondition, condtype selfv1alpha1.ConfigConditionType) (int, bool) {
	for i, cond := range conditions {
		if cond.Type == condtype {
			return i, true
		}
	}
	return 0, false
}
