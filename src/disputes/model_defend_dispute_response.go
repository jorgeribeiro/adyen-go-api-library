/*
Disputes API

API version: 30
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package disputes

import (
	"encoding/json"

	"github.com/adyen/adyen-go-api-library/v8/src/common"
)

// checks if the DefendDisputeResponse type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &DefendDisputeResponse{}

// DefendDisputeResponse struct for DefendDisputeResponse
type DefendDisputeResponse struct {
	DisputeServiceResult DisputeServiceResult `json:"disputeServiceResult"`
}

// NewDefendDisputeResponse instantiates a new DefendDisputeResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDefendDisputeResponse(disputeServiceResult DisputeServiceResult) *DefendDisputeResponse {
	this := DefendDisputeResponse{}
	this.DisputeServiceResult = disputeServiceResult
	return &this
}

// NewDefendDisputeResponseWithDefaults instantiates a new DefendDisputeResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDefendDisputeResponseWithDefaults() *DefendDisputeResponse {
	this := DefendDisputeResponse{}
	return &this
}

// GetDisputeServiceResult returns the DisputeServiceResult field value
func (o *DefendDisputeResponse) GetDisputeServiceResult() DisputeServiceResult {
	if o == nil {
		var ret DisputeServiceResult
		return ret
	}

	return o.DisputeServiceResult
}

// GetDisputeServiceResultOk returns a tuple with the DisputeServiceResult field value
// and a boolean to check if the value has been set.
func (o *DefendDisputeResponse) GetDisputeServiceResultOk() (*DisputeServiceResult, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DisputeServiceResult, true
}

// SetDisputeServiceResult sets field value
func (o *DefendDisputeResponse) SetDisputeServiceResult(v DisputeServiceResult) {
	o.DisputeServiceResult = v
}

func (o DefendDisputeResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DefendDisputeResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["disputeServiceResult"] = o.DisputeServiceResult
	return toSerialize, nil
}

type NullableDefendDisputeResponse struct {
	value *DefendDisputeResponse
	isSet bool
}

func (v NullableDefendDisputeResponse) Get() *DefendDisputeResponse {
	return v.value
}

func (v *NullableDefendDisputeResponse) Set(val *DefendDisputeResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableDefendDisputeResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableDefendDisputeResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDefendDisputeResponse(val *DefendDisputeResponse) *NullableDefendDisputeResponse {
	return &NullableDefendDisputeResponse{value: val, isSet: true}
}

func (v NullableDefendDisputeResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDefendDisputeResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}