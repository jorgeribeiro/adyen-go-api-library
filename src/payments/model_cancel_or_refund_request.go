/*
Adyen Payment API

API version: 68
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package payments

import (
	"encoding/json"

	"github.com/adyen/adyen-go-api-library/v16/src/common"
)

// checks if the CancelOrRefundRequest type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &CancelOrRefundRequest{}

// CancelOrRefundRequest struct for CancelOrRefundRequest
type CancelOrRefundRequest struct {
	// This field contains additional data, which may be required for a particular modification request.  The additionalData object consists of entries, each of which includes the key and value.
	AdditionalData *map[string]string `json:"additionalData,omitempty"`
	// The merchant account that is used to process the payment.
	MerchantAccount string            `json:"merchantAccount"`
	MpiData         *ThreeDSecureData `json:"mpiData,omitempty"`
	// The original merchant reference to cancel.
	OriginalMerchantReference *string `json:"originalMerchantReference,omitempty"`
	// The original pspReference of the payment to modify. This reference is returned in: * authorisation response * authorisation notification
	OriginalReference       string                   `json:"originalReference"`
	PlatformChargebackLogic *PlatformChargebackLogic `json:"platformChargebackLogic,omitempty"`
	// Your reference for the payment modification. This reference is visible in Customer Area and in reports. Maximum length: 80 characters.
	Reference *string `json:"reference,omitempty"`
	// The transaction reference provided by the PED. For point-of-sale integrations only.
	TenderReference *string `json:"tenderReference,omitempty"`
	// Unique terminal ID for the PED that originally processed the request. For point-of-sale integrations only.
	UniqueTerminalId *string `json:"uniqueTerminalId,omitempty"`
}

// NewCancelOrRefundRequest instantiates a new CancelOrRefundRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCancelOrRefundRequest(merchantAccount string, originalReference string) *CancelOrRefundRequest {
	this := CancelOrRefundRequest{}
	this.MerchantAccount = merchantAccount
	this.OriginalReference = originalReference
	return &this
}

// NewCancelOrRefundRequestWithDefaults instantiates a new CancelOrRefundRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCancelOrRefundRequestWithDefaults() *CancelOrRefundRequest {
	this := CancelOrRefundRequest{}
	return &this
}

// GetAdditionalData returns the AdditionalData field value if set, zero value otherwise.
func (o *CancelOrRefundRequest) GetAdditionalData() map[string]string {
	if o == nil || common.IsNil(o.AdditionalData) {
		var ret map[string]string
		return ret
	}
	return *o.AdditionalData
}

// GetAdditionalDataOk returns a tuple with the AdditionalData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CancelOrRefundRequest) GetAdditionalDataOk() (*map[string]string, bool) {
	if o == nil || common.IsNil(o.AdditionalData) {
		return nil, false
	}
	return o.AdditionalData, true
}

// HasAdditionalData returns a boolean if a field has been set.
func (o *CancelOrRefundRequest) HasAdditionalData() bool {
	if o != nil && !common.IsNil(o.AdditionalData) {
		return true
	}

	return false
}

// SetAdditionalData gets a reference to the given map[string]string and assigns it to the AdditionalData field.
func (o *CancelOrRefundRequest) SetAdditionalData(v map[string]string) {
	o.AdditionalData = &v
}

// GetMerchantAccount returns the MerchantAccount field value
func (o *CancelOrRefundRequest) GetMerchantAccount() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.MerchantAccount
}

// GetMerchantAccountOk returns a tuple with the MerchantAccount field value
// and a boolean to check if the value has been set.
func (o *CancelOrRefundRequest) GetMerchantAccountOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MerchantAccount, true
}

// SetMerchantAccount sets field value
func (o *CancelOrRefundRequest) SetMerchantAccount(v string) {
	o.MerchantAccount = v
}

// GetMpiData returns the MpiData field value if set, zero value otherwise.
func (o *CancelOrRefundRequest) GetMpiData() ThreeDSecureData {
	if o == nil || common.IsNil(o.MpiData) {
		var ret ThreeDSecureData
		return ret
	}
	return *o.MpiData
}

// GetMpiDataOk returns a tuple with the MpiData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CancelOrRefundRequest) GetMpiDataOk() (*ThreeDSecureData, bool) {
	if o == nil || common.IsNil(o.MpiData) {
		return nil, false
	}
	return o.MpiData, true
}

// HasMpiData returns a boolean if a field has been set.
func (o *CancelOrRefundRequest) HasMpiData() bool {
	if o != nil && !common.IsNil(o.MpiData) {
		return true
	}

	return false
}

// SetMpiData gets a reference to the given ThreeDSecureData and assigns it to the MpiData field.
func (o *CancelOrRefundRequest) SetMpiData(v ThreeDSecureData) {
	o.MpiData = &v
}

// GetOriginalMerchantReference returns the OriginalMerchantReference field value if set, zero value otherwise.
func (o *CancelOrRefundRequest) GetOriginalMerchantReference() string {
	if o == nil || common.IsNil(o.OriginalMerchantReference) {
		var ret string
		return ret
	}
	return *o.OriginalMerchantReference
}

// GetOriginalMerchantReferenceOk returns a tuple with the OriginalMerchantReference field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CancelOrRefundRequest) GetOriginalMerchantReferenceOk() (*string, bool) {
	if o == nil || common.IsNil(o.OriginalMerchantReference) {
		return nil, false
	}
	return o.OriginalMerchantReference, true
}

// HasOriginalMerchantReference returns a boolean if a field has been set.
func (o *CancelOrRefundRequest) HasOriginalMerchantReference() bool {
	if o != nil && !common.IsNil(o.OriginalMerchantReference) {
		return true
	}

	return false
}

// SetOriginalMerchantReference gets a reference to the given string and assigns it to the OriginalMerchantReference field.
func (o *CancelOrRefundRequest) SetOriginalMerchantReference(v string) {
	o.OriginalMerchantReference = &v
}

// GetOriginalReference returns the OriginalReference field value
func (o *CancelOrRefundRequest) GetOriginalReference() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.OriginalReference
}

// GetOriginalReferenceOk returns a tuple with the OriginalReference field value
// and a boolean to check if the value has been set.
func (o *CancelOrRefundRequest) GetOriginalReferenceOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OriginalReference, true
}

// SetOriginalReference sets field value
func (o *CancelOrRefundRequest) SetOriginalReference(v string) {
	o.OriginalReference = v
}

// GetPlatformChargebackLogic returns the PlatformChargebackLogic field value if set, zero value otherwise.
func (o *CancelOrRefundRequest) GetPlatformChargebackLogic() PlatformChargebackLogic {
	if o == nil || common.IsNil(o.PlatformChargebackLogic) {
		var ret PlatformChargebackLogic
		return ret
	}
	return *o.PlatformChargebackLogic
}

// GetPlatformChargebackLogicOk returns a tuple with the PlatformChargebackLogic field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CancelOrRefundRequest) GetPlatformChargebackLogicOk() (*PlatformChargebackLogic, bool) {
	if o == nil || common.IsNil(o.PlatformChargebackLogic) {
		return nil, false
	}
	return o.PlatformChargebackLogic, true
}

// HasPlatformChargebackLogic returns a boolean if a field has been set.
func (o *CancelOrRefundRequest) HasPlatformChargebackLogic() bool {
	if o != nil && !common.IsNil(o.PlatformChargebackLogic) {
		return true
	}

	return false
}

// SetPlatformChargebackLogic gets a reference to the given PlatformChargebackLogic and assigns it to the PlatformChargebackLogic field.
func (o *CancelOrRefundRequest) SetPlatformChargebackLogic(v PlatformChargebackLogic) {
	o.PlatformChargebackLogic = &v
}

// GetReference returns the Reference field value if set, zero value otherwise.
func (o *CancelOrRefundRequest) GetReference() string {
	if o == nil || common.IsNil(o.Reference) {
		var ret string
		return ret
	}
	return *o.Reference
}

// GetReferenceOk returns a tuple with the Reference field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CancelOrRefundRequest) GetReferenceOk() (*string, bool) {
	if o == nil || common.IsNil(o.Reference) {
		return nil, false
	}
	return o.Reference, true
}

// HasReference returns a boolean if a field has been set.
func (o *CancelOrRefundRequest) HasReference() bool {
	if o != nil && !common.IsNil(o.Reference) {
		return true
	}

	return false
}

// SetReference gets a reference to the given string and assigns it to the Reference field.
func (o *CancelOrRefundRequest) SetReference(v string) {
	o.Reference = &v
}

// GetTenderReference returns the TenderReference field value if set, zero value otherwise.
func (o *CancelOrRefundRequest) GetTenderReference() string {
	if o == nil || common.IsNil(o.TenderReference) {
		var ret string
		return ret
	}
	return *o.TenderReference
}

// GetTenderReferenceOk returns a tuple with the TenderReference field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CancelOrRefundRequest) GetTenderReferenceOk() (*string, bool) {
	if o == nil || common.IsNil(o.TenderReference) {
		return nil, false
	}
	return o.TenderReference, true
}

// HasTenderReference returns a boolean if a field has been set.
func (o *CancelOrRefundRequest) HasTenderReference() bool {
	if o != nil && !common.IsNil(o.TenderReference) {
		return true
	}

	return false
}

// SetTenderReference gets a reference to the given string and assigns it to the TenderReference field.
func (o *CancelOrRefundRequest) SetTenderReference(v string) {
	o.TenderReference = &v
}

// GetUniqueTerminalId returns the UniqueTerminalId field value if set, zero value otherwise.
func (o *CancelOrRefundRequest) GetUniqueTerminalId() string {
	if o == nil || common.IsNil(o.UniqueTerminalId) {
		var ret string
		return ret
	}
	return *o.UniqueTerminalId
}

// GetUniqueTerminalIdOk returns a tuple with the UniqueTerminalId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CancelOrRefundRequest) GetUniqueTerminalIdOk() (*string, bool) {
	if o == nil || common.IsNil(o.UniqueTerminalId) {
		return nil, false
	}
	return o.UniqueTerminalId, true
}

// HasUniqueTerminalId returns a boolean if a field has been set.
func (o *CancelOrRefundRequest) HasUniqueTerminalId() bool {
	if o != nil && !common.IsNil(o.UniqueTerminalId) {
		return true
	}

	return false
}

// SetUniqueTerminalId gets a reference to the given string and assigns it to the UniqueTerminalId field.
func (o *CancelOrRefundRequest) SetUniqueTerminalId(v string) {
	o.UniqueTerminalId = &v
}

func (o CancelOrRefundRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CancelOrRefundRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.AdditionalData) {
		toSerialize["additionalData"] = o.AdditionalData
	}
	toSerialize["merchantAccount"] = o.MerchantAccount
	if !common.IsNil(o.MpiData) {
		toSerialize["mpiData"] = o.MpiData
	}
	if !common.IsNil(o.OriginalMerchantReference) {
		toSerialize["originalMerchantReference"] = o.OriginalMerchantReference
	}
	toSerialize["originalReference"] = o.OriginalReference
	if !common.IsNil(o.PlatformChargebackLogic) {
		toSerialize["platformChargebackLogic"] = o.PlatformChargebackLogic
	}
	if !common.IsNil(o.Reference) {
		toSerialize["reference"] = o.Reference
	}
	if !common.IsNil(o.TenderReference) {
		toSerialize["tenderReference"] = o.TenderReference
	}
	if !common.IsNil(o.UniqueTerminalId) {
		toSerialize["uniqueTerminalId"] = o.UniqueTerminalId
	}
	return toSerialize, nil
}

type NullableCancelOrRefundRequest struct {
	value *CancelOrRefundRequest
	isSet bool
}

func (v NullableCancelOrRefundRequest) Get() *CancelOrRefundRequest {
	return v.value
}

func (v *NullableCancelOrRefundRequest) Set(val *CancelOrRefundRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCancelOrRefundRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCancelOrRefundRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCancelOrRefundRequest(val *CancelOrRefundRequest) *NullableCancelOrRefundRequest {
	return &NullableCancelOrRefundRequest{value: val, isSet: true}
}

func (v NullableCancelOrRefundRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCancelOrRefundRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
