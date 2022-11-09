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

type VPNSSLClientObservation struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type VPNSSLClientParameters struct {

	// The name of ssl vpn client to be created.
	// +kubebuilder:validation:Required
	SSLVPNClientName *string `json:"sslVpnClientName" tf:"ssl_vpn_client_name,omitempty"`

	// VPN ssl server id.
	// +kubebuilder:validation:Required
	SSLVPNServerID *string `json:"sslVpnServerId" tf:"ssl_vpn_server_id,omitempty"`
}

// VPNSSLClientSpec defines the desired state of VPNSSLClient
type VPNSSLClientSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     VPNSSLClientParameters `json:"forProvider"`
}

// VPNSSLClientStatus defines the observed state of VPNSSLClient.
type VPNSSLClientStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        VPNSSLClientObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// VPNSSLClient is the Schema for the VPNSSLClients API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,tencentcloudjet}
type VPNSSLClient struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              VPNSSLClientSpec   `json:"spec"`
	Status            VPNSSLClientStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// VPNSSLClientList contains a list of VPNSSLClients
type VPNSSLClientList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []VPNSSLClient `json:"items"`
}

// Repository type metadata.
var (
	VPNSSLClient_Kind             = "VPNSSLClient"
	VPNSSLClient_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: VPNSSLClient_Kind}.String()
	VPNSSLClient_KindAPIVersion   = VPNSSLClient_Kind + "." + CRDGroupVersion.String()
	VPNSSLClient_GroupVersionKind = CRDGroupVersion.WithKind(VPNSSLClient_Kind)
)

func init() {
	SchemeBuilder.Register(&VPNSSLClient{}, &VPNSSLClientList{})
}
