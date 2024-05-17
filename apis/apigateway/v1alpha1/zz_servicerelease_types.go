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

type ServiceReleaseInitParameters struct {

	// API gateway service environment name to be released. Valid values: test, prepub, release.
	// API gateway service environment name to be released. Valid values: `test`, `prepub`, `release`.
	EnvironmentName *string `json:"environmentName,omitempty" tf:"environment_name,omitempty"`

	// This release description of the API gateway service.
	// This release description of the API gateway service.
	ReleaseDesc *string `json:"releaseDesc,omitempty" tf:"release_desc,omitempty"`

	// The release version.
	// The release version.
	ReleaseVersion *string `json:"releaseVersion,omitempty" tf:"release_version,omitempty"`
}

type ServiceReleaseObservation struct {

	// API gateway service environment name to be released. Valid values: test, prepub, release.
	// API gateway service environment name to be released. Valid values: `test`, `prepub`, `release`.
	EnvironmentName *string `json:"environmentName,omitempty" tf:"environment_name,omitempty"`

	// ID of the resource.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// This release description of the API gateway service.
	// This release description of the API gateway service.
	ReleaseDesc *string `json:"releaseDesc,omitempty" tf:"release_desc,omitempty"`

	// The release version.
	// The release version.
	ReleaseVersion *string `json:"releaseVersion,omitempty" tf:"release_version,omitempty"`

	// ID of API gateway service.
	// ID of API gateway service.
	ServiceID *string `json:"serviceId,omitempty" tf:"service_id,omitempty"`
}

type ServiceReleaseParameters struct {

	// API gateway service environment name to be released. Valid values: test, prepub, release.
	// API gateway service environment name to be released. Valid values: `test`, `prepub`, `release`.
	// +kubebuilder:validation:Optional
	EnvironmentName *string `json:"environmentName,omitempty" tf:"environment_name,omitempty"`

	// This release description of the API gateway service.
	// This release description of the API gateway service.
	// +kubebuilder:validation:Optional
	ReleaseDesc *string `json:"releaseDesc,omitempty" tf:"release_desc,omitempty"`

	// The release version.
	// The release version.
	// +kubebuilder:validation:Optional
	ReleaseVersion *string `json:"releaseVersion,omitempty" tf:"release_version,omitempty"`

	// ID of API gateway service.
	// ID of API gateway service.
	// +crossplane:generate:reference:type=Service
	// +kubebuilder:validation:Optional
	ServiceID *string `json:"serviceId,omitempty" tf:"service_id,omitempty"`

	// Reference to a Service to populate serviceId.
	// +kubebuilder:validation:Optional
	ServiceIDRef *v1.Reference `json:"serviceIdRef,omitempty" tf:"-"`

	// Selector for a Service to populate serviceId.
	// +kubebuilder:validation:Optional
	ServiceIDSelector *v1.Selector `json:"serviceIdSelector,omitempty" tf:"-"`
}

// ServiceReleaseSpec defines the desired state of ServiceRelease
type ServiceReleaseSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ServiceReleaseParameters `json:"forProvider"`
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
	InitProvider ServiceReleaseInitParameters `json:"initProvider,omitempty"`
}

// ServiceReleaseStatus defines the observed state of ServiceRelease.
type ServiceReleaseStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ServiceReleaseObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// ServiceRelease is the Schema for the ServiceReleases API. Use this resource to create API gateway service release.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,tencentcloud}
type ServiceRelease struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.environmentName) || (has(self.initProvider) && has(self.initProvider.environmentName))",message="spec.forProvider.environmentName is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.releaseDesc) || (has(self.initProvider) && has(self.initProvider.releaseDesc))",message="spec.forProvider.releaseDesc is a required parameter"
	Spec   ServiceReleaseSpec   `json:"spec"`
	Status ServiceReleaseStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ServiceReleaseList contains a list of ServiceReleases
type ServiceReleaseList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ServiceRelease `json:"items"`
}

// Repository type metadata.
var (
	ServiceRelease_Kind             = "ServiceRelease"
	ServiceRelease_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: ServiceRelease_Kind}.String()
	ServiceRelease_KindAPIVersion   = ServiceRelease_Kind + "." + CRDGroupVersion.String()
	ServiceRelease_GroupVersionKind = CRDGroupVersion.WithKind(ServiceRelease_Kind)
)

func init() {
	SchemeBuilder.Register(&ServiceRelease{}, &ServiceReleaseList{})
}
