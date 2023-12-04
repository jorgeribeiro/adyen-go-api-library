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

// checks if the SplitConfigurationLogic type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &SplitConfigurationLogic{}

// SplitConfigurationLogic struct for SplitConfigurationLogic
type SplitConfigurationLogic struct {
	// Specifies the logic to apply when booking the transaction fees. Should be combined with adyenFees.  Possible values: **deductFromLiableAccount**, **deductFromOneBalanceAccount**.
	AcquiringFees        *string               `json:"acquiringFees,omitempty"`
	AdditionalCommission *AdditionalCommission `json:"additionalCommission,omitempty"`
	// Specifies the logic to apply when booking the transaction fees. Should be combined with schemeFee, interchange & adyenMarkup.  Possible values: **deductFromLiableAccount**, **deductFromOneBalanceAccount**.
	AdyenCommission *string `json:"adyenCommission,omitempty"`
	// Specifies the logic to apply when booking the transaction fees. Should be combined with acquiringFees.  Possible values: **deductFromLiableAccount**, **deductFromOneBalanceAccount**.
	AdyenFees *string `json:"adyenFees,omitempty"`
	// Specifies the logic to apply when booking the transaction fees. Should be combined with schemeFee, adyenCommission & interchange.  Possible values: **deductFromLiableAccount**, **deductFromOneBalanceAccount**.
	AdyenMarkup *string `json:"adyenMarkup,omitempty"`
	// Specifies the logic to apply when booking the chargeback amount.  Possible values: **deductFromLiableAccount**, **deductFromOneBalanceAccount**, **deductAccordingToSplitRatio**.
	Chargeback *string `json:"chargeback,omitempty"`
	// Specifies the logic to apply when allocating the chargeback costs.  Possible values: **deductFromLiableAccount**, **deductFromOneBalanceAccount**
	ChargebackCostAllocation *string    `json:"chargebackCostAllocation,omitempty"`
	Commission               Commission `json:"commission"`
	// Specifies the logic to apply when booking the transaction fees. Should be combined with schemeFee, adyenCommission & adyenMarkup.  Possible values: **deductFromLiableAccount**, **deductFromOneBalanceAccount**.
	Interchange *string `json:"interchange,omitempty"`
	// Specifies the logic to apply when booking the transaction fees. Cannot be combined with other fees.  Possible values: **deductFromLiableAccount**, **deductFromOneBalanceAccount**.
	PaymentFee *string `json:"paymentFee,omitempty"`
	// Specifies the logic to apply when booking the amount left over after currency conversion.  Possible values: **addToLiableAccount**, **addToOneBalanceAccount**.
	Remainder *string `json:"remainder,omitempty"`
	// Specifies the logic to apply when booking the transaction fees. Should be combined with interchange, adyenCommission & adyenMarkup.  Possible values: **deductFromLiableAccount**, **deductFromOneBalanceAccount**.
	SchemeFee *string `json:"schemeFee,omitempty"`
	// Unique identifier of the split logic that is applied when the split configuration conditions are met.
	SplitLogicId *string `json:"splitLogicId,omitempty"`
	// Specifies the logic to apply when booking the surcharge amount.  Possible values: **addToLiableAccount**, **addToOneBalanceAccount**
	Surcharge *string `json:"surcharge,omitempty"`
	// Specifies the logic to apply when booking tips (gratuity).  Possible values: **addToLiableAccount**, **addToOneBalanceAccount**.
	Tip *string `json:"tip,omitempty"`
}

// NewSplitConfigurationLogic instantiates a new SplitConfigurationLogic object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSplitConfigurationLogic(commission Commission) *SplitConfigurationLogic {
	this := SplitConfigurationLogic{}
	this.Commission = commission
	return &this
}

// NewSplitConfigurationLogicWithDefaults instantiates a new SplitConfigurationLogic object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSplitConfigurationLogicWithDefaults() *SplitConfigurationLogic {
	this := SplitConfigurationLogic{}
	return &this
}

// GetAcquiringFees returns the AcquiringFees field value if set, zero value otherwise.
func (o *SplitConfigurationLogic) GetAcquiringFees() string {
	if o == nil || common.IsNil(o.AcquiringFees) {
		var ret string
		return ret
	}
	return *o.AcquiringFees
}

// GetAcquiringFeesOk returns a tuple with the AcquiringFees field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SplitConfigurationLogic) GetAcquiringFeesOk() (*string, bool) {
	if o == nil || common.IsNil(o.AcquiringFees) {
		return nil, false
	}
	return o.AcquiringFees, true
}

// HasAcquiringFees returns a boolean if a field has been set.
func (o *SplitConfigurationLogic) HasAcquiringFees() bool {
	if o != nil && !common.IsNil(o.AcquiringFees) {
		return true
	}

	return false
}

// SetAcquiringFees gets a reference to the given string and assigns it to the AcquiringFees field.
func (o *SplitConfigurationLogic) SetAcquiringFees(v string) {
	o.AcquiringFees = &v
}

// GetAdditionalCommission returns the AdditionalCommission field value if set, zero value otherwise.
func (o *SplitConfigurationLogic) GetAdditionalCommission() AdditionalCommission {
	if o == nil || common.IsNil(o.AdditionalCommission) {
		var ret AdditionalCommission
		return ret
	}
	return *o.AdditionalCommission
}

// GetAdditionalCommissionOk returns a tuple with the AdditionalCommission field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SplitConfigurationLogic) GetAdditionalCommissionOk() (*AdditionalCommission, bool) {
	if o == nil || common.IsNil(o.AdditionalCommission) {
		return nil, false
	}
	return o.AdditionalCommission, true
}

// HasAdditionalCommission returns a boolean if a field has been set.
func (o *SplitConfigurationLogic) HasAdditionalCommission() bool {
	if o != nil && !common.IsNil(o.AdditionalCommission) {
		return true
	}

	return false
}

// SetAdditionalCommission gets a reference to the given AdditionalCommission and assigns it to the AdditionalCommission field.
func (o *SplitConfigurationLogic) SetAdditionalCommission(v AdditionalCommission) {
	o.AdditionalCommission = &v
}

// GetAdyenCommission returns the AdyenCommission field value if set, zero value otherwise.
func (o *SplitConfigurationLogic) GetAdyenCommission() string {
	if o == nil || common.IsNil(o.AdyenCommission) {
		var ret string
		return ret
	}
	return *o.AdyenCommission
}

// GetAdyenCommissionOk returns a tuple with the AdyenCommission field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SplitConfigurationLogic) GetAdyenCommissionOk() (*string, bool) {
	if o == nil || common.IsNil(o.AdyenCommission) {
		return nil, false
	}
	return o.AdyenCommission, true
}

// HasAdyenCommission returns a boolean if a field has been set.
func (o *SplitConfigurationLogic) HasAdyenCommission() bool {
	if o != nil && !common.IsNil(o.AdyenCommission) {
		return true
	}

	return false
}

// SetAdyenCommission gets a reference to the given string and assigns it to the AdyenCommission field.
func (o *SplitConfigurationLogic) SetAdyenCommission(v string) {
	o.AdyenCommission = &v
}

// GetAdyenFees returns the AdyenFees field value if set, zero value otherwise.
func (o *SplitConfigurationLogic) GetAdyenFees() string {
	if o == nil || common.IsNil(o.AdyenFees) {
		var ret string
		return ret
	}
	return *o.AdyenFees
}

// GetAdyenFeesOk returns a tuple with the AdyenFees field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SplitConfigurationLogic) GetAdyenFeesOk() (*string, bool) {
	if o == nil || common.IsNil(o.AdyenFees) {
		return nil, false
	}
	return o.AdyenFees, true
}

// HasAdyenFees returns a boolean if a field has been set.
func (o *SplitConfigurationLogic) HasAdyenFees() bool {
	if o != nil && !common.IsNil(o.AdyenFees) {
		return true
	}

	return false
}

// SetAdyenFees gets a reference to the given string and assigns it to the AdyenFees field.
func (o *SplitConfigurationLogic) SetAdyenFees(v string) {
	o.AdyenFees = &v
}

// GetAdyenMarkup returns the AdyenMarkup field value if set, zero value otherwise.
func (o *SplitConfigurationLogic) GetAdyenMarkup() string {
	if o == nil || common.IsNil(o.AdyenMarkup) {
		var ret string
		return ret
	}
	return *o.AdyenMarkup
}

// GetAdyenMarkupOk returns a tuple with the AdyenMarkup field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SplitConfigurationLogic) GetAdyenMarkupOk() (*string, bool) {
	if o == nil || common.IsNil(o.AdyenMarkup) {
		return nil, false
	}
	return o.AdyenMarkup, true
}

// HasAdyenMarkup returns a boolean if a field has been set.
func (o *SplitConfigurationLogic) HasAdyenMarkup() bool {
	if o != nil && !common.IsNil(o.AdyenMarkup) {
		return true
	}

	return false
}

// SetAdyenMarkup gets a reference to the given string and assigns it to the AdyenMarkup field.
func (o *SplitConfigurationLogic) SetAdyenMarkup(v string) {
	o.AdyenMarkup = &v
}

// GetChargeback returns the Chargeback field value if set, zero value otherwise.
func (o *SplitConfigurationLogic) GetChargeback() string {
	if o == nil || common.IsNil(o.Chargeback) {
		var ret string
		return ret
	}
	return *o.Chargeback
}

// GetChargebackOk returns a tuple with the Chargeback field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SplitConfigurationLogic) GetChargebackOk() (*string, bool) {
	if o == nil || common.IsNil(o.Chargeback) {
		return nil, false
	}
	return o.Chargeback, true
}

// HasChargeback returns a boolean if a field has been set.
func (o *SplitConfigurationLogic) HasChargeback() bool {
	if o != nil && !common.IsNil(o.Chargeback) {
		return true
	}

	return false
}

// SetChargeback gets a reference to the given string and assigns it to the Chargeback field.
func (o *SplitConfigurationLogic) SetChargeback(v string) {
	o.Chargeback = &v
}

// GetChargebackCostAllocation returns the ChargebackCostAllocation field value if set, zero value otherwise.
func (o *SplitConfigurationLogic) GetChargebackCostAllocation() string {
	if o == nil || common.IsNil(o.ChargebackCostAllocation) {
		var ret string
		return ret
	}
	return *o.ChargebackCostAllocation
}

// GetChargebackCostAllocationOk returns a tuple with the ChargebackCostAllocation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SplitConfigurationLogic) GetChargebackCostAllocationOk() (*string, bool) {
	if o == nil || common.IsNil(o.ChargebackCostAllocation) {
		return nil, false
	}
	return o.ChargebackCostAllocation, true
}

// HasChargebackCostAllocation returns a boolean if a field has been set.
func (o *SplitConfigurationLogic) HasChargebackCostAllocation() bool {
	if o != nil && !common.IsNil(o.ChargebackCostAllocation) {
		return true
	}

	return false
}

// SetChargebackCostAllocation gets a reference to the given string and assigns it to the ChargebackCostAllocation field.
func (o *SplitConfigurationLogic) SetChargebackCostAllocation(v string) {
	o.ChargebackCostAllocation = &v
}

// GetCommission returns the Commission field value
func (o *SplitConfigurationLogic) GetCommission() Commission {
	if o == nil {
		var ret Commission
		return ret
	}

	return o.Commission
}

// GetCommissionOk returns a tuple with the Commission field value
// and a boolean to check if the value has been set.
func (o *SplitConfigurationLogic) GetCommissionOk() (*Commission, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Commission, true
}

// SetCommission sets field value
func (o *SplitConfigurationLogic) SetCommission(v Commission) {
	o.Commission = v
}

// GetInterchange returns the Interchange field value if set, zero value otherwise.
func (o *SplitConfigurationLogic) GetInterchange() string {
	if o == nil || common.IsNil(o.Interchange) {
		var ret string
		return ret
	}
	return *o.Interchange
}

// GetInterchangeOk returns a tuple with the Interchange field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SplitConfigurationLogic) GetInterchangeOk() (*string, bool) {
	if o == nil || common.IsNil(o.Interchange) {
		return nil, false
	}
	return o.Interchange, true
}

// HasInterchange returns a boolean if a field has been set.
func (o *SplitConfigurationLogic) HasInterchange() bool {
	if o != nil && !common.IsNil(o.Interchange) {
		return true
	}

	return false
}

// SetInterchange gets a reference to the given string and assigns it to the Interchange field.
func (o *SplitConfigurationLogic) SetInterchange(v string) {
	o.Interchange = &v
}

// GetPaymentFee returns the PaymentFee field value if set, zero value otherwise.
func (o *SplitConfigurationLogic) GetPaymentFee() string {
	if o == nil || common.IsNil(o.PaymentFee) {
		var ret string
		return ret
	}
	return *o.PaymentFee
}

// GetPaymentFeeOk returns a tuple with the PaymentFee field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SplitConfigurationLogic) GetPaymentFeeOk() (*string, bool) {
	if o == nil || common.IsNil(o.PaymentFee) {
		return nil, false
	}
	return o.PaymentFee, true
}

// HasPaymentFee returns a boolean if a field has been set.
func (o *SplitConfigurationLogic) HasPaymentFee() bool {
	if o != nil && !common.IsNil(o.PaymentFee) {
		return true
	}

	return false
}

// SetPaymentFee gets a reference to the given string and assigns it to the PaymentFee field.
func (o *SplitConfigurationLogic) SetPaymentFee(v string) {
	o.PaymentFee = &v
}

// GetRemainder returns the Remainder field value if set, zero value otherwise.
func (o *SplitConfigurationLogic) GetRemainder() string {
	if o == nil || common.IsNil(o.Remainder) {
		var ret string
		return ret
	}
	return *o.Remainder
}

// GetRemainderOk returns a tuple with the Remainder field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SplitConfigurationLogic) GetRemainderOk() (*string, bool) {
	if o == nil || common.IsNil(o.Remainder) {
		return nil, false
	}
	return o.Remainder, true
}

// HasRemainder returns a boolean if a field has been set.
func (o *SplitConfigurationLogic) HasRemainder() bool {
	if o != nil && !common.IsNil(o.Remainder) {
		return true
	}

	return false
}

// SetRemainder gets a reference to the given string and assigns it to the Remainder field.
func (o *SplitConfigurationLogic) SetRemainder(v string) {
	o.Remainder = &v
}

// GetSchemeFee returns the SchemeFee field value if set, zero value otherwise.
func (o *SplitConfigurationLogic) GetSchemeFee() string {
	if o == nil || common.IsNil(o.SchemeFee) {
		var ret string
		return ret
	}
	return *o.SchemeFee
}

// GetSchemeFeeOk returns a tuple with the SchemeFee field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SplitConfigurationLogic) GetSchemeFeeOk() (*string, bool) {
	if o == nil || common.IsNil(o.SchemeFee) {
		return nil, false
	}
	return o.SchemeFee, true
}

// HasSchemeFee returns a boolean if a field has been set.
func (o *SplitConfigurationLogic) HasSchemeFee() bool {
	if o != nil && !common.IsNil(o.SchemeFee) {
		return true
	}

	return false
}

// SetSchemeFee gets a reference to the given string and assigns it to the SchemeFee field.
func (o *SplitConfigurationLogic) SetSchemeFee(v string) {
	o.SchemeFee = &v
}

// GetSplitLogicId returns the SplitLogicId field value if set, zero value otherwise.
func (o *SplitConfigurationLogic) GetSplitLogicId() string {
	if o == nil || common.IsNil(o.SplitLogicId) {
		var ret string
		return ret
	}
	return *o.SplitLogicId
}

// GetSplitLogicIdOk returns a tuple with the SplitLogicId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SplitConfigurationLogic) GetSplitLogicIdOk() (*string, bool) {
	if o == nil || common.IsNil(o.SplitLogicId) {
		return nil, false
	}
	return o.SplitLogicId, true
}

// HasSplitLogicId returns a boolean if a field has been set.
func (o *SplitConfigurationLogic) HasSplitLogicId() bool {
	if o != nil && !common.IsNil(o.SplitLogicId) {
		return true
	}

	return false
}

// SetSplitLogicId gets a reference to the given string and assigns it to the SplitLogicId field.
func (o *SplitConfigurationLogic) SetSplitLogicId(v string) {
	o.SplitLogicId = &v
}

// GetSurcharge returns the Surcharge field value if set, zero value otherwise.
func (o *SplitConfigurationLogic) GetSurcharge() string {
	if o == nil || common.IsNil(o.Surcharge) {
		var ret string
		return ret
	}
	return *o.Surcharge
}

// GetSurchargeOk returns a tuple with the Surcharge field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SplitConfigurationLogic) GetSurchargeOk() (*string, bool) {
	if o == nil || common.IsNil(o.Surcharge) {
		return nil, false
	}
	return o.Surcharge, true
}

// HasSurcharge returns a boolean if a field has been set.
func (o *SplitConfigurationLogic) HasSurcharge() bool {
	if o != nil && !common.IsNil(o.Surcharge) {
		return true
	}

	return false
}

// SetSurcharge gets a reference to the given string and assigns it to the Surcharge field.
func (o *SplitConfigurationLogic) SetSurcharge(v string) {
	o.Surcharge = &v
}

// GetTip returns the Tip field value if set, zero value otherwise.
func (o *SplitConfigurationLogic) GetTip() string {
	if o == nil || common.IsNil(o.Tip) {
		var ret string
		return ret
	}
	return *o.Tip
}

// GetTipOk returns a tuple with the Tip field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SplitConfigurationLogic) GetTipOk() (*string, bool) {
	if o == nil || common.IsNil(o.Tip) {
		return nil, false
	}
	return o.Tip, true
}

// HasTip returns a boolean if a field has been set.
func (o *SplitConfigurationLogic) HasTip() bool {
	if o != nil && !common.IsNil(o.Tip) {
		return true
	}

	return false
}

// SetTip gets a reference to the given string and assigns it to the Tip field.
func (o *SplitConfigurationLogic) SetTip(v string) {
	o.Tip = &v
}

func (o SplitConfigurationLogic) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SplitConfigurationLogic) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.AcquiringFees) {
		toSerialize["acquiringFees"] = o.AcquiringFees
	}
	if !common.IsNil(o.AdditionalCommission) {
		toSerialize["additionalCommission"] = o.AdditionalCommission
	}
	if !common.IsNil(o.AdyenCommission) {
		toSerialize["adyenCommission"] = o.AdyenCommission
	}
	if !common.IsNil(o.AdyenFees) {
		toSerialize["adyenFees"] = o.AdyenFees
	}
	if !common.IsNil(o.AdyenMarkup) {
		toSerialize["adyenMarkup"] = o.AdyenMarkup
	}
	if !common.IsNil(o.Chargeback) {
		toSerialize["chargeback"] = o.Chargeback
	}
	if !common.IsNil(o.ChargebackCostAllocation) {
		toSerialize["chargebackCostAllocation"] = o.ChargebackCostAllocation
	}
	toSerialize["commission"] = o.Commission
	if !common.IsNil(o.Interchange) {
		toSerialize["interchange"] = o.Interchange
	}
	if !common.IsNil(o.PaymentFee) {
		toSerialize["paymentFee"] = o.PaymentFee
	}
	if !common.IsNil(o.Remainder) {
		toSerialize["remainder"] = o.Remainder
	}
	if !common.IsNil(o.SchemeFee) {
		toSerialize["schemeFee"] = o.SchemeFee
	}
	if !common.IsNil(o.SplitLogicId) {
		toSerialize["splitLogicId"] = o.SplitLogicId
	}
	if !common.IsNil(o.Surcharge) {
		toSerialize["surcharge"] = o.Surcharge
	}
	if !common.IsNil(o.Tip) {
		toSerialize["tip"] = o.Tip
	}
	return toSerialize, nil
}

type NullableSplitConfigurationLogic struct {
	value *SplitConfigurationLogic
	isSet bool
}

func (v NullableSplitConfigurationLogic) Get() *SplitConfigurationLogic {
	return v.value
}

func (v *NullableSplitConfigurationLogic) Set(val *SplitConfigurationLogic) {
	v.value = val
	v.isSet = true
}

func (v NullableSplitConfigurationLogic) IsSet() bool {
	return v.isSet
}

func (v *NullableSplitConfigurationLogic) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSplitConfigurationLogic(val *SplitConfigurationLogic) *NullableSplitConfigurationLogic {
	return &NullableSplitConfigurationLogic{value: val, isSet: true}
}

func (v NullableSplitConfigurationLogic) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSplitConfigurationLogic) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

func (o *SplitConfigurationLogic) isValidAcquiringFees() bool {
	var allowedEnumValues = []string{"deductFromLiableAccount", "deductFromOneBalanceAccount"}
	for _, allowed := range allowedEnumValues {
		if o.GetAcquiringFees() == allowed {
			return true
		}
	}
	return false
}
func (o *SplitConfigurationLogic) isValidAdyenCommission() bool {
	var allowedEnumValues = []string{"deductFromLiableAccount", "deductFromOneBalanceAccount"}
	for _, allowed := range allowedEnumValues {
		if o.GetAdyenCommission() == allowed {
			return true
		}
	}
	return false
}
func (o *SplitConfigurationLogic) isValidAdyenFees() bool {
	var allowedEnumValues = []string{"deductFromLiableAccount", "deductFromOneBalanceAccount"}
	for _, allowed := range allowedEnumValues {
		if o.GetAdyenFees() == allowed {
			return true
		}
	}
	return false
}
func (o *SplitConfigurationLogic) isValidAdyenMarkup() bool {
	var allowedEnumValues = []string{"deductFromLiableAccount", "deductFromOneBalanceAccount"}
	for _, allowed := range allowedEnumValues {
		if o.GetAdyenMarkup() == allowed {
			return true
		}
	}
	return false
}
func (o *SplitConfigurationLogic) isValidChargeback() bool {
	var allowedEnumValues = []string{"deductFromLiableAccount", "deductFromOneBalanceAccount", "deductAccordingToSplitRatio"}
	for _, allowed := range allowedEnumValues {
		if o.GetChargeback() == allowed {
			return true
		}
	}
	return false
}
func (o *SplitConfigurationLogic) isValidChargebackCostAllocation() bool {
	var allowedEnumValues = []string{"deductFromLiableAccount", "deductFromOneBalanceAccount"}
	for _, allowed := range allowedEnumValues {
		if o.GetChargebackCostAllocation() == allowed {
			return true
		}
	}
	return false
}
func (o *SplitConfigurationLogic) isValidInterchange() bool {
	var allowedEnumValues = []string{"deductFromLiableAccount", "deductFromOneBalanceAccount"}
	for _, allowed := range allowedEnumValues {
		if o.GetInterchange() == allowed {
			return true
		}
	}
	return false
}
func (o *SplitConfigurationLogic) isValidPaymentFee() bool {
	var allowedEnumValues = []string{"deductFromLiableAccount", "deductFromOneBalanceAccount"}
	for _, allowed := range allowedEnumValues {
		if o.GetPaymentFee() == allowed {
			return true
		}
	}
	return false
}
func (o *SplitConfigurationLogic) isValidRemainder() bool {
	var allowedEnumValues = []string{"addToLiableAccount", "addToOneBalanceAccount"}
	for _, allowed := range allowedEnumValues {
		if o.GetRemainder() == allowed {
			return true
		}
	}
	return false
}
func (o *SplitConfigurationLogic) isValidSchemeFee() bool {
	var allowedEnumValues = []string{"deductFromLiableAccount", "deductFromOneBalanceAccount"}
	for _, allowed := range allowedEnumValues {
		if o.GetSchemeFee() == allowed {
			return true
		}
	}
	return false
}
func (o *SplitConfigurationLogic) isValidSurcharge() bool {
	var allowedEnumValues = []string{"addToLiableAccount", "addToOneBalanceAccount"}
	for _, allowed := range allowedEnumValues {
		if o.GetSurcharge() == allowed {
			return true
		}
	}
	return false
}
func (o *SplitConfigurationLogic) isValidTip() bool {
	var allowedEnumValues = []string{"addToLiableAccount", "addToOneBalanceAccount"}
	for _, allowed := range allowedEnumValues {
		if o.GetTip() == allowed {
			return true
		}
	}
	return false
}