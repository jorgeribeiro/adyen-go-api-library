/*
Legal Entity Management API

API version: 3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package legalentity

import (
	"encoding/json"

	"github.com/adyen/adyen-go-api-library/v6/src/common"
)

// checks if the OwnerEntity type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &OwnerEntity{}

// OwnerEntity struct for OwnerEntity
type OwnerEntity struct {
	// Unique identifier of the resource that owns the document. For `type` **legalEntity**, this value is the unique identifier of the [legal entity](https://docs.adyen.com/api-explorer/legalentity/latest/post/legalEntities#responses-200-id). For `type` **bankAccount**, this value is the unique identifier of the [transfer instrument](https://docs.adyen.com/api-explorer/legalentity/latest/post/transferInstruments#responses-200-id).
	Id string `json:"id"`
	// Type of resource that owns the document.  Possible values: **legalEntity**, **bankAccount**.
	Type string `json:"type"`
}

// NewOwnerEntity instantiates a new OwnerEntity object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOwnerEntity(id string, type_ string) *OwnerEntity {
	this := OwnerEntity{}
	this.Id = id
	this.Type = type_
	return &this
}

// NewOwnerEntityWithDefaults instantiates a new OwnerEntity object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOwnerEntityWithDefaults() *OwnerEntity {
	this := OwnerEntity{}
	return &this
}

// GetId returns the Id field value
func (o *OwnerEntity) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *OwnerEntity) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *OwnerEntity) SetId(v string) {
	o.Id = v
}

// GetType returns the Type field value
func (o *OwnerEntity) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *OwnerEntity) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *OwnerEntity) SetType(v string) {
	o.Type = v
}

func (o OwnerEntity) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OwnerEntity) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["type"] = o.Type
	return toSerialize, nil
}

type NullableOwnerEntity struct {
	value *OwnerEntity
	isSet bool
}

func (v NullableOwnerEntity) Get() *OwnerEntity {
	return v.value
}

func (v *NullableOwnerEntity) Set(val *OwnerEntity) {
	v.value = val
	v.isSet = true
}

func (v NullableOwnerEntity) IsSet() bool {
	return v.isSet
}

func (v *NullableOwnerEntity) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOwnerEntity(val *OwnerEntity) *NullableOwnerEntity {
	return &NullableOwnerEntity{value: val, isSet: true}
}

func (v NullableOwnerEntity) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOwnerEntity) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
