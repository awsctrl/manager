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
	metav1alpha1 "go.awsctrl.io/manager/apis/meta/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// LayerVersionSpec defines the desired state of LayerVersion
type LayerVersionSpec struct {
	metav1alpha1.CloudFormationMeta `json:",inline"`

	// CompatibleRuntimes http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lambda-layerversion.html#cfn-lambda-layerversion-compatibleruntimes
	CompatibleRuntimes []string `json:"compatibleRuntimes,omitempty" cloudformation:"CompatibleRuntimes"`

	// Content http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lambda-layerversion.html#cfn-lambda-layerversion-content
	Content LayerVersion_Content `json:"content,omitempty" cloudformation:"Content"`

	// Description http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lambda-layerversion.html#cfn-lambda-layerversion-description
	Description string `json:"description,omitempty" cloudformation:"Description,Parameter"`

	// LayerName http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lambda-layerversion.html#cfn-lambda-layerversion-layername
	LayerName string `json:"layerName,omitempty" cloudformation:"LayerName,Parameter"`

	// LicenseInfo http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lambda-layerversion.html#cfn-lambda-layerversion-licenseinfo
	LicenseInfo string `json:"licenseInfo,omitempty" cloudformation:"LicenseInfo,Parameter"`
}

// LayerVersion_Content defines the desired state of LayerVersionContent
type LayerVersion_Content struct {
	// S3Bucket http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lambda-layerversion-content.html#cfn-lambda-layerversion-content-s3bucket
	S3Bucket string `json:"s3Bucket,omitempty" cloudformation:"S3Bucket,Parameter"`

	// S3Key http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lambda-layerversion-content.html#cfn-lambda-layerversion-content-s3key
	S3Key string `json:"s3Key,omitempty" cloudformation:"S3Key,Parameter"`

	// S3ObjectVersion http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lambda-layerversion-content.html#cfn-lambda-layerversion-content-s3objectversion
	S3ObjectVersion string `json:"s3ObjectVersion,omitempty" cloudformation:"S3ObjectVersion,Parameter"`
}

// LayerVersionStatus defines the observed state of LayerVersion
type LayerVersionStatus struct {
	metav1alpha1.StatusMeta `json:",inline"`
}

// LayerVersionOutput defines the stack outputs
type LayerVersionOutput struct {
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lambda-layerversion.html
	Ref string `json:"ref,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:resource:categories=aws;lambda
// +kubebuilder:printcolumn:JSONPath=.status.status,description="status of the stack",name=Status,priority=0,type=string
// +kubebuilder:printcolumn:JSONPath=.status.message,description="reason for the stack status",name=Message,priority=1,type=string
// +kubebuilder:printcolumn:JSONPath=.status.stackID,description="CloudFormation Stack ID",name=StackID,priority=2,type=string

// LayerVersion is the Schema for the lambda LayerVersion API
type LayerVersion struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   LayerVersionSpec   `json:"spec,omitempty"`
	Status LayerVersionStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// LayerVersionList contains a list of Account
type LayerVersionList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`

	Items []LayerVersion `json:"items"`
}

func init() {
	SchemeBuilder.Register(&LayerVersion{}, &LayerVersionList{})
}
