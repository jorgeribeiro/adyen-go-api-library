/*
Management Webhooks

API version: 1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package managementwebhook

import (
	"encoding/json"

	"github.com/adyen/adyen-go-api-library/v7/src/common"
)

// checks if the AccountUpdateNotificationData type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &AccountUpdateNotificationData{}

// AccountUpdateNotificationData struct for AccountUpdateNotificationData
type AccountUpdateNotificationData struct {
	// Key-value pairs that specify what you can do with the merchant account and its settings. The key is a capability. For example, the **sendToTransferInstrument** is the capability required before you can pay out the funds of a merchant account to a [bank account](https://docs.adyen.com/api-explorer/legalentity/latest/post/transferInstruments). The value is an object containing the settings for the capability.
	Capabilities map[string]AccountCapabilityData `json:"capabilities"`
	// The unique identifier of the [legal entity](https://docs.adyen.com/api-explorer/legalentity/latest/post/legalEntities#responses-200-id).
	LegalEntityId *string `json:"legalEntityId,omitempty"`
	// The unique identifier of the merchant account.
	MerchantId string `json:"merchantId"`
	// The status of the merchant account.  Possible values:  * **PreActive**: The merchant account has been created. Users cannot access the merchant account in the Customer Area. The account cannot process payments. * **Active**: Users can access the merchant account in the Customer Area. If the company account is also **Active**, then payment processing and payouts are enabled. * **InactiveWithModifications**: Users can access the merchant account in the Customer Area. The account cannot process new payments but can still modify payments, for example issue refunds. The account can still receive payouts. * **Inactive**: Users can access the merchant account in the Customer Area. Payment processing and payouts are disabled. * **Closed**: The account is closed and this cannot be reversed. Users cannot log in. Payment processing and payouts are disabled.
	Status string `json:"status"`
}

// NewAccountUpdateNotificationData instantiates a new AccountUpdateNotificationData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAccountUpdateNotificationData(capabilities map[string]AccountCapabilityData, merchantId string, status string) *AccountUpdateNotificationData {
	this := AccountUpdateNotificationData{}
	this.Capabilities = capabilities
	this.MerchantId = merchantId
	this.Status = status
	return &this
}

// NewAccountUpdateNotificationDataWithDefaults instantiates a new AccountUpdateNotificationData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAccountUpdateNotificationDataWithDefaults() *AccountUpdateNotificationData {
	this := AccountUpdateNotificationData{}
	return &this
}

// GetCapabilities returns the Capabilities field value
func (o *AccountUpdateNotificationData) GetCapabilities() map[string]AccountCapabilityData {
	if o == nil {
		var ret map[string]AccountCapabilityData
		return ret
	}

	return o.Capabilities
}

// GetCapabilitiesOk returns a tuple with the Capabilities field value
// and a boolean to check if the value has been set.
func (o *AccountUpdateNotificationData) GetCapabilitiesOk() (*map[string]AccountCapabilityData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Capabilities, true
}

// SetCapabilities sets field value
func (o *AccountUpdateNotificationData) SetCapabilities(v map[string]AccountCapabilityData) {
	o.Capabilities = v
}

// GetLegalEntityId returns the LegalEntityId field value if set, zero value otherwise.
func (o *AccountUpdateNotificationData) GetLegalEntityId() string {
	if o == nil || common.IsNil(o.LegalEntityId) {
		var ret string
		return ret
	}
	return *o.LegalEntityId
}

// GetLegalEntityIdOk returns a tuple with the LegalEntityId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountUpdateNotificationData) GetLegalEntityIdOk() (*string, bool) {
	if o == nil || common.IsNil(o.LegalEntityId) {
		return nil, false
	}
	return o.LegalEntityId, true
}

// HasLegalEntityId returns a boolean if a field has been set.
func (o *AccountUpdateNotificationData) HasLegalEntityId() bool {
	if o != nil && !common.IsNil(o.LegalEntityId) {
		return true
	}

	return false
}

// SetLegalEntityId gets a reference to the given string and assigns it to the LegalEntityId field.
func (o *AccountUpdateNotificationData) SetLegalEntityId(v string) {
	o.LegalEntityId = &v
}

// GetMerchantId returns the MerchantId field value
func (o *AccountUpdateNotificationData) GetMerchantId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.MerchantId
}

// GetMerchantIdOk returns a tuple with the MerchantId field value
// and a boolean to check if the value has been set.
func (o *AccountUpdateNotificationData) GetMerchantIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MerchantId, true
}

// SetMerchantId sets field value
func (o *AccountUpdateNotificationData) SetMerchantId(v string) {
	o.MerchantId = v
}

// GetStatus returns the Status field value
func (o *AccountUpdateNotificationData) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *AccountUpdateNotificationData) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *AccountUpdateNotificationData) SetStatus(v string) {
	o.Status = v
}

func (o AccountUpdateNotificationData) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AccountUpdateNotificationData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["capabilities"] = o.Capabilities
	if !common.IsNil(o.LegalEntityId) {
		toSerialize["legalEntityId"] = o.LegalEntityId
	}
	toSerialize["merchantId"] = o.MerchantId
	toSerialize["status"] = o.Status
	return toSerialize, nil
}

type NullableAccountUpdateNotificationData struct {
	value *AccountUpdateNotificationData
	isSet bool
}

func (v NullableAccountUpdateNotificationData) Get() *AccountUpdateNotificationData {
	return v.value
}

func (v *NullableAccountUpdateNotificationData) Set(val *AccountUpdateNotificationData) {
	v.value = val
	v.isSet = true
}

func (v NullableAccountUpdateNotificationData) IsSet() bool {
	return v.isSet
}

func (v *NullableAccountUpdateNotificationData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAccountUpdateNotificationData(val *AccountUpdateNotificationData) *NullableAccountUpdateNotificationData {
	return &NullableAccountUpdateNotificationData{value: val, isSet: true}
}

func (v NullableAccountUpdateNotificationData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAccountUpdateNotificationData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
