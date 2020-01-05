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

// UserToGroupAdditionSpec defines the desired state of UserToGroupAddition
type UserToGroupAdditionSpec struct {
	metav1alpha1.CloudFormationMeta `json:",inline"`

	// GroupName http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iam-addusertogroup.html#cfn-iam-addusertogroup-groupname
	GroupName string `json:"groupName" cloudformation:"GroupName,Parameter"`

	// Users http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iam-addusertogroup.html#cfn-iam-addusertogroup-users
	Users []string `json:"users" cloudformation:"Users"`
}

// UserToGroupAdditionStatus defines the observed state of UserToGroupAddition
type UserToGroupAdditionStatus struct {
	metav1alpha1.StatusMeta `json:",inline"`
}

// UserToGroupAdditionOutput defines the stack outputs
type UserToGroupAdditionOutput struct {
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iam-addusertogroup.html
	Ref string `json:"ref,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:resource:categories=aws;iam
// +kubebuilder:printcolumn:JSONPath=.status.status,description="status of the stack",name=Status,priority=0,type=string
// +kubebuilder:printcolumn:JSONPath=.status.message,description="reason for the stack status",name=Message,priority=1,type=string
// +kubebuilder:printcolumn:JSONPath=.status.stackID,description="CloudFormation Stack ID",name=StackID,priority=2,type=string

// UserToGroupAddition is the Schema for the iam UserToGroupAddition API
type UserToGroupAddition struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   UserToGroupAdditionSpec   `json:"spec,omitempty"`
	Status UserToGroupAdditionStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// UserToGroupAdditionList contains a list of Account
type UserToGroupAdditionList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`

	Items []UserToGroupAddition `json:"items"`
}

func init() {
	SchemeBuilder.Register(&UserToGroupAddition{}, &UserToGroupAdditionList{})
}
