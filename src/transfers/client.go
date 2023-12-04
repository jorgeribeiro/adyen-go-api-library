/*
Transfers API

API version: 4
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package transfers

import (
    "github.com/adyen/adyen-go-api-library/v8/src/common"
)

// APIClient manages communication with the Transfers API API v4
// In most cases there should be only one, shared, APIClient.
type APIClient struct {
	common common.Service // Reuse a single struct instead of allocating one for each service on the heap.

	// API Services

	CapitalApi *CapitalApi

	TransactionsApi *TransactionsApi

	TransfersApi *TransfersApi
}

// NewAPIClient creates a new API client.
func NewAPIClient(client *common.Client) *APIClient {
	c := &APIClient{}
    c.common.Client = client
    c.common.BasePath = func() string {
        return client.Cfg.TransfersEndpoint
    }

	// API Services
	c.CapitalApi = (*CapitalApi)(&c.common)
	c.TransactionsApi = (*TransactionsApi)(&c.common)
	c.TransfersApi = (*TransfersApi)(&c.common)

	return c
}