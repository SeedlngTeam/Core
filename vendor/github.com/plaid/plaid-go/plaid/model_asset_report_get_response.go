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

// AssetReportGetResponse AssetReportGetResponse defines the response schema for `/asset_report/get`
type AssetReportGetResponse struct {
	Report AssetReport `json:"report"`
	// If the Asset Report generation was successful but identity information cannot be returned, this array will contain information about the errors causing identity information to be missing
	Warnings []Warning `json:"warnings"`
	// A unique identifier for the request, which can be used for troubleshooting. This identifier, like all Plaid identifiers, is case sensitive.
	RequestId string `json:"request_id"`
	AdditionalProperties map[string]interface{}
}

type _AssetReportGetResponse AssetReportGetResponse

// NewAssetReportGetResponse instantiates a new AssetReportGetResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAssetReportGetResponse(report AssetReport, warnings []Warning, requestId string) *AssetReportGetResponse {
	this := AssetReportGetResponse{}
	this.Report = report
	this.Warnings = warnings
	this.RequestId = requestId
	return &this
}

// NewAssetReportGetResponseWithDefaults instantiates a new AssetReportGetResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAssetReportGetResponseWithDefaults() *AssetReportGetResponse {
	this := AssetReportGetResponse{}
	return &this
}

// GetReport returns the Report field value
func (o *AssetReportGetResponse) GetReport() AssetReport {
	if o == nil {
		var ret AssetReport
		return ret
	}

	return o.Report
}

// GetReportOk returns a tuple with the Report field value
// and a boolean to check if the value has been set.
func (o *AssetReportGetResponse) GetReportOk() (*AssetReport, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Report, true
}

// SetReport sets field value
func (o *AssetReportGetResponse) SetReport(v AssetReport) {
	o.Report = v
}

// GetWarnings returns the Warnings field value
func (o *AssetReportGetResponse) GetWarnings() []Warning {
	if o == nil {
		var ret []Warning
		return ret
	}

	return o.Warnings
}

// GetWarningsOk returns a tuple with the Warnings field value
// and a boolean to check if the value has been set.
func (o *AssetReportGetResponse) GetWarningsOk() (*[]Warning, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Warnings, true
}

// SetWarnings sets field value
func (o *AssetReportGetResponse) SetWarnings(v []Warning) {
	o.Warnings = v
}

// GetRequestId returns the RequestId field value
func (o *AssetReportGetResponse) GetRequestId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RequestId
}

// GetRequestIdOk returns a tuple with the RequestId field value
// and a boolean to check if the value has been set.
func (o *AssetReportGetResponse) GetRequestIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.RequestId, true
}

// SetRequestId sets field value
func (o *AssetReportGetResponse) SetRequestId(v string) {
	o.RequestId = v
}

func (o AssetReportGetResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["report"] = o.Report
	}
	if true {
		toSerialize["warnings"] = o.Warnings
	}
	if true {
		toSerialize["request_id"] = o.RequestId
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *AssetReportGetResponse) UnmarshalJSON(bytes []byte) (err error) {
	varAssetReportGetResponse := _AssetReportGetResponse{}

	if err = json.Unmarshal(bytes, &varAssetReportGetResponse); err == nil {
		*o = AssetReportGetResponse(varAssetReportGetResponse)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "report")
		delete(additionalProperties, "warnings")
		delete(additionalProperties, "request_id")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableAssetReportGetResponse struct {
	value *AssetReportGetResponse
	isSet bool
}

func (v NullableAssetReportGetResponse) Get() *AssetReportGetResponse {
	return v.value
}

func (v *NullableAssetReportGetResponse) Set(val *AssetReportGetResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableAssetReportGetResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableAssetReportGetResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAssetReportGetResponse(val *AssetReportGetResponse) *NullableAssetReportGetResponse {
	return &NullableAssetReportGetResponse{value: val, isSet: true}
}

func (v NullableAssetReportGetResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAssetReportGetResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


