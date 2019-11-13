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

// Package Controller will generate the controllers/<group>/<resource>_controller.go
package controller

import (
	"fmt"
	"path/filepath"
	"strings"

	"awsctrl.io/generator/pkg/input"
	"awsctrl.io/generator/pkg/resource"
)

var _ input.File = &Controller{}

// Controller scaffolds the controllers/<group>/<resource>_controller.go
type Controller struct {
	input.Input

	// Resource is a resource in the API group
	Resource *resource.Resource
}

// GetInput implements input.File
func (in *Controller) GetInput() input.Input {
	if in.Path == "" {
		in.Path = strings.ToLower(filepath.Join("controllers", in.Resource.Group, fmt.Sprintf("%s_controller.go", in.Resource.Kind)))
	}
	in.TemplateBody = controllerTemplate
	return in.Input
}

// Validate validates the values
func (g *Controller) Validate() error {
	return g.Resource.Validate()
}

const controllerTemplate = `{{ .Boilerplate }}

package {{ .Resource.Group | lower }}

import (
	"context"
	"time"

	"github.com/go-logr/logr"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/runtime"

	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"

	v1alpha1 "awsctrl.io/apis/{{ .Resource.Group | lower }}/{{ .Resource.Version }}"
	cloudformationv1alpha1 "awsctrl.io/apis/cloudformation/v1alpha1"
	"awsctrl.io/controllers/generic"
)

var (
	// APIGVStr returns the group version for the resource
	APIGVStr = v1alpha1.GroupVersion.String()
)

// {{ .Resource.Kind }}Reconciler reconciles a {{ .Resource.Kind }} object
type {{ .Resource.Kind }}Reconciler struct {
	client.Client
	Log    logr.Logger
	Scheme *runtime.Scheme
}

// Load the Cloudformation Stack resource
// +kubebuilder:rbac:groups=cloudformation.awsctrl.io,resources=stacks,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=cloudformation.awsctrl.io,resources=stacks/status,verbs=get;update;patch

// Load the {{ .Resource.Group }} {{ .Resource.Kind }} resource
// +kubebuilder:rbac:groups={{ .Resource.Group | lower }}.awsctrl.io,resources={{ .Resource.Kind | lower | pluralize }},verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups={{ .Resource.Group | lower }}.awsctrl.io,resources={{ .Resource.Kind | lower | pluralize }}/status,verbs=get;update;patch

// Reconcile will make the desired state a reality
func (r *{{ .Resource.Kind }}Reconciler) Reconcile(req ctrl.Request) (ctrl.Result, error) {
	ctx := context.Background()
	log := r.Log.WithValues("{{ .Resource.Kind }}", req.NamespacedName)

	var err error
	var instance v1alpha1.{{ .Resource.Kind }}
	if err = r.Get(ctx, req.NamespacedName, &instance); err != nil {
		if errors.IsNotFound(err) {
			return ctrl.Result{}, nil
		}
		return ctrl.Result{}, err
	}

	var cfncontroller generic.Generic
	if cfncontroller, err = generic.New(r.Client, r.Scheme); err != nil {
		return ctrl.Result{}, err
	}

	var requeue time.Duration
	if requeue, err = cfncontroller.Reconcile(ctx, log, &instance); err != nil {
		return ctrl.Result{RequeueAfter: requeue}, err
	}

	return ctrl.Result{}, nil
}

// SetupWithManager will setup the controller
func (r *{{ .Resource.Kind }}Reconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&v1alpha1.{{ .Resource.Kind }}{}).
		Owns(&cloudformationv1alpha1.Stack{}).
		Complete(r)
}
`
