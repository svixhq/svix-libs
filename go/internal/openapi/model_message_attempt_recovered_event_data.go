/*
 * Svix API
 *
 * Welcome to the Svix API documentation!  Useful links: [Homepage](https://www.svix.com) | [Support email](mailto:support+docs@svix.com) | [Blog](https://www.svix.com/blog/) | [Slack Community](https://www.svix.com/slack/)  # Introduction  This is the reference documentation and schemas for the [Svix webhook service](https://www.svix.com) API. For tutorials and other documentation please refer to [the documentation](https://docs.svix.com).  ## Main concepts  In Svix you have four important entities you will be interacting with:  - `messages`: these are the webhooks being sent. They can have contents and a few other properties. - `application`: this is where `messages` are sent to. Usually you want to create one application for each user on your platform. - `endpoint`: endpoints are the URLs messages will be sent to. Each application can have multiple `endpoints` and each message sent to that application will be sent to all of them (unless they are not subscribed to the sent event type). - `event-type`: event types are identifiers denoting the type of the message being sent. Event types are primarily used to decide which events are sent to which endpoint.   ## Authentication  Get your authentication token (`AUTH_TOKEN`) from the [Svix dashboard](https://dashboard.svix.com) and use it as part of the `Authorization` header as such: `Authorization: Bearer ${AUTH_TOKEN}`. For more information on authentication, please refer to the [authentication token docs](https://docs.svix.com/api-keys).  <SecurityDefinitions />   ## Code samples  The code samples assume you already have the respective libraries installed and you know how to use them. For the latest information on how to do that, please refer to [the documentation](https://docs.svix.com/).   ## Idempotency  Svix supports [idempotency](https://en.wikipedia.org/wiki/Idempotence) for safely retrying requests without accidentally performing the same operation twice. This is useful when an API call is disrupted in transit and you do not receive a response.  To perform an idempotent request, pass the idempotency key in the `Idempotency-Key` header to the request. The idempotency key should be a unique value generated by the client. You can create the key in however way you like, though we suggest using UUID v4, or any other string with enough entropy to avoid collisions.  Svix's idempotency works by saving the resulting status code and body of the first request made for any given idempotency key for any successful request. Subsequent requests with the same key return the same result.  Please note that idempotency is only supported for `POST` requests.   ## Cross-Origin Resource Sharing  This API features Cross-Origin Resource Sharing (CORS) implemented in compliance with [W3C spec](https://www.w3.org/TR/cors/). And that allows cross-domain communication from the browser. All responses have a wildcard same-origin which makes them completely public and accessible to everyone, including any code on any site. 
 *
 * API version: 1.7.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// MessageAttemptRecoveredEventData Sent when a message delivery has failed (all of the retry attempts have been exhausted) as a \"message.attempt.exhausted\" type or after it's failed four times as a \"message.attempt.failing\" event.
type MessageAttemptRecoveredEventData struct {
	// The app's ID
	AppId string `json:"appId"`
	// The app's UID
	AppUid NullableString `json:"appUid,omitempty"`
	// The ep's ID
	EndpointId string `json:"endpointId"`
	LastAttempt MessageAttemptFailedData `json:"lastAttempt"`
	// The msg's UID
	MsgEventId NullableString `json:"msgEventId,omitempty"`
	// The msg's ID
	MsgId string `json:"msgId"`
}

// NewMessageAttemptRecoveredEventData instantiates a new MessageAttemptRecoveredEventData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMessageAttemptRecoveredEventData(appId string, endpointId string, lastAttempt MessageAttemptFailedData, msgId string) *MessageAttemptRecoveredEventData {
	this := MessageAttemptRecoveredEventData{}
	this.AppId = appId
	this.EndpointId = endpointId
	this.LastAttempt = lastAttempt
	this.MsgId = msgId
	return &this
}

// NewMessageAttemptRecoveredEventDataWithDefaults instantiates a new MessageAttemptRecoveredEventData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMessageAttemptRecoveredEventDataWithDefaults() *MessageAttemptRecoveredEventData {
	this := MessageAttemptRecoveredEventData{}
	return &this
}

// GetAppId returns the AppId field value
func (o *MessageAttemptRecoveredEventData) GetAppId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AppId
}

// GetAppIdOk returns a tuple with the AppId field value
// and a boolean to check if the value has been set.
func (o *MessageAttemptRecoveredEventData) GetAppIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.AppId, true
}

// SetAppId sets field value
func (o *MessageAttemptRecoveredEventData) SetAppId(v string) {
	o.AppId = v
}

// GetAppUid returns the AppUid field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MessageAttemptRecoveredEventData) GetAppUid() string {
	if o == nil || o.AppUid.Get() == nil {
		var ret string
		return ret
	}
	return *o.AppUid.Get()
}

// GetAppUidOk returns a tuple with the AppUid field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MessageAttemptRecoveredEventData) GetAppUidOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.AppUid.Get(), o.AppUid.IsSet()
}

// HasAppUid returns a boolean if a field has been set.
func (o *MessageAttemptRecoveredEventData) HasAppUid() bool {
	if o != nil && o.AppUid.IsSet() {
		return true
	}

	return false
}

// SetAppUid gets a reference to the given NullableString and assigns it to the AppUid field.
func (o *MessageAttemptRecoveredEventData) SetAppUid(v string) {
	o.AppUid.Set(&v)
}
// SetAppUidNil sets the value for AppUid to be an explicit nil
func (o *MessageAttemptRecoveredEventData) SetAppUidNil() {
	o.AppUid.Set(nil)
}

// UnsetAppUid ensures that no value is present for AppUid, not even an explicit nil
func (o *MessageAttemptRecoveredEventData) UnsetAppUid() {
	o.AppUid.Unset()
}

// GetEndpointId returns the EndpointId field value
func (o *MessageAttemptRecoveredEventData) GetEndpointId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.EndpointId
}

// GetEndpointIdOk returns a tuple with the EndpointId field value
// and a boolean to check if the value has been set.
func (o *MessageAttemptRecoveredEventData) GetEndpointIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.EndpointId, true
}

// SetEndpointId sets field value
func (o *MessageAttemptRecoveredEventData) SetEndpointId(v string) {
	o.EndpointId = v
}

// GetLastAttempt returns the LastAttempt field value
func (o *MessageAttemptRecoveredEventData) GetLastAttempt() MessageAttemptFailedData {
	if o == nil {
		var ret MessageAttemptFailedData
		return ret
	}

	return o.LastAttempt
}

// GetLastAttemptOk returns a tuple with the LastAttempt field value
// and a boolean to check if the value has been set.
func (o *MessageAttemptRecoveredEventData) GetLastAttemptOk() (*MessageAttemptFailedData, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.LastAttempt, true
}

// SetLastAttempt sets field value
func (o *MessageAttemptRecoveredEventData) SetLastAttempt(v MessageAttemptFailedData) {
	o.LastAttempt = v
}

// GetMsgEventId returns the MsgEventId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MessageAttemptRecoveredEventData) GetMsgEventId() string {
	if o == nil || o.MsgEventId.Get() == nil {
		var ret string
		return ret
	}
	return *o.MsgEventId.Get()
}

// GetMsgEventIdOk returns a tuple with the MsgEventId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MessageAttemptRecoveredEventData) GetMsgEventIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.MsgEventId.Get(), o.MsgEventId.IsSet()
}

// HasMsgEventId returns a boolean if a field has been set.
func (o *MessageAttemptRecoveredEventData) HasMsgEventId() bool {
	if o != nil && o.MsgEventId.IsSet() {
		return true
	}

	return false
}

// SetMsgEventId gets a reference to the given NullableString and assigns it to the MsgEventId field.
func (o *MessageAttemptRecoveredEventData) SetMsgEventId(v string) {
	o.MsgEventId.Set(&v)
}
// SetMsgEventIdNil sets the value for MsgEventId to be an explicit nil
func (o *MessageAttemptRecoveredEventData) SetMsgEventIdNil() {
	o.MsgEventId.Set(nil)
}

// UnsetMsgEventId ensures that no value is present for MsgEventId, not even an explicit nil
func (o *MessageAttemptRecoveredEventData) UnsetMsgEventId() {
	o.MsgEventId.Unset()
}

// GetMsgId returns the MsgId field value
func (o *MessageAttemptRecoveredEventData) GetMsgId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.MsgId
}

// GetMsgIdOk returns a tuple with the MsgId field value
// and a boolean to check if the value has been set.
func (o *MessageAttemptRecoveredEventData) GetMsgIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.MsgId, true
}

// SetMsgId sets field value
func (o *MessageAttemptRecoveredEventData) SetMsgId(v string) {
	o.MsgId = v
}

func (o MessageAttemptRecoveredEventData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["appId"] = o.AppId
	}
	if o.AppUid.IsSet() {
		toSerialize["appUid"] = o.AppUid.Get()
	}
	if true {
		toSerialize["endpointId"] = o.EndpointId
	}
	if true {
		toSerialize["lastAttempt"] = o.LastAttempt
	}
	if o.MsgEventId.IsSet() {
		toSerialize["msgEventId"] = o.MsgEventId.Get()
	}
	if true {
		toSerialize["msgId"] = o.MsgId
	}
	return json.Marshal(toSerialize)
}

type NullableMessageAttemptRecoveredEventData struct {
	value *MessageAttemptRecoveredEventData
	isSet bool
}

func (v NullableMessageAttemptRecoveredEventData) Get() *MessageAttemptRecoveredEventData {
	return v.value
}

func (v *NullableMessageAttemptRecoveredEventData) Set(val *MessageAttemptRecoveredEventData) {
	v.value = val
	v.isSet = true
}

func (v NullableMessageAttemptRecoveredEventData) IsSet() bool {
	return v.isSet
}

func (v *NullableMessageAttemptRecoveredEventData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMessageAttemptRecoveredEventData(val *MessageAttemptRecoveredEventData) *NullableMessageAttemptRecoveredEventData {
	return &NullableMessageAttemptRecoveredEventData{value: val, isSet: true}
}

func (v NullableMessageAttemptRecoveredEventData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMessageAttemptRecoveredEventData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


