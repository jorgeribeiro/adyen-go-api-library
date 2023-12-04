/*
Configuration API

API version: 2
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package balanceplatform

import (
	"encoding/json"

	"github.com/adyen/adyen-go-api-library/v8/src/common"
)

// checks if the PaymentInstrumentGroup type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &PaymentInstrumentGroup{}

// PaymentInstrumentGroup struct for PaymentInstrumentGroup
type PaymentInstrumentGroup struct {
	// The unique identifier of the [balance platform](https://docs.adyen.com/api-explorer/#/balanceplatform/latest/get/balancePlatforms/{id}__queryParam_id) to which the payment instrument group belongs.
	BalancePlatform string `json:"balancePlatform"`
	// Your description for the payment instrument group, maximum 300 characters.
	Description *string `json:"description,omitempty"`
	// The unique identifier of the payment instrument group.
	Id *string `json:"id,omitempty"`
	// Properties of the payment instrument group.
	Properties *map[string]string `json:"properties,omitempty"`
	// Your reference for the payment instrument group, maximum 150 characters.
	Reference *string `json:"reference,omitempty"`
	// The tx variant of the payment instrument group.
	TxVariant string `json:"txVariant"`
}

// NewPaymentInstrumentGroup instantiates a new PaymentInstrumentGroup object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPaymentInstrumentGroup(balancePlatform string, txVariant string) *PaymentInstrumentGroup {
	this := PaymentInstrumentGroup{}
	this.BalancePlatform = balancePlatform
	this.TxVariant = txVariant
	return &this
}

// NewPaymentInstrumentGroupWithDefaults instantiates a new PaymentInstrumentGroup object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPaymentInstrumentGroupWithDefaults() *PaymentInstrumentGroup {
	this := PaymentInstrumentGroup{}
	return &this
}

// GetBalancePlatform returns the BalancePlatform field value
func (o *PaymentInstrumentGroup) GetBalancePlatform() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.BalancePlatform
}

// GetBalancePlatformOk returns a tuple with the BalancePlatform field value
// and a boolean to check if the value has been set.
func (o *PaymentInstrumentGroup) GetBalancePlatformOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BalancePlatform, true
}

// SetBalancePlatform sets field value
func (o *PaymentInstrumentGroup) SetBalancePlatform(v string) {
	o.BalancePlatform = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *PaymentInstrumentGroup) GetDescription() string {
	if o == nil || common.IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentInstrumentGroup) GetDescriptionOk() (*string, bool) {
	if o == nil || common.IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *PaymentInstrumentGroup) HasDescription() bool {
	if o != nil && !common.IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *PaymentInstrumentGroup) SetDescription(v string) {
	o.Description = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *PaymentInstrumentGroup) GetId() string {
	if o == nil || common.IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentInstrumentGroup) GetIdOk() (*string, bool) {
	if o == nil || common.IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *PaymentInstrumentGroup) HasId() bool {
	if o != nil && !common.IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *PaymentInstrumentGroup) SetId(v string) {
	o.Id = &v
}

// GetProperties returns the Properties field value if set, zero value otherwise.
func (o *PaymentInstrumentGroup) GetProperties() map[string]string {
	if o == nil || common.IsNil(o.Properties) {
		var ret map[string]string
		return ret
	}
	return *o.Properties
}

// GetPropertiesOk returns a tuple with the Properties field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentInstrumentGroup) GetPropertiesOk() (*map[string]string, bool) {
	if o == nil || common.IsNil(o.Properties) {
		return nil, false
	}
	return o.Properties, true
}

// HasProperties returns a boolean if a field has been set.
func (o *PaymentInstrumentGroup) HasProperties() bool {
	if o != nil && !common.IsNil(o.Properties) {
		return true
	}

	return false
}

// SetProperties gets a reference to the given map[string]string and assigns it to the Properties field.
func (o *PaymentInstrumentGroup) SetProperties(v map[string]string) {
	o.Properties = &v
}

// GetReference returns the Reference field value if set, zero value otherwise.
func (o *PaymentInstrumentGroup) GetReference() string {
	if o == nil || common.IsNil(o.Reference) {
		var ret string
		return ret
	}
	return *o.Reference
}

// GetReferenceOk returns a tuple with the Reference field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentInstrumentGroup) GetReferenceOk() (*string, bool) {
	if o == nil || common.IsNil(o.Reference) {
		return nil, false
	}
	return o.Reference, true
}

// HasReference returns a boolean if a field has been set.
func (o *PaymentInstrumentGroup) HasReference() bool {
	if o != nil && !common.IsNil(o.Reference) {
		return true
	}

	return false
}

// SetReference gets a reference to the given string and assigns it to the Reference field.
func (o *PaymentInstrumentGroup) SetReference(v string) {
	o.Reference = &v
}

// GetTxVariant returns the TxVariant field value
func (o *PaymentInstrumentGroup) GetTxVariant() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TxVariant
}

// GetTxVariantOk returns a tuple with the TxVariant field value
// and a boolean to check if the value has been set.
func (o *PaymentInstrumentGroup) GetTxVariantOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TxVariant, true
}

// SetTxVariant sets field value
func (o *PaymentInstrumentGroup) SetTxVariant(v string) {
	o.TxVariant = v
}

func (o PaymentInstrumentGroup) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PaymentInstrumentGroup) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["balancePlatform"] = o.BalancePlatform
	if !common.IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !common.IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !common.IsNil(o.Properties) {
		toSerialize["properties"] = o.Properties
	}
	if !common.IsNil(o.Reference) {
		toSerialize["reference"] = o.Reference
	}
	toSerialize["txVariant"] = o.TxVariant
	return toSerialize, nil
}

type NullablePaymentInstrumentGroup struct {
	value *PaymentInstrumentGroup
	isSet bool
}

func (v NullablePaymentInstrumentGroup) Get() *PaymentInstrumentGroup {
	return v.value
}

func (v *NullablePaymentInstrumentGroup) Set(val *PaymentInstrumentGroup) {
	v.value = val
	v.isSet = true
}

func (v NullablePaymentInstrumentGroup) IsSet() bool {
	return v.isSet
}

func (v *NullablePaymentInstrumentGroup) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePaymentInstrumentGroup(val *PaymentInstrumentGroup) *NullablePaymentInstrumentGroup {
	return &NullablePaymentInstrumentGroup{value: val, isSet: true}
}

func (v NullablePaymentInstrumentGroup) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePaymentInstrumentGroup) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}