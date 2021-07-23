/*
 * Svix API
 *
 * Welcome to the Svix API documentation!  Useful links: [Homepage](https://www.svix.com) | [Support email](mailto:support+docs@svix.com) | [Blog](https://www.svix.com/blog/) | [Slack Community](https://www.svix.com/slack/)  # Introduction  This is the reference documentation and schemas for the Svix API. For tutorials and other documentation please refer to [the documentation](https://docs.svix.com).  ## Main concepts  In Svix you have four important entities you will be interacting with:  - `messages`: these are the webhooks being sent. They can have contents and a few other properties. - `application`: this is where `messages` are sent to. Usually you want to create one application for each of your users. - `endpoint`: endpoints are the URLs messages will be sent to. Each application can have multiple `endpoints` and each message sent to that application will be sent to all of them (unless they are not subscribed to the sent event type). - `event-type`: event types are identifiers denoting the type of the message being sent. Event types are primarily used to decide which events are sent to which endpoint.   ## Authentication  Get your authentication token (`AUTH_TOKEN`) from the [Svix dashboard](https://dashboard.svix.com) and use it as part of the `Authorization` header as such: `Authorization: Bearer ${AUTH_TOKEN}`.  <SecurityDefinitions />   ## Code samples  The code samples assume you already have the respective libraries installed and you know how to use them. For the latest information on how to do that, please refer to [the documentation](https://docs.svix.com/).   ## Cross-Origin Resource Sharing  This API features Cross-Origin Resource Sharing (CORS) implemented in compliance with [W3C spec](https://www.w3.org/TR/cors/). And that allows cross-domain communication from the browser. All responses have a wildcard same-origin which makes them completely public and accessible to everyone, including any code on any site. 
 *
 * API version: 1.4
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// EndpointStats struct for EndpointStats
type EndpointStats struct {
	Fail int32 `json:"fail"`
	Pending int32 `json:"pending"`
	Success int32 `json:"success"`
}

// NewEndpointStats instantiates a new EndpointStats object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEndpointStats(fail int32, pending int32, success int32) *EndpointStats {
	this := EndpointStats{}
	this.Fail = fail
	this.Pending = pending
	this.Success = success
	return &this
}

// NewEndpointStatsWithDefaults instantiates a new EndpointStats object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEndpointStatsWithDefaults() *EndpointStats {
	this := EndpointStats{}
	return &this
}

// GetFail returns the Fail field value
func (o *EndpointStats) GetFail() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Fail
}

// GetFailOk returns a tuple with the Fail field value
// and a boolean to check if the value has been set.
func (o *EndpointStats) GetFailOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Fail, true
}

// SetFail sets field value
func (o *EndpointStats) SetFail(v int32) {
	o.Fail = v
}

// GetPending returns the Pending field value
func (o *EndpointStats) GetPending() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Pending
}

// GetPendingOk returns a tuple with the Pending field value
// and a boolean to check if the value has been set.
func (o *EndpointStats) GetPendingOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Pending, true
}

// SetPending sets field value
func (o *EndpointStats) SetPending(v int32) {
	o.Pending = v
}

// GetSuccess returns the Success field value
func (o *EndpointStats) GetSuccess() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Success
}

// GetSuccessOk returns a tuple with the Success field value
// and a boolean to check if the value has been set.
func (o *EndpointStats) GetSuccessOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Success, true
}

// SetSuccess sets field value
func (o *EndpointStats) SetSuccess(v int32) {
	o.Success = v
}

func (o EndpointStats) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["fail"] = o.Fail
	}
	if true {
		toSerialize["pending"] = o.Pending
	}
	if true {
		toSerialize["success"] = o.Success
	}
	return json.Marshal(toSerialize)
}

type NullableEndpointStats struct {
	value *EndpointStats
	isSet bool
}

func (v NullableEndpointStats) Get() *EndpointStats {
	return v.value
}

func (v *NullableEndpointStats) Set(val *EndpointStats) {
	v.value = val
	v.isSet = true
}

func (v NullableEndpointStats) IsSet() bool {
	return v.isSet
}

func (v *NullableEndpointStats) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEndpointStats(val *EndpointStats) *NullableEndpointStats {
	return &NullableEndpointStats{value: val, isSet: true}
}

func (v NullableEndpointStats) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEndpointStats) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


