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
	"fmt"
	"reflect"
	"strings"

	metav1alpha1 "go.awsctrl.io/manager/apis/meta/v1alpha1"
	controllerutils "go.awsctrl.io/manager/controllers/utils"
	cfnencoder "go.awsctrl.io/manager/encoding/cloudformation"

	"github.com/awslabs/goformation/v4/cloudformation"
	"github.com/awslabs/goformation/v4/cloudformation/lambda"
	"k8s.io/client-go/dynamic"
)

// GetNotificationARNs is an autogenerated deepcopy function, will return notifications for stack
func (in *Function) GetNotificationARNs() []string {
	notifcations := []string{}
	for _, notifarn := range in.Spec.NotificationARNs {
		notifcations = append(notifcations, *notifarn)
	}
	return notifcations
}

// GetTemplate will return the JSON version of the CFN to use.
func (in *Function) GetTemplate(client dynamic.Interface) (string, error) {
	if client == nil {
		return "", fmt.Errorf("k8s client not loaded for template")
	}
	template := cloudformation.NewTemplate()

	template.Description = "AWS Controller - lambda.Function (ac-{TODO})"

	template.Outputs = map[string]interface{}{
		"ResourceRef": map[string]interface{}{
			"Value": cloudformation.Ref("Function"),
			"Export": map[string]interface{}{
				"Name": in.Name + "Ref",
			},
		},
		"Arn": map[string]interface{}{
			"Value":  cloudformation.GetAtt("Function", "Arn"),
			"Export": map[string]interface{}{"Name": in.Name + "Arn"},
		},
	}

	lambdaFunction := &lambda.Function{}

	if in.Spec.Timeout != lambdaFunction.Timeout {
		lambdaFunction.Timeout = in.Spec.Timeout
	}

	// TODO(christopherhein) move these to a defaulter
	lambdaFunctionKmsKeyRefItem := in.Spec.KmsKeyRef.DeepCopy()

	if lambdaFunctionKmsKeyRefItem.ObjectRef.Namespace == "" {
		lambdaFunctionKmsKeyRefItem.ObjectRef.Namespace = in.Namespace
	}

	in.Spec.KmsKeyRef = *lambdaFunctionKmsKeyRefItem
	kmsKeyArn, err := in.Spec.KmsKeyRef.String(client)
	if err != nil {
		return "", err
	}

	if kmsKeyArn != "" {
		lambdaFunction.KmsKeyArn = kmsKeyArn
	}

	if !reflect.DeepEqual(in.Spec.VpcConfig, Function_VpcConfig{}) {
		lambdaFunctionVpcConfig := lambda.Function_VpcConfig{}

		if len(in.Spec.VpcConfig.SecurityGroupRefs) > 0 {
			lambdaFunctionVpcConfigSecurityGroupRefs := []string{}

			for _, item := range in.Spec.VpcConfig.SecurityGroupRefs {
				lambdaFunctionVpcConfigSecurityGroupRefsItem := item.DeepCopy()

				if lambdaFunctionVpcConfigSecurityGroupRefsItem.ObjectRef.Namespace == "" {
					lambdaFunctionVpcConfigSecurityGroupRefsItem.ObjectRef.Namespace = in.Namespace
				}

			}

			lambdaFunctionVpcConfig.SecurityGroupIds = lambdaFunctionVpcConfigSecurityGroupRefs
		}

		if len(in.Spec.VpcConfig.SubnetRefs) > 0 {
			lambdaFunctionVpcConfigSubnetRefs := []string{}

			for _, item := range in.Spec.VpcConfig.SubnetRefs {
				lambdaFunctionVpcConfigSubnetRefsItem := item.DeepCopy()

				if lambdaFunctionVpcConfigSubnetRefsItem.ObjectRef.Namespace == "" {
					lambdaFunctionVpcConfigSubnetRefsItem.ObjectRef.Namespace = in.Namespace
				}

			}

			lambdaFunctionVpcConfig.SubnetIds = lambdaFunctionVpcConfigSubnetRefs
		}

		lambdaFunction.VpcConfig = &lambdaFunctionVpcConfig
	}

	if !reflect.DeepEqual(in.Spec.Code, Function_Code{}) {
		lambdaFunctionCode := lambda.Function_Code{}

		if in.Spec.Code.S3ObjectVersion != "" {
			lambdaFunctionCode.S3ObjectVersion = in.Spec.Code.S3ObjectVersion
		}

		if in.Spec.Code.ZipFile != "" {
			lambdaFunctionCode.ZipFile = in.Spec.Code.ZipFile
		}

		if in.Spec.Code.S3Bucket != "" {
			lambdaFunctionCode.S3Bucket = in.Spec.Code.S3Bucket
		}

		if in.Spec.Code.S3Key != "" {
			lambdaFunctionCode.S3Key = in.Spec.Code.S3Key
		}

		lambdaFunction.Code = &lambdaFunctionCode
	}

	if !reflect.DeepEqual(in.Spec.DeadLetterConfig, Function_DeadLetterConfig{}) {
		lambdaFunctionDeadLetterConfig := lambda.Function_DeadLetterConfig{}

		// TODO(christopherhein) move these to a defaulter
		lambdaFunctionDeadLetterConfigTargetRefItem := in.Spec.DeadLetterConfig.TargetRef.DeepCopy()

		if lambdaFunctionDeadLetterConfigTargetRefItem.ObjectRef.Namespace == "" {
			lambdaFunctionDeadLetterConfigTargetRefItem.ObjectRef.Namespace = in.Namespace
		}

		in.Spec.DeadLetterConfig.TargetRef = *lambdaFunctionDeadLetterConfigTargetRefItem
		targetArn, err := in.Spec.DeadLetterConfig.TargetRef.String(client)
		if err != nil {
			return "", err
		}

		if targetArn != "" {
			lambdaFunctionDeadLetterConfig.TargetArn = targetArn
		}

		lambdaFunction.DeadLetterConfig = &lambdaFunctionDeadLetterConfig
	}

	// TODO(christopherhein) move these to a defaulter
	if in.Spec.FunctionName == "" {
		lambdaFunction.FunctionName = in.Name
	}

	if in.Spec.FunctionName != "" {
		lambdaFunction.FunctionName = in.Spec.FunctionName
	}

	// TODO(christopherhein): implement tags this could be easy now that I have the mechanims of nested objects
	if in.Spec.ReservedConcurrentExecutions != lambdaFunction.ReservedConcurrentExecutions {
		lambdaFunction.ReservedConcurrentExecutions = in.Spec.ReservedConcurrentExecutions
	}

	if in.Spec.Description != "" {
		lambdaFunction.Description = in.Spec.Description
	}

	if in.Spec.Role != "" {
		lambdaFunction.Role = in.Spec.Role
	}

	if !reflect.DeepEqual(in.Spec.TracingConfig, Function_TracingConfig{}) {
		lambdaFunctionTracingConfig := lambda.Function_TracingConfig{}

		if in.Spec.TracingConfig.Mode != "" {
			lambdaFunctionTracingConfig.Mode = in.Spec.TracingConfig.Mode
		}

		lambdaFunction.TracingConfig = &lambdaFunctionTracingConfig
	}

	if in.Spec.Handler != "" {
		lambdaFunction.Handler = in.Spec.Handler
	}

	if !reflect.DeepEqual(in.Spec.Environment, Function_Environment{}) {
		lambdaFunctionEnvironment := lambda.Function_Environment{}

		if !reflect.DeepEqual(in.Spec.Environment.Variables, map[string]string{}) {
			lambdaFunctionEnvironment.Variables = in.Spec.Environment.Variables
		}

		lambdaFunction.Environment = &lambdaFunctionEnvironment
	}

	if len(in.Spec.Layers) > 0 {
		lambdaFunction.Layers = in.Spec.Layers
	}

	if in.Spec.MemorySize != lambdaFunction.MemorySize {
		lambdaFunction.MemorySize = in.Spec.MemorySize
	}

	if in.Spec.Runtime != "" {
		lambdaFunction.Runtime = in.Spec.Runtime
	}

	template.Resources = map[string]cloudformation.Resource{
		"Function": lambdaFunction,
	}

	// json, err := template.JSONWithOptions(&intrinsics.ProcessorOptions{NoEvaluateConditions: true})
	json, err := template.JSON()
	if err != nil {
		return "", err
	}

	return string(json), nil
}

// GetStackID will return stackID
func (in *Function) GetStackID() string {
	return in.Status.StackID
}

// GenerateStackName will generate a StackName
func (in *Function) GenerateStackName() string {
	return strings.Join([]string{"lambda", "function", in.GetName(), in.GetNamespace()}, "-")
}

// GetStackName will return stackName
func (in *Function) GetStackName() string {
	return in.Spec.StackName
}

// GetTemplateVersionLabel will return the stack template version
func (in *Function) GetTemplateVersionLabel() (value string, ok bool) {
	value, ok = in.Labels[controllerutils.StackTemplateVersionLabel]
	return
}

// GetParameters will return CFN Parameters
func (in *Function) GetParameters() map[string]string {
	params := map[string]string{}
	cfnencoder.MarshalTypes(params, in.Spec, "Parameter")
	return params
}

// GetCloudFormationMeta will return CFN meta object
func (in *Function) GetCloudFormationMeta() metav1alpha1.CloudFormationMeta {
	return in.Spec.CloudFormationMeta
}

// GetStatus will return the CFN Status
func (in *Function) GetStatus() metav1alpha1.ConditionStatus {
	return in.Status.Status
}

// SetStackID will put a stackID
func (in *Function) SetStackID(input string) {
	in.Status.StackID = input
	return
}

// SetStackName will return stackName
func (in *Function) SetStackName(input string) {
	in.Spec.StackName = input
	return
}

// SetTemplateVersionLabel will set the template version label
func (in *Function) SetTemplateVersionLabel() {
	if len(in.Labels) == 0 {
		in.Labels = map[string]string{}
	}

	in.Labels[controllerutils.StackTemplateVersionLabel] = controllerutils.ComputeHash(in.Spec)
}

// TemplateVersionChanged will return bool if template has changed
func (in *Function) TemplateVersionChanged() bool {
	// Ignore bool since it will still record changed
	label, _ := in.GetTemplateVersionLabel()
	return label != controllerutils.ComputeHash(in.Spec)
}

// SetStatus will set status for object
func (in *Function) SetStatus(status *metav1alpha1.StatusMeta) {
	in.Status.StatusMeta = *status
}
