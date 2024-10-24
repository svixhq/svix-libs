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

// checks if the TemplateUpdate type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TemplateUpdate{}

// TemplateUpdate struct for TemplateUpdate
type TemplateUpdate struct {
	Description *string `json:"description,omitempty"`
	FeatureFlag NullableString `json:"featureFlag,omitempty" validate:"regexp=^[a-zA-Z0-9\\\\-_.]+$"`
	FilterTypes []string `json:"filterTypes,omitempty"`
	Instructions *string `json:"instructions,omitempty"`
	InstructionsLink NullableString `json:"instructionsLink,omitempty"`
	Kind *TransformationTemplateKind `json:"kind,omitempty"`
	Logo string `json:"logo"`
	Name *string `json:"name,omitempty"`
	Transformation string `json:"transformation"`
}

type _TemplateUpdate TemplateUpdate

// NewTemplateUpdate instantiates a new TemplateUpdate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTemplateUpdate(logo string, transformation string) *TemplateUpdate {
	this := TemplateUpdate{}
	var description string = ""
	this.Description = &description
	var instructions string = ""
	this.Instructions = &instructions
	this.Logo = logo
	var name string = ""
	this.Name = &name
	this.Transformation = transformation
	return &this
}

// NewTemplateUpdateWithDefaults instantiates a new TemplateUpdate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTemplateUpdateWithDefaults() *TemplateUpdate {
	this := TemplateUpdate{}
	var description string = ""
	this.Description = &description
	var instructions string = ""
	this.Instructions = &instructions
	var name string = ""
	this.Name = &name
	return &this
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *TemplateUpdate) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TemplateUpdate) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *TemplateUpdate) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *TemplateUpdate) SetDescription(v string) {
	o.Description = &v
}

// GetFeatureFlag returns the FeatureFlag field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TemplateUpdate) GetFeatureFlag() string {
	if o == nil || IsNil(o.FeatureFlag.Get()) {
		var ret string
		return ret
	}
	return *o.FeatureFlag.Get()
}

// GetFeatureFlagOk returns a tuple with the FeatureFlag field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TemplateUpdate) GetFeatureFlagOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.FeatureFlag.Get(), o.FeatureFlag.IsSet()
}

// HasFeatureFlag returns a boolean if a field has been set.
func (o *TemplateUpdate) HasFeatureFlag() bool {
	if o != nil && o.FeatureFlag.IsSet() {
		return true
	}

	return false
}

// SetFeatureFlag gets a reference to the given NullableString and assigns it to the FeatureFlag field.
func (o *TemplateUpdate) SetFeatureFlag(v string) {
	o.FeatureFlag.Set(&v)
}
// SetFeatureFlagNil sets the value for FeatureFlag to be an explicit nil
func (o *TemplateUpdate) SetFeatureFlagNil() {
	o.FeatureFlag.Set(nil)
}

// UnsetFeatureFlag ensures that no value is present for FeatureFlag, not even an explicit nil
func (o *TemplateUpdate) UnsetFeatureFlag() {
	o.FeatureFlag.Unset()
}

// GetFilterTypes returns the FilterTypes field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TemplateUpdate) GetFilterTypes() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.FilterTypes
}

// GetFilterTypesOk returns a tuple with the FilterTypes field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TemplateUpdate) GetFilterTypesOk() ([]string, bool) {
	if o == nil || IsNil(o.FilterTypes) {
		return nil, false
	}
	return o.FilterTypes, true
}

// HasFilterTypes returns a boolean if a field has been set.
func (o *TemplateUpdate) HasFilterTypes() bool {
	if o != nil && !IsNil(o.FilterTypes) {
		return true
	}

	return false
}

// SetFilterTypes gets a reference to the given []string and assigns it to the FilterTypes field.
func (o *TemplateUpdate) SetFilterTypes(v []string) {
	o.FilterTypes = v
}

// GetInstructions returns the Instructions field value if set, zero value otherwise.
func (o *TemplateUpdate) GetInstructions() string {
	if o == nil || IsNil(o.Instructions) {
		var ret string
		return ret
	}
	return *o.Instructions
}

// GetInstructionsOk returns a tuple with the Instructions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TemplateUpdate) GetInstructionsOk() (*string, bool) {
	if o == nil || IsNil(o.Instructions) {
		return nil, false
	}
	return o.Instructions, true
}

// HasInstructions returns a boolean if a field has been set.
func (o *TemplateUpdate) HasInstructions() bool {
	if o != nil && !IsNil(o.Instructions) {
		return true
	}

	return false
}

// SetInstructions gets a reference to the given string and assigns it to the Instructions field.
func (o *TemplateUpdate) SetInstructions(v string) {
	o.Instructions = &v
}

// GetInstructionsLink returns the InstructionsLink field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TemplateUpdate) GetInstructionsLink() string {
	if o == nil || IsNil(o.InstructionsLink.Get()) {
		var ret string
		return ret
	}
	return *o.InstructionsLink.Get()
}

// GetInstructionsLinkOk returns a tuple with the InstructionsLink field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TemplateUpdate) GetInstructionsLinkOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.InstructionsLink.Get(), o.InstructionsLink.IsSet()
}

// HasInstructionsLink returns a boolean if a field has been set.
func (o *TemplateUpdate) HasInstructionsLink() bool {
	if o != nil && o.InstructionsLink.IsSet() {
		return true
	}

	return false
}

// SetInstructionsLink gets a reference to the given NullableString and assigns it to the InstructionsLink field.
func (o *TemplateUpdate) SetInstructionsLink(v string) {
	o.InstructionsLink.Set(&v)
}
// SetInstructionsLinkNil sets the value for InstructionsLink to be an explicit nil
func (o *TemplateUpdate) SetInstructionsLinkNil() {
	o.InstructionsLink.Set(nil)
}

// UnsetInstructionsLink ensures that no value is present for InstructionsLink, not even an explicit nil
func (o *TemplateUpdate) UnsetInstructionsLink() {
	o.InstructionsLink.Unset()
}

// GetKind returns the Kind field value if set, zero value otherwise.
func (o *TemplateUpdate) GetKind() TransformationTemplateKind {
	if o == nil || IsNil(o.Kind) {
		var ret TransformationTemplateKind
		return ret
	}
	return *o.Kind
}

// GetKindOk returns a tuple with the Kind field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TemplateUpdate) GetKindOk() (*TransformationTemplateKind, bool) {
	if o == nil || IsNil(o.Kind) {
		return nil, false
	}
	return o.Kind, true
}

// HasKind returns a boolean if a field has been set.
func (o *TemplateUpdate) HasKind() bool {
	if o != nil && !IsNil(o.Kind) {
		return true
	}

	return false
}

// SetKind gets a reference to the given TransformationTemplateKind and assigns it to the Kind field.
func (o *TemplateUpdate) SetKind(v TransformationTemplateKind) {
	o.Kind = &v
}

// GetLogo returns the Logo field value
func (o *TemplateUpdate) GetLogo() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Logo
}

// GetLogoOk returns a tuple with the Logo field value
// and a boolean to check if the value has been set.
func (o *TemplateUpdate) GetLogoOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Logo, true
}

// SetLogo sets field value
func (o *TemplateUpdate) SetLogo(v string) {
	o.Logo = v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *TemplateUpdate) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TemplateUpdate) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *TemplateUpdate) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *TemplateUpdate) SetName(v string) {
	o.Name = &v
}

// GetTransformation returns the Transformation field value
func (o *TemplateUpdate) GetTransformation() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Transformation
}

// GetTransformationOk returns a tuple with the Transformation field value
// and a boolean to check if the value has been set.
func (o *TemplateUpdate) GetTransformationOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Transformation, true
}

// SetTransformation sets field value
func (o *TemplateUpdate) SetTransformation(v string) {
	o.Transformation = v
}

func (o TemplateUpdate) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TemplateUpdate) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if o.FeatureFlag.IsSet() {
		toSerialize["featureFlag"] = o.FeatureFlag.Get()
	}
	if o.FilterTypes != nil {
		toSerialize["filterTypes"] = o.FilterTypes
	}
	if !IsNil(o.Instructions) {
		toSerialize["instructions"] = o.Instructions
	}
	if o.InstructionsLink.IsSet() {
		toSerialize["instructionsLink"] = o.InstructionsLink.Get()
	}
	if !IsNil(o.Kind) {
		toSerialize["kind"] = o.Kind
	}
	toSerialize["logo"] = o.Logo
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	toSerialize["transformation"] = o.Transformation
	return toSerialize, nil
}

func (o *TemplateUpdate) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"logo",
		"transformation",
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

	varTemplateUpdate := _TemplateUpdate{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varTemplateUpdate)

	if err != nil {
		return err
	}

	*o = TemplateUpdate(varTemplateUpdate)

	return err
}

type NullableTemplateUpdate struct {
	value *TemplateUpdate
	isSet bool
}

func (v NullableTemplateUpdate) Get() *TemplateUpdate {
	return v.value
}

func (v *NullableTemplateUpdate) Set(val *TemplateUpdate) {
	v.value = val
	v.isSet = true
}

func (v NullableTemplateUpdate) IsSet() bool {
	return v.isSet
}

func (v *NullableTemplateUpdate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTemplateUpdate(val *TemplateUpdate) *NullableTemplateUpdate {
	return &NullableTemplateUpdate{value: val, isSet: true}
}

func (v NullableTemplateUpdate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTemplateUpdate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


