/*
Management API

API version: 3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package management

import (
	"encoding/json"

	"github.com/adyen/adyen-go-api-library/v8/src/common"
)

// checks if the CreateMerchantApiCredentialRequest type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &CreateMerchantApiCredentialRequest{}

// CreateMerchantApiCredentialRequest struct for CreateMerchantApiCredentialRequest
type CreateMerchantApiCredentialRequest struct {
	// The list of [allowed origins](https://docs.adyen.com/development-resources/client-side-authentication#allowed-origins) for the new API credential.
	AllowedOrigins []string `json:"allowedOrigins,omitempty"`
	// Description of the API credential.
	Description *string `json:"description,omitempty"`
	// List of [roles](https://docs.adyen.com/development-resources/api-credentials#roles-1) for the API credential. Only roles assigned to 'ws@Company.<CompanyName>' can be assigned to other API credentials.
	Roles []string `json:"roles,omitempty"`
}

// NewCreateMerchantApiCredentialRequest instantiates a new CreateMerchantApiCredentialRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateMerchantApiCredentialRequest() *CreateMerchantApiCredentialRequest {
	this := CreateMerchantApiCredentialRequest{}
	return &this
}

// NewCreateMerchantApiCredentialRequestWithDefaults instantiates a new CreateMerchantApiCredentialRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateMerchantApiCredentialRequestWithDefaults() *CreateMerchantApiCredentialRequest {
	this := CreateMerchantApiCredentialRequest{}
	return &this
}

// GetAllowedOrigins returns the AllowedOrigins field value if set, zero value otherwise.
func (o *CreateMerchantApiCredentialRequest) GetAllowedOrigins() []string {
	if o == nil || common.IsNil(o.AllowedOrigins) {
		var ret []string
		return ret
	}
	return o.AllowedOrigins
}

// GetAllowedOriginsOk returns a tuple with the AllowedOrigins field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateMerchantApiCredentialRequest) GetAllowedOriginsOk() ([]string, bool) {
	if o == nil || common.IsNil(o.AllowedOrigins) {
		return nil, false
	}
	return o.AllowedOrigins, true
}

// HasAllowedOrigins returns a boolean if a field has been set.
func (o *CreateMerchantApiCredentialRequest) HasAllowedOrigins() bool {
	if o != nil && !common.IsNil(o.AllowedOrigins) {
		return true
	}

	return false
}

// SetAllowedOrigins gets a reference to the given []string and assigns it to the AllowedOrigins field.
func (o *CreateMerchantApiCredentialRequest) SetAllowedOrigins(v []string) {
	o.AllowedOrigins = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *CreateMerchantApiCredentialRequest) GetDescription() string {
	if o == nil || common.IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateMerchantApiCredentialRequest) GetDescriptionOk() (*string, bool) {
	if o == nil || common.IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *CreateMerchantApiCredentialRequest) HasDescription() bool {
	if o != nil && !common.IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *CreateMerchantApiCredentialRequest) SetDescription(v string) {
	o.Description = &v
}

// GetRoles returns the Roles field value if set, zero value otherwise.
func (o *CreateMerchantApiCredentialRequest) GetRoles() []string {
	if o == nil || common.IsNil(o.Roles) {
		var ret []string
		return ret
	}
	return o.Roles
}

// GetRolesOk returns a tuple with the Roles field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateMerchantApiCredentialRequest) GetRolesOk() ([]string, bool) {
	if o == nil || common.IsNil(o.Roles) {
		return nil, false
	}
	return o.Roles, true
}

// HasRoles returns a boolean if a field has been set.
func (o *CreateMerchantApiCredentialRequest) HasRoles() bool {
	if o != nil && !common.IsNil(o.Roles) {
		return true
	}

	return false
}

// SetRoles gets a reference to the given []string and assigns it to the Roles field.
func (o *CreateMerchantApiCredentialRequest) SetRoles(v []string) {
	o.Roles = v
}

func (o CreateMerchantApiCredentialRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateMerchantApiCredentialRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.AllowedOrigins) {
		toSerialize["allowedOrigins"] = o.AllowedOrigins
	}
	if !common.IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !common.IsNil(o.Roles) {
		toSerialize["roles"] = o.Roles
	}
	return toSerialize, nil
}

type NullableCreateMerchantApiCredentialRequest struct {
	value *CreateMerchantApiCredentialRequest
	isSet bool
}

func (v NullableCreateMerchantApiCredentialRequest) Get() *CreateMerchantApiCredentialRequest {
	return v.value
}

func (v *NullableCreateMerchantApiCredentialRequest) Set(val *CreateMerchantApiCredentialRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateMerchantApiCredentialRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateMerchantApiCredentialRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateMerchantApiCredentialRequest(val *CreateMerchantApiCredentialRequest) *NullableCreateMerchantApiCredentialRequest {
	return &NullableCreateMerchantApiCredentialRequest{value: val, isSet: true}
}

func (v NullableCreateMerchantApiCredentialRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateMerchantApiCredentialRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}