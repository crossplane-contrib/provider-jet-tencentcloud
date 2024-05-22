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

type InstanceInitParameters struct {

	// Auto-renew flag. 0 - default state (manual renewal); 1 - automatic renewal; 2 - explicit no automatic renewal.
	// Auto-renew flag. 0 - default state (manual renewal); 1 - automatic renewal; 2 - explicit no automatic renewal.
	AutoRenewFlag *float64 `json:"autoRenewFlag,omitempty" tf:"auto_renew_flag,omitempty"`

	// The available zone ID of an instance to be created, please refer to tencentcloud_redis_zone_config.list.
	// The available zone ID of an instance to be created, please refer to `tencentcloud_redis_zone_config.list`.
	AvailabilityZone *string `json:"availabilityZone,omitempty" tf:"availability_zone,omitempty"`

	// The charge type of instance. Valid values: PREPAID and POSTPAID. Default value is POSTPAID. Note: TencentCloud International only supports POSTPAID. Caution that update operation on this field will delete old instances and create new with new charge type.
	// The charge type of instance. Valid values: `PREPAID` and `POSTPAID`. Default value is `POSTPAID`. Note: TencentCloud International only supports `POSTPAID`. Caution that update operation on this field will delete old instances and create new with new charge type.
	ChargeType *string `json:"chargeType,omitempty" tf:"charge_type,omitempty"`

	// Indicate whether to delete Redis instance directly or not. Default is false. If set true, the instance will be deleted instead of staying recycle bin. Note: only works for PREPAID instance.
	// Indicate whether to delete Redis instance directly or not. Default is false. If set true, the instance will be deleted instead of staying recycle bin. Note: only works for `PREPAID` instance.
	ForceDelete *bool `json:"forceDelete,omitempty" tf:"force_delete,omitempty"`

	// IP address of an instance. When the operation_network is changeVip, this parameter needs to be configured.
	// IP address of an instance. When the `operation_network` is `changeVip`, this parameter needs to be configured.
	IP *string `json:"ip,omitempty" tf:"ip,omitempty"`

	// The memory volume of an available instance(in MB), please refer to tencentcloud_redis_zone_config.list[zone].shard_memories. When redis is standard type, it represents total memory size of the instance; when Redis is cluster type, it represents memory size of per sharding. 512MB is supported only in master-slave instance.
	// The memory volume of an available instance(in MB), please refer to `tencentcloud_redis_zone_config.list[zone].shard_memories`. When redis is standard type, it represents total memory size of the instance; when Redis is cluster type, it represents memory size of per sharding. `512MB` is supported only in master-slave instance.
	MemSize *float64 `json:"memSize,omitempty" tf:"mem_size,omitempty"`

	// Instance name.
	// Instance name.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Indicates whether the redis instance support no-auth access. NOTE: Only available in private cloud environment.
	// Indicates whether the redis instance support no-auth access. NOTE: Only available in private cloud environment.
	NoAuth *bool `json:"noAuth,omitempty" tf:"no_auth,omitempty"`

	// Refers to the category of the pre-modified network, including: changeVip: refers to switching the private network, including its intranet IPv4 address and port; changeVpc: refers to switching the subnet to which the private network belongs; changeBaseToVpc: refers to switching the basic network to a private network; changeVPort: refers to only modifying the instance network port.
	// Refers to the category of the pre-modified network, including: `changeVip`: refers to switching the private network, including its intranet IPv4 address and port; `changeVpc`: refers to switching the subnet to which the private network belongs; `changeBaseToVpc`: refers to switching the basic network to a private network; `changeVPort`: refers to only modifying the instance network port.
	OperationNetwork *string `json:"operationNetwork,omitempty" tf:"operation_network,omitempty"`

	// Specify params template id. If not set, will use default template.
	// Specify params template id. If not set, will use default template.
	ParamsTemplateID *string `json:"paramsTemplateId,omitempty" tf:"params_template_id,omitempty"`

	// The port used to access a redis instance. The default value is 6379. When the operation_network is changeVPort or changeVip, this parameter needs to be configured.
	// The port used to access a redis instance. The default value is 6379. When the `operation_network` is `changeVPort` or `changeVip`, this parameter needs to be configured.
	Port *float64 `json:"port,omitempty" tf:"port,omitempty"`

	// The tenancy (time unit is month) of the prepaid instance, NOTE: it only works when charge_type is set to PREPAID. Valid values are 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 24, 36.
	// The tenancy (time unit is month) of the prepaid instance, NOTE: it only works when charge_type is set to `PREPAID`. Valid values are `1`, `2`, `3`, `4`, `5`, `6`, `7`, `8`, `9`, `10`, `11`, `12`, `24`, `36`.
	PrepaidPeriod *float64 `json:"prepaidPeriod,omitempty" tf:"prepaid_period,omitempty"`

	// Specifies which project the instance should belong to.
	// Specifies which project the instance should belong to.
	ProjectID *float64 `json:"projectId,omitempty" tf:"project_id,omitempty"`

	// Original intranet IPv4 address retention time: unit: day, value range: 0, 1, 2, 3, 7, 15.
	// Original intranet IPv4 address retention time: unit: day, value range: `0`, `1`, `2`, `3`, `7`, `15`.
	Recycle *float64 `json:"recycle,omitempty" tf:"recycle,omitempty"`

	// The number of instance copies. This is not required for standalone and master slave versions and must equal to count of replica_zone_ids, Non-multi-AZ does not require replica_zone_ids; Redis memory version 4.0, 5.0, 6.2 standard architecture and cluster architecture support the number of copies in the range [1, 2, 3, 4, 5]; Redis 2.8 standard version and CKV standard version only support 1 copy.
	// The number of instance copies. This is not required for standalone and master slave versions and must equal to count of `replica_zone_ids`, Non-multi-AZ does not require `replica_zone_ids`; Redis memory version 4.0, 5.0, 6.2 standard architecture and cluster architecture support the number of copies in the range [1, 2, 3, 4, 5]; Redis 2.8 standard version and CKV standard version only support 1 copy.
	RedisReplicasNum *float64 `json:"redisReplicasNum,omitempty" tf:"redis_replicas_num,omitempty"`

	// The number of instance shards; this parameter does not need to be configured for standard version instances; for cluster version instances, the number of shards ranges from: [1, 3, 5, 8, 12, 16, 24 , 32, 40, 48, 64, 80, 96, 128].
	// The number of instance shards; this parameter does not need to be configured for standard version instances; for cluster version instances, the number of shards ranges from: [`1`, `3`, `5`, `8`, `12`, `16`, `24 `, `32`, `40`, `48`, `64`, `80`, `96`, `128`].
	RedisShardNum *float64 `json:"redisShardNum,omitempty" tf:"redis_shard_num,omitempty"`

	// ID of replica nodes available zone. This is not required for standalone and master slave versions. NOTE: Removing some of the same zone of replicas (e.g. removing 100001 of [100001, 100001, 100002]) will pick the first hit to remove.
	// ID of replica nodes available zone. This is not required for standalone and master slave versions. NOTE: Removing some of the same zone of replicas (e.g. removing 100001 of [100001, 100001, 100002]) will pick the first hit to remove.
	ReplicaZoneIds []*float64 `json:"replicaZoneIds,omitempty" tf:"replica_zone_ids,omitempty"`

	// Whether copy read-only is supported, Redis 2.8 Standard Edition and CKV Standard Edition do not support replica read-only, turn on replica read-only, the instance will automatically read and write separate, write requests are routed to the primary node, read requests are routed to the replica node, if you need to open replica read-only, the recommended number of replicas >=2.
	// Whether copy read-only is supported, Redis 2.8 Standard Edition and CKV Standard Edition do not support replica read-only, turn on replica read-only, the instance will automatically read and write separate, write requests are routed to the primary node, read requests are routed to the replica node, if you need to open replica read-only, the recommended number of replicas >=2.
	ReplicasReadOnly *bool `json:"replicasReadOnly,omitempty" tf:"replicas_read_only,omitempty"`

	// ID of security group. If both vpc_id and subnet_id are not set, this argument should not be set either.
	// ID of security group. If both vpc_id and subnet_id are not set, this argument should not be set either.
	// +listType=set
	SecurityGroups []*string `json:"securityGroups,omitempty" tf:"security_groups,omitempty"`

	// Specifies which subnet the instance should belong to. When the operation_network is changeVpc or changeBaseToVpc, this parameter needs to be configured.
	// Specifies which subnet the instance should belong to. When the `operation_network` is `changeVpc` or `changeBaseToVpc`, this parameter needs to be configured.
	SubnetID *string `json:"subnetId,omitempty" tf:"subnet_id,omitempty"`

	// Instance tags.
	// Instance tags.
	// +mapType=granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// It has been deprecated from version 1.33.1. Please use 'type_id' instead. Instance type. Available values: cluster_ckv,cluster_redis5.0,cluster_redis,master_slave_ckv,master_slave_redis4.0,master_slave_redis5.0,master_slave_redis,standalone_redis, specific region support specific types, need to refer data tencentcloud_redis_zone_config.
	// Instance type. Available values: `cluster_ckv`,`cluster_redis5.0`,`cluster_redis`,`master_slave_ckv`,`master_slave_redis4.0`,`master_slave_redis5.0`,`master_slave_redis`,`standalone_redis`, specific region support specific types, need to refer data `tencentcloud_redis_zone_config`.
	Type *string `json:"type,omitempty" tf:"type,omitempty"`

	// Instance type. Available values reference data source tencentcloud_redis_zone_config or document, toggle immediately when modified.
	// Instance type. Available values reference data source `tencentcloud_redis_zone_config` or [document](https://intl.cloud.tencent.com/document/product/239/32069), toggle immediately when modified.
	TypeID *float64 `json:"typeId,omitempty" tf:"type_id,omitempty"`

	// ID of the vpc with which the instance is to be associated. When the operation_network is changeVpc or changeBaseToVpc, this parameter needs to be configured.
	// ID of the vpc with which the instance is to be associated. When the `operation_network` is `changeVpc` or `changeBaseToVpc`, this parameter needs to be configured.
	VPCID *string `json:"vpcId,omitempty" tf:"vpc_id,omitempty"`

	// Switching mode: 1-maintenance time window switching, 2-immediate switching, default value 2.
	// Switching mode: `1`-maintenance time window switching, `2`-immediate switching, default value `2`.
	WaitSwitch *float64 `json:"waitSwitch,omitempty" tf:"wait_switch,omitempty"`
}

type InstanceObservation struct {

	// Auto-renew flag. 0 - default state (manual renewal); 1 - automatic renewal; 2 - explicit no automatic renewal.
	// Auto-renew flag. 0 - default state (manual renewal); 1 - automatic renewal; 2 - explicit no automatic renewal.
	AutoRenewFlag *float64 `json:"autoRenewFlag,omitempty" tf:"auto_renew_flag,omitempty"`

	// The available zone ID of an instance to be created, please refer to tencentcloud_redis_zone_config.list.
	// The available zone ID of an instance to be created, please refer to `tencentcloud_redis_zone_config.list`.
	AvailabilityZone *string `json:"availabilityZone,omitempty" tf:"availability_zone,omitempty"`

	// The charge type of instance. Valid values: PREPAID and POSTPAID. Default value is POSTPAID. Note: TencentCloud International only supports POSTPAID. Caution that update operation on this field will delete old instances and create new with new charge type.
	// The charge type of instance. Valid values: `PREPAID` and `POSTPAID`. Default value is `POSTPAID`. Note: TencentCloud International only supports `POSTPAID`. Caution that update operation on this field will delete old instances and create new with new charge type.
	ChargeType *string `json:"chargeType,omitempty" tf:"charge_type,omitempty"`

	// The time when the instance was created.
	// The time when the instance was created.
	CreateTime *string `json:"createTime,omitempty" tf:"create_time,omitempty"`

	// Indicate whether to delete Redis instance directly or not. Default is false. If set true, the instance will be deleted instead of staying recycle bin. Note: only works for PREPAID instance.
	// Indicate whether to delete Redis instance directly or not. Default is false. If set true, the instance will be deleted instead of staying recycle bin. Note: only works for `PREPAID` instance.
	ForceDelete *bool `json:"forceDelete,omitempty" tf:"force_delete,omitempty"`

	// ID of the resource.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// IP address of an instance. When the operation_network is changeVip, this parameter needs to be configured.
	// IP address of an instance. When the `operation_network` is `changeVip`, this parameter needs to be configured.
	IP *string `json:"ip,omitempty" tf:"ip,omitempty"`

	// The memory volume of an available instance(in MB), please refer to tencentcloud_redis_zone_config.list[zone].shard_memories. When redis is standard type, it represents total memory size of the instance; when Redis is cluster type, it represents memory size of per sharding. 512MB is supported only in master-slave instance.
	// The memory volume of an available instance(in MB), please refer to `tencentcloud_redis_zone_config.list[zone].shard_memories`. When redis is standard type, it represents total memory size of the instance; when Redis is cluster type, it represents memory size of per sharding. `512MB` is supported only in master-slave instance.
	MemSize *float64 `json:"memSize,omitempty" tf:"mem_size,omitempty"`

	// Instance name.
	// Instance name.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Indicates whether the redis instance support no-auth access. NOTE: Only available in private cloud environment.
	// Indicates whether the redis instance support no-auth access. NOTE: Only available in private cloud environment.
	NoAuth *bool `json:"noAuth,omitempty" tf:"no_auth,omitempty"`

	// Readonly Primary/Replica nodes.
	// Readonly Primary/Replica nodes.
	NodeInfo []NodeInfoObservation `json:"nodeInfo,omitempty" tf:"node_info,omitempty"`

	// Refers to the category of the pre-modified network, including: changeVip: refers to switching the private network, including its intranet IPv4 address and port; changeVpc: refers to switching the subnet to which the private network belongs; changeBaseToVpc: refers to switching the basic network to a private network; changeVPort: refers to only modifying the instance network port.
	// Refers to the category of the pre-modified network, including: `changeVip`: refers to switching the private network, including its intranet IPv4 address and port; `changeVpc`: refers to switching the subnet to which the private network belongs; `changeBaseToVpc`: refers to switching the basic network to a private network; `changeVPort`: refers to only modifying the instance network port.
	OperationNetwork *string `json:"operationNetwork,omitempty" tf:"operation_network,omitempty"`

	// Specify params template id. If not set, will use default template.
	// Specify params template id. If not set, will use default template.
	ParamsTemplateID *string `json:"paramsTemplateId,omitempty" tf:"params_template_id,omitempty"`

	// The port used to access a redis instance. The default value is 6379. When the operation_network is changeVPort or changeVip, this parameter needs to be configured.
	// The port used to access a redis instance. The default value is 6379. When the `operation_network` is `changeVPort` or `changeVip`, this parameter needs to be configured.
	Port *float64 `json:"port,omitempty" tf:"port,omitempty"`

	// The tenancy (time unit is month) of the prepaid instance, NOTE: it only works when charge_type is set to PREPAID. Valid values are 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 24, 36.
	// The tenancy (time unit is month) of the prepaid instance, NOTE: it only works when charge_type is set to `PREPAID`. Valid values are `1`, `2`, `3`, `4`, `5`, `6`, `7`, `8`, `9`, `10`, `11`, `12`, `24`, `36`.
	PrepaidPeriod *float64 `json:"prepaidPeriod,omitempty" tf:"prepaid_period,omitempty"`

	// Specifies which project the instance should belong to.
	// Specifies which project the instance should belong to.
	ProjectID *float64 `json:"projectId,omitempty" tf:"project_id,omitempty"`

	// Original intranet IPv4 address retention time: unit: day, value range: 0, 1, 2, 3, 7, 15.
	// Original intranet IPv4 address retention time: unit: day, value range: `0`, `1`, `2`, `3`, `7`, `15`.
	Recycle *float64 `json:"recycle,omitempty" tf:"recycle,omitempty"`

	// The number of instance copies. This is not required for standalone and master slave versions and must equal to count of replica_zone_ids, Non-multi-AZ does not require replica_zone_ids; Redis memory version 4.0, 5.0, 6.2 standard architecture and cluster architecture support the number of copies in the range [1, 2, 3, 4, 5]; Redis 2.8 standard version and CKV standard version only support 1 copy.
	// The number of instance copies. This is not required for standalone and master slave versions and must equal to count of `replica_zone_ids`, Non-multi-AZ does not require `replica_zone_ids`; Redis memory version 4.0, 5.0, 6.2 standard architecture and cluster architecture support the number of copies in the range [1, 2, 3, 4, 5]; Redis 2.8 standard version and CKV standard version only support 1 copy.
	RedisReplicasNum *float64 `json:"redisReplicasNum,omitempty" tf:"redis_replicas_num,omitempty"`

	// The number of instance shards; this parameter does not need to be configured for standard version instances; for cluster version instances, the number of shards ranges from: [1, 3, 5, 8, 12, 16, 24 , 32, 40, 48, 64, 80, 96, 128].
	// The number of instance shards; this parameter does not need to be configured for standard version instances; for cluster version instances, the number of shards ranges from: [`1`, `3`, `5`, `8`, `12`, `16`, `24 `, `32`, `40`, `48`, `64`, `80`, `96`, `128`].
	RedisShardNum *float64 `json:"redisShardNum,omitempty" tf:"redis_shard_num,omitempty"`

	// ID of replica nodes available zone. This is not required for standalone and master slave versions. NOTE: Removing some of the same zone of replicas (e.g. removing 100001 of [100001, 100001, 100002]) will pick the first hit to remove.
	// ID of replica nodes available zone. This is not required for standalone and master slave versions. NOTE: Removing some of the same zone of replicas (e.g. removing 100001 of [100001, 100001, 100002]) will pick the first hit to remove.
	ReplicaZoneIds []*float64 `json:"replicaZoneIds,omitempty" tf:"replica_zone_ids,omitempty"`

	// Whether copy read-only is supported, Redis 2.8 Standard Edition and CKV Standard Edition do not support replica read-only, turn on replica read-only, the instance will automatically read and write separate, write requests are routed to the primary node, read requests are routed to the replica node, if you need to open replica read-only, the recommended number of replicas >=2.
	// Whether copy read-only is supported, Redis 2.8 Standard Edition and CKV Standard Edition do not support replica read-only, turn on replica read-only, the instance will automatically read and write separate, write requests are routed to the primary node, read requests are routed to the replica node, if you need to open replica read-only, the recommended number of replicas >=2.
	ReplicasReadOnly *bool `json:"replicasReadOnly,omitempty" tf:"replicas_read_only,omitempty"`

	// ID of security group. If both vpc_id and subnet_id are not set, this argument should not be set either.
	// ID of security group. If both vpc_id and subnet_id are not set, this argument should not be set either.
	// +listType=set
	SecurityGroups []*string `json:"securityGroups,omitempty" tf:"security_groups,omitempty"`

	// Current status of an instance, maybe: init, processing, online, isolate and todelete.
	// Current status of an instance, maybe: init, processing, online, isolate and todelete.
	Status *string `json:"status,omitempty" tf:"status,omitempty"`

	// Specifies which subnet the instance should belong to. When the operation_network is changeVpc or changeBaseToVpc, this parameter needs to be configured.
	// Specifies which subnet the instance should belong to. When the `operation_network` is `changeVpc` or `changeBaseToVpc`, this parameter needs to be configured.
	SubnetID *string `json:"subnetId,omitempty" tf:"subnet_id,omitempty"`

	// Instance tags.
	// Instance tags.
	// +mapType=granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// It has been deprecated from version 1.33.1. Please use 'type_id' instead. Instance type. Available values: cluster_ckv,cluster_redis5.0,cluster_redis,master_slave_ckv,master_slave_redis4.0,master_slave_redis5.0,master_slave_redis,standalone_redis, specific region support specific types, need to refer data tencentcloud_redis_zone_config.
	// Instance type. Available values: `cluster_ckv`,`cluster_redis5.0`,`cluster_redis`,`master_slave_ckv`,`master_slave_redis4.0`,`master_slave_redis5.0`,`master_slave_redis`,`standalone_redis`, specific region support specific types, need to refer data `tencentcloud_redis_zone_config`.
	Type *string `json:"type,omitempty" tf:"type,omitempty"`

	// Instance type. Available values reference data source tencentcloud_redis_zone_config or document, toggle immediately when modified.
	// Instance type. Available values reference data source `tencentcloud_redis_zone_config` or [document](https://intl.cloud.tencent.com/document/product/239/32069), toggle immediately when modified.
	TypeID *float64 `json:"typeId,omitempty" tf:"type_id,omitempty"`

	// ID of the vpc with which the instance is to be associated. When the operation_network is changeVpc or changeBaseToVpc, this parameter needs to be configured.
	// ID of the vpc with which the instance is to be associated. When the `operation_network` is `changeVpc` or `changeBaseToVpc`, this parameter needs to be configured.
	VPCID *string `json:"vpcId,omitempty" tf:"vpc_id,omitempty"`

	// Switching mode: 1-maintenance time window switching, 2-immediate switching, default value 2.
	// Switching mode: `1`-maintenance time window switching, `2`-immediate switching, default value `2`.
	WaitSwitch *float64 `json:"waitSwitch,omitempty" tf:"wait_switch,omitempty"`
}

type InstanceParameters struct {

	// Auto-renew flag. 0 - default state (manual renewal); 1 - automatic renewal; 2 - explicit no automatic renewal.
	// Auto-renew flag. 0 - default state (manual renewal); 1 - automatic renewal; 2 - explicit no automatic renewal.
	// +kubebuilder:validation:Optional
	AutoRenewFlag *float64 `json:"autoRenewFlag,omitempty" tf:"auto_renew_flag,omitempty"`

	// The available zone ID of an instance to be created, please refer to tencentcloud_redis_zone_config.list.
	// The available zone ID of an instance to be created, please refer to `tencentcloud_redis_zone_config.list`.
	// +kubebuilder:validation:Optional
	AvailabilityZone *string `json:"availabilityZone,omitempty" tf:"availability_zone,omitempty"`

	// The charge type of instance. Valid values: PREPAID and POSTPAID. Default value is POSTPAID. Note: TencentCloud International only supports POSTPAID. Caution that update operation on this field will delete old instances and create new with new charge type.
	// The charge type of instance. Valid values: `PREPAID` and `POSTPAID`. Default value is `POSTPAID`. Note: TencentCloud International only supports `POSTPAID`. Caution that update operation on this field will delete old instances and create new with new charge type.
	// +kubebuilder:validation:Optional
	ChargeType *string `json:"chargeType,omitempty" tf:"charge_type,omitempty"`

	// Indicate whether to delete Redis instance directly or not. Default is false. If set true, the instance will be deleted instead of staying recycle bin. Note: only works for PREPAID instance.
	// Indicate whether to delete Redis instance directly or not. Default is false. If set true, the instance will be deleted instead of staying recycle bin. Note: only works for `PREPAID` instance.
	// +kubebuilder:validation:Optional
	ForceDelete *bool `json:"forceDelete,omitempty" tf:"force_delete,omitempty"`

	// IP address of an instance. When the operation_network is changeVip, this parameter needs to be configured.
	// IP address of an instance. When the `operation_network` is `changeVip`, this parameter needs to be configured.
	// +kubebuilder:validation:Optional
	IP *string `json:"ip,omitempty" tf:"ip,omitempty"`

	// The memory volume of an available instance(in MB), please refer to tencentcloud_redis_zone_config.list[zone].shard_memories. When redis is standard type, it represents total memory size of the instance; when Redis is cluster type, it represents memory size of per sharding. 512MB is supported only in master-slave instance.
	// The memory volume of an available instance(in MB), please refer to `tencentcloud_redis_zone_config.list[zone].shard_memories`. When redis is standard type, it represents total memory size of the instance; when Redis is cluster type, it represents memory size of per sharding. `512MB` is supported only in master-slave instance.
	// +kubebuilder:validation:Optional
	MemSize *float64 `json:"memSize,omitempty" tf:"mem_size,omitempty"`

	// Instance name.
	// Instance name.
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Indicates whether the redis instance support no-auth access. NOTE: Only available in private cloud environment.
	// Indicates whether the redis instance support no-auth access. NOTE: Only available in private cloud environment.
	// +kubebuilder:validation:Optional
	NoAuth *bool `json:"noAuth,omitempty" tf:"no_auth,omitempty"`

	// Refers to the category of the pre-modified network, including: changeVip: refers to switching the private network, including its intranet IPv4 address and port; changeVpc: refers to switching the subnet to which the private network belongs; changeBaseToVpc: refers to switching the basic network to a private network; changeVPort: refers to only modifying the instance network port.
	// Refers to the category of the pre-modified network, including: `changeVip`: refers to switching the private network, including its intranet IPv4 address and port; `changeVpc`: refers to switching the subnet to which the private network belongs; `changeBaseToVpc`: refers to switching the basic network to a private network; `changeVPort`: refers to only modifying the instance network port.
	// +kubebuilder:validation:Optional
	OperationNetwork *string `json:"operationNetwork,omitempty" tf:"operation_network,omitempty"`

	// Specify params template id. If not set, will use default template.
	// Specify params template id. If not set, will use default template.
	// +kubebuilder:validation:Optional
	ParamsTemplateID *string `json:"paramsTemplateId,omitempty" tf:"params_template_id,omitempty"`

	// Password for a Redis user, which should be 8 to 16 characters. NOTE: Only no_auth=true specified can make password empty.
	// Password for a Redis user, which should be 8 to 16 characters. NOTE: Only `no_auth=true` specified can make password empty.
	// +kubebuilder:validation:Optional
	PasswordSecretRef *v1.SecretKeySelector `json:"passwordSecretRef,omitempty" tf:"-"`

	// The port used to access a redis instance. The default value is 6379. When the operation_network is changeVPort or changeVip, this parameter needs to be configured.
	// The port used to access a redis instance. The default value is 6379. When the `operation_network` is `changeVPort` or `changeVip`, this parameter needs to be configured.
	// +kubebuilder:validation:Optional
	Port *float64 `json:"port,omitempty" tf:"port,omitempty"`

	// The tenancy (time unit is month) of the prepaid instance, NOTE: it only works when charge_type is set to PREPAID. Valid values are 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 24, 36.
	// The tenancy (time unit is month) of the prepaid instance, NOTE: it only works when charge_type is set to `PREPAID`. Valid values are `1`, `2`, `3`, `4`, `5`, `6`, `7`, `8`, `9`, `10`, `11`, `12`, `24`, `36`.
	// +kubebuilder:validation:Optional
	PrepaidPeriod *float64 `json:"prepaidPeriod,omitempty" tf:"prepaid_period,omitempty"`

	// Specifies which project the instance should belong to.
	// Specifies which project the instance should belong to.
	// +kubebuilder:validation:Optional
	ProjectID *float64 `json:"projectId,omitempty" tf:"project_id,omitempty"`

	// Original intranet IPv4 address retention time: unit: day, value range: 0, 1, 2, 3, 7, 15.
	// Original intranet IPv4 address retention time: unit: day, value range: `0`, `1`, `2`, `3`, `7`, `15`.
	// +kubebuilder:validation:Optional
	Recycle *float64 `json:"recycle,omitempty" tf:"recycle,omitempty"`

	// The number of instance copies. This is not required for standalone and master slave versions and must equal to count of replica_zone_ids, Non-multi-AZ does not require replica_zone_ids; Redis memory version 4.0, 5.0, 6.2 standard architecture and cluster architecture support the number of copies in the range [1, 2, 3, 4, 5]; Redis 2.8 standard version and CKV standard version only support 1 copy.
	// The number of instance copies. This is not required for standalone and master slave versions and must equal to count of `replica_zone_ids`, Non-multi-AZ does not require `replica_zone_ids`; Redis memory version 4.0, 5.0, 6.2 standard architecture and cluster architecture support the number of copies in the range [1, 2, 3, 4, 5]; Redis 2.8 standard version and CKV standard version only support 1 copy.
	// +kubebuilder:validation:Optional
	RedisReplicasNum *float64 `json:"redisReplicasNum,omitempty" tf:"redis_replicas_num,omitempty"`

	// The number of instance shards; this parameter does not need to be configured for standard version instances; for cluster version instances, the number of shards ranges from: [1, 3, 5, 8, 12, 16, 24 , 32, 40, 48, 64, 80, 96, 128].
	// The number of instance shards; this parameter does not need to be configured for standard version instances; for cluster version instances, the number of shards ranges from: [`1`, `3`, `5`, `8`, `12`, `16`, `24 `, `32`, `40`, `48`, `64`, `80`, `96`, `128`].
	// +kubebuilder:validation:Optional
	RedisShardNum *float64 `json:"redisShardNum,omitempty" tf:"redis_shard_num,omitempty"`

	// ID of replica nodes available zone. This is not required for standalone and master slave versions. NOTE: Removing some of the same zone of replicas (e.g. removing 100001 of [100001, 100001, 100002]) will pick the first hit to remove.
	// ID of replica nodes available zone. This is not required for standalone and master slave versions. NOTE: Removing some of the same zone of replicas (e.g. removing 100001 of [100001, 100001, 100002]) will pick the first hit to remove.
	// +kubebuilder:validation:Optional
	ReplicaZoneIds []*float64 `json:"replicaZoneIds,omitempty" tf:"replica_zone_ids,omitempty"`

	// Whether copy read-only is supported, Redis 2.8 Standard Edition and CKV Standard Edition do not support replica read-only, turn on replica read-only, the instance will automatically read and write separate, write requests are routed to the primary node, read requests are routed to the replica node, if you need to open replica read-only, the recommended number of replicas >=2.
	// Whether copy read-only is supported, Redis 2.8 Standard Edition and CKV Standard Edition do not support replica read-only, turn on replica read-only, the instance will automatically read and write separate, write requests are routed to the primary node, read requests are routed to the replica node, if you need to open replica read-only, the recommended number of replicas >=2.
	// +kubebuilder:validation:Optional
	ReplicasReadOnly *bool `json:"replicasReadOnly,omitempty" tf:"replicas_read_only,omitempty"`

	// ID of security group. If both vpc_id and subnet_id are not set, this argument should not be set either.
	// ID of security group. If both vpc_id and subnet_id are not set, this argument should not be set either.
	// +kubebuilder:validation:Optional
	// +listType=set
	SecurityGroups []*string `json:"securityGroups,omitempty" tf:"security_groups,omitempty"`

	// Specifies which subnet the instance should belong to. When the operation_network is changeVpc or changeBaseToVpc, this parameter needs to be configured.
	// Specifies which subnet the instance should belong to. When the `operation_network` is `changeVpc` or `changeBaseToVpc`, this parameter needs to be configured.
	// +kubebuilder:validation:Optional
	SubnetID *string `json:"subnetId,omitempty" tf:"subnet_id,omitempty"`

	// Instance tags.
	// Instance tags.
	// +kubebuilder:validation:Optional
	// +mapType=granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// It has been deprecated from version 1.33.1. Please use 'type_id' instead. Instance type. Available values: cluster_ckv,cluster_redis5.0,cluster_redis,master_slave_ckv,master_slave_redis4.0,master_slave_redis5.0,master_slave_redis,standalone_redis, specific region support specific types, need to refer data tencentcloud_redis_zone_config.
	// Instance type. Available values: `cluster_ckv`,`cluster_redis5.0`,`cluster_redis`,`master_slave_ckv`,`master_slave_redis4.0`,`master_slave_redis5.0`,`master_slave_redis`,`standalone_redis`, specific region support specific types, need to refer data `tencentcloud_redis_zone_config`.
	// +kubebuilder:validation:Optional
	Type *string `json:"type,omitempty" tf:"type,omitempty"`

	// Instance type. Available values reference data source tencentcloud_redis_zone_config or document, toggle immediately when modified.
	// Instance type. Available values reference data source `tencentcloud_redis_zone_config` or [document](https://intl.cloud.tencent.com/document/product/239/32069), toggle immediately when modified.
	// +kubebuilder:validation:Optional
	TypeID *float64 `json:"typeId,omitempty" tf:"type_id,omitempty"`

	// ID of the vpc with which the instance is to be associated. When the operation_network is changeVpc or changeBaseToVpc, this parameter needs to be configured.
	// ID of the vpc with which the instance is to be associated. When the `operation_network` is `changeVpc` or `changeBaseToVpc`, this parameter needs to be configured.
	// +kubebuilder:validation:Optional
	VPCID *string `json:"vpcId,omitempty" tf:"vpc_id,omitempty"`

	// Switching mode: 1-maintenance time window switching, 2-immediate switching, default value 2.
	// Switching mode: `1`-maintenance time window switching, `2`-immediate switching, default value `2`.
	// +kubebuilder:validation:Optional
	WaitSwitch *float64 `json:"waitSwitch,omitempty" tf:"wait_switch,omitempty"`
}

type NodeInfoInitParameters struct {
}

type NodeInfoObservation struct {

	// ID of the resource.
	ID *float64 `json:"id,omitempty" tf:"id,omitempty"`

	// Indicates whether the node is master.
	Master *bool `json:"master,omitempty" tf:"master,omitempty"`

	// ID of the availability zone of the master or replica node.
	ZoneID *float64 `json:"zoneId,omitempty" tf:"zone_id,omitempty"`
}

type NodeInfoParameters struct {
}

// InstanceSpec defines the desired state of Instance
type InstanceSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     InstanceParameters `json:"forProvider"`
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
	InitProvider InstanceInitParameters `json:"initProvider,omitempty"`
}

// InstanceStatus defines the observed state of Instance.
type InstanceStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        InstanceObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// Instance is the Schema for the Instances API. Provides a resource to create a Redis instance and set its attributes.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,tencentcloud}
type Instance struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.availabilityZone) || (has(self.initProvider) && has(self.initProvider.availabilityZone))",message="spec.forProvider.availabilityZone is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.memSize) || (has(self.initProvider) && has(self.initProvider.memSize))",message="spec.forProvider.memSize is a required parameter"
	Spec   InstanceSpec   `json:"spec"`
	Status InstanceStatus `json:"status,omitempty"`
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
