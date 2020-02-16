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

// FunctionSpec defines the desired state of Function
type FunctionSpec struct {
	metav1alpha1.CloudFormationMeta `json:",inline"`

	// Code http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lambda-function.html#cfn-lambda-function-code
	Code Function_Code `json:"code,omitempty" cloudformation:"Code"`

	// DeadLetterConfig http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lambda-function.html#cfn-lambda-function-deadletterconfig
	DeadLetterConfig Function_DeadLetterConfig `json:"deadLetterConfig,omitempty" cloudformation:"DeadLetterConfig"`

	// Description http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lambda-function.html#cfn-lambda-function-description
	Description string `json:"description,omitempty" cloudformation:"Description,Parameter"`

	// Environment http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lambda-function.html#cfn-lambda-function-environment
	Environment Function_Environment `json:"environment,omitempty" cloudformation:"Environment"`

	// FunctionName http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lambda-function.html#cfn-lambda-function-functionname
	FunctionName string `json:"functionName,omitempty" cloudformation:"FunctionName,Parameter"`

	// Handler http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lambda-function.html#cfn-lambda-function-handler
	Handler string `json:"handler,omitempty" cloudformation:"Handler,Parameter"`

	// KmsKeyRef http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lambda-function.html#cfn-lambda-function-kmskeyarn
	KmsKeyRef metav1alpha1.ObjectReference `json:"kmsKeyRef,omitempty" cloudformation:"KmsKeyArn,Parameter"`

	// Layers http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lambda-function.html#cfn-lambda-function-layers
	Layers []string `json:"layers,omitempty" cloudformation:"Layers"`

	// MemorySize http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lambda-function.html#cfn-lambda-function-memorysize
	MemorySize int `json:"memorySize,omitempty" cloudformation:"MemorySize,Parameter"`

	// ReservedConcurrentExecutions http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lambda-function.html#cfn-lambda-function-reservedconcurrentexecutions
	ReservedConcurrentExecutions int `json:"reservedConcurrentExecutions,omitempty" cloudformation:"ReservedConcurrentExecutions,Parameter"`

	// Role http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lambda-function.html#cfn-lambda-function-role
	Role string `json:"role,omitempty" cloudformation:"Role,Parameter"`

	// Runtime http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lambda-function.html#cfn-lambda-function-runtime
	Runtime string `json:"runtime,omitempty" cloudformation:"Runtime,Parameter"`

	// Timeout http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lambda-function.html#cfn-lambda-function-timeout
	Timeout int `json:"timeout,omitempty" cloudformation:"Timeout,Parameter"`

	// TracingConfig http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lambda-function.html#cfn-lambda-function-tracingconfig
	TracingConfig Function_TracingConfig `json:"tracingConfig,omitempty" cloudformation:"TracingConfig"`

	// VpcConfig http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lambda-function.html#cfn-lambda-function-vpcconfig
	VpcConfig Function_VpcConfig `json:"vpcConfig,omitempty" cloudformation:"VpcConfig"`
}

// Function_Code defines the desired state of FunctionCode
type Function_Code struct {
	// S3Bucket http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lambda-function-code.html#cfn-lambda-function-code-s3bucket
	S3Bucket string `json:"s3Bucket,omitempty" cloudformation:"S3Bucket,Parameter"`

	// S3Key http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lambda-function-code.html#cfn-lambda-function-code-s3key
	S3Key string `json:"s3Key,omitempty" cloudformation:"S3Key,Parameter"`

	// S3ObjectVersion http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lambda-function-code.html#cfn-lambda-function-code-s3objectversion
	S3ObjectVersion string `json:"s3ObjectVersion,omitempty" cloudformation:"S3ObjectVersion,Parameter"`

	// ZipFile http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lambda-function-code.html#cfn-lambda-function-code-zipfile
	ZipFile string `json:"zipFile,omitempty" cloudformation:"ZipFile,Parameter"`
}

// Function_DeadLetterConfig defines the desired state of FunctionDeadLetterConfig
type Function_DeadLetterConfig struct {
	// TargetRef http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lambda-function-deadletterconfig.html#cfn-lambda-function-deadletterconfig-targetarn
	TargetRef metav1alpha1.ObjectReference `json:"targetRef,omitempty" cloudformation:"TargetArn,Parameter"`
}

// Function_Environment defines the desired state of FunctionEnvironment
type Function_Environment struct {
	// Variables http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lambda-function-environment.html#cfn-lambda-function-environment-variables
	Variables map[string]string `json:"variables,omitempty" cloudformation:"Variables"`
}

// Function_TracingConfig defines the desired state of FunctionTracingConfig
type Function_TracingConfig struct {
	// Mode http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lambda-function-tracingconfig.html#cfn-lambda-function-tracingconfig-mode
	Mode string `json:"mode,omitempty" cloudformation:"Mode,Parameter"`
}

// Function_VpcConfig defines the desired state of FunctionVpcConfig
type Function_VpcConfig struct {
	// SecurityGroupRefs http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lambda-function-vpcconfig.html#cfn-lambda-function-vpcconfig-securitygroupids
	SecurityGroupRefs []metav1alpha1.ObjectReference `json:"securityGroupRefs,omitempty" cloudformation:"SecurityGroupIds"`

	// SubnetRefs http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lambda-function-vpcconfig.html#cfn-lambda-function-vpcconfig-subnetids
	SubnetRefs []metav1alpha1.ObjectReference `json:"subnetRefs,omitempty" cloudformation:"SubnetIds"`
}

// FunctionStatus defines the observed state of Function
type FunctionStatus struct {
	metav1alpha1.StatusMeta `json:",inline"`
}

// FunctionOutput defines the stack outputs
type FunctionOutput struct {
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lambda-function.html
	Ref string `json:"ref,omitempty"`

	// Arn defines the Arn
	Arn string `json:"arn,omitempty" cloudformation:"Arn,Output"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:resource:categories=aws;lambda
// +kubebuilder:printcolumn:JSONPath=.status.status,description="status of the stack",name=Status,priority=0,type=string
// +kubebuilder:printcolumn:JSONPath=.status.message,description="reason for the stack status",name=Message,priority=1,type=string
// +kubebuilder:printcolumn:JSONPath=.status.stackID,description="CloudFormation Stack ID",name=StackID,priority=2,type=string

// Function is the Schema for the lambda Function API
type Function struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   FunctionSpec   `json:"spec,omitempty"`
	Status FunctionStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// FunctionList contains a list of Account
type FunctionList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`

	Items []Function `json:"items"`
}

func init() {
	SchemeBuilder.Register(&Function{}, &FunctionList{})
}
