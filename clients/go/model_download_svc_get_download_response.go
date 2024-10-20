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

// checks if the DownloadSvcGetDownloadResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DownloadSvcGetDownloadResponse{}

// DownloadSvcGetDownloadResponse struct for DownloadSvcGetDownloadResponse
type DownloadSvcGetDownloadResponse struct {
	Download *DownloadSvcDownloadDetails `json:"download,omitempty"`
	Exists *bool `json:"exists,omitempty"`
}

// NewDownloadSvcGetDownloadResponse instantiates a new DownloadSvcGetDownloadResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDownloadSvcGetDownloadResponse() *DownloadSvcGetDownloadResponse {
	this := DownloadSvcGetDownloadResponse{}
	return &this
}

// NewDownloadSvcGetDownloadResponseWithDefaults instantiates a new DownloadSvcGetDownloadResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDownloadSvcGetDownloadResponseWithDefaults() *DownloadSvcGetDownloadResponse {
	this := DownloadSvcGetDownloadResponse{}
	return &this
}

// GetDownload returns the Download field value if set, zero value otherwise.
func (o *DownloadSvcGetDownloadResponse) GetDownload() DownloadSvcDownloadDetails {
	if o == nil || IsNil(o.Download) {
		var ret DownloadSvcDownloadDetails
		return ret
	}
	return *o.Download
}

// GetDownloadOk returns a tuple with the Download field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DownloadSvcGetDownloadResponse) GetDownloadOk() (*DownloadSvcDownloadDetails, bool) {
	if o == nil || IsNil(o.Download) {
		return nil, false
	}
	return o.Download, true
}

// HasDownload returns a boolean if a field has been set.
func (o *DownloadSvcGetDownloadResponse) HasDownload() bool {
	if o != nil && !IsNil(o.Download) {
		return true
	}

	return false
}

// SetDownload gets a reference to the given DownloadSvcDownloadDetails and assigns it to the Download field.
func (o *DownloadSvcGetDownloadResponse) SetDownload(v DownloadSvcDownloadDetails) {
	o.Download = &v
}

// GetExists returns the Exists field value if set, zero value otherwise.
func (o *DownloadSvcGetDownloadResponse) GetExists() bool {
	if o == nil || IsNil(o.Exists) {
		var ret bool
		return ret
	}
	return *o.Exists
}

// GetExistsOk returns a tuple with the Exists field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DownloadSvcGetDownloadResponse) GetExistsOk() (*bool, bool) {
	if o == nil || IsNil(o.Exists) {
		return nil, false
	}
	return o.Exists, true
}

// HasExists returns a boolean if a field has been set.
func (o *DownloadSvcGetDownloadResponse) HasExists() bool {
	if o != nil && !IsNil(o.Exists) {
		return true
	}

	return false
}

// SetExists gets a reference to the given bool and assigns it to the Exists field.
func (o *DownloadSvcGetDownloadResponse) SetExists(v bool) {
	o.Exists = &v
}

func (o DownloadSvcGetDownloadResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DownloadSvcGetDownloadResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Download) {
		toSerialize["download"] = o.Download
	}
	if !IsNil(o.Exists) {
		toSerialize["exists"] = o.Exists
	}
	return toSerialize, nil
}

type NullableDownloadSvcGetDownloadResponse struct {
	value *DownloadSvcGetDownloadResponse
	isSet bool
}

func (v NullableDownloadSvcGetDownloadResponse) Get() *DownloadSvcGetDownloadResponse {
	return v.value
}

func (v *NullableDownloadSvcGetDownloadResponse) Set(val *DownloadSvcGetDownloadResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableDownloadSvcGetDownloadResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableDownloadSvcGetDownloadResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDownloadSvcGetDownloadResponse(val *DownloadSvcGetDownloadResponse) *NullableDownloadSvcGetDownloadResponse {
	return &NullableDownloadSvcGetDownloadResponse{value: val, isSet: true}
}

func (v NullableDownloadSvcGetDownloadResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDownloadSvcGetDownloadResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


