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

package cloudformation

import (
	"context"
	"fmt"
	"strings"

	"k8s.io/apimachinery/pkg/types"

	cloudformationv1alpha1 "awsctrl.io/apis/cloudformation/v1alpha1"
	metav1alpha1 "awsctrl.io/apis/meta/v1alpha1"
	cloudformationutils "awsctrl.io/controllers/cloudformation/utils"
	controllerutils "awsctrl.io/controllers/utils"

	cfn "github.com/aws/aws-sdk-go/service/cloudformation"

	"github.com/iancoleman/strcase"
)

// createCFNStack will create a new CFN Stack
func (r *StackReconciler) createCFNStack(ctx context.Context, instance *cloudformationv1alpha1.Stack) error {
	output, err := cloudformationutils.CreateCFNStack(r.AWSClient, instance)
	if err != nil {
		return err
	}

	if err := r.updateStackTemplateVersion(ctx, types.NamespacedName{Namespace: instance.Namespace, Name: instance.Name}); err != nil {
		return err
	}

	return r.updateCFNStackStatus(ctx, instance, metav1alpha1.CreateInProgressStatus, "", string(*output.StackId), map[string]string{})
}

// updateCFNStack will update the CFN Stack
func (r *StackReconciler) updateCFNStack(ctx context.Context, instance *cloudformationv1alpha1.Stack) error {
	if err := cloudformationutils.UpdateCFNStack(r.AWSClient, instance); err != nil {
		return err
	}

	if err := r.updateStackTemplateVersion(ctx, types.NamespacedName{Namespace: instance.Namespace, Name: instance.Name}); err != nil {
		return err
	}

	return r.updateCFNStackStatus(ctx, instance, metav1alpha1.UpdateInProgressStatus, "", instance.Status.StackID, map[string]string{})
}

// deleteCFNStack will delete the CFN Stack
func (r *StackReconciler) deleteCFNStack(ctx context.Context, instance *cloudformationv1alpha1.Stack) error {
	if err := cloudformationutils.DeleteCFNStack(r.AWSClient, instance); err != nil {
		return err
	}

	return r.updateCFNStackStatus(ctx, instance, metav1alpha1.DeleteInProgressStatus, "", instance.Status.StackID, map[string]string{})
}

// describeCFNStackStatus will get the latest from CFN Stacks and update etcd
func (r *StackReconciler) describeCFNStackStatus(ctx context.Context, instance *cloudformationv1alpha1.Stack) (err error) {
	var outputs *cfn.DescribeStacksOutput
	if outputs, err = cloudformationutils.DescribeCFNStacks(r.AWSClient, instance); err != nil {
		return err
	}

	if len(outputs.Stacks) == 0 {
		return fmt.Errorf("could not find stack with name '%s'", instance.Name)
	}

	nsn := types.NamespacedName{Namespace: instance.GetNamespace(), Name: instance.GetName()}

	var instanceCopy cloudformationv1alpha1.Stack
	if err = r.Get(ctx, nsn, &instanceCopy); err != nil {
		return err
	}

	outputsMap := map[string]string{}
	for _, output := range outputs.Stacks[0].Outputs {
		outputsMap[string(*output.OutputKey)] = string(*output.OutputValue)
	}

	var reason string
	if outputs.Stacks[0].StackStatusReason != nil {
		reason = *outputs.Stacks[0].StackStatusReason
	}
	status := metav1alpha1.ConditionStatus(strcase.ToCamel(strings.ToLower(string(*outputs.Stacks[0].StackStatus))))

	return r.updateCFNStackStatus(ctx, &instanceCopy, status, reason, *outputs.Stacks[0].StackId, outputsMap)
}

// updateCFNStackStatus will update the stack status object
func (r *StackReconciler) updateCFNStackStatus(ctx context.Context, instance *cloudformationv1alpha1.Stack, status metav1alpha1.ConditionStatus, message, stackID string, outputs map[string]string) error {
	nsn := types.NamespacedName{Namespace: instance.GetNamespace(), Name: instance.GetName()}

	var instanceCopy cloudformationv1alpha1.Stack
	if err := r.Get(ctx, nsn, &instanceCopy); err != nil {
		return err
	}

	stackstatus := cloudformationv1alpha1.StackStatus{
		StatusMeta: metav1alpha1.StatusMeta{
			StackID:            stackID,
			Status:             status,
			Message:            &message,
			ObservedGeneration: instanceCopy.Generation,
		},
	}

	newOutputs := map[string]string{}
	if len(instanceCopy.Status.Outputs) == 0 {
		newOutputs = outputs
	} else {
		newOutputs = instanceCopy.Status.Outputs
		for key, value := range outputs {
			newOutputs[key] = value
		}
	}

	stackstatus.Outputs = newOutputs
	instanceCopy.Status = stackstatus

	return r.Status().Update(ctx, &instanceCopy)
}

// updateStackTemplateVersion will add the template version
func (r *StackReconciler) updateStackTemplateVersion(ctx context.Context, namespaceName types.NamespacedName) error {
	var instance cloudformationv1alpha1.Stack
	if err := r.Get(ctx, namespaceName, &instance); err != nil {
		return err
	}
	instanceCopy := instance.DeepCopy()

	if len(instanceCopy.Labels) == 0 {
		instanceCopy.Labels = map[string]string{}
	}

	instanceCopy.Labels[controllerutils.StackTemplateVersionLabel] = controllerutils.ComputeHash(instanceCopy.Spec)

	return r.Update(ctx, instanceCopy)
}
