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

type HostResourceInitParameters struct {
}

type HostResourceObservation struct {

	// The number of available CPU cores of the instance.
	CPUAvailableNum *float64 `json:"cpuAvailableNum,omitempty" tf:"cpu_available_num,omitempty"`

	// The number of total CPU cores of the instance.
	CPUTotalNum *float64 `json:"cpuTotalNum,omitempty" tf:"cpu_total_num,omitempty"`

	// Instance disk available capacity, unit in GB.
	DiskAvailableSize *float64 `json:"diskAvailableSize,omitempty" tf:"disk_available_size,omitempty"`

	// Instance disk total capacity, unit in GB.
	DiskTotalSize *float64 `json:"diskTotalSize,omitempty" tf:"disk_total_size,omitempty"`

	// Type of the disk.
	DiskType *string `json:"diskType,omitempty" tf:"disk_type,omitempty"`

	// Instance memory available capacity, unit in GB.
	MemoryAvailableSize *float64 `json:"memoryAvailableSize,omitempty" tf:"memory_available_size,omitempty"`

	// Instance memory total capacity, unit in GB.
	MemoryTotalSize *float64 `json:"memoryTotalSize,omitempty" tf:"memory_total_size,omitempty"`
}

type HostResourceParameters struct {
}

type InstanceInitParameters struct {

	// The available zone for the CDH instance.
	// The available zone for the CDH instance.
	AvailabilityZone *string `json:"availabilityZone,omitempty" tf:"availability_zone,omitempty"`

	// The charge type of instance. Valid values are PREPAID. The default is PREPAID.
	// The charge type of instance. Valid values are `PREPAID`. The default is `PREPAID`.
	ChargeType *string `json:"chargeType,omitempty" tf:"charge_type,omitempty"`

	// The name of the CDH instance. The max length of host_name is 60.
	// The name of the CDH instance. The max length of host_name is 60.
	HostName *string `json:"hostName,omitempty" tf:"host_name,omitempty"`

	// The type of the CDH instance.
	// The type of the CDH instance.
	HostType *string `json:"hostType,omitempty" tf:"host_type,omitempty"`

	// The tenancy (time unit is month) of the prepaid instance, NOTE: it only works when charge_type is set to PREPAID. Valid values are 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 24, 36.
	// The tenancy (time unit is month) of the prepaid instance, NOTE: it only works when charge_type is set to `PREPAID`. Valid values are `1`, `2`, `3`, `4`, `5`, `6`, `7`, `8`, `9`, `10`, `11`, `12`, `24`, `36`.
	PrepaidPeriod *float64 `json:"prepaidPeriod,omitempty" tf:"prepaid_period,omitempty"`

	// Auto renewal flag. Valid values: NOTIFY_AND_AUTO_RENEW: notify upon expiration and renew automatically, NOTIFY_AND_MANUAL_RENEW: notify upon expiration but do not renew automatically, DISABLE_NOTIFY_AND_MANUAL_RENEW: neither notify upon expiration nor renew automatically. Default value: NOTIFY_AND_MANUAL_RENEW. If this parameter is specified as NOTIFY_AND_AUTO_RENEW, the instance will be automatically renewed on a monthly basis if the account balance is sufficient. NOTE: it only works when charge_type is set to PREPAID.
	// Auto renewal flag. Valid values: `NOTIFY_AND_AUTO_RENEW`: notify upon expiration and renew automatically, `NOTIFY_AND_MANUAL_RENEW`: notify upon expiration but do not renew automatically, `DISABLE_NOTIFY_AND_MANUAL_RENEW`: neither notify upon expiration nor renew automatically. Default value: `NOTIFY_AND_MANUAL_RENEW`. If this parameter is specified as `NOTIFY_AND_AUTO_RENEW`, the instance will be automatically renewed on a monthly basis if the account balance is sufficient. NOTE: it only works when charge_type is set to `PREPAID`.
	PrepaidRenewFlag *string `json:"prepaidRenewFlag,omitempty" tf:"prepaid_renew_flag,omitempty"`

	// The project the instance belongs to, default to 0.
	// The project the instance belongs to, default to 0.
	ProjectID *float64 `json:"projectId,omitempty" tf:"project_id,omitempty"`
}

type InstanceObservation struct {

	// The available zone for the CDH instance.
	// The available zone for the CDH instance.
	AvailabilityZone *string `json:"availabilityZone,omitempty" tf:"availability_zone,omitempty"`

	// The charge type of instance. Valid values are PREPAID. The default is PREPAID.
	// The charge type of instance. Valid values are `PREPAID`. The default is `PREPAID`.
	ChargeType *string `json:"chargeType,omitempty" tf:"charge_type,omitempty"`

	// Create time of the instance.
	// Create time of the instance.
	CreateTime *string `json:"createTime,omitempty" tf:"create_time,omitempty"`

	// Id of CVM instances that have been created on the CDH instance.
	// Id of CVM instances that have been created on the CDH instance.
	CvmInstanceIds []*string `json:"cvmInstanceIds,omitempty" tf:"cvm_instance_ids,omitempty"`

	// Expired time of the instance.
	// Expired time of the instance.
	ExpiredTime *string `json:"expiredTime,omitempty" tf:"expired_time,omitempty"`

	// The name of the CDH instance. The max length of host_name is 60.
	// The name of the CDH instance. The max length of host_name is 60.
	HostName *string `json:"hostName,omitempty" tf:"host_name,omitempty"`

	// An information list of host resource. Each element contains the following attributes:
	// An information list of host resource. Each element contains the following attributes:
	HostResource []HostResourceObservation `json:"hostResource,omitempty" tf:"host_resource,omitempty"`

	// State of the CDH instance.
	// State of the CDH instance.
	HostState *string `json:"hostState,omitempty" tf:"host_state,omitempty"`

	// The type of the CDH instance.
	// The type of the CDH instance.
	HostType *string `json:"hostType,omitempty" tf:"host_type,omitempty"`

	// ID of the resource.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The tenancy (time unit is month) of the prepaid instance, NOTE: it only works when charge_type is set to PREPAID. Valid values are 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 24, 36.
	// The tenancy (time unit is month) of the prepaid instance, NOTE: it only works when charge_type is set to `PREPAID`. Valid values are `1`, `2`, `3`, `4`, `5`, `6`, `7`, `8`, `9`, `10`, `11`, `12`, `24`, `36`.
	PrepaidPeriod *float64 `json:"prepaidPeriod,omitempty" tf:"prepaid_period,omitempty"`

	// Auto renewal flag. Valid values: NOTIFY_AND_AUTO_RENEW: notify upon expiration and renew automatically, NOTIFY_AND_MANUAL_RENEW: notify upon expiration but do not renew automatically, DISABLE_NOTIFY_AND_MANUAL_RENEW: neither notify upon expiration nor renew automatically. Default value: NOTIFY_AND_MANUAL_RENEW. If this parameter is specified as NOTIFY_AND_AUTO_RENEW, the instance will be automatically renewed on a monthly basis if the account balance is sufficient. NOTE: it only works when charge_type is set to PREPAID.
	// Auto renewal flag. Valid values: `NOTIFY_AND_AUTO_RENEW`: notify upon expiration and renew automatically, `NOTIFY_AND_MANUAL_RENEW`: notify upon expiration but do not renew automatically, `DISABLE_NOTIFY_AND_MANUAL_RENEW`: neither notify upon expiration nor renew automatically. Default value: `NOTIFY_AND_MANUAL_RENEW`. If this parameter is specified as `NOTIFY_AND_AUTO_RENEW`, the instance will be automatically renewed on a monthly basis if the account balance is sufficient. NOTE: it only works when charge_type is set to `PREPAID`.
	PrepaidRenewFlag *string `json:"prepaidRenewFlag,omitempty" tf:"prepaid_renew_flag,omitempty"`

	// The project the instance belongs to, default to 0.
	// The project the instance belongs to, default to 0.
	ProjectID *float64 `json:"projectId,omitempty" tf:"project_id,omitempty"`
}

type InstanceParameters struct {

	// The available zone for the CDH instance.
	// The available zone for the CDH instance.
	// +kubebuilder:validation:Optional
	AvailabilityZone *string `json:"availabilityZone,omitempty" tf:"availability_zone,omitempty"`

	// The charge type of instance. Valid values are PREPAID. The default is PREPAID.
	// The charge type of instance. Valid values are `PREPAID`. The default is `PREPAID`.
	// +kubebuilder:validation:Optional
	ChargeType *string `json:"chargeType,omitempty" tf:"charge_type,omitempty"`

	// The name of the CDH instance. The max length of host_name is 60.
	// The name of the CDH instance. The max length of host_name is 60.
	// +kubebuilder:validation:Optional
	HostName *string `json:"hostName,omitempty" tf:"host_name,omitempty"`

	// The type of the CDH instance.
	// The type of the CDH instance.
	// +kubebuilder:validation:Optional
	HostType *string `json:"hostType,omitempty" tf:"host_type,omitempty"`

	// The tenancy (time unit is month) of the prepaid instance, NOTE: it only works when charge_type is set to PREPAID. Valid values are 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 24, 36.
	// The tenancy (time unit is month) of the prepaid instance, NOTE: it only works when charge_type is set to `PREPAID`. Valid values are `1`, `2`, `3`, `4`, `5`, `6`, `7`, `8`, `9`, `10`, `11`, `12`, `24`, `36`.
	// +kubebuilder:validation:Optional
	PrepaidPeriod *float64 `json:"prepaidPeriod,omitempty" tf:"prepaid_period,omitempty"`

	// Auto renewal flag. Valid values: NOTIFY_AND_AUTO_RENEW: notify upon expiration and renew automatically, NOTIFY_AND_MANUAL_RENEW: notify upon expiration but do not renew automatically, DISABLE_NOTIFY_AND_MANUAL_RENEW: neither notify upon expiration nor renew automatically. Default value: NOTIFY_AND_MANUAL_RENEW. If this parameter is specified as NOTIFY_AND_AUTO_RENEW, the instance will be automatically renewed on a monthly basis if the account balance is sufficient. NOTE: it only works when charge_type is set to PREPAID.
	// Auto renewal flag. Valid values: `NOTIFY_AND_AUTO_RENEW`: notify upon expiration and renew automatically, `NOTIFY_AND_MANUAL_RENEW`: notify upon expiration but do not renew automatically, `DISABLE_NOTIFY_AND_MANUAL_RENEW`: neither notify upon expiration nor renew automatically. Default value: `NOTIFY_AND_MANUAL_RENEW`. If this parameter is specified as `NOTIFY_AND_AUTO_RENEW`, the instance will be automatically renewed on a monthly basis if the account balance is sufficient. NOTE: it only works when charge_type is set to `PREPAID`.
	// +kubebuilder:validation:Optional
	PrepaidRenewFlag *string `json:"prepaidRenewFlag,omitempty" tf:"prepaid_renew_flag,omitempty"`

	// The project the instance belongs to, default to 0.
	// The project the instance belongs to, default to 0.
	// +kubebuilder:validation:Optional
	ProjectID *float64 `json:"projectId,omitempty" tf:"project_id,omitempty"`
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

// Instance is the Schema for the Instances API. Provides a resource to manage CDH instance.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,tencentcloud}
type Instance struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.availabilityZone) || (has(self.initProvider) && has(self.initProvider.availabilityZone))",message="spec.forProvider.availabilityZone is a required parameter"
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
