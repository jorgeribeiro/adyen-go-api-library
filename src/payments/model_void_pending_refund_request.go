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

// checks if the VoidPendingRefundRequest type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &VoidPendingRefundRequest{}

// VoidPendingRefundRequest struct for VoidPendingRefundRequest
type VoidPendingRefundRequest struct {
	// This field contains additional data, which may be required for a particular modification request.  The additionalData object consists of entries, each of which includes the key and value.
	AdditionalData *map[string]string `json:"additionalData,omitempty"`
	// The merchant account that is used to process the payment.
	MerchantAccount    string            `json:"merchantAccount"`
	ModificationAmount *Amount           `json:"modificationAmount,omitempty"`
	MpiData            *ThreeDSecureData `json:"mpiData,omitempty"`
	// The original merchant reference to cancel.
	OriginalMerchantReference *string `json:"originalMerchantReference,omitempty"`
	// The original pspReference of the payment to modify. This reference is returned in: * authorisation response * authorisation notification
	OriginalReference       *string                  `json:"originalReference,omitempty"`
	PlatformChargebackLogic *PlatformChargebackLogic `json:"platformChargebackLogic,omitempty"`
	// Your reference for the payment modification. This reference is visible in Customer Area and in reports. Maximum length: 80 characters.
	Reference *string `json:"reference,omitempty"`
	// An array of objects specifying how the amount should be split between accounts when using Adyen for Platforms. For more information, see how to split payments for [platforms](https://docs.adyen.com/platforms/automatic-split-configuration/).
	Splits []Split `json:"splits,omitempty"`
	// The transaction reference provided by the PED. For point-of-sale integrations only.
	TenderReference *string `json:"tenderReference,omitempty"`
	// Unique terminal ID for the PED that originally processed the request. For point-of-sale integrations only.
	UniqueTerminalId *string `json:"uniqueTerminalId,omitempty"`
}

// NewVoidPendingRefundRequest instantiates a new VoidPendingRefundRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVoidPendingRefundRequest(merchantAccount string) *VoidPendingRefundRequest {
	this := VoidPendingRefundRequest{}
	this.MerchantAccount = merchantAccount
	return &this
}

// NewVoidPendingRefundRequestWithDefaults instantiates a new VoidPendingRefundRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVoidPendingRefundRequestWithDefaults() *VoidPendingRefundRequest {
	this := VoidPendingRefundRequest{}
	return &this
}

// GetAdditionalData returns the AdditionalData field value if set, zero value otherwise.
func (o *VoidPendingRefundRequest) GetAdditionalData() map[string]string {
	if o == nil || common.IsNil(o.AdditionalData) {
		var ret map[string]string
		return ret
	}
	return *o.AdditionalData
}

// GetAdditionalDataOk returns a tuple with the AdditionalData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VoidPendingRefundRequest) GetAdditionalDataOk() (*map[string]string, bool) {
	if o == nil || common.IsNil(o.AdditionalData) {
		return nil, false
	}
	return o.AdditionalData, true
}

// HasAdditionalData returns a boolean if a field has been set.
func (o *VoidPendingRefundRequest) HasAdditionalData() bool {
	if o != nil && !common.IsNil(o.AdditionalData) {
		return true
	}

	return false
}

// SetAdditionalData gets a reference to the given map[string]string and assigns it to the AdditionalData field.
func (o *VoidPendingRefundRequest) SetAdditionalData(v map[string]string) {
	o.AdditionalData = &v
}

// GetMerchantAccount returns the MerchantAccount field value
func (o *VoidPendingRefundRequest) GetMerchantAccount() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.MerchantAccount
}

// GetMerchantAccountOk returns a tuple with the MerchantAccount field value
// and a boolean to check if the value has been set.
func (o *VoidPendingRefundRequest) GetMerchantAccountOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MerchantAccount, true
}

// SetMerchantAccount sets field value
func (o *VoidPendingRefundRequest) SetMerchantAccount(v string) {
	o.MerchantAccount = v
}

// GetModificationAmount returns the ModificationAmount field value if set, zero value otherwise.
func (o *VoidPendingRefundRequest) GetModificationAmount() Amount {
	if o == nil || common.IsNil(o.ModificationAmount) {
		var ret Amount
		return ret
	}
	return *o.ModificationAmount
}

// GetModificationAmountOk returns a tuple with the ModificationAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VoidPendingRefundRequest) GetModificationAmountOk() (*Amount, bool) {
	if o == nil || common.IsNil(o.ModificationAmount) {
		return nil, false
	}
	return o.ModificationAmount, true
}

// HasModificationAmount returns a boolean if a field has been set.
func (o *VoidPendingRefundRequest) HasModificationAmount() bool {
	if o != nil && !common.IsNil(o.ModificationAmount) {
		return true
	}

	return false
}

// SetModificationAmount gets a reference to the given Amount and assigns it to the ModificationAmount field.
func (o *VoidPendingRefundRequest) SetModificationAmount(v Amount) {
	o.ModificationAmount = &v
}

// GetMpiData returns the MpiData field value if set, zero value otherwise.
func (o *VoidPendingRefundRequest) GetMpiData() ThreeDSecureData {
	if o == nil || common.IsNil(o.MpiData) {
		var ret ThreeDSecureData
		return ret
	}
	return *o.MpiData
}

// GetMpiDataOk returns a tuple with the MpiData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VoidPendingRefundRequest) GetMpiDataOk() (*ThreeDSecureData, bool) {
	if o == nil || common.IsNil(o.MpiData) {
		return nil, false
	}
	return o.MpiData, true
}

// HasMpiData returns a boolean if a field has been set.
func (o *VoidPendingRefundRequest) HasMpiData() bool {
	if o != nil && !common.IsNil(o.MpiData) {
		return true
	}

	return false
}

// SetMpiData gets a reference to the given ThreeDSecureData and assigns it to the MpiData field.
func (o *VoidPendingRefundRequest) SetMpiData(v ThreeDSecureData) {
	o.MpiData = &v
}

// GetOriginalMerchantReference returns the OriginalMerchantReference field value if set, zero value otherwise.
func (o *VoidPendingRefundRequest) GetOriginalMerchantReference() string {
	if o == nil || common.IsNil(o.OriginalMerchantReference) {
		var ret string
		return ret
	}
	return *o.OriginalMerchantReference
}

// GetOriginalMerchantReferenceOk returns a tuple with the OriginalMerchantReference field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VoidPendingRefundRequest) GetOriginalMerchantReferenceOk() (*string, bool) {
	if o == nil || common.IsNil(o.OriginalMerchantReference) {
		return nil, false
	}
	return o.OriginalMerchantReference, true
}

// HasOriginalMerchantReference returns a boolean if a field has been set.
func (o *VoidPendingRefundRequest) HasOriginalMerchantReference() bool {
	if o != nil && !common.IsNil(o.OriginalMerchantReference) {
		return true
	}

	return false
}

// SetOriginalMerchantReference gets a reference to the given string and assigns it to the OriginalMerchantReference field.
func (o *VoidPendingRefundRequest) SetOriginalMerchantReference(v string) {
	o.OriginalMerchantReference = &v
}

// GetOriginalReference returns the OriginalReference field value if set, zero value otherwise.
func (o *VoidPendingRefundRequest) GetOriginalReference() string {
	if o == nil || common.IsNil(o.OriginalReference) {
		var ret string
		return ret
	}
	return *o.OriginalReference
}

// GetOriginalReferenceOk returns a tuple with the OriginalReference field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VoidPendingRefundRequest) GetOriginalReferenceOk() (*string, bool) {
	if o == nil || common.IsNil(o.OriginalReference) {
		return nil, false
	}
	return o.OriginalReference, true
}

// HasOriginalReference returns a boolean if a field has been set.
func (o *VoidPendingRefundRequest) HasOriginalReference() bool {
	if o != nil && !common.IsNil(o.OriginalReference) {
		return true
	}

	return false
}

// SetOriginalReference gets a reference to the given string and assigns it to the OriginalReference field.
func (o *VoidPendingRefundRequest) SetOriginalReference(v string) {
	o.OriginalReference = &v
}

// GetPlatformChargebackLogic returns the PlatformChargebackLogic field value if set, zero value otherwise.
func (o *VoidPendingRefundRequest) GetPlatformChargebackLogic() PlatformChargebackLogic {
	if o == nil || common.IsNil(o.PlatformChargebackLogic) {
		var ret PlatformChargebackLogic
		return ret
	}
	return *o.PlatformChargebackLogic
}

// GetPlatformChargebackLogicOk returns a tuple with the PlatformChargebackLogic field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VoidPendingRefundRequest) GetPlatformChargebackLogicOk() (*PlatformChargebackLogic, bool) {
	if o == nil || common.IsNil(o.PlatformChargebackLogic) {
		return nil, false
	}
	return o.PlatformChargebackLogic, true
}

// HasPlatformChargebackLogic returns a boolean if a field has been set.
func (o *VoidPendingRefundRequest) HasPlatformChargebackLogic() bool {
	if o != nil && !common.IsNil(o.PlatformChargebackLogic) {
		return true
	}

	return false
}

// SetPlatformChargebackLogic gets a reference to the given PlatformChargebackLogic and assigns it to the PlatformChargebackLogic field.
func (o *VoidPendingRefundRequest) SetPlatformChargebackLogic(v PlatformChargebackLogic) {
	o.PlatformChargebackLogic = &v
}

// GetReference returns the Reference field value if set, zero value otherwise.
func (o *VoidPendingRefundRequest) GetReference() string {
	if o == nil || common.IsNil(o.Reference) {
		var ret string
		return ret
	}
	return *o.Reference
}

// GetReferenceOk returns a tuple with the Reference field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VoidPendingRefundRequest) GetReferenceOk() (*string, bool) {
	if o == nil || common.IsNil(o.Reference) {
		return nil, false
	}
	return o.Reference, true
}

// HasReference returns a boolean if a field has been set.
func (o *VoidPendingRefundRequest) HasReference() bool {
	if o != nil && !common.IsNil(o.Reference) {
		return true
	}

	return false
}

// SetReference gets a reference to the given string and assigns it to the Reference field.
func (o *VoidPendingRefundRequest) SetReference(v string) {
	o.Reference = &v
}

// GetSplits returns the Splits field value if set, zero value otherwise.
func (o *VoidPendingRefundRequest) GetSplits() []Split {
	if o == nil || common.IsNil(o.Splits) {
		var ret []Split
		return ret
	}
	return o.Splits
}

// GetSplitsOk returns a tuple with the Splits field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VoidPendingRefundRequest) GetSplitsOk() ([]Split, bool) {
	if o == nil || common.IsNil(o.Splits) {
		return nil, false
	}
	return o.Splits, true
}

// HasSplits returns a boolean if a field has been set.
func (o *VoidPendingRefundRequest) HasSplits() bool {
	if o != nil && !common.IsNil(o.Splits) {
		return true
	}

	return false
}

// SetSplits gets a reference to the given []Split and assigns it to the Splits field.
func (o *VoidPendingRefundRequest) SetSplits(v []Split) {
	o.Splits = v
}

// GetTenderReference returns the TenderReference field value if set, zero value otherwise.
func (o *VoidPendingRefundRequest) GetTenderReference() string {
	if o == nil || common.IsNil(o.TenderReference) {
		var ret string
		return ret
	}
	return *o.TenderReference
}

// GetTenderReferenceOk returns a tuple with the TenderReference field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VoidPendingRefundRequest) GetTenderReferenceOk() (*string, bool) {
	if o == nil || common.IsNil(o.TenderReference) {
		return nil, false
	}
	return o.TenderReference, true
}

// HasTenderReference returns a boolean if a field has been set.
func (o *VoidPendingRefundRequest) HasTenderReference() bool {
	if o != nil && !common.IsNil(o.TenderReference) {
		return true
	}

	return false
}

// SetTenderReference gets a reference to the given string and assigns it to the TenderReference field.
func (o *VoidPendingRefundRequest) SetTenderReference(v string) {
	o.TenderReference = &v
}

// GetUniqueTerminalId returns the UniqueTerminalId field value if set, zero value otherwise.
func (o *VoidPendingRefundRequest) GetUniqueTerminalId() string {
	if o == nil || common.IsNil(o.UniqueTerminalId) {
		var ret string
		return ret
	}
	return *o.UniqueTerminalId
}

// GetUniqueTerminalIdOk returns a tuple with the UniqueTerminalId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VoidPendingRefundRequest) GetUniqueTerminalIdOk() (*string, bool) {
	if o == nil || common.IsNil(o.UniqueTerminalId) {
		return nil, false
	}
	return o.UniqueTerminalId, true
}

// HasUniqueTerminalId returns a boolean if a field has been set.
func (o *VoidPendingRefundRequest) HasUniqueTerminalId() bool {
	if o != nil && !common.IsNil(o.UniqueTerminalId) {
		return true
	}

	return false
}

// SetUniqueTerminalId gets a reference to the given string and assigns it to the UniqueTerminalId field.
func (o *VoidPendingRefundRequest) SetUniqueTerminalId(v string) {
	o.UniqueTerminalId = &v
}

func (o VoidPendingRefundRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o VoidPendingRefundRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.AdditionalData) {
		toSerialize["additionalData"] = o.AdditionalData
	}
	toSerialize["merchantAccount"] = o.MerchantAccount
	if !common.IsNil(o.ModificationAmount) {
		toSerialize["modificationAmount"] = o.ModificationAmount
	}
	if !common.IsNil(o.MpiData) {
		toSerialize["mpiData"] = o.MpiData
	}
	if !common.IsNil(o.OriginalMerchantReference) {
		toSerialize["originalMerchantReference"] = o.OriginalMerchantReference
	}
	if !common.IsNil(o.OriginalReference) {
		toSerialize["originalReference"] = o.OriginalReference
	}
	if !common.IsNil(o.PlatformChargebackLogic) {
		toSerialize["platformChargebackLogic"] = o.PlatformChargebackLogic
	}
	if !common.IsNil(o.Reference) {
		toSerialize["reference"] = o.Reference
	}
	if !common.IsNil(o.Splits) {
		toSerialize["splits"] = o.Splits
	}
	if !common.IsNil(o.TenderReference) {
		toSerialize["tenderReference"] = o.TenderReference
	}
	if !common.IsNil(o.UniqueTerminalId) {
		toSerialize["uniqueTerminalId"] = o.UniqueTerminalId
	}
	return toSerialize, nil
}

type NullableVoidPendingRefundRequest struct {
	value *VoidPendingRefundRequest
	isSet bool
}

func (v NullableVoidPendingRefundRequest) Get() *VoidPendingRefundRequest {
	return v.value
}

func (v *NullableVoidPendingRefundRequest) Set(val *VoidPendingRefundRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableVoidPendingRefundRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableVoidPendingRefundRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVoidPendingRefundRequest(val *VoidPendingRefundRequest) *NullableVoidPendingRefundRequest {
	return &NullableVoidPendingRefundRequest{value: val, isSet: true}
}

func (v NullableVoidPendingRefundRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVoidPendingRefundRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
