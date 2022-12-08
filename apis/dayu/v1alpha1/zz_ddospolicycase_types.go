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

type DdosPolicyCaseObservation struct {
	CreateTime *string `json:"createTime,omitempty" tf:"create_time,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	SceneID *string `json:"sceneId,omitempty" tf:"scene_id,omitempty"`
}

type DdosPolicyCaseParameters struct {

	// App protocol set of the DDoS policy case.
	// +kubebuilder:validation:Required
	AppProtocols []*string `json:"appProtocols" tf:"app_protocols,omitempty"`

	// App type of the DDoS policy case. Valid values: `WEB`, `GAME`, `APP` and `OTHER`.
	// +kubebuilder:validation:Required
	AppType *string `json:"appType" tf:"app_type,omitempty"`

	// Indicate whether the service involves overseas or not. Valid values: `no` and `yes`.
	// +kubebuilder:validation:Required
	HasAbroad *string `json:"hasAbroad" tf:"has_abroad,omitempty"`

	// Indicate whether the service actively initiates TCP requests or not. Valid values: `no` and `yes`.
	// +kubebuilder:validation:Required
	HasInitiateTCP *string `json:"hasInitiateTcp" tf:"has_initiate_tcp,omitempty"`

	// Indicate whether the actively initiate UDP requests or not. Valid values: `no` and `yes`.
	// +kubebuilder:validation:Optional
	HasInitiateUDP *string `json:"hasInitiateUdp,omitempty" tf:"has_initiate_udp,omitempty"`

	// Indicate whether the service involves VPN service or not. Valid values: `no` and `yes`.
	// +kubebuilder:validation:Optional
	HasVPN *string `json:"hasVpn,omitempty" tf:"has_vpn,omitempty"`

	// The max length of TCP message package, valid value length should be greater than 0 and less than 1500. It should be greater than `min_tcp_package_len`.
	// +kubebuilder:validation:Optional
	MaxTCPPackageLen *string `json:"maxTcpPackageLen,omitempty" tf:"max_tcp_package_len,omitempty"`

	// The max length of UDP message package, valid value length should be greater than 0 and less than 1500. It should be greater than `min_udp_package_len`.
	// +kubebuilder:validation:Optional
	MaxUDPPackageLen *string `json:"maxUdpPackageLen,omitempty" tf:"max_udp_package_len,omitempty"`

	// The minimum length of TCP message package, valid value length should be greater than 0 and less than 1500.
	// +kubebuilder:validation:Optional
	MinTCPPackageLen *string `json:"minTcpPackageLen,omitempty" tf:"min_tcp_package_len,omitempty"`

	// The minimum length of UDP message package, valid value length should be greater than 0 and less than 1500.
	// +kubebuilder:validation:Optional
	MinUDPPackageLen *string `json:"minUdpPackageLen,omitempty" tf:"min_udp_package_len,omitempty"`

	// The port that actively initiates TCP requests. Valid value ranges: (1~65535).
	// +kubebuilder:validation:Optional
	PeerTCPPort *string `json:"peerTcpPort,omitempty" tf:"peer_tcp_port,omitempty"`

	// The port that actively initiates UDP requests. Valid value ranges: (1~65535).
	// +kubebuilder:validation:Optional
	PeerUDPPort *string `json:"peerUdpPort,omitempty" tf:"peer_udp_port,omitempty"`

	// Platform set of the DDoS policy case.
	// +kubebuilder:validation:Required
	PlatformTypes []*string `json:"platformTypes" tf:"platform_types,omitempty"`

	// Type of the resource that the DDoS policy case works for. Valid values: `bgpip`, `bgp` and `bgp-multip`.
	// +kubebuilder:validation:Required
	ResourceType *string `json:"resourceType" tf:"resource_type,omitempty"`

	// End port of the TCP service. Valid value ranges: (0~65535). It must be greater than `tcp_start_port`.
	// +kubebuilder:validation:Required
	TCPEndPort *string `json:"tcpEndPort" tf:"tcp_end_port,omitempty"`

	// The fixed signature of TCP protocol load, valid value length is range from 1 to 512.
	// +kubebuilder:validation:Optional
	TCPFootprint *string `json:"tcpFootprint,omitempty" tf:"tcp_footprint,omitempty"`

	// Start port of the TCP service. Valid value ranges: (0~65535).
	// +kubebuilder:validation:Required
	TCPStartPort *string `json:"tcpStartPort" tf:"tcp_start_port,omitempty"`

	// End port of the UDP service. Valid value ranges: (0~65535). It must be greater than `udp_start_port`.
	// +kubebuilder:validation:Required
	UDPEndPort *string `json:"udpEndPort" tf:"udp_end_port,omitempty"`

	// The fixed signature of TCP protocol load, valid value length is range from 1 to 512.
	// +kubebuilder:validation:Optional
	UDPFootprint *string `json:"udpFootprint,omitempty" tf:"udp_footprint,omitempty"`

	// Start port of the UDP service. Valid value ranges: (0~65535).
	// +kubebuilder:validation:Required
	UDPStartPort *string `json:"udpStartPort" tf:"udp_start_port,omitempty"`

	// Web API url set.
	// +kubebuilder:validation:Required
	WebAPIUrls []*string `json:"webApiUrls" tf:"web_api_urls,omitempty"`
}

// DdosPolicyCaseSpec defines the desired state of DdosPolicyCase
type DdosPolicyCaseSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     DdosPolicyCaseParameters `json:"forProvider"`
}

// DdosPolicyCaseStatus defines the observed state of DdosPolicyCase.
type DdosPolicyCaseStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        DdosPolicyCaseObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// DdosPolicyCase is the Schema for the DdosPolicyCases API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,tencentcloudjet}
type DdosPolicyCase struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              DdosPolicyCaseSpec   `json:"spec"`
	Status            DdosPolicyCaseStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// DdosPolicyCaseList contains a list of DdosPolicyCases
type DdosPolicyCaseList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []DdosPolicyCase `json:"items"`
}

// Repository type metadata.
var (
	DdosPolicyCase_Kind             = "DdosPolicyCase"
	DdosPolicyCase_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: DdosPolicyCase_Kind}.String()
	DdosPolicyCase_KindAPIVersion   = DdosPolicyCase_Kind + "." + CRDGroupVersion.String()
	DdosPolicyCase_GroupVersionKind = CRDGroupVersion.WithKind(DdosPolicyCase_Kind)
)

func init() {
	SchemeBuilder.Register(&DdosPolicyCase{}, &DdosPolicyCaseList{})
}