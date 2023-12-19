/*
 * Adyen for Platforms: Notifications
 *
 * The Notification API sends notifications to the endpoints specified in a given subscription. Subscriptions are managed through the Notification Configuration API. The API specifications listed here detail the format of each notification.  For more information, refer to our [documentation](https://docs.adyen.com/platforms/notifications).
 *
 * API version: 6
 * Contact: support@adyen.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package platformsnotificationevents

import (
	"encoding/json"
)

// OperationStatus struct for OperationStatus
type OperationStatus struct {
	Message *Message `json:"message,omitempty"`
	// The status code.
	StatusCode *string `json:"statusCode,omitempty"`
}

// NewOperationStatus instantiates a new OperationStatus object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOperationStatus() *OperationStatus {
	this := OperationStatus{}
	return &this
}

// NewOperationStatusWithDefaults instantiates a new OperationStatus object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOperationStatusWithDefaults() *OperationStatus {
	this := OperationStatus{}
	return &this
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *OperationStatus) GetMessage() Message {
	if o == nil || o.Message == nil {
		var ret Message
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OperationStatus) GetMessageOk() (*Message, bool) {
	if o == nil || o.Message == nil {
		return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *OperationStatus) HasMessage() bool {
	if o != nil && o.Message != nil {
		return true
	}

	return false
}

// SetMessage gets a reference to the given Message and assigns it to the Message field.
func (o *OperationStatus) SetMessage(v Message) {
	o.Message = &v
}

// GetStatusCode returns the StatusCode field value if set, zero value otherwise.
func (o *OperationStatus) GetStatusCode() string {
	if o == nil || o.StatusCode == nil {
		var ret string
		return ret
	}
	return *o.StatusCode
}

// GetStatusCodeOk returns a tuple with the StatusCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OperationStatus) GetStatusCodeOk() (*string, bool) {
	if o == nil || o.StatusCode == nil {
		return nil, false
	}
	return o.StatusCode, true
}

// HasStatusCode returns a boolean if a field has been set.
func (o *OperationStatus) HasStatusCode() bool {
	if o != nil && o.StatusCode != nil {
		return true
	}

	return false
}

// SetStatusCode gets a reference to the given string and assigns it to the StatusCode field.
func (o *OperationStatus) SetStatusCode(v string) {
	o.StatusCode = &v
}

func (o OperationStatus) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Message != nil {
		toSerialize["message"] = o.Message
	}
	if o.StatusCode != nil {
		toSerialize["statusCode"] = o.StatusCode
	}
	return json.Marshal(toSerialize)
}

type NullableOperationStatus struct {
	value *OperationStatus
	isSet bool
}

func (v NullableOperationStatus) Get() *OperationStatus {
	return v.value
}

func (v *NullableOperationStatus) Set(val *OperationStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableOperationStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableOperationStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOperationStatus(val *OperationStatus) *NullableOperationStatus {
	return &NullableOperationStatus{value: val, isSet: true}
}

func (v NullableOperationStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOperationStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
