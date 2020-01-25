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

// UsagePlanKeySpec defines the desired state of UsagePlanKey
type UsagePlanKeySpec struct {
	metav1alpha1.CloudFormationMeta `json:",inline"`

	// KeyType http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigateway-usageplankey.html#cfn-apigateway-usageplankey-keytype
	KeyType string `json:"keyType,omitempty" cloudformation:"KeyType,Parameter"`

	// UsagePlanRef http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigateway-usageplankey.html#cfn-apigateway-usageplankey-usageplanid
	UsagePlanRef metav1alpha1.ObjectReference `json:"usagePlanRef,omitempty" cloudformation:"UsagePlanId,Parameter"`

	// KeyRef http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigateway-usageplankey.html#cfn-apigateway-usageplankey-keyid
	KeyRef metav1alpha1.ObjectReference `json:"keyRef,omitempty" cloudformation:"KeyId,Parameter"`
}

// UsagePlanKeyStatus defines the observed state of UsagePlanKey
type UsagePlanKeyStatus struct {
	metav1alpha1.StatusMeta `json:",inline"`
}

// UsagePlanKeyOutput defines the stack outputs
type UsagePlanKeyOutput struct {
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigateway-usageplankey.html
	Ref string `json:"ref,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:resource:categories=aws;apigateway
// +kubebuilder:printcolumn:JSONPath=.status.status,description="status of the stack",name=Status,priority=0,type=string
// +kubebuilder:printcolumn:JSONPath=.status.message,description="reason for the stack status",name=Message,priority=1,type=string
// +kubebuilder:printcolumn:JSONPath=.status.stackID,description="CloudFormation Stack ID",name=StackID,priority=2,type=string

// UsagePlanKey is the Schema for the apigateway UsagePlanKey API
type UsagePlanKey struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   UsagePlanKeySpec   `json:"spec,omitempty"`
	Status UsagePlanKeyStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// UsagePlanKeyList contains a list of Account
type UsagePlanKeyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`

	Items []UsagePlanKey `json:"items"`
}

func init() {
	SchemeBuilder.Register(&UsagePlanKey{}, &UsagePlanKeyList{})
}
