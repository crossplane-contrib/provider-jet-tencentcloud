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

type Layer7ListenerObservation struct {
	CreateTime *string `json:"createTime,omitempty" tf:"create_time,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	Status *float64 `json:"status,omitempty" tf:"status,omitempty"`
}

type Layer7ListenerParameters struct {

	// Authentication type of the layer7 listener. `0` is one-way authentication and `1` is mutual authentication. NOTES: Only supports listeners of `HTTPS` protocol.
	// +kubebuilder:validation:Optional
	AuthType *float64 `json:"authType,omitempty" tf:"auth_type,omitempty"`

	// Certificate ID of the layer7 listener. NOTES: Only supports listeners of `HTTPS` protocol.
	// +kubebuilder:validation:Optional
	CertificateID *string `json:"certificateId,omitempty" tf:"certificate_id,omitempty"`

	// ID of the client certificate. Set only when `auth_type` is specified as mutual authentication. NOTES: Only supports listeners of `HTTPS` protocol.
	// +kubebuilder:validation:Optional
	ClientCertificateID *string `json:"clientCertificateId,omitempty" tf:"client_certificate_id,omitempty"`

	// ID list of the client certificate. Set only when `auth_type` is specified as mutual authentication. NOTES: Only supports listeners of `HTTPS` protocol.
	// +kubebuilder:validation:Optional
	ClientCertificateIds []*string `json:"clientCertificateIds,omitempty" tf:"client_certificate_ids,omitempty"`

	// Protocol type of the forwarding. Valid value: `HTTP` and `HTTPS`. NOTES: Only supports listeners of `HTTPS` protocol.
	// +kubebuilder:validation:Optional
	ForwardProtocol *string `json:"forwardProtocol,omitempty" tf:"forward_protocol,omitempty"`

	// Name of the layer7 listener, the maximum length is 30.
	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// Port of the layer7 listener.
	// +kubebuilder:validation:Required
	Port *float64 `json:"port" tf:"port,omitempty"`

	// Protocol of the layer7 listener. Valid value: `HTTP` and `HTTPS`.
	// +kubebuilder:validation:Required
	Protocol *string `json:"protocol" tf:"protocol,omitempty"`

	// ID of the GAAP proxy.
	// +crossplane:generate:reference:type=Proxy
	// +kubebuilder:validation:Optional
	ProxyID *string `json:"proxyId,omitempty" tf:"proxy_id,omitempty"`

	// +kubebuilder:validation:Optional
	ProxyIDRef *v1.Reference `json:"proxyIdRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	ProxyIDSelector *v1.Selector `json:"proxyIdSelector,omitempty" tf:"-"`
}

// Layer7ListenerSpec defines the desired state of Layer7Listener
type Layer7ListenerSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     Layer7ListenerParameters `json:"forProvider"`
}

// Layer7ListenerStatus defines the observed state of Layer7Listener.
type Layer7ListenerStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        Layer7ListenerObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Layer7Listener is the Schema for the Layer7Listeners API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,tencentcloudjet}
type Layer7Listener struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              Layer7ListenerSpec   `json:"spec"`
	Status            Layer7ListenerStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// Layer7ListenerList contains a list of Layer7Listeners
type Layer7ListenerList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Layer7Listener `json:"items"`
}

// Repository type metadata.
var (
	Layer7Listener_Kind             = "Layer7Listener"
	Layer7Listener_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Layer7Listener_Kind}.String()
	Layer7Listener_KindAPIVersion   = Layer7Listener_Kind + "." + CRDGroupVersion.String()
	Layer7Listener_GroupVersionKind = CRDGroupVersion.WithKind(Layer7Listener_Kind)
)

func init() {
	SchemeBuilder.Register(&Layer7Listener{}, &Layer7ListenerList{})
}