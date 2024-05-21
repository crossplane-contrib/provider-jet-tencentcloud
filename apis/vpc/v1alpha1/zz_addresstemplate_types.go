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

type AddressTemplateInitParameters struct {

	// Address list. IP(10.0.0.1), CIDR(10.0.1.0/24), IP range(10.0.0.1-10.0.0.100) format are supported.
	// Address list. IP(`10.0.0.1`), CIDR(`10.0.1.0/24`), IP range(`10.0.0.1-10.0.0.100`) format are supported.
	// +listType=set
	Addresses []*string `json:"addresses,omitempty" tf:"addresses,omitempty"`

	// Name of the address template.
	// Name of the address template.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`
}

type AddressTemplateObservation struct {

	// Address list. IP(10.0.0.1), CIDR(10.0.1.0/24), IP range(10.0.0.1-10.0.0.100) format are supported.
	// Address list. IP(`10.0.0.1`), CIDR(`10.0.1.0/24`), IP range(`10.0.0.1-10.0.0.100`) format are supported.
	// +listType=set
	Addresses []*string `json:"addresses,omitempty" tf:"addresses,omitempty"`

	// ID of the resource.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Name of the address template.
	// Name of the address template.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`
}

type AddressTemplateParameters struct {

	// Address list. IP(10.0.0.1), CIDR(10.0.1.0/24), IP range(10.0.0.1-10.0.0.100) format are supported.
	// Address list. IP(`10.0.0.1`), CIDR(`10.0.1.0/24`), IP range(`10.0.0.1-10.0.0.100`) format are supported.
	// +kubebuilder:validation:Optional
	// +listType=set
	Addresses []*string `json:"addresses,omitempty" tf:"addresses,omitempty"`

	// Name of the address template.
	// Name of the address template.
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`
}

// AddressTemplateSpec defines the desired state of AddressTemplate
type AddressTemplateSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     AddressTemplateParameters `json:"forProvider"`
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
	InitProvider AddressTemplateInitParameters `json:"initProvider,omitempty"`
}

// AddressTemplateStatus defines the observed state of AddressTemplate.
type AddressTemplateStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        AddressTemplateObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// AddressTemplate is the Schema for the AddressTemplates API. Provides a resource to manage address template.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,tencentcloud}
type AddressTemplate struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.addresses) || (has(self.initProvider) && has(self.initProvider.addresses))",message="spec.forProvider.addresses is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.name) || (has(self.initProvider) && has(self.initProvider.name))",message="spec.forProvider.name is a required parameter"
	Spec   AddressTemplateSpec   `json:"spec"`
	Status AddressTemplateStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// AddressTemplateList contains a list of AddressTemplates
type AddressTemplateList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []AddressTemplate `json:"items"`
}

// Repository type metadata.
var (
	AddressTemplate_Kind             = "AddressTemplate"
	AddressTemplate_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: AddressTemplate_Kind}.String()
	AddressTemplate_KindAPIVersion   = AddressTemplate_Kind + "." + CRDGroupVersion.String()
	AddressTemplate_GroupVersionKind = CRDGroupVersion.WithKind(AddressTemplate_Kind)
)

func init() {
	SchemeBuilder.Register(&AddressTemplate{}, &AddressTemplateList{})
}
