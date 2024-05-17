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

package kms

import "github.com/crossplane/upjet/pkg/config"

const shortGroupKms = "kms"

// Configure configures the kms group
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("tencentcloud_kms_key", func(r *config.Resource) {
		r.ShortGroup = shortGroupKms
		r.Kind = "Key"
	})

	p.AddResourceConfigurator("tencentcloud_kms_external_key", func(r *config.Resource) {
		r.ShortGroup = shortGroupKms
		r.Kind = "ExternalKey"
	})
}
