/*
Transfers API

API version: 4
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package transfers

import (
	"encoding/json"

	"github.com/adyen/adyen-go-api-library/v16/src/common"
)

// checks if the TransactionEventViolation type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &TransactionEventViolation{}

// TransactionEventViolation struct for TransactionEventViolation
type TransactionEventViolation struct {
	// An explanation about why the transaction rule failed.
	Reason                *string                   `json:"reason,omitempty"`
	TransactionRule       *TransactionRuleReference `json:"transactionRule,omitempty"`
	TransactionRuleSource *TransactionRuleSource    `json:"transactionRuleSource,omitempty"`
}

// NewTransactionEventViolation instantiates a new TransactionEventViolation object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTransactionEventViolation() *TransactionEventViolation {
	this := TransactionEventViolation{}
	return &this
}

// NewTransactionEventViolationWithDefaults instantiates a new TransactionEventViolation object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTransactionEventViolationWithDefaults() *TransactionEventViolation {
	this := TransactionEventViolation{}
	return &this
}

// GetReason returns the Reason field value if set, zero value otherwise.
func (o *TransactionEventViolation) GetReason() string {
	if o == nil || common.IsNil(o.Reason) {
		var ret string
		return ret
	}
	return *o.Reason
}

// GetReasonOk returns a tuple with the Reason field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransactionEventViolation) GetReasonOk() (*string, bool) {
	if o == nil || common.IsNil(o.Reason) {
		return nil, false
	}
	return o.Reason, true
}

// HasReason returns a boolean if a field has been set.
func (o *TransactionEventViolation) HasReason() bool {
	if o != nil && !common.IsNil(o.Reason) {
		return true
	}

	return false
}

// SetReason gets a reference to the given string and assigns it to the Reason field.
func (o *TransactionEventViolation) SetReason(v string) {
	o.Reason = &v
}

// GetTransactionRule returns the TransactionRule field value if set, zero value otherwise.
func (o *TransactionEventViolation) GetTransactionRule() TransactionRuleReference {
	if o == nil || common.IsNil(o.TransactionRule) {
		var ret TransactionRuleReference
		return ret
	}
	return *o.TransactionRule
}

// GetTransactionRuleOk returns a tuple with the TransactionRule field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransactionEventViolation) GetTransactionRuleOk() (*TransactionRuleReference, bool) {
	if o == nil || common.IsNil(o.TransactionRule) {
		return nil, false
	}
	return o.TransactionRule, true
}

// HasTransactionRule returns a boolean if a field has been set.
func (o *TransactionEventViolation) HasTransactionRule() bool {
	if o != nil && !common.IsNil(o.TransactionRule) {
		return true
	}

	return false
}

// SetTransactionRule gets a reference to the given TransactionRuleReference and assigns it to the TransactionRule field.
func (o *TransactionEventViolation) SetTransactionRule(v TransactionRuleReference) {
	o.TransactionRule = &v
}

// GetTransactionRuleSource returns the TransactionRuleSource field value if set, zero value otherwise.
func (o *TransactionEventViolation) GetTransactionRuleSource() TransactionRuleSource {
	if o == nil || common.IsNil(o.TransactionRuleSource) {
		var ret TransactionRuleSource
		return ret
	}
	return *o.TransactionRuleSource
}

// GetTransactionRuleSourceOk returns a tuple with the TransactionRuleSource field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransactionEventViolation) GetTransactionRuleSourceOk() (*TransactionRuleSource, bool) {
	if o == nil || common.IsNil(o.TransactionRuleSource) {
		return nil, false
	}
	return o.TransactionRuleSource, true
}

// HasTransactionRuleSource returns a boolean if a field has been set.
func (o *TransactionEventViolation) HasTransactionRuleSource() bool {
	if o != nil && !common.IsNil(o.TransactionRuleSource) {
		return true
	}

	return false
}

// SetTransactionRuleSource gets a reference to the given TransactionRuleSource and assigns it to the TransactionRuleSource field.
func (o *TransactionEventViolation) SetTransactionRuleSource(v TransactionRuleSource) {
	o.TransactionRuleSource = &v
}

func (o TransactionEventViolation) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TransactionEventViolation) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.Reason) {
		toSerialize["reason"] = o.Reason
	}
	if !common.IsNil(o.TransactionRule) {
		toSerialize["transactionRule"] = o.TransactionRule
	}
	if !common.IsNil(o.TransactionRuleSource) {
		toSerialize["transactionRuleSource"] = o.TransactionRuleSource
	}
	return toSerialize, nil
}

type NullableTransactionEventViolation struct {
	value *TransactionEventViolation
	isSet bool
}

func (v NullableTransactionEventViolation) Get() *TransactionEventViolation {
	return v.value
}

func (v *NullableTransactionEventViolation) Set(val *TransactionEventViolation) {
	v.value = val
	v.isSet = true
}

func (v NullableTransactionEventViolation) IsSet() bool {
	return v.isSet
}

func (v *NullableTransactionEventViolation) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTransactionEventViolation(val *TransactionEventViolation) *NullableTransactionEventViolation {
	return &NullableTransactionEventViolation{value: val, isSet: true}
}

func (v NullableTransactionEventViolation) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTransactionEventViolation) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
