/*
Management Webhooks

API version: 3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package managementwebhook

import (
	"encoding/json"

	"github.com/adyen/adyen-go-api-library/v16/src/common"
)

// checks if the TerminalAssignmentNotificationResponse type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &TerminalAssignmentNotificationResponse{}

// TerminalAssignmentNotificationResponse struct for TerminalAssignmentNotificationResponse
type TerminalAssignmentNotificationResponse struct {
	// Respond with any **2xx** HTTP status code to [accept the webhook](https://docs.adyen.com/development-resources/webhooks#accept-notifications).
	NotificationResponse *string `json:"notificationResponse,omitempty"`
}

// NewTerminalAssignmentNotificationResponse instantiates a new TerminalAssignmentNotificationResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTerminalAssignmentNotificationResponse() *TerminalAssignmentNotificationResponse {
	this := TerminalAssignmentNotificationResponse{}
	return &this
}

// NewTerminalAssignmentNotificationResponseWithDefaults instantiates a new TerminalAssignmentNotificationResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTerminalAssignmentNotificationResponseWithDefaults() *TerminalAssignmentNotificationResponse {
	this := TerminalAssignmentNotificationResponse{}
	return &this
}

// GetNotificationResponse returns the NotificationResponse field value if set, zero value otherwise.
func (o *TerminalAssignmentNotificationResponse) GetNotificationResponse() string {
	if o == nil || common.IsNil(o.NotificationResponse) {
		var ret string
		return ret
	}
	return *o.NotificationResponse
}

// GetNotificationResponseOk returns a tuple with the NotificationResponse field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TerminalAssignmentNotificationResponse) GetNotificationResponseOk() (*string, bool) {
	if o == nil || common.IsNil(o.NotificationResponse) {
		return nil, false
	}
	return o.NotificationResponse, true
}

// HasNotificationResponse returns a boolean if a field has been set.
func (o *TerminalAssignmentNotificationResponse) HasNotificationResponse() bool {
	if o != nil && !common.IsNil(o.NotificationResponse) {
		return true
	}

	return false
}

// SetNotificationResponse gets a reference to the given string and assigns it to the NotificationResponse field.
func (o *TerminalAssignmentNotificationResponse) SetNotificationResponse(v string) {
	o.NotificationResponse = &v
}

func (o TerminalAssignmentNotificationResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TerminalAssignmentNotificationResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.NotificationResponse) {
		toSerialize["notificationResponse"] = o.NotificationResponse
	}
	return toSerialize, nil
}

type NullableTerminalAssignmentNotificationResponse struct {
	value *TerminalAssignmentNotificationResponse
	isSet bool
}

func (v NullableTerminalAssignmentNotificationResponse) Get() *TerminalAssignmentNotificationResponse {
	return v.value
}

func (v *NullableTerminalAssignmentNotificationResponse) Set(val *TerminalAssignmentNotificationResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableTerminalAssignmentNotificationResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableTerminalAssignmentNotificationResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTerminalAssignmentNotificationResponse(val *TerminalAssignmentNotificationResponse) *NullableTerminalAssignmentNotificationResponse {
	return &NullableTerminalAssignmentNotificationResponse{value: val, isSet: true}
}

func (v NullableTerminalAssignmentNotificationResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTerminalAssignmentNotificationResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}