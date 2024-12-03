/*
Management API

API version: 3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package management

import (
	"encoding/json"

	"github.com/adyen/adyen-go-api-library/v16/src/common"
)

// checks if the WeChatPayInfo type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &WeChatPayInfo{}

// WeChatPayInfo struct for WeChatPayInfo
type WeChatPayInfo struct {
	// The name of the contact person from merchant support.
	ContactPersonName string `json:"contactPersonName"`
	// The email address of merchant support.
	Email string `json:"email"`
}

// NewWeChatPayInfo instantiates a new WeChatPayInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWeChatPayInfo(contactPersonName string, email string) *WeChatPayInfo {
	this := WeChatPayInfo{}
	this.ContactPersonName = contactPersonName
	this.Email = email
	return &this
}

// NewWeChatPayInfoWithDefaults instantiates a new WeChatPayInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWeChatPayInfoWithDefaults() *WeChatPayInfo {
	this := WeChatPayInfo{}
	return &this
}

// GetContactPersonName returns the ContactPersonName field value
func (o *WeChatPayInfo) GetContactPersonName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ContactPersonName
}

// GetContactPersonNameOk returns a tuple with the ContactPersonName field value
// and a boolean to check if the value has been set.
func (o *WeChatPayInfo) GetContactPersonNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ContactPersonName, true
}

// SetContactPersonName sets field value
func (o *WeChatPayInfo) SetContactPersonName(v string) {
	o.ContactPersonName = v
}

// GetEmail returns the Email field value
func (o *WeChatPayInfo) GetEmail() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Email
}

// GetEmailOk returns a tuple with the Email field value
// and a boolean to check if the value has been set.
func (o *WeChatPayInfo) GetEmailOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Email, true
}

// SetEmail sets field value
func (o *WeChatPayInfo) SetEmail(v string) {
	o.Email = v
}

func (o WeChatPayInfo) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o WeChatPayInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["contactPersonName"] = o.ContactPersonName
	toSerialize["email"] = o.Email
	return toSerialize, nil
}

type NullableWeChatPayInfo struct {
	value *WeChatPayInfo
	isSet bool
}

func (v NullableWeChatPayInfo) Get() *WeChatPayInfo {
	return v.value
}

func (v *NullableWeChatPayInfo) Set(val *WeChatPayInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableWeChatPayInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableWeChatPayInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWeChatPayInfo(val *WeChatPayInfo) *NullableWeChatPayInfo {
	return &NullableWeChatPayInfo{value: val, isSet: true}
}

func (v NullableWeChatPayInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWeChatPayInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
