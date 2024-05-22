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

type SecurityPolicyInitParameters struct {

	// Default policy. Valid value: ACCEPT and DROP.
	// Default policy. Valid value: `ACCEPT` and `DROP`.
	Action *string `json:"action,omitempty" tf:"action,omitempty"`

	// Indicates whether policy is enable, default value is true.
	// Indicates whether policy is enable, default value is `true`.
	Enable *bool `json:"enable,omitempty" tf:"enable,omitempty"`

	// ID of the GAAP proxy.
	// ID of the GAAP proxy.
	// +crossplane:generate:reference:type=Proxy
	ProxyID *string `json:"proxyId,omitempty" tf:"proxy_id,omitempty"`

	// Reference to a Proxy to populate proxyId.
	// +kubebuilder:validation:Optional
	ProxyIDRef *v1.Reference `json:"proxyIdRef,omitempty" tf:"-"`

	// Selector for a Proxy to populate proxyId.
	// +kubebuilder:validation:Optional
	ProxyIDSelector *v1.Selector `json:"proxyIdSelector,omitempty" tf:"-"`
}

type SecurityPolicyObservation struct {

	// Default policy. Valid value: ACCEPT and DROP.
	// Default policy. Valid value: `ACCEPT` and `DROP`.
	Action *string `json:"action,omitempty" tf:"action,omitempty"`

	// Indicates whether policy is enable, default value is true.
	// Indicates whether policy is enable, default value is `true`.
	Enable *bool `json:"enable,omitempty" tf:"enable,omitempty"`

	// ID of the resource.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// ID of the GAAP proxy.
	// ID of the GAAP proxy.
	ProxyID *string `json:"proxyId,omitempty" tf:"proxy_id,omitempty"`
}

type SecurityPolicyParameters struct {

	// Default policy. Valid value: ACCEPT and DROP.
	// Default policy. Valid value: `ACCEPT` and `DROP`.
	// +kubebuilder:validation:Optional
	Action *string `json:"action,omitempty" tf:"action,omitempty"`

	// Indicates whether policy is enable, default value is true.
	// Indicates whether policy is enable, default value is `true`.
	// +kubebuilder:validation:Optional
	Enable *bool `json:"enable,omitempty" tf:"enable,omitempty"`

	// ID of the GAAP proxy.
	// ID of the GAAP proxy.
	// +crossplane:generate:reference:type=Proxy
	// +kubebuilder:validation:Optional
	ProxyID *string `json:"proxyId,omitempty" tf:"proxy_id,omitempty"`

	// Reference to a Proxy to populate proxyId.
	// +kubebuilder:validation:Optional
	ProxyIDRef *v1.Reference `json:"proxyIdRef,omitempty" tf:"-"`

	// Selector for a Proxy to populate proxyId.
	// +kubebuilder:validation:Optional
	ProxyIDSelector *v1.Selector `json:"proxyIdSelector,omitempty" tf:"-"`
}

// SecurityPolicySpec defines the desired state of SecurityPolicy
type SecurityPolicySpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     SecurityPolicyParameters `json:"forProvider"`
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
	InitProvider SecurityPolicyInitParameters `json:"initProvider,omitempty"`
}

// SecurityPolicyStatus defines the observed state of SecurityPolicy.
type SecurityPolicyStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        SecurityPolicyObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// SecurityPolicy is the Schema for the SecurityPolicys API. Provides a resource to create a security policy of GAAP proxy.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,tencentcloud}
type SecurityPolicy struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.action) || (has(self.initProvider) && has(self.initProvider.action))",message="spec.forProvider.action is a required parameter"
	Spec   SecurityPolicySpec   `json:"spec"`
	Status SecurityPolicyStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// SecurityPolicyList contains a list of SecurityPolicys
type SecurityPolicyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []SecurityPolicy `json:"items"`
}

// Repository type metadata.
var (
	SecurityPolicy_Kind             = "SecurityPolicy"
	SecurityPolicy_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: SecurityPolicy_Kind}.String()
	SecurityPolicy_KindAPIVersion   = SecurityPolicy_Kind + "." + CRDGroupVersion.String()
	SecurityPolicy_GroupVersionKind = CRDGroupVersion.WithKind(SecurityPolicy_Kind)
)

func init() {
	SchemeBuilder.Register(&SecurityPolicy{}, &SecurityPolicyList{})
}
