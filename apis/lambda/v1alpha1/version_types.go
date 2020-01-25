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

// VersionSpec defines the desired state of Version
type VersionSpec struct {
	metav1alpha1.CloudFormationMeta `json:",inline"`

	// CodeSha256 http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lambda-version.html#cfn-lambda-version-codesha256
	CodeSha256 string `json:"codeSha256,omitempty" cloudformation:"CodeSha256,Parameter"`

	// Description http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lambda-version.html#cfn-lambda-version-description
	Description string `json:"description,omitempty" cloudformation:"Description,Parameter"`

	// FunctionName http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lambda-version.html#cfn-lambda-version-functionname
	FunctionName string `json:"functionName,omitempty" cloudformation:"FunctionName,Parameter"`

	// ProvisionedConcurrencyConfig http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lambda-version.html#cfn-lambda-version-provisionedconcurrencyconfig
	ProvisionedConcurrencyConfig Version_ProvisionedConcurrencyConfiguration `json:"provisionedConcurrencyConfig,omitempty" cloudformation:"ProvisionedConcurrencyConfig"`
}

// Version_ProvisionedConcurrencyConfiguration defines the desired state of VersionProvisionedConcurrencyConfiguration
type Version_ProvisionedConcurrencyConfiguration struct {
	// ProvisionedConcurrentExecutions http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lambda-version-provisionedconcurrencyconfiguration.html#cfn-lambda-version-provisionedconcurrencyconfiguration-provisionedconcurrentexecutions
	ProvisionedConcurrentExecutions int `json:"provisionedConcurrentExecutions,omitempty" cloudformation:"ProvisionedConcurrentExecutions,Parameter"`
}

// VersionStatus defines the observed state of Version
type VersionStatus struct {
	metav1alpha1.StatusMeta `json:",inline"`
}

// VersionOutput defines the stack outputs
type VersionOutput struct {
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lambda-version.html
	Ref string `json:"ref,omitempty"`

	// Version defines the Version
	Version string `json:"version,omitempty" cloudformation:"Version,Output"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:resource:categories=aws;lambda
// +kubebuilder:printcolumn:JSONPath=.status.status,description="status of the stack",name=Status,priority=0,type=string
// +kubebuilder:printcolumn:JSONPath=.status.message,description="reason for the stack status",name=Message,priority=1,type=string
// +kubebuilder:printcolumn:JSONPath=.status.stackID,description="CloudFormation Stack ID",name=StackID,priority=2,type=string

// Version is the Schema for the lambda Version API
type Version struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   VersionSpec   `json:"spec,omitempty"`
	Status VersionStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// VersionList contains a list of Account
type VersionList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`

	Items []Version `json:"items"`
}

func init() {
	SchemeBuilder.Register(&Version{}, &VersionList{})
}
