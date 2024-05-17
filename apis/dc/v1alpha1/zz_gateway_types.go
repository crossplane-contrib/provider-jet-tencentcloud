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

type GatewayInitParameters struct {

	// Type of the gateway. Valid value: NORMAL and NAT. Default is NORMAL. NOTES: CCN only supports NORMAL and a VPC can create two DCGs, the one is NAT type and the other is non-NAT type.
	// Type of the gateway. Valid value: `NORMAL` and `NAT`. Default is `NORMAL`. NOTES: CCN only supports `NORMAL` and a VPC can create two DCGs, the one is NAT type and the other is non-NAT type.
	GatewayType *string `json:"gatewayType,omitempty" tf:"gateway_type,omitempty"`

	// Name of the DCG.
	// Name of the DCG.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// If the network_type value is VPC, the available value is VPC ID. But when the network_type value is CCN, the available value is CCN instance ID.
	// If the `network_type` value is `VPC`, the available value is VPC ID. But when the `network_type` value is `CCN`, the available value is CCN instance ID.
	NetworkInstanceID *string `json:"networkInstanceId,omitempty" tf:"network_instance_id,omitempty"`

	// Type of associated network. Valid value: VPC and CCN.
	// Type of associated network. Valid value: `VPC` and `CCN`.
	NetworkType *string `json:"networkType,omitempty" tf:"network_type,omitempty"`
}

type GatewayObservation struct {

	// Type of CCN route. Valid value: BGP and STATIC. The property is available when the DCG type is CCN gateway and BGP enabled.
	// Type of CCN route. Valid value: `BGP` and `STATIC`. The property is available when the DCG type is CCN gateway and BGP enabled.
	CnnRouteType *string `json:"cnnRouteType,omitempty" tf:"cnn_route_type,omitempty"`

	// Creation time of resource.
	// Creation time of resource.
	CreateTime *string `json:"createTime,omitempty" tf:"create_time,omitempty"`

	// Indicates whether the BGP is enabled.
	// Indicates whether the BGP is enabled.
	EnableBGP *bool `json:"enableBgp,omitempty" tf:"enable_bgp,omitempty"`

	// Type of the gateway. Valid value: NORMAL and NAT. Default is NORMAL. NOTES: CCN only supports NORMAL and a VPC can create two DCGs, the one is NAT type and the other is non-NAT type.
	// Type of the gateway. Valid value: `NORMAL` and `NAT`. Default is `NORMAL`. NOTES: CCN only supports `NORMAL` and a VPC can create two DCGs, the one is NAT type and the other is non-NAT type.
	GatewayType *string `json:"gatewayType,omitempty" tf:"gateway_type,omitempty"`

	// ID of the resource.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Name of the DCG.
	// Name of the DCG.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// If the network_type value is VPC, the available value is VPC ID. But when the network_type value is CCN, the available value is CCN instance ID.
	// If the `network_type` value is `VPC`, the available value is VPC ID. But when the `network_type` value is `CCN`, the available value is CCN instance ID.
	NetworkInstanceID *string `json:"networkInstanceId,omitempty" tf:"network_instance_id,omitempty"`

	// Type of associated network. Valid value: VPC and CCN.
	// Type of associated network. Valid value: `VPC` and `CCN`.
	NetworkType *string `json:"networkType,omitempty" tf:"network_type,omitempty"`
}

type GatewayParameters struct {

	// Type of the gateway. Valid value: NORMAL and NAT. Default is NORMAL. NOTES: CCN only supports NORMAL and a VPC can create two DCGs, the one is NAT type and the other is non-NAT type.
	// Type of the gateway. Valid value: `NORMAL` and `NAT`. Default is `NORMAL`. NOTES: CCN only supports `NORMAL` and a VPC can create two DCGs, the one is NAT type and the other is non-NAT type.
	// +kubebuilder:validation:Optional
	GatewayType *string `json:"gatewayType,omitempty" tf:"gateway_type,omitempty"`

	// Name of the DCG.
	// Name of the DCG.
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// If the network_type value is VPC, the available value is VPC ID. But when the network_type value is CCN, the available value is CCN instance ID.
	// If the `network_type` value is `VPC`, the available value is VPC ID. But when the `network_type` value is `CCN`, the available value is CCN instance ID.
	// +kubebuilder:validation:Optional
	NetworkInstanceID *string `json:"networkInstanceId,omitempty" tf:"network_instance_id,omitempty"`

	// Type of associated network. Valid value: VPC and CCN.
	// Type of associated network. Valid value: `VPC` and `CCN`.
	// +kubebuilder:validation:Optional
	NetworkType *string `json:"networkType,omitempty" tf:"network_type,omitempty"`
}

// GatewaySpec defines the desired state of Gateway
type GatewaySpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     GatewayParameters `json:"forProvider"`
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
	InitProvider GatewayInitParameters `json:"initProvider,omitempty"`
}

// GatewayStatus defines the observed state of Gateway.
type GatewayStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        GatewayObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Gateway is the Schema for the Gateways API. Provides a resource to creating direct connect gateway instance.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,tencentcloud}
type Gateway struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.name) || (has(self.initProvider) && has(self.initProvider.name))",message="spec.forProvider.name is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.networkInstanceId) || (has(self.initProvider) && has(self.initProvider.networkInstanceId))",message="spec.forProvider.networkInstanceId is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.networkType) || (has(self.initProvider) && has(self.initProvider.networkType))",message="spec.forProvider.networkType is a required parameter"
	Spec   GatewaySpec   `json:"spec"`
	Status GatewayStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// GatewayList contains a list of Gateways
type GatewayList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Gateway `json:"items"`
}

// Repository type metadata.
var (
	Gateway_Kind             = "Gateway"
	Gateway_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Gateway_Kind}.String()
	Gateway_KindAPIVersion   = Gateway_Kind + "." + CRDGroupVersion.String()
	Gateway_GroupVersionKind = CRDGroupVersion.WithKind(Gateway_Kind)
)

func init() {
	SchemeBuilder.Register(&Gateway{}, &GatewayList{})
}
