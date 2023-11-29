//go:build !ignore_autogenerated
// +build !ignore_autogenerated

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

// Code generated by controller-gen. DO NOT EDIT.

package v1alpha1

import (
	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EsACLObservation) DeepCopyInto(out *EsACLObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EsACLObservation.
func (in *EsACLObservation) DeepCopy() *EsACLObservation {
	if in == nil {
		return nil
	}
	out := new(EsACLObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EsACLParameters) DeepCopyInto(out *EsACLParameters) {
	*out = *in
	if in.BlackList != nil {
		in, out := &in.BlackList, &out.BlackList
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.WhiteList != nil {
		in, out := &in.WhiteList, &out.WhiteList
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EsACLParameters.
func (in *EsACLParameters) DeepCopy() *EsACLParameters {
	if in == nil {
		return nil
	}
	out := new(EsACLParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Instance) DeepCopyInto(out *Instance) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Instance.
func (in *Instance) DeepCopy() *Instance {
	if in == nil {
		return nil
	}
	out := new(Instance)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Instance) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *InstanceList) DeepCopyInto(out *InstanceList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Instance, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new InstanceList.
func (in *InstanceList) DeepCopy() *InstanceList {
	if in == nil {
		return nil
	}
	out := new(InstanceList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *InstanceList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *InstanceObservation) DeepCopyInto(out *InstanceObservation) {
	*out = *in
	if in.CreateTime != nil {
		in, out := &in.CreateTime, &out.CreateTime
		*out = new(string)
		**out = **in
	}
	if in.ElasticsearchDomain != nil {
		in, out := &in.ElasticsearchDomain, &out.ElasticsearchDomain
		*out = new(string)
		**out = **in
	}
	if in.ElasticsearchPort != nil {
		in, out := &in.ElasticsearchPort, &out.ElasticsearchPort
		*out = new(float64)
		**out = **in
	}
	if in.ElasticsearchVip != nil {
		in, out := &in.ElasticsearchVip, &out.ElasticsearchVip
		*out = new(string)
		**out = **in
	}
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(string)
		**out = **in
	}
	if in.KibanaURL != nil {
		in, out := &in.KibanaURL, &out.KibanaURL
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new InstanceObservation.
func (in *InstanceObservation) DeepCopy() *InstanceObservation {
	if in == nil {
		return nil
	}
	out := new(InstanceObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *InstanceParameters) DeepCopyInto(out *InstanceParameters) {
	*out = *in
	if in.AvailabilityZone != nil {
		in, out := &in.AvailabilityZone, &out.AvailabilityZone
		*out = new(string)
		**out = **in
	}
	if in.BasicSecurityType != nil {
		in, out := &in.BasicSecurityType, &out.BasicSecurityType
		*out = new(float64)
		**out = **in
	}
	if in.ChargePeriod != nil {
		in, out := &in.ChargePeriod, &out.ChargePeriod
		*out = new(float64)
		**out = **in
	}
	if in.ChargeType != nil {
		in, out := &in.ChargeType, &out.ChargeType
		*out = new(string)
		**out = **in
	}
	if in.DeployMode != nil {
		in, out := &in.DeployMode, &out.DeployMode
		*out = new(float64)
		**out = **in
	}
	if in.EsACL != nil {
		in, out := &in.EsACL, &out.EsACL
		*out = make([]EsACLParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.InstanceName != nil {
		in, out := &in.InstanceName, &out.InstanceName
		*out = new(string)
		**out = **in
	}
	if in.LicenseType != nil {
		in, out := &in.LicenseType, &out.LicenseType
		*out = new(string)
		**out = **in
	}
	if in.MultiZoneInfos != nil {
		in, out := &in.MultiZoneInfos, &out.MultiZoneInfos
		*out = make([]MultiZoneInfosParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.NodeInfoList != nil {
		in, out := &in.NodeInfoList, &out.NodeInfoList
		*out = make([]NodeInfoListParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	out.PasswordSecretRef = in.PasswordSecretRef
	if in.RenewFlag != nil {
		in, out := &in.RenewFlag, &out.RenewFlag
		*out = new(string)
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
	if in.Tags != nil {
		in, out := &in.Tags, &out.Tags
		*out = make(map[string]*string, len(*in))
		for key, val := range *in {
			var outVal *string
			if val == nil {
				(*out)[key] = nil
			} else {
				in, out := &val, &outVal
				*out = new(string)
				**out = **in
			}
			(*out)[key] = outVal
		}
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
	if in.Version != nil {
		in, out := &in.Version, &out.Version
		*out = new(string)
		**out = **in
	}
	if in.WebNodeTypeInfo != nil {
		in, out := &in.WebNodeTypeInfo, &out.WebNodeTypeInfo
		*out = make([]WebNodeTypeInfoParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new InstanceParameters.
func (in *InstanceParameters) DeepCopy() *InstanceParameters {
	if in == nil {
		return nil
	}
	out := new(InstanceParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *InstanceSpec) DeepCopyInto(out *InstanceSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new InstanceSpec.
func (in *InstanceSpec) DeepCopy() *InstanceSpec {
	if in == nil {
		return nil
	}
	out := new(InstanceSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *InstanceStatus) DeepCopyInto(out *InstanceStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	in.AtProvider.DeepCopyInto(&out.AtProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new InstanceStatus.
func (in *InstanceStatus) DeepCopy() *InstanceStatus {
	if in == nil {
		return nil
	}
	out := new(InstanceStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MultiZoneInfosObservation) DeepCopyInto(out *MultiZoneInfosObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MultiZoneInfosObservation.
func (in *MultiZoneInfosObservation) DeepCopy() *MultiZoneInfosObservation {
	if in == nil {
		return nil
	}
	out := new(MultiZoneInfosObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MultiZoneInfosParameters) DeepCopyInto(out *MultiZoneInfosParameters) {
	*out = *in
	if in.AvailabilityZone != nil {
		in, out := &in.AvailabilityZone, &out.AvailabilityZone
		*out = new(string)
		**out = **in
	}
	if in.SubnetID != nil {
		in, out := &in.SubnetID, &out.SubnetID
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MultiZoneInfosParameters.
func (in *MultiZoneInfosParameters) DeepCopy() *MultiZoneInfosParameters {
	if in == nil {
		return nil
	}
	out := new(MultiZoneInfosParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NodeInfoListObservation) DeepCopyInto(out *NodeInfoListObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NodeInfoListObservation.
func (in *NodeInfoListObservation) DeepCopy() *NodeInfoListObservation {
	if in == nil {
		return nil
	}
	out := new(NodeInfoListObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NodeInfoListParameters) DeepCopyInto(out *NodeInfoListParameters) {
	*out = *in
	if in.DiskSize != nil {
		in, out := &in.DiskSize, &out.DiskSize
		*out = new(float64)
		**out = **in
	}
	if in.DiskType != nil {
		in, out := &in.DiskType, &out.DiskType
		*out = new(string)
		**out = **in
	}
	if in.Encrypt != nil {
		in, out := &in.Encrypt, &out.Encrypt
		*out = new(bool)
		**out = **in
	}
	if in.NodeNum != nil {
		in, out := &in.NodeNum, &out.NodeNum
		*out = new(float64)
		**out = **in
	}
	if in.NodeType != nil {
		in, out := &in.NodeType, &out.NodeType
		*out = new(string)
		**out = **in
	}
	if in.Type != nil {
		in, out := &in.Type, &out.Type
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NodeInfoListParameters.
func (in *NodeInfoListParameters) DeepCopy() *NodeInfoListParameters {
	if in == nil {
		return nil
	}
	out := new(NodeInfoListParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WebNodeTypeInfoObservation) DeepCopyInto(out *WebNodeTypeInfoObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WebNodeTypeInfoObservation.
func (in *WebNodeTypeInfoObservation) DeepCopy() *WebNodeTypeInfoObservation {
	if in == nil {
		return nil
	}
	out := new(WebNodeTypeInfoObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WebNodeTypeInfoParameters) DeepCopyInto(out *WebNodeTypeInfoParameters) {
	*out = *in
	if in.NodeNum != nil {
		in, out := &in.NodeNum, &out.NodeNum
		*out = new(float64)
		**out = **in
	}
	if in.NodeType != nil {
		in, out := &in.NodeType, &out.NodeType
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WebNodeTypeInfoParameters.
func (in *WebNodeTypeInfoParameters) DeepCopy() *WebNodeTypeInfoParameters {
	if in == nil {
		return nil
	}
	out := new(WebNodeTypeInfoParameters)
	in.DeepCopyInto(out)
	return out
}
