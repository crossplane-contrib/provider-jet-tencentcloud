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

type GrafanaInstanceObservation struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	InstanceID *string `json:"instanceId,omitempty" tf:"instance_id,omitempty"`

	InstanceStatus *float64 `json:"instanceStatus,omitempty" tf:"instance_status,omitempty"`

	RootURL *string `json:"rootUrl,omitempty" tf:"root_url,omitempty"`
}

type GrafanaInstanceParameters struct {

	// Control whether grafana could be accessed by internet.
	// +kubebuilder:validation:Optional
	EnableInternet *bool `json:"enableInternet,omitempty" tf:"enable_internet,omitempty"`

	// Grafana server admin password.
	// +kubebuilder:validation:Optional
	GrafanaInitPassword *string `json:"grafanaInitPassword,omitempty" tf:"grafana_init_password,omitempty"`

	// Instance name.
	// +kubebuilder:validation:Required
	InstanceName *string `json:"instanceName" tf:"instance_name,omitempty"`

	// Subnet Id array.
	// +kubebuilder:validation:Optional
	SubnetIds []*string `json:"subnetIds,omitempty" tf:"subnet_ids,omitempty"`

	// Tag description list.
	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// Vpc Id.
	// +crossplane:generate:reference:type=github.com/crossplane-contrib/provider-tencentcloud/apis/vpc/v1alpha1.VPC
	// +kubebuilder:validation:Optional
	VPCID *string `json:"vpcId,omitempty" tf:"vpc_id,omitempty"`

	// +kubebuilder:validation:Optional
	VPCIDRef *v1.Reference `json:"vpcidRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	VPCIDSelector *v1.Selector `json:"vpcidSelector,omitempty" tf:"-"`
}

// GrafanaInstanceSpec defines the desired state of GrafanaInstance
type GrafanaInstanceSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     GrafanaInstanceParameters `json:"forProvider"`
}

// GrafanaInstanceStatus defines the observed state of GrafanaInstance.
type GrafanaInstanceStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        GrafanaInstanceObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// GrafanaInstance is the Schema for the GrafanaInstances API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,tencentcloudjet}
type GrafanaInstance struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              GrafanaInstanceSpec   `json:"spec"`
	Status            GrafanaInstanceStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// GrafanaInstanceList contains a list of GrafanaInstances
type GrafanaInstanceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []GrafanaInstance `json:"items"`
}

// Repository type metadata.
var (
	GrafanaInstance_Kind             = "GrafanaInstance"
	GrafanaInstance_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: GrafanaInstance_Kind}.String()
	GrafanaInstance_KindAPIVersion   = GrafanaInstance_Kind + "." + CRDGroupVersion.String()
	GrafanaInstance_GroupVersionKind = CRDGroupVersion.WithKind(GrafanaInstance_Kind)
)

func init() {
	SchemeBuilder.Register(&GrafanaInstance{}, &GrafanaInstanceList{})
}
