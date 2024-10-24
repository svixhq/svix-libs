/*
Svix API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 1.1.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the EndpointSecretOut type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EndpointSecretOut{}

// EndpointSecretOut struct for EndpointSecretOut
type EndpointSecretOut struct {
	// The endpoint's verification secret. If `null` is passed, a secret is automatically generated. Format: `base64` encoded random bytes optionally prefixed with `whsec_`. Recommended size: 24.
	Key string `json:"key" validate:"regexp=^(whsec_)?[a-zA-Z0-9+\\/=]{32,100}$"`
}

type _EndpointSecretOut EndpointSecretOut

// NewEndpointSecretOut instantiates a new EndpointSecretOut object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEndpointSecretOut(key string) *EndpointSecretOut {
	this := EndpointSecretOut{}
	this.Key = key
	return &this
}

// NewEndpointSecretOutWithDefaults instantiates a new EndpointSecretOut object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEndpointSecretOutWithDefaults() *EndpointSecretOut {
	this := EndpointSecretOut{}
	return &this
}

// GetKey returns the Key field value
func (o *EndpointSecretOut) GetKey() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Key
}

// GetKeyOk returns a tuple with the Key field value
// and a boolean to check if the value has been set.
func (o *EndpointSecretOut) GetKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Key, true
}

// SetKey sets field value
func (o *EndpointSecretOut) SetKey(v string) {
	o.Key = v
}

func (o EndpointSecretOut) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EndpointSecretOut) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["key"] = o.Key
	return toSerialize, nil
}

func (o *EndpointSecretOut) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"key",
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

	varEndpointSecretOut := _EndpointSecretOut{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varEndpointSecretOut)

	if err != nil {
		return err
	}

	*o = EndpointSecretOut(varEndpointSecretOut)

	return err
}

type NullableEndpointSecretOut struct {
	value *EndpointSecretOut
	isSet bool
}

func (v NullableEndpointSecretOut) Get() *EndpointSecretOut {
	return v.value
}

func (v *NullableEndpointSecretOut) Set(val *EndpointSecretOut) {
	v.value = val
	v.isSet = true
}

func (v NullableEndpointSecretOut) IsSet() bool {
	return v.isSet
}

func (v *NullableEndpointSecretOut) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEndpointSecretOut(val *EndpointSecretOut) *NullableEndpointSecretOut {
	return &NullableEndpointSecretOut{value: val, isSet: true}
}

func (v NullableEndpointSecretOut) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEndpointSecretOut) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


