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

type L4RuleV2InitParameters struct {

	// Business of the resource that the layer 4 rule works for. Valid values: `bgpip` and `net`.
	Business *string `json:"business,omitempty" tf:"business,omitempty"`

	// Resource id.
	ResourceID *string `json:"resourceId,omitempty" tf:"resource_id,omitempty"`

	// A list of layer 4 rules. Each element contains the following attributes:
	Rules []RulesInitParameters `json:"rules,omitempty" tf:"rules,omitempty"`

	// Resource vpn.
	VPN *string `json:"vpn,omitempty" tf:"vpn,omitempty"`

	// The virtual port of the layer 4 rule.
	VirtualPort *float64 `json:"virtualPort,omitempty" tf:"virtual_port,omitempty"`
}

type L4RuleV2Observation struct {

	// Business of the resource that the layer 4 rule works for. Valid values: `bgpip` and `net`.
	Business *string `json:"business,omitempty" tf:"business,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Resource id.
	ResourceID *string `json:"resourceId,omitempty" tf:"resource_id,omitempty"`

	// A list of layer 4 rules. Each element contains the following attributes:
	Rules []RulesObservation `json:"rules,omitempty" tf:"rules,omitempty"`

	// Resource vpn.
	VPN *string `json:"vpn,omitempty" tf:"vpn,omitempty"`

	// The virtual port of the layer 4 rule.
	VirtualPort *float64 `json:"virtualPort,omitempty" tf:"virtual_port,omitempty"`
}

type L4RuleV2Parameters struct {

	// Business of the resource that the layer 4 rule works for. Valid values: `bgpip` and `net`.
	// +kubebuilder:validation:Optional
	Business *string `json:"business,omitempty" tf:"business,omitempty"`

	// Resource id.
	// +kubebuilder:validation:Optional
	ResourceID *string `json:"resourceId,omitempty" tf:"resource_id,omitempty"`

	// A list of layer 4 rules. Each element contains the following attributes:
	// +kubebuilder:validation:Optional
	Rules []RulesParameters `json:"rules,omitempty" tf:"rules,omitempty"`

	// Resource vpn.
	// +kubebuilder:validation:Optional
	VPN *string `json:"vpn,omitempty" tf:"vpn,omitempty"`

	// The virtual port of the layer 4 rule.
	// +kubebuilder:validation:Optional
	VirtualPort *float64 `json:"virtualPort,omitempty" tf:"virtual_port,omitempty"`
}

type RulesInitParameters struct {

	// session hold switch.
	KeepEnable *bool `json:"keepEnable,omitempty" tf:"keep_enable,omitempty"`

	// The keeptime of the layer 4 rule.
	Keeptime *float64 `json:"keeptime,omitempty" tf:"keeptime,omitempty"`

	// LB type of the rule, `1` for weight cycling and `2` for IP hash.
	LBType *float64 `json:"lbType,omitempty" tf:"lb_type,omitempty"`

	// Protocol of the rule.
	Protocol *string `json:"protocol,omitempty" tf:"protocol,omitempty"`

	// Corresponding regional information.
	Region *float64 `json:"region,omitempty" tf:"region,omitempty"`

	// Remove the watermark state.
	RemoveSwitch *bool `json:"removeSwitch,omitempty" tf:"remove_switch,omitempty"`

	// Name of the rule.
	RuleName *string `json:"ruleName,omitempty" tf:"rule_name,omitempty"`

	// Source list of the rule.
	SourceList []RulesSourceListInitParameters `json:"sourceList,omitempty" tf:"source_list,omitempty"`

	// The source port of the layer 4 rule.
	SourcePort *float64 `json:"sourcePort,omitempty" tf:"source_port,omitempty"`

	// Source type, `1` for source of host, `2` for source of IP.
	SourceType *float64 `json:"sourceType,omitempty" tf:"source_type,omitempty"`

	// The virtual port of the layer 4 rule.
	VirtualPort *float64 `json:"virtualPort,omitempty" tf:"virtual_port,omitempty"`
}

type RulesObservation struct {

	// session hold switch.
	KeepEnable *bool `json:"keepEnable,omitempty" tf:"keep_enable,omitempty"`

	// The keeptime of the layer 4 rule.
	Keeptime *float64 `json:"keeptime,omitempty" tf:"keeptime,omitempty"`

	// LB type of the rule, `1` for weight cycling and `2` for IP hash.
	LBType *float64 `json:"lbType,omitempty" tf:"lb_type,omitempty"`

	// Protocol of the rule.
	Protocol *string `json:"protocol,omitempty" tf:"protocol,omitempty"`

	// Corresponding regional information.
	Region *float64 `json:"region,omitempty" tf:"region,omitempty"`

	// Remove the watermark state.
	RemoveSwitch *bool `json:"removeSwitch,omitempty" tf:"remove_switch,omitempty"`

	// Name of the rule.
	RuleName *string `json:"ruleName,omitempty" tf:"rule_name,omitempty"`

	// Source list of the rule.
	SourceList []RulesSourceListObservation `json:"sourceList,omitempty" tf:"source_list,omitempty"`

	// The source port of the layer 4 rule.
	SourcePort *float64 `json:"sourcePort,omitempty" tf:"source_port,omitempty"`

	// Source type, `1` for source of host, `2` for source of IP.
	SourceType *float64 `json:"sourceType,omitempty" tf:"source_type,omitempty"`

	// The virtual port of the layer 4 rule.
	VirtualPort *float64 `json:"virtualPort,omitempty" tf:"virtual_port,omitempty"`
}

type RulesParameters struct {

	// session hold switch.
	// +kubebuilder:validation:Optional
	KeepEnable *bool `json:"keepEnable" tf:"keep_enable,omitempty"`

	// The keeptime of the layer 4 rule.
	// +kubebuilder:validation:Optional
	Keeptime *float64 `json:"keeptime" tf:"keeptime,omitempty"`

	// LB type of the rule, `1` for weight cycling and `2` for IP hash.
	// +kubebuilder:validation:Optional
	LBType *float64 `json:"lbType" tf:"lb_type,omitempty"`

	// Protocol of the rule.
	// +kubebuilder:validation:Optional
	Protocol *string `json:"protocol" tf:"protocol,omitempty"`

	// Corresponding regional information.
	// +kubebuilder:validation:Optional
	Region *float64 `json:"region" tf:"region,omitempty"`

	// Remove the watermark state.
	// +kubebuilder:validation:Optional
	RemoveSwitch *bool `json:"removeSwitch" tf:"remove_switch,omitempty"`

	// Name of the rule.
	// +kubebuilder:validation:Optional
	RuleName *string `json:"ruleName" tf:"rule_name,omitempty"`

	// Source list of the rule.
	// +kubebuilder:validation:Optional
	SourceList []RulesSourceListParameters `json:"sourceList" tf:"source_list,omitempty"`

	// The source port of the layer 4 rule.
	// +kubebuilder:validation:Optional
	SourcePort *float64 `json:"sourcePort" tf:"source_port,omitempty"`

	// Source type, `1` for source of host, `2` for source of IP.
	// +kubebuilder:validation:Optional
	SourceType *float64 `json:"sourceType" tf:"source_type,omitempty"`

	// The virtual port of the layer 4 rule.
	// +kubebuilder:validation:Optional
	VirtualPort *float64 `json:"virtualPort" tf:"virtual_port,omitempty"`
}

type RulesSourceListInitParameters struct {

	// Source IP or domain.
	Source *string `json:"source,omitempty" tf:"source,omitempty"`

	// Weight of the source.
	Weight *float64 `json:"weight,omitempty" tf:"weight,omitempty"`
}

type RulesSourceListObservation struct {

	// Source IP or domain.
	Source *string `json:"source,omitempty" tf:"source,omitempty"`

	// Weight of the source.
	Weight *float64 `json:"weight,omitempty" tf:"weight,omitempty"`
}

type RulesSourceListParameters struct {

	// Source IP or domain.
	// +kubebuilder:validation:Optional
	Source *string `json:"source" tf:"source,omitempty"`

	// Weight of the source.
	// +kubebuilder:validation:Optional
	Weight *float64 `json:"weight" tf:"weight,omitempty"`
}

// L4RuleV2Spec defines the desired state of L4RuleV2
type L4RuleV2Spec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     L4RuleV2Parameters `json:"forProvider"`
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
	InitProvider L4RuleV2InitParameters `json:"initProvider,omitempty"`
}

// L4RuleV2Status defines the observed state of L4RuleV2.
type L4RuleV2Status struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        L4RuleV2Observation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// L4RuleV2 is the Schema for the L4RuleV2s API. <no value>
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,tencentcloud}
type L4RuleV2 struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.business) || (has(self.initProvider) && has(self.initProvider.business))",message="spec.forProvider.business is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.resourceId) || (has(self.initProvider) && has(self.initProvider.resourceId))",message="spec.forProvider.resourceId is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.rules) || (has(self.initProvider) && has(self.initProvider.rules))",message="spec.forProvider.rules is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.virtualPort) || (has(self.initProvider) && has(self.initProvider.virtualPort))",message="spec.forProvider.virtualPort is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.vpn) || (has(self.initProvider) && has(self.initProvider.vpn))",message="spec.forProvider.vpn is a required parameter"
	Spec   L4RuleV2Spec   `json:"spec"`
	Status L4RuleV2Status `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// L4RuleV2List contains a list of L4RuleV2s
type L4RuleV2List struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []L4RuleV2 `json:"items"`
}

// Repository type metadata.
var (
	L4RuleV2_Kind             = "L4RuleV2"
	L4RuleV2_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: L4RuleV2_Kind}.String()
	L4RuleV2_KindAPIVersion   = L4RuleV2_Kind + "." + CRDGroupVersion.String()
	L4RuleV2_GroupVersionKind = CRDGroupVersion.WithKind(L4RuleV2_Kind)
)

func init() {
	SchemeBuilder.Register(&L4RuleV2{}, &L4RuleV2List{})
}
