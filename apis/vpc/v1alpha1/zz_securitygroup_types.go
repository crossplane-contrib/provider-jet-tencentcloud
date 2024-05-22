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

type SecurityGroupInitParameters struct {

	// Description of the security group.
	// Description of the security group.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Name of the security group to be queried.
	// Name of the security group to be queried.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Project ID of the security group.
	// Project ID of the security group.
	ProjectID *float64 `json:"projectId,omitempty" tf:"project_id,omitempty"`

	// Tags of the security group.
	// Tags of the security group.
	// +mapType=granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

type SecurityGroupObservation struct {

	// Description of the security group.
	// Description of the security group.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// ID of the resource.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Name of the security group to be queried.
	// Name of the security group to be queried.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Project ID of the security group.
	// Project ID of the security group.
	ProjectID *float64 `json:"projectId,omitempty" tf:"project_id,omitempty"`

	// Tags of the security group.
	// Tags of the security group.
	// +mapType=granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

type SecurityGroupParameters struct {

	// Description of the security group.
	// Description of the security group.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Name of the security group to be queried.
	// Name of the security group to be queried.
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Project ID of the security group.
	// Project ID of the security group.
	// +kubebuilder:validation:Optional
	ProjectID *float64 `json:"projectId,omitempty" tf:"project_id,omitempty"`

	// Tags of the security group.
	// Tags of the security group.
	// +kubebuilder:validation:Optional
	// +mapType=granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

// SecurityGroupSpec defines the desired state of SecurityGroup
type SecurityGroupSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     SecurityGroupParameters `json:"forProvider"`
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
	InitProvider SecurityGroupInitParameters `json:"initProvider,omitempty"`
}

// SecurityGroupStatus defines the observed state of SecurityGroup.
type SecurityGroupStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        SecurityGroupObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// SecurityGroup is the Schema for the SecurityGroups API. Provides a resource to create security group.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,tencentcloud}
type SecurityGroup struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.name) || (has(self.initProvider) && has(self.initProvider.name))",message="spec.forProvider.name is a required parameter"
	Spec   SecurityGroupSpec   `json:"spec"`
	Status SecurityGroupStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// SecurityGroupList contains a list of SecurityGroups
type SecurityGroupList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []SecurityGroup `json:"items"`
}

// Repository type metadata.
var (
	SecurityGroup_Kind             = "SecurityGroup"
	SecurityGroup_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: SecurityGroup_Kind}.String()
	SecurityGroup_KindAPIVersion   = SecurityGroup_Kind + "." + CRDGroupVersion.String()
	SecurityGroup_GroupVersionKind = CRDGroupVersion.WithKind(SecurityGroup_Kind)
)

func init() {
	SchemeBuilder.Register(&SecurityGroup{}, &SecurityGroupList{})
}
