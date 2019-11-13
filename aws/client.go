/*
Copyright © 2019 AWS Controller author

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

package aws

import (
	"fmt"
	"sync"

	"github.com/aws/aws-sdk-go/aws"

	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/cloudformation"
	"github.com/aws/aws-sdk-go/service/cloudformation/cloudformationiface"

	selfv1alpha1 "awsctrl.io/apis/self/v1alpha1"
)

// AWS Sets and Gets clients based on the config
type AWS interface {
	// Configured will return an error if client isn't configured
	Configured() error

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
		mutex:      sync.Mutex{},
		clients:    clients,
		sessions:   sessions,
		configured: false,
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

	configured bool
}

// Configured will check if the client is configured
func (a *awsClient) Configured() error {
	if a.configured {
		return nil
	}
	return fmt.Errorf("aws client not loaded and configured, check for self.awsctrl.io/config")
}

// Configure will configure the clients
func (a *awsClient) Configure(config *selfv1alpha1.ConfigAWS) error {
	for _, region := range config.SupportedRegions {
		sess, err := session.NewSession(&aws.Config{Region: aws.String(region)})
		if err != nil {
			return err
		}
		a.SetSession(region, sess)
		a.SetClient(region, cloudformation.New(sess))
	}

	a.mutex.Lock()
	defer a.mutex.Unlock()
	a.rawConfig = config
	a.notficationARN = config.Queue.TopicARN
	a.defaultRegion = config.DefaultRegion
	a.configured = true

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

// GetNotificationARN will return the notification arn for subs
func (a *awsClient) GetNotificationARN() string {
	return a.notficationARN
}

// GetDefaultRegion will return the static default
func (a *awsClient) GetDefaultRegion() string {
	return a.defaultRegion
}
