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

type NotificationInitParameters struct {

	// A list of Notification Types that trigger notifications. Acceptable values are SCALE_OUT_FAILED, SCALE_IN_SUCCESSFUL, SCALE_IN_FAILED, REPLACE_UNHEALTHY_INSTANCE_SUCCESSFUL and REPLACE_UNHEALTHY_INSTANCE_FAILED.
	// A list of Notification Types that trigger notifications. Acceptable values are `SCALE_OUT_FAILED`, `SCALE_IN_SUCCESSFUL`, `SCALE_IN_FAILED`, `REPLACE_UNHEALTHY_INSTANCE_SUCCESSFUL` and `REPLACE_UNHEALTHY_INSTANCE_FAILED`.
	NotificationTypes []*string `json:"notificationTypes,omitempty" tf:"notification_types,omitempty"`

	// A group of user IDs to be notified.
	// A group of user IDs to be notified.
	NotificationUserGroupIds []*string `json:"notificationUserGroupIds,omitempty" tf:"notification_user_group_ids,omitempty"`
}

type NotificationObservation struct {

	// ID of the resource.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// A list of Notification Types that trigger notifications. Acceptable values are SCALE_OUT_FAILED, SCALE_IN_SUCCESSFUL, SCALE_IN_FAILED, REPLACE_UNHEALTHY_INSTANCE_SUCCESSFUL and REPLACE_UNHEALTHY_INSTANCE_FAILED.
	// A list of Notification Types that trigger notifications. Acceptable values are `SCALE_OUT_FAILED`, `SCALE_IN_SUCCESSFUL`, `SCALE_IN_FAILED`, `REPLACE_UNHEALTHY_INSTANCE_SUCCESSFUL` and `REPLACE_UNHEALTHY_INSTANCE_FAILED`.
	NotificationTypes []*string `json:"notificationTypes,omitempty" tf:"notification_types,omitempty"`

	// A group of user IDs to be notified.
	// A group of user IDs to be notified.
	NotificationUserGroupIds []*string `json:"notificationUserGroupIds,omitempty" tf:"notification_user_group_ids,omitempty"`

	// ID of a scaling group.
	// ID of a scaling group.
	ScalingGroupID *string `json:"scalingGroupId,omitempty" tf:"scaling_group_id,omitempty"`
}

type NotificationParameters struct {

	// A list of Notification Types that trigger notifications. Acceptable values are SCALE_OUT_FAILED, SCALE_IN_SUCCESSFUL, SCALE_IN_FAILED, REPLACE_UNHEALTHY_INSTANCE_SUCCESSFUL and REPLACE_UNHEALTHY_INSTANCE_FAILED.
	// A list of Notification Types that trigger notifications. Acceptable values are `SCALE_OUT_FAILED`, `SCALE_IN_SUCCESSFUL`, `SCALE_IN_FAILED`, `REPLACE_UNHEALTHY_INSTANCE_SUCCESSFUL` and `REPLACE_UNHEALTHY_INSTANCE_FAILED`.
	// +kubebuilder:validation:Optional
	NotificationTypes []*string `json:"notificationTypes,omitempty" tf:"notification_types,omitempty"`

	// A group of user IDs to be notified.
	// A group of user IDs to be notified.
	// +kubebuilder:validation:Optional
	NotificationUserGroupIds []*string `json:"notificationUserGroupIds,omitempty" tf:"notification_user_group_ids,omitempty"`

	// ID of a scaling group.
	// ID of a scaling group.
	// +crossplane:generate:reference:type=ScalingGroup
	// +kubebuilder:validation:Optional
	ScalingGroupID *string `json:"scalingGroupId,omitempty" tf:"scaling_group_id,omitempty"`

	// Reference to a ScalingGroup to populate scalingGroupId.
	// +kubebuilder:validation:Optional
	ScalingGroupIDRef *v1.Reference `json:"scalingGroupIdRef,omitempty" tf:"-"`

	// Selector for a ScalingGroup to populate scalingGroupId.
	// +kubebuilder:validation:Optional
	ScalingGroupIDSelector *v1.Selector `json:"scalingGroupIdSelector,omitempty" tf:"-"`
}

// NotificationSpec defines the desired state of Notification
type NotificationSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     NotificationParameters `json:"forProvider"`
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
	InitProvider NotificationInitParameters `json:"initProvider,omitempty"`
}

// NotificationStatus defines the observed state of Notification.
type NotificationStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        NotificationObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Notification is the Schema for the Notifications API. Provides a resource for an AS (Auto scaling) notification.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,tencentcloud}
type Notification struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.notificationTypes) || (has(self.initProvider) && has(self.initProvider.notificationTypes))",message="spec.forProvider.notificationTypes is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.notificationUserGroupIds) || (has(self.initProvider) && has(self.initProvider.notificationUserGroupIds))",message="spec.forProvider.notificationUserGroupIds is a required parameter"
	Spec   NotificationSpec   `json:"spec"`
	Status NotificationStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// NotificationList contains a list of Notifications
type NotificationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Notification `json:"items"`
}

// Repository type metadata.
var (
	Notification_Kind             = "Notification"
	Notification_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Notification_Kind}.String()
	Notification_KindAPIVersion   = Notification_Kind + "." + CRDGroupVersion.String()
	Notification_GroupVersionKind = CRDGroupVersion.WithKind(Notification_Kind)
)

func init() {
	SchemeBuilder.Register(&Notification{}, &NotificationList{})
}
