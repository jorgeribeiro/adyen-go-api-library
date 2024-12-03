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

// checks if the DayOfWeekRestriction type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &DayOfWeekRestriction{}

// DayOfWeekRestriction struct for DayOfWeekRestriction
type DayOfWeekRestriction struct {
	// Defines how the condition must be evaluated.
	Operation string `json:"operation"`
	// List of days of the week.  Possible values: **monday**, **tuesday**, **wednesday**, **thursday**, **friday**, **saturday**, **sunday**.
	Value []string `json:"value,omitempty"`
}

// NewDayOfWeekRestriction instantiates a new DayOfWeekRestriction object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDayOfWeekRestriction(operation string) *DayOfWeekRestriction {
	this := DayOfWeekRestriction{}
	this.Operation = operation
	return &this
}

// NewDayOfWeekRestrictionWithDefaults instantiates a new DayOfWeekRestriction object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDayOfWeekRestrictionWithDefaults() *DayOfWeekRestriction {
	this := DayOfWeekRestriction{}
	return &this
}

// GetOperation returns the Operation field value
func (o *DayOfWeekRestriction) GetOperation() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Operation
}

// GetOperationOk returns a tuple with the Operation field value
// and a boolean to check if the value has been set.
func (o *DayOfWeekRestriction) GetOperationOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Operation, true
}

// SetOperation sets field value
func (o *DayOfWeekRestriction) SetOperation(v string) {
	o.Operation = v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *DayOfWeekRestriction) GetValue() []string {
	if o == nil || common.IsNil(o.Value) {
		var ret []string
		return ret
	}
	return o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DayOfWeekRestriction) GetValueOk() ([]string, bool) {
	if o == nil || common.IsNil(o.Value) {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *DayOfWeekRestriction) HasValue() bool {
	if o != nil && !common.IsNil(o.Value) {
		return true
	}

	return false
}

// SetValue gets a reference to the given []string and assigns it to the Value field.
func (o *DayOfWeekRestriction) SetValue(v []string) {
	o.Value = v
}

func (o DayOfWeekRestriction) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DayOfWeekRestriction) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["operation"] = o.Operation
	if !common.IsNil(o.Value) {
		toSerialize["value"] = o.Value
	}
	return toSerialize, nil
}

type NullableDayOfWeekRestriction struct {
	value *DayOfWeekRestriction
	isSet bool
}

func (v NullableDayOfWeekRestriction) Get() *DayOfWeekRestriction {
	return v.value
}

func (v *NullableDayOfWeekRestriction) Set(val *DayOfWeekRestriction) {
	v.value = val
	v.isSet = true
}

func (v NullableDayOfWeekRestriction) IsSet() bool {
	return v.isSet
}

func (v *NullableDayOfWeekRestriction) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDayOfWeekRestriction(val *DayOfWeekRestriction) *NullableDayOfWeekRestriction {
	return &NullableDayOfWeekRestriction{value: val, isSet: true}
}

func (v NullableDayOfWeekRestriction) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDayOfWeekRestriction) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
