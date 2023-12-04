/*
Legal Entity Management API

API version: 3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package legalentity

import (
	"encoding/json"

	"github.com/adyen/adyen-go-api-library/v8/src/common"
)

// checks if the VerificationErrorRecursive type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &VerificationErrorRecursive{}

// VerificationErrorRecursive struct for VerificationErrorRecursive
type VerificationErrorRecursive struct {
	// Contains key-value pairs that specify the actions that the legal entity can do in your platform. The key is a capability required for your integration. For example, **issueCard** for Issuing.The value is an object containing the settings for the capability.
	Capabilities []string `json:"capabilities,omitempty"`
	// The general error code.
	Code *string `json:"code,omitempty"`
	// The general error message.
	Message *string `json:"message,omitempty"`
	// The type of error.
	Type *string `json:"type,omitempty"`
	// An object containing possible solutions to fix a verification error.
	RemediatingActions []RemediatingAction `json:"remediatingActions,omitempty"`
}

// NewVerificationErrorRecursive instantiates a new VerificationErrorRecursive object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVerificationErrorRecursive() *VerificationErrorRecursive {
	this := VerificationErrorRecursive{}
	return &this
}

// NewVerificationErrorRecursiveWithDefaults instantiates a new VerificationErrorRecursive object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVerificationErrorRecursiveWithDefaults() *VerificationErrorRecursive {
	this := VerificationErrorRecursive{}
	return &this
}

// GetCapabilities returns the Capabilities field value if set, zero value otherwise.
func (o *VerificationErrorRecursive) GetCapabilities() []string {
	if o == nil || common.IsNil(o.Capabilities) {
		var ret []string
		return ret
	}
	return o.Capabilities
}

// GetCapabilitiesOk returns a tuple with the Capabilities field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VerificationErrorRecursive) GetCapabilitiesOk() ([]string, bool) {
	if o == nil || common.IsNil(o.Capabilities) {
		return nil, false
	}
	return o.Capabilities, true
}

// HasCapabilities returns a boolean if a field has been set.
func (o *VerificationErrorRecursive) HasCapabilities() bool {
	if o != nil && !common.IsNil(o.Capabilities) {
		return true
	}

	return false
}

// SetCapabilities gets a reference to the given []string and assigns it to the Capabilities field.
func (o *VerificationErrorRecursive) SetCapabilities(v []string) {
	o.Capabilities = v
}

// GetCode returns the Code field value if set, zero value otherwise.
func (o *VerificationErrorRecursive) GetCode() string {
	if o == nil || common.IsNil(o.Code) {
		var ret string
		return ret
	}
	return *o.Code
}

// GetCodeOk returns a tuple with the Code field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VerificationErrorRecursive) GetCodeOk() (*string, bool) {
	if o == nil || common.IsNil(o.Code) {
		return nil, false
	}
	return o.Code, true
}

// HasCode returns a boolean if a field has been set.
func (o *VerificationErrorRecursive) HasCode() bool {
	if o != nil && !common.IsNil(o.Code) {
		return true
	}

	return false
}

// SetCode gets a reference to the given string and assigns it to the Code field.
func (o *VerificationErrorRecursive) SetCode(v string) {
	o.Code = &v
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *VerificationErrorRecursive) GetMessage() string {
	if o == nil || common.IsNil(o.Message) {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VerificationErrorRecursive) GetMessageOk() (*string, bool) {
	if o == nil || common.IsNil(o.Message) {
		return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *VerificationErrorRecursive) HasMessage() bool {
	if o != nil && !common.IsNil(o.Message) {
		return true
	}

	return false
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *VerificationErrorRecursive) SetMessage(v string) {
	o.Message = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *VerificationErrorRecursive) GetType() string {
	if o == nil || common.IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VerificationErrorRecursive) GetTypeOk() (*string, bool) {
	if o == nil || common.IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *VerificationErrorRecursive) HasType() bool {
	if o != nil && !common.IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *VerificationErrorRecursive) SetType(v string) {
	o.Type = &v
}

// GetRemediatingActions returns the RemediatingActions field value if set, zero value otherwise.
func (o *VerificationErrorRecursive) GetRemediatingActions() []RemediatingAction {
	if o == nil || common.IsNil(o.RemediatingActions) {
		var ret []RemediatingAction
		return ret
	}
	return o.RemediatingActions
}

// GetRemediatingActionsOk returns a tuple with the RemediatingActions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VerificationErrorRecursive) GetRemediatingActionsOk() ([]RemediatingAction, bool) {
	if o == nil || common.IsNil(o.RemediatingActions) {
		return nil, false
	}
	return o.RemediatingActions, true
}

// HasRemediatingActions returns a boolean if a field has been set.
func (o *VerificationErrorRecursive) HasRemediatingActions() bool {
	if o != nil && !common.IsNil(o.RemediatingActions) {
		return true
	}

	return false
}

// SetRemediatingActions gets a reference to the given []RemediatingAction and assigns it to the RemediatingActions field.
func (o *VerificationErrorRecursive) SetRemediatingActions(v []RemediatingAction) {
	o.RemediatingActions = v
}

func (o VerificationErrorRecursive) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o VerificationErrorRecursive) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.Capabilities) {
		toSerialize["capabilities"] = o.Capabilities
	}
	if !common.IsNil(o.Code) {
		toSerialize["code"] = o.Code
	}
	if !common.IsNil(o.Message) {
		toSerialize["message"] = o.Message
	}
	if !common.IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !common.IsNil(o.RemediatingActions) {
		toSerialize["remediatingActions"] = o.RemediatingActions
	}
	return toSerialize, nil
}

type NullableVerificationErrorRecursive struct {
	value *VerificationErrorRecursive
	isSet bool
}

func (v NullableVerificationErrorRecursive) Get() *VerificationErrorRecursive {
	return v.value
}

func (v *NullableVerificationErrorRecursive) Set(val *VerificationErrorRecursive) {
	v.value = val
	v.isSet = true
}

func (v NullableVerificationErrorRecursive) IsSet() bool {
	return v.isSet
}

func (v *NullableVerificationErrorRecursive) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVerificationErrorRecursive(val *VerificationErrorRecursive) *NullableVerificationErrorRecursive {
	return &NullableVerificationErrorRecursive{value: val, isSet: true}
}

func (v NullableVerificationErrorRecursive) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVerificationErrorRecursive) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

func (o *VerificationErrorRecursive) isValidType() bool {
	var allowedEnumValues = []string{"dataMissing", "dataReview", "invalidInput", "pendingStatus", "rejected"}
	for _, allowed := range allowedEnumValues {
		if o.GetType() == allowed {
			return true
		}
	}
	return false
}