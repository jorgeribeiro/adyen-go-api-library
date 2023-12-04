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

// checks if the MasterpassDetails type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &MasterpassDetails{}

// MasterpassDetails struct for MasterpassDetails
type MasterpassDetails struct {
	// The checkout attempt identifier.
	CheckoutAttemptId *string `json:"checkoutAttemptId,omitempty"`
	// The funding source that should be used when multiple sources are available. For Brazilian combo cards, by default the funding source is credit. To use debit, set this value to **debit**.
	FundingSource *string `json:"fundingSource,omitempty"`
	// The Masterpass transaction ID.
	MasterpassTransactionId string `json:"masterpassTransactionId"`
	// **masterpass**
	Type *string `json:"type,omitempty"`
}

// NewMasterpassDetails instantiates a new MasterpassDetails object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMasterpassDetails(masterpassTransactionId string) *MasterpassDetails {
	this := MasterpassDetails{}
	this.MasterpassTransactionId = masterpassTransactionId
	var type_ string = "masterpass"
	this.Type = &type_
	return &this
}

// NewMasterpassDetailsWithDefaults instantiates a new MasterpassDetails object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMasterpassDetailsWithDefaults() *MasterpassDetails {
	this := MasterpassDetails{}
	var type_ string = "masterpass"
	this.Type = &type_
	return &this
}

// GetCheckoutAttemptId returns the CheckoutAttemptId field value if set, zero value otherwise.
func (o *MasterpassDetails) GetCheckoutAttemptId() string {
	if o == nil || common.IsNil(o.CheckoutAttemptId) {
		var ret string
		return ret
	}
	return *o.CheckoutAttemptId
}

// GetCheckoutAttemptIdOk returns a tuple with the CheckoutAttemptId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MasterpassDetails) GetCheckoutAttemptIdOk() (*string, bool) {
	if o == nil || common.IsNil(o.CheckoutAttemptId) {
		return nil, false
	}
	return o.CheckoutAttemptId, true
}

// HasCheckoutAttemptId returns a boolean if a field has been set.
func (o *MasterpassDetails) HasCheckoutAttemptId() bool {
	if o != nil && !common.IsNil(o.CheckoutAttemptId) {
		return true
	}

	return false
}

// SetCheckoutAttemptId gets a reference to the given string and assigns it to the CheckoutAttemptId field.
func (o *MasterpassDetails) SetCheckoutAttemptId(v string) {
	o.CheckoutAttemptId = &v
}

// GetFundingSource returns the FundingSource field value if set, zero value otherwise.
func (o *MasterpassDetails) GetFundingSource() string {
	if o == nil || common.IsNil(o.FundingSource) {
		var ret string
		return ret
	}
	return *o.FundingSource
}

// GetFundingSourceOk returns a tuple with the FundingSource field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MasterpassDetails) GetFundingSourceOk() (*string, bool) {
	if o == nil || common.IsNil(o.FundingSource) {
		return nil, false
	}
	return o.FundingSource, true
}

// HasFundingSource returns a boolean if a field has been set.
func (o *MasterpassDetails) HasFundingSource() bool {
	if o != nil && !common.IsNil(o.FundingSource) {
		return true
	}

	return false
}

// SetFundingSource gets a reference to the given string and assigns it to the FundingSource field.
func (o *MasterpassDetails) SetFundingSource(v string) {
	o.FundingSource = &v
}

// GetMasterpassTransactionId returns the MasterpassTransactionId field value
func (o *MasterpassDetails) GetMasterpassTransactionId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.MasterpassTransactionId
}

// GetMasterpassTransactionIdOk returns a tuple with the MasterpassTransactionId field value
// and a boolean to check if the value has been set.
func (o *MasterpassDetails) GetMasterpassTransactionIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MasterpassTransactionId, true
}

// SetMasterpassTransactionId sets field value
func (o *MasterpassDetails) SetMasterpassTransactionId(v string) {
	o.MasterpassTransactionId = v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *MasterpassDetails) GetType() string {
	if o == nil || common.IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MasterpassDetails) GetTypeOk() (*string, bool) {
	if o == nil || common.IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *MasterpassDetails) HasType() bool {
	if o != nil && !common.IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *MasterpassDetails) SetType(v string) {
	o.Type = &v
}

func (o MasterpassDetails) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MasterpassDetails) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.CheckoutAttemptId) {
		toSerialize["checkoutAttemptId"] = o.CheckoutAttemptId
	}
	if !common.IsNil(o.FundingSource) {
		toSerialize["fundingSource"] = o.FundingSource
	}
	toSerialize["masterpassTransactionId"] = o.MasterpassTransactionId
	if !common.IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	return toSerialize, nil
}

type NullableMasterpassDetails struct {
	value *MasterpassDetails
	isSet bool
}

func (v NullableMasterpassDetails) Get() *MasterpassDetails {
	return v.value
}

func (v *NullableMasterpassDetails) Set(val *MasterpassDetails) {
	v.value = val
	v.isSet = true
}

func (v NullableMasterpassDetails) IsSet() bool {
	return v.isSet
}

func (v *NullableMasterpassDetails) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMasterpassDetails(val *MasterpassDetails) *NullableMasterpassDetails {
	return &NullableMasterpassDetails{value: val, isSet: true}
}

func (v NullableMasterpassDetails) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMasterpassDetails) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

func (o *MasterpassDetails) isValidFundingSource() bool {
	var allowedEnumValues = []string{"credit", "debit"}
	for _, allowed := range allowedEnumValues {
		if o.GetFundingSource() == allowed {
			return true
		}
	}
	return false
}
func (o *MasterpassDetails) isValidType() bool {
	var allowedEnumValues = []string{"masterpass"}
	for _, allowed := range allowedEnumValues {
		if o.GetType() == allowed {
			return true
		}
	}
	return false
}