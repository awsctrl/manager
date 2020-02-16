/*
Copyright © 2019 AWS Controller authors

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
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// ConfigSpec defines the desired state of Config
type ConfigSpec struct {
	// ClusterName is the identity for categorizing all resources under
	ClusterName string `json:"clusterName"`

	// AWS contains all the AWS Account specific details
	AWS ConfigAWS `json:"aws"`

	// Resources is a whitelist of Services and Resources to enable
	// +optional
	Resources []string `json:"resources,omitempty"`

	// Sync is the config for the syncing parameters
	// +optional
	Sync ConfigSync `json:"sync,omitempty"`

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
	// +optional
	AccountID string `json:"accountID,omitempty"`
}

// ConfigSync contains the sync configurations
type ConfigSync struct {
	// Enabled turns on the syncing
	Enabled bool `json:"enabled"`
}

// ConfigStatus defines the observed state of Config
type ConfigStatus struct {
	// Conditions defines the stages the deployment is in
	// +optional
	Conditions []ConfigStatusCondition `json:"conditions,omitempty"`
}

// ConfigStatusCondition defines the conditions the operator is in
type ConfigStatusCondition struct {
	// type of cluster condition, values in (\"Ready\")
	Type ConfigConditionType `json:"type"`

	// Status of the condition, one of (\"True\", \"False\", \"Unknown\")
	Status corev1.ConditionStatus `json:"status"`

	// LastTransitionTime
	LastTransitionTime metav1.Time `json:"lastTransitionTime"`

	// Message defines the human readable message for the condition
	// +optional
	Message string `json:"message,omitempty"`

	// Reason defines the Machine readable status
	// +optional
	Reason string `json:"reason,omitempty"`
}

// ConfigConditionType defines type for config condition type.
type ConfigConditionType string

const (
	// ConfigConditionReady represents the readiness of the AWS Controller.
	ConfigConditionReady ConfigConditionType = "Ready"

	// ConfigConditionPendingAWSConfiguration represents the controller not
	// being configured with AWS
	ConfigConditionPendingAWSConfiguration ConfigConditionType = "PendingAWSConfiguration"
)

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:resource:categories=aws;self
// +kubebuilder:printcolumn:JSONPath=.status.conditions[?(@.status == "True")].type,description="status of the stack",name=Status,priority=0,type=string

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
