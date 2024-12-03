/*
Transfer webhooks

API version: 4
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package transferwebhook

import (
	"encoding/json"

	"github.com/adyen/adyen-go-api-library/v16/src/common"
)

// checks if the ExternalReason type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &ExternalReason{}

// ExternalReason struct for ExternalReason
type ExternalReason struct {
	// The reason code.
	Code *string `json:"code,omitempty"`
	// The description of the reason code.
	Description *string `json:"description,omitempty"`
	// The namespace for the reason code.
	Namespace *string `json:"namespace,omitempty"`
}

// NewExternalReason instantiates a new ExternalReason object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewExternalReason() *ExternalReason {
	this := ExternalReason{}
	return &this
}

// NewExternalReasonWithDefaults instantiates a new ExternalReason object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewExternalReasonWithDefaults() *ExternalReason {
	this := ExternalReason{}
	return &this
}

// GetCode returns the Code field value if set, zero value otherwise.
func (o *ExternalReason) GetCode() string {
	if o == nil || common.IsNil(o.Code) {
		var ret string
		return ret
	}
	return *o.Code
}

// GetCodeOk returns a tuple with the Code field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExternalReason) GetCodeOk() (*string, bool) {
	if o == nil || common.IsNil(o.Code) {
		return nil, false
	}
	return o.Code, true
}

// HasCode returns a boolean if a field has been set.
func (o *ExternalReason) HasCode() bool {
	if o != nil && !common.IsNil(o.Code) {
		return true
	}

	return false
}

// SetCode gets a reference to the given string and assigns it to the Code field.
func (o *ExternalReason) SetCode(v string) {
	o.Code = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *ExternalReason) GetDescription() string {
	if o == nil || common.IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExternalReason) GetDescriptionOk() (*string, bool) {
	if o == nil || common.IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *ExternalReason) HasDescription() bool {
	if o != nil && !common.IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *ExternalReason) SetDescription(v string) {
	o.Description = &v
}

// GetNamespace returns the Namespace field value if set, zero value otherwise.
func (o *ExternalReason) GetNamespace() string {
	if o == nil || common.IsNil(o.Namespace) {
		var ret string
		return ret
	}
	return *o.Namespace
}

// GetNamespaceOk returns a tuple with the Namespace field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExternalReason) GetNamespaceOk() (*string, bool) {
	if o == nil || common.IsNil(o.Namespace) {
		return nil, false
	}
	return o.Namespace, true
}

// HasNamespace returns a boolean if a field has been set.
func (o *ExternalReason) HasNamespace() bool {
	if o != nil && !common.IsNil(o.Namespace) {
		return true
	}

	return false
}

// SetNamespace gets a reference to the given string and assigns it to the Namespace field.
func (o *ExternalReason) SetNamespace(v string) {
	o.Namespace = &v
}

func (o ExternalReason) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ExternalReason) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.Code) {
		toSerialize["code"] = o.Code
	}
	if !common.IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !common.IsNil(o.Namespace) {
		toSerialize["namespace"] = o.Namespace
	}
	return toSerialize, nil
}

type NullableExternalReason struct {
	value *ExternalReason
	isSet bool
}

func (v NullableExternalReason) Get() *ExternalReason {
	return v.value
}

func (v *NullableExternalReason) Set(val *ExternalReason) {
	v.value = val
	v.isSet = true
}

func (v NullableExternalReason) IsSet() bool {
	return v.isSet
}

func (v *NullableExternalReason) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableExternalReason(val *ExternalReason) *NullableExternalReason {
	return &NullableExternalReason{value: val, isSet: true}
}

func (v NullableExternalReason) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableExternalReason) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
