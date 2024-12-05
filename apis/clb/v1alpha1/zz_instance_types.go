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

	// It's only applicable to public network CLB instances. IP version. Values: IPV4, IPV6 and IPv6FullChain (case-insensitive). Default: IPV4. Note: IPV6 indicates IPv6 NAT64, while IPv6FullChain indicates IPv6.
	// It's only applicable to public network CLB instances. IP version. Values: `IPV4`, `IPV6` and `IPv6FullChain` (case-insensitive). Default: `IPV4`. Note: IPV6 indicates IPv6 NAT64, while IPv6FullChain indicates IPv6.
	AddressIPVersion *string `json:"addressIpVersion,omitempty" tf:"address_ip_version,omitempty"`

	// Bandwidth package id. If set, the internet_charge_type must be BANDWIDTH_PACKAGE.
	// Bandwidth package id. If set, the `internet_charge_type` must be `BANDWIDTH_PACKAGE`.
	BandwidthPackageID *string `json:"bandwidthPackageId,omitempty" tf:"bandwidth_package_id,omitempty"`

	// Name of the CLB. The name can only contain Chinese characters, English letters, numbers, underscore and hyphen '-'.
	// Name of the CLB. The name can only contain Chinese characters, English letters, numbers, underscore and hyphen '-'.
	ClbName *string `json:"clbName,omitempty" tf:"clb_name,omitempty"`

	// Cluster ID.
	// Cluster ID.
	ClusterID *string `json:"clusterId,omitempty" tf:"cluster_id,omitempty"`

	// Whether to enable delete protection.
	// Whether to enable delete protection.
	DeleteProtect *bool `json:"deleteProtect,omitempty" tf:"delete_protect,omitempty"`

	// If create dynamic vip CLB instance, true or false.
	// If create dynamic vip CLB instance, `true` or `false`.
	DynamicVip *bool `json:"dynamicVip,omitempty" tf:"dynamic_vip,omitempty"`

	// Max bandwidth out, only applicable to open CLB. Valid value ranges is [1, 2048]. Unit is MB.
	// Max bandwidth out, only applicable to open CLB. Valid value ranges is [1, 2048]. Unit is MB.
	InternetBandwidthMaxOut *float64 `json:"internetBandwidthMaxOut,omitempty" tf:"internet_bandwidth_max_out,omitempty"`

	// Internet charge type, only applicable to open CLB. Valid values are TRAFFIC_POSTPAID_BY_HOUR, BANDWIDTH_POSTPAID_BY_HOUR and BANDWIDTH_PACKAGE.
	// Internet charge type, only applicable to open CLB. Valid values are `TRAFFIC_POSTPAID_BY_HOUR`, `BANDWIDTH_POSTPAID_BY_HOUR` and `BANDWIDTH_PACKAGE`.
	InternetChargeType *string `json:"internetChargeType,omitempty" tf:"internet_charge_type,omitempty"`

	// Whether the target allow flow come from clb. If value is true, only check security group of clb, or check both clb and backend instance security group.
	// Whether the target allow flow come from clb. If value is true, only check security group of clb, or check both clb and backend instance security group.
	LoadBalancerPassToTarget *bool `json:"loadBalancerPassToTarget,omitempty" tf:"load_balancer_pass_to_target,omitempty"`

	// The id of log set.
	// The id of log set.
	LogSetID *string `json:"logSetId,omitempty" tf:"log_set_id,omitempty"`

	// The id of log topic.
	// The id of log topic.
	LogTopicID *string `json:"logTopicId,omitempty" tf:"log_topic_id,omitempty"`

	// Setting master zone id of cross available zone disaster recovery, only applicable to open CLB.
	// Setting master zone id of cross available zone disaster recovery, only applicable to open CLB.
	MasterZoneID *string `json:"masterZoneId,omitempty" tf:"master_zone_id,omitempty"`

	// Type of CLB instance. Valid values: OPEN and INTERNAL.
	// Type of CLB instance. Valid values: `OPEN` and `INTERNAL`.
	NetworkType *string `json:"networkType,omitempty" tf:"network_type,omitempty"`

	// ID of the project within the CLB instance, 0 - Default Project.
	// ID of the project within the CLB instance, `0` - Default Project.
	ProjectID *float64 `json:"projectId,omitempty" tf:"project_id,omitempty"`

	// This parameter is required to create LCU-supported instances. Values:SLA: Super Large 4. When you have activated Super Large models, SLA refers to Super Large 4; clb.c2.medium: Standard; clb.c3.small: Advanced 1; clb.c3.medium: Advanced 1; clb.c4.small: Super Large 1; clb.c4.medium: Super Large 2; clb.c4.large: Super Large 3; clb.c4.xlarge: Super Large 4. For more details, see Instance Specifications.
	// This parameter is required to create LCU-supported instances. Values:`SLA`: Super Large 4. When you have activated Super Large models, `SLA` refers to Super Large 4; `clb.c2.medium`: Standard; `clb.c3.small`: Advanced 1; `clb.c3.medium`: Advanced 1; `clb.c4.small`: Super Large 1; `clb.c4.medium`: Super Large 2; `clb.c4.large`: Super Large 3; `clb.c4.xlarge`: Super Large 4. For more details, see [Instance Specifications](https://intl.cloud.tencent.com/document/product/214/84689?from_cn_redirect=1).
	SLAType *string `json:"slaType,omitempty" tf:"sla_type,omitempty"`

	// Security groups of the CLB instance. Supports both OPEN and INTERNAL CLBs.
	// Security groups of the CLB instance. Supports both `OPEN` and `INTERNAL` CLBs.
	SecurityGroups []*string `json:"securityGroups,omitempty" tf:"security_groups,omitempty"`

	// Setting slave zone id of cross available zone disaster recovery, only applicable to open CLB. this zone will undertake traffic when the master is down.
	// Setting slave zone id of cross available zone disaster recovery, only applicable to open CLB. this zone will undertake traffic when the master is down.
	SlaveZoneID *string `json:"slaveZoneId,omitempty" tf:"slave_zone_id,omitempty"`

	// Snat Ip List, required with snat_pro=true. NOTE: This argument cannot be read and modified here because dynamic ip is untraceable, please import resource tencentcloud_clb_snat_ip to handle fixed ips.
	// Snat Ip List, required with `snat_pro=true`. NOTE: This argument cannot be read and modified here because dynamic ip is untraceable, please import resource `tencentcloud_clb_snat_ip` to handle fixed ips.
	SnatIps []SnatIpsInitParameters `json:"snatIps,omitempty" tf:"snat_ips,omitempty"`

	// Indicates whether Binding IPs of other VPCs feature switch.
	// Indicates whether Binding IPs of other VPCs feature switch.
	SnatPro *bool `json:"snatPro,omitempty" tf:"snat_pro,omitempty"`

	// In the case of purchasing a INTERNAL clb instance, the subnet id must be specified. The VIP of the INTERNAL clb instance will be generated from this subnet.
	// In the case of purchasing a `INTERNAL` clb instance, the subnet id must be specified. The VIP of the `INTERNAL` clb instance will be generated from this subnet.
	// +crossplane:generate:reference:type=github.com/crossplane-contrib/provider-tencentcloud/apis/vpc/v1alpha1.Subnet
	SubnetID *string `json:"subnetId,omitempty" tf:"subnet_id,omitempty"`

	// Reference to a Subnet in vpc to populate subnetId.
	// +kubebuilder:validation:Optional
	SubnetIDRef *v1.Reference `json:"subnetIdRef,omitempty" tf:"-"`

	// Selector for a Subnet in vpc to populate subnetId.
	// +kubebuilder:validation:Optional
	SubnetIDSelector *v1.Selector `json:"subnetIdSelector,omitempty" tf:"-"`

	// The available tags within this CLB.
	// The available tags within this CLB.
	// +mapType=granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// Region information of backend services are attached the CLB instance. Only supports OPEN CLBs.
	// Region information of backend services are attached the CLB instance. Only supports `OPEN` CLBs.
	TargetRegionInfoRegion *string `json:"targetRegionInfoRegion,omitempty" tf:"target_region_info_region,omitempty"`

	// Vpc information of backend services are attached the CLB instance. Only supports OPEN CLBs.
	// Vpc information of backend services are attached the CLB instance. Only supports `OPEN` CLBs.
	TargetRegionInfoVPCID *string `json:"targetRegionInfoVpcId,omitempty" tf:"target_region_info_vpc_id,omitempty"`

	// VPC ID of the CLB.
	// VPC ID of the CLB.
	// +crossplane:generate:reference:type=github.com/crossplane-contrib/provider-tencentcloud/apis/vpc/v1alpha1.VPC
	VPCID *string `json:"vpcId,omitempty" tf:"vpc_id,omitempty"`

	// Reference to a VPC in vpc to populate vpcId.
	// +kubebuilder:validation:Optional
	VPCIDRef *v1.Reference `json:"vpcIdRef,omitempty" tf:"-"`

	// Selector for a VPC in vpc to populate vpcId.
	// +kubebuilder:validation:Optional
	VPCIDSelector *v1.Selector `json:"vpcIdSelector,omitempty" tf:"-"`

	// Specifies the VIP for the application of a CLB instance. This parameter is optional. If you do not specify this parameter, the system automatically assigns a value for the parameter. IPv4 and IPv6 CLB instances support this parameter, but IPv6 NAT64 CLB instances do not.
	// Specifies the VIP for the application of a CLB instance. This parameter is optional. If you do not specify this parameter, the system automatically assigns a value for the parameter. IPv4 and IPv6 CLB instances support this parameter, but IPv6 NAT64 CLB instances do not.
	Vip *string `json:"vip,omitempty" tf:"vip,omitempty"`

	// Network operator, only applicable to open CLB. Valid values are CMCC(China Mobile), CTCC(Telecom), CUCC(China Unicom) and BGP. If this ISP is specified, network billing method can only use the bandwidth package billing (BANDWIDTH_PACKAGE).
	// Network operator, only applicable to open CLB. Valid values are `CMCC`(China Mobile), `CTCC`(Telecom), `CUCC`(China Unicom) and `BGP`. If this ISP is specified, network billing method can only use the bandwidth package billing (BANDWIDTH_PACKAGE).
	VipIsp *string `json:"vipIsp,omitempty" tf:"vip_isp,omitempty"`

	// Available zone id, only applicable to open CLB.
	// Available zone id, only applicable to open CLB.
	ZoneID *string `json:"zoneId,omitempty" tf:"zone_id,omitempty"`
}

type InstanceObservation struct {

	// It's only applicable to public network CLB instances. IP version. Values: IPV4, IPV6 and IPv6FullChain (case-insensitive). Default: IPV4. Note: IPV6 indicates IPv6 NAT64, while IPv6FullChain indicates IPv6.
	// It's only applicable to public network CLB instances. IP version. Values: `IPV4`, `IPV6` and `IPv6FullChain` (case-insensitive). Default: `IPV4`. Note: IPV6 indicates IPv6 NAT64, while IPv6FullChain indicates IPv6.
	AddressIPVersion *string `json:"addressIpVersion,omitempty" tf:"address_ip_version,omitempty"`

	// The IPv6 address of the load balancing instance.
	// The IPv6 address of the load balancing instance.
	AddressIPv6 *string `json:"addressIpv6,omitempty" tf:"address_ipv6,omitempty"`

	// Bandwidth package id. If set, the internet_charge_type must be BANDWIDTH_PACKAGE.
	// Bandwidth package id. If set, the `internet_charge_type` must be `BANDWIDTH_PACKAGE`.
	BandwidthPackageID *string `json:"bandwidthPackageId,omitempty" tf:"bandwidth_package_id,omitempty"`

	// Name of the CLB. The name can only contain Chinese characters, English letters, numbers, underscore and hyphen '-'.
	// Name of the CLB. The name can only contain Chinese characters, English letters, numbers, underscore and hyphen '-'.
	ClbName *string `json:"clbName,omitempty" tf:"clb_name,omitempty"`

	// The virtual service address table of the CLB.
	// The virtual service address table of the CLB.
	ClbVips []*string `json:"clbVips,omitempty" tf:"clb_vips,omitempty"`

	// Cluster ID.
	// Cluster ID.
	ClusterID *string `json:"clusterId,omitempty" tf:"cluster_id,omitempty"`

	// Whether to enable delete protection.
	// Whether to enable delete protection.
	DeleteProtect *bool `json:"deleteProtect,omitempty" tf:"delete_protect,omitempty"`

	// Domain name of the CLB instance.
	// Domain name of the CLB instance.
	Domain *string `json:"domain,omitempty" tf:"domain,omitempty"`

	// If create dynamic vip CLB instance, true or false.
	// If create dynamic vip CLB instance, `true` or `false`.
	DynamicVip *bool `json:"dynamicVip,omitempty" tf:"dynamic_vip,omitempty"`

	// ID of the resource.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// This field is meaningful when the IP address version is ipv6, IPv6Nat64 | IPv6FullChain.
	// This field is meaningful when the IP address version is ipv6, `IPv6Nat64` | `IPv6FullChain`.
	IPv6Mode *string `json:"ipv6Mode,omitempty" tf:"ipv6_mode,omitempty"`

	// Max bandwidth out, only applicable to open CLB. Valid value ranges is [1, 2048]. Unit is MB.
	// Max bandwidth out, only applicable to open CLB. Valid value ranges is [1, 2048]. Unit is MB.
	InternetBandwidthMaxOut *float64 `json:"internetBandwidthMaxOut,omitempty" tf:"internet_bandwidth_max_out,omitempty"`

	// Internet charge type, only applicable to open CLB. Valid values are TRAFFIC_POSTPAID_BY_HOUR, BANDWIDTH_POSTPAID_BY_HOUR and BANDWIDTH_PACKAGE.
	// Internet charge type, only applicable to open CLB. Valid values are `TRAFFIC_POSTPAID_BY_HOUR`, `BANDWIDTH_POSTPAID_BY_HOUR` and `BANDWIDTH_PACKAGE`.
	InternetChargeType *string `json:"internetChargeType,omitempty" tf:"internet_charge_type,omitempty"`

	// Whether the target allow flow come from clb. If value is true, only check security group of clb, or check both clb and backend instance security group.
	// Whether the target allow flow come from clb. If value is true, only check security group of clb, or check both clb and backend instance security group.
	LoadBalancerPassToTarget *bool `json:"loadBalancerPassToTarget,omitempty" tf:"load_balancer_pass_to_target,omitempty"`

	// The id of log set.
	// The id of log set.
	LogSetID *string `json:"logSetId,omitempty" tf:"log_set_id,omitempty"`

	// The id of log topic.
	// The id of log topic.
	LogTopicID *string `json:"logTopicId,omitempty" tf:"log_topic_id,omitempty"`

	// Setting master zone id of cross available zone disaster recovery, only applicable to open CLB.
	// Setting master zone id of cross available zone disaster recovery, only applicable to open CLB.
	MasterZoneID *string `json:"masterZoneId,omitempty" tf:"master_zone_id,omitempty"`

	// Type of CLB instance. Valid values: OPEN and INTERNAL.
	// Type of CLB instance. Valid values: `OPEN` and `INTERNAL`.
	NetworkType *string `json:"networkType,omitempty" tf:"network_type,omitempty"`

	// ID of the project within the CLB instance, 0 - Default Project.
	// ID of the project within the CLB instance, `0` - Default Project.
	ProjectID *float64 `json:"projectId,omitempty" tf:"project_id,omitempty"`

	// This parameter is required to create LCU-supported instances. Values:SLA: Super Large 4. When you have activated Super Large models, SLA refers to Super Large 4; clb.c2.medium: Standard; clb.c3.small: Advanced 1; clb.c3.medium: Advanced 1; clb.c4.small: Super Large 1; clb.c4.medium: Super Large 2; clb.c4.large: Super Large 3; clb.c4.xlarge: Super Large 4. For more details, see Instance Specifications.
	// This parameter is required to create LCU-supported instances. Values:`SLA`: Super Large 4. When you have activated Super Large models, `SLA` refers to Super Large 4; `clb.c2.medium`: Standard; `clb.c3.small`: Advanced 1; `clb.c3.medium`: Advanced 1; `clb.c4.small`: Super Large 1; `clb.c4.medium`: Super Large 2; `clb.c4.large`: Super Large 3; `clb.c4.xlarge`: Super Large 4. For more details, see [Instance Specifications](https://intl.cloud.tencent.com/document/product/214/84689?from_cn_redirect=1).
	SLAType *string `json:"slaType,omitempty" tf:"sla_type,omitempty"`

	// Security groups of the CLB instance. Supports both OPEN and INTERNAL CLBs.
	// Security groups of the CLB instance. Supports both `OPEN` and `INTERNAL` CLBs.
	SecurityGroups []*string `json:"securityGroups,omitempty" tf:"security_groups,omitempty"`

	// Setting slave zone id of cross available zone disaster recovery, only applicable to open CLB. this zone will undertake traffic when the master is down.
	// Setting slave zone id of cross available zone disaster recovery, only applicable to open CLB. this zone will undertake traffic when the master is down.
	SlaveZoneID *string `json:"slaveZoneId,omitempty" tf:"slave_zone_id,omitempty"`

	// Snat Ip List, required with snat_pro=true. NOTE: This argument cannot be read and modified here because dynamic ip is untraceable, please import resource tencentcloud_clb_snat_ip to handle fixed ips.
	// Snat Ip List, required with `snat_pro=true`. NOTE: This argument cannot be read and modified here because dynamic ip is untraceable, please import resource `tencentcloud_clb_snat_ip` to handle fixed ips.
	SnatIps []SnatIpsObservation `json:"snatIps,omitempty" tf:"snat_ips,omitempty"`

	// Indicates whether Binding IPs of other VPCs feature switch.
	// Indicates whether Binding IPs of other VPCs feature switch.
	SnatPro *bool `json:"snatPro,omitempty" tf:"snat_pro,omitempty"`

	// In the case of purchasing a INTERNAL clb instance, the subnet id must be specified. The VIP of the INTERNAL clb instance will be generated from this subnet.
	// In the case of purchasing a `INTERNAL` clb instance, the subnet id must be specified. The VIP of the `INTERNAL` clb instance will be generated from this subnet.
	SubnetID *string `json:"subnetId,omitempty" tf:"subnet_id,omitempty"`

	// The available tags within this CLB.
	// The available tags within this CLB.
	// +mapType=granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// Region information of backend services are attached the CLB instance. Only supports OPEN CLBs.
	// Region information of backend services are attached the CLB instance. Only supports `OPEN` CLBs.
	TargetRegionInfoRegion *string `json:"targetRegionInfoRegion,omitempty" tf:"target_region_info_region,omitempty"`

	// Vpc information of backend services are attached the CLB instance. Only supports OPEN CLBs.
	// Vpc information of backend services are attached the CLB instance. Only supports `OPEN` CLBs.
	TargetRegionInfoVPCID *string `json:"targetRegionInfoVpcId,omitempty" tf:"target_region_info_vpc_id,omitempty"`

	// VPC ID of the CLB.
	// VPC ID of the CLB.
	VPCID *string `json:"vpcId,omitempty" tf:"vpc_id,omitempty"`

	// Specifies the VIP for the application of a CLB instance. This parameter is optional. If you do not specify this parameter, the system automatically assigns a value for the parameter. IPv4 and IPv6 CLB instances support this parameter, but IPv6 NAT64 CLB instances do not.
	// Specifies the VIP for the application of a CLB instance. This parameter is optional. If you do not specify this parameter, the system automatically assigns a value for the parameter. IPv4 and IPv6 CLB instances support this parameter, but IPv6 NAT64 CLB instances do not.
	Vip *string `json:"vip,omitempty" tf:"vip,omitempty"`

	// Network operator, only applicable to open CLB. Valid values are CMCC(China Mobile), CTCC(Telecom), CUCC(China Unicom) and BGP. If this ISP is specified, network billing method can only use the bandwidth package billing (BANDWIDTH_PACKAGE).
	// Network operator, only applicable to open CLB. Valid values are `CMCC`(China Mobile), `CTCC`(Telecom), `CUCC`(China Unicom) and `BGP`. If this ISP is specified, network billing method can only use the bandwidth package billing (BANDWIDTH_PACKAGE).
	VipIsp *string `json:"vipIsp,omitempty" tf:"vip_isp,omitempty"`

	// Available zone id, only applicable to open CLB.
	// Available zone id, only applicable to open CLB.
	ZoneID *string `json:"zoneId,omitempty" tf:"zone_id,omitempty"`
}

type InstanceParameters struct {

	// It's only applicable to public network CLB instances. IP version. Values: IPV4, IPV6 and IPv6FullChain (case-insensitive). Default: IPV4. Note: IPV6 indicates IPv6 NAT64, while IPv6FullChain indicates IPv6.
	// It's only applicable to public network CLB instances. IP version. Values: `IPV4`, `IPV6` and `IPv6FullChain` (case-insensitive). Default: `IPV4`. Note: IPV6 indicates IPv6 NAT64, while IPv6FullChain indicates IPv6.
	// +kubebuilder:validation:Optional
	AddressIPVersion *string `json:"addressIpVersion,omitempty" tf:"address_ip_version,omitempty"`

	// Bandwidth package id. If set, the internet_charge_type must be BANDWIDTH_PACKAGE.
	// Bandwidth package id. If set, the `internet_charge_type` must be `BANDWIDTH_PACKAGE`.
	// +kubebuilder:validation:Optional
	BandwidthPackageID *string `json:"bandwidthPackageId,omitempty" tf:"bandwidth_package_id,omitempty"`

	// Name of the CLB. The name can only contain Chinese characters, English letters, numbers, underscore and hyphen '-'.
	// Name of the CLB. The name can only contain Chinese characters, English letters, numbers, underscore and hyphen '-'.
	// +kubebuilder:validation:Optional
	ClbName *string `json:"clbName,omitempty" tf:"clb_name,omitempty"`

	// Cluster ID.
	// Cluster ID.
	// +kubebuilder:validation:Optional
	ClusterID *string `json:"clusterId,omitempty" tf:"cluster_id,omitempty"`

	// Whether to enable delete protection.
	// Whether to enable delete protection.
	// +kubebuilder:validation:Optional
	DeleteProtect *bool `json:"deleteProtect,omitempty" tf:"delete_protect,omitempty"`

	// If create dynamic vip CLB instance, true or false.
	// If create dynamic vip CLB instance, `true` or `false`.
	// +kubebuilder:validation:Optional
	DynamicVip *bool `json:"dynamicVip,omitempty" tf:"dynamic_vip,omitempty"`

	// Max bandwidth out, only applicable to open CLB. Valid value ranges is [1, 2048]. Unit is MB.
	// Max bandwidth out, only applicable to open CLB. Valid value ranges is [1, 2048]. Unit is MB.
	// +kubebuilder:validation:Optional
	InternetBandwidthMaxOut *float64 `json:"internetBandwidthMaxOut,omitempty" tf:"internet_bandwidth_max_out,omitempty"`

	// Internet charge type, only applicable to open CLB. Valid values are TRAFFIC_POSTPAID_BY_HOUR, BANDWIDTH_POSTPAID_BY_HOUR and BANDWIDTH_PACKAGE.
	// Internet charge type, only applicable to open CLB. Valid values are `TRAFFIC_POSTPAID_BY_HOUR`, `BANDWIDTH_POSTPAID_BY_HOUR` and `BANDWIDTH_PACKAGE`.
	// +kubebuilder:validation:Optional
	InternetChargeType *string `json:"internetChargeType,omitempty" tf:"internet_charge_type,omitempty"`

	// Whether the target allow flow come from clb. If value is true, only check security group of clb, or check both clb and backend instance security group.
	// Whether the target allow flow come from clb. If value is true, only check security group of clb, or check both clb and backend instance security group.
	// +kubebuilder:validation:Optional
	LoadBalancerPassToTarget *bool `json:"loadBalancerPassToTarget,omitempty" tf:"load_balancer_pass_to_target,omitempty"`

	// The id of log set.
	// The id of log set.
	// +kubebuilder:validation:Optional
	LogSetID *string `json:"logSetId,omitempty" tf:"log_set_id,omitempty"`

	// The id of log topic.
	// The id of log topic.
	// +kubebuilder:validation:Optional
	LogTopicID *string `json:"logTopicId,omitempty" tf:"log_topic_id,omitempty"`

	// Setting master zone id of cross available zone disaster recovery, only applicable to open CLB.
	// Setting master zone id of cross available zone disaster recovery, only applicable to open CLB.
	// +kubebuilder:validation:Optional
	MasterZoneID *string `json:"masterZoneId,omitempty" tf:"master_zone_id,omitempty"`

	// Type of CLB instance. Valid values: OPEN and INTERNAL.
	// Type of CLB instance. Valid values: `OPEN` and `INTERNAL`.
	// +kubebuilder:validation:Optional
	NetworkType *string `json:"networkType,omitempty" tf:"network_type,omitempty"`

	// ID of the project within the CLB instance, 0 - Default Project.
	// ID of the project within the CLB instance, `0` - Default Project.
	// +kubebuilder:validation:Optional
	ProjectID *float64 `json:"projectId,omitempty" tf:"project_id,omitempty"`

	// This parameter is required to create LCU-supported instances. Values:SLA: Super Large 4. When you have activated Super Large models, SLA refers to Super Large 4; clb.c2.medium: Standard; clb.c3.small: Advanced 1; clb.c3.medium: Advanced 1; clb.c4.small: Super Large 1; clb.c4.medium: Super Large 2; clb.c4.large: Super Large 3; clb.c4.xlarge: Super Large 4. For more details, see Instance Specifications.
	// This parameter is required to create LCU-supported instances. Values:`SLA`: Super Large 4. When you have activated Super Large models, `SLA` refers to Super Large 4; `clb.c2.medium`: Standard; `clb.c3.small`: Advanced 1; `clb.c3.medium`: Advanced 1; `clb.c4.small`: Super Large 1; `clb.c4.medium`: Super Large 2; `clb.c4.large`: Super Large 3; `clb.c4.xlarge`: Super Large 4. For more details, see [Instance Specifications](https://intl.cloud.tencent.com/document/product/214/84689?from_cn_redirect=1).
	// +kubebuilder:validation:Optional
	SLAType *string `json:"slaType,omitempty" tf:"sla_type,omitempty"`

	// Security groups of the CLB instance. Supports both OPEN and INTERNAL CLBs.
	// Security groups of the CLB instance. Supports both `OPEN` and `INTERNAL` CLBs.
	// +kubebuilder:validation:Optional
	SecurityGroups []*string `json:"securityGroups,omitempty" tf:"security_groups,omitempty"`

	// Setting slave zone id of cross available zone disaster recovery, only applicable to open CLB. this zone will undertake traffic when the master is down.
	// Setting slave zone id of cross available zone disaster recovery, only applicable to open CLB. this zone will undertake traffic when the master is down.
	// +kubebuilder:validation:Optional
	SlaveZoneID *string `json:"slaveZoneId,omitempty" tf:"slave_zone_id,omitempty"`

	// Snat Ip List, required with snat_pro=true. NOTE: This argument cannot be read and modified here because dynamic ip is untraceable, please import resource tencentcloud_clb_snat_ip to handle fixed ips.
	// Snat Ip List, required with `snat_pro=true`. NOTE: This argument cannot be read and modified here because dynamic ip is untraceable, please import resource `tencentcloud_clb_snat_ip` to handle fixed ips.
	// +kubebuilder:validation:Optional
	SnatIps []SnatIpsParameters `json:"snatIps,omitempty" tf:"snat_ips,omitempty"`

	// Indicates whether Binding IPs of other VPCs feature switch.
	// Indicates whether Binding IPs of other VPCs feature switch.
	// +kubebuilder:validation:Optional
	SnatPro *bool `json:"snatPro,omitempty" tf:"snat_pro,omitempty"`

	// In the case of purchasing a INTERNAL clb instance, the subnet id must be specified. The VIP of the INTERNAL clb instance will be generated from this subnet.
	// In the case of purchasing a `INTERNAL` clb instance, the subnet id must be specified. The VIP of the `INTERNAL` clb instance will be generated from this subnet.
	// +crossplane:generate:reference:type=github.com/crossplane-contrib/provider-tencentcloud/apis/vpc/v1alpha1.Subnet
	// +kubebuilder:validation:Optional
	SubnetID *string `json:"subnetId,omitempty" tf:"subnet_id,omitempty"`

	// Reference to a Subnet in vpc to populate subnetId.
	// +kubebuilder:validation:Optional
	SubnetIDRef *v1.Reference `json:"subnetIdRef,omitempty" tf:"-"`

	// Selector for a Subnet in vpc to populate subnetId.
	// +kubebuilder:validation:Optional
	SubnetIDSelector *v1.Selector `json:"subnetIdSelector,omitempty" tf:"-"`

	// The available tags within this CLB.
	// The available tags within this CLB.
	// +kubebuilder:validation:Optional
	// +mapType=granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// Region information of backend services are attached the CLB instance. Only supports OPEN CLBs.
	// Region information of backend services are attached the CLB instance. Only supports `OPEN` CLBs.
	// +kubebuilder:validation:Optional
	TargetRegionInfoRegion *string `json:"targetRegionInfoRegion,omitempty" tf:"target_region_info_region,omitempty"`

	// Vpc information of backend services are attached the CLB instance. Only supports OPEN CLBs.
	// Vpc information of backend services are attached the CLB instance. Only supports `OPEN` CLBs.
	// +kubebuilder:validation:Optional
	TargetRegionInfoVPCID *string `json:"targetRegionInfoVpcId,omitempty" tf:"target_region_info_vpc_id,omitempty"`

	// VPC ID of the CLB.
	// VPC ID of the CLB.
	// +crossplane:generate:reference:type=github.com/crossplane-contrib/provider-tencentcloud/apis/vpc/v1alpha1.VPC
	// +kubebuilder:validation:Optional
	VPCID *string `json:"vpcId,omitempty" tf:"vpc_id,omitempty"`

	// Reference to a VPC in vpc to populate vpcId.
	// +kubebuilder:validation:Optional
	VPCIDRef *v1.Reference `json:"vpcIdRef,omitempty" tf:"-"`

	// Selector for a VPC in vpc to populate vpcId.
	// +kubebuilder:validation:Optional
	VPCIDSelector *v1.Selector `json:"vpcIdSelector,omitempty" tf:"-"`

	// Specifies the VIP for the application of a CLB instance. This parameter is optional. If you do not specify this parameter, the system automatically assigns a value for the parameter. IPv4 and IPv6 CLB instances support this parameter, but IPv6 NAT64 CLB instances do not.
	// Specifies the VIP for the application of a CLB instance. This parameter is optional. If you do not specify this parameter, the system automatically assigns a value for the parameter. IPv4 and IPv6 CLB instances support this parameter, but IPv6 NAT64 CLB instances do not.
	// +kubebuilder:validation:Optional
	Vip *string `json:"vip,omitempty" tf:"vip,omitempty"`

	// Network operator, only applicable to open CLB. Valid values are CMCC(China Mobile), CTCC(Telecom), CUCC(China Unicom) and BGP. If this ISP is specified, network billing method can only use the bandwidth package billing (BANDWIDTH_PACKAGE).
	// Network operator, only applicable to open CLB. Valid values are `CMCC`(China Mobile), `CTCC`(Telecom), `CUCC`(China Unicom) and `BGP`. If this ISP is specified, network billing method can only use the bandwidth package billing (BANDWIDTH_PACKAGE).
	// +kubebuilder:validation:Optional
	VipIsp *string `json:"vipIsp,omitempty" tf:"vip_isp,omitempty"`

	// Available zone id, only applicable to open CLB.
	// Available zone id, only applicable to open CLB.
	// +kubebuilder:validation:Optional
	ZoneID *string `json:"zoneId,omitempty" tf:"zone_id,omitempty"`
}

type SnatIpsInitParameters struct {

	// Snat IP address, If set to empty will auto allocated.
	// Snat IP address, If set to empty will auto allocated.
	IP *string `json:"ip,omitempty" tf:"ip,omitempty"`

	// In the case of purchasing a INTERNAL clb instance, the subnet id must be specified. The VIP of the INTERNAL clb instance will be generated from this subnet.
	// Snat subnet ID.
	SubnetID *string `json:"subnetId,omitempty" tf:"subnet_id,omitempty"`
}

type SnatIpsObservation struct {

	// Snat IP address, If set to empty will auto allocated.
	// Snat IP address, If set to empty will auto allocated.
	IP *string `json:"ip,omitempty" tf:"ip,omitempty"`

	// In the case of purchasing a INTERNAL clb instance, the subnet id must be specified. The VIP of the INTERNAL clb instance will be generated from this subnet.
	// Snat subnet ID.
	SubnetID *string `json:"subnetId,omitempty" tf:"subnet_id,omitempty"`
}

type SnatIpsParameters struct {

	// Snat IP address, If set to empty will auto allocated.
	// Snat IP address, If set to empty will auto allocated.
	// +kubebuilder:validation:Optional
	IP *string `json:"ip,omitempty" tf:"ip,omitempty"`

	// In the case of purchasing a INTERNAL clb instance, the subnet id must be specified. The VIP of the INTERNAL clb instance will be generated from this subnet.
	// Snat subnet ID.
	// +kubebuilder:validation:Optional
	SubnetID *string `json:"subnetId" tf:"subnet_id,omitempty"`
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

// Instance is the Schema for the Instances API. Provides a resource to create a CLB instance.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,tencentcloud}
type Instance struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.clbName) || (has(self.initProvider) && has(self.initProvider.clbName))",message="spec.forProvider.clbName is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.networkType) || (has(self.initProvider) && has(self.initProvider.networkType))",message="spec.forProvider.networkType is a required parameter"
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
