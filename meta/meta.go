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

// Package meta contains meta helpers
package meta

import (
	metav1alpha1 "awsctrl.io/apis/meta/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// StackObject defines defined functions for all stack objects
type StackObject interface {
	runtime.Object
	metav1.Object

	// GetNotificationARNs will return the list of Notifications
	GetNotificationARNs() []string

	// GetTemplate will return the JSON version of the CFN to use.
	GetTemplate() string

	// GenerateStackName will generate a stackName
	GenerateStackName() string

	// GetStackID will return stackID
	GetStackID() string

	// GetStackName will return stackName
	GetStackName() string

	// GetTemplateVersionLabel will return the template version label
	GetTemplateVersionLabel() (string, bool)

	// GetParameters will return the CFN Params
	GetParameters() map[string]string

	// GetCloudFormationMeta will return CFN meta object
	GetCloudFormationMeta() metav1alpha1.CloudFormationMeta

	// GetStatus will return the CFN Status
	GetStatus() metav1alpha1.ConditionStatus

	// SetStackID will put a stackID
	SetStackID(string)

	// SetStackName will put a stackName
	SetStackName(string)

	// SetTemplateVersionLabel will set the template version label
	SetTemplateVersionLabel()

	// TemplateVersionChanged will return if the template version has changed
	TemplateVersionChanged() bool

	// SetStatus will set status for object
	SetStatus(*metav1alpha1.StatusMeta)
}
