/*
Legal Entity Management API

API version: 3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package legalentity

import (
	"context"
	"net/http"
	"net/url"
	"strings"

	"github.com/adyen/adyen-go-api-library/v16/src/common"
)

// TaxEDeliveryConsentApi service
type TaxEDeliveryConsentApi common.Service

// All parameters accepted by TaxEDeliveryConsentApi.CheckStatusOfConsentForElectronicDeliveryOfTaxForms
type TaxEDeliveryConsentApiCheckStatusOfConsentForElectronicDeliveryOfTaxFormsInput struct {
	id string
}

/*
Prepare a request for CheckStatusOfConsentForElectronicDeliveryOfTaxForms
@param id The unique identifier of the legal entity. For sole proprietorships, this is the individual legal entity ID of the owner. For organizations, this is the ID of the organization.
@return TaxEDeliveryConsentApiCheckStatusOfConsentForElectronicDeliveryOfTaxFormsInput
*/
func (a *TaxEDeliveryConsentApi) CheckStatusOfConsentForElectronicDeliveryOfTaxFormsInput(id string) TaxEDeliveryConsentApiCheckStatusOfConsentForElectronicDeliveryOfTaxFormsInput {
	return TaxEDeliveryConsentApiCheckStatusOfConsentForElectronicDeliveryOfTaxFormsInput{
		id: id,
	}
}

/*
CheckStatusOfConsentForElectronicDeliveryOfTaxForms Check the status of consent for electronic delivery of tax forms

Returns the consent status for electronic delivery of tax forms.

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param r TaxEDeliveryConsentApiCheckStatusOfConsentForElectronicDeliveryOfTaxFormsInput - Request parameters, see CheckStatusOfConsentForElectronicDeliveryOfTaxFormsInput
@return CheckTaxElectronicDeliveryConsentResponse, *http.Response, error
*/
func (a *TaxEDeliveryConsentApi) CheckStatusOfConsentForElectronicDeliveryOfTaxForms(ctx context.Context, r TaxEDeliveryConsentApiCheckStatusOfConsentForElectronicDeliveryOfTaxFormsInput) (CheckTaxElectronicDeliveryConsentResponse, *http.Response, error) {
	res := &CheckTaxElectronicDeliveryConsentResponse{}
	path := "/legalEntities/{id}/checkTaxElectronicDeliveryConsent"
	path = strings.Replace(path, "{"+"id"+"}", url.PathEscape(common.ParameterValueToString(r.id, "id")), -1)
	queryParams := url.Values{}
	headerParams := make(map[string]string)
	httpRes, err := common.SendAPIRequest(
		ctx,
		a.Client,
		nil,
		res,
		http.MethodPost,
		a.BasePath()+path,
		queryParams,
		headerParams,
	)

	return *res, httpRes, err
}

// All parameters accepted by TaxEDeliveryConsentApi.SetConsentStatusForElectronicDeliveryOfTaxForms
type TaxEDeliveryConsentApiSetConsentStatusForElectronicDeliveryOfTaxFormsInput struct {
	id                                     string
	setTaxElectronicDeliveryConsentRequest *SetTaxElectronicDeliveryConsentRequest
}

func (r TaxEDeliveryConsentApiSetConsentStatusForElectronicDeliveryOfTaxFormsInput) SetTaxElectronicDeliveryConsentRequest(setTaxElectronicDeliveryConsentRequest SetTaxElectronicDeliveryConsentRequest) TaxEDeliveryConsentApiSetConsentStatusForElectronicDeliveryOfTaxFormsInput {
	r.setTaxElectronicDeliveryConsentRequest = &setTaxElectronicDeliveryConsentRequest
	return r
}

/*
Prepare a request for SetConsentStatusForElectronicDeliveryOfTaxForms
@param id The unique identifier of the legal entity. For sole proprietorships, this is the individual legal entity ID of the owner. For organizations, this is the ID of the organization.
@return TaxEDeliveryConsentApiSetConsentStatusForElectronicDeliveryOfTaxFormsInput
*/
func (a *TaxEDeliveryConsentApi) SetConsentStatusForElectronicDeliveryOfTaxFormsInput(id string) TaxEDeliveryConsentApiSetConsentStatusForElectronicDeliveryOfTaxFormsInput {
	return TaxEDeliveryConsentApiSetConsentStatusForElectronicDeliveryOfTaxFormsInput{
		id: id,
	}
}

/*
SetConsentStatusForElectronicDeliveryOfTaxForms Set the consent status for electronic delivery of tax forms

Set the consent status for electronic delivery of tax forms.

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param r TaxEDeliveryConsentApiSetConsentStatusForElectronicDeliveryOfTaxFormsInput - Request parameters, see SetConsentStatusForElectronicDeliveryOfTaxFormsInput
@return *http.Response, error
*/
func (a *TaxEDeliveryConsentApi) SetConsentStatusForElectronicDeliveryOfTaxForms(ctx context.Context, r TaxEDeliveryConsentApiSetConsentStatusForElectronicDeliveryOfTaxFormsInput) (*http.Response, error) {
	var res interface{}
	path := "/legalEntities/{id}/setTaxElectronicDeliveryConsent"
	path = strings.Replace(path, "{"+"id"+"}", url.PathEscape(common.ParameterValueToString(r.id, "id")), -1)
	queryParams := url.Values{}
	headerParams := make(map[string]string)
	httpRes, err := common.SendAPIRequest(
		ctx,
		a.Client,
		r.setTaxElectronicDeliveryConsentRequest,
		res,
		http.MethodPost,
		a.BasePath()+path,
		queryParams,
		headerParams,
	)

	return httpRes, err
}
