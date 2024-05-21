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

type EsACLInitParameters struct {

	// Blacklist of kibana access.
	// Blacklist of kibana access.
	// +listType=set
	BlackList []*string `json:"blackList,omitempty" tf:"black_list,omitempty"`

	// Whitelist of kibana access.
	// Whitelist of kibana access.
	// +listType=set
	WhiteList []*string `json:"whiteList,omitempty" tf:"white_list,omitempty"`
}

type EsACLObservation struct {

	// Blacklist of kibana access.
	// Blacklist of kibana access.
	// +listType=set
	BlackList []*string `json:"blackList,omitempty" tf:"black_list,omitempty"`

	// Whitelist of kibana access.
	// Whitelist of kibana access.
	// +listType=set
	WhiteList []*string `json:"whiteList,omitempty" tf:"white_list,omitempty"`
}

type EsACLParameters struct {

	// Blacklist of kibana access.
	// Blacklist of kibana access.
	// +kubebuilder:validation:Optional
	// +listType=set
	BlackList []*string `json:"blackList,omitempty" tf:"black_list,omitempty"`

	// Whitelist of kibana access.
	// Whitelist of kibana access.
	// +kubebuilder:validation:Optional
	// +listType=set
	WhiteList []*string `json:"whiteList,omitempty" tf:"white_list,omitempty"`
}

type InstanceInitParameters struct {

	// Availability zone. When create multi-az es, this parameter must be omitted or -.
	// Availability zone. When create multi-az es, this parameter must be omitted or `-`.
	AvailabilityZone *string `json:"availabilityZone,omitempty" tf:"availability_zone,omitempty"`

	// Whether to enable X-Pack security authentication in Basic Edition 6.8 and above. Valid values are 1 and 2. 1 is disabled, 2 is enabled, and default value is 1. Notice: this parameter is only take effect on basic license.
	// Whether to enable X-Pack security authentication in Basic Edition 6.8 and above. Valid values are `1` and `2`. `1` is disabled, `2` is enabled, and default value is `1`. Notice: this parameter is only take effect on `basic` license.
	BasicSecurityType *float64 `json:"basicSecurityType,omitempty" tf:"basic_security_type,omitempty"`

	// The tenancy of the prepaid instance, and uint is month. NOTE: it only works when charge_type is set to PREPAID.
	// The tenancy of the prepaid instance, and uint is month. NOTE: it only works when charge_type is set to `PREPAID`.
	ChargePeriod *float64 `json:"chargePeriod,omitempty" tf:"charge_period,omitempty"`

	// The charge type of instance. Valid values are PREPAID and POSTPAID_BY_HOUR.
	// The charge type of instance. Valid values are `PREPAID` and `POSTPAID_BY_HOUR`.
	ChargeType *string `json:"chargeType,omitempty" tf:"charge_type,omitempty"`

	// Cluster deployment mode. Valid values are 0 and 1. 0 is single-AZ deployment, and 1 is multi-AZ deployment. Default value is 0.
	// Cluster deployment mode. Valid values are `0` and `1`. `0` is single-AZ deployment, and `1` is multi-AZ deployment. Default value is `0`.
	DeployMode *float64 `json:"deployMode,omitempty" tf:"deploy_mode,omitempty"`

	// Kibana Access Control Configuration.
	// Kibana Access Control Configuration.
	EsACL []EsACLInitParameters `json:"esAcl,omitempty" tf:"es_acl,omitempty"`

	// Name of the instance, which can contain 1 to 50 English letters, Chinese characters, digits, dashes(-), or underscores(_).
	// Name of the instance, which can contain 1 to 50 English letters, Chinese characters, digits, dashes(-), or underscores(_).
	InstanceName *string `json:"instanceName,omitempty" tf:"instance_name,omitempty"`

	// License type. Valid values are oss, basic and platinum. The default value is platinum.
	// License type. Valid values are `oss`, `basic` and `platinum`. The default value is `platinum`.
	LicenseType *string `json:"licenseType,omitempty" tf:"license_type,omitempty"`

	// Details of AZs in multi-AZ deployment mode (which is required when deploy_mode is 1).
	// Details of AZs in multi-AZ deployment mode (which is required when deploy_mode is `1`).
	MultiZoneInfos []MultiZoneInfosInitParameters `json:"multiZoneInfos,omitempty" tf:"multi_zone_infos,omitempty"`

	// Node information list, which is used to describe the specification information of various types of nodes in the cluster, such as node type, node quantity, node specification, disk type, and disk size.
	// Node information list, which is used to describe the specification information of various types of nodes in the cluster, such as node type, node quantity, node specification, disk type, and disk size.
	NodeInfoList []NodeInfoListInitParameters `json:"nodeInfoList,omitempty" tf:"node_info_list,omitempty"`

	// When enabled, the instance will be renew automatically when it reach the end of the prepaid tenancy. Valid values are RENEW_FLAG_AUTO and RENEW_FLAG_MANUAL. NOTE: it only works when charge_type is set to PREPAID.
	// When enabled, the instance will be renew automatically when it reach the end of the prepaid tenancy. Valid values are `RENEW_FLAG_AUTO` and `RENEW_FLAG_MANUAL`. NOTE: it only works when charge_type is set to `PREPAID`.
	RenewFlag *string `json:"renewFlag,omitempty" tf:"renew_flag,omitempty"`

	// The ID of a VPC subnetwork. When create multi-az es, this parameter must be omitted or -.
	// The ID of a VPC subnetwork. When create multi-az es, this parameter must be omitted or `-`.
	// +crossplane:generate:reference:type=github.com/crossplane-contrib/provider-tencentcloud/apis/vpc/v1alpha1.Subnet
	SubnetID *string `json:"subnetId,omitempty" tf:"subnet_id,omitempty"`

	// Reference to a Subnet in vpc to populate subnetId.
	// +kubebuilder:validation:Optional
	SubnetIDRef *v1.Reference `json:"subnetIdRef,omitempty" tf:"-"`

	// Selector for a Subnet in vpc to populate subnetId.
	// +kubebuilder:validation:Optional
	SubnetIDSelector *v1.Selector `json:"subnetIdSelector,omitempty" tf:"-"`

	// A mapping of tags to assign to the instance. For tag limits, please refer to Use Limits.
	// A mapping of tags to assign to the instance. For tag limits, please refer to [Use Limits](https://intl.cloud.tencent.com/document/product/651/13354).
	// +mapType=granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// The ID of a VPC network.
	// The ID of a VPC network.
	// +crossplane:generate:reference:type=github.com/crossplane-contrib/provider-tencentcloud/apis/vpc/v1alpha1.VPC
	VPCID *string `json:"vpcId,omitempty" tf:"vpc_id,omitempty"`

	// Reference to a VPC in vpc to populate vpcId.
	// +kubebuilder:validation:Optional
	VPCIDRef *v1.Reference `json:"vpcIdRef,omitempty" tf:"-"`

	// Selector for a VPC in vpc to populate vpcId.
	// +kubebuilder:validation:Optional
	VPCIDSelector *v1.Selector `json:"vpcIdSelector,omitempty" tf:"-"`

	// Version of the instance. Valid values are 5.6.4, 6.4.3, 6.8.2, 7.5.1 and 7.10.1.
	// Version of the instance. Valid values are `5.6.4`, `6.4.3`, `6.8.2`, `7.5.1` and `7.10.1`.
	Version *string `json:"version,omitempty" tf:"version,omitempty"`

	// Visual node configuration.
	// Visual node configuration.
	WebNodeTypeInfo []WebNodeTypeInfoInitParameters `json:"webNodeTypeInfo,omitempty" tf:"web_node_type_info,omitempty"`
}

type InstanceObservation struct {

	// Availability zone. When create multi-az es, this parameter must be omitted or -.
	// Availability zone. When create multi-az es, this parameter must be omitted or `-`.
	AvailabilityZone *string `json:"availabilityZone,omitempty" tf:"availability_zone,omitempty"`

	// Whether to enable X-Pack security authentication in Basic Edition 6.8 and above. Valid values are 1 and 2. 1 is disabled, 2 is enabled, and default value is 1. Notice: this parameter is only take effect on basic license.
	// Whether to enable X-Pack security authentication in Basic Edition 6.8 and above. Valid values are `1` and `2`. `1` is disabled, `2` is enabled, and default value is `1`. Notice: this parameter is only take effect on `basic` license.
	BasicSecurityType *float64 `json:"basicSecurityType,omitempty" tf:"basic_security_type,omitempty"`

	// The tenancy of the prepaid instance, and uint is month. NOTE: it only works when charge_type is set to PREPAID.
	// The tenancy of the prepaid instance, and uint is month. NOTE: it only works when charge_type is set to `PREPAID`.
	ChargePeriod *float64 `json:"chargePeriod,omitempty" tf:"charge_period,omitempty"`

	// The charge type of instance. Valid values are PREPAID and POSTPAID_BY_HOUR.
	// The charge type of instance. Valid values are `PREPAID` and `POSTPAID_BY_HOUR`.
	ChargeType *string `json:"chargeType,omitempty" tf:"charge_type,omitempty"`

	// Instance creation time.
	// Instance creation time.
	CreateTime *string `json:"createTime,omitempty" tf:"create_time,omitempty"`

	// Cluster deployment mode. Valid values are 0 and 1. 0 is single-AZ deployment, and 1 is multi-AZ deployment. Default value is 0.
	// Cluster deployment mode. Valid values are `0` and `1`. `0` is single-AZ deployment, and `1` is multi-AZ deployment. Default value is `0`.
	DeployMode *float64 `json:"deployMode,omitempty" tf:"deploy_mode,omitempty"`

	// Elasticsearch domain name.
	// Elasticsearch domain name.
	ElasticsearchDomain *string `json:"elasticsearchDomain,omitempty" tf:"elasticsearch_domain,omitempty"`

	// Elasticsearch port.
	// Elasticsearch port.
	ElasticsearchPort *float64 `json:"elasticsearchPort,omitempty" tf:"elasticsearch_port,omitempty"`

	// Elasticsearch VIP.
	// Elasticsearch VIP.
	ElasticsearchVip *string `json:"elasticsearchVip,omitempty" tf:"elasticsearch_vip,omitempty"`

	// Kibana Access Control Configuration.
	// Kibana Access Control Configuration.
	EsACL []EsACLObservation `json:"esAcl,omitempty" tf:"es_acl,omitempty"`

	// ID of the resource.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Name of the instance, which can contain 1 to 50 English letters, Chinese characters, digits, dashes(-), or underscores(_).
	// Name of the instance, which can contain 1 to 50 English letters, Chinese characters, digits, dashes(-), or underscores(_).
	InstanceName *string `json:"instanceName,omitempty" tf:"instance_name,omitempty"`

	// Kibana access URL.
	// Kibana access URL.
	KibanaURL *string `json:"kibanaUrl,omitempty" tf:"kibana_url,omitempty"`

	// License type. Valid values are oss, basic and platinum. The default value is platinum.
	// License type. Valid values are `oss`, `basic` and `platinum`. The default value is `platinum`.
	LicenseType *string `json:"licenseType,omitempty" tf:"license_type,omitempty"`

	// Details of AZs in multi-AZ deployment mode (which is required when deploy_mode is 1).
	// Details of AZs in multi-AZ deployment mode (which is required when deploy_mode is `1`).
	MultiZoneInfos []MultiZoneInfosObservation `json:"multiZoneInfos,omitempty" tf:"multi_zone_infos,omitempty"`

	// Node information list, which is used to describe the specification information of various types of nodes in the cluster, such as node type, node quantity, node specification, disk type, and disk size.
	// Node information list, which is used to describe the specification information of various types of nodes in the cluster, such as node type, node quantity, node specification, disk type, and disk size.
	NodeInfoList []NodeInfoListObservation `json:"nodeInfoList,omitempty" tf:"node_info_list,omitempty"`

	// When enabled, the instance will be renew automatically when it reach the end of the prepaid tenancy. Valid values are RENEW_FLAG_AUTO and RENEW_FLAG_MANUAL. NOTE: it only works when charge_type is set to PREPAID.
	// When enabled, the instance will be renew automatically when it reach the end of the prepaid tenancy. Valid values are `RENEW_FLAG_AUTO` and `RENEW_FLAG_MANUAL`. NOTE: it only works when charge_type is set to `PREPAID`.
	RenewFlag *string `json:"renewFlag,omitempty" tf:"renew_flag,omitempty"`

	// The ID of a VPC subnetwork. When create multi-az es, this parameter must be omitted or -.
	// The ID of a VPC subnetwork. When create multi-az es, this parameter must be omitted or `-`.
	SubnetID *string `json:"subnetId,omitempty" tf:"subnet_id,omitempty"`

	// A mapping of tags to assign to the instance. For tag limits, please refer to Use Limits.
	// A mapping of tags to assign to the instance. For tag limits, please refer to [Use Limits](https://intl.cloud.tencent.com/document/product/651/13354).
	// +mapType=granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// The ID of a VPC network.
	// The ID of a VPC network.
	VPCID *string `json:"vpcId,omitempty" tf:"vpc_id,omitempty"`

	// Version of the instance. Valid values are 5.6.4, 6.4.3, 6.8.2, 7.5.1 and 7.10.1.
	// Version of the instance. Valid values are `5.6.4`, `6.4.3`, `6.8.2`, `7.5.1` and `7.10.1`.
	Version *string `json:"version,omitempty" tf:"version,omitempty"`

	// Visual node configuration.
	// Visual node configuration.
	WebNodeTypeInfo []WebNodeTypeInfoObservation `json:"webNodeTypeInfo,omitempty" tf:"web_node_type_info,omitempty"`
}

type InstanceParameters struct {

	// Availability zone. When create multi-az es, this parameter must be omitted or -.
	// Availability zone. When create multi-az es, this parameter must be omitted or `-`.
	// +kubebuilder:validation:Optional
	AvailabilityZone *string `json:"availabilityZone,omitempty" tf:"availability_zone,omitempty"`

	// Whether to enable X-Pack security authentication in Basic Edition 6.8 and above. Valid values are 1 and 2. 1 is disabled, 2 is enabled, and default value is 1. Notice: this parameter is only take effect on basic license.
	// Whether to enable X-Pack security authentication in Basic Edition 6.8 and above. Valid values are `1` and `2`. `1` is disabled, `2` is enabled, and default value is `1`. Notice: this parameter is only take effect on `basic` license.
	// +kubebuilder:validation:Optional
	BasicSecurityType *float64 `json:"basicSecurityType,omitempty" tf:"basic_security_type,omitempty"`

	// The tenancy of the prepaid instance, and uint is month. NOTE: it only works when charge_type is set to PREPAID.
	// The tenancy of the prepaid instance, and uint is month. NOTE: it only works when charge_type is set to `PREPAID`.
	// +kubebuilder:validation:Optional
	ChargePeriod *float64 `json:"chargePeriod,omitempty" tf:"charge_period,omitempty"`

	// The charge type of instance. Valid values are PREPAID and POSTPAID_BY_HOUR.
	// The charge type of instance. Valid values are `PREPAID` and `POSTPAID_BY_HOUR`.
	// +kubebuilder:validation:Optional
	ChargeType *string `json:"chargeType,omitempty" tf:"charge_type,omitempty"`

	// Cluster deployment mode. Valid values are 0 and 1. 0 is single-AZ deployment, and 1 is multi-AZ deployment. Default value is 0.
	// Cluster deployment mode. Valid values are `0` and `1`. `0` is single-AZ deployment, and `1` is multi-AZ deployment. Default value is `0`.
	// +kubebuilder:validation:Optional
	DeployMode *float64 `json:"deployMode,omitempty" tf:"deploy_mode,omitempty"`

	// Kibana Access Control Configuration.
	// Kibana Access Control Configuration.
	// +kubebuilder:validation:Optional
	EsACL []EsACLParameters `json:"esAcl,omitempty" tf:"es_acl,omitempty"`

	// Name of the instance, which can contain 1 to 50 English letters, Chinese characters, digits, dashes(-), or underscores(_).
	// Name of the instance, which can contain 1 to 50 English letters, Chinese characters, digits, dashes(-), or underscores(_).
	// +kubebuilder:validation:Optional
	InstanceName *string `json:"instanceName,omitempty" tf:"instance_name,omitempty"`

	// License type. Valid values are oss, basic and platinum. The default value is platinum.
	// License type. Valid values are `oss`, `basic` and `platinum`. The default value is `platinum`.
	// +kubebuilder:validation:Optional
	LicenseType *string `json:"licenseType,omitempty" tf:"license_type,omitempty"`

	// Details of AZs in multi-AZ deployment mode (which is required when deploy_mode is 1).
	// Details of AZs in multi-AZ deployment mode (which is required when deploy_mode is `1`).
	// +kubebuilder:validation:Optional
	MultiZoneInfos []MultiZoneInfosParameters `json:"multiZoneInfos,omitempty" tf:"multi_zone_infos,omitempty"`

	// Node information list, which is used to describe the specification information of various types of nodes in the cluster, such as node type, node quantity, node specification, disk type, and disk size.
	// Node information list, which is used to describe the specification information of various types of nodes in the cluster, such as node type, node quantity, node specification, disk type, and disk size.
	// +kubebuilder:validation:Optional
	NodeInfoList []NodeInfoListParameters `json:"nodeInfoList,omitempty" tf:"node_info_list,omitempty"`

	// Password to an instance, the password needs to be 8 to 16 characters, including at least two items ([a-z,A-Z], [0-9] and [-!@#$%&^*+=_:;,.?] special symbols.
	// Password to an instance, the password needs to be 8 to 16 characters, including at least two items ([a-z,A-Z], [0-9] and [-!@#$%&^*+=_:;,.?] special symbols.
	// +kubebuilder:validation:Optional
	PasswordSecretRef v1.SecretKeySelector `json:"passwordSecretRef" tf:"-"`

	// When enabled, the instance will be renew automatically when it reach the end of the prepaid tenancy. Valid values are RENEW_FLAG_AUTO and RENEW_FLAG_MANUAL. NOTE: it only works when charge_type is set to PREPAID.
	// When enabled, the instance will be renew automatically when it reach the end of the prepaid tenancy. Valid values are `RENEW_FLAG_AUTO` and `RENEW_FLAG_MANUAL`. NOTE: it only works when charge_type is set to `PREPAID`.
	// +kubebuilder:validation:Optional
	RenewFlag *string `json:"renewFlag,omitempty" tf:"renew_flag,omitempty"`

	// The ID of a VPC subnetwork. When create multi-az es, this parameter must be omitted or -.
	// The ID of a VPC subnetwork. When create multi-az es, this parameter must be omitted or `-`.
	// +crossplane:generate:reference:type=github.com/crossplane-contrib/provider-tencentcloud/apis/vpc/v1alpha1.Subnet
	// +kubebuilder:validation:Optional
	SubnetID *string `json:"subnetId,omitempty" tf:"subnet_id,omitempty"`

	// Reference to a Subnet in vpc to populate subnetId.
	// +kubebuilder:validation:Optional
	SubnetIDRef *v1.Reference `json:"subnetIdRef,omitempty" tf:"-"`

	// Selector for a Subnet in vpc to populate subnetId.
	// +kubebuilder:validation:Optional
	SubnetIDSelector *v1.Selector `json:"subnetIdSelector,omitempty" tf:"-"`

	// A mapping of tags to assign to the instance. For tag limits, please refer to Use Limits.
	// A mapping of tags to assign to the instance. For tag limits, please refer to [Use Limits](https://intl.cloud.tencent.com/document/product/651/13354).
	// +kubebuilder:validation:Optional
	// +mapType=granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// The ID of a VPC network.
	// The ID of a VPC network.
	// +crossplane:generate:reference:type=github.com/crossplane-contrib/provider-tencentcloud/apis/vpc/v1alpha1.VPC
	// +kubebuilder:validation:Optional
	VPCID *string `json:"vpcId,omitempty" tf:"vpc_id,omitempty"`

	// Reference to a VPC in vpc to populate vpcId.
	// +kubebuilder:validation:Optional
	VPCIDRef *v1.Reference `json:"vpcIdRef,omitempty" tf:"-"`

	// Selector for a VPC in vpc to populate vpcId.
	// +kubebuilder:validation:Optional
	VPCIDSelector *v1.Selector `json:"vpcIdSelector,omitempty" tf:"-"`

	// Version of the instance. Valid values are 5.6.4, 6.4.3, 6.8.2, 7.5.1 and 7.10.1.
	// Version of the instance. Valid values are `5.6.4`, `6.4.3`, `6.8.2`, `7.5.1` and `7.10.1`.
	// +kubebuilder:validation:Optional
	Version *string `json:"version,omitempty" tf:"version,omitempty"`

	// Visual node configuration.
	// Visual node configuration.
	// +kubebuilder:validation:Optional
	WebNodeTypeInfo []WebNodeTypeInfoParameters `json:"webNodeTypeInfo,omitempty" tf:"web_node_type_info,omitempty"`
}

type MultiZoneInfosInitParameters struct {

	// Availability zone. When create multi-az es, this parameter must be omitted or -.
	// Availability zone.
	AvailabilityZone *string `json:"availabilityZone,omitempty" tf:"availability_zone,omitempty"`

	// The ID of a VPC subnetwork. When create multi-az es, this parameter must be omitted or -.
	// The ID of a VPC subnetwork.
	SubnetID *string `json:"subnetId,omitempty" tf:"subnet_id,omitempty"`
}

type MultiZoneInfosObservation struct {

	// Availability zone. When create multi-az es, this parameter must be omitted or -.
	// Availability zone.
	AvailabilityZone *string `json:"availabilityZone,omitempty" tf:"availability_zone,omitempty"`

	// The ID of a VPC subnetwork. When create multi-az es, this parameter must be omitted or -.
	// The ID of a VPC subnetwork.
	SubnetID *string `json:"subnetId,omitempty" tf:"subnet_id,omitempty"`
}

type MultiZoneInfosParameters struct {

	// Availability zone. When create multi-az es, this parameter must be omitted or -.
	// Availability zone.
	// +kubebuilder:validation:Optional
	AvailabilityZone *string `json:"availabilityZone" tf:"availability_zone,omitempty"`

	// The ID of a VPC subnetwork. When create multi-az es, this parameter must be omitted or -.
	// The ID of a VPC subnetwork.
	// +kubebuilder:validation:Optional
	SubnetID *string `json:"subnetId" tf:"subnet_id,omitempty"`
}

type NodeInfoListInitParameters struct {

	// Node disk size. Unit is GB, and default value is 100.
	// Node disk size. Unit is GB, and default value is `100`.
	DiskSize *float64 `json:"diskSize,omitempty" tf:"disk_size,omitempty"`

	// Node disk type. Valid values are CLOUD_SSD and CLOUD_PREMIUM, CLOUD_HSSD. The default value is CLOUD_SSD.
	// Node disk type. Valid values are `CLOUD_SSD` and `CLOUD_PREMIUM`, `CLOUD_HSSD`. The default value is `CLOUD_SSD`.
	DiskType *string `json:"diskType,omitempty" tf:"disk_type,omitempty"`

	// Decides to encrypt this disk or not.
	// Decides to encrypt this disk or not.
	Encrypt *bool `json:"encrypt,omitempty" tf:"encrypt,omitempty"`

	// Number of nodes.
	// Number of nodes.
	NodeNum *float64 `json:"nodeNum,omitempty" tf:"node_num,omitempty"`

	// Node specification, and valid values refer to document of tencentcloud.
	// Node specification, and valid values refer to [document of tencentcloud](https://intl.cloud.tencent.com/document/product/845/18376).
	NodeType *string `json:"nodeType,omitempty" tf:"node_type,omitempty"`

	// Node type. Valid values are hotData, warmData and dedicatedMaster. The default value is 'hotData`.
	// Node type. Valid values are `hotData`, `warmData` and `dedicatedMaster`. The default value is 'hotData`.
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

type NodeInfoListObservation struct {

	// Node disk size. Unit is GB, and default value is 100.
	// Node disk size. Unit is GB, and default value is `100`.
	DiskSize *float64 `json:"diskSize,omitempty" tf:"disk_size,omitempty"`

	// Node disk type. Valid values are CLOUD_SSD and CLOUD_PREMIUM, CLOUD_HSSD. The default value is CLOUD_SSD.
	// Node disk type. Valid values are `CLOUD_SSD` and `CLOUD_PREMIUM`, `CLOUD_HSSD`. The default value is `CLOUD_SSD`.
	DiskType *string `json:"diskType,omitempty" tf:"disk_type,omitempty"`

	// Decides to encrypt this disk or not.
	// Decides to encrypt this disk or not.
	Encrypt *bool `json:"encrypt,omitempty" tf:"encrypt,omitempty"`

	// Number of nodes.
	// Number of nodes.
	NodeNum *float64 `json:"nodeNum,omitempty" tf:"node_num,omitempty"`

	// Node specification, and valid values refer to document of tencentcloud.
	// Node specification, and valid values refer to [document of tencentcloud](https://intl.cloud.tencent.com/document/product/845/18376).
	NodeType *string `json:"nodeType,omitempty" tf:"node_type,omitempty"`

	// Node type. Valid values are hotData, warmData and dedicatedMaster. The default value is 'hotData`.
	// Node type. Valid values are `hotData`, `warmData` and `dedicatedMaster`. The default value is 'hotData`.
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

type NodeInfoListParameters struct {

	// Node disk size. Unit is GB, and default value is 100.
	// Node disk size. Unit is GB, and default value is `100`.
	// +kubebuilder:validation:Optional
	DiskSize *float64 `json:"diskSize,omitempty" tf:"disk_size,omitempty"`

	// Node disk type. Valid values are CLOUD_SSD and CLOUD_PREMIUM, CLOUD_HSSD. The default value is CLOUD_SSD.
	// Node disk type. Valid values are `CLOUD_SSD` and `CLOUD_PREMIUM`, `CLOUD_HSSD`. The default value is `CLOUD_SSD`.
	// +kubebuilder:validation:Optional
	DiskType *string `json:"diskType,omitempty" tf:"disk_type,omitempty"`

	// Decides to encrypt this disk or not.
	// Decides to encrypt this disk or not.
	// +kubebuilder:validation:Optional
	Encrypt *bool `json:"encrypt,omitempty" tf:"encrypt,omitempty"`

	// Number of nodes.
	// Number of nodes.
	// +kubebuilder:validation:Optional
	NodeNum *float64 `json:"nodeNum" tf:"node_num,omitempty"`

	// Node specification, and valid values refer to document of tencentcloud.
	// Node specification, and valid values refer to [document of tencentcloud](https://intl.cloud.tencent.com/document/product/845/18376).
	// +kubebuilder:validation:Optional
	NodeType *string `json:"nodeType" tf:"node_type,omitempty"`

	// Node type. Valid values are hotData, warmData and dedicatedMaster. The default value is 'hotData`.
	// Node type. Valid values are `hotData`, `warmData` and `dedicatedMaster`. The default value is 'hotData`.
	// +kubebuilder:validation:Optional
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

type WebNodeTypeInfoInitParameters struct {

	// Number of nodes.
	// Visual node number.
	NodeNum *float64 `json:"nodeNum,omitempty" tf:"node_num,omitempty"`

	// Node specification, and valid values refer to document of tencentcloud.
	// Visual node specifications.
	NodeType *string `json:"nodeType,omitempty" tf:"node_type,omitempty"`
}

type WebNodeTypeInfoObservation struct {

	// Number of nodes.
	// Visual node number.
	NodeNum *float64 `json:"nodeNum,omitempty" tf:"node_num,omitempty"`

	// Node specification, and valid values refer to document of tencentcloud.
	// Visual node specifications.
	NodeType *string `json:"nodeType,omitempty" tf:"node_type,omitempty"`
}

type WebNodeTypeInfoParameters struct {

	// Number of nodes.
	// Visual node number.
	// +kubebuilder:validation:Optional
	NodeNum *float64 `json:"nodeNum" tf:"node_num,omitempty"`

	// Node specification, and valid values refer to document of tencentcloud.
	// Visual node specifications.
	// +kubebuilder:validation:Optional
	NodeType *string `json:"nodeType" tf:"node_type,omitempty"`
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

// Instance is the Schema for the Instances API. Provides an elasticsearch instance resource.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,tencentcloud}
type Instance struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.nodeInfoList) || (has(self.initProvider) && has(self.initProvider.nodeInfoList))",message="spec.forProvider.nodeInfoList is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.passwordSecretRef)",message="spec.forProvider.passwordSecretRef is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.version) || (has(self.initProvider) && has(self.initProvider.version))",message="spec.forProvider.version is a required parameter"
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
