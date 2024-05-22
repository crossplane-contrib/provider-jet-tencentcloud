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

type LogFileRetentionPeriodInitParameters struct {

	// The number of days to save, cannot exceed 30.
	// The number of days to save, cannot exceed 30.
	Days *float64 `json:"days,omitempty" tf:"days,omitempty"`

	// instance id.
	// instance id.
	InstanceID *string `json:"instanceId,omitempty" tf:"instance_id,omitempty"`
}

type LogFileRetentionPeriodObservation struct {

	// The number of days to save, cannot exceed 30.
	// The number of days to save, cannot exceed 30.
	Days *float64 `json:"days,omitempty" tf:"days,omitempty"`

	// ID of the resource.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// instance id.
	// instance id.
	InstanceID *string `json:"instanceId,omitempty" tf:"instance_id,omitempty"`
}

type LogFileRetentionPeriodParameters struct {

	// The number of days to save, cannot exceed 30.
	// The number of days to save, cannot exceed 30.
	// +kubebuilder:validation:Optional
	Days *float64 `json:"days,omitempty" tf:"days,omitempty"`

	// instance id.
	// instance id.
	// +kubebuilder:validation:Optional
	InstanceID *string `json:"instanceId,omitempty" tf:"instance_id,omitempty"`
}

// LogFileRetentionPeriodSpec defines the desired state of LogFileRetentionPeriod
type LogFileRetentionPeriodSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     LogFileRetentionPeriodParameters `json:"forProvider"`
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
	InitProvider LogFileRetentionPeriodInitParameters `json:"initProvider,omitempty"`
}

// LogFileRetentionPeriodStatus defines the observed state of LogFileRetentionPeriod.
type LogFileRetentionPeriodStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        LogFileRetentionPeriodObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// LogFileRetentionPeriod is the Schema for the LogFileRetentionPeriods API. Provides a resource to create a mariadb log_file_retention_period
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,tencentcloud}
type LogFileRetentionPeriod struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.days) || (has(self.initProvider) && has(self.initProvider.days))",message="spec.forProvider.days is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.instanceId) || (has(self.initProvider) && has(self.initProvider.instanceId))",message="spec.forProvider.instanceId is a required parameter"
	Spec   LogFileRetentionPeriodSpec   `json:"spec"`
	Status LogFileRetentionPeriodStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// LogFileRetentionPeriodList contains a list of LogFileRetentionPeriods
type LogFileRetentionPeriodList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []LogFileRetentionPeriod `json:"items"`
}

// Repository type metadata.
var (
	LogFileRetentionPeriod_Kind             = "LogFileRetentionPeriod"
	LogFileRetentionPeriod_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: LogFileRetentionPeriod_Kind}.String()
	LogFileRetentionPeriod_KindAPIVersion   = LogFileRetentionPeriod_Kind + "." + CRDGroupVersion.String()
	LogFileRetentionPeriod_GroupVersionKind = CRDGroupVersion.WithKind(LogFileRetentionPeriod_Kind)
)

func init() {
	SchemeBuilder.Register(&LogFileRetentionPeriod{}, &LogFileRetentionPeriodList{})
}
