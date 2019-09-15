package stack

import (
	"context"
	"fmt"
	"strings"

	cloudformationv1alpha1 "awsctrl.io/pkg/apis/cloudformation/v1alpha1"
	metav1alpha1 "awsctrl.io/pkg/apis/meta/v1alpha1"
	"awsctrl.io/pkg/controllerutils"

	awsclient "github.com/aws/aws-sdk-go/aws"
	cfn "github.com/aws/aws-sdk-go/service/cloudformation"
	"github.com/iancoleman/strcase"
)

func namer(name, namespace string) string {
	namerArr := []string{name, namespace}
	return strings.Join(namerArr, "-")
}

// createCFNStack will create the cloudformation stack
func (r *ReconcileStack) createCFNStack(ctx context.Context, instance *cloudformationv1alpha1.Stack) error {
	instanceCopy := instance.DeepCopy()
	region := instance.Spec.Region

	input := &cfn.CreateStackInput{}
	createStackInputs(instanceCopy, r.aws.GetNotificationARN(), input)

	output, err := r.aws.GetClient(region).CreateStack(input)
	if err != nil {
		return err
	}

	status := cloudformationv1alpha1.StackStatus{
		StatusMeta: metav1alpha1.StatusMeta{
			Status:  metav1alpha1.CreateInProgressStatus,
			StackID: string(*output.StackId),
			// This should be done with a custom version tool which onmy checks .Spec
			ObservedGeneration: instanceCopy.Generation,
		},
	}
	instanceCopy.Status = status

	return r.Status().Update(ctx, instanceCopy)
}

func (r *ReconcileStack) updateCFNStack(ctx context.Context, instance *cloudformationv1alpha1.Stack) error {
	if instance.Status.ObservedGeneration != instance.Generation {
		instanceCopy := instance.DeepCopy()

		input := &cfn.UpdateStackInput{}

		updateStackInputs(instanceCopy, r.aws.GetNotificationARN(), input)
		region := instanceCopy.Spec.Region

		_, err := r.aws.GetClient(region).UpdateStack(input)
		if err != nil {
			return err
		}

		instanceCopy.Status.Status = metav1alpha1.UpdateInProgressStatus
		instanceCopy.Status.ObservedGeneration = instanceCopy.Generation

		return r.Status().Update(ctx, instanceCopy)
	}
	return nil
}

func (r *ReconcileStack) deleteCFNStack(ctx context.Context, instance *cloudformationv1alpha1.Stack) error {
	instanceCopy := instance.DeepCopy()
	region := instanceCopy.Spec.Region

	input := &cfn.DeleteStackInput{}
	input.SetStackName(namer(instanceCopy.ObjectMeta.Name, instanceCopy.ObjectMeta.Namespace))

	_, err := r.aws.GetClient(region).DeleteStack(input)
	if err != nil {
		return err
	}

	deleteWaiter := &cfn.DescribeStacksInput{}
	deleteWaiter.SetStackName(instanceCopy.Status.StackID)

	err = r.aws.GetClient(region).WaitUntilStackDeleteComplete(deleteWaiter)
	if err != nil {
		return err
	}

	instanceCopy.ObjectMeta = controllerutils.RemoveFinalizer(instanceCopy.ObjectMeta, stackDeletionFinalizerName)

	return r.Update(ctx, instanceCopy)
}

func (r *ReconcileStack) describeCFNStackStatus(ctx context.Context, instance *cloudformationv1alpha1.Stack) error {
	instanceCopy := instance.DeepCopy()
	region := instanceCopy.Spec.Region

	input := &cfn.DescribeStacksInput{}
	input.SetStackName(instanceCopy.Status.StackID)

	outputs, err := r.aws.GetClient(region).DescribeStacks(input)
	if err != nil {
		return err
	}

	if len(outputs.Stacks) == 0 {
		return fmt.Errorf("could not find stack with name '%s'", instanceCopy.Name)
	}

	outputsMap := map[string]string{}
	for _, output := range outputs.Stacks[0].Outputs {
		outputsMap[string(*output.OutputKey)] = string(*output.OutputValue)
	}
	instanceCopy.Outputs = outputsMap

	if err = r.Client.Update(ctx, instanceCopy); err != nil {
		return err
	}

	statusCopy := instanceCopy.DeepCopy()

	status := cloudformationv1alpha1.StackStatus{
		StatusMeta: metav1alpha1.StatusMeta{
			// Figure out proper way to get the stack status and use my enum
			Status:             strcase.ToCamel(strings.ToLower(string(*outputs.Stacks[0].StackStatus))),
			Message:            outputs.Stacks[0].StackStatusReason,
			StackID:            string(*outputs.Stacks[0].StackId),
			ObservedGeneration: statusCopy.Generation,
		},
	}
	statusCopy.Status = status

	return r.Status().Update(ctx, statusCopy)
}

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
		notificationARNs = append(notificationARNs, awsclient.String(notificationARN))
	}
	input.SetNotificationARNs(notificationARNs)

	onFailure := stack.Spec.OnFailure
	if onFailure == "" {
		onFailure = "DELETE"
	}
	input.SetOnFailure(onFailure)

	// TODO: Comeback and make this more thought out with cluster name and a hash
	input.SetStackName(namer(stack.ObjectMeta.Name, stack.ObjectMeta.Namespace))
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
		notificationARNs = append(notificationARNs, awsclient.String(notificationARN))
	}
	input.SetNotificationARNs(notificationARNs)

	onFailure := stack.Spec.CloudFormationMeta.OnFailure
	if onFailure == "" {
		onFailure = "DELETE"
	}

	// TODO: Comeback and make this more thought out with cluster name and a hash
	input.SetStackName(namer(stack.ObjectMeta.Name, stack.ObjectMeta.Namespace))
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
