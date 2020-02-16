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

// AccessPointSpec defines the desired state of AccessPoint
type AccessPointSpec struct {
	metav1alpha1.CloudFormationMeta `json:",inline"`

	// Bucket http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-s3-accesspoint.html#cfn-s3-accesspoint-bucket
	Bucket string `json:"bucket,omitempty" cloudformation:"Bucket,Parameter"`

	// CreationDate http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-s3-accesspoint.html#cfn-s3-accesspoint-creationdate
	CreationDate string `json:"creationDate,omitempty" cloudformation:"CreationDate,Parameter"`

	// Name http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-s3-accesspoint.html#cfn-s3-accesspoint-name
	Name string `json:"name,omitempty" cloudformation:"Name,Parameter"`

	// NetworkOrigin http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-s3-accesspoint.html#cfn-s3-accesspoint-networkorigin
	NetworkOrigin string `json:"networkOrigin,omitempty" cloudformation:"NetworkOrigin,Parameter"`

	// Policy http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-s3-accesspoint.html#cfn-s3-accesspoint-policy
	Policy string `json:"policy,omitempty" cloudformation:"Policy,Parameter"`

	// PolicyStatus http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-s3-accesspoint.html#cfn-s3-accesspoint-policystatus
	PolicyStatus string `json:"policyStatus,omitempty" cloudformation:"PolicyStatus,Parameter"`

	// PublicAccessBlockConfiguration http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-s3-accesspoint.html#cfn-s3-accesspoint-publicaccessblockconfiguration
	PublicAccessBlockConfiguration AccessPoint_PublicAccessBlockConfiguration `json:"publicAccessBlockConfiguration,omitempty" cloudformation:"PublicAccessBlockConfiguration"`

	// VpcConfiguration http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-s3-accesspoint.html#cfn-s3-accesspoint-vpcconfiguration
	VpcConfiguration AccessPoint_VpcConfiguration `json:"vpcConfiguration,omitempty" cloudformation:"VpcConfiguration"`
}

// AccessPoint_PublicAccessBlockConfiguration defines the desired state of AccessPointPublicAccessBlockConfiguration
type AccessPoint_PublicAccessBlockConfiguration struct {
	// BlockPublicAcls http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-accesspoint-publicaccessblockconfiguration.html#cfn-s3-accesspoint-publicaccessblockconfiguration-blockpublicacls
	BlockPublicAcls bool `json:"blockPublicAcls,omitempty" cloudformation:"BlockPublicAcls,Parameter"`

	// BlockPublicPolicy http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-accesspoint-publicaccessblockconfiguration.html#cfn-s3-accesspoint-publicaccessblockconfiguration-blockpublicpolicy
	BlockPublicPolicy bool `json:"blockPublicPolicy,omitempty" cloudformation:"BlockPublicPolicy,Parameter"`

	// IgnorePublicAcls http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-accesspoint-publicaccessblockconfiguration.html#cfn-s3-accesspoint-publicaccessblockconfiguration-ignorepublicacls
	IgnorePublicAcls bool `json:"ignorePublicAcls,omitempty" cloudformation:"IgnorePublicAcls,Parameter"`

	// RestrictPublicBuckets http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-accesspoint-publicaccessblockconfiguration.html#cfn-s3-accesspoint-publicaccessblockconfiguration-restrictpublicbuckets
	RestrictPublicBuckets bool `json:"restrictPublicBuckets,omitempty" cloudformation:"RestrictPublicBuckets,Parameter"`
}

// AccessPoint_VpcConfiguration defines the desired state of AccessPointVpcConfiguration
type AccessPoint_VpcConfiguration struct {
	// VpcRef http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-accesspoint-vpcconfiguration.html#cfn-s3-accesspoint-vpcconfiguration-vpcid
	VpcRef metav1alpha1.ObjectReference `json:"vpcRef,omitempty" cloudformation:"VpcId,Parameter"`
}

// AccessPointStatus defines the observed state of AccessPoint
type AccessPointStatus struct {
	metav1alpha1.StatusMeta `json:",inline"`
}

// AccessPointOutput defines the stack outputs
type AccessPointOutput struct {
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-s3-accesspoint.html
	Ref string `json:"ref,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:resource:categories=aws;s3
// +kubebuilder:printcolumn:JSONPath=.status.status,description="status of the stack",name=Status,priority=0,type=string
// +kubebuilder:printcolumn:JSONPath=.status.message,description="reason for the stack status",name=Message,priority=1,type=string
// +kubebuilder:printcolumn:JSONPath=.status.stackID,description="CloudFormation Stack ID",name=StackID,priority=2,type=string

// AccessPoint is the Schema for the s3 AccessPoint API
type AccessPoint struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   AccessPointSpec   `json:"spec,omitempty"`
	Status AccessPointStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// AccessPointList contains a list of Account
type AccessPointList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`

	Items []AccessPoint `json:"items"`
}

func init() {
	SchemeBuilder.Register(&AccessPoint{}, &AccessPointList{})
}
