/*
Transfers API

API version: 4
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package transfers

import (
	"encoding/json"

	"github.com/adyen/adyen-go-api-library/v16/src/common"
)

// checks if the TransferInfo type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &TransferInfo{}

// TransferInfo struct for TransferInfo
type TransferInfo struct {
	Amount Amount `json:"amount"`
	// The unique identifier of the source [balance account](https://docs.adyen.com/api-explorer/balanceplatform/latest/post/balanceAccounts#responses-200-id).  If you want to make a transfer using a **virtual** **bankAccount** assigned to the balance account, you must specify the [payment instrument ID](https://docs.adyen.com/api-explorer/balanceplatform/latest/post/paymentInstruments#responses-200-id) of the **virtual** **bankAccount**. If you only specify a balance account ID, Adyen uses the default **physical** **bankAccount** payment instrument assigned to the balance account.
	BalanceAccountId *string `json:"balanceAccountId,omitempty"`
	// The category of the transfer.  Possible values:   - **bank**: a transfer involving a [transfer instrument](https://docs.adyen.com/api-explorer/#/legalentity/latest/post/transferInstruments__resParam_id) or a bank account.  - **card**: a transfer involving a third-party card.  - **internal**: a transfer between [balance accounts](https://docs.adyen.com/api-explorer/#/balanceplatform/latest/post/balanceAccounts__resParam_id) within your platform.  - **issuedCard**: a transfer initiated by a Adyen-issued card.  - **platformPayment**: funds movements related to payments that are acquired for your users.  - **topUp**: an incoming transfer initiated by your user to top up their balance account.
	Category     string             `json:"category"`
	Counterparty CounterpartyInfoV3 `json:"counterparty"`
	// Your description for the transfer. It is used by most banks as the transfer description. We recommend sending a maximum of 140 characters, otherwise the description may be truncated.  Supported characters: **[a-z] [A-Z] [0-9] / - ?** **: ( ) . , ' + Space**  Supported characters for **regular** and **fast** transfers to a US counterparty: **[a-z] [A-Z] [0-9] & $ % # @** **~ = + - _ ' \" ! ?**
	Description *string `json:"description,omitempty"`
	// The unique identifier of the source [payment instrument](https://docs.adyen.com/api-explorer/balanceplatform/latest/post/paymentInstruments#responses-200-id).  If you want to make a transfer using a **virtual** **bankAccount**, you must specify the payment instrument ID of the **virtual** **bankAccount**. If you only specify a balance account ID, Adyen uses the default **physical** **bankAccount** payment instrument assigned to the balance account.
	PaymentInstrumentId *string `json:"paymentInstrumentId,omitempty"`
	//  The list of priorities for the bank transfer. This sets the speed at which the transfer is sent and the fees that you have to pay. You can provide multiple priorities. Adyen will try to pay out using the priority you list first. If that's not possible, it moves on to the next option in the order of your provided priorities.   Possible values:  * **regular**: for normal, low-value transactions.  * **fast**: a faster way to transfer funds, but the fees are higher. Recommended for high-priority, low-value transactions.  * **wire**: the fastest way to transfer funds, but this has the highest fees. Recommended for high-priority, high-value transactions.  * **instant**: for instant funds transfers in [SEPA countries](https://www.ecb.europa.eu/paym/integration/retail/sepa/html/index.en.html).  * **crossBorder**: for high-value transfers to a recipient in a different country.  * **internal**: for transfers to an Adyen-issued business bank account (by bank account number/IBAN).  Required for transfers with `category` **bank**. For more details, see [fallback priorities](https://docs.adyen.com/payouts/payout-service/payout-to-users/#fallback-priorities).
	Priorities []string `json:"priorities,omitempty"`
	// The priority for the bank transfer. This sets the speed at which the transfer is sent and the fees that you have to pay. Required for transfers with `category` **bank**.  Possible values:  * **regular**: for normal, low-value transactions.  * **fast**: a faster way to transfer funds, but the fees are higher. Recommended for high-priority, low-value transactions.  * **wire**: the fastest way to transfer funds, but this has the highest fees. Recommended for high-priority, high-value transactions.  * **instant**: for instant funds transfers in [SEPA countries](https://www.ecb.europa.eu/paym/integration/retail/sepa/html/index.en.html).  * **crossBorder**: for high-value transfers to a recipient in a different country.  * **internal**: for transfers to an Adyen-issued business bank account (by bank account number/IBAN).
	Priority *string `json:"priority,omitempty"`
	// Your reference for the transfer, used internally within your platform. If you don't provide this in the request, Adyen generates a unique reference.
	Reference *string `json:"reference,omitempty"`
	//  A reference that is sent to the recipient. This reference is also sent in all webhooks related to the transfer, so you can use it to track statuses for both parties involved in the funds movement.   Supported characters: **a-z**, **A-Z**, **0-9**. The maximum length depends on the `category`.  - **internal**: 80 characters  - **bank**: 35 characters when transferring to an IBAN, 15 characters for others.
	ReferenceForBeneficiary *string                `json:"referenceForBeneficiary,omitempty"`
	Review                  *TransferRequestReview `json:"review,omitempty"`
	// The type of transfer.  Possible values:   - **bankTransfer**: for push transfers to a transfer instrument or a bank account. The `category` must be **bank**. - **internalTransfer**: for push transfers between balance accounts. The `category` must be **internal**. - **internalDirectDebit**: for pull transfers (direct debits) between balance accounts. The `category` must be **internal**.
	Type          *string                      `json:"type,omitempty"`
	UltimateParty *UltimatePartyIdentification `json:"ultimateParty,omitempty"`
}

// NewTransferInfo instantiates a new TransferInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTransferInfo(amount Amount, category string, counterparty CounterpartyInfoV3) *TransferInfo {
	this := TransferInfo{}
	this.Amount = amount
	this.Category = category
	this.Counterparty = counterparty
	return &this
}

// NewTransferInfoWithDefaults instantiates a new TransferInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTransferInfoWithDefaults() *TransferInfo {
	this := TransferInfo{}
	return &this
}

// GetAmount returns the Amount field value
func (o *TransferInfo) GetAmount() Amount {
	if o == nil {
		var ret Amount
		return ret
	}

	return o.Amount
}

// GetAmountOk returns a tuple with the Amount field value
// and a boolean to check if the value has been set.
func (o *TransferInfo) GetAmountOk() (*Amount, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Amount, true
}

// SetAmount sets field value
func (o *TransferInfo) SetAmount(v Amount) {
	o.Amount = v
}

// GetBalanceAccountId returns the BalanceAccountId field value if set, zero value otherwise.
func (o *TransferInfo) GetBalanceAccountId() string {
	if o == nil || common.IsNil(o.BalanceAccountId) {
		var ret string
		return ret
	}
	return *o.BalanceAccountId
}

// GetBalanceAccountIdOk returns a tuple with the BalanceAccountId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransferInfo) GetBalanceAccountIdOk() (*string, bool) {
	if o == nil || common.IsNil(o.BalanceAccountId) {
		return nil, false
	}
	return o.BalanceAccountId, true
}

// HasBalanceAccountId returns a boolean if a field has been set.
func (o *TransferInfo) HasBalanceAccountId() bool {
	if o != nil && !common.IsNil(o.BalanceAccountId) {
		return true
	}

	return false
}

// SetBalanceAccountId gets a reference to the given string and assigns it to the BalanceAccountId field.
func (o *TransferInfo) SetBalanceAccountId(v string) {
	o.BalanceAccountId = &v
}

// GetCategory returns the Category field value
func (o *TransferInfo) GetCategory() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Category
}

// GetCategoryOk returns a tuple with the Category field value
// and a boolean to check if the value has been set.
func (o *TransferInfo) GetCategoryOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Category, true
}

// SetCategory sets field value
func (o *TransferInfo) SetCategory(v string) {
	o.Category = v
}

// GetCounterparty returns the Counterparty field value
func (o *TransferInfo) GetCounterparty() CounterpartyInfoV3 {
	if o == nil {
		var ret CounterpartyInfoV3
		return ret
	}

	return o.Counterparty
}

// GetCounterpartyOk returns a tuple with the Counterparty field value
// and a boolean to check if the value has been set.
func (o *TransferInfo) GetCounterpartyOk() (*CounterpartyInfoV3, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Counterparty, true
}

// SetCounterparty sets field value
func (o *TransferInfo) SetCounterparty(v CounterpartyInfoV3) {
	o.Counterparty = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *TransferInfo) GetDescription() string {
	if o == nil || common.IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransferInfo) GetDescriptionOk() (*string, bool) {
	if o == nil || common.IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *TransferInfo) HasDescription() bool {
	if o != nil && !common.IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *TransferInfo) SetDescription(v string) {
	o.Description = &v
}

// GetPaymentInstrumentId returns the PaymentInstrumentId field value if set, zero value otherwise.
func (o *TransferInfo) GetPaymentInstrumentId() string {
	if o == nil || common.IsNil(o.PaymentInstrumentId) {
		var ret string
		return ret
	}
	return *o.PaymentInstrumentId
}

// GetPaymentInstrumentIdOk returns a tuple with the PaymentInstrumentId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransferInfo) GetPaymentInstrumentIdOk() (*string, bool) {
	if o == nil || common.IsNil(o.PaymentInstrumentId) {
		return nil, false
	}
	return o.PaymentInstrumentId, true
}

// HasPaymentInstrumentId returns a boolean if a field has been set.
func (o *TransferInfo) HasPaymentInstrumentId() bool {
	if o != nil && !common.IsNil(o.PaymentInstrumentId) {
		return true
	}

	return false
}

// SetPaymentInstrumentId gets a reference to the given string and assigns it to the PaymentInstrumentId field.
func (o *TransferInfo) SetPaymentInstrumentId(v string) {
	o.PaymentInstrumentId = &v
}

// GetPriorities returns the Priorities field value if set, zero value otherwise.
func (o *TransferInfo) GetPriorities() []string {
	if o == nil || common.IsNil(o.Priorities) {
		var ret []string
		return ret
	}
	return o.Priorities
}

// GetPrioritiesOk returns a tuple with the Priorities field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransferInfo) GetPrioritiesOk() ([]string, bool) {
	if o == nil || common.IsNil(o.Priorities) {
		return nil, false
	}
	return o.Priorities, true
}

// HasPriorities returns a boolean if a field has been set.
func (o *TransferInfo) HasPriorities() bool {
	if o != nil && !common.IsNil(o.Priorities) {
		return true
	}

	return false
}

// SetPriorities gets a reference to the given []string and assigns it to the Priorities field.
func (o *TransferInfo) SetPriorities(v []string) {
	o.Priorities = v
}

// GetPriority returns the Priority field value if set, zero value otherwise.
func (o *TransferInfo) GetPriority() string {
	if o == nil || common.IsNil(o.Priority) {
		var ret string
		return ret
	}
	return *o.Priority
}

// GetPriorityOk returns a tuple with the Priority field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransferInfo) GetPriorityOk() (*string, bool) {
	if o == nil || common.IsNil(o.Priority) {
		return nil, false
	}
	return o.Priority, true
}

// HasPriority returns a boolean if a field has been set.
func (o *TransferInfo) HasPriority() bool {
	if o != nil && !common.IsNil(o.Priority) {
		return true
	}

	return false
}

// SetPriority gets a reference to the given string and assigns it to the Priority field.
func (o *TransferInfo) SetPriority(v string) {
	o.Priority = &v
}

// GetReference returns the Reference field value if set, zero value otherwise.
func (o *TransferInfo) GetReference() string {
	if o == nil || common.IsNil(o.Reference) {
		var ret string
		return ret
	}
	return *o.Reference
}

// GetReferenceOk returns a tuple with the Reference field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransferInfo) GetReferenceOk() (*string, bool) {
	if o == nil || common.IsNil(o.Reference) {
		return nil, false
	}
	return o.Reference, true
}

// HasReference returns a boolean if a field has been set.
func (o *TransferInfo) HasReference() bool {
	if o != nil && !common.IsNil(o.Reference) {
		return true
	}

	return false
}

// SetReference gets a reference to the given string and assigns it to the Reference field.
func (o *TransferInfo) SetReference(v string) {
	o.Reference = &v
}

// GetReferenceForBeneficiary returns the ReferenceForBeneficiary field value if set, zero value otherwise.
func (o *TransferInfo) GetReferenceForBeneficiary() string {
	if o == nil || common.IsNil(o.ReferenceForBeneficiary) {
		var ret string
		return ret
	}
	return *o.ReferenceForBeneficiary
}

// GetReferenceForBeneficiaryOk returns a tuple with the ReferenceForBeneficiary field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransferInfo) GetReferenceForBeneficiaryOk() (*string, bool) {
	if o == nil || common.IsNil(o.ReferenceForBeneficiary) {
		return nil, false
	}
	return o.ReferenceForBeneficiary, true
}

// HasReferenceForBeneficiary returns a boolean if a field has been set.
func (o *TransferInfo) HasReferenceForBeneficiary() bool {
	if o != nil && !common.IsNil(o.ReferenceForBeneficiary) {
		return true
	}

	return false
}

// SetReferenceForBeneficiary gets a reference to the given string and assigns it to the ReferenceForBeneficiary field.
func (o *TransferInfo) SetReferenceForBeneficiary(v string) {
	o.ReferenceForBeneficiary = &v
}

// GetReview returns the Review field value if set, zero value otherwise.
func (o *TransferInfo) GetReview() TransferRequestReview {
	if o == nil || common.IsNil(o.Review) {
		var ret TransferRequestReview
		return ret
	}
	return *o.Review
}

// GetReviewOk returns a tuple with the Review field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransferInfo) GetReviewOk() (*TransferRequestReview, bool) {
	if o == nil || common.IsNil(o.Review) {
		return nil, false
	}
	return o.Review, true
}

// HasReview returns a boolean if a field has been set.
func (o *TransferInfo) HasReview() bool {
	if o != nil && !common.IsNil(o.Review) {
		return true
	}

	return false
}

// SetReview gets a reference to the given TransferRequestReview and assigns it to the Review field.
func (o *TransferInfo) SetReview(v TransferRequestReview) {
	o.Review = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *TransferInfo) GetType() string {
	if o == nil || common.IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransferInfo) GetTypeOk() (*string, bool) {
	if o == nil || common.IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *TransferInfo) HasType() bool {
	if o != nil && !common.IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *TransferInfo) SetType(v string) {
	o.Type = &v
}

// GetUltimateParty returns the UltimateParty field value if set, zero value otherwise.
func (o *TransferInfo) GetUltimateParty() UltimatePartyIdentification {
	if o == nil || common.IsNil(o.UltimateParty) {
		var ret UltimatePartyIdentification
		return ret
	}
	return *o.UltimateParty
}

// GetUltimatePartyOk returns a tuple with the UltimateParty field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransferInfo) GetUltimatePartyOk() (*UltimatePartyIdentification, bool) {
	if o == nil || common.IsNil(o.UltimateParty) {
		return nil, false
	}
	return o.UltimateParty, true
}

// HasUltimateParty returns a boolean if a field has been set.
func (o *TransferInfo) HasUltimateParty() bool {
	if o != nil && !common.IsNil(o.UltimateParty) {
		return true
	}

	return false
}

// SetUltimateParty gets a reference to the given UltimatePartyIdentification and assigns it to the UltimateParty field.
func (o *TransferInfo) SetUltimateParty(v UltimatePartyIdentification) {
	o.UltimateParty = &v
}

func (o TransferInfo) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TransferInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["amount"] = o.Amount
	if !common.IsNil(o.BalanceAccountId) {
		toSerialize["balanceAccountId"] = o.BalanceAccountId
	}
	toSerialize["category"] = o.Category
	toSerialize["counterparty"] = o.Counterparty
	if !common.IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !common.IsNil(o.PaymentInstrumentId) {
		toSerialize["paymentInstrumentId"] = o.PaymentInstrumentId
	}
	if !common.IsNil(o.Priorities) {
		toSerialize["priorities"] = o.Priorities
	}
	if !common.IsNil(o.Priority) {
		toSerialize["priority"] = o.Priority
	}
	if !common.IsNil(o.Reference) {
		toSerialize["reference"] = o.Reference
	}
	if !common.IsNil(o.ReferenceForBeneficiary) {
		toSerialize["referenceForBeneficiary"] = o.ReferenceForBeneficiary
	}
	if !common.IsNil(o.Review) {
		toSerialize["review"] = o.Review
	}
	if !common.IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !common.IsNil(o.UltimateParty) {
		toSerialize["ultimateParty"] = o.UltimateParty
	}
	return toSerialize, nil
}

type NullableTransferInfo struct {
	value *TransferInfo
	isSet bool
}

func (v NullableTransferInfo) Get() *TransferInfo {
	return v.value
}

func (v *NullableTransferInfo) Set(val *TransferInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableTransferInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableTransferInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTransferInfo(val *TransferInfo) *NullableTransferInfo {
	return &NullableTransferInfo{value: val, isSet: true}
}

func (v NullableTransferInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTransferInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

func (o *TransferInfo) isValidCategory() bool {
	var allowedEnumValues = []string{"bank", "card", "internal", "issuedCard", "platformPayment", "topUp"}
	for _, allowed := range allowedEnumValues {
		if o.GetCategory() == allowed {
			return true
		}
	}
	return false
}
func (o *TransferInfo) isValidPriority() bool {
	var allowedEnumValues = []string{"crossBorder", "fast", "instant", "internal", "regular", "wire"}
	for _, allowed := range allowedEnumValues {
		if o.GetPriority() == allowed {
			return true
		}
	}
	return false
}
func (o *TransferInfo) isValidType() bool {
	var allowedEnumValues = []string{"bankTransfer", "internalTransfer", "internalDirectDebit"}
	for _, allowed := range allowedEnumValues {
		if o.GetType() == allowed {
			return true
		}
	}
	return false
}
