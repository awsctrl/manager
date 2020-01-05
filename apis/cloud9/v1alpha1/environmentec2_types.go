/*
Copyright © 2019 AWS Controller authors

Licensed under the Apache License, Version 2.0 (the &#34;License&#34;);
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an &#34;AS IS&#34; BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package v1alpha1

import (
	metav1alpha1 "go.awsctrl.io/manager/apis/meta/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EnvironmentEC2Spec defines the desired state of EnvironmentEC2
type EnvironmentEC2Spec struct {
	metav1alpha1.CloudFormationMeta `json:",inline"`

	// AutomaticStopTimeMinutes http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cloud9-environmentec2.html#cfn-cloud9-environmentec2-automaticstoptimeminutes
	AutomaticStopTimeMinutes int `json:"automaticStopTimeMinutes,omitempty" cloudformation:"AutomaticStopTimeMinutes,Parameter"`

	// Subnet http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cloud9-environmentec2.html#cfn-cloud9-environmentec2-subnetid
	Subnet metav1alpha1.ObjectReference `json:"subnet,omitempty" cloudformation:"SubnetId,Parameter"`

	// InstanceType http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cloud9-environmentec2.html#cfn-cloud9-environmentec2-instancetype
	InstanceType string `json:"instanceType" cloudformation:"InstanceType,Parameter"`

	// Name http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cloud9-environmentec2.html#cfn-cloud9-environmentec2-name
	Name string `json:"name,omitempty" cloudformation:"Name,Parameter"`

	// Repositories http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cloud9-environmentec2.html#cfn-cloud9-environmentec2-repositories
	Repositories []EnvironmentEC2_Repository `json:"repositories,omitempty" cloudformation:"Repositories"`

	// Owner http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cloud9-environmentec2.html#cfn-cloud9-environmentec2-ownerarn
	Owner metav1alpha1.ObjectReference `json:"owner,omitempty" cloudformation:"OwnerArn,Parameter"`

	// Description http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cloud9-environmentec2.html#cfn-cloud9-environmentec2-description
	Description string `json:"description,omitempty" cloudformation:"Description,Parameter"`
}

// EnvironmentEC2_Repository defines the desired state of EnvironmentEC2Repository
type EnvironmentEC2_Repository struct {
	// PathComponent http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloud9-environmentec2-repository.html#cfn-cloud9-environmentec2-repository-pathcomponent
	PathComponent string `json:"pathComponent" cloudformation:"PathComponent,Parameter"`

	// RepositoryUrl http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloud9-environmentec2-repository.html#cfn-cloud9-environmentec2-repository-repositoryurl
	RepositoryUrl string `json:"repositoryUrl" cloudformation:"RepositoryUrl,Parameter"`
}

// EnvironmentEC2Status defines the observed state of EnvironmentEC2
type EnvironmentEC2Status struct {
	metav1alpha1.StatusMeta `json:",inline"`
}

// EnvironmentEC2Output defines the stack outputs
type EnvironmentEC2Output struct {
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cloud9-environmentec2.html
	Ref string `json:"ref,omitempty"`

	// Arn defines the Arn
	Arn string `json:"arn,omitempty" cloudformation:"Arn,Output"`

	// Name defines the Name
	Name string `json:"name,omitempty" cloudformation:"Name,Output"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:resource:categories=aws;cloud9
// +kubebuilder:printcolumn:JSONPath=.status.status,description="status of the stack",name=Status,priority=0,type=string
// +kubebuilder:printcolumn:JSONPath=.status.message,description="reason for the stack status",name=Message,priority=1,type=string
// +kubebuilder:printcolumn:JSONPath=.status.stackID,description="CloudFormation Stack ID",name=StackID,priority=2,type=string

// EnvironmentEC2 is the Schema for the cloud9 EnvironmentEC2 API
type EnvironmentEC2 struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   EnvironmentEC2Spec   `json:"spec,omitempty"`
	Status EnvironmentEC2Status `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// EnvironmentEC2List contains a list of Account
type EnvironmentEC2List struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`

	Items []EnvironmentEC2 `json:"items"`
}

func init() {
	SchemeBuilder.Register(&EnvironmentEC2{}, &EnvironmentEC2List{})
}
