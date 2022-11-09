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
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	InternalEndPoint *string `json:"internalEndPoint,omitempty" tf:"internal_end_point,omitempty"`

	PublicDomain *string `json:"publicDomain,omitempty" tf:"public_domain,omitempty"`

	PublicStatus *string `json:"publicStatus,omitempty" tf:"public_status,omitempty"`

	Replications []ReplicationsObservation `json:"replications,omitempty" tf:"replications,omitempty"`

	SecurityPolicy []SecurityPolicyObservation `json:"securityPolicy,omitempty" tf:"security_policy,omitempty"`

	Status *string `json:"status,omitempty" tf:"status,omitempty"`
}

type InstanceParameters struct {

	// Indicate to delete the COS bucket which is auto-created with the instance or not.
	// +kubebuilder:validation:Optional
	DeleteBucket *bool `json:"deleteBucket,omitempty" tf:"delete_bucket,omitempty"`

	// TCR types. Valid values are: `standard`, `basic`, `premium`.
	// +kubebuilder:validation:Required
	InstanceType *string `json:"instanceType" tf:"instance_type,omitempty"`

	// Name of the TCR instance.
	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// Control public network access.
	// +kubebuilder:validation:Optional
	OpenPublicOperation *bool `json:"openPublicOperation,omitempty" tf:"open_public_operation,omitempty"`

	// Specify List of instance Replications, premium only. The available [source region list](https://www.tencentcloud.com/document/api/1051/41101) is here.
	// +kubebuilder:validation:Optional
	Replications []ReplicationsParameters `json:"replications,omitempty" tf:"replications,omitempty"`

	// Public network access allowlist policies of the TCR instance. Only available when `open_public_operation` is `true`.
	// +kubebuilder:validation:Optional
	SecurityPolicy []SecurityPolicyParameters `json:"securityPolicy,omitempty" tf:"security_policy,omitempty"`

	// The available tags within this TCR instance.
	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

type ReplicationsObservation struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type ReplicationsParameters struct {

	// Replication region ID, check the example at the top of page to find out id of region.
	// +kubebuilder:validation:Optional
	RegionID *float64 `json:"regionId,omitempty" tf:"region_id,omitempty"`

	// Specify whether to sync TCR cloud tags to COS Bucket. NOTE: You have to specify when adding, modifying will be ignored for now.
	// +kubebuilder:validation:Optional
	SynTag *bool `json:"synTag,omitempty" tf:"syn_tag,omitempty"`
}

type SecurityPolicyObservation struct {
	Index *float64 `json:"index,omitempty" tf:"index,omitempty"`

	Version *string `json:"version,omitempty" tf:"version,omitempty"`
}

type SecurityPolicyParameters struct {

	// The public network IP address of the access source.
	// +kubebuilder:validation:Optional
	CidrBlock *string `json:"cidrBlock,omitempty" tf:"cidr_block,omitempty"`

	// Remarks of policy.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`
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