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

type ApiObservation struct {
	CreateTime *string `json:"createTime,omitempty" tf:"create_time,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	UpdateTime *string `json:"updateTime,omitempty" tf:"update_time,omitempty"`
}

type ApiParameters struct {

	// Custom API description.
	// +kubebuilder:validation:Optional
	APIDesc *string `json:"apiDesc,omitempty" tf:"api_desc,omitempty"`

	// Custom API name.
	// +kubebuilder:validation:Required
	APIName *string `json:"apiName" tf:"api_name,omitempty"`

	// API authentication type. Valid values: `SECRET` (key pair authentication),`NONE` (no authentication). Default value: `NONE`.
	// +kubebuilder:validation:Optional
	AuthType *string `json:"authType,omitempty" tf:"auth_type,omitempty"`

	// Whether to enable CORS. Default value: `true`.
	// +kubebuilder:validation:Optional
	EnableCors *bool `json:"enableCors,omitempty" tf:"enable_cors,omitempty"`

	// API QPS value. Enter a positive number to limit the API query rate per second `QPS`.
	// +kubebuilder:validation:Optional
	PreLimit *float64 `json:"preLimit,omitempty" tf:"pre_limit,omitempty"`

	// API frontend request type. Valid values: `HTTP`, `WEBSOCKET`. Default value: `HTTP`.
	// +kubebuilder:validation:Optional
	Protocol *string `json:"protocol,omitempty" tf:"protocol,omitempty"`

	// API QPS value. Enter a positive number to limit the API query rate per second `QPS`.
	// +kubebuilder:validation:Optional
	ReleaseLimit *float64 `json:"releaseLimit,omitempty" tf:"release_limit,omitempty"`

	// Request frontend method configuration. Valid values: `GET`,`POST`,`PUT`,`DELETE`,`HEAD`,`ANY`. Default value: `GET`.
	// +kubebuilder:validation:Optional
	RequestConfigMethod *string `json:"requestConfigMethod,omitempty" tf:"request_config_method,omitempty"`

	// Request frontend path configuration. Like `/user/getinfo`.
	// +kubebuilder:validation:Required
	RequestConfigPath *string `json:"requestConfigPath" tf:"request_config_path,omitempty"`

	// Frontend request parameters.
	// +kubebuilder:validation:Optional
	RequestParameters []RequestParametersParameters `json:"requestParameters,omitempty" tf:"request_parameters,omitempty"`

	// Custom error code configuration. Must keep at least one after set.
	// +kubebuilder:validation:Optional
	ResponseErrorCodes []ResponseErrorCodesParameters `json:"responseErrorCodes,omitempty" tf:"response_error_codes,omitempty"`

	// Response failure sample of custom response configuration.
	// +kubebuilder:validation:Optional
	ResponseFailExample *string `json:"responseFailExample,omitempty" tf:"response_fail_example,omitempty"`

	// Successful response sample of custom response configuration.
	// +kubebuilder:validation:Optional
	ResponseSuccessExample *string `json:"responseSuccessExample,omitempty" tf:"response_success_example,omitempty"`

	// Return type. Valid values: `HTML`, `JSON`, `TEXT`, `BINARY`, `XML`. Default value: `HTML`.
	// +kubebuilder:validation:Optional
	ResponseType *string `json:"responseType,omitempty" tf:"response_type,omitempty"`

	// API backend service request method, such as `GET`. If `service_config_type` is `HTTP`, this parameter will be required. The frontend `request_config_method` and backend method `service_config_method` can be different.
	// +kubebuilder:validation:Optional
	ServiceConfigMethod *string `json:"serviceConfigMethod,omitempty" tf:"service_config_method,omitempty"`

	// Returned information of API backend mocking. This parameter is required when `service_config_type` is `MOCK`.
	// +kubebuilder:validation:Optional
	ServiceConfigMockReturnMessage *string `json:"serviceConfigMockReturnMessage,omitempty" tf:"service_config_mock_return_message,omitempty"`

	// API backend service path, such as /path. If `service_config_type` is `HTTP`, this parameter will be required. The frontend `request_config_path` and backend path `service_config_path` can be different.
	// +kubebuilder:validation:Optional
	ServiceConfigPath *string `json:"serviceConfigPath,omitempty" tf:"service_config_path,omitempty"`

	// Backend type. This parameter takes effect when VPC is enabled. Currently, only `clb` is supported.
	// +kubebuilder:validation:Optional
	ServiceConfigProduct *string `json:"serviceConfigProduct,omitempty" tf:"service_config_product,omitempty"`

	// SCF function name. This parameter takes effect when `service_config_type` is `SCF`.
	// +kubebuilder:validation:Optional
	ServiceConfigScfFunctionName *string `json:"serviceConfigScfFunctionName,omitempty" tf:"service_config_scf_function_name,omitempty"`

	// SCF function namespace. This parameter takes effect when `service_config_type` is `SCF`.
	// +kubebuilder:validation:Optional
	ServiceConfigScfFunctionNamespace *string `json:"serviceConfigScfFunctionNamespace,omitempty" tf:"service_config_scf_function_namespace,omitempty"`

	// SCF function version. This parameter takes effect when `service_config_type` is `SCF`.
	// +kubebuilder:validation:Optional
	ServiceConfigScfFunctionQualifier *string `json:"serviceConfigScfFunctionQualifier,omitempty" tf:"service_config_scf_function_qualifier,omitempty"`

	// API backend service timeout period in seconds. Default value: `5`.
	// +kubebuilder:validation:Optional
	ServiceConfigTimeout *float64 `json:"serviceConfigTimeout,omitempty" tf:"service_config_timeout,omitempty"`

	// API backend service type. Valid values: `WEBSOCKET`, `HTTP`, `SCF`, `MOCK`. Default value: `HTTP`.
	// +kubebuilder:validation:Optional
	ServiceConfigType *string `json:"serviceConfigType,omitempty" tf:"service_config_type,omitempty"`

	// API backend service url. This parameter is required when `service_config_type` is `HTTP`.
	// +kubebuilder:validation:Optional
	ServiceConfigURL *string `json:"serviceConfigUrl,omitempty" tf:"service_config_url,omitempty"`

	// Unique VPC ID.
	// +kubebuilder:validation:Optional
	ServiceConfigVPCID *string `json:"serviceConfigVpcId,omitempty" tf:"service_config_vpc_id,omitempty"`

	// Which service this API belongs. Refer to resource `tencentcloud_api_gateway_service`.
	// +crossplane:generate:reference:type=Service
	// +kubebuilder:validation:Optional
	ServiceID *string `json:"serviceId,omitempty" tf:"service_id,omitempty"`

	// +kubebuilder:validation:Optional
	ServiceIDRef *v1.Reference `json:"serviceIdRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	ServiceIDSelector *v1.Selector `json:"serviceIdSelector,omitempty" tf:"-"`

	// API QPS value. Enter a positive number to limit the API query rate per second `QPS`.
	// +kubebuilder:validation:Optional
	TestLimit *float64 `json:"testLimit,omitempty" tf:"test_limit,omitempty"`
}

type RequestParametersObservation struct {
}

type RequestParametersParameters struct {

	// Parameter default value.
	// +kubebuilder:validation:Optional
	DefaultValue *string `json:"defaultValue,omitempty" tf:"default_value,omitempty"`

	// Parameter description.
	// +kubebuilder:validation:Optional
	Desc *string `json:"desc,omitempty" tf:"desc,omitempty"`

	// Parameter name.
	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// Parameter location.
	// +kubebuilder:validation:Required
	Position *string `json:"position" tf:"position,omitempty"`

	// If this parameter required. Default value: `false`.
	// +kubebuilder:validation:Optional
	Required *bool `json:"required,omitempty" tf:"required,omitempty"`

	// Parameter type.
	// +kubebuilder:validation:Required
	Type *string `json:"type" tf:"type,omitempty"`
}

type ResponseErrorCodesObservation struct {
}

type ResponseErrorCodesParameters struct {

	// Custom response configuration error code.
	// +kubebuilder:validation:Required
	Code *float64 `json:"code" tf:"code,omitempty"`

	// Custom error code conversion.
	// +kubebuilder:validation:Optional
	ConvertedCode *float64 `json:"convertedCode,omitempty" tf:"converted_code,omitempty"`

	// Parameter description.
	// +kubebuilder:validation:Optional
	Desc *string `json:"desc,omitempty" tf:"desc,omitempty"`

	// Custom response configuration error message.
	// +kubebuilder:validation:Required
	Msg *string `json:"msg" tf:"msg,omitempty"`

	// Whether to enable error code conversion. Default value: `false`.
	// +kubebuilder:validation:Optional
	NeedConvert *bool `json:"needConvert,omitempty" tf:"need_convert,omitempty"`
}

// ApiSpec defines the desired state of Api
type ApiSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ApiParameters `json:"forProvider"`
}

// ApiStatus defines the observed state of Api.
type ApiStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ApiObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Api is the Schema for the Apis API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,tencentcloudjet}
type Api struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ApiSpec   `json:"spec"`
	Status            ApiStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ApiList contains a list of Apis
type ApiList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Api `json:"items"`
}

// Repository type metadata.
var (
	Api_Kind             = "Api"
	Api_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Api_Kind}.String()
	Api_KindAPIVersion   = Api_Kind + "." + CRDGroupVersion.String()
	Api_GroupVersionKind = CRDGroupVersion.WithKind(Api_Kind)
)

func init() {
	SchemeBuilder.Register(&Api{}, &ApiList{})
}
