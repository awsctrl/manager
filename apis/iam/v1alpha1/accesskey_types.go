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

// AccessKeySpec defines the desired state of AccessKey
type AccessKeySpec struct {
	metav1alpha1.CloudFormationMeta `json:",inline"`

	// UserName http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iam-accesskey.html#cfn-iam-accesskey-username
	UserName string `json:"userName" cloudformation:"UserName,Parameter"`

	// Serial http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iam-accesskey.html#cfn-iam-accesskey-serial
	Serial int `json:"serial,omitempty" cloudformation:"Serial,Parameter"`

	// Status http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iam-accesskey.html#cfn-iam-accesskey-status
	Status string `json:"status,omitempty" cloudformation:"Status,Parameter"`
}

// AccessKeyStatus defines the observed state of AccessKey
type AccessKeyStatus struct {
	metav1alpha1.StatusMeta `json:",inline"`
}

// AccessKeyOutput defines the stack outputs
type AccessKeyOutput struct {
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iam-accesskey.html
	Ref string `json:"ref,omitempty"`

	// SecretAccessKey defines the SecretAccessKey
	SecretAccessKey string `json:"secretAccessKey,omitempty" cloudformation:"SecretAccessKey,Output"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:resource:categories=aws;iam
// +kubebuilder:printcolumn:JSONPath=.status.status,description="status of the stack",name=Status,priority=0,type=string
// +kubebuilder:printcolumn:JSONPath=.status.message,description="reason for the stack status",name=Message,priority=1,type=string
// +kubebuilder:printcolumn:JSONPath=.status.stackID,description="CloudFormation Stack ID",name=StackID,priority=2,type=string

// AccessKey is the Schema for the iam AccessKey API
type AccessKey struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   AccessKeySpec   `json:"spec,omitempty"`
	Status AccessKeyStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// AccessKeyList contains a list of Account
type AccessKeyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`

	Items []AccessKey `json:"items"`
}

func init() {
	SchemeBuilder.Register(&AccessKey{}, &AccessKeyList{})
}
