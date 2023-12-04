/*
Configuration API

API version: 2
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package balanceplatform

import (
	"encoding/json"

	"github.com/adyen/adyen-go-api-library/v8/src/common"
)

// checks if the SweepSchedule type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &SweepSchedule{}

// SweepSchedule struct for SweepSchedule
type SweepSchedule struct {
	// A [cron expression](https://en.wikipedia.org/wiki/Cron#CRON_expression) that is used to set the sweep schedule. The schedule uses the time zone of the balance account.  For example, **30 17 * * MON** schedules a sweep every Monday at 17:30.  The expression must have five values separated by a single space in the following order:  * Minute: **0-59**  * Hour: **0-23**  * Day of the month: **1-31**  * Month: **1-12** or **JAN-DEC**  * Day of the week: **0-7** (0 and 7 are Sunday) or **MON-SUN**.  The following non-standard characters are supported: **&ast;**, **L**, **#**, **W** and **_/_**. See [crontab guru](https://crontab.guru/) for more examples.  Required when `type` is **cron**.
	CronExpression *string `json:"cronExpression,omitempty"`
	// The schedule type.  Possible values:  * **cron**: push out funds based on a `cronExpression`.  * **daily**: push out funds daily at 07:00 AM CET.  * **weekly**: push out funds every Monday at 07:00 AM CET.  * **monthly**: push out funds every first of the month at 07:00 AM CET.  * **balance**: pull in funds instantly if the balance is less than or equal to the `triggerAmount`. You can only use this for sweeps of `type` **pull** and when the source is a `merchantAccount` or `transferInstrument`. If the source is transferInstrument, merchant account identifier is still required, with which you want to process the transaction.
	Type string `json:"type"`
}

// NewSweepSchedule instantiates a new SweepSchedule object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSweepSchedule(type_ string) *SweepSchedule {
	this := SweepSchedule{}
	this.Type = type_
	return &this
}

// NewSweepScheduleWithDefaults instantiates a new SweepSchedule object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSweepScheduleWithDefaults() *SweepSchedule {
	this := SweepSchedule{}
	return &this
}

// GetCronExpression returns the CronExpression field value if set, zero value otherwise.
func (o *SweepSchedule) GetCronExpression() string {
	if o == nil || common.IsNil(o.CronExpression) {
		var ret string
		return ret
	}
	return *o.CronExpression
}

// GetCronExpressionOk returns a tuple with the CronExpression field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SweepSchedule) GetCronExpressionOk() (*string, bool) {
	if o == nil || common.IsNil(o.CronExpression) {
		return nil, false
	}
	return o.CronExpression, true
}

// HasCronExpression returns a boolean if a field has been set.
func (o *SweepSchedule) HasCronExpression() bool {
	if o != nil && !common.IsNil(o.CronExpression) {
		return true
	}

	return false
}

// SetCronExpression gets a reference to the given string and assigns it to the CronExpression field.
func (o *SweepSchedule) SetCronExpression(v string) {
	o.CronExpression = &v
}

// GetType returns the Type field value
func (o *SweepSchedule) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *SweepSchedule) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *SweepSchedule) SetType(v string) {
	o.Type = v
}

func (o SweepSchedule) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SweepSchedule) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.CronExpression) {
		toSerialize["cronExpression"] = o.CronExpression
	}
	toSerialize["type"] = o.Type
	return toSerialize, nil
}

type NullableSweepSchedule struct {
	value *SweepSchedule
	isSet bool
}

func (v NullableSweepSchedule) Get() *SweepSchedule {
	return v.value
}

func (v *NullableSweepSchedule) Set(val *SweepSchedule) {
	v.value = val
	v.isSet = true
}

func (v NullableSweepSchedule) IsSet() bool {
	return v.isSet
}

func (v *NullableSweepSchedule) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSweepSchedule(val *SweepSchedule) *NullableSweepSchedule {
	return &NullableSweepSchedule{value: val, isSet: true}
}

func (v NullableSweepSchedule) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSweepSchedule) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

func (o *SweepSchedule) isValidType() bool {
	var allowedEnumValues = []string{"daily", "weekly", "monthly", "balance", "cron"}
	for _, allowed := range allowedEnumValues {
		if o.GetType() == allowed {
			return true
		}
	}
	return false
}