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

// checks if the UpdateSplitConfigurationRuleRequest type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &UpdateSplitConfigurationRuleRequest{}

// UpdateSplitConfigurationRuleRequest struct for UpdateSplitConfigurationRuleRequest
type UpdateSplitConfigurationRuleRequest struct {
	// The currency condition that defines whether the split logic applies. Its value must be a three-character [ISO currency code](https://en.wikipedia.org/wiki/ISO_4217).
	Currency string `json:"currency"`
	// The funding source condition of the payment method (only for cards).  Possible values: **credit**, **debit**, or **ANY**.
	FundingSource *string `json:"fundingSource,omitempty"`
	// The payment method condition that defines whether the split logic applies.  Possible values: * [Payment method variant](https://docs.adyen.com/development-resources/paymentmethodvariant): Apply the split logic for a specific payment method. * **ANY**: Apply the split logic for all available payment methods.
	PaymentMethod string `json:"paymentMethod"`
	// The sales channel condition that defines whether the split logic applies.  Possible values: * **Ecommerce**: Online transactions where the cardholder is present. * **ContAuth**: Card on file and/or subscription transactions, where the cardholder is known to the merchant (returning customer). * **Moto**: Mail-order and telephone-order transactions where the customer is in contact with the merchant via email or telephone. * **POS**: Point-of-sale transactions where the customer is physically present to make a payment using a secure payment terminal. * **ANY**: All sales channels.
	ShopperInteraction string `json:"shopperInteraction"`
}

// NewUpdateSplitConfigurationRuleRequest instantiates a new UpdateSplitConfigurationRuleRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateSplitConfigurationRuleRequest(currency string, paymentMethod string, shopperInteraction string) *UpdateSplitConfigurationRuleRequest {
	this := UpdateSplitConfigurationRuleRequest{}
	this.Currency = currency
	this.PaymentMethod = paymentMethod
	this.ShopperInteraction = shopperInteraction
	return &this
}

// NewUpdateSplitConfigurationRuleRequestWithDefaults instantiates a new UpdateSplitConfigurationRuleRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateSplitConfigurationRuleRequestWithDefaults() *UpdateSplitConfigurationRuleRequest {
	this := UpdateSplitConfigurationRuleRequest{}
	return &this
}

// GetCurrency returns the Currency field value
func (o *UpdateSplitConfigurationRuleRequest) GetCurrency() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Currency
}

// GetCurrencyOk returns a tuple with the Currency field value
// and a boolean to check if the value has been set.
func (o *UpdateSplitConfigurationRuleRequest) GetCurrencyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Currency, true
}

// SetCurrency sets field value
func (o *UpdateSplitConfigurationRuleRequest) SetCurrency(v string) {
	o.Currency = v
}

// GetFundingSource returns the FundingSource field value if set, zero value otherwise.
func (o *UpdateSplitConfigurationRuleRequest) GetFundingSource() string {
	if o == nil || common.IsNil(o.FundingSource) {
		var ret string
		return ret
	}
	return *o.FundingSource
}

// GetFundingSourceOk returns a tuple with the FundingSource field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateSplitConfigurationRuleRequest) GetFundingSourceOk() (*string, bool) {
	if o == nil || common.IsNil(o.FundingSource) {
		return nil, false
	}
	return o.FundingSource, true
}

// HasFundingSource returns a boolean if a field has been set.
func (o *UpdateSplitConfigurationRuleRequest) HasFundingSource() bool {
	if o != nil && !common.IsNil(o.FundingSource) {
		return true
	}

	return false
}

// SetFundingSource gets a reference to the given string and assigns it to the FundingSource field.
func (o *UpdateSplitConfigurationRuleRequest) SetFundingSource(v string) {
	o.FundingSource = &v
}

// GetPaymentMethod returns the PaymentMethod field value
func (o *UpdateSplitConfigurationRuleRequest) GetPaymentMethod() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PaymentMethod
}

// GetPaymentMethodOk returns a tuple with the PaymentMethod field value
// and a boolean to check if the value has been set.
func (o *UpdateSplitConfigurationRuleRequest) GetPaymentMethodOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PaymentMethod, true
}

// SetPaymentMethod sets field value
func (o *UpdateSplitConfigurationRuleRequest) SetPaymentMethod(v string) {
	o.PaymentMethod = v
}

// GetShopperInteraction returns the ShopperInteraction field value
func (o *UpdateSplitConfigurationRuleRequest) GetShopperInteraction() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ShopperInteraction
}

// GetShopperInteractionOk returns a tuple with the ShopperInteraction field value
// and a boolean to check if the value has been set.
func (o *UpdateSplitConfigurationRuleRequest) GetShopperInteractionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ShopperInteraction, true
}

// SetShopperInteraction sets field value
func (o *UpdateSplitConfigurationRuleRequest) SetShopperInteraction(v string) {
	o.ShopperInteraction = v
}

func (o UpdateSplitConfigurationRuleRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateSplitConfigurationRuleRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["currency"] = o.Currency
	if !common.IsNil(o.FundingSource) {
		toSerialize["fundingSource"] = o.FundingSource
	}
	toSerialize["paymentMethod"] = o.PaymentMethod
	toSerialize["shopperInteraction"] = o.ShopperInteraction
	return toSerialize, nil
}

type NullableUpdateSplitConfigurationRuleRequest struct {
	value *UpdateSplitConfigurationRuleRequest
	isSet bool
}

func (v NullableUpdateSplitConfigurationRuleRequest) Get() *UpdateSplitConfigurationRuleRequest {
	return v.value
}

func (v *NullableUpdateSplitConfigurationRuleRequest) Set(val *UpdateSplitConfigurationRuleRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateSplitConfigurationRuleRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateSplitConfigurationRuleRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateSplitConfigurationRuleRequest(val *UpdateSplitConfigurationRuleRequest) *NullableUpdateSplitConfigurationRuleRequest {
	return &NullableUpdateSplitConfigurationRuleRequest{value: val, isSet: true}
}

func (v NullableUpdateSplitConfigurationRuleRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateSplitConfigurationRuleRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
