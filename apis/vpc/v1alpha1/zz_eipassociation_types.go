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

type EipAssociationObservation struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type EipAssociationParameters struct {

	// The ID of EIP.
	// +crossplane:generate:reference:type=Eip
	// +kubebuilder:validation:Optional
	EIPID *string `json:"eipId,omitempty" tf:"eip_id,omitempty"`

	// +kubebuilder:validation:Optional
	EIPIDRef *v1.Reference `json:"eipidRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	EIPIDSelector *v1.Selector `json:"eipidSelector,omitempty" tf:"-"`

	// The CVM or CLB instance id going to bind with the EIP. This field is conflict with `network_interface_id` and `private_ip fields`.
	// +kubebuilder:validation:Optional
	InstanceID *string `json:"instanceId,omitempty" tf:"instance_id,omitempty"`

	// Indicates the network interface id like `eni-xxxxxx`. This field is conflict with `instance_id`.
	// +kubebuilder:validation:Optional
	NetworkInterfaceID *string `json:"networkInterfaceId,omitempty" tf:"network_interface_id,omitempty"`

	// Indicates an IP belongs to the `network_interface_id`. This field is conflict with `instance_id`.
	// +kubebuilder:validation:Optional
	PrivateIP *string `json:"privateIp,omitempty" tf:"private_ip,omitempty"`
}

// EipAssociationSpec defines the desired state of EipAssociation
type EipAssociationSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     EipAssociationParameters `json:"forProvider"`
}

// EipAssociationStatus defines the observed state of EipAssociation.
type EipAssociationStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        EipAssociationObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// EipAssociation is the Schema for the EipAssociations API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,tencentcloudjet}
type EipAssociation struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              EipAssociationSpec   `json:"spec"`
	Status            EipAssociationStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// EipAssociationList contains a list of EipAssociations
type EipAssociationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []EipAssociation `json:"items"`
}

// Repository type metadata.
var (
	EipAssociation_Kind             = "EipAssociation"
	EipAssociation_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: EipAssociation_Kind}.String()
	EipAssociation_KindAPIVersion   = EipAssociation_Kind + "." + CRDGroupVersion.String()
	EipAssociation_GroupVersionKind = CRDGroupVersion.WithKind(EipAssociation_Kind)
)

func init() {
	SchemeBuilder.Register(&EipAssociation{}, &EipAssociationList{})
}