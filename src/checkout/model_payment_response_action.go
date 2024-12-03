/*
Adyen Checkout API

API version: 71
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package checkout

import (
	"encoding/json"
	"fmt"
)

// PaymentResponseAction - Action to be taken for completing the payment.
type PaymentResponseAction struct {
	CheckoutAwaitAction                   *CheckoutAwaitAction
	CheckoutBankTransferAction            *CheckoutBankTransferAction
	CheckoutDelegatedAuthenticationAction *CheckoutDelegatedAuthenticationAction
	CheckoutNativeRedirectAction          *CheckoutNativeRedirectAction
	CheckoutQrCodeAction                  *CheckoutQrCodeAction
	CheckoutRedirectAction                *CheckoutRedirectAction
	CheckoutSDKAction                     *CheckoutSDKAction
	CheckoutThreeDS2Action                *CheckoutThreeDS2Action
	CheckoutVoucherAction                 *CheckoutVoucherAction
}

// CheckoutAwaitActionAsPaymentResponseAction is a convenience function that returns CheckoutAwaitAction wrapped in PaymentResponseAction
func CheckoutAwaitActionAsPaymentResponseAction(v *CheckoutAwaitAction) PaymentResponseAction {
	return PaymentResponseAction{
		CheckoutAwaitAction: v,
	}
}

// CheckoutBankTransferActionAsPaymentResponseAction is a convenience function that returns CheckoutBankTransferAction wrapped in PaymentResponseAction
func CheckoutBankTransferActionAsPaymentResponseAction(v *CheckoutBankTransferAction) PaymentResponseAction {
	return PaymentResponseAction{
		CheckoutBankTransferAction: v,
	}
}

// CheckoutDelegatedAuthenticationActionAsPaymentResponseAction is a convenience function that returns CheckoutDelegatedAuthenticationAction wrapped in PaymentResponseAction
func CheckoutDelegatedAuthenticationActionAsPaymentResponseAction(v *CheckoutDelegatedAuthenticationAction) PaymentResponseAction {
	return PaymentResponseAction{
		CheckoutDelegatedAuthenticationAction: v,
	}
}

// CheckoutNativeRedirectActionAsPaymentResponseAction is a convenience function that returns CheckoutNativeRedirectAction wrapped in PaymentResponseAction
func CheckoutNativeRedirectActionAsPaymentResponseAction(v *CheckoutNativeRedirectAction) PaymentResponseAction {
	return PaymentResponseAction{
		CheckoutNativeRedirectAction: v,
	}
}

// CheckoutQrCodeActionAsPaymentResponseAction is a convenience function that returns CheckoutQrCodeAction wrapped in PaymentResponseAction
func CheckoutQrCodeActionAsPaymentResponseAction(v *CheckoutQrCodeAction) PaymentResponseAction {
	return PaymentResponseAction{
		CheckoutQrCodeAction: v,
	}
}

// CheckoutRedirectActionAsPaymentResponseAction is a convenience function that returns CheckoutRedirectAction wrapped in PaymentResponseAction
func CheckoutRedirectActionAsPaymentResponseAction(v *CheckoutRedirectAction) PaymentResponseAction {
	return PaymentResponseAction{
		CheckoutRedirectAction: v,
	}
}

// CheckoutSDKActionAsPaymentResponseAction is a convenience function that returns CheckoutSDKAction wrapped in PaymentResponseAction
func CheckoutSDKActionAsPaymentResponseAction(v *CheckoutSDKAction) PaymentResponseAction {
	return PaymentResponseAction{
		CheckoutSDKAction: v,
	}
}

// CheckoutThreeDS2ActionAsPaymentResponseAction is a convenience function that returns CheckoutThreeDS2Action wrapped in PaymentResponseAction
func CheckoutThreeDS2ActionAsPaymentResponseAction(v *CheckoutThreeDS2Action) PaymentResponseAction {
	return PaymentResponseAction{
		CheckoutThreeDS2Action: v,
	}
}

// CheckoutVoucherActionAsPaymentResponseAction is a convenience function that returns CheckoutVoucherAction wrapped in PaymentResponseAction
func CheckoutVoucherActionAsPaymentResponseAction(v *CheckoutVoucherAction) PaymentResponseAction {
	return PaymentResponseAction{
		CheckoutVoucherAction: v,
	}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *PaymentResponseAction) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into CheckoutAwaitAction
	err = json.Unmarshal(data, &dst.CheckoutAwaitAction)
	if err == nil {
		jsonCheckoutAwaitAction, _ := json.Marshal(dst.CheckoutAwaitAction)
		if string(jsonCheckoutAwaitAction) == "{}" || !dst.CheckoutAwaitAction.isValidType() { // empty struct
			dst.CheckoutAwaitAction = nil
		} else {
			match++
		}
	} else {
		dst.CheckoutAwaitAction = nil
	}

	// try to unmarshal data into CheckoutBankTransferAction
	err = json.Unmarshal(data, &dst.CheckoutBankTransferAction)
	if err == nil {
		jsonCheckoutBankTransferAction, _ := json.Marshal(dst.CheckoutBankTransferAction)
		if string(jsonCheckoutBankTransferAction) == "{}" || !dst.CheckoutBankTransferAction.isValidType() { // empty struct
			dst.CheckoutBankTransferAction = nil
		} else {
			match++
		}
	} else {
		dst.CheckoutBankTransferAction = nil
	}

	// try to unmarshal data into CheckoutDelegatedAuthenticationAction
	err = json.Unmarshal(data, &dst.CheckoutDelegatedAuthenticationAction)
	if err == nil {
		jsonCheckoutDelegatedAuthenticationAction, _ := json.Marshal(dst.CheckoutDelegatedAuthenticationAction)
		if string(jsonCheckoutDelegatedAuthenticationAction) == "{}" || !dst.CheckoutDelegatedAuthenticationAction.isValidType() { // empty struct
			dst.CheckoutDelegatedAuthenticationAction = nil
		} else {
			match++
		}
	} else {
		dst.CheckoutDelegatedAuthenticationAction = nil
	}

	// try to unmarshal data into CheckoutNativeRedirectAction
	err = json.Unmarshal(data, &dst.CheckoutNativeRedirectAction)
	if err == nil {
		jsonCheckoutNativeRedirectAction, _ := json.Marshal(dst.CheckoutNativeRedirectAction)
		if string(jsonCheckoutNativeRedirectAction) == "{}" || !dst.CheckoutNativeRedirectAction.isValidType() { // empty struct
			dst.CheckoutNativeRedirectAction = nil
		} else {
			match++
		}
	} else {
		dst.CheckoutNativeRedirectAction = nil
	}

	// try to unmarshal data into CheckoutQrCodeAction
	err = json.Unmarshal(data, &dst.CheckoutQrCodeAction)
	if err == nil {
		jsonCheckoutQrCodeAction, _ := json.Marshal(dst.CheckoutQrCodeAction)
		if string(jsonCheckoutQrCodeAction) == "{}" || !dst.CheckoutQrCodeAction.isValidType() { // empty struct
			dst.CheckoutQrCodeAction = nil
		} else {
			match++
		}
	} else {
		dst.CheckoutQrCodeAction = nil
	}

	// try to unmarshal data into CheckoutRedirectAction
	err = json.Unmarshal(data, &dst.CheckoutRedirectAction)
	if err == nil {
		jsonCheckoutRedirectAction, _ := json.Marshal(dst.CheckoutRedirectAction)
		if string(jsonCheckoutRedirectAction) == "{}" || !dst.CheckoutRedirectAction.isValidType() { // empty struct
			dst.CheckoutRedirectAction = nil
		} else {
			match++
		}
	} else {
		dst.CheckoutRedirectAction = nil
	}

	// try to unmarshal data into CheckoutSDKAction
	err = json.Unmarshal(data, &dst.CheckoutSDKAction)
	if err == nil {
		jsonCheckoutSDKAction, _ := json.Marshal(dst.CheckoutSDKAction)
		if string(jsonCheckoutSDKAction) == "{}" || !dst.CheckoutSDKAction.isValidType() { // empty struct
			dst.CheckoutSDKAction = nil
		} else {
			match++
		}
	} else {
		dst.CheckoutSDKAction = nil
	}

	// try to unmarshal data into CheckoutThreeDS2Action
	err = json.Unmarshal(data, &dst.CheckoutThreeDS2Action)
	if err == nil {
		jsonCheckoutThreeDS2Action, _ := json.Marshal(dst.CheckoutThreeDS2Action)
		if string(jsonCheckoutThreeDS2Action) == "{}" || !dst.CheckoutThreeDS2Action.isValidType() { // empty struct
			dst.CheckoutThreeDS2Action = nil
		} else {
			match++
		}
	} else {
		dst.CheckoutThreeDS2Action = nil
	}

	// try to unmarshal data into CheckoutVoucherAction
	err = json.Unmarshal(data, &dst.CheckoutVoucherAction)
	if err == nil {
		jsonCheckoutVoucherAction, _ := json.Marshal(dst.CheckoutVoucherAction)
		if string(jsonCheckoutVoucherAction) == "{}" || !dst.CheckoutVoucherAction.isValidType() { // empty struct
			dst.CheckoutVoucherAction = nil
		} else {
			match++
		}
	} else {
		dst.CheckoutVoucherAction = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.CheckoutAwaitAction = nil
		dst.CheckoutBankTransferAction = nil
		dst.CheckoutDelegatedAuthenticationAction = nil
		dst.CheckoutNativeRedirectAction = nil
		dst.CheckoutQrCodeAction = nil
		dst.CheckoutRedirectAction = nil
		dst.CheckoutSDKAction = nil
		dst.CheckoutThreeDS2Action = nil
		dst.CheckoutVoucherAction = nil

		return fmt.Errorf("data matches more than one schema in oneOf(PaymentResponseAction)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(PaymentResponseAction)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src PaymentResponseAction) MarshalJSON() ([]byte, error) {
	if src.CheckoutAwaitAction != nil {
		return json.Marshal(&src.CheckoutAwaitAction)
	}

	if src.CheckoutBankTransferAction != nil {
		return json.Marshal(&src.CheckoutBankTransferAction)
	}

	if src.CheckoutDelegatedAuthenticationAction != nil {
		return json.Marshal(&src.CheckoutDelegatedAuthenticationAction)
	}

	if src.CheckoutNativeRedirectAction != nil {
		return json.Marshal(&src.CheckoutNativeRedirectAction)
	}

	if src.CheckoutQrCodeAction != nil {
		return json.Marshal(&src.CheckoutQrCodeAction)
	}

	if src.CheckoutRedirectAction != nil {
		return json.Marshal(&src.CheckoutRedirectAction)
	}

	if src.CheckoutSDKAction != nil {
		return json.Marshal(&src.CheckoutSDKAction)
	}

	if src.CheckoutThreeDS2Action != nil {
		return json.Marshal(&src.CheckoutThreeDS2Action)
	}

	if src.CheckoutVoucherAction != nil {
		return json.Marshal(&src.CheckoutVoucherAction)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *PaymentResponseAction) GetActualInstance() interface{} {
	if obj == nil {
		return nil
	}
	if obj.CheckoutAwaitAction != nil {
		return obj.CheckoutAwaitAction
	}

	if obj.CheckoutBankTransferAction != nil {
		return obj.CheckoutBankTransferAction
	}

	if obj.CheckoutDelegatedAuthenticationAction != nil {
		return obj.CheckoutDelegatedAuthenticationAction
	}

	if obj.CheckoutNativeRedirectAction != nil {
		return obj.CheckoutNativeRedirectAction
	}

	if obj.CheckoutQrCodeAction != nil {
		return obj.CheckoutQrCodeAction
	}

	if obj.CheckoutRedirectAction != nil {
		return obj.CheckoutRedirectAction
	}

	if obj.CheckoutSDKAction != nil {
		return obj.CheckoutSDKAction
	}

	if obj.CheckoutThreeDS2Action != nil {
		return obj.CheckoutThreeDS2Action
	}

	if obj.CheckoutVoucherAction != nil {
		return obj.CheckoutVoucherAction
	}

	// all schemas are nil
	return nil
}

type NullablePaymentResponseAction struct {
	value *PaymentResponseAction
	isSet bool
}

func (v NullablePaymentResponseAction) Get() *PaymentResponseAction {
	return v.value
}

func (v *NullablePaymentResponseAction) Set(val *PaymentResponseAction) {
	v.value = val
	v.isSet = true
}

func (v NullablePaymentResponseAction) IsSet() bool {
	return v.isSet
}

func (v *NullablePaymentResponseAction) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePaymentResponseAction(val *PaymentResponseAction) *NullablePaymentResponseAction {
	return &NullablePaymentResponseAction{value: val, isSet: true}
}

func (v NullablePaymentResponseAction) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePaymentResponseAction) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
