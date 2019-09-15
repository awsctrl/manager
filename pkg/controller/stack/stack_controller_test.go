package stack

import (
	"context"
	"testing"
	"time"

	cloudformationv1alpha1 "awsctrl.io/pkg/apis/cloudformation/v1alpha1"
	"awsctrl.io/pkg/testutils"
	"github.com/onsi/gomega"
	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/manager"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"
)

var c client.Client

var expectedRequest = reconcile.Request{NamespacedName: types.NamespacedName{Name: "foo", Namespace: "default"}}

var stackKey = types.NamespacedName{Name: "foo", Namespace: "default"}

const timeout = time.Second * 5

func TestReconcile(t *testing.T) {
	g := gomega.NewGomegaWithT(t)
	instance := &cloudformationv1alpha1.Stack{
		ObjectMeta: metav1.ObjectMeta{
			Name:      "foo",
			Namespace: "default",
		},
		Spec: cloudformationv1alpha1.StackSpec{
			Parameters: map[string]string{},
		},
	}

	// Setup the Manager and Controller.  Wrap the Controller Reconcile function so it writes each request to a
	// channel when it is finished.
	mgr, err := manager.New(cfg, manager.Options{})
	g.Expect(err).NotTo(gomega.HaveOccurred())
	c = mgr.GetClient()

	recFn, requests := SetupTestReconcile(NewReconciler(mgr, testutils.NewAWS(), testutils.NewToken(), "default"))
	g.Expect(add(mgr, recFn)).To(gomega.Succeed())

	stopMgr, mgrStopped := StartTestManager(mgr, g)

	defer func() {
		close(stopMgr)
		mgrStopped.Wait()
	}()

	// Create the Stack object and expect the Reconcile and Deployment to be created
	err = c.Create(context.TODO(), instance)
	// The instance object may not be a valid object because it might be missing some required fields.
	// Please modify the instance object by adding required fields and then remove the following if statement.
	if errors.IsInvalid(err) {
		t.Logf("failed to create object, got an invalid object error: %v", err)
		return
	}
	g.Expect(err).NotTo(gomega.HaveOccurred())
	defer c.Delete(context.TODO(), instance)
	g.Eventually(requests, timeout).Should(gomega.Receive(gomega.Equal(expectedRequest)))

	// Delete the Stack and expect Reconcile
	g.Expect(c.Delete(context.TODO(), instance)).To(gomega.Succeed())
	g.Eventually(requests, timeout).Should(gomega.Receive(gomega.Equal(expectedRequest)))

	// Manually delete Stack since GC isn't enabled in the test control plane
	g.Eventually(func() error { return c.Delete(context.TODO(), instance) }, timeout).
		Should(gomega.MatchError("stacks.cloudformation.awsctrl.io \"foo\" not found"))
}
