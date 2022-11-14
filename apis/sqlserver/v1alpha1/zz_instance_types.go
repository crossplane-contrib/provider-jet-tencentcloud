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

type InstanceObservation struct {
	CreateTime *string `json:"createTime,omitempty" tf:"create_time,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	RoFlag *string `json:"roFlag,omitempty" tf:"ro_flag,omitempty"`

	Status *float64 `json:"status,omitempty" tf:"status,omitempty"`

	Vip *string `json:"vip,omitempty" tf:"vip,omitempty"`

	Vport *float64 `json:"vport,omitempty" tf:"vport,omitempty"`
}

type InstanceParameters struct {

	// Automatic renewal sign. 0 for normal renewal, 1 for automatic renewal (Default). Only valid when purchasing a prepaid instance.
	// +kubebuilder:validation:Optional
	AutoRenew *float64 `json:"autoRenew,omitempty" tf:"auto_renew,omitempty"`

	// Whether to use the voucher automatically; 1 for yes, 0 for no, the default is 0.
	// +kubebuilder:validation:Optional
	AutoVoucher *float64 `json:"autoVoucher,omitempty" tf:"auto_voucher,omitempty"`

	// Availability zone.
	// +kubebuilder:validation:Optional
	AvailabilityZone *string `json:"availabilityZone,omitempty" tf:"availability_zone,omitempty"`

	// Pay type of the SQL Server instance. Available values `PREPAID`, `POSTPAID_BY_HOUR`.
	// +kubebuilder:validation:Optional
	ChargeType *string `json:"chargeType,omitempty" tf:"charge_type,omitempty"`

	// Version of the SQL Server database engine. Allowed values are `2008R2`(SQL Server 2008 Enterprise), `2012SP3`(SQL Server 2012 Enterprise), `2016SP1` (SQL Server 2016 Enterprise), `201602`(SQL Server 2016 Standard) and `2017`(SQL Server 2017 Enterprise). Default is `2008R2`.
	// +kubebuilder:validation:Optional
	EngineVersion *string `json:"engineVersion,omitempty" tf:"engine_version,omitempty"`

	// Instance type. `DUAL` (dual-server high availability), `CLUSTER` (cluster). Default is `DUAL`.
	// +kubebuilder:validation:Optional
	HaType *string `json:"haType,omitempty" tf:"ha_type,omitempty"`

	// Start time of the maintenance in one day, format like `HH:mm`.
	// +kubebuilder:validation:Optional
	MaintenanceStartTime *string `json:"maintenanceStartTime,omitempty" tf:"maintenance_start_time,omitempty"`

	// The timespan of maintenance in one day, unit is hour.
	// +kubebuilder:validation:Optional
	MaintenanceTimeSpan *float64 `json:"maintenanceTimeSpan,omitempty" tf:"maintenance_time_span,omitempty"`

	// A list of integer indicates weekly maintenance. For example, [2,7] presents do weekly maintenance on every Tuesday and Sunday.
	// +kubebuilder:validation:Optional
	MaintenanceWeekSet []*float64 `json:"maintenanceWeekSet,omitempty" tf:"maintenance_week_set,omitempty"`

	// Memory size (in GB). Allowed value must be larger than `memory` that data source `tencentcloud_sqlserver_specinfos` provides.
	// +kubebuilder:validation:Required
	Memory *float64 `json:"memory" tf:"memory,omitempty"`

	// Indicate whether to deploy across availability zones.
	// +kubebuilder:validation:Optional
	MultiZones *bool `json:"multiZones,omitempty" tf:"multi_zones,omitempty"`

	// Name of the SQL Server instance.
	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// Purchase instance period in month. The value does not exceed 48.
	// +kubebuilder:validation:Optional
	Period *float64 `json:"period,omitempty" tf:"period,omitempty"`

	// Project ID, default value is 0.
	// +kubebuilder:validation:Optional
	ProjectID *float64 `json:"projectId,omitempty" tf:"project_id,omitempty"`

	// Security group bound to the instance.
	// +kubebuilder:validation:Optional
	SecurityGroups []*string `json:"securityGroups,omitempty" tf:"security_groups,omitempty"`

	// Disk size (in GB). Allowed value must be a multiple of 10. The storage must be set with the limit of `storage_min` and `storage_max` which data source `tencentcloud_sqlserver_specinfos` provides.
	// +kubebuilder:validation:Required
	Storage *float64 `json:"storage" tf:"storage,omitempty"`

	// ID of subnet.
	// +crossplane:generate:reference:type=github.com/crossplane-contrib/provider-tencentcloud/apis/vpc/v1alpha1.Subnet
	// +kubebuilder:validation:Optional
	SubnetID *string `json:"subnetId,omitempty" tf:"subnet_id,omitempty"`

	// +kubebuilder:validation:Optional
	SubnetIDRef *v1.Reference `json:"subnetIdRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	SubnetIDSelector *v1.Selector `json:"subnetIdSelector,omitempty" tf:"-"`

	// The tags of the SQL Server.
	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// ID of VPC.
	// +crossplane:generate:reference:type=github.com/crossplane-contrib/provider-tencentcloud/apis/vpc/v1alpha1.VPC
	// +kubebuilder:validation:Optional
	VPCID *string `json:"vpcId,omitempty" tf:"vpc_id,omitempty"`

	// +kubebuilder:validation:Optional
	VPCIDRef *v1.Reference `json:"vpcidRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	VPCIDSelector *v1.Selector `json:"vpcidSelector,omitempty" tf:"-"`

	// An array of voucher IDs, currently only one can be used for a single order.
	// +kubebuilder:validation:Optional
	VoucherIds []*string `json:"voucherIds,omitempty" tf:"voucher_ids,omitempty"`
}

// InstanceSpec defines the desired state of Instance
type InstanceSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     InstanceParameters `json:"forProvider"`
}

// InstanceStatus defines the observed state of Instance.
type InstanceStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        InstanceObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Instance is the Schema for the Instances API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,tencentcloudjet}
type Instance struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              InstanceSpec   `json:"spec"`
	Status            InstanceStatus `json:"status,omitempty"`
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
