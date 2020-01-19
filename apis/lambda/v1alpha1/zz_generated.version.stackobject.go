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
	"github.com/awslabs/goformation/v4/cloudformation/lambda"
	"k8s.io/client-go/dynamic"
)

// GetNotificationARNs is an autogenerated deepcopy function, will return notifications for stack
func (in *Version) GetNotificationARNs() []string {
	notifcations := []string{}
	for _, notifarn := range in.Spec.NotificationARNs {
		notifcations = append(notifcations, *notifarn)
	}
	return notifcations
}

// GetTemplate will return the JSON version of the CFN to use.
func (in *Version) GetTemplate(client dynamic.Interface) (string, error) {
	if client == nil {
		return "", fmt.Errorf("k8s client not loaded for template")
	}
	template := cloudformation.NewTemplate()

	template.Description = "AWS Controller - lambda.Version (ac-{TODO})"

	template.Outputs = map[string]interface{}{
		"ResourceRef": map[string]interface{}{
			"Value": cloudformation.Ref("Version"),
		},
		"Version": map[string]interface{}{
			"Value": cloudformation.GetAtt("Version", "Version"),
		},
	}

	lambdaVersion := &lambda.Version{}

	if !reflect.DeepEqual(in.Spec.ProvisionedConcurrencyConfig, Version_ProvisionedConcurrencyConfiguration{}) {
		lambdaVersionProvisionedConcurrencyConfiguration := lambda.Version_ProvisionedConcurrencyConfiguration{}

		if in.Spec.ProvisionedConcurrencyConfig.ProvisionedConcurrentExecutions != lambdaVersionProvisionedConcurrencyConfiguration.ProvisionedConcurrentExecutions {
			lambdaVersionProvisionedConcurrencyConfiguration.ProvisionedConcurrentExecutions = in.Spec.ProvisionedConcurrencyConfig.ProvisionedConcurrentExecutions
		}

		lambdaVersion.ProvisionedConcurrencyConfig = &lambdaVersionProvisionedConcurrencyConfiguration
	}

	if in.Spec.CodeSha256 != "" {
		lambdaVersion.CodeSha256 = in.Spec.CodeSha256
	}

	if in.Spec.Description != "" {
		lambdaVersion.Description = in.Spec.Description
	}

	if in.Spec.FunctionName != "" {
		lambdaVersion.FunctionName = in.Spec.FunctionName
	}

	template.Resources = map[string]cloudformation.Resource{
		"Version": lambdaVersion,
	}

	// json, err := template.JSONWithOptions(&intrinsics.ProcessorOptions{NoEvaluateConditions: true})
	json, err := template.JSON()
	if err != nil {
		return "", err
	}

	return string(json), nil
}

// GetStackID will return stackID
func (in *Version) GetStackID() string {
	return in.Status.StackID
}

// GenerateStackName will generate a StackName
func (in *Version) GenerateStackName() string {
	return strings.Join([]string{"lambda", "version", in.GetName(), in.GetNamespace()}, "-")
}

// GetStackName will return stackName
func (in *Version) GetStackName() string {
	return in.Spec.StackName
}

// GetTemplateVersionLabel will return the stack template version
func (in *Version) GetTemplateVersionLabel() (value string, ok bool) {
	value, ok = in.Labels[controllerutils.StackTemplateVersionLabel]
	return
}

// GetParameters will return CFN Parameters
func (in *Version) GetParameters() map[string]string {
	params := map[string]string{}
	cfnencoder.MarshalTypes(params, in.Spec, "Parameter")
	return params
}

// GetCloudFormationMeta will return CFN meta object
func (in *Version) GetCloudFormationMeta() metav1alpha1.CloudFormationMeta {
	return in.Spec.CloudFormationMeta
}

// GetStatus will return the CFN Status
func (in *Version) GetStatus() metav1alpha1.ConditionStatus {
	return in.Status.Status
}

// SetStackID will put a stackID
func (in *Version) SetStackID(input string) {
	in.Status.StackID = input
	return
}

// SetStackName will return stackName
func (in *Version) SetStackName(input string) {
	in.Spec.StackName = input
	return
}

// SetTemplateVersionLabel will set the template version label
func (in *Version) SetTemplateVersionLabel() {
	if len(in.Labels) == 0 {
		in.Labels = map[string]string{}
	}

	in.Labels[controllerutils.StackTemplateVersionLabel] = controllerutils.ComputeHash(in.Spec)
}

// TemplateVersionChanged will return bool if template has changed
func (in *Version) TemplateVersionChanged() bool {
	// Ignore bool since it will still record changed
	label, _ := in.GetTemplateVersionLabel()
	return label != controllerutils.ComputeHash(in.Spec)
}

// SetStatus will set status for object
func (in *Version) SetStatus(status *metav1alpha1.StatusMeta) {
	in.Status.StatusMeta = *status
}
