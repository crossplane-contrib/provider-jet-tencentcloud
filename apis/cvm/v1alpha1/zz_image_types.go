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

type ImageInitParameters struct {

	// Cloud disk ID list, When creating a whole machine image based on an instance, specify the data disk ID contained in the image.
	// Cloud disk ID list, When creating a whole machine image based on an instance, specify the data disk ID contained in the image.
	// +listType=set
	DataDiskIds []*string `json:"dataDiskIds,omitempty" tf:"data_disk_ids,omitempty"`

	// Set whether to force shutdown during mirroring. The default value is false, when set to true, it means that the mirror will be made after shutdown.
	// Set whether to force shutdown during mirroring. The default value is `false`, when set to true, it means that the mirror will be made after shutdown.
	ForcePoweroff *bool `json:"forcePoweroff,omitempty" tf:"force_poweroff,omitempty"`

	// Image Description.
	// Image Description.
	ImageDescription *string `json:"imageDescription,omitempty" tf:"image_description,omitempty"`

	// Set image family. Example value: business-daily-update.
	// Set image family. Example value: `business-daily-update`.
	ImageFamily *string `json:"imageFamily,omitempty" tf:"image_family,omitempty"`

	// Image name.
	// Image name.
	ImageName *string `json:"imageName,omitempty" tf:"image_name,omitempty"`

	// Cloud server instance ID.
	// Cloud server instance ID.
	InstanceID *string `json:"instanceId,omitempty" tf:"instance_id,omitempty"`

	// Cloud disk snapshot ID list; creating a mirror based on a snapshot must include a system disk snapshot. It cannot be passed in simultaneously with InstanceId.
	// Cloud disk snapshot ID list; creating a mirror based on a snapshot must include a system disk snapshot. It cannot be passed in simultaneously with InstanceId.
	// +listType=set
	SnapshotIds []*string `json:"snapshotIds,omitempty" tf:"snapshot_ids,omitempty"`

	// Sysprep function under Windows. When creating a Windows image, you can select true or false to enable or disable the Syspre function.
	// Sysprep function under Windows. When creating a Windows image, you can select true or false to enable or disable the Syspre function.
	Sysprep *bool `json:"sysprep,omitempty" tf:"sysprep,omitempty"`

	// Tags of the image.
	// Tags of the image.
	// +mapType=granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

type ImageObservation struct {

	// Cloud disk ID list, When creating a whole machine image based on an instance, specify the data disk ID contained in the image.
	// Cloud disk ID list, When creating a whole machine image based on an instance, specify the data disk ID contained in the image.
	// +listType=set
	DataDiskIds []*string `json:"dataDiskIds,omitempty" tf:"data_disk_ids,omitempty"`

	// Set whether to force shutdown during mirroring. The default value is false, when set to true, it means that the mirror will be made after shutdown.
	// Set whether to force shutdown during mirroring. The default value is `false`, when set to true, it means that the mirror will be made after shutdown.
	ForcePoweroff *bool `json:"forcePoweroff,omitempty" tf:"force_poweroff,omitempty"`

	// ID of the resource.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Image Description.
	// Image Description.
	ImageDescription *string `json:"imageDescription,omitempty" tf:"image_description,omitempty"`

	// Set image family. Example value: business-daily-update.
	// Set image family. Example value: `business-daily-update`.
	ImageFamily *string `json:"imageFamily,omitempty" tf:"image_family,omitempty"`

	// Image name.
	// Image name.
	ImageName *string `json:"imageName,omitempty" tf:"image_name,omitempty"`

	// Cloud server instance ID.
	// Cloud server instance ID.
	InstanceID *string `json:"instanceId,omitempty" tf:"instance_id,omitempty"`

	// Cloud disk snapshot ID list; creating a mirror based on a snapshot must include a system disk snapshot. It cannot be passed in simultaneously with InstanceId.
	// Cloud disk snapshot ID list; creating a mirror based on a snapshot must include a system disk snapshot. It cannot be passed in simultaneously with InstanceId.
	// +listType=set
	SnapshotIds []*string `json:"snapshotIds,omitempty" tf:"snapshot_ids,omitempty"`

	// Sysprep function under Windows. When creating a Windows image, you can select true or false to enable or disable the Syspre function.
	// Sysprep function under Windows. When creating a Windows image, you can select true or false to enable or disable the Syspre function.
	Sysprep *bool `json:"sysprep,omitempty" tf:"sysprep,omitempty"`

	// Tags of the image.
	// Tags of the image.
	// +mapType=granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

type ImageParameters struct {

	// Cloud disk ID list, When creating a whole machine image based on an instance, specify the data disk ID contained in the image.
	// Cloud disk ID list, When creating a whole machine image based on an instance, specify the data disk ID contained in the image.
	// +kubebuilder:validation:Optional
	// +listType=set
	DataDiskIds []*string `json:"dataDiskIds,omitempty" tf:"data_disk_ids,omitempty"`

	// Set whether to force shutdown during mirroring. The default value is false, when set to true, it means that the mirror will be made after shutdown.
	// Set whether to force shutdown during mirroring. The default value is `false`, when set to true, it means that the mirror will be made after shutdown.
	// +kubebuilder:validation:Optional
	ForcePoweroff *bool `json:"forcePoweroff,omitempty" tf:"force_poweroff,omitempty"`

	// Image Description.
	// Image Description.
	// +kubebuilder:validation:Optional
	ImageDescription *string `json:"imageDescription,omitempty" tf:"image_description,omitempty"`

	// Set image family. Example value: business-daily-update.
	// Set image family. Example value: `business-daily-update`.
	// +kubebuilder:validation:Optional
	ImageFamily *string `json:"imageFamily,omitempty" tf:"image_family,omitempty"`

	// Image name.
	// Image name.
	// +kubebuilder:validation:Optional
	ImageName *string `json:"imageName,omitempty" tf:"image_name,omitempty"`

	// Cloud server instance ID.
	// Cloud server instance ID.
	// +kubebuilder:validation:Optional
	InstanceID *string `json:"instanceId,omitempty" tf:"instance_id,omitempty"`

	// Cloud disk snapshot ID list; creating a mirror based on a snapshot must include a system disk snapshot. It cannot be passed in simultaneously with InstanceId.
	// Cloud disk snapshot ID list; creating a mirror based on a snapshot must include a system disk snapshot. It cannot be passed in simultaneously with InstanceId.
	// +kubebuilder:validation:Optional
	// +listType=set
	SnapshotIds []*string `json:"snapshotIds,omitempty" tf:"snapshot_ids,omitempty"`

	// Sysprep function under Windows. When creating a Windows image, you can select true or false to enable or disable the Syspre function.
	// Sysprep function under Windows. When creating a Windows image, you can select true or false to enable or disable the Syspre function.
	// +kubebuilder:validation:Optional
	Sysprep *bool `json:"sysprep,omitempty" tf:"sysprep,omitempty"`

	// Tags of the image.
	// Tags of the image.
	// +kubebuilder:validation:Optional
	// +mapType=granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

// ImageSpec defines the desired state of Image
type ImageSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ImageParameters `json:"forProvider"`
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
	InitProvider ImageInitParameters `json:"initProvider,omitempty"`
}

// ImageStatus defines the observed state of Image.
type ImageStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ImageObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// Image is the Schema for the Images API. Provide a resource to manage image.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,tencentcloud}
type Image struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.imageName) || (has(self.initProvider) && has(self.initProvider.imageName))",message="spec.forProvider.imageName is a required parameter"
	Spec   ImageSpec   `json:"spec"`
	Status ImageStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ImageList contains a list of Images
type ImageList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Image `json:"items"`
}

// Repository type metadata.
var (
	Image_Kind             = "Image"
	Image_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Image_Kind}.String()
	Image_KindAPIVersion   = Image_Kind + "." + CRDGroupVersion.String()
	Image_GroupVersionKind = CRDGroupVersion.WithKind(Image_Kind)
)

func init() {
	SchemeBuilder.Register(&Image{}, &ImageList{})
}
