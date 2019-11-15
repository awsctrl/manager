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

package cloudformation_test

import (
	"context"
	"time"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"k8s.io/apimachinery/pkg/types"

	cloudformationv1alpha1 "awsctrl.io/apis/cloudformation/v1alpha1"
	metav1alpha1 "awsctrl.io/apis/meta/v1alpha1"
	selfv1alpha1 "awsctrl.io/apis/self/v1alpha1"
	cloudformationutils "awsctrl.io/controllers/cloudformation/utils"
	controllerutils "awsctrl.io/controllers/utils"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var _ = Describe("Run Controller", func() {
	const timeout = time.Second * 30
	const interval = time.Second * 1

	Context("Run directly without existing job", func() {
		It("Should create successfully", func() {
			Expect(1).To(Equal(1))
		})
	})

	Context("Run a new Stack", func() {
		It("Should create successfully", func() {
			configkey := types.NamespacedName{Name: "config", Namespace: "default"}
			config := &selfv1alpha1.Config{
				ObjectMeta: metav1.ObjectMeta{
					Name:      configkey.Name,
					Namespace: configkey.Namespace,
				},
				Spec: selfv1alpha1.ConfigSpec{
					AWS: selfv1alpha1.ConfigAWS{
						SupportedRegions: []string{"us-west-2"},
						DefaultRegion:    "us-west-2",
					},
					Resources: []string{"cloudformation:stack"},
				},
			}
			Expect(k8sClient.Create(context.Background(), config)).Should(Succeed())

			stackkey := types.NamespacedName{Name: "test-stack", Namespace: "default"}
			stack := &cloudformationv1alpha1.Stack{
				ObjectMeta: metav1.ObjectMeta{
					Name:      stackkey.Name,
					Namespace: stackkey.Namespace,
				},
				Spec: cloudformationv1alpha1.StackSpec{
					Parameters: map[string]string{
						"name": "test-stack",
					},
					CloudFormationMeta: metav1alpha1.CloudFormationMeta{
						Region: "us-west-2",
					},
					TemplateBody: "",
				},
			}

			Expect(k8sClient.Create(context.Background(), stack)).Should(Succeed())
			time.Sleep(time.Second * 5)
			defer func() {
				Expect(k8sClient.Delete(context.Background(), stack)).Should(Succeed())
				Expect(k8sClient.Delete(context.Background(), config)).Should(Succeed())
				time.Sleep(time.Second * 5)
			}()

			By("Adding CFNFinalizer")
			Eventually(func() bool {
				f := &cloudformationv1alpha1.Stack{}
				k8sClient.Get(context.Background(), stackkey, f)
				return f.GetFinalizers()[0] == cloudformationutils.StackDeletionFinalizerName
			}, timeout, interval).Should(BeTrue())

			By("Adding ClientRequestToken")
			Eventually(func() bool {
				f := &cloudformationv1alpha1.Stack{}
				k8sClient.Get(context.Background(), stackkey, f)
				return f.Spec.ClientRequestToken != ""
			}, timeout, interval).Should(BeTrue())

			By("Adding NotificationARN")
			Eventually(func() bool {
				f := &cloudformationv1alpha1.Stack{}
				k8sClient.Get(context.Background(), stackkey, f)
				return len(f.GetNotificationARNs()) == 1
			}, timeout, interval).Should(BeTrue())

			By("Creating CFN Stack")
			Eventually(func() bool {
				f := &cloudformationv1alpha1.Stack{}
				k8sClient.Get(context.Background(), stackkey, f)
				return f.Status.StackID != ""
			}, timeout, interval).Should(BeTrue())

			By("Setting Template Version")
			Eventually(func() bool {
				f := &cloudformationv1alpha1.Stack{}
				k8sClient.Get(context.Background(), stackkey, f)
				return f.Labels[controllerutils.StackTemplateVersionLabel] != ""
			}, timeout, interval).Should(BeTrue())

			By("Updating CFN Stack")
			Eventually(func() bool {
				f := &cloudformationv1alpha1.Stack{}
				k8sClient.Get(context.Background(), stackkey, f)
				return f.Status.Status == metav1alpha1.UpdateCompleteStatus
			}, timeout, interval).Should(BeTrue())

			By("Describing Completed CFN Stack")
			Eventually(func() bool {
				f := &cloudformationv1alpha1.Stack{}
				k8sClient.Get(context.Background(), stackkey, f)
				return f.Status.Outputs["Name"] == "test"
			}, timeout, interval).Should(BeTrue())
		})
	})
})
