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

// ResolveReferences of this HttpDomain.
func (mg *HttpDomain) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.ListenerID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.ListenerIDRef,
		Selector:     mg.Spec.ForProvider.ListenerIDSelector,
		To: reference.To{
			List:    &Layer7ListenerList{},
			Managed: &Layer7Listener{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.ListenerID")
	}
	mg.Spec.ForProvider.ListenerID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.ListenerIDRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this HttpRule.
func (mg *HttpRule) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.ListenerID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.ListenerIDRef,
		Selector:     mg.Spec.ForProvider.ListenerIDSelector,
		To: reference.To{
			List:    &Layer7ListenerList{},
			Managed: &Layer7Listener{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.ListenerID")
	}
	mg.Spec.ForProvider.ListenerID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.ListenerIDRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this Layer4Listener.
func (mg *Layer4Listener) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.ProxyID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.ProxyIDRef,
		Selector:     mg.Spec.ForProvider.ProxyIDSelector,
		To: reference.To{
			List:    &ProxyList{},
			Managed: &Proxy{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.ProxyID")
	}
	mg.Spec.ForProvider.ProxyID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.ProxyIDRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this Layer7Listener.
func (mg *Layer7Listener) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.ProxyID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.ProxyIDRef,
		Selector:     mg.Spec.ForProvider.ProxyIDSelector,
		To: reference.To{
			List:    &ProxyList{},
			Managed: &Proxy{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.ProxyID")
	}
	mg.Spec.ForProvider.ProxyID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.ProxyIDRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this SecurityPolicy.
func (mg *SecurityPolicy) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.ProxyID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.ProxyIDRef,
		Selector:     mg.Spec.ForProvider.ProxyIDSelector,
		To: reference.To{
			List:    &ProxyList{},
			Managed: &Proxy{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.ProxyID")
	}
	mg.Spec.ForProvider.ProxyID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.ProxyIDRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this SecurityRule.
func (mg *SecurityRule) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.PolicyID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.PolicyIDRef,
		Selector:     mg.Spec.ForProvider.PolicyIDSelector,
		To: reference.To{
			List:    &SecurityPolicyList{},
			Managed: &SecurityPolicy{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.PolicyID")
	}
	mg.Spec.ForProvider.PolicyID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.PolicyIDRef = rsp.ResolvedReference

	return nil
}
