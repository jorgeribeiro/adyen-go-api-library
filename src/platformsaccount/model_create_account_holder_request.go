/*
 * Adyen for Platforms: Account API
 *
 * The Account API provides endpoints for managing account-related entities on your platform. These related entities include account holders, accounts, bank accounts, shareholders, and KYC-related documents. The management operations include actions such as creation, retrieval, updating, and deletion of them.  For more information, refer to our [documentation](https://docs.adyen.com/platforms). ## Authentication To connect to the Account API, you must use basic authentication credentials of your web service user. If you don't have one, please contact the [Adyen Support Team](https://support.adyen.com/hc/en-us/requests/new). Then use its credentials to authenticate your request, for example:  ``` curl -U \"ws@MarketPlace.YourMarketPlace\":\"YourWsPassword\" \\ -H \"Content-Type: application/json\" \\ ... ``` Note that when going live, you need to generate new web service user credentials to access the [live endpoints](https://docs.adyen.com/development-resources/live-endpoints).  ## Versioning The Account API supports versioning of its endpoints through a version suffix in the endpoint URL. This suffix has the following format: \"vXX\", where XX is the version number.  For example: ``` https://cal-test.adyen.com/cal/services/Account/v6/createAccountHolder ```
 *
 * API version: 6
 * Contact: support@adyen.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package platformsaccount

// CreateAccountHolderRequest struct for CreateAccountHolderRequest
type CreateAccountHolderRequest struct {
	// The desired code of the prospective account holder. > Must be between three (3) and fifty (50) characters long. Only letters, digits, and hyphens (-) are permitted.
	AccountHolderCode    string               `json:"accountHolderCode"`
	AccountHolderDetails AccountHolderDetails `json:"accountHolderDetails"`
	// If set to true, an account with the default options is created for this account holder. **Default Value:** true
	CreateDefaultAccount bool `json:"createDefaultAccount,omitempty"`
	// A description of the prospective account holder, maximum 256 characters. You can use alphanumeric characters (A-Z, a-z, 0-9), white spaces, and underscores `_`.
	Description string `json:"description,omitempty"`
	// The entity type. Permitted values: `Business`, `Individual`  If an account holder is 'Business', then `accountHolderDetails.businessDetails` must be provided, as well as at least one entry in the `accountHolderDetails.businessDetails.shareholders` list.  If an account holder is 'Individual', then `accountHolderDetails.individualDetails` must be provided.
	LegalEntity string `json:"legalEntity"`
	// The three-character [ISO currency code](https://docs.adyen.com/development-resources/currency-codes), with which the prospective account holder primarily deals.
	PrimaryCurrency string `json:"primaryCurrency,omitempty"`
	// The starting [processing tier](https://docs.adyen.com/platforms/onboarding-and-verification/precheck-kyc-information) for the prospective account holder.
	ProcessingTier int32 `json:"processingTier,omitempty"`
	// The identifier of the profile that applies to this entity.
	VerificationProfile string `json:"verificationProfile,omitempty"`
}
