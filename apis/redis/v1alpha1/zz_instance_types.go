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

type InstanceObservation struct {
	CreateTime *string `json:"createTime,omitempty" tf:"create_time,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	IP *string `json:"ip,omitempty" tf:"ip,omitempty"`

	NodeInfo []NodeInfoObservation `json:"nodeInfo,omitempty" tf:"node_info,omitempty"`

	Status *string `json:"status,omitempty" tf:"status,omitempty"`
}

type InstanceParameters struct {

	// Auto-renew flag. 0 - default state (manual renewal); 1 - automatic renewal; 2 - explicit no automatic renewal.
	// +kubebuilder:validation:Optional
	AutoRenewFlag *float64 `json:"autoRenewFlag,omitempty" tf:"auto_renew_flag,omitempty"`

	// The available zone ID of an instance to be created, please refer to `tencentcloud_redis_zone_config.list`.
	// +kubebuilder:validation:Required
	AvailabilityZone *string `json:"availabilityZone" tf:"availability_zone,omitempty"`

	// The charge type of instance. Valid values: `PREPAID` and `POSTPAID`. Default value is `POSTPAID`. Note: TencentCloud International only supports `POSTPAID`. Caution that update operation on this field will delete old instances and create new with new charge type.
	// +kubebuilder:validation:Optional
	ChargeType *string `json:"chargeType,omitempty" tf:"charge_type,omitempty"`

	// Indicate whether to delete Redis instance directly or not. Default is false. If set true, the instance will be deleted instead of staying recycle bin. Note: only works for `PREPAID` instance.
	// +kubebuilder:validation:Optional
	ForceDelete *bool `json:"forceDelete,omitempty" tf:"force_delete,omitempty"`

	// The memory volume of an available instance(in MB), please refer to `tencentcloud_redis_zone_config.list[zone].shard_memories`. When redis is standard type, it represents total memory size of the instance; when Redis is cluster type, it represents memory size of per sharding.
	// +kubebuilder:validation:Required
	MemSize *float64 `json:"memSize" tf:"mem_size,omitempty"`

	// Instance name.
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Indicates whether the redis instance support no-auth access. NOTE: Only available in private cloud environment.
	// +kubebuilder:validation:Optional
	NoAuth *bool `json:"noAuth,omitempty" tf:"no_auth,omitempty"`

	// Specify params template id. If not set, will use default template.
	// +kubebuilder:validation:Optional
	ParamsTemplateID *string `json:"paramsTemplateId,omitempty" tf:"params_template_id,omitempty"`

	// Password for a Redis user, which should be 8 to 16 characters. NOTE: Only `no_auth=true` specified can make password empty.
	// +kubebuilder:validation:Optional
	PasswordSecretRef *v1.SecretKeySelector `json:"passwordSecretRef,omitempty" tf:"-"`

	// The port used to access a redis instance. The default value is 6379. And this value can't be changed after creation, or the Redis instance will be recreated.
	// +kubebuilder:validation:Optional
	Port *float64 `json:"port,omitempty" tf:"port,omitempty"`

	// The tenancy (time unit is month) of the prepaid instance, NOTE: it only works when charge_type is set to `PREPAID`. Valid values are `1`, `2`, `3`, `4`, `5`, `6`, `7`, `8`, `9`, `10`, `11`, `12`, `24`, `36`.
	// +kubebuilder:validation:Optional
	PrepaidPeriod *float64 `json:"prepaidPeriod,omitempty" tf:"prepaid_period,omitempty"`

	// Specifies which project the instance should belong to.
	// +kubebuilder:validation:Optional
	ProjectID *float64 `json:"projectId,omitempty" tf:"project_id,omitempty"`

	// The number of instance copies. This is not required for standalone and master slave versions and must equal to count of `replica_zone_ids`.
	// +kubebuilder:validation:Optional
	RedisReplicasNum *float64 `json:"redisReplicasNum,omitempty" tf:"redis_replicas_num,omitempty"`

	// The number of instance shard, default is 1. This is not required for standalone and master slave versions.
	// +kubebuilder:validation:Optional
	RedisShardNum *float64 `json:"redisShardNum,omitempty" tf:"redis_shard_num,omitempty"`

	// ID of replica nodes available zone. This is not required for standalone and master slave versions. NOTE: Removing some of the same zone of replicas (e.g. removing 100001 of [100001, 100001, 100002]) will pick the first hit to remove.
	// +kubebuilder:validation:Optional
	ReplicaZoneIds []*float64 `json:"replicaZoneIds,omitempty" tf:"replica_zone_ids,omitempty"`

	// Whether copy read-only is supported, Redis 2.8 Standard Edition and CKV Standard Edition do not support replica read-only, turn on replica read-only, the instance will automatically read and write separate, write requests are routed to the primary node, read requests are routed to the replica node, if you need to open replica read-only, the recommended number of replicas >=2.
	// +kubebuilder:validation:Optional
	ReplicasReadOnly *bool `json:"replicasReadOnly,omitempty" tf:"replicas_read_only,omitempty"`

	// ID of security group. If both vpc_id and subnet_id are not set, this argument should not be set either.
	// +kubebuilder:validation:Optional
	SecurityGroups []*string `json:"securityGroups,omitempty" tf:"security_groups,omitempty"`

	// Specifies which subnet the instance should belong to.
	// +kubebuilder:validation:Optional
	SubnetID *string `json:"subnetId,omitempty" tf:"subnet_id,omitempty"`

	// Instance tags.
	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// Instance type. Available values: `cluster_ckv`,`cluster_redis5.0`,`cluster_redis`,`master_slave_ckv`,`master_slave_redis4.0`,`master_slave_redis5.0`,`master_slave_redis`,`standalone_redis`, specific region support specific types, need to refer data `tencentcloud_redis_zone_config`.
	// +kubebuilder:validation:Optional
	Type *string `json:"type,omitempty" tf:"type,omitempty"`

	// Instance type. Available values reference data source `tencentcloud_redis_zone_config` or [document](https://intl.cloud.tencent.com/document/product/239/32069).
	// +kubebuilder:validation:Optional
	TypeID *float64 `json:"typeId,omitempty" tf:"type_id,omitempty"`

	// ID of the vpc with which the instance is to be associated.
	// +kubebuilder:validation:Optional
	VPCID *string `json:"vpcId,omitempty" tf:"vpc_id,omitempty"`
}

type NodeInfoObservation struct {
	ID *float64 `json:"id,omitempty" tf:"id,omitempty"`

	Master *bool `json:"master,omitempty" tf:"master,omitempty"`

	ZoneID *float64 `json:"zoneId,omitempty" tf:"zone_id,omitempty"`
}

type NodeInfoParameters struct {
}

// InstanceSpec defines the desired state of Instance
type InstanceSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     InstanceParameters `json:"forProvider"`
}

// InstanceStatus defines the observed state of Instance.
type InstanceStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        InstanceObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Instance is the Schema for the Instances API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,tencentcloudjet}
type Instance struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              InstanceSpec   `json:"spec"`
	Status            InstanceStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// InstanceList contains a list of Instances
type InstanceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Instance `json:"items"`
}

// Repository type metadata.
var (
	Instance_Kind             = "Instance"
	Instance_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Instance_Kind}.String()
	Instance_KindAPIVersion   = Instance_Kind + "." + CRDGroupVersion.String()
	Instance_GroupVersionKind = CRDGroupVersion.WithKind(Instance_Kind)
)

func init() {
	SchemeBuilder.Register(&Instance{}, &InstanceList{})
}
