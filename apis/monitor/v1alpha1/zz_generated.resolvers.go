/*
Copyright 2022 Upbound Inc.
*/
// Code generated by angryjet. DO NOT EDIT.

package v1alpha1

import (
	"context"
	v1alpha1 "github.com/crossplane-contrib/provider-tencentcloud/apis/vpc/v1alpha1"
	reference "github.com/crossplane/crossplane-runtime/pkg/reference"
	errors "github.com/pkg/errors"
	client "sigs.k8s.io/controller-runtime/pkg/client"
)

// ResolveReferences of this AlarmPolicy.
func (mg *AlarmPolicy) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var mrsp reference.MultiResolutionResponse
	var err error

	mrsp, err = r.ResolveMultiple(ctx, reference.MultiResolutionRequest{
		CurrentValues: reference.FromPtrValues(mg.Spec.ForProvider.NoticeIds),
		Extract:       reference.ExternalName(),
		References:    mg.Spec.ForProvider.NoticeIdsRefs,
		Selector:      mg.Spec.ForProvider.NoticeIdsSelector,
		To: reference.To{
			List:    &AlarmNoticeList{},
			Managed: &AlarmNotice{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.NoticeIds")
	}
	mg.Spec.ForProvider.NoticeIds = reference.ToPtrValues(mrsp.ResolvedValues)
	mg.Spec.ForProvider.NoticeIdsRefs = mrsp.ResolvedReferences

	mrsp, err = r.ResolveMultiple(ctx, reference.MultiResolutionRequest{
		CurrentValues: reference.FromPtrValues(mg.Spec.InitProvider.NoticeIds),
		Extract:       reference.ExternalName(),
		References:    mg.Spec.InitProvider.NoticeIdsRefs,
		Selector:      mg.Spec.InitProvider.NoticeIdsSelector,
		To: reference.To{
			List:    &AlarmNoticeList{},
			Managed: &AlarmNotice{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.NoticeIds")
	}
	mg.Spec.InitProvider.NoticeIds = reference.ToPtrValues(mrsp.ResolvedValues)
	mg.Spec.InitProvider.NoticeIdsRefs = mrsp.ResolvedReferences

	return nil
}

// ResolveReferences of this GrafanaInstance.
func (mg *GrafanaInstance) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.VPCID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.VPCIDRef,
		Selector:     mg.Spec.ForProvider.VPCIDSelector,
		To: reference.To{
			List:    &v1alpha1.VPCList{},
			Managed: &v1alpha1.VPC{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.VPCID")
	}
	mg.Spec.ForProvider.VPCID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.VPCIDRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.VPCID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.InitProvider.VPCIDRef,
		Selector:     mg.Spec.InitProvider.VPCIDSelector,
		To: reference.To{
			List:    &v1alpha1.VPCList{},
			Managed: &v1alpha1.VPC{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.VPCID")
	}
	mg.Spec.InitProvider.VPCID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.VPCIDRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this GrafanaIntegration.
func (mg *GrafanaIntegration) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.InstanceID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.InstanceIDRef,
		Selector:     mg.Spec.ForProvider.InstanceIDSelector,
		To: reference.To{
			List:    &GrafanaInstanceList{},
			Managed: &GrafanaInstance{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.InstanceID")
	}
	mg.Spec.ForProvider.InstanceID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.InstanceIDRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.InstanceID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.InitProvider.InstanceIDRef,
		Selector:     mg.Spec.InitProvider.InstanceIDSelector,
		To: reference.To{
			List:    &GrafanaInstanceList{},
			Managed: &GrafanaInstance{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.InstanceID")
	}
	mg.Spec.InitProvider.InstanceID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.InstanceIDRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this GrafanaNotificationChannel.
func (mg *GrafanaNotificationChannel) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.InstanceID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.InstanceIDRef,
		Selector:     mg.Spec.ForProvider.InstanceIDSelector,
		To: reference.To{
			List:    &GrafanaInstanceList{},
			Managed: &GrafanaInstance{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.InstanceID")
	}
	mg.Spec.ForProvider.InstanceID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.InstanceIDRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.InstanceID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.InitProvider.InstanceIDRef,
		Selector:     mg.Spec.InitProvider.InstanceIDSelector,
		To: reference.To{
			List:    &GrafanaInstanceList{},
			Managed: &GrafanaInstance{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.InstanceID")
	}
	mg.Spec.InitProvider.InstanceID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.InstanceIDRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this GrafanaPlugin.
func (mg *GrafanaPlugin) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.InstanceID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.InstanceIDRef,
		Selector:     mg.Spec.ForProvider.InstanceIDSelector,
		To: reference.To{
			List:    &GrafanaInstanceList{},
			Managed: &GrafanaInstance{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.InstanceID")
	}
	mg.Spec.ForProvider.InstanceID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.InstanceIDRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.InstanceID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.InitProvider.InstanceIDRef,
		Selector:     mg.Spec.InitProvider.InstanceIDSelector,
		To: reference.To{
			List:    &GrafanaInstanceList{},
			Managed: &GrafanaInstance{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.InstanceID")
	}
	mg.Spec.InitProvider.InstanceID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.InstanceIDRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this GrafanaSSOAccount.
func (mg *GrafanaSSOAccount) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.InstanceID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.InstanceIDRef,
		Selector:     mg.Spec.ForProvider.InstanceIDSelector,
		To: reference.To{
			List:    &GrafanaInstanceList{},
			Managed: &GrafanaInstance{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.InstanceID")
	}
	mg.Spec.ForProvider.InstanceID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.InstanceIDRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.InstanceID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.InitProvider.InstanceIDRef,
		Selector:     mg.Spec.InitProvider.InstanceIDSelector,
		To: reference.To{
			List:    &GrafanaInstanceList{},
			Managed: &GrafanaInstance{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.InstanceID")
	}
	mg.Spec.InitProvider.InstanceID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.InstanceIDRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this PolicyBindingObject.
func (mg *PolicyBindingObject) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.PolicyID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.PolicyIDRef,
		Selector:     mg.Spec.ForProvider.PolicyIDSelector,
		To: reference.To{
			List:    &AlarmPolicyList{},
			Managed: &AlarmPolicy{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.PolicyID")
	}
	mg.Spec.ForProvider.PolicyID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.PolicyIDRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.PolicyID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.InitProvider.PolicyIDRef,
		Selector:     mg.Spec.InitProvider.PolicyIDSelector,
		To: reference.To{
			List:    &AlarmPolicyList{},
			Managed: &AlarmPolicy{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.PolicyID")
	}
	mg.Spec.InitProvider.PolicyID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.PolicyIDRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this TmpAlertRule.
func (mg *TmpAlertRule) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.InstanceID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.InstanceIDRef,
		Selector:     mg.Spec.ForProvider.InstanceIDSelector,
		To: reference.To{
			List:    &TmpInstanceList{},
			Managed: &TmpInstance{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.InstanceID")
	}
	mg.Spec.ForProvider.InstanceID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.InstanceIDRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.InstanceID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.InitProvider.InstanceIDRef,
		Selector:     mg.Spec.InitProvider.InstanceIDSelector,
		To: reference.To{
			List:    &TmpInstanceList{},
			Managed: &TmpInstance{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.InstanceID")
	}
	mg.Spec.InitProvider.InstanceID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.InstanceIDRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this TmpCvmAgent.
func (mg *TmpCvmAgent) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.InstanceID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.InstanceIDRef,
		Selector:     mg.Spec.ForProvider.InstanceIDSelector,
		To: reference.To{
			List:    &TmpInstanceList{},
			Managed: &TmpInstance{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.InstanceID")
	}
	mg.Spec.ForProvider.InstanceID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.InstanceIDRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.InstanceID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.InitProvider.InstanceIDRef,
		Selector:     mg.Spec.InitProvider.InstanceIDSelector,
		To: reference.To{
			List:    &TmpInstanceList{},
			Managed: &TmpInstance{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.InstanceID")
	}
	mg.Spec.InitProvider.InstanceID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.InstanceIDRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this TmpExporterIntegration.
func (mg *TmpExporterIntegration) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.InstanceID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.InstanceIDRef,
		Selector:     mg.Spec.ForProvider.InstanceIDSelector,
		To: reference.To{
			List:    &TmpInstanceList{},
			Managed: &TmpInstance{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.InstanceID")
	}
	mg.Spec.ForProvider.InstanceID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.InstanceIDRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.InstanceID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.InitProvider.InstanceIDRef,
		Selector:     mg.Spec.InitProvider.InstanceIDSelector,
		To: reference.To{
			List:    &TmpInstanceList{},
			Managed: &TmpInstance{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.InstanceID")
	}
	mg.Spec.InitProvider.InstanceID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.InstanceIDRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this TmpInstance.
func (mg *TmpInstance) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.SubnetID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.SubnetIDRef,
		Selector:     mg.Spec.ForProvider.SubnetIDSelector,
		To: reference.To{
			List:    &v1alpha1.SubnetList{},
			Managed: &v1alpha1.Subnet{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.SubnetID")
	}
	mg.Spec.ForProvider.SubnetID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.SubnetIDRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.VPCID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.VPCIDRef,
		Selector:     mg.Spec.ForProvider.VPCIDSelector,
		To: reference.To{
			List:    &v1alpha1.VPCList{},
			Managed: &v1alpha1.VPC{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.VPCID")
	}
	mg.Spec.ForProvider.VPCID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.VPCIDRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.SubnetID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.InitProvider.SubnetIDRef,
		Selector:     mg.Spec.InitProvider.SubnetIDSelector,
		To: reference.To{
			List:    &v1alpha1.SubnetList{},
			Managed: &v1alpha1.Subnet{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.SubnetID")
	}
	mg.Spec.InitProvider.SubnetID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.SubnetIDRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.VPCID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.InitProvider.VPCIDRef,
		Selector:     mg.Spec.InitProvider.VPCIDSelector,
		To: reference.To{
			List:    &v1alpha1.VPCList{},
			Managed: &v1alpha1.VPC{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.VPCID")
	}
	mg.Spec.InitProvider.VPCID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.VPCIDRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this TmpScrapeJob.
func (mg *TmpScrapeJob) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.InstanceID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.InstanceIDRef,
		Selector:     mg.Spec.ForProvider.InstanceIDSelector,
		To: reference.To{
			List:    &TmpInstanceList{},
			Managed: &TmpInstance{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.InstanceID")
	}
	mg.Spec.ForProvider.InstanceID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.InstanceIDRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.InstanceID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.InitProvider.InstanceIDRef,
		Selector:     mg.Spec.InitProvider.InstanceIDSelector,
		To: reference.To{
			List:    &TmpInstanceList{},
			Managed: &TmpInstance{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.InstanceID")
	}
	mg.Spec.InitProvider.InstanceID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.InstanceIDRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this TmpTkeAlertPolicy.
func (mg *TmpTkeAlertPolicy) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.InstanceID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.InstanceIDRef,
		Selector:     mg.Spec.ForProvider.InstanceIDSelector,
		To: reference.To{
			List:    &TmpInstanceList{},
			Managed: &TmpInstance{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.InstanceID")
	}
	mg.Spec.ForProvider.InstanceID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.InstanceIDRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.InstanceID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.InitProvider.InstanceIDRef,
		Selector:     mg.Spec.InitProvider.InstanceIDSelector,
		To: reference.To{
			List:    &TmpInstanceList{},
			Managed: &TmpInstance{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.InstanceID")
	}
	mg.Spec.InitProvider.InstanceID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.InstanceIDRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this TmpTkeClusterAgent.
func (mg *TmpTkeClusterAgent) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.InstanceID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.InstanceIDRef,
		Selector:     mg.Spec.ForProvider.InstanceIDSelector,
		To: reference.To{
			List:    &TmpInstanceList{},
			Managed: &TmpInstance{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.InstanceID")
	}
	mg.Spec.ForProvider.InstanceID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.InstanceIDRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.InstanceID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.InitProvider.InstanceIDRef,
		Selector:     mg.Spec.InitProvider.InstanceIDSelector,
		To: reference.To{
			List:    &TmpInstanceList{},
			Managed: &TmpInstance{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.InstanceID")
	}
	mg.Spec.InitProvider.InstanceID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.InstanceIDRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this TmpTkeConfig.
func (mg *TmpTkeConfig) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.InstanceID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.InstanceIDRef,
		Selector:     mg.Spec.ForProvider.InstanceIDSelector,
		To: reference.To{
			List:    &TmpInstanceList{},
			Managed: &TmpInstance{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.InstanceID")
	}
	mg.Spec.ForProvider.InstanceID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.InstanceIDRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.InstanceID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.InitProvider.InstanceIDRef,
		Selector:     mg.Spec.InitProvider.InstanceIDSelector,
		To: reference.To{
			List:    &TmpInstanceList{},
			Managed: &TmpInstance{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.InstanceID")
	}
	mg.Spec.InitProvider.InstanceID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.InstanceIDRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this TmpTkeGlobalNotification.
func (mg *TmpTkeGlobalNotification) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.InstanceID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.InstanceIDRef,
		Selector:     mg.Spec.ForProvider.InstanceIDSelector,
		To: reference.To{
			List:    &TmpInstanceList{},
			Managed: &TmpInstance{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.InstanceID")
	}
	mg.Spec.ForProvider.InstanceID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.InstanceIDRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.InstanceID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.InitProvider.InstanceIDRef,
		Selector:     mg.Spec.InitProvider.InstanceIDSelector,
		To: reference.To{
			List:    &TmpInstanceList{},
			Managed: &TmpInstance{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.InstanceID")
	}
	mg.Spec.InitProvider.InstanceID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.InstanceIDRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this TmpTkeRecordRuleYaml.
func (mg *TmpTkeRecordRuleYaml) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.InstanceID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.InstanceIDRef,
		Selector:     mg.Spec.ForProvider.InstanceIDSelector,
		To: reference.To{
			List:    &TmpInstanceList{},
			Managed: &TmpInstance{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.InstanceID")
	}
	mg.Spec.ForProvider.InstanceID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.InstanceIDRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.InstanceID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.InitProvider.InstanceIDRef,
		Selector:     mg.Spec.InitProvider.InstanceIDSelector,
		To: reference.To{
			List:    &TmpInstanceList{},
			Managed: &TmpInstance{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.InstanceID")
	}
	mg.Spec.InitProvider.InstanceID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.InstanceIDRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this TmpTkeTemplateAttachment.
func (mg *TmpTkeTemplateAttachment) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.TemplateID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.TemplateIDRef,
		Selector:     mg.Spec.ForProvider.TemplateIDSelector,
		To: reference.To{
			List:    &TmpTkeTemplateList{},
			Managed: &TmpTkeTemplate{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.TemplateID")
	}
	mg.Spec.ForProvider.TemplateID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.TemplateIDRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.TemplateID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.InitProvider.TemplateIDRef,
		Selector:     mg.Spec.InitProvider.TemplateIDSelector,
		To: reference.To{
			List:    &TmpTkeTemplateList{},
			Managed: &TmpTkeTemplate{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.TemplateID")
	}
	mg.Spec.InitProvider.TemplateID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.TemplateIDRef = rsp.ResolvedReference

	return nil
}
