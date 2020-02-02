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

// DeploymentSpec defines the desired state of Deployment
type DeploymentSpec struct {
	metav1alpha1.CloudFormationMeta `json:",inline"`

	// DeploymentCanarySettings http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigateway-deployment.html#cfn-apigateway-deployment-deploymentcanarysettings
	DeploymentCanarySettings Deployment_DeploymentCanarySettings `json:"deploymentCanarySettings,omitempty" cloudformation:"DeploymentCanarySettings"`

	// Description http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigateway-deployment.html#cfn-apigateway-deployment-description
	Description string `json:"description,omitempty" cloudformation:"Description,Parameter"`

	// RestApiRef http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigateway-deployment.html#cfn-apigateway-deployment-restapiid
	RestApiRef metav1alpha1.ObjectReference `json:"restApiRef,omitempty" cloudformation:"RestApiId,Parameter"`

	// StageDescription http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigateway-deployment.html#cfn-apigateway-deployment-stagedescription
	StageDescription Deployment_StageDescription `json:"stageDescription,omitempty" cloudformation:"StageDescription"`

	// StageName http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigateway-deployment.html#cfn-apigateway-deployment-stagename
	StageName string `json:"stageName,omitempty" cloudformation:"StageName,Parameter"`
}

// Deployment_MethodSetting defines the desired state of DeploymentMethodSetting
type Deployment_MethodSetting struct {
	// ThrottlingRateLimit http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-apigateway-deployment-stagedescription-methodsetting.html#cfn-apigateway-deployment-stagedescription-methodsetting-throttlingratelimit
	ThrottlingRateLimit int `json:"throttlingRateLimit,omitempty" cloudformation:"ThrottlingRateLimit,Parameter"`

	// CachingEnabled http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-apigateway-deployment-stagedescription-methodsetting.html#cfn-apigateway-deployment-stagedescription-methodsetting-cachingenabled
	CachingEnabled bool `json:"cachingEnabled,omitempty" cloudformation:"CachingEnabled,Parameter"`

	// HttpMethod http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-apigateway-deployment-stagedescription-methodsetting.html#cfn-apigateway-deployment-stagedescription-methodsetting-httpmethod
	HttpMethod string `json:"httpMethod,omitempty" cloudformation:"HttpMethod,Parameter"`

	// MetricsEnabled http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-apigateway-deployment-stagedescription-methodsetting.html#cfn-apigateway-deployment-stagedescription-methodsetting-metricsenabled
	MetricsEnabled bool `json:"metricsEnabled,omitempty" cloudformation:"MetricsEnabled,Parameter"`

	// ResourcePath http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-apigateway-deployment-stagedescription-methodsetting.html#cfn-apigateway-deployment-stagedescription-methodsetting-resourcepath
	ResourcePath string `json:"resourcePath,omitempty" cloudformation:"ResourcePath,Parameter"`

	// ThrottlingBurstLimit http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-apigateway-deployment-stagedescription-methodsetting.html#cfn-apigateway-deployment-stagedescription-methodsetting-throttlingburstlimit
	ThrottlingBurstLimit int `json:"throttlingBurstLimit,omitempty" cloudformation:"ThrottlingBurstLimit,Parameter"`

	// CacheDataEncrypted http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-apigateway-deployment-stagedescription-methodsetting.html#cfn-apigateway-deployment-stagedescription-methodsetting-cachedataencrypted
	CacheDataEncrypted bool `json:"cacheDataEncrypted,omitempty" cloudformation:"CacheDataEncrypted,Parameter"`

	// LoggingLevel http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-apigateway-deployment-stagedescription-methodsetting.html#cfn-apigateway-deployment-stagedescription-methodsetting-logginglevel
	LoggingLevel string `json:"loggingLevel,omitempty" cloudformation:"LoggingLevel,Parameter"`

	// CacheTtlInSeconds http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-apigateway-deployment-stagedescription-methodsetting.html#cfn-apigateway-deployment-stagedescription-methodsetting-cachettlinseconds
	CacheTtlInSeconds int `json:"cacheTtlInSeconds,omitempty" cloudformation:"CacheTtlInSeconds,Parameter"`

	// DataTraceEnabled http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-apigateway-deployment-stagedescription-methodsetting.html#cfn-apigateway-deployment-stagedescription-methodsetting-datatraceenabled
	DataTraceEnabled bool `json:"dataTraceEnabled,omitempty" cloudformation:"DataTraceEnabled,Parameter"`
}

// Deployment_AccessLogSetting defines the desired state of DeploymentAccessLogSetting
type Deployment_AccessLogSetting struct {
	// DestinationRef http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-apigateway-deployment-accesslogsetting.html#cfn-apigateway-deployment-accesslogsetting-destinationarn
	DestinationRef metav1alpha1.ObjectReference `json:"destinationRef,omitempty" cloudformation:"DestinationArn,Parameter"`

	// Format http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-apigateway-deployment-accesslogsetting.html#cfn-apigateway-deployment-accesslogsetting-format
	Format string `json:"format,omitempty" cloudformation:"Format,Parameter"`
}

// Deployment_DeploymentCanarySettings defines the desired state of DeploymentDeploymentCanarySettings
type Deployment_DeploymentCanarySettings struct {
	// PercentTraffic http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-apigateway-deployment-deploymentcanarysettings.html#cfn-apigateway-deployment-deploymentcanarysettings-percenttraffic
	PercentTraffic int `json:"percentTraffic,omitempty" cloudformation:"PercentTraffic,Parameter"`

	// StageVariableOverrides http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-apigateway-deployment-deploymentcanarysettings.html#cfn-apigateway-deployment-deploymentcanarysettings-stagevariableoverrides
	StageVariableOverrides map[string]string `json:"stageVariableOverrides,omitempty" cloudformation:"StageVariableOverrides"`

	// UseStageCache http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-apigateway-deployment-deploymentcanarysettings.html#cfn-apigateway-deployment-deploymentcanarysettings-usestagecache
	UseStageCache bool `json:"useStageCache,omitempty" cloudformation:"UseStageCache,Parameter"`
}

// Deployment_StageDescription defines the desired state of DeploymentStageDescription
type Deployment_StageDescription struct {
	// ThrottlingRateLimit http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-apigateway-deployment-stagedescription.html#cfn-apigateway-deployment-stagedescription-throttlingratelimit
	ThrottlingRateLimit int `json:"throttlingRateLimit,omitempty" cloudformation:"ThrottlingRateLimit,Parameter"`

	// CacheDataEncrypted http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-apigateway-deployment-stagedescription.html#cfn-apigateway-deployment-stagedescription-cachedataencrypted
	CacheDataEncrypted bool `json:"cacheDataEncrypted,omitempty" cloudformation:"CacheDataEncrypted,Parameter"`

	// TracingEnabled http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-apigateway-deployment-stagedescription.html#cfn-apigateway-deployment-stagedescription-tracingenabled
	TracingEnabled bool `json:"tracingEnabled,omitempty" cloudformation:"TracingEnabled,Parameter"`

	// CachingEnabled http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-apigateway-deployment-stagedescription.html#cfn-apigateway-deployment-stagedescription-cachingenabled
	CachingEnabled bool `json:"cachingEnabled,omitempty" cloudformation:"CachingEnabled,Parameter"`

	// DataTraceEnabled http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-apigateway-deployment-stagedescription.html#cfn-apigateway-deployment-stagedescription-datatraceenabled
	DataTraceEnabled bool `json:"dataTraceEnabled,omitempty" cloudformation:"DataTraceEnabled,Parameter"`

	// LoggingLevel http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-apigateway-deployment-stagedescription.html#cfn-apigateway-deployment-stagedescription-logginglevel
	LoggingLevel string `json:"loggingLevel,omitempty" cloudformation:"LoggingLevel,Parameter"`

	// ClientCertificateRef http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-apigateway-deployment-stagedescription.html#cfn-apigateway-deployment-stagedescription-clientcertificateid
	ClientCertificateRef metav1alpha1.ObjectReference `json:"clientCertificateRef,omitempty" cloudformation:"ClientCertificateId,Parameter"`

	// AccessLogSetting http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-apigateway-deployment-stagedescription.html#cfn-apigateway-deployment-stagedescription-accesslogsetting
	AccessLogSetting Deployment_AccessLogSetting `json:"accessLogSetting,omitempty" cloudformation:"AccessLogSetting"`

	// DocumentationVersion http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-apigateway-deployment-stagedescription.html#cfn-apigateway-deployment-stagedescription-documentationversion
	DocumentationVersion string `json:"documentationVersion,omitempty" cloudformation:"DocumentationVersion,Parameter"`

	// MetricsEnabled http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-apigateway-deployment-stagedescription.html#cfn-apigateway-deployment-stagedescription-metricsenabled
	MetricsEnabled bool `json:"metricsEnabled,omitempty" cloudformation:"MetricsEnabled,Parameter"`

	// ThrottlingBurstLimit http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-apigateway-deployment-stagedescription.html#cfn-apigateway-deployment-stagedescription-throttlingburstlimit
	ThrottlingBurstLimit int `json:"throttlingBurstLimit,omitempty" cloudformation:"ThrottlingBurstLimit,Parameter"`

	// CanarySetting http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-apigateway-deployment-stagedescription.html#cfn-apigateway-deployment-stagedescription-canarysetting
	CanarySetting Deployment_CanarySetting `json:"canarySetting,omitempty" cloudformation:"CanarySetting"`

	// CacheClusterEnabled http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-apigateway-deployment-stagedescription.html#cfn-apigateway-deployment-stagedescription-cacheclusterenabled
	CacheClusterEnabled bool `json:"cacheClusterEnabled,omitempty" cloudformation:"CacheClusterEnabled,Parameter"`

	// Description http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-apigateway-deployment-stagedescription.html#cfn-apigateway-deployment-stagedescription-description
	Description string `json:"description,omitempty" cloudformation:"Description,Parameter"`

	// Variables http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-apigateway-deployment-stagedescription.html#cfn-apigateway-deployment-stagedescription-variables
	Variables map[string]string `json:"variables,omitempty" cloudformation:"Variables"`

	// CacheClusterSize http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-apigateway-deployment-stagedescription.html#cfn-apigateway-deployment-stagedescription-cacheclustersize
	CacheClusterSize string `json:"cacheClusterSize,omitempty" cloudformation:"CacheClusterSize,Parameter"`

	// CacheTtlInSeconds http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-apigateway-deployment-stagedescription.html#cfn-apigateway-deployment-stagedescription-cachettlinseconds
	CacheTtlInSeconds int `json:"cacheTtlInSeconds,omitempty" cloudformation:"CacheTtlInSeconds,Parameter"`

	// MethodSettings http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-apigateway-deployment-stagedescription.html#cfn-apigateway-deployment-stagedescription-methodsettings
	MethodSettings []Deployment_MethodSetting `json:"methodSettings,omitempty" cloudformation:"MethodSettings"`
}

// Deployment_CanarySetting defines the desired state of DeploymentCanarySetting
type Deployment_CanarySetting struct {
	// PercentTraffic http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-apigateway-deployment-canarysetting.html#cfn-apigateway-deployment-canarysetting-percenttraffic
	PercentTraffic int `json:"percentTraffic,omitempty" cloudformation:"PercentTraffic,Parameter"`

	// StageVariableOverrides http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-apigateway-deployment-canarysetting.html#cfn-apigateway-deployment-canarysetting-stagevariableoverrides
	StageVariableOverrides map[string]string `json:"stageVariableOverrides,omitempty" cloudformation:"StageVariableOverrides"`

	// UseStageCache http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-apigateway-deployment-canarysetting.html#cfn-apigateway-deployment-canarysetting-usestagecache
	UseStageCache bool `json:"useStageCache,omitempty" cloudformation:"UseStageCache,Parameter"`
}

// DeploymentStatus defines the observed state of Deployment
type DeploymentStatus struct {
	metav1alpha1.StatusMeta `json:",inline"`
}

// DeploymentOutput defines the stack outputs
type DeploymentOutput struct {
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigateway-deployment.html
	Ref string `json:"ref,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:resource:categories=aws;apigateway
// +kubebuilder:printcolumn:JSONPath=.status.status,description="status of the stack",name=Status,priority=0,type=string
// +kubebuilder:printcolumn:JSONPath=.status.message,description="reason for the stack status",name=Message,priority=1,type=string
// +kubebuilder:printcolumn:JSONPath=.status.stackID,description="CloudFormation Stack ID",name=StackID,priority=2,type=string

// Deployment is the Schema for the apigateway Deployment API
type Deployment struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   DeploymentSpec   `json:"spec,omitempty"`
	Status DeploymentStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// DeploymentList contains a list of Account
type DeploymentList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`

	Items []Deployment `json:"items"`
}

func init() {
	SchemeBuilder.Register(&Deployment{}, &DeploymentList{})
}
