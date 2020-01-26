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

// MethodSpec defines the desired state of Method
type MethodSpec struct {
	metav1alpha1.CloudFormationMeta `json:",inline"`

	// AuthorizationType http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigateway-method.html#cfn-apigateway-method-authorizationtype
	AuthorizationType string `json:"authorizationType,omitempty" cloudformation:"AuthorizationType,Parameter"`

	// AuthorizerRef http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigateway-method.html#cfn-apigateway-method-authorizerid
	AuthorizerRef metav1alpha1.ObjectReference `json:"authorizerRef,omitempty" cloudformation:"AuthorizerId,Parameter"`

	// MethodResponses http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigateway-method.html#cfn-apigateway-method-methodresponses
	MethodResponses []Method_MethodResponse `json:"methodResponses,omitempty" cloudformation:"MethodResponses"`

	// AuthorizationScopes http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigateway-method.html#cfn-apigateway-method-authorizationscopes
	AuthorizationScopes []string `json:"authorizationScopes,omitempty" cloudformation:"AuthorizationScopes"`

	// Integration http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigateway-method.html#cfn-apigateway-method-integration
	Integration Method_Integration `json:"integration,omitempty" cloudformation:"Integration"`

	// OperationName http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigateway-method.html#cfn-apigateway-method-operationname
	OperationName string `json:"operationName,omitempty" cloudformation:"OperationName,Parameter"`

	// ResourceRef http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigateway-method.html#cfn-apigateway-method-resourceid
	ResourceRef metav1alpha1.ObjectReference `json:"resourceRef,omitempty" cloudformation:"ResourceId,Parameter"`

	// ApiKeyRequired http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigateway-method.html#cfn-apigateway-method-apikeyrequired
	ApiKeyRequired bool `json:"apiKeyRequired,omitempty" cloudformation:"ApiKeyRequired,Parameter"`

	// RequestParameters http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigateway-method.html#cfn-apigateway-method-requestparameters
	RequestParameters map[string]bool `json:"requestParameters,omitempty" cloudformation:"RequestParameters"`

	// RestApiRef http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigateway-method.html#cfn-apigateway-method-restapiid
	RestApiRef metav1alpha1.ObjectReference `json:"restApiRef,omitempty" cloudformation:"RestApiId,Parameter"`

	// HttpMethod http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigateway-method.html#cfn-apigateway-method-httpmethod
	HttpMethod string `json:"httpMethod,omitempty" cloudformation:"HttpMethod,Parameter"`

	// RequestModels http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigateway-method.html#cfn-apigateway-method-requestmodels
	RequestModels map[string]string `json:"requestModels,omitempty" cloudformation:"RequestModels"`

	// RequestValidatorRef http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigateway-method.html#cfn-apigateway-method-requestvalidatorid
	RequestValidatorRef metav1alpha1.ObjectReference `json:"requestValidatorRef,omitempty" cloudformation:"RequestValidatorId,Parameter"`
}

// Method_Integration defines the desired state of MethodIntegration
type Method_Integration struct {
	// Uri http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-apitgateway-method-integration.html#cfn-apigateway-method-integration-uri
	Uri string `json:"uri,omitempty" cloudformation:"Uri,Parameter"`

	// ConnectionType http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-apitgateway-method-integration.html#cfn-apigateway-method-integration-connectiontype
	ConnectionType string `json:"connectionType,omitempty" cloudformation:"ConnectionType,Parameter"`

	// Type http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-apitgateway-method-integration.html#cfn-apigateway-method-integration-type
	Type string `json:"type,omitempty" cloudformation:"Type,Parameter"`

	// CacheNamespace http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-apitgateway-method-integration.html#cfn-apigateway-method-integration-cachenamespace
	CacheNamespace string `json:"cacheNamespace,omitempty" cloudformation:"CacheNamespace,Parameter"`

	// ContentHandling http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-apitgateway-method-integration.html#cfn-apigateway-method-integration-contenthandling
	ContentHandling string `json:"contentHandling,omitempty" cloudformation:"ContentHandling,Parameter"`

	// RequestParameters http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-apitgateway-method-integration.html#cfn-apigateway-method-integration-requestparameters
	RequestParameters map[string]string `json:"requestParameters,omitempty" cloudformation:"RequestParameters"`

	// Credentials http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-apitgateway-method-integration.html#cfn-apigateway-method-integration-credentials
	Credentials string `json:"credentials,omitempty" cloudformation:"Credentials,Parameter"`

	// PassthroughBehavior http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-apitgateway-method-integration.html#cfn-apigateway-method-integration-passthroughbehavior
	PassthroughBehavior string `json:"passthroughBehavior,omitempty" cloudformation:"PassthroughBehavior,Parameter"`

	// ConnectionRef http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-apitgateway-method-integration.html#cfn-apigateway-method-integration-connectionid
	ConnectionRef metav1alpha1.ObjectReference `json:"connectionRef,omitempty" cloudformation:"ConnectionId,Parameter"`

	// IntegrationHttpMethod http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-apitgateway-method-integration.html#cfn-apigateway-method-integration-integrationhttpmethod
	IntegrationHttpMethod string `json:"integrationHttpMethod,omitempty" cloudformation:"IntegrationHttpMethod,Parameter"`

	// RequestTemplates http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-apitgateway-method-integration.html#cfn-apigateway-method-integration-requesttemplates
	RequestTemplates map[string]string `json:"requestTemplates,omitempty" cloudformation:"RequestTemplates"`

	// TimeoutInMillis http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-apitgateway-method-integration.html#cfn-apigateway-method-integration-timeoutinmillis
	TimeoutInMillis int `json:"timeoutInMillis,omitempty" cloudformation:"TimeoutInMillis,Parameter"`

	// CacheKeyParameters http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-apitgateway-method-integration.html#cfn-apigateway-method-integration-cachekeyparameters
	CacheKeyParameters []string `json:"cacheKeyParameters,omitempty" cloudformation:"CacheKeyParameters"`

	// IntegrationResponses http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-apitgateway-method-integration.html#cfn-apigateway-method-integration-integrationresponses
	IntegrationResponses []Method_IntegrationResponse `json:"integrationResponses,omitempty" cloudformation:"IntegrationResponses"`
}

// Method_IntegrationResponse defines the desired state of MethodIntegrationResponse
type Method_IntegrationResponse struct {
	// ContentHandling http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-apitgateway-method-integration-integrationresponse.html#cfn-apigateway-method-integrationresponse-contenthandling
	ContentHandling string `json:"contentHandling,omitempty" cloudformation:"ContentHandling,Parameter"`

	// ResponseParameters http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-apitgateway-method-integration-integrationresponse.html#cfn-apigateway-method-integration-integrationresponse-responseparameters
	ResponseParameters map[string]string `json:"responseParameters,omitempty" cloudformation:"ResponseParameters"`

	// ResponseTemplates http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-apitgateway-method-integration-integrationresponse.html#cfn-apigateway-method-integration-integrationresponse-responsetemplates
	ResponseTemplates map[string]string `json:"responseTemplates,omitempty" cloudformation:"ResponseTemplates"`

	// SelectionPattern http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-apitgateway-method-integration-integrationresponse.html#cfn-apigateway-method-integration-integrationresponse-selectionpattern
	SelectionPattern string `json:"selectionPattern,omitempty" cloudformation:"SelectionPattern,Parameter"`

	// StatusCode http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-apitgateway-method-integration-integrationresponse.html#cfn-apigateway-method-integration-integrationresponse-statuscode
	StatusCode string `json:"statusCode,omitempty" cloudformation:"StatusCode,Parameter"`
}

// Method_MethodResponse defines the desired state of MethodMethodResponse
type Method_MethodResponse struct {
	// ResponseModels http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-apitgateway-method-methodresponse.html#cfn-apigateway-method-methodresponse-responsemodels
	ResponseModels map[string]string `json:"responseModels,omitempty" cloudformation:"ResponseModels"`

	// ResponseParameters http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-apitgateway-method-methodresponse.html#cfn-apigateway-method-methodresponse-responseparameters
	ResponseParameters map[string]bool `json:"responseParameters,omitempty" cloudformation:"ResponseParameters"`

	// StatusCode http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-apitgateway-method-methodresponse.html#cfn-apigateway-method-methodresponse-statuscode
	StatusCode string `json:"statusCode,omitempty" cloudformation:"StatusCode,Parameter"`
}

// MethodStatus defines the observed state of Method
type MethodStatus struct {
	metav1alpha1.StatusMeta `json:",inline"`
}

// MethodOutput defines the stack outputs
type MethodOutput struct {
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigateway-method.html
	Ref string `json:"ref,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:resource:categories=aws;apigateway
// +kubebuilder:printcolumn:JSONPath=.status.status,description="status of the stack",name=Status,priority=0,type=string
// +kubebuilder:printcolumn:JSONPath=.status.message,description="reason for the stack status",name=Message,priority=1,type=string
// +kubebuilder:printcolumn:JSONPath=.status.stackID,description="CloudFormation Stack ID",name=StackID,priority=2,type=string

// Method is the Schema for the apigateway Method API
type Method struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   MethodSpec   `json:"spec,omitempty"`
	Status MethodStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// MethodList contains a list of Account
type MethodList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`

	Items []Method `json:"items"`
}

func init() {
	SchemeBuilder.Register(&Method{}, &MethodList{})
}
