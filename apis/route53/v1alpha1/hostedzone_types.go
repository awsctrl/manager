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

// HostedZoneSpec defines the desired state of HostedZone
type HostedZoneSpec struct {
	metav1alpha1.CloudFormationMeta `json:",inline"`

	// QueryLoggingConfig http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-route53-hostedzone.html#cfn-route53-hostedzone-queryloggingconfig
	QueryLoggingConfig HostedZone_QueryLoggingConfig `json:"queryLoggingConfig,omitempty" cloudformation:"QueryLoggingConfig"`

	// VPCs http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-route53-hostedzone.html#cfn-route53-hostedzone-vpcs
	VPCs []HostedZone_VPC `json:"vPCs,omitempty" cloudformation:"VPCs"`

	// HostedZoneConfig http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-route53-hostedzone.html#cfn-route53-hostedzone-hostedzoneconfig
	HostedZoneConfig HostedZone_HostedZoneConfig `json:"hostedZoneConfig,omitempty" cloudformation:"HostedZoneConfig"`

	// HostedZoneTags http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-route53-hostedzone.html#cfn-route53-hostedzone-hostedzonetags
	HostedZoneTags []HostedZone_HostedZoneTag `json:"hostedZoneTags,omitempty" cloudformation:"HostedZoneTags"`

	// Name http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-route53-hostedzone.html#cfn-route53-hostedzone-name
	Name string `json:"name" cloudformation:"Name,Parameter"`
}

// HostedZone_QueryLoggingConfig defines the desired state of HostedZoneQueryLoggingConfig
type HostedZone_QueryLoggingConfig struct {
	// CloudWatchLogsLogGroup http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-route53-hostedzone-queryloggingconfig.html#cfn-route53-hostedzone-queryloggingconfig-cloudwatchlogsloggrouparn
	CloudWatchLogsLogGroup metav1alpha1.ObjectReference `json:"cloudWatchLogsLogGroup" cloudformation:"CloudWatchLogsLogGroupArn,Parameter"`
}

// HostedZone_HostedZoneConfig defines the desired state of HostedZoneHostedZoneConfig
type HostedZone_HostedZoneConfig struct {
	// Comment http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-route53-hostedzone-hostedzoneconfig.html#cfn-route53-hostedzone-hostedzoneconfig-comment
	Comment string `json:"comment,omitempty" cloudformation:"Comment,Parameter"`
}

// HostedZone_HostedZoneTag defines the desired state of HostedZoneHostedZoneTag
type HostedZone_HostedZoneTag struct {
	// Key http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-route53-hostedzone-hostedzonetags.html#cfn-route53-hostedzonetags-key
	Key string `json:"key" cloudformation:"Key,Parameter"`

	// Value http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-route53-hostedzone-hostedzonetags.html#cfn-route53-hostedzonetags-value
	Value string `json:"value" cloudformation:"Value,Parameter"`
}

// HostedZone_VPC defines the desired state of HostedZoneVPC
type HostedZone_VPC struct {
	// VPC http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-route53-hostedzone-hostedzonevpcs.html#cfn-route53-hostedzone-hostedzonevpcs-vpcid
	VPC metav1alpha1.ObjectReference `json:"vPC" cloudformation:"VPCId,Parameter"`

	// VPCRegion http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-route53-hostedzone-hostedzonevpcs.html#cfn-route53-hostedzone-hostedzonevpcs-vpcregion
	VPCRegion string `json:"vPCRegion" cloudformation:"VPCRegion,Parameter"`
}

// HostedZoneStatus defines the observed state of HostedZone
type HostedZoneStatus struct {
	metav1alpha1.StatusMeta `json:",inline"`
}

// HostedZoneOutput defines the stack outputs
type HostedZoneOutput struct {
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-route53-hostedzone.html
	Ref string `json:"ref,omitempty"`

	// NameServers defines the NameServers
	NameServers string `json:"nameServers,omitempty" cloudformation:"NameServers,Output"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:resource:categories=aws;route53
// +kubebuilder:printcolumn:JSONPath=.status.status,description="status of the stack",name=Status,priority=0,type=string
// +kubebuilder:printcolumn:JSONPath=.status.message,description="reason for the stack status",name=Message,priority=1,type=string
// +kubebuilder:printcolumn:JSONPath=.status.stackID,description="CloudFormation Stack ID",name=StackID,priority=2,type=string

// HostedZone is the Schema for the route53 HostedZone API
type HostedZone struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   HostedZoneSpec   `json:"spec,omitempty"`
	Status HostedZoneStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// HostedZoneList contains a list of Account
type HostedZoneList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`

	Items []HostedZone `json:"items"`
}

func init() {
	SchemeBuilder.Register(&HostedZone{}, &HostedZoneList{})
}
