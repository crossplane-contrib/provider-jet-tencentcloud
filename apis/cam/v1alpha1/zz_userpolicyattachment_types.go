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

type UserPolicyAttachmentInitParameters struct {

	// Name of the attached CAM user as uniq key.
	// Name of the attached CAM user as uniq key.
	UserName *string `json:"userName,omitempty" tf:"user_name,omitempty"`
}

type UserPolicyAttachmentObservation struct {

	// Mode of Creation of the CAM user policy attachment. 1 means the CAM policy attachment is created by production, and the others indicate syntax strategy ways.
	// Mode of Creation of the CAM user policy attachment. `1` means the CAM policy attachment is created by production, and the others indicate syntax strategy ways.
	CreateMode *float64 `json:"createMode,omitempty" tf:"create_mode,omitempty"`

	// Create time of the CAM user policy attachment.
	// Create time of the CAM user policy attachment.
	CreateTime *string `json:"createTime,omitempty" tf:"create_time,omitempty"`

	// ID of the resource.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// ID of the policy.
	// ID of the policy.
	PolicyID *string `json:"policyId,omitempty" tf:"policy_id,omitempty"`

	// Name of the policy.
	// Name of the policy.
	PolicyName *string `json:"policyName,omitempty" tf:"policy_name,omitempty"`

	// Type of the policy strategy. User means customer strategy and QCS means preset strategy.
	// Type of the policy strategy. `User` means customer strategy and `QCS` means preset strategy.
	PolicyType *string `json:"policyType,omitempty" tf:"policy_type,omitempty"`

	// It has been deprecated from version 1.59.5. Use user_name instead. ID of the attached CAM user.
	// ID of the attached CAM user.
	UserID *string `json:"userId,omitempty" tf:"user_id,omitempty"`

	// Name of the attached CAM user as uniq key.
	// Name of the attached CAM user as uniq key.
	UserName *string `json:"userName,omitempty" tf:"user_name,omitempty"`
}

type UserPolicyAttachmentParameters struct {

	// ID of the policy.
	// ID of the policy.
	// +crossplane:generate:reference:type=Policy
	// +kubebuilder:validation:Optional
	PolicyID *string `json:"policyId,omitempty" tf:"policy_id,omitempty"`

	// Reference to a Policy to populate policyId.
	// +kubebuilder:validation:Optional
	PolicyIDRef *v1.Reference `json:"policyIdRef,omitempty" tf:"-"`

	// Selector for a Policy to populate policyId.
	// +kubebuilder:validation:Optional
	PolicyIDSelector *v1.Selector `json:"policyIdSelector,omitempty" tf:"-"`

	// It has been deprecated from version 1.59.5. Use user_name instead. ID of the attached CAM user.
	// ID of the attached CAM user.
	// +crossplane:generate:reference:type=User
	// +kubebuilder:validation:Optional
	UserID *string `json:"userId,omitempty" tf:"user_id,omitempty"`

	// Reference to a User to populate userId.
	// +kubebuilder:validation:Optional
	UserIDRef *v1.Reference `json:"userIdRef,omitempty" tf:"-"`

	// Selector for a User to populate userId.
	// +kubebuilder:validation:Optional
	UserIDSelector *v1.Selector `json:"userIdSelector,omitempty" tf:"-"`

	// Name of the attached CAM user as uniq key.
	// Name of the attached CAM user as uniq key.
	// +kubebuilder:validation:Optional
	UserName *string `json:"userName,omitempty" tf:"user_name,omitempty"`
}

// UserPolicyAttachmentSpec defines the desired state of UserPolicyAttachment
type UserPolicyAttachmentSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     UserPolicyAttachmentParameters `json:"forProvider"`
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
	InitProvider UserPolicyAttachmentInitParameters `json:"initProvider,omitempty"`
}

// UserPolicyAttachmentStatus defines the observed state of UserPolicyAttachment.
type UserPolicyAttachmentStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        UserPolicyAttachmentObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// UserPolicyAttachment is the Schema for the UserPolicyAttachments API. Provides a resource to create a CAM user policy attachment.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,tencentcloud}
type UserPolicyAttachment struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              UserPolicyAttachmentSpec   `json:"spec"`
	Status            UserPolicyAttachmentStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// UserPolicyAttachmentList contains a list of UserPolicyAttachments
type UserPolicyAttachmentList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []UserPolicyAttachment `json:"items"`
}

// Repository type metadata.
var (
	UserPolicyAttachment_Kind             = "UserPolicyAttachment"
	UserPolicyAttachment_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: UserPolicyAttachment_Kind}.String()
	UserPolicyAttachment_KindAPIVersion   = UserPolicyAttachment_Kind + "." + CRDGroupVersion.String()
	UserPolicyAttachment_GroupVersionKind = CRDGroupVersion.WithKind(UserPolicyAttachment_Kind)
)

func init() {
	SchemeBuilder.Register(&UserPolicyAttachment{}, &UserPolicyAttachmentList{})
}
