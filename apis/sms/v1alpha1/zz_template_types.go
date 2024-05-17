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

type TemplateInitParameters struct {

	// Whether it is Global SMS: 0: Mainland China SMS; 1: Global SMS.
	// Whether it is Global SMS: 0: Mainland China SMS; 1: Global SMS.
	International *float64 `json:"international,omitempty" tf:"international,omitempty"`

	// Template remarks, such as reason for application and use case.
	// Template remarks, such as reason for application and use case.
	Remark *string `json:"remark,omitempty" tf:"remark,omitempty"`

	// SMS type. 0: regular SMS, 1: marketing SMS.
	// SMS type. 0: regular SMS, 1: marketing SMS.
	SMSType *float64 `json:"smsType,omitempty" tf:"sms_type,omitempty"`

	// Message Template Content.
	// Message Template Content.
	TemplateContent *string `json:"templateContent,omitempty" tf:"template_content,omitempty"`

	// Message Template name, which must be unique.
	// Message Template name, which must be unique.
	TemplateName *string `json:"templateName,omitempty" tf:"template_name,omitempty"`
}

type TemplateObservation struct {

	// ID of the resource.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Whether it is Global SMS: 0: Mainland China SMS; 1: Global SMS.
	// Whether it is Global SMS: 0: Mainland China SMS; 1: Global SMS.
	International *float64 `json:"international,omitempty" tf:"international,omitempty"`

	// Template remarks, such as reason for application and use case.
	// Template remarks, such as reason for application and use case.
	Remark *string `json:"remark,omitempty" tf:"remark,omitempty"`

	// SMS type. 0: regular SMS, 1: marketing SMS.
	// SMS type. 0: regular SMS, 1: marketing SMS.
	SMSType *float64 `json:"smsType,omitempty" tf:"sms_type,omitempty"`

	// Message Template Content.
	// Message Template Content.
	TemplateContent *string `json:"templateContent,omitempty" tf:"template_content,omitempty"`

	// Message Template name, which must be unique.
	// Message Template name, which must be unique.
	TemplateName *string `json:"templateName,omitempty" tf:"template_name,omitempty"`
}

type TemplateParameters struct {

	// Whether it is Global SMS: 0: Mainland China SMS; 1: Global SMS.
	// Whether it is Global SMS: 0: Mainland China SMS; 1: Global SMS.
	// +kubebuilder:validation:Optional
	International *float64 `json:"international,omitempty" tf:"international,omitempty"`

	// Template remarks, such as reason for application and use case.
	// Template remarks, such as reason for application and use case.
	// +kubebuilder:validation:Optional
	Remark *string `json:"remark,omitempty" tf:"remark,omitempty"`

	// SMS type. 0: regular SMS, 1: marketing SMS.
	// SMS type. 0: regular SMS, 1: marketing SMS.
	// +kubebuilder:validation:Optional
	SMSType *float64 `json:"smsType,omitempty" tf:"sms_type,omitempty"`

	// Message Template Content.
	// Message Template Content.
	// +kubebuilder:validation:Optional
	TemplateContent *string `json:"templateContent,omitempty" tf:"template_content,omitempty"`

	// Message Template name, which must be unique.
	// Message Template name, which must be unique.
	// +kubebuilder:validation:Optional
	TemplateName *string `json:"templateName,omitempty" tf:"template_name,omitempty"`
}

// TemplateSpec defines the desired state of Template
type TemplateSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     TemplateParameters `json:"forProvider"`
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
	InitProvider TemplateInitParameters `json:"initProvider,omitempty"`
}

// TemplateStatus defines the observed state of Template.
type TemplateStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        TemplateObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Template is the Schema for the Templates API. Provides a resource to create a sms template
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,tencentcloud}
type Template struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.international) || (has(self.initProvider) && has(self.initProvider.international))",message="spec.forProvider.international is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.remark) || (has(self.initProvider) && has(self.initProvider.remark))",message="spec.forProvider.remark is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.smsType) || (has(self.initProvider) && has(self.initProvider.smsType))",message="spec.forProvider.smsType is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.templateContent) || (has(self.initProvider) && has(self.initProvider.templateContent))",message="spec.forProvider.templateContent is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.templateName) || (has(self.initProvider) && has(self.initProvider.templateName))",message="spec.forProvider.templateName is a required parameter"
	Spec   TemplateSpec   `json:"spec"`
	Status TemplateStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// TemplateList contains a list of Templates
type TemplateList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Template `json:"items"`
}

// Repository type metadata.
var (
	Template_Kind             = "Template"
	Template_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Template_Kind}.String()
	Template_KindAPIVersion   = Template_Kind + "." + CRDGroupVersion.String()
	Template_GroupVersionKind = CRDGroupVersion.WithKind(Template_Kind)
)

func init() {
	SchemeBuilder.Register(&Template{}, &TemplateList{})
}
