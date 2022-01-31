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

// IncomeVerificationPaystubGetResponse IncomeVerificationPaystubGetResponse defines the response schema for `/income/verification/paystub/get`.
type IncomeVerificationPaystubGetResponse struct {
	Paystub Paystub `json:"paystub"`
	Error NullableError `json:"error,omitempty"`
	// A unique identifier for the request, which can be used for troubleshooting. This identifier, like all Plaid identifiers, is case sensitive.
	RequestId string `json:"request_id"`
	AdditionalProperties map[string]interface{}
}

type _IncomeVerificationPaystubGetResponse IncomeVerificationPaystubGetResponse

// NewIncomeVerificationPaystubGetResponse instantiates a new IncomeVerificationPaystubGetResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIncomeVerificationPaystubGetResponse(paystub Paystub, requestId string) *IncomeVerificationPaystubGetResponse {
	this := IncomeVerificationPaystubGetResponse{}
	this.Paystub = paystub
	this.RequestId = requestId
	return &this
}

// NewIncomeVerificationPaystubGetResponseWithDefaults instantiates a new IncomeVerificationPaystubGetResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIncomeVerificationPaystubGetResponseWithDefaults() *IncomeVerificationPaystubGetResponse {
	this := IncomeVerificationPaystubGetResponse{}
	return &this
}

// GetPaystub returns the Paystub field value
func (o *IncomeVerificationPaystubGetResponse) GetPaystub() Paystub {
	if o == nil {
		var ret Paystub
		return ret
	}

	return o.Paystub
}

// GetPaystubOk returns a tuple with the Paystub field value
// and a boolean to check if the value has been set.
func (o *IncomeVerificationPaystubGetResponse) GetPaystubOk() (*Paystub, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Paystub, true
}

// SetPaystub sets field value
func (o *IncomeVerificationPaystubGetResponse) SetPaystub(v Paystub) {
	o.Paystub = v
}

// GetError returns the Error field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IncomeVerificationPaystubGetResponse) GetError() Error {
	if o == nil || o.Error.Get() == nil {
		var ret Error
		return ret
	}
	return *o.Error.Get()
}

// GetErrorOk returns a tuple with the Error field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IncomeVerificationPaystubGetResponse) GetErrorOk() (*Error, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Error.Get(), o.Error.IsSet()
}

// HasError returns a boolean if a field has been set.
func (o *IncomeVerificationPaystubGetResponse) HasError() bool {
	if o != nil && o.Error.IsSet() {
		return true
	}

	return false
}

// SetError gets a reference to the given NullableError and assigns it to the Error field.
func (o *IncomeVerificationPaystubGetResponse) SetError(v Error) {
	o.Error.Set(&v)
}
// SetErrorNil sets the value for Error to be an explicit nil
func (o *IncomeVerificationPaystubGetResponse) SetErrorNil() {
	o.Error.Set(nil)
}

// UnsetError ensures that no value is present for Error, not even an explicit nil
func (o *IncomeVerificationPaystubGetResponse) UnsetError() {
	o.Error.Unset()
}

// GetRequestId returns the RequestId field value
func (o *IncomeVerificationPaystubGetResponse) GetRequestId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RequestId
}

// GetRequestIdOk returns a tuple with the RequestId field value
// and a boolean to check if the value has been set.
func (o *IncomeVerificationPaystubGetResponse) GetRequestIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.RequestId, true
}

// SetRequestId sets field value
func (o *IncomeVerificationPaystubGetResponse) SetRequestId(v string) {
	o.RequestId = v
}

func (o IncomeVerificationPaystubGetResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["paystub"] = o.Paystub
	}
	if o.Error.IsSet() {
		toSerialize["error"] = o.Error.Get()
	}
	if true {
		toSerialize["request_id"] = o.RequestId
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *IncomeVerificationPaystubGetResponse) UnmarshalJSON(bytes []byte) (err error) {
	varIncomeVerificationPaystubGetResponse := _IncomeVerificationPaystubGetResponse{}

	if err = json.Unmarshal(bytes, &varIncomeVerificationPaystubGetResponse); err == nil {
		*o = IncomeVerificationPaystubGetResponse(varIncomeVerificationPaystubGetResponse)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "paystub")
		delete(additionalProperties, "error")
		delete(additionalProperties, "request_id")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableIncomeVerificationPaystubGetResponse struct {
	value *IncomeVerificationPaystubGetResponse
	isSet bool
}

func (v NullableIncomeVerificationPaystubGetResponse) Get() *IncomeVerificationPaystubGetResponse {
	return v.value
}

func (v *NullableIncomeVerificationPaystubGetResponse) Set(val *IncomeVerificationPaystubGetResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableIncomeVerificationPaystubGetResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableIncomeVerificationPaystubGetResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIncomeVerificationPaystubGetResponse(val *IncomeVerificationPaystubGetResponse) *NullableIncomeVerificationPaystubGetResponse {
	return &NullableIncomeVerificationPaystubGetResponse{value: val, isSet: true}
}

func (v NullableIncomeVerificationPaystubGetResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIncomeVerificationPaystubGetResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


