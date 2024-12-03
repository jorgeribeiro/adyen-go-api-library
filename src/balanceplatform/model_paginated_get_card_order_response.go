/*
Configuration API

API version: 2
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package balanceplatform

import (
	"encoding/json"

	"github.com/adyen/adyen-go-api-library/v16/src/common"
)

// checks if the PaginatedGetCardOrderResponse type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &PaginatedGetCardOrderResponse{}

// PaginatedGetCardOrderResponse struct for PaginatedGetCardOrderResponse
type PaginatedGetCardOrderResponse struct {
	// Contains objects with information about card orders.
	CardOrders []CardOrder `json:"cardOrders,omitempty"`
	// Indicates whether there are more items on the next page.
	HasNext bool `json:"hasNext"`
	// Indicates whether there are more items on the previous page.
	HasPrevious bool `json:"hasPrevious"`
}

// NewPaginatedGetCardOrderResponse instantiates a new PaginatedGetCardOrderResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPaginatedGetCardOrderResponse(hasNext bool, hasPrevious bool) *PaginatedGetCardOrderResponse {
	this := PaginatedGetCardOrderResponse{}
	this.HasNext = hasNext
	this.HasPrevious = hasPrevious
	return &this
}

// NewPaginatedGetCardOrderResponseWithDefaults instantiates a new PaginatedGetCardOrderResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPaginatedGetCardOrderResponseWithDefaults() *PaginatedGetCardOrderResponse {
	this := PaginatedGetCardOrderResponse{}
	return &this
}

// GetCardOrders returns the CardOrders field value if set, zero value otherwise.
func (o *PaginatedGetCardOrderResponse) GetCardOrders() []CardOrder {
	if o == nil || common.IsNil(o.CardOrders) {
		var ret []CardOrder
		return ret
	}
	return o.CardOrders
}

// GetCardOrdersOk returns a tuple with the CardOrders field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaginatedGetCardOrderResponse) GetCardOrdersOk() ([]CardOrder, bool) {
	if o == nil || common.IsNil(o.CardOrders) {
		return nil, false
	}
	return o.CardOrders, true
}

// HasCardOrders returns a boolean if a field has been set.
func (o *PaginatedGetCardOrderResponse) HasCardOrders() bool {
	if o != nil && !common.IsNil(o.CardOrders) {
		return true
	}

	return false
}

// SetCardOrders gets a reference to the given []CardOrder and assigns it to the CardOrders field.
func (o *PaginatedGetCardOrderResponse) SetCardOrders(v []CardOrder) {
	o.CardOrders = v
}

// GetHasNext returns the HasNext field value
func (o *PaginatedGetCardOrderResponse) GetHasNext() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.HasNext
}

// GetHasNextOk returns a tuple with the HasNext field value
// and a boolean to check if the value has been set.
func (o *PaginatedGetCardOrderResponse) GetHasNextOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.HasNext, true
}

// SetHasNext sets field value
func (o *PaginatedGetCardOrderResponse) SetHasNext(v bool) {
	o.HasNext = v
}

// GetHasPrevious returns the HasPrevious field value
func (o *PaginatedGetCardOrderResponse) GetHasPrevious() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.HasPrevious
}

// GetHasPreviousOk returns a tuple with the HasPrevious field value
// and a boolean to check if the value has been set.
func (o *PaginatedGetCardOrderResponse) GetHasPreviousOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.HasPrevious, true
}

// SetHasPrevious sets field value
func (o *PaginatedGetCardOrderResponse) SetHasPrevious(v bool) {
	o.HasPrevious = v
}

func (o PaginatedGetCardOrderResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PaginatedGetCardOrderResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.CardOrders) {
		toSerialize["cardOrders"] = o.CardOrders
	}
	toSerialize["hasNext"] = o.HasNext
	toSerialize["hasPrevious"] = o.HasPrevious
	return toSerialize, nil
}

type NullablePaginatedGetCardOrderResponse struct {
	value *PaginatedGetCardOrderResponse
	isSet bool
}

func (v NullablePaginatedGetCardOrderResponse) Get() *PaginatedGetCardOrderResponse {
	return v.value
}

func (v *NullablePaginatedGetCardOrderResponse) Set(val *PaginatedGetCardOrderResponse) {
	v.value = val
	v.isSet = true
}

func (v NullablePaginatedGetCardOrderResponse) IsSet() bool {
	return v.isSet
}

func (v *NullablePaginatedGetCardOrderResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePaginatedGetCardOrderResponse(val *PaginatedGetCardOrderResponse) *NullablePaginatedGetCardOrderResponse {
	return &NullablePaginatedGetCardOrderResponse{value: val, isSet: true}
}

func (v NullablePaginatedGetCardOrderResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePaginatedGetCardOrderResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
