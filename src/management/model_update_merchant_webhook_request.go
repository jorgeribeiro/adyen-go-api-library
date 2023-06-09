/*
Management API

API version: 1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package management

import (
	"encoding/json"

	"github.com/adyen/adyen-go-api-library/v7/src/common"
)

// checks if the UpdateMerchantWebhookRequest type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &UpdateMerchantWebhookRequest{}

// UpdateMerchantWebhookRequest struct for UpdateMerchantWebhookRequest
type UpdateMerchantWebhookRequest struct {
	// Indicates if expired SSL certificates are accepted. Default value: **false**.
	AcceptsExpiredCertificate *bool `json:"acceptsExpiredCertificate,omitempty"`
	// Indicates if self-signed SSL certificates are accepted. Default value: **false**.
	AcceptsSelfSignedCertificate *bool `json:"acceptsSelfSignedCertificate,omitempty"`
	// Indicates if untrusted SSL certificates are accepted. Default value: **false**.
	AcceptsUntrustedRootCertificate *bool `json:"acceptsUntrustedRootCertificate,omitempty"`
	// Indicates if the webhook configuration is active. The field must be **true** for us to send webhooks about events related an account.
	Active             *bool               `json:"active,omitempty"`
	AdditionalSettings *AdditionalSettings `json:"additionalSettings,omitempty"`
	// Format or protocol for receiving webhooks. Possible values: * **soap** * **http** * **json**
	CommunicationFormat *string `json:"communicationFormat,omitempty"`
	// Your description for this webhook configuration.
	Description *string `json:"description,omitempty"`
	// Network type for Terminal API notification webhooks. Possible values: * **public** * **local**  Default Value: **public**.
	NetworkType *string `json:"networkType,omitempty"`
	// Password to access the webhook URL.
	Password *string `json:"password,omitempty"`
	// Indicates if the SOAP action header needs to be populated. Default value: **false**.  Only applies if `communicationFormat`: **soap**.
	PopulateSoapActionHeader *bool `json:"populateSoapActionHeader,omitempty"`
	// SSL version to access the public webhook URL specified in the `url` field. Possible values: * **TLSv1.3** * **TLSv1.2** * **HTTP** - Only allowed on Test environment.  If not specified, the webhook will use `sslVersion`: **TLSv1.2**.
	SslVersion *string `json:"sslVersion,omitempty"`
	// Public URL where webhooks will be sent, for example **https://www.domain.com/webhook-endpoint**.
	Url *string `json:"url,omitempty"`
	// Username to access the webhook URL.
	Username *string `json:"username,omitempty"`
}

// NewUpdateMerchantWebhookRequest instantiates a new UpdateMerchantWebhookRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateMerchantWebhookRequest() *UpdateMerchantWebhookRequest {
	this := UpdateMerchantWebhookRequest{}
	return &this
}

// NewUpdateMerchantWebhookRequestWithDefaults instantiates a new UpdateMerchantWebhookRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateMerchantWebhookRequestWithDefaults() *UpdateMerchantWebhookRequest {
	this := UpdateMerchantWebhookRequest{}
	return &this
}

// GetAcceptsExpiredCertificate returns the AcceptsExpiredCertificate field value if set, zero value otherwise.
func (o *UpdateMerchantWebhookRequest) GetAcceptsExpiredCertificate() bool {
	if o == nil || common.IsNil(o.AcceptsExpiredCertificate) {
		var ret bool
		return ret
	}
	return *o.AcceptsExpiredCertificate
}

// GetAcceptsExpiredCertificateOk returns a tuple with the AcceptsExpiredCertificate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateMerchantWebhookRequest) GetAcceptsExpiredCertificateOk() (*bool, bool) {
	if o == nil || common.IsNil(o.AcceptsExpiredCertificate) {
		return nil, false
	}
	return o.AcceptsExpiredCertificate, true
}

// HasAcceptsExpiredCertificate returns a boolean if a field has been set.
func (o *UpdateMerchantWebhookRequest) HasAcceptsExpiredCertificate() bool {
	if o != nil && !common.IsNil(o.AcceptsExpiredCertificate) {
		return true
	}

	return false
}

// SetAcceptsExpiredCertificate gets a reference to the given bool and assigns it to the AcceptsExpiredCertificate field.
func (o *UpdateMerchantWebhookRequest) SetAcceptsExpiredCertificate(v bool) {
	o.AcceptsExpiredCertificate = &v
}

// GetAcceptsSelfSignedCertificate returns the AcceptsSelfSignedCertificate field value if set, zero value otherwise.
func (o *UpdateMerchantWebhookRequest) GetAcceptsSelfSignedCertificate() bool {
	if o == nil || common.IsNil(o.AcceptsSelfSignedCertificate) {
		var ret bool
		return ret
	}
	return *o.AcceptsSelfSignedCertificate
}

// GetAcceptsSelfSignedCertificateOk returns a tuple with the AcceptsSelfSignedCertificate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateMerchantWebhookRequest) GetAcceptsSelfSignedCertificateOk() (*bool, bool) {
	if o == nil || common.IsNil(o.AcceptsSelfSignedCertificate) {
		return nil, false
	}
	return o.AcceptsSelfSignedCertificate, true
}

// HasAcceptsSelfSignedCertificate returns a boolean if a field has been set.
func (o *UpdateMerchantWebhookRequest) HasAcceptsSelfSignedCertificate() bool {
	if o != nil && !common.IsNil(o.AcceptsSelfSignedCertificate) {
		return true
	}

	return false
}

// SetAcceptsSelfSignedCertificate gets a reference to the given bool and assigns it to the AcceptsSelfSignedCertificate field.
func (o *UpdateMerchantWebhookRequest) SetAcceptsSelfSignedCertificate(v bool) {
	o.AcceptsSelfSignedCertificate = &v
}

// GetAcceptsUntrustedRootCertificate returns the AcceptsUntrustedRootCertificate field value if set, zero value otherwise.
func (o *UpdateMerchantWebhookRequest) GetAcceptsUntrustedRootCertificate() bool {
	if o == nil || common.IsNil(o.AcceptsUntrustedRootCertificate) {
		var ret bool
		return ret
	}
	return *o.AcceptsUntrustedRootCertificate
}

// GetAcceptsUntrustedRootCertificateOk returns a tuple with the AcceptsUntrustedRootCertificate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateMerchantWebhookRequest) GetAcceptsUntrustedRootCertificateOk() (*bool, bool) {
	if o == nil || common.IsNil(o.AcceptsUntrustedRootCertificate) {
		return nil, false
	}
	return o.AcceptsUntrustedRootCertificate, true
}

// HasAcceptsUntrustedRootCertificate returns a boolean if a field has been set.
func (o *UpdateMerchantWebhookRequest) HasAcceptsUntrustedRootCertificate() bool {
	if o != nil && !common.IsNil(o.AcceptsUntrustedRootCertificate) {
		return true
	}

	return false
}

// SetAcceptsUntrustedRootCertificate gets a reference to the given bool and assigns it to the AcceptsUntrustedRootCertificate field.
func (o *UpdateMerchantWebhookRequest) SetAcceptsUntrustedRootCertificate(v bool) {
	o.AcceptsUntrustedRootCertificate = &v
}

// GetActive returns the Active field value if set, zero value otherwise.
func (o *UpdateMerchantWebhookRequest) GetActive() bool {
	if o == nil || common.IsNil(o.Active) {
		var ret bool
		return ret
	}
	return *o.Active
}

// GetActiveOk returns a tuple with the Active field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateMerchantWebhookRequest) GetActiveOk() (*bool, bool) {
	if o == nil || common.IsNil(o.Active) {
		return nil, false
	}
	return o.Active, true
}

// HasActive returns a boolean if a field has been set.
func (o *UpdateMerchantWebhookRequest) HasActive() bool {
	if o != nil && !common.IsNil(o.Active) {
		return true
	}

	return false
}

// SetActive gets a reference to the given bool and assigns it to the Active field.
func (o *UpdateMerchantWebhookRequest) SetActive(v bool) {
	o.Active = &v
}

// GetAdditionalSettings returns the AdditionalSettings field value if set, zero value otherwise.
func (o *UpdateMerchantWebhookRequest) GetAdditionalSettings() AdditionalSettings {
	if o == nil || common.IsNil(o.AdditionalSettings) {
		var ret AdditionalSettings
		return ret
	}
	return *o.AdditionalSettings
}

// GetAdditionalSettingsOk returns a tuple with the AdditionalSettings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateMerchantWebhookRequest) GetAdditionalSettingsOk() (*AdditionalSettings, bool) {
	if o == nil || common.IsNil(o.AdditionalSettings) {
		return nil, false
	}
	return o.AdditionalSettings, true
}

// HasAdditionalSettings returns a boolean if a field has been set.
func (o *UpdateMerchantWebhookRequest) HasAdditionalSettings() bool {
	if o != nil && !common.IsNil(o.AdditionalSettings) {
		return true
	}

	return false
}

// SetAdditionalSettings gets a reference to the given AdditionalSettings and assigns it to the AdditionalSettings field.
func (o *UpdateMerchantWebhookRequest) SetAdditionalSettings(v AdditionalSettings) {
	o.AdditionalSettings = &v
}

// GetCommunicationFormat returns the CommunicationFormat field value if set, zero value otherwise.
func (o *UpdateMerchantWebhookRequest) GetCommunicationFormat() string {
	if o == nil || common.IsNil(o.CommunicationFormat) {
		var ret string
		return ret
	}
	return *o.CommunicationFormat
}

// GetCommunicationFormatOk returns a tuple with the CommunicationFormat field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateMerchantWebhookRequest) GetCommunicationFormatOk() (*string, bool) {
	if o == nil || common.IsNil(o.CommunicationFormat) {
		return nil, false
	}
	return o.CommunicationFormat, true
}

// HasCommunicationFormat returns a boolean if a field has been set.
func (o *UpdateMerchantWebhookRequest) HasCommunicationFormat() bool {
	if o != nil && !common.IsNil(o.CommunicationFormat) {
		return true
	}

	return false
}

// SetCommunicationFormat gets a reference to the given string and assigns it to the CommunicationFormat field.
func (o *UpdateMerchantWebhookRequest) SetCommunicationFormat(v string) {
	o.CommunicationFormat = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *UpdateMerchantWebhookRequest) GetDescription() string {
	if o == nil || common.IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateMerchantWebhookRequest) GetDescriptionOk() (*string, bool) {
	if o == nil || common.IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *UpdateMerchantWebhookRequest) HasDescription() bool {
	if o != nil && !common.IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *UpdateMerchantWebhookRequest) SetDescription(v string) {
	o.Description = &v
}

// GetNetworkType returns the NetworkType field value if set, zero value otherwise.
func (o *UpdateMerchantWebhookRequest) GetNetworkType() string {
	if o == nil || common.IsNil(o.NetworkType) {
		var ret string
		return ret
	}
	return *o.NetworkType
}

// GetNetworkTypeOk returns a tuple with the NetworkType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateMerchantWebhookRequest) GetNetworkTypeOk() (*string, bool) {
	if o == nil || common.IsNil(o.NetworkType) {
		return nil, false
	}
	return o.NetworkType, true
}

// HasNetworkType returns a boolean if a field has been set.
func (o *UpdateMerchantWebhookRequest) HasNetworkType() bool {
	if o != nil && !common.IsNil(o.NetworkType) {
		return true
	}

	return false
}

// SetNetworkType gets a reference to the given string and assigns it to the NetworkType field.
func (o *UpdateMerchantWebhookRequest) SetNetworkType(v string) {
	o.NetworkType = &v
}

// GetPassword returns the Password field value if set, zero value otherwise.
func (o *UpdateMerchantWebhookRequest) GetPassword() string {
	if o == nil || common.IsNil(o.Password) {
		var ret string
		return ret
	}
	return *o.Password
}

// GetPasswordOk returns a tuple with the Password field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateMerchantWebhookRequest) GetPasswordOk() (*string, bool) {
	if o == nil || common.IsNil(o.Password) {
		return nil, false
	}
	return o.Password, true
}

// HasPassword returns a boolean if a field has been set.
func (o *UpdateMerchantWebhookRequest) HasPassword() bool {
	if o != nil && !common.IsNil(o.Password) {
		return true
	}

	return false
}

// SetPassword gets a reference to the given string and assigns it to the Password field.
func (o *UpdateMerchantWebhookRequest) SetPassword(v string) {
	o.Password = &v
}

// GetPopulateSoapActionHeader returns the PopulateSoapActionHeader field value if set, zero value otherwise.
func (o *UpdateMerchantWebhookRequest) GetPopulateSoapActionHeader() bool {
	if o == nil || common.IsNil(o.PopulateSoapActionHeader) {
		var ret bool
		return ret
	}
	return *o.PopulateSoapActionHeader
}

// GetPopulateSoapActionHeaderOk returns a tuple with the PopulateSoapActionHeader field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateMerchantWebhookRequest) GetPopulateSoapActionHeaderOk() (*bool, bool) {
	if o == nil || common.IsNil(o.PopulateSoapActionHeader) {
		return nil, false
	}
	return o.PopulateSoapActionHeader, true
}

// HasPopulateSoapActionHeader returns a boolean if a field has been set.
func (o *UpdateMerchantWebhookRequest) HasPopulateSoapActionHeader() bool {
	if o != nil && !common.IsNil(o.PopulateSoapActionHeader) {
		return true
	}

	return false
}

// SetPopulateSoapActionHeader gets a reference to the given bool and assigns it to the PopulateSoapActionHeader field.
func (o *UpdateMerchantWebhookRequest) SetPopulateSoapActionHeader(v bool) {
	o.PopulateSoapActionHeader = &v
}

// GetSslVersion returns the SslVersion field value if set, zero value otherwise.
func (o *UpdateMerchantWebhookRequest) GetSslVersion() string {
	if o == nil || common.IsNil(o.SslVersion) {
		var ret string
		return ret
	}
	return *o.SslVersion
}

// GetSslVersionOk returns a tuple with the SslVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateMerchantWebhookRequest) GetSslVersionOk() (*string, bool) {
	if o == nil || common.IsNil(o.SslVersion) {
		return nil, false
	}
	return o.SslVersion, true
}

// HasSslVersion returns a boolean if a field has been set.
func (o *UpdateMerchantWebhookRequest) HasSslVersion() bool {
	if o != nil && !common.IsNil(o.SslVersion) {
		return true
	}

	return false
}

// SetSslVersion gets a reference to the given string and assigns it to the SslVersion field.
func (o *UpdateMerchantWebhookRequest) SetSslVersion(v string) {
	o.SslVersion = &v
}

// GetUrl returns the Url field value if set, zero value otherwise.
func (o *UpdateMerchantWebhookRequest) GetUrl() string {
	if o == nil || common.IsNil(o.Url) {
		var ret string
		return ret
	}
	return *o.Url
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateMerchantWebhookRequest) GetUrlOk() (*string, bool) {
	if o == nil || common.IsNil(o.Url) {
		return nil, false
	}
	return o.Url, true
}

// HasUrl returns a boolean if a field has been set.
func (o *UpdateMerchantWebhookRequest) HasUrl() bool {
	if o != nil && !common.IsNil(o.Url) {
		return true
	}

	return false
}

// SetUrl gets a reference to the given string and assigns it to the Url field.
func (o *UpdateMerchantWebhookRequest) SetUrl(v string) {
	o.Url = &v
}

// GetUsername returns the Username field value if set, zero value otherwise.
func (o *UpdateMerchantWebhookRequest) GetUsername() string {
	if o == nil || common.IsNil(o.Username) {
		var ret string
		return ret
	}
	return *o.Username
}

// GetUsernameOk returns a tuple with the Username field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateMerchantWebhookRequest) GetUsernameOk() (*string, bool) {
	if o == nil || common.IsNil(o.Username) {
		return nil, false
	}
	return o.Username, true
}

// HasUsername returns a boolean if a field has been set.
func (o *UpdateMerchantWebhookRequest) HasUsername() bool {
	if o != nil && !common.IsNil(o.Username) {
		return true
	}

	return false
}

// SetUsername gets a reference to the given string and assigns it to the Username field.
func (o *UpdateMerchantWebhookRequest) SetUsername(v string) {
	o.Username = &v
}

func (o UpdateMerchantWebhookRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateMerchantWebhookRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.AcceptsExpiredCertificate) {
		toSerialize["acceptsExpiredCertificate"] = o.AcceptsExpiredCertificate
	}
	if !common.IsNil(o.AcceptsSelfSignedCertificate) {
		toSerialize["acceptsSelfSignedCertificate"] = o.AcceptsSelfSignedCertificate
	}
	if !common.IsNil(o.AcceptsUntrustedRootCertificate) {
		toSerialize["acceptsUntrustedRootCertificate"] = o.AcceptsUntrustedRootCertificate
	}
	if !common.IsNil(o.Active) {
		toSerialize["active"] = o.Active
	}
	if !common.IsNil(o.AdditionalSettings) {
		toSerialize["additionalSettings"] = o.AdditionalSettings
	}
	if !common.IsNil(o.CommunicationFormat) {
		toSerialize["communicationFormat"] = o.CommunicationFormat
	}
	if !common.IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !common.IsNil(o.NetworkType) {
		toSerialize["networkType"] = o.NetworkType
	}
	if !common.IsNil(o.Password) {
		toSerialize["password"] = o.Password
	}
	if !common.IsNil(o.PopulateSoapActionHeader) {
		toSerialize["populateSoapActionHeader"] = o.PopulateSoapActionHeader
	}
	if !common.IsNil(o.SslVersion) {
		toSerialize["sslVersion"] = o.SslVersion
	}
	if !common.IsNil(o.Url) {
		toSerialize["url"] = o.Url
	}
	if !common.IsNil(o.Username) {
		toSerialize["username"] = o.Username
	}
	return toSerialize, nil
}

type NullableUpdateMerchantWebhookRequest struct {
	value *UpdateMerchantWebhookRequest
	isSet bool
}

func (v NullableUpdateMerchantWebhookRequest) Get() *UpdateMerchantWebhookRequest {
	return v.value
}

func (v *NullableUpdateMerchantWebhookRequest) Set(val *UpdateMerchantWebhookRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateMerchantWebhookRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateMerchantWebhookRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateMerchantWebhookRequest(val *UpdateMerchantWebhookRequest) *NullableUpdateMerchantWebhookRequest {
	return &NullableUpdateMerchantWebhookRequest{value: val, isSet: true}
}

func (v NullableUpdateMerchantWebhookRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateMerchantWebhookRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

func (o *UpdateMerchantWebhookRequest) isValidCommunicationFormat() bool {
	var allowedEnumValues = []string{"http", "json", "soap"}
	for _, allowed := range allowedEnumValues {
		if o.GetCommunicationFormat() == allowed {
			return true
		}
	}
	return false
}
func (o *UpdateMerchantWebhookRequest) isValidNetworkType() bool {
	var allowedEnumValues = []string{"local", "public"}
	for _, allowed := range allowedEnumValues {
		if o.GetNetworkType() == allowed {
			return true
		}
	}
	return false
}
func (o *UpdateMerchantWebhookRequest) isValidSslVersion() bool {
	var allowedEnumValues = []string{"HTTP", "SSL", "SSLv3", "TLS", "TLSv1", "TLSv1.1", "TLSv1.2", "TLSv1.3"}
	for _, allowed := range allowedEnumValues {
		if o.GetSslVersion() == allowed {
			return true
		}
	}
	return false
}
