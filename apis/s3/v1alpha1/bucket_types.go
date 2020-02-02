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

// BucketSpec defines the desired state of Bucket
type BucketSpec struct {
	metav1alpha1.CloudFormationMeta `json:",inline"`

	// ObjectLockConfiguration http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket.html#cfn-s3-bucket-objectlockconfiguration
	ObjectLockConfiguration Bucket_ObjectLockConfiguration `json:"objectLockConfiguration,omitempty" cloudformation:"ObjectLockConfiguration"`

	// AccessControl http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket.html#cfn-s3-bucket-accesscontrol
	AccessControl string `json:"accessControl,omitempty" cloudformation:"AccessControl,Parameter"`

	// AccelerateConfiguration http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket.html#cfn-s3-bucket-accelerateconfiguration
	AccelerateConfiguration Bucket_AccelerateConfiguration `json:"accelerateConfiguration,omitempty" cloudformation:"AccelerateConfiguration"`

	// CorsConfiguration http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket.html#cfn-s3-bucket-crossoriginconfig
	CorsConfiguration Bucket_CorsConfiguration `json:"corsConfiguration,omitempty" cloudformation:"CorsConfiguration"`

	// MetricsConfigurations http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket.html#cfn-s3-bucket-metricsconfigurations
	MetricsConfigurations []Bucket_MetricsConfiguration `json:"metricsConfigurations,omitempty" cloudformation:"MetricsConfigurations"`

	// WebsiteConfiguration http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket.html#cfn-s3-bucket-websiteconfiguration
	WebsiteConfiguration Bucket_WebsiteConfiguration `json:"websiteConfiguration,omitempty" cloudformation:"WebsiteConfiguration"`

	// BucketEncryption http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket.html#cfn-s3-bucket-bucketencryption
	BucketEncryption Bucket_BucketEncryption `json:"bucketEncryption,omitempty" cloudformation:"BucketEncryption"`

	// LoggingConfiguration http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket.html#cfn-s3-bucket-loggingconfig
	LoggingConfiguration Bucket_LoggingConfiguration `json:"loggingConfiguration,omitempty" cloudformation:"LoggingConfiguration"`

	// AnalyticsConfigurations http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket.html#cfn-s3-bucket-analyticsconfigurations
	AnalyticsConfigurations []Bucket_AnalyticsConfiguration `json:"analyticsConfigurations,omitempty" cloudformation:"AnalyticsConfigurations"`

	// LifecycleConfiguration http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket.html#cfn-s3-bucket-lifecycleconfig
	LifecycleConfiguration Bucket_LifecycleConfiguration `json:"lifecycleConfiguration,omitempty" cloudformation:"LifecycleConfiguration"`

	// BucketName http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket.html#cfn-s3-bucket-name
	BucketName string `json:"bucketName,omitempty" cloudformation:"BucketName,Parameter"`

	// InventoryConfigurations http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket.html#cfn-s3-bucket-inventoryconfigurations
	InventoryConfigurations []Bucket_InventoryConfiguration `json:"inventoryConfigurations,omitempty" cloudformation:"InventoryConfigurations"`

	// NotificationConfiguration http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket.html#cfn-s3-bucket-notification
	NotificationConfiguration Bucket_NotificationConfiguration `json:"notificationConfiguration,omitempty" cloudformation:"NotificationConfiguration"`

	// ObjectLockEnabled http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket.html#cfn-s3-bucket-objectlockenabled
	ObjectLockEnabled bool `json:"objectLockEnabled,omitempty" cloudformation:"ObjectLockEnabled,Parameter"`

	// PublicAccessBlockConfiguration http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket.html#cfn-s3-bucket-publicaccessblockconfiguration
	PublicAccessBlockConfiguration Bucket_PublicAccessBlockConfiguration `json:"publicAccessBlockConfiguration,omitempty" cloudformation:"PublicAccessBlockConfiguration"`

	// VersioningConfiguration http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket.html#cfn-s3-bucket-versioning
	VersioningConfiguration Bucket_VersioningConfiguration `json:"versioningConfiguration,omitempty" cloudformation:"VersioningConfiguration"`

	// ReplicationConfiguration http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket.html#cfn-s3-bucket-replicationconfiguration
	ReplicationConfiguration Bucket_ReplicationConfiguration `json:"replicationConfiguration,omitempty" cloudformation:"ReplicationConfiguration"`
}

// Bucket_EncryptionConfiguration defines the desired state of BucketEncryptionConfiguration
type Bucket_EncryptionConfiguration struct {
	// ReplicaKmsKeyRef http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket-encryptionconfiguration.html#cfn-s3-bucket-encryptionconfiguration-replicakmskeyid
	ReplicaKmsKeyRef metav1alpha1.ObjectReference `json:"replicaKmsKeyRef,omitempty" cloudformation:"ReplicaKmsKeyID,Parameter"`
}

// Bucket_ReplicationRule defines the desired state of BucketReplicationRule
type Bucket_ReplicationRule struct {
	// Prefix http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket-replicationconfiguration-rules.html#cfn-s3-bucket-replicationconfiguration-rules-prefix
	Prefix string `json:"prefix,omitempty" cloudformation:"Prefix,Parameter"`

	// SourceSelectionCriteria http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket-replicationconfiguration-rules.html#cfn-s3-bucket-replicationrule-sourceselectioncriteria
	SourceSelectionCriteria Bucket_SourceSelectionCriteria `json:"sourceSelectionCriteria,omitempty" cloudformation:"SourceSelectionCriteria"`

	// Status http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket-replicationconfiguration-rules.html#cfn-s3-bucket-replicationconfiguration-rules-status
	Status string `json:"status,omitempty" cloudformation:"Status,Parameter"`

	// Destination http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket-replicationconfiguration-rules.html#cfn-s3-bucket-replicationconfiguration-rules-destination
	Destination Bucket_ReplicationDestination `json:"destination,omitempty" cloudformation:"Destination"`

	// Ref http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket-replicationconfiguration-rules.html#cfn-s3-bucket-replicationconfiguration-rules-id
	Ref metav1alpha1.ObjectReference `json:"ref,omitempty" cloudformation:"Id,Parameter"`
}

// Bucket_Rule defines the desired state of BucketRule
type Bucket_Rule struct {
	// NoncurrentVersionTransition http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket-lifecycleconfig-rule.html#cfn-s3-bucket-lifecycleconfig-rule-noncurrentversiontransition
	NoncurrentVersionTransition Bucket_NoncurrentVersionTransition `json:"noncurrentVersionTransition,omitempty" cloudformation:"NoncurrentVersionTransition"`

	// NoncurrentVersionTransitions http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket-lifecycleconfig-rule.html#cfn-s3-bucket-lifecycleconfig-rule-noncurrentversiontransitions
	NoncurrentVersionTransitions []Bucket_NoncurrentVersionTransition `json:"noncurrentVersionTransitions,omitempty" cloudformation:"NoncurrentVersionTransitions"`

	// Status http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket-lifecycleconfig-rule.html#cfn-s3-bucket-lifecycleconfig-rule-status
	Status string `json:"status,omitempty" cloudformation:"Status,Parameter"`

	// Ref http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket-lifecycleconfig-rule.html#cfn-s3-bucket-lifecycleconfig-rule-id
	Ref metav1alpha1.ObjectReference `json:"ref,omitempty" cloudformation:"Id,Parameter"`

	// Transition http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket-lifecycleconfig-rule.html#cfn-s3-bucket-lifecycleconfig-rule-transition
	Transition Bucket_Transition `json:"transition,omitempty" cloudformation:"Transition"`

	// AbortIncompleteMultipartUpload http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket-lifecycleconfig-rule.html#cfn-s3-bucket-rule-abortincompletemultipartupload
	AbortIncompleteMultipartUpload Bucket_AbortIncompleteMultipartUpload `json:"abortIncompleteMultipartUpload,omitempty" cloudformation:"AbortIncompleteMultipartUpload"`

	// ExpirationInDays http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket-lifecycleconfig-rule.html#cfn-s3-bucket-lifecycleconfig-rule-expirationindays
	ExpirationInDays int `json:"expirationInDays,omitempty" cloudformation:"ExpirationInDays,Parameter"`

	// NoncurrentVersionExpirationInDays http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket-lifecycleconfig-rule.html#cfn-s3-bucket-lifecycleconfig-rule-noncurrentversionexpirationindays
	NoncurrentVersionExpirationInDays int `json:"noncurrentVersionExpirationInDays,omitempty" cloudformation:"NoncurrentVersionExpirationInDays,Parameter"`

	// TagFilters http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket-lifecycleconfig-rule.html#cfn-s3-bucket-rule-tagfilters
	TagFilters []Bucket_TagFilter `json:"tagFilters,omitempty" cloudformation:"TagFilters"`

	// Transitions http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket-lifecycleconfig-rule.html#cfn-s3-bucket-lifecycleconfig-rule-transitions
	Transitions []Bucket_Transition `json:"transitions,omitempty" cloudformation:"Transitions"`

	// ExpirationDate http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket-lifecycleconfig-rule.html#cfn-s3-bucket-lifecycleconfig-rule-expirationdate
	ExpirationDate string `json:"expirationDate,omitempty" cloudformation:"ExpirationDate,Parameter"`

	// Prefix http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket-lifecycleconfig-rule.html#cfn-s3-bucket-lifecycleconfig-rule-prefix
	Prefix string `json:"prefix,omitempty" cloudformation:"Prefix,Parameter"`
}

// Bucket_TopicConfiguration defines the desired state of BucketTopicConfiguration
type Bucket_TopicConfiguration struct {
	// Event http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket-notificationconfig-topicconfig.html#cfn-s3-bucket-notificationconfig-topicconfig-event
	Event string `json:"event,omitempty" cloudformation:"Event,Parameter"`

	// Filter http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket-notificationconfig-topicconfig.html#cfn-s3-bucket-notificationconfig-topicconfig-filter
	Filter Bucket_NotificationFilter `json:"filter,omitempty" cloudformation:"Filter"`

	// Topic http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket-notificationconfig-topicconfig.html#cfn-s3-bucket-notificationconfig-topicconfig-topic
	Topic string `json:"topic,omitempty" cloudformation:"Topic,Parameter"`
}

// Bucket_NoncurrentVersionTransition defines the desired state of BucketNoncurrentVersionTransition
type Bucket_NoncurrentVersionTransition struct {
	// StorageClass http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket-lifecycleconfig-rule-noncurrentversiontransition.html#cfn-s3-bucket-lifecycleconfig-rule-noncurrentversiontransition-storageclass
	StorageClass string `json:"storageClass,omitempty" cloudformation:"StorageClass,Parameter"`

	// TransitionInDays http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket-lifecycleconfig-rule-noncurrentversiontransition.html#cfn-s3-bucket-lifecycleconfig-rule-noncurrentversiontransition-transitionindays
	TransitionInDays int `json:"transitionInDays,omitempty" cloudformation:"TransitionInDays,Parameter"`
}

// Bucket_CorsRule defines the desired state of BucketCorsRule
type Bucket_CorsRule struct {
	// AllowedOrigins http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket-cors-corsrule.html#cfn-s3-bucket-cors-corsrule-allowedorigins
	AllowedOrigins []string `json:"allowedOrigins,omitempty" cloudformation:"AllowedOrigins"`

	// ExposedHeaders http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket-cors-corsrule.html#cfn-s3-bucket-cors-corsrule-exposedheaders
	ExposedHeaders []string `json:"exposedHeaders,omitempty" cloudformation:"ExposedHeaders"`

	// Ref http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket-cors-corsrule.html#cfn-s3-bucket-cors-corsrule-id
	Ref metav1alpha1.ObjectReference `json:"ref,omitempty" cloudformation:"Id,Parameter"`

	// MaxAge http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket-cors-corsrule.html#cfn-s3-bucket-cors-corsrule-maxage
	MaxAge int `json:"maxAge,omitempty" cloudformation:"MaxAge,Parameter"`

	// AllowedHeaders http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket-cors-corsrule.html#cfn-s3-bucket-cors-corsrule-allowedheaders
	AllowedHeaders []string `json:"allowedHeaders,omitempty" cloudformation:"AllowedHeaders"`

	// AllowedMethods http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket-cors-corsrule.html#cfn-s3-bucket-cors-corsrule-allowedmethods
	AllowedMethods []string `json:"allowedMethods,omitempty" cloudformation:"AllowedMethods"`
}

// Bucket_MetricsConfiguration defines the desired state of BucketMetricsConfiguration
type Bucket_MetricsConfiguration struct {
	// Ref http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket-metricsconfiguration.html#cfn-s3-bucket-metricsconfiguration-id
	Ref metav1alpha1.ObjectReference `json:"ref,omitempty" cloudformation:"Id,Parameter"`

	// Prefix http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket-metricsconfiguration.html#cfn-s3-bucket-metricsconfiguration-prefix
	Prefix string `json:"prefix,omitempty" cloudformation:"Prefix,Parameter"`

	// TagFilters http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket-metricsconfiguration.html#cfn-s3-bucket-metricsconfiguration-tagfilters
	TagFilters []Bucket_TagFilter `json:"tagFilters,omitempty" cloudformation:"TagFilters"`
}

// Bucket_ObjectLockConfiguration defines the desired state of BucketObjectLockConfiguration
type Bucket_ObjectLockConfiguration struct {
	// ObjectLockEnabled http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket-objectlockconfiguration.html#cfn-s3-bucket-objectlockconfiguration-objectlockenabled
	ObjectLockEnabled string `json:"objectLockEnabled,omitempty" cloudformation:"ObjectLockEnabled,Parameter"`

	// Rule http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket-objectlockconfiguration.html#cfn-s3-bucket-objectlockconfiguration-rule
	Rule Bucket_ObjectLockRule `json:"rule,omitempty" cloudformation:"Rule"`
}

// Bucket_ObjectLockRule defines the desired state of BucketObjectLockRule
type Bucket_ObjectLockRule struct {
	// DefaultRetention http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket-objectlockrule.html#cfn-s3-bucket-objectlockrule-defaultretention
	DefaultRetention Bucket_DefaultRetention `json:"defaultRetention,omitempty" cloudformation:"DefaultRetention"`
}

// Bucket_NotificationConfiguration defines the desired state of BucketNotificationConfiguration
type Bucket_NotificationConfiguration struct {
	// TopicConfigurations http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket-notificationconfig.html#cfn-s3-bucket-notificationconfig-topicconfig
	TopicConfigurations []Bucket_TopicConfiguration `json:"topicConfigurations,omitempty" cloudformation:"TopicConfigurations"`

	// LambdaConfigurations http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket-notificationconfig.html#cfn-s3-bucket-notificationconfig-lambdaconfig
	LambdaConfigurations []Bucket_LambdaConfiguration `json:"lambdaConfigurations,omitempty" cloudformation:"LambdaConfigurations"`

	// QueueConfigurations http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket-notificationconfig.html#cfn-s3-bucket-notificationconfig-queueconfig
	QueueConfigurations []Bucket_QueueConfiguration `json:"queueConfigurations,omitempty" cloudformation:"QueueConfigurations"`
}

// Bucket_ReplicationDestination defines the desired state of BucketReplicationDestination
type Bucket_ReplicationDestination struct {
	// Account http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket-replicationconfiguration-rules-destination.html#cfn-s3-bucket-replicationdestination-account
	Account string `json:"account,omitempty" cloudformation:"Account,Parameter"`

	// Bucket http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket-replicationconfiguration-rules-destination.html#cfn-s3-bucket-replicationconfiguration-rules-destination-bucket
	Bucket string `json:"bucket,omitempty" cloudformation:"Bucket,Parameter"`

	// EncryptionConfiguration http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket-replicationconfiguration-rules-destination.html#cfn-s3-bucket-replicationdestination-encryptionconfiguration
	EncryptionConfiguration Bucket_EncryptionConfiguration `json:"encryptionConfiguration,omitempty" cloudformation:"EncryptionConfiguration"`

	// StorageClass http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket-replicationconfiguration-rules-destination.html#cfn-s3-bucket-replicationconfiguration-rules-destination-storageclass
	StorageClass string `json:"storageClass,omitempty" cloudformation:"StorageClass,Parameter"`

	// AccessControlTranslation http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket-replicationconfiguration-rules-destination.html#cfn-s3-bucket-replicationdestination-accesscontroltranslation
	AccessControlTranslation Bucket_AccessControlTranslation `json:"accessControlTranslation,omitempty" cloudformation:"AccessControlTranslation"`
}

// Bucket_TagFilter defines the desired state of BucketTagFilter
type Bucket_TagFilter struct {
	// Key http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket-tagfilter.html#cfn-s3-bucket-tagfilter-key
	Key string `json:"key,omitempty" cloudformation:"Key,Parameter"`

	// Value http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket-tagfilter.html#cfn-s3-bucket-tagfilter-value
	Value string `json:"value,omitempty" cloudformation:"Value,Parameter"`
}

// Bucket_SseKmsEncryptedObjects defines the desired state of BucketSseKmsEncryptedObjects
type Bucket_SseKmsEncryptedObjects struct {
	// Status http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket-ssekmsencryptedobjects.html#cfn-s3-bucket-ssekmsencryptedobjects-status
	Status string `json:"status,omitempty" cloudformation:"Status,Parameter"`
}

// Bucket_CorsConfiguration defines the desired state of BucketCorsConfiguration
type Bucket_CorsConfiguration struct {
	// CorsRules http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket-cors.html#cfn-s3-bucket-cors-corsrule
	CorsRules []Bucket_CorsRule `json:"corsRules,omitempty" cloudformation:"CorsRules"`
}

// Bucket_VersioningConfiguration defines the desired state of BucketVersioningConfiguration
type Bucket_VersioningConfiguration struct {
	// Status http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket-versioningconfig.html#cfn-s3-bucket-versioningconfig-status
	Status string `json:"status,omitempty" cloudformation:"Status,Parameter"`
}

// Bucket_RedirectAllRequestsTo defines the desired state of BucketRedirectAllRequestsTo
type Bucket_RedirectAllRequestsTo struct {
	// Protocol http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-websiteconfiguration-redirectallrequeststo.html#cfn-s3-websiteconfiguration-redirectallrequeststo-protocol
	Protocol string `json:"protocol,omitempty" cloudformation:"Protocol,Parameter"`

	// HostName http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-websiteconfiguration-redirectallrequeststo.html#cfn-s3-websiteconfiguration-redirectallrequeststo-hostname
	HostName string `json:"hostName,omitempty" cloudformation:"HostName,Parameter"`
}

// Bucket_AnalyticsConfiguration defines the desired state of BucketAnalyticsConfiguration
type Bucket_AnalyticsConfiguration struct {
	// Ref http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket-analyticsconfiguration.html#cfn-s3-bucket-analyticsconfiguration-id
	Ref metav1alpha1.ObjectReference `json:"ref,omitempty" cloudformation:"Id,Parameter"`

	// Prefix http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket-analyticsconfiguration.html#cfn-s3-bucket-analyticsconfiguration-prefix
	Prefix string `json:"prefix,omitempty" cloudformation:"Prefix,Parameter"`

	// StorageClassAnalysis http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket-analyticsconfiguration.html#cfn-s3-bucket-analyticsconfiguration-storageclassanalysis
	StorageClassAnalysis Bucket_StorageClassAnalysis `json:"storageClassAnalysis,omitempty" cloudformation:"StorageClassAnalysis"`

	// TagFilters http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket-analyticsconfiguration.html#cfn-s3-bucket-analyticsconfiguration-tagfilters
	TagFilters []Bucket_TagFilter `json:"tagFilters,omitempty" cloudformation:"TagFilters"`
}

// Bucket_StorageClassAnalysis defines the desired state of BucketStorageClassAnalysis
type Bucket_StorageClassAnalysis struct {
	// DataExport http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket-storageclassanalysis.html#cfn-s3-bucket-storageclassanalysis-dataexport
	DataExport Bucket_DataExport `json:"dataExport,omitempty" cloudformation:"DataExport"`
}

// Bucket_AbortIncompleteMultipartUpload defines the desired state of BucketAbortIncompleteMultipartUpload
type Bucket_AbortIncompleteMultipartUpload struct {
	// DaysAfterInitiation http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket-abortincompletemultipartupload.html#cfn-s3-bucket-abortincompletemultipartupload-daysafterinitiation
	DaysAfterInitiation int `json:"daysAfterInitiation,omitempty" cloudformation:"DaysAfterInitiation,Parameter"`
}

// Bucket_ServerSideEncryptionByDefault defines the desired state of BucketServerSideEncryptionByDefault
type Bucket_ServerSideEncryptionByDefault struct {
	// SSEAlgorithm http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket-serversideencryptionbydefault.html#cfn-s3-bucket-serversideencryptionbydefault-ssealgorithm
	SSEAlgorithm string `json:"sSEAlgorithm,omitempty" cloudformation:"SSEAlgorithm,Parameter"`

	// KMSMasterKeyRef http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket-serversideencryptionbydefault.html#cfn-s3-bucket-serversideencryptionbydefault-kmsmasterkeyid
	KMSMasterKeyRef metav1alpha1.ObjectReference `json:"kMSMasterKeyRef,omitempty" cloudformation:"KMSMasterKeyID,Parameter"`
}

// Bucket_RedirectRule defines the desired state of BucketRedirectRule
type Bucket_RedirectRule struct {
	// Protocol http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-websiteconfiguration-routingrules-redirectrule.html#cfn-s3-websiteconfiguration-redirectrule-protocol
	Protocol string `json:"protocol,omitempty" cloudformation:"Protocol,Parameter"`

	// ReplaceKeyPrefixWith http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-websiteconfiguration-routingrules-redirectrule.html#cfn-s3-websiteconfiguration-redirectrule-replacekeyprefixwith
	ReplaceKeyPrefixWith string `json:"replaceKeyPrefixWith,omitempty" cloudformation:"ReplaceKeyPrefixWith,Parameter"`

	// ReplaceKeyWith http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-websiteconfiguration-routingrules-redirectrule.html#cfn-s3-websiteconfiguration-redirectrule-replacekeywith
	ReplaceKeyWith string `json:"replaceKeyWith,omitempty" cloudformation:"ReplaceKeyWith,Parameter"`

	// HostName http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-websiteconfiguration-routingrules-redirectrule.html#cfn-s3-websiteconfiguration-redirectrule-hostname
	HostName string `json:"hostName,omitempty" cloudformation:"HostName,Parameter"`

	// HttpRedirectCode http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-websiteconfiguration-routingrules-redirectrule.html#cfn-s3-websiteconfiguration-redirectrule-httpredirectcode
	HttpRedirectCode string `json:"httpRedirectCode,omitempty" cloudformation:"HttpRedirectCode,Parameter"`
}

// Bucket_Destination defines the desired state of BucketDestination
type Bucket_Destination struct {
	// BucketAccountRef http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket-destination.html#cfn-s3-bucket-destination-bucketaccountid
	BucketAccountRef metav1alpha1.ObjectReference `json:"bucketAccountRef,omitempty" cloudformation:"BucketAccountId,Parameter"`

	// BucketRef http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket-destination.html#cfn-s3-bucket-destination-bucketarn
	BucketRef metav1alpha1.ObjectReference `json:"bucketRef,omitempty" cloudformation:"BucketArn,Parameter"`

	// Format http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket-destination.html#cfn-s3-bucket-destination-format
	Format string `json:"format,omitempty" cloudformation:"Format,Parameter"`

	// Prefix http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket-destination.html#cfn-s3-bucket-destination-prefix
	Prefix string `json:"prefix,omitempty" cloudformation:"Prefix,Parameter"`
}

// Bucket_QueueConfiguration defines the desired state of BucketQueueConfiguration
type Bucket_QueueConfiguration struct {
	// Queue http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket-notificationconfig-queueconfig.html#cfn-s3-bucket-notificationconfig-queueconfig-queue
	Queue string `json:"queue,omitempty" cloudformation:"Queue,Parameter"`

	// Event http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket-notificationconfig-queueconfig.html#cfn-s3-bucket-notificationconfig-queueconfig-event
	Event string `json:"event,omitempty" cloudformation:"Event,Parameter"`

	// Filter http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket-notificationconfig-queueconfig.html#cfn-s3-bucket-notificationconfig-queueconfig-filter
	Filter Bucket_NotificationFilter `json:"filter,omitempty" cloudformation:"Filter"`
}

// Bucket_NotificationFilter defines the desired state of BucketNotificationFilter
type Bucket_NotificationFilter struct {
	// S3Key http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket-notificationconfiguration-config-filter.html#cfn-s3-bucket-notificationconfiguraiton-config-filter-s3key
	S3Key Bucket_S3KeyFilter `json:"s3Key,omitempty" cloudformation:"S3Key"`
}

// Bucket_WebsiteConfiguration defines the desired state of BucketWebsiteConfiguration
type Bucket_WebsiteConfiguration struct {
	// IndexDocument http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-websiteconfiguration.html#cfn-s3-websiteconfiguration-indexdocument
	IndexDocument string `json:"indexDocument,omitempty" cloudformation:"IndexDocument,Parameter"`

	// RedirectAllRequestsTo http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-websiteconfiguration.html#cfn-s3-websiteconfiguration-redirectallrequeststo
	RedirectAllRequestsTo Bucket_RedirectAllRequestsTo `json:"redirectAllRequestsTo,omitempty" cloudformation:"RedirectAllRequestsTo"`

	// RoutingRules http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-websiteconfiguration.html#cfn-s3-websiteconfiguration-routingrules
	RoutingRules []Bucket_RoutingRule `json:"routingRules,omitempty" cloudformation:"RoutingRules"`

	// ErrorDocument http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-websiteconfiguration.html#cfn-s3-websiteconfiguration-errordocument
	ErrorDocument string `json:"errorDocument,omitempty" cloudformation:"ErrorDocument,Parameter"`
}

// Bucket_FilterRule defines the desired state of BucketFilterRule
type Bucket_FilterRule struct {
	// Name http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket-notificationconfiguration-config-filter-s3key-rules.html#cfn-s3-bucket-notificationconfiguraiton-config-filter-s3key-rules-name
	Name string `json:"name,omitempty" cloudformation:"Name,Parameter"`

	// Value http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket-notificationconfiguration-config-filter-s3key-rules.html#cfn-s3-bucket-notificationconfiguraiton-config-filter-s3key-rules-value
	Value string `json:"value,omitempty" cloudformation:"Value,Parameter"`
}

// Bucket_SourceSelectionCriteria defines the desired state of BucketSourceSelectionCriteria
type Bucket_SourceSelectionCriteria struct {
	// SseKmsEncryptedObjects http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket-sourceselectioncriteria.html#cfn-s3-bucket-sourceselectioncriteria-ssekmsencryptedobjects
	SseKmsEncryptedObjects Bucket_SseKmsEncryptedObjects `json:"sseKmsEncryptedObjects,omitempty" cloudformation:"SseKmsEncryptedObjects"`
}

// Bucket_RoutingRuleCondition defines the desired state of BucketRoutingRuleCondition
type Bucket_RoutingRuleCondition struct {
	// HttpErrorCodeReturnedEquals http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-websiteconfiguration-routingrules-routingrulecondition.html#cfn-s3-websiteconfiguration-routingrules-routingrulecondition-httperrorcodereturnedequals
	HttpErrorCodeReturnedEquals string `json:"httpErrorCodeReturnedEquals,omitempty" cloudformation:"HttpErrorCodeReturnedEquals,Parameter"`

	// KeyPrefixEquals http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-websiteconfiguration-routingrules-routingrulecondition.html#cfn-s3-websiteconfiguration-routingrules-routingrulecondition-keyprefixequals
	KeyPrefixEquals string `json:"keyPrefixEquals,omitempty" cloudformation:"KeyPrefixEquals,Parameter"`
}

// Bucket_ReplicationConfiguration defines the desired state of BucketReplicationConfiguration
type Bucket_ReplicationConfiguration struct {
	// Role http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket-replicationconfiguration.html#cfn-s3-bucket-replicationconfiguration-role
	Role string `json:"role,omitempty" cloudformation:"Role,Parameter"`

	// Rules http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket-replicationconfiguration.html#cfn-s3-bucket-replicationconfiguration-rules
	Rules []Bucket_ReplicationRule `json:"rules,omitempty" cloudformation:"Rules"`
}

// Bucket_PublicAccessBlockConfiguration defines the desired state of BucketPublicAccessBlockConfiguration
type Bucket_PublicAccessBlockConfiguration struct {
	// RestrictPublicBuckets http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket-publicaccessblockconfiguration.html#cfn-s3-bucket-publicaccessblockconfiguration-restrictpublicbuckets
	RestrictPublicBuckets bool `json:"restrictPublicBuckets,omitempty" cloudformation:"RestrictPublicBuckets,Parameter"`

	// BlockPublicAcls http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket-publicaccessblockconfiguration.html#cfn-s3-bucket-publicaccessblockconfiguration-blockpublicacls
	BlockPublicAcls bool `json:"blockPublicAcls,omitempty" cloudformation:"BlockPublicAcls,Parameter"`

	// BlockPublicPolicy http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket-publicaccessblockconfiguration.html#cfn-s3-bucket-publicaccessblockconfiguration-blockpublicpolicy
	BlockPublicPolicy bool `json:"blockPublicPolicy,omitempty" cloudformation:"BlockPublicPolicy,Parameter"`

	// IgnorePublicAcls http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket-publicaccessblockconfiguration.html#cfn-s3-bucket-publicaccessblockconfiguration-ignorepublicacls
	IgnorePublicAcls bool `json:"ignorePublicAcls,omitempty" cloudformation:"IgnorePublicAcls,Parameter"`
}

// Bucket_AccessControlTranslation defines the desired state of BucketAccessControlTranslation
type Bucket_AccessControlTranslation struct {
	// Owner http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket-accesscontroltranslation.html#cfn-s3-bucket-accesscontroltranslation-owner
	Owner string `json:"owner,omitempty" cloudformation:"Owner,Parameter"`
}

// Bucket_InventoryConfiguration defines the desired state of BucketInventoryConfiguration
type Bucket_InventoryConfiguration struct {
	// Enabled http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket-inventoryconfiguration.html#cfn-s3-bucket-inventoryconfiguration-enabled
	Enabled bool `json:"enabled,omitempty" cloudformation:"Enabled,Parameter"`

	// Ref http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket-inventoryconfiguration.html#cfn-s3-bucket-inventoryconfiguration-id
	Ref metav1alpha1.ObjectReference `json:"ref,omitempty" cloudformation:"Id,Parameter"`

	// IncludedObjectVersions http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket-inventoryconfiguration.html#cfn-s3-bucket-inventoryconfiguration-includedobjectversions
	IncludedObjectVersions string `json:"includedObjectVersions,omitempty" cloudformation:"IncludedObjectVersions,Parameter"`

	// OptionalFields http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket-inventoryconfiguration.html#cfn-s3-bucket-inventoryconfiguration-optionalfields
	OptionalFields []string `json:"optionalFields,omitempty" cloudformation:"OptionalFields"`

	// Prefix http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket-inventoryconfiguration.html#cfn-s3-bucket-inventoryconfiguration-prefix
	Prefix string `json:"prefix,omitempty" cloudformation:"Prefix,Parameter"`

	// ScheduleFrequency http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket-inventoryconfiguration.html#cfn-s3-bucket-inventoryconfiguration-schedulefrequency
	ScheduleFrequency string `json:"scheduleFrequency,omitempty" cloudformation:"ScheduleFrequency,Parameter"`

	// Destination http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket-inventoryconfiguration.html#cfn-s3-bucket-inventoryconfiguration-destination
	Destination Bucket_Destination `json:"destination,omitempty" cloudformation:"Destination"`
}

// Bucket_LoggingConfiguration defines the desired state of BucketLoggingConfiguration
type Bucket_LoggingConfiguration struct {
	// DestinationBucketName http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket-loggingconfig.html#cfn-s3-bucket-loggingconfig-destinationbucketname
	DestinationBucketName string `json:"destinationBucketName,omitempty" cloudformation:"DestinationBucketName,Parameter"`

	// LogFilePrefix http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket-loggingconfig.html#cfn-s3-bucket-loggingconfig-logfileprefix
	LogFilePrefix string `json:"logFilePrefix,omitempty" cloudformation:"LogFilePrefix,Parameter"`
}

// Bucket_Transition defines the desired state of BucketTransition
type Bucket_Transition struct {
	// StorageClass http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket-lifecycleconfig-rule-transition.html#cfn-s3-bucket-lifecycleconfig-rule-transition-storageclass
	StorageClass string `json:"storageClass,omitempty" cloudformation:"StorageClass,Parameter"`

	// TransitionDate http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket-lifecycleconfig-rule-transition.html#cfn-s3-bucket-lifecycleconfig-rule-transition-transitiondate
	TransitionDate string `json:"transitionDate,omitempty" cloudformation:"TransitionDate,Parameter"`

	// TransitionInDays http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket-lifecycleconfig-rule-transition.html#cfn-s3-bucket-lifecycleconfig-rule-transition-transitionindays
	TransitionInDays int `json:"transitionInDays,omitempty" cloudformation:"TransitionInDays,Parameter"`
}

// Bucket_DataExport defines the desired state of BucketDataExport
type Bucket_DataExport struct {
	// Destination http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket-dataexport.html#cfn-s3-bucket-dataexport-destination
	Destination Bucket_Destination `json:"destination,omitempty" cloudformation:"Destination"`

	// OutputSchemaVersion http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket-dataexport.html#cfn-s3-bucket-dataexport-outputschemaversion
	OutputSchemaVersion string `json:"outputSchemaVersion,omitempty" cloudformation:"OutputSchemaVersion,Parameter"`
}

// Bucket_AccelerateConfiguration defines the desired state of BucketAccelerateConfiguration
type Bucket_AccelerateConfiguration struct {
	// AccelerationStatus http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket-accelerateconfiguration.html#cfn-s3-bucket-accelerateconfiguration-accelerationstatus
	AccelerationStatus string `json:"accelerationStatus,omitempty" cloudformation:"AccelerationStatus,Parameter"`
}

// Bucket_LambdaConfiguration defines the desired state of BucketLambdaConfiguration
type Bucket_LambdaConfiguration struct {
	// Event http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket-notificationconfig-lambdaconfig.html#cfn-s3-bucket-notificationconfig-lambdaconfig-event
	Event string `json:"event,omitempty" cloudformation:"Event,Parameter"`

	// Filter http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket-notificationconfig-lambdaconfig.html#cfn-s3-bucket-notificationconfig-lambdaconfig-filter
	Filter Bucket_NotificationFilter `json:"filter,omitempty" cloudformation:"Filter"`

	// Function http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket-notificationconfig-lambdaconfig.html#cfn-s3-bucket-notificationconfig-lambdaconfig-function
	Function string `json:"function,omitempty" cloudformation:"Function,Parameter"`
}

// Bucket_LifecycleConfiguration defines the desired state of BucketLifecycleConfiguration
type Bucket_LifecycleConfiguration struct {
	// Rules http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket-lifecycleconfig.html#cfn-s3-bucket-lifecycleconfig-rules
	Rules []Bucket_Rule `json:"rules,omitempty" cloudformation:"Rules"`
}

// Bucket_BucketEncryption defines the desired state of BucketBucketEncryption
type Bucket_BucketEncryption struct {
	// ServerSideEncryptionConfiguration http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket-bucketencryption.html#cfn-s3-bucket-bucketencryption-serversideencryptionconfiguration
	ServerSideEncryptionConfiguration []Bucket_ServerSideEncryptionRule `json:"serverSideEncryptionConfiguration,omitempty" cloudformation:"ServerSideEncryptionConfiguration"`
}

// Bucket_DefaultRetention defines the desired state of BucketDefaultRetention
type Bucket_DefaultRetention struct {
	// Days http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket-defaultretention.html#cfn-s3-bucket-defaultretention-days
	Days int `json:"days,omitempty" cloudformation:"Days,Parameter"`

	// Mode http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket-defaultretention.html#cfn-s3-bucket-defaultretention-mode
	Mode string `json:"mode,omitempty" cloudformation:"Mode,Parameter"`

	// Years http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket-defaultretention.html#cfn-s3-bucket-defaultretention-years
	Years int `json:"years,omitempty" cloudformation:"Years,Parameter"`
}

// Bucket_RoutingRule defines the desired state of BucketRoutingRule
type Bucket_RoutingRule struct {
	// RedirectRule http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-websiteconfiguration-routingrules.html#cfn-s3-websiteconfiguration-routingrules-redirectrule
	RedirectRule Bucket_RedirectRule `json:"redirectRule,omitempty" cloudformation:"RedirectRule"`

	// RoutingRuleCondition http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-websiteconfiguration-routingrules.html#cfn-s3-websiteconfiguration-routingrules-routingrulecondition
	RoutingRuleCondition Bucket_RoutingRuleCondition `json:"routingRuleCondition,omitempty" cloudformation:"RoutingRuleCondition"`
}

// Bucket_S3KeyFilter defines the desired state of BucketS3KeyFilter
type Bucket_S3KeyFilter struct {
	// Rules http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket-notificationconfiguration-config-filter-s3key.html#cfn-s3-bucket-notificationconfiguraiton-config-filter-s3key-rules
	Rules []Bucket_FilterRule `json:"rules,omitempty" cloudformation:"Rules"`
}

// Bucket_ServerSideEncryptionRule defines the desired state of BucketServerSideEncryptionRule
type Bucket_ServerSideEncryptionRule struct {
	// ServerSideEncryptionByDefault http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket-serversideencryptionrule.html#cfn-s3-bucket-serversideencryptionrule-serversideencryptionbydefault
	ServerSideEncryptionByDefault Bucket_ServerSideEncryptionByDefault `json:"serverSideEncryptionByDefault,omitempty" cloudformation:"ServerSideEncryptionByDefault"`
}

// BucketStatus defines the observed state of Bucket
type BucketStatus struct {
	metav1alpha1.StatusMeta `json:",inline"`
}

// BucketOutput defines the stack outputs
type BucketOutput struct {
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket.html
	Ref string `json:"ref,omitempty"`

	// Arn defines the Arn
	Arn string `json:"arn,omitempty" cloudformation:"Arn,Output"`

	// DomainName defines the DomainName
	DomainName string `json:"domainName,omitempty" cloudformation:"DomainName,Output"`

	// DualStackDomainName defines the DualStackDomainName
	DualStackDomainName string `json:"dualStackDomainName,omitempty" cloudformation:"DualStackDomainName,Output"`

	// RegionalDomainName defines the RegionalDomainName
	RegionalDomainName string `json:"regionalDomainName,omitempty" cloudformation:"RegionalDomainName,Output"`

	// WebsiteURL defines the WebsiteURL
	WebsiteURL string `json:"websiteURL,omitempty" cloudformation:"WebsiteURL,Output"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:resource:categories=aws;s3
// +kubebuilder:printcolumn:JSONPath=.status.status,description="status of the stack",name=Status,priority=0,type=string
// +kubebuilder:printcolumn:JSONPath=.status.message,description="reason for the stack status",name=Message,priority=1,type=string
// +kubebuilder:printcolumn:JSONPath=.status.stackID,description="CloudFormation Stack ID",name=StackID,priority=2,type=string

// Bucket is the Schema for the s3 Bucket API
type Bucket struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   BucketSpec   `json:"spec,omitempty"`
	Status BucketStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// BucketList contains a list of Account
type BucketList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`

	Items []Bucket `json:"items"`
}

func init() {
	SchemeBuilder.Register(&Bucket{}, &BucketList{})
}
