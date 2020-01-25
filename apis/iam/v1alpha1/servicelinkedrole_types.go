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

// ServiceLinkedRoleSpec defines the desired state of ServiceLinkedRole
type ServiceLinkedRoleSpec struct {
	metav1alpha1.CloudFormationMeta `json:",inline"`

	// AWSServiceName http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iam-servicelinkedrole.html#cfn-iam-servicelinkedrole-awsservicename
	AWSServiceName string `json:"aWSServiceName,omitempty" cloudformation:"AWSServiceName,Parameter"`

	// CustomSuffix http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iam-servicelinkedrole.html#cfn-iam-servicelinkedrole-customsuffix
	CustomSuffix string `json:"customSuffix,omitempty" cloudformation:"CustomSuffix,Parameter"`

	// Description http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iam-servicelinkedrole.html#cfn-iam-servicelinkedrole-description
	Description string `json:"description,omitempty" cloudformation:"Description,Parameter"`
}

// ServiceLinkedRoleStatus defines the observed state of ServiceLinkedRole
type ServiceLinkedRoleStatus struct {
	metav1alpha1.StatusMeta `json:",inline"`
}

// ServiceLinkedRoleOutput defines the stack outputs
type ServiceLinkedRoleOutput struct {
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iam-servicelinkedrole.html
	Ref string `json:"ref,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:resource:categories=aws;iam
// +kubebuilder:printcolumn:JSONPath=.status.status,description="status of the stack",name=Status,priority=0,type=string
// +kubebuilder:printcolumn:JSONPath=.status.message,description="reason for the stack status",name=Message,priority=1,type=string
// +kubebuilder:printcolumn:JSONPath=.status.stackID,description="CloudFormation Stack ID",name=StackID,priority=2,type=string

// ServiceLinkedRole is the Schema for the iam ServiceLinkedRole API
type ServiceLinkedRole struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ServiceLinkedRoleSpec   `json:"spec,omitempty"`
	Status ServiceLinkedRoleStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ServiceLinkedRoleList contains a list of Account
type ServiceLinkedRoleList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`

	Items []ServiceLinkedRole `json:"items"`
}

func init() {
	SchemeBuilder.Register(&ServiceLinkedRole{}, &ServiceLinkedRoleList{})
}
