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

type CcBlackWhiteIpsInitParameters struct {

	// Blacklist and whitelist IP addresses.
	// Blacklist and whitelist IP addresses.
	BlackWhiteIP *string `json:"blackWhiteIp,omitempty" tf:"black_white_ip,omitempty"`

	// Create time.
	// Create time.
	CreateTime *string `json:"createTime,omitempty" tf:"create_time,omitempty"`

	// Domain.
	// Domain.
	Domain *string `json:"domain,omitempty" tf:"domain,omitempty"`

	// Modify time.
	// Modify time.
	ModifyTime *string `json:"modifyTime,omitempty" tf:"modify_time,omitempty"`

	// Protocol.
	// Protocol.
	Protocol *string `json:"protocol,omitempty" tf:"protocol,omitempty"`

	// IP type, value [black(blacklist IP), white (whitelist IP)].
	// IP type, value [black(blacklist IP), white (whitelist IP)].
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

type CcBlackWhiteIpsObservation struct {

	// Blacklist and whitelist IP addresses.
	// Blacklist and whitelist IP addresses.
	BlackWhiteIP *string `json:"blackWhiteIp,omitempty" tf:"black_white_ip,omitempty"`

	// Create time.
	// Create time.
	CreateTime *string `json:"createTime,omitempty" tf:"create_time,omitempty"`

	// Domain.
	// Domain.
	Domain *string `json:"domain,omitempty" tf:"domain,omitempty"`

	// Modify time.
	// Modify time.
	ModifyTime *string `json:"modifyTime,omitempty" tf:"modify_time,omitempty"`

	// Protocol.
	// Protocol.
	Protocol *string `json:"protocol,omitempty" tf:"protocol,omitempty"`

	// IP type, value [black(blacklist IP), white (whitelist IP)].
	// IP type, value [black(blacklist IP), white (whitelist IP)].
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

type CcBlackWhiteIpsParameters struct {

	// Blacklist and whitelist IP addresses.
	// Blacklist and whitelist IP addresses.
	// +kubebuilder:validation:Optional
	BlackWhiteIP *string `json:"blackWhiteIp" tf:"black_white_ip,omitempty"`

	// Create time.
	// Create time.
	// +kubebuilder:validation:Optional
	CreateTime *string `json:"createTime,omitempty" tf:"create_time,omitempty"`

	// Domain.
	// Domain.
	// +kubebuilder:validation:Optional
	Domain *string `json:"domain" tf:"domain,omitempty"`

	// Modify time.
	// Modify time.
	// +kubebuilder:validation:Optional
	ModifyTime *string `json:"modifyTime,omitempty" tf:"modify_time,omitempty"`

	// Protocol.
	// Protocol.
	// +kubebuilder:validation:Optional
	Protocol *string `json:"protocol" tf:"protocol,omitempty"`

	// IP type, value [black(blacklist IP), white (whitelist IP)].
	// IP type, value [black(blacklist IP), white (whitelist IP)].
	// +kubebuilder:validation:Optional
	Type *string `json:"type" tf:"type,omitempty"`
}

type CcGeoIPPolicysInitParameters struct {

	// User action, drop or arg.
	// User action, drop or arg.
	Action *string `json:"action,omitempty" tf:"action,omitempty"`

	// The list of region IDs that the user selects to block.
	// The list of region IDs that the user selects to block.
	AreaList []*float64 `json:"areaList,omitempty" tf:"area_list,omitempty"`

	// Create time.
	// Create time.
	CreateTime *string `json:"createTime,omitempty" tf:"create_time,omitempty"`

	// Domain.
	// domain.
	Domain *string `json:"domain,omitempty" tf:"domain,omitempty"`

	// Modify time.
	// Modify time.
	ModifyTime *string `json:"modifyTime,omitempty" tf:"modify_time,omitempty"`

	// Protocol.
	// Protocol, preferably HTTP, HTTPS.
	Protocol *string `json:"protocol,omitempty" tf:"protocol,omitempty"`

	// Regional types, divided into china, oversea and customized.
	// Regional types, divided into china, oversea and customized.
	RegionType *string `json:"regionType,omitempty" tf:"region_type,omitempty"`
}

type CcGeoIPPolicysObservation struct {

	// User action, drop or arg.
	// User action, drop or arg.
	Action *string `json:"action,omitempty" tf:"action,omitempty"`

	// The list of region IDs that the user selects to block.
	// The list of region IDs that the user selects to block.
	AreaList []*float64 `json:"areaList,omitempty" tf:"area_list,omitempty"`

	// Create time.
	// Create time.
	CreateTime *string `json:"createTime,omitempty" tf:"create_time,omitempty"`

	// Domain.
	// domain.
	Domain *string `json:"domain,omitempty" tf:"domain,omitempty"`

	// Modify time.
	// Modify time.
	ModifyTime *string `json:"modifyTime,omitempty" tf:"modify_time,omitempty"`

	// Protocol.
	// Protocol, preferably HTTP, HTTPS.
	Protocol *string `json:"protocol,omitempty" tf:"protocol,omitempty"`

	// Regional types, divided into china, oversea and customized.
	// Regional types, divided into china, oversea and customized.
	RegionType *string `json:"regionType,omitempty" tf:"region_type,omitempty"`
}

type CcGeoIPPolicysParameters struct {

	// User action, drop or arg.
	// User action, drop or arg.
	// +kubebuilder:validation:Optional
	Action *string `json:"action" tf:"action,omitempty"`

	// The list of region IDs that the user selects to block.
	// The list of region IDs that the user selects to block.
	// +kubebuilder:validation:Optional
	AreaList []*float64 `json:"areaList,omitempty" tf:"area_list,omitempty"`

	// Create time.
	// Create time.
	// +kubebuilder:validation:Optional
	CreateTime *string `json:"createTime,omitempty" tf:"create_time,omitempty"`

	// Domain.
	// domain.
	// +kubebuilder:validation:Optional
	Domain *string `json:"domain" tf:"domain,omitempty"`

	// Modify time.
	// Modify time.
	// +kubebuilder:validation:Optional
	ModifyTime *string `json:"modifyTime,omitempty" tf:"modify_time,omitempty"`

	// Protocol.
	// Protocol, preferably HTTP, HTTPS.
	// +kubebuilder:validation:Optional
	Protocol *string `json:"protocol" tf:"protocol,omitempty"`

	// Regional types, divided into china, oversea and customized.
	// Regional types, divided into china, oversea and customized.
	// +kubebuilder:validation:Optional
	RegionType *string `json:"regionType" tf:"region_type,omitempty"`
}

type CcPolicyV2InitParameters struct {

	// Business of resource instance. bgpip indicates anti-anti-ip ip; bgp means exclusive package; bgp-multip means shared packet; net indicates anti-anti-ip pro version.
	// Business of resource instance. bgpip indicates anti-anti-ip ip; bgp means exclusive package; bgp-multip means shared packet; net indicates anti-anti-ip pro version.
	Business *string `json:"business,omitempty" tf:"business,omitempty"`

	// Blacklist and whitelist.
	// Blacklist and whitelist.
	CcBlackWhiteIps []CcBlackWhiteIpsInitParameters `json:"ccBlackWhiteIps,omitempty" tf:"cc_black_white_ips,omitempty"`

	// Details of the CC region blocking policy list.
	// Details of the CC region blocking policy list.
	CcGeoIPPolicys []CcGeoIPPolicysInitParameters `json:"ccGeoIpPolicys,omitempty" tf:"cc_geo_ip_policys,omitempty"`

	// CC Precision Protection List.
	// CC Precision Protection List.
	CcPrecisionPolicys []CcPrecisionPolicysInitParameters `json:"ccPrecisionPolicys,omitempty" tf:"cc_precision_policys,omitempty"`

	// CC frequency throttling policy.
	// CC frequency throttling policy.
	CcPrecisionReqLimits []CcPrecisionReqLimitsInitParameters `json:"ccPrecisionReqLimits,omitempty" tf:"cc_precision_req_limits,omitempty"`

	// The ID of the resource instance.
	// The ID of the resource instance.
	ResourceID *string `json:"resourceId,omitempty" tf:"resource_id,omitempty"`

	// List of protection threshold configurations.
	// List of protection threshold configurations.
	Thresholds []ThresholdsInitParameters `json:"thresholds,omitempty" tf:"thresholds,omitempty"`
}

type CcPolicyV2Observation struct {

	// Business of resource instance. bgpip indicates anti-anti-ip ip; bgp means exclusive package; bgp-multip means shared packet; net indicates anti-anti-ip pro version.
	// Business of resource instance. bgpip indicates anti-anti-ip ip; bgp means exclusive package; bgp-multip means shared packet; net indicates anti-anti-ip pro version.
	Business *string `json:"business,omitempty" tf:"business,omitempty"`

	// Blacklist and whitelist.
	// Blacklist and whitelist.
	CcBlackWhiteIps []CcBlackWhiteIpsObservation `json:"ccBlackWhiteIps,omitempty" tf:"cc_black_white_ips,omitempty"`

	// Details of the CC region blocking policy list.
	// Details of the CC region blocking policy list.
	CcGeoIPPolicys []CcGeoIPPolicysObservation `json:"ccGeoIpPolicys,omitempty" tf:"cc_geo_ip_policys,omitempty"`

	// CC Precision Protection List.
	// CC Precision Protection List.
	CcPrecisionPolicys []CcPrecisionPolicysObservation `json:"ccPrecisionPolicys,omitempty" tf:"cc_precision_policys,omitempty"`

	// CC frequency throttling policy.
	// CC frequency throttling policy.
	CcPrecisionReqLimits []CcPrecisionReqLimitsObservation `json:"ccPrecisionReqLimits,omitempty" tf:"cc_precision_req_limits,omitempty"`

	// ID of the resource.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The ID of the resource instance.
	// The ID of the resource instance.
	ResourceID *string `json:"resourceId,omitempty" tf:"resource_id,omitempty"`

	// List of protection threshold configurations.
	// List of protection threshold configurations.
	Thresholds []ThresholdsObservation `json:"thresholds,omitempty" tf:"thresholds,omitempty"`
}

type CcPolicyV2Parameters struct {

	// Business of resource instance. bgpip indicates anti-anti-ip ip; bgp means exclusive package; bgp-multip means shared packet; net indicates anti-anti-ip pro version.
	// Business of resource instance. bgpip indicates anti-anti-ip ip; bgp means exclusive package; bgp-multip means shared packet; net indicates anti-anti-ip pro version.
	// +kubebuilder:validation:Optional
	Business *string `json:"business,omitempty" tf:"business,omitempty"`

	// Blacklist and whitelist.
	// Blacklist and whitelist.
	// +kubebuilder:validation:Optional
	CcBlackWhiteIps []CcBlackWhiteIpsParameters `json:"ccBlackWhiteIps,omitempty" tf:"cc_black_white_ips,omitempty"`

	// Details of the CC region blocking policy list.
	// Details of the CC region blocking policy list.
	// +kubebuilder:validation:Optional
	CcGeoIPPolicys []CcGeoIPPolicysParameters `json:"ccGeoIpPolicys,omitempty" tf:"cc_geo_ip_policys,omitempty"`

	// CC Precision Protection List.
	// CC Precision Protection List.
	// +kubebuilder:validation:Optional
	CcPrecisionPolicys []CcPrecisionPolicysParameters `json:"ccPrecisionPolicys,omitempty" tf:"cc_precision_policys,omitempty"`

	// CC frequency throttling policy.
	// CC frequency throttling policy.
	// +kubebuilder:validation:Optional
	CcPrecisionReqLimits []CcPrecisionReqLimitsParameters `json:"ccPrecisionReqLimits,omitempty" tf:"cc_precision_req_limits,omitempty"`

	// The ID of the resource instance.
	// The ID of the resource instance.
	// +kubebuilder:validation:Optional
	ResourceID *string `json:"resourceId,omitempty" tf:"resource_id,omitempty"`

	// List of protection threshold configurations.
	// List of protection threshold configurations.
	// +kubebuilder:validation:Optional
	Thresholds []ThresholdsParameters `json:"thresholds,omitempty" tf:"thresholds,omitempty"`
}

type CcPrecisionPolicysInitParameters struct {

	// Domain.
	// Domain.
	Domain *string `json:"domain,omitempty" tf:"domain,omitempty"`

	// Ip address.
	// Ip address.
	IP *string `json:"ip,omitempty" tf:"ip,omitempty"`

	// Policy mode (discard or captcha).
	// Policy mode (discard or captcha).
	PolicyAction *string `json:"policyAction,omitempty" tf:"policy_action,omitempty"`

	// A list of policies.
	// A list of policies.
	Policys []PolicysInitParameters `json:"policys,omitempty" tf:"policys,omitempty"`

	// Protocol.
	// Protocol.
	Protocol *string `json:"protocol,omitempty" tf:"protocol,omitempty"`
}

type CcPrecisionPolicysObservation struct {

	// Domain.
	// Domain.
	Domain *string `json:"domain,omitempty" tf:"domain,omitempty"`

	// Ip address.
	// Ip address.
	IP *string `json:"ip,omitempty" tf:"ip,omitempty"`

	// Policy mode (discard or captcha).
	// Policy mode (discard or captcha).
	PolicyAction *string `json:"policyAction,omitempty" tf:"policy_action,omitempty"`

	// ID of the resource.
	// Policy Id.
	PolicyID *string `json:"policyId,omitempty" tf:"policy_id,omitempty"`

	// A list of policies.
	// A list of policies.
	Policys []PolicysObservation `json:"policys,omitempty" tf:"policys,omitempty"`

	// Protocol.
	// Protocol.
	Protocol *string `json:"protocol,omitempty" tf:"protocol,omitempty"`
}

type CcPrecisionPolicysParameters struct {

	// Domain.
	// Domain.
	// +kubebuilder:validation:Optional
	Domain *string `json:"domain" tf:"domain,omitempty"`

	// Ip address.
	// Ip address.
	// +kubebuilder:validation:Optional
	IP *string `json:"ip" tf:"ip,omitempty"`

	// Policy mode (discard or captcha).
	// Policy mode (discard or captcha).
	// +kubebuilder:validation:Optional
	PolicyAction *string `json:"policyAction" tf:"policy_action,omitempty"`

	// A list of policies.
	// A list of policies.
	// +kubebuilder:validation:Optional
	Policys []PolicysParameters `json:"policys" tf:"policys,omitempty"`

	// Protocol.
	// Protocol.
	// +kubebuilder:validation:Optional
	Protocol *string `json:"protocol" tf:"protocol,omitempty"`
}

type CcPrecisionReqLimitsInitParameters struct {

	// Domain.
	// Domain.
	Domain *string `json:"domain,omitempty" tf:"domain,omitempty"`

	// Protection rating, the optional value of default means default policy, loose means loose, and strict means strict.
	// Protection rating, the optional value of default means default policy, loose means loose, and strict means strict.
	Level *string `json:"level,omitempty" tf:"level,omitempty"`

	// A list of policies.
	// The CC Frequency Limit Policy Item field.
	Policys []CcPrecisionReqLimitsPolicysInitParameters `json:"policys,omitempty" tf:"policys,omitempty"`

	// Protocol.
	// Protocol, preferably HTTP, HTTPS.
	Protocol *string `json:"protocol,omitempty" tf:"protocol,omitempty"`
}

type CcPrecisionReqLimitsObservation struct {

	// Domain.
	// Domain.
	Domain *string `json:"domain,omitempty" tf:"domain,omitempty"`

	// Ip address.
	// IP address.
	IP *string `json:"ip,omitempty" tf:"ip,omitempty"`

	// ID of the resource.
	// Instance id.
	InstanceID *string `json:"instanceId,omitempty" tf:"instance_id,omitempty"`

	// Protection rating, the optional value of default means default policy, loose means loose, and strict means strict.
	// Protection rating, the optional value of default means default policy, loose means loose, and strict means strict.
	Level *string `json:"level,omitempty" tf:"level,omitempty"`

	// A list of policies.
	// The CC Frequency Limit Policy Item field.
	Policys []CcPrecisionReqLimitsPolicysObservation `json:"policys,omitempty" tf:"policys,omitempty"`

	// Protocol.
	// Protocol, preferably HTTP, HTTPS.
	Protocol *string `json:"protocol,omitempty" tf:"protocol,omitempty"`
}

type CcPrecisionReqLimitsParameters struct {

	// Domain.
	// Domain.
	// +kubebuilder:validation:Optional
	Domain *string `json:"domain" tf:"domain,omitempty"`

	// Protection rating, the optional value of default means default policy, loose means loose, and strict means strict.
	// Protection rating, the optional value of default means default policy, loose means loose, and strict means strict.
	// +kubebuilder:validation:Optional
	Level *string `json:"level" tf:"level,omitempty"`

	// A list of policies.
	// The CC Frequency Limit Policy Item field.
	// +kubebuilder:validation:Optional
	Policys []CcPrecisionReqLimitsPolicysParameters `json:"policys" tf:"policys,omitempty"`

	// Protocol.
	// Protocol, preferably HTTP, HTTPS.
	// +kubebuilder:validation:Optional
	Protocol *string `json:"protocol" tf:"protocol,omitempty"`
}

type CcPrecisionReqLimitsPolicysInitParameters struct {

	// User action, drop or arg.
	// The frequency limit policy mode, the optional value of arg indicates the verification code, and drop indicates the discard.
	Action *string `json:"action,omitempty" tf:"action,omitempty"`

	// Cookies, one of the three policy entries can only be filled in.
	// Cookies, one of the three policy entries can only be filled in.
	Cookie *string `json:"cookie,omitempty" tf:"cookie,omitempty"`

	// The duration of the frequency limit policy can be taken from 1 to 86400 per second.
	// The duration of the frequency limit policy can be taken from 1 to 86400 per second.
	ExecuteDuration *float64 `json:"executeDuration,omitempty" tf:"execute_duration,omitempty"`

	// The policy item is compared, and the optional value include indicates inclusion, and equal means equal.
	// The policy item is compared, and the optional value include indicates inclusion, and equal means equal.
	Mode *string `json:"mode,omitempty" tf:"mode,omitempty"`

	// Statistical period, take values 1, 10, 30, 60, in seconds.
	// Statistical period, take values 1, 10, 30, 60, in seconds.
	Period *float64 `json:"period,omitempty" tf:"period,omitempty"`

	// The number of requests, the value is 1 to 20000.
	// The number of requests, the value is 1 to 20000.
	RequestNum *float64 `json:"requestNum,omitempty" tf:"request_num,omitempty"`

	// Uri, one of the three policy entries can only be filled in.
	// Uri, one of the three policy entries can only be filled in.
	URI *string `json:"uri,omitempty" tf:"uri,omitempty"`

	// User-Agent, only one of the three policy entries can be filled in.
	// User-Agent, only one of the three policy entries can be filled in.
	UserAgent *string `json:"userAgent,omitempty" tf:"user_agent,omitempty"`
}

type CcPrecisionReqLimitsPolicysObservation struct {

	// User action, drop or arg.
	// The frequency limit policy mode, the optional value of arg indicates the verification code, and drop indicates the discard.
	Action *string `json:"action,omitempty" tf:"action,omitempty"`

	// Cookies, one of the three policy entries can only be filled in.
	// Cookies, one of the three policy entries can only be filled in.
	Cookie *string `json:"cookie,omitempty" tf:"cookie,omitempty"`

	// The duration of the frequency limit policy can be taken from 1 to 86400 per second.
	// The duration of the frequency limit policy can be taken from 1 to 86400 per second.
	ExecuteDuration *float64 `json:"executeDuration,omitempty" tf:"execute_duration,omitempty"`

	// The policy item is compared, and the optional value include indicates inclusion, and equal means equal.
	// The policy item is compared, and the optional value include indicates inclusion, and equal means equal.
	Mode *string `json:"mode,omitempty" tf:"mode,omitempty"`

	// Statistical period, take values 1, 10, 30, 60, in seconds.
	// Statistical period, take values 1, 10, 30, 60, in seconds.
	Period *float64 `json:"period,omitempty" tf:"period,omitempty"`

	// The number of requests, the value is 1 to 20000.
	// The number of requests, the value is 1 to 20000.
	RequestNum *float64 `json:"requestNum,omitempty" tf:"request_num,omitempty"`

	// Uri, one of the three policy entries can only be filled in.
	// Uri, one of the three policy entries can only be filled in.
	URI *string `json:"uri,omitempty" tf:"uri,omitempty"`

	// User-Agent, only one of the three policy entries can be filled in.
	// User-Agent, only one of the three policy entries can be filled in.
	UserAgent *string `json:"userAgent,omitempty" tf:"user_agent,omitempty"`
}

type CcPrecisionReqLimitsPolicysParameters struct {

	// User action, drop or arg.
	// The frequency limit policy mode, the optional value of arg indicates the verification code, and drop indicates the discard.
	// +kubebuilder:validation:Optional
	Action *string `json:"action" tf:"action,omitempty"`

	// Cookies, one of the three policy entries can only be filled in.
	// Cookies, one of the three policy entries can only be filled in.
	// +kubebuilder:validation:Optional
	Cookie *string `json:"cookie,omitempty" tf:"cookie,omitempty"`

	// The duration of the frequency limit policy can be taken from 1 to 86400 per second.
	// The duration of the frequency limit policy can be taken from 1 to 86400 per second.
	// +kubebuilder:validation:Optional
	ExecuteDuration *float64 `json:"executeDuration" tf:"execute_duration,omitempty"`

	// The policy item is compared, and the optional value include indicates inclusion, and equal means equal.
	// The policy item is compared, and the optional value include indicates inclusion, and equal means equal.
	// +kubebuilder:validation:Optional
	Mode *string `json:"mode" tf:"mode,omitempty"`

	// Statistical period, take values 1, 10, 30, 60, in seconds.
	// Statistical period, take values 1, 10, 30, 60, in seconds.
	// +kubebuilder:validation:Optional
	Period *float64 `json:"period" tf:"period,omitempty"`

	// The number of requests, the value is 1 to 20000.
	// The number of requests, the value is 1 to 20000.
	// +kubebuilder:validation:Optional
	RequestNum *float64 `json:"requestNum" tf:"request_num,omitempty"`

	// Uri, one of the three policy entries can only be filled in.
	// Uri, one of the three policy entries can only be filled in.
	// +kubebuilder:validation:Optional
	URI *string `json:"uri,omitempty" tf:"uri,omitempty"`

	// User-Agent, only one of the three policy entries can be filled in.
	// User-Agent, only one of the three policy entries can be filled in.
	// +kubebuilder:validation:Optional
	UserAgent *string `json:"userAgent,omitempty" tf:"user_agent,omitempty"`
}

type PolicysInitParameters struct {

	// Configuration item types, currently only support value.
	// Configuration item types, currently only support value.
	FieldName *string `json:"fieldName,omitempty" tf:"field_name,omitempty"`

	// Configuration fields with the desirable values cgi, ua, cookie, referer, accept, srcip.
	// Configuration fields with the desirable values cgi, ua, cookie, referer, accept, srcip.
	FieldType *string `json:"fieldType,omitempty" tf:"field_type,omitempty"`

	// Configure the value.
	// Configure the value.
	Value *string `json:"value,omitempty" tf:"value,omitempty"`

	// Configure the item-value comparison mode, which can be taken as the value of evaluate, not_equal, include.
	// Configure the item-value comparison mode, which can be taken as the value of evaluate, not_equal, include.
	ValueOperator *string `json:"valueOperator,omitempty" tf:"value_operator,omitempty"`
}

type PolicysObservation struct {

	// Configuration item types, currently only support value.
	// Configuration item types, currently only support value.
	FieldName *string `json:"fieldName,omitempty" tf:"field_name,omitempty"`

	// Configuration fields with the desirable values cgi, ua, cookie, referer, accept, srcip.
	// Configuration fields with the desirable values cgi, ua, cookie, referer, accept, srcip.
	FieldType *string `json:"fieldType,omitempty" tf:"field_type,omitempty"`

	// Configure the value.
	// Configure the value.
	Value *string `json:"value,omitempty" tf:"value,omitempty"`

	// Configure the item-value comparison mode, which can be taken as the value of evaluate, not_equal, include.
	// Configure the item-value comparison mode, which can be taken as the value of evaluate, not_equal, include.
	ValueOperator *string `json:"valueOperator,omitempty" tf:"value_operator,omitempty"`
}

type PolicysParameters struct {

	// Configuration item types, currently only support value.
	// Configuration item types, currently only support value.
	// +kubebuilder:validation:Optional
	FieldName *string `json:"fieldName" tf:"field_name,omitempty"`

	// Configuration fields with the desirable values cgi, ua, cookie, referer, accept, srcip.
	// Configuration fields with the desirable values cgi, ua, cookie, referer, accept, srcip.
	// +kubebuilder:validation:Optional
	FieldType *string `json:"fieldType" tf:"field_type,omitempty"`

	// Configure the value.
	// Configure the value.
	// +kubebuilder:validation:Optional
	Value *string `json:"value" tf:"value,omitempty"`

	// Configure the item-value comparison mode, which can be taken as the value of evaluate, not_equal, include.
	// Configure the item-value comparison mode, which can be taken as the value of evaluate, not_equal, include.
	// +kubebuilder:validation:Optional
	ValueOperator *string `json:"valueOperator" tf:"value_operator,omitempty"`
}

type ThresholdsInitParameters struct {

	// Domain.
	// domain.
	Domain *string `json:"domain,omitempty" tf:"domain,omitempty"`

	// Cleaning threshold, -1 indicates that the default mode is turned on.
	// Cleaning threshold, -1 indicates that the `default` mode is turned on.
	Threshold *float64 `json:"threshold,omitempty" tf:"threshold,omitempty"`
}

type ThresholdsObservation struct {

	// Domain.
	// domain.
	Domain *string `json:"domain,omitempty" tf:"domain,omitempty"`

	// Cleaning threshold, -1 indicates that the default mode is turned on.
	// Cleaning threshold, -1 indicates that the `default` mode is turned on.
	Threshold *float64 `json:"threshold,omitempty" tf:"threshold,omitempty"`
}

type ThresholdsParameters struct {

	// Domain.
	// domain.
	// +kubebuilder:validation:Optional
	Domain *string `json:"domain" tf:"domain,omitempty"`

	// Cleaning threshold, -1 indicates that the default mode is turned on.
	// Cleaning threshold, -1 indicates that the `default` mode is turned on.
	// +kubebuilder:validation:Optional
	Threshold *float64 `json:"threshold" tf:"threshold,omitempty"`
}

// CcPolicyV2Spec defines the desired state of CcPolicyV2
type CcPolicyV2Spec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     CcPolicyV2Parameters `json:"forProvider"`
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
	InitProvider CcPolicyV2InitParameters `json:"initProvider,omitempty"`
}

// CcPolicyV2Status defines the observed state of CcPolicyV2.
type CcPolicyV2Status struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        CcPolicyV2Observation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// CcPolicyV2 is the Schema for the CcPolicyV2s API. Use this resource to create a dayu CC policy
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,tencentcloud}
type CcPolicyV2 struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.business) || (has(self.initProvider) && has(self.initProvider.business))",message="spec.forProvider.business is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.resourceId) || (has(self.initProvider) && has(self.initProvider.resourceId))",message="spec.forProvider.resourceId is a required parameter"
	Spec   CcPolicyV2Spec   `json:"spec"`
	Status CcPolicyV2Status `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// CcPolicyV2List contains a list of CcPolicyV2s
type CcPolicyV2List struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []CcPolicyV2 `json:"items"`
}

// Repository type metadata.
var (
	CcPolicyV2_Kind             = "CcPolicyV2"
	CcPolicyV2_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: CcPolicyV2_Kind}.String()
	CcPolicyV2_KindAPIVersion   = CcPolicyV2_Kind + "." + CRDGroupVersion.String()
	CcPolicyV2_GroupVersionKind = CRDGroupVersion.WithKind(CcPolicyV2_Kind)
)

func init() {
	SchemeBuilder.Register(&CcPolicyV2{}, &CcPolicyV2List{})
}
