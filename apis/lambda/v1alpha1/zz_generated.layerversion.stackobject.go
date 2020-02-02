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
func (in *LayerVersion) GetNotificationARNs() []string {
	notifcations := []string{}
	for _, notifarn := range in.Spec.NotificationARNs {
		notifcations = append(notifcations, *notifarn)
	}
	return notifcations
}

// GetTemplate will return the JSON version of the CFN to use.
func (in *LayerVersion) GetTemplate(client dynamic.Interface) (string, error) {
	if client == nil {
		return "", fmt.Errorf("k8s client not loaded for template")
	}
	template := cloudformation.NewTemplate()

	template.Description = "AWS Controller - lambda.LayerVersion (ac-{TODO})"

	template.Outputs = map[string]interface{}{
		"ResourceRef": map[string]interface{}{
			"Value": cloudformation.Ref("LayerVersion"),
			"Export": map[string]interface{}{
				"Name": in.Name + "Ref",
			},
		},
	}

	lambdaLayerVersion := &lambda.LayerVersion{}

	if in.Spec.LayerName != "" {
		lambdaLayerVersion.LayerName = in.Spec.LayerName
	}

	if in.Spec.LicenseInfo != "" {
		lambdaLayerVersion.LicenseInfo = in.Spec.LicenseInfo
	}

	if len(in.Spec.CompatibleRuntimes) > 0 {
		lambdaLayerVersion.CompatibleRuntimes = in.Spec.CompatibleRuntimes
	}

	if !reflect.DeepEqual(in.Spec.Content, LayerVersion_Content{}) {
		lambdaLayerVersionContent := lambda.LayerVersion_Content{}

		if in.Spec.Content.S3Bucket != "" {
			lambdaLayerVersionContent.S3Bucket = in.Spec.Content.S3Bucket
		}

		if in.Spec.Content.S3Key != "" {
			lambdaLayerVersionContent.S3Key = in.Spec.Content.S3Key
		}

		if in.Spec.Content.S3ObjectVersion != "" {
			lambdaLayerVersionContent.S3ObjectVersion = in.Spec.Content.S3ObjectVersion
		}

		lambdaLayerVersion.Content = &lambdaLayerVersionContent
	}

	if in.Spec.Description != "" {
		lambdaLayerVersion.Description = in.Spec.Description
	}

	template.Resources = map[string]cloudformation.Resource{
		"LayerVersion": lambdaLayerVersion,
	}

	// json, err := template.JSONWithOptions(&intrinsics.ProcessorOptions{NoEvaluateConditions: true})
	json, err := template.JSON()
	if err != nil {
		return "", err
	}

	return string(json), nil
}

// GetStackID will return stackID
func (in *LayerVersion) GetStackID() string {
	return in.Status.StackID
}

// GenerateStackName will generate a StackName
func (in *LayerVersion) GenerateStackName() string {
	return strings.Join([]string{"lambda", "layerversion", in.GetName(), in.GetNamespace()}, "-")
}

// GetStackName will return stackName
func (in *LayerVersion) GetStackName() string {
	return in.Spec.StackName
}

// GetTemplateVersionLabel will return the stack template version
func (in *LayerVersion) GetTemplateVersionLabel() (value string, ok bool) {
	value, ok = in.Labels[controllerutils.StackTemplateVersionLabel]
	return
}

// GetParameters will return CFN Parameters
func (in *LayerVersion) GetParameters() map[string]string {
	params := map[string]string{}
	cfnencoder.MarshalTypes(params, in.Spec, "Parameter")
	return params
}

// GetCloudFormationMeta will return CFN meta object
func (in *LayerVersion) GetCloudFormationMeta() metav1alpha1.CloudFormationMeta {
	return in.Spec.CloudFormationMeta
}

// GetStatus will return the CFN Status
func (in *LayerVersion) GetStatus() metav1alpha1.ConditionStatus {
	return in.Status.Status
}

// SetStackID will put a stackID
func (in *LayerVersion) SetStackID(input string) {
	in.Status.StackID = input
	return
}

// SetStackName will return stackName
func (in *LayerVersion) SetStackName(input string) {
	in.Spec.StackName = input
	return
}

// SetTemplateVersionLabel will set the template version label
func (in *LayerVersion) SetTemplateVersionLabel() {
	if len(in.Labels) == 0 {
		in.Labels = map[string]string{}
	}

	in.Labels[controllerutils.StackTemplateVersionLabel] = controllerutils.ComputeHash(in.Spec)
}

// TemplateVersionChanged will return bool if template has changed
func (in *LayerVersion) TemplateVersionChanged() bool {
	// Ignore bool since it will still record changed
	label, _ := in.GetTemplateVersionLabel()
	return label != controllerutils.ComputeHash(in.Spec)
}

// SetStatus will set status for object
func (in *LayerVersion) SetStatus(status *metav1alpha1.StatusMeta) {
	in.Status.StatusMeta = *status
}
