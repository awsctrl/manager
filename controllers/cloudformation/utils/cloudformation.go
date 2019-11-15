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

package utils

import (
	"strings"

	cloudformationv1alpha1 "awsctrl.io/apis/cloudformation/v1alpha1"
	"awsctrl.io/aws"
	controllerutils "awsctrl.io/controllers/utils"

	awssdk "github.com/aws/aws-sdk-go/aws"
	cfn "github.com/aws/aws-sdk-go/service/cloudformation"
)

var (
	StackDeletionFinalizerName = "stack.cloudformation.awsctrl.io/deletion"
)

// TemplateVersionChange checks if the template has changed and returns true if it has
func TemplateVersionChanged(instance *cloudformationv1alpha1.Stack) bool {
	return instance.Labels[controllerutils.StackTemplateVersionLabel] != controllerutils.ComputeHash(instance.Spec)
}

// Name generates the name for the CFN Templates
func Name(name, namespace string) string {
	NameArr := []string{name, namespace}
	return strings.Join(NameArr, "-")
}

// CreateCFNStack will create the cloudformation stack
func CreateCFNStack(awsclient aws.AWS, instance *cloudformationv1alpha1.Stack) (*cfn.CreateStackOutput, error) {
	region := instance.Spec.Region

	input := &cfn.CreateStackInput{}
	createStackInputs(instance, awsclient.GetNotificationARN(), input)

	return awsclient.GetClient(region).CreateStack(input)
}

// UpdateCFNStack will update an existing cloudformation stack
func UpdateCFNStack(awsclient aws.AWS, instance *cloudformationv1alpha1.Stack) error {
	if instance.Status.ObservedGeneration != instance.Generation {
		input := &cfn.UpdateStackInput{}

		updateStackInputs(instance, awsclient.GetNotificationARN(), input)
		region := instance.Spec.Region

		_, err := awsclient.GetClient(region).UpdateStack(input)
		return err
	}
	return nil
}

// DeleteCFNStack will delete an existing cloudformation stack
func DeleteCFNStack(awsclient aws.AWS, instance *cloudformationv1alpha1.Stack) error {
	region := instance.Spec.Region

	input := &cfn.DeleteStackInput{}
	input.SetStackName(Name(instance.ObjectMeta.Name, instance.ObjectMeta.Namespace))

	_, err := awsclient.GetClient(region).DeleteStack(input)
	return err
}

// DescribeCFNStacks will describe all stack with name
func DescribeCFNStacks(awsclient aws.AWS, instance *cloudformationv1alpha1.Stack) (*cfn.DescribeStacksOutput, error) {
	region := instance.Spec.Region

	input := &cfn.DescribeStacksInput{}
	input.SetStackName(instance.Status.StackID)

	return awsclient.GetClient(region).DescribeStacks(input)
}

// ======= Private functions

func createStackInputs(stack *cloudformationv1alpha1.Stack, notificationARN string, input *cfn.CreateStackInput) {
	input.SetCapabilities(stack.Spec.Capabilities)

	parameters := []*cfn.Parameter{}
	for key, value := range stack.Spec.Parameters {
		param := &cfn.Parameter{}
		param.SetParameterKey(key)
		param.SetParameterValue(value)
		parameters = append(parameters, param)
	}
	input.SetParameters(parameters)

	input.SetClientRequestToken(stack.Spec.ClientRequestToken)
	input.SetEnableTerminationProtection(stack.Spec.TerminationProtection)

	notificationARNs := []*string{}
	for _, notificationARN := range stack.Spec.NotificationARNs {
		notificationARNs = append(notificationARNs, notificationARN)
	}
	if notificationARN != "" {
		notificationARNs = append(notificationARNs, awssdk.String(notificationARN))
	}
	input.SetNotificationARNs(notificationARNs)

	onFailure := stack.Spec.OnFailure
	if onFailure == "" {
		onFailure = "DELETE"
	}
	input.SetOnFailure(onFailure)

	// TODO: Comeback and make this more thought out with cluster name and a hash
	input.SetStackName(Name(stack.ObjectMeta.Name, stack.ObjectMeta.Namespace))
	input.SetTemplateBody(stack.Spec.TemplateBody)

	tags := []*cfn.Tag{}
	for key, value := range stack.Spec.Tags {
		tag := &cfn.Tag{}
		tag.SetKey(key)
		tag.SetValue(value)
		tags = append(tags, tag)
	}
	input.SetTags(tags)
}

func updateStackInputs(stack *cloudformationv1alpha1.Stack, notificationARN string, input *cfn.UpdateStackInput) {
	input.SetCapabilities(stack.Spec.Capabilities)

	parameters := []*cfn.Parameter{}
	for key, value := range stack.Spec.Parameters {
		param := &cfn.Parameter{}
		param.SetParameterKey(key)
		param.SetParameterValue(value)
		parameters = append(parameters, param)
	}
	input.SetParameters(parameters)

	// ClientRequestToken I expected to abe a stack specific token which would
	// help stop other stacks from using the same name but it doesn't seem to
	// work like that
	// input.SetClientRequestToken(stack.Spec.ClientRequestToken)

	notificationARNs := []*string{}
	for _, notificationARN := range stack.Spec.NotificationARNs {
		notificationARNs = append(notificationARNs, notificationARN)
	}
	if notificationARN != "" {
		notificationARNs = append(notificationARNs, awssdk.String(notificationARN))
	}
	input.SetNotificationARNs(notificationARNs)

	onFailure := stack.Spec.CloudFormationMeta.OnFailure
	if onFailure == "" {
		onFailure = "DELETE"
	}

	// TODO: Comeback and make this more thought out with cluster name and a hash
	input.SetStackName(Name(stack.ObjectMeta.Name, stack.ObjectMeta.Namespace))
	input.SetTemplateBody(stack.Spec.TemplateBody)

	tags := []*cfn.Tag{}
	for key, value := range stack.Spec.Tags {
		tag := &cfn.Tag{}
		tag.SetKey(key)
		tag.SetValue(value)
		tags = append(tags, tag)
	}
	input.SetTags(tags)
}
