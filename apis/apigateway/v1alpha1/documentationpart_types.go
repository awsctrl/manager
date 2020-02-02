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

// DocumentationPartSpec defines the desired state of DocumentationPart
type DocumentationPartSpec struct {
	metav1alpha1.CloudFormationMeta `json:",inline"`

	// Properties http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigateway-documentationpart.html#cfn-apigateway-documentationpart-properties
	Properties string `json:"properties,omitempty" cloudformation:"Properties,Parameter"`

	// RestApiRef http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigateway-documentationpart.html#cfn-apigateway-documentationpart-restapiid
	RestApiRef metav1alpha1.ObjectReference `json:"restApiRef,omitempty" cloudformation:"RestApiId,Parameter"`

	// Location http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigateway-documentationpart.html#cfn-apigateway-documentationpart-location
	Location DocumentationPart_Location `json:"location,omitempty" cloudformation:"Location"`
}

// DocumentationPart_Location defines the desired state of DocumentationPartLocation
type DocumentationPart_Location struct {
	// Type http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-apigateway-documentationpart-location.html#cfn-apigateway-documentationpart-location-type
	Type string `json:"type,omitempty" cloudformation:"Type,Parameter"`

	// Method http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-apigateway-documentationpart-location.html#cfn-apigateway-documentationpart-location-method
	Method string `json:"method,omitempty" cloudformation:"Method,Parameter"`

	// Name http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-apigateway-documentationpart-location.html#cfn-apigateway-documentationpart-location-name
	Name string `json:"name,omitempty" cloudformation:"Name,Parameter"`

	// Path http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-apigateway-documentationpart-location.html#cfn-apigateway-documentationpart-location-path
	Path string `json:"path,omitempty" cloudformation:"Path,Parameter"`

	// StatusCode http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-apigateway-documentationpart-location.html#cfn-apigateway-documentationpart-location-statuscode
	StatusCode string `json:"statusCode,omitempty" cloudformation:"StatusCode,Parameter"`
}

// DocumentationPartStatus defines the observed state of DocumentationPart
type DocumentationPartStatus struct {
	metav1alpha1.StatusMeta `json:",inline"`
}

// DocumentationPartOutput defines the stack outputs
type DocumentationPartOutput struct {
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigateway-documentationpart.html
	Ref string `json:"ref,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:resource:categories=aws;apigateway
// +kubebuilder:printcolumn:JSONPath=.status.status,description="status of the stack",name=Status,priority=0,type=string
// +kubebuilder:printcolumn:JSONPath=.status.message,description="reason for the stack status",name=Message,priority=1,type=string
// +kubebuilder:printcolumn:JSONPath=.status.stackID,description="CloudFormation Stack ID",name=StackID,priority=2,type=string

// DocumentationPart is the Schema for the apigateway DocumentationPart API
type DocumentationPart struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   DocumentationPartSpec   `json:"spec,omitempty"`
	Status DocumentationPartStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// DocumentationPartList contains a list of Account
type DocumentationPartList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`

	Items []DocumentationPart `json:"items"`
}

func init() {
	SchemeBuilder.Register(&DocumentationPart{}, &DocumentationPartList{})
}
