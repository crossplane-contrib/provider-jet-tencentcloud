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

type ScfNamespaceInitParameters struct {

	// Description of the SCF namespace.
	// Description of the SCF namespace.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Name of the SCF namespace.
	// Name of the SCF namespace.
	Namespace *string `json:"namespace,omitempty" tf:"namespace,omitempty"`
}

type ScfNamespaceObservation struct {

	// SCF namespace creation time.
	// SCF namespace creation time.
	CreateTime *string `json:"createTime,omitempty" tf:"create_time,omitempty"`

	// Description of the SCF namespace.
	// Description of the SCF namespace.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// ID of the resource.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// SCF namespace last modified time.
	// SCF namespace last modified time.
	ModifyTime *string `json:"modifyTime,omitempty" tf:"modify_time,omitempty"`

	// Name of the SCF namespace.
	// Name of the SCF namespace.
	Namespace *string `json:"namespace,omitempty" tf:"namespace,omitempty"`

	// SCF namespace type.
	// SCF namespace type.
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

type ScfNamespaceParameters struct {

	// Description of the SCF namespace.
	// Description of the SCF namespace.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Name of the SCF namespace.
	// Name of the SCF namespace.
	// +kubebuilder:validation:Optional
	Namespace *string `json:"namespace,omitempty" tf:"namespace,omitempty"`
}

// ScfNamespaceSpec defines the desired state of ScfNamespace
type ScfNamespaceSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ScfNamespaceParameters `json:"forProvider"`
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
	InitProvider ScfNamespaceInitParameters `json:"initProvider,omitempty"`
}

// ScfNamespaceStatus defines the observed state of ScfNamespace.
type ScfNamespaceStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ScfNamespaceObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// ScfNamespace is the Schema for the ScfNamespaces API. Provide a resource to create a SCF namespace.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,tencentcloud}
type ScfNamespace struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.__namespace__) || (has(self.initProvider) && has(self.initProvider.__namespace__))",message="spec.forProvider.namespace is a required parameter"
	Spec   ScfNamespaceSpec   `json:"spec"`
	Status ScfNamespaceStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ScfNamespaceList contains a list of ScfNamespaces
type ScfNamespaceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ScfNamespace `json:"items"`
}

// Repository type metadata.
var (
	ScfNamespace_Kind             = "ScfNamespace"
	ScfNamespace_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: ScfNamespace_Kind}.String()
	ScfNamespace_KindAPIVersion   = ScfNamespace_Kind + "." + CRDGroupVersion.String()
	ScfNamespace_GroupVersionKind = CRDGroupVersion.WithKind(ScfNamespace_Kind)
)

func init() {
	SchemeBuilder.Register(&ScfNamespace{}, &ScfNamespaceList{})
}
