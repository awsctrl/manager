package config

import (
	"context"
	"testing"
	"time"

	selfv1alpha1 "awsctrl.io/pkg/apis/self/v1alpha1"
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

const timeout = time.Second * 5

func TestReconcile(t *testing.T) {
	g := gomega.NewGomegaWithT(t)
	instance := &selfv1alpha1.Config{
		ObjectMeta: metav1.ObjectMeta{
			Name:      "foo",
			Namespace: "default",
		},
		Spec: selfv1alpha1.ConfigSpec{
			Resources: []string{},
			AWS: selfv1alpha1.ConfigAWS{
				SupportedRegions: []string{},
			},
		},
	}

	// Setup the Manager and Controller.  Wrap the Controller Reconcile function so it writes each request to a
	// channel when it is finished.
	mgr, err := manager.New(cfg, manager.Options{})
	g.Expect(err).NotTo(gomega.HaveOccurred())
	c = mgr.GetClient()

	recFn, requests := SetupTestReconcile(newReconciler(mgr))
	g.Expect(add(mgr, recFn)).To(gomega.Succeed())

	stopMgr, mgrStopped := StartTestManager(mgr, g)

	defer func() {
		close(stopMgr)
		mgrStopped.Wait()
	}()

	// Create the Config object and expect the Reconcile and Deployment to be created
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

	// Delete the Config and expect Reconcile
	g.Expect(c.Delete(context.TODO(), instance)).To(gomega.Succeed())
	g.Eventually(requests, timeout).Should(gomega.Receive(gomega.Equal(expectedRequest)))

	// Manually delete Config since GC isn't enabled in the test control plane
	g.Eventually(func() error { return c.Delete(context.TODO(), instance) }, timeout).
		Should(gomega.MatchError("configs.self.awsctrl.io \"foo\" not found"))
}
