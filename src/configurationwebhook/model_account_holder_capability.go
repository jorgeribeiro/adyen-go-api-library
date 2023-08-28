/*
Configuration webhooks

API version: 1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package configurationwebhook

import (
	"encoding/json"

	"github.com/adyen/adyen-go-api-library/v7/src/common"
)

// checks if the AccountHolderCapability type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &AccountHolderCapability{}

// AccountHolderCapability struct for AccountHolderCapability
type AccountHolderCapability struct {
	// Indicates whether the capability is allowed. Adyen sets this to **true** if the verification is successful and the account holder is permitted to use the capability.
	Allowed *bool `json:"allowed,omitempty"`
	// The capability level that is allowed for the account holder.  Possible values: **notApplicable**, **low**, **medium**, **high**.
	AllowedLevel    *string             `json:"allowedLevel,omitempty"`
	AllowedSettings *CapabilitySettings `json:"allowedSettings,omitempty"`
	// Indicates whether the capability is enabled. If **false**, the capability is temporarily disabled for the account holder.
	Enabled *bool `json:"enabled,omitempty"`
	// Contains verification errors and the actions that you can take to resolve them.
	Problems []CapabilityProblem `json:"problems,omitempty"`
	// Indicates whether the capability is requested. To check whether the account holder is permitted to use the capability, refer to the `allowed` field.
	Requested *bool `json:"requested,omitempty"`
	// The requested level of the capability. Some capabilities, such as those used in [card issuing](https://docs.adyen.com/issuing/add-capabilities#capability-levels), have different levels. Levels increase the capability, but also require additional checks and increased monitoring.  Possible values: **notApplicable**, **low**, **medium**, **high**.
	RequestedLevel    *string             `json:"requestedLevel,omitempty"`
	RequestedSettings *CapabilitySettings `json:"requestedSettings,omitempty"`
	// Contains the status of the transfer instruments associated with this capability.
	TransferInstruments []AccountSupportingEntityCapability `json:"transferInstruments,omitempty"`
	// The status of the verification checks for the capability.  Possible values:  * **pending**: Adyen is running the verification.  * **invalid**: The verification failed. Check if the `errors` array contains more information.  * **valid**: The verification has been successfully completed.  * **rejected**: Adyen has verified the information, but found reasons to not allow the capability.
	VerificationStatus *string `json:"verificationStatus,omitempty"`
}

// NewAccountHolderCapability instantiates a new AccountHolderCapability object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAccountHolderCapability() *AccountHolderCapability {
	this := AccountHolderCapability{}
	return &this
}

// NewAccountHolderCapabilityWithDefaults instantiates a new AccountHolderCapability object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAccountHolderCapabilityWithDefaults() *AccountHolderCapability {
	this := AccountHolderCapability{}
	return &this
}

// GetAllowed returns the Allowed field value if set, zero value otherwise.
func (o *AccountHolderCapability) GetAllowed() bool {
	if o == nil || common.IsNil(o.Allowed) {
		var ret bool
		return ret
	}
	return *o.Allowed
}

// GetAllowedOk returns a tuple with the Allowed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountHolderCapability) GetAllowedOk() (*bool, bool) {
	if o == nil || common.IsNil(o.Allowed) {
		return nil, false
	}
	return o.Allowed, true
}

// HasAllowed returns a boolean if a field has been set.
func (o *AccountHolderCapability) HasAllowed() bool {
	if o != nil && !common.IsNil(o.Allowed) {
		return true
	}

	return false
}

// SetAllowed gets a reference to the given bool and assigns it to the Allowed field.
func (o *AccountHolderCapability) SetAllowed(v bool) {
	o.Allowed = &v
}

// GetAllowedLevel returns the AllowedLevel field value if set, zero value otherwise.
func (o *AccountHolderCapability) GetAllowedLevel() string {
	if o == nil || common.IsNil(o.AllowedLevel) {
		var ret string
		return ret
	}
	return *o.AllowedLevel
}

// GetAllowedLevelOk returns a tuple with the AllowedLevel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountHolderCapability) GetAllowedLevelOk() (*string, bool) {
	if o == nil || common.IsNil(o.AllowedLevel) {
		return nil, false
	}
	return o.AllowedLevel, true
}

// HasAllowedLevel returns a boolean if a field has been set.
func (o *AccountHolderCapability) HasAllowedLevel() bool {
	if o != nil && !common.IsNil(o.AllowedLevel) {
		return true
	}

	return false
}

// SetAllowedLevel gets a reference to the given string and assigns it to the AllowedLevel field.
func (o *AccountHolderCapability) SetAllowedLevel(v string) {
	o.AllowedLevel = &v
}

// GetAllowedSettings returns the AllowedSettings field value if set, zero value otherwise.
func (o *AccountHolderCapability) GetAllowedSettings() CapabilitySettings {
	if o == nil || common.IsNil(o.AllowedSettings) {
		var ret CapabilitySettings
		return ret
	}
	return *o.AllowedSettings
}

// GetAllowedSettingsOk returns a tuple with the AllowedSettings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountHolderCapability) GetAllowedSettingsOk() (*CapabilitySettings, bool) {
	if o == nil || common.IsNil(o.AllowedSettings) {
		return nil, false
	}
	return o.AllowedSettings, true
}

// HasAllowedSettings returns a boolean if a field has been set.
func (o *AccountHolderCapability) HasAllowedSettings() bool {
	if o != nil && !common.IsNil(o.AllowedSettings) {
		return true
	}

	return false
}

// SetAllowedSettings gets a reference to the given CapabilitySettings and assigns it to the AllowedSettings field.
func (o *AccountHolderCapability) SetAllowedSettings(v CapabilitySettings) {
	o.AllowedSettings = &v
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *AccountHolderCapability) GetEnabled() bool {
	if o == nil || common.IsNil(o.Enabled) {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountHolderCapability) GetEnabledOk() (*bool, bool) {
	if o == nil || common.IsNil(o.Enabled) {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *AccountHolderCapability) HasEnabled() bool {
	if o != nil && !common.IsNil(o.Enabled) {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *AccountHolderCapability) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetProblems returns the Problems field value if set, zero value otherwise.
func (o *AccountHolderCapability) GetProblems() []CapabilityProblem {
	if o == nil || common.IsNil(o.Problems) {
		var ret []CapabilityProblem
		return ret
	}
	return o.Problems
}

// GetProblemsOk returns a tuple with the Problems field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountHolderCapability) GetProblemsOk() ([]CapabilityProblem, bool) {
	if o == nil || common.IsNil(o.Problems) {
		return nil, false
	}
	return o.Problems, true
}

// HasProblems returns a boolean if a field has been set.
func (o *AccountHolderCapability) HasProblems() bool {
	if o != nil && !common.IsNil(o.Problems) {
		return true
	}

	return false
}

// SetProblems gets a reference to the given []CapabilityProblem and assigns it to the Problems field.
func (o *AccountHolderCapability) SetProblems(v []CapabilityProblem) {
	o.Problems = v
}

// GetRequested returns the Requested field value if set, zero value otherwise.
func (o *AccountHolderCapability) GetRequested() bool {
	if o == nil || common.IsNil(o.Requested) {
		var ret bool
		return ret
	}
	return *o.Requested
}

// GetRequestedOk returns a tuple with the Requested field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountHolderCapability) GetRequestedOk() (*bool, bool) {
	if o == nil || common.IsNil(o.Requested) {
		return nil, false
	}
	return o.Requested, true
}

// HasRequested returns a boolean if a field has been set.
func (o *AccountHolderCapability) HasRequested() bool {
	if o != nil && !common.IsNil(o.Requested) {
		return true
	}

	return false
}

// SetRequested gets a reference to the given bool and assigns it to the Requested field.
func (o *AccountHolderCapability) SetRequested(v bool) {
	o.Requested = &v
}

// GetRequestedLevel returns the RequestedLevel field value if set, zero value otherwise.
func (o *AccountHolderCapability) GetRequestedLevel() string {
	if o == nil || common.IsNil(o.RequestedLevel) {
		var ret string
		return ret
	}
	return *o.RequestedLevel
}

// GetRequestedLevelOk returns a tuple with the RequestedLevel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountHolderCapability) GetRequestedLevelOk() (*string, bool) {
	if o == nil || common.IsNil(o.RequestedLevel) {
		return nil, false
	}
	return o.RequestedLevel, true
}

// HasRequestedLevel returns a boolean if a field has been set.
func (o *AccountHolderCapability) HasRequestedLevel() bool {
	if o != nil && !common.IsNil(o.RequestedLevel) {
		return true
	}

	return false
}

// SetRequestedLevel gets a reference to the given string and assigns it to the RequestedLevel field.
func (o *AccountHolderCapability) SetRequestedLevel(v string) {
	o.RequestedLevel = &v
}

// GetRequestedSettings returns the RequestedSettings field value if set, zero value otherwise.
func (o *AccountHolderCapability) GetRequestedSettings() CapabilitySettings {
	if o == nil || common.IsNil(o.RequestedSettings) {
		var ret CapabilitySettings
		return ret
	}
	return *o.RequestedSettings
}

// GetRequestedSettingsOk returns a tuple with the RequestedSettings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountHolderCapability) GetRequestedSettingsOk() (*CapabilitySettings, bool) {
	if o == nil || common.IsNil(o.RequestedSettings) {
		return nil, false
	}
	return o.RequestedSettings, true
}

// HasRequestedSettings returns a boolean if a field has been set.
func (o *AccountHolderCapability) HasRequestedSettings() bool {
	if o != nil && !common.IsNil(o.RequestedSettings) {
		return true
	}

	return false
}

// SetRequestedSettings gets a reference to the given CapabilitySettings and assigns it to the RequestedSettings field.
func (o *AccountHolderCapability) SetRequestedSettings(v CapabilitySettings) {
	o.RequestedSettings = &v
}

// GetTransferInstruments returns the TransferInstruments field value if set, zero value otherwise.
func (o *AccountHolderCapability) GetTransferInstruments() []AccountSupportingEntityCapability {
	if o == nil || common.IsNil(o.TransferInstruments) {
		var ret []AccountSupportingEntityCapability
		return ret
	}
	return o.TransferInstruments
}

// GetTransferInstrumentsOk returns a tuple with the TransferInstruments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountHolderCapability) GetTransferInstrumentsOk() ([]AccountSupportingEntityCapability, bool) {
	if o == nil || common.IsNil(o.TransferInstruments) {
		return nil, false
	}
	return o.TransferInstruments, true
}

// HasTransferInstruments returns a boolean if a field has been set.
func (o *AccountHolderCapability) HasTransferInstruments() bool {
	if o != nil && !common.IsNil(o.TransferInstruments) {
		return true
	}

	return false
}

// SetTransferInstruments gets a reference to the given []AccountSupportingEntityCapability and assigns it to the TransferInstruments field.
func (o *AccountHolderCapability) SetTransferInstruments(v []AccountSupportingEntityCapability) {
	o.TransferInstruments = v
}

// GetVerificationStatus returns the VerificationStatus field value if set, zero value otherwise.
func (o *AccountHolderCapability) GetVerificationStatus() string {
	if o == nil || common.IsNil(o.VerificationStatus) {
		var ret string
		return ret
	}
	return *o.VerificationStatus
}

// GetVerificationStatusOk returns a tuple with the VerificationStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountHolderCapability) GetVerificationStatusOk() (*string, bool) {
	if o == nil || common.IsNil(o.VerificationStatus) {
		return nil, false
	}
	return o.VerificationStatus, true
}

// HasVerificationStatus returns a boolean if a field has been set.
func (o *AccountHolderCapability) HasVerificationStatus() bool {
	if o != nil && !common.IsNil(o.VerificationStatus) {
		return true
	}

	return false
}

// SetVerificationStatus gets a reference to the given string and assigns it to the VerificationStatus field.
func (o *AccountHolderCapability) SetVerificationStatus(v string) {
	o.VerificationStatus = &v
}

func (o AccountHolderCapability) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AccountHolderCapability) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.Allowed) {
		toSerialize["allowed"] = o.Allowed
	}
	if !common.IsNil(o.AllowedLevel) {
		toSerialize["allowedLevel"] = o.AllowedLevel
	}
	if !common.IsNil(o.AllowedSettings) {
		toSerialize["allowedSettings"] = o.AllowedSettings
	}
	if !common.IsNil(o.Enabled) {
		toSerialize["enabled"] = o.Enabled
	}
	if !common.IsNil(o.Problems) {
		toSerialize["problems"] = o.Problems
	}
	if !common.IsNil(o.Requested) {
		toSerialize["requested"] = o.Requested
	}
	if !common.IsNil(o.RequestedLevel) {
		toSerialize["requestedLevel"] = o.RequestedLevel
	}
	if !common.IsNil(o.RequestedSettings) {
		toSerialize["requestedSettings"] = o.RequestedSettings
	}
	if !common.IsNil(o.TransferInstruments) {
		toSerialize["transferInstruments"] = o.TransferInstruments
	}
	if !common.IsNil(o.VerificationStatus) {
		toSerialize["verificationStatus"] = o.VerificationStatus
	}
	return toSerialize, nil
}

type NullableAccountHolderCapability struct {
	value *AccountHolderCapability
	isSet bool
}

func (v NullableAccountHolderCapability) Get() *AccountHolderCapability {
	return v.value
}

func (v *NullableAccountHolderCapability) Set(val *AccountHolderCapability) {
	v.value = val
	v.isSet = true
}

func (v NullableAccountHolderCapability) IsSet() bool {
	return v.isSet
}

func (v *NullableAccountHolderCapability) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAccountHolderCapability(val *AccountHolderCapability) *NullableAccountHolderCapability {
	return &NullableAccountHolderCapability{value: val, isSet: true}
}

func (v NullableAccountHolderCapability) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAccountHolderCapability) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

func (o *AccountHolderCapability) isValidAllowedLevel() bool {
	var allowedEnumValues = []string{"high", "low", "medium", "notApplicable"}
	for _, allowed := range allowedEnumValues {
		if o.GetAllowedLevel() == allowed {
			return true
		}
	}
	return false
}
func (o *AccountHolderCapability) isValidRequestedLevel() bool {
	var allowedEnumValues = []string{"high", "low", "medium", "notApplicable"}
	for _, allowed := range allowedEnumValues {
		if o.GetRequestedLevel() == allowed {
			return true
		}
	}
	return false
}
func (o *AccountHolderCapability) isValidVerificationStatus() bool {
	var allowedEnumValues = []string{"invalid", "pending", "rejected", "valid"}
	for _, allowed := range allowedEnumValues {
		if o.GetVerificationStatus() == allowed {
			return true
		}
	}
	return false
}
