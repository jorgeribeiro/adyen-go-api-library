/*
Adyen Checkout API

API version: 71
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package checkout

import (
	"encoding/json"

	"github.com/adyen/adyen-go-api-library/v16/src/common"
)

// checks if the ThreeDSRequestData type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &ThreeDSRequestData{}

// ThreeDSRequestData struct for ThreeDSRequestData
type ThreeDSRequestData struct {
	// Dimensions of the 3DS2 challenge window to be displayed to the cardholder.  Possible values:  * **01** - size of 250x400  * **02** - size of 390x400 * **03** - size of 500x600 * **04** - size of 600x400 * **05** - Fullscreen
	ChallengeWindowSize *string `json:"challengeWindowSize,omitempty"`
	// Flag for data only flow.
	DataOnly *string `json:"dataOnly,omitempty"`
	// Indicates if [native 3D Secure authentication](https://docs.adyen.com/online-payments/3d-secure/native-3ds2) should be used when available.  Possible values: * **preferred**: Use native 3D Secure authentication when available. * **disabled**: Only use the redirect 3D Secure authentication flow.
	NativeThreeDS *string `json:"nativeThreeDS,omitempty"`
	// The version of 3D Secure to use.  Possible values:  * **2.1.0** * **2.2.0**
	ThreeDSVersion *string `json:"threeDSVersion,omitempty"`
}

// NewThreeDSRequestData instantiates a new ThreeDSRequestData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewThreeDSRequestData() *ThreeDSRequestData {
	this := ThreeDSRequestData{}
	return &this
}

// NewThreeDSRequestDataWithDefaults instantiates a new ThreeDSRequestData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewThreeDSRequestDataWithDefaults() *ThreeDSRequestData {
	this := ThreeDSRequestData{}
	return &this
}

// GetChallengeWindowSize returns the ChallengeWindowSize field value if set, zero value otherwise.
func (o *ThreeDSRequestData) GetChallengeWindowSize() string {
	if o == nil || common.IsNil(o.ChallengeWindowSize) {
		var ret string
		return ret
	}
	return *o.ChallengeWindowSize
}

// GetChallengeWindowSizeOk returns a tuple with the ChallengeWindowSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ThreeDSRequestData) GetChallengeWindowSizeOk() (*string, bool) {
	if o == nil || common.IsNil(o.ChallengeWindowSize) {
		return nil, false
	}
	return o.ChallengeWindowSize, true
}

// HasChallengeWindowSize returns a boolean if a field has been set.
func (o *ThreeDSRequestData) HasChallengeWindowSize() bool {
	if o != nil && !common.IsNil(o.ChallengeWindowSize) {
		return true
	}

	return false
}

// SetChallengeWindowSize gets a reference to the given string and assigns it to the ChallengeWindowSize field.
func (o *ThreeDSRequestData) SetChallengeWindowSize(v string) {
	o.ChallengeWindowSize = &v
}

// GetDataOnly returns the DataOnly field value if set, zero value otherwise.
func (o *ThreeDSRequestData) GetDataOnly() string {
	if o == nil || common.IsNil(o.DataOnly) {
		var ret string
		return ret
	}
	return *o.DataOnly
}

// GetDataOnlyOk returns a tuple with the DataOnly field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ThreeDSRequestData) GetDataOnlyOk() (*string, bool) {
	if o == nil || common.IsNil(o.DataOnly) {
		return nil, false
	}
	return o.DataOnly, true
}

// HasDataOnly returns a boolean if a field has been set.
func (o *ThreeDSRequestData) HasDataOnly() bool {
	if o != nil && !common.IsNil(o.DataOnly) {
		return true
	}

	return false
}

// SetDataOnly gets a reference to the given string and assigns it to the DataOnly field.
func (o *ThreeDSRequestData) SetDataOnly(v string) {
	o.DataOnly = &v
}

// GetNativeThreeDS returns the NativeThreeDS field value if set, zero value otherwise.
func (o *ThreeDSRequestData) GetNativeThreeDS() string {
	if o == nil || common.IsNil(o.NativeThreeDS) {
		var ret string
		return ret
	}
	return *o.NativeThreeDS
}

// GetNativeThreeDSOk returns a tuple with the NativeThreeDS field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ThreeDSRequestData) GetNativeThreeDSOk() (*string, bool) {
	if o == nil || common.IsNil(o.NativeThreeDS) {
		return nil, false
	}
	return o.NativeThreeDS, true
}

// HasNativeThreeDS returns a boolean if a field has been set.
func (o *ThreeDSRequestData) HasNativeThreeDS() bool {
	if o != nil && !common.IsNil(o.NativeThreeDS) {
		return true
	}

	return false
}

// SetNativeThreeDS gets a reference to the given string and assigns it to the NativeThreeDS field.
func (o *ThreeDSRequestData) SetNativeThreeDS(v string) {
	o.NativeThreeDS = &v
}

// GetThreeDSVersion returns the ThreeDSVersion field value if set, zero value otherwise.
func (o *ThreeDSRequestData) GetThreeDSVersion() string {
	if o == nil || common.IsNil(o.ThreeDSVersion) {
		var ret string
		return ret
	}
	return *o.ThreeDSVersion
}

// GetThreeDSVersionOk returns a tuple with the ThreeDSVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ThreeDSRequestData) GetThreeDSVersionOk() (*string, bool) {
	if o == nil || common.IsNil(o.ThreeDSVersion) {
		return nil, false
	}
	return o.ThreeDSVersion, true
}

// HasThreeDSVersion returns a boolean if a field has been set.
func (o *ThreeDSRequestData) HasThreeDSVersion() bool {
	if o != nil && !common.IsNil(o.ThreeDSVersion) {
		return true
	}

	return false
}

// SetThreeDSVersion gets a reference to the given string and assigns it to the ThreeDSVersion field.
func (o *ThreeDSRequestData) SetThreeDSVersion(v string) {
	o.ThreeDSVersion = &v
}

func (o ThreeDSRequestData) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ThreeDSRequestData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.ChallengeWindowSize) {
		toSerialize["challengeWindowSize"] = o.ChallengeWindowSize
	}
	if !common.IsNil(o.DataOnly) {
		toSerialize["dataOnly"] = o.DataOnly
	}
	if !common.IsNil(o.NativeThreeDS) {
		toSerialize["nativeThreeDS"] = o.NativeThreeDS
	}
	if !common.IsNil(o.ThreeDSVersion) {
		toSerialize["threeDSVersion"] = o.ThreeDSVersion
	}
	return toSerialize, nil
}

type NullableThreeDSRequestData struct {
	value *ThreeDSRequestData
	isSet bool
}

func (v NullableThreeDSRequestData) Get() *ThreeDSRequestData {
	return v.value
}

func (v *NullableThreeDSRequestData) Set(val *ThreeDSRequestData) {
	v.value = val
	v.isSet = true
}

func (v NullableThreeDSRequestData) IsSet() bool {
	return v.isSet
}

func (v *NullableThreeDSRequestData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableThreeDSRequestData(val *ThreeDSRequestData) *NullableThreeDSRequestData {
	return &NullableThreeDSRequestData{value: val, isSet: true}
}

func (v NullableThreeDSRequestData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableThreeDSRequestData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

func (o *ThreeDSRequestData) isValidChallengeWindowSize() bool {
	var allowedEnumValues = []string{"01", "02", "03", "04", "05"}
	for _, allowed := range allowedEnumValues {
		if o.GetChallengeWindowSize() == allowed {
			return true
		}
	}
	return false
}
func (o *ThreeDSRequestData) isValidDataOnly() bool {
	var allowedEnumValues = []string{"false", "true"}
	for _, allowed := range allowedEnumValues {
		if o.GetDataOnly() == allowed {
			return true
		}
	}
	return false
}
func (o *ThreeDSRequestData) isValidNativeThreeDS() bool {
	var allowedEnumValues = []string{"preferred", "disabled"}
	for _, allowed := range allowedEnumValues {
		if o.GetNativeThreeDS() == allowed {
			return true
		}
	}
	return false
}
func (o *ThreeDSRequestData) isValidThreeDSVersion() bool {
	var allowedEnumValues = []string{"2.1.0", "2.2.0"}
	for _, allowed := range allowedEnumValues {
		if o.GetThreeDSVersion() == allowed {
			return true
		}
	}
	return false
}
