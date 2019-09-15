package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// CloudFormationMeta is the default CloudFormation spec metadata
type CloudFormationMeta struct {
	// +optional
	StackName string `json:"stackName,omitempty" description:"stack name for resource for recalling connected resources"`

	// +optional
	Region string `json:"region,omitempty" description:"region which the stack should be created in"`

	// +optional
	// NotificationARNs the Simple Notification Service (SNS) topic ARNs to
	// publish stack related events.
	NotificationARNs []*string `json:"notificationARNs,omitempty" description:"the Simple Notification Service (SNS) topic ARNs to publish stack related events"`

	// +optional
	// OnFailure determines what action will be taken if stack creation fails.
	// This must be one of: DO_NOTHING, ROLLBACK, or DELETE.
	OnFailure string `json:"onFailure,omitempty" description:"determines what action will be taken if stack creation fails"`

	// Tags key-value pairs to associate with this stack. AWS CloudFormation also
	// propagates these tags to the resources created in the stack.
	// +optional
	Tags map[string]string `json:"tags,omitempty" description:"key-value pairs to associate with this stack"`

	// +optional
	// TerminationProtection whether to enable termination protection on
	// the specified stack. If a user attempts to delete a stack with termination
	// protection enabled, the operation fails and the stack remains unchanged.
	TerminationProtection bool `json:"terminationProtection,omitempty" description:"whether to enable termination protection"`
}

// StatusInterface allows you to return your stack info without having know the object
type StatusInterface interface {
	GetStackID() string
}

// StatusMeta is the default CloudFormation status metadata
type StatusMeta struct {
	// ObservedGeneration is the version of the manifest which has been applied
	// +optional
	ObservedGeneration int64 `json:"generation,omitempty" description:"version of the manifest which has been applied"`

	// Status is the status of the condition
	// +optional
	Status string `json:"status,omitempty" description:"status of the condition, one of True, False, Unknown"`

	// +optional
	Message *string `json:"message,omitempty,omitempty" description:"human-readable message indicating details about last transition"`

	// +optional
	LastHeartbeatTime *metav1.Time `json:"lastHeartbeatTime,omitempty" description:"last time we got an update on a given condition"`

	// +optional
	LastTransitionTime *metav1.Time `json:"lastTransitionTime,omitempty" description:"last time the condition transit from one status to another"`

	// +optional
	StackID string `json:"stackID,omitempty" description:"CloudFormation Stack ID for referencing"`

	// +optional
	StackName string `json:"stackName,omitempty" description:"CloudFormation Stack Name for referencing"`
}

// FailureEvent describes what a stack does on failure
type FailureEvent string

const (
	// DoNothingFailureEvent When a stack fails don't do anything, this leaves the previously created resources
	DoNothingFailureEvent = "DO_NOTHING"

	// RollbackFailureEvent When a stack fails rollback all resources but leave the stack
	RollbackFailureEvent = "ROLLBACK"

	// DeleteFailureEvent When a stack fails delete the stack and all resources, DEFAULT action,
	// this allows the control loop to continue retrying automatically
	DeleteFailureEvent = "DELETE"
)

// ConditionStatus is the current condition status
type ConditionStatus string

const (
	// CreateCompleteStatus represents CREATE_COMPLETE CloudFormation status
	CreateCompleteStatus = "CreateComplete"

	// CreateInProgressStatus represents CREATE_IN_PROGRESS CloudFormation status
	CreateInProgressStatus = "CreateInProgress"

	// CreateFailedStatus represents CREATE_FAILED CloudFormation status
	CreateFailedStatus = "CreateFailed"

	// DeleteCompleteStatus represents DELETE_COMPLETE CloudFormation status
	DeleteCompleteStatus = "DeleteComplete"

	// DeleteFailedStatus represents DELETE_FAILED CloudFormation status
	DeleteFailedStatus = "DeleteFailed"

	// DeleteInProgressStatus represents DELETE_IN_PROGRESS CloudFormation status
	DeleteInProgressStatus = "DeleteInProgress"

	// ReviewInProgressStatus represents REVIEW_IN_PROGRESS CloudFormation status
	ReviewInProgressStatus = "ReviewInProgress"

	// RollbackCompleteStatus represents ROLLBACK_COMPLETE CloudFormation status
	RollbackCompleteStatus = "RollbackComplete"

	// RollbackFailedStatus represents ROLLBACK_FAILED CloudFormation status
	RollbackFailedStatus = "RollbackFailed"

	// RollbackInProgressStatus represents ROLLBACK_IN_PROGRESS CloudFormation status
	RollbackInProgressStatus = "RollbackInProgress"

	// UpdateCompleteStatus represents UPDATE_COMPLETE CloudFormation status
	UpdateCompleteStatus = "UpdateComplete"

	// UpdateCompleteCleanupInProgressStatus represents UPDATE_COMPLETE_CLEANUP_IN_PROGRESS CloudFormation status
	UpdateCompleteCleanupInProgressStatus = "UpdateCompleteCleanupInProgress"

	// UpdateInProgressStatus represents UPDATE_IN_PROGRESS CloudFormation status
	UpdateInProgressStatus = "UpdateInProgress"

	// UpdateRollbackCompleteStatus represents UPDATE_ROLLBACK_COMPLETE CloudFormation status
	UpdateRollbackCompleteStatus = "UpdateRollbackComplete"

	// UpdateRollbackCompleteCleanupInProgressStatus represents UPDATE_ROLLBACK_COMPLETE_CLEANUP_IN_PROGRESS CloudFormation status
	UpdateRollbackCompleteCleanupInProgressStatus = "UpdateRollbackCompleteCleanupInProgress"

	// UpdateRollbackFailedStatus represents UPDATE_ROLLBACK_FAILED CloudFormation status
	UpdateRollbackFailedStatus = "UpdateRollbackFailed"

	// UpdateRollbackInProgressStatus represents UPDATE_ROLLBACK_IN_PROGRESS CloudFormation status
	UpdateRollbackInProgressStatus = "UpdateRollbackInProgress"
)
