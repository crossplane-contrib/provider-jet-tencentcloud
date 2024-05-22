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

type RouteEntryInitParameters struct {

	// The RouteEntry's target network segment.
	// The RouteEntry's target network segment.
	CidrBlock *string `json:"cidrBlock,omitempty" tf:"cidr_block,omitempty"`

	// The route entry's next hub. CVM instance ID or VPC router interface ID.
	// The route entry's next hub. CVM instance ID or VPC router interface ID.
	NextHub *string `json:"nextHub,omitempty" tf:"next_hub,omitempty"`

	// The next hop type. Valid values: public_gateway,vpn_gateway,sslvpn_gateway,dc_gateway,peering_connection,nat_gateway,havip,local_gateway and instance. instance points to CVM Instance.
	// The next hop type. Valid values: `public_gateway`,`vpn_gateway`,`sslvpn_gateway`,`dc_gateway`,`peering_connection`,`nat_gateway`,`havip`,`local_gateway` and `instance`. `instance` points to CVM Instance.
	NextType *string `json:"nextType,omitempty" tf:"next_type,omitempty"`

	// The ID of the route table.
	// The ID of the route table.
	// +crossplane:generate:reference:type=RouteTable
	RouteTableID *string `json:"routeTableId,omitempty" tf:"route_table_id,omitempty"`

	// Reference to a RouteTable to populate routeTableId.
	// +kubebuilder:validation:Optional
	RouteTableIDRef *v1.Reference `json:"routeTableIdRef,omitempty" tf:"-"`

	// Selector for a RouteTable to populate routeTableId.
	// +kubebuilder:validation:Optional
	RouteTableIDSelector *v1.Selector `json:"routeTableIdSelector,omitempty" tf:"-"`

	// The VPC ID.
	// The VPC ID.
	// +crossplane:generate:reference:type=VPC
	VPCID *string `json:"vpcId,omitempty" tf:"vpc_id,omitempty"`

	// Reference to a VPC to populate vpcId.
	// +kubebuilder:validation:Optional
	VPCIDRef *v1.Reference `json:"vpcIdRef,omitempty" tf:"-"`

	// Selector for a VPC to populate vpcId.
	// +kubebuilder:validation:Optional
	VPCIDSelector *v1.Selector `json:"vpcIdSelector,omitempty" tf:"-"`
}

type RouteEntryObservation struct {

	// The RouteEntry's target network segment.
	// The RouteEntry's target network segment.
	CidrBlock *string `json:"cidrBlock,omitempty" tf:"cidr_block,omitempty"`

	// ID of the resource.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The route entry's next hub. CVM instance ID or VPC router interface ID.
	// The route entry's next hub. CVM instance ID or VPC router interface ID.
	NextHub *string `json:"nextHub,omitempty" tf:"next_hub,omitempty"`

	// The next hop type. Valid values: public_gateway,vpn_gateway,sslvpn_gateway,dc_gateway,peering_connection,nat_gateway,havip,local_gateway and instance. instance points to CVM Instance.
	// The next hop type. Valid values: `public_gateway`,`vpn_gateway`,`sslvpn_gateway`,`dc_gateway`,`peering_connection`,`nat_gateway`,`havip`,`local_gateway` and `instance`. `instance` points to CVM Instance.
	NextType *string `json:"nextType,omitempty" tf:"next_type,omitempty"`

	// The ID of the route table.
	// The ID of the route table.
	RouteTableID *string `json:"routeTableId,omitempty" tf:"route_table_id,omitempty"`

	// The VPC ID.
	// The VPC ID.
	VPCID *string `json:"vpcId,omitempty" tf:"vpc_id,omitempty"`
}

type RouteEntryParameters struct {

	// The RouteEntry's target network segment.
	// The RouteEntry's target network segment.
	// +kubebuilder:validation:Optional
	CidrBlock *string `json:"cidrBlock,omitempty" tf:"cidr_block,omitempty"`

	// The route entry's next hub. CVM instance ID or VPC router interface ID.
	// The route entry's next hub. CVM instance ID or VPC router interface ID.
	// +kubebuilder:validation:Optional
	NextHub *string `json:"nextHub,omitempty" tf:"next_hub,omitempty"`

	// The next hop type. Valid values: public_gateway,vpn_gateway,sslvpn_gateway,dc_gateway,peering_connection,nat_gateway,havip,local_gateway and instance. instance points to CVM Instance.
	// The next hop type. Valid values: `public_gateway`,`vpn_gateway`,`sslvpn_gateway`,`dc_gateway`,`peering_connection`,`nat_gateway`,`havip`,`local_gateway` and `instance`. `instance` points to CVM Instance.
	// +kubebuilder:validation:Optional
	NextType *string `json:"nextType,omitempty" tf:"next_type,omitempty"`

	// The ID of the route table.
	// The ID of the route table.
	// +crossplane:generate:reference:type=RouteTable
	// +kubebuilder:validation:Optional
	RouteTableID *string `json:"routeTableId,omitempty" tf:"route_table_id,omitempty"`

	// Reference to a RouteTable to populate routeTableId.
	// +kubebuilder:validation:Optional
	RouteTableIDRef *v1.Reference `json:"routeTableIdRef,omitempty" tf:"-"`

	// Selector for a RouteTable to populate routeTableId.
	// +kubebuilder:validation:Optional
	RouteTableIDSelector *v1.Selector `json:"routeTableIdSelector,omitempty" tf:"-"`

	// The VPC ID.
	// The VPC ID.
	// +crossplane:generate:reference:type=VPC
	// +kubebuilder:validation:Optional
	VPCID *string `json:"vpcId,omitempty" tf:"vpc_id,omitempty"`

	// Reference to a VPC to populate vpcId.
	// +kubebuilder:validation:Optional
	VPCIDRef *v1.Reference `json:"vpcIdRef,omitempty" tf:"-"`

	// Selector for a VPC to populate vpcId.
	// +kubebuilder:validation:Optional
	VPCIDSelector *v1.Selector `json:"vpcIdSelector,omitempty" tf:"-"`
}

// RouteEntrySpec defines the desired state of RouteEntry
type RouteEntrySpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     RouteEntryParameters `json:"forProvider"`
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
	InitProvider RouteEntryInitParameters `json:"initProvider,omitempty"`
}

// RouteEntryStatus defines the observed state of RouteEntry.
type RouteEntryStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        RouteEntryObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// RouteEntry is the Schema for the RouteEntrys API. Provides a resource to create a routing entry in a VPC routing table.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,tencentcloud}
type RouteEntry struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.cidrBlock) || (has(self.initProvider) && has(self.initProvider.cidrBlock))",message="spec.forProvider.cidrBlock is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.nextHub) || (has(self.initProvider) && has(self.initProvider.nextHub))",message="spec.forProvider.nextHub is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.nextType) || (has(self.initProvider) && has(self.initProvider.nextType))",message="spec.forProvider.nextType is a required parameter"
	Spec   RouteEntrySpec   `json:"spec"`
	Status RouteEntryStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// RouteEntryList contains a list of RouteEntrys
type RouteEntryList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []RouteEntry `json:"items"`
}

// Repository type metadata.
var (
	RouteEntry_Kind             = "RouteEntry"
	RouteEntry_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: RouteEntry_Kind}.String()
	RouteEntry_KindAPIVersion   = RouteEntry_Kind + "." + CRDGroupVersion.String()
	RouteEntry_GroupVersionKind = CRDGroupVersion.WithKind(RouteEntry_Kind)
)

func init() {
	SchemeBuilder.Register(&RouteEntry{}, &RouteEntryList{})
}
