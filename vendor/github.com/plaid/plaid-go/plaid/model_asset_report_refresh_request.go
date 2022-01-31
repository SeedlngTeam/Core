/*
 * The Plaid API
 *
 * The Plaid REST API. Please see https://plaid.com/docs/api for more details.
 *
 * API version: 2020-09-14_1.33.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package plaid

import (
	"encoding/json"
)

// AssetReportRefreshRequest AssetReportRefreshRequest defines the request schema for `/asset_report/refresh`
type AssetReportRefreshRequest struct {
	// Your Plaid API `client_id`. The `client_id` is required and may be provided either in the `PLAID-CLIENT-ID` header or as part of a request body.
	ClientId *string `json:"client_id,omitempty"`
	// Your Plaid API `secret`. The `secret` is required and may be provided either in the `PLAID-SECRET` header or as part of a request body.
	Secret *string `json:"secret,omitempty"`
	// The `asset_report_token` returned by the original call to `/asset_report/create`
	AssetReportToken string `json:"asset_report_token"`
	// The maximum number of days of history to include in the Asset Report. Must be an integer. If not specified, the value from the original call to `/asset_report/create` will be used.
	DaysRequested *int32 `json:"days_requested,omitempty"`
	Options *AssetReportRefreshRequestOptions `json:"options,omitempty"`
}

// NewAssetReportRefreshRequest instantiates a new AssetReportRefreshRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAssetReportRefreshRequest(assetReportToken string) *AssetReportRefreshRequest {
	this := AssetReportRefreshRequest{}
	this.AssetReportToken = assetReportToken
	return &this
}

// NewAssetReportRefreshRequestWithDefaults instantiates a new AssetReportRefreshRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAssetReportRefreshRequestWithDefaults() *AssetReportRefreshRequest {
	this := AssetReportRefreshRequest{}
	return &this
}

// GetClientId returns the ClientId field value if set, zero value otherwise.
func (o *AssetReportRefreshRequest) GetClientId() string {
	if o == nil || o.ClientId == nil {
		var ret string
		return ret
	}
	return *o.ClientId
}

// GetClientIdOk returns a tuple with the ClientId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetReportRefreshRequest) GetClientIdOk() (*string, bool) {
	if o == nil || o.ClientId == nil {
		return nil, false
	}
	return o.ClientId, true
}

// HasClientId returns a boolean if a field has been set.
func (o *AssetReportRefreshRequest) HasClientId() bool {
	if o != nil && o.ClientId != nil {
		return true
	}

	return false
}

// SetClientId gets a reference to the given string and assigns it to the ClientId field.
func (o *AssetReportRefreshRequest) SetClientId(v string) {
	o.ClientId = &v
}

// GetSecret returns the Secret field value if set, zero value otherwise.
func (o *AssetReportRefreshRequest) GetSecret() string {
	if o == nil || o.Secret == nil {
		var ret string
		return ret
	}
	return *o.Secret
}

// GetSecretOk returns a tuple with the Secret field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetReportRefreshRequest) GetSecretOk() (*string, bool) {
	if o == nil || o.Secret == nil {
		return nil, false
	}
	return o.Secret, true
}

// HasSecret returns a boolean if a field has been set.
func (o *AssetReportRefreshRequest) HasSecret() bool {
	if o != nil && o.Secret != nil {
		return true
	}

	return false
}

// SetSecret gets a reference to the given string and assigns it to the Secret field.
func (o *AssetReportRefreshRequest) SetSecret(v string) {
	o.Secret = &v
}

// GetAssetReportToken returns the AssetReportToken field value
func (o *AssetReportRefreshRequest) GetAssetReportToken() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AssetReportToken
}

// GetAssetReportTokenOk returns a tuple with the AssetReportToken field value
// and a boolean to check if the value has been set.
func (o *AssetReportRefreshRequest) GetAssetReportTokenOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.AssetReportToken, true
}

// SetAssetReportToken sets field value
func (o *AssetReportRefreshRequest) SetAssetReportToken(v string) {
	o.AssetReportToken = v
}

// GetDaysRequested returns the DaysRequested field value if set, zero value otherwise.
func (o *AssetReportRefreshRequest) GetDaysRequested() int32 {
	if o == nil || o.DaysRequested == nil {
		var ret int32
		return ret
	}
	return *o.DaysRequested
}

// GetDaysRequestedOk returns a tuple with the DaysRequested field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetReportRefreshRequest) GetDaysRequestedOk() (*int32, bool) {
	if o == nil || o.DaysRequested == nil {
		return nil, false
	}
	return o.DaysRequested, true
}

// HasDaysRequested returns a boolean if a field has been set.
func (o *AssetReportRefreshRequest) HasDaysRequested() bool {
	if o != nil && o.DaysRequested != nil {
		return true
	}

	return false
}

// SetDaysRequested gets a reference to the given int32 and assigns it to the DaysRequested field.
func (o *AssetReportRefreshRequest) SetDaysRequested(v int32) {
	o.DaysRequested = &v
}

// GetOptions returns the Options field value if set, zero value otherwise.
func (o *AssetReportRefreshRequest) GetOptions() AssetReportRefreshRequestOptions {
	if o == nil || o.Options == nil {
		var ret AssetReportRefreshRequestOptions
		return ret
	}
	return *o.Options
}

// GetOptionsOk returns a tuple with the Options field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetReportRefreshRequest) GetOptionsOk() (*AssetReportRefreshRequestOptions, bool) {
	if o == nil || o.Options == nil {
		return nil, false
	}
	return o.Options, true
}

// HasOptions returns a boolean if a field has been set.
func (o *AssetReportRefreshRequest) HasOptions() bool {
	if o != nil && o.Options != nil {
		return true
	}

	return false
}

// SetOptions gets a reference to the given AssetReportRefreshRequestOptions and assigns it to the Options field.
func (o *AssetReportRefreshRequest) SetOptions(v AssetReportRefreshRequestOptions) {
	o.Options = &v
}

func (o AssetReportRefreshRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ClientId != nil {
		toSerialize["client_id"] = o.ClientId
	}
	if o.Secret != nil {
		toSerialize["secret"] = o.Secret
	}
	if true {
		toSerialize["asset_report_token"] = o.AssetReportToken
	}
	if o.DaysRequested != nil {
		toSerialize["days_requested"] = o.DaysRequested
	}
	if o.Options != nil {
		toSerialize["options"] = o.Options
	}
	return json.Marshal(toSerialize)
}

type NullableAssetReportRefreshRequest struct {
	value *AssetReportRefreshRequest
	isSet bool
}

func (v NullableAssetReportRefreshRequest) Get() *AssetReportRefreshRequest {
	return v.value
}

func (v *NullableAssetReportRefreshRequest) Set(val *AssetReportRefreshRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableAssetReportRefreshRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableAssetReportRefreshRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAssetReportRefreshRequest(val *AssetReportRefreshRequest) *NullableAssetReportRefreshRequest {
	return &NullableAssetReportRefreshRequest{value: val, isSet: true}
}

func (v NullableAssetReportRefreshRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAssetReportRefreshRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


