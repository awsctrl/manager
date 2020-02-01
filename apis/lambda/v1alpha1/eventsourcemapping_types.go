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

// EventSourceMappingSpec defines the desired state of EventSourceMapping
type EventSourceMappingSpec struct {
	metav1alpha1.CloudFormationMeta `json:",inline"`

	// DestinationConfig http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lambda-eventsourcemapping.html#cfn-lambda-eventsourcemapping-destinationconfig
	DestinationConfig EventSourceMapping_DestinationConfig `json:"destinationConfig,omitempty" cloudformation:"DestinationConfig"`

	// StartingPosition http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lambda-eventsourcemapping.html#cfn-lambda-eventsourcemapping-startingposition
	StartingPosition string `json:"startingPosition,omitempty" cloudformation:"StartingPosition,Parameter"`

	// BatchSize http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lambda-eventsourcemapping.html#cfn-lambda-eventsourcemapping-batchsize
	BatchSize int `json:"batchSize,omitempty" cloudformation:"BatchSize,Parameter"`

	// Enabled http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lambda-eventsourcemapping.html#cfn-lambda-eventsourcemapping-enabled
	Enabled bool `json:"enabled,omitempty" cloudformation:"Enabled,Parameter"`

	// EventSourceRef http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lambda-eventsourcemapping.html#cfn-lambda-eventsourcemapping-eventsourcearn
	EventSourceRef metav1alpha1.ObjectReference `json:"eventSourceRef,omitempty" cloudformation:"EventSourceArn,Parameter"`

	// FunctionName http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lambda-eventsourcemapping.html#cfn-lambda-eventsourcemapping-functionname
	FunctionName string `json:"functionName,omitempty" cloudformation:"FunctionName,Parameter"`

	// BisectBatchOnFunctionError http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lambda-eventsourcemapping.html#cfn-lambda-eventsourcemapping-bisectbatchonfunctionerror
	BisectBatchOnFunctionError bool `json:"bisectBatchOnFunctionError,omitempty" cloudformation:"BisectBatchOnFunctionError,Parameter"`

	// MaximumBatchingWindowInSeconds http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lambda-eventsourcemapping.html#cfn-lambda-eventsourcemapping-maximumbatchingwindowinseconds
	MaximumBatchingWindowInSeconds int `json:"maximumBatchingWindowInSeconds,omitempty" cloudformation:"MaximumBatchingWindowInSeconds,Parameter"`

	// MaximumRecordAgeInSeconds http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lambda-eventsourcemapping.html#cfn-lambda-eventsourcemapping-maximumrecordageinseconds
	MaximumRecordAgeInSeconds int `json:"maximumRecordAgeInSeconds,omitempty" cloudformation:"MaximumRecordAgeInSeconds,Parameter"`

	// MaximumRetryAttempts http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lambda-eventsourcemapping.html#cfn-lambda-eventsourcemapping-maximumretryattempts
	MaximumRetryAttempts int `json:"maximumRetryAttempts,omitempty" cloudformation:"MaximumRetryAttempts,Parameter"`

	// ParallelizationFactor http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lambda-eventsourcemapping.html#cfn-lambda-eventsourcemapping-parallelizationfactor
	ParallelizationFactor int `json:"parallelizationFactor,omitempty" cloudformation:"ParallelizationFactor,Parameter"`
}

// EventSourceMapping_OnFailure defines the desired state of EventSourceMappingOnFailure
type EventSourceMapping_OnFailure struct {
	// Destination http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lambda-eventsourcemapping-onfailure.html#cfn-lambda-eventsourcemapping-onfailure-destination
	Destination string `json:"destination,omitempty" cloudformation:"Destination,Parameter"`
}

// EventSourceMapping_DestinationConfig defines the desired state of EventSourceMappingDestinationConfig
type EventSourceMapping_DestinationConfig struct {
	// OnFailure http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lambda-eventsourcemapping-destinationconfig.html#cfn-lambda-eventsourcemapping-destinationconfig-onfailure
	OnFailure EventSourceMapping_OnFailure `json:"onFailure,omitempty" cloudformation:"OnFailure"`
}

// EventSourceMappingStatus defines the observed state of EventSourceMapping
type EventSourceMappingStatus struct {
	metav1alpha1.StatusMeta `json:",inline"`
}

// EventSourceMappingOutput defines the stack outputs
type EventSourceMappingOutput struct {
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lambda-eventsourcemapping.html
	Ref string `json:"ref,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:resource:categories=aws;lambda
// +kubebuilder:printcolumn:JSONPath=.status.status,description="status of the stack",name=Status,priority=0,type=string
// +kubebuilder:printcolumn:JSONPath=.status.message,description="reason for the stack status",name=Message,priority=1,type=string
// +kubebuilder:printcolumn:JSONPath=.status.stackID,description="CloudFormation Stack ID",name=StackID,priority=2,type=string

// EventSourceMapping is the Schema for the lambda EventSourceMapping API
type EventSourceMapping struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   EventSourceMappingSpec   `json:"spec,omitempty"`
	Status EventSourceMappingStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// EventSourceMappingList contains a list of Account
type EventSourceMappingList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`

	Items []EventSourceMapping `json:"items"`
}

func init() {
	SchemeBuilder.Register(&EventSourceMapping{}, &EventSourceMappingList{})
}
