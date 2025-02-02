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

// checks if the AppPortalAccessOut type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AppPortalAccessOut{}

// AppPortalAccessOut struct for AppPortalAccessOut
type AppPortalAccessOut struct {
	Token string `json:"token"`
	Url string `json:"url"`
}

type _AppPortalAccessOut AppPortalAccessOut

// NewAppPortalAccessOut instantiates a new AppPortalAccessOut object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAppPortalAccessOut(token string, url string) *AppPortalAccessOut {
	this := AppPortalAccessOut{}
	this.Token = token
	this.Url = url
	return &this
}

// NewAppPortalAccessOutWithDefaults instantiates a new AppPortalAccessOut object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAppPortalAccessOutWithDefaults() *AppPortalAccessOut {
	this := AppPortalAccessOut{}
	return &this
}

// GetToken returns the Token field value
func (o *AppPortalAccessOut) GetToken() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Token
}

// GetTokenOk returns a tuple with the Token field value
// and a boolean to check if the value has been set.
func (o *AppPortalAccessOut) GetTokenOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Token, true
}

// SetToken sets field value
func (o *AppPortalAccessOut) SetToken(v string) {
	o.Token = v
}

// GetUrl returns the Url field value
func (o *AppPortalAccessOut) GetUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Url
}

// GetUrlOk returns a tuple with the Url field value
// and a boolean to check if the value has been set.
func (o *AppPortalAccessOut) GetUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Url, true
}

// SetUrl sets field value
func (o *AppPortalAccessOut) SetUrl(v string) {
	o.Url = v
}

func (o AppPortalAccessOut) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AppPortalAccessOut) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["token"] = o.Token
	toSerialize["url"] = o.Url
	return toSerialize, nil
}

func (o *AppPortalAccessOut) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"token",
		"url",
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

	varAppPortalAccessOut := _AppPortalAccessOut{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varAppPortalAccessOut)

	if err != nil {
		return err
	}

	*o = AppPortalAccessOut(varAppPortalAccessOut)

	return err
}

type NullableAppPortalAccessOut struct {
	value *AppPortalAccessOut
	isSet bool
}

func (v NullableAppPortalAccessOut) Get() *AppPortalAccessOut {
	return v.value
}

func (v *NullableAppPortalAccessOut) Set(val *AppPortalAccessOut) {
	v.value = val
	v.isSet = true
}

func (v NullableAppPortalAccessOut) IsSet() bool {
	return v.isSet
}

func (v *NullableAppPortalAccessOut) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAppPortalAccessOut(val *AppPortalAccessOut) *NullableAppPortalAccessOut {
	return &NullableAppPortalAccessOut{value: val, isSet: true}
}

func (v NullableAppPortalAccessOut) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAppPortalAccessOut) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


