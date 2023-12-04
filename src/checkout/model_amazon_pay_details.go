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

// checks if the AmazonPayDetails type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &AmazonPayDetails{}

// AmazonPayDetails struct for AmazonPayDetails
type AmazonPayDetails struct {
	// This is the `amazonPayToken` that you obtained from the [Get Checkout Session](https://amazon-pay-acquirer-guide.s3-eu-west-1.amazonaws.com/v1/amazon-pay-api-v2/checkout-session.html#get-checkout-session) response. This token is used for API only integration specifically.
	AmazonPayToken *string `json:"amazonPayToken,omitempty"`
	// The checkout attempt identifier.
	CheckoutAttemptId *string `json:"checkoutAttemptId,omitempty"`
	// The `checkoutSessionId` is used to identify the checkout session at the Amazon Pay side. This field is required only for drop-in and components integration, where it replaces the amazonPayToken.
	CheckoutSessionId *string `json:"checkoutSessionId,omitempty"`
	// **amazonpay**
	Type *string `json:"type,omitempty"`
}

// NewAmazonPayDetails instantiates a new AmazonPayDetails object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAmazonPayDetails() *AmazonPayDetails {
	this := AmazonPayDetails{}
	var type_ string = "amazonpay"
	this.Type = &type_
	return &this
}

// NewAmazonPayDetailsWithDefaults instantiates a new AmazonPayDetails object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAmazonPayDetailsWithDefaults() *AmazonPayDetails {
	this := AmazonPayDetails{}
	var type_ string = "amazonpay"
	this.Type = &type_
	return &this
}

// GetAmazonPayToken returns the AmazonPayToken field value if set, zero value otherwise.
func (o *AmazonPayDetails) GetAmazonPayToken() string {
	if o == nil || common.IsNil(o.AmazonPayToken) {
		var ret string
		return ret
	}
	return *o.AmazonPayToken
}

// GetAmazonPayTokenOk returns a tuple with the AmazonPayToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AmazonPayDetails) GetAmazonPayTokenOk() (*string, bool) {
	if o == nil || common.IsNil(o.AmazonPayToken) {
		return nil, false
	}
	return o.AmazonPayToken, true
}

// HasAmazonPayToken returns a boolean if a field has been set.
func (o *AmazonPayDetails) HasAmazonPayToken() bool {
	if o != nil && !common.IsNil(o.AmazonPayToken) {
		return true
	}

	return false
}

// SetAmazonPayToken gets a reference to the given string and assigns it to the AmazonPayToken field.
func (o *AmazonPayDetails) SetAmazonPayToken(v string) {
	o.AmazonPayToken = &v
}

// GetCheckoutAttemptId returns the CheckoutAttemptId field value if set, zero value otherwise.
func (o *AmazonPayDetails) GetCheckoutAttemptId() string {
	if o == nil || common.IsNil(o.CheckoutAttemptId) {
		var ret string
		return ret
	}
	return *o.CheckoutAttemptId
}

// GetCheckoutAttemptIdOk returns a tuple with the CheckoutAttemptId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AmazonPayDetails) GetCheckoutAttemptIdOk() (*string, bool) {
	if o == nil || common.IsNil(o.CheckoutAttemptId) {
		return nil, false
	}
	return o.CheckoutAttemptId, true
}

// HasCheckoutAttemptId returns a boolean if a field has been set.
func (o *AmazonPayDetails) HasCheckoutAttemptId() bool {
	if o != nil && !common.IsNil(o.CheckoutAttemptId) {
		return true
	}

	return false
}

// SetCheckoutAttemptId gets a reference to the given string and assigns it to the CheckoutAttemptId field.
func (o *AmazonPayDetails) SetCheckoutAttemptId(v string) {
	o.CheckoutAttemptId = &v
}

// GetCheckoutSessionId returns the CheckoutSessionId field value if set, zero value otherwise.
func (o *AmazonPayDetails) GetCheckoutSessionId() string {
	if o == nil || common.IsNil(o.CheckoutSessionId) {
		var ret string
		return ret
	}
	return *o.CheckoutSessionId
}

// GetCheckoutSessionIdOk returns a tuple with the CheckoutSessionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AmazonPayDetails) GetCheckoutSessionIdOk() (*string, bool) {
	if o == nil || common.IsNil(o.CheckoutSessionId) {
		return nil, false
	}
	return o.CheckoutSessionId, true
}

// HasCheckoutSessionId returns a boolean if a field has been set.
func (o *AmazonPayDetails) HasCheckoutSessionId() bool {
	if o != nil && !common.IsNil(o.CheckoutSessionId) {
		return true
	}

	return false
}

// SetCheckoutSessionId gets a reference to the given string and assigns it to the CheckoutSessionId field.
func (o *AmazonPayDetails) SetCheckoutSessionId(v string) {
	o.CheckoutSessionId = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *AmazonPayDetails) GetType() string {
	if o == nil || common.IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AmazonPayDetails) GetTypeOk() (*string, bool) {
	if o == nil || common.IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *AmazonPayDetails) HasType() bool {
	if o != nil && !common.IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *AmazonPayDetails) SetType(v string) {
	o.Type = &v
}

func (o AmazonPayDetails) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AmazonPayDetails) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.AmazonPayToken) {
		toSerialize["amazonPayToken"] = o.AmazonPayToken
	}
	if !common.IsNil(o.CheckoutAttemptId) {
		toSerialize["checkoutAttemptId"] = o.CheckoutAttemptId
	}
	if !common.IsNil(o.CheckoutSessionId) {
		toSerialize["checkoutSessionId"] = o.CheckoutSessionId
	}
	if !common.IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	return toSerialize, nil
}

type NullableAmazonPayDetails struct {
	value *AmazonPayDetails
	isSet bool
}

func (v NullableAmazonPayDetails) Get() *AmazonPayDetails {
	return v.value
}

func (v *NullableAmazonPayDetails) Set(val *AmazonPayDetails) {
	v.value = val
	v.isSet = true
}

func (v NullableAmazonPayDetails) IsSet() bool {
	return v.isSet
}

func (v *NullableAmazonPayDetails) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAmazonPayDetails(val *AmazonPayDetails) *NullableAmazonPayDetails {
	return &NullableAmazonPayDetails{value: val, isSet: true}
}

func (v NullableAmazonPayDetails) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAmazonPayDetails) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

func (o *AmazonPayDetails) isValidType() bool {
	var allowedEnumValues = []string{"amazonpay"}
	for _, allowed := range allowedEnumValues {
		if o.GetType() == allowed {
			return true
		}
	}
	return false
}