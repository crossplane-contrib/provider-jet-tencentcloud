// SPDX-FileCopyrightText: 2023 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

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

type DynamicIndexInitParameters struct {

	// Whether to take effect. Default value: true.
	// index automatic configuration switch.
	Status *bool `json:"status,omitempty" tf:"status,omitempty"`
}

type DynamicIndexObservation struct {

	// Whether to take effect. Default value: true.
	// index automatic configuration switch.
	Status *bool `json:"status,omitempty" tf:"status,omitempty"`
}

type DynamicIndexParameters struct {

	// Whether to take effect. Default value: true.
	// index automatic configuration switch.
	// +kubebuilder:validation:Optional
	Status *bool `json:"status" tf:"status,omitempty"`
}

type FullTextInitParameters struct {

	// Case sensitivity.
	// Case sensitivity.
	CaseSensitive *bool `json:"caseSensitive,omitempty" tf:"case_sensitive,omitempty"`

	// Whether Chinese characters are contained.
	// Whether Chinese characters are contained.
	ContainZH *bool `json:"containZH,omitempty" tf:"contain_z_h,omitempty"`

	// Full-Text index delimiter. Each character in the string represents a delimiter.
	// Full-Text index delimiter. Each character in the string represents a delimiter.
	Tokenizer *string `json:"tokenizer,omitempty" tf:"tokenizer,omitempty"`
}

type FullTextObservation struct {

	// Case sensitivity.
	// Case sensitivity.
	CaseSensitive *bool `json:"caseSensitive,omitempty" tf:"case_sensitive,omitempty"`

	// Whether Chinese characters are contained.
	// Whether Chinese characters are contained.
	ContainZH *bool `json:"containZH,omitempty" tf:"contain_z_h,omitempty"`

	// Full-Text index delimiter. Each character in the string represents a delimiter.
	// Full-Text index delimiter. Each character in the string represents a delimiter.
	Tokenizer *string `json:"tokenizer,omitempty" tf:"tokenizer,omitempty"`
}

type FullTextParameters struct {

	// Case sensitivity.
	// Case sensitivity.
	// +kubebuilder:validation:Optional
	CaseSensitive *bool `json:"caseSensitive" tf:"case_sensitive,omitempty"`

	// Whether Chinese characters are contained.
	// Whether Chinese characters are contained.
	// +kubebuilder:validation:Optional
	ContainZH *bool `json:"containZH" tf:"contain_z_h,omitempty"`

	// Full-Text index delimiter. Each character in the string represents a delimiter.
	// Full-Text index delimiter. Each character in the string represents a delimiter.
	// +kubebuilder:validation:Optional
	Tokenizer *string `json:"tokenizer" tf:"tokenizer,omitempty"`
}

type IndexInitParameters struct {

	// Internal field marker of full-text index. Default value: false. Valid value: false: excluding internal fields; true: including internal fields.
	// Internal field marker of full-text index. Default value: false. Valid value: false: excluding internal fields; true: including internal fields.
	IncludeInternalFields *bool `json:"includeInternalFields,omitempty" tf:"include_internal_fields,omitempty"`

	// Metadata flag. Default value: 0. Valid value: 0: full-text index (including the metadata field with key-value index enabled); 1: full-text index (including all metadata fields); 2: full-text index (excluding metadata fields)..
	// Metadata flag. Default value: 0. Valid value: 0: full-text index (including the metadata field with key-value index enabled); 1: full-text index (including all metadata fields); 2: full-text index (excluding metadata fields)..
	MetadataFlag *float64 `json:"metadataFlag,omitempty" tf:"metadata_flag,omitempty"`

	// Index rule.
	// Index rule.
	Rule []RuleInitParameters `json:"rule,omitempty" tf:"rule,omitempty"`

	// Whether to take effect. Default value: true.
	// Whether to take effect. Default value: true.
	Status *bool `json:"status,omitempty" tf:"status,omitempty"`
}

type IndexObservation struct {

	// ID of the resource.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Internal field marker of full-text index. Default value: false. Valid value: false: excluding internal fields; true: including internal fields.
	// Internal field marker of full-text index. Default value: false. Valid value: false: excluding internal fields; true: including internal fields.
	IncludeInternalFields *bool `json:"includeInternalFields,omitempty" tf:"include_internal_fields,omitempty"`

	// Metadata flag. Default value: 0. Valid value: 0: full-text index (including the metadata field with key-value index enabled); 1: full-text index (including all metadata fields); 2: full-text index (excluding metadata fields)..
	// Metadata flag. Default value: 0. Valid value: 0: full-text index (including the metadata field with key-value index enabled); 1: full-text index (including all metadata fields); 2: full-text index (excluding metadata fields)..
	MetadataFlag *float64 `json:"metadataFlag,omitempty" tf:"metadata_flag,omitempty"`

	// Index rule.
	// Index rule.
	Rule []RuleObservation `json:"rule,omitempty" tf:"rule,omitempty"`

	// Whether to take effect. Default value: true.
	// Whether to take effect. Default value: true.
	Status *bool `json:"status,omitempty" tf:"status,omitempty"`

	// Log topic ID.
	// Log topic ID.
	TopicID *string `json:"topicId,omitempty" tf:"topic_id,omitempty"`
}

type IndexParameters struct {

	// Internal field marker of full-text index. Default value: false. Valid value: false: excluding internal fields; true: including internal fields.
	// Internal field marker of full-text index. Default value: false. Valid value: false: excluding internal fields; true: including internal fields.
	// +kubebuilder:validation:Optional
	IncludeInternalFields *bool `json:"includeInternalFields,omitempty" tf:"include_internal_fields,omitempty"`

	// Metadata flag. Default value: 0. Valid value: 0: full-text index (including the metadata field with key-value index enabled); 1: full-text index (including all metadata fields); 2: full-text index (excluding metadata fields)..
	// Metadata flag. Default value: 0. Valid value: 0: full-text index (including the metadata field with key-value index enabled); 1: full-text index (including all metadata fields); 2: full-text index (excluding metadata fields)..
	// +kubebuilder:validation:Optional
	MetadataFlag *float64 `json:"metadataFlag,omitempty" tf:"metadata_flag,omitempty"`

	// Index rule.
	// Index rule.
	// +kubebuilder:validation:Optional
	Rule []RuleParameters `json:"rule,omitempty" tf:"rule,omitempty"`

	// Whether to take effect. Default value: true.
	// Whether to take effect. Default value: true.
	// +kubebuilder:validation:Optional
	Status *bool `json:"status,omitempty" tf:"status,omitempty"`

	// Log topic ID.
	// Log topic ID.
	// +crossplane:generate:reference:type=Topic
	// +kubebuilder:validation:Optional
	TopicID *string `json:"topicId,omitempty" tf:"topic_id,omitempty"`

	// Reference to a Topic to populate topicId.
	// +kubebuilder:validation:Optional
	TopicIDRef *v1.Reference `json:"topicIdRef,omitempty" tf:"-"`

	// Selector for a Topic to populate topicId.
	// +kubebuilder:validation:Optional
	TopicIDSelector *v1.Selector `json:"topicIdSelector,omitempty" tf:"-"`
}

type KeyValueInitParameters struct {

	// Case sensitivity.
	// Case sensitivity.
	CaseSensitive *bool `json:"caseSensitive,omitempty" tf:"case_sensitive,omitempty"`

	// Key-Value pair information of the index to be created. Up to 100 key-value pairs can be configured.
	// Key-Value pair information of the index to be created. Up to 100 key-value pairs can be configured.
	KeyValues []KeyValuesInitParameters `json:"keyValues,omitempty" tf:"key_values,omitempty"`
}

type KeyValueObservation struct {

	// Case sensitivity.
	// Case sensitivity.
	CaseSensitive *bool `json:"caseSensitive,omitempty" tf:"case_sensitive,omitempty"`

	// Key-Value pair information of the index to be created. Up to 100 key-value pairs can be configured.
	// Key-Value pair information of the index to be created. Up to 100 key-value pairs can be configured.
	KeyValues []KeyValuesObservation `json:"keyValues,omitempty" tf:"key_values,omitempty"`
}

type KeyValueParameters struct {

	// Case sensitivity.
	// Case sensitivity.
	// +kubebuilder:validation:Optional
	CaseSensitive *bool `json:"caseSensitive" tf:"case_sensitive,omitempty"`

	// Key-Value pair information of the index to be created. Up to 100 key-value pairs can be configured.
	// Key-Value pair information of the index to be created. Up to 100 key-value pairs can be configured.
	// +kubebuilder:validation:Optional
	KeyValues []KeyValuesParameters `json:"keyValues,omitempty" tf:"key_values,omitempty"`
}

type KeyValuesInitParameters struct {

	// When a key value or metafield index needs to be configured for a field, the metafield Key does not need to be prefixed with TAG. and is consistent with the one when logs are uploaded. TAG. will be prefixed automatically for display in the console..
	// When a key value or metafield index needs to be configured for a field, the metafield Key does not need to be prefixed with __TAG__. and is consistent with the one when logs are uploaded. __TAG__. will be prefixed automatically for display in the console..
	Key *string `json:"key,omitempty" tf:"key,omitempty"`

	// Field index description information.
	// Field index description information.
	Value []ValueInitParameters `json:"value,omitempty" tf:"value,omitempty"`
}

type KeyValuesObservation struct {

	// When a key value or metafield index needs to be configured for a field, the metafield Key does not need to be prefixed with TAG. and is consistent with the one when logs are uploaded. TAG. will be prefixed automatically for display in the console..
	// When a key value or metafield index needs to be configured for a field, the metafield Key does not need to be prefixed with __TAG__. and is consistent with the one when logs are uploaded. __TAG__. will be prefixed automatically for display in the console..
	Key *string `json:"key,omitempty" tf:"key,omitempty"`

	// Field index description information.
	// Field index description information.
	Value []ValueObservation `json:"value,omitempty" tf:"value,omitempty"`
}

type KeyValuesParameters struct {

	// When a key value or metafield index needs to be configured for a field, the metafield Key does not need to be prefixed with TAG. and is consistent with the one when logs are uploaded. TAG. will be prefixed automatically for display in the console..
	// When a key value or metafield index needs to be configured for a field, the metafield Key does not need to be prefixed with __TAG__. and is consistent with the one when logs are uploaded. __TAG__. will be prefixed automatically for display in the console..
	// +kubebuilder:validation:Optional
	Key *string `json:"key" tf:"key,omitempty"`

	// Field index description information.
	// Field index description information.
	// +kubebuilder:validation:Optional
	Value []ValueParameters `json:"value,omitempty" tf:"value,omitempty"`
}

type KeyValuesValueInitParameters struct {

	// Whether Chinese characters are contained.
	// Whether Chinese characters are contained.
	ContainZH *bool `json:"containZH,omitempty" tf:"contain_z_h,omitempty"`

	// Whether the analysis feature is enabled for the field.
	// Whether the analysis feature is enabled for the field.
	SQLFlag *bool `json:"sqlFlag,omitempty" tf:"sql_flag,omitempty"`

	// Full-Text index delimiter. Each character in the string represents a delimiter.
	// Field delimiter, which is meaningful only if the field type is text. Each character in the entered string represents a delimiter.
	Tokenizer *string `json:"tokenizer,omitempty" tf:"tokenizer,omitempty"`

	// Field type. Valid values: long, text, double.
	// Field type. Valid values: long, text, double.
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

type KeyValuesValueObservation struct {

	// Whether Chinese characters are contained.
	// Whether Chinese characters are contained.
	ContainZH *bool `json:"containZH,omitempty" tf:"contain_z_h,omitempty"`

	// Whether the analysis feature is enabled for the field.
	// Whether the analysis feature is enabled for the field.
	SQLFlag *bool `json:"sqlFlag,omitempty" tf:"sql_flag,omitempty"`

	// Full-Text index delimiter. Each character in the string represents a delimiter.
	// Field delimiter, which is meaningful only if the field type is text. Each character in the entered string represents a delimiter.
	Tokenizer *string `json:"tokenizer,omitempty" tf:"tokenizer,omitempty"`

	// Field type. Valid values: long, text, double.
	// Field type. Valid values: long, text, double.
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

type KeyValuesValueParameters struct {

	// Whether Chinese characters are contained.
	// Whether Chinese characters are contained.
	// +kubebuilder:validation:Optional
	ContainZH *bool `json:"containZH,omitempty" tf:"contain_z_h,omitempty"`

	// Whether the analysis feature is enabled for the field.
	// Whether the analysis feature is enabled for the field.
	// +kubebuilder:validation:Optional
	SQLFlag *bool `json:"sqlFlag,omitempty" tf:"sql_flag,omitempty"`

	// Full-Text index delimiter. Each character in the string represents a delimiter.
	// Field delimiter, which is meaningful only if the field type is text. Each character in the entered string represents a delimiter.
	// +kubebuilder:validation:Optional
	Tokenizer *string `json:"tokenizer,omitempty" tf:"tokenizer,omitempty"`

	// Field type. Valid values: long, text, double.
	// Field type. Valid values: long, text, double.
	// +kubebuilder:validation:Optional
	Type *string `json:"type" tf:"type,omitempty"`
}

type RuleInitParameters struct {

	// The key value index is automatically configured. If it is empty, it means that the function is not enabled.
	// The key value index is automatically configured. If it is empty, it means that the function is not enabled.
	DynamicIndex []DynamicIndexInitParameters `json:"dynamicIndex,omitempty" tf:"dynamic_index,omitempty"`

	// Full-Text index configuration.
	// Full-Text index configuration.
	FullText []FullTextInitParameters `json:"fullText,omitempty" tf:"full_text,omitempty"`

	// Key-Value index configuration.
	// Key-Value index configuration.
	KeyValue []KeyValueInitParameters `json:"keyValue,omitempty" tf:"key_value,omitempty"`

	// Metafield index configuration.
	// Metafield index configuration.
	Tag []TagInitParameters `json:"tag,omitempty" tf:"tag,omitempty"`
}

type RuleObservation struct {

	// The key value index is automatically configured. If it is empty, it means that the function is not enabled.
	// The key value index is automatically configured. If it is empty, it means that the function is not enabled.
	DynamicIndex []DynamicIndexObservation `json:"dynamicIndex,omitempty" tf:"dynamic_index,omitempty"`

	// Full-Text index configuration.
	// Full-Text index configuration.
	FullText []FullTextObservation `json:"fullText,omitempty" tf:"full_text,omitempty"`

	// Key-Value index configuration.
	// Key-Value index configuration.
	KeyValue []KeyValueObservation `json:"keyValue,omitempty" tf:"key_value,omitempty"`

	// Metafield index configuration.
	// Metafield index configuration.
	Tag []TagObservation `json:"tag,omitempty" tf:"tag,omitempty"`
}

type RuleParameters struct {

	// The key value index is automatically configured. If it is empty, it means that the function is not enabled.
	// The key value index is automatically configured. If it is empty, it means that the function is not enabled.
	// +kubebuilder:validation:Optional
	DynamicIndex []DynamicIndexParameters `json:"dynamicIndex,omitempty" tf:"dynamic_index,omitempty"`

	// Full-Text index configuration.
	// Full-Text index configuration.
	// +kubebuilder:validation:Optional
	FullText []FullTextParameters `json:"fullText,omitempty" tf:"full_text,omitempty"`

	// Key-Value index configuration.
	// Key-Value index configuration.
	// +kubebuilder:validation:Optional
	KeyValue []KeyValueParameters `json:"keyValue,omitempty" tf:"key_value,omitempty"`

	// Metafield index configuration.
	// Metafield index configuration.
	// +kubebuilder:validation:Optional
	Tag []TagParameters `json:"tag,omitempty" tf:"tag,omitempty"`
}

type TagInitParameters struct {

	// Case sensitivity.
	// Case sensitivity.
	CaseSensitive *bool `json:"caseSensitive,omitempty" tf:"case_sensitive,omitempty"`

	// Key-Value pair information of the index to be created. Up to 100 key-value pairs can be configured.
	// Key-Value pair information of the index to be created. Up to 100 key-value pairs can be configured.
	KeyValues []TagKeyValuesInitParameters `json:"keyValues,omitempty" tf:"key_values,omitempty"`
}

type TagKeyValuesInitParameters struct {

	// When a key value or metafield index needs to be configured for a field, the metafield Key does not need to be prefixed with TAG. and is consistent with the one when logs are uploaded. TAG. will be prefixed automatically for display in the console..
	// When a key value or metafield index needs to be configured for a field, the metafield Key does not need to be prefixed with __TAG__. and is consistent with the one when logs are uploaded. __TAG__. will be prefixed automatically for display in the console..
	Key *string `json:"key,omitempty" tf:"key,omitempty"`

	// Field index description information.
	// Field index description information.
	Value []KeyValuesValueInitParameters `json:"value,omitempty" tf:"value,omitempty"`
}

type TagKeyValuesObservation struct {

	// When a key value or metafield index needs to be configured for a field, the metafield Key does not need to be prefixed with TAG. and is consistent with the one when logs are uploaded. TAG. will be prefixed automatically for display in the console..
	// When a key value or metafield index needs to be configured for a field, the metafield Key does not need to be prefixed with __TAG__. and is consistent with the one when logs are uploaded. __TAG__. will be prefixed automatically for display in the console..
	Key *string `json:"key,omitempty" tf:"key,omitempty"`

	// Field index description information.
	// Field index description information.
	Value []KeyValuesValueObservation `json:"value,omitempty" tf:"value,omitempty"`
}

type TagKeyValuesParameters struct {

	// When a key value or metafield index needs to be configured for a field, the metafield Key does not need to be prefixed with TAG. and is consistent with the one when logs are uploaded. TAG. will be prefixed automatically for display in the console..
	// When a key value or metafield index needs to be configured for a field, the metafield Key does not need to be prefixed with __TAG__. and is consistent with the one when logs are uploaded. __TAG__. will be prefixed automatically for display in the console..
	// +kubebuilder:validation:Optional
	Key *string `json:"key" tf:"key,omitempty"`

	// Field index description information.
	// Field index description information.
	// +kubebuilder:validation:Optional
	Value []KeyValuesValueParameters `json:"value,omitempty" tf:"value,omitempty"`
}

type TagObservation struct {

	// Case sensitivity.
	// Case sensitivity.
	CaseSensitive *bool `json:"caseSensitive,omitempty" tf:"case_sensitive,omitempty"`

	// Key-Value pair information of the index to be created. Up to 100 key-value pairs can be configured.
	// Key-Value pair information of the index to be created. Up to 100 key-value pairs can be configured.
	KeyValues []TagKeyValuesObservation `json:"keyValues,omitempty" tf:"key_values,omitempty"`
}

type TagParameters struct {

	// Case sensitivity.
	// Case sensitivity.
	// +kubebuilder:validation:Optional
	CaseSensitive *bool `json:"caseSensitive" tf:"case_sensitive,omitempty"`

	// Key-Value pair information of the index to be created. Up to 100 key-value pairs can be configured.
	// Key-Value pair information of the index to be created. Up to 100 key-value pairs can be configured.
	// +kubebuilder:validation:Optional
	KeyValues []TagKeyValuesParameters `json:"keyValues,omitempty" tf:"key_values,omitempty"`
}

type ValueInitParameters struct {

	// Whether Chinese characters are contained.
	// Whether Chinese characters are contained.
	ContainZH *bool `json:"containZH,omitempty" tf:"contain_z_h,omitempty"`

	// Whether the analysis feature is enabled for the field.
	// Whether the analysis feature is enabled for the field.
	SQLFlag *bool `json:"sqlFlag,omitempty" tf:"sql_flag,omitempty"`

	// Full-Text index delimiter. Each character in the string represents a delimiter.
	// Field delimiter, which is meaningful only if the field type is text. Each character in the entered string represents a delimiter.
	Tokenizer *string `json:"tokenizer,omitempty" tf:"tokenizer,omitempty"`

	// Field type. Valid values: long, text, double.
	// Field type. Valid values: long, text, double.
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

type ValueObservation struct {

	// Whether Chinese characters are contained.
	// Whether Chinese characters are contained.
	ContainZH *bool `json:"containZH,omitempty" tf:"contain_z_h,omitempty"`

	// Whether the analysis feature is enabled for the field.
	// Whether the analysis feature is enabled for the field.
	SQLFlag *bool `json:"sqlFlag,omitempty" tf:"sql_flag,omitempty"`

	// Full-Text index delimiter. Each character in the string represents a delimiter.
	// Field delimiter, which is meaningful only if the field type is text. Each character in the entered string represents a delimiter.
	Tokenizer *string `json:"tokenizer,omitempty" tf:"tokenizer,omitempty"`

	// Field type. Valid values: long, text, double.
	// Field type. Valid values: long, text, double.
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

type ValueParameters struct {

	// Whether Chinese characters are contained.
	// Whether Chinese characters are contained.
	// +kubebuilder:validation:Optional
	ContainZH *bool `json:"containZH,omitempty" tf:"contain_z_h,omitempty"`

	// Whether the analysis feature is enabled for the field.
	// Whether the analysis feature is enabled for the field.
	// +kubebuilder:validation:Optional
	SQLFlag *bool `json:"sqlFlag,omitempty" tf:"sql_flag,omitempty"`

	// Full-Text index delimiter. Each character in the string represents a delimiter.
	// Field delimiter, which is meaningful only if the field type is text. Each character in the entered string represents a delimiter.
	// +kubebuilder:validation:Optional
	Tokenizer *string `json:"tokenizer,omitempty" tf:"tokenizer,omitempty"`

	// Field type. Valid values: long, text, double.
	// Field type. Valid values: long, text, double.
	// +kubebuilder:validation:Optional
	Type *string `json:"type" tf:"type,omitempty"`
}

// IndexSpec defines the desired state of Index
type IndexSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     IndexParameters `json:"forProvider"`
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
	InitProvider IndexInitParameters `json:"initProvider,omitempty"`
}

// IndexStatus defines the observed state of Index.
type IndexStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        IndexObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Index is the Schema for the Indexs API. Provides a resource to create a cls index.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,tencentcloud}
type Index struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              IndexSpec   `json:"spec"`
	Status            IndexStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// IndexList contains a list of Indexs
type IndexList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Index `json:"items"`
}

// Repository type metadata.
var (
	Index_Kind             = "Index"
	Index_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Index_Kind}.String()
	Index_KindAPIVersion   = Index_Kind + "." + CRDGroupVersion.String()
	Index_GroupVersionKind = CRDGroupVersion.WithKind(Index_Kind)
)

func init() {
	SchemeBuilder.Register(&Index{}, &IndexList{})
}
