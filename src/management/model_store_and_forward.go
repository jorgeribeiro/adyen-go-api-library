/*
Management API

API version: 3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package management

import (
	"encoding/json"

	"github.com/adyen/adyen-go-api-library/v9/src/common"
)

// checks if the StoreAndForward type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &StoreAndForward{}

// StoreAndForward struct for StoreAndForward
type StoreAndForward struct {
	// The maximum amount that the terminal accepts for a single store-and-forward payment.
	MaxAmount []MinorUnitsMonetaryValue `json:"maxAmount,omitempty"`
	// The maximum number of store-and-forward transactions per terminal that you can process while offline.
	MaxPayments        *int32              `json:"maxPayments,omitempty"`
	SupportedCardTypes *SupportedCardTypes `json:"supportedCardTypes,omitempty"`
}

// NewStoreAndForward instantiates a new StoreAndForward object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStoreAndForward() *StoreAndForward {
	this := StoreAndForward{}
	return &this
}

// NewStoreAndForwardWithDefaults instantiates a new StoreAndForward object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStoreAndForwardWithDefaults() *StoreAndForward {
	this := StoreAndForward{}
	return &this
}

// GetMaxAmount returns the MaxAmount field value if set, zero value otherwise.
func (o *StoreAndForward) GetMaxAmount() []MinorUnitsMonetaryValue {
	if o == nil || common.IsNil(o.MaxAmount) {
		var ret []MinorUnitsMonetaryValue
		return ret
	}
	return o.MaxAmount
}

// GetMaxAmountOk returns a tuple with the MaxAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StoreAndForward) GetMaxAmountOk() ([]MinorUnitsMonetaryValue, bool) {
	if o == nil || common.IsNil(o.MaxAmount) {
		return nil, false
	}
	return o.MaxAmount, true
}

// HasMaxAmount returns a boolean if a field has been set.
func (o *StoreAndForward) HasMaxAmount() bool {
	if o != nil && !common.IsNil(o.MaxAmount) {
		return true
	}

	return false
}

// SetMaxAmount gets a reference to the given []MinorUnitsMonetaryValue and assigns it to the MaxAmount field.
func (o *StoreAndForward) SetMaxAmount(v []MinorUnitsMonetaryValue) {
	o.MaxAmount = v
}

// GetMaxPayments returns the MaxPayments field value if set, zero value otherwise.
func (o *StoreAndForward) GetMaxPayments() int32 {
	if o == nil || common.IsNil(o.MaxPayments) {
		var ret int32
		return ret
	}
	return *o.MaxPayments
}

// GetMaxPaymentsOk returns a tuple with the MaxPayments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StoreAndForward) GetMaxPaymentsOk() (*int32, bool) {
	if o == nil || common.IsNil(o.MaxPayments) {
		return nil, false
	}
	return o.MaxPayments, true
}

// HasMaxPayments returns a boolean if a field has been set.
func (o *StoreAndForward) HasMaxPayments() bool {
	if o != nil && !common.IsNil(o.MaxPayments) {
		return true
	}

	return false
}

// SetMaxPayments gets a reference to the given int32 and assigns it to the MaxPayments field.
func (o *StoreAndForward) SetMaxPayments(v int32) {
	o.MaxPayments = &v
}

// GetSupportedCardTypes returns the SupportedCardTypes field value if set, zero value otherwise.
func (o *StoreAndForward) GetSupportedCardTypes() SupportedCardTypes {
	if o == nil || common.IsNil(o.SupportedCardTypes) {
		var ret SupportedCardTypes
		return ret
	}
	return *o.SupportedCardTypes
}

// GetSupportedCardTypesOk returns a tuple with the SupportedCardTypes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StoreAndForward) GetSupportedCardTypesOk() (*SupportedCardTypes, bool) {
	if o == nil || common.IsNil(o.SupportedCardTypes) {
		return nil, false
	}
	return o.SupportedCardTypes, true
}

// HasSupportedCardTypes returns a boolean if a field has been set.
func (o *StoreAndForward) HasSupportedCardTypes() bool {
	if o != nil && !common.IsNil(o.SupportedCardTypes) {
		return true
	}

	return false
}

// SetSupportedCardTypes gets a reference to the given SupportedCardTypes and assigns it to the SupportedCardTypes field.
func (o *StoreAndForward) SetSupportedCardTypes(v SupportedCardTypes) {
	o.SupportedCardTypes = &v
}

func (o StoreAndForward) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o StoreAndForward) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.MaxAmount) {
		toSerialize["maxAmount"] = o.MaxAmount
	}
	if !common.IsNil(o.MaxPayments) {
		toSerialize["maxPayments"] = o.MaxPayments
	}
	if !common.IsNil(o.SupportedCardTypes) {
		toSerialize["supportedCardTypes"] = o.SupportedCardTypes
	}
	return toSerialize, nil
}

type NullableStoreAndForward struct {
	value *StoreAndForward
	isSet bool
}

func (v NullableStoreAndForward) Get() *StoreAndForward {
	return v.value
}

func (v *NullableStoreAndForward) Set(val *StoreAndForward) {
	v.value = val
	v.isSet = true
}

func (v NullableStoreAndForward) IsSet() bool {
	return v.isSet
}

func (v *NullableStoreAndForward) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStoreAndForward(val *StoreAndForward) *NullableStoreAndForward {
	return &NullableStoreAndForward{value: val, isSet: true}
}

func (v NullableStoreAndForward) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStoreAndForward) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}