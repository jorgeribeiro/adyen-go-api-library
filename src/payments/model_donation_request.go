/*
Adyen Payment API

API version: 68
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package payments

import (
	"encoding/json"

	"github.com/adyen/adyen-go-api-library/v8/src/common"
)

// checks if the DonationRequest type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &DonationRequest{}

// DonationRequest struct for DonationRequest
type DonationRequest struct {
	// The Adyen account name of the charity.
	DonationAccount string `json:"donationAccount"`
	// The merchant account that is used to process the payment.
	MerchantAccount    string `json:"merchantAccount"`
	ModificationAmount Amount `json:"modificationAmount"`
	// The original pspReference of the payment to modify. This reference is returned in: * authorisation response * authorisation notification
	OriginalReference       *string                  `json:"originalReference,omitempty"`
	PlatformChargebackLogic *PlatformChargebackLogic `json:"platformChargebackLogic,omitempty"`
	// Your reference for the payment modification. This reference is visible in Customer Area and in reports. Maximum length: 80 characters.
	Reference *string `json:"reference,omitempty"`
}

// NewDonationRequest instantiates a new DonationRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDonationRequest(donationAccount string, merchantAccount string, modificationAmount Amount) *DonationRequest {
	this := DonationRequest{}
	this.DonationAccount = donationAccount
	this.MerchantAccount = merchantAccount
	this.ModificationAmount = modificationAmount
	return &this
}

// NewDonationRequestWithDefaults instantiates a new DonationRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDonationRequestWithDefaults() *DonationRequest {
	this := DonationRequest{}
	return &this
}

// GetDonationAccount returns the DonationAccount field value
func (o *DonationRequest) GetDonationAccount() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DonationAccount
}

// GetDonationAccountOk returns a tuple with the DonationAccount field value
// and a boolean to check if the value has been set.
func (o *DonationRequest) GetDonationAccountOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DonationAccount, true
}

// SetDonationAccount sets field value
func (o *DonationRequest) SetDonationAccount(v string) {
	o.DonationAccount = v
}

// GetMerchantAccount returns the MerchantAccount field value
func (o *DonationRequest) GetMerchantAccount() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.MerchantAccount
}

// GetMerchantAccountOk returns a tuple with the MerchantAccount field value
// and a boolean to check if the value has been set.
func (o *DonationRequest) GetMerchantAccountOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MerchantAccount, true
}

// SetMerchantAccount sets field value
func (o *DonationRequest) SetMerchantAccount(v string) {
	o.MerchantAccount = v
}

// GetModificationAmount returns the ModificationAmount field value
func (o *DonationRequest) GetModificationAmount() Amount {
	if o == nil {
		var ret Amount
		return ret
	}

	return o.ModificationAmount
}

// GetModificationAmountOk returns a tuple with the ModificationAmount field value
// and a boolean to check if the value has been set.
func (o *DonationRequest) GetModificationAmountOk() (*Amount, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ModificationAmount, true
}

// SetModificationAmount sets field value
func (o *DonationRequest) SetModificationAmount(v Amount) {
	o.ModificationAmount = v
}

// GetOriginalReference returns the OriginalReference field value if set, zero value otherwise.
func (o *DonationRequest) GetOriginalReference() string {
	if o == nil || common.IsNil(o.OriginalReference) {
		var ret string
		return ret
	}
	return *o.OriginalReference
}

// GetOriginalReferenceOk returns a tuple with the OriginalReference field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DonationRequest) GetOriginalReferenceOk() (*string, bool) {
	if o == nil || common.IsNil(o.OriginalReference) {
		return nil, false
	}
	return o.OriginalReference, true
}

// HasOriginalReference returns a boolean if a field has been set.
func (o *DonationRequest) HasOriginalReference() bool {
	if o != nil && !common.IsNil(o.OriginalReference) {
		return true
	}

	return false
}

// SetOriginalReference gets a reference to the given string and assigns it to the OriginalReference field.
func (o *DonationRequest) SetOriginalReference(v string) {
	o.OriginalReference = &v
}

// GetPlatformChargebackLogic returns the PlatformChargebackLogic field value if set, zero value otherwise.
func (o *DonationRequest) GetPlatformChargebackLogic() PlatformChargebackLogic {
	if o == nil || common.IsNil(o.PlatformChargebackLogic) {
		var ret PlatformChargebackLogic
		return ret
	}
	return *o.PlatformChargebackLogic
}

// GetPlatformChargebackLogicOk returns a tuple with the PlatformChargebackLogic field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DonationRequest) GetPlatformChargebackLogicOk() (*PlatformChargebackLogic, bool) {
	if o == nil || common.IsNil(o.PlatformChargebackLogic) {
		return nil, false
	}
	return o.PlatformChargebackLogic, true
}

// HasPlatformChargebackLogic returns a boolean if a field has been set.
func (o *DonationRequest) HasPlatformChargebackLogic() bool {
	if o != nil && !common.IsNil(o.PlatformChargebackLogic) {
		return true
	}

	return false
}

// SetPlatformChargebackLogic gets a reference to the given PlatformChargebackLogic and assigns it to the PlatformChargebackLogic field.
func (o *DonationRequest) SetPlatformChargebackLogic(v PlatformChargebackLogic) {
	o.PlatformChargebackLogic = &v
}

// GetReference returns the Reference field value if set, zero value otherwise.
func (o *DonationRequest) GetReference() string {
	if o == nil || common.IsNil(o.Reference) {
		var ret string
		return ret
	}
	return *o.Reference
}

// GetReferenceOk returns a tuple with the Reference field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DonationRequest) GetReferenceOk() (*string, bool) {
	if o == nil || common.IsNil(o.Reference) {
		return nil, false
	}
	return o.Reference, true
}

// HasReference returns a boolean if a field has been set.
func (o *DonationRequest) HasReference() bool {
	if o != nil && !common.IsNil(o.Reference) {
		return true
	}

	return false
}

// SetReference gets a reference to the given string and assigns it to the Reference field.
func (o *DonationRequest) SetReference(v string) {
	o.Reference = &v
}

func (o DonationRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DonationRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["donationAccount"] = o.DonationAccount
	toSerialize["merchantAccount"] = o.MerchantAccount
	toSerialize["modificationAmount"] = o.ModificationAmount
	if !common.IsNil(o.OriginalReference) {
		toSerialize["originalReference"] = o.OriginalReference
	}
	if !common.IsNil(o.PlatformChargebackLogic) {
		toSerialize["platformChargebackLogic"] = o.PlatformChargebackLogic
	}
	if !common.IsNil(o.Reference) {
		toSerialize["reference"] = o.Reference
	}
	return toSerialize, nil
}

type NullableDonationRequest struct {
	value *DonationRequest
	isSet bool
}

func (v NullableDonationRequest) Get() *DonationRequest {
	return v.value
}

func (v *NullableDonationRequest) Set(val *DonationRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableDonationRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableDonationRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDonationRequest(val *DonationRequest) *NullableDonationRequest {
	return &NullableDonationRequest{value: val, isSet: true}
}

func (v NullableDonationRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDonationRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}