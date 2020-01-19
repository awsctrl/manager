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

// GatewayResponseSpec defines the desired state of GatewayResponse
type GatewayResponseSpec struct {
	metav1alpha1.CloudFormationMeta `json:",inline"`

	// ResponseType http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigateway-gatewayresponse.html#cfn-apigateway-gatewayresponse-responsetype
	ResponseType string `json:"responseType" cloudformation:"ResponseType,Parameter"`

	// RestApi http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigateway-gatewayresponse.html#cfn-apigateway-gatewayresponse-restapiid
	RestApi metav1alpha1.ObjectReference `json:"restApi" cloudformation:"RestApiId,Parameter"`

	// StatusCode http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigateway-gatewayresponse.html#cfn-apigateway-gatewayresponse-statuscode
	StatusCode string `json:"statusCode,omitempty" cloudformation:"StatusCode,Parameter"`

	// ResponseParameters http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigateway-gatewayresponse.html#cfn-apigateway-gatewayresponse-responseparameters
	ResponseParameters map[string]string `json:"responseParameters,omitempty" cloudformation:"ResponseParameters"`

	// ResponseTemplates http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigateway-gatewayresponse.html#cfn-apigateway-gatewayresponse-responsetemplates
	ResponseTemplates map[string]string `json:"responseTemplates,omitempty" cloudformation:"ResponseTemplates"`
}

// GatewayResponseStatus defines the observed state of GatewayResponse
type GatewayResponseStatus struct {
	metav1alpha1.StatusMeta `json:",inline"`
}

// GatewayResponseOutput defines the stack outputs
type GatewayResponseOutput struct {
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigateway-gatewayresponse.html
	Ref string `json:"ref,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:resource:categories=aws;apigateway
// +kubebuilder:printcolumn:JSONPath=.status.status,description="status of the stack",name=Status,priority=0,type=string
// +kubebuilder:printcolumn:JSONPath=.status.message,description="reason for the stack status",name=Message,priority=1,type=string
// +kubebuilder:printcolumn:JSONPath=.status.stackID,description="CloudFormation Stack ID",name=StackID,priority=2,type=string

// GatewayResponse is the Schema for the apigateway GatewayResponse API
type GatewayResponse struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   GatewayResponseSpec   `json:"spec,omitempty"`
	Status GatewayResponseStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// GatewayResponseList contains a list of Account
type GatewayResponseList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`

	Items []GatewayResponse `json:"items"`
}

func init() {
	SchemeBuilder.Register(&GatewayResponse{}, &GatewayResponseList{})
}
