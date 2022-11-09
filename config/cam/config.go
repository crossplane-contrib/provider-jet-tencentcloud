/*
Copyright 2021 The Crossplane Authors.

Licensed under the Apache License, shortGroupCamersion 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language goshortGroupCamerning permissions and
limitations under the License.
*/

package cam

import (
	tjconfig "github.com/crossplane/terrajet/pkg/config"
)

const shortGroupCam = "cam"

// Configure configures the cam group
func Configure(p *tjconfig.Provider) {
	p.AddResourceConfigurator("tencentcloud_cam_role", func(r *tjconfig.Resource) {
		r.ExternalName = tjconfig.IdentifierFromProvider
		r.ShortGroup = shortGroupCam
		r.Kind = "Role"
	})

	p.AddResourceConfigurator("tencentcloud_cam_role_sso", func(r *tjconfig.Resource) {
		r.ExternalName = tjconfig.IdentifierFromProvider
		r.ShortGroup = shortGroupCam
		r.Kind = "RoleSSO"
	})

	p.AddResourceConfigurator("tencentcloud_cam_policy", func(r *tjconfig.Resource) {
		r.ExternalName = tjconfig.IdentifierFromProvider
		r.ShortGroup = shortGroupCam
		r.Kind = "Policy"
	})

	p.AddResourceConfigurator("tencentcloud_cam_role_policy_attachment", func(r *tjconfig.Resource) {
		r.ExternalName = tjconfig.IdentifierFromProvider
		r.ShortGroup = shortGroupCam
		r.Kind = "RolePolicyAttachment"
		r.References["role_id"] = tjconfig.Reference{
			Type: "Role",
		}
		r.References["policy_id"] = tjconfig.Reference{
			Type: "Policy",
		}
	})

	p.AddResourceConfigurator("tencentcloud_cam_oidc_sso", func(r *tjconfig.Resource) {
		r.ExternalName = tjconfig.IdentifierFromProvider
		r.ShortGroup = shortGroupCam
		r.Kind = "OidcSSO"
	})

	p.AddResourceConfigurator("tencentcloud_cam_group", func(r *tjconfig.Resource) {
		r.ExternalName = tjconfig.IdentifierFromProvider
		r.ShortGroup = shortGroupCam
		r.Kind = "Group"
	})

	p.AddResourceConfigurator("tencentcloud_cam_user", func(r *tjconfig.Resource) {
		r.ExternalName = tjconfig.IdentifierFromProvider
		r.ShortGroup = shortGroupCam
		r.Kind = "User"
	})

	p.AddResourceConfigurator("tencentcloud_cam_group_membership", func(r *tjconfig.Resource) {
		r.ExternalName = tjconfig.IdentifierFromProvider
		r.ShortGroup = shortGroupCam
		r.Kind = "GroupMembership"
		r.References["group_id"] = tjconfig.Reference{
			Type: "Group",
		}
	})

	p.AddResourceConfigurator("tencentcloud_cam_group_policy_attachment", func(r *tjconfig.Resource) {
		r.ExternalName = tjconfig.IdentifierFromProvider
		r.ShortGroup = shortGroupCam
		r.Kind = "GroupPolicyAttachment"
		r.References["group_id"] = tjconfig.Reference{
			Type: "Group",
		}
		r.References["policy_id"] = tjconfig.Reference{
			Type: "Policy",
		}
	})

	p.AddResourceConfigurator("tencentcloud_cam_user_policy_attachment", func(r *tjconfig.Resource) {
		r.ExternalName = tjconfig.IdentifierFromProvider
		r.ShortGroup = shortGroupCam
		r.Kind = "UserPolicyAttachment"
		r.References["user_id"] = tjconfig.Reference{
			Type: "User",
		}
		r.References["policy_id"] = tjconfig.Reference{
			Type: "Policy",
		}
	})

	p.AddResourceConfigurator("tencentcloud_cam_saml_proshortGroupCamider", func(r *tjconfig.Resource) {
		r.ExternalName = tjconfig.IdentifierFromProvider
		r.ShortGroup = shortGroupCam
		r.Kind = "SamlProshortGroupCamider"
	})

}
