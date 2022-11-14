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

type BindingObjectsObservation struct {
	DimensionsJSON *string `json:"dimensionsJson,omitempty" tf:"dimensions_json,omitempty"`

	IsShielded *float64 `json:"isShielded,omitempty" tf:"is_shielded,omitempty"`

	Region *string `json:"region,omitempty" tf:"region,omitempty"`

	UniqueID *string `json:"uniqueId,omitempty" tf:"unique_id,omitempty"`
}

type BindingObjectsParameters struct {
}

type PolicyGroupConditionsObservation struct {
}

type PolicyGroupConditionsParameters struct {

	// Alarm sending cycle per second. <0 does not fire, `0` only fires once, and >0 fires every triggerTime second.
	// +kubebuilder:validation:Required
	AlarmNotifyPeriod *float64 `json:"alarmNotifyPeriod" tf:"alarm_notify_period,omitempty"`

	// Alarm sending convergence type. `0` continuous alarm, `1` index alarm.
	// +kubebuilder:validation:Required
	AlarmNotifyType *float64 `json:"alarmNotifyType" tf:"alarm_notify_type,omitempty"`

	// Data aggregation cycle (unit of second), if the metric has a default value can not be filled, refer to `data.tencentcloud_monitor_policy_conditions(period_keys)`.
	// +kubebuilder:validation:Optional
	CalcPeriod *float64 `json:"calcPeriod,omitempty" tf:"calc_period,omitempty"`

	// Compare type. Valid value ranges: [1~12]. `1` means more than, `2` means greater than or equal, `3` means less than, `4` means less than or equal to, `5` means equal, `6` means not equal, `7` means days rose, `8` means days fell, `9` means weeks rose, `10` means weeks fell, `11` means period rise, `12` means period fell, refer to `data.tencentcloud_monitor_policy_conditions(calc_type_keys)`.
	// +kubebuilder:validation:Optional
	CalcType *float64 `json:"calcType,omitempty" tf:"calc_type,omitempty"`

	// Threshold value, refer to `data.tencentcloud_monitor_policy_conditions(calc_value_*)`.
	// +kubebuilder:validation:Optional
	CalcValue *float64 `json:"calcValue,omitempty" tf:"calc_value,omitempty"`

	// The rule triggers an alert that lasts for several detection cycles, refer to `data.tencentcloud_monitor_policy_conditions(period_num_keys)`.
	// +kubebuilder:validation:Optional
	ContinuePeriod *float64 `json:"continuePeriod,omitempty" tf:"continue_period,omitempty"`

	// Id of the metric, refer to `data.tencentcloud_monitor_policy_conditions(metric_id)`.
	// +kubebuilder:validation:Required
	MetricID *float64 `json:"metricId" tf:"metric_id,omitempty"`
}

type PolicyGroupEventConditionsObservation struct {
}

type PolicyGroupEventConditionsParameters struct {

	// Alarm sending cycle per second. <0 does not fire, `0` only fires once, and >0 fires every triggerTime second.
	// +kubebuilder:validation:Required
	AlarmNotifyPeriod *float64 `json:"alarmNotifyPeriod" tf:"alarm_notify_period,omitempty"`

	// Alarm sending convergence type. `0` continuous alarm, `1` index alarm.
	// +kubebuilder:validation:Required
	AlarmNotifyType *float64 `json:"alarmNotifyType" tf:"alarm_notify_type,omitempty"`

	// The ID of this event metric, refer to `data.tencentcloud_monitor_policy_conditions(event_id).
	// +kubebuilder:validation:Required
	EventID *float64 `json:"eventId" tf:"event_id,omitempty"`
}

type PolicyGroupObservation struct {
	BindingObjects []BindingObjectsObservation `json:"bindingObjects,omitempty" tf:"binding_objects,omitempty"`

	DimensionGroup []*string `json:"dimensionGroup,omitempty" tf:"dimension_group,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	LastEditUin *string `json:"lastEditUin,omitempty" tf:"last_edit_uin,omitempty"`

	Receivers []PolicyGroupReceiversObservation `json:"receivers,omitempty" tf:"receivers,omitempty"`

	SupportRegions []*string `json:"supportRegions,omitempty" tf:"support_regions,omitempty"`

	UpdateTime *string `json:"updateTime,omitempty" tf:"update_time,omitempty"`
}

type PolicyGroupParameters struct {

	// A list of threshold rules. Each element contains the following attributes:
	// +kubebuilder:validation:Optional
	Conditions []PolicyGroupConditionsParameters `json:"conditions,omitempty" tf:"conditions,omitempty"`

	// A list of event rules. Each element contains the following attributes:
	// +kubebuilder:validation:Optional
	EventConditions []PolicyGroupEventConditionsParameters `json:"eventConditions,omitempty" tf:"event_conditions,omitempty"`

	// Policy group name, length should between 1 and 20.
	// +kubebuilder:validation:Required
	GroupName *string `json:"groupName" tf:"group_name,omitempty"`

	// The and or relation of indicator alarm rule. Valid values: `0`, `1`. `0` represents or rule (if any rule is met, the alarm will be raised), `1` represents and rule (if all rules are met, the alarm will be raised).The default is 0.
	// +kubebuilder:validation:Optional
	IsUnionRule *float64 `json:"isUnionRule,omitempty" tf:"is_union_rule,omitempty"`

	// Policy view name, eg:`cvm_device`,`BANDWIDTHPACKAGE`, refer to `data.tencentcloud_monitor_policy_conditions(policy_view_name)`.
	// +kubebuilder:validation:Required
	PolicyViewName *string `json:"policyViewName" tf:"policy_view_name,omitempty"`

	// The project id to which the policy group belongs, default is `0`.
	// +kubebuilder:validation:Optional
	ProjectID *float64 `json:"projectId,omitempty" tf:"project_id,omitempty"`

	// Policy group's remark information.
	// +kubebuilder:validation:Required
	Remark *string `json:"remark" tf:"remark,omitempty"`
}

type PolicyGroupReceiversObservation struct {
	EndTime *float64 `json:"endTime,omitempty" tf:"end_time,omitempty"`

	NeedSendNotice *float64 `json:"needSendNotice,omitempty" tf:"need_send_notice,omitempty"`

	NotifyWay []*string `json:"notifyWay,omitempty" tf:"notify_way,omitempty"`

	PersonInterval *float64 `json:"personInterval,omitempty" tf:"person_interval,omitempty"`

	ReceiveLanguage *string `json:"receiveLanguage,omitempty" tf:"receive_language,omitempty"`

	ReceiverGroupList []*float64 `json:"receiverGroupList,omitempty" tf:"receiver_group_list,omitempty"`

	ReceiverType *string `json:"receiverType,omitempty" tf:"receiver_type,omitempty"`

	ReceiverUserList []*float64 `json:"receiverUserList,omitempty" tf:"receiver_user_list,omitempty"`

	RecoverNotify []*string `json:"recoverNotify,omitempty" tf:"recover_notify,omitempty"`

	RoundInterval *float64 `json:"roundInterval,omitempty" tf:"round_interval,omitempty"`

	RoundNumber *float64 `json:"roundNumber,omitempty" tf:"round_number,omitempty"`

	SendFor []*string `json:"sendFor,omitempty" tf:"send_for,omitempty"`

	StartTime *float64 `json:"startTime,omitempty" tf:"start_time,omitempty"`

	UIDList []*float64 `json:"uidList,omitempty" tf:"uid_list,omitempty"`
}

type PolicyGroupReceiversParameters struct {
}

// PolicyGroupSpec defines the desired state of PolicyGroup
type PolicyGroupSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     PolicyGroupParameters `json:"forProvider"`
}

// PolicyGroupStatus defines the observed state of PolicyGroup.
type PolicyGroupStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        PolicyGroupObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// PolicyGroup is the Schema for the PolicyGroups API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,tencentcloudjet}
type PolicyGroup struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              PolicyGroupSpec   `json:"spec"`
	Status            PolicyGroupStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// PolicyGroupList contains a list of PolicyGroups
type PolicyGroupList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []PolicyGroup `json:"items"`
}

// Repository type metadata.
var (
	PolicyGroup_Kind             = "PolicyGroup"
	PolicyGroup_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: PolicyGroup_Kind}.String()
	PolicyGroup_KindAPIVersion   = PolicyGroup_Kind + "." + CRDGroupVersion.String()
	PolicyGroup_GroupVersionKind = CRDGroupVersion.WithKind(PolicyGroup_Kind)
)

func init() {
	SchemeBuilder.Register(&PolicyGroup{}, &PolicyGroupList{})
}
