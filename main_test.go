/*
 * Adyen API Client
 *
 * Contact: support@adyen.com
 */

package main

import (
	"testing"

	"github.com/adyen/adyen-go-api-library/src/api"
	"github.com/adyen/adyen-go-api-library/src/checkout"
	"github.com/adyen/adyen-go-api-library/src/common"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_main(t *testing.T) {
	t.Run("Create a new APIClient", func(t *testing.T) {

		client := api.NewAPIClient(&common.Config{
			Environment: "TEST",
		})
		require.NotNil(t, client)
		require.NotNil(t, client.Checkout)
		require.NotNil(t, client.Checkout.Client)
		require.NotNil(t, client.Checkout.Client.Cfg)
		require.Equal(t, common.Test, client.Checkout.Client.Cfg.Environment)
		assert.Equal(t, "https://checkout-test.adyen.com/checkout/v52", client.Checkout.BasePath())

		t.Run("Create a API request that should fail", func(t *testing.T) {
			res, httpRes, err := client.Checkout.PaymentMethodsPost(&checkout.PaymentMethodsRequest{})
			require.NotNil(t, err)
			assert.Equal(t, "401 Unauthorized", err.Error())
			require.NotNil(t, res)
			assert.Equal(t, checkout.PaymentMethodsResponse{}, res)
			require.NotNil(t, httpRes)
			assert.Equal(t, 401, httpRes.StatusCode)
			assert.Equal(t, "401 Unauthorized", httpRes.Status)
		})
	})
}