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

// ApiKeySpec defines the desired state of ApiKey
type ApiKeySpec struct {
	metav1alpha1.CloudFormationMeta `json:",inline"`

	// CustomerRef http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigateway-apikey.html#cfn-apigateway-apikey-customerid
	CustomerRef metav1alpha1.ObjectReference `json:"customerRef,omitempty" cloudformation:"CustomerId,Parameter"`

	// Description http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigateway-apikey.html#cfn-apigateway-apikey-description
	Description string `json:"description,omitempty" cloudformation:"Description,Parameter"`

	// Enabled http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigateway-apikey.html#cfn-apigateway-apikey-enabled
	Enabled bool `json:"enabled,omitempty" cloudformation:"Enabled,Parameter"`

	// GenerateDistinctId http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigateway-apikey.html#cfn-apigateway-apikey-generatedistinctid
	GenerateDistinctId bool `json:"generateDistinctId,omitempty" cloudformation:"GenerateDistinctId,Parameter"`

	// Name http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigateway-apikey.html#cfn-apigateway-apikey-name
	Name string `json:"name,omitempty" cloudformation:"Name,Parameter"`

	// StageKeys http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigateway-apikey.html#cfn-apigateway-apikey-stagekeys
	StageKeys []ApiKey_StageKey `json:"stageKeys,omitempty" cloudformation:"StageKeys"`

	// Value http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigateway-apikey.html#cfn-apigateway-apikey-value
	Value string `json:"value,omitempty" cloudformation:"Value,Parameter"`
}

// ApiKey_StageKey defines the desired state of ApiKeyStageKey
type ApiKey_StageKey struct {
	// RestApiRef http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-apigateway-apikey-stagekey.html#cfn-apigateway-apikey-stagekey-restapiid
	RestApiRef metav1alpha1.ObjectReference `json:"restApiRef,omitempty" cloudformation:"RestApiId,Parameter"`

	// StageName http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-apigateway-apikey-stagekey.html#cfn-apigateway-apikey-stagekey-stagename
	StageName string `json:"stageName,omitempty" cloudformation:"StageName,Parameter"`
}

// ApiKeyStatus defines the observed state of ApiKey
type ApiKeyStatus struct {
	metav1alpha1.StatusMeta `json:",inline"`
}

// ApiKeyOutput defines the stack outputs
type ApiKeyOutput struct {
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigateway-apikey.html
	Ref string `json:"ref,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:resource:categories=aws;apigateway
// +kubebuilder:printcolumn:JSONPath=.status.status,description="status of the stack",name=Status,priority=0,type=string
// +kubebuilder:printcolumn:JSONPath=.status.message,description="reason for the stack status",name=Message,priority=1,type=string
// +kubebuilder:printcolumn:JSONPath=.status.stackID,description="CloudFormation Stack ID",name=StackID,priority=2,type=string

// ApiKey is the Schema for the apigateway ApiKey API
type ApiKey struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ApiKeySpec   `json:"spec,omitempty"`
	Status ApiKeyStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ApiKeyList contains a list of Account
type ApiKeyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`

	Items []ApiKey `json:"items"`
}

func init() {
	SchemeBuilder.Register(&ApiKey{}, &ApiKeyList{})
}
