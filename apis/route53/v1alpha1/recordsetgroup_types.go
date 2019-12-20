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

// RecordSetGroupSpec defines the desired state of RecordSetGroup
type RecordSetGroupSpec struct {
	metav1alpha1.CloudFormationMeta `json:",inline"`

	// HostedZone http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-route53-recordsetgroup.html#cfn-route53-recordsetgroup-hostedzoneid
	HostedZone metav1alpha1.ObjectReference `json:"hostedZone,omitempty" cloudformation:"HostedZoneId,Parameter"`

	// HostedZoneName http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-route53-recordsetgroup.html#cfn-route53-recordsetgroup-hostedzonename
	HostedZoneName string `json:"hostedZoneName,omitempty" cloudformation:"HostedZoneName,Parameter"`

	// RecordSets http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-route53-recordsetgroup.html#cfn-route53-recordsetgroup-recordsets
	RecordSets []RecordSetGroup_RecordSet `json:"recordSets,omitempty" cloudformation:"RecordSets"`

	// Comment http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-route53-recordsetgroup.html#cfn-route53-recordsetgroup-comment
	Comment string `json:"comment,omitempty" cloudformation:"Comment,Parameter"`
}

// RecordSetGroup_AliasTarget defines the desired state of RecordSetGroupAliasTarget
type RecordSetGroup_AliasTarget struct {
	// HostedZone http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-route53-aliastarget.html#cfn-route53-aliastarget-hostedzoneid
	HostedZone metav1alpha1.ObjectReference `json:"hostedZone" cloudformation:"HostedZoneId,Parameter"`

	// DNSName http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-route53-aliastarget.html#cfn-route53-aliastarget-dnshostname
	DNSName string `json:"dNSName" cloudformation:"DNSName,Parameter"`

	// EvaluateTargetHealth http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-route53-aliastarget.html#cfn-route53-aliastarget-evaluatetargethealth
	EvaluateTargetHealth bool `json:"evaluateTargetHealth,omitempty" cloudformation:"EvaluateTargetHealth,Parameter"`
}

// RecordSetGroup_GeoLocation defines the desired state of RecordSetGroupGeoLocation
type RecordSetGroup_GeoLocation struct {
	// SubdivisionCode http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-route53-recordset-geolocation.html#cfn-route53-recordset-geolocation-subdivisioncode
	SubdivisionCode string `json:"subdivisionCode,omitempty" cloudformation:"SubdivisionCode,Parameter"`

	// ContinentCode http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-route53-recordset-geolocation.html#cfn-route53-recordsetgroup-geolocation-continentcode
	ContinentCode string `json:"continentCode,omitempty" cloudformation:"ContinentCode,Parameter"`

	// CountryCode http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-route53-recordset-geolocation.html#cfn-route53-recordset-geolocation-countrycode
	CountryCode string `json:"countryCode,omitempty" cloudformation:"CountryCode,Parameter"`
}

// RecordSetGroup_RecordSet defines the desired state of RecordSetGroupRecordSet
type RecordSetGroup_RecordSet struct {
	// SetIdentifier http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-route53-recordset.html#cfn-route53-recordset-setidentifier
	SetIdentifier string `json:"setIdentifier,omitempty" cloudformation:"SetIdentifier,Parameter"`

	// GeoLocation http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-route53-recordset.html#cfn-route53-recordset-geolocation
	GeoLocation RecordSetGroup_GeoLocation `json:"geoLocation,omitempty" cloudformation:"GeoLocation"`

	// AliasTarget http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-route53-recordset.html#cfn-route53-recordset-aliastarget
	AliasTarget RecordSetGroup_AliasTarget `json:"aliasTarget,omitempty" cloudformation:"AliasTarget"`

	// Region http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-route53-recordset.html#cfn-route53-recordset-region
	Region string `json:"region,omitempty" cloudformation:"Region,Parameter"`

	// ResourceRecords http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-route53-recordset.html#cfn-route53-recordset-resourcerecords
	ResourceRecords []string `json:"resourceRecords,omitempty" cloudformation:"ResourceRecords"`

	// MultiValueAnswer http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-route53-recordset.html#cfn-route53-recordset-multivalueanswer
	MultiValueAnswer bool `json:"multiValueAnswer,omitempty" cloudformation:"MultiValueAnswer,Parameter"`

	// HostedZoneName http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-route53-recordset.html#cfn-route53-recordset-hostedzonename
	HostedZoneName string `json:"hostedZoneName,omitempty" cloudformation:"HostedZoneName,Parameter"`

	// Comment http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-route53-recordset.html#cfn-route53-recordset-comment
	Comment string `json:"comment,omitempty" cloudformation:"Comment,Parameter"`

	// HostedZone http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-route53-recordset.html#cfn-route53-recordset-hostedzoneid
	HostedZone metav1alpha1.ObjectReference `json:"hostedZone,omitempty" cloudformation:"HostedZoneId,Parameter"`

	// HealthCheck http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-route53-recordset.html#cfn-route53-recordset-healthcheckid
	HealthCheck metav1alpha1.ObjectReference `json:"healthCheck,omitempty" cloudformation:"HealthCheckId,Parameter"`

	// Weight http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-route53-recordset.html#cfn-route53-recordset-weight
	Weight int `json:"weight,omitempty" cloudformation:"Weight,Parameter"`

	// Name http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-route53-recordset.html#cfn-route53-recordset-name
	Name string `json:"name" cloudformation:"Name,Parameter"`

	// TTL http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-route53-recordset.html#cfn-route53-recordset-ttl
	TTL string `json:"tTL,omitempty" cloudformation:"TTL,Parameter"`

	// Failover http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-route53-recordset.html#cfn-route53-recordset-failover
	Failover string `json:"failover,omitempty" cloudformation:"Failover,Parameter"`

	// Type http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-route53-recordset.html#cfn-route53-recordset-type
	Type string `json:"type" cloudformation:"Type,Parameter"`
}

// RecordSetGroupStatus defines the observed state of RecordSetGroup
type RecordSetGroupStatus struct {
	metav1alpha1.StatusMeta `json:",inline"`
}

// RecordSetGroupOutput defines the stack outputs
type RecordSetGroupOutput struct {
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-route53-recordsetgroup.html
	Ref string `json:"ref,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:resource:categories=aws;route53
// +kubebuilder:printcolumn:JSONPath=.status.status,description="status of the stack",name=Status,priority=0,type=string
// +kubebuilder:printcolumn:JSONPath=.status.message,description="reason for the stack status",name=Message,priority=1,type=string
// +kubebuilder:printcolumn:JSONPath=.status.stackID,description="CloudFormation Stack ID",name=StackID,priority=2,type=string

// RecordSetGroup is the Schema for the route53 RecordSetGroup API
type RecordSetGroup struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   RecordSetGroupSpec   `json:"spec,omitempty"`
	Status RecordSetGroupStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// RecordSetGroupList contains a list of Account
type RecordSetGroupList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`

	Items []RecordSetGroup `json:"items"`
}

func init() {
	SchemeBuilder.Register(&RecordSetGroup{}, &RecordSetGroupList{})
}
