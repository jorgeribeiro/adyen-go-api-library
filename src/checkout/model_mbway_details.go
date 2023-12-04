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

// checks if the MbwayDetails type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &MbwayDetails{}

// MbwayDetails struct for MbwayDetails
type MbwayDetails struct {
	// The checkout attempt identifier.
	CheckoutAttemptId *string `json:"checkoutAttemptId,omitempty"`
	//
	ShopperEmail string `json:"shopperEmail"`
	//
	TelephoneNumber string `json:"telephoneNumber"`
	// **mbway**
	Type *string `json:"type,omitempty"`
}

// NewMbwayDetails instantiates a new MbwayDetails object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMbwayDetails(shopperEmail string, telephoneNumber string) *MbwayDetails {
	this := MbwayDetails{}
	this.ShopperEmail = shopperEmail
	this.TelephoneNumber = telephoneNumber
	var type_ string = "mbway"
	this.Type = &type_
	return &this
}

// NewMbwayDetailsWithDefaults instantiates a new MbwayDetails object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMbwayDetailsWithDefaults() *MbwayDetails {
	this := MbwayDetails{}
	var type_ string = "mbway"
	this.Type = &type_
	return &this
}

// GetCheckoutAttemptId returns the CheckoutAttemptId field value if set, zero value otherwise.
func (o *MbwayDetails) GetCheckoutAttemptId() string {
	if o == nil || common.IsNil(o.CheckoutAttemptId) {
		var ret string
		return ret
	}
	return *o.CheckoutAttemptId
}

// GetCheckoutAttemptIdOk returns a tuple with the CheckoutAttemptId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MbwayDetails) GetCheckoutAttemptIdOk() (*string, bool) {
	if o == nil || common.IsNil(o.CheckoutAttemptId) {
		return nil, false
	}
	return o.CheckoutAttemptId, true
}

// HasCheckoutAttemptId returns a boolean if a field has been set.
func (o *MbwayDetails) HasCheckoutAttemptId() bool {
	if o != nil && !common.IsNil(o.CheckoutAttemptId) {
		return true
	}

	return false
}

// SetCheckoutAttemptId gets a reference to the given string and assigns it to the CheckoutAttemptId field.
func (o *MbwayDetails) SetCheckoutAttemptId(v string) {
	o.CheckoutAttemptId = &v
}

// GetShopperEmail returns the ShopperEmail field value
func (o *MbwayDetails) GetShopperEmail() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ShopperEmail
}

// GetShopperEmailOk returns a tuple with the ShopperEmail field value
// and a boolean to check if the value has been set.
func (o *MbwayDetails) GetShopperEmailOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ShopperEmail, true
}

// SetShopperEmail sets field value
func (o *MbwayDetails) SetShopperEmail(v string) {
	o.ShopperEmail = v
}

// GetTelephoneNumber returns the TelephoneNumber field value
func (o *MbwayDetails) GetTelephoneNumber() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TelephoneNumber
}

// GetTelephoneNumberOk returns a tuple with the TelephoneNumber field value
// and a boolean to check if the value has been set.
func (o *MbwayDetails) GetTelephoneNumberOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TelephoneNumber, true
}

// SetTelephoneNumber sets field value
func (o *MbwayDetails) SetTelephoneNumber(v string) {
	o.TelephoneNumber = v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *MbwayDetails) GetType() string {
	if o == nil || common.IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MbwayDetails) GetTypeOk() (*string, bool) {
	if o == nil || common.IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *MbwayDetails) HasType() bool {
	if o != nil && !common.IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *MbwayDetails) SetType(v string) {
	o.Type = &v
}

func (o MbwayDetails) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MbwayDetails) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.CheckoutAttemptId) {
		toSerialize["checkoutAttemptId"] = o.CheckoutAttemptId
	}
	toSerialize["shopperEmail"] = o.ShopperEmail
	toSerialize["telephoneNumber"] = o.TelephoneNumber
	if !common.IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	return toSerialize, nil
}

type NullableMbwayDetails struct {
	value *MbwayDetails
	isSet bool
}

func (v NullableMbwayDetails) Get() *MbwayDetails {
	return v.value
}

func (v *NullableMbwayDetails) Set(val *MbwayDetails) {
	v.value = val
	v.isSet = true
}

func (v NullableMbwayDetails) IsSet() bool {
	return v.isSet
}

func (v *NullableMbwayDetails) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMbwayDetails(val *MbwayDetails) *NullableMbwayDetails {
	return &NullableMbwayDetails{value: val, isSet: true}
}

func (v NullableMbwayDetails) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMbwayDetails) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

func (o *MbwayDetails) isValidType() bool {
	var allowedEnumValues = []string{"mbway"}
	for _, allowed := range allowedEnumValues {
		if o.GetType() == allowed {
			return true
		}
	}
	return false
}