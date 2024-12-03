/*
Configuration API

API version: 2
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package balanceplatform

import (
	"encoding/json"
	"time"

	"github.com/adyen/adyen-go-api-library/v16/src/common"
)

// checks if the CardOrderItem type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &CardOrderItem{}

// CardOrderItem struct for CardOrderItem
type CardOrderItem struct {
	// The unique identifier of the balance platform.
	BalancePlatform *string                      `json:"balancePlatform,omitempty"`
	Card            *CardOrderItemDeliveryStatus `json:"card,omitempty"`
	// The unique identifier of the card order item.
	CardOrderItemId *string `json:"cardOrderItemId,omitempty"`
	// The date and time when the event was triggered, in ISO 8601 extended format. For example, **2020-12-18T10:15:30+01:00**.
	CreationDate *time.Time `json:"creationDate,omitempty"`
	// The ID of the resource.
	Id *string `json:"id,omitempty"`
	// The unique identifier of the payment instrument related to the card order item.
	PaymentInstrumentId *string                      `json:"paymentInstrumentId,omitempty"`
	Pin                 *CardOrderItemDeliveryStatus `json:"pin,omitempty"`
	// The shipping method used to deliver the card or the PIN.
	ShippingMethod *string `json:"shippingMethod,omitempty"`
}

// NewCardOrderItem instantiates a new CardOrderItem object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCardOrderItem() *CardOrderItem {
	this := CardOrderItem{}
	return &this
}

// NewCardOrderItemWithDefaults instantiates a new CardOrderItem object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCardOrderItemWithDefaults() *CardOrderItem {
	this := CardOrderItem{}
	return &this
}

// GetBalancePlatform returns the BalancePlatform field value if set, zero value otherwise.
func (o *CardOrderItem) GetBalancePlatform() string {
	if o == nil || common.IsNil(o.BalancePlatform) {
		var ret string
		return ret
	}
	return *o.BalancePlatform
}

// GetBalancePlatformOk returns a tuple with the BalancePlatform field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CardOrderItem) GetBalancePlatformOk() (*string, bool) {
	if o == nil || common.IsNil(o.BalancePlatform) {
		return nil, false
	}
	return o.BalancePlatform, true
}

// HasBalancePlatform returns a boolean if a field has been set.
func (o *CardOrderItem) HasBalancePlatform() bool {
	if o != nil && !common.IsNil(o.BalancePlatform) {
		return true
	}

	return false
}

// SetBalancePlatform gets a reference to the given string and assigns it to the BalancePlatform field.
func (o *CardOrderItem) SetBalancePlatform(v string) {
	o.BalancePlatform = &v
}

// GetCard returns the Card field value if set, zero value otherwise.
func (o *CardOrderItem) GetCard() CardOrderItemDeliveryStatus {
	if o == nil || common.IsNil(o.Card) {
		var ret CardOrderItemDeliveryStatus
		return ret
	}
	return *o.Card
}

// GetCardOk returns a tuple with the Card field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CardOrderItem) GetCardOk() (*CardOrderItemDeliveryStatus, bool) {
	if o == nil || common.IsNil(o.Card) {
		return nil, false
	}
	return o.Card, true
}

// HasCard returns a boolean if a field has been set.
func (o *CardOrderItem) HasCard() bool {
	if o != nil && !common.IsNil(o.Card) {
		return true
	}

	return false
}

// SetCard gets a reference to the given CardOrderItemDeliveryStatus and assigns it to the Card field.
func (o *CardOrderItem) SetCard(v CardOrderItemDeliveryStatus) {
	o.Card = &v
}

// GetCardOrderItemId returns the CardOrderItemId field value if set, zero value otherwise.
func (o *CardOrderItem) GetCardOrderItemId() string {
	if o == nil || common.IsNil(o.CardOrderItemId) {
		var ret string
		return ret
	}
	return *o.CardOrderItemId
}

// GetCardOrderItemIdOk returns a tuple with the CardOrderItemId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CardOrderItem) GetCardOrderItemIdOk() (*string, bool) {
	if o == nil || common.IsNil(o.CardOrderItemId) {
		return nil, false
	}
	return o.CardOrderItemId, true
}

// HasCardOrderItemId returns a boolean if a field has been set.
func (o *CardOrderItem) HasCardOrderItemId() bool {
	if o != nil && !common.IsNil(o.CardOrderItemId) {
		return true
	}

	return false
}

// SetCardOrderItemId gets a reference to the given string and assigns it to the CardOrderItemId field.
func (o *CardOrderItem) SetCardOrderItemId(v string) {
	o.CardOrderItemId = &v
}

// GetCreationDate returns the CreationDate field value if set, zero value otherwise.
func (o *CardOrderItem) GetCreationDate() time.Time {
	if o == nil || common.IsNil(o.CreationDate) {
		var ret time.Time
		return ret
	}
	return *o.CreationDate
}

// GetCreationDateOk returns a tuple with the CreationDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CardOrderItem) GetCreationDateOk() (*time.Time, bool) {
	if o == nil || common.IsNil(o.CreationDate) {
		return nil, false
	}
	return o.CreationDate, true
}

// HasCreationDate returns a boolean if a field has been set.
func (o *CardOrderItem) HasCreationDate() bool {
	if o != nil && !common.IsNil(o.CreationDate) {
		return true
	}

	return false
}

// SetCreationDate gets a reference to the given time.Time and assigns it to the CreationDate field.
func (o *CardOrderItem) SetCreationDate(v time.Time) {
	o.CreationDate = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *CardOrderItem) GetId() string {
	if o == nil || common.IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CardOrderItem) GetIdOk() (*string, bool) {
	if o == nil || common.IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *CardOrderItem) HasId() bool {
	if o != nil && !common.IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *CardOrderItem) SetId(v string) {
	o.Id = &v
}

// GetPaymentInstrumentId returns the PaymentInstrumentId field value if set, zero value otherwise.
func (o *CardOrderItem) GetPaymentInstrumentId() string {
	if o == nil || common.IsNil(o.PaymentInstrumentId) {
		var ret string
		return ret
	}
	return *o.PaymentInstrumentId
}

// GetPaymentInstrumentIdOk returns a tuple with the PaymentInstrumentId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CardOrderItem) GetPaymentInstrumentIdOk() (*string, bool) {
	if o == nil || common.IsNil(o.PaymentInstrumentId) {
		return nil, false
	}
	return o.PaymentInstrumentId, true
}

// HasPaymentInstrumentId returns a boolean if a field has been set.
func (o *CardOrderItem) HasPaymentInstrumentId() bool {
	if o != nil && !common.IsNil(o.PaymentInstrumentId) {
		return true
	}

	return false
}

// SetPaymentInstrumentId gets a reference to the given string and assigns it to the PaymentInstrumentId field.
func (o *CardOrderItem) SetPaymentInstrumentId(v string) {
	o.PaymentInstrumentId = &v
}

// GetPin returns the Pin field value if set, zero value otherwise.
func (o *CardOrderItem) GetPin() CardOrderItemDeliveryStatus {
	if o == nil || common.IsNil(o.Pin) {
		var ret CardOrderItemDeliveryStatus
		return ret
	}
	return *o.Pin
}

// GetPinOk returns a tuple with the Pin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CardOrderItem) GetPinOk() (*CardOrderItemDeliveryStatus, bool) {
	if o == nil || common.IsNil(o.Pin) {
		return nil, false
	}
	return o.Pin, true
}

// HasPin returns a boolean if a field has been set.
func (o *CardOrderItem) HasPin() bool {
	if o != nil && !common.IsNil(o.Pin) {
		return true
	}

	return false
}

// SetPin gets a reference to the given CardOrderItemDeliveryStatus and assigns it to the Pin field.
func (o *CardOrderItem) SetPin(v CardOrderItemDeliveryStatus) {
	o.Pin = &v
}

// GetShippingMethod returns the ShippingMethod field value if set, zero value otherwise.
func (o *CardOrderItem) GetShippingMethod() string {
	if o == nil || common.IsNil(o.ShippingMethod) {
		var ret string
		return ret
	}
	return *o.ShippingMethod
}

// GetShippingMethodOk returns a tuple with the ShippingMethod field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CardOrderItem) GetShippingMethodOk() (*string, bool) {
	if o == nil || common.IsNil(o.ShippingMethod) {
		return nil, false
	}
	return o.ShippingMethod, true
}

// HasShippingMethod returns a boolean if a field has been set.
func (o *CardOrderItem) HasShippingMethod() bool {
	if o != nil && !common.IsNil(o.ShippingMethod) {
		return true
	}

	return false
}

// SetShippingMethod gets a reference to the given string and assigns it to the ShippingMethod field.
func (o *CardOrderItem) SetShippingMethod(v string) {
	o.ShippingMethod = &v
}

func (o CardOrderItem) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CardOrderItem) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.BalancePlatform) {
		toSerialize["balancePlatform"] = o.BalancePlatform
	}
	if !common.IsNil(o.Card) {
		toSerialize["card"] = o.Card
	}
	if !common.IsNil(o.CardOrderItemId) {
		toSerialize["cardOrderItemId"] = o.CardOrderItemId
	}
	if !common.IsNil(o.CreationDate) {
		toSerialize["creationDate"] = o.CreationDate
	}
	if !common.IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !common.IsNil(o.PaymentInstrumentId) {
		toSerialize["paymentInstrumentId"] = o.PaymentInstrumentId
	}
	if !common.IsNil(o.Pin) {
		toSerialize["pin"] = o.Pin
	}
	if !common.IsNil(o.ShippingMethod) {
		toSerialize["shippingMethod"] = o.ShippingMethod
	}
	return toSerialize, nil
}

type NullableCardOrderItem struct {
	value *CardOrderItem
	isSet bool
}

func (v NullableCardOrderItem) Get() *CardOrderItem {
	return v.value
}

func (v *NullableCardOrderItem) Set(val *CardOrderItem) {
	v.value = val
	v.isSet = true
}

func (v NullableCardOrderItem) IsSet() bool {
	return v.isSet
}

func (v *NullableCardOrderItem) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCardOrderItem(val *CardOrderItem) *NullableCardOrderItem {
	return &NullableCardOrderItem{value: val, isSet: true}
}

func (v NullableCardOrderItem) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCardOrderItem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
