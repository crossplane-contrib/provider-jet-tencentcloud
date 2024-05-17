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

type AclInitParameters struct {

	// The default is *, which means that any host can access it. Support filling in IP or network segment, and support ;separation.
	// The default is *, which means that any host can access it. Support filling in IP or network segment, and support `;`separation.
	Host *string `json:"host,omitempty" tf:"host,omitempty"`

	// ACL operation mode. Valid values: UNKNOWN, ANY, ALL, READ, WRITE, CREATE, DELETE, ALTER, DESCRIBE, CLUSTER_ACTION, DESCRIBE_CONFIGS and ALTER_CONFIGS.
	// ACL operation mode. Valid values: `UNKNOWN`, `ANY`, `ALL`, `READ`, `WRITE`, `CREATE`, `DELETE`, `ALTER`, `DESCRIBE`, `CLUSTER_ACTION`, `DESCRIBE_CONFIGS` and `ALTER_CONFIGS`.
	OperationType *string `json:"operationType,omitempty" tf:"operation_type,omitempty"`

	// ACL permission type. Valid values: UNKNOWN, ANY, DENY, ALLOW. and ALLOW by default. Currently, CKafka supports ALLOW (equivalent to allow list), and other fields will be used for future ACLs compatible with open-source Kafka.
	// ACL permission type. Valid values: `UNKNOWN`, `ANY`, `DENY`, `ALLOW`. and `ALLOW` by default. Currently, CKafka supports `ALLOW` (equivalent to allow list), and other fields will be used for future ACLs compatible with open-source Kafka.
	PermissionType *string `json:"permissionType,omitempty" tf:"permission_type,omitempty"`

	// User list. The default value is *, which means that any user can access. The current user can only be one included in the user list. For example: root meaning user root can access.
	// User list. The default value is `*`, which means that any user can access. The current user can only be one included in the user list. For example: `root` meaning user root can access.
	Principal *string `json:"principal,omitempty" tf:"principal,omitempty"`

	// ACL resource name, which is related to resource_type. For example, if resource_type is TOPIC, this field indicates the topic name; if resource_type is GROUP, this field indicates the group name.
	// ACL resource name, which is related to `resource_type`. For example, if `resource_type` is `TOPIC`, this field indicates the topic name; if `resource_type` is `GROUP`, this field indicates the group name.
	ResourceName *string `json:"resourceName,omitempty" tf:"resource_name,omitempty"`

	// ACL resource type. Valid values are UNKNOWN, ANY, TOPIC, GROUP, CLUSTER, TRANSACTIONAL_ID. and TOPIC by default. Currently, only TOPIC is available, and other fields will be used for future ACLs compatible with open-source Kafka.
	// ACL resource type. Valid values are `UNKNOWN`, `ANY`, `TOPIC`, `GROUP`, `CLUSTER`, `TRANSACTIONAL_ID`. and `TOPIC` by default. Currently, only `TOPIC` is available, and other fields will be used for future ACLs compatible with open-source Kafka.
	ResourceType *string `json:"resourceType,omitempty" tf:"resource_type,omitempty"`
}

type AclObservation struct {

	// The default is *, which means that any host can access it. Support filling in IP or network segment, and support ;separation.
	// The default is *, which means that any host can access it. Support filling in IP or network segment, and support `;`separation.
	Host *string `json:"host,omitempty" tf:"host,omitempty"`

	// ID of the resource.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// ID of the ckafka instance.
	// ID of the ckafka instance.
	InstanceID *string `json:"instanceId,omitempty" tf:"instance_id,omitempty"`

	// ACL operation mode. Valid values: UNKNOWN, ANY, ALL, READ, WRITE, CREATE, DELETE, ALTER, DESCRIBE, CLUSTER_ACTION, DESCRIBE_CONFIGS and ALTER_CONFIGS.
	// ACL operation mode. Valid values: `UNKNOWN`, `ANY`, `ALL`, `READ`, `WRITE`, `CREATE`, `DELETE`, `ALTER`, `DESCRIBE`, `CLUSTER_ACTION`, `DESCRIBE_CONFIGS` and `ALTER_CONFIGS`.
	OperationType *string `json:"operationType,omitempty" tf:"operation_type,omitempty"`

	// ACL permission type. Valid values: UNKNOWN, ANY, DENY, ALLOW. and ALLOW by default. Currently, CKafka supports ALLOW (equivalent to allow list), and other fields will be used for future ACLs compatible with open-source Kafka.
	// ACL permission type. Valid values: `UNKNOWN`, `ANY`, `DENY`, `ALLOW`. and `ALLOW` by default. Currently, CKafka supports `ALLOW` (equivalent to allow list), and other fields will be used for future ACLs compatible with open-source Kafka.
	PermissionType *string `json:"permissionType,omitempty" tf:"permission_type,omitempty"`

	// User list. The default value is *, which means that any user can access. The current user can only be one included in the user list. For example: root meaning user root can access.
	// User list. The default value is `*`, which means that any user can access. The current user can only be one included in the user list. For example: `root` meaning user root can access.
	Principal *string `json:"principal,omitempty" tf:"principal,omitempty"`

	// ACL resource name, which is related to resource_type. For example, if resource_type is TOPIC, this field indicates the topic name; if resource_type is GROUP, this field indicates the group name.
	// ACL resource name, which is related to `resource_type`. For example, if `resource_type` is `TOPIC`, this field indicates the topic name; if `resource_type` is `GROUP`, this field indicates the group name.
	ResourceName *string `json:"resourceName,omitempty" tf:"resource_name,omitempty"`

	// ACL resource type. Valid values are UNKNOWN, ANY, TOPIC, GROUP, CLUSTER, TRANSACTIONAL_ID. and TOPIC by default. Currently, only TOPIC is available, and other fields will be used for future ACLs compatible with open-source Kafka.
	// ACL resource type. Valid values are `UNKNOWN`, `ANY`, `TOPIC`, `GROUP`, `CLUSTER`, `TRANSACTIONAL_ID`. and `TOPIC` by default. Currently, only `TOPIC` is available, and other fields will be used for future ACLs compatible with open-source Kafka.
	ResourceType *string `json:"resourceType,omitempty" tf:"resource_type,omitempty"`
}

type AclParameters struct {

	// The default is *, which means that any host can access it. Support filling in IP or network segment, and support ;separation.
	// The default is *, which means that any host can access it. Support filling in IP or network segment, and support `;`separation.
	// +kubebuilder:validation:Optional
	Host *string `json:"host,omitempty" tf:"host,omitempty"`

	// ID of the ckafka instance.
	// ID of the ckafka instance.
	// +crossplane:generate:reference:type=Instance
	// +kubebuilder:validation:Optional
	InstanceID *string `json:"instanceId,omitempty" tf:"instance_id,omitempty"`

	// Reference to a Instance to populate instanceId.
	// +kubebuilder:validation:Optional
	InstanceIDRef *v1.Reference `json:"instanceIdRef,omitempty" tf:"-"`

	// Selector for a Instance to populate instanceId.
	// +kubebuilder:validation:Optional
	InstanceIDSelector *v1.Selector `json:"instanceIdSelector,omitempty" tf:"-"`

	// ACL operation mode. Valid values: UNKNOWN, ANY, ALL, READ, WRITE, CREATE, DELETE, ALTER, DESCRIBE, CLUSTER_ACTION, DESCRIBE_CONFIGS and ALTER_CONFIGS.
	// ACL operation mode. Valid values: `UNKNOWN`, `ANY`, `ALL`, `READ`, `WRITE`, `CREATE`, `DELETE`, `ALTER`, `DESCRIBE`, `CLUSTER_ACTION`, `DESCRIBE_CONFIGS` and `ALTER_CONFIGS`.
	// +kubebuilder:validation:Optional
	OperationType *string `json:"operationType,omitempty" tf:"operation_type,omitempty"`

	// ACL permission type. Valid values: UNKNOWN, ANY, DENY, ALLOW. and ALLOW by default. Currently, CKafka supports ALLOW (equivalent to allow list), and other fields will be used for future ACLs compatible with open-source Kafka.
	// ACL permission type. Valid values: `UNKNOWN`, `ANY`, `DENY`, `ALLOW`. and `ALLOW` by default. Currently, CKafka supports `ALLOW` (equivalent to allow list), and other fields will be used for future ACLs compatible with open-source Kafka.
	// +kubebuilder:validation:Optional
	PermissionType *string `json:"permissionType,omitempty" tf:"permission_type,omitempty"`

	// User list. The default value is *, which means that any user can access. The current user can only be one included in the user list. For example: root meaning user root can access.
	// User list. The default value is `*`, which means that any user can access. The current user can only be one included in the user list. For example: `root` meaning user root can access.
	// +kubebuilder:validation:Optional
	Principal *string `json:"principal,omitempty" tf:"principal,omitempty"`

	// ACL resource name, which is related to resource_type. For example, if resource_type is TOPIC, this field indicates the topic name; if resource_type is GROUP, this field indicates the group name.
	// ACL resource name, which is related to `resource_type`. For example, if `resource_type` is `TOPIC`, this field indicates the topic name; if `resource_type` is `GROUP`, this field indicates the group name.
	// +kubebuilder:validation:Optional
	ResourceName *string `json:"resourceName,omitempty" tf:"resource_name,omitempty"`

	// ACL resource type. Valid values are UNKNOWN, ANY, TOPIC, GROUP, CLUSTER, TRANSACTIONAL_ID. and TOPIC by default. Currently, only TOPIC is available, and other fields will be used for future ACLs compatible with open-source Kafka.
	// ACL resource type. Valid values are `UNKNOWN`, `ANY`, `TOPIC`, `GROUP`, `CLUSTER`, `TRANSACTIONAL_ID`. and `TOPIC` by default. Currently, only `TOPIC` is available, and other fields will be used for future ACLs compatible with open-source Kafka.
	// +kubebuilder:validation:Optional
	ResourceType *string `json:"resourceType,omitempty" tf:"resource_type,omitempty"`
}

// AclSpec defines the desired state of Acl
type AclSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     AclParameters `json:"forProvider"`
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
	InitProvider AclInitParameters `json:"initProvider,omitempty"`
}

// AclStatus defines the observed state of Acl.
type AclStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        AclObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Acl is the Schema for the Acls API. Provides a resource to create a Ckafka Acl.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,tencentcloud}
type Acl struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.operationType) || (has(self.initProvider) && has(self.initProvider.operationType))",message="spec.forProvider.operationType is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.resourceName) || (has(self.initProvider) && has(self.initProvider.resourceName))",message="spec.forProvider.resourceName is a required parameter"
	Spec   AclSpec   `json:"spec"`
	Status AclStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// AclList contains a list of Acls
type AclList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Acl `json:"items"`
}

// Repository type metadata.
var (
	Acl_Kind             = "Acl"
	Acl_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Acl_Kind}.String()
	Acl_KindAPIVersion   = Acl_Kind + "." + CRDGroupVersion.String()
	Acl_GroupVersionKind = CRDGroupVersion.WithKind(Acl_Kind)
)

func init() {
	SchemeBuilder.Register(&Acl{}, &AclList{})
}
