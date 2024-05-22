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

type IdlInitParameters struct {

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

	// IDL file content of the TcaplusDB table.
	// IDL file content of the TcaplusDB table.
	FileContent *string `json:"fileContent,omitempty" tf:"file_content,omitempty"`

	// File ext type of the IDL file. If file_type is PROTO, file_ext_type must be 'proto'; If file_type is TDR, file_ext_type must be 'xml'.
	// File ext type of the IDL file. If `file_type` is `PROTO`, `file_ext_type` must be 'proto'; If `file_type` is `TDR`, `file_ext_type` must be 'xml'.
	FileExtType *string `json:"fileExtType,omitempty" tf:"file_ext_type,omitempty"`

	// Name of the IDL file.
	// Name of the IDL file.
	FileName *string `json:"fileName,omitempty" tf:"file_name,omitempty"`

	// Type of the IDL file. Valid values are PROTO and TDR.
	// Type of the IDL file. Valid values are PROTO and TDR.
	FileType *string `json:"fileType,omitempty" tf:"file_type,omitempty"`

	// ID of the table group to which the IDL file belongs.
	// ID of the table group to which the IDL file belongs.
	// +crossplane:generate:reference:type=TableGroup
	TablegroupID *string `json:"tablegroupId,omitempty" tf:"tablegroup_id,omitempty"`

	// Reference to a TableGroup to populate tablegroupId.
	// +kubebuilder:validation:Optional
	TablegroupIDRef *v1.Reference `json:"tablegroupIdRef,omitempty" tf:"-"`

	// Selector for a TableGroup to populate tablegroupId.
	// +kubebuilder:validation:Optional
	TablegroupIDSelector *v1.Selector `json:"tablegroupIdSelector,omitempty" tf:"-"`
}

type IdlObservation struct {

	// ID of the TcaplusDB cluster to which the table group belongs.
	// ID of the TcaplusDB cluster to which the table group belongs.
	ClusterID *string `json:"clusterId,omitempty" tf:"cluster_id,omitempty"`

	// IDL file content of the TcaplusDB table.
	// IDL file content of the TcaplusDB table.
	FileContent *string `json:"fileContent,omitempty" tf:"file_content,omitempty"`

	// File ext type of the IDL file. If file_type is PROTO, file_ext_type must be 'proto'; If file_type is TDR, file_ext_type must be 'xml'.
	// File ext type of the IDL file. If `file_type` is `PROTO`, `file_ext_type` must be 'proto'; If `file_type` is `TDR`, `file_ext_type` must be 'xml'.
	FileExtType *string `json:"fileExtType,omitempty" tf:"file_ext_type,omitempty"`

	// Name of the IDL file.
	// Name of the IDL file.
	FileName *string `json:"fileName,omitempty" tf:"file_name,omitempty"`

	// Type of the IDL file. Valid values are PROTO and TDR.
	// Type of the IDL file. Valid values are PROTO and TDR.
	FileType *string `json:"fileType,omitempty" tf:"file_type,omitempty"`

	// ID of the resource.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Table info of the IDL.
	// Table info of the IDL.
	TableInfos []TableInfosObservation `json:"tableInfos,omitempty" tf:"table_infos,omitempty"`

	// ID of the table group to which the IDL file belongs.
	// ID of the table group to which the IDL file belongs.
	TablegroupID *string `json:"tablegroupId,omitempty" tf:"tablegroup_id,omitempty"`
}

type IdlParameters struct {

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

	// IDL file content of the TcaplusDB table.
	// IDL file content of the TcaplusDB table.
	// +kubebuilder:validation:Optional
	FileContent *string `json:"fileContent,omitempty" tf:"file_content,omitempty"`

	// File ext type of the IDL file. If file_type is PROTO, file_ext_type must be 'proto'; If file_type is TDR, file_ext_type must be 'xml'.
	// File ext type of the IDL file. If `file_type` is `PROTO`, `file_ext_type` must be 'proto'; If `file_type` is `TDR`, `file_ext_type` must be 'xml'.
	// +kubebuilder:validation:Optional
	FileExtType *string `json:"fileExtType,omitempty" tf:"file_ext_type,omitempty"`

	// Name of the IDL file.
	// Name of the IDL file.
	// +kubebuilder:validation:Optional
	FileName *string `json:"fileName,omitempty" tf:"file_name,omitempty"`

	// Type of the IDL file. Valid values are PROTO and TDR.
	// Type of the IDL file. Valid values are PROTO and TDR.
	// +kubebuilder:validation:Optional
	FileType *string `json:"fileType,omitempty" tf:"file_type,omitempty"`

	// ID of the table group to which the IDL file belongs.
	// ID of the table group to which the IDL file belongs.
	// +crossplane:generate:reference:type=TableGroup
	// +kubebuilder:validation:Optional
	TablegroupID *string `json:"tablegroupId,omitempty" tf:"tablegroup_id,omitempty"`

	// Reference to a TableGroup to populate tablegroupId.
	// +kubebuilder:validation:Optional
	TablegroupIDRef *v1.Reference `json:"tablegroupIdRef,omitempty" tf:"-"`

	// Selector for a TableGroup to populate tablegroupId.
	// +kubebuilder:validation:Optional
	TablegroupIDSelector *v1.Selector `json:"tablegroupIdSelector,omitempty" tf:"-"`
}

type TableInfosInitParameters struct {
}

type TableInfosObservation struct {

	// Error messages for creating IDL file.
	Error *string `json:"error,omitempty" tf:"error,omitempty"`

	// Index key set of the TcaplusDB table.
	IndexKeySet *string `json:"indexKeySet,omitempty" tf:"index_key_set,omitempty"`

	// Primary key fields of the TcaplusDB table.
	KeyFields *string `json:"keyFields,omitempty" tf:"key_fields,omitempty"`

	// Total size of primary key field of the TcaplusDB table.
	SumKeyFieldSize *float64 `json:"sumKeyFieldSize,omitempty" tf:"sum_key_field_size,omitempty"`

	// Total size of non-primary key fields of the TcaplusDB table.
	SumValueFieldSize *float64 `json:"sumValueFieldSize,omitempty" tf:"sum_value_field_size,omitempty"`

	// Name of the TcaplusDB table.
	TableName *string `json:"tableName,omitempty" tf:"table_name,omitempty"`

	// Non-primary key fields of the TcaplusDB table.
	ValueFields *string `json:"valueFields,omitempty" tf:"value_fields,omitempty"`
}

type TableInfosParameters struct {
}

// IdlSpec defines the desired state of Idl
type IdlSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     IdlParameters `json:"forProvider"`
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
	InitProvider IdlInitParameters `json:"initProvider,omitempty"`
}

// IdlStatus defines the observed state of Idl.
type IdlStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        IdlObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// Idl is the Schema for the Idls API. Use this resource to create TcaplusDB IDL file.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,tencentcloud}
type Idl struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.fileContent) || (has(self.initProvider) && has(self.initProvider.fileContent))",message="spec.forProvider.fileContent is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.fileExtType) || (has(self.initProvider) && has(self.initProvider.fileExtType))",message="spec.forProvider.fileExtType is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.fileName) || (has(self.initProvider) && has(self.initProvider.fileName))",message="spec.forProvider.fileName is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.fileType) || (has(self.initProvider) && has(self.initProvider.fileType))",message="spec.forProvider.fileType is a required parameter"
	Spec   IdlSpec   `json:"spec"`
	Status IdlStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// IdlList contains a list of Idls
type IdlList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Idl `json:"items"`
}

// Repository type metadata.
var (
	Idl_Kind             = "Idl"
	Idl_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Idl_Kind}.String()
	Idl_KindAPIVersion   = Idl_Kind + "." + CRDGroupVersion.String()
	Idl_GroupVersionKind = CRDGroupVersion.WithKind(Idl_Kind)
)

func init() {
	SchemeBuilder.Register(&Idl{}, &IdlList{})
}
