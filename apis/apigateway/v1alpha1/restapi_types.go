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

// RestApiSpec defines the desired state of RestApi
type RestApiSpec struct {
	metav1alpha1.CloudFormationMeta `json:",inline"`

	// Body http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigateway-restapi.html#cfn-apigateway-restapi-body
	Body string `json:"body,omitempty" cloudformation:"Body,Parameter"`

	// BodyS3Location http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigateway-restapi.html#cfn-apigateway-restapi-bodys3location
	BodyS3Location RestApi_S3Location `json:"bodyS3Location,omitempty" cloudformation:"BodyS3Location"`

	// Parameters http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigateway-restapi.html#cfn-apigateway-restapi-parameters
	Parameters map[string]string `json:"parameters,omitempty" cloudformation:"Parameters"`

	// MinimumCompressionSize http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigateway-restapi.html#cfn-apigateway-restapi-minimumcompressionsize
	MinimumCompressionSize int `json:"minimumCompressionSize,omitempty" cloudformation:"MinimumCompressionSize,Parameter"`

	// EndpointConfiguration http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigateway-restapi.html#cfn-apigateway-restapi-endpointconfiguration
	EndpointConfiguration RestApi_EndpointConfiguration `json:"endpointConfiguration,omitempty" cloudformation:"EndpointConfiguration"`

	// FailOnWarnings http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigateway-restapi.html#cfn-apigateway-restapi-failonwarnings
	FailOnWarnings bool `json:"failOnWarnings,omitempty" cloudformation:"FailOnWarnings,Parameter"`

	// ApiKeySourceType http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigateway-restapi.html#cfn-apigateway-restapi-apikeysourcetype
	ApiKeySourceType string `json:"apiKeySourceType,omitempty" cloudformation:"ApiKeySourceType,Parameter"`

	// Description http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigateway-restapi.html#cfn-apigateway-restapi-description
	Description string `json:"description,omitempty" cloudformation:"Description,Parameter"`

	// Name http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigateway-restapi.html#cfn-apigateway-restapi-name
	Name string `json:"name,omitempty" cloudformation:"Name,Parameter"`

	// Policy http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigateway-restapi.html#cfn-apigateway-restapi-policy
	Policy string `json:"policy,omitempty" cloudformation:"Policy,Parameter"`

	// BinaryMediaTypes http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigateway-restapi.html#cfn-apigateway-restapi-binarymediatypes
	BinaryMediaTypes []string `json:"binaryMediaTypes,omitempty" cloudformation:"BinaryMediaTypes"`

	// CloneFrom http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigateway-restapi.html#cfn-apigateway-restapi-clonefrom
	CloneFrom string `json:"cloneFrom,omitempty" cloudformation:"CloneFrom,Parameter"`
}

// RestApi_EndpointConfiguration defines the desired state of RestApiEndpointConfiguration
type RestApi_EndpointConfiguration struct {
	// Types http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-apigateway-restapi-endpointconfiguration.html#cfn-apigateway-restapi-endpointconfiguration-types
	Types []string `json:"types,omitempty" cloudformation:"Types"`

	// VpcEndpoint http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-apigateway-restapi-endpointconfiguration.html#cfn-apigateway-restapi-endpointconfiguration-vpcendpointids
	VpcEndpoint []metav1alpha1.ObjectReference `json:"vpcEndpoint,omitempty" cloudformation:"VpcEndpointIds"`
}

// RestApi_S3Location defines the desired state of RestApiS3Location
type RestApi_S3Location struct {
	// Bucket http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-apigateway-restapi-s3location.html#cfn-apigateway-restapi-s3location-bucket
	Bucket string `json:"bucket,omitempty" cloudformation:"Bucket,Parameter"`

	// ETag http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-apigateway-restapi-s3location.html#cfn-apigateway-restapi-s3location-etag
	ETag string `json:"eTag,omitempty" cloudformation:"ETag,Parameter"`

	// Key http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-apigateway-restapi-s3location.html#cfn-apigateway-restapi-s3location-key
	Key string `json:"key,omitempty" cloudformation:"Key,Parameter"`

	// Version http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-apigateway-restapi-s3location.html#cfn-apigateway-restapi-s3location-version
	Version string `json:"version,omitempty" cloudformation:"Version,Parameter"`
}

// RestApiStatus defines the observed state of RestApi
type RestApiStatus struct {
	metav1alpha1.StatusMeta `json:",inline"`
}

// RestApiOutput defines the stack outputs
type RestApiOutput struct {
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigateway-restapi.html
	Ref string `json:"ref,omitempty"`

	// RootResourceId defines the RootResourceId
	RootResourceId string `json:"rootResourceId,omitempty" cloudformation:"RootResourceId,Output"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:resource:categories=aws;apigateway
// +kubebuilder:printcolumn:JSONPath=.status.status,description="status of the stack",name=Status,priority=0,type=string
// +kubebuilder:printcolumn:JSONPath=.status.message,description="reason for the stack status",name=Message,priority=1,type=string
// +kubebuilder:printcolumn:JSONPath=.status.stackID,description="CloudFormation Stack ID",name=StackID,priority=2,type=string

// RestApi is the Schema for the apigateway RestApi API
type RestApi struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   RestApiSpec   `json:"spec,omitempty"`
	Status RestApiStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// RestApiList contains a list of Account
type RestApiList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`

	Items []RestApi `json:"items"`
}

func init() {
	SchemeBuilder.Register(&RestApi{}, &RestApiList{})
}
