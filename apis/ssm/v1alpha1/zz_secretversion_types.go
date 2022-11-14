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

type SecretVersionObservation struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type SecretVersionParameters struct {

	// The base64-encoded binary secret. secret_binary and secret_string must be set only one, and the maximum support is 4096 bytes. When secret status is `Disabled`, this field will not update anymore.
	// +kubebuilder:validation:Optional
	SecretBinary *string `json:"secretBinary,omitempty" tf:"secret_binary,omitempty"`

	// Name of secret which cannot be repeated in the same region. The maximum length is 128 bytes. The name can only contain English letters, numbers, underscore and hyphen '-'. The first character must be a letter or number.
	// +kubebuilder:validation:Required
	SecretName *string `json:"secretName" tf:"secret_name,omitempty"`

	// The string text of secret. secret_binary and secret_string must be set only one, and the maximum support is 4096 bytes. When secret status is `Disabled`, this field will not update anymore.
	// +kubebuilder:validation:Optional
	SecretString *string `json:"secretString,omitempty" tf:"secret_string,omitempty"`

	// Version of secret. The maximum length is 64 bytes. The version_id can only contain English letters, numbers, underscore and hyphen '-'. The first character must be a letter or number.
	// +kubebuilder:validation:Required
	VersionID *string `json:"versionId" tf:"version_id,omitempty"`
}

// SecretVersionSpec defines the desired state of SecretVersion
type SecretVersionSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     SecretVersionParameters `json:"forProvider"`
}

// SecretVersionStatus defines the observed state of SecretVersion.
type SecretVersionStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        SecretVersionObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// SecretVersion is the Schema for the SecretVersions API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,tencentcloudjet}
type SecretVersion struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              SecretVersionSpec   `json:"spec"`
	Status            SecretVersionStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// SecretVersionList contains a list of SecretVersions
type SecretVersionList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []SecretVersion `json:"items"`
}

// Repository type metadata.
var (
	SecretVersion_Kind             = "SecretVersion"
	SecretVersion_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: SecretVersion_Kind}.String()
	SecretVersion_KindAPIVersion   = SecretVersion_Kind + "." + CRDGroupVersion.String()
	SecretVersion_GroupVersionKind = CRDGroupVersion.WithKind(SecretVersion_Kind)
)

func init() {
	SchemeBuilder.Register(&SecretVersion{}, &SecretVersionList{})
}
