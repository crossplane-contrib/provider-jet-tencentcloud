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

type ScaleWorkerDataDiskObservation struct {
}

type ScaleWorkerDataDiskParameters struct {

	// Indicate whether to auto format and mount or not. Default is `false`.
	// +kubebuilder:validation:Optional
	AutoFormatAndMount *bool `json:"autoFormatAndMount,omitempty" tf:"auto_format_and_mount,omitempty"`

	// Volume of disk in GB. Default is `0`.
	// +kubebuilder:validation:Optional
	DiskSize *float64 `json:"diskSize,omitempty" tf:"disk_size,omitempty"`

	// Types of disk, available values: `CLOUD_PREMIUM` and `CLOUD_SSD` and `CLOUD_HSSD` and `CLOUD_TSSD`.
	// +kubebuilder:validation:Optional
	DiskType *string `json:"diskType,omitempty" tf:"disk_type,omitempty"`

	// File system, e.g. `ext3/ext4/xfs`.
	// +kubebuilder:validation:Optional
	FileSystem *string `json:"fileSystem,omitempty" tf:"file_system,omitempty"`

	// Mount target.
	// +kubebuilder:validation:Optional
	MountTarget *string `json:"mountTarget,omitempty" tf:"mount_target,omitempty"`
}

type ScaleWorkerObservation struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	WorkerInstancesList []ScaleWorkerWorkerInstancesListObservation `json:"workerInstancesList,omitempty" tf:"worker_instances_list,omitempty"`
}

type ScaleWorkerParameters struct {

	// ID of the cluster.
	// +crossplane:generate:reference:type=Cluster
	// +kubebuilder:validation:Optional
	ClusterID *string `json:"clusterId,omitempty" tf:"cluster_id,omitempty"`

	// +kubebuilder:validation:Optional
	ClusterIDRef *v1.Reference `json:"clusterIdRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	ClusterIDSelector *v1.Selector `json:"clusterIdSelector,omitempty" tf:"-"`

	// Configurations of data disk.
	// +kubebuilder:validation:Optional
	DataDisk []ScaleWorkerDataDiskParameters `json:"dataDisk,omitempty" tf:"data_disk,omitempty"`

	// Indicate to set desired pod number in current node. Valid when the cluster enable customized pod cidr.
	// +kubebuilder:validation:Optional
	DesiredPodNum *float64 `json:"desiredPodNum,omitempty" tf:"desired_pod_num,omitempty"`

	// Docker graph path. Default is `/var/lib/docker`.
	// +kubebuilder:validation:Optional
	DockerGraphPath *string `json:"dockerGraphPath,omitempty" tf:"docker_graph_path,omitempty"`

	// Custom parameter information related to the node.
	// +kubebuilder:validation:Optional
	ExtraArgs []*string `json:"extraArgs,omitempty" tf:"extra_args,omitempty"`

	// Labels of kubernetes scale worker created nodes.
	// +kubebuilder:validation:Optional
	Labels map[string]*string `json:"labels,omitempty" tf:"labels,omitempty"`

	// Mount target. Default is not mounting.
	// +kubebuilder:validation:Optional
	MountTarget *string `json:"mountTarget,omitempty" tf:"mount_target,omitempty"`

	// Sets whether the joining node participates in the schedule. Default is '0'. Participate in scheduling.
	// +kubebuilder:validation:Optional
	Unschedulable *float64 `json:"unschedulable,omitempty" tf:"unschedulable,omitempty"`

	// Deploy the machine configuration information of the 'WORK' service, and create <=20 units for common users.
	// +kubebuilder:validation:Required
	WorkerConfig []ScaleWorkerWorkerConfigParameters `json:"workerConfig" tf:"worker_config,omitempty"`
}

type ScaleWorkerWorkerConfigDataDiskObservation struct {
}

type ScaleWorkerWorkerConfigDataDiskParameters struct {

	// Indicate whether to auto format and mount or not. Default is `false`.
	// +kubebuilder:validation:Optional
	AutoFormatAndMount *bool `json:"autoFormatAndMount,omitempty" tf:"auto_format_and_mount,omitempty"`

	// The name of the device or partition to mount.
	// +kubebuilder:validation:Optional
	DiskPartition *string `json:"diskPartition,omitempty" tf:"disk_partition,omitempty"`

	// Volume of disk in GB. Default is `0`.
	// +kubebuilder:validation:Optional
	DiskSize *float64 `json:"diskSize,omitempty" tf:"disk_size,omitempty"`

	// Types of disk, available values: `CLOUD_PREMIUM` and `CLOUD_SSD` and `CLOUD_HSSD` and `CLOUD_TSSD`.
	// +kubebuilder:validation:Optional
	DiskType *string `json:"diskType,omitempty" tf:"disk_type,omitempty"`

	// Indicates whether to encrypt data disk, default `false`.
	// +kubebuilder:validation:Optional
	Encrypt *bool `json:"encrypt,omitempty" tf:"encrypt,omitempty"`

	// File system, e.g. `ext3/ext4/xfs`.
	// +kubebuilder:validation:Optional
	FileSystem *string `json:"fileSystem,omitempty" tf:"file_system,omitempty"`

	// ID of the custom CMK in the format of UUID or `kms-abcd1234`. This parameter is used to encrypt cloud disks.
	// +kubebuilder:validation:Optional
	KMSKeyID *string `json:"kmsKeyId,omitempty" tf:"kms_key_id,omitempty"`

	// Mount target.
	// +kubebuilder:validation:Optional
	MountTarget *string `json:"mountTarget,omitempty" tf:"mount_target,omitempty"`

	// Data disk snapshot ID.
	// +kubebuilder:validation:Optional
	SnapshotID *string `json:"snapshotId,omitempty" tf:"snapshot_id,omitempty"`
}

type ScaleWorkerWorkerConfigObservation struct {
}

type ScaleWorkerWorkerConfigParameters struct {

	// Indicates which availability zone will be used.
	// +kubebuilder:validation:Optional
	AvailabilityZone *string `json:"availabilityZone,omitempty" tf:"availability_zone,omitempty"`

	// bandwidth package id. if user is standard user, then the bandwidth_package_id is needed, or default has bandwidth_package_id.
	// +kubebuilder:validation:Optional
	BandwidthPackageID *string `json:"bandwidthPackageId,omitempty" tf:"bandwidth_package_id,omitempty"`

	// CAM role name authorized to access.
	// +kubebuilder:validation:Optional
	CamRoleName *string `json:"camRoleName,omitempty" tf:"cam_role_name,omitempty"`

	// Number of cvm.
	// +kubebuilder:validation:Optional
	Count *float64 `json:"count,omitempty" tf:"count,omitempty"`

	// Configurations of data disk.
	// +kubebuilder:validation:Optional
	DataDisk []ScaleWorkerWorkerConfigDataDiskParameters `json:"dataDisk,omitempty" tf:"data_disk,omitempty"`

	// Indicate to set desired pod number in node. valid when enable_customized_pod_cidr=true, and it override `[globe_]desired_pod_num` for current node. Either all the fields `desired_pod_num` or none.
	// +kubebuilder:validation:Optional
	DesiredPodNum *float64 `json:"desiredPodNum,omitempty" tf:"desired_pod_num,omitempty"`

	// Disaster recover groups to which a CVM instance belongs. Only support maximum 1.
	// +kubebuilder:validation:Optional
	DisasterRecoverGroupIds []*string `json:"disasterRecoverGroupIds,omitempty" tf:"disaster_recover_group_ids,omitempty"`

	// To specify whether to enable cloud monitor service. Default is TRUE.
	// +kubebuilder:validation:Optional
	EnhancedMonitorService *bool `json:"enhancedMonitorService,omitempty" tf:"enhanced_monitor_service,omitempty"`

	// To specify whether to enable cloud security service. Default is TRUE.
	// +kubebuilder:validation:Optional
	EnhancedSecurityService *bool `json:"enhancedSecurityService,omitempty" tf:"enhanced_security_service,omitempty"`

	// The host name of the attached instance. Dot (.) and dash (-) cannot be used as the first and last characters of HostName and cannot be used consecutively. Windows example: The length of the name character is [2, 15], letters (capitalization is not restricted), numbers and dashes (-) are allowed, dots (.) are not supported, and not all numbers are allowed. Examples of other types (Linux, etc.): The character length is [2, 60], and multiple dots are allowed. There is a segment between the dots. Each segment allows letters (with no limitation on capitalization), numbers and dashes (-).
	// +kubebuilder:validation:Optional
	Hostname *string `json:"hostname,omitempty" tf:"hostname,omitempty"`

	// The valid image id, format of img-xxx.
	// +kubebuilder:validation:Optional
	ImgID *string `json:"imgId,omitempty" tf:"img_id,omitempty"`

	// The charge type of instance. Valid values are `PREPAID` and `POSTPAID_BY_HOUR`. The default is `POSTPAID_BY_HOUR`. Note: TencentCloud International only supports `POSTPAID_BY_HOUR`, `PREPAID` instance will not terminated after cluster deleted, and may not allow to delete before expired.
	// +kubebuilder:validation:Optional
	InstanceChargeType *string `json:"instanceChargeType,omitempty" tf:"instance_charge_type,omitempty"`

	// The tenancy (time unit is month) of the prepaid instance. NOTE: it only works when instance_charge_type is set to `PREPAID`. Valid values are `1`, `2`, `3`, `4`, `5`, `6`, `7`, `8`, `9`, `10`, `11`, `12`, `24`, `36`.
	// +kubebuilder:validation:Optional
	InstanceChargeTypePrepaidPeriod *float64 `json:"instanceChargeTypePrepaidPeriod,omitempty" tf:"instance_charge_type_prepaid_period,omitempty"`

	// Auto renewal flag. Valid values: `NOTIFY_AND_AUTO_RENEW`: notify upon expiration and renew automatically, `NOTIFY_AND_MANUAL_RENEW`: notify upon expiration but do not renew automatically, `DISABLE_NOTIFY_AND_MANUAL_RENEW`: neither notify upon expiration nor renew automatically. Default value: `NOTIFY_AND_MANUAL_RENEW`. If this parameter is specified as `NOTIFY_AND_AUTO_RENEW`, the instance will be automatically renewed on a monthly basis if the account balance is sufficient. NOTE: it only works when instance_charge_type is set to `PREPAID`.
	// +kubebuilder:validation:Optional
	InstanceChargeTypePrepaidRenewFlag *string `json:"instanceChargeTypePrepaidRenewFlag,omitempty" tf:"instance_charge_type_prepaid_renew_flag,omitempty"`

	// Name of the CVMs.
	// +kubebuilder:validation:Optional
	InstanceName *string `json:"instanceName,omitempty" tf:"instance_name,omitempty"`

	// Specified types of CVM instance.
	// +kubebuilder:validation:Required
	InstanceType *string `json:"instanceType" tf:"instance_type,omitempty"`

	// Charge types for network traffic. Available values include `TRAFFIC_POSTPAID_BY_HOUR`.
	// +kubebuilder:validation:Optional
	InternetChargeType *string `json:"internetChargeType,omitempty" tf:"internet_charge_type,omitempty"`

	// Max bandwidth of Internet access in Mbps. Default is 0.
	// +kubebuilder:validation:Optional
	InternetMaxBandwidthOut *float64 `json:"internetMaxBandwidthOut,omitempty" tf:"internet_max_bandwidth_out,omitempty"`

	// ID list of keys, should be set if `password` not set.
	// +kubebuilder:validation:Optional
	KeyIds []*string `json:"keyIds,omitempty" tf:"key_ids,omitempty"`

	// Password to access, should be set if `key_ids` not set.
	// +kubebuilder:validation:Optional
	PasswordSecretRef *v1.SecretKeySelector `json:"passwordSecretRef,omitempty" tf:"-"`

	// Specify whether to assign an Internet IP address.
	// +kubebuilder:validation:Optional
	PublicIPAssigned *bool `json:"publicIpAssigned,omitempty" tf:"public_ip_assigned,omitempty"`

	// Security groups to which a CVM instance belongs.
	// +kubebuilder:validation:Optional
	SecurityGroupIds []*string `json:"securityGroupIds,omitempty" tf:"security_group_ids,omitempty"`

	// Private network ID.
	// +kubebuilder:validation:Required
	SubnetID *string `json:"subnetId" tf:"subnet_id,omitempty"`

	// Volume of system disk in GB. Default is `50`.
	// +kubebuilder:validation:Optional
	SystemDiskSize *float64 `json:"systemDiskSize,omitempty" tf:"system_disk_size,omitempty"`

	// System disk type. For more information on limits of system disk types, see [Storage Overview](https://intl.cloud.tencent.com/document/product/213/4952). Valid values: `LOCAL_BASIC`: local disk, `LOCAL_SSD`: local SSD disk, `CLOUD_SSD`: SSD, `CLOUD_PREMIUM`: Premium Cloud Storage. NOTE: `CLOUD_BASIC`, `LOCAL_BASIC` and `LOCAL_SSD` are deprecated.
	// +kubebuilder:validation:Optional
	SystemDiskType *string `json:"systemDiskType,omitempty" tf:"system_disk_type,omitempty"`

	// ase64-encoded User Data text, the length limit is 16KB.
	// +kubebuilder:validation:Optional
	UserData *string `json:"userData,omitempty" tf:"user_data,omitempty"`
}

type ScaleWorkerWorkerInstancesListObservation struct {
	FailedReason *string `json:"failedReason,omitempty" tf:"failed_reason,omitempty"`

	InstanceID *string `json:"instanceId,omitempty" tf:"instance_id,omitempty"`

	InstanceRole *string `json:"instanceRole,omitempty" tf:"instance_role,omitempty"`

	InstanceState *string `json:"instanceState,omitempty" tf:"instance_state,omitempty"`

	LanIP *string `json:"lanIp,omitempty" tf:"lan_ip,omitempty"`
}

type ScaleWorkerWorkerInstancesListParameters struct {
}

// ScaleWorkerSpec defines the desired state of ScaleWorker
type ScaleWorkerSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ScaleWorkerParameters `json:"forProvider"`
}

// ScaleWorkerStatus defines the observed state of ScaleWorker.
type ScaleWorkerStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ScaleWorkerObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// ScaleWorker is the Schema for the ScaleWorkers API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,tencentcloudjet}
type ScaleWorker struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ScaleWorkerSpec   `json:"spec"`
	Status            ScaleWorkerStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ScaleWorkerList contains a list of ScaleWorkers
type ScaleWorkerList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ScaleWorker `json:"items"`
}

// Repository type metadata.
var (
	ScaleWorker_Kind             = "ScaleWorker"
	ScaleWorker_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: ScaleWorker_Kind}.String()
	ScaleWorker_KindAPIVersion   = ScaleWorker_Kind + "." + CRDGroupVersion.String()
	ScaleWorker_GroupVersionKind = CRDGroupVersion.WithKind(ScaleWorker_Kind)
)

func init() {
	SchemeBuilder.Register(&ScaleWorker{}, &ScaleWorkerList{})
}