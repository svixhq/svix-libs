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

// checks if the ReplayOut type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ReplayOut{}

// ReplayOut struct for ReplayOut
type ReplayOut struct {
	Id string `json:"id"`
	Status BackgroundTaskStatus `json:"status"`
	Task BackgroundTaskType `json:"task"`
}

type _ReplayOut ReplayOut

// NewReplayOut instantiates a new ReplayOut object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewReplayOut(id string, status BackgroundTaskStatus, task BackgroundTaskType) *ReplayOut {
	this := ReplayOut{}
	this.Id = id
	this.Status = status
	this.Task = task
	return &this
}

// NewReplayOutWithDefaults instantiates a new ReplayOut object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewReplayOutWithDefaults() *ReplayOut {
	this := ReplayOut{}
	return &this
}

// GetId returns the Id field value
func (o *ReplayOut) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *ReplayOut) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *ReplayOut) SetId(v string) {
	o.Id = v
}

// GetStatus returns the Status field value
func (o *ReplayOut) GetStatus() BackgroundTaskStatus {
	if o == nil {
		var ret BackgroundTaskStatus
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *ReplayOut) GetStatusOk() (*BackgroundTaskStatus, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *ReplayOut) SetStatus(v BackgroundTaskStatus) {
	o.Status = v
}

// GetTask returns the Task field value
func (o *ReplayOut) GetTask() BackgroundTaskType {
	if o == nil {
		var ret BackgroundTaskType
		return ret
	}

	return o.Task
}

// GetTaskOk returns a tuple with the Task field value
// and a boolean to check if the value has been set.
func (o *ReplayOut) GetTaskOk() (*BackgroundTaskType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Task, true
}

// SetTask sets field value
func (o *ReplayOut) SetTask(v BackgroundTaskType) {
	o.Task = v
}

func (o ReplayOut) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ReplayOut) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["status"] = o.Status
	toSerialize["task"] = o.Task
	return toSerialize, nil
}

func (o *ReplayOut) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"status",
		"task",
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

	varReplayOut := _ReplayOut{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varReplayOut)

	if err != nil {
		return err
	}

	*o = ReplayOut(varReplayOut)

	return err
}

type NullableReplayOut struct {
	value *ReplayOut
	isSet bool
}

func (v NullableReplayOut) Get() *ReplayOut {
	return v.value
}

func (v *NullableReplayOut) Set(val *ReplayOut) {
	v.value = val
	v.isSet = true
}

func (v NullableReplayOut) IsSet() bool {
	return v.isSet
}

func (v *NullableReplayOut) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableReplayOut(val *ReplayOut) *NullableReplayOut {
	return &NullableReplayOut{value: val, isSet: true}
}

func (v NullableReplayOut) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableReplayOut) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


