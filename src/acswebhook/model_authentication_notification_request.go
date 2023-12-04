/*
Authentication webhooks

API version: 1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package acswebhook

import (
	"encoding/json"

	"github.com/adyen/adyen-go-api-library/v8/src/common"
)

// checks if the AuthenticationNotificationRequest type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &AuthenticationNotificationRequest{}

// AuthenticationNotificationRequest struct for AuthenticationNotificationRequest
type AuthenticationNotificationRequest struct {
	Data AuthenticationNotificationData `json:"data"`
	// The environment from which the webhook originated.  Possible values: **test**, **live**.
	Environment string `json:"environment"`
	// Type of notification.
	Type string `json:"type"`
}

// NewAuthenticationNotificationRequest instantiates a new AuthenticationNotificationRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAuthenticationNotificationRequest(data AuthenticationNotificationData, environment string, type_ string) *AuthenticationNotificationRequest {
	this := AuthenticationNotificationRequest{}
	this.Data = data
	this.Environment = environment
	this.Type = type_
	return &this
}

// NewAuthenticationNotificationRequestWithDefaults instantiates a new AuthenticationNotificationRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAuthenticationNotificationRequestWithDefaults() *AuthenticationNotificationRequest {
	this := AuthenticationNotificationRequest{}
	return &this
}

// GetData returns the Data field value
func (o *AuthenticationNotificationRequest) GetData() AuthenticationNotificationData {
	if o == nil {
		var ret AuthenticationNotificationData
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *AuthenticationNotificationRequest) GetDataOk() (*AuthenticationNotificationData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *AuthenticationNotificationRequest) SetData(v AuthenticationNotificationData) {
	o.Data = v
}

// GetEnvironment returns the Environment field value
func (o *AuthenticationNotificationRequest) GetEnvironment() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Environment
}

// GetEnvironmentOk returns a tuple with the Environment field value
// and a boolean to check if the value has been set.
func (o *AuthenticationNotificationRequest) GetEnvironmentOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Environment, true
}

// SetEnvironment sets field value
func (o *AuthenticationNotificationRequest) SetEnvironment(v string) {
	o.Environment = v
}

// GetType returns the Type field value
func (o *AuthenticationNotificationRequest) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *AuthenticationNotificationRequest) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *AuthenticationNotificationRequest) SetType(v string) {
	o.Type = v
}

func (o AuthenticationNotificationRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AuthenticationNotificationRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	toSerialize["environment"] = o.Environment
	toSerialize["type"] = o.Type
	return toSerialize, nil
}

type NullableAuthenticationNotificationRequest struct {
	value *AuthenticationNotificationRequest
	isSet bool
}

func (v NullableAuthenticationNotificationRequest) Get() *AuthenticationNotificationRequest {
	return v.value
}

func (v *NullableAuthenticationNotificationRequest) Set(val *AuthenticationNotificationRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableAuthenticationNotificationRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableAuthenticationNotificationRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAuthenticationNotificationRequest(val *AuthenticationNotificationRequest) *NullableAuthenticationNotificationRequest {
	return &NullableAuthenticationNotificationRequest{value: val, isSet: true}
}

func (v NullableAuthenticationNotificationRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAuthenticationNotificationRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

func (o *AuthenticationNotificationRequest) isValidType() bool {
	var allowedEnumValues = []string{"balancePlatform.authentication.created"}
	for _, allowed := range allowedEnumValues {
		if o.GetType() == allowed {
			return true
		}
	}
	return false
}