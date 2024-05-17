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

type DimensionsInitParameters struct {

	// Represents a collection of dimensions of an object instance, json format.eg:'{"unInstanceId":"ins-ot3cq4bi"}'.
	// Represents a collection of dimensions of an object instance, json format.eg:'{"unInstanceId":"ins-ot3cq4bi"}'.
	DimensionsJSON *string `json:"dimensionsJson,omitempty" tf:"dimensions_json,omitempty"`
}

type DimensionsObservation struct {

	// Represents a collection of dimensions of an object instance, json format.eg:'{"unInstanceId":"ins-ot3cq4bi"}'.
	// Represents a collection of dimensions of an object instance, json format.eg:'{"unInstanceId":"ins-ot3cq4bi"}'.
	DimensionsJSON *string `json:"dimensionsJson,omitempty" tf:"dimensions_json,omitempty"`

	// ID of the resource.
	// Object unique ID.
	UniqueID *string `json:"uniqueId,omitempty" tf:"unique_id,omitempty"`
}

type DimensionsParameters struct {

	// Represents a collection of dimensions of an object instance, json format.eg:'{"unInstanceId":"ins-ot3cq4bi"}'.
	// Represents a collection of dimensions of an object instance, json format.eg:'{"unInstanceId":"ins-ot3cq4bi"}'.
	// +kubebuilder:validation:Optional
	DimensionsJSON *string `json:"dimensionsJson" tf:"dimensions_json,omitempty"`
}

type PolicyBindingObjectInitParameters struct {

	// A list objects. Each element contains the following attributes:
	// A list objects. Each element contains the following attributes:
	Dimensions []DimensionsInitParameters `json:"dimensions,omitempty" tf:"dimensions,omitempty"`
}

type PolicyBindingObjectObservation struct {

	// A list objects. Each element contains the following attributes:
	// A list objects. Each element contains the following attributes:
	Dimensions []DimensionsObservation `json:"dimensions,omitempty" tf:"dimensions,omitempty"`

	// ID of the resource.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Alarm policy ID for binding objects.
	// Alarm policy ID for binding objects.
	PolicyID *string `json:"policyId,omitempty" tf:"policy_id,omitempty"`
}

type PolicyBindingObjectParameters struct {

	// A list objects. Each element contains the following attributes:
	// A list objects. Each element contains the following attributes:
	// +kubebuilder:validation:Optional
	Dimensions []DimensionsParameters `json:"dimensions,omitempty" tf:"dimensions,omitempty"`

	// Alarm policy ID for binding objects.
	// Alarm policy ID for binding objects.
	// +crossplane:generate:reference:type=AlarmPolicy
	// +kubebuilder:validation:Optional
	PolicyID *string `json:"policyId,omitempty" tf:"policy_id,omitempty"`

	// Reference to a AlarmPolicy to populate policyId.
	// +kubebuilder:validation:Optional
	PolicyIDRef *v1.Reference `json:"policyIdRef,omitempty" tf:"-"`

	// Selector for a AlarmPolicy to populate policyId.
	// +kubebuilder:validation:Optional
	PolicyIDSelector *v1.Selector `json:"policyIdSelector,omitempty" tf:"-"`
}

// PolicyBindingObjectSpec defines the desired state of PolicyBindingObject
type PolicyBindingObjectSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     PolicyBindingObjectParameters `json:"forProvider"`
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
	InitProvider PolicyBindingObjectInitParameters `json:"initProvider,omitempty"`
}

// PolicyBindingObjectStatus defines the observed state of PolicyBindingObject.
type PolicyBindingObjectStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        PolicyBindingObjectObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// PolicyBindingObject is the Schema for the PolicyBindingObjects API. Provides a resource for bind objects to a alarm policy resource.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,tencentcloud}
type PolicyBindingObject struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.dimensions) || (has(self.initProvider) && has(self.initProvider.dimensions))",message="spec.forProvider.dimensions is a required parameter"
	Spec   PolicyBindingObjectSpec   `json:"spec"`
	Status PolicyBindingObjectStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// PolicyBindingObjectList contains a list of PolicyBindingObjects
type PolicyBindingObjectList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []PolicyBindingObject `json:"items"`
}

// Repository type metadata.
var (
	PolicyBindingObject_Kind             = "PolicyBindingObject"
	PolicyBindingObject_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: PolicyBindingObject_Kind}.String()
	PolicyBindingObject_KindAPIVersion   = PolicyBindingObject_Kind + "." + CRDGroupVersion.String()
	PolicyBindingObject_GroupVersionKind = CRDGroupVersion.WithKind(PolicyBindingObject_Kind)
)

func init() {
	SchemeBuilder.Register(&PolicyBindingObject{}, &PolicyBindingObjectList{})
}
