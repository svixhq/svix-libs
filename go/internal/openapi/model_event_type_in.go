/*
 * Svix API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 1.1.1
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// EventTypeIn struct for EventTypeIn
type EventTypeIn struct {
	Archived *bool `json:"archived,omitempty"`
	Deprecated *bool `json:"deprecated,omitempty"`
	Description string `json:"description"`
	FeatureFlag NullableString `json:"featureFlag,omitempty"`
	// The event type group's name
	GroupName NullableString `json:"groupName,omitempty"`
	// The event type's name
	Name string `json:"name"`
	// The schema for the event type for a specific version as a JSON schema.
	Schemas map[string]map[string]interface{} `json:"schemas,omitempty"`
}

// NewEventTypeIn instantiates a new EventTypeIn object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEventTypeIn(description string, name string) *EventTypeIn {
	this := EventTypeIn{}
	var archived bool = false
	this.Archived = &archived
	var deprecated bool = false
	this.Deprecated = &deprecated
	this.Description = description
	this.Name = name
	return &this
}

// NewEventTypeInWithDefaults instantiates a new EventTypeIn object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEventTypeInWithDefaults() *EventTypeIn {
	this := EventTypeIn{}
	var archived bool = false
	this.Archived = &archived
	var deprecated bool = false
	this.Deprecated = &deprecated
	return &this
}

// GetArchived returns the Archived field value if set, zero value otherwise.
func (o *EventTypeIn) GetArchived() bool {
	if o == nil || o.Archived == nil {
		var ret bool
		return ret
	}
	return *o.Archived
}

// GetArchivedOk returns a tuple with the Archived field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventTypeIn) GetArchivedOk() (*bool, bool) {
	if o == nil || o.Archived == nil {
		return nil, false
	}
	return o.Archived, true
}

// HasArchived returns a boolean if a field has been set.
func (o *EventTypeIn) HasArchived() bool {
	if o != nil && o.Archived != nil {
		return true
	}

	return false
}

// SetArchived gets a reference to the given bool and assigns it to the Archived field.
func (o *EventTypeIn) SetArchived(v bool) {
	o.Archived = &v
}

// GetDeprecated returns the Deprecated field value if set, zero value otherwise.
func (o *EventTypeIn) GetDeprecated() bool {
	if o == nil || o.Deprecated == nil {
		var ret bool
		return ret
	}
	return *o.Deprecated
}

// GetDeprecatedOk returns a tuple with the Deprecated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventTypeIn) GetDeprecatedOk() (*bool, bool) {
	if o == nil || o.Deprecated == nil {
		return nil, false
	}
	return o.Deprecated, true
}

// HasDeprecated returns a boolean if a field has been set.
func (o *EventTypeIn) HasDeprecated() bool {
	if o != nil && o.Deprecated != nil {
		return true
	}

	return false
}

// SetDeprecated gets a reference to the given bool and assigns it to the Deprecated field.
func (o *EventTypeIn) SetDeprecated(v bool) {
	o.Deprecated = &v
}

// GetDescription returns the Description field value
func (o *EventTypeIn) GetDescription() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Description
}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
func (o *EventTypeIn) GetDescriptionOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Description, true
}

// SetDescription sets field value
func (o *EventTypeIn) SetDescription(v string) {
	o.Description = v
}

// GetFeatureFlag returns the FeatureFlag field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EventTypeIn) GetFeatureFlag() string {
	if o == nil || o.FeatureFlag.Get() == nil {
		var ret string
		return ret
	}
	return *o.FeatureFlag.Get()
}

// GetFeatureFlagOk returns a tuple with the FeatureFlag field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EventTypeIn) GetFeatureFlagOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.FeatureFlag.Get(), o.FeatureFlag.IsSet()
}

// HasFeatureFlag returns a boolean if a field has been set.
func (o *EventTypeIn) HasFeatureFlag() bool {
	if o != nil && o.FeatureFlag.IsSet() {
		return true
	}

	return false
}

// SetFeatureFlag gets a reference to the given NullableString and assigns it to the FeatureFlag field.
func (o *EventTypeIn) SetFeatureFlag(v string) {
	o.FeatureFlag.Set(&v)
}
// SetFeatureFlagNil sets the value for FeatureFlag to be an explicit nil
func (o *EventTypeIn) SetFeatureFlagNil() {
	o.FeatureFlag.Set(nil)
}

// UnsetFeatureFlag ensures that no value is present for FeatureFlag, not even an explicit nil
func (o *EventTypeIn) UnsetFeatureFlag() {
	o.FeatureFlag.Unset()
}

// GetGroupName returns the GroupName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EventTypeIn) GetGroupName() string {
	if o == nil || o.GroupName.Get() == nil {
		var ret string
		return ret
	}
	return *o.GroupName.Get()
}

// GetGroupNameOk returns a tuple with the GroupName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EventTypeIn) GetGroupNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.GroupName.Get(), o.GroupName.IsSet()
}

// HasGroupName returns a boolean if a field has been set.
func (o *EventTypeIn) HasGroupName() bool {
	if o != nil && o.GroupName.IsSet() {
		return true
	}

	return false
}

// SetGroupName gets a reference to the given NullableString and assigns it to the GroupName field.
func (o *EventTypeIn) SetGroupName(v string) {
	o.GroupName.Set(&v)
}
// SetGroupNameNil sets the value for GroupName to be an explicit nil
func (o *EventTypeIn) SetGroupNameNil() {
	o.GroupName.Set(nil)
}

// UnsetGroupName ensures that no value is present for GroupName, not even an explicit nil
func (o *EventTypeIn) UnsetGroupName() {
	o.GroupName.Unset()
}

// GetName returns the Name field value
func (o *EventTypeIn) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *EventTypeIn) GetNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *EventTypeIn) SetName(v string) {
	o.Name = v
}

// GetSchemas returns the Schemas field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EventTypeIn) GetSchemas() map[string]map[string]interface{} {
	if o == nil  {
		var ret map[string]map[string]interface{}
		return ret
	}
	return o.Schemas
}

// GetSchemasOk returns a tuple with the Schemas field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EventTypeIn) GetSchemasOk() (*map[string]map[string]interface{}, bool) {
	if o == nil || o.Schemas == nil {
		return nil, false
	}
	return &o.Schemas, true
}

// HasSchemas returns a boolean if a field has been set.
func (o *EventTypeIn) HasSchemas() bool {
	if o != nil && o.Schemas != nil {
		return true
	}

	return false
}

// SetSchemas gets a reference to the given map[string]map[string]interface{} and assigns it to the Schemas field.
func (o *EventTypeIn) SetSchemas(v map[string]map[string]interface{}) {
	o.Schemas = v
}

func (o EventTypeIn) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Archived != nil {
		toSerialize["archived"] = o.Archived
	}
	if o.Deprecated != nil {
		toSerialize["deprecated"] = o.Deprecated
	}
	if true {
		toSerialize["description"] = o.Description
	}
	if o.FeatureFlag.IsSet() {
		toSerialize["featureFlag"] = o.FeatureFlag.Get()
	}
	if o.GroupName.IsSet() {
		toSerialize["groupName"] = o.GroupName.Get()
	}
	if true {
		toSerialize["name"] = o.Name
	}
	if o.Schemas != nil {
		toSerialize["schemas"] = o.Schemas
	}
	return json.Marshal(toSerialize)
}

type NullableEventTypeIn struct {
	value *EventTypeIn
	isSet bool
}

func (v NullableEventTypeIn) Get() *EventTypeIn {
	return v.value
}

func (v *NullableEventTypeIn) Set(val *EventTypeIn) {
	v.value = val
	v.isSet = true
}

func (v NullableEventTypeIn) IsSet() bool {
	return v.isSet
}

func (v *NullableEventTypeIn) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEventTypeIn(val *EventTypeIn) *NullableEventTypeIn {
	return &NullableEventTypeIn{value: val, isSet: true}
}

func (v NullableEventTypeIn) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEventTypeIn) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


