/*
Management API

API version: 3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package management

import (
	"encoding/json"

	"github.com/adyen/adyen-go-api-library/v16/src/common"
)

// checks if the ReleaseUpdateDetails type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &ReleaseUpdateDetails{}

// ReleaseUpdateDetails struct for ReleaseUpdateDetails
type ReleaseUpdateDetails struct {
	// Type of terminal action: Update Release.
	Type *string `json:"type,omitempty"`
	// Boolean flag that tells if the terminal should update at the first next maintenance call. If false, terminal will update on its configured reboot time.
	UpdateAtFirstMaintenanceCall *bool `json:"updateAtFirstMaintenanceCall,omitempty"`
}

// NewReleaseUpdateDetails instantiates a new ReleaseUpdateDetails object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewReleaseUpdateDetails() *ReleaseUpdateDetails {
	this := ReleaseUpdateDetails{}
	var type_ string = "ReleaseUpdate"
	this.Type = &type_
	return &this
}

// NewReleaseUpdateDetailsWithDefaults instantiates a new ReleaseUpdateDetails object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewReleaseUpdateDetailsWithDefaults() *ReleaseUpdateDetails {
	this := ReleaseUpdateDetails{}
	var type_ string = "ReleaseUpdate"
	this.Type = &type_
	return &this
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *ReleaseUpdateDetails) GetType() string {
	if o == nil || common.IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReleaseUpdateDetails) GetTypeOk() (*string, bool) {
	if o == nil || common.IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *ReleaseUpdateDetails) HasType() bool {
	if o != nil && !common.IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *ReleaseUpdateDetails) SetType(v string) {
	o.Type = &v
}

// GetUpdateAtFirstMaintenanceCall returns the UpdateAtFirstMaintenanceCall field value if set, zero value otherwise.
func (o *ReleaseUpdateDetails) GetUpdateAtFirstMaintenanceCall() bool {
	if o == nil || common.IsNil(o.UpdateAtFirstMaintenanceCall) {
		var ret bool
		return ret
	}
	return *o.UpdateAtFirstMaintenanceCall
}

// GetUpdateAtFirstMaintenanceCallOk returns a tuple with the UpdateAtFirstMaintenanceCall field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReleaseUpdateDetails) GetUpdateAtFirstMaintenanceCallOk() (*bool, bool) {
	if o == nil || common.IsNil(o.UpdateAtFirstMaintenanceCall) {
		return nil, false
	}
	return o.UpdateAtFirstMaintenanceCall, true
}

// HasUpdateAtFirstMaintenanceCall returns a boolean if a field has been set.
func (o *ReleaseUpdateDetails) HasUpdateAtFirstMaintenanceCall() bool {
	if o != nil && !common.IsNil(o.UpdateAtFirstMaintenanceCall) {
		return true
	}

	return false
}

// SetUpdateAtFirstMaintenanceCall gets a reference to the given bool and assigns it to the UpdateAtFirstMaintenanceCall field.
func (o *ReleaseUpdateDetails) SetUpdateAtFirstMaintenanceCall(v bool) {
	o.UpdateAtFirstMaintenanceCall = &v
}

func (o ReleaseUpdateDetails) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ReleaseUpdateDetails) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !common.IsNil(o.UpdateAtFirstMaintenanceCall) {
		toSerialize["updateAtFirstMaintenanceCall"] = o.UpdateAtFirstMaintenanceCall
	}
	return toSerialize, nil
}

type NullableReleaseUpdateDetails struct {
	value *ReleaseUpdateDetails
	isSet bool
}

func (v NullableReleaseUpdateDetails) Get() *ReleaseUpdateDetails {
	return v.value
}

func (v *NullableReleaseUpdateDetails) Set(val *ReleaseUpdateDetails) {
	v.value = val
	v.isSet = true
}

func (v NullableReleaseUpdateDetails) IsSet() bool {
	return v.isSet
}

func (v *NullableReleaseUpdateDetails) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableReleaseUpdateDetails(val *ReleaseUpdateDetails) *NullableReleaseUpdateDetails {
	return &NullableReleaseUpdateDetails{value: val, isSet: true}
}

func (v NullableReleaseUpdateDetails) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableReleaseUpdateDetails) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

func (o *ReleaseUpdateDetails) isValidType() bool {
	var allowedEnumValues = []string{"ReleaseUpdate"}
	for _, allowed := range allowedEnumValues {
		if o.GetType() == allowed {
			return true
		}
	}
	return false
}
