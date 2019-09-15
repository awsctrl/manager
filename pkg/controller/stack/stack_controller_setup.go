package stack

import (
	"context"

	cloudformationv1alpha1 "awsctrl.io/pkg/apis/cloudformation/v1alpha1"
	selfv1alpha1 "awsctrl.io/pkg/apis/self/v1alpha1"
	"awsctrl.io/pkg/controllerutils"
)

func (r *ReconcileStack) addCFNFinalizer(ctx context.Context, instance *cloudformationv1alpha1.Stack) error {
	instanceCopy := instance.DeepCopy()
	instanceCopy.ObjectMeta = controllerutils.AddFinalizer(instanceCopy.ObjectMeta, stackDeletionFinalizerName)
	return r.Update(ctx, instanceCopy)
}

func (r *ReconcileStack) addNotificationARN(ctx context.Context, instance *cloudformationv1alpha1.Stack, config *selfv1alpha1.Config) error {
	instanceCopy := instance.DeepCopy()
	instanceCopy.Spec.NotificationARNs = append(instanceCopy.Spec.NotificationARNs, &config.Spec.AWS.Queue.TopicARN)

	return r.Update(ctx, instanceCopy)
}

// generateClientRequestToken will generate the client request token
func (r *ReconcileStack) generateClientRequestToken(ctx context.Context, instance *cloudformationv1alpha1.Stack) error {
	instanceCopy := instance.DeepCopy()
	instanceCopy.Spec.ClientRequestToken = r.clientToken.Generate()

	return r.Update(ctx, instanceCopy)
}
