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

package scf

import "github.com/crossplane/upjet/pkg/config"

const shortGroupScf = "scf"

// Configure configures the scf group
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("tencentcloud_scf_function", func(r *config.Resource) {
		r.ExternalName = config.IdentifierFromProvider
		r.ShortGroup = shortGroupScf
		r.Kind = "Function"
	})

	p.AddResourceConfigurator("tencentcloud_scf_namespace", func(r *config.Resource) {
		r.ExternalName = config.IdentifierFromProvider
		r.ShortGroup = shortGroupScf
		r.Kind = "ScfNamespace"
	})

	p.AddResourceConfigurator("tencentcloud_scf_layer", func(r *config.Resource) {
		r.ExternalName = config.IdentifierFromProvider
		r.ShortGroup = shortGroupScf
		r.Kind = "Layer"
	})
}
