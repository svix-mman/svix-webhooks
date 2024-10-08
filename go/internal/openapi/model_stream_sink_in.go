/*
Svix API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 1.1.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"gopkg.in/validator.v2"
	"fmt"
)

// StreamSinkIn - struct for StreamSinkIn
type StreamSinkIn struct {
	StreamSinkInOneOf *StreamSinkInOneOf
	StreamSinkInOneOf1 *StreamSinkInOneOf1
	StreamSinkInOneOf2 *StreamSinkInOneOf2
	StreamSinkInOneOf3 *StreamSinkInOneOf3
	StreamSinkInOneOf4 *StreamSinkInOneOf4
	StreamSinkInOneOf5 *StreamSinkInOneOf5
	StreamSinkInOneOf6 *StreamSinkInOneOf6
}

// StreamSinkInOneOfAsStreamSinkIn is a convenience function that returns StreamSinkInOneOf wrapped in StreamSinkIn
func StreamSinkInOneOfAsStreamSinkIn(v *StreamSinkInOneOf) StreamSinkIn {
	return StreamSinkIn{
		StreamSinkInOneOf: v,
	}
}

// StreamSinkInOneOf1AsStreamSinkIn is a convenience function that returns StreamSinkInOneOf1 wrapped in StreamSinkIn
func StreamSinkInOneOf1AsStreamSinkIn(v *StreamSinkInOneOf1) StreamSinkIn {
	return StreamSinkIn{
		StreamSinkInOneOf1: v,
	}
}

// StreamSinkInOneOf2AsStreamSinkIn is a convenience function that returns StreamSinkInOneOf2 wrapped in StreamSinkIn
func StreamSinkInOneOf2AsStreamSinkIn(v *StreamSinkInOneOf2) StreamSinkIn {
	return StreamSinkIn{
		StreamSinkInOneOf2: v,
	}
}

// StreamSinkInOneOf3AsStreamSinkIn is a convenience function that returns StreamSinkInOneOf3 wrapped in StreamSinkIn
func StreamSinkInOneOf3AsStreamSinkIn(v *StreamSinkInOneOf3) StreamSinkIn {
	return StreamSinkIn{
		StreamSinkInOneOf3: v,
	}
}

// StreamSinkInOneOf4AsStreamSinkIn is a convenience function that returns StreamSinkInOneOf4 wrapped in StreamSinkIn
func StreamSinkInOneOf4AsStreamSinkIn(v *StreamSinkInOneOf4) StreamSinkIn {
	return StreamSinkIn{
		StreamSinkInOneOf4: v,
	}
}

// StreamSinkInOneOf5AsStreamSinkIn is a convenience function that returns StreamSinkInOneOf5 wrapped in StreamSinkIn
func StreamSinkInOneOf5AsStreamSinkIn(v *StreamSinkInOneOf5) StreamSinkIn {
	return StreamSinkIn{
		StreamSinkInOneOf5: v,
	}
}

// StreamSinkInOneOf6AsStreamSinkIn is a convenience function that returns StreamSinkInOneOf6 wrapped in StreamSinkIn
func StreamSinkInOneOf6AsStreamSinkIn(v *StreamSinkInOneOf6) StreamSinkIn {
	return StreamSinkIn{
		StreamSinkInOneOf6: v,
	}
}


// Unmarshal JSON data into one of the pointers in the struct
func (dst *StreamSinkIn) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into StreamSinkInOneOf
	err = newStrictDecoder(data).Decode(&dst.StreamSinkInOneOf)
	if err == nil {
		jsonStreamSinkInOneOf, _ := json.Marshal(dst.StreamSinkInOneOf)
		if string(jsonStreamSinkInOneOf) == "{}" { // empty struct
			dst.StreamSinkInOneOf = nil
		} else {
			if err = validator.Validate(dst.StreamSinkInOneOf); err != nil {
				dst.StreamSinkInOneOf = nil
			} else {
				match++
			}
		}
	} else {
		dst.StreamSinkInOneOf = nil
	}

	// try to unmarshal data into StreamSinkInOneOf1
	err = newStrictDecoder(data).Decode(&dst.StreamSinkInOneOf1)
	if err == nil {
		jsonStreamSinkInOneOf1, _ := json.Marshal(dst.StreamSinkInOneOf1)
		if string(jsonStreamSinkInOneOf1) == "{}" { // empty struct
			dst.StreamSinkInOneOf1 = nil
		} else {
			if err = validator.Validate(dst.StreamSinkInOneOf1); err != nil {
				dst.StreamSinkInOneOf1 = nil
			} else {
				match++
			}
		}
	} else {
		dst.StreamSinkInOneOf1 = nil
	}

	// try to unmarshal data into StreamSinkInOneOf2
	err = newStrictDecoder(data).Decode(&dst.StreamSinkInOneOf2)
	if err == nil {
		jsonStreamSinkInOneOf2, _ := json.Marshal(dst.StreamSinkInOneOf2)
		if string(jsonStreamSinkInOneOf2) == "{}" { // empty struct
			dst.StreamSinkInOneOf2 = nil
		} else {
			if err = validator.Validate(dst.StreamSinkInOneOf2); err != nil {
				dst.StreamSinkInOneOf2 = nil
			} else {
				match++
			}
		}
	} else {
		dst.StreamSinkInOneOf2 = nil
	}

	// try to unmarshal data into StreamSinkInOneOf3
	err = newStrictDecoder(data).Decode(&dst.StreamSinkInOneOf3)
	if err == nil {
		jsonStreamSinkInOneOf3, _ := json.Marshal(dst.StreamSinkInOneOf3)
		if string(jsonStreamSinkInOneOf3) == "{}" { // empty struct
			dst.StreamSinkInOneOf3 = nil
		} else {
			if err = validator.Validate(dst.StreamSinkInOneOf3); err != nil {
				dst.StreamSinkInOneOf3 = nil
			} else {
				match++
			}
		}
	} else {
		dst.StreamSinkInOneOf3 = nil
	}

	// try to unmarshal data into StreamSinkInOneOf4
	err = newStrictDecoder(data).Decode(&dst.StreamSinkInOneOf4)
	if err == nil {
		jsonStreamSinkInOneOf4, _ := json.Marshal(dst.StreamSinkInOneOf4)
		if string(jsonStreamSinkInOneOf4) == "{}" { // empty struct
			dst.StreamSinkInOneOf4 = nil
		} else {
			if err = validator.Validate(dst.StreamSinkInOneOf4); err != nil {
				dst.StreamSinkInOneOf4 = nil
			} else {
				match++
			}
		}
	} else {
		dst.StreamSinkInOneOf4 = nil
	}

	// try to unmarshal data into StreamSinkInOneOf5
	err = newStrictDecoder(data).Decode(&dst.StreamSinkInOneOf5)
	if err == nil {
		jsonStreamSinkInOneOf5, _ := json.Marshal(dst.StreamSinkInOneOf5)
		if string(jsonStreamSinkInOneOf5) == "{}" { // empty struct
			dst.StreamSinkInOneOf5 = nil
		} else {
			if err = validator.Validate(dst.StreamSinkInOneOf5); err != nil {
				dst.StreamSinkInOneOf5 = nil
			} else {
				match++
			}
		}
	} else {
		dst.StreamSinkInOneOf5 = nil
	}

	// try to unmarshal data into StreamSinkInOneOf6
	err = newStrictDecoder(data).Decode(&dst.StreamSinkInOneOf6)
	if err == nil {
		jsonStreamSinkInOneOf6, _ := json.Marshal(dst.StreamSinkInOneOf6)
		if string(jsonStreamSinkInOneOf6) == "{}" { // empty struct
			dst.StreamSinkInOneOf6 = nil
		} else {
			if err = validator.Validate(dst.StreamSinkInOneOf6); err != nil {
				dst.StreamSinkInOneOf6 = nil
			} else {
				match++
			}
		}
	} else {
		dst.StreamSinkInOneOf6 = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.StreamSinkInOneOf = nil
		dst.StreamSinkInOneOf1 = nil
		dst.StreamSinkInOneOf2 = nil
		dst.StreamSinkInOneOf3 = nil
		dst.StreamSinkInOneOf4 = nil
		dst.StreamSinkInOneOf5 = nil
		dst.StreamSinkInOneOf6 = nil

		return fmt.Errorf("data matches more than one schema in oneOf(StreamSinkIn)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(StreamSinkIn)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src StreamSinkIn) MarshalJSON() ([]byte, error) {
	if src.StreamSinkInOneOf != nil {
		return json.Marshal(&src.StreamSinkInOneOf)
	}

	if src.StreamSinkInOneOf1 != nil {
		return json.Marshal(&src.StreamSinkInOneOf1)
	}

	if src.StreamSinkInOneOf2 != nil {
		return json.Marshal(&src.StreamSinkInOneOf2)
	}

	if src.StreamSinkInOneOf3 != nil {
		return json.Marshal(&src.StreamSinkInOneOf3)
	}

	if src.StreamSinkInOneOf4 != nil {
		return json.Marshal(&src.StreamSinkInOneOf4)
	}

	if src.StreamSinkInOneOf5 != nil {
		return json.Marshal(&src.StreamSinkInOneOf5)
	}

	if src.StreamSinkInOneOf6 != nil {
		return json.Marshal(&src.StreamSinkInOneOf6)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *StreamSinkIn) GetActualInstance() (interface{}) {
	if obj == nil {
		return nil
	}
	if obj.StreamSinkInOneOf != nil {
		return obj.StreamSinkInOneOf
	}

	if obj.StreamSinkInOneOf1 != nil {
		return obj.StreamSinkInOneOf1
	}

	if obj.StreamSinkInOneOf2 != nil {
		return obj.StreamSinkInOneOf2
	}

	if obj.StreamSinkInOneOf3 != nil {
		return obj.StreamSinkInOneOf3
	}

	if obj.StreamSinkInOneOf4 != nil {
		return obj.StreamSinkInOneOf4
	}

	if obj.StreamSinkInOneOf5 != nil {
		return obj.StreamSinkInOneOf5
	}

	if obj.StreamSinkInOneOf6 != nil {
		return obj.StreamSinkInOneOf6
	}

	// all schemas are nil
	return nil
}

type NullableStreamSinkIn struct {
	value *StreamSinkIn
	isSet bool
}

func (v NullableStreamSinkIn) Get() *StreamSinkIn {
	return v.value
}

func (v *NullableStreamSinkIn) Set(val *StreamSinkIn) {
	v.value = val
	v.isSet = true
}

func (v NullableStreamSinkIn) IsSet() bool {
	return v.isSet
}

func (v *NullableStreamSinkIn) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStreamSinkIn(val *StreamSinkIn) *NullableStreamSinkIn {
	return &NullableStreamSinkIn{value: val, isSet: true}
}

func (v NullableStreamSinkIn) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStreamSinkIn) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


