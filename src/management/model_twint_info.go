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

// checks if the TwintInfo type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &TwintInfo{}

// TwintInfo struct for TwintInfo
type TwintInfo struct {
	// Twint logo. Format: Base64-encoded string.
	Logo string `json:"logo"`
}

// NewTwintInfo instantiates a new TwintInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTwintInfo(logo string) *TwintInfo {
	this := TwintInfo{}
	this.Logo = logo
	return &this
}

// NewTwintInfoWithDefaults instantiates a new TwintInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTwintInfoWithDefaults() *TwintInfo {
	this := TwintInfo{}
	return &this
}

// GetLogo returns the Logo field value
func (o *TwintInfo) GetLogo() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Logo
}

// GetLogoOk returns a tuple with the Logo field value
// and a boolean to check if the value has been set.
func (o *TwintInfo) GetLogoOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Logo, true
}

// SetLogo sets field value
func (o *TwintInfo) SetLogo(v string) {
	o.Logo = v
}

func (o TwintInfo) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TwintInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["logo"] = o.Logo
	return toSerialize, nil
}

type NullableTwintInfo struct {
	value *TwintInfo
	isSet bool
}

func (v NullableTwintInfo) Get() *TwintInfo {
	return v.value
}

func (v *NullableTwintInfo) Set(val *TwintInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableTwintInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableTwintInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTwintInfo(val *TwintInfo) *NullableTwintInfo {
	return &NullableTwintInfo{value: val, isSet: true}
}

func (v NullableTwintInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTwintInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}