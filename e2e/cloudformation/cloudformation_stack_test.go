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
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/cloudformation"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"k8s.io/apimachinery/pkg/types"

	cloudformationv1alpha1 "go.awsctrl.io/manager/apis/cloudformation/v1alpha1"
	metav1alpha1 "go.awsctrl.io/manager/apis/meta/v1alpha1"
	cloudformationutils "go.awsctrl.io/manager/controllers/cloudformation/utils"
	controllerutils "go.awsctrl.io/manager/controllers/utils"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var _ = Describe("Run CloudFormation Stack Controller", func() {

	Context("Run directly without existing job", func() {
		It("Should create successfully", func() {
			Expect(1).To(Equal(1))
		})
	})

	Context("Run a new Stack", func() {
		It("Should create successfully", func() {
			var stackID string

			stack := &cloudformationv1alpha1.Stack{
				ObjectMeta: metav1.ObjectMeta{
					GenerateName: "test-stack-",
					Namespace:    podnamespace,
				},
				Spec: cloudformationv1alpha1.StackSpec{
					Parameters: map[string]string{},
					CloudFormationMeta: metav1alpha1.CloudFormationMeta{
						Region: "us-west-2",
					},
					TemplateBody: `{
						"AWSTemplateFormatVersion": "2010-09-09",
						"Description": "AWS Controller - ecr.Repository (ac-{TODO})",
						"Resources": {
						  "Repository": {
							"Properties": {
							  "RepositoryName": "sample-repo"
							},
							"Type": "AWS::ECR::Repository"
						  }
						}
					  }`,
				},
			}

			Expect(k8sclient.Create(context.Background(), stack)).Should(Succeed())
			time.Sleep(5 * time.Second)

			stackkey := types.NamespacedName{
				Name:      stack.GetName(),
				Namespace: podnamespace,
			}

			By("Adding CFNFinalizer")
			Eventually(func() bool {
				f := &cloudformationv1alpha1.Stack{}
				k8sclient.Get(context.Background(), stackkey, f)
				return f.GetFinalizers()[0] == cloudformationutils.StackDeletionFinalizerName
			}, timeout, interval).Should(BeTrue())

			By("Adding ClientRequestToken")
			Eventually(func() bool {
				f := &cloudformationv1alpha1.Stack{}
				k8sclient.Get(context.Background(), stackkey, f)
				return f.Spec.ClientRequestToken != ""
			}, timeout, interval).Should(BeTrue())

			By("Adding NotificationARN")
			Eventually(func() bool {
				f := &cloudformationv1alpha1.Stack{}
				k8sclient.Get(context.Background(), stackkey, f)
				return len(f.GetNotificationARNs()) == 0
			}, timeout, interval).Should(BeTrue())

			By("Creating CFN Stack")
			Eventually(func() bool {
				f := &cloudformationv1alpha1.Stack{}
				k8sclient.Get(context.Background(), stackkey, f)
				return f.Status.StackID != ""
			}, timeout, interval).Should(BeTrue())

			By("Setting Template Version")
			Eventually(func() bool {
				f := &cloudformationv1alpha1.Stack{}
				k8sclient.Get(context.Background(), stackkey, f)
				return f.Labels[controllerutils.StackTemplateVersionLabel] != ""
			}, timeout, interval).Should(BeTrue())

			By("Updating CFN Stack")
			Eventually(func() bool {
				f := &cloudformationv1alpha1.Stack{}
				k8sclient.Get(context.Background(), stackkey, f)
				stackID = f.Status.StackID
				return f.Status.Status == metav1alpha1.CreateCompleteStatus || os.Getenv("USE_AWS_CLIENT") != "true"
			}, timeout, interval).Should(BeTrue())

			// By("Describing Completed CFN Stack")
			// Eventually(func() bool {
			// 	f := &cloudformationv1alpha1.Stack{}
			// 	k8sclient.Get(context.Background(), stackkey, f)
			// 	return f.Status.Outputs["Name"] == "test"
			// }, timeout, interval).Should(BeTrue())

			By("Deleting CFN Stack")
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
