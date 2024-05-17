/*
Copyright 2022 Upbound Inc.
*/
// Code generated by angryjet. DO NOT EDIT.

package v1alpha1

import (
	"context"
	reference "github.com/crossplane/crossplane-runtime/pkg/reference"
	errors "github.com/pkg/errors"
	client "sigs.k8s.io/controller-runtime/pkg/client"
)

// ResolveReferences of this Attachment.
func (mg *Attachment) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.ScalingGroupID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.ScalingGroupIDRef,
		Selector:     mg.Spec.ForProvider.ScalingGroupIDSelector,
		To: reference.To{
			List:    &ScalingGroupList{},
			Managed: &ScalingGroup{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.ScalingGroupID")
	}
	mg.Spec.ForProvider.ScalingGroupID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.ScalingGroupIDRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this LifecycleHook.
func (mg *LifecycleHook) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.ScalingGroupID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.ScalingGroupIDRef,
		Selector:     mg.Spec.ForProvider.ScalingGroupIDSelector,
		To: reference.To{
			List:    &ScalingGroupList{},
			Managed: &ScalingGroup{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.ScalingGroupID")
	}
	mg.Spec.ForProvider.ScalingGroupID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.ScalingGroupIDRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this Notification.
func (mg *Notification) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.ScalingGroupID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.ScalingGroupIDRef,
		Selector:     mg.Spec.ForProvider.ScalingGroupIDSelector,
		To: reference.To{
			List:    &ScalingGroupList{},
			Managed: &ScalingGroup{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.ScalingGroupID")
	}
	mg.Spec.ForProvider.ScalingGroupID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.ScalingGroupIDRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this ScalingGroup.
func (mg *ScalingGroup) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.ConfigurationID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.ConfigurationIDRef,
		Selector:     mg.Spec.ForProvider.ConfigurationIDSelector,
		To: reference.To{
			List:    &ScalingConfigList{},
			Managed: &ScalingConfig{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.ConfigurationID")
	}
	mg.Spec.ForProvider.ConfigurationID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.ConfigurationIDRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this ScalingPolicy.
func (mg *ScalingPolicy) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.ScalingGroupID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.ScalingGroupIDRef,
		Selector:     mg.Spec.ForProvider.ScalingGroupIDSelector,
		To: reference.To{
			List:    &ScalingGroupList{},
			Managed: &ScalingGroup{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.ScalingGroupID")
	}
	mg.Spec.ForProvider.ScalingGroupID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.ScalingGroupIDRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this Schedule.
func (mg *Schedule) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.ScalingGroupID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.ScalingGroupIDRef,
		Selector:     mg.Spec.ForProvider.ScalingGroupIDSelector,
		To: reference.To{
			List:    &ScalingGroupList{},
			Managed: &ScalingGroup{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.ScalingGroupID")
	}
	mg.Spec.ForProvider.ScalingGroupID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.ScalingGroupIDRef = rsp.ResolvedReference

	return nil
}
