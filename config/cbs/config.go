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

package cbs

import "github.com/crossplane/upjet/pkg/config"

const shortGroupCbs = "cbs"

// Configure configures the cbs group
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("tencentcloud_cbs_storage", func(r *config.Resource) {
		r.ShortGroup = shortGroupCbs
		r.Kind = "Storage"
	})

	p.AddResourceConfigurator("tencentcloud_cbs_storage_set", func(r *config.Resource) {
		r.ShortGroup = shortGroupCbs
		r.Kind = "StorageSet"
	})

	p.AddResourceConfigurator("tencentcloud_cbs_storage_attachment", func(r *config.Resource) {
		r.ShortGroup = shortGroupCbs
		r.Kind = "StorageAttachment"
		r.References["storage_id"] = config.Reference{
			Type: "Storage",
		}
		r.References["instance_id"] = config.Reference{
			Type: "github.com/crossplane-contrib/provider-tencentcloud/apis/cvm/v1alpha1.Instance",
		}
	})

	p.AddResourceConfigurator("tencentcloud_cbs_snapshot", func(r *config.Resource) {
		r.ShortGroup = shortGroupCbs
		r.Kind = "Snapshot"
		r.References["storage_id"] = config.Reference{
			Type: "Storage",
		}
	})

	p.AddResourceConfigurator("tencentcloud_cbs_snapshot_policy", func(r *config.Resource) {
		r.ShortGroup = shortGroupCbs
		r.Kind = "SnapshotPolicy"
	})

	p.AddResourceConfigurator("tencentcloud_cbs_snapshot_policy_attachment", func(r *config.Resource) {
		r.ShortGroup = shortGroupCbs
		r.Kind = "SnapshotPolicyAttachment"
		r.References["storage_id"] = config.Reference{
			Type: "Storage",
		}
		r.References["snapshot_policy_id"] = config.Reference{
			Type: "SnapshotPolicy",
		}
	})

}
