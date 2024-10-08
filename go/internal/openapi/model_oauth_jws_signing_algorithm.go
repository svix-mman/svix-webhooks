/*
Svix API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 1.1.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"fmt"
)

// OauthJwsSigningAlgorithm the model 'OauthJwsSigningAlgorithm'
type OauthJwsSigningAlgorithm string

// List of OauthJwsSigningAlgorithm
const (
	OAUTHJWSSIGNINGALGORITHM_RS256 OauthJwsSigningAlgorithm = "RS256"
)

// All allowed values of OauthJwsSigningAlgorithm enum
var AllowedOauthJwsSigningAlgorithmEnumValues = []OauthJwsSigningAlgorithm{
	"RS256",
}

func (v *OauthJwsSigningAlgorithm) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := OauthJwsSigningAlgorithm(value)
	for _, existing := range AllowedOauthJwsSigningAlgorithmEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid OauthJwsSigningAlgorithm", value)
}

// NewOauthJwsSigningAlgorithmFromValue returns a pointer to a valid OauthJwsSigningAlgorithm
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewOauthJwsSigningAlgorithmFromValue(v string) (*OauthJwsSigningAlgorithm, error) {
	ev := OauthJwsSigningAlgorithm(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for OauthJwsSigningAlgorithm: valid values are %v", v, AllowedOauthJwsSigningAlgorithmEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v OauthJwsSigningAlgorithm) IsValid() bool {
	for _, existing := range AllowedOauthJwsSigningAlgorithmEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to OauthJwsSigningAlgorithm value
func (v OauthJwsSigningAlgorithm) Ptr() *OauthJwsSigningAlgorithm {
	return &v
}

type NullableOauthJwsSigningAlgorithm struct {
	value *OauthJwsSigningAlgorithm
	isSet bool
}

func (v NullableOauthJwsSigningAlgorithm) Get() *OauthJwsSigningAlgorithm {
	return v.value
}

func (v *NullableOauthJwsSigningAlgorithm) Set(val *OauthJwsSigningAlgorithm) {
	v.value = val
	v.isSet = true
}

func (v NullableOauthJwsSigningAlgorithm) IsSet() bool {
	return v.isSet
}

func (v *NullableOauthJwsSigningAlgorithm) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOauthJwsSigningAlgorithm(val *OauthJwsSigningAlgorithm) *NullableOauthJwsSigningAlgorithm {
	return &NullableOauthJwsSigningAlgorithm{value: val, isSet: true}
}

func (v NullableOauthJwsSigningAlgorithm) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOauthJwsSigningAlgorithm) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

