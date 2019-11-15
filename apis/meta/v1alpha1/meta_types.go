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

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// CloudFormationMeta is the default CloudFormation spec metadata
type CloudFormationMeta struct {
	// +optional
	StackName string `json:"stackName,omitempty"`

	// +optional
	Region string `json:"region,omitempty"`

	// +optional
	// NotificationARNs the Simple Notification Service (SNS) topic ARNs to
	// publish stack related events.
	NotificationARNs []*string `json:"notificationARNs,omitempty"`

	// +optional
	// OnFailure determines what action will be taken if stack creation fails.
	// This must be one of: DO_NOTHING, ROLLBACK, or DELETE.
	OnFailure string `json:"onFailure,omitempty"`

	// Tags key-value pairs to associate with this stack. AWS CloudFormation also
	// propagates these tags to the resources created in the stack.
	// +optional
	Tags map[string]string `json:"tags,omitempty"`

	// TerminationProtection whether to enable termination protection on
	// the specified stack. If a user attempts to delete a stack with termination
	// protection enabled, the operation fails and the stack remains unchanged.
	// +optional
	TerminationProtection bool `json:"terminationProtection,omitempty"`
}

// StatusMeta is the default CloudFormation status metadata
type StatusMeta struct {
	// ObservedGeneration is the version of the manifest which has been applied
	// +optional
	ObservedGeneration int64 `json:"generation,omitempty"`

	// Status is the status of the condition
	// +optional
	Status ConditionStatus `json:"status,omitempty"`

	// +optional
	Message *string `json:"message,omitempty,omitempty"`

	// +optional
	LastHeartbeatTime *metav1.Time `json:"lastHeartbeatTime,omitempty"`

	// +optional
	LastTransitionTime *metav1.Time `json:"lastTransitionTime,omitempty"`

	// +optional
	StackID string `json:"stackID,omitempty"`

	// +optional
	StackName string `json:"stackName,omitempty"`

	// +optional
	Outputs map[string]string `json:"outputs,omitempty"`
}

// FailureEvent describes what a stack does on failure
type FailureEvent string

const (
	// DoNothingFailureEvent When a stack fails don't do anything, this leaves the previously created resources
	DoNothingFailureEvent FailureEvent = "DO_NOTHING"

	// RollbackFailureEvent When a stack fails rollback all resources but leave the stack
	RollbackFailureEvent FailureEvent = "ROLLBACK"

	// DeleteFailureEvent When a stack fails delete the stack and all resources, DEFAULT action,
	// this allows the control loop to continue retrying automatically
	DeleteFailureEvent FailureEvent = "DELETE"
)

// ConditionStatus is the current condition status
type ConditionStatus string

const (
	// CreateCompleteStatus represents CREATE_COMPLETE CloudFormation status
	CreateCompleteStatus ConditionStatus = "CreateComplete"

	// CreateInProgressStatus represents CREATE_IN_PROGRESS CloudFormation status
	CreateInProgressStatus ConditionStatus = "CreateInProgress"

	// CreateFailedStatus represents CREATE_FAILED CloudFormation status
	CreateFailedStatus ConditionStatus = "CreateFailed"

	// DeleteCompleteStatus represents DELETE_COMPLETE CloudFormation status
	DeleteCompleteStatus ConditionStatus = "DeleteComplete"

	// DeleteFailedStatus represents DELETE_FAILED CloudFormation status
	DeleteFailedStatus ConditionStatus = "DeleteFailed"

	// DeleteInProgressStatus represents DELETE_IN_PROGRESS CloudFormation status
	DeleteInProgressStatus ConditionStatus = "DeleteInProgress"

	// ReviewInProgressStatus represents REVIEW_IN_PROGRESS CloudFormation status
	ReviewInProgressStatus ConditionStatus = "ReviewInProgress"

	// RollbackCompleteStatus represents ROLLBACK_COMPLETE CloudFormation status
	RollbackCompleteStatus ConditionStatus = "RollbackComplete"

	// RollbackFailedStatus represents ROLLBACK_FAILED CloudFormation status
	RollbackFailedStatus ConditionStatus = "RollbackFailed"

	// RollbackInProgressStatus represents ROLLBACK_IN_PROGRESS CloudFormation status
	RollbackInProgressStatus ConditionStatus = "RollbackInProgress"

	// UpdateCompleteStatus represents UPDATE_COMPLETE CloudFormation status
	UpdateCompleteStatus ConditionStatus = "UpdateComplete"

	// UpdateCompleteCleanupInProgressStatus represents UPDATE_COMPLETE_CLEANUP_IN_PROGRESS CloudFormation status
	UpdateCompleteCleanupInProgressStatus ConditionStatus = "UpdateCompleteCleanupInProgress"

	// UpdateInProgressStatus represents UPDATE_IN_PROGRESS CloudFormation status
	UpdateInProgressStatus ConditionStatus = "UpdateInProgress"

	// UpdateRollbackCompleteStatus represents UPDATE_ROLLBACK_COMPLETE CloudFormation status
	UpdateRollbackCompleteStatus ConditionStatus = "UpdateRollbackComplete"

	// UpdateRollbackCompleteCleanupInProgressStatus represents UPDATE_ROLLBACK_COMPLETE_CLEANUP_IN_PROGRESS CloudFormation status
	UpdateRollbackCompleteCleanupInProgressStatus ConditionStatus = "UpdateRollbackCompleteCleanupInProgress"

	// UpdateRollbackFailedStatus represents UPDATE_ROLLBACK_FAILED CloudFormation status
	UpdateRollbackFailedStatus ConditionStatus = "UpdateRollbackFailed"

	// UpdateRollbackInProgressStatus represents UPDATE_ROLLBACK_IN_PROGRESS CloudFormation status
	UpdateRollbackInProgressStatus ConditionStatus = "UpdateRollbackInProgress"
)
