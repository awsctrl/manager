package queue

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	cloudformationv1alpha1 "awsctrl.io/pkg/apis/cloudformation/v1alpha1"
	selfv1alpha1 "awsctrl.io/pkg/apis/self/v1alpha1"
	"awsctrl.io/pkg/event"

	"github.com/aws/aws-sdk-go/aws"
	awsclient "github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sns"
	"github.com/aws/aws-sdk-go/service/sns/snsiface"
	"github.com/aws/aws-sdk-go/service/sqs"
	"github.com/aws/aws-sdk-go/service/sqs/sqsiface"

	k8scache "k8s.io/client-go/tools/cache"

	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/cache"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/manager"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"
	logf "sigs.k8s.io/controller-runtime/pkg/runtime/log"
)

var (
	log = logf.Log.WithName("queue")
)

// Queue will define all the functions of the queue
type Queue interface {
	Start(<-chan struct{}) error
	Reconcile(*event.Event) error
}

type queueInformer struct {
	client.Client
	cache.Cache

	region    string
	queueName string
	queueURL  string
	queueARN  string
	topicARN  string
	subARN    string

	sess      *session.Session
	sqsclient sqsiface.SQSAPI
	snsclient snsiface.SNSAPI

	Handler func(reconcile.Request) (reconcile.Result, error)
}

// New will create a new client to watch
func New(mgr manager.Manager, handler func(reconcile.Request) (reconcile.Result, error)) Queue {
	return &queueInformer{
		Client:  mgr.GetClient(),
		Cache:   mgr.GetCache(),
		Handler: handler,
	}
}

// Start will start the queue manager
func (q *queueInformer) Start(stopCh <-chan struct{}) error {
	gvk := schema.GroupVersionKind{Group: "self.awsctrl.io", Version: "v1alpha1", Kind: "Config"}
	configInformer, err := q.Cache.GetInformerForKind(gvk)

	if ok := k8scache.WaitForCacheSync(stopCh, configInformer.HasSynced); !ok {
		log.Info("failed to wait for caches to sync")
	}

	ctx := context.Background()

	var config selfv1alpha1.Config
	if err := q.Client.Get(ctx, types.NamespacedName{Name: "config", Namespace: os.Getenv("POD_NAMESPACE")}, &config); err != nil {
		return err
	}

	region := config.Spec.AWS.DefaultRegion

	sess, err := session.NewSession(&awsclient.Config{Region: awsclient.String(region)})
	if err != nil {
		log.Error(err, "error building aws session")
		return err
	}

	q.sess = sess
	q.region = region
	q.sqsclient = sqs.New(sess)
	q.snsclient = sns.New(sess)
	q.queueName = config.Spec.ClusterName
	q.queueURL = config.Spec.AWS.Queue.QueueURL
	q.queueARN = config.Spec.AWS.Queue.QueueARN
	q.topicARN = config.Spec.AWS.Queue.TopicARN
	q.subARN = config.Spec.AWS.Queue.SubARN

	if q.queueURL == "" && q.queueARN == "" && q.topicARN == "" && q.subARN == "" {
		if err = q.createSQSQueue(); err != nil {
			return err
		}

		if err = q.createSNSTopic(); err != nil {
			return err
		}

		if err = q.setPolicy(); err != nil {
			return err
		}

		config.Spec.AWS.Queue.QueueURL = q.queueURL
		config.Spec.AWS.Queue.QueueARN = q.queueARN
		config.Spec.AWS.Queue.TopicARN = q.topicARN
		config.Spec.AWS.Queue.SubARN = q.subARN
	}

	config.Spec.AWS.Queue.Name = q.queueName
	config.Spec.AWS.Queue.Region = q.region

	if err := q.Client.Update(ctx, &config); err != nil {
		return err
	}

	go func() {
		for {
			select {
			case <-stopCh:
				log.Info("shutting down SQS listener")
				return
			default:
			}

			input := &sqs.ReceiveMessageInput{}
			input.SetQueueUrl(q.queueURL)
			input.SetAttributeNames(aws.StringSlice([]string{"SentTimestamp"}))
			input.SetMaxNumberOfMessages(1)
			input.SetMessageAttributeNames(aws.StringSlice([]string{"All"}))
			input.SetWaitTimeSeconds(10)

			output, err := q.sqsclient.ReceiveMessage(input)
			if err != nil {
				log.Info("error pulling messages off the queue", "error", err)
				return
			}

			for _, message := range output.Messages {
				evtMessage := &event.Message{}
				err := json.Unmarshal([]byte(*message.Body), evtMessage)
				if err != nil {
					log.Info("error unmarshalling the message body", "error", err)
					break
				}

				evt := &event.Event{}
				err = event.Unmarshal(evtMessage.Message, evt)
				if err != nil {
					log.Info("error unmarshalling the message", "error", err)
					break
				}

				if err := q.Reconcile(evt); err != nil {
					log.Info("error deleting the message", "error", err)
					break
				}

				deleteInput := &sqs.DeleteMessageInput{}
				deleteInput.SetQueueUrl(q.queueURL)
				deleteInput.SetReceiptHandle(*message.ReceiptHandle)
				_, err = q.sqsclient.DeleteMessage(deleteInput)
				if err != nil {
					log.Info("error deleting the message", "error", err)
					break
				}
			}
		}
	}()
	return nil
}

func (q *queueInformer) Reconcile(evt *event.Event) error {
	ctx := context.Background()

	if evt.StackId == "" {
		return fmt.Errorf("")
	}

	var instances cloudformationv1alpha1.StackList
	if err := q.Client.List(ctx, &instances, client.MatchingField("spec.status.stackID", evt.StackId)); err != nil {
		if errors.IsNotFound(err) {
			return nil
		}
		return err
	}

	for _, stack := range instances.Items {
		_, err := q.Handler(reconcile.Request{NamespacedName: types.NamespacedName{Namespace: stack.Namespace, Name: stack.Name}})
		if err != nil {
			log.Info("reconcile", "stack", stack, "error", err)
		}
	}

	return nil
}

func (q *queueInformer) createQueue() error {
	input := &sqs.CreateQueueInput{}
	input.SetQueueName(q.queueName)
	output, err := q.sqsclient.CreateQueue(input)
	if err != nil {
		return err
	}
	q.queueURL = *output.QueueUrl
	return nil
}

func (q *queueInformer) getQueueURL() error {
	input := &sqs.GetQueueUrlInput{}
	input.SetQueueName(q.queueName)
	output, err := q.sqsclient.GetQueueUrl(input)
	if err != nil {
		return err
	}
	q.queueURL = *output.QueueUrl
	return nil
}

func (q *queueInformer) createSQSQueue() error {
	if q.queueName == "" {
		return fmt.Errorf("SQS Queue name can't be blank")
	}

	if err := q.getQueueURL(); err != nil {
		log.Info("getting queue URL", "error", err)

		if err = q.createQueue(); err != nil {
			return err
		}
	}

	queueQueryInputs := &sqs.GetQueueAttributesInput{}
	queueQueryInputs.SetQueueUrl(q.queueURL)
	queueQueryInputs.SetAttributeNames([]*string{aws.String("All")})

	sqsQueueOutput, err := q.sqsclient.GetQueueAttributes(queueQueryInputs)
	if err != nil {
		return err
	}
	q.queueARN = *sqsQueueOutput.Attributes["QueueArn"]

	return nil
}

func (q *queueInformer) createSNSTopic() error {
	topicInputs := &sns.CreateTopicInput{}
	topicInputs.SetName(q.queueName)
	output, err := q.snsclient.CreateTopic(topicInputs)
	if err != nil {
		return err
	}
	q.topicARN = *output.TopicArn

	subInput := &sns.SubscribeInput{}
	subInput.SetTopicArn(q.topicARN)
	subInput.SetEndpoint(q.queueARN)
	subInput.SetProtocol("sqs")

	subOutput, err := q.snsclient.Subscribe(subInput)
	if err != nil {
		return err
	}
	q.subARN = *subOutput.SubscriptionArn

	return nil
}

func (q *queueInformer) setPolicy() error {
	policy := newPolicy(q.queueARN, q.queueName, []string{q.topicARN})
	policyb, err := json.Marshal(policy)
	if err != nil {
		return err
	}

	input := &sqs.SetQueueAttributesInput{}
	input.SetQueueUrl(q.queueURL)
	input.SetAttributes(map[string]*string{"Policy": aws.String(string(policyb))})
	_, err = q.sqsclient.SetQueueAttributes(input)
	if err != nil {
		return err
	}
	return nil
}

func newPolicy(queueARN, name string, topicARNs []string) Policy {
	statements := []Statement{}
	for _, topicARN := range topicARNs {
		statements = append(statements, Statement{
			Sid:       topicARN + " statment",
			Effect:    "Allow",
			Principal: "*",
			Action:    []string{"SQS:ReceiveMessage", "SQS:SendMessage"},
			Resource:  queueARN,
			Condition: Condition{
				ArnEquals: ArnEquals{
					AwsSourceArn: topicARN,
				},
			},
		})
	}

	return Policy{
		Version:   "2012-10-17",
		ID:        queueARN + "/" + name + "-policy",
		Statement: statements,
	}
}
