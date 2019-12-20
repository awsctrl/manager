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
	"time"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"k8s.io/apimachinery/pkg/types"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/cloudformation"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	apigatewayv1alpha1 "go.awsctrl.io/manager/apis/apigateway/v1alpha1"
	cloudformationv1alpha1 "go.awsctrl.io/manager/apis/cloudformation/v1alpha1"

	metav1alpha1 "go.awsctrl.io/manager/apis/meta/v1alpha1"
)

// RunAccountSpecs allows all instance E2E tests to run
var _ = Describe("Run Apigateway Account Controller", func() {
	const timeout = time.Second * 60
	const interval = time.Second * 1

	Context("Without Account{} existing", func() {

		It("Should create apigateway.Account{}", func() {
			var stackID string
			var stackName string
			var stack *cloudformationv1alpha1.Stack
			k8sclient := k8smanager.GetClient()
			Expect(k8sclient).ToNot(BeNil())

			key := types.NamespacedName{
				Name:      "sample-name",
				Namespace: "default",
			}

			instance := &apigatewayv1alpha1.Account{
				ObjectMeta: metav1.ObjectMeta{
					Name:      key.Name,
					Namespace: key.Namespace,
				},
				Spec: apigatewayv1alpha1.AccountSpec{
					CloudWatchRoleArn: fmt.Sprintf("arn:aws:iam::%s:role/awsctrl-apigateway-instance", os.Getenv("AWS_ACCOUNT_ID")),
				},
			}
			By("Creating new Apigateway Account")
			Expect(k8sclient.Create(context.Background(), instance)).Should(Succeed())

			By("Expecting CreateComplete")
			Eventually(func() bool {
				By("Getting latest Apigateway Account")
				instance = &apigatewayv1alpha1.Account{}
				Expect(k8sclient.Get(context.Background(), key, instance)).Should(Succeed())

				stackID = instance.GetStackID()
				stackName = instance.GetStackName()

				return instance.Status.Status == metav1alpha1.CreateCompleteStatus ||
					(instance.Status.Status == metav1alpha1.UpdateCompleteStatus && os.Getenv("USE_AWS_CLIENT") != "true")
			}, timeout, interval).Should(BeTrue())

			By("Checking object OwnerShip")
			Eventually(func() bool {
				stackkey := types.NamespacedName{
					Name:      stackName,
					Namespace: key.Namespace,
				}

				stack = &cloudformationv1alpha1.Stack{}
				Expect(k8sclient.Get(context.Background(), stackkey, stack)).Should(Succeed())

				expectedOwnerReference := v1.OwnerReference{
					Kind:       instance.Kind,
					APIVersion: instance.APIVersion,
					UID:        instance.UID,
					Name:       instance.Name,
				}

				ownerrefs := stack.GetOwnerReferences()
				Expect(len(ownerrefs)).To(Equal(1))

				return ownerrefs[0].Name == expectedOwnerReference.Name
			}, timeout, interval).Should(BeTrue())

			By("Deleting Apigateway Account")
			Expect(k8sclient.Delete(context.Background(), instance)).Should(Succeed())

			By("Deleting Account Stack")
			Expect(k8sclient.Delete(context.Background(), stack)).Should(Succeed())

			By("Expecting metav1alpha1.DeleteCompleteStatus")
			Eventually(func() bool {
				output, err := awsclient.GetClient("us-west-2").DescribeStacks(&cloudformation.DescribeStacksInput{StackName: aws.String(stackID)})
				Expect(err).To(BeNil())
				stackoutput := output.Stacks[0].StackStatus
				return *stackoutput == "DELETE_COMPLETE" || os.Getenv("AWS_ACCOUNT_ID") != "true"
			}, timeout, interval).Should(BeTrue())
		})
	})
})
