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

// KYCCheckStatusData struct for KYCCheckStatusData
type KYCCheckStatusData struct {
	// A list of the fields required for execution of the check.
	RequiredFields *[]string `json:"requiredFields,omitempty"`
	// The status of the check. >Permitted Values: `DATA_PROVIDED`, `PASSED`, `PENDING`, `AWAITING_DATA`, `RETRY_LIMIT_REACHED`, `INVALID_DATA`, `FAILED`.
	Status  string           `json:"status"`
	Summary *KYCCheckSummary `json:"summary,omitempty"`
	// The type of check. >Permitted Values: `COMPANY_VERIFICATION`, `IDENTITY_VERIFICATION`, `PASSPORT_VERIFICATION`, `BANK_ACCOUNT_VERIFICATION`, `NONPROFIT_VERIFICATION`, `CARD_VERIFICATION`.
	Type string `json:"type"`
}

// NewKYCCheckStatusData instantiates a new KYCCheckStatusData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKYCCheckStatusData(status string, type_ string) *KYCCheckStatusData {
	this := KYCCheckStatusData{}
	this.Status = status
	this.Type = type_
	return &this
}

// NewKYCCheckStatusDataWithDefaults instantiates a new KYCCheckStatusData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKYCCheckStatusDataWithDefaults() *KYCCheckStatusData {
	this := KYCCheckStatusData{}
	return &this
}

// GetRequiredFields returns the RequiredFields field value if set, zero value otherwise.
func (o *KYCCheckStatusData) GetRequiredFields() []string {
	if o == nil || o.RequiredFields == nil {
		var ret []string
		return ret
	}
	return *o.RequiredFields
}

// GetRequiredFieldsOk returns a tuple with the RequiredFields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KYCCheckStatusData) GetRequiredFieldsOk() (*[]string, bool) {
	if o == nil || o.RequiredFields == nil {
		return nil, false
	}
	return o.RequiredFields, true
}

// HasRequiredFields returns a boolean if a field has been set.
func (o *KYCCheckStatusData) HasRequiredFields() bool {
	if o != nil && o.RequiredFields != nil {
		return true
	}

	return false
}

// SetRequiredFields gets a reference to the given []string and assigns it to the RequiredFields field.
func (o *KYCCheckStatusData) SetRequiredFields(v []string) {
	o.RequiredFields = &v
}

// GetStatus returns the Status field value
func (o *KYCCheckStatusData) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *KYCCheckStatusData) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *KYCCheckStatusData) SetStatus(v string) {
	o.Status = v
}

// GetSummary returns the Summary field value if set, zero value otherwise.
func (o *KYCCheckStatusData) GetSummary() KYCCheckSummary {
	if o == nil || o.Summary == nil {
		var ret KYCCheckSummary
		return ret
	}
	return *o.Summary
}

// GetSummaryOk returns a tuple with the Summary field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KYCCheckStatusData) GetSummaryOk() (*KYCCheckSummary, bool) {
	if o == nil || o.Summary == nil {
		return nil, false
	}
	return o.Summary, true
}

// HasSummary returns a boolean if a field has been set.
func (o *KYCCheckStatusData) HasSummary() bool {
	if o != nil && o.Summary != nil {
		return true
	}

	return false
}

// SetSummary gets a reference to the given KYCCheckSummary and assigns it to the Summary field.
func (o *KYCCheckStatusData) SetSummary(v KYCCheckSummary) {
	o.Summary = &v
}

// GetType returns the Type field value
func (o *KYCCheckStatusData) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *KYCCheckStatusData) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *KYCCheckStatusData) SetType(v string) {
	o.Type = v
}

func (o KYCCheckStatusData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.RequiredFields != nil {
		toSerialize["requiredFields"] = o.RequiredFields
	}
	if true {
		toSerialize["status"] = o.Status
	}
	if o.Summary != nil {
		toSerialize["summary"] = o.Summary
	}
	if true {
		toSerialize["type"] = o.Type
	}
	return json.Marshal(toSerialize)
}

type NullableKYCCheckStatusData struct {
	value *KYCCheckStatusData
	isSet bool
}

func (v NullableKYCCheckStatusData) Get() *KYCCheckStatusData {
	return v.value
}

func (v *NullableKYCCheckStatusData) Set(val *KYCCheckStatusData) {
	v.value = val
	v.isSet = true
}

func (v NullableKYCCheckStatusData) IsSet() bool {
	return v.isSet
}

func (v *NullableKYCCheckStatusData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKYCCheckStatusData(val *KYCCheckStatusData) *NullableKYCCheckStatusData {
	return &NullableKYCCheckStatusData{value: val, isSet: true}
}

func (v NullableKYCCheckStatusData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKYCCheckStatusData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
