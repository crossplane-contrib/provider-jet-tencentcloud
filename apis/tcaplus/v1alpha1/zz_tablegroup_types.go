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

type TableGroupInitParameters struct {

	// ID of the TcaplusDB cluster to which the table group belongs.
	// ID of the TcaplusDB cluster to which the table group belongs.
	// +crossplane:generate:reference:type=Cluster
	ClusterID *string `json:"clusterId,omitempty" tf:"cluster_id,omitempty"`

	// Reference to a Cluster to populate clusterId.
	// +kubebuilder:validation:Optional
	ClusterIDRef *v1.Reference `json:"clusterIdRef,omitempty" tf:"-"`

	// Selector for a Cluster to populate clusterId.
	// +kubebuilder:validation:Optional
	ClusterIDSelector *v1.Selector `json:"clusterIdSelector,omitempty" tf:"-"`

	// Name of the TcaplusDB table group. Name length should be between 1 and 30.
	// Name of the TcaplusDB table group. Name length should be between 1 and 30.
	TablegroupName *string `json:"tablegroupName,omitempty" tf:"tablegroup_name,omitempty"`
}

type TableGroupObservation struct {

	// ID of the TcaplusDB cluster to which the table group belongs.
	// ID of the TcaplusDB cluster to which the table group belongs.
	ClusterID *string `json:"clusterId,omitempty" tf:"cluster_id,omitempty"`

	// Create time of the TcaplusDB table group.
	// Create time of the TcaplusDB table group.
	CreateTime *string `json:"createTime,omitempty" tf:"create_time,omitempty"`

	// ID of the resource.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Number of tables.
	// Number of tables.
	TableCount *float64 `json:"tableCount,omitempty" tf:"table_count,omitempty"`

	// Name of the TcaplusDB table group. Name length should be between 1 and 30.
	// Name of the TcaplusDB table group. Name length should be between 1 and 30.
	TablegroupName *string `json:"tablegroupName,omitempty" tf:"tablegroup_name,omitempty"`

	// Total storage size (MB).
	// Total storage size (MB).
	TotalSize *float64 `json:"totalSize,omitempty" tf:"total_size,omitempty"`
}

type TableGroupParameters struct {

	// ID of the TcaplusDB cluster to which the table group belongs.
	// ID of the TcaplusDB cluster to which the table group belongs.
	// +crossplane:generate:reference:type=Cluster
	// +kubebuilder:validation:Optional
	ClusterID *string `json:"clusterId,omitempty" tf:"cluster_id,omitempty"`

	// Reference to a Cluster to populate clusterId.
	// +kubebuilder:validation:Optional
	ClusterIDRef *v1.Reference `json:"clusterIdRef,omitempty" tf:"-"`

	// Selector for a Cluster to populate clusterId.
	// +kubebuilder:validation:Optional
	ClusterIDSelector *v1.Selector `json:"clusterIdSelector,omitempty" tf:"-"`

	// Name of the TcaplusDB table group. Name length should be between 1 and 30.
	// Name of the TcaplusDB table group. Name length should be between 1 and 30.
	// +kubebuilder:validation:Optional
	TablegroupName *string `json:"tablegroupName,omitempty" tf:"tablegroup_name,omitempty"`
}

// TableGroupSpec defines the desired state of TableGroup
type TableGroupSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     TableGroupParameters `json:"forProvider"`
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
	InitProvider TableGroupInitParameters `json:"initProvider,omitempty"`
}

// TableGroupStatus defines the observed state of TableGroup.
type TableGroupStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        TableGroupObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// TableGroup is the Schema for the TableGroups API. Use this resource to create TcaplusDB table group.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,tencentcloud}
type TableGroup struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.tablegroupName) || (has(self.initProvider) && has(self.initProvider.tablegroupName))",message="spec.forProvider.tablegroupName is a required parameter"
	Spec   TableGroupSpec   `json:"spec"`
	Status TableGroupStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// TableGroupList contains a list of TableGroups
type TableGroupList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []TableGroup `json:"items"`
}

// Repository type metadata.
var (
	TableGroup_Kind             = "TableGroup"
	TableGroup_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: TableGroup_Kind}.String()
	TableGroup_KindAPIVersion   = TableGroup_Kind + "." + CRDGroupVersion.String()
	TableGroup_GroupVersionKind = CRDGroupVersion.WithKind(TableGroup_Kind)
)

func init() {
	SchemeBuilder.Register(&TableGroup{}, &TableGroupList{})
}
