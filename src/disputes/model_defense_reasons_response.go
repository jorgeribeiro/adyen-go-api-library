/*
Disputes API

API version: 30
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package disputes

import (
	"encoding/json"

	"github.com/adyen/adyen-go-api-library/v16/src/common"
)

// checks if the DefenseReasonsResponse type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &DefenseReasonsResponse{}

// DefenseReasonsResponse struct for DefenseReasonsResponse
type DefenseReasonsResponse struct {
	// The defense reasons that can be used to defend the dispute.
	DefenseReasons       []DefenseReason      `json:"defenseReasons,omitempty"`
	DisputeServiceResult DisputeServiceResult `json:"disputeServiceResult"`
}

// NewDefenseReasonsResponse instantiates a new DefenseReasonsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDefenseReasonsResponse(disputeServiceResult DisputeServiceResult) *DefenseReasonsResponse {
	this := DefenseReasonsResponse{}
	this.DisputeServiceResult = disputeServiceResult
	return &this
}

// NewDefenseReasonsResponseWithDefaults instantiates a new DefenseReasonsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDefenseReasonsResponseWithDefaults() *DefenseReasonsResponse {
	this := DefenseReasonsResponse{}
	return &this
}

// GetDefenseReasons returns the DefenseReasons field value if set, zero value otherwise.
func (o *DefenseReasonsResponse) GetDefenseReasons() []DefenseReason {
	if o == nil || common.IsNil(o.DefenseReasons) {
		var ret []DefenseReason
		return ret
	}
	return o.DefenseReasons
}

// GetDefenseReasonsOk returns a tuple with the DefenseReasons field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DefenseReasonsResponse) GetDefenseReasonsOk() ([]DefenseReason, bool) {
	if o == nil || common.IsNil(o.DefenseReasons) {
		return nil, false
	}
	return o.DefenseReasons, true
}

// HasDefenseReasons returns a boolean if a field has been set.
func (o *DefenseReasonsResponse) HasDefenseReasons() bool {
	if o != nil && !common.IsNil(o.DefenseReasons) {
		return true
	}

	return false
}

// SetDefenseReasons gets a reference to the given []DefenseReason and assigns it to the DefenseReasons field.
func (o *DefenseReasonsResponse) SetDefenseReasons(v []DefenseReason) {
	o.DefenseReasons = v
}

// GetDisputeServiceResult returns the DisputeServiceResult field value
func (o *DefenseReasonsResponse) GetDisputeServiceResult() DisputeServiceResult {
	if o == nil {
		var ret DisputeServiceResult
		return ret
	}

	return o.DisputeServiceResult
}

// GetDisputeServiceResultOk returns a tuple with the DisputeServiceResult field value
// and a boolean to check if the value has been set.
func (o *DefenseReasonsResponse) GetDisputeServiceResultOk() (*DisputeServiceResult, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DisputeServiceResult, true
}

// SetDisputeServiceResult sets field value
func (o *DefenseReasonsResponse) SetDisputeServiceResult(v DisputeServiceResult) {
	o.DisputeServiceResult = v
}

func (o DefenseReasonsResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DefenseReasonsResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.DefenseReasons) {
		toSerialize["defenseReasons"] = o.DefenseReasons
	}
	toSerialize["disputeServiceResult"] = o.DisputeServiceResult
	return toSerialize, nil
}

type NullableDefenseReasonsResponse struct {
	value *DefenseReasonsResponse
	isSet bool
}

func (v NullableDefenseReasonsResponse) Get() *DefenseReasonsResponse {
	return v.value
}

func (v *NullableDefenseReasonsResponse) Set(val *DefenseReasonsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableDefenseReasonsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableDefenseReasonsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDefenseReasonsResponse(val *DefenseReasonsResponse) *NullableDefenseReasonsResponse {
	return &NullableDefenseReasonsResponse{value: val, isSet: true}
}

func (v NullableDefenseReasonsResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDefenseReasonsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
