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

type FileSystemInitParameters struct {

	// ID of a access group.
	// ID of a access group.
	// +crossplane:generate:reference:type=AccessGroup
	AccessGroupID *string `json:"accessGroupId,omitempty" tf:"access_group_id,omitempty"`

	// Reference to a AccessGroup to populate accessGroupId.
	// +kubebuilder:validation:Optional
	AccessGroupIDRef *v1.Reference `json:"accessGroupIdRef,omitempty" tf:"-"`

	// Selector for a AccessGroup to populate accessGroupId.
	// +kubebuilder:validation:Optional
	AccessGroupIDSelector *v1.Selector `json:"accessGroupIdSelector,omitempty" tf:"-"`

	// The available zone that the file system locates at.
	// The available zone that the file system locates at.
	AvailabilityZone *string `json:"availabilityZone,omitempty" tf:"availability_zone,omitempty"`

	// File system capacity, in GiB . For Standard Turbo, the minimum purchase required is 40,960 GiB (40 TiB) and the expansion increment is 20,480 GiB (20 TiB). For High-Performance Turbo, the minimum purchase required is 20,480 GiB (20 TiB) and the expansion increment is 10,240 GiB (10 TiB).
	// File system capacity, in GiB (required for the Turbo series). For Standard Turbo, the minimum purchase required is 40,960 GiB (40 TiB) and the expansion increment is 20,480 GiB (20 TiB). For High-Performance Turbo, the minimum purchase required is 20,480 GiB (20 TiB) and the expansion increment is 10,240 GiB (10 TiB).
	Capacity *float64 `json:"capacity,omitempty" tf:"capacity,omitempty"`

	// CCN instance ID .
	// CCN instance ID (required if the network type is CCN).
	CcnID *string `json:"ccnId,omitempty" tf:"ccn_id,omitempty"`

	// CCN IP range used by the CFS , which cannot conflict with other IP ranges bound in CCN.
	// CCN IP range used by the CFS (required if the network type is CCN), which cannot conflict with other IP ranges bound in CCN.
	CidrBlock *string `json:"cidrBlock,omitempty" tf:"cidr_block,omitempty"`

	// IP of mount point.
	// IP of mount point.
	MountIP *string `json:"mountIp,omitempty" tf:"mount_ip,omitempty"`

	// Name of a file system.
	// Name of a file system.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Network type, Default VPC. Valid values: VPC and CCN. Select VPC for a Standard or High-Performance file system, and CCN for a Standard Turbo or High-Performance Turbo one.
	// Network type, Default `VPC`. Valid values: `VPC` and `CCN`. Select `VPC` for a Standard or High-Performance file system, and `CCN` for a Standard Turbo or High-Performance Turbo one.
	NetInterface *string `json:"netInterface,omitempty" tf:"net_interface,omitempty"`

	// File system protocol. Valid values: NFS, CIFS, TURBO. If this parameter is left empty, NFS is used by default. For the Turbo series, you must set this parameter to TURBO.
	// File system protocol. Valid values: `NFS`, `CIFS`, `TURBO`. If this parameter is left empty, `NFS` is used by default. For the Turbo series, you must set this parameter to `TURBO`.
	Protocol *string `json:"protocol,omitempty" tf:"protocol,omitempty"`

	// Storage type of the file system. Valid values: SD (Standard), HP (High-Performance), TB (Standard Turbo), and TP (High-Performance Turbo). Default value: SD.
	// Storage type of the file system. Valid values: `SD` (Standard), `HP` (High-Performance), `TB` (Standard Turbo), and `TP` (High-Performance Turbo). Default value: `SD`.
	StorageType *string `json:"storageType,omitempty" tf:"storage_type,omitempty"`

	// ID of a subnet.
	// ID of a subnet.
	// +crossplane:generate:reference:type=github.com/crossplane-contrib/provider-tencentcloud/apis/vpc/v1alpha1.Subnet
	SubnetID *string `json:"subnetId,omitempty" tf:"subnet_id,omitempty"`

	// Reference to a Subnet in vpc to populate subnetId.
	// +kubebuilder:validation:Optional
	SubnetIDRef *v1.Reference `json:"subnetIdRef,omitempty" tf:"-"`

	// Selector for a Subnet in vpc to populate subnetId.
	// +kubebuilder:validation:Optional
	SubnetIDSelector *v1.Selector `json:"subnetIdSelector,omitempty" tf:"-"`

	// Instance tags.
	// Instance tags.
	// +mapType=granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// ID of a VPC network.
	// ID of a VPC network.
	// +crossplane:generate:reference:type=github.com/crossplane-contrib/provider-tencentcloud/apis/vpc/v1alpha1.VPC
	VPCID *string `json:"vpcId,omitempty" tf:"vpc_id,omitempty"`

	// Reference to a VPC in vpc to populate vpcId.
	// +kubebuilder:validation:Optional
	VPCIDRef *v1.Reference `json:"vpcIdRef,omitempty" tf:"-"`

	// Selector for a VPC in vpc to populate vpcId.
	// +kubebuilder:validation:Optional
	VPCIDSelector *v1.Selector `json:"vpcIdSelector,omitempty" tf:"-"`
}

type FileSystemObservation struct {

	// ID of a access group.
	// ID of a access group.
	AccessGroupID *string `json:"accessGroupId,omitempty" tf:"access_group_id,omitempty"`

	// The available zone that the file system locates at.
	// The available zone that the file system locates at.
	AvailabilityZone *string `json:"availabilityZone,omitempty" tf:"availability_zone,omitempty"`

	// File system capacity, in GiB . For Standard Turbo, the minimum purchase required is 40,960 GiB (40 TiB) and the expansion increment is 20,480 GiB (20 TiB). For High-Performance Turbo, the minimum purchase required is 20,480 GiB (20 TiB) and the expansion increment is 10,240 GiB (10 TiB).
	// File system capacity, in GiB (required for the Turbo series). For Standard Turbo, the minimum purchase required is 40,960 GiB (40 TiB) and the expansion increment is 20,480 GiB (20 TiB). For High-Performance Turbo, the minimum purchase required is 20,480 GiB (20 TiB) and the expansion increment is 10,240 GiB (10 TiB).
	Capacity *float64 `json:"capacity,omitempty" tf:"capacity,omitempty"`

	// CCN instance ID .
	// CCN instance ID (required if the network type is CCN).
	CcnID *string `json:"ccnId,omitempty" tf:"ccn_id,omitempty"`

	// CCN IP range used by the CFS , which cannot conflict with other IP ranges bound in CCN.
	// CCN IP range used by the CFS (required if the network type is CCN), which cannot conflict with other IP ranges bound in CCN.
	CidrBlock *string `json:"cidrBlock,omitempty" tf:"cidr_block,omitempty"`

	// Create time of the file system.
	// Create time of the file system.
	CreateTime *string `json:"createTime,omitempty" tf:"create_time,omitempty"`

	// Mount root-directory.
	// Mount root-directory.
	FsID *string `json:"fsId,omitempty" tf:"fs_id,omitempty"`

	// ID of the resource.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// IP of mount point.
	// IP of mount point.
	MountIP *string `json:"mountIp,omitempty" tf:"mount_ip,omitempty"`

	// Name of a file system.
	// Name of a file system.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Network type, Default VPC. Valid values: VPC and CCN. Select VPC for a Standard or High-Performance file system, and CCN for a Standard Turbo or High-Performance Turbo one.
	// Network type, Default `VPC`. Valid values: `VPC` and `CCN`. Select `VPC` for a Standard or High-Performance file system, and `CCN` for a Standard Turbo or High-Performance Turbo one.
	NetInterface *string `json:"netInterface,omitempty" tf:"net_interface,omitempty"`

	// File system protocol. Valid values: NFS, CIFS, TURBO. If this parameter is left empty, NFS is used by default. For the Turbo series, you must set this parameter to TURBO.
	// File system protocol. Valid values: `NFS`, `CIFS`, `TURBO`. If this parameter is left empty, `NFS` is used by default. For the Turbo series, you must set this parameter to `TURBO`.
	Protocol *string `json:"protocol,omitempty" tf:"protocol,omitempty"`

	// Storage type of the file system. Valid values: SD (Standard), HP (High-Performance), TB (Standard Turbo), and TP (High-Performance Turbo). Default value: SD.
	// Storage type of the file system. Valid values: `SD` (Standard), `HP` (High-Performance), `TB` (Standard Turbo), and `TP` (High-Performance Turbo). Default value: `SD`.
	StorageType *string `json:"storageType,omitempty" tf:"storage_type,omitempty"`

	// ID of a subnet.
	// ID of a subnet.
	SubnetID *string `json:"subnetId,omitempty" tf:"subnet_id,omitempty"`

	// Instance tags.
	// Instance tags.
	// +mapType=granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// ID of a VPC network.
	// ID of a VPC network.
	VPCID *string `json:"vpcId,omitempty" tf:"vpc_id,omitempty"`
}

type FileSystemParameters struct {

	// ID of a access group.
	// ID of a access group.
	// +crossplane:generate:reference:type=AccessGroup
	// +kubebuilder:validation:Optional
	AccessGroupID *string `json:"accessGroupId,omitempty" tf:"access_group_id,omitempty"`

	// Reference to a AccessGroup to populate accessGroupId.
	// +kubebuilder:validation:Optional
	AccessGroupIDRef *v1.Reference `json:"accessGroupIdRef,omitempty" tf:"-"`

	// Selector for a AccessGroup to populate accessGroupId.
	// +kubebuilder:validation:Optional
	AccessGroupIDSelector *v1.Selector `json:"accessGroupIdSelector,omitempty" tf:"-"`

	// The available zone that the file system locates at.
	// The available zone that the file system locates at.
	// +kubebuilder:validation:Optional
	AvailabilityZone *string `json:"availabilityZone,omitempty" tf:"availability_zone,omitempty"`

	// File system capacity, in GiB . For Standard Turbo, the minimum purchase required is 40,960 GiB (40 TiB) and the expansion increment is 20,480 GiB (20 TiB). For High-Performance Turbo, the minimum purchase required is 20,480 GiB (20 TiB) and the expansion increment is 10,240 GiB (10 TiB).
	// File system capacity, in GiB (required for the Turbo series). For Standard Turbo, the minimum purchase required is 40,960 GiB (40 TiB) and the expansion increment is 20,480 GiB (20 TiB). For High-Performance Turbo, the minimum purchase required is 20,480 GiB (20 TiB) and the expansion increment is 10,240 GiB (10 TiB).
	// +kubebuilder:validation:Optional
	Capacity *float64 `json:"capacity,omitempty" tf:"capacity,omitempty"`

	// CCN instance ID .
	// CCN instance ID (required if the network type is CCN).
	// +kubebuilder:validation:Optional
	CcnID *string `json:"ccnId,omitempty" tf:"ccn_id,omitempty"`

	// CCN IP range used by the CFS , which cannot conflict with other IP ranges bound in CCN.
	// CCN IP range used by the CFS (required if the network type is CCN), which cannot conflict with other IP ranges bound in CCN.
	// +kubebuilder:validation:Optional
	CidrBlock *string `json:"cidrBlock,omitempty" tf:"cidr_block,omitempty"`

	// IP of mount point.
	// IP of mount point.
	// +kubebuilder:validation:Optional
	MountIP *string `json:"mountIp,omitempty" tf:"mount_ip,omitempty"`

	// Name of a file system.
	// Name of a file system.
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Network type, Default VPC. Valid values: VPC and CCN. Select VPC for a Standard or High-Performance file system, and CCN for a Standard Turbo or High-Performance Turbo one.
	// Network type, Default `VPC`. Valid values: `VPC` and `CCN`. Select `VPC` for a Standard or High-Performance file system, and `CCN` for a Standard Turbo or High-Performance Turbo one.
	// +kubebuilder:validation:Optional
	NetInterface *string `json:"netInterface,omitempty" tf:"net_interface,omitempty"`

	// File system protocol. Valid values: NFS, CIFS, TURBO. If this parameter is left empty, NFS is used by default. For the Turbo series, you must set this parameter to TURBO.
	// File system protocol. Valid values: `NFS`, `CIFS`, `TURBO`. If this parameter is left empty, `NFS` is used by default. For the Turbo series, you must set this parameter to `TURBO`.
	// +kubebuilder:validation:Optional
	Protocol *string `json:"protocol,omitempty" tf:"protocol,omitempty"`

	// Storage type of the file system. Valid values: SD (Standard), HP (High-Performance), TB (Standard Turbo), and TP (High-Performance Turbo). Default value: SD.
	// Storage type of the file system. Valid values: `SD` (Standard), `HP` (High-Performance), `TB` (Standard Turbo), and `TP` (High-Performance Turbo). Default value: `SD`.
	// +kubebuilder:validation:Optional
	StorageType *string `json:"storageType,omitempty" tf:"storage_type,omitempty"`

	// ID of a subnet.
	// ID of a subnet.
	// +crossplane:generate:reference:type=github.com/crossplane-contrib/provider-tencentcloud/apis/vpc/v1alpha1.Subnet
	// +kubebuilder:validation:Optional
	SubnetID *string `json:"subnetId,omitempty" tf:"subnet_id,omitempty"`

	// Reference to a Subnet in vpc to populate subnetId.
	// +kubebuilder:validation:Optional
	SubnetIDRef *v1.Reference `json:"subnetIdRef,omitempty" tf:"-"`

	// Selector for a Subnet in vpc to populate subnetId.
	// +kubebuilder:validation:Optional
	SubnetIDSelector *v1.Selector `json:"subnetIdSelector,omitempty" tf:"-"`

	// Instance tags.
	// Instance tags.
	// +kubebuilder:validation:Optional
	// +mapType=granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// ID of a VPC network.
	// ID of a VPC network.
	// +crossplane:generate:reference:type=github.com/crossplane-contrib/provider-tencentcloud/apis/vpc/v1alpha1.VPC
	// +kubebuilder:validation:Optional
	VPCID *string `json:"vpcId,omitempty" tf:"vpc_id,omitempty"`

	// Reference to a VPC in vpc to populate vpcId.
	// +kubebuilder:validation:Optional
	VPCIDRef *v1.Reference `json:"vpcIdRef,omitempty" tf:"-"`

	// Selector for a VPC in vpc to populate vpcId.
	// +kubebuilder:validation:Optional
	VPCIDSelector *v1.Selector `json:"vpcIdSelector,omitempty" tf:"-"`
}

// FileSystemSpec defines the desired state of FileSystem
type FileSystemSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     FileSystemParameters `json:"forProvider"`
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
	InitProvider FileSystemInitParameters `json:"initProvider,omitempty"`
}

// FileSystemStatus defines the observed state of FileSystem.
type FileSystemStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        FileSystemObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// FileSystem is the Schema for the FileSystems API. Provides a resource to create a cloud file system(CFS).
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,tencentcloud}
type FileSystem struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.availabilityZone) || (has(self.initProvider) && has(self.initProvider.availabilityZone))",message="spec.forProvider.availabilityZone is a required parameter"
	Spec   FileSystemSpec   `json:"spec"`
	Status FileSystemStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// FileSystemList contains a list of FileSystems
type FileSystemList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []FileSystem `json:"items"`
}

// Repository type metadata.
var (
	FileSystem_Kind             = "FileSystem"
	FileSystem_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: FileSystem_Kind}.String()
	FileSystem_KindAPIVersion   = FileSystem_Kind + "." + CRDGroupVersion.String()
	FileSystem_GroupVersionKind = CRDGroupVersion.WithKind(FileSystem_Kind)
)

func init() {
	SchemeBuilder.Register(&FileSystem{}, &FileSystemList{})
}
