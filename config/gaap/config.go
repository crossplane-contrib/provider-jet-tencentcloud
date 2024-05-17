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

package gaap

import "github.com/crossplane/upjet/pkg/config"

const shortGroupGaap = "gaap"

// Configure configures the gaap group
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("tencentcloud_gaap_proxy", func(r *config.Resource) {
		r.ShortGroup = shortGroupGaap
		r.Kind = "Proxy"
	})

	p.AddResourceConfigurator("tencentcloud_gaap_certificate", func(r *config.Resource) {
		r.ShortGroup = shortGroupGaap
		r.Kind = "Certificate"
	})

	p.AddResourceConfigurator("tencentcloud_gaap_layer4_listener", func(r *config.Resource) {
		r.ShortGroup = shortGroupGaap
		r.Kind = "Layer4Listener"
		r.References["proxy_id"] = config.Reference{
			Type: "Proxy",
		}
	})

	p.AddResourceConfigurator("tencentcloud_gaap_layer7_listener", func(r *config.Resource) {
		r.ShortGroup = shortGroupGaap
		r.Kind = "Layer7Listener"
		r.References["proxy_id"] = config.Reference{
			Type: "Proxy",
		}
	})

	p.AddResourceConfigurator("tencentcloud_gaap_realserver", func(r *config.Resource) {
		r.ShortGroup = shortGroupGaap
		r.Kind = "Realserver"
	})

	p.AddResourceConfigurator("tencentcloud_gaap_http_domain", func(r *config.Resource) {
		r.ShortGroup = shortGroupGaap
		r.Kind = "HttpDomain"
		r.References["listener_id"] = config.Reference{
			Type: "Layer7Listener",
		}
	})

	p.AddResourceConfigurator("tencentcloud_gaap_http_rule", func(r *config.Resource) {
		r.ShortGroup = shortGroupGaap
		r.Kind = "HttpRule"
		r.References["listener_id"] = config.Reference{
			Type: "Layer7Listener",
		}
	})

	p.AddResourceConfigurator("tencentcloud_gaap_security_policy", func(r *config.Resource) {
		r.ShortGroup = shortGroupGaap
		r.Kind = "SecurityPolicy"
		r.References["proxy_id"] = config.Reference{
			Type: "Proxy",
		}
	})

	p.AddResourceConfigurator("tencentcloud_gaap_security_rule", func(r *config.Resource) {
		r.ShortGroup = shortGroupGaap
		r.Kind = "SecurityRule"
		r.References["policy_id"] = config.Reference{
			Type: "SecurityPolicy",
		}
	})

}
