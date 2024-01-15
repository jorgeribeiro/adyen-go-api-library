/*
Management API

API version: 3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package management

import (
	"encoding/json"

	"github.com/adyen/adyen-go-api-library/v8/src/common"
)

// checks if the TapToPay type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &TapToPay{}

// TapToPay struct for TapToPay
type TapToPay struct {
	// The text shown on the screen during the Tap to Pay transaction.
	MerchantDisplayName *string `json:"merchantDisplayName,omitempty"`
}

// NewTapToPay instantiates a new TapToPay object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTapToPay() *TapToPay {
	this := TapToPay{}
	return &this
}

// NewTapToPayWithDefaults instantiates a new TapToPay object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTapToPayWithDefaults() *TapToPay {
	this := TapToPay{}
	return &this
}

// GetMerchantDisplayName returns the MerchantDisplayName field value if set, zero value otherwise.
func (o *TapToPay) GetMerchantDisplayName() string {
	if o == nil || common.IsNil(o.MerchantDisplayName) {
		var ret string
		return ret
	}
	return *o.MerchantDisplayName
}

// GetMerchantDisplayNameOk returns a tuple with the MerchantDisplayName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TapToPay) GetMerchantDisplayNameOk() (*string, bool) {
	if o == nil || common.IsNil(o.MerchantDisplayName) {
		return nil, false
	}
	return o.MerchantDisplayName, true
}

// HasMerchantDisplayName returns a boolean if a field has been set.
func (o *TapToPay) HasMerchantDisplayName() bool {
	if o != nil && !common.IsNil(o.MerchantDisplayName) {
		return true
	}

	return false
}

// SetMerchantDisplayName gets a reference to the given string and assigns it to the MerchantDisplayName field.
func (o *TapToPay) SetMerchantDisplayName(v string) {
	o.MerchantDisplayName = &v
}

func (o TapToPay) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TapToPay) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.MerchantDisplayName) {
		toSerialize["merchantDisplayName"] = o.MerchantDisplayName
	}
	return toSerialize, nil
}

type NullableTapToPay struct {
	value *TapToPay
	isSet bool
}

func (v NullableTapToPay) Get() *TapToPay {
	return v.value
}

func (v *NullableTapToPay) Set(val *TapToPay) {
	v.value = val
	v.isSet = true
}

func (v NullableTapToPay) IsSet() bool {
	return v.isSet
}

func (v *NullableTapToPay) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTapToPay(val *TapToPay) *NullableTapToPay {
	return &NullableTapToPay{value: val, isSet: true}
}

func (v NullableTapToPay) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTapToPay) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
