/*
 * Adyen API Client
 *
 * Contact: support@adyen.com
 */

package main

import (
	"net/http"
	"net/url"
	"os"
	"testing"

	"github.com/adyen/adyen-go-api-library/src/common"
	"github.com/adyen/adyen-go-api-library/src/checkout"
	"github.com/joho/godotenv"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_main(t *testing.T) {
	t.Run("Create a new APIClient", func(t *testing.T) {

		client := NewAPIClient(&common.Config{
			Environment: "TEST",
		})
		require.NotNil(t, client)
		require.NotNil(t, client.client)
		require.NotNil(t, client.client.Cfg)
		require.NotNil(t, client.client.Cfg.HTTPClient)
		require.NotNil(t, client.Checkout)
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

		godotenv.Load("./../.env")

		var (
			MerchantAccount = os.Getenv("ADYEN_MERCHANT")
			APIKey          = os.Getenv("ADYEN_API_KEY")
		)

		client = NewAPIClientWithAPIKey(APIKey, "TEST")

		t.Run("Create a API request that should pass", func(t *testing.T) {

			res, httpRes, err := client.Checkout.PaymentMethodsPost(&checkout.PaymentMethodsRequest{
				MerchantAccount: MerchantAccount,
			})

			require.Nil(t, err)
			require.NotNil(t, res)
			assert.True(t, len(res.Groups) > 1)
			assert.True(t, len(res.PaymentMethods) > 1)
			require.NotNil(t, httpRes)
			assert.Equal(t, 200, httpRes.StatusCode)
			assert.Equal(t, "200 OK", httpRes.Status)
		})

		//creating the proxyURL
		proxyURL, _ := url.Parse("http://localhost:7000")
		transport := &http.Transport{
			Proxy: http.ProxyURL(proxyURL),
		}
		client = NewAPIClient(&common.Config{
			HTTPClient: &http.Client{
				Transport: transport,
			},
			Environment: "TEST",
			ApiKey:      "your api key",
		})
	})
}