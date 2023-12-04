/*
Adyen Checkout API

API version: 71
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package checkout

import (
	"encoding/json"

	"github.com/adyen/adyen-go-api-library/v8/src/common"
)

// checks if the MobilePayDetails type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &MobilePayDetails{}

// MobilePayDetails struct for MobilePayDetails
type MobilePayDetails struct {
	// The checkout attempt identifier.
	CheckoutAttemptId *string `json:"checkoutAttemptId,omitempty"`
	// **mobilepay**
	Type *string `json:"type,omitempty"`
}

// NewMobilePayDetails instantiates a new MobilePayDetails object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMobilePayDetails() *MobilePayDetails {
	this := MobilePayDetails{}
	var type_ string = "mobilepay"
	this.Type = &type_
	return &this
}

// NewMobilePayDetailsWithDefaults instantiates a new MobilePayDetails object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMobilePayDetailsWithDefaults() *MobilePayDetails {
	this := MobilePayDetails{}
	var type_ string = "mobilepay"
	this.Type = &type_
	return &this
}

// GetCheckoutAttemptId returns the CheckoutAttemptId field value if set, zero value otherwise.
func (o *MobilePayDetails) GetCheckoutAttemptId() string {
	if o == nil || common.IsNil(o.CheckoutAttemptId) {
		var ret string
		return ret
	}
	return *o.CheckoutAttemptId
}

// GetCheckoutAttemptIdOk returns a tuple with the CheckoutAttemptId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MobilePayDetails) GetCheckoutAttemptIdOk() (*string, bool) {
	if o == nil || common.IsNil(o.CheckoutAttemptId) {
		return nil, false
	}
	return o.CheckoutAttemptId, true
}

// HasCheckoutAttemptId returns a boolean if a field has been set.
func (o *MobilePayDetails) HasCheckoutAttemptId() bool {
	if o != nil && !common.IsNil(o.CheckoutAttemptId) {
		return true
	}

	return false
}

// SetCheckoutAttemptId gets a reference to the given string and assigns it to the CheckoutAttemptId field.
func (o *MobilePayDetails) SetCheckoutAttemptId(v string) {
	o.CheckoutAttemptId = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *MobilePayDetails) GetType() string {
	if o == nil || common.IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MobilePayDetails) GetTypeOk() (*string, bool) {
	if o == nil || common.IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *MobilePayDetails) HasType() bool {
	if o != nil && !common.IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *MobilePayDetails) SetType(v string) {
	o.Type = &v
}

func (o MobilePayDetails) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MobilePayDetails) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.CheckoutAttemptId) {
		toSerialize["checkoutAttemptId"] = o.CheckoutAttemptId
	}
	if !common.IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	return toSerialize, nil
}

type NullableMobilePayDetails struct {
	value *MobilePayDetails
	isSet bool
}

func (v NullableMobilePayDetails) Get() *MobilePayDetails {
	return v.value
}

func (v *NullableMobilePayDetails) Set(val *MobilePayDetails) {
	v.value = val
	v.isSet = true
}

func (v NullableMobilePayDetails) IsSet() bool {
	return v.isSet
}

func (v *NullableMobilePayDetails) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMobilePayDetails(val *MobilePayDetails) *NullableMobilePayDetails {
	return &NullableMobilePayDetails{value: val, isSet: true}
}

func (v NullableMobilePayDetails) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMobilePayDetails) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

func (o *MobilePayDetails) isValidType() bool {
	var allowedEnumValues = []string{"mobilepay"}
	for _, allowed := range allowedEnumValues {
		if o.GetType() == allowed {
			return true
		}
	}
	return false
}