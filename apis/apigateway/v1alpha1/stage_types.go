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

// StageSpec defines the desired state of Stage
type StageSpec struct {
	metav1alpha1.CloudFormationMeta `json:",inline"`

	// AccessLogSetting http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigateway-stage.html#cfn-apigateway-stage-accesslogsetting
	AccessLogSetting Stage_AccessLogSetting `json:"accessLogSetting,omitempty" cloudformation:"AccessLogSetting"`

	// CacheClusterEnabled http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigateway-stage.html#cfn-apigateway-stage-cacheclusterenabled
	CacheClusterEnabled bool `json:"cacheClusterEnabled,omitempty" cloudformation:"CacheClusterEnabled,Parameter"`

	// CacheClusterSize http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigateway-stage.html#cfn-apigateway-stage-cacheclustersize
	CacheClusterSize string `json:"cacheClusterSize,omitempty" cloudformation:"CacheClusterSize,Parameter"`

	// CanarySetting http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigateway-stage.html#cfn-apigateway-stage-canarysetting
	CanarySetting Stage_CanarySetting `json:"canarySetting,omitempty" cloudformation:"CanarySetting"`

	// ClientCertificateRef http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigateway-stage.html#cfn-apigateway-stage-clientcertificateid
	ClientCertificateRef metav1alpha1.ObjectReference `json:"clientCertificateRef,omitempty" cloudformation:"ClientCertificateId,Parameter"`

	// DeploymentRef http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigateway-stage.html#cfn-apigateway-stage-deploymentid
	DeploymentRef metav1alpha1.ObjectReference `json:"deploymentRef,omitempty" cloudformation:"DeploymentId,Parameter"`

	// Description http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigateway-stage.html#cfn-apigateway-stage-description
	Description string `json:"description,omitempty" cloudformation:"Description,Parameter"`

	// DocumentationVersion http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigateway-stage.html#cfn-apigateway-stage-documentationversion
	DocumentationVersion string `json:"documentationVersion,omitempty" cloudformation:"DocumentationVersion,Parameter"`

	// MethodSettings http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigateway-stage.html#cfn-apigateway-stage-methodsettings
	MethodSettings []Stage_MethodSetting `json:"methodSettings,omitempty" cloudformation:"MethodSettings"`

	// RestApiRef http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigateway-stage.html#cfn-apigateway-stage-restapiid
	RestApiRef metav1alpha1.ObjectReference `json:"restApiRef,omitempty" cloudformation:"RestApiId,Parameter"`

	// StageName http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigateway-stage.html#cfn-apigateway-stage-stagename
	StageName string `json:"stageName,omitempty" cloudformation:"StageName,Parameter"`

	// TracingEnabled http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigateway-stage.html#cfn-apigateway-stage-tracingenabled
	TracingEnabled bool `json:"tracingEnabled,omitempty" cloudformation:"TracingEnabled,Parameter"`

	// Variables http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigateway-stage.html#cfn-apigateway-stage-variables
	Variables map[string]string `json:"variables,omitempty" cloudformation:"Variables"`
}

// Stage_AccessLogSetting defines the desired state of StageAccessLogSetting
type Stage_AccessLogSetting struct {
	// DestinationRef http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-apigateway-stage-accesslogsetting.html#cfn-apigateway-stage-accesslogsetting-destinationarn
	DestinationRef metav1alpha1.ObjectReference `json:"destinationRef,omitempty" cloudformation:"DestinationArn,Parameter"`

	// Format http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-apigateway-stage-accesslogsetting.html#cfn-apigateway-stage-accesslogsetting-format
	Format string `json:"format,omitempty" cloudformation:"Format,Parameter"`
}

// Stage_CanarySetting defines the desired state of StageCanarySetting
type Stage_CanarySetting struct {
	// DeploymentRef http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-apigateway-stage-canarysetting.html#cfn-apigateway-stage-canarysetting-deploymentid
	DeploymentRef metav1alpha1.ObjectReference `json:"deploymentRef,omitempty" cloudformation:"DeploymentId,Parameter"`

	// PercentTraffic http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-apigateway-stage-canarysetting.html#cfn-apigateway-stage-canarysetting-percenttraffic
	PercentTraffic int `json:"percentTraffic,omitempty" cloudformation:"PercentTraffic,Parameter"`

	// StageVariableOverrides http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-apigateway-stage-canarysetting.html#cfn-apigateway-stage-canarysetting-stagevariableoverrides
	StageVariableOverrides map[string]string `json:"stageVariableOverrides,omitempty" cloudformation:"StageVariableOverrides"`

	// UseStageCache http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-apigateway-stage-canarysetting.html#cfn-apigateway-stage-canarysetting-usestagecache
	UseStageCache bool `json:"useStageCache,omitempty" cloudformation:"UseStageCache,Parameter"`
}

// Stage_MethodSetting defines the desired state of StageMethodSetting
type Stage_MethodSetting struct {
	// CacheDataEncrypted http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-apitgateway-stage-methodsetting.html#cfn-apigateway-stage-methodsetting-cachedataencrypted
	CacheDataEncrypted bool `json:"cacheDataEncrypted,omitempty" cloudformation:"CacheDataEncrypted,Parameter"`

	// CacheTtlInSeconds http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-apitgateway-stage-methodsetting.html#cfn-apigateway-stage-methodsetting-cachettlinseconds
	CacheTtlInSeconds int `json:"cacheTtlInSeconds,omitempty" cloudformation:"CacheTtlInSeconds,Parameter"`

	// CachingEnabled http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-apitgateway-stage-methodsetting.html#cfn-apigateway-stage-methodsetting-cachingenabled
	CachingEnabled bool `json:"cachingEnabled,omitempty" cloudformation:"CachingEnabled,Parameter"`

	// DataTraceEnabled http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-apitgateway-stage-methodsetting.html#cfn-apigateway-stage-methodsetting-datatraceenabled
	DataTraceEnabled bool `json:"dataTraceEnabled,omitempty" cloudformation:"DataTraceEnabled,Parameter"`

	// HttpMethod http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-apitgateway-stage-methodsetting.html#cfn-apigateway-stage-methodsetting-httpmethod
	HttpMethod string `json:"httpMethod,omitempty" cloudformation:"HttpMethod,Parameter"`

	// LoggingLevel http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-apitgateway-stage-methodsetting.html#cfn-apigateway-stage-methodsetting-logginglevel
	LoggingLevel string `json:"loggingLevel,omitempty" cloudformation:"LoggingLevel,Parameter"`

	// MetricsEnabled http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-apitgateway-stage-methodsetting.html#cfn-apigateway-stage-methodsetting-metricsenabled
	MetricsEnabled bool `json:"metricsEnabled,omitempty" cloudformation:"MetricsEnabled,Parameter"`

	// ResourcePath http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-apitgateway-stage-methodsetting.html#cfn-apigateway-stage-methodsetting-resourcepath
	ResourcePath string `json:"resourcePath,omitempty" cloudformation:"ResourcePath,Parameter"`

	// ThrottlingBurstLimit http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-apitgateway-stage-methodsetting.html#cfn-apigateway-stage-methodsetting-throttlingburstlimit
	ThrottlingBurstLimit int `json:"throttlingBurstLimit,omitempty" cloudformation:"ThrottlingBurstLimit,Parameter"`

	// ThrottlingRateLimit http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-apitgateway-stage-methodsetting.html#cfn-apigateway-stage-methodsetting-throttlingratelimit
	ThrottlingRateLimit int `json:"throttlingRateLimit,omitempty" cloudformation:"ThrottlingRateLimit,Parameter"`
}

// StageStatus defines the observed state of Stage
type StageStatus struct {
	metav1alpha1.StatusMeta `json:",inline"`
}

// StageOutput defines the stack outputs
type StageOutput struct {
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigateway-stage.html
	Ref string `json:"ref,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:resource:categories=aws;apigateway
// +kubebuilder:printcolumn:JSONPath=.status.status,description="status of the stack",name=Status,priority=0,type=string
// +kubebuilder:printcolumn:JSONPath=.status.message,description="reason for the stack status",name=Message,priority=1,type=string
// +kubebuilder:printcolumn:JSONPath=.status.stackID,description="CloudFormation Stack ID",name=StackID,priority=2,type=string

// Stage is the Schema for the apigateway Stage API
type Stage struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   StageSpec   `json:"spec,omitempty"`
	Status StageStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// StageList contains a list of Account
type StageList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`

	Items []Stage `json:"items"`
}

func init() {
	SchemeBuilder.Register(&Stage{}, &StageList{})
}
