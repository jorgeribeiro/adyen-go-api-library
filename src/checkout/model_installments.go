/*
Adyen Checkout API

API version: 71
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package checkout

import (
	"encoding/json"

	"github.com/adyen/adyen-go-api-library/v16/src/common"
)

// checks if the Installments type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &Installments{}

// Installments struct for Installments
type Installments struct {
	// Defines the bonus percentage, refund percentage or if the transaction is Buy now Pay later. Used for [card installments in Mexico](https://docs.adyen.com/payment-methods/cards/credit-card-installments/#getting-paid-mexico)
	Extra *int32 `json:"extra,omitempty"`
	// The installment plan, used for [card installments in Japan](https://docs.adyen.com/payment-methods/cards/credit-card-installments#make-a-payment-japan). and [Mexico](https://docs.adyen.com/payment-methods/cards/credit-card-installments/#getting-paid-mexico). By default, this is set to **regular**.
	Plan *string `json:"plan,omitempty"`
	// Defines the number of installments. Usually, the maximum allowed number of installments is capped. For example, it may not be possible to split a payment in more than 24 installments. The acquirer sets this upper limit, so its value may vary. This value can be zero for Installments processed in Mexico.
	Value int32 `json:"value"`
}

// NewInstallments instantiates a new Installments object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInstallments(value int32) *Installments {
	this := Installments{}
	this.Value = value
	return &this
}

// NewInstallmentsWithDefaults instantiates a new Installments object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInstallmentsWithDefaults() *Installments {
	this := Installments{}
	return &this
}

// GetExtra returns the Extra field value if set, zero value otherwise.
func (o *Installments) GetExtra() int32 {
	if o == nil || common.IsNil(o.Extra) {
		var ret int32
		return ret
	}
	return *o.Extra
}

// GetExtraOk returns a tuple with the Extra field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Installments) GetExtraOk() (*int32, bool) {
	if o == nil || common.IsNil(o.Extra) {
		return nil, false
	}
	return o.Extra, true
}

// HasExtra returns a boolean if a field has been set.
func (o *Installments) HasExtra() bool {
	if o != nil && !common.IsNil(o.Extra) {
		return true
	}

	return false
}

// SetExtra gets a reference to the given int32 and assigns it to the Extra field.
func (o *Installments) SetExtra(v int32) {
	o.Extra = &v
}

// GetPlan returns the Plan field value if set, zero value otherwise.
func (o *Installments) GetPlan() string {
	if o == nil || common.IsNil(o.Plan) {
		var ret string
		return ret
	}
	return *o.Plan
}

// GetPlanOk returns a tuple with the Plan field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Installments) GetPlanOk() (*string, bool) {
	if o == nil || common.IsNil(o.Plan) {
		return nil, false
	}
	return o.Plan, true
}

// HasPlan returns a boolean if a field has been set.
func (o *Installments) HasPlan() bool {
	if o != nil && !common.IsNil(o.Plan) {
		return true
	}

	return false
}

// SetPlan gets a reference to the given string and assigns it to the Plan field.
func (o *Installments) SetPlan(v string) {
	o.Plan = &v
}

// GetValue returns the Value field value
func (o *Installments) GetValue() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Value
}

// GetValueOk returns a tuple with the Value field value
// and a boolean to check if the value has been set.
func (o *Installments) GetValueOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Value, true
}

// SetValue sets field value
func (o *Installments) SetValue(v int32) {
	o.Value = v
}

func (o Installments) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Installments) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.Extra) {
		toSerialize["extra"] = o.Extra
	}
	if !common.IsNil(o.Plan) {
		toSerialize["plan"] = o.Plan
	}
	toSerialize["value"] = o.Value
	return toSerialize, nil
}

type NullableInstallments struct {
	value *Installments
	isSet bool
}

func (v NullableInstallments) Get() *Installments {
	return v.value
}

func (v *NullableInstallments) Set(val *Installments) {
	v.value = val
	v.isSet = true
}

func (v NullableInstallments) IsSet() bool {
	return v.isSet
}

func (v *NullableInstallments) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInstallments(val *Installments) *NullableInstallments {
	return &NullableInstallments{value: val, isSet: true}
}

func (v NullableInstallments) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInstallments) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

func (o *Installments) isValidPlan() bool {
	var allowedEnumValues = []string{"buynow_paylater", "interes_refund_prctg", "interest_bonus", "nointeres_refund_prctg", "nointerest_bonus", "refund_prctg", "regular", "revolving", "with_interest"}
	for _, allowed := range allowedEnumValues {
		if o.GetPlan() == allowed {
			return true
		}
	}
	return false
}
