/*
 * Svix API
 *
 * Welcome to the Svix API documentation!  Useful links: [Homepage](https://www.svix.com) | [Support email](mailto:support+docs@svix.com) | [Blog](https://www.svix.com/blog/) | [Slack Community](https://www.svix.com/slack/)  # Introduction  This is the reference documentation and schemas for the [Svix webhook service](https://www.svix.com) API. For tutorials and other documentation please refer to [the documentation](https://docs.svix.com).  ## Main concepts  In Svix you have four important entities you will be interacting with:  - `messages`: these are the webhooks being sent. They can have contents and a few other properties. - `application`: this is where `messages` are sent to. Usually you want to create one application for each user on your platform. - `endpoint`: endpoints are the URLs messages will be sent to. Each application can have multiple `endpoints` and each message sent to that application will be sent to all of them (unless they are not subscribed to the sent event type). - `event-type`: event types are identifiers denoting the type of the message being sent. Event types are primarily used to decide which events are sent to which endpoint.   ## Authentication  Get your authentication token (`AUTH_TOKEN`) from the [Svix dashboard](https://dashboard.svix.com) and use it as part of the `Authorization` header as such: `Authorization: Bearer ${AUTH_TOKEN}`.  <SecurityDefinitions />   ## Code samples  The code samples assume you already have the respective libraries installed and you know how to use them. For the latest information on how to do that, please refer to [the documentation](https://docs.svix.com/).   ## Idempotency  Svix supports [idempotency](https://en.wikipedia.org/wiki/Idempotence) for safely retrying requests without accidentally performing the same operation twice. This is useful when an API call is disrupted in transit and you do not receive a response.  To perform an idempotent request, pass the idempotency key in the `Idempotency-Key` header to the request. The idempotency key should be a unique value generated by the client. You can create the key in however way you like, though we suggest using UUID v4, or any other string with enough entropy to avoid collisions.  Svix's idempotency works by saving the resulting status code and body of the first request made for any given idempotency key for any successful request. Subsequent requests with the same key return the same result.  Please note that idempotency is only supported for `POST` requests.   ## Cross-Origin Resource Sharing  This API features Cross-Origin Resource Sharing (CORS) implemented in compliance with [W3C spec](https://www.w3.org/TR/cors/). And that allows cross-domain communication from the browser. All responses have a wildcard same-origin which makes them completely public and accessible to everyone, including any code on any site. 
 *
 * API version: 1.4
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// AppUsageStatsOut struct for AppUsageStatsOut
type AppUsageStatsOut struct {
	Id string `json:"id"`
	Status BackgroundTaskStatus `json:"status"`
	Task *string `json:"task,omitempty"`
}

// NewAppUsageStatsOut instantiates a new AppUsageStatsOut object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAppUsageStatsOut(id string, status BackgroundTaskStatus) *AppUsageStatsOut {
	this := AppUsageStatsOut{}
	this.Id = id
	this.Status = status
	var task string = "application.stats"
	this.Task = &task
	return &this
}

// NewAppUsageStatsOutWithDefaults instantiates a new AppUsageStatsOut object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAppUsageStatsOutWithDefaults() *AppUsageStatsOut {
	this := AppUsageStatsOut{}
	var task string = "application.stats"
	this.Task = &task
	return &this
}

// GetId returns the Id field value
func (o *AppUsageStatsOut) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *AppUsageStatsOut) GetIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *AppUsageStatsOut) SetId(v string) {
	o.Id = v
}

// GetStatus returns the Status field value
func (o *AppUsageStatsOut) GetStatus() BackgroundTaskStatus {
	if o == nil {
		var ret BackgroundTaskStatus
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *AppUsageStatsOut) GetStatusOk() (*BackgroundTaskStatus, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *AppUsageStatsOut) SetStatus(v BackgroundTaskStatus) {
	o.Status = v
}

// GetTask returns the Task field value if set, zero value otherwise.
func (o *AppUsageStatsOut) GetTask() string {
	if o == nil || o.Task == nil {
		var ret string
		return ret
	}
	return *o.Task
}

// GetTaskOk returns a tuple with the Task field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppUsageStatsOut) GetTaskOk() (*string, bool) {
	if o == nil || o.Task == nil {
		return nil, false
	}
	return o.Task, true
}

// HasTask returns a boolean if a field has been set.
func (o *AppUsageStatsOut) HasTask() bool {
	if o != nil && o.Task != nil {
		return true
	}

	return false
}

// SetTask gets a reference to the given string and assigns it to the Task field.
func (o *AppUsageStatsOut) SetTask(v string) {
	o.Task = &v
}

func (o AppUsageStatsOut) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["id"] = o.Id
	}
	if true {
		toSerialize["status"] = o.Status
	}
	if o.Task != nil {
		toSerialize["task"] = o.Task
	}
	return json.Marshal(toSerialize)
}

type NullableAppUsageStatsOut struct {
	value *AppUsageStatsOut
	isSet bool
}

func (v NullableAppUsageStatsOut) Get() *AppUsageStatsOut {
	return v.value
}

func (v *NullableAppUsageStatsOut) Set(val *AppUsageStatsOut) {
	v.value = val
	v.isSet = true
}

func (v NullableAppUsageStatsOut) IsSet() bool {
	return v.isSet
}

func (v *NullableAppUsageStatsOut) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAppUsageStatsOut(val *AppUsageStatsOut) *NullableAppUsageStatsOut {
	return &NullableAppUsageStatsOut{value: val, isSet: true}
}

func (v NullableAppUsageStatsOut) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAppUsageStatsOut) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


