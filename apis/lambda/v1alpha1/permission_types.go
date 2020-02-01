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

// PermissionSpec defines the desired state of Permission
type PermissionSpec struct {
	metav1alpha1.CloudFormationMeta `json:",inline"`

	// Action http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lambda-permission.html#cfn-lambda-permission-action
	Action string `json:"action,omitempty" cloudformation:"Action,Parameter"`

	// EventSourceToken http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lambda-permission.html#cfn-lambda-permission-eventsourcetoken
	EventSourceToken string `json:"eventSourceToken,omitempty" cloudformation:"EventSourceToken,Parameter"`

	// FunctionName http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lambda-permission.html#cfn-lambda-permission-functionname
	FunctionName string `json:"functionName,omitempty" cloudformation:"FunctionName,Parameter"`

	// Principal http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lambda-permission.html#cfn-lambda-permission-principal
	Principal string `json:"principal,omitempty" cloudformation:"Principal,Parameter"`

	// SourceAccount http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lambda-permission.html#cfn-lambda-permission-sourceaccount
	SourceAccount string `json:"sourceAccount,omitempty" cloudformation:"SourceAccount,Parameter"`

	// SourceRef http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lambda-permission.html#cfn-lambda-permission-sourcearn
	SourceRef metav1alpha1.ObjectReference `json:"sourceRef,omitempty" cloudformation:"SourceArn,Parameter"`
}

// PermissionStatus defines the observed state of Permission
type PermissionStatus struct {
	metav1alpha1.StatusMeta `json:",inline"`
}

// PermissionOutput defines the stack outputs
type PermissionOutput struct {
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lambda-permission.html
	Ref string `json:"ref,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:resource:categories=aws;lambda
// +kubebuilder:printcolumn:JSONPath=.status.status,description="status of the stack",name=Status,priority=0,type=string
// +kubebuilder:printcolumn:JSONPath=.status.message,description="reason for the stack status",name=Message,priority=1,type=string
// +kubebuilder:printcolumn:JSONPath=.status.stackID,description="CloudFormation Stack ID",name=StackID,priority=2,type=string

// Permission is the Schema for the lambda Permission API
type Permission struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   PermissionSpec   `json:"spec,omitempty"`
	Status PermissionStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// PermissionList contains a list of Account
type PermissionList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`

	Items []Permission `json:"items"`
}

func init() {
	SchemeBuilder.Register(&Permission{}, &PermissionList{})
}
