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

type RolePolicyAttachmentInitParameters struct {
}

type RolePolicyAttachmentObservation struct {

	// Mode of Creation of the CAM role policy attachment. 1 means the CAM policy attachment is created by production, and the others indicate syntax strategy ways.
	// Mode of Creation of the CAM role policy attachment. `1` means the CAM policy attachment is created by production, and the others indicate syntax strategy ways.
	CreateMode *float64 `json:"createMode,omitempty" tf:"create_mode,omitempty"`

	// The create time of the CAM role policy attachment.
	// The create time of the CAM role policy attachment.
	CreateTime *string `json:"createTime,omitempty" tf:"create_time,omitempty"`

	// ID of the resource.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Name of the policy.
	// Name of the policy.
	PolicyName *string `json:"policyName,omitempty" tf:"policy_name,omitempty"`

	// Type of the policy strategy. User means customer strategy and QCS means preset strategy.
	// Type of the policy strategy. `User` means customer strategy and `QCS` means preset strategy.
	PolicyType *string `json:"policyType,omitempty" tf:"policy_type,omitempty"`

	// Name of the attached CAM role.
	// Name of the attached CAM role.
	RoleName *string `json:"roleName,omitempty" tf:"role_name,omitempty"`
}

type RolePolicyAttachmentParameters struct {

	// Name of the policy.
	// Name of the policy.
	// +crossplane:generate:reference:type=Policy
	// +kubebuilder:validation:Optional
	PolicyName *string `json:"policyName,omitempty" tf:"policy_name,omitempty"`

	// Reference to a Policy to populate policyName.
	// +kubebuilder:validation:Optional
	PolicyNameRef *v1.Reference `json:"policyNameRef,omitempty" tf:"-"`

	// Selector for a Policy to populate policyName.
	// +kubebuilder:validation:Optional
	PolicyNameSelector *v1.Selector `json:"policyNameSelector,omitempty" tf:"-"`

	// Name of the attached CAM role.
	// Name of the attached CAM role.
	// +crossplane:generate:reference:type=Role
	// +kubebuilder:validation:Optional
	RoleName *string `json:"roleName,omitempty" tf:"role_name,omitempty"`

	// Reference to a Role to populate roleName.
	// +kubebuilder:validation:Optional
	RoleNameRef *v1.Reference `json:"roleNameRef,omitempty" tf:"-"`

	// Selector for a Role to populate roleName.
	// +kubebuilder:validation:Optional
	RoleNameSelector *v1.Selector `json:"roleNameSelector,omitempty" tf:"-"`
}

// RolePolicyAttachmentSpec defines the desired state of RolePolicyAttachment
type RolePolicyAttachmentSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     RolePolicyAttachmentParameters `json:"forProvider"`
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
	InitProvider RolePolicyAttachmentInitParameters `json:"initProvider,omitempty"`
}

// RolePolicyAttachmentStatus defines the observed state of RolePolicyAttachment.
type RolePolicyAttachmentStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        RolePolicyAttachmentObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// RolePolicyAttachment is the Schema for the RolePolicyAttachments API. Provides a resource to create a CAM role policy attachment.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,tencentcloud}
type RolePolicyAttachment struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              RolePolicyAttachmentSpec   `json:"spec"`
	Status            RolePolicyAttachmentStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// RolePolicyAttachmentList contains a list of RolePolicyAttachments
type RolePolicyAttachmentList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []RolePolicyAttachment `json:"items"`
}

// Repository type metadata.
var (
	RolePolicyAttachment_Kind             = "RolePolicyAttachment"
	RolePolicyAttachment_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: RolePolicyAttachment_Kind}.String()
	RolePolicyAttachment_KindAPIVersion   = RolePolicyAttachment_Kind + "." + CRDGroupVersion.String()
	RolePolicyAttachment_GroupVersionKind = CRDGroupVersion.WithKind(RolePolicyAttachment_Kind)
)

func init() {
	SchemeBuilder.Register(&RolePolicyAttachment{}, &RolePolicyAttachmentList{})
}
