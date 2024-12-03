/*
Adyen Recurring API

API version: 68
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package recurring

import (
	"encoding/json"

	"github.com/adyen/adyen-go-api-library/v16/src/common"
)

// checks if the ScheduleAccountUpdaterResult type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &ScheduleAccountUpdaterResult{}

// ScheduleAccountUpdaterResult struct for ScheduleAccountUpdaterResult
type ScheduleAccountUpdaterResult struct {
	// Adyen's 16-character unique reference associated with the transaction. This value is globally unique; quote it when communicating with us about this request.
	PspReference string `json:"pspReference"`
	// The result of scheduling an Account Updater. If scheduling was successful, this field returns **Success**; otherwise it contains the error message.
	Result string `json:"result"`
}

// NewScheduleAccountUpdaterResult instantiates a new ScheduleAccountUpdaterResult object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewScheduleAccountUpdaterResult(pspReference string, result string) *ScheduleAccountUpdaterResult {
	this := ScheduleAccountUpdaterResult{}
	this.PspReference = pspReference
	this.Result = result
	return &this
}

// NewScheduleAccountUpdaterResultWithDefaults instantiates a new ScheduleAccountUpdaterResult object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewScheduleAccountUpdaterResultWithDefaults() *ScheduleAccountUpdaterResult {
	this := ScheduleAccountUpdaterResult{}
	return &this
}

// GetPspReference returns the PspReference field value
func (o *ScheduleAccountUpdaterResult) GetPspReference() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PspReference
}

// GetPspReferenceOk returns a tuple with the PspReference field value
// and a boolean to check if the value has been set.
func (o *ScheduleAccountUpdaterResult) GetPspReferenceOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PspReference, true
}

// SetPspReference sets field value
func (o *ScheduleAccountUpdaterResult) SetPspReference(v string) {
	o.PspReference = v
}

// GetResult returns the Result field value
func (o *ScheduleAccountUpdaterResult) GetResult() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Result
}

// GetResultOk returns a tuple with the Result field value
// and a boolean to check if the value has been set.
func (o *ScheduleAccountUpdaterResult) GetResultOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Result, true
}

// SetResult sets field value
func (o *ScheduleAccountUpdaterResult) SetResult(v string) {
	o.Result = v
}

func (o ScheduleAccountUpdaterResult) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ScheduleAccountUpdaterResult) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["pspReference"] = o.PspReference
	toSerialize["result"] = o.Result
	return toSerialize, nil
}

type NullableScheduleAccountUpdaterResult struct {
	value *ScheduleAccountUpdaterResult
	isSet bool
}

func (v NullableScheduleAccountUpdaterResult) Get() *ScheduleAccountUpdaterResult {
	return v.value
}

func (v *NullableScheduleAccountUpdaterResult) Set(val *ScheduleAccountUpdaterResult) {
	v.value = val
	v.isSet = true
}

func (v NullableScheduleAccountUpdaterResult) IsSet() bool {
	return v.isSet
}

func (v *NullableScheduleAccountUpdaterResult) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableScheduleAccountUpdaterResult(val *ScheduleAccountUpdaterResult) *NullableScheduleAccountUpdaterResult {
	return &NullableScheduleAccountUpdaterResult{value: val, isSet: true}
}

func (v NullableScheduleAccountUpdaterResult) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableScheduleAccountUpdaterResult) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
