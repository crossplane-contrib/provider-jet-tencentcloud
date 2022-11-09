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

type VPCObservation struct {
	CreateTime *string `json:"createTime,omitempty" tf:"create_time,omitempty"`

	DefaultRouteTableID *string `json:"defaultRouteTableId,omitempty" tf:"default_route_table_id,omitempty"`

	DockerAssistantCidrs []*string `json:"dockerAssistantCidrs,omitempty" tf:"docker_assistant_cidrs,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	IsDefault *bool `json:"isDefault,omitempty" tf:"is_default,omitempty"`
}

type VPCParameters struct {

	// List of Assistant CIDR, NOTE: Only `NORMAL` typed CIDRs included, check the Docker CIDR by readonly `assistant_docker_cidrs`.
	// +kubebuilder:validation:Optional
	AssistantCidrs []*string `json:"assistantCidrs,omitempty" tf:"assistant_cidrs,omitempty"`

	// A network address block which should be a subnet of the three internal network segments (10.0.0.0/16, 172.16.0.0/12 and 192.168.0.0/16).
	// +kubebuilder:validation:Required
	CidrBlock *string `json:"cidrBlock" tf:"cidr_block,omitempty"`

	// The DNS server list of the VPC. And you can specify 0 to 5 servers to this list.
	// +kubebuilder:validation:Optional
	DNSServers []*string `json:"dnsServers,omitempty" tf:"dns_servers,omitempty"`

	// Indicates whether VPC multicast is enabled. The default value is 'true'.
	// +kubebuilder:validation:Optional
	IsMulticast *bool `json:"isMulticast,omitempty" tf:"is_multicast,omitempty"`

	// The name of the VPC.
	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// Tags of the VPC.
	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

// VPCSpec defines the desired state of VPC
type VPCSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     VPCParameters `json:"forProvider"`
}

// VPCStatus defines the observed state of VPC.
type VPCStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        VPCObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// VPC is the Schema for the VPCs API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,tencentcloudjet}
type VPC struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              VPCSpec   `json:"spec"`
	Status            VPCStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// VPCList contains a list of VPCs
type VPCList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []VPC `json:"items"`
}

// Repository type metadata.
var (
	VPC_Kind             = "VPC"
	VPC_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: VPC_Kind}.String()
	VPC_KindAPIVersion   = VPC_Kind + "." + CRDGroupVersion.String()
	VPC_GroupVersionKind = CRDGroupVersion.WithKind(VPC_Kind)
)

func init() {
	SchemeBuilder.Register(&VPC{}, &VPCList{})
}
