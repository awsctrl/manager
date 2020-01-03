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

	cloudformationv1alpha1 "go.awsctrl.io/manager/apis/cloudformation/v1alpha1"
	cloudformationutils "go.awsctrl.io/manager/controllers/cloudformation/utils"
	controllerutils "go.awsctrl.io/manager/controllers/utils"
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

// generateClientRequestToken will generate the client request token
func (r *StackReconciler) generateClientRequestToken(ctx context.Context, instance *cloudformationv1alpha1.Stack) error {
	instanceCopy := instance.DeepCopy()
	instanceCopy.Spec.ClientRequestToken = r.TokenClient.Generate()

	return r.Update(ctx, instanceCopy)
}
