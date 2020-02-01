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

// DomainNameSpec defines the desired state of DomainName
type DomainNameSpec struct {
	metav1alpha1.CloudFormationMeta `json:",inline"`

	// DomainName http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigateway-domainname.html#cfn-apigateway-domainname-domainname
	DomainName string `json:"domainName,omitempty" cloudformation:"DomainName,Parameter"`

	// EndpointConfiguration http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigateway-domainname.html#cfn-apigateway-domainname-endpointconfiguration
	EndpointConfiguration DomainName_EndpointConfiguration `json:"endpointConfiguration,omitempty" cloudformation:"EndpointConfiguration"`

	// RegionalCertificateRef http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigateway-domainname.html#cfn-apigateway-domainname-regionalcertificatearn
	RegionalCertificateRef metav1alpha1.ObjectReference `json:"regionalCertificateRef,omitempty" cloudformation:"RegionalCertificateArn,Parameter"`

	// SecurityPolicy http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigateway-domainname.html#cfn-apigateway-domainname-securitypolicy
	SecurityPolicy string `json:"securityPolicy,omitempty" cloudformation:"SecurityPolicy,Parameter"`

	// CertificateRef http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigateway-domainname.html#cfn-apigateway-domainname-certificatearn
	CertificateRef metav1alpha1.ObjectReference `json:"certificateRef,omitempty" cloudformation:"CertificateArn,Parameter"`
}

// DomainName_EndpointConfiguration defines the desired state of DomainNameEndpointConfiguration
type DomainName_EndpointConfiguration struct {
	// Types http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-apigateway-domainname-endpointconfiguration.html#cfn-apigateway-domainname-endpointconfiguration-types
	Types []string `json:"types,omitempty" cloudformation:"Types"`
}

// DomainNameStatus defines the observed state of DomainName
type DomainNameStatus struct {
	metav1alpha1.StatusMeta `json:",inline"`
}

// DomainNameOutput defines the stack outputs
type DomainNameOutput struct {
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigateway-domainname.html
	Ref string `json:"ref,omitempty"`

	// DistributionDomainName defines the DistributionDomainName
	DistributionDomainName string `json:"distributionDomainName,omitempty" cloudformation:"DistributionDomainName,Output"`

	// DistributionHostedZoneId defines the DistributionHostedZoneId
	DistributionHostedZoneId string `json:"distributionHostedZoneId,omitempty" cloudformation:"DistributionHostedZoneId,Output"`

	// RegionalDomainName defines the RegionalDomainName
	RegionalDomainName string `json:"regionalDomainName,omitempty" cloudformation:"RegionalDomainName,Output"`

	// RegionalHostedZoneId defines the RegionalHostedZoneId
	RegionalHostedZoneId string `json:"regionalHostedZoneId,omitempty" cloudformation:"RegionalHostedZoneId,Output"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:resource:categories=aws;apigateway
// +kubebuilder:printcolumn:JSONPath=.status.status,description="status of the stack",name=Status,priority=0,type=string
// +kubebuilder:printcolumn:JSONPath=.status.message,description="reason for the stack status",name=Message,priority=1,type=string
// +kubebuilder:printcolumn:JSONPath=.status.stackID,description="CloudFormation Stack ID",name=StackID,priority=2,type=string

// DomainName is the Schema for the apigateway DomainName API
type DomainName struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   DomainNameSpec   `json:"spec,omitempty"`
	Status DomainNameStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// DomainNameList contains a list of Account
type DomainNameList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`

	Items []DomainName `json:"items"`
}

func init() {
	SchemeBuilder.Register(&DomainName{}, &DomainNameList{})
}
