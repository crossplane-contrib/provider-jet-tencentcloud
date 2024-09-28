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

type LifecycleCommandInitParameters struct {

	// Remote command ID. It is required to execute a command.
	// Remote command ID. It is required to execute a command.
	CommandID *string `json:"commandId,omitempty" tf:"command_id,omitempty"`

	// Custom parameter. The field type is JSON encoded string. For example, {"varA": "222"}.
	// Custom parameter. The field type is JSON encoded string. For example, {"varA": "222"}.
	Parameters *string `json:"parameters,omitempty" tf:"parameters,omitempty"`
}

type LifecycleCommandObservation struct {

	// Remote command ID. It is required to execute a command.
	// Remote command ID. It is required to execute a command.
	CommandID *string `json:"commandId,omitempty" tf:"command_id,omitempty"`

	// Custom parameter. The field type is JSON encoded string. For example, {"varA": "222"}.
	// Custom parameter. The field type is JSON encoded string. For example, {"varA": "222"}.
	Parameters *string `json:"parameters,omitempty" tf:"parameters,omitempty"`
}

type LifecycleCommandParameters struct {

	// Remote command ID. It is required to execute a command.
	// Remote command ID. It is required to execute a command.
	// +kubebuilder:validation:Optional
	CommandID *string `json:"commandId" tf:"command_id,omitempty"`

	// Custom parameter. The field type is JSON encoded string. For example, {"varA": "222"}.
	// Custom parameter. The field type is JSON encoded string. For example, {"varA": "222"}.
	// +kubebuilder:validation:Optional
	Parameters *string `json:"parameters,omitempty" tf:"parameters,omitempty"`
}

type LifecycleHookInitParameters struct {

	// Defines the action the AS group should take when the lifecycle hook timeout elapses or if an unexpected failure occurs. Valid values: CONTINUE and ABANDON. The default value is CONTINUE.
	// Defines the action the AS group should take when the lifecycle hook timeout elapses or if an unexpected failure occurs. Valid values: `CONTINUE` and `ABANDON`. The default value is `CONTINUE`.
	DefaultResult *string `json:"defaultResult,omitempty" tf:"default_result,omitempty"`

	// Defines the amount of time, in seconds, that can elapse before the lifecycle hook times out. Valid value ranges: (30~7200). and default value is 300.
	// Defines the amount of time, in seconds, that can elapse before the lifecycle hook times out. Valid value ranges: (30~7200). and default value is `300`.
	HeartbeatTimeout *float64 `json:"heartbeatTimeout,omitempty" tf:"heartbeat_timeout,omitempty"`

	// Remote command execution object. NotificationTarget and LifecycleCommand cannot be specified at the same time.
	// Remote command execution object. `NotificationTarget` and `LifecycleCommand` cannot be specified at the same time.
	LifecycleCommand []LifecycleCommandInitParameters `json:"lifecycleCommand,omitempty" tf:"lifecycle_command,omitempty"`

	// The name of the lifecycle hook.
	// The name of the lifecycle hook.
	LifecycleHookName *string `json:"lifecycleHookName,omitempty" tf:"lifecycle_hook_name,omitempty"`

	// The instance state to which you want to attach the lifecycle hook. Valid values: INSTANCE_LAUNCHING and INSTANCE_TERMINATING.
	// The instance state to which you want to attach the lifecycle hook. Valid values: `INSTANCE_LAUNCHING` and `INSTANCE_TERMINATING`.
	LifecycleTransition *string `json:"lifecycleTransition,omitempty" tf:"lifecycle_transition,omitempty"`

	// Contains additional information that you want to include any time AS sends a message to the notification target.
	// Contains additional information that you want to include any time AS sends a message to the notification target.
	NotificationMetadata *string `json:"notificationMetadata,omitempty" tf:"notification_metadata,omitempty"`

	// For CMQ_QUEUE type, a name of queue must be set.
	// For CMQ_QUEUE type, a name of queue must be set.
	NotificationQueueName *string `json:"notificationQueueName,omitempty" tf:"notification_queue_name,omitempty"`

	// Target type. Valid values: CMQ_QUEUE, CMQ_TOPIC, TDMQ_CMQ_QUEUE, TDMQ_CMQ_TOPIC.
	// Target type. Valid values: `CMQ_QUEUE`, `CMQ_TOPIC`, `TDMQ_CMQ_QUEUE`, `TDMQ_CMQ_TOPIC`.
	NotificationTargetType *string `json:"notificationTargetType,omitempty" tf:"notification_target_type,omitempty"`

	// For CMQ_TOPIC type, a name of topic must be set.
	// For CMQ_TOPIC type, a name of topic must be set.
	NotificationTopicName *string `json:"notificationTopicName,omitempty" tf:"notification_topic_name,omitempty"`

	// ID of a scaling group.
	// ID of a scaling group.
	// +crossplane:generate:reference:type=ScalingGroup
	ScalingGroupID *string `json:"scalingGroupId,omitempty" tf:"scaling_group_id,omitempty"`

	// Reference to a ScalingGroup to populate scalingGroupId.
	// +kubebuilder:validation:Optional
	ScalingGroupIDRef *v1.Reference `json:"scalingGroupIdRef,omitempty" tf:"-"`

	// Selector for a ScalingGroup to populate scalingGroupId.
	// +kubebuilder:validation:Optional
	ScalingGroupIDSelector *v1.Selector `json:"scalingGroupIdSelector,omitempty" tf:"-"`
}

type LifecycleHookObservation struct {

	// Defines the action the AS group should take when the lifecycle hook timeout elapses or if an unexpected failure occurs. Valid values: CONTINUE and ABANDON. The default value is CONTINUE.
	// Defines the action the AS group should take when the lifecycle hook timeout elapses or if an unexpected failure occurs. Valid values: `CONTINUE` and `ABANDON`. The default value is `CONTINUE`.
	DefaultResult *string `json:"defaultResult,omitempty" tf:"default_result,omitempty"`

	// Defines the amount of time, in seconds, that can elapse before the lifecycle hook times out. Valid value ranges: (30~7200). and default value is 300.
	// Defines the amount of time, in seconds, that can elapse before the lifecycle hook times out. Valid value ranges: (30~7200). and default value is `300`.
	HeartbeatTimeout *float64 `json:"heartbeatTimeout,omitempty" tf:"heartbeat_timeout,omitempty"`

	// ID of the resource.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Remote command execution object. NotificationTarget and LifecycleCommand cannot be specified at the same time.
	// Remote command execution object. `NotificationTarget` and `LifecycleCommand` cannot be specified at the same time.
	LifecycleCommand []LifecycleCommandObservation `json:"lifecycleCommand,omitempty" tf:"lifecycle_command,omitempty"`

	// The name of the lifecycle hook.
	// The name of the lifecycle hook.
	LifecycleHookName *string `json:"lifecycleHookName,omitempty" tf:"lifecycle_hook_name,omitempty"`

	// The instance state to which you want to attach the lifecycle hook. Valid values: INSTANCE_LAUNCHING and INSTANCE_TERMINATING.
	// The instance state to which you want to attach the lifecycle hook. Valid values: `INSTANCE_LAUNCHING` and `INSTANCE_TERMINATING`.
	LifecycleTransition *string `json:"lifecycleTransition,omitempty" tf:"lifecycle_transition,omitempty"`

	// Contains additional information that you want to include any time AS sends a message to the notification target.
	// Contains additional information that you want to include any time AS sends a message to the notification target.
	NotificationMetadata *string `json:"notificationMetadata,omitempty" tf:"notification_metadata,omitempty"`

	// For CMQ_QUEUE type, a name of queue must be set.
	// For CMQ_QUEUE type, a name of queue must be set.
	NotificationQueueName *string `json:"notificationQueueName,omitempty" tf:"notification_queue_name,omitempty"`

	// Target type. Valid values: CMQ_QUEUE, CMQ_TOPIC, TDMQ_CMQ_QUEUE, TDMQ_CMQ_TOPIC.
	// Target type. Valid values: `CMQ_QUEUE`, `CMQ_TOPIC`, `TDMQ_CMQ_QUEUE`, `TDMQ_CMQ_TOPIC`.
	NotificationTargetType *string `json:"notificationTargetType,omitempty" tf:"notification_target_type,omitempty"`

	// For CMQ_TOPIC type, a name of topic must be set.
	// For CMQ_TOPIC type, a name of topic must be set.
	NotificationTopicName *string `json:"notificationTopicName,omitempty" tf:"notification_topic_name,omitempty"`

	// ID of a scaling group.
	// ID of a scaling group.
	ScalingGroupID *string `json:"scalingGroupId,omitempty" tf:"scaling_group_id,omitempty"`
}

type LifecycleHookParameters struct {

	// Defines the action the AS group should take when the lifecycle hook timeout elapses or if an unexpected failure occurs. Valid values: CONTINUE and ABANDON. The default value is CONTINUE.
	// Defines the action the AS group should take when the lifecycle hook timeout elapses or if an unexpected failure occurs. Valid values: `CONTINUE` and `ABANDON`. The default value is `CONTINUE`.
	// +kubebuilder:validation:Optional
	DefaultResult *string `json:"defaultResult,omitempty" tf:"default_result,omitempty"`

	// Defines the amount of time, in seconds, that can elapse before the lifecycle hook times out. Valid value ranges: (30~7200). and default value is 300.
	// Defines the amount of time, in seconds, that can elapse before the lifecycle hook times out. Valid value ranges: (30~7200). and default value is `300`.
	// +kubebuilder:validation:Optional
	HeartbeatTimeout *float64 `json:"heartbeatTimeout,omitempty" tf:"heartbeat_timeout,omitempty"`

	// Remote command execution object. NotificationTarget and LifecycleCommand cannot be specified at the same time.
	// Remote command execution object. `NotificationTarget` and `LifecycleCommand` cannot be specified at the same time.
	// +kubebuilder:validation:Optional
	LifecycleCommand []LifecycleCommandParameters `json:"lifecycleCommand,omitempty" tf:"lifecycle_command,omitempty"`

	// The name of the lifecycle hook.
	// The name of the lifecycle hook.
	// +kubebuilder:validation:Optional
	LifecycleHookName *string `json:"lifecycleHookName,omitempty" tf:"lifecycle_hook_name,omitempty"`

	// The instance state to which you want to attach the lifecycle hook. Valid values: INSTANCE_LAUNCHING and INSTANCE_TERMINATING.
	// The instance state to which you want to attach the lifecycle hook. Valid values: `INSTANCE_LAUNCHING` and `INSTANCE_TERMINATING`.
	// +kubebuilder:validation:Optional
	LifecycleTransition *string `json:"lifecycleTransition,omitempty" tf:"lifecycle_transition,omitempty"`

	// Contains additional information that you want to include any time AS sends a message to the notification target.
	// Contains additional information that you want to include any time AS sends a message to the notification target.
	// +kubebuilder:validation:Optional
	NotificationMetadata *string `json:"notificationMetadata,omitempty" tf:"notification_metadata,omitempty"`

	// For CMQ_QUEUE type, a name of queue must be set.
	// For CMQ_QUEUE type, a name of queue must be set.
	// +kubebuilder:validation:Optional
	NotificationQueueName *string `json:"notificationQueueName,omitempty" tf:"notification_queue_name,omitempty"`

	// Target type. Valid values: CMQ_QUEUE, CMQ_TOPIC, TDMQ_CMQ_QUEUE, TDMQ_CMQ_TOPIC.
	// Target type. Valid values: `CMQ_QUEUE`, `CMQ_TOPIC`, `TDMQ_CMQ_QUEUE`, `TDMQ_CMQ_TOPIC`.
	// +kubebuilder:validation:Optional
	NotificationTargetType *string `json:"notificationTargetType,omitempty" tf:"notification_target_type,omitempty"`

	// For CMQ_TOPIC type, a name of topic must be set.
	// For CMQ_TOPIC type, a name of topic must be set.
	// +kubebuilder:validation:Optional
	NotificationTopicName *string `json:"notificationTopicName,omitempty" tf:"notification_topic_name,omitempty"`

	// ID of a scaling group.
	// ID of a scaling group.
	// +crossplane:generate:reference:type=ScalingGroup
	// +kubebuilder:validation:Optional
	ScalingGroupID *string `json:"scalingGroupId,omitempty" tf:"scaling_group_id,omitempty"`

	// Reference to a ScalingGroup to populate scalingGroupId.
	// +kubebuilder:validation:Optional
	ScalingGroupIDRef *v1.Reference `json:"scalingGroupIdRef,omitempty" tf:"-"`

	// Selector for a ScalingGroup to populate scalingGroupId.
	// +kubebuilder:validation:Optional
	ScalingGroupIDSelector *v1.Selector `json:"scalingGroupIdSelector,omitempty" tf:"-"`
}

// LifecycleHookSpec defines the desired state of LifecycleHook
type LifecycleHookSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     LifecycleHookParameters `json:"forProvider"`
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
	InitProvider LifecycleHookInitParameters `json:"initProvider,omitempty"`
}

// LifecycleHookStatus defines the observed state of LifecycleHook.
type LifecycleHookStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        LifecycleHookObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// LifecycleHook is the Schema for the LifecycleHooks API. Provides a resource for an AS (Auto scaling) lifecycle hook.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,tencentcloud}
type LifecycleHook struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.lifecycleHookName) || (has(self.initProvider) && has(self.initProvider.lifecycleHookName))",message="spec.forProvider.lifecycleHookName is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.lifecycleTransition) || (has(self.initProvider) && has(self.initProvider.lifecycleTransition))",message="spec.forProvider.lifecycleTransition is a required parameter"
	Spec   LifecycleHookSpec   `json:"spec"`
	Status LifecycleHookStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// LifecycleHookList contains a list of LifecycleHooks
type LifecycleHookList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []LifecycleHook `json:"items"`
}

// Repository type metadata.
var (
	LifecycleHook_Kind             = "LifecycleHook"
	LifecycleHook_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: LifecycleHook_Kind}.String()
	LifecycleHook_KindAPIVersion   = LifecycleHook_Kind + "." + CRDGroupVersion.String()
	LifecycleHook_GroupVersionKind = CRDGroupVersion.WithKind(LifecycleHook_Kind)
)

func init() {
	SchemeBuilder.Register(&LifecycleHook{}, &LifecycleHookList{})
}
