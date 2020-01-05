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
	"k8s.io/client-go/dynamic"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/manager"

	apigatewayv1alpha1 "go.awsctrl.io/manager/apis/apigateway/v1alpha1"
	"go.awsctrl.io/manager/controllers/apigateway"

	certificatemanagerv1alpha1 "go.awsctrl.io/manager/apis/certificatemanager/v1alpha1"
	"go.awsctrl.io/manager/controllers/certificatemanager"

	cloud9v1alpha1 "go.awsctrl.io/manager/apis/cloud9/v1alpha1"
	"go.awsctrl.io/manager/controllers/cloud9"

	ecrv1alpha1 "go.awsctrl.io/manager/apis/ecr/v1alpha1"
	"go.awsctrl.io/manager/controllers/ecr"

	iamv1alpha1 "go.awsctrl.io/manager/apis/iam/v1alpha1"
	"go.awsctrl.io/manager/controllers/iam"

	route53v1alpha1 "go.awsctrl.io/manager/apis/route53/v1alpha1"
	"go.awsctrl.io/manager/controllers/route53"
)

// AddAllSchemes will configure all the schemes
func AddAllSchemes(scheme *runtime.Scheme) error {

	_ = apigatewayv1alpha1.AddToScheme(scheme)

	_ = certificatemanagerv1alpha1.AddToScheme(scheme)

	_ = cloud9v1alpha1.AddToScheme(scheme)

	_ = ecrv1alpha1.AddToScheme(scheme)

	_ = iamv1alpha1.AddToScheme(scheme)

	_ = route53v1alpha1.AddToScheme(scheme)

	return nil
}

// SetupControllers will configure your manager with all controllers
func SetupControllers(mgr manager.Manager, dynamicClient dynamic.Interface) (reconciler string, err error) {

	if err = (&route53.HostedZoneReconciler{
		Client:    mgr.GetClient(),
		Interface: dynamicClient,
		Log:       ctrl.Log.WithName("controllers").WithName("route53").WithName("hostedzone"),
		Scheme:    mgr.GetScheme(),
	}).SetupWithManager(mgr); err != nil {
		return "route53:hostedzone", err
	}

	if err = (&apigateway.AccountReconciler{
		Client:    mgr.GetClient(),
		Interface: dynamicClient,
		Log:       ctrl.Log.WithName("controllers").WithName("apigateway").WithName("account"),
		Scheme:    mgr.GetScheme(),
	}).SetupWithManager(mgr); err != nil {
		return "apigateway:account", err
	}

	if err = (&apigateway.DomainNameReconciler{
		Client:    mgr.GetClient(),
		Interface: dynamicClient,
		Log:       ctrl.Log.WithName("controllers").WithName("apigateway").WithName("domainname"),
		Scheme:    mgr.GetScheme(),
	}).SetupWithManager(mgr); err != nil {
		return "apigateway:domainname", err
	}

	if err = (&apigateway.ModelReconciler{
		Client:    mgr.GetClient(),
		Interface: dynamicClient,
		Log:       ctrl.Log.WithName("controllers").WithName("apigateway").WithName("model"),
		Scheme:    mgr.GetScheme(),
	}).SetupWithManager(mgr); err != nil {
		return "apigateway:model", err
	}

	if err = (&apigateway.VpcLinkReconciler{
		Client:    mgr.GetClient(),
		Interface: dynamicClient,
		Log:       ctrl.Log.WithName("controllers").WithName("apigateway").WithName("vpclink"),
		Scheme:    mgr.GetScheme(),
	}).SetupWithManager(mgr); err != nil {
		return "apigateway:vpclink", err
	}

	if err = (&route53.RecordSetReconciler{
		Client:    mgr.GetClient(),
		Interface: dynamicClient,
		Log:       ctrl.Log.WithName("controllers").WithName("route53").WithName("recordset"),
		Scheme:    mgr.GetScheme(),
	}).SetupWithManager(mgr); err != nil {
		return "route53:recordset", err
	}

	if err = (&ecr.RepositoryReconciler{
		Client:    mgr.GetClient(),
		Interface: dynamicClient,
		Log:       ctrl.Log.WithName("controllers").WithName("ecr").WithName("repository"),
		Scheme:    mgr.GetScheme(),
	}).SetupWithManager(mgr); err != nil {
		return "ecr:repository", err
	}

	if err = (&iam.RoleReconciler{
		Client:    mgr.GetClient(),
		Interface: dynamicClient,
		Log:       ctrl.Log.WithName("controllers").WithName("iam").WithName("role"),
		Scheme:    mgr.GetScheme(),
	}).SetupWithManager(mgr); err != nil {
		return "iam:role", err
	}

	if err = (&apigateway.DocumentationPartReconciler{
		Client:    mgr.GetClient(),
		Interface: dynamicClient,
		Log:       ctrl.Log.WithName("controllers").WithName("apigateway").WithName("documentationpart"),
		Scheme:    mgr.GetScheme(),
	}).SetupWithManager(mgr); err != nil {
		return "apigateway:documentationpart", err
	}

	if err = (&iam.AccessKeyReconciler{
		Client:    mgr.GetClient(),
		Interface: dynamicClient,
		Log:       ctrl.Log.WithName("controllers").WithName("iam").WithName("accesskey"),
		Scheme:    mgr.GetScheme(),
	}).SetupWithManager(mgr); err != nil {
		return "iam:accesskey", err
	}

	if err = (&iam.ServiceLinkedRoleReconciler{
		Client:    mgr.GetClient(),
		Interface: dynamicClient,
		Log:       ctrl.Log.WithName("controllers").WithName("iam").WithName("servicelinkedrole"),
		Scheme:    mgr.GetScheme(),
	}).SetupWithManager(mgr); err != nil {
		return "iam:servicelinkedrole", err
	}

	if err = (&apigateway.AuthorizerReconciler{
		Client:    mgr.GetClient(),
		Interface: dynamicClient,
		Log:       ctrl.Log.WithName("controllers").WithName("apigateway").WithName("authorizer"),
		Scheme:    mgr.GetScheme(),
	}).SetupWithManager(mgr); err != nil {
		return "apigateway:authorizer", err
	}

	if err = (&iam.InstanceProfileReconciler{
		Client:    mgr.GetClient(),
		Interface: dynamicClient,
		Log:       ctrl.Log.WithName("controllers").WithName("iam").WithName("instanceprofile"),
		Scheme:    mgr.GetScheme(),
	}).SetupWithManager(mgr); err != nil {
		return "iam:instanceprofile", err
	}

	if err = (&certificatemanager.CertificateReconciler{
		Client:    mgr.GetClient(),
		Interface: dynamicClient,
		Log:       ctrl.Log.WithName("controllers").WithName("certificatemanager").WithName("certificate"),
		Scheme:    mgr.GetScheme(),
	}).SetupWithManager(mgr); err != nil {
		return "certificatemanager:certificate", err
	}

	if err = (&apigateway.UsagePlanReconciler{
		Client:    mgr.GetClient(),
		Interface: dynamicClient,
		Log:       ctrl.Log.WithName("controllers").WithName("apigateway").WithName("usageplan"),
		Scheme:    mgr.GetScheme(),
	}).SetupWithManager(mgr); err != nil {
		return "apigateway:usageplan", err
	}

	if err = (&apigateway.StageReconciler{
		Client:    mgr.GetClient(),
		Interface: dynamicClient,
		Log:       ctrl.Log.WithName("controllers").WithName("apigateway").WithName("stage"),
		Scheme:    mgr.GetScheme(),
	}).SetupWithManager(mgr); err != nil {
		return "apigateway:stage", err
	}

	if err = (&apigateway.MethodReconciler{
		Client:    mgr.GetClient(),
		Interface: dynamicClient,
		Log:       ctrl.Log.WithName("controllers").WithName("apigateway").WithName("method"),
		Scheme:    mgr.GetScheme(),
	}).SetupWithManager(mgr); err != nil {
		return "apigateway:method", err
	}

	if err = (&iam.UserToGroupAdditionReconciler{
		Client:    mgr.GetClient(),
		Interface: dynamicClient,
		Log:       ctrl.Log.WithName("controllers").WithName("iam").WithName("usertogroupaddition"),
		Scheme:    mgr.GetScheme(),
	}).SetupWithManager(mgr); err != nil {
		return "iam:usertogroupaddition", err
	}

	if err = (&apigateway.RestApiReconciler{
		Client:    mgr.GetClient(),
		Interface: dynamicClient,
		Log:       ctrl.Log.WithName("controllers").WithName("apigateway").WithName("restapi"),
		Scheme:    mgr.GetScheme(),
	}).SetupWithManager(mgr); err != nil {
		return "apigateway:restapi", err
	}

	if err = (&route53.RecordSetGroupReconciler{
		Client:    mgr.GetClient(),
		Interface: dynamicClient,
		Log:       ctrl.Log.WithName("controllers").WithName("route53").WithName("recordsetgroup"),
		Scheme:    mgr.GetScheme(),
	}).SetupWithManager(mgr); err != nil {
		return "route53:recordsetgroup", err
	}

	if err = (&iam.ManagedPolicyReconciler{
		Client:    mgr.GetClient(),
		Interface: dynamicClient,
		Log:       ctrl.Log.WithName("controllers").WithName("iam").WithName("managedpolicy"),
		Scheme:    mgr.GetScheme(),
	}).SetupWithManager(mgr); err != nil {
		return "iam:managedpolicy", err
	}

	if err = (&apigateway.DeploymentReconciler{
		Client:    mgr.GetClient(),
		Interface: dynamicClient,
		Log:       ctrl.Log.WithName("controllers").WithName("apigateway").WithName("deployment"),
		Scheme:    mgr.GetScheme(),
	}).SetupWithManager(mgr); err != nil {
		return "apigateway:deployment", err
	}

	if err = (&apigateway.ApiKeyReconciler{
		Client:    mgr.GetClient(),
		Interface: dynamicClient,
		Log:       ctrl.Log.WithName("controllers").WithName("apigateway").WithName("apikey"),
		Scheme:    mgr.GetScheme(),
	}).SetupWithManager(mgr); err != nil {
		return "apigateway:apikey", err
	}

	if err = (&apigateway.ClientCertificateReconciler{
		Client:    mgr.GetClient(),
		Interface: dynamicClient,
		Log:       ctrl.Log.WithName("controllers").WithName("apigateway").WithName("clientcertificate"),
		Scheme:    mgr.GetScheme(),
	}).SetupWithManager(mgr); err != nil {
		return "apigateway:clientcertificate", err
	}

	if err = (&iam.PolicyReconciler{
		Client:    mgr.GetClient(),
		Interface: dynamicClient,
		Log:       ctrl.Log.WithName("controllers").WithName("iam").WithName("policy"),
		Scheme:    mgr.GetScheme(),
	}).SetupWithManager(mgr); err != nil {
		return "iam:policy", err
	}

	if err = (&iam.GroupReconciler{
		Client:    mgr.GetClient(),
		Interface: dynamicClient,
		Log:       ctrl.Log.WithName("controllers").WithName("iam").WithName("group"),
		Scheme:    mgr.GetScheme(),
	}).SetupWithManager(mgr); err != nil {
		return "iam:group", err
	}

	if err = (&cloud9.EnvironmentEC2Reconciler{
		Client:    mgr.GetClient(),
		Interface: dynamicClient,
		Log:       ctrl.Log.WithName("controllers").WithName("cloud9").WithName("environmentec2"),
		Scheme:    mgr.GetScheme(),
	}).SetupWithManager(mgr); err != nil {
		return "cloud9:environmentec2", err
	}

	if err = (&apigateway.UsagePlanKeyReconciler{
		Client:    mgr.GetClient(),
		Interface: dynamicClient,
		Log:       ctrl.Log.WithName("controllers").WithName("apigateway").WithName("usageplankey"),
		Scheme:    mgr.GetScheme(),
	}).SetupWithManager(mgr); err != nil {
		return "apigateway:usageplankey", err
	}

	if err = (&apigateway.ResourceReconciler{
		Client:    mgr.GetClient(),
		Interface: dynamicClient,
		Log:       ctrl.Log.WithName("controllers").WithName("apigateway").WithName("resource"),
		Scheme:    mgr.GetScheme(),
	}).SetupWithManager(mgr); err != nil {
		return "apigateway:resource", err
	}

	if err = (&apigateway.GatewayResponseReconciler{
		Client:    mgr.GetClient(),
		Interface: dynamicClient,
		Log:       ctrl.Log.WithName("controllers").WithName("apigateway").WithName("gatewayresponse"),
		Scheme:    mgr.GetScheme(),
	}).SetupWithManager(mgr); err != nil {
		return "apigateway:gatewayresponse", err
	}

	if err = (&apigateway.DocumentationVersionReconciler{
		Client:    mgr.GetClient(),
		Interface: dynamicClient,
		Log:       ctrl.Log.WithName("controllers").WithName("apigateway").WithName("documentationversion"),
		Scheme:    mgr.GetScheme(),
	}).SetupWithManager(mgr); err != nil {
		return "apigateway:documentationversion", err
	}

	if err = (&iam.UserReconciler{
		Client:    mgr.GetClient(),
		Interface: dynamicClient,
		Log:       ctrl.Log.WithName("controllers").WithName("iam").WithName("user"),
		Scheme:    mgr.GetScheme(),
	}).SetupWithManager(mgr); err != nil {
		return "iam:user", err
	}

	if err = (&route53.HealthCheckReconciler{
		Client:    mgr.GetClient(),
		Interface: dynamicClient,
		Log:       ctrl.Log.WithName("controllers").WithName("route53").WithName("healthcheck"),
		Scheme:    mgr.GetScheme(),
	}).SetupWithManager(mgr); err != nil {
		return "route53:healthcheck", err
	}

	if err = (&apigateway.RequestValidatorReconciler{
		Client:    mgr.GetClient(),
		Interface: dynamicClient,
		Log:       ctrl.Log.WithName("controllers").WithName("apigateway").WithName("requestvalidator"),
		Scheme:    mgr.GetScheme(),
	}).SetupWithManager(mgr); err != nil {
		return "apigateway:requestvalidator", err
	}

	if err = (&apigateway.BasePathMappingReconciler{
		Client:    mgr.GetClient(),
		Interface: dynamicClient,
		Log:       ctrl.Log.WithName("controllers").WithName("apigateway").WithName("basepathmapping"),
		Scheme:    mgr.GetScheme(),
	}).SetupWithManager(mgr); err != nil {
		return "apigateway:basepathmapping", err
	}

	return reconciler, nil
}
