/*
Management API

API version: 3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package management

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"

	"github.com/adyen/adyen-go-api-library/v16/src/common"
)

// UsersCompanyLevelApi service
type UsersCompanyLevelApi common.Service

// All parameters accepted by UsersCompanyLevelApi.CreateNewUser
type UsersCompanyLevelApiCreateNewUserInput struct {
	companyId                string
	createCompanyUserRequest *CreateCompanyUserRequest
}

func (r UsersCompanyLevelApiCreateNewUserInput) CreateCompanyUserRequest(createCompanyUserRequest CreateCompanyUserRequest) UsersCompanyLevelApiCreateNewUserInput {
	r.createCompanyUserRequest = &createCompanyUserRequest
	return r
}

/*
Prepare a request for CreateNewUser
@param companyId The unique identifier of the company account.
@return UsersCompanyLevelApiCreateNewUserInput
*/
func (a *UsersCompanyLevelApi) CreateNewUserInput(companyId string) UsersCompanyLevelApiCreateNewUserInput {
	return UsersCompanyLevelApiCreateNewUserInput{
		companyId: companyId,
	}
}

/*
CreateNewUser Create a new user

Creates the user for the `companyId` identified in the path.

To make this request, your API credential must have the following [role](https://docs.adyen.com/development-resources/api-credentials#api-permissions):
* Management API—Users read and write


@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param r UsersCompanyLevelApiCreateNewUserInput - Request parameters, see CreateNewUserInput
@return CreateCompanyUserResponse, *http.Response, error
*/
func (a *UsersCompanyLevelApi) CreateNewUser(ctx context.Context, r UsersCompanyLevelApiCreateNewUserInput) (CreateCompanyUserResponse, *http.Response, error) {
	res := &CreateCompanyUserResponse{}
	path := "/companies/{companyId}/users"
	path = strings.Replace(path, "{"+"companyId"+"}", url.PathEscape(common.ParameterValueToString(r.companyId, "companyId")), -1)
	queryParams := url.Values{}
	headerParams := make(map[string]string)
	httpRes, err := common.SendAPIRequest(
		ctx,
		a.Client,
		r.createCompanyUserRequest,
		res,
		http.MethodPost,
		a.BasePath()+path,
		queryParams,
		headerParams,
	)

	if httpRes == nil {
		return *res, httpRes, err
	}

	var serviceError common.RestServiceError
	if httpRes.StatusCode == 400 {
		body, _ := ioutil.ReadAll(httpRes.Body)
		decodeError := json.Unmarshal([]byte(body), &serviceError)
		if decodeError != nil {
			return *res, httpRes, decodeError
		}
		return *res, httpRes, serviceError
	}
	if httpRes.StatusCode == 401 {
		body, _ := ioutil.ReadAll(httpRes.Body)
		decodeError := json.Unmarshal([]byte(body), &serviceError)
		if decodeError != nil {
			return *res, httpRes, decodeError
		}
		return *res, httpRes, serviceError
	}
	if httpRes.StatusCode == 403 {
		body, _ := ioutil.ReadAll(httpRes.Body)
		decodeError := json.Unmarshal([]byte(body), &serviceError)
		if decodeError != nil {
			return *res, httpRes, decodeError
		}
		return *res, httpRes, serviceError
	}
	if httpRes.StatusCode == 422 {
		body, _ := ioutil.ReadAll(httpRes.Body)
		decodeError := json.Unmarshal([]byte(body), &serviceError)
		if decodeError != nil {
			return *res, httpRes, decodeError
		}
		return *res, httpRes, serviceError
	}
	if httpRes.StatusCode == 500 {
		body, _ := ioutil.ReadAll(httpRes.Body)
		decodeError := json.Unmarshal([]byte(body), &serviceError)
		if decodeError != nil {
			return *res, httpRes, decodeError
		}
		return *res, httpRes, serviceError
	}

	return *res, httpRes, err
}

// All parameters accepted by UsersCompanyLevelApi.GetUserDetails
type UsersCompanyLevelApiGetUserDetailsInput struct {
	companyId string
	userId    string
}

/*
Prepare a request for GetUserDetails
@param companyId The unique identifier of the company account.@param userId The unique identifier of the user.
@return UsersCompanyLevelApiGetUserDetailsInput
*/
func (a *UsersCompanyLevelApi) GetUserDetailsInput(companyId string, userId string) UsersCompanyLevelApiGetUserDetailsInput {
	return UsersCompanyLevelApiGetUserDetailsInput{
		companyId: companyId,
		userId:    userId,
	}
}

/*
GetUserDetails Get user details

Returns user details for the `userId` and the `companyId` identified in the path.

To make this request, your API credential must have the following [role](https://docs.adyen.com/development-resources/api-credentials#api-permissions):
* Management API—Users read and write


@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param r UsersCompanyLevelApiGetUserDetailsInput - Request parameters, see GetUserDetailsInput
@return CompanyUser, *http.Response, error
*/
func (a *UsersCompanyLevelApi) GetUserDetails(ctx context.Context, r UsersCompanyLevelApiGetUserDetailsInput) (CompanyUser, *http.Response, error) {
	res := &CompanyUser{}
	path := "/companies/{companyId}/users/{userId}"
	path = strings.Replace(path, "{"+"companyId"+"}", url.PathEscape(common.ParameterValueToString(r.companyId, "companyId")), -1)
	path = strings.Replace(path, "{"+"userId"+"}", url.PathEscape(common.ParameterValueToString(r.userId, "userId")), -1)
	queryParams := url.Values{}
	headerParams := make(map[string]string)
	httpRes, err := common.SendAPIRequest(
		ctx,
		a.Client,
		nil,
		res,
		http.MethodGet,
		a.BasePath()+path,
		queryParams,
		headerParams,
	)

	if httpRes == nil {
		return *res, httpRes, err
	}

	var serviceError common.RestServiceError
	if httpRes.StatusCode == 400 {
		body, _ := ioutil.ReadAll(httpRes.Body)
		decodeError := json.Unmarshal([]byte(body), &serviceError)
		if decodeError != nil {
			return *res, httpRes, decodeError
		}
		return *res, httpRes, serviceError
	}
	if httpRes.StatusCode == 401 {
		body, _ := ioutil.ReadAll(httpRes.Body)
		decodeError := json.Unmarshal([]byte(body), &serviceError)
		if decodeError != nil {
			return *res, httpRes, decodeError
		}
		return *res, httpRes, serviceError
	}
	if httpRes.StatusCode == 403 {
		body, _ := ioutil.ReadAll(httpRes.Body)
		decodeError := json.Unmarshal([]byte(body), &serviceError)
		if decodeError != nil {
			return *res, httpRes, decodeError
		}
		return *res, httpRes, serviceError
	}
	if httpRes.StatusCode == 422 {
		body, _ := ioutil.ReadAll(httpRes.Body)
		decodeError := json.Unmarshal([]byte(body), &serviceError)
		if decodeError != nil {
			return *res, httpRes, decodeError
		}
		return *res, httpRes, serviceError
	}
	if httpRes.StatusCode == 500 {
		body, _ := ioutil.ReadAll(httpRes.Body)
		decodeError := json.Unmarshal([]byte(body), &serviceError)
		if decodeError != nil {
			return *res, httpRes, decodeError
		}
		return *res, httpRes, serviceError
	}

	return *res, httpRes, err
}

// All parameters accepted by UsersCompanyLevelApi.ListUsers
type UsersCompanyLevelApiListUsersInput struct {
	companyId  string
	pageNumber *int32
	pageSize   *int32
	username   *string
}

// The number of the page to return.
func (r UsersCompanyLevelApiListUsersInput) PageNumber(pageNumber int32) UsersCompanyLevelApiListUsersInput {
	r.pageNumber = &pageNumber
	return r
}

// The number of items to have on a page. Maximum value is **100**. The default is **10** items on a page.
func (r UsersCompanyLevelApiListUsersInput) PageSize(pageSize int32) UsersCompanyLevelApiListUsersInput {
	r.pageSize = &pageSize
	return r
}

// The partial or complete username to select all users that match.
func (r UsersCompanyLevelApiListUsersInput) Username(username string) UsersCompanyLevelApiListUsersInput {
	r.username = &username
	return r
}

/*
Prepare a request for ListUsers
@param companyId The unique identifier of the company account.
@return UsersCompanyLevelApiListUsersInput
*/
func (a *UsersCompanyLevelApi) ListUsersInput(companyId string) UsersCompanyLevelApiListUsersInput {
	return UsersCompanyLevelApiListUsersInput{
		companyId: companyId,
	}
}

/*
ListUsers Get a list of users

Returns the list of users for the `companyId` identified in the path.

To make this request, your API credential must have the following [role](https://docs.adyen.com/development-resources/api-credentials#api-permissions):
* Management API—Users read and write


@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param r UsersCompanyLevelApiListUsersInput - Request parameters, see ListUsersInput
@return ListCompanyUsersResponse, *http.Response, error
*/
func (a *UsersCompanyLevelApi) ListUsers(ctx context.Context, r UsersCompanyLevelApiListUsersInput) (ListCompanyUsersResponse, *http.Response, error) {
	res := &ListCompanyUsersResponse{}
	path := "/companies/{companyId}/users"
	path = strings.Replace(path, "{"+"companyId"+"}", url.PathEscape(common.ParameterValueToString(r.companyId, "companyId")), -1)
	queryParams := url.Values{}
	headerParams := make(map[string]string)
	if r.pageNumber != nil {
		common.ParameterAddToQuery(queryParams, "pageNumber", r.pageNumber, "")
	}
	if r.pageSize != nil {
		common.ParameterAddToQuery(queryParams, "pageSize", r.pageSize, "")
	}
	if r.username != nil {
		common.ParameterAddToQuery(queryParams, "username", r.username, "")
	}
	httpRes, err := common.SendAPIRequest(
		ctx,
		a.Client,
		nil,
		res,
		http.MethodGet,
		a.BasePath()+path,
		queryParams,
		headerParams,
	)

	if httpRes == nil {
		return *res, httpRes, err
	}

	var serviceError common.RestServiceError
	if httpRes.StatusCode == 400 {
		body, _ := ioutil.ReadAll(httpRes.Body)
		decodeError := json.Unmarshal([]byte(body), &serviceError)
		if decodeError != nil {
			return *res, httpRes, decodeError
		}
		return *res, httpRes, serviceError
	}
	if httpRes.StatusCode == 401 {
		body, _ := ioutil.ReadAll(httpRes.Body)
		decodeError := json.Unmarshal([]byte(body), &serviceError)
		if decodeError != nil {
			return *res, httpRes, decodeError
		}
		return *res, httpRes, serviceError
	}
	if httpRes.StatusCode == 403 {
		body, _ := ioutil.ReadAll(httpRes.Body)
		decodeError := json.Unmarshal([]byte(body), &serviceError)
		if decodeError != nil {
			return *res, httpRes, decodeError
		}
		return *res, httpRes, serviceError
	}
	if httpRes.StatusCode == 422 {
		body, _ := ioutil.ReadAll(httpRes.Body)
		decodeError := json.Unmarshal([]byte(body), &serviceError)
		if decodeError != nil {
			return *res, httpRes, decodeError
		}
		return *res, httpRes, serviceError
	}
	if httpRes.StatusCode == 500 {
		body, _ := ioutil.ReadAll(httpRes.Body)
		decodeError := json.Unmarshal([]byte(body), &serviceError)
		if decodeError != nil {
			return *res, httpRes, decodeError
		}
		return *res, httpRes, serviceError
	}

	return *res, httpRes, err
}

// All parameters accepted by UsersCompanyLevelApi.UpdateUserDetails
type UsersCompanyLevelApiUpdateUserDetailsInput struct {
	companyId                string
	userId                   string
	updateCompanyUserRequest *UpdateCompanyUserRequest
}

func (r UsersCompanyLevelApiUpdateUserDetailsInput) UpdateCompanyUserRequest(updateCompanyUserRequest UpdateCompanyUserRequest) UsersCompanyLevelApiUpdateUserDetailsInput {
	r.updateCompanyUserRequest = &updateCompanyUserRequest
	return r
}

/*
Prepare a request for UpdateUserDetails
@param companyId The unique identifier of the company account.@param userId The unique identifier of the user.
@return UsersCompanyLevelApiUpdateUserDetailsInput
*/
func (a *UsersCompanyLevelApi) UpdateUserDetailsInput(companyId string, userId string) UsersCompanyLevelApiUpdateUserDetailsInput {
	return UsersCompanyLevelApiUpdateUserDetailsInput{
		companyId: companyId,
		userId:    userId,
	}
}

/*
UpdateUserDetails Update user details

Updates user details for the `userId` and the `companyId` identified in the path.

To make this request, your API credential must have the following [role](https://docs.adyen.com/development-resources/api-credentials#api-permissions):
* Management API—Users read and write


@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param r UsersCompanyLevelApiUpdateUserDetailsInput - Request parameters, see UpdateUserDetailsInput
@return CompanyUser, *http.Response, error
*/
func (a *UsersCompanyLevelApi) UpdateUserDetails(ctx context.Context, r UsersCompanyLevelApiUpdateUserDetailsInput) (CompanyUser, *http.Response, error) {
	res := &CompanyUser{}
	path := "/companies/{companyId}/users/{userId}"
	path = strings.Replace(path, "{"+"companyId"+"}", url.PathEscape(common.ParameterValueToString(r.companyId, "companyId")), -1)
	path = strings.Replace(path, "{"+"userId"+"}", url.PathEscape(common.ParameterValueToString(r.userId, "userId")), -1)
	queryParams := url.Values{}
	headerParams := make(map[string]string)
	httpRes, err := common.SendAPIRequest(
		ctx,
		a.Client,
		r.updateCompanyUserRequest,
		res,
		http.MethodPatch,
		a.BasePath()+path,
		queryParams,
		headerParams,
	)

	if httpRes == nil {
		return *res, httpRes, err
	}

	var serviceError common.RestServiceError
	if httpRes.StatusCode == 400 {
		body, _ := ioutil.ReadAll(httpRes.Body)
		decodeError := json.Unmarshal([]byte(body), &serviceError)
		if decodeError != nil {
			return *res, httpRes, decodeError
		}
		return *res, httpRes, serviceError
	}
	if httpRes.StatusCode == 401 {
		body, _ := ioutil.ReadAll(httpRes.Body)
		decodeError := json.Unmarshal([]byte(body), &serviceError)
		if decodeError != nil {
			return *res, httpRes, decodeError
		}
		return *res, httpRes, serviceError
	}
	if httpRes.StatusCode == 403 {
		body, _ := ioutil.ReadAll(httpRes.Body)
		decodeError := json.Unmarshal([]byte(body), &serviceError)
		if decodeError != nil {
			return *res, httpRes, decodeError
		}
		return *res, httpRes, serviceError
	}
	if httpRes.StatusCode == 422 {
		body, _ := ioutil.ReadAll(httpRes.Body)
		decodeError := json.Unmarshal([]byte(body), &serviceError)
		if decodeError != nil {
			return *res, httpRes, decodeError
		}
		return *res, httpRes, serviceError
	}
	if httpRes.StatusCode == 500 {
		body, _ := ioutil.ReadAll(httpRes.Body)
		decodeError := json.Unmarshal([]byte(body), &serviceError)
		if decodeError != nil {
			return *res, httpRes, decodeError
		}
		return *res, httpRes, serviceError
	}

	return *res, httpRes, err
}
