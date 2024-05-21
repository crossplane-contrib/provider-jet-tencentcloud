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

type StrategyAttachmentInitParameters struct {

	// The API that needs to be bound.
	// The API that needs to be bound.
	// +crossplane:generate:reference:type=Api
	BindAPIID *string `json:"bindApiId,omitempty" tf:"bind_api_id,omitempty"`

	// Reference to a Api to populate bindApiId.
	// +kubebuilder:validation:Optional
	BindAPIIDRef *v1.Reference `json:"bindApiIdRef,omitempty" tf:"-"`

	// Selector for a Api to populate bindApiId.
	// +kubebuilder:validation:Optional
	BindAPIIDSelector *v1.Selector `json:"bindApiIdSelector,omitempty" tf:"-"`

	// The environment of the strategy association. Valid values: test, release, prepub.
	// The environment of the strategy association. Valid values: `test`, `release`, `prepub`.
	EnvironmentName *string `json:"environmentName,omitempty" tf:"environment_name,omitempty"`

	// The ID of the API gateway service.
	// The ID of the API gateway service.
	// +crossplane:generate:reference:type=Service
	ServiceID *string `json:"serviceId,omitempty" tf:"service_id,omitempty"`

	// Reference to a Service to populate serviceId.
	// +kubebuilder:validation:Optional
	ServiceIDRef *v1.Reference `json:"serviceIdRef,omitempty" tf:"-"`

	// Selector for a Service to populate serviceId.
	// +kubebuilder:validation:Optional
	ServiceIDSelector *v1.Selector `json:"serviceIdSelector,omitempty" tf:"-"`

	// The ID of the API gateway strategy.
	// The ID of the API gateway strategy.
	// +crossplane:generate:reference:type=IpStrategy
	StrategyID *string `json:"strategyId,omitempty" tf:"strategy_id,omitempty"`

	// Reference to a IpStrategy to populate strategyId.
	// +kubebuilder:validation:Optional
	StrategyIDRef *v1.Reference `json:"strategyIdRef,omitempty" tf:"-"`

	// Selector for a IpStrategy to populate strategyId.
	// +kubebuilder:validation:Optional
	StrategyIDSelector *v1.Selector `json:"strategyIdSelector,omitempty" tf:"-"`
}

type StrategyAttachmentObservation struct {

	// The API that needs to be bound.
	// The API that needs to be bound.
	BindAPIID *string `json:"bindApiId,omitempty" tf:"bind_api_id,omitempty"`

	// The environment of the strategy association. Valid values: test, release, prepub.
	// The environment of the strategy association. Valid values: `test`, `release`, `prepub`.
	EnvironmentName *string `json:"environmentName,omitempty" tf:"environment_name,omitempty"`

	// ID of the resource.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The ID of the API gateway service.
	// The ID of the API gateway service.
	ServiceID *string `json:"serviceId,omitempty" tf:"service_id,omitempty"`

	// The ID of the API gateway strategy.
	// The ID of the API gateway strategy.
	StrategyID *string `json:"strategyId,omitempty" tf:"strategy_id,omitempty"`
}

type StrategyAttachmentParameters struct {

	// The API that needs to be bound.
	// The API that needs to be bound.
	// +crossplane:generate:reference:type=Api
	// +kubebuilder:validation:Optional
	BindAPIID *string `json:"bindApiId,omitempty" tf:"bind_api_id,omitempty"`

	// Reference to a Api to populate bindApiId.
	// +kubebuilder:validation:Optional
	BindAPIIDRef *v1.Reference `json:"bindApiIdRef,omitempty" tf:"-"`

	// Selector for a Api to populate bindApiId.
	// +kubebuilder:validation:Optional
	BindAPIIDSelector *v1.Selector `json:"bindApiIdSelector,omitempty" tf:"-"`

	// The environment of the strategy association. Valid values: test, release, prepub.
	// The environment of the strategy association. Valid values: `test`, `release`, `prepub`.
	// +kubebuilder:validation:Optional
	EnvironmentName *string `json:"environmentName,omitempty" tf:"environment_name,omitempty"`

	// The ID of the API gateway service.
	// The ID of the API gateway service.
	// +crossplane:generate:reference:type=Service
	// +kubebuilder:validation:Optional
	ServiceID *string `json:"serviceId,omitempty" tf:"service_id,omitempty"`

	// Reference to a Service to populate serviceId.
	// +kubebuilder:validation:Optional
	ServiceIDRef *v1.Reference `json:"serviceIdRef,omitempty" tf:"-"`

	// Selector for a Service to populate serviceId.
	// +kubebuilder:validation:Optional
	ServiceIDSelector *v1.Selector `json:"serviceIdSelector,omitempty" tf:"-"`

	// The ID of the API gateway strategy.
	// The ID of the API gateway strategy.
	// +crossplane:generate:reference:type=IpStrategy
	// +kubebuilder:validation:Optional
	StrategyID *string `json:"strategyId,omitempty" tf:"strategy_id,omitempty"`

	// Reference to a IpStrategy to populate strategyId.
	// +kubebuilder:validation:Optional
	StrategyIDRef *v1.Reference `json:"strategyIdRef,omitempty" tf:"-"`

	// Selector for a IpStrategy to populate strategyId.
	// +kubebuilder:validation:Optional
	StrategyIDSelector *v1.Selector `json:"strategyIdSelector,omitempty" tf:"-"`
}

// StrategyAttachmentSpec defines the desired state of StrategyAttachment
type StrategyAttachmentSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     StrategyAttachmentParameters `json:"forProvider"`
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
	InitProvider StrategyAttachmentInitParameters `json:"initProvider,omitempty"`
}

// StrategyAttachmentStatus defines the observed state of StrategyAttachment.
type StrategyAttachmentStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        StrategyAttachmentObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// StrategyAttachment is the Schema for the StrategyAttachments API. Use this resource to create IP strategy attachment of API gateway.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,tencentcloud}
type StrategyAttachment struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.environmentName) || (has(self.initProvider) && has(self.initProvider.environmentName))",message="spec.forProvider.environmentName is a required parameter"
	Spec   StrategyAttachmentSpec   `json:"spec"`
	Status StrategyAttachmentStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// StrategyAttachmentList contains a list of StrategyAttachments
type StrategyAttachmentList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []StrategyAttachment `json:"items"`
}

// Repository type metadata.
var (
	StrategyAttachment_Kind             = "StrategyAttachment"
	StrategyAttachment_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: StrategyAttachment_Kind}.String()
	StrategyAttachment_KindAPIVersion   = StrategyAttachment_Kind + "." + CRDGroupVersion.String()
	StrategyAttachment_GroupVersionKind = CRDGroupVersion.WithKind(StrategyAttachment_Kind)
)

func init() {
	SchemeBuilder.Register(&StrategyAttachment{}, &StrategyAttachmentList{})
}
