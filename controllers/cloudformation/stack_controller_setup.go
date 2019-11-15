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

	cloudformationv1alpha1 "awsctrl.io/apis/cloudformation/v1alpha1"
	selfv1alpha1 "awsctrl.io/apis/self/v1alpha1"
	cloudformationutils "awsctrl.io/controllers/cloudformation/utils"
	controllerutils "awsctrl.io/controllers/utils"
)

// addCFNFinalizer will add the deletion finalizer
func (r *StackReconciler) addCFNFinalizer(ctx context.Context, instance *cloudformationv1alpha1.Stack) error {
	instanceCopy := instance.DeepCopy()
	instanceCopy.ObjectMeta = controllerutils.AddFinalizer(instanceCopy.ObjectMeta, cloudformationutils.StackDeletionFinalizerName)
	return r.Update(ctx, instanceCopy)
}

// removeCFNFinalizer will remove the deletion finalizer
func (r *StackReconciler) removeCFNFinalizer(ctx context.Context, instance *cloudformationv1alpha1.Stack) error {
	instanceCopy := instance.DeepCopy()
	instanceCopy.ObjectMeta = controllerutils.RemoveFinalizer(instanceCopy.ObjectMeta, cloudformationutils.StackDeletionFinalizerName)
	return r.Update(ctx, instanceCopy)
}

// addNotificationARN will append the root notifcation arn for updates
func (r *StackReconciler) addNotificationARN(ctx context.Context, instance *cloudformationv1alpha1.Stack, config *selfv1alpha1.Config) error {
	instanceCopy := instance.DeepCopy()
	instanceCopy.Spec.NotificationARNs = append(instanceCopy.Spec.NotificationARNs, &config.Spec.AWS.Queue.TopicARN)

	return r.Update(ctx, instanceCopy)
}

// generateClientRequestToken will generate the client request token
func (r *StackReconciler) generateClientRequestToken(ctx context.Context, instance *cloudformationv1alpha1.Stack) error {
	instanceCopy := instance.DeepCopy()
	instanceCopy.Spec.ClientRequestToken = r.TokenClient.Generate()

	return r.Update(ctx, instanceCopy)
}
