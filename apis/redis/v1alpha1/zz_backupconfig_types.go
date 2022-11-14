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

type BackupConfigObservation struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type BackupConfigParameters struct {

	// Specifys which day the backup action should take place. Valid values: `Monday`, `Tuesday`, `Wednesday`, `Thursday`, `Friday`, `Saturday` and `Sunday`.
	// +kubebuilder:validation:Optional
	BackupPeriod []*string `json:"backupPeriod,omitempty" tf:"backup_period,omitempty"`

	// Specifys what time the backup action should take place. And the time interval should be one hour.
	// +kubebuilder:validation:Required
	BackupTime *string `json:"backupTime" tf:"backup_time,omitempty"`

	// ID of a redis instance to which the policy will be applied.
	// +crossplane:generate:reference:type=Instance
	// +kubebuilder:validation:Optional
	RedisID *string `json:"redisId,omitempty" tf:"redis_id,omitempty"`

	// +kubebuilder:validation:Optional
	RedisIDRef *v1.Reference `json:"redisIdRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	RedisIDSelector *v1.Selector `json:"redisIdSelector,omitempty" tf:"-"`
}

// BackupConfigSpec defines the desired state of BackupConfig
type BackupConfigSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     BackupConfigParameters `json:"forProvider"`
}

// BackupConfigStatus defines the observed state of BackupConfig.
type BackupConfigStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        BackupConfigObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// BackupConfig is the Schema for the BackupConfigs API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,tencentcloudjet}
type BackupConfig struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              BackupConfigSpec   `json:"spec"`
	Status            BackupConfigStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// BackupConfigList contains a list of BackupConfigs
type BackupConfigList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []BackupConfig `json:"items"`
}

// Repository type metadata.
var (
	BackupConfig_Kind             = "BackupConfig"
	BackupConfig_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: BackupConfig_Kind}.String()
	BackupConfig_KindAPIVersion   = BackupConfig_Kind + "." + CRDGroupVersion.String()
	BackupConfig_GroupVersionKind = CRDGroupVersion.WithKind(BackupConfig_Kind)
)

func init() {
	SchemeBuilder.Register(&BackupConfig{}, &BackupConfigList{})
}
