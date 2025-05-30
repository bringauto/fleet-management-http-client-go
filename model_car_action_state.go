/*
BringAuto Fleet Management v2 API

Specification for BringAuto fleet backend HTTP API

API version: 4.1.0
Contact: fleet@bringauto.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the CarActionState type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CarActionState{}

// CarActionState Car Action State object structure
type CarActionState struct {
	Id *int32 `json:"id,omitempty"`
	CarId int32 `json:"carId"`
	// A Unix timestamp in milliseconds. The timestamp is used to determine the time of creation of an object.
	Timestamp *int64 `json:"timestamp,omitempty"`
	ActionStatus CarActionStatus `json:"actionStatus"`
}

type _CarActionState CarActionState

// NewCarActionState instantiates a new CarActionState object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCarActionState(carId int32, actionStatus CarActionStatus) *CarActionState {
	this := CarActionState{}
	this.CarId = carId
	this.ActionStatus = actionStatus
	return &this
}

// NewCarActionStateWithDefaults instantiates a new CarActionState object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCarActionStateWithDefaults() *CarActionState {
	this := CarActionState{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *CarActionState) GetId() int32 {
	if o == nil || IsNil(o.Id) {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CarActionState) GetIdOk() (*int32, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *CarActionState) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *CarActionState) SetId(v int32) {
	o.Id = &v
}

// GetCarId returns the CarId field value
func (o *CarActionState) GetCarId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.CarId
}

// GetCarIdOk returns a tuple with the CarId field value
// and a boolean to check if the value has been set.
func (o *CarActionState) GetCarIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CarId, true
}

// SetCarId sets field value
func (o *CarActionState) SetCarId(v int32) {
	o.CarId = v
}

// GetTimestamp returns the Timestamp field value if set, zero value otherwise.
func (o *CarActionState) GetTimestamp() int64 {
	if o == nil || IsNil(o.Timestamp) {
		var ret int64
		return ret
	}
	return *o.Timestamp
}

// GetTimestampOk returns a tuple with the Timestamp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CarActionState) GetTimestampOk() (*int64, bool) {
	if o == nil || IsNil(o.Timestamp) {
		return nil, false
	}
	return o.Timestamp, true
}

// HasTimestamp returns a boolean if a field has been set.
func (o *CarActionState) HasTimestamp() bool {
	if o != nil && !IsNil(o.Timestamp) {
		return true
	}

	return false
}

// SetTimestamp gets a reference to the given int64 and assigns it to the Timestamp field.
func (o *CarActionState) SetTimestamp(v int64) {
	o.Timestamp = &v
}

// GetActionStatus returns the ActionStatus field value
func (o *CarActionState) GetActionStatus() CarActionStatus {
	if o == nil {
		var ret CarActionStatus
		return ret
	}

	return o.ActionStatus
}

// GetActionStatusOk returns a tuple with the ActionStatus field value
// and a boolean to check if the value has been set.
func (o *CarActionState) GetActionStatusOk() (*CarActionStatus, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ActionStatus, true
}

// SetActionStatus sets field value
func (o *CarActionState) SetActionStatus(v CarActionStatus) {
	o.ActionStatus = v
}

func (o CarActionState) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CarActionState) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	toSerialize["carId"] = o.CarId
	if !IsNil(o.Timestamp) {
		toSerialize["timestamp"] = o.Timestamp
	}
	toSerialize["actionStatus"] = o.ActionStatus
	return toSerialize, nil
}

func (o *CarActionState) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"carId",
		"actionStatus",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varCarActionState := _CarActionState{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCarActionState)

	if err != nil {
		return err
	}

	*o = CarActionState(varCarActionState)

	return err
}

type NullableCarActionState struct {
	value *CarActionState
	isSet bool
}

func (v NullableCarActionState) Get() *CarActionState {
	return v.value
}

func (v *NullableCarActionState) Set(val *CarActionState) {
	v.value = val
	v.isSet = true
}

func (v NullableCarActionState) IsSet() bool {
	return v.isSet
}

func (v *NullableCarActionState) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCarActionState(val *CarActionState) *NullableCarActionState {
	return &NullableCarActionState{value: val, isSet: true}
}

func (v NullableCarActionState) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCarActionState) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


