/*
Superplatform

On-premise AI platform and microservices ecosystem.

API version: 0.2
Contact: sales@singulatron.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// checks if the RegistrySvcListServiceDefinitionsResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RegistrySvcListServiceDefinitionsResponse{}

// RegistrySvcListServiceDefinitionsResponse struct for RegistrySvcListServiceDefinitionsResponse
type RegistrySvcListServiceDefinitionsResponse struct {
	ServiceDefinitions []RegistrySvcServiceDefinition `json:"serviceDefinitions,omitempty"`
}

// NewRegistrySvcListServiceDefinitionsResponse instantiates a new RegistrySvcListServiceDefinitionsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRegistrySvcListServiceDefinitionsResponse() *RegistrySvcListServiceDefinitionsResponse {
	this := RegistrySvcListServiceDefinitionsResponse{}
	return &this
}

// NewRegistrySvcListServiceDefinitionsResponseWithDefaults instantiates a new RegistrySvcListServiceDefinitionsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRegistrySvcListServiceDefinitionsResponseWithDefaults() *RegistrySvcListServiceDefinitionsResponse {
	this := RegistrySvcListServiceDefinitionsResponse{}
	return &this
}

// GetServiceDefinitions returns the ServiceDefinitions field value if set, zero value otherwise.
func (o *RegistrySvcListServiceDefinitionsResponse) GetServiceDefinitions() []RegistrySvcServiceDefinition {
	if o == nil || IsNil(o.ServiceDefinitions) {
		var ret []RegistrySvcServiceDefinition
		return ret
	}
	return o.ServiceDefinitions
}

// GetServiceDefinitionsOk returns a tuple with the ServiceDefinitions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RegistrySvcListServiceDefinitionsResponse) GetServiceDefinitionsOk() ([]RegistrySvcServiceDefinition, bool) {
	if o == nil || IsNil(o.ServiceDefinitions) {
		return nil, false
	}
	return o.ServiceDefinitions, true
}

// HasServiceDefinitions returns a boolean if a field has been set.
func (o *RegistrySvcListServiceDefinitionsResponse) HasServiceDefinitions() bool {
	if o != nil && !IsNil(o.ServiceDefinitions) {
		return true
	}

	return false
}

// SetServiceDefinitions gets a reference to the given []RegistrySvcServiceDefinition and assigns it to the ServiceDefinitions field.
func (o *RegistrySvcListServiceDefinitionsResponse) SetServiceDefinitions(v []RegistrySvcServiceDefinition) {
	o.ServiceDefinitions = v
}

func (o RegistrySvcListServiceDefinitionsResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RegistrySvcListServiceDefinitionsResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ServiceDefinitions) {
		toSerialize["serviceDefinitions"] = o.ServiceDefinitions
	}
	return toSerialize, nil
}

type NullableRegistrySvcListServiceDefinitionsResponse struct {
	value *RegistrySvcListServiceDefinitionsResponse
	isSet bool
}

func (v NullableRegistrySvcListServiceDefinitionsResponse) Get() *RegistrySvcListServiceDefinitionsResponse {
	return v.value
}

func (v *NullableRegistrySvcListServiceDefinitionsResponse) Set(val *RegistrySvcListServiceDefinitionsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableRegistrySvcListServiceDefinitionsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableRegistrySvcListServiceDefinitionsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRegistrySvcListServiceDefinitionsResponse(val *RegistrySvcListServiceDefinitionsResponse) *NullableRegistrySvcListServiceDefinitionsResponse {
	return &NullableRegistrySvcListServiceDefinitionsResponse{value: val, isSet: true}
}

func (v NullableRegistrySvcListServiceDefinitionsResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRegistrySvcListServiceDefinitionsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

