/*
Copyright 2022 Upbound Inc.
*/

// Code generated by upjet. DO NOT EDIT.

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type OidcSSOInitParameters struct {

	// Authorization request Endpoint, OpenID Connect identity provider authorization address. Corresponds to the value of the authorization_endpoint field in the Openid-configuration provided by the Enterprise IdP.
	// Authorization request Endpoint, OpenID Connect identity provider authorization address. Corresponds to the value of the `authorization_endpoint` field in the Openid-configuration provided by the Enterprise IdP.
	AuthorizationEndpoint *string `json:"authorizationEndpoint,omitempty" tf:"authorization_endpoint,omitempty"`

	// Client ID, the client ID registered with the OpenID Connect identity provider.
	// Client ID, the client ID registered with the OpenID Connect identity provider.
	ClientID *string `json:"clientId,omitempty" tf:"client_id,omitempty"`

	// The signature public key requires base64_encode. Verify the public key signed by the OpenID Connect identity provider ID Token. For the security of your account, we recommend that you rotate the signed public key regularly.
	// The signature public key requires base64_encode. Verify the public key signed by the OpenID Connect identity provider ID Token. For the security of your account, we recommend that you rotate the signed public key regularly.
	IdentityKey *string `json:"identityKey,omitempty" tf:"identity_key,omitempty"`

	// Identity provider URL. OpenID Connect identity provider identity.Corresponds to the value of the issuer field in the Openid-configuration provided by the Enterprise IdP.
	// Identity provider URL. OpenID Connect identity provider identity.Corresponds to the value of the `issuer` field in the Openid-configuration provided by the Enterprise IdP.
	IdentityURL *string `json:"identityUrl,omitempty" tf:"identity_url,omitempty"`

	// Map field names. Which field in the IdP's id_token maps to the user name of the subuser, usually the sub or name field.
	// Map field names. Which field in the IdP's id_token maps to the user name of the subuser, usually the sub or name field.
	MappingFiled *string `json:"mappingFiled,omitempty" tf:"mapping_filed,omitempty"`

	// Authorize the request Forsonse mode. Authorization request return mode, form_post and frogment two optional modes, recommended to select form_post mode.
	// Authorize the request Forsonse mode. Authorization request return mode, form_post and frogment two optional modes, recommended to select form_post mode.
	ResponseMode *string `json:"responseMode,omitempty" tf:"response_mode,omitempty"`

	// Authorization requests The Response type, with a fixed value id_token.
	// Authorization requests The Response type, with a fixed value id_token.
	ResponseType *string `json:"responseType,omitempty" tf:"response_type,omitempty"`

	// Authorize the request Scope. openid; email; profile; Authorization request information scope. The default is required openid.
	// Authorize the request Scope. openid; email; profile; Authorization request information scope. The default is required openid.
	// +listType=set
	Scope []*string `json:"scope,omitempty" tf:"scope,omitempty"`
}

type OidcSSOObservation struct {

	// Authorization request Endpoint, OpenID Connect identity provider authorization address. Corresponds to the value of the authorization_endpoint field in the Openid-configuration provided by the Enterprise IdP.
	// Authorization request Endpoint, OpenID Connect identity provider authorization address. Corresponds to the value of the `authorization_endpoint` field in the Openid-configuration provided by the Enterprise IdP.
	AuthorizationEndpoint *string `json:"authorizationEndpoint,omitempty" tf:"authorization_endpoint,omitempty"`

	// Client ID, the client ID registered with the OpenID Connect identity provider.
	// Client ID, the client ID registered with the OpenID Connect identity provider.
	ClientID *string `json:"clientId,omitempty" tf:"client_id,omitempty"`

	// ID of the resource.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The signature public key requires base64_encode. Verify the public key signed by the OpenID Connect identity provider ID Token. For the security of your account, we recommend that you rotate the signed public key regularly.
	// The signature public key requires base64_encode. Verify the public key signed by the OpenID Connect identity provider ID Token. For the security of your account, we recommend that you rotate the signed public key regularly.
	IdentityKey *string `json:"identityKey,omitempty" tf:"identity_key,omitempty"`

	// Identity provider URL. OpenID Connect identity provider identity.Corresponds to the value of the issuer field in the Openid-configuration provided by the Enterprise IdP.
	// Identity provider URL. OpenID Connect identity provider identity.Corresponds to the value of the `issuer` field in the Openid-configuration provided by the Enterprise IdP.
	IdentityURL *string `json:"identityUrl,omitempty" tf:"identity_url,omitempty"`

	// Map field names. Which field in the IdP's id_token maps to the user name of the subuser, usually the sub or name field.
	// Map field names. Which field in the IdP's id_token maps to the user name of the subuser, usually the sub or name field.
	MappingFiled *string `json:"mappingFiled,omitempty" tf:"mapping_filed,omitempty"`

	// Authorize the request Forsonse mode. Authorization request return mode, form_post and frogment two optional modes, recommended to select form_post mode.
	// Authorize the request Forsonse mode. Authorization request return mode, form_post and frogment two optional modes, recommended to select form_post mode.
	ResponseMode *string `json:"responseMode,omitempty" tf:"response_mode,omitempty"`

	// Authorization requests The Response type, with a fixed value id_token.
	// Authorization requests The Response type, with a fixed value id_token.
	ResponseType *string `json:"responseType,omitempty" tf:"response_type,omitempty"`

	// Authorize the request Scope. openid; email; profile; Authorization request information scope. The default is required openid.
	// Authorize the request Scope. openid; email; profile; Authorization request information scope. The default is required openid.
	// +listType=set
	Scope []*string `json:"scope,omitempty" tf:"scope,omitempty"`
}

type OidcSSOParameters struct {

	// Authorization request Endpoint, OpenID Connect identity provider authorization address. Corresponds to the value of the authorization_endpoint field in the Openid-configuration provided by the Enterprise IdP.
	// Authorization request Endpoint, OpenID Connect identity provider authorization address. Corresponds to the value of the `authorization_endpoint` field in the Openid-configuration provided by the Enterprise IdP.
	// +kubebuilder:validation:Optional
	AuthorizationEndpoint *string `json:"authorizationEndpoint,omitempty" tf:"authorization_endpoint,omitempty"`

	// Client ID, the client ID registered with the OpenID Connect identity provider.
	// Client ID, the client ID registered with the OpenID Connect identity provider.
	// +kubebuilder:validation:Optional
	ClientID *string `json:"clientId,omitempty" tf:"client_id,omitempty"`

	// The signature public key requires base64_encode. Verify the public key signed by the OpenID Connect identity provider ID Token. For the security of your account, we recommend that you rotate the signed public key regularly.
	// The signature public key requires base64_encode. Verify the public key signed by the OpenID Connect identity provider ID Token. For the security of your account, we recommend that you rotate the signed public key regularly.
	// +kubebuilder:validation:Optional
	IdentityKey *string `json:"identityKey,omitempty" tf:"identity_key,omitempty"`

	// Identity provider URL. OpenID Connect identity provider identity.Corresponds to the value of the issuer field in the Openid-configuration provided by the Enterprise IdP.
	// Identity provider URL. OpenID Connect identity provider identity.Corresponds to the value of the `issuer` field in the Openid-configuration provided by the Enterprise IdP.
	// +kubebuilder:validation:Optional
	IdentityURL *string `json:"identityUrl,omitempty" tf:"identity_url,omitempty"`

	// Map field names. Which field in the IdP's id_token maps to the user name of the subuser, usually the sub or name field.
	// Map field names. Which field in the IdP's id_token maps to the user name of the subuser, usually the sub or name field.
	// +kubebuilder:validation:Optional
	MappingFiled *string `json:"mappingFiled,omitempty" tf:"mapping_filed,omitempty"`

	// Authorize the request Forsonse mode. Authorization request return mode, form_post and frogment two optional modes, recommended to select form_post mode.
	// Authorize the request Forsonse mode. Authorization request return mode, form_post and frogment two optional modes, recommended to select form_post mode.
	// +kubebuilder:validation:Optional
	ResponseMode *string `json:"responseMode,omitempty" tf:"response_mode,omitempty"`

	// Authorization requests The Response type, with a fixed value id_token.
	// Authorization requests The Response type, with a fixed value id_token.
	// +kubebuilder:validation:Optional
	ResponseType *string `json:"responseType,omitempty" tf:"response_type,omitempty"`

	// Authorize the request Scope. openid; email; profile; Authorization request information scope. The default is required openid.
	// Authorize the request Scope. openid; email; profile; Authorization request information scope. The default is required openid.
	// +kubebuilder:validation:Optional
	// +listType=set
	Scope []*string `json:"scope,omitempty" tf:"scope,omitempty"`
}

// OidcSSOSpec defines the desired state of OidcSSO
type OidcSSOSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     OidcSSOParameters `json:"forProvider"`
	// THIS IS A BETA FIELD. It will be honored
	// unless the Management Policies feature flag is disabled.
	// InitProvider holds the same fields as ForProvider, with the exception
	// of Identifier and other resource reference fields. The fields that are
	// in InitProvider are merged into ForProvider when the resource is created.
	// The same fields are also added to the terraform ignore_changes hook, to
	// avoid updating them after creation. This is useful for fields that are
	// required on creation, but we do not desire to update them after creation,
	// for example because of an external controller is managing them, like an
	// autoscaler.
	InitProvider OidcSSOInitParameters `json:"initProvider,omitempty"`
}

// OidcSSOStatus defines the observed state of OidcSSO.
type OidcSSOStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        OidcSSOObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// OidcSSO is the Schema for the OidcSSOs API. Provides a resource to create a CAM-OIDC-SSO.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,tencentcloud}
type OidcSSO struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.authorizationEndpoint) || (has(self.initProvider) && has(self.initProvider.authorizationEndpoint))",message="spec.forProvider.authorizationEndpoint is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.clientId) || (has(self.initProvider) && has(self.initProvider.clientId))",message="spec.forProvider.clientId is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.identityKey) || (has(self.initProvider) && has(self.initProvider.identityKey))",message="spec.forProvider.identityKey is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.identityUrl) || (has(self.initProvider) && has(self.initProvider.identityUrl))",message="spec.forProvider.identityUrl is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.mappingFiled) || (has(self.initProvider) && has(self.initProvider.mappingFiled))",message="spec.forProvider.mappingFiled is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.responseMode) || (has(self.initProvider) && has(self.initProvider.responseMode))",message="spec.forProvider.responseMode is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.responseType) || (has(self.initProvider) && has(self.initProvider.responseType))",message="spec.forProvider.responseType is a required parameter"
	Spec   OidcSSOSpec   `json:"spec"`
	Status OidcSSOStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// OidcSSOList contains a list of OidcSSOs
type OidcSSOList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []OidcSSO `json:"items"`
}

// Repository type metadata.
var (
	OidcSSO_Kind             = "OidcSSO"
	OidcSSO_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: OidcSSO_Kind}.String()
	OidcSSO_KindAPIVersion   = OidcSSO_Kind + "." + CRDGroupVersion.String()
	OidcSSO_GroupVersionKind = CRDGroupVersion.WithKind(OidcSSO_Kind)
)

func init() {
	SchemeBuilder.Register(&OidcSSO{}, &OidcSSOList{})
}
