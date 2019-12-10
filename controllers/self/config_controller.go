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

	"github.com/go-logr/logr"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/types"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"

	selfv1alpha1 "go.awsctrl.io/manager/apis/self/v1alpha1"
	"go.awsctrl.io/manager/aws"
)

var (
	configDeletionFinalizerName = "config.self.awsctrl.io/deletion"
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

	var config selfv1alpha1.Config
	if err := r.Get(ctx, types.NamespacedName{Namespace: r.PodNamespace, Name: r.ConfigName}, &config); err != nil {
		return ctrl.Result{}, err
	}

	log.Info("Configuring the AWS Client")
	if err := r.AWSClient.Configure(&config.Spec.AWS); err != nil {
		return ctrl.Result{}, err
	}

	return ctrl.Result{}, nil
}

// SetupWithManager will setup the controller
func (r *ConfigReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&selfv1alpha1.Config{}).
		Complete(r)
}
