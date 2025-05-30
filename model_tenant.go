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

// checks if the Tenant type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Tenant{}

// Tenant Tenant owning a subset of the entities on the server.
type Tenant struct {
	Id *int32 `json:"id,omitempty"`
	// Tenant name
	Name string `json:"name"`
}

type _Tenant Tenant

// NewTenant instantiates a new Tenant object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTenant(name string) *Tenant {
	this := Tenant{}
	this.Name = name
	return &this
}

// NewTenantWithDefaults instantiates a new Tenant object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTenantWithDefaults() *Tenant {
	this := Tenant{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *Tenant) GetId() int32 {
	if o == nil || IsNil(o.Id) {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Tenant) GetIdOk() (*int32, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *Tenant) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *Tenant) SetId(v int32) {
	o.Id = &v
}

// GetName returns the Name field value
func (o *Tenant) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *Tenant) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *Tenant) SetName(v string) {
	o.Name = v
}

func (o Tenant) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Tenant) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	toSerialize["name"] = o.Name
	return toSerialize, nil
}

func (o *Tenant) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"name",
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

	varTenant := _Tenant{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varTenant)

	if err != nil {
		return err
	}

	*o = Tenant(varTenant)

	return err
}

type NullableTenant struct {
	value *Tenant
	isSet bool
}

func (v NullableTenant) Get() *Tenant {
	return v.value
}

func (v *NullableTenant) Set(val *Tenant) {
	v.value = val
	v.isSet = true
}

func (v NullableTenant) IsSet() bool {
	return v.isSet
}

func (v *NullableTenant) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTenant(val *Tenant) *NullableTenant {
	return &NullableTenant{value: val, isSet: true}
}

func (v NullableTenant) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTenant) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


