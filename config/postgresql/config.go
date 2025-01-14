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

package postgresql

import "github.com/crossplane/upjet/pkg/config"

const shortGroupPostgresql = "postgresql"

// Configure configures the postgresql group
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("tencentcloud_postgresql_instance", func(r *config.Resource) {
		r.ShortGroup = shortGroupPostgresql
		r.Kind = "Instance"
		r.References["vpc_id"] = config.Reference{
			Type: "github.com/crossplane-contrib/provider-tencentcloud/apis/vpc/v1alpha1.VPC",
		}
		r.References["subnet_id"] = config.Reference{
			Type: "github.com/crossplane-contrib/provider-tencentcloud/apis/vpc/v1alpha1.Subnet",
		}
		if t, ok := r.TerraformResource.Schema["db_major_vesion"]; ok {
			t.Optional = false
			t.Computed = true
		}
	})

	p.AddResourceConfigurator("tencentcloud_postgresql_readonly_instance", func(r *config.Resource) {
		r.ShortGroup = shortGroupPostgresql
		r.Kind = "ReadonlyInstance"
		r.References["vpc_id"] = config.Reference{
			Type: "github.com/crossplane-contrib/provider-tencentcloud/apis/vpc/v1alpha1.VPC",
		}
		r.References["subnet_id"] = config.Reference{
			Type: "github.com/crossplane-contrib/provider-tencentcloud/apis/vpc/v1alpha1.Subnet",
		}
		r.References["master_instance_id"] = config.Reference{
			Type: "Instance",
		}
		r.References["security_groups_ids"] = config.Reference{
			Type: "github.com/crossplane-contrib/provider-tencentcloud/apis/vpc/v1alpha1.SecurityGroup",
		}
	})

	p.AddResourceConfigurator("tencentcloud_postgresql_readonly_group", func(r *config.Resource) {
		r.ShortGroup = shortGroupPostgresql
		r.Kind = "ReadonlyGroup"
		r.References["vpc_id"] = config.Reference{
			Type: "github.com/crossplane-contrib/provider-tencentcloud/apis/vpc/v1alpha1.VPC",
		}
		r.References["subnet_id"] = config.Reference{
			Type: "github.com/crossplane-contrib/provider-tencentcloud/apis/vpc/v1alpha1.Subnet",
		}
	})

	p.AddResourceConfigurator("tencentcloud_postgresql_readonly_attachment", func(r *config.Resource) {
		r.ShortGroup = shortGroupPostgresql
		r.Kind = "ReadonlyAttachment"
		r.References["db_instance_id"] = config.Reference{
			Type: "ReadonlyInstance",
		}
		r.References["read_only_group_id"] = config.Reference{
			Type: "ReadonlyGroup",
		}
	})
}
