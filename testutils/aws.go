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

package testutils

import (
	"fmt"
	"sync"

	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/cloudformation/cloudformationiface"
	selfv1alpha1 "go.awsctrl.io/manager/apis/self/v1alpha1"
)

// NewAWS will initialize a mocked AWS Client
func NewAWS() *AWS {
	return &AWS{
		Config:  true,
		clients: map[string]cloudformationiface.CloudFormationAPI{"us-west-2": NewCFN("")},
	}
}

// AWS wraps all the test mocking for the AWS Client
type AWS struct {
	Config  bool
	clients map[string]cloudformationiface.CloudFormationAPI
	mux     sync.Mutex
}

// Configured will return nil to simulate a configured client
func (a *AWS) Configured() error {
	return nil
}

// Configure will setup the mocking client using the self.Config Custom Resource
func (a *AWS) Configure(*selfv1alpha1.ConfigAWS) error {
	if !a.Config {
		return fmt.Errorf("error occured")
	}
	return nil
}

// SetClient will add a new AWS Region client so that the controller can span more regions
func (a *AWS) SetClient(region string, iface cloudformationiface.CloudFormationAPI) bool {
	a.mux.Lock()
	defer a.mux.Unlock()
	a.clients = map[string]cloudformationiface.CloudFormationAPI{"us-west-2": iface}
	return true
}

// GetClients will return slice with a single AWS CloudFormation Client for us-west-2
func (a *AWS) GetClients() map[string]cloudformationiface.CloudFormationAPI {
	return a.clients
}

// GetClient will return a single AWS CloudFormation Client for us-west-2
func (a *AWS) GetClient(region string) cloudformationiface.CloudFormationAPI {
	return a.GetClients()["us-west-2"]
}

// SetSession will set a session by region
func (a *AWS) SetSession(region string, sess *session.Session) bool {
	return true
}

// GetSessions will return all sessions
func (a *AWS) GetSessions() map[string]*session.Session {
	return map[string]*session.Session{"us-west-2": &session.Session{}}
}

// GetSession will return a session for a region
func (a *AWS) GetSession(region string) *session.Session {
	return a.GetSessions()["us-west-2"]
}

// GetNotificationARN will return the SQS CloudFormation ARN
func (a *AWS) GetNotificationARN() string {
	return "arn:aws:sns::topic/awsctrl"
}

// GetDefaultRegion will return the base region to deploy into
func (a *AWS) GetDefaultRegion() string {
	return "us-west-2"
}
