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

// AliasSpec defines the desired state of Alias
type AliasSpec struct {
	metav1alpha1.CloudFormationMeta `json:",inline"`

	// FunctionVersion http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lambda-alias.html#cfn-lambda-alias-functionversion
	FunctionVersion string `json:"functionVersion" cloudformation:"FunctionVersion,Parameter"`

	// Name http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lambda-alias.html#cfn-lambda-alias-name
	Name string `json:"name" cloudformation:"Name,Parameter"`

	// ProvisionedConcurrencyConfig http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lambda-alias.html#cfn-lambda-alias-provisionedconcurrencyconfig
	ProvisionedConcurrencyConfig Alias_ProvisionedConcurrencyConfiguration `json:"provisionedConcurrencyConfig,omitempty" cloudformation:"ProvisionedConcurrencyConfig"`

	// RoutingConfig http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lambda-alias.html#cfn-lambda-alias-routingconfig
	RoutingConfig Alias_AliasRoutingConfiguration `json:"routingConfig,omitempty" cloudformation:"RoutingConfig"`

	// Description http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lambda-alias.html#cfn-lambda-alias-description
	Description string `json:"description,omitempty" cloudformation:"Description,Parameter"`

	// FunctionName http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lambda-alias.html#cfn-lambda-alias-functionname
	FunctionName string `json:"functionName" cloudformation:"FunctionName,Parameter"`
}

// Alias_ProvisionedConcurrencyConfiguration defines the desired state of AliasProvisionedConcurrencyConfiguration
type Alias_ProvisionedConcurrencyConfiguration struct {
	// ProvisionedConcurrentExecutions http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lambda-alias-provisionedconcurrencyconfiguration.html#cfn-lambda-alias-provisionedconcurrencyconfiguration-provisionedconcurrentexecutions
	ProvisionedConcurrentExecutions int `json:"provisionedConcurrentExecutions" cloudformation:"ProvisionedConcurrentExecutions,Parameter"`
}

// Alias_AliasRoutingConfiguration defines the desired state of AliasAliasRoutingConfiguration
type Alias_AliasRoutingConfiguration struct {
	// AdditionalVersionWeights http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lambda-alias-aliasroutingconfiguration.html#cfn-lambda-alias-aliasroutingconfiguration-additionalversionweights
	AdditionalVersionWeights []Alias_VersionWeight `json:"additionalVersionWeights" cloudformation:"AdditionalVersionWeights"`
}

// Alias_VersionWeight defines the desired state of AliasVersionWeight
type Alias_VersionWeight struct {
	// FunctionWeight http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lambda-alias-versionweight.html#cfn-lambda-alias-versionweight-functionweight
	FunctionWeight string `json:"functionWeight" cloudformation:"FunctionWeight,Parameter"`

	// FunctionVersion http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lambda-alias-versionweight.html#cfn-lambda-alias-versionweight-functionversion
	FunctionVersion string `json:"functionVersion" cloudformation:"FunctionVersion,Parameter"`
}

// AliasStatus defines the observed state of Alias
type AliasStatus struct {
	metav1alpha1.StatusMeta `json:",inline"`
}

// AliasOutput defines the stack outputs
type AliasOutput struct {
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lambda-alias.html
	Ref string `json:"ref,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:resource:categories=aws;lambda
// +kubebuilder:printcolumn:JSONPath=.status.status,description="status of the stack",name=Status,priority=0,type=string
// +kubebuilder:printcolumn:JSONPath=.status.message,description="reason for the stack status",name=Message,priority=1,type=string
// +kubebuilder:printcolumn:JSONPath=.status.stackID,description="CloudFormation Stack ID",name=StackID,priority=2,type=string

// Alias is the Schema for the lambda Alias API
type Alias struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   AliasSpec   `json:"spec,omitempty"`
	Status AliasStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// AliasList contains a list of Account
type AliasList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`

	Items []Alias `json:"items"`
}

func init() {
	SchemeBuilder.Register(&Alias{}, &AliasList{})
}
