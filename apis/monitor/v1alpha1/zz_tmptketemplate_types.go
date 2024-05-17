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

type RecordRulesInitParameters struct {

	// Config.
	// Config.
	Config *string `json:"config,omitempty" tf:"config,omitempty"`

	// Name.
	// Name.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Used for the argument, if the configuration comes to the template, the template id.
	// Used for the argument, if the configuration comes to the template, the template id.
	TemplateID *string `json:"templateId,omitempty" tf:"template_id,omitempty"`
}

type RecordRulesObservation struct {

	// Config.
	// Config.
	Config *string `json:"config,omitempty" tf:"config,omitempty"`

	// Name.
	// Name.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Used for the argument, if the configuration comes to the template, the template id.
	// Used for the argument, if the configuration comes to the template, the template id.
	TemplateID *string `json:"templateId,omitempty" tf:"template_id,omitempty"`
}

type RecordRulesParameters struct {

	// Config.
	// Config.
	// +kubebuilder:validation:Optional
	Config *string `json:"config" tf:"config,omitempty"`

	// Name.
	// Name.
	// +kubebuilder:validation:Optional
	Name *string `json:"name" tf:"name,omitempty"`

	// Used for the argument, if the configuration comes to the template, the template id.
	// Used for the argument, if the configuration comes to the template, the template id.
	// +kubebuilder:validation:Optional
	TemplateID *string `json:"templateId,omitempty" tf:"template_id,omitempty"`
}

type TemplateInitParameters struct {

	// Template description.
	// Template description.
	Describe *string `json:"describe,omitempty" tf:"describe,omitempty"`

	// Whether the system-supplied default template is used for outgoing references.
	// Whether the system-supplied default template is used for outgoing references.
	IsDefault *bool `json:"isDefault,omitempty" tf:"is_default,omitempty"`

	// Template dimensions, the following types are supported instance instance level, cluster cluster level.
	// Template dimensions, the following types are supported `instance` instance level, `cluster` cluster level.
	Level *string `json:"level,omitempty" tf:"level,omitempty"`

	// Name.
	// Template name.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Effective when Level is a cluster, A list of PodMonitors rules in the template.
	// Effective when Level is a cluster, A list of PodMonitors rules in the template.
	PodMonitors []TemplatePodMonitorsInitParameters `json:"podMonitors,omitempty" tf:"pod_monitors,omitempty"`

	// Effective when Level is a cluster, A list of RawJobs rules in the template.
	// Effective when Level is a cluster, A list of RawJobs rules in the template.
	RawJobs []TemplateRawJobsInitParameters `json:"rawJobs,omitempty" tf:"raw_jobs,omitempty"`

	// Effective when Level is instance, A list of aggregation rules in the template.
	// Effective when Level is instance, A list of aggregation rules in the template.
	RecordRules []RecordRulesInitParameters `json:"recordRules,omitempty" tf:"record_rules,omitempty"`

	// Effective when Level is a cluster, A list of ServiceMonitor rules in the template.
	// Effective when Level is a cluster, A list of ServiceMonitor rules in the template.
	ServiceMonitors []TemplateServiceMonitorsInitParameters `json:"serviceMonitors,omitempty" tf:"service_monitors,omitempty"`

	// Used for the argument, if the configuration comes to the template, the template id.
	// The ID of the template, which is used for the outgoing reference.
	TemplateID *string `json:"templateId,omitempty" tf:"template_id,omitempty"`

	// Last updated, for outgoing references.
	// Last updated, for outgoing references.
	UpdateTime *string `json:"updateTime,omitempty" tf:"update_time,omitempty"`

	// Whether the system-supplied default template is used for outgoing references.
	// Whether the system-supplied default template is used for outgoing references.
	Version *string `json:"version,omitempty" tf:"version,omitempty"`
}

type TemplateObservation struct {

	// Template description.
	// Template description.
	Describe *string `json:"describe,omitempty" tf:"describe,omitempty"`

	// Whether the system-supplied default template is used for outgoing references.
	// Whether the system-supplied default template is used for outgoing references.
	IsDefault *bool `json:"isDefault,omitempty" tf:"is_default,omitempty"`

	// Template dimensions, the following types are supported instance instance level, cluster cluster level.
	// Template dimensions, the following types are supported `instance` instance level, `cluster` cluster level.
	Level *string `json:"level,omitempty" tf:"level,omitempty"`

	// Name.
	// Template name.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Effective when Level is a cluster, A list of PodMonitors rules in the template.
	// Effective when Level is a cluster, A list of PodMonitors rules in the template.
	PodMonitors []TemplatePodMonitorsObservation `json:"podMonitors,omitempty" tf:"pod_monitors,omitempty"`

	// Effective when Level is a cluster, A list of RawJobs rules in the template.
	// Effective when Level is a cluster, A list of RawJobs rules in the template.
	RawJobs []TemplateRawJobsObservation `json:"rawJobs,omitempty" tf:"raw_jobs,omitempty"`

	// Effective when Level is instance, A list of aggregation rules in the template.
	// Effective when Level is instance, A list of aggregation rules in the template.
	RecordRules []RecordRulesObservation `json:"recordRules,omitempty" tf:"record_rules,omitempty"`

	// Effective when Level is a cluster, A list of ServiceMonitor rules in the template.
	// Effective when Level is a cluster, A list of ServiceMonitor rules in the template.
	ServiceMonitors []TemplateServiceMonitorsObservation `json:"serviceMonitors,omitempty" tf:"service_monitors,omitempty"`

	// Used for the argument, if the configuration comes to the template, the template id.
	// The ID of the template, which is used for the outgoing reference.
	TemplateID *string `json:"templateId,omitempty" tf:"template_id,omitempty"`

	// Last updated, for outgoing references.
	// Last updated, for outgoing references.
	UpdateTime *string `json:"updateTime,omitempty" tf:"update_time,omitempty"`

	// Whether the system-supplied default template is used for outgoing references.
	// Whether the system-supplied default template is used for outgoing references.
	Version *string `json:"version,omitempty" tf:"version,omitempty"`
}

type TemplateParameters struct {

	// Template description.
	// Template description.
	// +kubebuilder:validation:Optional
	Describe *string `json:"describe,omitempty" tf:"describe,omitempty"`

	// Whether the system-supplied default template is used for outgoing references.
	// Whether the system-supplied default template is used for outgoing references.
	// +kubebuilder:validation:Optional
	IsDefault *bool `json:"isDefault,omitempty" tf:"is_default,omitempty"`

	// Template dimensions, the following types are supported instance instance level, cluster cluster level.
	// Template dimensions, the following types are supported `instance` instance level, `cluster` cluster level.
	// +kubebuilder:validation:Optional
	Level *string `json:"level" tf:"level,omitempty"`

	// Name.
	// Template name.
	// +kubebuilder:validation:Optional
	Name *string `json:"name" tf:"name,omitempty"`

	// Effective when Level is a cluster, A list of PodMonitors rules in the template.
	// Effective when Level is a cluster, A list of PodMonitors rules in the template.
	// +kubebuilder:validation:Optional
	PodMonitors []TemplatePodMonitorsParameters `json:"podMonitors,omitempty" tf:"pod_monitors,omitempty"`

	// Effective when Level is a cluster, A list of RawJobs rules in the template.
	// Effective when Level is a cluster, A list of RawJobs rules in the template.
	// +kubebuilder:validation:Optional
	RawJobs []TemplateRawJobsParameters `json:"rawJobs,omitempty" tf:"raw_jobs,omitempty"`

	// Effective when Level is instance, A list of aggregation rules in the template.
	// Effective when Level is instance, A list of aggregation rules in the template.
	// +kubebuilder:validation:Optional
	RecordRules []RecordRulesParameters `json:"recordRules,omitempty" tf:"record_rules,omitempty"`

	// Effective when Level is a cluster, A list of ServiceMonitor rules in the template.
	// Effective when Level is a cluster, A list of ServiceMonitor rules in the template.
	// +kubebuilder:validation:Optional
	ServiceMonitors []TemplateServiceMonitorsParameters `json:"serviceMonitors,omitempty" tf:"service_monitors,omitempty"`

	// Used for the argument, if the configuration comes to the template, the template id.
	// The ID of the template, which is used for the outgoing reference.
	// +kubebuilder:validation:Optional
	TemplateID *string `json:"templateId,omitempty" tf:"template_id,omitempty"`

	// Last updated, for outgoing references.
	// Last updated, for outgoing references.
	// +kubebuilder:validation:Optional
	UpdateTime *string `json:"updateTime,omitempty" tf:"update_time,omitempty"`

	// Whether the system-supplied default template is used for outgoing references.
	// Whether the system-supplied default template is used for outgoing references.
	// +kubebuilder:validation:Optional
	Version *string `json:"version,omitempty" tf:"version,omitempty"`
}

type TemplatePodMonitorsInitParameters struct {

	// Config.
	// Config.
	Config *string `json:"config,omitempty" tf:"config,omitempty"`

	// Name.
	// Name.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Used for the argument, if the configuration comes to the template, the template id.
	// Used for the argument, if the configuration comes to the template, the template id.
	TemplateID *string `json:"templateId,omitempty" tf:"template_id,omitempty"`
}

type TemplatePodMonitorsObservation struct {

	// Config.
	// Config.
	Config *string `json:"config,omitempty" tf:"config,omitempty"`

	// Name.
	// Name.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Used for the argument, if the configuration comes to the template, the template id.
	// Used for the argument, if the configuration comes to the template, the template id.
	TemplateID *string `json:"templateId,omitempty" tf:"template_id,omitempty"`
}

type TemplatePodMonitorsParameters struct {

	// Config.
	// Config.
	// +kubebuilder:validation:Optional
	Config *string `json:"config" tf:"config,omitempty"`

	// Name.
	// Name.
	// +kubebuilder:validation:Optional
	Name *string `json:"name" tf:"name,omitempty"`

	// Used for the argument, if the configuration comes to the template, the template id.
	// Used for the argument, if the configuration comes to the template, the template id.
	// +kubebuilder:validation:Optional
	TemplateID *string `json:"templateId,omitempty" tf:"template_id,omitempty"`
}

type TemplateRawJobsInitParameters struct {

	// Config.
	// Config.
	Config *string `json:"config,omitempty" tf:"config,omitempty"`

	// Name.
	// Name.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Used for the argument, if the configuration comes to the template, the template id.
	// Used for the argument, if the configuration comes to the template, the template id.
	TemplateID *string `json:"templateId,omitempty" tf:"template_id,omitempty"`
}

type TemplateRawJobsObservation struct {

	// Config.
	// Config.
	Config *string `json:"config,omitempty" tf:"config,omitempty"`

	// Name.
	// Name.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Used for the argument, if the configuration comes to the template, the template id.
	// Used for the argument, if the configuration comes to the template, the template id.
	TemplateID *string `json:"templateId,omitempty" tf:"template_id,omitempty"`
}

type TemplateRawJobsParameters struct {

	// Config.
	// Config.
	// +kubebuilder:validation:Optional
	Config *string `json:"config" tf:"config,omitempty"`

	// Name.
	// Name.
	// +kubebuilder:validation:Optional
	Name *string `json:"name" tf:"name,omitempty"`

	// Used for the argument, if the configuration comes to the template, the template id.
	// Used for the argument, if the configuration comes to the template, the template id.
	// +kubebuilder:validation:Optional
	TemplateID *string `json:"templateId,omitempty" tf:"template_id,omitempty"`
}

type TemplateServiceMonitorsInitParameters struct {

	// Config.
	// Config.
	Config *string `json:"config,omitempty" tf:"config,omitempty"`

	// Name.
	// Name.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Used for the argument, if the configuration comes to the template, the template id.
	// Used for the argument, if the configuration comes to the template, the template id.
	TemplateID *string `json:"templateId,omitempty" tf:"template_id,omitempty"`
}

type TemplateServiceMonitorsObservation struct {

	// Config.
	// Config.
	Config *string `json:"config,omitempty" tf:"config,omitempty"`

	// Name.
	// Name.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Used for the argument, if the configuration comes to the template, the template id.
	// Used for the argument, if the configuration comes to the template, the template id.
	TemplateID *string `json:"templateId,omitempty" tf:"template_id,omitempty"`
}

type TemplateServiceMonitorsParameters struct {

	// Config.
	// Config.
	// +kubebuilder:validation:Optional
	Config *string `json:"config" tf:"config,omitempty"`

	// Name.
	// Name.
	// +kubebuilder:validation:Optional
	Name *string `json:"name" tf:"name,omitempty"`

	// Used for the argument, if the configuration comes to the template, the template id.
	// Used for the argument, if the configuration comes to the template, the template id.
	// +kubebuilder:validation:Optional
	TemplateID *string `json:"templateId,omitempty" tf:"template_id,omitempty"`
}

type TmpTkeTemplateInitParameters struct {

	// Template settings.
	// Template settings.
	Template []TemplateInitParameters `json:"template,omitempty" tf:"template,omitempty"`
}

type TmpTkeTemplateObservation struct {

	// ID of the resource.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Template settings.
	// Template settings.
	Template []TemplateObservation `json:"template,omitempty" tf:"template,omitempty"`
}

type TmpTkeTemplateParameters struct {

	// Template settings.
	// Template settings.
	// +kubebuilder:validation:Optional
	Template []TemplateParameters `json:"template,omitempty" tf:"template,omitempty"`
}

// TmpTkeTemplateSpec defines the desired state of TmpTkeTemplate
type TmpTkeTemplateSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     TmpTkeTemplateParameters `json:"forProvider"`
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
	InitProvider TmpTkeTemplateInitParameters `json:"initProvider,omitempty"`
}

// TmpTkeTemplateStatus defines the observed state of TmpTkeTemplate.
type TmpTkeTemplateStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        TmpTkeTemplateObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// TmpTkeTemplate is the Schema for the TmpTkeTemplates API. Provides a resource to create a tmp tke template
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,tencentcloud}
type TmpTkeTemplate struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.template) || (has(self.initProvider) && has(self.initProvider.template))",message="spec.forProvider.template is a required parameter"
	Spec   TmpTkeTemplateSpec   `json:"spec"`
	Status TmpTkeTemplateStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// TmpTkeTemplateList contains a list of TmpTkeTemplates
type TmpTkeTemplateList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []TmpTkeTemplate `json:"items"`
}

// Repository type metadata.
var (
	TmpTkeTemplate_Kind             = "TmpTkeTemplate"
	TmpTkeTemplate_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: TmpTkeTemplate_Kind}.String()
	TmpTkeTemplate_KindAPIVersion   = TmpTkeTemplate_Kind + "." + CRDGroupVersion.String()
	TmpTkeTemplate_GroupVersionKind = CRDGroupVersion.WithKind(TmpTkeTemplate_Kind)
)

func init() {
	SchemeBuilder.Register(&TmpTkeTemplate{}, &TmpTkeTemplateList{})
}
