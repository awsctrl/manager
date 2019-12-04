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

// Package controllermanager sets up the controller manager
package controllermanager

import (
	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"

	apigatewayv1alpha1 "go.awsctrl.io/manager/apis/apigateway/v1alpha1"
	"go.awsctrl.io/manager/controllers/apigateway"

	cloud9v1alpha1 "go.awsctrl.io/manager/apis/cloud9/v1alpha1"
	"go.awsctrl.io/manager/controllers/cloud9"

	"sigs.k8s.io/controller-runtime/pkg/manager"
)

// AddAllSchemes will configure all the schemes
func AddAllSchemes(scheme *runtime.Scheme) error {

	_ = apigatewayv1alpha1.AddToScheme(scheme)

	_ = cloud9v1alpha1.AddToScheme(scheme)

	return nil
}

// SetupControllers will configure your manager with all controllers
func SetupControllers(mgr manager.Manager) (reconciler string, err error) {

	if err = (&cloud9.EnvironmentEC2Reconciler{
		Client: mgr.GetClient(),
		Log:    ctrl.Log.WithName("controllers").WithName("cloud9").WithName("environmentec2"),
		Scheme: mgr.GetScheme(),
	}).SetupWithManager(mgr); err != nil {
		return "cloud9:environmentec2", err
	}

	if err = (&apigateway.AccountReconciler{
		Client: mgr.GetClient(),
		Log:    ctrl.Log.WithName("controllers").WithName("apigateway").WithName("account"),
		Scheme: mgr.GetScheme(),
	}).SetupWithManager(mgr); err != nil {
		return "apigateway:account", err
	}

	return reconciler, nil
}
