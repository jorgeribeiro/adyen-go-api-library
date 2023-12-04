/*
Configuration webhooks

API version: 1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package configurationwebhook

import (
	"encoding/json"

	"github.com/adyen/adyen-go-api-library/v8/src/common"
)

// checks if the SweepConfigurationNotificationRequest type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &SweepConfigurationNotificationRequest{}

// SweepConfigurationNotificationRequest struct for SweepConfigurationNotificationRequest
type SweepConfigurationNotificationRequest struct {
	Data SweepConfigurationNotificationData `json:"data"`
	// The environment from which the webhook originated.  Possible values: **test**, **live**.
	Environment string `json:"environment"`
	// Type of webhook.
	Type string `json:"type"`
}

// NewSweepConfigurationNotificationRequest instantiates a new SweepConfigurationNotificationRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSweepConfigurationNotificationRequest(data SweepConfigurationNotificationData, environment string, type_ string) *SweepConfigurationNotificationRequest {
	this := SweepConfigurationNotificationRequest{}
	this.Data = data
	this.Environment = environment
	this.Type = type_
	return &this
}

// NewSweepConfigurationNotificationRequestWithDefaults instantiates a new SweepConfigurationNotificationRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSweepConfigurationNotificationRequestWithDefaults() *SweepConfigurationNotificationRequest {
	this := SweepConfigurationNotificationRequest{}
	return &this
}

// GetData returns the Data field value
func (o *SweepConfigurationNotificationRequest) GetData() SweepConfigurationNotificationData {
	if o == nil {
		var ret SweepConfigurationNotificationData
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *SweepConfigurationNotificationRequest) GetDataOk() (*SweepConfigurationNotificationData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *SweepConfigurationNotificationRequest) SetData(v SweepConfigurationNotificationData) {
	o.Data = v
}

// GetEnvironment returns the Environment field value
func (o *SweepConfigurationNotificationRequest) GetEnvironment() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Environment
}

// GetEnvironmentOk returns a tuple with the Environment field value
// and a boolean to check if the value has been set.
func (o *SweepConfigurationNotificationRequest) GetEnvironmentOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Environment, true
}

// SetEnvironment sets field value
func (o *SweepConfigurationNotificationRequest) SetEnvironment(v string) {
	o.Environment = v
}

// GetType returns the Type field value
func (o *SweepConfigurationNotificationRequest) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *SweepConfigurationNotificationRequest) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *SweepConfigurationNotificationRequest) SetType(v string) {
	o.Type = v
}

func (o SweepConfigurationNotificationRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SweepConfigurationNotificationRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	toSerialize["environment"] = o.Environment
	toSerialize["type"] = o.Type
	return toSerialize, nil
}

type NullableSweepConfigurationNotificationRequest struct {
	value *SweepConfigurationNotificationRequest
	isSet bool
}

func (v NullableSweepConfigurationNotificationRequest) Get() *SweepConfigurationNotificationRequest {
	return v.value
}

func (v *NullableSweepConfigurationNotificationRequest) Set(val *SweepConfigurationNotificationRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableSweepConfigurationNotificationRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableSweepConfigurationNotificationRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSweepConfigurationNotificationRequest(val *SweepConfigurationNotificationRequest) *NullableSweepConfigurationNotificationRequest {
	return &NullableSweepConfigurationNotificationRequest{value: val, isSet: true}
}

func (v NullableSweepConfigurationNotificationRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSweepConfigurationNotificationRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

func (o *SweepConfigurationNotificationRequest) isValidType() bool {
	var allowedEnumValues = []string{"balancePlatform.balanceAccountSweep.created", "balancePlatform.balanceAccountSweep.updated", "balancePlatform.balanceAccountSweep.deleted"}
	for _, allowed := range allowedEnumValues {
		if o.GetType() == allowed {
			return true
		}
	}
	return false
}