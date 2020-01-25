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

// UsagePlanSpec defines the desired state of UsagePlan
type UsagePlanSpec struct {
	metav1alpha1.CloudFormationMeta `json:",inline"`

	// Description http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigateway-usageplan.html#cfn-apigateway-usageplan-description
	Description string `json:"description,omitempty" cloudformation:"Description,Parameter"`

	// Quota http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigateway-usageplan.html#cfn-apigateway-usageplan-quota
	Quota UsagePlan_QuotaSettings `json:"quota,omitempty" cloudformation:"Quota"`

	// Throttle http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigateway-usageplan.html#cfn-apigateway-usageplan-throttle
	Throttle UsagePlan_ThrottleSettings `json:"throttle,omitempty" cloudformation:"Throttle"`

	// UsagePlanName http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigateway-usageplan.html#cfn-apigateway-usageplan-usageplanname
	UsagePlanName string `json:"usagePlanName,omitempty" cloudformation:"UsagePlanName,Parameter"`

	// ApiStages http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigateway-usageplan.html#cfn-apigateway-usageplan-apistages
	ApiStages []UsagePlan_ApiStage `json:"apiStages,omitempty" cloudformation:"ApiStages"`
}

// UsagePlan_ThrottleSettings defines the desired state of UsagePlanThrottleSettings
type UsagePlan_ThrottleSettings struct {
	// BurstLimit http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-apigateway-usageplan-throttlesettings.html#cfn-apigateway-usageplan-throttlesettings-burstlimit
	BurstLimit int `json:"burstLimit,omitempty" cloudformation:"BurstLimit,Parameter"`

	// RateLimit http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-apigateway-usageplan-throttlesettings.html#cfn-apigateway-usageplan-throttlesettings-ratelimit
	RateLimit int `json:"rateLimit,omitempty" cloudformation:"RateLimit,Parameter"`
}

// UsagePlan_QuotaSettings defines the desired state of UsagePlanQuotaSettings
type UsagePlan_QuotaSettings struct {
	// Limit http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-apigateway-usageplan-quotasettings.html#cfn-apigateway-usageplan-quotasettings-limit
	Limit int `json:"limit,omitempty" cloudformation:"Limit,Parameter"`

	// Offset http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-apigateway-usageplan-quotasettings.html#cfn-apigateway-usageplan-quotasettings-offset
	Offset int `json:"offset,omitempty" cloudformation:"Offset,Parameter"`

	// Period http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-apigateway-usageplan-quotasettings.html#cfn-apigateway-usageplan-quotasettings-period
	Period string `json:"period,omitempty" cloudformation:"Period,Parameter"`
}

// UsagePlan_ApiStage defines the desired state of UsagePlanApiStage
type UsagePlan_ApiStage struct {
	// Throttle http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-apigateway-usageplan-apistage.html#cfn-apigateway-usageplan-apistage-throttle
	Throttle map[string]UsagePlan_ThrottleSettings `json:"throttle,omitempty" cloudformation:"Throttle"`

	// ApiRef http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-apigateway-usageplan-apistage.html#cfn-apigateway-usageplan-apistage-apiid
	ApiRef metav1alpha1.ObjectReference `json:"apiRef,omitempty" cloudformation:"ApiId,Parameter"`

	// Stage http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-apigateway-usageplan-apistage.html#cfn-apigateway-usageplan-apistage-stage
	Stage string `json:"stage,omitempty" cloudformation:"Stage,Parameter"`
}

// UsagePlanStatus defines the observed state of UsagePlan
type UsagePlanStatus struct {
	metav1alpha1.StatusMeta `json:",inline"`
}

// UsagePlanOutput defines the stack outputs
type UsagePlanOutput struct {
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigateway-usageplan.html
	Ref string `json:"ref,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:resource:categories=aws;apigateway
// +kubebuilder:printcolumn:JSONPath=.status.status,description="status of the stack",name=Status,priority=0,type=string
// +kubebuilder:printcolumn:JSONPath=.status.message,description="reason for the stack status",name=Message,priority=1,type=string
// +kubebuilder:printcolumn:JSONPath=.status.stackID,description="CloudFormation Stack ID",name=StackID,priority=2,type=string

// UsagePlan is the Schema for the apigateway UsagePlan API
type UsagePlan struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   UsagePlanSpec   `json:"spec,omitempty"`
	Status UsagePlanStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// UsagePlanList contains a list of Account
type UsagePlanList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`

	Items []UsagePlan `json:"items"`
}

func init() {
	SchemeBuilder.Register(&UsagePlan{}, &UsagePlanList{})
}
