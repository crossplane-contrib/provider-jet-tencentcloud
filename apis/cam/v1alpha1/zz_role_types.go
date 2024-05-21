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

type RoleInitParameters struct {

	// Indicates whether the CAM role can login or not.
	// Indicates whether the CAM role can login or not.
	ConsoleLogin *bool `json:"consoleLogin,omitempty" tf:"console_login,omitempty"`

	// Description of the CAM role.
	// Description of the CAM role.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Document of the CAM role. The syntax refers to CAM POLICY. The elements in json claimed supporting two types as string and array only support type array; 2.
	// Document of the CAM role. The syntax refers to [CAM POLICY](https://intl.cloud.tencent.com/document/product/598/10604). The elements in json claimed supporting two types as `string` and `array` only support type `array`; 2.
	Document *string `json:"document,omitempty" tf:"document,omitempty"`

	// A list of tags used to associate different resources.
	// A list of tags used to associate different resources.
	// +mapType=granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

type RoleObservation struct {

	// Indicates whether the CAM role can login or not.
	// Indicates whether the CAM role can login or not.
	ConsoleLogin *bool `json:"consoleLogin,omitempty" tf:"console_login,omitempty"`

	// Create time of the CAM role.
	// Create time of the CAM role.
	CreateTime *string `json:"createTime,omitempty" tf:"create_time,omitempty"`

	// Description of the CAM role.
	// Description of the CAM role.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Document of the CAM role. The syntax refers to CAM POLICY. The elements in json claimed supporting two types as string and array only support type array; 2.
	// Document of the CAM role. The syntax refers to [CAM POLICY](https://intl.cloud.tencent.com/document/product/598/10604). The elements in json claimed supporting two types as `string` and `array` only support type `array`; 2.
	Document *string `json:"document,omitempty" tf:"document,omitempty"`

	// ID of the resource.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// A list of tags used to associate different resources.
	// A list of tags used to associate different resources.
	// +mapType=granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// The last update time of the CAM role.
	// The last update time of the CAM role.
	UpdateTime *string `json:"updateTime,omitempty" tf:"update_time,omitempty"`
}

type RoleParameters struct {

	// Indicates whether the CAM role can login or not.
	// Indicates whether the CAM role can login or not.
	// +kubebuilder:validation:Optional
	ConsoleLogin *bool `json:"consoleLogin,omitempty" tf:"console_login,omitempty"`

	// Description of the CAM role.
	// Description of the CAM role.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Document of the CAM role. The syntax refers to CAM POLICY. The elements in json claimed supporting two types as string and array only support type array; 2.
	// Document of the CAM role. The syntax refers to [CAM POLICY](https://intl.cloud.tencent.com/document/product/598/10604). The elements in json claimed supporting two types as `string` and `array` only support type `array`; 2.
	// +kubebuilder:validation:Optional
	Document *string `json:"document,omitempty" tf:"document,omitempty"`

	// A list of tags used to associate different resources.
	// A list of tags used to associate different resources.
	// +kubebuilder:validation:Optional
	// +mapType=granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

// RoleSpec defines the desired state of Role
type RoleSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     RoleParameters `json:"forProvider"`
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
	InitProvider RoleInitParameters `json:"initProvider,omitempty"`
}

// RoleStatus defines the observed state of Role.
type RoleStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        RoleObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// Role is the Schema for the Roles API. Provides a resource to create a CAM role.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,tencentcloud}
type Role struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.document) || (has(self.initProvider) && has(self.initProvider.document))",message="spec.forProvider.document is a required parameter"
	Spec   RoleSpec   `json:"spec"`
	Status RoleStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// RoleList contains a list of Roles
type RoleList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Role `json:"items"`
}

// Repository type metadata.
var (
	Role_Kind             = "Role"
	Role_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Role_Kind}.String()
	Role_KindAPIVersion   = Role_Kind + "." + CRDGroupVersion.String()
	Role_GroupVersionKind = CRDGroupVersion.WithKind(Role_Kind)
)

func init() {
	SchemeBuilder.Register(&Role{}, &RoleList{})
}
