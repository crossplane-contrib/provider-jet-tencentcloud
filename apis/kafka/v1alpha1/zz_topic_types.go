// SPDX-FileCopyrightText: 2023 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

/*
Copyright 2022 Upbound Inc.
*/

// Code generated by upjet. DO NOT EDIT.

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type TopicInitParameters struct {

	// Clear log policy, log clear mode, default is delete. delete: logs are deleted according to the storage time. compact: logs are compressed according to the key. compact, delete: logs are compressed according to the key and will be deleted according to the storage time.
	// Clear log policy, log clear mode, default is `delete`. `delete`: logs are deleted according to the storage time. `compact`: logs are compressed according to the key. `compact, delete`: logs are compressed according to the key and will be deleted according to the storage time.
	CleanUpPolicy *string `json:"cleanUpPolicy,omitempty" tf:"clean_up_policy,omitempty"`

	// Whether to open the ip whitelist, true: open, false: close.
	// Whether to open the ip whitelist, `true`: open, `false`: close.
	EnableWhiteList *bool `json:"enableWhiteList,omitempty" tf:"enable_white_list,omitempty"`

	// Ip whitelist, quota limit, required when enableWhileList=true.
	// Ip whitelist, quota limit, required when enableWhileList=true.
	IPWhiteList []*string `json:"ipWhiteList,omitempty" tf:"ip_white_list,omitempty"`

	// Max message bytes. min: 1024 Byte(1KB), max: 8388608 Byte(8MB).
	// Max message bytes. min: 1024 Byte(1KB), max: 8388608 Byte(8MB).
	MaxMessageBytes *float64 `json:"maxMessageBytes,omitempty" tf:"max_message_bytes,omitempty"`

	// The subject note. It must start with a letter, and the remaining part can contain letters, numbers and dashes (-).
	// The subject note. It must start with a letter, and the remaining part can contain letters, numbers and dashes (-).
	Note *string `json:"note,omitempty" tf:"note,omitempty"`

	// The number of partition.
	// The number of partition.
	PartitionNum *float64 `json:"partitionNum,omitempty" tf:"partition_num,omitempty"`

	// The number of replica.
	// The number of replica.
	ReplicaNum *float64 `json:"replicaNum,omitempty" tf:"replica_num,omitempty"`

	// Message can be selected. Retention time, unit is ms, the current minimum value is 60000ms.
	// Message can be selected. Retention time, unit is ms, the current minimum value is 60000ms.
	Retention *float64 `json:"retention,omitempty" tf:"retention,omitempty"`

	// Segment scrolling time, in ms, the current minimum is 3600000ms.
	// Segment scrolling time, in ms, the current minimum is 3600000ms.
	Segment *float64 `json:"segment,omitempty" tf:"segment,omitempty"`

	// Min number of sync replicas, Default is 1.
	// Min number of sync replicas, Default is `1`.
	SyncReplicaMinNum *float64 `json:"syncReplicaMinNum,omitempty" tf:"sync_replica_min_num,omitempty"`

	// Name of the CKafka topic. It must start with a letter, the rest can contain letters, numbers and dashes(-).
	// Name of the CKafka topic. It must start with a letter, the rest can contain letters, numbers and dashes(-).
	TopicName *string `json:"topicName,omitempty" tf:"topic_name,omitempty"`

	// Whether to allow unsynchronized replicas to be selected as leader, default is false, true: allowed, false: not allowed.
	// Whether to allow unsynchronized replicas to be selected as leader, default is `false`, `true: `allowed, `false`: not allowed.
	UncleanLeaderElectionEnable *bool `json:"uncleanLeaderElectionEnable,omitempty" tf:"unclean_leader_election_enable,omitempty"`
}

type TopicObservation struct {

	// Clear log policy, log clear mode, default is delete. delete: logs are deleted according to the storage time. compact: logs are compressed according to the key. compact, delete: logs are compressed according to the key and will be deleted according to the storage time.
	// Clear log policy, log clear mode, default is `delete`. `delete`: logs are deleted according to the storage time. `compact`: logs are compressed according to the key. `compact, delete`: logs are compressed according to the key and will be deleted according to the storage time.
	CleanUpPolicy *string `json:"cleanUpPolicy,omitempty" tf:"clean_up_policy,omitempty"`

	// Create time of the CKafka topic.
	// Create time of the CKafka topic.
	CreateTime *string `json:"createTime,omitempty" tf:"create_time,omitempty"`

	// Whether to open the ip whitelist, true: open, false: close.
	// Whether to open the ip whitelist, `true`: open, `false`: close.
	EnableWhiteList *bool `json:"enableWhiteList,omitempty" tf:"enable_white_list,omitempty"`

	// Data backup cos bucket: the bucket address that is dumped to cos.
	// Data backup cos bucket: the bucket address that is dumped to cos.
	ForwardCosBucket *string `json:"forwardCosBucket,omitempty" tf:"forward_cos_bucket,omitempty"`

	// Periodic frequency of data backup to cos.
	// Periodic frequency of data backup to cos.
	ForwardInterval *float64 `json:"forwardInterval,omitempty" tf:"forward_interval,omitempty"`

	// Data backup cos status. Valid values: 0, 1. 1: do not open data backup, 0: open data backup.
	// Data backup cos status. Valid values: `0`, `1`. `1`: do not open data backup, `0`: open data backup.
	ForwardStatus *float64 `json:"forwardStatus,omitempty" tf:"forward_status,omitempty"`

	// ID of the resource.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Ip whitelist, quota limit, required when enableWhileList=true.
	// Ip whitelist, quota limit, required when enableWhileList=true.
	IPWhiteList []*string `json:"ipWhiteList,omitempty" tf:"ip_white_list,omitempty"`

	// Ckafka instance ID.
	// Ckafka instance ID.
	InstanceID *string `json:"instanceId,omitempty" tf:"instance_id,omitempty"`

	// Max message bytes. min: 1024 Byte(1KB), max: 8388608 Byte(8MB).
	// Max message bytes. min: 1024 Byte(1KB), max: 8388608 Byte(8MB).
	MaxMessageBytes *float64 `json:"maxMessageBytes,omitempty" tf:"max_message_bytes,omitempty"`

	// Message storage location.
	// Message storage location.
	MessageStorageLocation *string `json:"messageStorageLocation,omitempty" tf:"message_storage_location,omitempty"`

	// The subject note. It must start with a letter, and the remaining part can contain letters, numbers and dashes (-).
	// The subject note. It must start with a letter, and the remaining part can contain letters, numbers and dashes (-).
	Note *string `json:"note,omitempty" tf:"note,omitempty"`

	// The number of partition.
	// The number of partition.
	PartitionNum *float64 `json:"partitionNum,omitempty" tf:"partition_num,omitempty"`

	// The number of replica.
	// The number of replica.
	ReplicaNum *float64 `json:"replicaNum,omitempty" tf:"replica_num,omitempty"`

	// Message can be selected. Retention time, unit is ms, the current minimum value is 60000ms.
	// Message can be selected. Retention time, unit is ms, the current minimum value is 60000ms.
	Retention *float64 `json:"retention,omitempty" tf:"retention,omitempty"`

	// Segment scrolling time, in ms, the current minimum is 3600000ms.
	// Segment scrolling time, in ms, the current minimum is 3600000ms.
	Segment *float64 `json:"segment,omitempty" tf:"segment,omitempty"`

	// Number of bytes rolled by shard.
	// Number of bytes rolled by shard.
	SegmentBytes *float64 `json:"segmentBytes,omitempty" tf:"segment_bytes,omitempty"`

	// Min number of sync replicas, Default is 1.
	// Min number of sync replicas, Default is `1`.
	SyncReplicaMinNum *float64 `json:"syncReplicaMinNum,omitempty" tf:"sync_replica_min_num,omitempty"`

	// Name of the CKafka topic. It must start with a letter, the rest can contain letters, numbers and dashes(-).
	// Name of the CKafka topic. It must start with a letter, the rest can contain letters, numbers and dashes(-).
	TopicName *string `json:"topicName,omitempty" tf:"topic_name,omitempty"`

	// Whether to allow unsynchronized replicas to be selected as leader, default is false, true: allowed, false: not allowed.
	// Whether to allow unsynchronized replicas to be selected as leader, default is `false`, `true: `allowed, `false`: not allowed.
	UncleanLeaderElectionEnable *bool `json:"uncleanLeaderElectionEnable,omitempty" tf:"unclean_leader_election_enable,omitempty"`
}

type TopicParameters struct {

	// Clear log policy, log clear mode, default is delete. delete: logs are deleted according to the storage time. compact: logs are compressed according to the key. compact, delete: logs are compressed according to the key and will be deleted according to the storage time.
	// Clear log policy, log clear mode, default is `delete`. `delete`: logs are deleted according to the storage time. `compact`: logs are compressed according to the key. `compact, delete`: logs are compressed according to the key and will be deleted according to the storage time.
	// +kubebuilder:validation:Optional
	CleanUpPolicy *string `json:"cleanUpPolicy,omitempty" tf:"clean_up_policy,omitempty"`

	// Whether to open the ip whitelist, true: open, false: close.
	// Whether to open the ip whitelist, `true`: open, `false`: close.
	// +kubebuilder:validation:Optional
	EnableWhiteList *bool `json:"enableWhiteList,omitempty" tf:"enable_white_list,omitempty"`

	// Ip whitelist, quota limit, required when enableWhileList=true.
	// Ip whitelist, quota limit, required when enableWhileList=true.
	// +kubebuilder:validation:Optional
	IPWhiteList []*string `json:"ipWhiteList,omitempty" tf:"ip_white_list,omitempty"`

	// Ckafka instance ID.
	// Ckafka instance ID.
	// +crossplane:generate:reference:type=Instance
	// +kubebuilder:validation:Optional
	InstanceID *string `json:"instanceId,omitempty" tf:"instance_id,omitempty"`

	// Reference to a Instance to populate instanceId.
	// +kubebuilder:validation:Optional
	InstanceIDRef *v1.Reference `json:"instanceIdRef,omitempty" tf:"-"`

	// Selector for a Instance to populate instanceId.
	// +kubebuilder:validation:Optional
	InstanceIDSelector *v1.Selector `json:"instanceIdSelector,omitempty" tf:"-"`

	// Max message bytes. min: 1024 Byte(1KB), max: 8388608 Byte(8MB).
	// Max message bytes. min: 1024 Byte(1KB), max: 8388608 Byte(8MB).
	// +kubebuilder:validation:Optional
	MaxMessageBytes *float64 `json:"maxMessageBytes,omitempty" tf:"max_message_bytes,omitempty"`

	// The subject note. It must start with a letter, and the remaining part can contain letters, numbers and dashes (-).
	// The subject note. It must start with a letter, and the remaining part can contain letters, numbers and dashes (-).
	// +kubebuilder:validation:Optional
	Note *string `json:"note,omitempty" tf:"note,omitempty"`

	// The number of partition.
	// The number of partition.
	// +kubebuilder:validation:Optional
	PartitionNum *float64 `json:"partitionNum,omitempty" tf:"partition_num,omitempty"`

	// The number of replica.
	// The number of replica.
	// +kubebuilder:validation:Optional
	ReplicaNum *float64 `json:"replicaNum,omitempty" tf:"replica_num,omitempty"`

	// Message can be selected. Retention time, unit is ms, the current minimum value is 60000ms.
	// Message can be selected. Retention time, unit is ms, the current minimum value is 60000ms.
	// +kubebuilder:validation:Optional
	Retention *float64 `json:"retention,omitempty" tf:"retention,omitempty"`

	// Segment scrolling time, in ms, the current minimum is 3600000ms.
	// Segment scrolling time, in ms, the current minimum is 3600000ms.
	// +kubebuilder:validation:Optional
	Segment *float64 `json:"segment,omitempty" tf:"segment,omitempty"`

	// Min number of sync replicas, Default is 1.
	// Min number of sync replicas, Default is `1`.
	// +kubebuilder:validation:Optional
	SyncReplicaMinNum *float64 `json:"syncReplicaMinNum,omitempty" tf:"sync_replica_min_num,omitempty"`

	// Name of the CKafka topic. It must start with a letter, the rest can contain letters, numbers and dashes(-).
	// Name of the CKafka topic. It must start with a letter, the rest can contain letters, numbers and dashes(-).
	// +kubebuilder:validation:Optional
	TopicName *string `json:"topicName,omitempty" tf:"topic_name,omitempty"`

	// Whether to allow unsynchronized replicas to be selected as leader, default is false, true: allowed, false: not allowed.
	// Whether to allow unsynchronized replicas to be selected as leader, default is `false`, `true: `allowed, `false`: not allowed.
	// +kubebuilder:validation:Optional
	UncleanLeaderElectionEnable *bool `json:"uncleanLeaderElectionEnable,omitempty" tf:"unclean_leader_election_enable,omitempty"`
}

// TopicSpec defines the desired state of Topic
type TopicSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     TopicParameters `json:"forProvider"`
	// THIS IS A BETA FIELD. It will be honored
	// unless the Management Policies feature flag is disabled.
	// InitProvider holds the same fields as ForProvider, with the exception
	// of Identifier and other resource reference fields. The fields that are
	// in InitProvider are merged into ForProvider when the resource is created.
	// The same fields are also added to the terraform ignore_changes hook, to
	// avoid updating them after creation. This is useful for fields that are
	// required on creation, but we do not desire to update them after creation,
	// for example because of an external controller is managing them, like an
	// autoscaler.
	InitProvider TopicInitParameters `json:"initProvider,omitempty"`
}

// TopicStatus defines the observed state of Topic.
type TopicStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        TopicObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Topic is the Schema for the Topics API. Use this resource to create ckafka topic.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,tencentcloud}
type Topic struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.partitionNum) || (has(self.initProvider) && has(self.initProvider.partitionNum))",message="spec.forProvider.partitionNum is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.replicaNum) || (has(self.initProvider) && has(self.initProvider.replicaNum))",message="spec.forProvider.replicaNum is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.topicName) || (has(self.initProvider) && has(self.initProvider.topicName))",message="spec.forProvider.topicName is a required parameter"
	Spec   TopicSpec   `json:"spec"`
	Status TopicStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// TopicList contains a list of Topics
type TopicList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Topic `json:"items"`
}

// Repository type metadata.
var (
	Topic_Kind             = "Topic"
	Topic_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Topic_Kind}.String()
	Topic_KindAPIVersion   = Topic_Kind + "." + CRDGroupVersion.String()
	Topic_GroupVersionKind = CRDGroupVersion.WithKind(Topic_Kind)
)

func init() {
	SchemeBuilder.Register(&Topic{}, &TopicList{})
}
