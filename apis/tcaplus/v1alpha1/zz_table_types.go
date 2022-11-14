/*
Copyright 2021 The Crossplane Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by terrajet. DO NOT EDIT.

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type TableObservation struct {
	CreateTime *string `json:"createTime,omitempty" tf:"create_time,omitempty"`

	Error *string `json:"error,omitempty" tf:"error,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	Status *string `json:"status,omitempty" tf:"status,omitempty"`

	TableSize *float64 `json:"tableSize,omitempty" tf:"table_size,omitempty"`
}

type TableParameters struct {

	// ID of the TcaplusDB cluster to which the table belongs.
	// +crossplane:generate:reference:type=Cluster
	// +kubebuilder:validation:Optional
	ClusterID *string `json:"clusterId,omitempty" tf:"cluster_id,omitempty"`

	// +kubebuilder:validation:Optional
	ClusterIDRef *v1.Reference `json:"clusterIdRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	ClusterIDSelector *v1.Selector `json:"clusterIdSelector,omitempty" tf:"-"`

	// Description of the TcaplusDB table.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// ID of the IDL File.
	// +kubebuilder:validation:Required
	IdlID *string `json:"idlId" tf:"idl_id,omitempty"`

	// Reserved read capacity units of the TcaplusDB table.
	// +kubebuilder:validation:Required
	ReservedReadCu *float64 `json:"reservedReadCu" tf:"reserved_read_cu,omitempty"`

	// Reserved storage capacity of the TcaplusDB table (unit: GB).
	// +kubebuilder:validation:Required
	ReservedVolume *float64 `json:"reservedVolume" tf:"reserved_volume,omitempty"`

	// Reserved write capacity units of the TcaplusDB table.
	// +kubebuilder:validation:Required
	ReservedWriteCu *float64 `json:"reservedWriteCu" tf:"reserved_write_cu,omitempty"`

	// IDL type of the TcaplusDB table. Valid values: `PROTO` and `TDR`.
	// +kubebuilder:validation:Required
	TableIdlType *string `json:"tableIdlType" tf:"table_idl_type,omitempty"`

	// Name of the TcaplusDB table.
	// +kubebuilder:validation:Required
	TableName *string `json:"tableName" tf:"table_name,omitempty"`

	// Type of the TcaplusDB table. Valid values are `GENERIC` and `LIST`.
	// +kubebuilder:validation:Required
	TableType *string `json:"tableType" tf:"table_type,omitempty"`

	// ID of the table group to which the table belongs.
	// +crossplane:generate:reference:type=TableGroup
	// +kubebuilder:validation:Optional
	TablegroupID *string `json:"tablegroupId,omitempty" tf:"tablegroup_id,omitempty"`

	// +kubebuilder:validation:Optional
	TablegroupIDRef *v1.Reference `json:"tablegroupIdRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	TablegroupIDSelector *v1.Selector `json:"tablegroupIdSelector,omitempty" tf:"-"`
}

// TableSpec defines the desired state of Table
type TableSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     TableParameters `json:"forProvider"`
}

// TableStatus defines the observed state of Table.
type TableStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        TableObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Table is the Schema for the Tables API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,tencentcloudjet}
type Table struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              TableSpec   `json:"spec"`
	Status            TableStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// TableList contains a list of Tables
type TableList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Table `json:"items"`
}

// Repository type metadata.
var (
	Table_Kind             = "Table"
	Table_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Table_Kind}.String()
	Table_KindAPIVersion   = Table_Kind + "." + CRDGroupVersion.String()
	Table_GroupVersionKind = CRDGroupVersion.WithKind(Table_Kind)
)

func init() {
	SchemeBuilder.Register(&Table{}, &TableList{})
}
