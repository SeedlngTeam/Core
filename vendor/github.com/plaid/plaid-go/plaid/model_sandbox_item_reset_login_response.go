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

// SandboxItemResetLoginResponse SandboxItemResetLoginResponse defines the response schema for `/sandbox/item/reset_login`
type SandboxItemResetLoginResponse struct {
	// `true` if the call succeeded
	ResetLogin bool `json:"reset_login"`
	// A unique identifier for the request, which can be used for troubleshooting. This identifier, like all Plaid identifiers, is case sensitive.
	RequestId string `json:"request_id"`
	AdditionalProperties map[string]interface{}
}

type _SandboxItemResetLoginResponse SandboxItemResetLoginResponse

// NewSandboxItemResetLoginResponse instantiates a new SandboxItemResetLoginResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSandboxItemResetLoginResponse(resetLogin bool, requestId string) *SandboxItemResetLoginResponse {
	this := SandboxItemResetLoginResponse{}
	this.ResetLogin = resetLogin
	this.RequestId = requestId
	return &this
}

// NewSandboxItemResetLoginResponseWithDefaults instantiates a new SandboxItemResetLoginResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSandboxItemResetLoginResponseWithDefaults() *SandboxItemResetLoginResponse {
	this := SandboxItemResetLoginResponse{}
	return &this
}

// GetResetLogin returns the ResetLogin field value
func (o *SandboxItemResetLoginResponse) GetResetLogin() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.ResetLogin
}

// GetResetLoginOk returns a tuple with the ResetLogin field value
// and a boolean to check if the value has been set.
func (o *SandboxItemResetLoginResponse) GetResetLoginOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ResetLogin, true
}

// SetResetLogin sets field value
func (o *SandboxItemResetLoginResponse) SetResetLogin(v bool) {
	o.ResetLogin = v
}

// GetRequestId returns the RequestId field value
func (o *SandboxItemResetLoginResponse) GetRequestId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RequestId
}

// GetRequestIdOk returns a tuple with the RequestId field value
// and a boolean to check if the value has been set.
func (o *SandboxItemResetLoginResponse) GetRequestIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.RequestId, true
}

// SetRequestId sets field value
func (o *SandboxItemResetLoginResponse) SetRequestId(v string) {
	o.RequestId = v
}

func (o SandboxItemResetLoginResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["reset_login"] = o.ResetLogin
	}
	if true {
		toSerialize["request_id"] = o.RequestId
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *SandboxItemResetLoginResponse) UnmarshalJSON(bytes []byte) (err error) {
	varSandboxItemResetLoginResponse := _SandboxItemResetLoginResponse{}

	if err = json.Unmarshal(bytes, &varSandboxItemResetLoginResponse); err == nil {
		*o = SandboxItemResetLoginResponse(varSandboxItemResetLoginResponse)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "reset_login")
		delete(additionalProperties, "request_id")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableSandboxItemResetLoginResponse struct {
	value *SandboxItemResetLoginResponse
	isSet bool
}

func (v NullableSandboxItemResetLoginResponse) Get() *SandboxItemResetLoginResponse {
	return v.value
}

func (v *NullableSandboxItemResetLoginResponse) Set(val *SandboxItemResetLoginResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableSandboxItemResetLoginResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableSandboxItemResetLoginResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSandboxItemResetLoginResponse(val *SandboxItemResetLoginResponse) *NullableSandboxItemResetLoginResponse {
	return &NullableSandboxItemResetLoginResponse{value: val, isSet: true}
}

func (v NullableSandboxItemResetLoginResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSandboxItemResetLoginResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


