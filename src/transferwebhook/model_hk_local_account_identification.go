/*
Transfer webhooks

API version: 4
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package transferwebhook

import (
	"encoding/json"

	"github.com/adyen/adyen-go-api-library/v8/src/common"
)

// checks if the HKLocalAccountIdentification type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &HKLocalAccountIdentification{}

// HKLocalAccountIdentification struct for HKLocalAccountIdentification
type HKLocalAccountIdentification struct {
	// The 9- to 15-character bank account number (alphanumeric), without separators or whitespace. Starts with the 3-digit branch code.
	AccountNumber string `json:"accountNumber"`
	// The 3-digit clearing code, without separators or whitespace.
	ClearingCode string `json:"clearingCode"`
	// **hkLocal**
	Type string `json:"type"`
}

// NewHKLocalAccountIdentification instantiates a new HKLocalAccountIdentification object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHKLocalAccountIdentification(accountNumber string, clearingCode string, type_ string) *HKLocalAccountIdentification {
	this := HKLocalAccountIdentification{}
	this.AccountNumber = accountNumber
	this.ClearingCode = clearingCode
	this.Type = type_
	return &this
}

// NewHKLocalAccountIdentificationWithDefaults instantiates a new HKLocalAccountIdentification object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHKLocalAccountIdentificationWithDefaults() *HKLocalAccountIdentification {
	this := HKLocalAccountIdentification{}
	var type_ string = "hkLocal"
	this.Type = type_
	return &this
}

// GetAccountNumber returns the AccountNumber field value
func (o *HKLocalAccountIdentification) GetAccountNumber() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AccountNumber
}

// GetAccountNumberOk returns a tuple with the AccountNumber field value
// and a boolean to check if the value has been set.
func (o *HKLocalAccountIdentification) GetAccountNumberOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AccountNumber, true
}

// SetAccountNumber sets field value
func (o *HKLocalAccountIdentification) SetAccountNumber(v string) {
	o.AccountNumber = v
}

// GetClearingCode returns the ClearingCode field value
func (o *HKLocalAccountIdentification) GetClearingCode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClearingCode
}

// GetClearingCodeOk returns a tuple with the ClearingCode field value
// and a boolean to check if the value has been set.
func (o *HKLocalAccountIdentification) GetClearingCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClearingCode, true
}

// SetClearingCode sets field value
func (o *HKLocalAccountIdentification) SetClearingCode(v string) {
	o.ClearingCode = v
}

// GetType returns the Type field value
func (o *HKLocalAccountIdentification) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *HKLocalAccountIdentification) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *HKLocalAccountIdentification) SetType(v string) {
	o.Type = v
}

func (o HKLocalAccountIdentification) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o HKLocalAccountIdentification) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["accountNumber"] = o.AccountNumber
	toSerialize["clearingCode"] = o.ClearingCode
	toSerialize["type"] = o.Type
	return toSerialize, nil
}

type NullableHKLocalAccountIdentification struct {
	value *HKLocalAccountIdentification
	isSet bool
}

func (v NullableHKLocalAccountIdentification) Get() *HKLocalAccountIdentification {
	return v.value
}

func (v *NullableHKLocalAccountIdentification) Set(val *HKLocalAccountIdentification) {
	v.value = val
	v.isSet = true
}

func (v NullableHKLocalAccountIdentification) IsSet() bool {
	return v.isSet
}

func (v *NullableHKLocalAccountIdentification) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHKLocalAccountIdentification(val *HKLocalAccountIdentification) *NullableHKLocalAccountIdentification {
	return &NullableHKLocalAccountIdentification{value: val, isSet: true}
}

func (v NullableHKLocalAccountIdentification) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHKLocalAccountIdentification) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

func (o *HKLocalAccountIdentification) isValidType() bool {
	var allowedEnumValues = []string{"hkLocal"}
	for _, allowed := range allowedEnumValues {
		if o.GetType() == allowed {
			return true
		}
	}
	return false
}