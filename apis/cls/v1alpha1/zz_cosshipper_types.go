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

type CompressInitParameters struct {

	// Compression format. Valid values: gzip, lzop, none (no compression).
	// Compression format. Valid values: gzip, lzop, none (no compression).
	Format *string `json:"format,omitempty" tf:"format,omitempty"`
}

type CompressObservation struct {

	// Compression format. Valid values: gzip, lzop, none (no compression).
	// Compression format. Valid values: gzip, lzop, none (no compression).
	Format *string `json:"format,omitempty" tf:"format,omitempty"`
}

type CompressParameters struct {

	// Compression format. Valid values: gzip, lzop, none (no compression).
	// Compression format. Valid values: gzip, lzop, none (no compression).
	// +kubebuilder:validation:Optional
	Format *string `json:"format" tf:"format,omitempty"`
}

type ContentInitParameters struct {

	// CSV format content description.Note: this field may return null, indicating that no valid values can be obtained.
	// CSV format content description.Note: this field may return null, indicating that no valid values can be obtained.
	Csv []CsvInitParameters `json:"csv,omitempty" tf:"csv,omitempty"`

	// Compression format. Valid values: gzip, lzop, none (no compression).
	// Content format. Valid values: json, csv.
	Format *string `json:"format,omitempty" tf:"format,omitempty"`

	// JSON format content description.Note: this field may return null, indicating that no valid values can be obtained.
	// JSON format content description.Note: this field may return null, indicating that no valid values can be obtained.
	JSON []JSONInitParameters `json:"json,omitempty" tf:"json,omitempty"`
}

type ContentObservation struct {

	// CSV format content description.Note: this field may return null, indicating that no valid values can be obtained.
	// CSV format content description.Note: this field may return null, indicating that no valid values can be obtained.
	Csv []CsvObservation `json:"csv,omitempty" tf:"csv,omitempty"`

	// Compression format. Valid values: gzip, lzop, none (no compression).
	// Content format. Valid values: json, csv.
	Format *string `json:"format,omitempty" tf:"format,omitempty"`

	// JSON format content description.Note: this field may return null, indicating that no valid values can be obtained.
	// JSON format content description.Note: this field may return null, indicating that no valid values can be obtained.
	JSON []JSONObservation `json:"json,omitempty" tf:"json,omitempty"`
}

type ContentParameters struct {

	// CSV format content description.Note: this field may return null, indicating that no valid values can be obtained.
	// CSV format content description.Note: this field may return null, indicating that no valid values can be obtained.
	// +kubebuilder:validation:Optional
	Csv []CsvParameters `json:"csv,omitempty" tf:"csv,omitempty"`

	// Compression format. Valid values: gzip, lzop, none (no compression).
	// Content format. Valid values: json, csv.
	// +kubebuilder:validation:Optional
	Format *string `json:"format" tf:"format,omitempty"`

	// JSON format content description.Note: this field may return null, indicating that no valid values can be obtained.
	// JSON format content description.Note: this field may return null, indicating that no valid values can be obtained.
	// +kubebuilder:validation:Optional
	JSON []JSONParameters `json:"json,omitempty" tf:"json,omitempty"`
}

type CosShipperInitParameters struct {

	// Destination bucket in the shipping rule to be created.
	// Destination bucket in the shipping rule to be created.
	Bucket *string `json:"bucket,omitempty" tf:"bucket,omitempty"`

	// Compression configuration of shipped log.
	// Compression configuration of shipped log.
	Compress []CompressInitParameters `json:"compress,omitempty" tf:"compress,omitempty"`

	// Format configuration of shipped log content.
	// Format configuration of shipped log content.
	Content []ContentInitParameters `json:"content,omitempty" tf:"content,omitempty"`

	// Filter rules for shipped logs. Only logs matching the rules can be shipped. All rules are in the AND relationship, and up to five rules can be added. If the array is empty, no filtering will be performed, and all logs will be shipped.
	// Filter rules for shipped logs. Only logs matching the rules can be shipped. All rules are in the AND relationship, and up to five rules can be added. If the array is empty, no filtering will be performed, and all logs will be shipped.
	FilterRules []FilterRulesInitParameters `json:"filterRules,omitempty" tf:"filter_rules,omitempty"`

	// Shipping time interval in seconds. Default value: 300. Value range: 300~900.
	// Shipping time interval in seconds. Default value: 300. Value range: 300~900.
	Interval *float64 `json:"interval,omitempty" tf:"interval,omitempty"`

	// Maximum size of a file to be shipped, in MB. Default value: 256. Value range: 100~256.
	// Maximum size of a file to be shipped, in MB. Default value: 256. Value range: 100~256.
	MaxSize *float64 `json:"maxSize,omitempty" tf:"max_size,omitempty"`

	// Partition rule of shipped log, which can be represented in strftime time format.
	// Partition rule of shipped log, which can be represented in strftime time format.
	Partition *string `json:"partition,omitempty" tf:"partition,omitempty"`

	// Prefix of the shipping directory in the shipping rule to be created.
	// Prefix of the shipping directory in the shipping rule to be created.
	Prefix *string `json:"prefix,omitempty" tf:"prefix,omitempty"`

	// Shipping rule name.
	// Shipping rule name.
	ShipperName *string `json:"shipperName,omitempty" tf:"shipper_name,omitempty"`

	// ID of the log topic to which the shipping rule to be created belongs.
	// ID of the log topic to which the shipping rule to be created belongs.
	// +crossplane:generate:reference:type=Topic
	TopicID *string `json:"topicId,omitempty" tf:"topic_id,omitempty"`

	// Reference to a Topic to populate topicId.
	// +kubebuilder:validation:Optional
	TopicIDRef *v1.Reference `json:"topicIdRef,omitempty" tf:"-"`

	// Selector for a Topic to populate topicId.
	// +kubebuilder:validation:Optional
	TopicIDSelector *v1.Selector `json:"topicIdSelector,omitempty" tf:"-"`
}

type CosShipperObservation struct {

	// Destination bucket in the shipping rule to be created.
	// Destination bucket in the shipping rule to be created.
	Bucket *string `json:"bucket,omitempty" tf:"bucket,omitempty"`

	// Compression configuration of shipped log.
	// Compression configuration of shipped log.
	Compress []CompressObservation `json:"compress,omitempty" tf:"compress,omitempty"`

	// Format configuration of shipped log content.
	// Format configuration of shipped log content.
	Content []ContentObservation `json:"content,omitempty" tf:"content,omitempty"`

	// Filter rules for shipped logs. Only logs matching the rules can be shipped. All rules are in the AND relationship, and up to five rules can be added. If the array is empty, no filtering will be performed, and all logs will be shipped.
	// Filter rules for shipped logs. Only logs matching the rules can be shipped. All rules are in the AND relationship, and up to five rules can be added. If the array is empty, no filtering will be performed, and all logs will be shipped.
	FilterRules []FilterRulesObservation `json:"filterRules,omitempty" tf:"filter_rules,omitempty"`

	// ID of the resource.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Shipping time interval in seconds. Default value: 300. Value range: 300~900.
	// Shipping time interval in seconds. Default value: 300. Value range: 300~900.
	Interval *float64 `json:"interval,omitempty" tf:"interval,omitempty"`

	// Maximum size of a file to be shipped, in MB. Default value: 256. Value range: 100~256.
	// Maximum size of a file to be shipped, in MB. Default value: 256. Value range: 100~256.
	MaxSize *float64 `json:"maxSize,omitempty" tf:"max_size,omitempty"`

	// Partition rule of shipped log, which can be represented in strftime time format.
	// Partition rule of shipped log, which can be represented in strftime time format.
	Partition *string `json:"partition,omitempty" tf:"partition,omitempty"`

	// Prefix of the shipping directory in the shipping rule to be created.
	// Prefix of the shipping directory in the shipping rule to be created.
	Prefix *string `json:"prefix,omitempty" tf:"prefix,omitempty"`

	// Shipping rule name.
	// Shipping rule name.
	ShipperName *string `json:"shipperName,omitempty" tf:"shipper_name,omitempty"`

	// ID of the log topic to which the shipping rule to be created belongs.
	// ID of the log topic to which the shipping rule to be created belongs.
	TopicID *string `json:"topicId,omitempty" tf:"topic_id,omitempty"`
}

type CosShipperParameters struct {

	// Destination bucket in the shipping rule to be created.
	// Destination bucket in the shipping rule to be created.
	// +kubebuilder:validation:Optional
	Bucket *string `json:"bucket,omitempty" tf:"bucket,omitempty"`

	// Compression configuration of shipped log.
	// Compression configuration of shipped log.
	// +kubebuilder:validation:Optional
	Compress []CompressParameters `json:"compress,omitempty" tf:"compress,omitempty"`

	// Format configuration of shipped log content.
	// Format configuration of shipped log content.
	// +kubebuilder:validation:Optional
	Content []ContentParameters `json:"content,omitempty" tf:"content,omitempty"`

	// Filter rules for shipped logs. Only logs matching the rules can be shipped. All rules are in the AND relationship, and up to five rules can be added. If the array is empty, no filtering will be performed, and all logs will be shipped.
	// Filter rules for shipped logs. Only logs matching the rules can be shipped. All rules are in the AND relationship, and up to five rules can be added. If the array is empty, no filtering will be performed, and all logs will be shipped.
	// +kubebuilder:validation:Optional
	FilterRules []FilterRulesParameters `json:"filterRules,omitempty" tf:"filter_rules,omitempty"`

	// Shipping time interval in seconds. Default value: 300. Value range: 300~900.
	// Shipping time interval in seconds. Default value: 300. Value range: 300~900.
	// +kubebuilder:validation:Optional
	Interval *float64 `json:"interval,omitempty" tf:"interval,omitempty"`

	// Maximum size of a file to be shipped, in MB. Default value: 256. Value range: 100~256.
	// Maximum size of a file to be shipped, in MB. Default value: 256. Value range: 100~256.
	// +kubebuilder:validation:Optional
	MaxSize *float64 `json:"maxSize,omitempty" tf:"max_size,omitempty"`

	// Partition rule of shipped log, which can be represented in strftime time format.
	// Partition rule of shipped log, which can be represented in strftime time format.
	// +kubebuilder:validation:Optional
	Partition *string `json:"partition,omitempty" tf:"partition,omitempty"`

	// Prefix of the shipping directory in the shipping rule to be created.
	// Prefix of the shipping directory in the shipping rule to be created.
	// +kubebuilder:validation:Optional
	Prefix *string `json:"prefix,omitempty" tf:"prefix,omitempty"`

	// Shipping rule name.
	// Shipping rule name.
	// +kubebuilder:validation:Optional
	ShipperName *string `json:"shipperName,omitempty" tf:"shipper_name,omitempty"`

	// ID of the log topic to which the shipping rule to be created belongs.
	// ID of the log topic to which the shipping rule to be created belongs.
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

type CsvInitParameters struct {

	// Field delimiter.
	// Field delimiter.
	Delimiter *string `json:"delimiter,omitempty" tf:"delimiter,omitempty"`

	// Field delimiter.
	// Field delimiter.
	EscapeChar *string `json:"escapeChar,omitempty" tf:"escape_char,omitempty"`

	// Names of keys.Note: this field may return null, indicating that no valid values can be obtained.
	// Names of keys.Note: this field may return null, indicating that no valid values can be obtained.
	// +listType=set
	Keys []*string `json:"keys,omitempty" tf:"keys,omitempty"`

	// Content used to populate non-existing fields.
	// Content used to populate non-existing fields.
	NonExistingField *string `json:"nonExistingField,omitempty" tf:"non_existing_field,omitempty"`

	// Whether to print key on the first row of the CSV file.
	// Whether to print key on the first row of the CSV file.
	PrintKey *bool `json:"printKey,omitempty" tf:"print_key,omitempty"`
}

type CsvObservation struct {

	// Field delimiter.
	// Field delimiter.
	Delimiter *string `json:"delimiter,omitempty" tf:"delimiter,omitempty"`

	// Field delimiter.
	// Field delimiter.
	EscapeChar *string `json:"escapeChar,omitempty" tf:"escape_char,omitempty"`

	// Names of keys.Note: this field may return null, indicating that no valid values can be obtained.
	// Names of keys.Note: this field may return null, indicating that no valid values can be obtained.
	// +listType=set
	Keys []*string `json:"keys,omitempty" tf:"keys,omitempty"`

	// Content used to populate non-existing fields.
	// Content used to populate non-existing fields.
	NonExistingField *string `json:"nonExistingField,omitempty" tf:"non_existing_field,omitempty"`

	// Whether to print key on the first row of the CSV file.
	// Whether to print key on the first row of the CSV file.
	PrintKey *bool `json:"printKey,omitempty" tf:"print_key,omitempty"`
}

type CsvParameters struct {

	// Field delimiter.
	// Field delimiter.
	// +kubebuilder:validation:Optional
	Delimiter *string `json:"delimiter" tf:"delimiter,omitempty"`

	// Field delimiter.
	// Field delimiter.
	// +kubebuilder:validation:Optional
	EscapeChar *string `json:"escapeChar" tf:"escape_char,omitempty"`

	// Names of keys.Note: this field may return null, indicating that no valid values can be obtained.
	// Names of keys.Note: this field may return null, indicating that no valid values can be obtained.
	// +kubebuilder:validation:Optional
	// +listType=set
	Keys []*string `json:"keys" tf:"keys,omitempty"`

	// Content used to populate non-existing fields.
	// Content used to populate non-existing fields.
	// +kubebuilder:validation:Optional
	NonExistingField *string `json:"nonExistingField" tf:"non_existing_field,omitempty"`

	// Whether to print key on the first row of the CSV file.
	// Whether to print key on the first row of the CSV file.
	// +kubebuilder:validation:Optional
	PrintKey *bool `json:"printKey" tf:"print_key,omitempty"`
}

type FilterRulesInitParameters struct {

	// Filter rule key.
	// Filter rule key.
	Key *string `json:"key,omitempty" tf:"key,omitempty"`

	// Filter rule.
	// Filter rule.
	Regex *string `json:"regex,omitempty" tf:"regex,omitempty"`

	// Filter rule value.
	// Filter rule value.
	Value *string `json:"value,omitempty" tf:"value,omitempty"`
}

type FilterRulesObservation struct {

	// Filter rule key.
	// Filter rule key.
	Key *string `json:"key,omitempty" tf:"key,omitempty"`

	// Filter rule.
	// Filter rule.
	Regex *string `json:"regex,omitempty" tf:"regex,omitempty"`

	// Filter rule value.
	// Filter rule value.
	Value *string `json:"value,omitempty" tf:"value,omitempty"`
}

type FilterRulesParameters struct {

	// Filter rule key.
	// Filter rule key.
	// +kubebuilder:validation:Optional
	Key *string `json:"key" tf:"key,omitempty"`

	// Filter rule.
	// Filter rule.
	// +kubebuilder:validation:Optional
	Regex *string `json:"regex" tf:"regex,omitempty"`

	// Filter rule value.
	// Filter rule value.
	// +kubebuilder:validation:Optional
	Value *string `json:"value" tf:"value,omitempty"`
}

type JSONInitParameters struct {

	// Enablement flag.
	// Enablement flag.
	EnableTag *bool `json:"enableTag,omitempty" tf:"enable_tag,omitempty"`

	// Metadata information list
	// Note: this field may return null, indicating that no valid values can be obtained..
	// Metadata information list
	// Note: this field may return null, indicating that no valid values can be obtained..
	// +listType=set
	MetaFields []*string `json:"metaFields,omitempty" tf:"meta_fields,omitempty"`
}

type JSONObservation struct {

	// Enablement flag.
	// Enablement flag.
	EnableTag *bool `json:"enableTag,omitempty" tf:"enable_tag,omitempty"`

	// Metadata information list
	// Note: this field may return null, indicating that no valid values can be obtained..
	// Metadata information list
	// Note: this field may return null, indicating that no valid values can be obtained..
	// +listType=set
	MetaFields []*string `json:"metaFields,omitempty" tf:"meta_fields,omitempty"`
}

type JSONParameters struct {

	// Enablement flag.
	// Enablement flag.
	// +kubebuilder:validation:Optional
	EnableTag *bool `json:"enableTag" tf:"enable_tag,omitempty"`

	// Metadata information list
	// Note: this field may return null, indicating that no valid values can be obtained..
	// Metadata information list
	// Note: this field may return null, indicating that no valid values can be obtained..
	// +kubebuilder:validation:Optional
	// +listType=set
	MetaFields []*string `json:"metaFields" tf:"meta_fields,omitempty"`
}

// CosShipperSpec defines the desired state of CosShipper
type CosShipperSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     CosShipperParameters `json:"forProvider"`
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
	InitProvider CosShipperInitParameters `json:"initProvider,omitempty"`
}

// CosShipperStatus defines the observed state of CosShipper.
type CosShipperStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        CosShipperObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// CosShipper is the Schema for the CosShippers API. Provides a resource to create a cls cos shipper.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,tencentcloud}
type CosShipper struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.bucket) || (has(self.initProvider) && has(self.initProvider.bucket))",message="spec.forProvider.bucket is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.prefix) || (has(self.initProvider) && has(self.initProvider.prefix))",message="spec.forProvider.prefix is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.shipperName) || (has(self.initProvider) && has(self.initProvider.shipperName))",message="spec.forProvider.shipperName is a required parameter"
	Spec   CosShipperSpec   `json:"spec"`
	Status CosShipperStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// CosShipperList contains a list of CosShippers
type CosShipperList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []CosShipper `json:"items"`
}

// Repository type metadata.
var (
	CosShipper_Kind             = "CosShipper"
	CosShipper_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: CosShipper_Kind}.String()
	CosShipper_KindAPIVersion   = CosShipper_Kind + "." + CRDGroupVersion.String()
	CosShipper_GroupVersionKind = CRDGroupVersion.WithKind(CosShipper_Kind)
)

func init() {
	SchemeBuilder.Register(&CosShipper{}, &CosShipperList{})
}
