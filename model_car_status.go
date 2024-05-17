/*
BringAuto Fleet Management v2 API

Specification for BringAuto fleet backend HTTP API

API version: 2.3.1
Contact: fleet@bringauto.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"fmt"
)

// CarStatus Car Status enum
type CarStatus string

// List of CarStatus
const (
	IDLE CarStatus = "idle"
	CHARGING CarStatus = "charging"
	OUT_OF_ORDER CarStatus = "out_of_order"
	DRIVING CarStatus = "driving"
	IN_STOP CarStatus = "in_stop"
	PAUSED_BY_PHONE CarStatus = "paused_by_phone"
	PAUSED_BY_OBSTACLE CarStatus = "paused_by_obstacle"
	PAUSED_BY_BUTTON CarStatus = "paused_by_button"
)

// All allowed values of CarStatus enum
var AllowedCarStatusEnumValues = []CarStatus{
	"idle",
	"charging",
	"out_of_order",
	"driving",
	"in_stop",
	"paused_by_phone",
	"paused_by_obstacle",
	"paused_by_button",
}

func (v *CarStatus) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := CarStatus(value)
	for _, existing := range AllowedCarStatusEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid CarStatus", value)
}

// NewCarStatusFromValue returns a pointer to a valid CarStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewCarStatusFromValue(v string) (*CarStatus, error) {
	ev := CarStatus(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for CarStatus: valid values are %v", v, AllowedCarStatusEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v CarStatus) IsValid() bool {
	for _, existing := range AllowedCarStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to CarStatus value
func (v CarStatus) Ptr() *CarStatus {
	return &v
}

type NullableCarStatus struct {
	value *CarStatus
	isSet bool
}

func (v NullableCarStatus) Get() *CarStatus {
	return v.value
}

func (v *NullableCarStatus) Set(val *CarStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableCarStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableCarStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCarStatus(val *CarStatus) *NullableCarStatus {
	return &NullableCarStatus{value: val, isSet: true}
}

func (v NullableCarStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCarStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

