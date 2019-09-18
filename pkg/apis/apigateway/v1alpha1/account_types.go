package v1alpha1

import (
	metav1alpha1 "awsctrl.io/pkg/apis/meta/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// AccountSpec defines the desired state of Account
type AccountSpec struct {
	metav1alpha1.CloudFormationMeta `json:",inline"`

	// CloudWatchRoleRef http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigateway-account.html#cfn-apigateway-account-cloudwatchrolearn
	CloudWatchRoleRef string `json:"cloudWatchRoleRef,omitempty" cloudformation:"CloudWatchRoleArn,Parameter"`
}

// AccountStatus defines the observed state of Account
type AccountStatus struct {
	metav1alpha1.StatusMeta `json:",inline"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:subresource:status
// +kubebuilder:resource

// Account is the Schema for the accounts API
// +k8s:openapi-gen=true
type Account struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec    AccountSpec       `json:"spec,omitempty"`
	Status  AccountStatus     `json:"status,omitempty"`
	Outputs map[string]string `json:"outputs,omitempty"`
}

// GetTemplate will return the JSON version of the CFN to use.
func (r Account) GetTemplate() string {
	return "{\"AWSTemplateFormatVersion\":\"2010-09-09\",\"Description\":\"AWS Service Operator - AWS::ApiGateway::Account (aso-0owrq417x)\",\"Parameters\":{\"CloudWatchRoleArn\":{\"Type\":\"String\",\"Description\":\"http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigateway-account.html#cfn-apigateway-account-cloudwatchrolearn\"}},\"Resources\":{\"Resource\":{\"Type\":\"AWS::ApiGateway::Account\",\"Properties\":{\"CloudWatchRoleArn\":{\"Ref\":\"CloudWatchRoleArn\"}}}}}"
}

// GetStackID will return stackID
func (r Account) GetStackID() string {
	return r.Status.StatusMeta.StackID
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AccountList contains a list of Account
type AccountList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Account `json:"items"`
}

func init() {
	SchemeBuilder.Register(&Account{}, &AccountList{})
}
