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

type SAMLProviderInitParameters struct {

	// The description of the CAM SAML provider.
	// The description of the CAM SAML provider.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// The meta data document of the CAM SAML provider.
	// The meta data document of the CAM SAML provider.
	MetaData *string `json:"metaData,omitempty" tf:"meta_data,omitempty"`

	// Name of CAM SAML provider.
	// Name of CAM SAML provider.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`
}

type SAMLProviderObservation struct {

	// The create time of the CAM SAML provider.
	// The create time of the CAM SAML provider.
	CreateTime *string `json:"createTime,omitempty" tf:"create_time,omitempty"`

	// The description of the CAM SAML provider.
	// The description of the CAM SAML provider.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// ID of the resource.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The meta data document of the CAM SAML provider.
	// The meta data document of the CAM SAML provider.
	MetaData *string `json:"metaData,omitempty" tf:"meta_data,omitempty"`

	// Name of CAM SAML provider.
	// Name of CAM SAML provider.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The ARN of the CAM SAML provider.
	// The ARN of the CAM SAML provider.
	ProviderArn *string `json:"providerArn,omitempty" tf:"provider_arn,omitempty"`

	// The last update time of the CAM SAML provider.
	// The last update time of the CAM SAML provider.
	UpdateTime *string `json:"updateTime,omitempty" tf:"update_time,omitempty"`
}

type SAMLProviderParameters struct {

	// The description of the CAM SAML provider.
	// The description of the CAM SAML provider.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// The meta data document of the CAM SAML provider.
	// The meta data document of the CAM SAML provider.
	// +kubebuilder:validation:Optional
	MetaData *string `json:"metaData,omitempty" tf:"meta_data,omitempty"`

	// Name of CAM SAML provider.
	// Name of CAM SAML provider.
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`
}

// SAMLProviderSpec defines the desired state of SAMLProvider
type SAMLProviderSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     SAMLProviderParameters `json:"forProvider"`
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
	InitProvider SAMLProviderInitParameters `json:"initProvider,omitempty"`
}

// SAMLProviderStatus defines the observed state of SAMLProvider.
type SAMLProviderStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        SAMLProviderObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// SAMLProvider is the Schema for the SAMLProviders API. Provides a resource to create a CAM SAML provider.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,tencentcloud}
type SAMLProvider struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.description) || (has(self.initProvider) && has(self.initProvider.description))",message="spec.forProvider.description is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.metaData) || (has(self.initProvider) && has(self.initProvider.metaData))",message="spec.forProvider.metaData is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.name) || (has(self.initProvider) && has(self.initProvider.name))",message="spec.forProvider.name is a required parameter"
	Spec   SAMLProviderSpec   `json:"spec"`
	Status SAMLProviderStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// SAMLProviderList contains a list of SAMLProviders
type SAMLProviderList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []SAMLProvider `json:"items"`
}

// Repository type metadata.
var (
	SAMLProvider_Kind             = "SAMLProvider"
	SAMLProvider_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: SAMLProvider_Kind}.String()
	SAMLProvider_KindAPIVersion   = SAMLProvider_Kind + "." + CRDGroupVersion.String()
	SAMLProvider_GroupVersionKind = CRDGroupVersion.WithKind(SAMLProvider_Kind)
)

func init() {
	SchemeBuilder.Register(&SAMLProvider{}, &SAMLProviderList{})
}
