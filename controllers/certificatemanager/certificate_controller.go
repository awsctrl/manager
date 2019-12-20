/*
Copyright © 2019 AWS Controller authors

Licensed under the Apache License, Version 2.0 (the &#34;License&#34;);
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an &#34;AS IS&#34; BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package certificatemanager

import (
	"context"
	"time"

	"github.com/go-logr/logr"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/client-go/dynamic"

	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"

	v1alpha1 "go.awsctrl.io/manager/apis/certificatemanager/v1alpha1"
	cloudformationv1alpha1 "go.awsctrl.io/manager/apis/cloudformation/v1alpha1"
	"go.awsctrl.io/manager/controllers/generic"
)

// CertificateReconciler reconciles a Certificate object
type CertificateReconciler struct {
	client.Client
	dynamic.Interface
	Log    logr.Logger
	Scheme *runtime.Scheme
}

// Load the Cloudformation Stack resource
// +kubebuilder:rbac:groups=cloudformation.awsctrl.io,resources=stacks,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=cloudformation.awsctrl.io,resources=stacks/status,verbs=get;update;patch

// Load the certificatemanager Certificate resource
// +kubebuilder:rbac:groups=certificatemanager.awsctrl.io,resources=certificates,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=certificatemanager.awsctrl.io,resources=certificates/status,verbs=get;update;patch

// Reconcile will make the desired state a reality
func (r *CertificateReconciler) Reconcile(req ctrl.Request) (ctrl.Result, error) {
	ctx := context.Background()
	log := r.Log.WithValues("Certificate", req.NamespacedName)

	var err error
	var instance v1alpha1.Certificate
	if err = r.Get(ctx, req.NamespacedName, &instance); err != nil {
		if errors.IsNotFound(err) {
			return ctrl.Result{}, nil
		}
		return ctrl.Result{}, err
	}

	var cfncontroller generic.Generic
	if cfncontroller, err = generic.New(r.Client, r.Interface, r.Scheme); err != nil {
		return ctrl.Result{}, err
	}

	var requeue time.Duration
	if requeue, err = cfncontroller.Reconcile(ctx, log, &instance); err != nil {
		return ctrl.Result{RequeueAfter: requeue}, err
	}

	return ctrl.Result{}, nil
}

// SetupWithManager will setup the controller
func (r *CertificateReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&v1alpha1.Certificate{}).
		Owns(&cloudformationv1alpha1.Stack{}).
		Complete(r)
}
