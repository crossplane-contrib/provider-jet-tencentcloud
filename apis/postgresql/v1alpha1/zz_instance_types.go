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

type BackupPlanInitParameters struct {

	// List of backup period per week, available values: monday, tuesday, wednesday, thursday, friday, saturday, sunday. NOTE: At least specify two days.
	// List of backup period per week, available values: `monday`, `tuesday`, `wednesday`, `thursday`, `friday`, `saturday`, `sunday`. NOTE: At least specify two days.
	BackupPeriod []*string `json:"backupPeriod,omitempty" tf:"backup_period,omitempty"`

	// Specify days of the retention.
	// Specify days of the retention.
	BaseBackupRetentionPeriod *float64 `json:"baseBackupRetentionPeriod,omitempty" tf:"base_backup_retention_period,omitempty"`

	// Specify latest backup start time, format hh:mm:ss.
	// Specify latest backup start time, format `hh:mm:ss`.
	MaxBackupStartTime *string `json:"maxBackupStartTime,omitempty" tf:"max_backup_start_time,omitempty"`

	// Specify earliest backup start time, format hh:mm:ss.
	// Specify earliest backup start time, format `hh:mm:ss`.
	MinBackupStartTime *string `json:"minBackupStartTime,omitempty" tf:"min_backup_start_time,omitempty"`
}

type BackupPlanObservation struct {

	// List of backup period per week, available values: monday, tuesday, wednesday, thursday, friday, saturday, sunday. NOTE: At least specify two days.
	// List of backup period per week, available values: `monday`, `tuesday`, `wednesday`, `thursday`, `friday`, `saturday`, `sunday`. NOTE: At least specify two days.
	BackupPeriod []*string `json:"backupPeriod,omitempty" tf:"backup_period,omitempty"`

	// Specify days of the retention.
	// Specify days of the retention.
	BaseBackupRetentionPeriod *float64 `json:"baseBackupRetentionPeriod,omitempty" tf:"base_backup_retention_period,omitempty"`

	// Specify latest backup start time, format hh:mm:ss.
	// Specify latest backup start time, format `hh:mm:ss`.
	MaxBackupStartTime *string `json:"maxBackupStartTime,omitempty" tf:"max_backup_start_time,omitempty"`

	// Specify earliest backup start time, format hh:mm:ss.
	// Specify earliest backup start time, format `hh:mm:ss`.
	MinBackupStartTime *string `json:"minBackupStartTime,omitempty" tf:"min_backup_start_time,omitempty"`
}

type BackupPlanParameters struct {

	// List of backup period per week, available values: monday, tuesday, wednesday, thursday, friday, saturday, sunday. NOTE: At least specify two days.
	// List of backup period per week, available values: `monday`, `tuesday`, `wednesday`, `thursday`, `friday`, `saturday`, `sunday`. NOTE: At least specify two days.
	// +kubebuilder:validation:Optional
	BackupPeriod []*string `json:"backupPeriod,omitempty" tf:"backup_period,omitempty"`

	// Specify days of the retention.
	// Specify days of the retention.
	// +kubebuilder:validation:Optional
	BaseBackupRetentionPeriod *float64 `json:"baseBackupRetentionPeriod,omitempty" tf:"base_backup_retention_period,omitempty"`

	// Specify latest backup start time, format hh:mm:ss.
	// Specify latest backup start time, format `hh:mm:ss`.
	// +kubebuilder:validation:Optional
	MaxBackupStartTime *string `json:"maxBackupStartTime,omitempty" tf:"max_backup_start_time,omitempty"`

	// Specify earliest backup start time, format hh:mm:ss.
	// Specify earliest backup start time, format `hh:mm:ss`.
	// +kubebuilder:validation:Optional
	MinBackupStartTime *string `json:"minBackupStartTime,omitempty" tf:"min_backup_start_time,omitempty"`
}

type DBNodeSetInitParameters struct {

	// Indicates node type, available values:Primary, Standby. Default: Standby.
	// Indicates node type, available values:`Primary`, `Standby`. Default: `Standby`.
	Role *string `json:"role,omitempty" tf:"role,omitempty"`

	// Indicates the node available zone.
	// Indicates the node available zone.
	Zone *string `json:"zone,omitempty" tf:"zone,omitempty"`
}

type DBNodeSetObservation struct {

	// Indicates node type, available values:Primary, Standby. Default: Standby.
	// Indicates node type, available values:`Primary`, `Standby`. Default: `Standby`.
	Role *string `json:"role,omitempty" tf:"role,omitempty"`

	// Indicates the node available zone.
	// Indicates the node available zone.
	Zone *string `json:"zone,omitempty" tf:"zone,omitempty"`
}

type DBNodeSetParameters struct {

	// Indicates node type, available values:Primary, Standby. Default: Standby.
	// Indicates node type, available values:`Primary`, `Standby`. Default: `Standby`.
	// +kubebuilder:validation:Optional
	Role *string `json:"role,omitempty" tf:"role,omitempty"`

	// Indicates the node available zone.
	// Indicates the node available zone.
	// +kubebuilder:validation:Optional
	Zone *string `json:"zone" tf:"zone,omitempty"`
}

type InstanceInitParameters struct {

	// Auto renew flag, 1 for enabled. NOTES: Only support prepaid instance.
	// Auto renew flag, `1` for enabled. NOTES: Only support prepaid instance.
	AutoRenewFlag *float64 `json:"autoRenewFlag,omitempty" tf:"auto_renew_flag,omitempty"`

	// Whether to use voucher, 1 for enabled.
	// Whether to use voucher, `1` for enabled.
	AutoVoucher *float64 `json:"autoVoucher,omitempty" tf:"auto_voucher,omitempty"`

	// Availability zone. NOTE: This field could not be modified, please use db_node_set instead of modification. The changes on this field will be suppressed when using the db_node_set.
	// Availability zone. NOTE: This field could not be modified, please use `db_node_set` instead of modification. The changes on this field will be suppressed when using the `db_node_set`.
	AvailabilityZone *string `json:"availabilityZone,omitempty" tf:"availability_zone,omitempty"`

	// Specify DB backup plan.
	// Specify DB backup plan.
	BackupPlan []BackupPlanInitParameters `json:"backupPlan,omitempty" tf:"backup_plan,omitempty"`

	// Number of CPU cores. Allowed value must be equal cpu that data source tencentcloud_postgresql_specinfos provides.
	// Number of CPU cores. Allowed value must be equal `cpu` that data source `tencentcloud_postgresql_specinfos` provides.
	CPU *float64 `json:"cpu,omitempty" tf:"cpu,omitempty"`

	// Pay type of the postgresql instance. Values POSTPAID_BY_HOUR (Default), PREPAID. It only support to update the type from POSTPAID_BY_HOUR to PREPAID.
	// Pay type of the postgresql instance. Values `POSTPAID_BY_HOUR` (Default), `PREPAID`. It only support to update the type from `POSTPAID_BY_HOUR` to `PREPAID`.
	ChargeType *string `json:"chargeType,omitempty" tf:"charge_type,omitempty"`

	// Charset of the root account. Valid values are UTF8,LATIN1.
	// Charset of the root account. Valid values are `UTF8`,`LATIN1`.
	Charset *string `json:"charset,omitempty" tf:"charset,omitempty"`

	// PostgreSQL kernel version number. If it is specified, an instance running kernel DBKernelVersion will be created. It supports updating the minor kernel version immediately.
	// PostgreSQL kernel version number. If it is specified, an instance running kernel DBKernelVersion will be created. It supports updating the minor kernel version immediately.
	DBKernelVersion *string `json:"dbKernelVersion,omitempty" tf:"db_kernel_version,omitempty"`

	// PostgreSQL major version number. Valid values: 10, 11, 12, 13. If it is specified, an instance running the latest kernel of PostgreSQL DBMajorVersion will be created.
	// PostgreSQL major version number. Valid values: 10, 11, 12, 13. If it is specified, an instance running the latest kernel of PostgreSQL DBMajorVersion will be created.
	DBMajorVersion *string `json:"dbMajorVersion,omitempty" tf:"db_major_version,omitempty"`

	// db_major_vesion will be deprecated, use db_major_version instead. PostgreSQL major version number. Valid values: 10, 11, 12, 13. If it is specified, an instance running the latest kernel of PostgreSQL DBMajorVersion will be created.
	// PostgreSQL major version number. Valid values: 10, 11, 12, 13. If it is specified, an instance running the latest kernel of PostgreSQL DBMajorVersion will be created.
	DBMajorVesion *string `json:"dbMajorVesion,omitempty" tf:"db_major_vesion,omitempty"`

	// Specify instance node info for disaster migration.
	// Specify instance node info for disaster migration.
	DBNodeSet []DBNodeSetInitParameters `json:"dbNodeSet,omitempty" tf:"db_node_set,omitempty"`

	// Version of the postgresql database engine. Valid values: 10.4, 11.8, 12.4.
	// Version of the postgresql database engine. Valid values: `10.4`, `11.8`, `12.4`.
	EngineVersion *string `json:"engineVersion,omitempty" tf:"engine_version,omitempty"`

	// KeyId of the custom key.
	// KeyId of the custom key.
	KMSKeyID *string `json:"kmsKeyId,omitempty" tf:"kms_key_id,omitempty"`

	// Region of the custom key.
	// Region of the custom key.
	KMSRegion *string `json:"kmsRegion,omitempty" tf:"kms_region,omitempty"`

	// max_standby_archive_delay applies when WAL data is being read from WAL archive (and is therefore not current). Units are milliseconds if not specified.
	// max_standby_archive_delay applies when WAL data is being read from WAL archive (and is therefore not current). Units are milliseconds if not specified.
	MaxStandbyArchiveDelay *float64 `json:"maxStandbyArchiveDelay,omitempty" tf:"max_standby_archive_delay,omitempty"`

	// max_standby_streaming_delay applies when WAL data is being received via streaming replication. Units are milliseconds if not specified.
	// max_standby_streaming_delay applies when WAL data is being received via streaming replication. Units are milliseconds if not specified.
	MaxStandbyStreamingDelay *float64 `json:"maxStandbyStreamingDelay,omitempty" tf:"max_standby_streaming_delay,omitempty"`

	// Memory size(in GB). Allowed value must be larger than memory that data source tencentcloud_postgresql_specinfos provides.
	// Memory size(in GB). Allowed value must be larger than `memory` that data source `tencentcloud_postgresql_specinfos` provides.
	Memory *float64 `json:"memory,omitempty" tf:"memory,omitempty"`

	// Name of the postgresql instance.
	// Name of the postgresql instance.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Whether to support data transparent encryption, 1: yes, 0: no (default).
	// Whether to support data transparent encryption, 1: yes, 0: no (default).
	NeedSupportTde *float64 `json:"needSupportTde,omitempty" tf:"need_support_tde,omitempty"`

	// Specify Prepaid period in month. Default 1. Values: 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 24, 36. This field is valid only when creating a PREPAID type instance, or updating the charge type from POSTPAID_BY_HOUR to PREPAID.
	// Specify Prepaid period in month. Default `1`. Values: `1`, `2`, `3`, `4`, `5`, `6`, `7`, `8`, `9`, `10`, `11`, `12`, `24`, `36`. This field is valid only when creating a `PREPAID` type instance, or updating the charge type from `POSTPAID_BY_HOUR` to `PREPAID`.
	Period *float64 `json:"period,omitempty" tf:"period,omitempty"`

	// Project id, default value is 0.
	// Project id, default value is `0`.
	ProjectID *float64 `json:"projectId,omitempty" tf:"project_id,omitempty"`

	// Indicates whether to enable the access to an instance from public network or not.
	// Indicates whether to enable the access to an instance from public network or not.
	PublicAccessSwitch *bool `json:"publicAccessSwitch,omitempty" tf:"public_access_switch,omitempty"`

	// Instance root account name. This parameter is optional, Default value is root.
	// Instance root account name. This parameter is optional, Default value is `root`.
	RootUser *string `json:"rootUser,omitempty" tf:"root_user,omitempty"`

	// ID of security group. If both vpc_id and subnet_id are not set, this argument should not be set either.
	// ID of security group. If both vpc_id and subnet_id are not set, this argument should not be set either.
	// +listType=set
	SecurityGroups []*string `json:"securityGroups,omitempty" tf:"security_groups,omitempty"`

	// Volume size(in GB). Allowed value must be a multiple of 10. The storage must be set with the limit of storage_min and storage_max which data source tencentcloud_postgresql_specinfos provides.
	// Volume size(in GB). Allowed value must be a multiple of 10. The storage must be set with the limit of `storage_min` and `storage_max` which data source `tencentcloud_postgresql_specinfos` provides.
	Storage *float64 `json:"storage,omitempty" tf:"storage,omitempty"`

	// ID of subnet.
	// ID of subnet.
	// +crossplane:generate:reference:type=github.com/crossplane-contrib/provider-tencentcloud/apis/vpc/v1alpha1.Subnet
	SubnetID *string `json:"subnetId,omitempty" tf:"subnet_id,omitempty"`

	// Reference to a Subnet in vpc to populate subnetId.
	// +kubebuilder:validation:Optional
	SubnetIDRef *v1.Reference `json:"subnetIdRef,omitempty" tf:"-"`

	// Selector for a Subnet in vpc to populate subnetId.
	// +kubebuilder:validation:Optional
	SubnetIDSelector *v1.Selector `json:"subnetIdSelector,omitempty" tf:"-"`

	// The available tags within this postgresql.
	// The available tags within this postgresql.
	// +mapType=granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// ID of VPC.
	// ID of VPC.
	// +crossplane:generate:reference:type=github.com/crossplane-contrib/provider-tencentcloud/apis/vpc/v1alpha1.VPC
	VPCID *string `json:"vpcId,omitempty" tf:"vpc_id,omitempty"`

	// Reference to a VPC in vpc to populate vpcId.
	// +kubebuilder:validation:Optional
	VPCIDRef *v1.Reference `json:"vpcIdRef,omitempty" tf:"-"`

	// Selector for a VPC in vpc to populate vpcId.
	// +kubebuilder:validation:Optional
	VPCIDSelector *v1.Selector `json:"vpcIdSelector,omitempty" tf:"-"`

	// Specify Voucher Ids if auto_voucher was 1, only support using 1 vouchers for now.
	// Specify Voucher Ids if `auto_voucher` was `1`, only support using 1 vouchers for now.
	VoucherIds []*string `json:"voucherIds,omitempty" tf:"voucher_ids,omitempty"`
}

type InstanceObservation struct {

	// Auto renew flag, 1 for enabled. NOTES: Only support prepaid instance.
	// Auto renew flag, `1` for enabled. NOTES: Only support prepaid instance.
	AutoRenewFlag *float64 `json:"autoRenewFlag,omitempty" tf:"auto_renew_flag,omitempty"`

	// Whether to use voucher, 1 for enabled.
	// Whether to use voucher, `1` for enabled.
	AutoVoucher *float64 `json:"autoVoucher,omitempty" tf:"auto_voucher,omitempty"`

	// Availability zone. NOTE: This field could not be modified, please use db_node_set instead of modification. The changes on this field will be suppressed when using the db_node_set.
	// Availability zone. NOTE: This field could not be modified, please use `db_node_set` instead of modification. The changes on this field will be suppressed when using the `db_node_set`.
	AvailabilityZone *string `json:"availabilityZone,omitempty" tf:"availability_zone,omitempty"`

	// Specify DB backup plan.
	// Specify DB backup plan.
	BackupPlan []BackupPlanObservation `json:"backupPlan,omitempty" tf:"backup_plan,omitempty"`

	// Number of CPU cores. Allowed value must be equal cpu that data source tencentcloud_postgresql_specinfos provides.
	// Number of CPU cores. Allowed value must be equal `cpu` that data source `tencentcloud_postgresql_specinfos` provides.
	CPU *float64 `json:"cpu,omitempty" tf:"cpu,omitempty"`

	// Pay type of the postgresql instance. Values POSTPAID_BY_HOUR (Default), PREPAID. It only support to update the type from POSTPAID_BY_HOUR to PREPAID.
	// Pay type of the postgresql instance. Values `POSTPAID_BY_HOUR` (Default), `PREPAID`. It only support to update the type from `POSTPAID_BY_HOUR` to `PREPAID`.
	ChargeType *string `json:"chargeType,omitempty" tf:"charge_type,omitempty"`

	// Charset of the root account. Valid values are UTF8,LATIN1.
	// Charset of the root account. Valid values are `UTF8`,`LATIN1`.
	Charset *string `json:"charset,omitempty" tf:"charset,omitempty"`

	// Create time of the postgresql instance.
	// Create time of the postgresql instance.
	CreateTime *string `json:"createTime,omitempty" tf:"create_time,omitempty"`

	// PostgreSQL kernel version number. If it is specified, an instance running kernel DBKernelVersion will be created. It supports updating the minor kernel version immediately.
	// PostgreSQL kernel version number. If it is specified, an instance running kernel DBKernelVersion will be created. It supports updating the minor kernel version immediately.
	DBKernelVersion *string `json:"dbKernelVersion,omitempty" tf:"db_kernel_version,omitempty"`

	// PostgreSQL major version number. Valid values: 10, 11, 12, 13. If it is specified, an instance running the latest kernel of PostgreSQL DBMajorVersion will be created.
	// PostgreSQL major version number. Valid values: 10, 11, 12, 13. If it is specified, an instance running the latest kernel of PostgreSQL DBMajorVersion will be created.
	DBMajorVersion *string `json:"dbMajorVersion,omitempty" tf:"db_major_version,omitempty"`

	// db_major_vesion will be deprecated, use db_major_version instead. PostgreSQL major version number. Valid values: 10, 11, 12, 13. If it is specified, an instance running the latest kernel of PostgreSQL DBMajorVersion will be created.
	// PostgreSQL major version number. Valid values: 10, 11, 12, 13. If it is specified, an instance running the latest kernel of PostgreSQL DBMajorVersion will be created.
	DBMajorVesion *string `json:"dbMajorVesion,omitempty" tf:"db_major_vesion,omitempty"`

	// Specify instance node info for disaster migration.
	// Specify instance node info for disaster migration.
	DBNodeSet []DBNodeSetObservation `json:"dbNodeSet,omitempty" tf:"db_node_set,omitempty"`

	// Version of the postgresql database engine. Valid values: 10.4, 11.8, 12.4.
	// Version of the postgresql database engine. Valid values: `10.4`, `11.8`, `12.4`.
	EngineVersion *string `json:"engineVersion,omitempty" tf:"engine_version,omitempty"`

	// ID of the resource.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// KeyId of the custom key.
	// KeyId of the custom key.
	KMSKeyID *string `json:"kmsKeyId,omitempty" tf:"kms_key_id,omitempty"`

	// Region of the custom key.
	// Region of the custom key.
	KMSRegion *string `json:"kmsRegion,omitempty" tf:"kms_region,omitempty"`

	// max_standby_archive_delay applies when WAL data is being read from WAL archive (and is therefore not current). Units are milliseconds if not specified.
	// max_standby_archive_delay applies when WAL data is being read from WAL archive (and is therefore not current). Units are milliseconds if not specified.
	MaxStandbyArchiveDelay *float64 `json:"maxStandbyArchiveDelay,omitempty" tf:"max_standby_archive_delay,omitempty"`

	// max_standby_streaming_delay applies when WAL data is being received via streaming replication. Units are milliseconds if not specified.
	// max_standby_streaming_delay applies when WAL data is being received via streaming replication. Units are milliseconds if not specified.
	MaxStandbyStreamingDelay *float64 `json:"maxStandbyStreamingDelay,omitempty" tf:"max_standby_streaming_delay,omitempty"`

	// Memory size(in GB). Allowed value must be larger than memory that data source tencentcloud_postgresql_specinfos provides.
	// Memory size(in GB). Allowed value must be larger than `memory` that data source `tencentcloud_postgresql_specinfos` provides.
	Memory *float64 `json:"memory,omitempty" tf:"memory,omitempty"`

	// Name of the postgresql instance.
	// Name of the postgresql instance.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Whether to support data transparent encryption, 1: yes, 0: no (default).
	// Whether to support data transparent encryption, 1: yes, 0: no (default).
	NeedSupportTde *float64 `json:"needSupportTde,omitempty" tf:"need_support_tde,omitempty"`

	// Specify Prepaid period in month. Default 1. Values: 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 24, 36. This field is valid only when creating a PREPAID type instance, or updating the charge type from POSTPAID_BY_HOUR to PREPAID.
	// Specify Prepaid period in month. Default `1`. Values: `1`, `2`, `3`, `4`, `5`, `6`, `7`, `8`, `9`, `10`, `11`, `12`, `24`, `36`. This field is valid only when creating a `PREPAID` type instance, or updating the charge type from `POSTPAID_BY_HOUR` to `PREPAID`.
	Period *float64 `json:"period,omitempty" tf:"period,omitempty"`

	// IP for private access.
	// IP for private access.
	PrivateAccessIP *string `json:"privateAccessIp,omitempty" tf:"private_access_ip,omitempty"`

	// Port for private access.
	// Port for private access.
	PrivateAccessPort *float64 `json:"privateAccessPort,omitempty" tf:"private_access_port,omitempty"`

	// Project id, default value is 0.
	// Project id, default value is `0`.
	ProjectID *float64 `json:"projectId,omitempty" tf:"project_id,omitempty"`

	// Host for public access.
	// Host for public access.
	PublicAccessHost *string `json:"publicAccessHost,omitempty" tf:"public_access_host,omitempty"`

	// Port for public access.
	// Port for public access.
	PublicAccessPort *float64 `json:"publicAccessPort,omitempty" tf:"public_access_port,omitempty"`

	// Indicates whether to enable the access to an instance from public network or not.
	// Indicates whether to enable the access to an instance from public network or not.
	PublicAccessSwitch *bool `json:"publicAccessSwitch,omitempty" tf:"public_access_switch,omitempty"`

	// Instance root account name. This parameter is optional, Default value is root.
	// Instance root account name. This parameter is optional, Default value is `root`.
	RootUser *string `json:"rootUser,omitempty" tf:"root_user,omitempty"`

	// ID of security group. If both vpc_id and subnet_id are not set, this argument should not be set either.
	// ID of security group. If both vpc_id and subnet_id are not set, this argument should not be set either.
	// +listType=set
	SecurityGroups []*string `json:"securityGroups,omitempty" tf:"security_groups,omitempty"`

	// Volume size(in GB). Allowed value must be a multiple of 10. The storage must be set with the limit of storage_min and storage_max which data source tencentcloud_postgresql_specinfos provides.
	// Volume size(in GB). Allowed value must be a multiple of 10. The storage must be set with the limit of `storage_min` and `storage_max` which data source `tencentcloud_postgresql_specinfos` provides.
	Storage *float64 `json:"storage,omitempty" tf:"storage,omitempty"`

	// ID of subnet.
	// ID of subnet.
	SubnetID *string `json:"subnetId,omitempty" tf:"subnet_id,omitempty"`

	// The available tags within this postgresql.
	// The available tags within this postgresql.
	// +mapType=granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// Uid of the postgresql instance.
	// Uid of the postgresql instance.
	UID *float64 `json:"uid,omitempty" tf:"uid,omitempty"`

	// ID of VPC.
	// ID of VPC.
	VPCID *string `json:"vpcId,omitempty" tf:"vpc_id,omitempty"`

	// Specify Voucher Ids if auto_voucher was 1, only support using 1 vouchers for now.
	// Specify Voucher Ids if `auto_voucher` was `1`, only support using 1 vouchers for now.
	VoucherIds []*string `json:"voucherIds,omitempty" tf:"voucher_ids,omitempty"`
}

type InstanceParameters struct {

	// Auto renew flag, 1 for enabled. NOTES: Only support prepaid instance.
	// Auto renew flag, `1` for enabled. NOTES: Only support prepaid instance.
	// +kubebuilder:validation:Optional
	AutoRenewFlag *float64 `json:"autoRenewFlag,omitempty" tf:"auto_renew_flag,omitempty"`

	// Whether to use voucher, 1 for enabled.
	// Whether to use voucher, `1` for enabled.
	// +kubebuilder:validation:Optional
	AutoVoucher *float64 `json:"autoVoucher,omitempty" tf:"auto_voucher,omitempty"`

	// Availability zone. NOTE: This field could not be modified, please use db_node_set instead of modification. The changes on this field will be suppressed when using the db_node_set.
	// Availability zone. NOTE: This field could not be modified, please use `db_node_set` instead of modification. The changes on this field will be suppressed when using the `db_node_set`.
	// +kubebuilder:validation:Optional
	AvailabilityZone *string `json:"availabilityZone,omitempty" tf:"availability_zone,omitempty"`

	// Specify DB backup plan.
	// Specify DB backup plan.
	// +kubebuilder:validation:Optional
	BackupPlan []BackupPlanParameters `json:"backupPlan,omitempty" tf:"backup_plan,omitempty"`

	// Number of CPU cores. Allowed value must be equal cpu that data source tencentcloud_postgresql_specinfos provides.
	// Number of CPU cores. Allowed value must be equal `cpu` that data source `tencentcloud_postgresql_specinfos` provides.
	// +kubebuilder:validation:Optional
	CPU *float64 `json:"cpu,omitempty" tf:"cpu,omitempty"`

	// Pay type of the postgresql instance. Values POSTPAID_BY_HOUR (Default), PREPAID. It only support to update the type from POSTPAID_BY_HOUR to PREPAID.
	// Pay type of the postgresql instance. Values `POSTPAID_BY_HOUR` (Default), `PREPAID`. It only support to update the type from `POSTPAID_BY_HOUR` to `PREPAID`.
	// +kubebuilder:validation:Optional
	ChargeType *string `json:"chargeType,omitempty" tf:"charge_type,omitempty"`

	// Charset of the root account. Valid values are UTF8,LATIN1.
	// Charset of the root account. Valid values are `UTF8`,`LATIN1`.
	// +kubebuilder:validation:Optional
	Charset *string `json:"charset,omitempty" tf:"charset,omitempty"`

	// PostgreSQL kernel version number. If it is specified, an instance running kernel DBKernelVersion will be created. It supports updating the minor kernel version immediately.
	// PostgreSQL kernel version number. If it is specified, an instance running kernel DBKernelVersion will be created. It supports updating the minor kernel version immediately.
	// +kubebuilder:validation:Optional
	DBKernelVersion *string `json:"dbKernelVersion,omitempty" tf:"db_kernel_version,omitempty"`

	// PostgreSQL major version number. Valid values: 10, 11, 12, 13. If it is specified, an instance running the latest kernel of PostgreSQL DBMajorVersion will be created.
	// PostgreSQL major version number. Valid values: 10, 11, 12, 13. If it is specified, an instance running the latest kernel of PostgreSQL DBMajorVersion will be created.
	// +kubebuilder:validation:Optional
	DBMajorVersion *string `json:"dbMajorVersion,omitempty" tf:"db_major_version,omitempty"`

	// db_major_vesion will be deprecated, use db_major_version instead. PostgreSQL major version number. Valid values: 10, 11, 12, 13. If it is specified, an instance running the latest kernel of PostgreSQL DBMajorVersion will be created.
	// PostgreSQL major version number. Valid values: 10, 11, 12, 13. If it is specified, an instance running the latest kernel of PostgreSQL DBMajorVersion will be created.
	// +kubebuilder:validation:Optional
	DBMajorVesion *string `json:"dbMajorVesion,omitempty" tf:"db_major_vesion,omitempty"`

	// Specify instance node info for disaster migration.
	// Specify instance node info for disaster migration.
	// +kubebuilder:validation:Optional
	DBNodeSet []DBNodeSetParameters `json:"dbNodeSet,omitempty" tf:"db_node_set,omitempty"`

	// Version of the postgresql database engine. Valid values: 10.4, 11.8, 12.4.
	// Version of the postgresql database engine. Valid values: `10.4`, `11.8`, `12.4`.
	// +kubebuilder:validation:Optional
	EngineVersion *string `json:"engineVersion,omitempty" tf:"engine_version,omitempty"`

	// KeyId of the custom key.
	// KeyId of the custom key.
	// +kubebuilder:validation:Optional
	KMSKeyID *string `json:"kmsKeyId,omitempty" tf:"kms_key_id,omitempty"`

	// Region of the custom key.
	// Region of the custom key.
	// +kubebuilder:validation:Optional
	KMSRegion *string `json:"kmsRegion,omitempty" tf:"kms_region,omitempty"`

	// max_standby_archive_delay applies when WAL data is being read from WAL archive (and is therefore not current). Units are milliseconds if not specified.
	// max_standby_archive_delay applies when WAL data is being read from WAL archive (and is therefore not current). Units are milliseconds if not specified.
	// +kubebuilder:validation:Optional
	MaxStandbyArchiveDelay *float64 `json:"maxStandbyArchiveDelay,omitempty" tf:"max_standby_archive_delay,omitempty"`

	// max_standby_streaming_delay applies when WAL data is being received via streaming replication. Units are milliseconds if not specified.
	// max_standby_streaming_delay applies when WAL data is being received via streaming replication. Units are milliseconds if not specified.
	// +kubebuilder:validation:Optional
	MaxStandbyStreamingDelay *float64 `json:"maxStandbyStreamingDelay,omitempty" tf:"max_standby_streaming_delay,omitempty"`

	// Memory size(in GB). Allowed value must be larger than memory that data source tencentcloud_postgresql_specinfos provides.
	// Memory size(in GB). Allowed value must be larger than `memory` that data source `tencentcloud_postgresql_specinfos` provides.
	// +kubebuilder:validation:Optional
	Memory *float64 `json:"memory,omitempty" tf:"memory,omitempty"`

	// Name of the postgresql instance.
	// Name of the postgresql instance.
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Whether to support data transparent encryption, 1: yes, 0: no (default).
	// Whether to support data transparent encryption, 1: yes, 0: no (default).
	// +kubebuilder:validation:Optional
	NeedSupportTde *float64 `json:"needSupportTde,omitempty" tf:"need_support_tde,omitempty"`

	// Specify Prepaid period in month. Default 1. Values: 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 24, 36. This field is valid only when creating a PREPAID type instance, or updating the charge type from POSTPAID_BY_HOUR to PREPAID.
	// Specify Prepaid period in month. Default `1`. Values: `1`, `2`, `3`, `4`, `5`, `6`, `7`, `8`, `9`, `10`, `11`, `12`, `24`, `36`. This field is valid only when creating a `PREPAID` type instance, or updating the charge type from `POSTPAID_BY_HOUR` to `PREPAID`.
	// +kubebuilder:validation:Optional
	Period *float64 `json:"period,omitempty" tf:"period,omitempty"`

	// Project id, default value is 0.
	// Project id, default value is `0`.
	// +kubebuilder:validation:Optional
	ProjectID *float64 `json:"projectId,omitempty" tf:"project_id,omitempty"`

	// Indicates whether to enable the access to an instance from public network or not.
	// Indicates whether to enable the access to an instance from public network or not.
	// +kubebuilder:validation:Optional
	PublicAccessSwitch *bool `json:"publicAccessSwitch,omitempty" tf:"public_access_switch,omitempty"`

	// Password of root account. This parameter can be specified when you purchase master instances, but it should be ignored when you purchase read-only instances or disaster recovery instances.
	// Password of root account. This parameter can be specified when you purchase master instances, but it should be ignored when you purchase read-only instances or disaster recovery instances.
	// +kubebuilder:validation:Optional
	RootPasswordSecretRef v1.SecretKeySelector `json:"rootPasswordSecretRef" tf:"-"`

	// Instance root account name. This parameter is optional, Default value is root.
	// Instance root account name. This parameter is optional, Default value is `root`.
	// +kubebuilder:validation:Optional
	RootUser *string `json:"rootUser,omitempty" tf:"root_user,omitempty"`

	// ID of security group. If both vpc_id and subnet_id are not set, this argument should not be set either.
	// ID of security group. If both vpc_id and subnet_id are not set, this argument should not be set either.
	// +kubebuilder:validation:Optional
	// +listType=set
	SecurityGroups []*string `json:"securityGroups,omitempty" tf:"security_groups,omitempty"`

	// Volume size(in GB). Allowed value must be a multiple of 10. The storage must be set with the limit of storage_min and storage_max which data source tencentcloud_postgresql_specinfos provides.
	// Volume size(in GB). Allowed value must be a multiple of 10. The storage must be set with the limit of `storage_min` and `storage_max` which data source `tencentcloud_postgresql_specinfos` provides.
	// +kubebuilder:validation:Optional
	Storage *float64 `json:"storage,omitempty" tf:"storage,omitempty"`

	// ID of subnet.
	// ID of subnet.
	// +crossplane:generate:reference:type=github.com/crossplane-contrib/provider-tencentcloud/apis/vpc/v1alpha1.Subnet
	// +kubebuilder:validation:Optional
	SubnetID *string `json:"subnetId,omitempty" tf:"subnet_id,omitempty"`

	// Reference to a Subnet in vpc to populate subnetId.
	// +kubebuilder:validation:Optional
	SubnetIDRef *v1.Reference `json:"subnetIdRef,omitempty" tf:"-"`

	// Selector for a Subnet in vpc to populate subnetId.
	// +kubebuilder:validation:Optional
	SubnetIDSelector *v1.Selector `json:"subnetIdSelector,omitempty" tf:"-"`

	// The available tags within this postgresql.
	// The available tags within this postgresql.
	// +kubebuilder:validation:Optional
	// +mapType=granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// ID of VPC.
	// ID of VPC.
	// +crossplane:generate:reference:type=github.com/crossplane-contrib/provider-tencentcloud/apis/vpc/v1alpha1.VPC
	// +kubebuilder:validation:Optional
	VPCID *string `json:"vpcId,omitempty" tf:"vpc_id,omitempty"`

	// Reference to a VPC in vpc to populate vpcId.
	// +kubebuilder:validation:Optional
	VPCIDRef *v1.Reference `json:"vpcIdRef,omitempty" tf:"-"`

	// Selector for a VPC in vpc to populate vpcId.
	// +kubebuilder:validation:Optional
	VPCIDSelector *v1.Selector `json:"vpcIdSelector,omitempty" tf:"-"`

	// Specify Voucher Ids if auto_voucher was 1, only support using 1 vouchers for now.
	// Specify Voucher Ids if `auto_voucher` was `1`, only support using 1 vouchers for now.
	// +kubebuilder:validation:Optional
	VoucherIds []*string `json:"voucherIds,omitempty" tf:"voucher_ids,omitempty"`
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

// Instance is the Schema for the Instances API. Use this resource to create postgresql instance.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,tencentcloud}
type Instance struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.availabilityZone) || (has(self.initProvider) && has(self.initProvider.availabilityZone))",message="spec.forProvider.availabilityZone is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.memory) || (has(self.initProvider) && has(self.initProvider.memory))",message="spec.forProvider.memory is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.name) || (has(self.initProvider) && has(self.initProvider.name))",message="spec.forProvider.name is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.rootPasswordSecretRef)",message="spec.forProvider.rootPasswordSecretRef is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.storage) || (has(self.initProvider) && has(self.initProvider.storage))",message="spec.forProvider.storage is a required parameter"
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
