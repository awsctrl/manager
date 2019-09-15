package v1alpha1

import (
	metav1alpha1 "awsctrl.io/pkg/apis/meta/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// StackSpec defines the desired state of Stack
type StackSpec struct {
	metav1alpha1.CloudFormationMeta `json:",inline"`

	// Capabilities In some cases, you must explicity acknowledge that your stack
	// template contains certain capabilities in order for AWS CloudFormation to
	// create the stack.
	// +optional
	Capabilities []*string `json:"capabilities,omitempty"`

	// Parameters A list of Parameter structures that specify input parameters for
	// the stack. For more information, see the Parameter data type.
	Parameters map[string]string `json:"parameters"`

	// ClientRequestToken A unique identifier for this CreateStack request.
	// Specify this token if you plan to retry requests so that AWS CloudFormation
	// knows that you're not attempting to create a stack with the same name. You
	// might retry CreateStack requests to ensure that AWS CloudFormation
	// successfully received them.
	// +optional
	ClientRequestToken string `json:"clientRequestToken,omitempty"`

	// TemplateBody Structure containing the template body with a minimum length
	// of 1 byte and a maximum length of 51,200 bytes.
	TemplateBody string `json:"templateBody"`
}

// StackStatus defines the observed state of Stack
type StackStatus struct {
	metav1alpha1.StatusMeta `json:",inline"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:subresource:status
// +kubebuilder:resource

// Stack is the Schema for the stacks API
// +k8s:openapi-gen=true
type Stack struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   StackSpec   `json:"spec,omitempty"`
	Status StackStatus `json:"status,omitempty"`

	Outputs map[string]string `json:"outputs,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// StackList contains a list of Stack
type StackList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Stack `json:"items"`
}

func init() {
	SchemeBuilder.Register(&Stack{}, &StackList{})
}
