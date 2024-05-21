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

type HourdbInstanceInitParameters struct {

	// db engine version, default to Percona 5.7.17.
	// db engine version, default to Percona 5.7.17.
	DBVersionID *string `json:"dbVersionId,omitempty" tf:"db_version_id,omitempty"`

	// DCN source instance ID.
	// DCN source instance ID.
	DcnInstanceID *string `json:"dcnInstanceId,omitempty" tf:"dcn_instance_id,omitempty"`

	// DCN source region.
	// DCN source region.
	DcnRegion *string `json:"dcnRegion,omitempty" tf:"dcn_region,omitempty"`

	// Whether to open the extranet access.
	// Whether to open the extranet access.
	ExtranetAccess *bool `json:"extranetAccess,omitempty" tf:"extranet_access,omitempty"`

	// Whether to support IPv6.
	// Whether to support IPv6.
	IPv6Flag *float64 `json:"ipv6Flag,omitempty" tf:"ipv6_flag,omitempty"`

	// name of this instance.
	// name of this instance.
	InstanceName *string `json:"instanceName,omitempty" tf:"instance_name,omitempty"`

	// project id.
	// project id.
	ProjectID *float64 `json:"projectId,omitempty" tf:"project_id,omitempty"`

	// resource tags.
	// resource tags.
	ResourceTags []ResourceTagsInitParameters `json:"resourceTags,omitempty" tf:"resource_tags,omitempty"`

	// security group id.
	// security group id.
	SecurityGroupID *string `json:"securityGroupId,omitempty" tf:"security_group_id,omitempty"`

	// instance shard count.
	// instance shard count.
	ShardCount *float64 `json:"shardCount,omitempty" tf:"shard_count,omitempty"`

	// memory(GB) for each shard. It can be obtained by querying api DescribeShardSpec.
	// memory(GB) for each shard. It can be obtained by querying api DescribeShardSpec.
	ShardMemory *float64 `json:"shardMemory,omitempty" tf:"shard_memory,omitempty"`

	// node count for each shard. It can be obtained by querying api DescribeShardSpec.
	// node count for each shard. It can be obtained by querying api DescribeShardSpec.
	ShardNodeCount *float64 `json:"shardNodeCount,omitempty" tf:"shard_node_count,omitempty"`

	// storage(GB) for each shard. It can be obtained by querying api DescribeShardSpec.
	// storage(GB) for each shard. It can be obtained by querying api DescribeShardSpec.
	ShardStorage *float64 `json:"shardStorage,omitempty" tf:"shard_storage,omitempty"`

	// subnet id, its required when vpcId is set.
	// subnet id, its required when vpcId is set.
	// +crossplane:generate:reference:type=github.com/crossplane-contrib/provider-tencentcloud/apis/vpc/v1alpha1.Subnet
	SubnetID *string `json:"subnetId,omitempty" tf:"subnet_id,omitempty"`

	// Reference to a Subnet in vpc to populate subnetId.
	// +kubebuilder:validation:Optional
	SubnetIDRef *v1.Reference `json:"subnetIdRef,omitempty" tf:"-"`

	// Selector for a Subnet in vpc to populate subnetId.
	// +kubebuilder:validation:Optional
	SubnetIDSelector *v1.Selector `json:"subnetIdSelector,omitempty" tf:"-"`

	// vpc id.
	// vpc id.
	// +crossplane:generate:reference:type=github.com/crossplane-contrib/provider-tencentcloud/apis/vpc/v1alpha1.VPC
	VPCID *string `json:"vpcId,omitempty" tf:"vpc_id,omitempty"`

	// Reference to a VPC in vpc to populate vpcId.
	// +kubebuilder:validation:Optional
	VPCIDRef *v1.Reference `json:"vpcIdRef,omitempty" tf:"-"`

	// Selector for a VPC in vpc to populate vpcId.
	// +kubebuilder:validation:Optional
	VPCIDSelector *v1.Selector `json:"vpcIdSelector,omitempty" tf:"-"`

	// The field is required to specify VIP.
	// The field is required to specify VIP.
	Vip *string `json:"vip,omitempty" tf:"vip,omitempty"`

	// The field is required to specify VIPv6.
	// The field is required to specify VIPv6.
	Vipv6 *string `json:"vipv6,omitempty" tf:"vipv6,omitempty"`

	// available zone.
	// available zone.
	// +listType=set
	Zones []*string `json:"zones,omitempty" tf:"zones,omitempty"`
}

type HourdbInstanceObservation struct {

	// db engine version, default to Percona 5.7.17.
	// db engine version, default to Percona 5.7.17.
	DBVersionID *string `json:"dbVersionId,omitempty" tf:"db_version_id,omitempty"`

	// DCN source instance ID.
	// DCN source instance ID.
	DcnInstanceID *string `json:"dcnInstanceId,omitempty" tf:"dcn_instance_id,omitempty"`

	// DCN source region.
	// DCN source region.
	DcnRegion *string `json:"dcnRegion,omitempty" tf:"dcn_region,omitempty"`

	// Whether to open the extranet access.
	// Whether to open the extranet access.
	ExtranetAccess *bool `json:"extranetAccess,omitempty" tf:"extranet_access,omitempty"`

	// ID of the resource.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Whether to support IPv6.
	// Whether to support IPv6.
	IPv6Flag *float64 `json:"ipv6Flag,omitempty" tf:"ipv6_flag,omitempty"`

	// name of this instance.
	// name of this instance.
	InstanceName *string `json:"instanceName,omitempty" tf:"instance_name,omitempty"`

	// project id.
	// project id.
	ProjectID *float64 `json:"projectId,omitempty" tf:"project_id,omitempty"`

	// resource tags.
	// resource tags.
	ResourceTags []ResourceTagsObservation `json:"resourceTags,omitempty" tf:"resource_tags,omitempty"`

	// security group id.
	// security group id.
	SecurityGroupID *string `json:"securityGroupId,omitempty" tf:"security_group_id,omitempty"`

	// instance shard count.
	// instance shard count.
	ShardCount *float64 `json:"shardCount,omitempty" tf:"shard_count,omitempty"`

	// memory(GB) for each shard. It can be obtained by querying api DescribeShardSpec.
	// memory(GB) for each shard. It can be obtained by querying api DescribeShardSpec.
	ShardMemory *float64 `json:"shardMemory,omitempty" tf:"shard_memory,omitempty"`

	// node count for each shard. It can be obtained by querying api DescribeShardSpec.
	// node count for each shard. It can be obtained by querying api DescribeShardSpec.
	ShardNodeCount *float64 `json:"shardNodeCount,omitempty" tf:"shard_node_count,omitempty"`

	// storage(GB) for each shard. It can be obtained by querying api DescribeShardSpec.
	// storage(GB) for each shard. It can be obtained by querying api DescribeShardSpec.
	ShardStorage *float64 `json:"shardStorage,omitempty" tf:"shard_storage,omitempty"`

	// subnet id, its required when vpcId is set.
	// subnet id, its required when vpcId is set.
	SubnetID *string `json:"subnetId,omitempty" tf:"subnet_id,omitempty"`

	// vpc id.
	// vpc id.
	VPCID *string `json:"vpcId,omitempty" tf:"vpc_id,omitempty"`

	// The field is required to specify VIP.
	// The field is required to specify VIP.
	Vip *string `json:"vip,omitempty" tf:"vip,omitempty"`

	// The field is required to specify VIPv6.
	// The field is required to specify VIPv6.
	Vipv6 *string `json:"vipv6,omitempty" tf:"vipv6,omitempty"`

	// Intranet port.
	// Intranet port.
	Vport *float64 `json:"vport,omitempty" tf:"vport,omitempty"`

	// available zone.
	// available zone.
	// +listType=set
	Zones []*string `json:"zones,omitempty" tf:"zones,omitempty"`
}

type HourdbInstanceParameters struct {

	// db engine version, default to Percona 5.7.17.
	// db engine version, default to Percona 5.7.17.
	// +kubebuilder:validation:Optional
	DBVersionID *string `json:"dbVersionId,omitempty" tf:"db_version_id,omitempty"`

	// DCN source instance ID.
	// DCN source instance ID.
	// +kubebuilder:validation:Optional
	DcnInstanceID *string `json:"dcnInstanceId,omitempty" tf:"dcn_instance_id,omitempty"`

	// DCN source region.
	// DCN source region.
	// +kubebuilder:validation:Optional
	DcnRegion *string `json:"dcnRegion,omitempty" tf:"dcn_region,omitempty"`

	// Whether to open the extranet access.
	// Whether to open the extranet access.
	// +kubebuilder:validation:Optional
	ExtranetAccess *bool `json:"extranetAccess,omitempty" tf:"extranet_access,omitempty"`

	// Whether to support IPv6.
	// Whether to support IPv6.
	// +kubebuilder:validation:Optional
	IPv6Flag *float64 `json:"ipv6Flag,omitempty" tf:"ipv6_flag,omitempty"`

	// name of this instance.
	// name of this instance.
	// +kubebuilder:validation:Optional
	InstanceName *string `json:"instanceName,omitempty" tf:"instance_name,omitempty"`

	// project id.
	// project id.
	// +kubebuilder:validation:Optional
	ProjectID *float64 `json:"projectId,omitempty" tf:"project_id,omitempty"`

	// resource tags.
	// resource tags.
	// +kubebuilder:validation:Optional
	ResourceTags []ResourceTagsParameters `json:"resourceTags,omitempty" tf:"resource_tags,omitempty"`

	// security group id.
	// security group id.
	// +kubebuilder:validation:Optional
	SecurityGroupID *string `json:"securityGroupId,omitempty" tf:"security_group_id,omitempty"`

	// instance shard count.
	// instance shard count.
	// +kubebuilder:validation:Optional
	ShardCount *float64 `json:"shardCount,omitempty" tf:"shard_count,omitempty"`

	// memory(GB) for each shard. It can be obtained by querying api DescribeShardSpec.
	// memory(GB) for each shard. It can be obtained by querying api DescribeShardSpec.
	// +kubebuilder:validation:Optional
	ShardMemory *float64 `json:"shardMemory,omitempty" tf:"shard_memory,omitempty"`

	// node count for each shard. It can be obtained by querying api DescribeShardSpec.
	// node count for each shard. It can be obtained by querying api DescribeShardSpec.
	// +kubebuilder:validation:Optional
	ShardNodeCount *float64 `json:"shardNodeCount,omitempty" tf:"shard_node_count,omitempty"`

	// storage(GB) for each shard. It can be obtained by querying api DescribeShardSpec.
	// storage(GB) for each shard. It can be obtained by querying api DescribeShardSpec.
	// +kubebuilder:validation:Optional
	ShardStorage *float64 `json:"shardStorage,omitempty" tf:"shard_storage,omitempty"`

	// subnet id, its required when vpcId is set.
	// subnet id, its required when vpcId is set.
	// +crossplane:generate:reference:type=github.com/crossplane-contrib/provider-tencentcloud/apis/vpc/v1alpha1.Subnet
	// +kubebuilder:validation:Optional
	SubnetID *string `json:"subnetId,omitempty" tf:"subnet_id,omitempty"`

	// Reference to a Subnet in vpc to populate subnetId.
	// +kubebuilder:validation:Optional
	SubnetIDRef *v1.Reference `json:"subnetIdRef,omitempty" tf:"-"`

	// Selector for a Subnet in vpc to populate subnetId.
	// +kubebuilder:validation:Optional
	SubnetIDSelector *v1.Selector `json:"subnetIdSelector,omitempty" tf:"-"`

	// vpc id.
	// vpc id.
	// +crossplane:generate:reference:type=github.com/crossplane-contrib/provider-tencentcloud/apis/vpc/v1alpha1.VPC
	// +kubebuilder:validation:Optional
	VPCID *string `json:"vpcId,omitempty" tf:"vpc_id,omitempty"`

	// Reference to a VPC in vpc to populate vpcId.
	// +kubebuilder:validation:Optional
	VPCIDRef *v1.Reference `json:"vpcIdRef,omitempty" tf:"-"`

	// Selector for a VPC in vpc to populate vpcId.
	// +kubebuilder:validation:Optional
	VPCIDSelector *v1.Selector `json:"vpcIdSelector,omitempty" tf:"-"`

	// The field is required to specify VIP.
	// The field is required to specify VIP.
	// +kubebuilder:validation:Optional
	Vip *string `json:"vip,omitempty" tf:"vip,omitempty"`

	// The field is required to specify VIPv6.
	// The field is required to specify VIPv6.
	// +kubebuilder:validation:Optional
	Vipv6 *string `json:"vipv6,omitempty" tf:"vipv6,omitempty"`

	// available zone.
	// available zone.
	// +kubebuilder:validation:Optional
	// +listType=set
	Zones []*string `json:"zones,omitempty" tf:"zones,omitempty"`
}

type ResourceTagsInitParameters struct {

	// tag key.
	// tag key.
	TagKey *string `json:"tagKey,omitempty" tf:"tag_key,omitempty"`

	// tag value.
	// tag value.
	TagValue *string `json:"tagValue,omitempty" tf:"tag_value,omitempty"`
}

type ResourceTagsObservation struct {

	// tag key.
	// tag key.
	TagKey *string `json:"tagKey,omitempty" tf:"tag_key,omitempty"`

	// tag value.
	// tag value.
	TagValue *string `json:"tagValue,omitempty" tf:"tag_value,omitempty"`
}

type ResourceTagsParameters struct {

	// tag key.
	// tag key.
	// +kubebuilder:validation:Optional
	TagKey *string `json:"tagKey" tf:"tag_key,omitempty"`

	// tag value.
	// tag value.
	// +kubebuilder:validation:Optional
	TagValue *string `json:"tagValue" tf:"tag_value,omitempty"`
}

// HourdbInstanceSpec defines the desired state of HourdbInstance
type HourdbInstanceSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     HourdbInstanceParameters `json:"forProvider"`
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
	InitProvider HourdbInstanceInitParameters `json:"initProvider,omitempty"`
}

// HourdbInstanceStatus defines the observed state of HourdbInstance.
type HourdbInstanceStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        HourdbInstanceObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// HourdbInstance is the Schema for the HourdbInstances API. Provides a resource to create a dcdb hourdb_instance
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,tencentcloud}
type HourdbInstance struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.shardCount) || (has(self.initProvider) && has(self.initProvider.shardCount))",message="spec.forProvider.shardCount is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.shardMemory) || (has(self.initProvider) && has(self.initProvider.shardMemory))",message="spec.forProvider.shardMemory is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.shardNodeCount) || (has(self.initProvider) && has(self.initProvider.shardNodeCount))",message="spec.forProvider.shardNodeCount is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.shardStorage) || (has(self.initProvider) && has(self.initProvider.shardStorage))",message="spec.forProvider.shardStorage is a required parameter"
	Spec   HourdbInstanceSpec   `json:"spec"`
	Status HourdbInstanceStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// HourdbInstanceList contains a list of HourdbInstances
type HourdbInstanceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []HourdbInstance `json:"items"`
}

// Repository type metadata.
var (
	HourdbInstance_Kind             = "HourdbInstance"
	HourdbInstance_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: HourdbInstance_Kind}.String()
	HourdbInstance_KindAPIVersion   = HourdbInstance_Kind + "." + CRDGroupVersion.String()
	HourdbInstance_GroupVersionKind = CRDGroupVersion.WithKind(HourdbInstance_Kind)
)

func init() {
	SchemeBuilder.Register(&HourdbInstance{}, &HourdbInstanceList{})
}
