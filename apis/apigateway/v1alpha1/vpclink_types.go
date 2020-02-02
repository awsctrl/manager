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

// VpcLinkSpec defines the desired state of VpcLink
type VpcLinkSpec struct {
	metav1alpha1.CloudFormationMeta `json:",inline"`

	// TargetRefs http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigateway-vpclink.html#cfn-apigateway-vpclink-targetarns
	TargetRefs []metav1alpha1.ObjectReference `json:"targetRefs,omitempty" cloudformation:"TargetArns"`

	// Description http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigateway-vpclink.html#cfn-apigateway-vpclink-description
	Description string `json:"description,omitempty" cloudformation:"Description,Parameter"`

	// Name http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigateway-vpclink.html#cfn-apigateway-vpclink-name
	Name string `json:"name,omitempty" cloudformation:"Name,Parameter"`
}

// VpcLinkStatus defines the observed state of VpcLink
type VpcLinkStatus struct {
	metav1alpha1.StatusMeta `json:",inline"`
}

// VpcLinkOutput defines the stack outputs
type VpcLinkOutput struct {
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigateway-vpclink.html
	Ref string `json:"ref,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:resource:categories=aws;apigateway
// +kubebuilder:printcolumn:JSONPath=.status.status,description="status of the stack",name=Status,priority=0,type=string
// +kubebuilder:printcolumn:JSONPath=.status.message,description="reason for the stack status",name=Message,priority=1,type=string
// +kubebuilder:printcolumn:JSONPath=.status.stackID,description="CloudFormation Stack ID",name=StackID,priority=2,type=string

// VpcLink is the Schema for the apigateway VpcLink API
type VpcLink struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   VpcLinkSpec   `json:"spec,omitempty"`
	Status VpcLinkStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// VpcLinkList contains a list of Account
type VpcLinkList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`

	Items []VpcLink `json:"items"`
}

func init() {
	SchemeBuilder.Register(&VpcLink{}, &VpcLinkList{})
}
