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

// ManagedPolicySpec defines the desired state of ManagedPolicy
type ManagedPolicySpec struct {
	metav1alpha1.CloudFormationMeta `json:",inline"`

	// Users http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iam-managedpolicy.html#cfn-iam-managedpolicy-users
	Users []string `json:"users,omitempty" cloudformation:"Users"`

	// Description http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iam-managedpolicy.html#cfn-iam-managedpolicy-description
	Description string `json:"description,omitempty" cloudformation:"Description,Parameter"`

	// Groups http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iam-managedpolicy.html#cfn-iam-managedpolicy-groups
	Groups []string `json:"groups,omitempty" cloudformation:"Groups"`

	// ManagedPolicyName http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iam-managedpolicy.html#cfn-iam-managedpolicy-managedpolicyname
	ManagedPolicyName string `json:"managedPolicyName,omitempty" cloudformation:"ManagedPolicyName,Parameter"`

	// Path http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iam-managedpolicy.html#cfn-ec2-dhcpoptions-path
	Path string `json:"path,omitempty" cloudformation:"Path,Parameter"`

	// PolicyDocument http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iam-managedpolicy.html#cfn-iam-managedpolicy-policydocument
	PolicyDocument string `json:"policyDocument" cloudformation:"PolicyDocument,Parameter"`

	// Roles http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iam-managedpolicy.html#cfn-iam-managedpolicy-roles
	Roles []string `json:"roles,omitempty" cloudformation:"Roles"`
}

// ManagedPolicyStatus defines the observed state of ManagedPolicy
type ManagedPolicyStatus struct {
	metav1alpha1.StatusMeta `json:",inline"`
}

// ManagedPolicyOutput defines the stack outputs
type ManagedPolicyOutput struct {
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iam-managedpolicy.html
	Ref string `json:"ref,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:resource:categories=aws;iam
// +kubebuilder:printcolumn:JSONPath=.status.status,description="status of the stack",name=Status,priority=0,type=string
// +kubebuilder:printcolumn:JSONPath=.status.message,description="reason for the stack status",name=Message,priority=1,type=string
// +kubebuilder:printcolumn:JSONPath=.status.stackID,description="CloudFormation Stack ID",name=StackID,priority=2,type=string

// ManagedPolicy is the Schema for the iam ManagedPolicy API
type ManagedPolicy struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ManagedPolicySpec   `json:"spec,omitempty"`
	Status ManagedPolicyStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ManagedPolicyList contains a list of Account
type ManagedPolicyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`

	Items []ManagedPolicy `json:"items"`
}

func init() {
	SchemeBuilder.Register(&ManagedPolicy{}, &ManagedPolicyList{})
}
