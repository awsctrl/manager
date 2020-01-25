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
	s3v1alpha1 "go.awsctrl.io/manager/apis/s3/v1alpha1"

	metav1alpha1 "go.awsctrl.io/manager/apis/meta/v1alpha1"
)

// RunBucketSpecs allows all instance E2E tests to run
var _ = Describe("Run s3 Bucket Controller", func() {

	Context("Without Bucket{} existing", func() {

		It("Should create s3.Bucket{}", func() {
			var stackID string
			var stackName string
			var stack *cloudformationv1alpha1.Stack
			k8sclient := k8smanager.GetClient()
			Expect(k8sclient).ToNot(BeNil())

			instance := &s3v1alpha1.Bucket{
				ObjectMeta: metav1.ObjectMeta{
					GenerateName: "sample-bucket-",
					Namespace:    podnamespace,
				},
				Spec: s3v1alpha1.BucketSpec{
					AccessControl: "PublicRead",
					MetricsConfigurations: []s3v1alpha1.Bucket_MetricsConfiguration{
						s3v1alpha1.Bucket_MetricsConfiguration{
							Ref: metav1alpha1.ObjectReference{
								Id: "EntireBucket",
							},
						},
					},
					WebsiteConfiguration: s3v1alpha1.Bucket_WebsiteConfiguration{
						IndexDocument: "index.html",
						ErrorDocument: "error.html",
						RoutingRules: []s3v1alpha1.Bucket_RoutingRule{
							s3v1alpha1.Bucket_RoutingRule{
								RoutingRuleCondition: s3v1alpha1.Bucket_RoutingRuleCondition{
									HttpErrorCodeReturnedEquals: "404",
									KeyPrefixEquals:             "out1/",
								},
								RedirectRule: s3v1alpha1.Bucket_RedirectRule{
									HostName:             "ec2-11-22-333-44.compute-1.amazonaws.com",
									ReplaceKeyPrefixWith: "report-404/",
								},
							},
						},
					},
				},
			}
			By("Creating new s3 Bucket")
			Expect(k8sclient.Create(context.Background(), instance)).Should(Succeed())

			key := types.NamespacedName{
				Name:      instance.GetName(),
				Namespace: podnamespace,
			}

			By("Expecting CreateComplete")
			Eventually(func() bool {
				By("Getting latest s3 Bucket")
				instance = &s3v1alpha1.Bucket{}
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
				err := k8sclient.Get(context.Background(), stackkey, stack)
				if err != nil {
					return false
				}

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

			By("Deleting s3 Bucket")
			Expect(k8sclient.Delete(context.Background(), instance)).Should(Succeed())

			By("Deleting Bucket Stack")
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
