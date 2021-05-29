/*
 * Svix
 *
 * The Svix server API documentation
 *
 * API version: 1.4
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"time"
)

// EndpointMessageOut struct for EndpointMessageOut
type EndpointMessageOut struct {
	EventType string `json:"eventType"`
	EventId *string `json:"eventId,omitempty"`
	Payload map[string]interface{} `json:"payload"`
	Id string `json:"id"`
	Timestamp time.Time `json:"timestamp"`
	Status MessageStatus `json:"status"`
}

// NewEndpointMessageOut instantiates a new EndpointMessageOut object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEndpointMessageOut(eventType string, payload map[string]interface{}, id string, timestamp time.Time, status MessageStatus, ) *EndpointMessageOut {
	this := EndpointMessageOut{}
	this.EventType = eventType
	this.Payload = payload
	this.Id = id
	this.Timestamp = timestamp
	this.Status = status
	return &this
}

// NewEndpointMessageOutWithDefaults instantiates a new EndpointMessageOut object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEndpointMessageOutWithDefaults() *EndpointMessageOut {
	this := EndpointMessageOut{}
	return &this
}

// GetEventType returns the EventType field value
func (o *EndpointMessageOut) GetEventType() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.EventType
}

// GetEventTypeOk returns a tuple with the EventType field value
// and a boolean to check if the value has been set.
func (o *EndpointMessageOut) GetEventTypeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.EventType, true
}

// SetEventType sets field value
func (o *EndpointMessageOut) SetEventType(v string) {
	o.EventType = v
}

// GetEventId returns the EventId field value if set, zero value otherwise.
func (o *EndpointMessageOut) GetEventId() string {
	if o == nil || o.EventId == nil {
		var ret string
		return ret
	}
	return *o.EventId
}

// GetEventIdOk returns a tuple with the EventId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EndpointMessageOut) GetEventIdOk() (*string, bool) {
	if o == nil || o.EventId == nil {
		return nil, false
	}
	return o.EventId, true
}

// HasEventId returns a boolean if a field has been set.
func (o *EndpointMessageOut) HasEventId() bool {
	if o != nil && o.EventId != nil {
		return true
	}

	return false
}

// SetEventId gets a reference to the given string and assigns it to the EventId field.
func (o *EndpointMessageOut) SetEventId(v string) {
	o.EventId = &v
}

// GetPayload returns the Payload field value
func (o *EndpointMessageOut) GetPayload() map[string]interface{} {
	if o == nil  {
		var ret map[string]interface{}
		return ret
	}

	return o.Payload
}

// GetPayloadOk returns a tuple with the Payload field value
// and a boolean to check if the value has been set.
func (o *EndpointMessageOut) GetPayloadOk() (*map[string]interface{}, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Payload, true
}

// SetPayload sets field value
func (o *EndpointMessageOut) SetPayload(v map[string]interface{}) {
	o.Payload = v
}

// GetId returns the Id field value
func (o *EndpointMessageOut) GetId() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *EndpointMessageOut) GetIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *EndpointMessageOut) SetId(v string) {
	o.Id = v
}

// GetTimestamp returns the Timestamp field value
func (o *EndpointMessageOut) GetTimestamp() time.Time {
	if o == nil  {
		var ret time.Time
		return ret
	}

	return o.Timestamp
}

// GetTimestampOk returns a tuple with the Timestamp field value
// and a boolean to check if the value has been set.
func (o *EndpointMessageOut) GetTimestampOk() (*time.Time, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Timestamp, true
}

// SetTimestamp sets field value
func (o *EndpointMessageOut) SetTimestamp(v time.Time) {
	o.Timestamp = v
}

// GetStatus returns the Status field value
func (o *EndpointMessageOut) GetStatus() MessageStatus {
	if o == nil  {
		var ret MessageStatus
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *EndpointMessageOut) GetStatusOk() (*MessageStatus, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *EndpointMessageOut) SetStatus(v MessageStatus) {
	o.Status = v
}

func (o EndpointMessageOut) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["eventType"] = o.EventType
	}
	if o.EventId != nil {
		toSerialize["eventId"] = o.EventId
	}
	if true {
		toSerialize["payload"] = o.Payload
	}
	if true {
		toSerialize["id"] = o.Id
	}
	if true {
		toSerialize["timestamp"] = o.Timestamp
	}
	if true {
		toSerialize["status"] = o.Status
	}
	return json.Marshal(toSerialize)
}

type NullableEndpointMessageOut struct {
	value *EndpointMessageOut
	isSet bool
}

func (v NullableEndpointMessageOut) Get() *EndpointMessageOut {
	return v.value
}

func (v *NullableEndpointMessageOut) Set(val *EndpointMessageOut) {
	v.value = val
	v.isSet = true
}

func (v NullableEndpointMessageOut) IsSet() bool {
	return v.isSet
}

func (v *NullableEndpointMessageOut) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEndpointMessageOut(val *EndpointMessageOut) *NullableEndpointMessageOut {
	return &NullableEndpointMessageOut{value: val, isSet: true}
}

func (v NullableEndpointMessageOut) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEndpointMessageOut) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


