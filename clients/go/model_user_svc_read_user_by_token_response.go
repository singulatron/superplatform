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

// checks if the UserSvcReadUserByTokenResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UserSvcReadUserByTokenResponse{}

// UserSvcReadUserByTokenResponse struct for UserSvcReadUserByTokenResponse
type UserSvcReadUserByTokenResponse struct {
	User *UserSvcUser `json:"user,omitempty"`
}

// NewUserSvcReadUserByTokenResponse instantiates a new UserSvcReadUserByTokenResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserSvcReadUserByTokenResponse() *UserSvcReadUserByTokenResponse {
	this := UserSvcReadUserByTokenResponse{}
	return &this
}

// NewUserSvcReadUserByTokenResponseWithDefaults instantiates a new UserSvcReadUserByTokenResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserSvcReadUserByTokenResponseWithDefaults() *UserSvcReadUserByTokenResponse {
	this := UserSvcReadUserByTokenResponse{}
	return &this
}

// GetUser returns the User field value if set, zero value otherwise.
func (o *UserSvcReadUserByTokenResponse) GetUser() UserSvcUser {
	if o == nil || IsNil(o.User) {
		var ret UserSvcUser
		return ret
	}
	return *o.User
}

// GetUserOk returns a tuple with the User field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserSvcReadUserByTokenResponse) GetUserOk() (*UserSvcUser, bool) {
	if o == nil || IsNil(o.User) {
		return nil, false
	}
	return o.User, true
}

// HasUser returns a boolean if a field has been set.
func (o *UserSvcReadUserByTokenResponse) HasUser() bool {
	if o != nil && !IsNil(o.User) {
		return true
	}

	return false
}

// SetUser gets a reference to the given UserSvcUser and assigns it to the User field.
func (o *UserSvcReadUserByTokenResponse) SetUser(v UserSvcUser) {
	o.User = &v
}

func (o UserSvcReadUserByTokenResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UserSvcReadUserByTokenResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.User) {
		toSerialize["user"] = o.User
	}
	return toSerialize, nil
}

type NullableUserSvcReadUserByTokenResponse struct {
	value *UserSvcReadUserByTokenResponse
	isSet bool
}

func (v NullableUserSvcReadUserByTokenResponse) Get() *UserSvcReadUserByTokenResponse {
	return v.value
}

func (v *NullableUserSvcReadUserByTokenResponse) Set(val *UserSvcReadUserByTokenResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableUserSvcReadUserByTokenResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableUserSvcReadUserByTokenResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserSvcReadUserByTokenResponse(val *UserSvcReadUserByTokenResponse) *NullableUserSvcReadUserByTokenResponse {
	return &NullableUserSvcReadUserByTokenResponse{value: val, isSet: true}
}

func (v NullableUserSvcReadUserByTokenResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserSvcReadUserByTokenResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

