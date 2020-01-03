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

// HealthCheckSpec defines the desired state of HealthCheck
type HealthCheckSpec struct {
	metav1alpha1.CloudFormationMeta `json:",inline"`

	// HealthCheckConfig http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-route53-healthcheck.html#cfn-route53-healthcheck-healthcheckconfig
	HealthCheckConfig HealthCheck_HealthCheckConfig `json:"healthCheckConfig" cloudformation:"HealthCheckConfig"`

	// HealthCheckTags http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-route53-healthcheck.html#cfn-route53-healthcheck-healthchecktags
	HealthCheckTags []HealthCheck_HealthCheckTag `json:"healthCheckTags,omitempty" cloudformation:"HealthCheckTags"`
}

// HealthCheck_HealthCheckTag defines the desired state of HealthCheckHealthCheckTag
type HealthCheck_HealthCheckTag struct {
	// Key http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-route53-healthcheck-healthchecktag.html#cfn-route53-healthchecktags-key
	Key string `json:"key" cloudformation:"Key,Parameter"`

	// Value http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-route53-healthcheck-healthchecktag.html#cfn-route53-healthchecktags-value
	Value string `json:"value" cloudformation:"Value,Parameter"`
}

// HealthCheck_HealthCheckConfig defines the desired state of HealthCheckHealthCheckConfig
type HealthCheck_HealthCheckConfig struct {
	// MeasureLatency http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-route53-healthcheck-healthcheckconfig.html#cfn-route53-healthcheck-healthcheckconfig-measurelatency
	MeasureLatency bool `json:"measureLatency,omitempty" cloudformation:"MeasureLatency,Parameter"`

	// FailureThreshold http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-route53-healthcheck-healthcheckconfig.html#cfn-route53-healthcheck-healthcheckconfig-failurethreshold
	FailureThreshold int `json:"failureThreshold,omitempty" cloudformation:"FailureThreshold,Parameter"`

	// EnableSNI http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-route53-healthcheck-healthcheckconfig.html#cfn-route53-healthcheck-healthcheckconfig-enablesni
	EnableSNI bool `json:"enableSNI,omitempty" cloudformation:"EnableSNI,Parameter"`

	// Port http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-route53-healthcheck-healthcheckconfig.html#cfn-route53-healthcheck-healthcheckconfig-port
	Port int `json:"port,omitempty" cloudformation:"Port,Parameter"`

	// Regions http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-route53-healthcheck-healthcheckconfig.html#cfn-route53-healthcheck-healthcheckconfig-regions
	Regions []string `json:"regions,omitempty" cloudformation:"Regions"`

	// Type http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-route53-healthcheck-healthcheckconfig.html#cfn-route53-healthcheck-healthcheckconfig-type
	Type string `json:"type" cloudformation:"Type,Parameter"`

	// InsufficientDataHealthStatus http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-route53-healthcheck-healthcheckconfig.html#cfn-route53-healthcheck-healthcheckconfig-insufficientdatahealthstatus
	InsufficientDataHealthStatus string `json:"insufficientDataHealthStatus,omitempty" cloudformation:"InsufficientDataHealthStatus,Parameter"`

	// Inverted http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-route53-healthcheck-healthcheckconfig.html#cfn-route53-healthcheck-healthcheckconfig-inverted
	Inverted bool `json:"inverted,omitempty" cloudformation:"Inverted,Parameter"`

	// ResourcePath http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-route53-healthcheck-healthcheckconfig.html#cfn-route53-healthcheck-healthcheckconfig-resourcepath
	ResourcePath string `json:"resourcePath,omitempty" cloudformation:"ResourcePath,Parameter"`

	// RequestInterval http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-route53-healthcheck-healthcheckconfig.html#cfn-route53-healthcheck-healthcheckconfig-requestinterval
	RequestInterval int `json:"requestInterval,omitempty" cloudformation:"RequestInterval,Parameter"`

	// SearchString http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-route53-healthcheck-healthcheckconfig.html#cfn-route53-healthcheck-healthcheckconfig-searchstring
	SearchString string `json:"searchString,omitempty" cloudformation:"SearchString,Parameter"`

	// AlarmIdentifier http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-route53-healthcheck-healthcheckconfig.html#cfn-route53-healthcheck-healthcheckconfig-alarmidentifier
	AlarmIdentifier HealthCheck_AlarmIdentifier `json:"alarmIdentifier,omitempty" cloudformation:"AlarmIdentifier"`

	// FullyQualifiedDomainName http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-route53-healthcheck-healthcheckconfig.html#cfn-route53-healthcheck-healthcheckconfig-fullyqualifieddomainname
	FullyQualifiedDomainName string `json:"fullyQualifiedDomainName,omitempty" cloudformation:"FullyQualifiedDomainName,Parameter"`

	// HealthThreshold http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-route53-healthcheck-healthcheckconfig.html#cfn-route53-healthcheck-healthcheckconfig-healththreshold
	HealthThreshold int `json:"healthThreshold,omitempty" cloudformation:"HealthThreshold,Parameter"`

	// ChildHealthChecks http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-route53-healthcheck-healthcheckconfig.html#cfn-route53-healthcheck-healthcheckconfig-childhealthchecks
	ChildHealthChecks []string `json:"childHealthChecks,omitempty" cloudformation:"ChildHealthChecks"`

	// IPAddress http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-route53-healthcheck-healthcheckconfig.html#cfn-route53-healthcheck-healthcheckconfig-ipaddress
	IPAddress string `json:"iPAddress,omitempty" cloudformation:"IPAddress,Parameter"`
}

// HealthCheck_AlarmIdentifier defines the desired state of HealthCheckAlarmIdentifier
type HealthCheck_AlarmIdentifier struct {
	// Region http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-route53-healthcheck-alarmidentifier.html#cfn-route53-healthcheck-alarmidentifier-region
	Region string `json:"region" cloudformation:"Region,Parameter"`

	// Name http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-route53-healthcheck-alarmidentifier.html#cfn-route53-healthcheck-alarmidentifier-name
	Name string `json:"name" cloudformation:"Name,Parameter"`
}

// HealthCheckStatus defines the observed state of HealthCheck
type HealthCheckStatus struct {
	metav1alpha1.StatusMeta `json:",inline"`
}

// HealthCheckOutput defines the stack outputs
type HealthCheckOutput struct {
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-route53-healthcheck.html
	Ref string `json:"ref,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:resource:categories=aws;route53
// +kubebuilder:printcolumn:JSONPath=.status.status,description="status of the stack",name=Status,priority=0,type=string
// +kubebuilder:printcolumn:JSONPath=.status.message,description="reason for the stack status",name=Message,priority=1,type=string
// +kubebuilder:printcolumn:JSONPath=.status.stackID,description="CloudFormation Stack ID",name=StackID,priority=2,type=string

// HealthCheck is the Schema for the route53 HealthCheck API
type HealthCheck struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   HealthCheckSpec   `json:"spec,omitempty"`
	Status HealthCheckStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// HealthCheckList contains a list of Account
type HealthCheckList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`

	Items []HealthCheck `json:"items"`
}

func init() {
	SchemeBuilder.Register(&HealthCheck{}, &HealthCheckList{})
}
