/*
 * Auth API
 *
 * Use the Auth API to manage tokens for secure access to IONOS Cloud  APIs (Auth API, Cloud API, Reseller API, Activity Log API, and others).
 *
 * API version: 1.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package ionoscloud

import (
	"encoding/json"
)

// Tokens struct for Tokens
type Tokens struct {
	// Array of items in that collection.
	Tokens *[]Token `json:"tokens,omitempty"`
}

// GetTokens returns the Tokens field value
// If the value is explicit nil, the zero value for []Token will be returned
func (o *Tokens) GetTokens() *[]Token {
	if o == nil {
		return nil
	}

	return o.Tokens

}

// GetTokensOk returns a tuple with the Tokens field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Tokens) GetTokensOk() (*[]Token, bool) {
	if o == nil {
		return nil, false
	}

	return o.Tokens, true
}

// SetTokens sets field value
func (o *Tokens) SetTokens(v []Token) {

	o.Tokens = &v

}

// HasTokens returns a boolean if a field has been set.
func (o *Tokens) HasTokens() bool {
	if o != nil && o.Tokens != nil {
		return true
	}

	return false
}

func (o Tokens) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}

	if o.Tokens != nil {
		toSerialize["tokens"] = o.Tokens
	}
	return json.Marshal(toSerialize)
}

type NullableTokens struct {
	value *Tokens
	isSet bool
}

func (v NullableTokens) Get() *Tokens {
	return v.value
}

func (v *NullableTokens) Set(val *Tokens) {
	v.value = val
	v.isSet = true
}

func (v NullableTokens) IsSet() bool {
	return v.isSet
}

func (v *NullableTokens) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTokens(val *Tokens) *NullableTokens {
	return &NullableTokens{value: val, isSet: true}
}

func (v NullableTokens) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTokens) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
