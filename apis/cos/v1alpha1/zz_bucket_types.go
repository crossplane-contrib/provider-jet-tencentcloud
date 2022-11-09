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

// Code generated by terrajet. DO NOT EDIT.

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type BucketObservation struct {
	CosBucketURL *string `json:"cosBucketUrl,omitempty" tf:"cos_bucket_url,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type BucketParameters struct {

	// The canned ACL to apply. Valid values: private, public-read, and public-read-write. Defaults to private.
	// +kubebuilder:validation:Optional
	ACL *string `json:"acl,omitempty" tf:"acl,omitempty"`

	// ACL XML body for multiple grant info. NOTE: this argument will overwrite `acl`. Check https://intl.cloud.tencent.com/document/product/436/7737 for more detail.
	// +kubebuilder:validation:Optional
	ACLBody *string `json:"aclBody,omitempty" tf:"acl_body,omitempty"`

	// The name of a bucket to be created. Bucket format should be [custom name]-[appid], for example `mycos-1258798060`.
	// +kubebuilder:validation:Required
	Bucket *string `json:"bucket" tf:"bucket,omitempty"`

	// A rule of Cross-Origin Resource Sharing (documented below).
	// +kubebuilder:validation:Optional
	CorsRules []CorsRulesParameters `json:"corsRules,omitempty" tf:"cors_rules,omitempty"`

	// The server-side encryption algorithm to use. Valid value is `AES256`.
	// +kubebuilder:validation:Optional
	EncryptionAlgorithm *string `json:"encryptionAlgorithm,omitempty" tf:"encryption_algorithm,omitempty"`

	// Force cleanup all objects before delete bucket.
	// +kubebuilder:validation:Optional
	ForceClean *bool `json:"forceClean,omitempty" tf:"force_clean,omitempty"`

	// A configuration of object lifecycle management (documented below).
	// +kubebuilder:validation:Optional
	LifecycleRules []LifecycleRulesParameters `json:"lifecycleRules,omitempty" tf:"lifecycle_rules,omitempty"`

	// Indicate the access log of this bucket to be saved or not. Default is `false`. If set `true`, the access log will be saved with `log_target_bucket`. To enable log, the full access of log service must be granted. [Full Access Role Policy](https://intl.cloud.tencent.com/document/product/436/16920).
	// +kubebuilder:validation:Optional
	LogEnable *bool `json:"logEnable,omitempty" tf:"log_enable,omitempty"`

	// The prefix log name which saves the access log of this bucket per 5 minutes. Eg. `MyLogPrefix/`. The log access file format is `log_target_bucket`/`log_prefix`{YYYY}/{MM}/{DD}/{time}_{random}_{index}.gz. Only valid when `log_enable` is `true`.
	// +kubebuilder:validation:Optional
	LogPrefix *string `json:"logPrefix,omitempty" tf:"log_prefix,omitempty"`

	// The target bucket name which saves the access log of this bucket per 5 minutes. The log access file format is `log_target_bucket`/`log_prefix`{YYYY}/{MM}/{DD}/{time}_{random}_{index}.gz. Only valid when `log_enable` is `true`. User must have full access on this bucket.
	// +kubebuilder:validation:Optional
	LogTargetBucket *string `json:"logTargetBucket,omitempty" tf:"log_target_bucket,omitempty"`

	// Indicates whether to create a bucket of multi available zone. NOTE: If set to true, the versioning must enable.
	// +kubebuilder:validation:Optional
	MultiAz *bool `json:"multiAz,omitempty" tf:"multi_az,omitempty"`

	// Bucket Origin Domain settings.
	// +kubebuilder:validation:Optional
	OriginDomainRules []OriginDomainRulesParameters `json:"originDomainRules,omitempty" tf:"origin_domain_rules,omitempty"`

	// Bucket Origin-Pull settings.
	// +kubebuilder:validation:Optional
	OriginPullRules []OriginPullRulesParameters `json:"originPullRules,omitempty" tf:"origin_pull_rules,omitempty"`

	// Request initiator identifier, format: `qcs::cam::uin/<owneruin>:uin/<subuin>`. NOTE: only `versioning_enable` is true can configure this argument.
	// +kubebuilder:validation:Optional
	ReplicaRole *string `json:"replicaRole,omitempty" tf:"replica_role,omitempty"`

	// List of replica rule. NOTE: only `versioning_enable` is true and `replica_role` set can configure this argument.
	// +kubebuilder:validation:Optional
	ReplicaRules []ReplicaRulesParameters `json:"replicaRules,omitempty" tf:"replica_rules,omitempty"`

	// The tags of a bucket.
	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// Enable bucket versioning.
	// +kubebuilder:validation:Optional
	VersioningEnable *bool `json:"versioningEnable,omitempty" tf:"versioning_enable,omitempty"`

	// A website object(documented below).
	// +kubebuilder:validation:Optional
	Website []WebsiteParameters `json:"website,omitempty" tf:"website,omitempty"`
}

type CorsRulesObservation struct {
}

type CorsRulesParameters struct {

	// Specifies which headers are allowed.
	// +kubebuilder:validation:Required
	AllowedHeaders []*string `json:"allowedHeaders" tf:"allowed_headers,omitempty"`

	// Specifies which methods are allowed. Can be `GET`, `PUT`, `POST`, `DELETE` or `HEAD`.
	// +kubebuilder:validation:Required
	AllowedMethods []*string `json:"allowedMethods" tf:"allowed_methods,omitempty"`

	// Specifies which origins are allowed.
	// +kubebuilder:validation:Required
	AllowedOrigins []*string `json:"allowedOrigins" tf:"allowed_origins,omitempty"`

	// Specifies expose header in the response.
	// +kubebuilder:validation:Optional
	ExposeHeaders []*string `json:"exposeHeaders,omitempty" tf:"expose_headers,omitempty"`

	// Specifies time in seconds that browser can cache the response for a preflight request.
	// +kubebuilder:validation:Optional
	MaxAgeSeconds *float64 `json:"maxAgeSeconds,omitempty" tf:"max_age_seconds,omitempty"`
}

type ExpirationObservation struct {
}

type ExpirationParameters struct {

	// Specifies the date after which you want the corresponding action to take effect.
	// +kubebuilder:validation:Optional
	Date *string `json:"date,omitempty" tf:"date,omitempty"`

	// Specifies the number of days after object creation when the specific rule action takes effect.
	// +kubebuilder:validation:Optional
	Days *float64 `json:"days,omitempty" tf:"days,omitempty"`

	// Indicates whether the delete marker of an expired object will be removed.
	// +kubebuilder:validation:Optional
	DeleteMarker *bool `json:"deleteMarker,omitempty" tf:"delete_marker,omitempty"`
}

type LifecycleRulesObservation struct {
}

type LifecycleRulesParameters struct {

	// Specifies a period in the object's expire (documented below).
	// +kubebuilder:validation:Optional
	Expiration []ExpirationParameters `json:"expiration,omitempty" tf:"expiration,omitempty"`

	// Object key prefix identifying one or more objects to which the rule applies.
	// +kubebuilder:validation:Required
	FilterPrefix *string `json:"filterPrefix" tf:"filter_prefix,omitempty"`

	// A unique identifier for the rule. It can be up to 255 characters.
	// +kubebuilder:validation:Optional
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Specifies when non current object versions shall expire.
	// +kubebuilder:validation:Optional
	NonCurrentExpiration []NonCurrentExpirationParameters `json:"nonCurrentExpiration,omitempty" tf:"non_current_expiration,omitempty"`

	// Specifies a period in the non current object's transitions.
	// +kubebuilder:validation:Optional
	NonCurrentTransition []NonCurrentTransitionParameters `json:"nonCurrentTransition,omitempty" tf:"non_current_transition,omitempty"`

	// Specifies a period in the object's transitions (documented below).
	// +kubebuilder:validation:Optional
	Transition []TransitionParameters `json:"transition,omitempty" tf:"transition,omitempty"`
}

type NonCurrentExpirationObservation struct {
}

type NonCurrentExpirationParameters struct {

	// Number of days after non current object creation when the specific rule action takes effect. The maximum value is 3650.
	// +kubebuilder:validation:Optional
	NonCurrentDays *float64 `json:"nonCurrentDays,omitempty" tf:"non_current_days,omitempty"`
}

type NonCurrentTransitionObservation struct {
}

type NonCurrentTransitionParameters struct {

	// Number of days after non current object creation when the specific rule action takes effect.
	// +kubebuilder:validation:Optional
	NonCurrentDays *float64 `json:"nonCurrentDays,omitempty" tf:"non_current_days,omitempty"`

	// Specifies the storage class to which you want the non current object to transition. Available values include `STANDARD`, `STANDARD_IA` and `ARCHIVE`.
	// +kubebuilder:validation:Required
	StorageClass *string `json:"storageClass" tf:"storage_class,omitempty"`
}

type OriginDomainRulesObservation struct {
}

type OriginDomainRulesParameters struct {

	// Specify domain host.
	// +kubebuilder:validation:Required
	Domain *string `json:"domain" tf:"domain,omitempty"`

	// Domain status, default: `ENABLED`.
	// +kubebuilder:validation:Optional
	Status *string `json:"status,omitempty" tf:"status,omitempty"`

	// Specify origin domain type, available values: `REST`, `WEBSITE`, `ACCELERATE`, default: `REST`.
	// +kubebuilder:validation:Optional
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

type OriginPullRulesObservation struct {
}

type OriginPullRulesParameters struct {

	// Specifies the custom headers that you can add for COS to access your origin server.
	// +kubebuilder:validation:Optional
	CustomHTTPHeaders map[string]*string `json:"customHttpHeaders,omitempty" tf:"custom_http_headers,omitempty"`

	// Specifies the pass through headers when accessing the origin server.
	// +kubebuilder:validation:Optional
	FollowHTTPHeaders []*string `json:"followHttpHeaders,omitempty" tf:"follow_http_headers,omitempty"`

	// Specifies whether to pass through COS request query string when accessing the origin server.
	// +kubebuilder:validation:Optional
	FollowQueryString *bool `json:"followQueryString,omitempty" tf:"follow_query_string,omitempty"`

	// Specifies whether to follow 3XX redirect to another origin server to pull data from.
	// +kubebuilder:validation:Optional
	FollowRedirection *bool `json:"followRedirection,omitempty" tf:"follow_redirection,omitempty"`

	// Allows only a domain name or IP address. You can optionally append a port number to the address.
	// +kubebuilder:validation:Required
	Host *string `json:"host" tf:"host,omitempty"`

	// Triggers the origin-pull rule when the requested file name matches this prefix.
	// +kubebuilder:validation:Optional
	Prefix *string `json:"prefix,omitempty" tf:"prefix,omitempty"`

	// Priority of origin-pull rules, do not set the same value for multiple rules.
	// +kubebuilder:validation:Required
	Priority *float64 `json:"priority" tf:"priority,omitempty"`

	// the protocol used for COS to access the specified origin server. The available value include `HTTP`, `HTTPS` and `FOLLOW`.
	// +kubebuilder:validation:Optional
	Protocol *string `json:"protocol,omitempty" tf:"protocol,omitempty"`

	// If `true`, COS will not return 3XX status code when pulling data from an origin server. Current available zone: ap-beijing, ap-shanghai, ap-singapore, ap-mumbai.
	// +kubebuilder:validation:Optional
	SyncBackToSource *bool `json:"syncBackToSource,omitempty" tf:"sync_back_to_source,omitempty"`
}

type ReplicaRulesObservation struct {
}

type ReplicaRulesParameters struct {

	// Destination bucket identifier, format: `qcs::cos:<region>::<bucketname-appid>`. NOTE: destination bucket must enable versioning.
	// +kubebuilder:validation:Required
	DestinationBucket *string `json:"destinationBucket" tf:"destination_bucket,omitempty"`

	// Storage class of destination, available values: `STANDARD`, `INTELLIGENT_TIERING`, `STANDARD_IA`. default is following current class of destination.
	// +kubebuilder:validation:Optional
	DestinationStorageClass *string `json:"destinationStorageClass,omitempty" tf:"destination_storage_class,omitempty"`

	// Name of a specific rule.
	// +kubebuilder:validation:Optional
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Prefix matching policy. Policies cannot overlap; otherwise, an error will be returned. To match the root directory, leave this parameter empty.
	// +kubebuilder:validation:Optional
	Prefix *string `json:"prefix,omitempty" tf:"prefix,omitempty"`

	// Status identifier, available values: `Enabled`, `Disabled`.
	// +kubebuilder:validation:Required
	Status *string `json:"status" tf:"status,omitempty"`
}

type TransitionObservation struct {
}

type TransitionParameters struct {

	// Specifies the date after which you want the corresponding action to take effect.
	// +kubebuilder:validation:Optional
	Date *string `json:"date,omitempty" tf:"date,omitempty"`

	// Specifies the number of days after object creation when the specific rule action takes effect.
	// +kubebuilder:validation:Optional
	Days *float64 `json:"days,omitempty" tf:"days,omitempty"`

	// Specifies the storage class to which you want the object to transition. Available values include `STANDARD`, `STANDARD_IA` and `ARCHIVE`.
	// +kubebuilder:validation:Required
	StorageClass *string `json:"storageClass" tf:"storage_class,omitempty"`
}

type WebsiteObservation struct {
}

type WebsiteParameters struct {

	// An absolute path to the document to return in case of a 4XX error.
	// +kubebuilder:validation:Optional
	ErrorDocument *string `json:"errorDocument,omitempty" tf:"error_document,omitempty"`

	// COS returns this index document when requests are made to the root domain or any of the subfolders.
	// +kubebuilder:validation:Optional
	IndexDocument *string `json:"indexDocument,omitempty" tf:"index_document,omitempty"`
}

// BucketSpec defines the desired state of Bucket
type BucketSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     BucketParameters `json:"forProvider"`
}

// BucketStatus defines the observed state of Bucket.
type BucketStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        BucketObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Bucket is the Schema for the Buckets API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,tencentcloudjet}
type Bucket struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              BucketSpec   `json:"spec"`
	Status            BucketStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// BucketList contains a list of Buckets
type BucketList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Bucket `json:"items"`
}

// Repository type metadata.
var (
	Bucket_Kind             = "Bucket"
	Bucket_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Bucket_Kind}.String()
	Bucket_KindAPIVersion   = Bucket_Kind + "." + CRDGroupVersion.String()
	Bucket_GroupVersionKind = CRDGroupVersion.WithKind(Bucket_Kind)
)

func init() {
	SchemeBuilder.Register(&Bucket{}, &BucketList{})
}
