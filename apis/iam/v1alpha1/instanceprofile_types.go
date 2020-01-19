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

// InstanceProfileSpec defines the desired state of InstanceProfile
type InstanceProfileSpec struct {
	metav1alpha1.CloudFormationMeta `json:",inline"`

	// Roles http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iam-instanceprofile.html#cfn-iam-instanceprofile-roles
	Roles []string `json:"roles" cloudformation:"Roles"`

	// InstanceProfileName http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iam-instanceprofile.html#cfn-iam-instanceprofile-instanceprofilename
	InstanceProfileName string `json:"instanceProfileName,omitempty" cloudformation:"InstanceProfileName,Parameter"`

	// Path http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iam-instanceprofile.html#cfn-iam-instanceprofile-path
	Path string `json:"path,omitempty" cloudformation:"Path,Parameter"`
}

// InstanceProfileStatus defines the observed state of InstanceProfile
type InstanceProfileStatus struct {
	metav1alpha1.StatusMeta `json:",inline"`
}

// InstanceProfileOutput defines the stack outputs
type InstanceProfileOutput struct {
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iam-instanceprofile.html
	Ref string `json:"ref,omitempty"`

	// Arn defines the Arn
	Arn string `json:"arn,omitempty" cloudformation:"Arn,Output"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:resource:categories=aws;iam
// +kubebuilder:printcolumn:JSONPath=.status.status,description="status of the stack",name=Status,priority=0,type=string
// +kubebuilder:printcolumn:JSONPath=.status.message,description="reason for the stack status",name=Message,priority=1,type=string
// +kubebuilder:printcolumn:JSONPath=.status.stackID,description="CloudFormation Stack ID",name=StackID,priority=2,type=string

// InstanceProfile is the Schema for the iam InstanceProfile API
type InstanceProfile struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   InstanceProfileSpec   `json:"spec,omitempty"`
	Status InstanceProfileStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// InstanceProfileList contains a list of Account
type InstanceProfileList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`

	Items []InstanceProfile `json:"items"`
}

func init() {
	SchemeBuilder.Register(&InstanceProfile{}, &InstanceProfileList{})
}
