/*
 * VM Auto Scaling Service (part of CloudAPI)
 *
 * Provides Endpoints to manage the Autoscaling feature by IONOS cloud
 *
 * API version: 1-SDK.1
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package ionoscloud

import (
	"encoding/json"
)

// ServerCollection struct for ServerCollection
type ServerCollection struct {
	// Absolute URL to the resource's representation
	Href *string `json:"href,omitempty"`
	// Unique identifier for the resource
	Id *string `json:"id,omitempty"`
	// Type of resource
	Type *string `json:"type,omitempty"`
	Items *[]Server `json:"items,omitempty"`
}



// GetHref returns the Href field value
// If the value is explicit nil, the zero value for string will be returned
func (o *ServerCollection) GetHref() *string {
	if o == nil {
		return nil
	}


	return o.Href

}

// GetHrefOk returns a tuple with the Href field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ServerCollection) GetHrefOk() (*string, bool) {
	if o == nil {
		return nil, false
	}


	return o.Href, true
}

// SetHref sets field value
func (o *ServerCollection) SetHref(v string) {


	o.Href = &v

}

// HasHref returns a boolean if a field has been set.
func (o *ServerCollection) HasHref() bool {
	if o != nil && o.Href != nil {
		return true
	}

	return false
}



// GetId returns the Id field value
// If the value is explicit nil, the zero value for string will be returned
func (o *ServerCollection) GetId() *string {
	if o == nil {
		return nil
	}


	return o.Id

}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ServerCollection) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}


	return o.Id, true
}

// SetId sets field value
func (o *ServerCollection) SetId(v string) {


	o.Id = &v

}

// HasId returns a boolean if a field has been set.
func (o *ServerCollection) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}



// GetType returns the Type field value
// If the value is explicit nil, the zero value for string will be returned
func (o *ServerCollection) GetType() *string {
	if o == nil {
		return nil
	}


	return o.Type

}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ServerCollection) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}


	return o.Type, true
}

// SetType sets field value
func (o *ServerCollection) SetType(v string) {


	o.Type = &v

}

// HasType returns a boolean if a field has been set.
func (o *ServerCollection) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}



// GetItems returns the Items field value
// If the value is explicit nil, the zero value for []Server will be returned
func (o *ServerCollection) GetItems() *[]Server {
	if o == nil {
		return nil
	}


	return o.Items

}

// GetItemsOk returns a tuple with the Items field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ServerCollection) GetItemsOk() (*[]Server, bool) {
	if o == nil {
		return nil, false
	}


	return o.Items, true
}

// SetItems sets field value
func (o *ServerCollection) SetItems(v []Server) {


	o.Items = &v

}

// HasItems returns a boolean if a field has been set.
func (o *ServerCollection) HasItems() bool {
	if o != nil && o.Items != nil {
		return true
	}

	return false
}


func (o ServerCollection) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}

	if o.Href != nil {
		toSerialize["href"] = o.Href
	}
	

	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	

	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	

	if o.Items != nil {
		toSerialize["items"] = o.Items
	}
	
	return json.Marshal(toSerialize)
}

type NullableServerCollection struct {
	value *ServerCollection
	isSet bool
}

func (v NullableServerCollection) Get() *ServerCollection {
	return v.value
}

func (v *NullableServerCollection) Set(val *ServerCollection) {
	v.value = val
	v.isSet = true
}

func (v NullableServerCollection) IsSet() bool {
	return v.isSet
}

func (v *NullableServerCollection) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableServerCollection(val *ServerCollection) *NullableServerCollection {
	return &NullableServerCollection{value: val, isSet: true}
}

func (v NullableServerCollection) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableServerCollection) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


