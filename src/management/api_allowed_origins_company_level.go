/*
Management API

API version: 1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package management

import (
	"context"
	"encoding/json"
	"io/ioutil"
	_nethttp "net/http"
	"net/url"
	"strings"

	"github.com/adyen/adyen-go-api-library/v7/src/common"
)

// AllowedOriginsCompanyLevelApi AllowedOriginsCompanyLevelApi service
type AllowedOriginsCompanyLevelApi common.Service

type AllowedOriginsCompanyLevelApiCreateAllowedOriginConfig struct {
	ctx             context.Context
	companyId       string
	apiCredentialId string
	allowedOrigin   *AllowedOrigin
}

func (r AllowedOriginsCompanyLevelApiCreateAllowedOriginConfig) AllowedOrigin(allowedOrigin AllowedOrigin) AllowedOriginsCompanyLevelApiCreateAllowedOriginConfig {
	r.allowedOrigin = &allowedOrigin
	return r
}

/*
CreateAllowedOrigin Create an allowed origin

Adds a new [allowed origin](https://docs.adyen.com/development-resources/client-side-authentication#allowed-origins) to the API credential's list of allowed origins.

To make this request, your API credential must have the following [roles](https://docs.adyen.com/development-resources/api-credentials#api-permissions):
* Management API—API credentials read and write

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param companyId The unique identifier of the company account.
	@param apiCredentialId Unique identifier of the API credential.
	@return AllowedOriginsCompanyLevelApiCreateAllowedOriginConfig
*/
func (a *AllowedOriginsCompanyLevelApi) CreateAllowedOriginConfig(ctx context.Context, companyId string, apiCredentialId string) AllowedOriginsCompanyLevelApiCreateAllowedOriginConfig {
	return AllowedOriginsCompanyLevelApiCreateAllowedOriginConfig{
		ctx:             ctx,
		companyId:       companyId,
		apiCredentialId: apiCredentialId,
	}
}

/*
Create an allowed origin
Adds a new [allowed origin](https://docs.adyen.com/development-resources/client-side-authentication#allowed-origins) to the API credential&#39;s list of allowed origins.  To make this request, your API credential must have the following [roles](https://docs.adyen.com/development-resources/api-credentials#api-permissions): * Management API—API credentials read and write
 * @param companyId The unique identifier of the company account.
 * @param apiCredentialId Unique identifier of the API credential.
 * @param req AllowedOrigin - reference of AllowedOrigin).
 * @param ctxs ...context.Context - optional, for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@return AllowedOriginsResponse
*/

func (a *AllowedOriginsCompanyLevelApi) CreateAllowedOrigin(r AllowedOriginsCompanyLevelApiCreateAllowedOriginConfig) (AllowedOriginsResponse, *_nethttp.Response, error) {
	var serviceError common.RestServiceError
	res := &AllowedOriginsResponse{}
	path := "/companies/{companyId}/apiCredentials/{apiCredentialId}/allowedOrigins"
	path = strings.Replace(path, "{"+"companyId"+"}", url.PathEscape(common.ParameterValueToString(r.companyId, "companyId")), -1)
	path = strings.Replace(path, "{"+"apiCredentialId"+"}", url.PathEscape(common.ParameterValueToString(r.apiCredentialId, "apiCredentialId")), -1)
	queryParams := url.Values{}
	headerParams := make(map[string]string)
	httpRes, err := common.SendAPIRequest(
		r.ctx,
		a.Client,
		r.allowedOrigin,
		res,
		_nethttp.MethodPost,
		a.BasePath()+path,
		queryParams,
		headerParams,
	)
	defer httpRes.Body.Close()

	if httpRes.StatusCode == 400 {
		// Read the response body
		body, _ := ioutil.ReadAll(httpRes.Body)
		_ = json.Unmarshal([]byte(body), &serviceError)
		return *res, httpRes, serviceError
	}

	if httpRes.StatusCode == 401 {
		// Read the response body
		body, _ := ioutil.ReadAll(httpRes.Body)
		_ = json.Unmarshal([]byte(body), &serviceError)
		return *res, httpRes, serviceError
	}

	if httpRes.StatusCode == 403 {
		// Read the response body
		body, _ := ioutil.ReadAll(httpRes.Body)
		_ = json.Unmarshal([]byte(body), &serviceError)
		return *res, httpRes, serviceError
	}

	if httpRes.StatusCode == 422 {
		// Read the response body
		body, _ := ioutil.ReadAll(httpRes.Body)
		_ = json.Unmarshal([]byte(body), &serviceError)
		return *res, httpRes, serviceError
	}

	if httpRes.StatusCode == 500 {
		// Read the response body
		body, _ := ioutil.ReadAll(httpRes.Body)
		_ = json.Unmarshal([]byte(body), &serviceError)
		return *res, httpRes, serviceError
	}
	return *res, httpRes, err
}

type AllowedOriginsCompanyLevelApiDeleteAllowedOriginConfig struct {
	ctx             context.Context
	companyId       string
	apiCredentialId string
	originId        string
}

/*
DeleteAllowedOrigin Delete an allowed origin

Removes the [allowed origin](https://docs.adyen.com/development-resources/client-side-authentication#allowed-origins) identified in the path. As soon as an allowed origin is removed, we no longer accept client-side requests from that domain.

To make this request, your API credential must have the following [roles](https://docs.adyen.com/development-resources/api-credentials#api-permissions):
* Management API—API credentials read and write

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param companyId The unique identifier of the company account.
	@param apiCredentialId Unique identifier of the API credential.
	@param originId Unique identifier of the allowed origin.
	@return AllowedOriginsCompanyLevelApiDeleteAllowedOriginConfig
*/
func (a *AllowedOriginsCompanyLevelApi) DeleteAllowedOriginConfig(ctx context.Context, companyId string, apiCredentialId string, originId string) AllowedOriginsCompanyLevelApiDeleteAllowedOriginConfig {
	return AllowedOriginsCompanyLevelApiDeleteAllowedOriginConfig{
		ctx:             ctx,
		companyId:       companyId,
		apiCredentialId: apiCredentialId,
		originId:        originId,
	}
}

/*
Delete an allowed origin
Removes the [allowed origin](https://docs.adyen.com/development-resources/client-side-authentication#allowed-origins) identified in the path. As soon as an allowed origin is removed, we no longer accept client-side requests from that domain.  To make this request, your API credential must have the following [roles](https://docs.adyen.com/development-resources/api-credentials#api-permissions): * Management API—API credentials read and write
 * @param companyId The unique identifier of the company account.
 * @param apiCredentialId Unique identifier of the API credential.
 * @param originId Unique identifier of the allowed origin.
 * @param ctxs ...context.Context - optional, for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
*/

func (a *AllowedOriginsCompanyLevelApi) DeleteAllowedOrigin(r AllowedOriginsCompanyLevelApiDeleteAllowedOriginConfig) (*_nethttp.Response, error) {
	var serviceError common.RestServiceError
	var res interface{}
	path := "/companies/{companyId}/apiCredentials/{apiCredentialId}/allowedOrigins/{originId}"
	path = strings.Replace(path, "{"+"companyId"+"}", url.PathEscape(common.ParameterValueToString(r.companyId, "companyId")), -1)
	path = strings.Replace(path, "{"+"apiCredentialId"+"}", url.PathEscape(common.ParameterValueToString(r.apiCredentialId, "apiCredentialId")), -1)
	path = strings.Replace(path, "{"+"originId"+"}", url.PathEscape(common.ParameterValueToString(r.originId, "originId")), -1)
	queryParams := url.Values{}
	headerParams := make(map[string]string)
	httpRes, err := common.SendAPIRequest(
		r.ctx,
		a.Client,
		nil,
		res,
		_nethttp.MethodDelete,
		a.BasePath()+path,
		queryParams,
		headerParams,
	)
	defer httpRes.Body.Close()

	if httpRes.StatusCode == 400 {
		// Read the response body
		body, _ := ioutil.ReadAll(httpRes.Body)
		_ = json.Unmarshal([]byte(body), &serviceError)
		return httpRes, serviceError
	}

	if httpRes.StatusCode == 401 {
		// Read the response body
		body, _ := ioutil.ReadAll(httpRes.Body)
		_ = json.Unmarshal([]byte(body), &serviceError)
		return httpRes, serviceError
	}

	if httpRes.StatusCode == 403 {
		// Read the response body
		body, _ := ioutil.ReadAll(httpRes.Body)
		_ = json.Unmarshal([]byte(body), &serviceError)
		return httpRes, serviceError
	}

	if httpRes.StatusCode == 422 {
		// Read the response body
		body, _ := ioutil.ReadAll(httpRes.Body)
		_ = json.Unmarshal([]byte(body), &serviceError)
		return httpRes, serviceError
	}

	if httpRes.StatusCode == 500 {
		// Read the response body
		body, _ := ioutil.ReadAll(httpRes.Body)
		_ = json.Unmarshal([]byte(body), &serviceError)
		return httpRes, serviceError
	}
	return httpRes, err
}

type AllowedOriginsCompanyLevelApiGetAllowedOriginConfig struct {
	ctx             context.Context
	companyId       string
	apiCredentialId string
	originId        string
}

/*
GetAllowedOrigin Get an allowed origin

Returns the [allowed origin](https://docs.adyen.com/development-resources/client-side-authentication#allowed-origins) identified in the path.

To make this request, your API credential must have the following [roles](https://docs.adyen.com/development-resources/api-credentials#api-permissions):
* Management API—API credentials read and write

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param companyId The unique identifier of the company account.
	@param apiCredentialId Unique identifier of the API credential.
	@param originId Unique identifier of the allowed origin.
	@return AllowedOriginsCompanyLevelApiGetAllowedOriginConfig
*/
func (a *AllowedOriginsCompanyLevelApi) GetAllowedOriginConfig(ctx context.Context, companyId string, apiCredentialId string, originId string) AllowedOriginsCompanyLevelApiGetAllowedOriginConfig {
	return AllowedOriginsCompanyLevelApiGetAllowedOriginConfig{
		ctx:             ctx,
		companyId:       companyId,
		apiCredentialId: apiCredentialId,
		originId:        originId,
	}
}

/*
Get an allowed origin
Returns the [allowed origin](https://docs.adyen.com/development-resources/client-side-authentication#allowed-origins) identified in the path.  To make this request, your API credential must have the following [roles](https://docs.adyen.com/development-resources/api-credentials#api-permissions): * Management API—API credentials read and write
 * @param companyId The unique identifier of the company account.
 * @param apiCredentialId Unique identifier of the API credential.
 * @param originId Unique identifier of the allowed origin.
 * @param ctxs ...context.Context - optional, for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@return AllowedOrigin
*/

func (a *AllowedOriginsCompanyLevelApi) GetAllowedOrigin(r AllowedOriginsCompanyLevelApiGetAllowedOriginConfig) (AllowedOrigin, *_nethttp.Response, error) {
	var serviceError common.RestServiceError
	res := &AllowedOrigin{}
	path := "/companies/{companyId}/apiCredentials/{apiCredentialId}/allowedOrigins/{originId}"
	path = strings.Replace(path, "{"+"companyId"+"}", url.PathEscape(common.ParameterValueToString(r.companyId, "companyId")), -1)
	path = strings.Replace(path, "{"+"apiCredentialId"+"}", url.PathEscape(common.ParameterValueToString(r.apiCredentialId, "apiCredentialId")), -1)
	path = strings.Replace(path, "{"+"originId"+"}", url.PathEscape(common.ParameterValueToString(r.originId, "originId")), -1)
	queryParams := url.Values{}
	headerParams := make(map[string]string)
	httpRes, err := common.SendAPIRequest(
		r.ctx,
		a.Client,
		nil,
		res,
		_nethttp.MethodGet,
		a.BasePath()+path,
		queryParams,
		headerParams,
	)
	defer httpRes.Body.Close()

	if httpRes.StatusCode == 400 {
		// Read the response body
		body, _ := ioutil.ReadAll(httpRes.Body)
		_ = json.Unmarshal([]byte(body), &serviceError)
		return *res, httpRes, serviceError
	}

	if httpRes.StatusCode == 401 {
		// Read the response body
		body, _ := ioutil.ReadAll(httpRes.Body)
		_ = json.Unmarshal([]byte(body), &serviceError)
		return *res, httpRes, serviceError
	}

	if httpRes.StatusCode == 403 {
		// Read the response body
		body, _ := ioutil.ReadAll(httpRes.Body)
		_ = json.Unmarshal([]byte(body), &serviceError)
		return *res, httpRes, serviceError
	}

	if httpRes.StatusCode == 422 {
		// Read the response body
		body, _ := ioutil.ReadAll(httpRes.Body)
		_ = json.Unmarshal([]byte(body), &serviceError)
		return *res, httpRes, serviceError
	}

	if httpRes.StatusCode == 500 {
		// Read the response body
		body, _ := ioutil.ReadAll(httpRes.Body)
		_ = json.Unmarshal([]byte(body), &serviceError)
		return *res, httpRes, serviceError
	}
	return *res, httpRes, err
}

type AllowedOriginsCompanyLevelApiListAllowedOriginsConfig struct {
	ctx             context.Context
	companyId       string
	apiCredentialId string
}

/*
ListAllowedOrigins Get a list of allowed origins

Returns the list of [allowed origins](https://docs.adyen.com/development-resources/client-side-authentication#allowed-origins) for the API credential identified in the path.

To make this request, your API credential must have the following [roles](https://docs.adyen.com/development-resources/api-credentials#api-permissions):
* Management API—API credentials read and write

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param companyId The unique identifier of the company account.
	@param apiCredentialId Unique identifier of the API credential.
	@return AllowedOriginsCompanyLevelApiListAllowedOriginsConfig
*/
func (a *AllowedOriginsCompanyLevelApi) ListAllowedOriginsConfig(ctx context.Context, companyId string, apiCredentialId string) AllowedOriginsCompanyLevelApiListAllowedOriginsConfig {
	return AllowedOriginsCompanyLevelApiListAllowedOriginsConfig{
		ctx:             ctx,
		companyId:       companyId,
		apiCredentialId: apiCredentialId,
	}
}

/*
Get a list of allowed origins
Returns the list of [allowed origins](https://docs.adyen.com/development-resources/client-side-authentication#allowed-origins) for the API credential identified in the path.  To make this request, your API credential must have the following [roles](https://docs.adyen.com/development-resources/api-credentials#api-permissions): * Management API—API credentials read and write
 * @param companyId The unique identifier of the company account.
 * @param apiCredentialId Unique identifier of the API credential.
 * @param ctxs ...context.Context - optional, for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@return AllowedOriginsResponse
*/

func (a *AllowedOriginsCompanyLevelApi) ListAllowedOrigins(r AllowedOriginsCompanyLevelApiListAllowedOriginsConfig) (AllowedOriginsResponse, *_nethttp.Response, error) {
	var serviceError common.RestServiceError
	res := &AllowedOriginsResponse{}
	path := "/companies/{companyId}/apiCredentials/{apiCredentialId}/allowedOrigins"
	path = strings.Replace(path, "{"+"companyId"+"}", url.PathEscape(common.ParameterValueToString(r.companyId, "companyId")), -1)
	path = strings.Replace(path, "{"+"apiCredentialId"+"}", url.PathEscape(common.ParameterValueToString(r.apiCredentialId, "apiCredentialId")), -1)
	queryParams := url.Values{}
	headerParams := make(map[string]string)
	httpRes, err := common.SendAPIRequest(
		r.ctx,
		a.Client,
		nil,
		res,
		_nethttp.MethodGet,
		a.BasePath()+path,
		queryParams,
		headerParams,
	)
	defer httpRes.Body.Close()

	if httpRes.StatusCode == 400 {
		// Read the response body
		body, _ := ioutil.ReadAll(httpRes.Body)
		_ = json.Unmarshal([]byte(body), &serviceError)
		return *res, httpRes, serviceError
	}

	if httpRes.StatusCode == 401 {
		// Read the response body
		body, _ := ioutil.ReadAll(httpRes.Body)
		_ = json.Unmarshal([]byte(body), &serviceError)
		return *res, httpRes, serviceError
	}

	if httpRes.StatusCode == 403 {
		// Read the response body
		body, _ := ioutil.ReadAll(httpRes.Body)
		_ = json.Unmarshal([]byte(body), &serviceError)
		return *res, httpRes, serviceError
	}

	if httpRes.StatusCode == 422 {
		// Read the response body
		body, _ := ioutil.ReadAll(httpRes.Body)
		_ = json.Unmarshal([]byte(body), &serviceError)
		return *res, httpRes, serviceError
	}

	if httpRes.StatusCode == 500 {
		// Read the response body
		body, _ := ioutil.ReadAll(httpRes.Body)
		_ = json.Unmarshal([]byte(body), &serviceError)
		return *res, httpRes, serviceError
	}
	return *res, httpRes, err
}
