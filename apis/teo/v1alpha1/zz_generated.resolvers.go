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
// Code generated by angryjet. DO NOT EDIT.

package v1alpha1

import (
	"context"

	reference "github.com/crossplane/crossplane-runtime/pkg/reference"
	errors "github.com/pkg/errors"
	client "sigs.k8s.io/controller-runtime/pkg/client"
)

// ResolveReferences of this ApplicationProxy.
func (mg *ApplicationProxy) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.ZoneID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.ZoneIDRef,
		Selector:     mg.Spec.ForProvider.ZoneIDSelector,
		To: reference.To{
			List:    &ZoneList{},
			Managed: &Zone{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.ZoneID")
	}
	mg.Spec.ForProvider.ZoneID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.ZoneIDRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this ApplicationProxyRule.
func (mg *ApplicationProxyRule) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.ZoneID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.ZoneIDRef,
		Selector:     mg.Spec.ForProvider.ZoneIDSelector,
		To: reference.To{
			List:    &ZoneList{},
			Managed: &Zone{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.ZoneID")
	}
	mg.Spec.ForProvider.ZoneID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.ZoneIDRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this CustomErrorPageolicy.
func (mg *CustomErrorPageolicy) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.ZoneID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.ZoneIDRef,
		Selector:     mg.Spec.ForProvider.ZoneIDSelector,
		To: reference.To{
			List:    &ZoneList{},
			Managed: &Zone{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.ZoneID")
	}
	mg.Spec.ForProvider.ZoneID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.ZoneIDRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this DDosPolicy.
func (mg *DDosPolicy) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.ZoneID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.ZoneIDRef,
		Selector:     mg.Spec.ForProvider.ZoneIDSelector,
		To: reference.To{
			List:    &ZoneList{},
			Managed: &Zone{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.ZoneID")
	}
	mg.Spec.ForProvider.ZoneID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.ZoneIDRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this DnsRecord.
func (mg *DnsRecord) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.ZoneID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.ZoneIDRef,
		Selector:     mg.Spec.ForProvider.ZoneIDSelector,
		To: reference.To{
			List:    &ZoneList{},
			Managed: &Zone{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.ZoneID")
	}
	mg.Spec.ForProvider.ZoneID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.ZoneIDRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this DnsSec.
func (mg *DnsSec) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.ZoneID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.ZoneIDRef,
		Selector:     mg.Spec.ForProvider.ZoneIDSelector,
		To: reference.To{
			List:    &ZoneList{},
			Managed: &Zone{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.ZoneID")
	}
	mg.Spec.ForProvider.ZoneID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.ZoneIDRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this OriginGroup.
func (mg *OriginGroup) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.ZoneID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.ZoneIDRef,
		Selector:     mg.Spec.ForProvider.ZoneIDSelector,
		To: reference.To{
			List:    &ZoneList{},
			Managed: &Zone{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.ZoneID")
	}
	mg.Spec.ForProvider.ZoneID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.ZoneIDRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this OriginGrouprity.
func (mg *OriginGrouprity) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.OriginGroupID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.OriginGroupIDRef,
		Selector:     mg.Spec.ForProvider.OriginGroupIDSelector,
		To: reference.To{
			List:    &OriginGroupList{},
			Managed: &OriginGroup{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.OriginGroupID")
	}
	mg.Spec.ForProvider.OriginGroupID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.OriginGroupIDRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this RuleEngine.
func (mg *RuleEngine) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.ZoneID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.ZoneIDRef,
		Selector:     mg.Spec.ForProvider.ZoneIDSelector,
		To: reference.To{
			List:    &ZoneList{},
			Managed: &Zone{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.ZoneID")
	}
	mg.Spec.ForProvider.ZoneID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.ZoneIDRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this RuleEnginePriority.
func (mg *RuleEnginePriority) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.ZoneID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.ZoneIDRef,
		Selector:     mg.Spec.ForProvider.ZoneIDSelector,
		To: reference.To{
			List:    &ZoneList{},
			Managed: &Zone{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.ZoneID")
	}
	mg.Spec.ForProvider.ZoneID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.ZoneIDRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this SecurityPolicy.
func (mg *SecurityPolicy) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.ZoneID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.ZoneIDRef,
		Selector:     mg.Spec.ForProvider.ZoneIDSelector,
		To: reference.To{
			List:    &ZoneList{},
			Managed: &Zone{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.ZoneID")
	}
	mg.Spec.ForProvider.ZoneID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.ZoneIDRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this ZoneSetting.
func (mg *ZoneSetting) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.ZoneID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.ZoneIDRef,
		Selector:     mg.Spec.ForProvider.ZoneIDSelector,
		To: reference.To{
			List:    &ZoneList{},
			Managed: &Zone{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.ZoneID")
	}
	mg.Spec.ForProvider.ZoneID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.ZoneIDRef = rsp.ResolvedReference

	return nil
}
