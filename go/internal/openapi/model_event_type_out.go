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
	"time"
)

// EventTypeOut struct for EventTypeOut
type EventTypeOut struct {
	Archived *bool `json:"archived,omitempty"`
	CreatedAt time.Time `json:"createdAt"`
	Deprecated bool `json:"deprecated"`
	Description string `json:"description"`
	FeatureFlag NullableString `json:"featureFlag,omitempty"`
	// The event type group's name
	GroupName NullableString `json:"groupName,omitempty"`
	// The event type's name
	Name string `json:"name"`
	// The schema for the event type for a specific version as a JSON schema.
	Schemas map[string]map[string]interface{} `json:"schemas,omitempty"`
	UpdatedAt time.Time `json:"updatedAt"`
}

// NewEventTypeOut instantiates a new EventTypeOut object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEventTypeOut(createdAt time.Time, deprecated bool, description string, name string, updatedAt time.Time) *EventTypeOut {
	this := EventTypeOut{}
	var archived bool = false
	this.Archived = &archived
	this.CreatedAt = createdAt
	this.Deprecated = deprecated
	this.Description = description
	this.Name = name
	this.UpdatedAt = updatedAt
	return &this
}

// NewEventTypeOutWithDefaults instantiates a new EventTypeOut object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEventTypeOutWithDefaults() *EventTypeOut {
	this := EventTypeOut{}
	var archived bool = false
	this.Archived = &archived
	return &this
}

// GetArchived returns the Archived field value if set, zero value otherwise.
func (o *EventTypeOut) GetArchived() bool {
	if o == nil || o.Archived == nil {
		var ret bool
		return ret
	}
	return *o.Archived
}

// GetArchivedOk returns a tuple with the Archived field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventTypeOut) GetArchivedOk() (*bool, bool) {
	if o == nil || o.Archived == nil {
		return nil, false
	}
	return o.Archived, true
}

// HasArchived returns a boolean if a field has been set.
func (o *EventTypeOut) HasArchived() bool {
	if o != nil && o.Archived != nil {
		return true
	}

	return false
}

// SetArchived gets a reference to the given bool and assigns it to the Archived field.
func (o *EventTypeOut) SetArchived(v bool) {
	o.Archived = &v
}

// GetCreatedAt returns the CreatedAt field value
func (o *EventTypeOut) GetCreatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *EventTypeOut) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *EventTypeOut) SetCreatedAt(v time.Time) {
	o.CreatedAt = v
}

// GetDeprecated returns the Deprecated field value
func (o *EventTypeOut) GetDeprecated() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Deprecated
}

// GetDeprecatedOk returns a tuple with the Deprecated field value
// and a boolean to check if the value has been set.
func (o *EventTypeOut) GetDeprecatedOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Deprecated, true
}

// SetDeprecated sets field value
func (o *EventTypeOut) SetDeprecated(v bool) {
	o.Deprecated = v
}

// GetDescription returns the Description field value
func (o *EventTypeOut) GetDescription() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Description
}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
func (o *EventTypeOut) GetDescriptionOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Description, true
}

// SetDescription sets field value
func (o *EventTypeOut) SetDescription(v string) {
	o.Description = v
}

// GetFeatureFlag returns the FeatureFlag field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EventTypeOut) GetFeatureFlag() string {
	if o == nil || o.FeatureFlag.Get() == nil {
		var ret string
		return ret
	}
	return *o.FeatureFlag.Get()
}

// GetFeatureFlagOk returns a tuple with the FeatureFlag field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EventTypeOut) GetFeatureFlagOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.FeatureFlag.Get(), o.FeatureFlag.IsSet()
}

// HasFeatureFlag returns a boolean if a field has been set.
func (o *EventTypeOut) HasFeatureFlag() bool {
	if o != nil && o.FeatureFlag.IsSet() {
		return true
	}

	return false
}

// SetFeatureFlag gets a reference to the given NullableString and assigns it to the FeatureFlag field.
func (o *EventTypeOut) SetFeatureFlag(v string) {
	o.FeatureFlag.Set(&v)
}
// SetFeatureFlagNil sets the value for FeatureFlag to be an explicit nil
func (o *EventTypeOut) SetFeatureFlagNil() {
	o.FeatureFlag.Set(nil)
}

// UnsetFeatureFlag ensures that no value is present for FeatureFlag, not even an explicit nil
func (o *EventTypeOut) UnsetFeatureFlag() {
	o.FeatureFlag.Unset()
}

// GetGroupName returns the GroupName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EventTypeOut) GetGroupName() string {
	if o == nil || o.GroupName.Get() == nil {
		var ret string
		return ret
	}
	return *o.GroupName.Get()
}

// GetGroupNameOk returns a tuple with the GroupName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EventTypeOut) GetGroupNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.GroupName.Get(), o.GroupName.IsSet()
}

// HasGroupName returns a boolean if a field has been set.
func (o *EventTypeOut) HasGroupName() bool {
	if o != nil && o.GroupName.IsSet() {
		return true
	}

	return false
}

// SetGroupName gets a reference to the given NullableString and assigns it to the GroupName field.
func (o *EventTypeOut) SetGroupName(v string) {
	o.GroupName.Set(&v)
}
// SetGroupNameNil sets the value for GroupName to be an explicit nil
func (o *EventTypeOut) SetGroupNameNil() {
	o.GroupName.Set(nil)
}

// UnsetGroupName ensures that no value is present for GroupName, not even an explicit nil
func (o *EventTypeOut) UnsetGroupName() {
	o.GroupName.Unset()
}

// GetName returns the Name field value
func (o *EventTypeOut) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *EventTypeOut) GetNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *EventTypeOut) SetName(v string) {
	o.Name = v
}

// GetSchemas returns the Schemas field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EventTypeOut) GetSchemas() map[string]map[string]interface{} {
	if o == nil  {
		var ret map[string]map[string]interface{}
		return ret
	}
	return o.Schemas
}

// GetSchemasOk returns a tuple with the Schemas field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EventTypeOut) GetSchemasOk() (*map[string]map[string]interface{}, bool) {
	if o == nil || o.Schemas == nil {
		return nil, false
	}
	return &o.Schemas, true
}

// HasSchemas returns a boolean if a field has been set.
func (o *EventTypeOut) HasSchemas() bool {
	if o != nil && o.Schemas != nil {
		return true
	}

	return false
}

// SetSchemas gets a reference to the given map[string]map[string]interface{} and assigns it to the Schemas field.
func (o *EventTypeOut) SetSchemas(v map[string]map[string]interface{}) {
	o.Schemas = v
}

// GetUpdatedAt returns the UpdatedAt field value
func (o *EventTypeOut) GetUpdatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value
// and a boolean to check if the value has been set.
func (o *EventTypeOut) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.UpdatedAt, true
}

// SetUpdatedAt sets field value
func (o *EventTypeOut) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = v
}

func (o EventTypeOut) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Archived != nil {
		toSerialize["archived"] = o.Archived
	}
	if true {
		toSerialize["createdAt"] = o.CreatedAt
	}
	if true {
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
	if true {
		toSerialize["updatedAt"] = o.UpdatedAt
	}
	return json.Marshal(toSerialize)
}

type NullableEventTypeOut struct {
	value *EventTypeOut
	isSet bool
}

func (v NullableEventTypeOut) Get() *EventTypeOut {
	return v.value
}

func (v *NullableEventTypeOut) Set(val *EventTypeOut) {
	v.value = val
	v.isSet = true
}

func (v NullableEventTypeOut) IsSet() bool {
	return v.isSet
}

func (v *NullableEventTypeOut) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEventTypeOut(val *EventTypeOut) *NullableEventTypeOut {
	return &NullableEventTypeOut{value: val, isSet: true}
}

func (v NullableEventTypeOut) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEventTypeOut) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


