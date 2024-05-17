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

type L7RuleInitParameters struct {

	// Domain that the layer 7 rule works for. Valid string length ranges from 0 to 80.
	// Domain that the layer 7 rule works for. Valid string length ranges from 0 to 80.
	Domain *string `json:"domain,omitempty" tf:"domain,omitempty"`

	// HTTP Status Code. The default is 26. Valid value ranges: [1~31]. 1 means the return value '1xx' is health. 2 means the return value '2xx' is health. 4 means the return value '3xx' is health. 8 means the return value '4xx' is health. 16 means the return value '5xx' is health. If you want multiple return codes to indicate health, need to add the corresponding values.
	// HTTP Status Code. The default is `26`. Valid value ranges: [1~31]. `1` means the return value '1xx' is health. `2` means the return value '2xx' is health. `4` means the return value '3xx' is health. `8` means the return value '4xx' is health. `16` means the return value '5xx' is health. If you want multiple return codes to indicate health, need to add the corresponding values.
	HealthCheckCode *float64 `json:"healthCheckCode,omitempty" tf:"health_check_code,omitempty"`

	// Health threshold of health check, and the default is 3. If a success result is returned for the health check 3 consecutive times, indicates that the forwarding is normal. The value range is [2-10].
	// Health threshold of health check, and the default is `3`. If a success result is returned for the health check 3 consecutive times, indicates that the forwarding is normal. The value range is [2-10].
	HealthCheckHealthNum *float64 `json:"healthCheckHealthNum,omitempty" tf:"health_check_health_num,omitempty"`

	// Interval time of health check. Valid value ranges: [10~60]sec. The default is 15 sec.
	// Interval time of health check. Valid value ranges: [10~60]sec. The default is 15 sec.
	HealthCheckInterval *float64 `json:"healthCheckInterval,omitempty" tf:"health_check_interval,omitempty"`

	// Methods of health check. The default is 'HEAD', the available value are 'HEAD' and 'GET'.
	// Methods of health check. The default is 'HEAD', the available value are 'HEAD' and 'GET'.
	HealthCheckMethod *string `json:"healthCheckMethod,omitempty" tf:"health_check_method,omitempty"`

	// Path of health check. The default is /.
	// Path of health check. The default is `/`.
	HealthCheckPath *string `json:"healthCheckPath,omitempty" tf:"health_check_path,omitempty"`

	// Indicates whether health check is enabled. The default is false.
	// Indicates whether health check is enabled. The default is `false`.
	HealthCheckSwitch *bool `json:"healthCheckSwitch,omitempty" tf:"health_check_switch,omitempty"`

	// Unhealthy threshold of health check, and the default is 3. If the unhealthy result is returned 3 consecutive times, indicates that the forwarding is abnormal. The value range is [2-10].
	// Unhealthy threshold of health check, and the default is `3`. If the unhealthy result is returned 3 consecutive times, indicates that the forwarding is abnormal. The value range is [2-10].
	HealthCheckUnhealthNum *float64 `json:"healthCheckUnhealthNum,omitempty" tf:"health_check_unhealth_num,omitempty"`

	// Name of the rule.
	// Name of the rule.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Protocol of the rule. Valid values: http, https.
	// Protocol of the rule. Valid values: `http`, `https`.
	Protocol *string `json:"protocol,omitempty" tf:"protocol,omitempty"`

	// ID of the resource that the layer 7 rule works for.
	// ID of the resource that the layer 7 rule works for.
	ResourceID *string `json:"resourceId,omitempty" tf:"resource_id,omitempty"`

	// Type of the resource that the layer 7 rule works for, valid value is bgpip.
	// Type of the resource that the layer 7 rule works for, valid value is `bgpip`.
	ResourceType *string `json:"resourceType,omitempty" tf:"resource_type,omitempty"`

	// SSL ID, when the protocol is https, the field should be set with valid SSL id.
	// SSL ID, when the `protocol` is `https`, the field should be set with valid SSL id.
	SSLID *string `json:"sslId,omitempty" tf:"ssl_id,omitempty"`

	// Source list of the rule, it can be a set of ip sources or a set of domain sources. The number of items ranges from 1 to 16.
	// Source list of the rule, it can be a set of ip sources or a set of domain sources. The number of items ranges from 1 to 16.
	SourceList []*string `json:"sourceList,omitempty" tf:"source_list,omitempty"`

	// Source type, 1 for source of host, 2 for source of IP.
	// Source type, `1` for source of host, `2` for source of IP.
	SourceType *float64 `json:"sourceType,omitempty" tf:"source_type,omitempty"`

	// Indicate the rule will take effect or not.
	// Indicate the rule will take effect or not.
	Switch *bool `json:"switch,omitempty" tf:"switch,omitempty"`
}

type L7RuleObservation struct {

	// Domain that the layer 7 rule works for. Valid string length ranges from 0 to 80.
	// Domain that the layer 7 rule works for. Valid string length ranges from 0 to 80.
	Domain *string `json:"domain,omitempty" tf:"domain,omitempty"`

	// HTTP Status Code. The default is 26. Valid value ranges: [1~31]. 1 means the return value '1xx' is health. 2 means the return value '2xx' is health. 4 means the return value '3xx' is health. 8 means the return value '4xx' is health. 16 means the return value '5xx' is health. If you want multiple return codes to indicate health, need to add the corresponding values.
	// HTTP Status Code. The default is `26`. Valid value ranges: [1~31]. `1` means the return value '1xx' is health. `2` means the return value '2xx' is health. `4` means the return value '3xx' is health. `8` means the return value '4xx' is health. `16` means the return value '5xx' is health. If you want multiple return codes to indicate health, need to add the corresponding values.
	HealthCheckCode *float64 `json:"healthCheckCode,omitempty" tf:"health_check_code,omitempty"`

	// Health threshold of health check, and the default is 3. If a success result is returned for the health check 3 consecutive times, indicates that the forwarding is normal. The value range is [2-10].
	// Health threshold of health check, and the default is `3`. If a success result is returned for the health check 3 consecutive times, indicates that the forwarding is normal. The value range is [2-10].
	HealthCheckHealthNum *float64 `json:"healthCheckHealthNum,omitempty" tf:"health_check_health_num,omitempty"`

	// Interval time of health check. Valid value ranges: [10~60]sec. The default is 15 sec.
	// Interval time of health check. Valid value ranges: [10~60]sec. The default is 15 sec.
	HealthCheckInterval *float64 `json:"healthCheckInterval,omitempty" tf:"health_check_interval,omitempty"`

	// Methods of health check. The default is 'HEAD', the available value are 'HEAD' and 'GET'.
	// Methods of health check. The default is 'HEAD', the available value are 'HEAD' and 'GET'.
	HealthCheckMethod *string `json:"healthCheckMethod,omitempty" tf:"health_check_method,omitempty"`

	// Path of health check. The default is /.
	// Path of health check. The default is `/`.
	HealthCheckPath *string `json:"healthCheckPath,omitempty" tf:"health_check_path,omitempty"`

	// Indicates whether health check is enabled. The default is false.
	// Indicates whether health check is enabled. The default is `false`.
	HealthCheckSwitch *bool `json:"healthCheckSwitch,omitempty" tf:"health_check_switch,omitempty"`

	// Unhealthy threshold of health check, and the default is 3. If the unhealthy result is returned 3 consecutive times, indicates that the forwarding is abnormal. The value range is [2-10].
	// Unhealthy threshold of health check, and the default is `3`. If the unhealthy result is returned 3 consecutive times, indicates that the forwarding is abnormal. The value range is [2-10].
	HealthCheckUnhealthNum *float64 `json:"healthCheckUnhealthNum,omitempty" tf:"health_check_unhealth_num,omitempty"`

	// ID of the resource.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Name of the rule.
	// Name of the rule.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Protocol of the rule. Valid values: http, https.
	// Protocol of the rule. Valid values: `http`, `https`.
	Protocol *string `json:"protocol,omitempty" tf:"protocol,omitempty"`

	// ID of the resource that the layer 7 rule works for.
	// ID of the resource that the layer 7 rule works for.
	ResourceID *string `json:"resourceId,omitempty" tf:"resource_id,omitempty"`

	// Type of the resource that the layer 7 rule works for, valid value is bgpip.
	// Type of the resource that the layer 7 rule works for, valid value is `bgpip`.
	ResourceType *string `json:"resourceType,omitempty" tf:"resource_type,omitempty"`

	// ID of the layer 7 rule.
	// ID of the layer 7 rule.
	RuleID *string `json:"ruleId,omitempty" tf:"rule_id,omitempty"`

	// SSL ID, when the protocol is https, the field should be set with valid SSL id.
	// SSL ID, when the `protocol` is `https`, the field should be set with valid SSL id.
	SSLID *string `json:"sslId,omitempty" tf:"ssl_id,omitempty"`

	// Source list of the rule, it can be a set of ip sources or a set of domain sources. The number of items ranges from 1 to 16.
	// Source list of the rule, it can be a set of ip sources or a set of domain sources. The number of items ranges from 1 to 16.
	SourceList []*string `json:"sourceList,omitempty" tf:"source_list,omitempty"`

	// Source type, 1 for source of host, 2 for source of IP.
	// Source type, `1` for source of host, `2` for source of IP.
	SourceType *float64 `json:"sourceType,omitempty" tf:"source_type,omitempty"`

	// Status of the rule. 0 for create/modify success, 2 for create/modify fail, 3 for delete success, 5 for delete failed, 6 for waiting to be created/modified, 7 for waiting to be deleted and 8 for waiting to get SSL ID.
	// Status of the rule. `0` for create/modify success, `2` for create/modify fail, `3` for delete success, `5` for delete failed, `6` for waiting to be created/modified, `7` for waiting to be deleted and 8 for waiting to get SSL ID.
	Status *float64 `json:"status,omitempty" tf:"status,omitempty"`

	// Indicate the rule will take effect or not.
	// Indicate the rule will take effect or not.
	Switch *bool `json:"switch,omitempty" tf:"switch,omitempty"`
}

type L7RuleParameters struct {

	// Domain that the layer 7 rule works for. Valid string length ranges from 0 to 80.
	// Domain that the layer 7 rule works for. Valid string length ranges from 0 to 80.
	// +kubebuilder:validation:Optional
	Domain *string `json:"domain,omitempty" tf:"domain,omitempty"`

	// HTTP Status Code. The default is 26. Valid value ranges: [1~31]. 1 means the return value '1xx' is health. 2 means the return value '2xx' is health. 4 means the return value '3xx' is health. 8 means the return value '4xx' is health. 16 means the return value '5xx' is health. If you want multiple return codes to indicate health, need to add the corresponding values.
	// HTTP Status Code. The default is `26`. Valid value ranges: [1~31]. `1` means the return value '1xx' is health. `2` means the return value '2xx' is health. `4` means the return value '3xx' is health. `8` means the return value '4xx' is health. `16` means the return value '5xx' is health. If you want multiple return codes to indicate health, need to add the corresponding values.
	// +kubebuilder:validation:Optional
	HealthCheckCode *float64 `json:"healthCheckCode,omitempty" tf:"health_check_code,omitempty"`

	// Health threshold of health check, and the default is 3. If a success result is returned for the health check 3 consecutive times, indicates that the forwarding is normal. The value range is [2-10].
	// Health threshold of health check, and the default is `3`. If a success result is returned for the health check 3 consecutive times, indicates that the forwarding is normal. The value range is [2-10].
	// +kubebuilder:validation:Optional
	HealthCheckHealthNum *float64 `json:"healthCheckHealthNum,omitempty" tf:"health_check_health_num,omitempty"`

	// Interval time of health check. Valid value ranges: [10~60]sec. The default is 15 sec.
	// Interval time of health check. Valid value ranges: [10~60]sec. The default is 15 sec.
	// +kubebuilder:validation:Optional
	HealthCheckInterval *float64 `json:"healthCheckInterval,omitempty" tf:"health_check_interval,omitempty"`

	// Methods of health check. The default is 'HEAD', the available value are 'HEAD' and 'GET'.
	// Methods of health check. The default is 'HEAD', the available value are 'HEAD' and 'GET'.
	// +kubebuilder:validation:Optional
	HealthCheckMethod *string `json:"healthCheckMethod,omitempty" tf:"health_check_method,omitempty"`

	// Path of health check. The default is /.
	// Path of health check. The default is `/`.
	// +kubebuilder:validation:Optional
	HealthCheckPath *string `json:"healthCheckPath,omitempty" tf:"health_check_path,omitempty"`

	// Indicates whether health check is enabled. The default is false.
	// Indicates whether health check is enabled. The default is `false`.
	// +kubebuilder:validation:Optional
	HealthCheckSwitch *bool `json:"healthCheckSwitch,omitempty" tf:"health_check_switch,omitempty"`

	// Unhealthy threshold of health check, and the default is 3. If the unhealthy result is returned 3 consecutive times, indicates that the forwarding is abnormal. The value range is [2-10].
	// Unhealthy threshold of health check, and the default is `3`. If the unhealthy result is returned 3 consecutive times, indicates that the forwarding is abnormal. The value range is [2-10].
	// +kubebuilder:validation:Optional
	HealthCheckUnhealthNum *float64 `json:"healthCheckUnhealthNum,omitempty" tf:"health_check_unhealth_num,omitempty"`

	// Name of the rule.
	// Name of the rule.
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Protocol of the rule. Valid values: http, https.
	// Protocol of the rule. Valid values: `http`, `https`.
	// +kubebuilder:validation:Optional
	Protocol *string `json:"protocol,omitempty" tf:"protocol,omitempty"`

	// ID of the resource that the layer 7 rule works for.
	// ID of the resource that the layer 7 rule works for.
	// +kubebuilder:validation:Optional
	ResourceID *string `json:"resourceId,omitempty" tf:"resource_id,omitempty"`

	// Type of the resource that the layer 7 rule works for, valid value is bgpip.
	// Type of the resource that the layer 7 rule works for, valid value is `bgpip`.
	// +kubebuilder:validation:Optional
	ResourceType *string `json:"resourceType,omitempty" tf:"resource_type,omitempty"`

	// SSL ID, when the protocol is https, the field should be set with valid SSL id.
	// SSL ID, when the `protocol` is `https`, the field should be set with valid SSL id.
	// +kubebuilder:validation:Optional
	SSLID *string `json:"sslId,omitempty" tf:"ssl_id,omitempty"`

	// Source list of the rule, it can be a set of ip sources or a set of domain sources. The number of items ranges from 1 to 16.
	// Source list of the rule, it can be a set of ip sources or a set of domain sources. The number of items ranges from 1 to 16.
	// +kubebuilder:validation:Optional
	SourceList []*string `json:"sourceList,omitempty" tf:"source_list,omitempty"`

	// Source type, 1 for source of host, 2 for source of IP.
	// Source type, `1` for source of host, `2` for source of IP.
	// +kubebuilder:validation:Optional
	SourceType *float64 `json:"sourceType,omitempty" tf:"source_type,omitempty"`

	// Indicate the rule will take effect or not.
	// Indicate the rule will take effect or not.
	// +kubebuilder:validation:Optional
	Switch *bool `json:"switch,omitempty" tf:"switch,omitempty"`
}

// L7RuleSpec defines the desired state of L7Rule
type L7RuleSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     L7RuleParameters `json:"forProvider"`
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
	InitProvider L7RuleInitParameters `json:"initProvider,omitempty"`
}

// L7RuleStatus defines the observed state of L7Rule.
type L7RuleStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        L7RuleObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// L7Rule is the Schema for the L7Rules API. Use this resource to create dayu layer 7 rule
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,tencentcloud}
type L7Rule struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.domain) || (has(self.initProvider) && has(self.initProvider.domain))",message="spec.forProvider.domain is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.name) || (has(self.initProvider) && has(self.initProvider.name))",message="spec.forProvider.name is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.protocol) || (has(self.initProvider) && has(self.initProvider.protocol))",message="spec.forProvider.protocol is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.resourceId) || (has(self.initProvider) && has(self.initProvider.resourceId))",message="spec.forProvider.resourceId is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.resourceType) || (has(self.initProvider) && has(self.initProvider.resourceType))",message="spec.forProvider.resourceType is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.sourceList) || (has(self.initProvider) && has(self.initProvider.sourceList))",message="spec.forProvider.sourceList is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.sourceType) || (has(self.initProvider) && has(self.initProvider.sourceType))",message="spec.forProvider.sourceType is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.switch) || (has(self.initProvider) && has(self.initProvider.switch))",message="spec.forProvider.switch is a required parameter"
	Spec   L7RuleSpec   `json:"spec"`
	Status L7RuleStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// L7RuleList contains a list of L7Rules
type L7RuleList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []L7Rule `json:"items"`
}

// Repository type metadata.
var (
	L7Rule_Kind             = "L7Rule"
	L7Rule_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: L7Rule_Kind}.String()
	L7Rule_KindAPIVersion   = L7Rule_Kind + "." + CRDGroupVersion.String()
	L7Rule_GroupVersionKind = CRDGroupVersion.WithKind(L7Rule_Kind)
)

func init() {
	SchemeBuilder.Register(&L7Rule{}, &L7RuleList{})
}
