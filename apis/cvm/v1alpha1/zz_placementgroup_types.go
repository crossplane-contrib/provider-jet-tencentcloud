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

type PlacementGroupInitParameters struct {

	// Name of the placement group, 1-60 characters in length.
	// Name of the placement group, 1-60 characters in length.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Type of the placement group. Valid values: HOST, SW and RACK.
	// Type of the placement group. Valid values: `HOST`, `SW` and `RACK`.
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

type PlacementGroupObservation struct {

	// Creation time of the placement group.
	// Creation time of the placement group.
	CreateTime *string `json:"createTime,omitempty" tf:"create_time,omitempty"`

	// Number of hosts in the placement group.
	// Number of hosts in the placement group.
	CurrentNum *float64 `json:"currentNum,omitempty" tf:"current_num,omitempty"`

	// Maximum number of hosts in the placement group.
	// Maximum number of hosts in the placement group.
	CvmQuotaTotal *float64 `json:"cvmQuotaTotal,omitempty" tf:"cvm_quota_total,omitempty"`

	// ID of the resource.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Name of the placement group, 1-60 characters in length.
	// Name of the placement group, 1-60 characters in length.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Type of the placement group. Valid values: HOST, SW and RACK.
	// Type of the placement group. Valid values: `HOST`, `SW` and `RACK`.
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

type PlacementGroupParameters struct {

	// Name of the placement group, 1-60 characters in length.
	// Name of the placement group, 1-60 characters in length.
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Type of the placement group. Valid values: HOST, SW and RACK.
	// Type of the placement group. Valid values: `HOST`, `SW` and `RACK`.
	// +kubebuilder:validation:Optional
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

// PlacementGroupSpec defines the desired state of PlacementGroup
type PlacementGroupSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     PlacementGroupParameters `json:"forProvider"`
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
	InitProvider PlacementGroupInitParameters `json:"initProvider,omitempty"`
}

// PlacementGroupStatus defines the observed state of PlacementGroup.
type PlacementGroupStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        PlacementGroupObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// PlacementGroup is the Schema for the PlacementGroups API. Provide a resource to create a placement group.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,tencentcloud}
type PlacementGroup struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.name) || (has(self.initProvider) && has(self.initProvider.name))",message="spec.forProvider.name is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.type) || (has(self.initProvider) && has(self.initProvider.type))",message="spec.forProvider.type is a required parameter"
	Spec   PlacementGroupSpec   `json:"spec"`
	Status PlacementGroupStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// PlacementGroupList contains a list of PlacementGroups
type PlacementGroupList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []PlacementGroup `json:"items"`
}

// Repository type metadata.
var (
	PlacementGroup_Kind             = "PlacementGroup"
	PlacementGroup_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: PlacementGroup_Kind}.String()
	PlacementGroup_KindAPIVersion   = PlacementGroup_Kind + "." + CRDGroupVersion.String()
	PlacementGroup_GroupVersionKind = CRDGroupVersion.WithKind(PlacementGroup_Kind)
)

func init() {
	SchemeBuilder.Register(&PlacementGroup{}, &PlacementGroupList{})
}
