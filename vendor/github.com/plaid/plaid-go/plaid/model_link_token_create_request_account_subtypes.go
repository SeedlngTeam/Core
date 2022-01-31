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

// LinkTokenCreateRequestAccountSubtypes By default, Link will only display account types that are compatible with all products supplied in the `products` parameter of `/link/token/create`. You can further limit the accounts shown in Link by using `account_filters` to specify the account subtypes to be shown in Link. Only the specified subtypes will be shown. This filtering applies to both the Account Select view (if enabled) and the Institution Select view. Institutions that do not support the selected subtypes will be omitted from Link. To indicate that all subtypes should be shown, use the value `\"all\"`. If the `account_filters` filter is used, any account type for which a filter is not specified will be entirely omitted from Link.  For a full list of valid types and subtypes, see the [Account schema](https://plaid.com/docs/api/accounts#accounts-schema).  For institutions using OAuth, the filter will not affect the list of institutions or accounts shown by the bank in the OAuth window. 
type LinkTokenCreateRequestAccountSubtypes struct {
	// A filter to apply to `depository`-type accounts
	Depository *map[string]map[string]interface{} `json:"depository,omitempty"`
	// A filter to apply to `credit`-type accounts
	Credit *map[string]map[string]interface{} `json:"credit,omitempty"`
	// A filter to apply to `loan`-type accounts
	Loan *map[string]map[string]interface{} `json:"loan,omitempty"`
	// A filter to apply to `investment`-type accounts
	Investment *map[string]map[string]interface{} `json:"investment,omitempty"`
}

// NewLinkTokenCreateRequestAccountSubtypes instantiates a new LinkTokenCreateRequestAccountSubtypes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLinkTokenCreateRequestAccountSubtypes() *LinkTokenCreateRequestAccountSubtypes {
	this := LinkTokenCreateRequestAccountSubtypes{}
	return &this
}

// NewLinkTokenCreateRequestAccountSubtypesWithDefaults instantiates a new LinkTokenCreateRequestAccountSubtypes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLinkTokenCreateRequestAccountSubtypesWithDefaults() *LinkTokenCreateRequestAccountSubtypes {
	this := LinkTokenCreateRequestAccountSubtypes{}
	return &this
}

// GetDepository returns the Depository field value if set, zero value otherwise.
func (o *LinkTokenCreateRequestAccountSubtypes) GetDepository() map[string]map[string]interface{} {
	if o == nil || o.Depository == nil {
		var ret map[string]map[string]interface{}
		return ret
	}
	return *o.Depository
}

// GetDepositoryOk returns a tuple with the Depository field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LinkTokenCreateRequestAccountSubtypes) GetDepositoryOk() (*map[string]map[string]interface{}, bool) {
	if o == nil || o.Depository == nil {
		return nil, false
	}
	return o.Depository, true
}

// HasDepository returns a boolean if a field has been set.
func (o *LinkTokenCreateRequestAccountSubtypes) HasDepository() bool {
	if o != nil && o.Depository != nil {
		return true
	}

	return false
}

// SetDepository gets a reference to the given map[string]map[string]interface{} and assigns it to the Depository field.
func (o *LinkTokenCreateRequestAccountSubtypes) SetDepository(v map[string]map[string]interface{}) {
	o.Depository = &v
}

// GetCredit returns the Credit field value if set, zero value otherwise.
func (o *LinkTokenCreateRequestAccountSubtypes) GetCredit() map[string]map[string]interface{} {
	if o == nil || o.Credit == nil {
		var ret map[string]map[string]interface{}
		return ret
	}
	return *o.Credit
}

// GetCreditOk returns a tuple with the Credit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LinkTokenCreateRequestAccountSubtypes) GetCreditOk() (*map[string]map[string]interface{}, bool) {
	if o == nil || o.Credit == nil {
		return nil, false
	}
	return o.Credit, true
}

// HasCredit returns a boolean if a field has been set.
func (o *LinkTokenCreateRequestAccountSubtypes) HasCredit() bool {
	if o != nil && o.Credit != nil {
		return true
	}

	return false
}

// SetCredit gets a reference to the given map[string]map[string]interface{} and assigns it to the Credit field.
func (o *LinkTokenCreateRequestAccountSubtypes) SetCredit(v map[string]map[string]interface{}) {
	o.Credit = &v
}

// GetLoan returns the Loan field value if set, zero value otherwise.
func (o *LinkTokenCreateRequestAccountSubtypes) GetLoan() map[string]map[string]interface{} {
	if o == nil || o.Loan == nil {
		var ret map[string]map[string]interface{}
		return ret
	}
	return *o.Loan
}

// GetLoanOk returns a tuple with the Loan field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LinkTokenCreateRequestAccountSubtypes) GetLoanOk() (*map[string]map[string]interface{}, bool) {
	if o == nil || o.Loan == nil {
		return nil, false
	}
	return o.Loan, true
}

// HasLoan returns a boolean if a field has been set.
func (o *LinkTokenCreateRequestAccountSubtypes) HasLoan() bool {
	if o != nil && o.Loan != nil {
		return true
	}

	return false
}

// SetLoan gets a reference to the given map[string]map[string]interface{} and assigns it to the Loan field.
func (o *LinkTokenCreateRequestAccountSubtypes) SetLoan(v map[string]map[string]interface{}) {
	o.Loan = &v
}

// GetInvestment returns the Investment field value if set, zero value otherwise.
func (o *LinkTokenCreateRequestAccountSubtypes) GetInvestment() map[string]map[string]interface{} {
	if o == nil || o.Investment == nil {
		var ret map[string]map[string]interface{}
		return ret
	}
	return *o.Investment
}

// GetInvestmentOk returns a tuple with the Investment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LinkTokenCreateRequestAccountSubtypes) GetInvestmentOk() (*map[string]map[string]interface{}, bool) {
	if o == nil || o.Investment == nil {
		return nil, false
	}
	return o.Investment, true
}

// HasInvestment returns a boolean if a field has been set.
func (o *LinkTokenCreateRequestAccountSubtypes) HasInvestment() bool {
	if o != nil && o.Investment != nil {
		return true
	}

	return false
}

// SetInvestment gets a reference to the given map[string]map[string]interface{} and assigns it to the Investment field.
func (o *LinkTokenCreateRequestAccountSubtypes) SetInvestment(v map[string]map[string]interface{}) {
	o.Investment = &v
}

func (o LinkTokenCreateRequestAccountSubtypes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Depository != nil {
		toSerialize["depository"] = o.Depository
	}
	if o.Credit != nil {
		toSerialize["credit"] = o.Credit
	}
	if o.Loan != nil {
		toSerialize["loan"] = o.Loan
	}
	if o.Investment != nil {
		toSerialize["investment"] = o.Investment
	}
	return json.Marshal(toSerialize)
}

type NullableLinkTokenCreateRequestAccountSubtypes struct {
	value *LinkTokenCreateRequestAccountSubtypes
	isSet bool
}

func (v NullableLinkTokenCreateRequestAccountSubtypes) Get() *LinkTokenCreateRequestAccountSubtypes {
	return v.value
}

func (v *NullableLinkTokenCreateRequestAccountSubtypes) Set(val *LinkTokenCreateRequestAccountSubtypes) {
	v.value = val
	v.isSet = true
}

func (v NullableLinkTokenCreateRequestAccountSubtypes) IsSet() bool {
	return v.isSet
}

func (v *NullableLinkTokenCreateRequestAccountSubtypes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLinkTokenCreateRequestAccountSubtypes(val *LinkTokenCreateRequestAccountSubtypes) *NullableLinkTokenCreateRequestAccountSubtypes {
	return &NullableLinkTokenCreateRequestAccountSubtypes{value: val, isSet: true}
}

func (v NullableLinkTokenCreateRequestAccountSubtypes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLinkTokenCreateRequestAccountSubtypes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


