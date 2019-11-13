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
	"github.com/aws/aws-sdk-go/service/cloudformation"
	cfn "github.com/aws/aws-sdk-go/service/cloudformation"
	"github.com/aws/aws-sdk-go/service/cloudformation/cloudformationiface"
)

var (
	stackID = "arn:aws:cloudformation:us-west-2:123456789012:stack/sample-default/60870DA4-CBEB-473D-AECA-51D3FB0DB72A"
)

type mockCloudFormationClient struct {
	cloudformationiface.CloudFormationAPI
}

func NewCFN() cloudformationiface.CloudFormationAPI {
	return &mockCloudFormationClient{}
}

func (m *mockCloudFormationClient) CreateStack(input *cfn.CreateStackInput) (*cfn.CreateStackOutput, error) {
	output := &cfn.CreateStackOutput{}
	output.SetStackId(stackID)
	return output, nil
}

func (m *mockCloudFormationClient) UpdateStack(input *cfn.UpdateStackInput) (*cfn.UpdateStackOutput, error) {
	output := &cfn.UpdateStackOutput{}
	output.SetStackId(stackID)
	return output, nil
}

func (m *mockCloudFormationClient) DescribeStacks(input *cfn.DescribeStacksInput) (*cfn.DescribeStacksOutput, error) {
	describeStackOutput := &cfn.DescribeStacksOutput{}
	stack := &cfn.Stack{}
	output := &cfn.Output{}
	output.SetOutputKey("Name")
	output.SetOutputValue("test")
	stack.SetOutputs([]*cfn.Output{output})
	stack.SetStackStatus("UPDATE_COMPLETE")
	stack.SetStackStatusReason("User initiated")
	stack.SetStackId(stackID)
	describeStackOutput.SetStacks([]*cfn.Stack{stack})
	return describeStackOutput, nil
}

func (m *mockCloudFormationClient) DeleteStack(input *cloudformation.DeleteStackInput) (*cloudformation.DeleteStackOutput, error) {
	return &cloudformation.DeleteStackOutput{}, nil
}

func (m *mockCloudFormationClient) WaitUntilStackDeleteComplete(input *cloudformation.DescribeStacksInput) error {
	return nil
}
