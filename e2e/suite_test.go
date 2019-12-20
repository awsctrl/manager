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

package e2e_test

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"testing"
	"time"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/gexec"

	"k8s.io/apimachinery/pkg/types"
	"k8s.io/client-go/kubernetes/scheme"
	"k8s.io/client-go/rest"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/envtest"
	logf "sigs.k8s.io/controller-runtime/pkg/log"
	"sigs.k8s.io/controller-runtime/pkg/log/zap"

	"go.awsctrl.io/manager/aws"
	"go.awsctrl.io/manager/controllers/cloudformation"
	"go.awsctrl.io/manager/controllers/controllermanager"
	"go.awsctrl.io/manager/controllers/self"
	"go.awsctrl.io/manager/testutils"
	"go.awsctrl.io/manager/token"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	cloudformationv1alpha1 "go.awsctrl.io/manager/apis/cloudformation/v1alpha1"
	selfv1alpha1 "go.awsctrl.io/manager/apis/self/v1alpha1"
)

// These tests use Ginkgo (BDD-style Go testing framework). Refer to
// http://onsi.github.io/ginkgo/ to learn more about Ginkgo.

var (
	cfg          *rest.Config
	k8sclient    client.Client
	k8smanager   ctrl.Manager
	testenv      *envtest.Environment
	awsclient    aws.AWS
	configname   string = "config"
	podnamespace string = "default"
)

func TestAPIs(t *testing.T) {
	RegisterFailHandler(Fail)

	RunSpecsWithDefaultAndCustomReporters(t,
		"Controller Suite",
		[]Reporter{envtest.NewlineReporter{}})
}

var _ = BeforeSuite(func(done Done) {
	logf.SetLogger(zap.LoggerTo(GinkgoWriter, true))

	By("bootstrapping test environment")
	testenv = &envtest.Environment{
		CRDDirectoryPaths: []string{filepath.Join("..", "config", "crd", "bases")},
	}

	var err error
	cfg, err = testenv.Start()
	Expect(err).ToNot(HaveOccurred())
	Expect(cfg).ToNot(BeNil())

	err = scheme.AddToScheme(scheme.Scheme)
	Expect(err).NotTo(HaveOccurred())

	err = selfv1alpha1.AddToScheme(scheme.Scheme)
	Expect(err).NotTo(HaveOccurred())

	err = cloudformationv1alpha1.AddToScheme(scheme.Scheme)
	Expect(err).NotTo(HaveOccurred())

	err = controllermanager.AddAllSchemes(scheme.Scheme)
	Expect(err).NotTo(HaveOccurred())

	k8smanager, err = ctrl.NewManager(cfg, ctrl.Options{
		Scheme: scheme.Scheme,
	})
	Expect(err).ToNot(HaveOccurred())

	if os.Getenv("USE_AWS_CLIENT") == "true" {
		awsclient = aws.New()
	} else {
		awsclient = testutils.NewAWS()
	}

	err = (&self.ConfigReconciler{
		Client:       k8smanager.GetClient(),
		Log:          ctrl.Log.WithName("controllers").WithName("self").WithName("config"),
		Scheme:       k8smanager.GetScheme(),
		ConfigName:   configname,
		PodNamespace: podnamespace,
		AWSClient:    awsclient,
	}).SetupWithManager(k8smanager)
	Expect(err).ToNot(HaveOccurred())

	err = (&cloudformation.StackReconciler{
		Client:       k8smanager.GetClient(),
		Log:          ctrl.Log.WithName("controllers").WithName("cloudformation").WithName("stack"),
		Scheme:       k8smanager.GetScheme(),
		ConfigName:   configname,
		PodNamespace: podnamespace,
		AWSClient:    awsclient,
		TokenClient:  token.New(),
	}).SetupWithManager(k8smanager)
	Expect(err).ToNot(HaveOccurred())

	_, err = controllermanager.SetupControllers(k8smanager)
	Expect(err).ToNot(HaveOccurred())

	go func() {
		err = k8smanager.Start(ctrl.SetupSignalHandler())
		Expect(err).ToNot(HaveOccurred())
	}()

	k8sclient = k8smanager.GetClient()
	Expect(k8sclient).ToNot(BeNil())

	configkey := types.NamespacedName{
		Name:      configname,
		Namespace: podnamespace,
	}

	config := &selfv1alpha1.Config{
		ObjectMeta: metav1.ObjectMeta{
			Name:      configkey.Name,
			Namespace: configkey.Namespace,
		},
		Spec: selfv1alpha1.ConfigSpec{
			ClusterName: "test-cluster",
			Resources:   []string{},
			AWS: selfv1alpha1.ConfigAWS{
				DefaultRegion:    "us-west-2",
				AccountID:        os.Getenv("AWS_ACCOUNT_ID"),
				SupportedRegions: []string{"us-west-2"},
				Queue: selfv1alpha1.ConfigQueue{
					TopicARN: fmt.Sprintf("arn:aws:sns:us-west-2:%s:awsctrl", os.Getenv("AWS_ACCOUNT_ID")),
				},
			},
		},
	}
	Expect(k8sclient.Create(context.Background(), config)).Should(Succeed())

	close(done)
}, 60)

var _ = AfterSuite(func() {
	By("tearing down the test environment")
	gexec.KillAndWait(5 * time.Second)
	err := testenv.Stop()
	Expect(err).ToNot(HaveOccurred())
})
