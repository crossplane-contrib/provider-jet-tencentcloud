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

type CustomDomainInitParameters struct {

	// Unique certificate ID of the custom domain name to be bound. You can choose to upload for the protocol attribute value https or http&https.
	// Unique certificate ID of the custom domain name to be bound. You can choose to upload for the `protocol` attribute value `https` or `http&https`.
	CertificateID *string `json:"certificateId,omitempty" tf:"certificate_id,omitempty"`

	// Default domain name.
	// Default domain name.
	DefaultDomain *string `json:"defaultDomain,omitempty" tf:"default_domain,omitempty"`

	// Whether the default path mapping is used. The default value is true. When it is false, it means custom path mapping. In this case, the path_mappings attribute is required.
	// Whether the default path mapping is used. The default value is `true`. When it is `false`, it means custom path mapping. In this case, the `path_mappings` attribute is required.
	IsDefaultMapping *bool `json:"isDefaultMapping,omitempty" tf:"is_default_mapping,omitempty"`

	// Whether to force HTTP requests to jump to HTTPS, default to false. When the parameter is true, the API gateway will redirect all HTTP protocol requests using the custom domain name to the HTTPS protocol for forwarding.
	// Whether to force HTTP requests to jump to HTTPS, default to false. When the parameter is true, the API gateway will redirect all HTTP protocol requests using the custom domain name to the HTTPS protocol for forwarding.
	IsForcedHTTPS *bool `json:"isForcedHttps,omitempty" tf:"is_forced_https,omitempty"`

	// Network type. Valid values: OUTER, INNER.
	// Network type. Valid values: `OUTER`, `INNER`.
	NetType *string `json:"netType,omitempty" tf:"net_type,omitempty"`

	// Custom domain name path mapping. The data format is: path#environment. Optional values for the environment are test, prepub, and release.
	// Custom domain name path mapping. The data format is: `path#environment`. Optional values for the environment are `test`, `prepub`, and `release`.
	PathMappings []*string `json:"pathMappings,omitempty" tf:"path_mappings,omitempty"`

	// Protocol supported by service. Valid values: http, https, http&https.
	// Protocol supported by service. Valid values: `http`, `https`, `http&https`.
	Protocol *string `json:"protocol,omitempty" tf:"protocol,omitempty"`

	// Custom domain name to be bound.
	// Custom domain name to be bound.
	SubDomain *string `json:"subDomain,omitempty" tf:"sub_domain,omitempty"`
}

type CustomDomainObservation struct {

	// Unique certificate ID of the custom domain name to be bound. You can choose to upload for the protocol attribute value https or http&https.
	// Unique certificate ID of the custom domain name to be bound. You can choose to upload for the `protocol` attribute value `https` or `http&https`.
	CertificateID *string `json:"certificateId,omitempty" tf:"certificate_id,omitempty"`

	// Default domain name.
	// Default domain name.
	DefaultDomain *string `json:"defaultDomain,omitempty" tf:"default_domain,omitempty"`

	// ID of the resource.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Whether the default path mapping is used. The default value is true. When it is false, it means custom path mapping. In this case, the path_mappings attribute is required.
	// Whether the default path mapping is used. The default value is `true`. When it is `false`, it means custom path mapping. In this case, the `path_mappings` attribute is required.
	IsDefaultMapping *bool `json:"isDefaultMapping,omitempty" tf:"is_default_mapping,omitempty"`

	// Whether to force HTTP requests to jump to HTTPS, default to false. When the parameter is true, the API gateway will redirect all HTTP protocol requests using the custom domain name to the HTTPS protocol for forwarding.
	// Whether to force HTTP requests to jump to HTTPS, default to false. When the parameter is true, the API gateway will redirect all HTTP protocol requests using the custom domain name to the HTTPS protocol for forwarding.
	IsForcedHTTPS *bool `json:"isForcedHttps,omitempty" tf:"is_forced_https,omitempty"`

	// Network type. Valid values: OUTER, INNER.
	// Network type. Valid values: `OUTER`, `INNER`.
	NetType *string `json:"netType,omitempty" tf:"net_type,omitempty"`

	// Custom domain name path mapping. The data format is: path#environment. Optional values for the environment are test, prepub, and release.
	// Custom domain name path mapping. The data format is: `path#environment`. Optional values for the environment are `test`, `prepub`, and `release`.
	PathMappings []*string `json:"pathMappings,omitempty" tf:"path_mappings,omitempty"`

	// Protocol supported by service. Valid values: http, https, http&https.
	// Protocol supported by service. Valid values: `http`, `https`, `http&https`.
	Protocol *string `json:"protocol,omitempty" tf:"protocol,omitempty"`

	// Unique service ID.
	// Unique service ID.
	ServiceID *string `json:"serviceId,omitempty" tf:"service_id,omitempty"`

	// Domain name resolution status. 1 means normal analysis, 0 means parsing failed.
	// Domain name resolution status. `1` means normal analysis, `0` means parsing failed.
	Status *float64 `json:"status,omitempty" tf:"status,omitempty"`

	// Custom domain name to be bound.
	// Custom domain name to be bound.
	SubDomain *string `json:"subDomain,omitempty" tf:"sub_domain,omitempty"`
}

type CustomDomainParameters struct {

	// Unique certificate ID of the custom domain name to be bound. You can choose to upload for the protocol attribute value https or http&https.
	// Unique certificate ID of the custom domain name to be bound. You can choose to upload for the `protocol` attribute value `https` or `http&https`.
	// +kubebuilder:validation:Optional
	CertificateID *string `json:"certificateId,omitempty" tf:"certificate_id,omitempty"`

	// Default domain name.
	// Default domain name.
	// +kubebuilder:validation:Optional
	DefaultDomain *string `json:"defaultDomain,omitempty" tf:"default_domain,omitempty"`

	// Whether the default path mapping is used. The default value is true. When it is false, it means custom path mapping. In this case, the path_mappings attribute is required.
	// Whether the default path mapping is used. The default value is `true`. When it is `false`, it means custom path mapping. In this case, the `path_mappings` attribute is required.
	// +kubebuilder:validation:Optional
	IsDefaultMapping *bool `json:"isDefaultMapping,omitempty" tf:"is_default_mapping,omitempty"`

	// Whether to force HTTP requests to jump to HTTPS, default to false. When the parameter is true, the API gateway will redirect all HTTP protocol requests using the custom domain name to the HTTPS protocol for forwarding.
	// Whether to force HTTP requests to jump to HTTPS, default to false. When the parameter is true, the API gateway will redirect all HTTP protocol requests using the custom domain name to the HTTPS protocol for forwarding.
	// +kubebuilder:validation:Optional
	IsForcedHTTPS *bool `json:"isForcedHttps,omitempty" tf:"is_forced_https,omitempty"`

	// Network type. Valid values: OUTER, INNER.
	// Network type. Valid values: `OUTER`, `INNER`.
	// +kubebuilder:validation:Optional
	NetType *string `json:"netType,omitempty" tf:"net_type,omitempty"`

	// Custom domain name path mapping. The data format is: path#environment. Optional values for the environment are test, prepub, and release.
	// Custom domain name path mapping. The data format is: `path#environment`. Optional values for the environment are `test`, `prepub`, and `release`.
	// +kubebuilder:validation:Optional
	PathMappings []*string `json:"pathMappings,omitempty" tf:"path_mappings,omitempty"`

	// Protocol supported by service. Valid values: http, https, http&https.
	// Protocol supported by service. Valid values: `http`, `https`, `http&https`.
	// +kubebuilder:validation:Optional
	Protocol *string `json:"protocol,omitempty" tf:"protocol,omitempty"`

	// Unique service ID.
	// Unique service ID.
	// +crossplane:generate:reference:type=Service
	// +kubebuilder:validation:Optional
	ServiceID *string `json:"serviceId,omitempty" tf:"service_id,omitempty"`

	// Reference to a Service to populate serviceId.
	// +kubebuilder:validation:Optional
	ServiceIDRef *v1.Reference `json:"serviceIdRef,omitempty" tf:"-"`

	// Selector for a Service to populate serviceId.
	// +kubebuilder:validation:Optional
	ServiceIDSelector *v1.Selector `json:"serviceIdSelector,omitempty" tf:"-"`

	// Custom domain name to be bound.
	// Custom domain name to be bound.
	// +kubebuilder:validation:Optional
	SubDomain *string `json:"subDomain,omitempty" tf:"sub_domain,omitempty"`
}

// CustomDomainSpec defines the desired state of CustomDomain
type CustomDomainSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     CustomDomainParameters `json:"forProvider"`
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
	InitProvider CustomDomainInitParameters `json:"initProvider,omitempty"`
}

// CustomDomainStatus defines the observed state of CustomDomain.
type CustomDomainStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        CustomDomainObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// CustomDomain is the Schema for the CustomDomains API. Use this resource to create custom domain of API gateway.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,tencentcloud}
type CustomDomain struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.defaultDomain) || (has(self.initProvider) && has(self.initProvider.defaultDomain))",message="spec.forProvider.defaultDomain is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.netType) || (has(self.initProvider) && has(self.initProvider.netType))",message="spec.forProvider.netType is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.protocol) || (has(self.initProvider) && has(self.initProvider.protocol))",message="spec.forProvider.protocol is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.subDomain) || (has(self.initProvider) && has(self.initProvider.subDomain))",message="spec.forProvider.subDomain is a required parameter"
	Spec   CustomDomainSpec   `json:"spec"`
	Status CustomDomainStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// CustomDomainList contains a list of CustomDomains
type CustomDomainList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []CustomDomain `json:"items"`
}

// Repository type metadata.
var (
	CustomDomain_Kind             = "CustomDomain"
	CustomDomain_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: CustomDomain_Kind}.String()
	CustomDomain_KindAPIVersion   = CustomDomain_Kind + "." + CRDGroupVersion.String()
	CustomDomain_GroupVersionKind = CRDGroupVersion.WithKind(CustomDomain_Kind)
)

func init() {
	SchemeBuilder.Register(&CustomDomain{}, &CustomDomainList{})
}
