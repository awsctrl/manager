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
	"os"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"k8s.io/apimachinery/pkg/types"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/cloudformation"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	cloudformationv1alpha1 "go.awsctrl.io/manager/apis/cloudformation/v1alpha1"
	route53v1alpha1 "go.awsctrl.io/manager/apis/route53/v1alpha1"

	metav1alpha1 "go.awsctrl.io/manager/apis/meta/v1alpha1"
)

// RunAccountSpecs allows all recordset E2E tests to run
var _ = Describe("Run Route53 RecordSet Controller", func() {

	Context("Without RecordSet{} existing", func() {

		It("Should create route53.RecordSet{}", func() {
			var stackID string
			var stackName string
			var stack *cloudformationv1alpha1.Stack
			k8sclient := k8smanager.GetClient()
			Expect(k8sclient).ToNot(BeNil())

			hostedzone := &route53v1alpha1.HostedZone{
				ObjectMeta: metav1.ObjectMeta{
					GenerateName: "sample-hostedzone-",
					Namespace:    podnamespace,
				},
				Spec: route53v1alpha1.HostedZoneSpec{
					Name: "e2e.awsctrl.io",
					HostedZoneConfig: route53v1alpha1.HostedZone_HostedZoneConfig{
						Comment: "e2e created hosted zone",
					},
				},
			}
			By("Creating new Route53 HostedZone")
			Expect(k8sclient.Create(context.Background(), hostedzone)).Should(Succeed())

			hostedzonekey := types.NamespacedName{
				Name:      hostedzone.GetName(),
				Namespace: podnamespace,
			}

			instance := &route53v1alpha1.RecordSet{
				ObjectMeta: metav1.ObjectMeta{
					GenerateName: "sample-recordset-",
					Namespace:    podnamespace,
				},
				Spec: route53v1alpha1.RecordSetSpec{
					Type:            "A",
					Name:            "test.e2e.awsctrl.io",
					TTL:             "300",
					ResourceRecords: []string{"104.198.14.52"},
					HostedZoneRef: metav1alpha1.ObjectReference{
						ObjectRef: metav1alpha1.ObjectRef{
							Kind:       "HostedZone",
							APIVersion: "route53.awsctrl.io/v1alpha1",
							Name:       hostedzonekey.Name,
							Key:        "ResourceRef",
						},
					},
				},
			}
			By("Creating new Route53 RecordSet")
			Expect(k8sclient.Create(context.Background(), instance)).Should(Succeed())

			key := types.NamespacedName{
				Name:      instance.GetName(),
				Namespace: podnamespace,
			}

			By("Expecting CreateComplete")
			Eventually(func() bool {
				By("Getting latest Route53 RecordSet")
				instance = &route53v1alpha1.RecordSet{}
				err := k8sclient.Get(context.Background(), key, instance)
				if err != nil {
					return false
				}

				stackID = instance.GetStackID()
				stackName = instance.GetStackName()

				return instance.Status.Status == metav1alpha1.CreateCompleteStatus ||
					(os.Getenv("USE_AWS_CLIENT") != "true" && instance.Status.Status != "")
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

			By("Deleting Route53 RecordSet")
			Expect(k8sclient.Delete(context.Background(), instance)).Should(Succeed())

			By("Deleting Route53 RecordSet Stack")
			Expect(k8sclient.Delete(context.Background(), stack)).Should(Succeed())

			By("Expecting metav1alpha1.DeleteCompleteStatus")
			Eventually(func() bool {
				if os.Getenv("USE_AWS_CLIENT") != "true" {
					return true
				}

				output, err := awsclient.GetClient("us-west-2").DescribeStacks(&cloudformation.DescribeStacksInput{StackName: aws.String(stackID)})
				Expect(err).To(BeNil())
				stackoutput := output.Stacks[0].StackStatus
				return *stackoutput == "DELETE_COMPLETE"
			}, timeout, interval).Should(BeTrue())
		})
	})
})
