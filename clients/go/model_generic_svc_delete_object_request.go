/*
Singulatron

Run and develop self-hosted AI apps. Your programmable in-house GPT. The Firebase for the AI age.

API version: 0.2
Contact: sales@singulatron.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// checks if the GenericSvcDeleteObjectRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GenericSvcDeleteObjectRequest{}

// GenericSvcDeleteObjectRequest struct for GenericSvcDeleteObjectRequest
type GenericSvcDeleteObjectRequest struct {
	Conditions []DatastoreCondition `json:"conditions,omitempty"`
	Table *string `json:"table,omitempty"`
}

// NewGenericSvcDeleteObjectRequest instantiates a new GenericSvcDeleteObjectRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGenericSvcDeleteObjectRequest() *GenericSvcDeleteObjectRequest {
	this := GenericSvcDeleteObjectRequest{}
	return &this
}

// NewGenericSvcDeleteObjectRequestWithDefaults instantiates a new GenericSvcDeleteObjectRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGenericSvcDeleteObjectRequestWithDefaults() *GenericSvcDeleteObjectRequest {
	this := GenericSvcDeleteObjectRequest{}
	return &this
}

// GetConditions returns the Conditions field value if set, zero value otherwise.
func (o *GenericSvcDeleteObjectRequest) GetConditions() []DatastoreCondition {
	if o == nil || IsNil(o.Conditions) {
		var ret []DatastoreCondition
		return ret
	}
	return o.Conditions
}

// GetConditionsOk returns a tuple with the Conditions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GenericSvcDeleteObjectRequest) GetConditionsOk() ([]DatastoreCondition, bool) {
	if o == nil || IsNil(o.Conditions) {
		return nil, false
	}
	return o.Conditions, true
}

// HasConditions returns a boolean if a field has been set.
func (o *GenericSvcDeleteObjectRequest) HasConditions() bool {
	if o != nil && !IsNil(o.Conditions) {
		return true
	}

	return false
}

// SetConditions gets a reference to the given []DatastoreCondition and assigns it to the Conditions field.
func (o *GenericSvcDeleteObjectRequest) SetConditions(v []DatastoreCondition) {
	o.Conditions = v
}

// GetTable returns the Table field value if set, zero value otherwise.
func (o *GenericSvcDeleteObjectRequest) GetTable() string {
	if o == nil || IsNil(o.Table) {
		var ret string
		return ret
	}
	return *o.Table
}

// GetTableOk returns a tuple with the Table field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GenericSvcDeleteObjectRequest) GetTableOk() (*string, bool) {
	if o == nil || IsNil(o.Table) {
		return nil, false
	}
	return o.Table, true
}

// HasTable returns a boolean if a field has been set.
func (o *GenericSvcDeleteObjectRequest) HasTable() bool {
	if o != nil && !IsNil(o.Table) {
		return true
	}

	return false
}

// SetTable gets a reference to the given string and assigns it to the Table field.
func (o *GenericSvcDeleteObjectRequest) SetTable(v string) {
	o.Table = &v
}

func (o GenericSvcDeleteObjectRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GenericSvcDeleteObjectRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Conditions) {
		toSerialize["conditions"] = o.Conditions
	}
	if !IsNil(o.Table) {
		toSerialize["table"] = o.Table
	}
	return toSerialize, nil
}

type NullableGenericSvcDeleteObjectRequest struct {
	value *GenericSvcDeleteObjectRequest
	isSet bool
}

func (v NullableGenericSvcDeleteObjectRequest) Get() *GenericSvcDeleteObjectRequest {
	return v.value
}

func (v *NullableGenericSvcDeleteObjectRequest) Set(val *GenericSvcDeleteObjectRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableGenericSvcDeleteObjectRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableGenericSvcDeleteObjectRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGenericSvcDeleteObjectRequest(val *GenericSvcDeleteObjectRequest) *NullableGenericSvcDeleteObjectRequest {
	return &NullableGenericSvcDeleteObjectRequest{value: val, isSet: true}
}

func (v NullableGenericSvcDeleteObjectRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGenericSvcDeleteObjectRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

