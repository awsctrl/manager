/*
Copyright © 2019 AWS Controller author

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package v1alpha1

import (
	goversion "go.hein.dev/go-version"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// ConfigSpec defines the desired state of Config
type ConfigSpec struct {
	// ClusterName is the identity for categorizing all resources under
	ClusterName string `json:"clusterName"`

	// AWS contains all the AWS Account specific details
	AWS ConfigAWS `json:"aws"`

	// Resources is a whitelist of Services and Resources to enable
	Resources []string `json:"resources"`

	// Version stores the operator version information
	// +optional
	Version *goversion.Info `json:"version,omitempty"`
}

// ConfigAWS contains all the AWS specific details
type ConfigAWS struct {
	// DefaultRegion configures which region a stack should be provisioned into by default
	DefaultRegion string `json:"defaultRegion"`

	// SupportedRegions configures which regions CRDs can deploy into
	SupportedRegions []string `json:"supportedRegions"`

	// AccountID defines the account which each resource is connected to
	AccountID string `json:"accountID"`

	// Queue will set up the params for the queue system
	// +optional
	Queue ConfigQueue `json:"queue,omitempty"`
}

// ConfigQueue holds all the configurations for the SQS & SNS configs
type ConfigQueue struct {
	// Region is where SQS and SNS should be provisioned
	// +optional
	Region string `json:"region"`

	// Name is the name of the SQS queue to be used
	// +optional
	Name string `json:"name"`

	// TopicARN is the name ARN for cloudformation to talk to.
	// +optional
	TopicARN string `json:"topicARN,omitempty"`

	// SubARN is the name ARN for sqs to talk to.
	// +optional
	SubARN string `json:"subARN,omitempty"`

	// QueueARN defines a preconfigured queue vs creating them on boot
	// +optional
	QueueARN string `json:"queueARN"`

	// QueueURL defines the URL for the Queue
	// +optional
	QueueURL string `json:"queueURL"`
}

// ConfigStatus defines the observed state of Config
type ConfigStatus struct {
	// Conditions defines the stages the deployment is in
	// +optional
	Conditions []ConfigStatusConditions `json:"conditions,omitempty"`
}

// ConfigStatusConditions defines the conditions the operator is in
type ConfigStatusConditions struct {
	// Message defines the human readable message for the condition
	// +optional
	Message string `json:"message,omitempty"`

	// Reason defines the Machine readable status
	// +optional
	Reason ReasonCondition `json:"reason,omitempty"`
}

// ReasonCondition defines the state changes
type ReasonCondition string

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:resource:categories=aws;all;self

// Config is the Schema for the Configs API
type Config struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ConfigSpec   `json:"spec,omitempty"`
	Status ConfigStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ConfigList contains a list of Config
type ConfigList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Config `json:"items"`
}

func init() {
	SchemeBuilder.Register(&Config{}, &ConfigList{})
}
