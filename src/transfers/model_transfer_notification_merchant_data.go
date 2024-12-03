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

// checks if the TransferNotificationMerchantData type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &TransferNotificationMerchantData{}

// TransferNotificationMerchantData struct for TransferNotificationMerchantData
type TransferNotificationMerchantData struct {
	// The unique identifier of the merchant's acquirer.
	AcquirerId *string `json:"acquirerId,omitempty"`
	// The city where the merchant is located.
	City *string `json:"city,omitempty"`
	// The country where the merchant is located.
	Country *string `json:"country,omitempty"`
	// The merchant category code.
	Mcc *string `json:"mcc,omitempty"`
	// The merchant identifier.
	MerchantId *string `json:"merchantId,omitempty"`
	// The name of the merchant's shop or service.
	Name *string `json:"name,omitempty"`
	// The merchant postal code.
	PostalCode *string `json:"postalCode,omitempty"`
}

// NewTransferNotificationMerchantData instantiates a new TransferNotificationMerchantData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTransferNotificationMerchantData() *TransferNotificationMerchantData {
	this := TransferNotificationMerchantData{}
	return &this
}

// NewTransferNotificationMerchantDataWithDefaults instantiates a new TransferNotificationMerchantData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTransferNotificationMerchantDataWithDefaults() *TransferNotificationMerchantData {
	this := TransferNotificationMerchantData{}
	return &this
}

// GetAcquirerId returns the AcquirerId field value if set, zero value otherwise.
func (o *TransferNotificationMerchantData) GetAcquirerId() string {
	if o == nil || common.IsNil(o.AcquirerId) {
		var ret string
		return ret
	}
	return *o.AcquirerId
}

// GetAcquirerIdOk returns a tuple with the AcquirerId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransferNotificationMerchantData) GetAcquirerIdOk() (*string, bool) {
	if o == nil || common.IsNil(o.AcquirerId) {
		return nil, false
	}
	return o.AcquirerId, true
}

// HasAcquirerId returns a boolean if a field has been set.
func (o *TransferNotificationMerchantData) HasAcquirerId() bool {
	if o != nil && !common.IsNil(o.AcquirerId) {
		return true
	}

	return false
}

// SetAcquirerId gets a reference to the given string and assigns it to the AcquirerId field.
func (o *TransferNotificationMerchantData) SetAcquirerId(v string) {
	o.AcquirerId = &v
}

// GetCity returns the City field value if set, zero value otherwise.
func (o *TransferNotificationMerchantData) GetCity() string {
	if o == nil || common.IsNil(o.City) {
		var ret string
		return ret
	}
	return *o.City
}

// GetCityOk returns a tuple with the City field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransferNotificationMerchantData) GetCityOk() (*string, bool) {
	if o == nil || common.IsNil(o.City) {
		return nil, false
	}
	return o.City, true
}

// HasCity returns a boolean if a field has been set.
func (o *TransferNotificationMerchantData) HasCity() bool {
	if o != nil && !common.IsNil(o.City) {
		return true
	}

	return false
}

// SetCity gets a reference to the given string and assigns it to the City field.
func (o *TransferNotificationMerchantData) SetCity(v string) {
	o.City = &v
}

// GetCountry returns the Country field value if set, zero value otherwise.
func (o *TransferNotificationMerchantData) GetCountry() string {
	if o == nil || common.IsNil(o.Country) {
		var ret string
		return ret
	}
	return *o.Country
}

// GetCountryOk returns a tuple with the Country field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransferNotificationMerchantData) GetCountryOk() (*string, bool) {
	if o == nil || common.IsNil(o.Country) {
		return nil, false
	}
	return o.Country, true
}

// HasCountry returns a boolean if a field has been set.
func (o *TransferNotificationMerchantData) HasCountry() bool {
	if o != nil && !common.IsNil(o.Country) {
		return true
	}

	return false
}

// SetCountry gets a reference to the given string and assigns it to the Country field.
func (o *TransferNotificationMerchantData) SetCountry(v string) {
	o.Country = &v
}

// GetMcc returns the Mcc field value if set, zero value otherwise.
func (o *TransferNotificationMerchantData) GetMcc() string {
	if o == nil || common.IsNil(o.Mcc) {
		var ret string
		return ret
	}
	return *o.Mcc
}

// GetMccOk returns a tuple with the Mcc field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransferNotificationMerchantData) GetMccOk() (*string, bool) {
	if o == nil || common.IsNil(o.Mcc) {
		return nil, false
	}
	return o.Mcc, true
}

// HasMcc returns a boolean if a field has been set.
func (o *TransferNotificationMerchantData) HasMcc() bool {
	if o != nil && !common.IsNil(o.Mcc) {
		return true
	}

	return false
}

// SetMcc gets a reference to the given string and assigns it to the Mcc field.
func (o *TransferNotificationMerchantData) SetMcc(v string) {
	o.Mcc = &v
}

// GetMerchantId returns the MerchantId field value if set, zero value otherwise.
func (o *TransferNotificationMerchantData) GetMerchantId() string {
	if o == nil || common.IsNil(o.MerchantId) {
		var ret string
		return ret
	}
	return *o.MerchantId
}

// GetMerchantIdOk returns a tuple with the MerchantId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransferNotificationMerchantData) GetMerchantIdOk() (*string, bool) {
	if o == nil || common.IsNil(o.MerchantId) {
		return nil, false
	}
	return o.MerchantId, true
}

// HasMerchantId returns a boolean if a field has been set.
func (o *TransferNotificationMerchantData) HasMerchantId() bool {
	if o != nil && !common.IsNil(o.MerchantId) {
		return true
	}

	return false
}

// SetMerchantId gets a reference to the given string and assigns it to the MerchantId field.
func (o *TransferNotificationMerchantData) SetMerchantId(v string) {
	o.MerchantId = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *TransferNotificationMerchantData) GetName() string {
	if o == nil || common.IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransferNotificationMerchantData) GetNameOk() (*string, bool) {
	if o == nil || common.IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *TransferNotificationMerchantData) HasName() bool {
	if o != nil && !common.IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *TransferNotificationMerchantData) SetName(v string) {
	o.Name = &v
}

// GetPostalCode returns the PostalCode field value if set, zero value otherwise.
func (o *TransferNotificationMerchantData) GetPostalCode() string {
	if o == nil || common.IsNil(o.PostalCode) {
		var ret string
		return ret
	}
	return *o.PostalCode
}

// GetPostalCodeOk returns a tuple with the PostalCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransferNotificationMerchantData) GetPostalCodeOk() (*string, bool) {
	if o == nil || common.IsNil(o.PostalCode) {
		return nil, false
	}
	return o.PostalCode, true
}

// HasPostalCode returns a boolean if a field has been set.
func (o *TransferNotificationMerchantData) HasPostalCode() bool {
	if o != nil && !common.IsNil(o.PostalCode) {
		return true
	}

	return false
}

// SetPostalCode gets a reference to the given string and assigns it to the PostalCode field.
func (o *TransferNotificationMerchantData) SetPostalCode(v string) {
	o.PostalCode = &v
}

func (o TransferNotificationMerchantData) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TransferNotificationMerchantData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.AcquirerId) {
		toSerialize["acquirerId"] = o.AcquirerId
	}
	if !common.IsNil(o.City) {
		toSerialize["city"] = o.City
	}
	if !common.IsNil(o.Country) {
		toSerialize["country"] = o.Country
	}
	if !common.IsNil(o.Mcc) {
		toSerialize["mcc"] = o.Mcc
	}
	if !common.IsNil(o.MerchantId) {
		toSerialize["merchantId"] = o.MerchantId
	}
	if !common.IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !common.IsNil(o.PostalCode) {
		toSerialize["postalCode"] = o.PostalCode
	}
	return toSerialize, nil
}

type NullableTransferNotificationMerchantData struct {
	value *TransferNotificationMerchantData
	isSet bool
}

func (v NullableTransferNotificationMerchantData) Get() *TransferNotificationMerchantData {
	return v.value
}

func (v *NullableTransferNotificationMerchantData) Set(val *TransferNotificationMerchantData) {
	v.value = val
	v.isSet = true
}

func (v NullableTransferNotificationMerchantData) IsSet() bool {
	return v.isSet
}

func (v *NullableTransferNotificationMerchantData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTransferNotificationMerchantData(val *TransferNotificationMerchantData) *NullableTransferNotificationMerchantData {
	return &NullableTransferNotificationMerchantData{value: val, isSet: true}
}

func (v NullableTransferNotificationMerchantData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTransferNotificationMerchantData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
