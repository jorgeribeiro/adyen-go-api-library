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

// checks if the MerchantData type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &MerchantData{}

// MerchantData struct for MerchantData
type MerchantData struct {
	// The unique identifier of the merchant's acquirer.
	AcquirerId *string `json:"acquirerId,omitempty"`
	// The merchant category code.
	Mcc *string `json:"mcc,omitempty"`
	// The merchant identifier.
	MerchantId   *string       `json:"merchantId,omitempty"`
	NameLocation *NameLocation `json:"nameLocation,omitempty"`
	// The merchant postal code.
	PostalCode *string `json:"postalCode,omitempty"`
}

// NewMerchantData instantiates a new MerchantData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMerchantData() *MerchantData {
	this := MerchantData{}
	return &this
}

// NewMerchantDataWithDefaults instantiates a new MerchantData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMerchantDataWithDefaults() *MerchantData {
	this := MerchantData{}
	return &this
}

// GetAcquirerId returns the AcquirerId field value if set, zero value otherwise.
func (o *MerchantData) GetAcquirerId() string {
	if o == nil || common.IsNil(o.AcquirerId) {
		var ret string
		return ret
	}
	return *o.AcquirerId
}

// GetAcquirerIdOk returns a tuple with the AcquirerId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MerchantData) GetAcquirerIdOk() (*string, bool) {
	if o == nil || common.IsNil(o.AcquirerId) {
		return nil, false
	}
	return o.AcquirerId, true
}

// HasAcquirerId returns a boolean if a field has been set.
func (o *MerchantData) HasAcquirerId() bool {
	if o != nil && !common.IsNil(o.AcquirerId) {
		return true
	}

	return false
}

// SetAcquirerId gets a reference to the given string and assigns it to the AcquirerId field.
func (o *MerchantData) SetAcquirerId(v string) {
	o.AcquirerId = &v
}

// GetMcc returns the Mcc field value if set, zero value otherwise.
func (o *MerchantData) GetMcc() string {
	if o == nil || common.IsNil(o.Mcc) {
		var ret string
		return ret
	}
	return *o.Mcc
}

// GetMccOk returns a tuple with the Mcc field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MerchantData) GetMccOk() (*string, bool) {
	if o == nil || common.IsNil(o.Mcc) {
		return nil, false
	}
	return o.Mcc, true
}

// HasMcc returns a boolean if a field has been set.
func (o *MerchantData) HasMcc() bool {
	if o != nil && !common.IsNil(o.Mcc) {
		return true
	}

	return false
}

// SetMcc gets a reference to the given string and assigns it to the Mcc field.
func (o *MerchantData) SetMcc(v string) {
	o.Mcc = &v
}

// GetMerchantId returns the MerchantId field value if set, zero value otherwise.
func (o *MerchantData) GetMerchantId() string {
	if o == nil || common.IsNil(o.MerchantId) {
		var ret string
		return ret
	}
	return *o.MerchantId
}

// GetMerchantIdOk returns a tuple with the MerchantId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MerchantData) GetMerchantIdOk() (*string, bool) {
	if o == nil || common.IsNil(o.MerchantId) {
		return nil, false
	}
	return o.MerchantId, true
}

// HasMerchantId returns a boolean if a field has been set.
func (o *MerchantData) HasMerchantId() bool {
	if o != nil && !common.IsNil(o.MerchantId) {
		return true
	}

	return false
}

// SetMerchantId gets a reference to the given string and assigns it to the MerchantId field.
func (o *MerchantData) SetMerchantId(v string) {
	o.MerchantId = &v
}

// GetNameLocation returns the NameLocation field value if set, zero value otherwise.
func (o *MerchantData) GetNameLocation() NameLocation {
	if o == nil || common.IsNil(o.NameLocation) {
		var ret NameLocation
		return ret
	}
	return *o.NameLocation
}

// GetNameLocationOk returns a tuple with the NameLocation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MerchantData) GetNameLocationOk() (*NameLocation, bool) {
	if o == nil || common.IsNil(o.NameLocation) {
		return nil, false
	}
	return o.NameLocation, true
}

// HasNameLocation returns a boolean if a field has been set.
func (o *MerchantData) HasNameLocation() bool {
	if o != nil && !common.IsNil(o.NameLocation) {
		return true
	}

	return false
}

// SetNameLocation gets a reference to the given NameLocation and assigns it to the NameLocation field.
func (o *MerchantData) SetNameLocation(v NameLocation) {
	o.NameLocation = &v
}

// GetPostalCode returns the PostalCode field value if set, zero value otherwise.
func (o *MerchantData) GetPostalCode() string {
	if o == nil || common.IsNil(o.PostalCode) {
		var ret string
		return ret
	}
	return *o.PostalCode
}

// GetPostalCodeOk returns a tuple with the PostalCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MerchantData) GetPostalCodeOk() (*string, bool) {
	if o == nil || common.IsNil(o.PostalCode) {
		return nil, false
	}
	return o.PostalCode, true
}

// HasPostalCode returns a boolean if a field has been set.
func (o *MerchantData) HasPostalCode() bool {
	if o != nil && !common.IsNil(o.PostalCode) {
		return true
	}

	return false
}

// SetPostalCode gets a reference to the given string and assigns it to the PostalCode field.
func (o *MerchantData) SetPostalCode(v string) {
	o.PostalCode = &v
}

func (o MerchantData) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MerchantData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.AcquirerId) {
		toSerialize["acquirerId"] = o.AcquirerId
	}
	if !common.IsNil(o.Mcc) {
		toSerialize["mcc"] = o.Mcc
	}
	if !common.IsNil(o.MerchantId) {
		toSerialize["merchantId"] = o.MerchantId
	}
	if !common.IsNil(o.NameLocation) {
		toSerialize["nameLocation"] = o.NameLocation
	}
	if !common.IsNil(o.PostalCode) {
		toSerialize["postalCode"] = o.PostalCode
	}
	return toSerialize, nil
}

type NullableMerchantData struct {
	value *MerchantData
	isSet bool
}

func (v NullableMerchantData) Get() *MerchantData {
	return v.value
}

func (v *NullableMerchantData) Set(val *MerchantData) {
	v.value = val
	v.isSet = true
}

func (v NullableMerchantData) IsSet() bool {
	return v.isSet
}

func (v *NullableMerchantData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMerchantData(val *MerchantData) *NullableMerchantData {
	return &NullableMerchantData{value: val, isSet: true}
}

func (v NullableMerchantData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMerchantData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
