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

// EventInvokeConfigSpec defines the desired state of EventInvokeConfig
type EventInvokeConfigSpec struct {
	metav1alpha1.CloudFormationMeta `json:",inline"`

	// DestinationConfig http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lambda-eventinvokeconfig.html#cfn-lambda-eventinvokeconfig-destinationconfig
	DestinationConfig EventInvokeConfig_DestinationConfig `json:"destinationConfig,omitempty" cloudformation:"DestinationConfig"`

	// FunctionName http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lambda-eventinvokeconfig.html#cfn-lambda-eventinvokeconfig-functionname
	FunctionName string `json:"functionName,omitempty" cloudformation:"FunctionName,Parameter"`

	// MaximumEventAgeInSeconds http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lambda-eventinvokeconfig.html#cfn-lambda-eventinvokeconfig-maximumeventageinseconds
	MaximumEventAgeInSeconds int `json:"maximumEventAgeInSeconds,omitempty" cloudformation:"MaximumEventAgeInSeconds,Parameter"`

	// MaximumRetryAttempts http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lambda-eventinvokeconfig.html#cfn-lambda-eventinvokeconfig-maximumretryattempts
	MaximumRetryAttempts int `json:"maximumRetryAttempts,omitempty" cloudformation:"MaximumRetryAttempts,Parameter"`

	// Qualifier http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lambda-eventinvokeconfig.html#cfn-lambda-eventinvokeconfig-qualifier
	Qualifier string `json:"qualifier,omitempty" cloudformation:"Qualifier,Parameter"`
}

// EventInvokeConfig_OnSuccess defines the desired state of EventInvokeConfigOnSuccess
type EventInvokeConfig_OnSuccess struct {
	// Destination http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lambda-eventinvokeconfig-destinationconfig-onsuccess.html#cfn-lambda-eventinvokeconfig-destinationconfig-onsuccess-destination
	Destination string `json:"destination,omitempty" cloudformation:"Destination,Parameter"`
}

// EventInvokeConfig_DestinationConfig defines the desired state of EventInvokeConfigDestinationConfig
type EventInvokeConfig_DestinationConfig struct {
	// OnFailure http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lambda-eventinvokeconfig-destinationconfig.html#cfn-lambda-eventinvokeconfig-destinationconfig-onfailure
	OnFailure EventInvokeConfig_OnFailure `json:"onFailure,omitempty" cloudformation:"OnFailure"`

	// OnSuccess http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lambda-eventinvokeconfig-destinationconfig.html#cfn-lambda-eventinvokeconfig-destinationconfig-onsuccess
	OnSuccess EventInvokeConfig_OnSuccess `json:"onSuccess,omitempty" cloudformation:"OnSuccess"`
}

// EventInvokeConfig_OnFailure defines the desired state of EventInvokeConfigOnFailure
type EventInvokeConfig_OnFailure struct {
	// Destination http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lambda-eventinvokeconfig-destinationconfig-onfailure.html#cfn-lambda-eventinvokeconfig-destinationconfig-onfailure-destination
	Destination string `json:"destination,omitempty" cloudformation:"Destination,Parameter"`
}

// EventInvokeConfigStatus defines the observed state of EventInvokeConfig
type EventInvokeConfigStatus struct {
	metav1alpha1.StatusMeta `json:",inline"`
}

// EventInvokeConfigOutput defines the stack outputs
type EventInvokeConfigOutput struct {
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lambda-eventinvokeconfig.html
	Ref string `json:"ref,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:resource:categories=aws;lambda
// +kubebuilder:printcolumn:JSONPath=.status.status,description="status of the stack",name=Status,priority=0,type=string
// +kubebuilder:printcolumn:JSONPath=.status.message,description="reason for the stack status",name=Message,priority=1,type=string
// +kubebuilder:printcolumn:JSONPath=.status.stackID,description="CloudFormation Stack ID",name=StackID,priority=2,type=string

// EventInvokeConfig is the Schema for the lambda EventInvokeConfig API
type EventInvokeConfig struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   EventInvokeConfigSpec   `json:"spec,omitempty"`
	Status EventInvokeConfigStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// EventInvokeConfigList contains a list of Account
type EventInvokeConfigList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`

	Items []EventInvokeConfig `json:"items"`
}

func init() {
	SchemeBuilder.Register(&EventInvokeConfig{}, &EventInvokeConfigList{})
}
