// SPDX-FileCopyrightText: 2023 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

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

type InstanceSetInitParameters struct {

	// Associate a public IP address with an instance in a VPC or Classic. Boolean value, Default is false.
	// Associate a public IP address with an instance in a VPC or Classic. Boolean value, Default is false.
	AllocatePublicIP *bool `json:"allocatePublicIp,omitempty" tf:"allocate_public_ip,omitempty"`

	// The available zone for the CVM instance.
	// The available zone for the CVM instance.
	AvailabilityZone *string `json:"availabilityZone,omitempty" tf:"availability_zone,omitempty"`

	// bandwidth package id. if user is standard user, then the bandwidth_package_id is needed, or default has bandwidth_package_id.
	// bandwidth package id. if user is standard user, then the bandwidth_package_id is needed, or default has bandwidth_package_id.
	BandwidthPackageID *string `json:"bandwidthPackageId,omitempty" tf:"bandwidth_package_id,omitempty"`

	// CAM role name authorized to access.
	// CAM role name authorized to access.
	CamRoleName *string `json:"camRoleName,omitempty" tf:"cam_role_name,omitempty"`

	// Disable enhance service for monitor, it is enabled by default. When this options is set, monitor agent won't be installed. Modifying will cause the instance reset.
	// Disable enhance service for monitor, it is enabled by default. When this options is set, monitor agent won't be installed. Modifying will cause the instance reset.
	DisableMonitorService *bool `json:"disableMonitorService,omitempty" tf:"disable_monitor_service,omitempty"`

	// Disable enhance service for security, it is enabled by default. When this options is set, security agent won't be installed. Modifying will cause the instance reset.
	// Disable enhance service for security, it is enabled by default. When this options is set, security agent won't be installed. Modifying will cause the instance reset.
	DisableSecurityService *bool `json:"disableSecurityService,omitempty" tf:"disable_security_service,omitempty"`

	// instance ids list to exclude.
	// instance ids list to exclude.
	ExcludeInstanceIds []*string `json:"excludeInstanceIds,omitempty" tf:"exclude_instance_ids,omitempty"`

	// The hostname of the instance. Windows instance: The name should be a combination of 2 to 15 characters comprised of letters (case insensitive), numbers, and hyphens (-). Period (.) is not supported, and the name cannot be a string of pure numbers. Other types (such as Linux) of instances: The name should be a combination of 2 to 60 characters, supporting multiple periods (.). The piece between two periods is composed of letters (case insensitive), numbers, and hyphens (-). Modifying will cause the instance reset.
	// The hostname of the instance. Windows instance: The name should be a combination of 2 to 15 characters comprised of letters (case insensitive), numbers, and hyphens (-). Period (.) is not supported, and the name cannot be a string of pure numbers. Other types (such as Linux) of instances: The name should be a combination of 2 to 60 characters, supporting multiple periods (.). The piece between two periods is composed of letters (case insensitive), numbers, and hyphens (-). Modifying will cause the instance reset.
	Hostname *string `json:"hostname,omitempty" tf:"hostname,omitempty"`

	// The image to use for the instance. Changing image_id will cause the instance reset.
	// The image to use for the instance. Changing `image_id` will cause the instance reset.
	ImageID *string `json:"imageId,omitempty" tf:"image_id,omitempty"`

	// The charge type of instance. Only support POSTPAID_BY_HOUR.
	// The charge type of instance. Only support `POSTPAID_BY_HOUR`.
	InstanceChargeType *string `json:"instanceChargeType,omitempty" tf:"instance_charge_type,omitempty"`

	// The number of instances to be purchased. Value range:[1,100]; default value: 1.
	// The number of instances to be purchased. Value range:[1,100]; default value: 1.
	InstanceCount *float64 `json:"instanceCount,omitempty" tf:"instance_count,omitempty"`

	// The name of the instance.
	// The name of the instance.
	InstanceName *string `json:"instanceName,omitempty" tf:"instance_name,omitempty"`

	// The type of the instance.
	// The type of the instance.
	InstanceType *string `json:"instanceType,omitempty" tf:"instance_type,omitempty"`

	// Internet charge type of the instance, Valid values are BANDWIDTH_PREPAID, TRAFFIC_POSTPAID_BY_HOUR, BANDWIDTH_POSTPAID_BY_HOUR and BANDWIDTH_PACKAGE. This value does not need to be set when allocate_public_ip is false.
	// Internet charge type of the instance, Valid values are `BANDWIDTH_PREPAID`, `TRAFFIC_POSTPAID_BY_HOUR`, `BANDWIDTH_POSTPAID_BY_HOUR` and `BANDWIDTH_PACKAGE`. This value does not need to be set when `allocate_public_ip` is false.
	InternetChargeType *string `json:"internetChargeType,omitempty" tf:"internet_charge_type,omitempty"`

	// Maximum outgoing bandwidth to the public network, measured in Mbps (Mega bits per second). This value does not need to be set when allocate_public_ip is false.
	// Maximum outgoing bandwidth to the public network, measured in Mbps (Mega bits per second). This value does not need to be set when `allocate_public_ip` is false.
	InternetMaxBandwidthOut *float64 `json:"internetMaxBandwidthOut,omitempty" tf:"internet_max_bandwidth_out,omitempty"`

	// Whether to keep image login or not, default is false. When the image type is private or shared or imported, this parameter can be set true. Modifying will cause the instance reset.
	// Whether to keep image login or not, default is `false`. When the image type is private or shared or imported, this parameter can be set `true`. Modifying will cause the instance reset.
	KeepImageLogin *bool `json:"keepImageLogin,omitempty" tf:"keep_image_login,omitempty"`

	// The key pair to use for the instance, it looks like skey-16jig7tx. Modifying will cause the instance reset.
	// The key pair to use for the instance, it looks like `skey-16jig7tx`. Modifying will cause the instance reset.
	KeyName *string `json:"keyName,omitempty" tf:"key_name,omitempty"`

	// The ID of a placement group.
	// The ID of a placement group.
	PlacementGroupID *string `json:"placementGroupId,omitempty" tf:"placement_group_id,omitempty"`

	// The private IP to be assigned to this instance, must be in the provided subnet and available.
	// The private IP to be assigned to this instance, must be in the provided subnet and available.
	PrivateIP *string `json:"privateIp,omitempty" tf:"private_ip,omitempty"`

	// The project the instance belongs to, default to 0.
	// The project the instance belongs to, default to 0.
	ProjectID *float64 `json:"projectId,omitempty" tf:"project_id,omitempty"`

	// A list of security group IDs to associate with.
	// A list of security group IDs to associate with.
	SecurityGroups []*string `json:"securityGroups,omitempty" tf:"security_groups,omitempty"`

	// System disk snapshot ID used to initialize the system disk. When system disk type is LOCAL_BASIC and LOCAL_SSD, disk id is not supported.
	// System disk snapshot ID used to initialize the system disk. When system disk type is `LOCAL_BASIC` and `LOCAL_SSD`, disk id is not supported.
	SystemDiskID *string `json:"systemDiskId,omitempty" tf:"system_disk_id,omitempty"`

	// Size of the system disk. Valid value ranges: (50~1000). and unit is GB. Default is 50GB. If modified, the instance may force stop.
	// Size of the system disk. Valid value ranges: (50~1000). and unit is GB. Default is 50GB. If modified, the instance may force stop.
	SystemDiskSize *float64 `json:"systemDiskSize,omitempty" tf:"system_disk_size,omitempty"`

	// System disk type. For more information on limits of system disk types, see Storage Overview. Valid values: LOCAL_BASIC: local disk, LOCAL_SSD: local SSD disk, CLOUD_SSD: SSD, CLOUD_PREMIUM: Premium Cloud Storage, CLOUD_BSSD: Basic SSD. NOTE: If modified, the instance may force stop.
	// System disk type. For more information on limits of system disk types, see [Storage Overview](https://intl.cloud.tencent.com/document/product/213/4952). Valid values: `LOCAL_BASIC`: local disk, `LOCAL_SSD`: local SSD disk, `CLOUD_SSD`: SSD, `CLOUD_PREMIUM`: Premium Cloud Storage, `CLOUD_BSSD`: Basic SSD. NOTE: If modified, the instance may force stop.
	SystemDiskType *string `json:"systemDiskType,omitempty" tf:"system_disk_type,omitempty"`

	// The user data to be injected into this instance. Must be base64 encoded and up to 16 KB.
	// The user data to be injected into this instance. Must be base64 encoded and up to 16 KB.
	UserData *string `json:"userData,omitempty" tf:"user_data,omitempty"`

	// The user data to be injected into this instance, in plain text. Conflicts with user_data. Up to 16 KB after base64 encoded.
	// The user data to be injected into this instance, in plain text. Conflicts with `user_data`. Up to 16 KB after base64 encoded.
	UserDataRaw *string `json:"userDataRaw,omitempty" tf:"user_data_raw,omitempty"`
}

type InstanceSetObservation struct {

	// Associate a public IP address with an instance in a VPC or Classic. Boolean value, Default is false.
	// Associate a public IP address with an instance in a VPC or Classic. Boolean value, Default is false.
	AllocatePublicIP *bool `json:"allocatePublicIp,omitempty" tf:"allocate_public_ip,omitempty"`

	// The available zone for the CVM instance.
	// The available zone for the CVM instance.
	AvailabilityZone *string `json:"availabilityZone,omitempty" tf:"availability_zone,omitempty"`

	// bandwidth package id. if user is standard user, then the bandwidth_package_id is needed, or default has bandwidth_package_id.
	// bandwidth package id. if user is standard user, then the bandwidth_package_id is needed, or default has bandwidth_package_id.
	BandwidthPackageID *string `json:"bandwidthPackageId,omitempty" tf:"bandwidth_package_id,omitempty"`

	// CAM role name authorized to access.
	// CAM role name authorized to access.
	CamRoleName *string `json:"camRoleName,omitempty" tf:"cam_role_name,omitempty"`

	// Create time of the instance.
	// Create time of the instance.
	CreateTime *string `json:"createTime,omitempty" tf:"create_time,omitempty"`

	// Disable enhance service for monitor, it is enabled by default. When this options is set, monitor agent won't be installed. Modifying will cause the instance reset.
	// Disable enhance service for monitor, it is enabled by default. When this options is set, monitor agent won't be installed. Modifying will cause the instance reset.
	DisableMonitorService *bool `json:"disableMonitorService,omitempty" tf:"disable_monitor_service,omitempty"`

	// Disable enhance service for security, it is enabled by default. When this options is set, security agent won't be installed. Modifying will cause the instance reset.
	// Disable enhance service for security, it is enabled by default. When this options is set, security agent won't be installed. Modifying will cause the instance reset.
	DisableSecurityService *bool `json:"disableSecurityService,omitempty" tf:"disable_security_service,omitempty"`

	// instance ids list to exclude.
	// instance ids list to exclude.
	ExcludeInstanceIds []*string `json:"excludeInstanceIds,omitempty" tf:"exclude_instance_ids,omitempty"`

	// Expired time of the instance.
	// Expired time of the instance.
	ExpiredTime *string `json:"expiredTime,omitempty" tf:"expired_time,omitempty"`

	// The hostname of the instance. Windows instance: The name should be a combination of 2 to 15 characters comprised of letters (case insensitive), numbers, and hyphens (-). Period (.) is not supported, and the name cannot be a string of pure numbers. Other types (such as Linux) of instances: The name should be a combination of 2 to 60 characters, supporting multiple periods (.). The piece between two periods is composed of letters (case insensitive), numbers, and hyphens (-). Modifying will cause the instance reset.
	// The hostname of the instance. Windows instance: The name should be a combination of 2 to 15 characters comprised of letters (case insensitive), numbers, and hyphens (-). Period (.) is not supported, and the name cannot be a string of pure numbers. Other types (such as Linux) of instances: The name should be a combination of 2 to 60 characters, supporting multiple periods (.). The piece between two periods is composed of letters (case insensitive), numbers, and hyphens (-). Modifying will cause the instance reset.
	Hostname *string `json:"hostname,omitempty" tf:"hostname,omitempty"`

	// ID of the resource.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The image to use for the instance. Changing image_id will cause the instance reset.
	// The image to use for the instance. Changing `image_id` will cause the instance reset.
	ImageID *string `json:"imageId,omitempty" tf:"image_id,omitempty"`

	// The charge type of instance. Only support POSTPAID_BY_HOUR.
	// The charge type of instance. Only support `POSTPAID_BY_HOUR`.
	InstanceChargeType *string `json:"instanceChargeType,omitempty" tf:"instance_charge_type,omitempty"`

	// The number of instances to be purchased. Value range:[1,100]; default value: 1.
	// The number of instances to be purchased. Value range:[1,100]; default value: 1.
	InstanceCount *float64 `json:"instanceCount,omitempty" tf:"instance_count,omitempty"`

	// instance id list.
	// instance id list.
	InstanceIds []*string `json:"instanceIds,omitempty" tf:"instance_ids,omitempty"`

	// The name of the instance.
	// The name of the instance.
	InstanceName *string `json:"instanceName,omitempty" tf:"instance_name,omitempty"`

	// Current status of the instance.
	// Current status of the instance.
	InstanceStatus *string `json:"instanceStatus,omitempty" tf:"instance_status,omitempty"`

	// The type of the instance.
	// The type of the instance.
	InstanceType *string `json:"instanceType,omitempty" tf:"instance_type,omitempty"`

	// Internet charge type of the instance, Valid values are BANDWIDTH_PREPAID, TRAFFIC_POSTPAID_BY_HOUR, BANDWIDTH_POSTPAID_BY_HOUR and BANDWIDTH_PACKAGE. This value does not need to be set when allocate_public_ip is false.
	// Internet charge type of the instance, Valid values are `BANDWIDTH_PREPAID`, `TRAFFIC_POSTPAID_BY_HOUR`, `BANDWIDTH_POSTPAID_BY_HOUR` and `BANDWIDTH_PACKAGE`. This value does not need to be set when `allocate_public_ip` is false.
	InternetChargeType *string `json:"internetChargeType,omitempty" tf:"internet_charge_type,omitempty"`

	// Maximum outgoing bandwidth to the public network, measured in Mbps (Mega bits per second). This value does not need to be set when allocate_public_ip is false.
	// Maximum outgoing bandwidth to the public network, measured in Mbps (Mega bits per second). This value does not need to be set when `allocate_public_ip` is false.
	InternetMaxBandwidthOut *float64 `json:"internetMaxBandwidthOut,omitempty" tf:"internet_max_bandwidth_out,omitempty"`

	// Whether to keep image login or not, default is false. When the image type is private or shared or imported, this parameter can be set true. Modifying will cause the instance reset.
	// Whether to keep image login or not, default is `false`. When the image type is private or shared or imported, this parameter can be set `true`. Modifying will cause the instance reset.
	KeepImageLogin *bool `json:"keepImageLogin,omitempty" tf:"keep_image_login,omitempty"`

	// The key pair to use for the instance, it looks like skey-16jig7tx. Modifying will cause the instance reset.
	// The key pair to use for the instance, it looks like `skey-16jig7tx`. Modifying will cause the instance reset.
	KeyName *string `json:"keyName,omitempty" tf:"key_name,omitempty"`

	// The ID of a placement group.
	// The ID of a placement group.
	PlacementGroupID *string `json:"placementGroupId,omitempty" tf:"placement_group_id,omitempty"`

	// The private IP to be assigned to this instance, must be in the provided subnet and available.
	// The private IP to be assigned to this instance, must be in the provided subnet and available.
	PrivateIP *string `json:"privateIp,omitempty" tf:"private_ip,omitempty"`

	// The project the instance belongs to, default to 0.
	// The project the instance belongs to, default to 0.
	ProjectID *float64 `json:"projectId,omitempty" tf:"project_id,omitempty"`

	// Public IP of the instance.
	// Public IP of the instance.
	PublicIP *string `json:"publicIp,omitempty" tf:"public_ip,omitempty"`

	// A list of security group IDs to associate with.
	// A list of security group IDs to associate with.
	SecurityGroups []*string `json:"securityGroups,omitempty" tf:"security_groups,omitempty"`

	// The ID of a VPC subnet. If you want to create instances in a VPC network, this parameter must be set.
	// The ID of a VPC subnet. If you want to create instances in a VPC network, this parameter must be set.
	SubnetID *string `json:"subnetId,omitempty" tf:"subnet_id,omitempty"`

	// System disk snapshot ID used to initialize the system disk. When system disk type is LOCAL_BASIC and LOCAL_SSD, disk id is not supported.
	// System disk snapshot ID used to initialize the system disk. When system disk type is `LOCAL_BASIC` and `LOCAL_SSD`, disk id is not supported.
	SystemDiskID *string `json:"systemDiskId,omitempty" tf:"system_disk_id,omitempty"`

	// Size of the system disk. Valid value ranges: (50~1000). and unit is GB. Default is 50GB. If modified, the instance may force stop.
	// Size of the system disk. Valid value ranges: (50~1000). and unit is GB. Default is 50GB. If modified, the instance may force stop.
	SystemDiskSize *float64 `json:"systemDiskSize,omitempty" tf:"system_disk_size,omitempty"`

	// System disk type. For more information on limits of system disk types, see Storage Overview. Valid values: LOCAL_BASIC: local disk, LOCAL_SSD: local SSD disk, CLOUD_SSD: SSD, CLOUD_PREMIUM: Premium Cloud Storage, CLOUD_BSSD: Basic SSD. NOTE: If modified, the instance may force stop.
	// System disk type. For more information on limits of system disk types, see [Storage Overview](https://intl.cloud.tencent.com/document/product/213/4952). Valid values: `LOCAL_BASIC`: local disk, `LOCAL_SSD`: local SSD disk, `CLOUD_SSD`: SSD, `CLOUD_PREMIUM`: Premium Cloud Storage, `CLOUD_BSSD`: Basic SSD. NOTE: If modified, the instance may force stop.
	SystemDiskType *string `json:"systemDiskType,omitempty" tf:"system_disk_type,omitempty"`

	// The user data to be injected into this instance. Must be base64 encoded and up to 16 KB.
	// The user data to be injected into this instance. Must be base64 encoded and up to 16 KB.
	UserData *string `json:"userData,omitempty" tf:"user_data,omitempty"`

	// The user data to be injected into this instance, in plain text. Conflicts with user_data. Up to 16 KB after base64 encoded.
	// The user data to be injected into this instance, in plain text. Conflicts with `user_data`. Up to 16 KB after base64 encoded.
	UserDataRaw *string `json:"userDataRaw,omitempty" tf:"user_data_raw,omitempty"`

	// The ID of a VPC network. If you want to create instances in a VPC network, this parameter must be set.
	// The ID of a VPC network. If you want to create instances in a VPC network, this parameter must be set.
	VPCID *string `json:"vpcId,omitempty" tf:"vpc_id,omitempty"`
}

type InstanceSetParameters struct {

	// Associate a public IP address with an instance in a VPC or Classic. Boolean value, Default is false.
	// Associate a public IP address with an instance in a VPC or Classic. Boolean value, Default is false.
	// +kubebuilder:validation:Optional
	AllocatePublicIP *bool `json:"allocatePublicIp,omitempty" tf:"allocate_public_ip,omitempty"`

	// The available zone for the CVM instance.
	// The available zone for the CVM instance.
	// +kubebuilder:validation:Optional
	AvailabilityZone *string `json:"availabilityZone,omitempty" tf:"availability_zone,omitempty"`

	// bandwidth package id. if user is standard user, then the bandwidth_package_id is needed, or default has bandwidth_package_id.
	// bandwidth package id. if user is standard user, then the bandwidth_package_id is needed, or default has bandwidth_package_id.
	// +kubebuilder:validation:Optional
	BandwidthPackageID *string `json:"bandwidthPackageId,omitempty" tf:"bandwidth_package_id,omitempty"`

	// CAM role name authorized to access.
	// CAM role name authorized to access.
	// +kubebuilder:validation:Optional
	CamRoleName *string `json:"camRoleName,omitempty" tf:"cam_role_name,omitempty"`

	// Disable enhance service for monitor, it is enabled by default. When this options is set, monitor agent won't be installed. Modifying will cause the instance reset.
	// Disable enhance service for monitor, it is enabled by default. When this options is set, monitor agent won't be installed. Modifying will cause the instance reset.
	// +kubebuilder:validation:Optional
	DisableMonitorService *bool `json:"disableMonitorService,omitempty" tf:"disable_monitor_service,omitempty"`

	// Disable enhance service for security, it is enabled by default. When this options is set, security agent won't be installed. Modifying will cause the instance reset.
	// Disable enhance service for security, it is enabled by default. When this options is set, security agent won't be installed. Modifying will cause the instance reset.
	// +kubebuilder:validation:Optional
	DisableSecurityService *bool `json:"disableSecurityService,omitempty" tf:"disable_security_service,omitempty"`

	// instance ids list to exclude.
	// instance ids list to exclude.
	// +kubebuilder:validation:Optional
	ExcludeInstanceIds []*string `json:"excludeInstanceIds,omitempty" tf:"exclude_instance_ids,omitempty"`

	// The hostname of the instance. Windows instance: The name should be a combination of 2 to 15 characters comprised of letters (case insensitive), numbers, and hyphens (-). Period (.) is not supported, and the name cannot be a string of pure numbers. Other types (such as Linux) of instances: The name should be a combination of 2 to 60 characters, supporting multiple periods (.). The piece between two periods is composed of letters (case insensitive), numbers, and hyphens (-). Modifying will cause the instance reset.
	// The hostname of the instance. Windows instance: The name should be a combination of 2 to 15 characters comprised of letters (case insensitive), numbers, and hyphens (-). Period (.) is not supported, and the name cannot be a string of pure numbers. Other types (such as Linux) of instances: The name should be a combination of 2 to 60 characters, supporting multiple periods (.). The piece between two periods is composed of letters (case insensitive), numbers, and hyphens (-). Modifying will cause the instance reset.
	// +kubebuilder:validation:Optional
	Hostname *string `json:"hostname,omitempty" tf:"hostname,omitempty"`

	// The image to use for the instance. Changing image_id will cause the instance reset.
	// The image to use for the instance. Changing `image_id` will cause the instance reset.
	// +kubebuilder:validation:Optional
	ImageID *string `json:"imageId,omitempty" tf:"image_id,omitempty"`

	// The charge type of instance. Only support POSTPAID_BY_HOUR.
	// The charge type of instance. Only support `POSTPAID_BY_HOUR`.
	// +kubebuilder:validation:Optional
	InstanceChargeType *string `json:"instanceChargeType,omitempty" tf:"instance_charge_type,omitempty"`

	// The number of instances to be purchased. Value range:[1,100]; default value: 1.
	// The number of instances to be purchased. Value range:[1,100]; default value: 1.
	// +kubebuilder:validation:Optional
	InstanceCount *float64 `json:"instanceCount,omitempty" tf:"instance_count,omitempty"`

	// The name of the instance.
	// The name of the instance.
	// +kubebuilder:validation:Optional
	InstanceName *string `json:"instanceName,omitempty" tf:"instance_name,omitempty"`

	// The type of the instance.
	// The type of the instance.
	// +kubebuilder:validation:Optional
	InstanceType *string `json:"instanceType,omitempty" tf:"instance_type,omitempty"`

	// Internet charge type of the instance, Valid values are BANDWIDTH_PREPAID, TRAFFIC_POSTPAID_BY_HOUR, BANDWIDTH_POSTPAID_BY_HOUR and BANDWIDTH_PACKAGE. This value does not need to be set when allocate_public_ip is false.
	// Internet charge type of the instance, Valid values are `BANDWIDTH_PREPAID`, `TRAFFIC_POSTPAID_BY_HOUR`, `BANDWIDTH_POSTPAID_BY_HOUR` and `BANDWIDTH_PACKAGE`. This value does not need to be set when `allocate_public_ip` is false.
	// +kubebuilder:validation:Optional
	InternetChargeType *string `json:"internetChargeType,omitempty" tf:"internet_charge_type,omitempty"`

	// Maximum outgoing bandwidth to the public network, measured in Mbps (Mega bits per second). This value does not need to be set when allocate_public_ip is false.
	// Maximum outgoing bandwidth to the public network, measured in Mbps (Mega bits per second). This value does not need to be set when `allocate_public_ip` is false.
	// +kubebuilder:validation:Optional
	InternetMaxBandwidthOut *float64 `json:"internetMaxBandwidthOut,omitempty" tf:"internet_max_bandwidth_out,omitempty"`

	// Whether to keep image login or not, default is false. When the image type is private or shared or imported, this parameter can be set true. Modifying will cause the instance reset.
	// Whether to keep image login or not, default is `false`. When the image type is private or shared or imported, this parameter can be set `true`. Modifying will cause the instance reset.
	// +kubebuilder:validation:Optional
	KeepImageLogin *bool `json:"keepImageLogin,omitempty" tf:"keep_image_login,omitempty"`

	// The key pair to use for the instance, it looks like skey-16jig7tx. Modifying will cause the instance reset.
	// The key pair to use for the instance, it looks like `skey-16jig7tx`. Modifying will cause the instance reset.
	// +kubebuilder:validation:Optional
	KeyName *string `json:"keyName,omitempty" tf:"key_name,omitempty"`

	// Password for the instance. In order for the new password to take effect, the instance will be restarted after the password change. Modifying will cause the instance reset.
	// Password for the instance. In order for the new password to take effect, the instance will be restarted after the password change. Modifying will cause the instance reset.
	// +kubebuilder:validation:Optional
	PasswordSecretRef *v1.SecretKeySelector `json:"passwordSecretRef,omitempty" tf:"-"`

	// The ID of a placement group.
	// The ID of a placement group.
	// +kubebuilder:validation:Optional
	PlacementGroupID *string `json:"placementGroupId,omitempty" tf:"placement_group_id,omitempty"`

	// The private IP to be assigned to this instance, must be in the provided subnet and available.
	// The private IP to be assigned to this instance, must be in the provided subnet and available.
	// +kubebuilder:validation:Optional
	PrivateIP *string `json:"privateIp,omitempty" tf:"private_ip,omitempty"`

	// The project the instance belongs to, default to 0.
	// The project the instance belongs to, default to 0.
	// +kubebuilder:validation:Optional
	ProjectID *float64 `json:"projectId,omitempty" tf:"project_id,omitempty"`

	// A list of security group IDs to associate with.
	// A list of security group IDs to associate with.
	// +kubebuilder:validation:Optional
	SecurityGroups []*string `json:"securityGroups,omitempty" tf:"security_groups,omitempty"`

	// The ID of a VPC subnet. If you want to create instances in a VPC network, this parameter must be set.
	// The ID of a VPC subnet. If you want to create instances in a VPC network, this parameter must be set.
	// +crossplane:generate:reference:type=github.com/crossplane-contrib/provider-tencentcloud/apis/vpc/v1alpha1.Subnet
	// +kubebuilder:validation:Optional
	SubnetID *string `json:"subnetId,omitempty" tf:"subnet_id,omitempty"`

	// Reference to a Subnet in vpc to populate subnetId.
	// +kubebuilder:validation:Optional
	SubnetIDRef *v1.Reference `json:"subnetIdRef,omitempty" tf:"-"`

	// Selector for a Subnet in vpc to populate subnetId.
	// +kubebuilder:validation:Optional
	SubnetIDSelector *v1.Selector `json:"subnetIdSelector,omitempty" tf:"-"`

	// System disk snapshot ID used to initialize the system disk. When system disk type is LOCAL_BASIC and LOCAL_SSD, disk id is not supported.
	// System disk snapshot ID used to initialize the system disk. When system disk type is `LOCAL_BASIC` and `LOCAL_SSD`, disk id is not supported.
	// +kubebuilder:validation:Optional
	SystemDiskID *string `json:"systemDiskId,omitempty" tf:"system_disk_id,omitempty"`

	// Size of the system disk. Valid value ranges: (50~1000). and unit is GB. Default is 50GB. If modified, the instance may force stop.
	// Size of the system disk. Valid value ranges: (50~1000). and unit is GB. Default is 50GB. If modified, the instance may force stop.
	// +kubebuilder:validation:Optional
	SystemDiskSize *float64 `json:"systemDiskSize,omitempty" tf:"system_disk_size,omitempty"`

	// System disk type. For more information on limits of system disk types, see Storage Overview. Valid values: LOCAL_BASIC: local disk, LOCAL_SSD: local SSD disk, CLOUD_SSD: SSD, CLOUD_PREMIUM: Premium Cloud Storage, CLOUD_BSSD: Basic SSD. NOTE: If modified, the instance may force stop.
	// System disk type. For more information on limits of system disk types, see [Storage Overview](https://intl.cloud.tencent.com/document/product/213/4952). Valid values: `LOCAL_BASIC`: local disk, `LOCAL_SSD`: local SSD disk, `CLOUD_SSD`: SSD, `CLOUD_PREMIUM`: Premium Cloud Storage, `CLOUD_BSSD`: Basic SSD. NOTE: If modified, the instance may force stop.
	// +kubebuilder:validation:Optional
	SystemDiskType *string `json:"systemDiskType,omitempty" tf:"system_disk_type,omitempty"`

	// The user data to be injected into this instance. Must be base64 encoded and up to 16 KB.
	// The user data to be injected into this instance. Must be base64 encoded and up to 16 KB.
	// +kubebuilder:validation:Optional
	UserData *string `json:"userData,omitempty" tf:"user_data,omitempty"`

	// The user data to be injected into this instance, in plain text. Conflicts with user_data. Up to 16 KB after base64 encoded.
	// The user data to be injected into this instance, in plain text. Conflicts with `user_data`. Up to 16 KB after base64 encoded.
	// +kubebuilder:validation:Optional
	UserDataRaw *string `json:"userDataRaw,omitempty" tf:"user_data_raw,omitempty"`

	// The ID of a VPC network. If you want to create instances in a VPC network, this parameter must be set.
	// The ID of a VPC network. If you want to create instances in a VPC network, this parameter must be set.
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

// InstanceSetSpec defines the desired state of InstanceSet
type InstanceSetSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     InstanceSetParameters `json:"forProvider"`
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
	InitProvider InstanceSetInitParameters `json:"initProvider,omitempty"`
}

// InstanceSetStatus defines the observed state of InstanceSet.
type InstanceSetStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        InstanceSetObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// InstanceSet is the Schema for the InstanceSets API. Provides a CVM instance set resource.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,tencentcloud}
type InstanceSet struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.availabilityZone) || (has(self.initProvider) && has(self.initProvider.availabilityZone))",message="spec.forProvider.availabilityZone is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.imageId) || (has(self.initProvider) && has(self.initProvider.imageId))",message="spec.forProvider.imageId is a required parameter"
	Spec   InstanceSetSpec   `json:"spec"`
	Status InstanceSetStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// InstanceSetList contains a list of InstanceSets
type InstanceSetList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []InstanceSet `json:"items"`
}

// Repository type metadata.
var (
	InstanceSet_Kind             = "InstanceSet"
	InstanceSet_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: InstanceSet_Kind}.String()
	InstanceSet_KindAPIVersion   = InstanceSet_Kind + "." + CRDGroupVersion.String()
	InstanceSet_GroupVersionKind = CRDGroupVersion.WithKind(InstanceSet_Kind)
)

func init() {
	SchemeBuilder.Register(&InstanceSet{}, &InstanceSetList{})
}
