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

// ActionCollection struct for ActionCollection
type ActionCollection struct {
	// Absolute URL to the resource's representation
	Href *string `json:"href,omitempty"`
	// Unique identifier for the resource
	Id *string `json:"id,omitempty"`
	// Type of resource
	Type *string `json:"type,omitempty"`
	Items *[]Action `json:"items,omitempty"`
}



// GetHref returns the Href field value
// If the value is explicit nil, the zero value for string will be returned
func (o *ActionCollection) GetHref() *string {
	if o == nil {
		return nil
	}


	return o.Href

}

// GetHrefOk returns a tuple with the Href field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ActionCollection) GetHrefOk() (*string, bool) {
	if o == nil {
		return nil, false
	}


	return o.Href, true
}

// SetHref sets field value
func (o *ActionCollection) SetHref(v string) {


	o.Href = &v

}

// HasHref returns a boolean if a field has been set.
func (o *ActionCollection) HasHref() bool {
	if o != nil && o.Href != nil {
		return true
	}

	return false
}



// GetId returns the Id field value
// If the value is explicit nil, the zero value for string will be returned
func (o *ActionCollection) GetId() *string {
	if o == nil {
		return nil
	}


	return o.Id

}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ActionCollection) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}


	return o.Id, true
}

// SetId sets field value
func (o *ActionCollection) SetId(v string) {


	o.Id = &v

}

// HasId returns a boolean if a field has been set.
func (o *ActionCollection) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}



// GetType returns the Type field value
// If the value is explicit nil, the zero value for string will be returned
func (o *ActionCollection) GetType() *string {
	if o == nil {
		return nil
	}


	return o.Type

}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ActionCollection) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}


	return o.Type, true
}

// SetType sets field value
func (o *ActionCollection) SetType(v string) {


	o.Type = &v

}

// HasType returns a boolean if a field has been set.
func (o *ActionCollection) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}



// GetItems returns the Items field value
// If the value is explicit nil, the zero value for []Action will be returned
func (o *ActionCollection) GetItems() *[]Action {
	if o == nil {
		return nil
	}


	return o.Items

}

// GetItemsOk returns a tuple with the Items field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ActionCollection) GetItemsOk() (*[]Action, bool) {
	if o == nil {
		return nil, false
	}


	return o.Items, true
}

// SetItems sets field value
func (o *ActionCollection) SetItems(v []Action) {


	o.Items = &v

}

// HasItems returns a boolean if a field has been set.
func (o *ActionCollection) HasItems() bool {
	if o != nil && o.Items != nil {
		return true
	}

	return false
}


func (o ActionCollection) MarshalJSON() ([]byte, error) {
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

type NullableActionCollection struct {
	value *ActionCollection
	isSet bool
}

func (v NullableActionCollection) Get() *ActionCollection {
	return v.value
}

func (v *NullableActionCollection) Set(val *ActionCollection) {
	v.value = val
	v.isSet = true
}

func (v NullableActionCollection) IsSet() bool {
	return v.isSet
}

func (v *NullableActionCollection) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableActionCollection(val *ActionCollection) *NullableActionCollection {
	return &NullableActionCollection{value: val, isSet: true}
}

func (v NullableActionCollection) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableActionCollection) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


