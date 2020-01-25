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

// RoleSpec defines the desired state of Role
type RoleSpec struct {
	metav1alpha1.CloudFormationMeta `json:",inline"`

	// ManagedPolicyRefs http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iam-role.html#cfn-iam-role-managepolicyarns
	ManagedPolicyRefs []metav1alpha1.ObjectReference `json:"managedPolicyRefs,omitempty" cloudformation:"ManagedPolicyArns"`

	// MaxSessionDuration http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iam-role.html#cfn-iam-role-maxsessionduration
	MaxSessionDuration int `json:"maxSessionDuration,omitempty" cloudformation:"MaxSessionDuration,Parameter"`

	// Policies http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iam-role.html#cfn-iam-role-policies
	Policies []Role_Policy `json:"policies,omitempty" cloudformation:"Policies"`

	// AssumeRolePolicyDocument http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iam-role.html#cfn-iam-role-assumerolepolicydocument
	AssumeRolePolicyDocument string `json:"assumeRolePolicyDocument,omitempty" cloudformation:"AssumeRolePolicyDocument,Parameter"`

	// Description http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iam-role.html#cfn-iam-role-description
	Description string `json:"description,omitempty" cloudformation:"Description,Parameter"`

	// RoleName http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iam-role.html#cfn-iam-role-rolename
	RoleName string `json:"roleName,omitempty" cloudformation:"RoleName,Parameter"`

	// PermissionsBoundary http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iam-role.html#cfn-iam-role-permissionsboundary
	PermissionsBoundary string `json:"permissionsBoundary,omitempty" cloudformation:"PermissionsBoundary,Parameter"`

	// Path http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iam-role.html#cfn-iam-role-path
	Path string `json:"path,omitempty" cloudformation:"Path,Parameter"`
}

// Role_Policy defines the desired state of RolePolicy
type Role_Policy struct {
	// PolicyDocument http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iam-policy.html#cfn-iam-policies-policydocument
	PolicyDocument string `json:"policyDocument,omitempty" cloudformation:"PolicyDocument,Parameter"`

	// PolicyName http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iam-policy.html#cfn-iam-policies-policyname
	PolicyName string `json:"policyName,omitempty" cloudformation:"PolicyName,Parameter"`
}

// RoleStatus defines the observed state of Role
type RoleStatus struct {
	metav1alpha1.StatusMeta `json:",inline"`
}

// RoleOutput defines the stack outputs
type RoleOutput struct {
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iam-role.html
	Ref string `json:"ref,omitempty"`

	// Arn defines the Arn
	Arn string `json:"arn,omitempty" cloudformation:"Arn,Output"`

	// RoleId defines the RoleId
	RoleId string `json:"roleId,omitempty" cloudformation:"RoleId,Output"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:resource:categories=aws;iam
// +kubebuilder:printcolumn:JSONPath=.status.status,description="status of the stack",name=Status,priority=0,type=string
// +kubebuilder:printcolumn:JSONPath=.status.message,description="reason for the stack status",name=Message,priority=1,type=string
// +kubebuilder:printcolumn:JSONPath=.status.stackID,description="CloudFormation Stack ID",name=StackID,priority=2,type=string

// Role is the Schema for the iam Role API
type Role struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   RoleSpec   `json:"spec,omitempty"`
	Status RoleStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// RoleList contains a list of Account
type RoleList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`

	Items []Role `json:"items"`
}

func init() {
	SchemeBuilder.Register(&Role{}, &RoleList{})
}
