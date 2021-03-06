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

package v1alpha1

import (
	"fmt"
	"reflect"
	"strings"

	metav1alpha1 "go.awsctrl.io/manager/apis/meta/v1alpha1"
	controllerutils "go.awsctrl.io/manager/controllers/utils"
	cfnencoder "go.awsctrl.io/manager/encoding/cloudformation"

	"github.com/awslabs/goformation/v4/cloudformation"
	"github.com/awslabs/goformation/v4/cloudformation/apigateway"
	"k8s.io/client-go/dynamic"
)

// GetNotificationARNs is an autogenerated deepcopy function, will return notifications for stack
func (in *Stage) GetNotificationARNs() []string {
	notifcations := []string{}
	for _, notifarn := range in.Spec.NotificationARNs {
		notifcations = append(notifcations, *notifarn)
	}
	return notifcations
}

// GetTemplate will return the JSON version of the CFN to use.
func (in *Stage) GetTemplate(client dynamic.Interface) (string, error) {
	if client == nil {
		return "", fmt.Errorf("k8s client not loaded for template")
	}
	template := cloudformation.NewTemplate()

	template.Description = "AWS Controller - apigateway.Stage (ac-{TODO})"

	template.Outputs = map[string]interface{}{
		"ResourceRef": map[string]interface{}{
			"Value": cloudformation.Ref("Stage"),
			"Export": map[string]interface{}{
				"Name": in.Name + "Ref",
			},
		},
	}

	apigatewayStage := &apigateway.Stage{}

	if !reflect.DeepEqual(in.Spec.AccessLogSetting, Stage_AccessLogSetting{}) {
		apigatewayStageAccessLogSetting := apigateway.Stage_AccessLogSetting{}

		// TODO(christopherhein) move these to a defaulter
		apigatewayStageAccessLogSettingDestinationRefItem := in.Spec.AccessLogSetting.DestinationRef.DeepCopy()

		if apigatewayStageAccessLogSettingDestinationRefItem.ObjectRef.Namespace == "" {
			apigatewayStageAccessLogSettingDestinationRefItem.ObjectRef.Namespace = in.Namespace
		}

		in.Spec.AccessLogSetting.DestinationRef = *apigatewayStageAccessLogSettingDestinationRefItem
		destinationArn, err := in.Spec.AccessLogSetting.DestinationRef.String(client)
		if err != nil {
			return "", err
		}

		if destinationArn != "" {
			apigatewayStageAccessLogSetting.DestinationArn = destinationArn
		}

		if in.Spec.AccessLogSetting.Format != "" {
			apigatewayStageAccessLogSetting.Format = in.Spec.AccessLogSetting.Format
		}

		apigatewayStage.AccessLogSetting = &apigatewayStageAccessLogSetting
	}

	if in.Spec.CacheClusterEnabled || !in.Spec.CacheClusterEnabled {
		apigatewayStage.CacheClusterEnabled = in.Spec.CacheClusterEnabled
	}

	if in.Spec.CacheClusterSize != "" {
		apigatewayStage.CacheClusterSize = in.Spec.CacheClusterSize
	}

	if !reflect.DeepEqual(in.Spec.CanarySetting, Stage_CanarySetting{}) {
		apigatewayStageCanarySetting := apigateway.Stage_CanarySetting{}

		// TODO(christopherhein) move these to a defaulter
		apigatewayStageCanarySettingDeploymentRefItem := in.Spec.CanarySetting.DeploymentRef.DeepCopy()

		if apigatewayStageCanarySettingDeploymentRefItem.ObjectRef.Namespace == "" {
			apigatewayStageCanarySettingDeploymentRefItem.ObjectRef.Namespace = in.Namespace
		}

		in.Spec.CanarySetting.DeploymentRef = *apigatewayStageCanarySettingDeploymentRefItem
		deploymentId, err := in.Spec.CanarySetting.DeploymentRef.String(client)
		if err != nil {
			return "", err
		}

		if deploymentId != "" {
			apigatewayStageCanarySetting.DeploymentId = deploymentId
		}

		if float64(in.Spec.CanarySetting.PercentTraffic) != apigatewayStageCanarySetting.PercentTraffic {
			apigatewayStageCanarySetting.PercentTraffic = float64(in.Spec.CanarySetting.PercentTraffic)
		}

		if !reflect.DeepEqual(in.Spec.CanarySetting.StageVariableOverrides, map[string]string{}) {
			apigatewayStageCanarySetting.StageVariableOverrides = in.Spec.CanarySetting.StageVariableOverrides
		}

		if in.Spec.CanarySetting.UseStageCache || !in.Spec.CanarySetting.UseStageCache {
			apigatewayStageCanarySetting.UseStageCache = in.Spec.CanarySetting.UseStageCache
		}

		apigatewayStage.CanarySetting = &apigatewayStageCanarySetting
	}

	// TODO(christopherhein) move these to a defaulter
	apigatewayStageClientCertificateRefItem := in.Spec.ClientCertificateRef.DeepCopy()

	if apigatewayStageClientCertificateRefItem.ObjectRef.Namespace == "" {
		apigatewayStageClientCertificateRefItem.ObjectRef.Namespace = in.Namespace
	}

	in.Spec.ClientCertificateRef = *apigatewayStageClientCertificateRefItem
	clientCertificateId, err := in.Spec.ClientCertificateRef.String(client)
	if err != nil {
		return "", err
	}

	if clientCertificateId != "" {
		apigatewayStage.ClientCertificateId = clientCertificateId
	}

	// TODO(christopherhein) move these to a defaulter
	apigatewayStageDeploymentRefItem := in.Spec.DeploymentRef.DeepCopy()

	if apigatewayStageDeploymentRefItem.ObjectRef.Namespace == "" {
		apigatewayStageDeploymentRefItem.ObjectRef.Namespace = in.Namespace
	}

	in.Spec.DeploymentRef = *apigatewayStageDeploymentRefItem
	deploymentId, err := in.Spec.DeploymentRef.String(client)
	if err != nil {
		return "", err
	}

	if deploymentId != "" {
		apigatewayStage.DeploymentId = deploymentId
	}

	if in.Spec.Description != "" {
		apigatewayStage.Description = in.Spec.Description
	}

	if in.Spec.DocumentationVersion != "" {
		apigatewayStage.DocumentationVersion = in.Spec.DocumentationVersion
	}

	apigatewayStageMethodSettings := []apigateway.Stage_MethodSetting{}

	for _, item := range in.Spec.MethodSettings {
		apigatewayStageMethodSetting := apigateway.Stage_MethodSetting{}

		if item.CacheDataEncrypted || !item.CacheDataEncrypted {
			apigatewayStageMethodSetting.CacheDataEncrypted = item.CacheDataEncrypted
		}

		if item.CacheTtlInSeconds != apigatewayStageMethodSetting.CacheTtlInSeconds {
			apigatewayStageMethodSetting.CacheTtlInSeconds = item.CacheTtlInSeconds
		}

		if item.CachingEnabled || !item.CachingEnabled {
			apigatewayStageMethodSetting.CachingEnabled = item.CachingEnabled
		}

		if item.DataTraceEnabled || !item.DataTraceEnabled {
			apigatewayStageMethodSetting.DataTraceEnabled = item.DataTraceEnabled
		}

		if item.HttpMethod != "" {
			apigatewayStageMethodSetting.HttpMethod = item.HttpMethod
		}

		if item.LoggingLevel != "" {
			apigatewayStageMethodSetting.LoggingLevel = item.LoggingLevel
		}

		if item.MetricsEnabled || !item.MetricsEnabled {
			apigatewayStageMethodSetting.MetricsEnabled = item.MetricsEnabled
		}

		if item.ResourcePath != "" {
			apigatewayStageMethodSetting.ResourcePath = item.ResourcePath
		}

		if item.ThrottlingBurstLimit != apigatewayStageMethodSetting.ThrottlingBurstLimit {
			apigatewayStageMethodSetting.ThrottlingBurstLimit = item.ThrottlingBurstLimit
		}

		if float64(item.ThrottlingRateLimit) != apigatewayStageMethodSetting.ThrottlingRateLimit {
			apigatewayStageMethodSetting.ThrottlingRateLimit = float64(item.ThrottlingRateLimit)
		}

	}

	if len(apigatewayStageMethodSettings) > 0 {
		apigatewayStage.MethodSettings = apigatewayStageMethodSettings
	}
	// TODO(christopherhein) move these to a defaulter
	apigatewayStageRestApiRefItem := in.Spec.RestApiRef.DeepCopy()

	if apigatewayStageRestApiRefItem.ObjectRef.Namespace == "" {
		apigatewayStageRestApiRefItem.ObjectRef.Namespace = in.Namespace
	}

	in.Spec.RestApiRef = *apigatewayStageRestApiRefItem
	restApiId, err := in.Spec.RestApiRef.String(client)
	if err != nil {
		return "", err
	}

	if restApiId != "" {
		apigatewayStage.RestApiId = restApiId
	}

	// TODO(christopherhein) move these to a defaulter
	if in.Spec.StageName == "" {
		apigatewayStage.StageName = in.Name
	}

	if in.Spec.StageName != "" {
		apigatewayStage.StageName = in.Spec.StageName
	}

	// TODO(christopherhein): implement tags this could be easy now that I have the mechanims of nested objects
	if in.Spec.TracingEnabled || !in.Spec.TracingEnabled {
		apigatewayStage.TracingEnabled = in.Spec.TracingEnabled
	}

	if !reflect.DeepEqual(in.Spec.Variables, map[string]string{}) {
		apigatewayStage.Variables = in.Spec.Variables
	}

	template.Resources = map[string]cloudformation.Resource{
		"Stage": apigatewayStage,
	}

	// json, err := template.JSONWithOptions(&intrinsics.ProcessorOptions{NoEvaluateConditions: true})
	json, err := template.JSON()
	if err != nil {
		return "", err
	}

	return string(json), nil
}

// GetStackID will return stackID
func (in *Stage) GetStackID() string {
	return in.Status.StackID
}

// GenerateStackName will generate a StackName
func (in *Stage) GenerateStackName() string {
	return strings.Join([]string{"apigateway", "stage", in.GetName(), in.GetNamespace()}, "-")
}

// GetStackName will return stackName
func (in *Stage) GetStackName() string {
	return in.Spec.StackName
}

// GetTemplateVersionLabel will return the stack template version
func (in *Stage) GetTemplateVersionLabel() (value string, ok bool) {
	value, ok = in.Labels[controllerutils.StackTemplateVersionLabel]
	return
}

// GetParameters will return CFN Parameters
func (in *Stage) GetParameters() map[string]string {
	params := map[string]string{}
	cfnencoder.MarshalTypes(params, in.Spec, "Parameter")
	return params
}

// GetCloudFormationMeta will return CFN meta object
func (in *Stage) GetCloudFormationMeta() metav1alpha1.CloudFormationMeta {
	return in.Spec.CloudFormationMeta
}

// GetStatus will return the CFN Status
func (in *Stage) GetStatus() metav1alpha1.ConditionStatus {
	return in.Status.Status
}

// SetStackID will put a stackID
func (in *Stage) SetStackID(input string) {
	in.Status.StackID = input
	return
}

// SetStackName will return stackName
func (in *Stage) SetStackName(input string) {
	in.Spec.StackName = input
	return
}

// SetTemplateVersionLabel will set the template version label
func (in *Stage) SetTemplateVersionLabel() {
	if len(in.Labels) == 0 {
		in.Labels = map[string]string{}
	}

	in.Labels[controllerutils.StackTemplateVersionLabel] = controllerutils.ComputeHash(in.Spec)
}

// TemplateVersionChanged will return bool if template has changed
func (in *Stage) TemplateVersionChanged() bool {
	// Ignore bool since it will still record changed
	label, _ := in.GetTemplateVersionLabel()
	return label != controllerutils.ComputeHash(in.Spec)
}

// SetStatus will set status for object
func (in *Stage) SetStatus(status *metav1alpha1.StatusMeta) {
	in.Status.StatusMeta = *status
}
