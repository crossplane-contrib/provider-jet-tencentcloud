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

type LogTopicInitParameters struct {

	// Log topic of CLB instance.
	// Log topic of CLB instance.
	// +crossplane:generate:reference:type=LogSet
	LogSetID *string `json:"logSetId,omitempty" tf:"log_set_id,omitempty"`

	// Reference to a LogSet to populate logSetId.
	// +kubebuilder:validation:Optional
	LogSetIDRef *v1.Reference `json:"logSetIdRef,omitempty" tf:"-"`

	// Selector for a LogSet to populate logSetId.
	// +kubebuilder:validation:Optional
	LogSetIDSelector *v1.Selector `json:"logSetIdSelector,omitempty" tf:"-"`

	// Log topic of CLB instance.
	// Log topic of CLB instance.
	TopicName *string `json:"topicName,omitempty" tf:"topic_name,omitempty"`
}

type LogTopicObservation struct {

	// Log topic creation time.
	// Log topic creation time.
	CreateTime *string `json:"createTime,omitempty" tf:"create_time,omitempty"`

	// ID of the resource.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Log topic of CLB instance.
	// Log topic of CLB instance.
	LogSetID *string `json:"logSetId,omitempty" tf:"log_set_id,omitempty"`

	// The status of log topic.
	// The status of log topic.
	Status *bool `json:"status,omitempty" tf:"status,omitempty"`

	// Log topic of CLB instance.
	// Log topic of CLB instance.
	TopicName *string `json:"topicName,omitempty" tf:"topic_name,omitempty"`
}

type LogTopicParameters struct {

	// Log topic of CLB instance.
	// Log topic of CLB instance.
	// +crossplane:generate:reference:type=LogSet
	// +kubebuilder:validation:Optional
	LogSetID *string `json:"logSetId,omitempty" tf:"log_set_id,omitempty"`

	// Reference to a LogSet to populate logSetId.
	// +kubebuilder:validation:Optional
	LogSetIDRef *v1.Reference `json:"logSetIdRef,omitempty" tf:"-"`

	// Selector for a LogSet to populate logSetId.
	// +kubebuilder:validation:Optional
	LogSetIDSelector *v1.Selector `json:"logSetIdSelector,omitempty" tf:"-"`

	// Log topic of CLB instance.
	// Log topic of CLB instance.
	// +kubebuilder:validation:Optional
	TopicName *string `json:"topicName,omitempty" tf:"topic_name,omitempty"`
}

// LogTopicSpec defines the desired state of LogTopic
type LogTopicSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     LogTopicParameters `json:"forProvider"`
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
	InitProvider LogTopicInitParameters `json:"initProvider,omitempty"`
}

// LogTopicStatus defines the observed state of LogTopic.
type LogTopicStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        LogTopicObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// LogTopic is the Schema for the LogTopics API. Provides a resource to create a CLB instance topic.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,tencentcloud}
type LogTopic struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.topicName) || (has(self.initProvider) && has(self.initProvider.topicName))",message="spec.forProvider.topicName is a required parameter"
	Spec   LogTopicSpec   `json:"spec"`
	Status LogTopicStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// LogTopicList contains a list of LogTopics
type LogTopicList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []LogTopic `json:"items"`
}

// Repository type metadata.
var (
	LogTopic_Kind             = "LogTopic"
	LogTopic_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: LogTopic_Kind}.String()
	LogTopic_KindAPIVersion   = LogTopic_Kind + "." + CRDGroupVersion.String()
	LogTopic_GroupVersionKind = CRDGroupVersion.WithKind(LogTopic_Kind)
)

func init() {
	SchemeBuilder.Register(&LogTopic{}, &LogTopicList{})
}
