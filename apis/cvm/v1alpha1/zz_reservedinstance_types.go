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

type ReservedInstanceInitParameters struct {

	// Configuration ID of the reserved instance.
	// Configuration ID of the reserved instance.
	ConfigID *string `json:"configId,omitempty" tf:"config_id,omitempty"`

	// Number of reserved instances to be purchased.
	// Number of reserved instances to be purchased.
	InstanceCount *float64 `json:"instanceCount,omitempty" tf:"instance_count,omitempty"`

	// Reserved Instance display name.
	// - If you do not specify an instance display name, 'Unnamed' is displayed by default.
	// - Up to 60 characters (including pattern strings) are supported.
	// Reserved Instance display name.
	// - If you do not specify an instance display name, 'Unnamed' is displayed by default.
	// - Up to 60 characters (including pattern strings) are supported.
	ReservedInstanceName *string `json:"reservedInstanceName,omitempty" tf:"reserved_instance_name,omitempty"`
}

type ReservedInstanceObservation struct {

	// Configuration ID of the reserved instance.
	// Configuration ID of the reserved instance.
	ConfigID *string `json:"configId,omitempty" tf:"config_id,omitempty"`

	// Expiry time of the RI.
	// Expiry time of the RI.
	EndTime *string `json:"endTime,omitempty" tf:"end_time,omitempty"`

	// ID of the resource.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Number of reserved instances to be purchased.
	// Number of reserved instances to be purchased.
	InstanceCount *float64 `json:"instanceCount,omitempty" tf:"instance_count,omitempty"`

	// Reserved Instance display name.
	// - If you do not specify an instance display name, 'Unnamed' is displayed by default.
	// - Up to 60 characters (including pattern strings) are supported.
	// Reserved Instance display name.
	// - If you do not specify an instance display name, 'Unnamed' is displayed by default.
	// - Up to 60 characters (including pattern strings) are supported.
	ReservedInstanceName *string `json:"reservedInstanceName,omitempty" tf:"reserved_instance_name,omitempty"`

	// Start time of the RI.
	// Start time of the RI.
	StartTime *string `json:"startTime,omitempty" tf:"start_time,omitempty"`

	// Status of the RI at the time of purchase.
	// Status of the RI at the time of purchase.
	Status *string `json:"status,omitempty" tf:"status,omitempty"`
}

type ReservedInstanceParameters struct {

	// Configuration ID of the reserved instance.
	// Configuration ID of the reserved instance.
	// +kubebuilder:validation:Optional
	ConfigID *string `json:"configId,omitempty" tf:"config_id,omitempty"`

	// Number of reserved instances to be purchased.
	// Number of reserved instances to be purchased.
	// +kubebuilder:validation:Optional
	InstanceCount *float64 `json:"instanceCount,omitempty" tf:"instance_count,omitempty"`

	// Reserved Instance display name.
	// - If you do not specify an instance display name, 'Unnamed' is displayed by default.
	// - Up to 60 characters (including pattern strings) are supported.
	// Reserved Instance display name.
	// - If you do not specify an instance display name, 'Unnamed' is displayed by default.
	// - Up to 60 characters (including pattern strings) are supported.
	// +kubebuilder:validation:Optional
	ReservedInstanceName *string `json:"reservedInstanceName,omitempty" tf:"reserved_instance_name,omitempty"`
}

// ReservedInstanceSpec defines the desired state of ReservedInstance
type ReservedInstanceSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ReservedInstanceParameters `json:"forProvider"`
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
	InitProvider ReservedInstanceInitParameters `json:"initProvider,omitempty"`
}

// ReservedInstanceStatus defines the observed state of ReservedInstance.
type ReservedInstanceStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ReservedInstanceObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// ReservedInstance is the Schema for the ReservedInstances API. Provides a reserved instance resource.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,tencentcloud}
type ReservedInstance struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.configId) || (has(self.initProvider) && has(self.initProvider.configId))",message="spec.forProvider.configId is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.instanceCount) || (has(self.initProvider) && has(self.initProvider.instanceCount))",message="spec.forProvider.instanceCount is a required parameter"
	Spec   ReservedInstanceSpec   `json:"spec"`
	Status ReservedInstanceStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ReservedInstanceList contains a list of ReservedInstances
type ReservedInstanceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ReservedInstance `json:"items"`
}

// Repository type metadata.
var (
	ReservedInstance_Kind             = "ReservedInstance"
	ReservedInstance_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: ReservedInstance_Kind}.String()
	ReservedInstance_KindAPIVersion   = ReservedInstance_Kind + "." + CRDGroupVersion.String()
	ReservedInstance_GroupVersionKind = CRDGroupVersion.WithKind(ReservedInstance_Kind)
)

func init() {
	SchemeBuilder.Register(&ReservedInstance{}, &ReservedInstanceList{})
}
