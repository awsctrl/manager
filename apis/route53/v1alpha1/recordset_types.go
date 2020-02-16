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

// RecordSetSpec defines the desired state of RecordSet
type RecordSetSpec struct {
	metav1alpha1.CloudFormationMeta `json:",inline"`

	// AliasTarget http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-route53-recordset.html#cfn-route53-recordset-aliastarget
	AliasTarget RecordSet_AliasTarget `json:"aliasTarget,omitempty" cloudformation:"AliasTarget"`

	// Comment http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-route53-recordset.html#cfn-route53-recordset-comment
	Comment string `json:"comment,omitempty" cloudformation:"Comment,Parameter"`

	// Failover http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-route53-recordset.html#cfn-route53-recordset-failover
	Failover string `json:"failover,omitempty" cloudformation:"Failover,Parameter"`

	// GeoLocation http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-route53-recordset.html#cfn-route53-recordset-geolocation
	GeoLocation RecordSet_GeoLocation `json:"geoLocation,omitempty" cloudformation:"GeoLocation"`

	// HealthCheckRef http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-route53-recordset.html#cfn-route53-recordset-healthcheckid
	HealthCheckRef metav1alpha1.ObjectReference `json:"healthCheckRef,omitempty" cloudformation:"HealthCheckId,Parameter"`

	// HostedZoneRef http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-route53-recordset.html#cfn-route53-recordset-hostedzoneid
	HostedZoneRef metav1alpha1.ObjectReference `json:"hostedZoneRef,omitempty" cloudformation:"HostedZoneId,Parameter"`

	// HostedZoneName http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-route53-recordset.html#cfn-route53-recordset-hostedzonename
	HostedZoneName string `json:"hostedZoneName,omitempty" cloudformation:"HostedZoneName,Parameter"`

	// MultiValueAnswer http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-route53-recordset.html#cfn-route53-recordset-multivalueanswer
	MultiValueAnswer bool `json:"multiValueAnswer,omitempty" cloudformation:"MultiValueAnswer,Parameter"`

	// Name http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-route53-recordset.html#cfn-route53-recordset-name
	Name string `json:"name,omitempty" cloudformation:"Name,Parameter"`

	// Region http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-route53-recordset.html#cfn-route53-recordset-region
	Region string `json:"region,omitempty" cloudformation:"Region,Parameter"`

	// ResourceRecords http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-route53-recordset.html#cfn-route53-recordset-resourcerecords
	ResourceRecords []string `json:"resourceRecords,omitempty" cloudformation:"ResourceRecords"`

	// SetIdentifier http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-route53-recordset.html#cfn-route53-recordset-setidentifier
	SetIdentifier string `json:"setIdentifier,omitempty" cloudformation:"SetIdentifier,Parameter"`

	// TTL http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-route53-recordset.html#cfn-route53-recordset-ttl
	TTL string `json:"tTL,omitempty" cloudformation:"TTL,Parameter"`

	// Type http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-route53-recordset.html#cfn-route53-recordset-type
	Type string `json:"type,omitempty" cloudformation:"Type,Parameter"`

	// Weight http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-route53-recordset.html#cfn-route53-recordset-weight
	Weight int `json:"weight,omitempty" cloudformation:"Weight,Parameter"`
}

// RecordSet_AliasTarget defines the desired state of RecordSetAliasTarget
type RecordSet_AliasTarget struct {
	// DNSName http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-route53-aliastarget.html#cfn-route53-aliastarget-dnshostname
	DNSName string `json:"dNSName,omitempty" cloudformation:"DNSName,Parameter"`

	// EvaluateTargetHealth http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-route53-aliastarget.html#cfn-route53-aliastarget-evaluatetargethealth
	EvaluateTargetHealth bool `json:"evaluateTargetHealth,omitempty" cloudformation:"EvaluateTargetHealth,Parameter"`

	// HostedZoneRef http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-route53-aliastarget.html#cfn-route53-aliastarget-hostedzoneid
	HostedZoneRef metav1alpha1.ObjectReference `json:"hostedZoneRef,omitempty" cloudformation:"HostedZoneId,Parameter"`
}

// RecordSet_GeoLocation defines the desired state of RecordSetGeoLocation
type RecordSet_GeoLocation struct {
	// ContinentCode http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-route53-recordset-geolocation.html#cfn-route53-recordset-geolocation-continentcode
	ContinentCode string `json:"continentCode,omitempty" cloudformation:"ContinentCode,Parameter"`

	// CountryCode http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-route53-recordset-geolocation.html#cfn-route53-recordset-geolocation-countrycode
	CountryCode string `json:"countryCode,omitempty" cloudformation:"CountryCode,Parameter"`

	// SubdivisionCode http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-route53-recordset-geolocation.html#cfn-route53-recordset-geolocation-subdivisioncode
	SubdivisionCode string `json:"subdivisionCode,omitempty" cloudformation:"SubdivisionCode,Parameter"`
}

// RecordSetStatus defines the observed state of RecordSet
type RecordSetStatus struct {
	metav1alpha1.StatusMeta `json:",inline"`
}

// RecordSetOutput defines the stack outputs
type RecordSetOutput struct {
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-route53-recordset.html
	Ref string `json:"ref,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:resource:categories=aws;route53
// +kubebuilder:printcolumn:JSONPath=.status.status,description="status of the stack",name=Status,priority=0,type=string
// +kubebuilder:printcolumn:JSONPath=.status.message,description="reason for the stack status",name=Message,priority=1,type=string
// +kubebuilder:printcolumn:JSONPath=.status.stackID,description="CloudFormation Stack ID",name=StackID,priority=2,type=string

// RecordSet is the Schema for the route53 RecordSet API
type RecordSet struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   RecordSetSpec   `json:"spec,omitempty"`
	Status RecordSetStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// RecordSetList contains a list of Account
type RecordSetList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`

	Items []RecordSet `json:"items"`
}

func init() {
	SchemeBuilder.Register(&RecordSet{}, &RecordSetList{})
}
