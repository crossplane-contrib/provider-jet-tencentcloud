//go:build !ignore_autogenerated

/*
Copyright 2022 Upbound Inc.
*/

// Code generated by controller-gen. DO NOT EDIT.

package v1alpha1

import (
	"github.com/crossplane/crossplane-runtime/apis/common/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Cluster) DeepCopyInto(out *Cluster) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Cluster.
func (in *Cluster) DeepCopy() *Cluster {
	if in == nil {
		return nil
	}
	out := new(Cluster)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Cluster) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterInitParameters) DeepCopyInto(out *ClusterInitParameters) {
	*out = *in
	if in.ClusterName != nil {
		in, out := &in.ClusterName, &out.ClusterName
		*out = new(string)
		**out = **in
	}
	if in.IdlType != nil {
		in, out := &in.IdlType, &out.IdlType
		*out = new(string)
		**out = **in
	}
	if in.OldPasswordExpireLast != nil {
		in, out := &in.OldPasswordExpireLast, &out.OldPasswordExpireLast
		*out = new(float64)
		**out = **in
	}
	if in.SubnetID != nil {
		in, out := &in.SubnetID, &out.SubnetID
		*out = new(string)
		**out = **in
	}
	if in.SubnetIDRef != nil {
		in, out := &in.SubnetIDRef, &out.SubnetIDRef
		*out = new(v1.Reference)
		(*in).DeepCopyInto(*out)
	}
	if in.SubnetIDSelector != nil {
		in, out := &in.SubnetIDSelector, &out.SubnetIDSelector
		*out = new(v1.Selector)
		(*in).DeepCopyInto(*out)
	}
	if in.VPCID != nil {
		in, out := &in.VPCID, &out.VPCID
		*out = new(string)
		**out = **in
	}
	if in.VPCIDRef != nil {
		in, out := &in.VPCIDRef, &out.VPCIDRef
		*out = new(v1.Reference)
		(*in).DeepCopyInto(*out)
	}
	if in.VPCIDSelector != nil {
		in, out := &in.VPCIDSelector, &out.VPCIDSelector
		*out = new(v1.Selector)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterInitParameters.
func (in *ClusterInitParameters) DeepCopy() *ClusterInitParameters {
	if in == nil {
		return nil
	}
	out := new(ClusterInitParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterList) DeepCopyInto(out *ClusterList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Cluster, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterList.
func (in *ClusterList) DeepCopy() *ClusterList {
	if in == nil {
		return nil
	}
	out := new(ClusterList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ClusterList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterObservation) DeepCopyInto(out *ClusterObservation) {
	*out = *in
	if in.APIAccessID != nil {
		in, out := &in.APIAccessID, &out.APIAccessID
		*out = new(string)
		**out = **in
	}
	if in.APIAccessIP != nil {
		in, out := &in.APIAccessIP, &out.APIAccessIP
		*out = new(string)
		**out = **in
	}
	if in.APIAccessPort != nil {
		in, out := &in.APIAccessPort, &out.APIAccessPort
		*out = new(float64)
		**out = **in
	}
	if in.ClusterName != nil {
		in, out := &in.ClusterName, &out.ClusterName
		*out = new(string)
		**out = **in
	}
	if in.CreateTime != nil {
		in, out := &in.CreateTime, &out.CreateTime
		*out = new(string)
		**out = **in
	}
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(string)
		**out = **in
	}
	if in.IdlType != nil {
		in, out := &in.IdlType, &out.IdlType
		*out = new(string)
		**out = **in
	}
	if in.NetworkType != nil {
		in, out := &in.NetworkType, &out.NetworkType
		*out = new(string)
		**out = **in
	}
	if in.OldPasswordExpireLast != nil {
		in, out := &in.OldPasswordExpireLast, &out.OldPasswordExpireLast
		*out = new(float64)
		**out = **in
	}
	if in.OldPasswordExpireTime != nil {
		in, out := &in.OldPasswordExpireTime, &out.OldPasswordExpireTime
		*out = new(string)
		**out = **in
	}
	if in.PasswordStatus != nil {
		in, out := &in.PasswordStatus, &out.PasswordStatus
		*out = new(string)
		**out = **in
	}
	if in.SubnetID != nil {
		in, out := &in.SubnetID, &out.SubnetID
		*out = new(string)
		**out = **in
	}
	if in.VPCID != nil {
		in, out := &in.VPCID, &out.VPCID
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterObservation.
func (in *ClusterObservation) DeepCopy() *ClusterObservation {
	if in == nil {
		return nil
	}
	out := new(ClusterObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterParameters) DeepCopyInto(out *ClusterParameters) {
	*out = *in
	if in.ClusterName != nil {
		in, out := &in.ClusterName, &out.ClusterName
		*out = new(string)
		**out = **in
	}
	if in.IdlType != nil {
		in, out := &in.IdlType, &out.IdlType
		*out = new(string)
		**out = **in
	}
	if in.OldPasswordExpireLast != nil {
		in, out := &in.OldPasswordExpireLast, &out.OldPasswordExpireLast
		*out = new(float64)
		**out = **in
	}
	out.PasswordSecretRef = in.PasswordSecretRef
	if in.SubnetID != nil {
		in, out := &in.SubnetID, &out.SubnetID
		*out = new(string)
		**out = **in
	}
	if in.SubnetIDRef != nil {
		in, out := &in.SubnetIDRef, &out.SubnetIDRef
		*out = new(v1.Reference)
		(*in).DeepCopyInto(*out)
	}
	if in.SubnetIDSelector != nil {
		in, out := &in.SubnetIDSelector, &out.SubnetIDSelector
		*out = new(v1.Selector)
		(*in).DeepCopyInto(*out)
	}
	if in.VPCID != nil {
		in, out := &in.VPCID, &out.VPCID
		*out = new(string)
		**out = **in
	}
	if in.VPCIDRef != nil {
		in, out := &in.VPCIDRef, &out.VPCIDRef
		*out = new(v1.Reference)
		(*in).DeepCopyInto(*out)
	}
	if in.VPCIDSelector != nil {
		in, out := &in.VPCIDSelector, &out.VPCIDSelector
		*out = new(v1.Selector)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterParameters.
func (in *ClusterParameters) DeepCopy() *ClusterParameters {
	if in == nil {
		return nil
	}
	out := new(ClusterParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterSpec) DeepCopyInto(out *ClusterSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
	in.InitProvider.DeepCopyInto(&out.InitProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterSpec.
func (in *ClusterSpec) DeepCopy() *ClusterSpec {
	if in == nil {
		return nil
	}
	out := new(ClusterSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterStatus) DeepCopyInto(out *ClusterStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	in.AtProvider.DeepCopyInto(&out.AtProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterStatus.
func (in *ClusterStatus) DeepCopy() *ClusterStatus {
	if in == nil {
		return nil
	}
	out := new(ClusterStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Idl) DeepCopyInto(out *Idl) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Idl.
func (in *Idl) DeepCopy() *Idl {
	if in == nil {
		return nil
	}
	out := new(Idl)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Idl) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IdlInitParameters) DeepCopyInto(out *IdlInitParameters) {
	*out = *in
	if in.ClusterID != nil {
		in, out := &in.ClusterID, &out.ClusterID
		*out = new(string)
		**out = **in
	}
	if in.ClusterIDRef != nil {
		in, out := &in.ClusterIDRef, &out.ClusterIDRef
		*out = new(v1.Reference)
		(*in).DeepCopyInto(*out)
	}
	if in.ClusterIDSelector != nil {
		in, out := &in.ClusterIDSelector, &out.ClusterIDSelector
		*out = new(v1.Selector)
		(*in).DeepCopyInto(*out)
	}
	if in.FileContent != nil {
		in, out := &in.FileContent, &out.FileContent
		*out = new(string)
		**out = **in
	}
	if in.FileExtType != nil {
		in, out := &in.FileExtType, &out.FileExtType
		*out = new(string)
		**out = **in
	}
	if in.FileName != nil {
		in, out := &in.FileName, &out.FileName
		*out = new(string)
		**out = **in
	}
	if in.FileType != nil {
		in, out := &in.FileType, &out.FileType
		*out = new(string)
		**out = **in
	}
	if in.TablegroupID != nil {
		in, out := &in.TablegroupID, &out.TablegroupID
		*out = new(string)
		**out = **in
	}
	if in.TablegroupIDRef != nil {
		in, out := &in.TablegroupIDRef, &out.TablegroupIDRef
		*out = new(v1.Reference)
		(*in).DeepCopyInto(*out)
	}
	if in.TablegroupIDSelector != nil {
		in, out := &in.TablegroupIDSelector, &out.TablegroupIDSelector
		*out = new(v1.Selector)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IdlInitParameters.
func (in *IdlInitParameters) DeepCopy() *IdlInitParameters {
	if in == nil {
		return nil
	}
	out := new(IdlInitParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IdlList) DeepCopyInto(out *IdlList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Idl, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IdlList.
func (in *IdlList) DeepCopy() *IdlList {
	if in == nil {
		return nil
	}
	out := new(IdlList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *IdlList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IdlObservation) DeepCopyInto(out *IdlObservation) {
	*out = *in
	if in.ClusterID != nil {
		in, out := &in.ClusterID, &out.ClusterID
		*out = new(string)
		**out = **in
	}
	if in.FileContent != nil {
		in, out := &in.FileContent, &out.FileContent
		*out = new(string)
		**out = **in
	}
	if in.FileExtType != nil {
		in, out := &in.FileExtType, &out.FileExtType
		*out = new(string)
		**out = **in
	}
	if in.FileName != nil {
		in, out := &in.FileName, &out.FileName
		*out = new(string)
		**out = **in
	}
	if in.FileType != nil {
		in, out := &in.FileType, &out.FileType
		*out = new(string)
		**out = **in
	}
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(string)
		**out = **in
	}
	if in.TableInfos != nil {
		in, out := &in.TableInfos, &out.TableInfos
		*out = make([]TableInfosObservation, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.TablegroupID != nil {
		in, out := &in.TablegroupID, &out.TablegroupID
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IdlObservation.
func (in *IdlObservation) DeepCopy() *IdlObservation {
	if in == nil {
		return nil
	}
	out := new(IdlObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IdlParameters) DeepCopyInto(out *IdlParameters) {
	*out = *in
	if in.ClusterID != nil {
		in, out := &in.ClusterID, &out.ClusterID
		*out = new(string)
		**out = **in
	}
	if in.ClusterIDRef != nil {
		in, out := &in.ClusterIDRef, &out.ClusterIDRef
		*out = new(v1.Reference)
		(*in).DeepCopyInto(*out)
	}
	if in.ClusterIDSelector != nil {
		in, out := &in.ClusterIDSelector, &out.ClusterIDSelector
		*out = new(v1.Selector)
		(*in).DeepCopyInto(*out)
	}
	if in.FileContent != nil {
		in, out := &in.FileContent, &out.FileContent
		*out = new(string)
		**out = **in
	}
	if in.FileExtType != nil {
		in, out := &in.FileExtType, &out.FileExtType
		*out = new(string)
		**out = **in
	}
	if in.FileName != nil {
		in, out := &in.FileName, &out.FileName
		*out = new(string)
		**out = **in
	}
	if in.FileType != nil {
		in, out := &in.FileType, &out.FileType
		*out = new(string)
		**out = **in
	}
	if in.TablegroupID != nil {
		in, out := &in.TablegroupID, &out.TablegroupID
		*out = new(string)
		**out = **in
	}
	if in.TablegroupIDRef != nil {
		in, out := &in.TablegroupIDRef, &out.TablegroupIDRef
		*out = new(v1.Reference)
		(*in).DeepCopyInto(*out)
	}
	if in.TablegroupIDSelector != nil {
		in, out := &in.TablegroupIDSelector, &out.TablegroupIDSelector
		*out = new(v1.Selector)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IdlParameters.
func (in *IdlParameters) DeepCopy() *IdlParameters {
	if in == nil {
		return nil
	}
	out := new(IdlParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IdlSpec) DeepCopyInto(out *IdlSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
	in.InitProvider.DeepCopyInto(&out.InitProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IdlSpec.
func (in *IdlSpec) DeepCopy() *IdlSpec {
	if in == nil {
		return nil
	}
	out := new(IdlSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IdlStatus) DeepCopyInto(out *IdlStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	in.AtProvider.DeepCopyInto(&out.AtProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IdlStatus.
func (in *IdlStatus) DeepCopy() *IdlStatus {
	if in == nil {
		return nil
	}
	out := new(IdlStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Table) DeepCopyInto(out *Table) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Table.
func (in *Table) DeepCopy() *Table {
	if in == nil {
		return nil
	}
	out := new(Table)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Table) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TableGroup) DeepCopyInto(out *TableGroup) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TableGroup.
func (in *TableGroup) DeepCopy() *TableGroup {
	if in == nil {
		return nil
	}
	out := new(TableGroup)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *TableGroup) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TableGroupInitParameters) DeepCopyInto(out *TableGroupInitParameters) {
	*out = *in
	if in.ClusterID != nil {
		in, out := &in.ClusterID, &out.ClusterID
		*out = new(string)
		**out = **in
	}
	if in.ClusterIDRef != nil {
		in, out := &in.ClusterIDRef, &out.ClusterIDRef
		*out = new(v1.Reference)
		(*in).DeepCopyInto(*out)
	}
	if in.ClusterIDSelector != nil {
		in, out := &in.ClusterIDSelector, &out.ClusterIDSelector
		*out = new(v1.Selector)
		(*in).DeepCopyInto(*out)
	}
	if in.TablegroupName != nil {
		in, out := &in.TablegroupName, &out.TablegroupName
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TableGroupInitParameters.
func (in *TableGroupInitParameters) DeepCopy() *TableGroupInitParameters {
	if in == nil {
		return nil
	}
	out := new(TableGroupInitParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TableGroupList) DeepCopyInto(out *TableGroupList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]TableGroup, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TableGroupList.
func (in *TableGroupList) DeepCopy() *TableGroupList {
	if in == nil {
		return nil
	}
	out := new(TableGroupList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *TableGroupList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TableGroupObservation) DeepCopyInto(out *TableGroupObservation) {
	*out = *in
	if in.ClusterID != nil {
		in, out := &in.ClusterID, &out.ClusterID
		*out = new(string)
		**out = **in
	}
	if in.CreateTime != nil {
		in, out := &in.CreateTime, &out.CreateTime
		*out = new(string)
		**out = **in
	}
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(string)
		**out = **in
	}
	if in.TableCount != nil {
		in, out := &in.TableCount, &out.TableCount
		*out = new(float64)
		**out = **in
	}
	if in.TablegroupName != nil {
		in, out := &in.TablegroupName, &out.TablegroupName
		*out = new(string)
		**out = **in
	}
	if in.TotalSize != nil {
		in, out := &in.TotalSize, &out.TotalSize
		*out = new(float64)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TableGroupObservation.
func (in *TableGroupObservation) DeepCopy() *TableGroupObservation {
	if in == nil {
		return nil
	}
	out := new(TableGroupObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TableGroupParameters) DeepCopyInto(out *TableGroupParameters) {
	*out = *in
	if in.ClusterID != nil {
		in, out := &in.ClusterID, &out.ClusterID
		*out = new(string)
		**out = **in
	}
	if in.ClusterIDRef != nil {
		in, out := &in.ClusterIDRef, &out.ClusterIDRef
		*out = new(v1.Reference)
		(*in).DeepCopyInto(*out)
	}
	if in.ClusterIDSelector != nil {
		in, out := &in.ClusterIDSelector, &out.ClusterIDSelector
		*out = new(v1.Selector)
		(*in).DeepCopyInto(*out)
	}
	if in.TablegroupName != nil {
		in, out := &in.TablegroupName, &out.TablegroupName
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TableGroupParameters.
func (in *TableGroupParameters) DeepCopy() *TableGroupParameters {
	if in == nil {
		return nil
	}
	out := new(TableGroupParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TableGroupSpec) DeepCopyInto(out *TableGroupSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
	in.InitProvider.DeepCopyInto(&out.InitProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TableGroupSpec.
func (in *TableGroupSpec) DeepCopy() *TableGroupSpec {
	if in == nil {
		return nil
	}
	out := new(TableGroupSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TableGroupStatus) DeepCopyInto(out *TableGroupStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	in.AtProvider.DeepCopyInto(&out.AtProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TableGroupStatus.
func (in *TableGroupStatus) DeepCopy() *TableGroupStatus {
	if in == nil {
		return nil
	}
	out := new(TableGroupStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TableInfosInitParameters) DeepCopyInto(out *TableInfosInitParameters) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TableInfosInitParameters.
func (in *TableInfosInitParameters) DeepCopy() *TableInfosInitParameters {
	if in == nil {
		return nil
	}
	out := new(TableInfosInitParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TableInfosObservation) DeepCopyInto(out *TableInfosObservation) {
	*out = *in
	if in.Error != nil {
		in, out := &in.Error, &out.Error
		*out = new(string)
		**out = **in
	}
	if in.IndexKeySet != nil {
		in, out := &in.IndexKeySet, &out.IndexKeySet
		*out = new(string)
		**out = **in
	}
	if in.KeyFields != nil {
		in, out := &in.KeyFields, &out.KeyFields
		*out = new(string)
		**out = **in
	}
	if in.SumKeyFieldSize != nil {
		in, out := &in.SumKeyFieldSize, &out.SumKeyFieldSize
		*out = new(float64)
		**out = **in
	}
	if in.SumValueFieldSize != nil {
		in, out := &in.SumValueFieldSize, &out.SumValueFieldSize
		*out = new(float64)
		**out = **in
	}
	if in.TableName != nil {
		in, out := &in.TableName, &out.TableName
		*out = new(string)
		**out = **in
	}
	if in.ValueFields != nil {
		in, out := &in.ValueFields, &out.ValueFields
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TableInfosObservation.
func (in *TableInfosObservation) DeepCopy() *TableInfosObservation {
	if in == nil {
		return nil
	}
	out := new(TableInfosObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TableInfosParameters) DeepCopyInto(out *TableInfosParameters) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TableInfosParameters.
func (in *TableInfosParameters) DeepCopy() *TableInfosParameters {
	if in == nil {
		return nil
	}
	out := new(TableInfosParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TableInitParameters) DeepCopyInto(out *TableInitParameters) {
	*out = *in
	if in.ClusterID != nil {
		in, out := &in.ClusterID, &out.ClusterID
		*out = new(string)
		**out = **in
	}
	if in.ClusterIDRef != nil {
		in, out := &in.ClusterIDRef, &out.ClusterIDRef
		*out = new(v1.Reference)
		(*in).DeepCopyInto(*out)
	}
	if in.ClusterIDSelector != nil {
		in, out := &in.ClusterIDSelector, &out.ClusterIDSelector
		*out = new(v1.Selector)
		(*in).DeepCopyInto(*out)
	}
	if in.Description != nil {
		in, out := &in.Description, &out.Description
		*out = new(string)
		**out = **in
	}
	if in.IdlID != nil {
		in, out := &in.IdlID, &out.IdlID
		*out = new(string)
		**out = **in
	}
	if in.ReservedReadCu != nil {
		in, out := &in.ReservedReadCu, &out.ReservedReadCu
		*out = new(float64)
		**out = **in
	}
	if in.ReservedVolume != nil {
		in, out := &in.ReservedVolume, &out.ReservedVolume
		*out = new(float64)
		**out = **in
	}
	if in.ReservedWriteCu != nil {
		in, out := &in.ReservedWriteCu, &out.ReservedWriteCu
		*out = new(float64)
		**out = **in
	}
	if in.TableIdlType != nil {
		in, out := &in.TableIdlType, &out.TableIdlType
		*out = new(string)
		**out = **in
	}
	if in.TableName != nil {
		in, out := &in.TableName, &out.TableName
		*out = new(string)
		**out = **in
	}
	if in.TableType != nil {
		in, out := &in.TableType, &out.TableType
		*out = new(string)
		**out = **in
	}
	if in.TablegroupID != nil {
		in, out := &in.TablegroupID, &out.TablegroupID
		*out = new(string)
		**out = **in
	}
	if in.TablegroupIDRef != nil {
		in, out := &in.TablegroupIDRef, &out.TablegroupIDRef
		*out = new(v1.Reference)
		(*in).DeepCopyInto(*out)
	}
	if in.TablegroupIDSelector != nil {
		in, out := &in.TablegroupIDSelector, &out.TablegroupIDSelector
		*out = new(v1.Selector)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TableInitParameters.
func (in *TableInitParameters) DeepCopy() *TableInitParameters {
	if in == nil {
		return nil
	}
	out := new(TableInitParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TableList) DeepCopyInto(out *TableList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Table, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TableList.
func (in *TableList) DeepCopy() *TableList {
	if in == nil {
		return nil
	}
	out := new(TableList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *TableList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TableObservation) DeepCopyInto(out *TableObservation) {
	*out = *in
	if in.ClusterID != nil {
		in, out := &in.ClusterID, &out.ClusterID
		*out = new(string)
		**out = **in
	}
	if in.CreateTime != nil {
		in, out := &in.CreateTime, &out.CreateTime
		*out = new(string)
		**out = **in
	}
	if in.Description != nil {
		in, out := &in.Description, &out.Description
		*out = new(string)
		**out = **in
	}
	if in.Error != nil {
		in, out := &in.Error, &out.Error
		*out = new(string)
		**out = **in
	}
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(string)
		**out = **in
	}
	if in.IdlID != nil {
		in, out := &in.IdlID, &out.IdlID
		*out = new(string)
		**out = **in
	}
	if in.ReservedReadCu != nil {
		in, out := &in.ReservedReadCu, &out.ReservedReadCu
		*out = new(float64)
		**out = **in
	}
	if in.ReservedVolume != nil {
		in, out := &in.ReservedVolume, &out.ReservedVolume
		*out = new(float64)
		**out = **in
	}
	if in.ReservedWriteCu != nil {
		in, out := &in.ReservedWriteCu, &out.ReservedWriteCu
		*out = new(float64)
		**out = **in
	}
	if in.Status != nil {
		in, out := &in.Status, &out.Status
		*out = new(string)
		**out = **in
	}
	if in.TableIdlType != nil {
		in, out := &in.TableIdlType, &out.TableIdlType
		*out = new(string)
		**out = **in
	}
	if in.TableName != nil {
		in, out := &in.TableName, &out.TableName
		*out = new(string)
		**out = **in
	}
	if in.TableSize != nil {
		in, out := &in.TableSize, &out.TableSize
		*out = new(float64)
		**out = **in
	}
	if in.TableType != nil {
		in, out := &in.TableType, &out.TableType
		*out = new(string)
		**out = **in
	}
	if in.TablegroupID != nil {
		in, out := &in.TablegroupID, &out.TablegroupID
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TableObservation.
func (in *TableObservation) DeepCopy() *TableObservation {
	if in == nil {
		return nil
	}
	out := new(TableObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TableParameters) DeepCopyInto(out *TableParameters) {
	*out = *in
	if in.ClusterID != nil {
		in, out := &in.ClusterID, &out.ClusterID
		*out = new(string)
		**out = **in
	}
	if in.ClusterIDRef != nil {
		in, out := &in.ClusterIDRef, &out.ClusterIDRef
		*out = new(v1.Reference)
		(*in).DeepCopyInto(*out)
	}
	if in.ClusterIDSelector != nil {
		in, out := &in.ClusterIDSelector, &out.ClusterIDSelector
		*out = new(v1.Selector)
		(*in).DeepCopyInto(*out)
	}
	if in.Description != nil {
		in, out := &in.Description, &out.Description
		*out = new(string)
		**out = **in
	}
	if in.IdlID != nil {
		in, out := &in.IdlID, &out.IdlID
		*out = new(string)
		**out = **in
	}
	if in.ReservedReadCu != nil {
		in, out := &in.ReservedReadCu, &out.ReservedReadCu
		*out = new(float64)
		**out = **in
	}
	if in.ReservedVolume != nil {
		in, out := &in.ReservedVolume, &out.ReservedVolume
		*out = new(float64)
		**out = **in
	}
	if in.ReservedWriteCu != nil {
		in, out := &in.ReservedWriteCu, &out.ReservedWriteCu
		*out = new(float64)
		**out = **in
	}
	if in.TableIdlType != nil {
		in, out := &in.TableIdlType, &out.TableIdlType
		*out = new(string)
		**out = **in
	}
	if in.TableName != nil {
		in, out := &in.TableName, &out.TableName
		*out = new(string)
		**out = **in
	}
	if in.TableType != nil {
		in, out := &in.TableType, &out.TableType
		*out = new(string)
		**out = **in
	}
	if in.TablegroupID != nil {
		in, out := &in.TablegroupID, &out.TablegroupID
		*out = new(string)
		**out = **in
	}
	if in.TablegroupIDRef != nil {
		in, out := &in.TablegroupIDRef, &out.TablegroupIDRef
		*out = new(v1.Reference)
		(*in).DeepCopyInto(*out)
	}
	if in.TablegroupIDSelector != nil {
		in, out := &in.TablegroupIDSelector, &out.TablegroupIDSelector
		*out = new(v1.Selector)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TableParameters.
func (in *TableParameters) DeepCopy() *TableParameters {
	if in == nil {
		return nil
	}
	out := new(TableParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TableSpec) DeepCopyInto(out *TableSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
	in.InitProvider.DeepCopyInto(&out.InitProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TableSpec.
func (in *TableSpec) DeepCopy() *TableSpec {
	if in == nil {
		return nil
	}
	out := new(TableSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TableStatus) DeepCopyInto(out *TableStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	in.AtProvider.DeepCopyInto(&out.AtProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TableStatus.
func (in *TableStatus) DeepCopy() *TableStatus {
	if in == nil {
		return nil
	}
	out := new(TableStatus)
	in.DeepCopyInto(out)
	return out
}
