package aws

import (
	"sync"

	"github.com/aws/aws-sdk-go/aws"

	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/cloudformation"
	"github.com/aws/aws-sdk-go/service/cloudformation/cloudformationiface"

	selfv1alpha1 "awsctrl.io/pkg/apis/self/v1alpha1"
)

// AWS Sets and Gets clients based on the config
type AWS interface {
	// Configure will setup the clients
	Configure(*selfv1alpha1.ConfigAWS) error

	// SetClient will set a client by region
	SetClient(string, cloudformationiface.CloudFormationAPI) bool

	// GetClients will return all clients
	GetClients() map[string]cloudformationiface.CloudFormationAPI

	// GetClient will return a client for a region
	GetClient(string) cloudformationiface.CloudFormationAPI

	// SetSession will set a session by region
	SetSession(string, *session.Session) bool

	// GetSessions will return all sessions
	GetSessions() map[string]*session.Session

	// GetSession will return a session for a region
	GetSession(string) *session.Session

	// GetNotificationARN will return the notification arn for subs
	GetNotificationARN() string

	// GetDefaultRegion will return the static default
	GetDefaultRegion() string
}

// New return AWS client
func New() AWS {
	clients := map[string]cloudformationiface.CloudFormationAPI{}
	sessions := map[string]*session.Session{}

	return &awsClient{
		mutex:    sync.Mutex{},
		clients:  clients,
		sessions: sessions,
	}
}

// awsClient will hold all the config data for AWS
type awsClient struct {
	mutex sync.Mutex

	clients  map[string]cloudformationiface.CloudFormationAPI
	sessions map[string]*session.Session

	notficationARN string
	defaultRegion  string

	rawConfig *selfv1alpha1.ConfigAWS
}

// Configure will configure the clients
func (a *awsClient) Configure(config *selfv1alpha1.ConfigAWS) error {
	a.rawConfig = config

	for _, region := range config.SupportedRegions {
		sess, err := session.NewSession(&aws.Config{Region: aws.String(region)})
		if err != nil {
			return err
		}
		a.SetSession(region, sess)
		a.SetClient(region, cloudformation.New(sess))
	}

	a.notficationARN = config.Queue.TopicARN
	a.defaultRegion = config.DefaultRegion

	return nil
}

// SetClient will set a client by region
func (a *awsClient) SetClient(region string, iface cloudformationiface.CloudFormationAPI) bool {
	a.mutex.Lock()
	defer a.mutex.Unlock()
	a.clients[region] = iface
	return true
}

// GetClients will return all clients
func (a *awsClient) GetClients() map[string]cloudformationiface.CloudFormationAPI {
	return a.clients
}

// GetClient will return a client for a region
func (a *awsClient) GetClient(region string) cloudformationiface.CloudFormationAPI {
	if region == "" {
		region = a.GetDefaultRegion()
	}
	return a.GetClients()[region]
}

// SetSession will set a session by region
func (a *awsClient) SetSession(region string, sess *session.Session) bool {
	a.mutex.Lock()
	defer a.mutex.Unlock()
	a.sessions[region] = sess
	return true
}

// GetSessions will return all sessions
func (a *awsClient) GetSessions() map[string]*session.Session {
	return a.sessions
}

// GetSession will return a session for a region
func (a *awsClient) GetSession(region string) *session.Session {
	if region == "" {
		region = a.GetDefaultRegion()
	}
	return a.GetSessions()[region]
}

func (a *awsClient) GetNotificationARN() string {
	return a.notficationARN
}

func (a *awsClient) GetDefaultRegion() string {
	return a.defaultRegion
}
