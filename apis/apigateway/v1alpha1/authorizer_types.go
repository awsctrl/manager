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

// AuthorizerSpec defines the desired state of Authorizer
type AuthorizerSpec struct {
	metav1alpha1.CloudFormationMeta `json:",inline"`

	// AuthorizerResultTtlInSeconds http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigateway-authorizer.html#cfn-apigateway-authorizer-authorizerresultttlinseconds
	AuthorizerResultTtlInSeconds int `json:"authorizerResultTtlInSeconds,omitempty" cloudformation:"AuthorizerResultTtlInSeconds,Parameter"`

	// AuthorizerUri http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigateway-authorizer.html#cfn-apigateway-authorizer-authorizeruri
	AuthorizerUri string `json:"authorizerUri,omitempty" cloudformation:"AuthorizerUri,Parameter"`

	// IdentitySource http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigateway-authorizer.html#cfn-apigateway-authorizer-identitysource
	IdentitySource string `json:"identitySource,omitempty" cloudformation:"IdentitySource,Parameter"`

	// IdentityValidationExpression http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigateway-authorizer.html#cfn-apigateway-authorizer-identityvalidationexpression
	IdentityValidationExpression string `json:"identityValidationExpression,omitempty" cloudformation:"IdentityValidationExpression,Parameter"`

	// ProviderRefs http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigateway-authorizer.html#cfn-apigateway-authorizer-providerarns
	ProviderRefs []metav1alpha1.ObjectReference `json:"providerRefs,omitempty" cloudformation:"ProviderARNs"`

	// Type http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigateway-authorizer.html#cfn-apigateway-authorizer-type
	Type string `json:"type,omitempty" cloudformation:"Type,Parameter"`

	// AuthType http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigateway-authorizer.html#cfn-apigateway-authorizer-authtype
	AuthType string `json:"authType,omitempty" cloudformation:"AuthType,Parameter"`

	// RestApiRef http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigateway-authorizer.html#cfn-apigateway-authorizer-restapiid
	RestApiRef metav1alpha1.ObjectReference `json:"restApiRef,omitempty" cloudformation:"RestApiId,Parameter"`

	// AuthorizerCredentials http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigateway-authorizer.html#cfn-apigateway-authorizer-authorizercredentials
	AuthorizerCredentials string `json:"authorizerCredentials,omitempty" cloudformation:"AuthorizerCredentials,Parameter"`

	// Name http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigateway-authorizer.html#cfn-apigateway-authorizer-name
	Name string `json:"name,omitempty" cloudformation:"Name,Parameter"`
}

// AuthorizerStatus defines the observed state of Authorizer
type AuthorizerStatus struct {
	metav1alpha1.StatusMeta `json:",inline"`
}

// AuthorizerOutput defines the stack outputs
type AuthorizerOutput struct {
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigateway-authorizer.html
	Ref string `json:"ref,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:resource:categories=aws;apigateway
// +kubebuilder:printcolumn:JSONPath=.status.status,description="status of the stack",name=Status,priority=0,type=string
// +kubebuilder:printcolumn:JSONPath=.status.message,description="reason for the stack status",name=Message,priority=1,type=string
// +kubebuilder:printcolumn:JSONPath=.status.stackID,description="CloudFormation Stack ID",name=StackID,priority=2,type=string

// Authorizer is the Schema for the apigateway Authorizer API
type Authorizer struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   AuthorizerSpec   `json:"spec,omitempty"`
	Status AuthorizerStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// AuthorizerList contains a list of Account
type AuthorizerList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`

	Items []Authorizer `json:"items"`
}

func init() {
	SchemeBuilder.Register(&Authorizer{}, &AuthorizerList{})
}
