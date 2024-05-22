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

type SubApplicationInitParameters struct {

	// Sub application description.
	// Sub application description.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Sub application name, which can contain up to 64 letters, digits, underscores, and hyphens (such as test_ABC-123) and must be unique under a user.
	// Sub application name, which can contain up to 64 letters, digits, underscores, and hyphens (such as test_ABC-123) and must be unique under a user.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Sub appliaction status.
	// Sub appliaction status.
	Status *string `json:"status,omitempty" tf:"status,omitempty"`
}

type SubApplicationObservation struct {

	// The time when the sub application was created.
	// The time when the sub application was created.
	CreateTime *string `json:"createTime,omitempty" tf:"create_time,omitempty"`

	// Sub application description.
	// Sub application description.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// ID of the resource.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Sub application name, which can contain up to 64 letters, digits, underscores, and hyphens (such as test_ABC-123) and must be unique under a user.
	// Sub application name, which can contain up to 64 letters, digits, underscores, and hyphens (such as test_ABC-123) and must be unique under a user.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Sub appliaction status.
	// Sub appliaction status.
	Status *string `json:"status,omitempty" tf:"status,omitempty"`
}

type SubApplicationParameters struct {

	// Sub application description.
	// Sub application description.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Sub application name, which can contain up to 64 letters, digits, underscores, and hyphens (such as test_ABC-123) and must be unique under a user.
	// Sub application name, which can contain up to 64 letters, digits, underscores, and hyphens (such as test_ABC-123) and must be unique under a user.
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Sub appliaction status.
	// Sub appliaction status.
	// +kubebuilder:validation:Optional
	Status *string `json:"status,omitempty" tf:"status,omitempty"`
}

// SubApplicationSpec defines the desired state of SubApplication
type SubApplicationSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     SubApplicationParameters `json:"forProvider"`
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
	InitProvider SubApplicationInitParameters `json:"initProvider,omitempty"`
}

// SubApplicationStatus defines the observed state of SubApplication.
type SubApplicationStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        SubApplicationObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// SubApplication is the Schema for the SubApplications API. Provide a resource to create a VOD sub application.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,tencentcloud}
type SubApplication struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.name) || (has(self.initProvider) && has(self.initProvider.name))",message="spec.forProvider.name is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.status) || (has(self.initProvider) && has(self.initProvider.status))",message="spec.forProvider.status is a required parameter"
	Spec   SubApplicationSpec   `json:"spec"`
	Status SubApplicationStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// SubApplicationList contains a list of SubApplications
type SubApplicationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []SubApplication `json:"items"`
}

// Repository type metadata.
var (
	SubApplication_Kind             = "SubApplication"
	SubApplication_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: SubApplication_Kind}.String()
	SubApplication_KindAPIVersion   = SubApplication_Kind + "." + CRDGroupVersion.String()
	SubApplication_GroupVersionKind = CRDGroupVersion.WithKind(SubApplication_Kind)
)

func init() {
	SchemeBuilder.Register(&SubApplication{}, &SubApplicationList{})
}
