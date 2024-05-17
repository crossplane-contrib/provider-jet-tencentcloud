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

type AgentsInitParameters struct {

	// An id identify the cluster, like cls-xxxxxx.
	// An id identify the cluster, like `cls-xxxxxx`.
	ClusterID *string `json:"clusterId,omitempty" tf:"cluster_id,omitempty"`

	// Type of cluster.
	// Type of cluster.
	ClusterType *string `json:"clusterType,omitempty" tf:"cluster_type,omitempty"`

	// Whether to enable the public network CLB.
	// Whether to enable the public network CLB.
	EnableExternal *bool `json:"enableExternal,omitempty" tf:"enable_external,omitempty"`

	// All metrics collected by the cluster will carry these labels.
	// All metrics collected by the cluster will carry these labels.
	ExternalLabels []ExternalLabelsInitParameters `json:"externalLabels,omitempty" tf:"external_labels,omitempty"`

	// Pod configuration for components deployed in the cluster.
	// Pod configuration for components deployed in the cluster.
	InClusterPodConfig []InClusterPodConfigInitParameters `json:"inClusterPodConfig,omitempty" tf:"in_cluster_pod_config,omitempty"`

	// Whether to install the default collection configuration.
	// Whether to install the default collection configuration.
	NotInstallBasicScrape *bool `json:"notInstallBasicScrape,omitempty" tf:"not_install_basic_scrape,omitempty"`

	// Whether to collect indicators, true means drop all indicators, false means collect default indicators.
	// Whether to collect indicators, true means drop all indicators, false means collect default indicators.
	NotScrape *bool `json:"notScrape,omitempty" tf:"not_scrape,omitempty"`

	// Limitation of region.
	// Limitation of region.
	Region *string `json:"region,omitempty" tf:"region,omitempty"`
}

type AgentsObservation struct {

	// An id identify the cluster, like cls-xxxxxx.
	// An id identify the cluster, like `cls-xxxxxx`.
	ClusterID *string `json:"clusterId,omitempty" tf:"cluster_id,omitempty"`

	// Indicator name.
	// the name of the cluster.
	ClusterName *string `json:"clusterName,omitempty" tf:"cluster_name,omitempty"`

	// Type of cluster.
	// Type of cluster.
	ClusterType *string `json:"clusterType,omitempty" tf:"cluster_type,omitempty"`

	// Whether to enable the public network CLB.
	// Whether to enable the public network CLB.
	EnableExternal *bool `json:"enableExternal,omitempty" tf:"enable_external,omitempty"`

	// All metrics collected by the cluster will carry these labels.
	// All metrics collected by the cluster will carry these labels.
	ExternalLabels []ExternalLabelsObservation `json:"externalLabels,omitempty" tf:"external_labels,omitempty"`

	// Pod configuration for components deployed in the cluster.
	// Pod configuration for components deployed in the cluster.
	InClusterPodConfig []InClusterPodConfigObservation `json:"inClusterPodConfig,omitempty" tf:"in_cluster_pod_config,omitempty"`

	// Whether to install the default collection configuration.
	// Whether to install the default collection configuration.
	NotInstallBasicScrape *bool `json:"notInstallBasicScrape,omitempty" tf:"not_install_basic_scrape,omitempty"`

	// Whether to collect indicators, true means drop all indicators, false means collect default indicators.
	// Whether to collect indicators, true means drop all indicators, false means collect default indicators.
	NotScrape *bool `json:"notScrape,omitempty" tf:"not_scrape,omitempty"`

	// Limitation of region.
	// Limitation of region.
	Region *string `json:"region,omitempty" tf:"region,omitempty"`

	// agent state, `normal`, `abnormal`.
	Status *string `json:"status,omitempty" tf:"status,omitempty"`
}

type AgentsParameters struct {

	// An id identify the cluster, like cls-xxxxxx.
	// An id identify the cluster, like `cls-xxxxxx`.
	// +kubebuilder:validation:Optional
	ClusterID *string `json:"clusterId" tf:"cluster_id,omitempty"`

	// Type of cluster.
	// Type of cluster.
	// +kubebuilder:validation:Optional
	ClusterType *string `json:"clusterType" tf:"cluster_type,omitempty"`

	// Whether to enable the public network CLB.
	// Whether to enable the public network CLB.
	// +kubebuilder:validation:Optional
	EnableExternal *bool `json:"enableExternal" tf:"enable_external,omitempty"`

	// All metrics collected by the cluster will carry these labels.
	// All metrics collected by the cluster will carry these labels.
	// +kubebuilder:validation:Optional
	ExternalLabels []ExternalLabelsParameters `json:"externalLabels,omitempty" tf:"external_labels,omitempty"`

	// Pod configuration for components deployed in the cluster.
	// Pod configuration for components deployed in the cluster.
	// +kubebuilder:validation:Optional
	InClusterPodConfig []InClusterPodConfigParameters `json:"inClusterPodConfig,omitempty" tf:"in_cluster_pod_config,omitempty"`

	// Whether to install the default collection configuration.
	// Whether to install the default collection configuration.
	// +kubebuilder:validation:Optional
	NotInstallBasicScrape *bool `json:"notInstallBasicScrape,omitempty" tf:"not_install_basic_scrape,omitempty"`

	// Whether to collect indicators, true means drop all indicators, false means collect default indicators.
	// Whether to collect indicators, true means drop all indicators, false means collect default indicators.
	// +kubebuilder:validation:Optional
	NotScrape *bool `json:"notScrape,omitempty" tf:"not_scrape,omitempty"`

	// Limitation of region.
	// Limitation of region.
	// +kubebuilder:validation:Optional
	Region *string `json:"region" tf:"region,omitempty"`
}

type ExternalLabelsInitParameters struct {

	// Indicator name.
	// Indicator name.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Index value.
	// Index value.
	Value *string `json:"value,omitempty" tf:"value,omitempty"`
}

type ExternalLabelsObservation struct {

	// Indicator name.
	// Indicator name.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Index value.
	// Index value.
	Value *string `json:"value,omitempty" tf:"value,omitempty"`
}

type ExternalLabelsParameters struct {

	// Indicator name.
	// Indicator name.
	// +kubebuilder:validation:Optional
	Name *string `json:"name" tf:"name,omitempty"`

	// Index value.
	// Index value.
	// +kubebuilder:validation:Optional
	Value *string `json:"value,omitempty" tf:"value,omitempty"`
}

type InClusterPodConfigInitParameters struct {

	// Whether to use HostNetWork.
	// Whether to use HostNetWork.
	HostNet *bool `json:"hostNet,omitempty" tf:"host_net,omitempty"`

	// Specify the pod to run the node.
	// Specify the pod to run the node.
	NodeSelector []NodeSelectorInitParameters `json:"nodeSelector,omitempty" tf:"node_selector,omitempty"`

	// Tolerate Stain.
	// Tolerate Stain.
	Tolerations []TolerationsInitParameters `json:"tolerations,omitempty" tf:"tolerations,omitempty"`
}

type InClusterPodConfigObservation struct {

	// Whether to use HostNetWork.
	// Whether to use HostNetWork.
	HostNet *bool `json:"hostNet,omitempty" tf:"host_net,omitempty"`

	// Specify the pod to run the node.
	// Specify the pod to run the node.
	NodeSelector []NodeSelectorObservation `json:"nodeSelector,omitempty" tf:"node_selector,omitempty"`

	// Tolerate Stain.
	// Tolerate Stain.
	Tolerations []TolerationsObservation `json:"tolerations,omitempty" tf:"tolerations,omitempty"`
}

type InClusterPodConfigParameters struct {

	// Whether to use HostNetWork.
	// Whether to use HostNetWork.
	// +kubebuilder:validation:Optional
	HostNet *bool `json:"hostNet" tf:"host_net,omitempty"`

	// Specify the pod to run the node.
	// Specify the pod to run the node.
	// +kubebuilder:validation:Optional
	NodeSelector []NodeSelectorParameters `json:"nodeSelector,omitempty" tf:"node_selector,omitempty"`

	// Tolerate Stain.
	// Tolerate Stain.
	// +kubebuilder:validation:Optional
	Tolerations []TolerationsParameters `json:"tolerations,omitempty" tf:"tolerations,omitempty"`
}

type NodeSelectorInitParameters struct {

	// Indicator name.
	// The pod configuration name of the component deployed in the cluster.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Index value.
	// Pod configuration values for components deployed in the cluster.
	Value *string `json:"value,omitempty" tf:"value,omitempty"`
}

type NodeSelectorObservation struct {

	// Indicator name.
	// The pod configuration name of the component deployed in the cluster.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Index value.
	// Pod configuration values for components deployed in the cluster.
	Value *string `json:"value,omitempty" tf:"value,omitempty"`
}

type NodeSelectorParameters struct {

	// Indicator name.
	// The pod configuration name of the component deployed in the cluster.
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Index value.
	// Pod configuration values for components deployed in the cluster.
	// +kubebuilder:validation:Optional
	Value *string `json:"value,omitempty" tf:"value,omitempty"`
}

type TmpTkeClusterAgentInitParameters struct {

	// agent list.
	// agent list.
	Agents []AgentsInitParameters `json:"agents,omitempty" tf:"agents,omitempty"`
}

type TmpTkeClusterAgentObservation struct {

	// agent list.
	// agent list.
	Agents []AgentsObservation `json:"agents,omitempty" tf:"agents,omitempty"`

	// ID of the resource.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Instance Id.
	// Instance Id.
	InstanceID *string `json:"instanceId,omitempty" tf:"instance_id,omitempty"`
}

type TmpTkeClusterAgentParameters struct {

	// agent list.
	// agent list.
	// +kubebuilder:validation:Optional
	Agents []AgentsParameters `json:"agents,omitempty" tf:"agents,omitempty"`

	// Instance Id.
	// Instance Id.
	// +crossplane:generate:reference:type=TmpInstance
	// +kubebuilder:validation:Optional
	InstanceID *string `json:"instanceId,omitempty" tf:"instance_id,omitempty"`

	// Reference to a TmpInstance to populate instanceId.
	// +kubebuilder:validation:Optional
	InstanceIDRef *v1.Reference `json:"instanceIdRef,omitempty" tf:"-"`

	// Selector for a TmpInstance to populate instanceId.
	// +kubebuilder:validation:Optional
	InstanceIDSelector *v1.Selector `json:"instanceIdSelector,omitempty" tf:"-"`
}

type TolerationsInitParameters struct {

	// blemish effect to match.
	// blemish effect to match.
	Effect *string `json:"effect,omitempty" tf:"effect,omitempty"`

	// The taint key to which the tolerance applies.
	// The taint key to which the tolerance applies.
	Key *string `json:"key,omitempty" tf:"key,omitempty"`

	// key-value relationship.
	// key-value relationship.
	Operator *string `json:"operator,omitempty" tf:"operator,omitempty"`
}

type TolerationsObservation struct {

	// blemish effect to match.
	// blemish effect to match.
	Effect *string `json:"effect,omitempty" tf:"effect,omitempty"`

	// The taint key to which the tolerance applies.
	// The taint key to which the tolerance applies.
	Key *string `json:"key,omitempty" tf:"key,omitempty"`

	// key-value relationship.
	// key-value relationship.
	Operator *string `json:"operator,omitempty" tf:"operator,omitempty"`
}

type TolerationsParameters struct {

	// blemish effect to match.
	// blemish effect to match.
	// +kubebuilder:validation:Optional
	Effect *string `json:"effect,omitempty" tf:"effect,omitempty"`

	// The taint key to which the tolerance applies.
	// The taint key to which the tolerance applies.
	// +kubebuilder:validation:Optional
	Key *string `json:"key,omitempty" tf:"key,omitempty"`

	// key-value relationship.
	// key-value relationship.
	// +kubebuilder:validation:Optional
	Operator *string `json:"operator,omitempty" tf:"operator,omitempty"`
}

// TmpTkeClusterAgentSpec defines the desired state of TmpTkeClusterAgent
type TmpTkeClusterAgentSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     TmpTkeClusterAgentParameters `json:"forProvider"`
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
	InitProvider TmpTkeClusterAgentInitParameters `json:"initProvider,omitempty"`
}

// TmpTkeClusterAgentStatus defines the observed state of TmpTkeClusterAgent.
type TmpTkeClusterAgentStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        TmpTkeClusterAgentObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// TmpTkeClusterAgent is the Schema for the TmpTkeClusterAgents API. Provides a resource to create a tmp tke cluster agent
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,tencentcloud}
type TmpTkeClusterAgent struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.agents) || (has(self.initProvider) && has(self.initProvider.agents))",message="spec.forProvider.agents is a required parameter"
	Spec   TmpTkeClusterAgentSpec   `json:"spec"`
	Status TmpTkeClusterAgentStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// TmpTkeClusterAgentList contains a list of TmpTkeClusterAgents
type TmpTkeClusterAgentList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []TmpTkeClusterAgent `json:"items"`
}

// Repository type metadata.
var (
	TmpTkeClusterAgent_Kind             = "TmpTkeClusterAgent"
	TmpTkeClusterAgent_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: TmpTkeClusterAgent_Kind}.String()
	TmpTkeClusterAgent_KindAPIVersion   = TmpTkeClusterAgent_Kind + "." + CRDGroupVersion.String()
	TmpTkeClusterAgent_GroupVersionKind = CRDGroupVersion.WithKind(TmpTkeClusterAgent_Kind)
)

func init() {
	SchemeBuilder.Register(&TmpTkeClusterAgent{}, &TmpTkeClusterAgentList{})
}
