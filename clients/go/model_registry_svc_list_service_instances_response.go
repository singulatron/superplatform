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

// checks if the RegistrySvcListServiceInstancesResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RegistrySvcListServiceInstancesResponse{}

// RegistrySvcListServiceInstancesResponse struct for RegistrySvcListServiceInstancesResponse
type RegistrySvcListServiceInstancesResponse struct {
	ServiceInstances []RegistrySvcServiceInstance `json:"serviceInstances,omitempty"`
}

// NewRegistrySvcListServiceInstancesResponse instantiates a new RegistrySvcListServiceInstancesResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRegistrySvcListServiceInstancesResponse() *RegistrySvcListServiceInstancesResponse {
	this := RegistrySvcListServiceInstancesResponse{}
	return &this
}

// NewRegistrySvcListServiceInstancesResponseWithDefaults instantiates a new RegistrySvcListServiceInstancesResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRegistrySvcListServiceInstancesResponseWithDefaults() *RegistrySvcListServiceInstancesResponse {
	this := RegistrySvcListServiceInstancesResponse{}
	return &this
}

// GetServiceInstances returns the ServiceInstances field value if set, zero value otherwise.
func (o *RegistrySvcListServiceInstancesResponse) GetServiceInstances() []RegistrySvcServiceInstance {
	if o == nil || IsNil(o.ServiceInstances) {
		var ret []RegistrySvcServiceInstance
		return ret
	}
	return o.ServiceInstances
}

// GetServiceInstancesOk returns a tuple with the ServiceInstances field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RegistrySvcListServiceInstancesResponse) GetServiceInstancesOk() ([]RegistrySvcServiceInstance, bool) {
	if o == nil || IsNil(o.ServiceInstances) {
		return nil, false
	}
	return o.ServiceInstances, true
}

// HasServiceInstances returns a boolean if a field has been set.
func (o *RegistrySvcListServiceInstancesResponse) HasServiceInstances() bool {
	if o != nil && !IsNil(o.ServiceInstances) {
		return true
	}

	return false
}

// SetServiceInstances gets a reference to the given []RegistrySvcServiceInstance and assigns it to the ServiceInstances field.
func (o *RegistrySvcListServiceInstancesResponse) SetServiceInstances(v []RegistrySvcServiceInstance) {
	o.ServiceInstances = v
}

func (o RegistrySvcListServiceInstancesResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RegistrySvcListServiceInstancesResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ServiceInstances) {
		toSerialize["serviceInstances"] = o.ServiceInstances
	}
	return toSerialize, nil
}

type NullableRegistrySvcListServiceInstancesResponse struct {
	value *RegistrySvcListServiceInstancesResponse
	isSet bool
}

func (v NullableRegistrySvcListServiceInstancesResponse) Get() *RegistrySvcListServiceInstancesResponse {
	return v.value
}

func (v *NullableRegistrySvcListServiceInstancesResponse) Set(val *RegistrySvcListServiceInstancesResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableRegistrySvcListServiceInstancesResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableRegistrySvcListServiceInstancesResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRegistrySvcListServiceInstancesResponse(val *RegistrySvcListServiceInstancesResponse) *NullableRegistrySvcListServiceInstancesResponse {
	return &NullableRegistrySvcListServiceInstancesResponse{value: val, isSet: true}
}

func (v NullableRegistrySvcListServiceInstancesResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRegistrySvcListServiceInstancesResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

