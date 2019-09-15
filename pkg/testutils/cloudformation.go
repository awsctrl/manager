package testutils

import (
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

func (m *mockCloudFormationClient) DescribeStacks(*cfn.DescribeStacksInput) (*cfn.DescribeStacksOutput, error) {
	describeStackOutput := &cfn.DescribeStacksOutput{}
	stack := &cfn.Stack{}
	output := &cfn.Output{}
	output.SetOutputKey("Name")
	output.SetOutputValue("test")
	stack.SetOutputs([]*cfn.Output{output})
	stack.SetStackStatus("UPDATE_IN_PROGRESS")
	stack.SetStackStatusReason("User initiated")
	stack.SetStackId(stackID)
	describeStackOutput.SetStacks([]*cfn.Stack{stack})
	return describeStackOutput, nil
}
