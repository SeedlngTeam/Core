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
	"time"
)

// ProductStatus A representation of the status health of a request type. Auth requests, Balance requests, Identity requests, Investments requests, Liabilities requests, Transactions updates, Investments updates, Liabilities updates, and Item logins each have their own status object.
type ProductStatus struct {
	// `HEALTHY`: the majority of requests are successful `DEGRADED`: only some requests are successful `DOWN`: all requests are failing
	Status string `json:"status"`
	// [ISO 8601](https://wikipedia.org/wiki/ISO_8601) formatted timestamp of the last status change for the institution. 
	LastStatusChange time.Time `json:"last_status_change"`
	Breakdown ProductStatusBreakdown `json:"breakdown"`
	AdditionalProperties map[string]interface{}
}

type _ProductStatus ProductStatus

// NewProductStatus instantiates a new ProductStatus object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProductStatus(status string, lastStatusChange time.Time, breakdown ProductStatusBreakdown) *ProductStatus {
	this := ProductStatus{}
	this.Status = status
	this.LastStatusChange = lastStatusChange
	this.Breakdown = breakdown
	return &this
}

// NewProductStatusWithDefaults instantiates a new ProductStatus object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProductStatusWithDefaults() *ProductStatus {
	this := ProductStatus{}
	return &this
}

// GetStatus returns the Status field value
func (o *ProductStatus) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *ProductStatus) GetStatusOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *ProductStatus) SetStatus(v string) {
	o.Status = v
}

// GetLastStatusChange returns the LastStatusChange field value
func (o *ProductStatus) GetLastStatusChange() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.LastStatusChange
}

// GetLastStatusChangeOk returns a tuple with the LastStatusChange field value
// and a boolean to check if the value has been set.
func (o *ProductStatus) GetLastStatusChangeOk() (*time.Time, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.LastStatusChange, true
}

// SetLastStatusChange sets field value
func (o *ProductStatus) SetLastStatusChange(v time.Time) {
	o.LastStatusChange = v
}

// GetBreakdown returns the Breakdown field value
func (o *ProductStatus) GetBreakdown() ProductStatusBreakdown {
	if o == nil {
		var ret ProductStatusBreakdown
		return ret
	}

	return o.Breakdown
}

// GetBreakdownOk returns a tuple with the Breakdown field value
// and a boolean to check if the value has been set.
func (o *ProductStatus) GetBreakdownOk() (*ProductStatusBreakdown, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Breakdown, true
}

// SetBreakdown sets field value
func (o *ProductStatus) SetBreakdown(v ProductStatusBreakdown) {
	o.Breakdown = v
}

func (o ProductStatus) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["status"] = o.Status
	}
	if true {
		toSerialize["last_status_change"] = o.LastStatusChange
	}
	if true {
		toSerialize["breakdown"] = o.Breakdown
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *ProductStatus) UnmarshalJSON(bytes []byte) (err error) {
	varProductStatus := _ProductStatus{}

	if err = json.Unmarshal(bytes, &varProductStatus); err == nil {
		*o = ProductStatus(varProductStatus)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "status")
		delete(additionalProperties, "last_status_change")
		delete(additionalProperties, "breakdown")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableProductStatus struct {
	value *ProductStatus
	isSet bool
}

func (v NullableProductStatus) Get() *ProductStatus {
	return v.value
}

func (v *NullableProductStatus) Set(val *ProductStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableProductStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableProductStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProductStatus(val *ProductStatus) *NullableProductStatus {
	return &NullableProductStatus{value: val, isSet: true}
}

func (v NullableProductStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProductStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


