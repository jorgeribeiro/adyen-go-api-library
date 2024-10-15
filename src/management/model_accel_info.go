/*
Management API

API version: 3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package management

import (
	"encoding/json"

	"github.com/adyen/adyen-go-api-library/v12/src/common"
)

// checks if the AccelInfo type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &AccelInfo{}

// AccelInfo struct for AccelInfo
type AccelInfo struct {
	// The type of transactions processed over this payment method.  Allowed values: - **pos** for in-person payments.  - **billpay** for subscription payments, both the initial payment and the later recurring payments. These transactions have `recurringProcessingModel` **Subscription**.  - **ecom** for all other card not present transactions. This includes non-recurring transactions and transactions with `recurringProcessingModel` **CardOnFile** or **UnscheduledCardOnFile**.
	ProcessingType         string                      `json:"processingType"`
	TransactionDescription *TransactionDescriptionInfo `json:"transactionDescription,omitempty"`
}

// NewAccelInfo instantiates a new AccelInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAccelInfo(processingType string) *AccelInfo {
	this := AccelInfo{}
	this.ProcessingType = processingType
	return &this
}

// NewAccelInfoWithDefaults instantiates a new AccelInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAccelInfoWithDefaults() *AccelInfo {
	this := AccelInfo{}
	return &this
}

// GetProcessingType returns the ProcessingType field value
func (o *AccelInfo) GetProcessingType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ProcessingType
}

// GetProcessingTypeOk returns a tuple with the ProcessingType field value
// and a boolean to check if the value has been set.
func (o *AccelInfo) GetProcessingTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ProcessingType, true
}

// SetProcessingType sets field value
func (o *AccelInfo) SetProcessingType(v string) {
	o.ProcessingType = v
}

// GetTransactionDescription returns the TransactionDescription field value if set, zero value otherwise.
func (o *AccelInfo) GetTransactionDescription() TransactionDescriptionInfo {
	if o == nil || common.IsNil(o.TransactionDescription) {
		var ret TransactionDescriptionInfo
		return ret
	}
	return *o.TransactionDescription
}

// GetTransactionDescriptionOk returns a tuple with the TransactionDescription field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccelInfo) GetTransactionDescriptionOk() (*TransactionDescriptionInfo, bool) {
	if o == nil || common.IsNil(o.TransactionDescription) {
		return nil, false
	}
	return o.TransactionDescription, true
}

// HasTransactionDescription returns a boolean if a field has been set.
func (o *AccelInfo) HasTransactionDescription() bool {
	if o != nil && !common.IsNil(o.TransactionDescription) {
		return true
	}

	return false
}

// SetTransactionDescription gets a reference to the given TransactionDescriptionInfo and assigns it to the TransactionDescription field.
func (o *AccelInfo) SetTransactionDescription(v TransactionDescriptionInfo) {
	o.TransactionDescription = &v
}

func (o AccelInfo) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AccelInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["processingType"] = o.ProcessingType
	if !common.IsNil(o.TransactionDescription) {
		toSerialize["transactionDescription"] = o.TransactionDescription
	}
	return toSerialize, nil
}

type NullableAccelInfo struct {
	value *AccelInfo
	isSet bool
}

func (v NullableAccelInfo) Get() *AccelInfo {
	return v.value
}

func (v *NullableAccelInfo) Set(val *AccelInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableAccelInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableAccelInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAccelInfo(val *AccelInfo) *NullableAccelInfo {
	return &NullableAccelInfo{value: val, isSet: true}
}

func (v NullableAccelInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAccelInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

func (o *AccelInfo) isValidProcessingType() bool {
	var allowedEnumValues = []string{"billpay", "ecom", "pos"}
	for _, allowed := range allowedEnumValues {
		if o.GetProcessingType() == allowed {
			return true
		}
	}
	return false
}