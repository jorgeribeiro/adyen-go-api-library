/*
Adyen Checkout API

API version: 71
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package checkout

import (
	"encoding/json"

	"github.com/adyen/adyen-go-api-library/v16/src/common"
)

// checks if the Amounts type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &Amounts{}

// Amounts struct for Amounts
type Amounts struct {
	// The three-character [ISO currency code](https://docs.adyen.com/development-resources/currency-codes/).
	Currency string `json:"currency"`
	// The amounts of the donation (in [minor units](https://docs.adyen.com/development-resources/currency-codes/)).
	Values []int64 `json:"values"`
}

// NewAmounts instantiates a new Amounts object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAmounts(currency string, values []int64) *Amounts {
	this := Amounts{}
	this.Currency = currency
	this.Values = values
	return &this
}

// NewAmountsWithDefaults instantiates a new Amounts object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAmountsWithDefaults() *Amounts {
	this := Amounts{}
	return &this
}

// GetCurrency returns the Currency field value
func (o *Amounts) GetCurrency() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Currency
}

// GetCurrencyOk returns a tuple with the Currency field value
// and a boolean to check if the value has been set.
func (o *Amounts) GetCurrencyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Currency, true
}

// SetCurrency sets field value
func (o *Amounts) SetCurrency(v string) {
	o.Currency = v
}

// GetValues returns the Values field value
func (o *Amounts) GetValues() []int64 {
	if o == nil {
		var ret []int64
		return ret
	}

	return o.Values
}

// GetValuesOk returns a tuple with the Values field value
// and a boolean to check if the value has been set.
func (o *Amounts) GetValuesOk() ([]int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.Values, true
}

// SetValues sets field value
func (o *Amounts) SetValues(v []int64) {
	o.Values = v
}

func (o Amounts) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Amounts) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["currency"] = o.Currency
	toSerialize["values"] = o.Values
	return toSerialize, nil
}

type NullableAmounts struct {
	value *Amounts
	isSet bool
}

func (v NullableAmounts) Get() *Amounts {
	return v.value
}

func (v *NullableAmounts) Set(val *Amounts) {
	v.value = val
	v.isSet = true
}

func (v NullableAmounts) IsSet() bool {
	return v.isSet
}

func (v *NullableAmounts) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAmounts(val *Amounts) *NullableAmounts {
	return &NullableAmounts{value: val, isSet: true}
}

func (v NullableAmounts) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAmounts) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
