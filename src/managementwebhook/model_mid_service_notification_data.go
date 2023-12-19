/*
Management Webhooks

API version: 3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package managementwebhook

import (
	"encoding/json"

	"github.com/adyen/adyen-go-api-library/v8/src/common"
)

// checks if the MidServiceNotificationData type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &MidServiceNotificationData{}

// MidServiceNotificationData struct for MidServiceNotificationData
type MidServiceNotificationData struct {
	// Indicates whether receiving payments is allowed. This value is set to **true** by Adyen after screening your merchant account.
	Allowed *bool `json:"allowed,omitempty"`
	// Indicates whether the payment method is enabled (**true**) or disabled (**false**).
	Enabled *bool `json:"enabled,omitempty"`
	// The unique identifier of the resource.
	Id string `json:"id"`
	// The unique identifier of the merchant account.
	MerchantId string `json:"merchantId"`
	// Your reference for the payment method.
	Reference *string `json:"reference,omitempty"`
	// The status of the request to add a payment method. Possible values:  * **success**: the payment method was added. * **failure**: the request failed. * **capabilityPending**: the **receivePayments** capability is not allowed.
	Status string `json:"status"`
	// The unique identifier of the [store](https://docs.adyen.com/api-explorer/#/ManagementService/latest/post/merchants/{id}/paymentMethodSettings__reqParam_storeId), if any.
	StoreId *string `json:"storeId,omitempty"`
	// Payment method [variant](https://docs.adyen.com/development-resources/paymentmethodvariant#management-api).
	Type string `json:"type"`
	// Payment method status. Possible values: * **valid** * **pending** * **invalid** * **rejected**
	VerificationStatus *string `json:"verificationStatus,omitempty"`
}

// NewMidServiceNotificationData instantiates a new MidServiceNotificationData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMidServiceNotificationData(id string, merchantId string, status string, type_ string) *MidServiceNotificationData {
	this := MidServiceNotificationData{}
	this.Id = id
	this.MerchantId = merchantId
	this.Status = status
	this.Type = type_
	return &this
}

// NewMidServiceNotificationDataWithDefaults instantiates a new MidServiceNotificationData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMidServiceNotificationDataWithDefaults() *MidServiceNotificationData {
	this := MidServiceNotificationData{}
	return &this
}

// GetAllowed returns the Allowed field value if set, zero value otherwise.
func (o *MidServiceNotificationData) GetAllowed() bool {
	if o == nil || common.IsNil(o.Allowed) {
		var ret bool
		return ret
	}
	return *o.Allowed
}

// GetAllowedOk returns a tuple with the Allowed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MidServiceNotificationData) GetAllowedOk() (*bool, bool) {
	if o == nil || common.IsNil(o.Allowed) {
		return nil, false
	}
	return o.Allowed, true
}

// HasAllowed returns a boolean if a field has been set.
func (o *MidServiceNotificationData) HasAllowed() bool {
	if o != nil && !common.IsNil(o.Allowed) {
		return true
	}

	return false
}

// SetAllowed gets a reference to the given bool and assigns it to the Allowed field.
func (o *MidServiceNotificationData) SetAllowed(v bool) {
	o.Allowed = &v
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *MidServiceNotificationData) GetEnabled() bool {
	if o == nil || common.IsNil(o.Enabled) {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MidServiceNotificationData) GetEnabledOk() (*bool, bool) {
	if o == nil || common.IsNil(o.Enabled) {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *MidServiceNotificationData) HasEnabled() bool {
	if o != nil && !common.IsNil(o.Enabled) {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *MidServiceNotificationData) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetId returns the Id field value
func (o *MidServiceNotificationData) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *MidServiceNotificationData) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *MidServiceNotificationData) SetId(v string) {
	o.Id = v
}

// GetMerchantId returns the MerchantId field value
func (o *MidServiceNotificationData) GetMerchantId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.MerchantId
}

// GetMerchantIdOk returns a tuple with the MerchantId field value
// and a boolean to check if the value has been set.
func (o *MidServiceNotificationData) GetMerchantIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MerchantId, true
}

// SetMerchantId sets field value
func (o *MidServiceNotificationData) SetMerchantId(v string) {
	o.MerchantId = v
}

// GetReference returns the Reference field value if set, zero value otherwise.
func (o *MidServiceNotificationData) GetReference() string {
	if o == nil || common.IsNil(o.Reference) {
		var ret string
		return ret
	}
	return *o.Reference
}

// GetReferenceOk returns a tuple with the Reference field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MidServiceNotificationData) GetReferenceOk() (*string, bool) {
	if o == nil || common.IsNil(o.Reference) {
		return nil, false
	}
	return o.Reference, true
}

// HasReference returns a boolean if a field has been set.
func (o *MidServiceNotificationData) HasReference() bool {
	if o != nil && !common.IsNil(o.Reference) {
		return true
	}

	return false
}

// SetReference gets a reference to the given string and assigns it to the Reference field.
func (o *MidServiceNotificationData) SetReference(v string) {
	o.Reference = &v
}

// GetStatus returns the Status field value
func (o *MidServiceNotificationData) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *MidServiceNotificationData) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *MidServiceNotificationData) SetStatus(v string) {
	o.Status = v
}

// GetStoreId returns the StoreId field value if set, zero value otherwise.
func (o *MidServiceNotificationData) GetStoreId() string {
	if o == nil || common.IsNil(o.StoreId) {
		var ret string
		return ret
	}
	return *o.StoreId
}

// GetStoreIdOk returns a tuple with the StoreId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MidServiceNotificationData) GetStoreIdOk() (*string, bool) {
	if o == nil || common.IsNil(o.StoreId) {
		return nil, false
	}
	return o.StoreId, true
}

// HasStoreId returns a boolean if a field has been set.
func (o *MidServiceNotificationData) HasStoreId() bool {
	if o != nil && !common.IsNil(o.StoreId) {
		return true
	}

	return false
}

// SetStoreId gets a reference to the given string and assigns it to the StoreId field.
func (o *MidServiceNotificationData) SetStoreId(v string) {
	o.StoreId = &v
}

// GetType returns the Type field value
func (o *MidServiceNotificationData) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *MidServiceNotificationData) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *MidServiceNotificationData) SetType(v string) {
	o.Type = v
}

// GetVerificationStatus returns the VerificationStatus field value if set, zero value otherwise.
func (o *MidServiceNotificationData) GetVerificationStatus() string {
	if o == nil || common.IsNil(o.VerificationStatus) {
		var ret string
		return ret
	}
	return *o.VerificationStatus
}

// GetVerificationStatusOk returns a tuple with the VerificationStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MidServiceNotificationData) GetVerificationStatusOk() (*string, bool) {
	if o == nil || common.IsNil(o.VerificationStatus) {
		return nil, false
	}
	return o.VerificationStatus, true
}

// HasVerificationStatus returns a boolean if a field has been set.
func (o *MidServiceNotificationData) HasVerificationStatus() bool {
	if o != nil && !common.IsNil(o.VerificationStatus) {
		return true
	}

	return false
}

// SetVerificationStatus gets a reference to the given string and assigns it to the VerificationStatus field.
func (o *MidServiceNotificationData) SetVerificationStatus(v string) {
	o.VerificationStatus = &v
}

func (o MidServiceNotificationData) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MidServiceNotificationData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.Allowed) {
		toSerialize["allowed"] = o.Allowed
	}
	if !common.IsNil(o.Enabled) {
		toSerialize["enabled"] = o.Enabled
	}
	toSerialize["id"] = o.Id
	toSerialize["merchantId"] = o.MerchantId
	if !common.IsNil(o.Reference) {
		toSerialize["reference"] = o.Reference
	}
	toSerialize["status"] = o.Status
	if !common.IsNil(o.StoreId) {
		toSerialize["storeId"] = o.StoreId
	}
	toSerialize["type"] = o.Type
	if !common.IsNil(o.VerificationStatus) {
		toSerialize["verificationStatus"] = o.VerificationStatus
	}
	return toSerialize, nil
}

type NullableMidServiceNotificationData struct {
	value *MidServiceNotificationData
	isSet bool
}

func (v NullableMidServiceNotificationData) Get() *MidServiceNotificationData {
	return v.value
}

func (v *NullableMidServiceNotificationData) Set(val *MidServiceNotificationData) {
	v.value = val
	v.isSet = true
}

func (v NullableMidServiceNotificationData) IsSet() bool {
	return v.isSet
}

func (v *NullableMidServiceNotificationData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMidServiceNotificationData(val *MidServiceNotificationData) *NullableMidServiceNotificationData {
	return &NullableMidServiceNotificationData{value: val, isSet: true}
}

func (v NullableMidServiceNotificationData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMidServiceNotificationData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

func (o *MidServiceNotificationData) isValidStatus() bool {
	var allowedEnumValues = []string{"success", "failure", "capabilityPending", "dataRequired", "updatesExpected"}
	for _, allowed := range allowedEnumValues {
		if o.GetStatus() == allowed {
			return true
		}
	}
	return false
}
func (o *MidServiceNotificationData) isValidVerificationStatus() bool {
	var allowedEnumValues = []string{"valid", "pending", "invalid", "rejected"}
	for _, allowed := range allowedEnumValues {
		if o.GetVerificationStatus() == allowed {
			return true
		}
	}
	return false
}
