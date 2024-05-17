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

type OriginGroupInitParameters struct {

	// Type of the origin group, this field should be set when OriginType is self, otherwise leave it empty. Valid values: area: select an origin by using Geo info of the client IP and Area field in Records; weight: weighted select an origin by using Weight field in Records; proto: config by HTTP protocol.
	// Type of the origin group, this field should be set when `OriginType` is self, otherwise leave it empty. Valid values: `area`: select an origin by using Geo info of the client IP and `Area` field in Records; `weight`: weighted select an origin by using `Weight` field in Records; `proto`: config by HTTP protocol.
	ConfigurationType *string `json:"configurationType,omitempty" tf:"configuration_type,omitempty"`

	// OriginGroup Name.
	// OriginGroup Name.
	OriginGroupName *string `json:"originGroupName,omitempty" tf:"origin_group_name,omitempty"`

	// Origin site records.
	// Origin site records.
	OriginRecords []OriginRecordsInitParameters `json:"originRecords,omitempty" tf:"origin_records,omitempty"`

	// Type of the origin site. Valid values: self: self-build website; cos: tencent cos; third_party: third party cos.
	// Type of the origin site. Valid values: `self`: self-build website; `cos`: tencent cos; `third_party`: third party cos.
	OriginType *string `json:"originType,omitempty" tf:"origin_type,omitempty"`
}

type OriginGroupObservation struct {

	// Type of the origin group, this field should be set when OriginType is self, otherwise leave it empty. Valid values: area: select an origin by using Geo info of the client IP and Area field in Records; weight: weighted select an origin by using Weight field in Records; proto: config by HTTP protocol.
	// Type of the origin group, this field should be set when `OriginType` is self, otherwise leave it empty. Valid values: `area`: select an origin by using Geo info of the client IP and `Area` field in Records; `weight`: weighted select an origin by using `Weight` field in Records; `proto`: config by HTTP protocol.
	ConfigurationType *string `json:"configurationType,omitempty" tf:"configuration_type,omitempty"`

	// ID of the resource.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// OriginGroup ID.
	// OriginGroup ID.
	OriginGroupID *string `json:"originGroupId,omitempty" tf:"origin_group_id,omitempty"`

	// OriginGroup Name.
	// OriginGroup Name.
	OriginGroupName *string `json:"originGroupName,omitempty" tf:"origin_group_name,omitempty"`

	// Origin site records.
	// Origin site records.
	OriginRecords []OriginRecordsObservation `json:"originRecords,omitempty" tf:"origin_records,omitempty"`

	// Type of the origin site. Valid values: self: self-build website; cos: tencent cos; third_party: third party cos.
	// Type of the origin site. Valid values: `self`: self-build website; `cos`: tencent cos; `third_party`: third party cos.
	OriginType *string `json:"originType,omitempty" tf:"origin_type,omitempty"`

	// Last modification date.
	// Last modification date.
	UpdateTime *string `json:"updateTime,omitempty" tf:"update_time,omitempty"`

	// Site ID.
	// Site ID.
	ZoneID *string `json:"zoneId,omitempty" tf:"zone_id,omitempty"`
}

type OriginGroupParameters struct {

	// Type of the origin group, this field should be set when OriginType is self, otherwise leave it empty. Valid values: area: select an origin by using Geo info of the client IP and Area field in Records; weight: weighted select an origin by using Weight field in Records; proto: config by HTTP protocol.
	// Type of the origin group, this field should be set when `OriginType` is self, otherwise leave it empty. Valid values: `area`: select an origin by using Geo info of the client IP and `Area` field in Records; `weight`: weighted select an origin by using `Weight` field in Records; `proto`: config by HTTP protocol.
	// +kubebuilder:validation:Optional
	ConfigurationType *string `json:"configurationType,omitempty" tf:"configuration_type,omitempty"`

	// OriginGroup Name.
	// OriginGroup Name.
	// +kubebuilder:validation:Optional
	OriginGroupName *string `json:"originGroupName,omitempty" tf:"origin_group_name,omitempty"`

	// Origin site records.
	// Origin site records.
	// +kubebuilder:validation:Optional
	OriginRecords []OriginRecordsParameters `json:"originRecords,omitempty" tf:"origin_records,omitempty"`

	// Type of the origin site. Valid values: self: self-build website; cos: tencent cos; third_party: third party cos.
	// Type of the origin site. Valid values: `self`: self-build website; `cos`: tencent cos; `third_party`: third party cos.
	// +kubebuilder:validation:Optional
	OriginType *string `json:"originType,omitempty" tf:"origin_type,omitempty"`

	// Site ID.
	// Site ID.
	// +crossplane:generate:reference:type=Zone
	// +kubebuilder:validation:Optional
	ZoneID *string `json:"zoneId,omitempty" tf:"zone_id,omitempty"`

	// Reference to a Zone to populate zoneId.
	// +kubebuilder:validation:Optional
	ZoneIDRef *v1.Reference `json:"zoneIdRef,omitempty" tf:"-"`

	// Selector for a Zone to populate zoneId.
	// +kubebuilder:validation:Optional
	ZoneIDSelector *v1.Selector `json:"zoneIdSelector,omitempty" tf:"-"`
}

type OriginRecordsInitParameters struct {

	// Indicating origin sites area when Type field is area. An empty List indicate the default area. Valid value:- Asia, Americas, Europe, Africa or Oceania.
	// Indicating origin sites area when `Type` field is `area`. An empty List indicate the default area. Valid value:- Asia, Americas, Europe, Africa or Oceania.
	Area []*string `json:"area,omitempty" tf:"area,omitempty"`

	// Port of the origin site. Valid value range: 1-65535.
	// Port of the origin site. Valid value range: 1-65535.
	Port *float64 `json:"port,omitempty" tf:"port,omitempty"`

	// Whether origin site is using private authentication. Only valid when OriginType is third_party.
	// Whether origin site is using private authentication. Only valid when `OriginType` is `third_party`.
	Private *bool `json:"private,omitempty" tf:"private,omitempty"`

	// Parameters for private authentication. Only valid when Private is true.
	// Parameters for private authentication. Only valid when `Private` is `true`.
	PrivateParameter []PrivateParameterInitParameters `json:"privateParameter,omitempty" tf:"private_parameter,omitempty"`

	// Record value, which could be an IPv4/IPv6 address or a domain.
	// Record value, which could be an IPv4/IPv6 address or a domain.
	Record *string `json:"record,omitempty" tf:"record,omitempty"`

	// Indicating origin sites weight when Type field is weight. Valid value range: 1-100. Sum of all weights should be 100.
	// Indicating origin sites weight when `Type` field is `weight`. Valid value range: 1-100. Sum of all weights should be 100.
	Weight *float64 `json:"weight,omitempty" tf:"weight,omitempty"`
}

type OriginRecordsObservation struct {

	// Indicating origin sites area when Type field is area. An empty List indicate the default area. Valid value:- Asia, Americas, Europe, Africa or Oceania.
	// Indicating origin sites area when `Type` field is `area`. An empty List indicate the default area. Valid value:- Asia, Americas, Europe, Africa or Oceania.
	Area []*string `json:"area,omitempty" tf:"area,omitempty"`

	// Port of the origin site. Valid value range: 1-65535.
	// Port of the origin site. Valid value range: 1-65535.
	Port *float64 `json:"port,omitempty" tf:"port,omitempty"`

	// Whether origin site is using private authentication. Only valid when OriginType is third_party.
	// Whether origin site is using private authentication. Only valid when `OriginType` is `third_party`.
	Private *bool `json:"private,omitempty" tf:"private,omitempty"`

	// Parameters for private authentication. Only valid when Private is true.
	// Parameters for private authentication. Only valid when `Private` is `true`.
	PrivateParameter []PrivateParameterObservation `json:"privateParameter,omitempty" tf:"private_parameter,omitempty"`

	// Record value, which could be an IPv4/IPv6 address or a domain.
	// Record value, which could be an IPv4/IPv6 address or a domain.
	Record *string `json:"record,omitempty" tf:"record,omitempty"`

	// ID of the resource.
	// Record Id.
	RecordID *string `json:"recordId,omitempty" tf:"record_id,omitempty"`

	// Indicating origin sites weight when Type field is weight. Valid value range: 1-100. Sum of all weights should be 100.
	// Indicating origin sites weight when `Type` field is `weight`. Valid value range: 1-100. Sum of all weights should be 100.
	Weight *float64 `json:"weight,omitempty" tf:"weight,omitempty"`
}

type OriginRecordsParameters struct {

	// Indicating origin sites area when Type field is area. An empty List indicate the default area. Valid value:- Asia, Americas, Europe, Africa or Oceania.
	// Indicating origin sites area when `Type` field is `area`. An empty List indicate the default area. Valid value:- Asia, Americas, Europe, Africa or Oceania.
	// +kubebuilder:validation:Optional
	Area []*string `json:"area,omitempty" tf:"area,omitempty"`

	// Port of the origin site. Valid value range: 1-65535.
	// Port of the origin site. Valid value range: 1-65535.
	// +kubebuilder:validation:Optional
	Port *float64 `json:"port" tf:"port,omitempty"`

	// Whether origin site is using private authentication. Only valid when OriginType is third_party.
	// Whether origin site is using private authentication. Only valid when `OriginType` is `third_party`.
	// +kubebuilder:validation:Optional
	Private *bool `json:"private,omitempty" tf:"private,omitempty"`

	// Parameters for private authentication. Only valid when Private is true.
	// Parameters for private authentication. Only valid when `Private` is `true`.
	// +kubebuilder:validation:Optional
	PrivateParameter []PrivateParameterParameters `json:"privateParameter,omitempty" tf:"private_parameter,omitempty"`

	// Record value, which could be an IPv4/IPv6 address or a domain.
	// Record value, which could be an IPv4/IPv6 address or a domain.
	// +kubebuilder:validation:Optional
	Record *string `json:"record" tf:"record,omitempty"`

	// Indicating origin sites weight when Type field is weight. Valid value range: 1-100. Sum of all weights should be 100.
	// Indicating origin sites weight when `Type` field is `weight`. Valid value range: 1-100. Sum of all weights should be 100.
	// +kubebuilder:validation:Optional
	Weight *float64 `json:"weight,omitempty" tf:"weight,omitempty"`
}

type PrivateParameterInitParameters struct {

	// Parameter Name. Valid values: AccessKeyId: Access Key ID; SecretAccessKey: Secret Access Key.
	// Parameter Name. Valid values: `AccessKeyId`: Access Key ID; `SecretAccessKey`: Secret Access Key.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Parameter value.
	// Parameter value.
	Value *string `json:"value,omitempty" tf:"value,omitempty"`
}

type PrivateParameterObservation struct {

	// Parameter Name. Valid values: AccessKeyId: Access Key ID; SecretAccessKey: Secret Access Key.
	// Parameter Name. Valid values: `AccessKeyId`: Access Key ID; `SecretAccessKey`: Secret Access Key.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Parameter value.
	// Parameter value.
	Value *string `json:"value,omitempty" tf:"value,omitempty"`
}

type PrivateParameterParameters struct {

	// Parameter Name. Valid values: AccessKeyId: Access Key ID; SecretAccessKey: Secret Access Key.
	// Parameter Name. Valid values: `AccessKeyId`: Access Key ID; `SecretAccessKey`: Secret Access Key.
	// +kubebuilder:validation:Optional
	Name *string `json:"name" tf:"name,omitempty"`

	// Parameter value.
	// Parameter value.
	// +kubebuilder:validation:Optional
	Value *string `json:"value" tf:"value,omitempty"`
}

// OriginGroupSpec defines the desired state of OriginGroup
type OriginGroupSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     OriginGroupParameters `json:"forProvider"`
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
	InitProvider OriginGroupInitParameters `json:"initProvider,omitempty"`
}

// OriginGroupStatus defines the observed state of OriginGroup.
type OriginGroupStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        OriginGroupObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// OriginGroup is the Schema for the OriginGroups API. Provides a resource to create a teo origin_group
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,tencentcloud}
type OriginGroup struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.configurationType) || (has(self.initProvider) && has(self.initProvider.configurationType))",message="spec.forProvider.configurationType is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.originGroupName) || (has(self.initProvider) && has(self.initProvider.originGroupName))",message="spec.forProvider.originGroupName is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.originRecords) || (has(self.initProvider) && has(self.initProvider.originRecords))",message="spec.forProvider.originRecords is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.originType) || (has(self.initProvider) && has(self.initProvider.originType))",message="spec.forProvider.originType is a required parameter"
	Spec   OriginGroupSpec   `json:"spec"`
	Status OriginGroupStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// OriginGroupList contains a list of OriginGroups
type OriginGroupList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []OriginGroup `json:"items"`
}

// Repository type metadata.
var (
	OriginGroup_Kind             = "OriginGroup"
	OriginGroup_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: OriginGroup_Kind}.String()
	OriginGroup_KindAPIVersion   = OriginGroup_Kind + "." + CRDGroupVersion.String()
	OriginGroup_GroupVersionKind = CRDGroupVersion.WithKind(OriginGroup_Kind)
)

func init() {
	SchemeBuilder.Register(&OriginGroup{}, &OriginGroupList{})
}
