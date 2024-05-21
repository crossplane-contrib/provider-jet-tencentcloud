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

type NetInfoListInitParameters struct {
}

type NetInfoListObservation struct {

	// Ip address of the net info.
	IP *string `json:"ip,omitempty" tf:"ip,omitempty"`

	// Port of the net info.
	Port *float64 `json:"port,omitempty" tf:"port,omitempty"`
}

type NetInfoListParameters struct {
}

type ReadonlyGroupInitParameters struct {

	// Primary instance ID.
	// Primary instance ID.
	MasterDBInstanceID *string `json:"masterDbInstanceId,omitempty" tf:"master_db_instance_id,omitempty"`

	// Delay threshold in ms.
	// Delay threshold in ms.
	MaxReplayLag *float64 `json:"maxReplayLag,omitempty" tf:"max_replay_lag,omitempty"`

	// Delayed log size threshold in MB.
	// Delayed log size threshold in MB.
	MaxReplayLatency *float64 `json:"maxReplayLatency,omitempty" tf:"max_replay_latency,omitempty"`

	// The minimum number of read-only replicas that must be retained in an RO group.
	// The minimum number of read-only replicas that must be retained in an RO group.
	MinDelayEliminateReserve *float64 `json:"minDelayEliminateReserve,omitempty" tf:"min_delay_eliminate_reserve,omitempty"`

	// RO group name.
	// RO group name.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Project ID.
	// Project ID.
	ProjectID *float64 `json:"projectId,omitempty" tf:"project_id,omitempty"`

	// Whether to remove a read-only replica from an RO group if the delay between the read-only replica and the primary instance exceeds the threshold. Valid values: 0 (no), 1 (yes).
	// Whether to remove a read-only replica from an RO group if the delay between the read-only replica and the primary instance exceeds the threshold. Valid values: 0 (no), 1 (yes).
	ReplayLagEliminate *float64 `json:"replayLagEliminate,omitempty" tf:"replay_lag_eliminate,omitempty"`

	// Whether to remove a read-only replica from an RO group if the sync log size difference between the read-only replica and the primary instance exceeds the threshold. Valid values: 0 (no), 1 (yes).
	// Whether to remove a read-only replica from an RO group if the sync log size difference between the read-only replica and the primary instance exceeds the threshold. Valid values: 0 (no), 1 (yes).
	ReplayLatencyEliminate *float64 `json:"replayLatencyEliminate,omitempty" tf:"replay_latency_eliminate,omitempty"`

	// ID of security group. If both vpc_id and subnet_id are not set, this argument should not be set either.
	// ID of security group. If both vpc_id and subnet_id are not set, this argument should not be set either.
	// +listType=set
	SecurityGroupsIds []*string `json:"securityGroupsIds,omitempty" tf:"security_groups_ids,omitempty"`

	// VPC subnet ID.
	// VPC subnet ID.
	// +crossplane:generate:reference:type=github.com/crossplane-contrib/provider-tencentcloud/apis/vpc/v1alpha1.Subnet
	SubnetID *string `json:"subnetId,omitempty" tf:"subnet_id,omitempty"`

	// Reference to a Subnet in vpc to populate subnetId.
	// +kubebuilder:validation:Optional
	SubnetIDRef *v1.Reference `json:"subnetIdRef,omitempty" tf:"-"`

	// Selector for a Subnet in vpc to populate subnetId.
	// +kubebuilder:validation:Optional
	SubnetIDSelector *v1.Selector `json:"subnetIdSelector,omitempty" tf:"-"`

	// VPC ID.
	// VPC ID.
	// +crossplane:generate:reference:type=github.com/crossplane-contrib/provider-tencentcloud/apis/vpc/v1alpha1.VPC
	VPCID *string `json:"vpcId,omitempty" tf:"vpc_id,omitempty"`

	// Reference to a VPC in vpc to populate vpcId.
	// +kubebuilder:validation:Optional
	VPCIDRef *v1.Reference `json:"vpcIdRef,omitempty" tf:"-"`

	// Selector for a VPC in vpc to populate vpcId.
	// +kubebuilder:validation:Optional
	VPCIDSelector *v1.Selector `json:"vpcIdSelector,omitempty" tf:"-"`
}

type ReadonlyGroupObservation struct {

	// Create time of the postgresql instance.
	// Create time of the postgresql instance.
	CreateTime *string `json:"createTime,omitempty" tf:"create_time,omitempty"`

	// ID of the resource.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Primary instance ID.
	// Primary instance ID.
	MasterDBInstanceID *string `json:"masterDbInstanceId,omitempty" tf:"master_db_instance_id,omitempty"`

	// Delay threshold in ms.
	// Delay threshold in ms.
	MaxReplayLag *float64 `json:"maxReplayLag,omitempty" tf:"max_replay_lag,omitempty"`

	// Delayed log size threshold in MB.
	// Delayed log size threshold in MB.
	MaxReplayLatency *float64 `json:"maxReplayLatency,omitempty" tf:"max_replay_latency,omitempty"`

	// The minimum number of read-only replicas that must be retained in an RO group.
	// The minimum number of read-only replicas that must be retained in an RO group.
	MinDelayEliminateReserve *float64 `json:"minDelayEliminateReserve,omitempty" tf:"min_delay_eliminate_reserve,omitempty"`

	// RO group name.
	// RO group name.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// List of db instance net info.
	// List of db instance net info.
	NetInfoList []NetInfoListObservation `json:"netInfoList,omitempty" tf:"net_info_list,omitempty"`

	// Project ID.
	// Project ID.
	ProjectID *float64 `json:"projectId,omitempty" tf:"project_id,omitempty"`

	// Whether to remove a read-only replica from an RO group if the delay between the read-only replica and the primary instance exceeds the threshold. Valid values: 0 (no), 1 (yes).
	// Whether to remove a read-only replica from an RO group if the delay between the read-only replica and the primary instance exceeds the threshold. Valid values: 0 (no), 1 (yes).
	ReplayLagEliminate *float64 `json:"replayLagEliminate,omitempty" tf:"replay_lag_eliminate,omitempty"`

	// Whether to remove a read-only replica from an RO group if the sync log size difference between the read-only replica and the primary instance exceeds the threshold. Valid values: 0 (no), 1 (yes).
	// Whether to remove a read-only replica from an RO group if the sync log size difference between the read-only replica and the primary instance exceeds the threshold. Valid values: 0 (no), 1 (yes).
	ReplayLatencyEliminate *float64 `json:"replayLatencyEliminate,omitempty" tf:"replay_latency_eliminate,omitempty"`

	// ID of security group. If both vpc_id and subnet_id are not set, this argument should not be set either.
	// ID of security group. If both vpc_id and subnet_id are not set, this argument should not be set either.
	// +listType=set
	SecurityGroupsIds []*string `json:"securityGroupsIds,omitempty" tf:"security_groups_ids,omitempty"`

	// VPC subnet ID.
	// VPC subnet ID.
	SubnetID *string `json:"subnetId,omitempty" tf:"subnet_id,omitempty"`

	// VPC ID.
	// VPC ID.
	VPCID *string `json:"vpcId,omitempty" tf:"vpc_id,omitempty"`
}

type ReadonlyGroupParameters struct {

	// Primary instance ID.
	// Primary instance ID.
	// +kubebuilder:validation:Optional
	MasterDBInstanceID *string `json:"masterDbInstanceId,omitempty" tf:"master_db_instance_id,omitempty"`

	// Delay threshold in ms.
	// Delay threshold in ms.
	// +kubebuilder:validation:Optional
	MaxReplayLag *float64 `json:"maxReplayLag,omitempty" tf:"max_replay_lag,omitempty"`

	// Delayed log size threshold in MB.
	// Delayed log size threshold in MB.
	// +kubebuilder:validation:Optional
	MaxReplayLatency *float64 `json:"maxReplayLatency,omitempty" tf:"max_replay_latency,omitempty"`

	// The minimum number of read-only replicas that must be retained in an RO group.
	// The minimum number of read-only replicas that must be retained in an RO group.
	// +kubebuilder:validation:Optional
	MinDelayEliminateReserve *float64 `json:"minDelayEliminateReserve,omitempty" tf:"min_delay_eliminate_reserve,omitempty"`

	// RO group name.
	// RO group name.
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Project ID.
	// Project ID.
	// +kubebuilder:validation:Optional
	ProjectID *float64 `json:"projectId,omitempty" tf:"project_id,omitempty"`

	// Whether to remove a read-only replica from an RO group if the delay between the read-only replica and the primary instance exceeds the threshold. Valid values: 0 (no), 1 (yes).
	// Whether to remove a read-only replica from an RO group if the delay between the read-only replica and the primary instance exceeds the threshold. Valid values: 0 (no), 1 (yes).
	// +kubebuilder:validation:Optional
	ReplayLagEliminate *float64 `json:"replayLagEliminate,omitempty" tf:"replay_lag_eliminate,omitempty"`

	// Whether to remove a read-only replica from an RO group if the sync log size difference between the read-only replica and the primary instance exceeds the threshold. Valid values: 0 (no), 1 (yes).
	// Whether to remove a read-only replica from an RO group if the sync log size difference between the read-only replica and the primary instance exceeds the threshold. Valid values: 0 (no), 1 (yes).
	// +kubebuilder:validation:Optional
	ReplayLatencyEliminate *float64 `json:"replayLatencyEliminate,omitempty" tf:"replay_latency_eliminate,omitempty"`

	// ID of security group. If both vpc_id and subnet_id are not set, this argument should not be set either.
	// ID of security group. If both vpc_id and subnet_id are not set, this argument should not be set either.
	// +kubebuilder:validation:Optional
	// +listType=set
	SecurityGroupsIds []*string `json:"securityGroupsIds,omitempty" tf:"security_groups_ids,omitempty"`

	// VPC subnet ID.
	// VPC subnet ID.
	// +crossplane:generate:reference:type=github.com/crossplane-contrib/provider-tencentcloud/apis/vpc/v1alpha1.Subnet
	// +kubebuilder:validation:Optional
	SubnetID *string `json:"subnetId,omitempty" tf:"subnet_id,omitempty"`

	// Reference to a Subnet in vpc to populate subnetId.
	// +kubebuilder:validation:Optional
	SubnetIDRef *v1.Reference `json:"subnetIdRef,omitempty" tf:"-"`

	// Selector for a Subnet in vpc to populate subnetId.
	// +kubebuilder:validation:Optional
	SubnetIDSelector *v1.Selector `json:"subnetIdSelector,omitempty" tf:"-"`

	// VPC ID.
	// VPC ID.
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

// ReadonlyGroupSpec defines the desired state of ReadonlyGroup
type ReadonlyGroupSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ReadonlyGroupParameters `json:"forProvider"`
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
	InitProvider ReadonlyGroupInitParameters `json:"initProvider,omitempty"`
}

// ReadonlyGroupStatus defines the observed state of ReadonlyGroup.
type ReadonlyGroupStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ReadonlyGroupObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// ReadonlyGroup is the Schema for the ReadonlyGroups API. Use this resource to create postgresql readonly group.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,tencentcloud}
type ReadonlyGroup struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.masterDbInstanceId) || (has(self.initProvider) && has(self.initProvider.masterDbInstanceId))",message="spec.forProvider.masterDbInstanceId is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.maxReplayLag) || (has(self.initProvider) && has(self.initProvider.maxReplayLag))",message="spec.forProvider.maxReplayLag is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.maxReplayLatency) || (has(self.initProvider) && has(self.initProvider.maxReplayLatency))",message="spec.forProvider.maxReplayLatency is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.minDelayEliminateReserve) || (has(self.initProvider) && has(self.initProvider.minDelayEliminateReserve))",message="spec.forProvider.minDelayEliminateReserve is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.name) || (has(self.initProvider) && has(self.initProvider.name))",message="spec.forProvider.name is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.projectId) || (has(self.initProvider) && has(self.initProvider.projectId))",message="spec.forProvider.projectId is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.replayLagEliminate) || (has(self.initProvider) && has(self.initProvider.replayLagEliminate))",message="spec.forProvider.replayLagEliminate is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.replayLatencyEliminate) || (has(self.initProvider) && has(self.initProvider.replayLatencyEliminate))",message="spec.forProvider.replayLatencyEliminate is a required parameter"
	Spec   ReadonlyGroupSpec   `json:"spec"`
	Status ReadonlyGroupStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ReadonlyGroupList contains a list of ReadonlyGroups
type ReadonlyGroupList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ReadonlyGroup `json:"items"`
}

// Repository type metadata.
var (
	ReadonlyGroup_Kind             = "ReadonlyGroup"
	ReadonlyGroup_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: ReadonlyGroup_Kind}.String()
	ReadonlyGroup_KindAPIVersion   = ReadonlyGroup_Kind + "." + CRDGroupVersion.String()
	ReadonlyGroup_GroupVersionKind = CRDGroupVersion.WithKind(ReadonlyGroup_Kind)
)

func init() {
	SchemeBuilder.Register(&ReadonlyGroup{}, &ReadonlyGroupList{})
}
