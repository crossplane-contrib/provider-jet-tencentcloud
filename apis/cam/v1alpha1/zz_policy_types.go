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

type PolicyInitParameters struct {

	// Description of the CAM policy.
	// Description of the CAM policy.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Document of the CAM policy. The syntax refers to CAM POLICY. The elements in JSON claimed supporting two types as string and array only support type array; 2.
	// Document of the CAM policy. The syntax refers to [CAM POLICY](https://intl.cloud.tencent.com/document/product/598/10604). The elements in JSON claimed supporting two types as `string` and `array` only support type `array`; 2.
	Document *string `json:"document,omitempty" tf:"document,omitempty"`
}

type PolicyObservation struct {

	// Create time of the CAM policy.
	// Create time of the CAM policy.
	CreateTime *string `json:"createTime,omitempty" tf:"create_time,omitempty"`

	// Description of the CAM policy.
	// Description of the CAM policy.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Document of the CAM policy. The syntax refers to CAM POLICY. The elements in JSON claimed supporting two types as string and array only support type array; 2.
	// Document of the CAM policy. The syntax refers to [CAM POLICY](https://intl.cloud.tencent.com/document/product/598/10604). The elements in JSON claimed supporting two types as `string` and `array` only support type `array`; 2.
	Document *string `json:"document,omitempty" tf:"document,omitempty"`

	// ID of the resource.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Type of the policy strategy. Valid values: 1, 2.  1 means customer strategy and 2 means preset strategy.
	// Type of the policy strategy. Valid values: `1`, `2`.  `1` means customer strategy and `2` means preset strategy.
	Type *float64 `json:"type,omitempty" tf:"type,omitempty"`

	// The last update time of the CAM policy.
	// The last update time of the CAM policy.
	UpdateTime *string `json:"updateTime,omitempty" tf:"update_time,omitempty"`
}

type PolicyParameters struct {

	// Description of the CAM policy.
	// Description of the CAM policy.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Document of the CAM policy. The syntax refers to CAM POLICY. The elements in JSON claimed supporting two types as string and array only support type array; 2.
	// Document of the CAM policy. The syntax refers to [CAM POLICY](https://intl.cloud.tencent.com/document/product/598/10604). The elements in JSON claimed supporting two types as `string` and `array` only support type `array`; 2.
	// +kubebuilder:validation:Optional
	Document *string `json:"document,omitempty" tf:"document,omitempty"`
}

// PolicySpec defines the desired state of Policy
type PolicySpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     PolicyParameters `json:"forProvider"`
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
	InitProvider PolicyInitParameters `json:"initProvider,omitempty"`
}

// PolicyStatus defines the observed state of Policy.
type PolicyStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        PolicyObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Policy is the Schema for the Policys API. Provides a resource to create a CAM policy.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,tencentcloud}
type Policy struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.document) || (has(self.initProvider) && has(self.initProvider.document))",message="spec.forProvider.document is a required parameter"
	Spec   PolicySpec   `json:"spec"`
	Status PolicyStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// PolicyList contains a list of Policys
type PolicyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Policy `json:"items"`
}

// Repository type metadata.
var (
	Policy_Kind             = "Policy"
	Policy_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Policy_Kind}.String()
	Policy_KindAPIVersion   = Policy_Kind + "." + CRDGroupVersion.String()
	Policy_GroupVersionKind = CRDGroupVersion.WithKind(Policy_Kind)
)

func init() {
	SchemeBuilder.Register(&Policy{}, &PolicyList{})
}
