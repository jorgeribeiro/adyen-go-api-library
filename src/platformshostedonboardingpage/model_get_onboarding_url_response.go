/*
 * Adyen for Platforms: Hosted Onboarding
 *
 * Hosted Onboarding provides endpoints for using the Hosted Onboarding Page. The related entities include account holders only.  For more information, refer to our [documentation](https://docs.adyen.com/platforms/onboarding-and-verification/hosted-onboarding-page). ## Authentication To connect to the HOP API, you must use basic authentication credentials of your web service user. If you don't have one, please contact the [Adyen Support Team](https://support.adyen.com/hc/en-us/requests/new). Then use its credentials to authenticate your request, for example:  ``` curl -U \"ws@MarketPlace.YourMarketPlace\":\"YourWsPassword\" \\ -H \"Content-Type: application/json\" \\ ... ``` Note that when going live, you need to generate new web service user credentials to access the [live endpoints](https://docs.adyen.com/development-resources/live-endpoints).  ## Versioning The HOP API supports versioning of its endpoints through a version suffix in the endpoint URL. This suffix has the following format: \"vXX\", where XX is the version number.  For example: ``` https://cal-test.adyen.com/cal/services/Hop/v6/getOnboardingUrl ```
 *
 * API version: 6
 * Contact: support@adyen.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package platformshostedonboardingpage
// GetOnboardingUrlResponse struct for GetOnboardingUrlResponse
type GetOnboardingUrlResponse struct {
	// Contains field validation errors that would prevent requests from being processed.
	InvalidFields *[]ErrorFieldType `json:"invalidFields,omitempty"`
	// The reference of a request. Can be used to uniquely identify the request.
	PspReference string `json:"pspReference,omitempty"`
	// The URL to the Hosted Onboarding Page where you should redirect your sub-merchant. This URL must be used within 30 seconds and can only be used once.
	RedirectUrl string `json:"redirectUrl,omitempty"`
	// The result code.
	ResultCode string `json:"resultCode,omitempty"`
}