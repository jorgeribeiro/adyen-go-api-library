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

// checks if the BillingAddress type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &BillingAddress{}

// BillingAddress struct for BillingAddress
type BillingAddress struct {
	// The name of the city. Maximum length: 3000 characters.
	City string `json:"city"`
	// The two-character ISO-3166-1 alpha-2 country code. For example, **US**. > If you don't know the country or are not collecting the country from the shopper, provide `country` as `ZZ`.
	Country string `json:"country"`
	// The number or name of the house. Maximum length: 3000 characters.
	HouseNumberOrName string `json:"houseNumberOrName"`
	// A maximum of five digits for an address in the US, or a maximum of ten characters for an address in all other countries.
	PostalCode string `json:"postalCode"`
	// The two-character ISO 3166-2 state or province code. For example, **CA** in the US or **ON** in Canada. > Required for the US and Canada.
	StateOrProvince *string `json:"stateOrProvince,omitempty"`
	// The name of the street. Maximum length: 3000 characters. > The house number should not be included in this field; it should be separately provided via `houseNumberOrName`.
	Street string `json:"street"`
}

// NewBillingAddress instantiates a new BillingAddress object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBillingAddress(city string, country string, houseNumberOrName string, postalCode string, street string) *BillingAddress {
	this := BillingAddress{}
	this.City = city
	this.Country = country
	this.HouseNumberOrName = houseNumberOrName
	this.PostalCode = postalCode
	this.Street = street
	return &this
}

// NewBillingAddressWithDefaults instantiates a new BillingAddress object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBillingAddressWithDefaults() *BillingAddress {
	this := BillingAddress{}
	return &this
}

// GetCity returns the City field value
func (o *BillingAddress) GetCity() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.City
}

// GetCityOk returns a tuple with the City field value
// and a boolean to check if the value has been set.
func (o *BillingAddress) GetCityOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.City, true
}

// SetCity sets field value
func (o *BillingAddress) SetCity(v string) {
	o.City = v
}

// GetCountry returns the Country field value
func (o *BillingAddress) GetCountry() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Country
}

// GetCountryOk returns a tuple with the Country field value
// and a boolean to check if the value has been set.
func (o *BillingAddress) GetCountryOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Country, true
}

// SetCountry sets field value
func (o *BillingAddress) SetCountry(v string) {
	o.Country = v
}

// GetHouseNumberOrName returns the HouseNumberOrName field value
func (o *BillingAddress) GetHouseNumberOrName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.HouseNumberOrName
}

// GetHouseNumberOrNameOk returns a tuple with the HouseNumberOrName field value
// and a boolean to check if the value has been set.
func (o *BillingAddress) GetHouseNumberOrNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.HouseNumberOrName, true
}

// SetHouseNumberOrName sets field value
func (o *BillingAddress) SetHouseNumberOrName(v string) {
	o.HouseNumberOrName = v
}

// GetPostalCode returns the PostalCode field value
func (o *BillingAddress) GetPostalCode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PostalCode
}

// GetPostalCodeOk returns a tuple with the PostalCode field value
// and a boolean to check if the value has been set.
func (o *BillingAddress) GetPostalCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PostalCode, true
}

// SetPostalCode sets field value
func (o *BillingAddress) SetPostalCode(v string) {
	o.PostalCode = v
}

// GetStateOrProvince returns the StateOrProvince field value if set, zero value otherwise.
func (o *BillingAddress) GetStateOrProvince() string {
	if o == nil || common.IsNil(o.StateOrProvince) {
		var ret string
		return ret
	}
	return *o.StateOrProvince
}

// GetStateOrProvinceOk returns a tuple with the StateOrProvince field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BillingAddress) GetStateOrProvinceOk() (*string, bool) {
	if o == nil || common.IsNil(o.StateOrProvince) {
		return nil, false
	}
	return o.StateOrProvince, true
}

// HasStateOrProvince returns a boolean if a field has been set.
func (o *BillingAddress) HasStateOrProvince() bool {
	if o != nil && !common.IsNil(o.StateOrProvince) {
		return true
	}

	return false
}

// SetStateOrProvince gets a reference to the given string and assigns it to the StateOrProvince field.
func (o *BillingAddress) SetStateOrProvince(v string) {
	o.StateOrProvince = &v
}

// GetStreet returns the Street field value
func (o *BillingAddress) GetStreet() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Street
}

// GetStreetOk returns a tuple with the Street field value
// and a boolean to check if the value has been set.
func (o *BillingAddress) GetStreetOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Street, true
}

// SetStreet sets field value
func (o *BillingAddress) SetStreet(v string) {
	o.Street = v
}

func (o BillingAddress) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BillingAddress) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["city"] = o.City
	toSerialize["country"] = o.Country
	toSerialize["houseNumberOrName"] = o.HouseNumberOrName
	toSerialize["postalCode"] = o.PostalCode
	if !common.IsNil(o.StateOrProvince) {
		toSerialize["stateOrProvince"] = o.StateOrProvince
	}
	toSerialize["street"] = o.Street
	return toSerialize, nil
}

type NullableBillingAddress struct {
	value *BillingAddress
	isSet bool
}

func (v NullableBillingAddress) Get() *BillingAddress {
	return v.value
}

func (v *NullableBillingAddress) Set(val *BillingAddress) {
	v.value = val
	v.isSet = true
}

func (v NullableBillingAddress) IsSet() bool {
	return v.isSet
}

func (v *NullableBillingAddress) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBillingAddress(val *BillingAddress) *NullableBillingAddress {
	return &NullableBillingAddress{value: val, isSet: true}
}

func (v NullableBillingAddress) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBillingAddress) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
