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

// TemplateCollection struct for TemplateCollection
type TemplateCollection struct {
	// Absolute URL to the resource's representation
	Href *string `json:"href,omitempty"`
	// Unique identifier for the resource
	Id *string `json:"id,omitempty"`
	// Type of resource
	Type *string `json:"type,omitempty"`
	Items *[]Resource `json:"items,omitempty"`
}



// GetHref returns the Href field value
// If the value is explicit nil, the zero value for string will be returned
func (o *TemplateCollection) GetHref() *string {
	if o == nil {
		return nil
	}


	return o.Href

}

// GetHrefOk returns a tuple with the Href field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TemplateCollection) GetHrefOk() (*string, bool) {
	if o == nil {
		return nil, false
	}


	return o.Href, true
}

// SetHref sets field value
func (o *TemplateCollection) SetHref(v string) {


	o.Href = &v

}

// HasHref returns a boolean if a field has been set.
func (o *TemplateCollection) HasHref() bool {
	if o != nil && o.Href != nil {
		return true
	}

	return false
}



// GetId returns the Id field value
// If the value is explicit nil, the zero value for string will be returned
func (o *TemplateCollection) GetId() *string {
	if o == nil {
		return nil
	}


	return o.Id

}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TemplateCollection) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}


	return o.Id, true
}

// SetId sets field value
func (o *TemplateCollection) SetId(v string) {


	o.Id = &v

}

// HasId returns a boolean if a field has been set.
func (o *TemplateCollection) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}



// GetType returns the Type field value
// If the value is explicit nil, the zero value for string will be returned
func (o *TemplateCollection) GetType() *string {
	if o == nil {
		return nil
	}


	return o.Type

}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TemplateCollection) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}


	return o.Type, true
}

// SetType sets field value
func (o *TemplateCollection) SetType(v string) {


	o.Type = &v

}

// HasType returns a boolean if a field has been set.
func (o *TemplateCollection) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}



// GetItems returns the Items field value
// If the value is explicit nil, the zero value for []Resource will be returned
func (o *TemplateCollection) GetItems() *[]Resource {
	if o == nil {
		return nil
	}


	return o.Items

}

// GetItemsOk returns a tuple with the Items field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TemplateCollection) GetItemsOk() (*[]Resource, bool) {
	if o == nil {
		return nil, false
	}


	return o.Items, true
}

// SetItems sets field value
func (o *TemplateCollection) SetItems(v []Resource) {


	o.Items = &v

}

// HasItems returns a boolean if a field has been set.
func (o *TemplateCollection) HasItems() bool {
	if o != nil && o.Items != nil {
		return true
	}

	return false
}


func (o TemplateCollection) MarshalJSON() ([]byte, error) {
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

type NullableTemplateCollection struct {
	value *TemplateCollection
	isSet bool
}

func (v NullableTemplateCollection) Get() *TemplateCollection {
	return v.value
}

func (v *NullableTemplateCollection) Set(val *TemplateCollection) {
	v.value = val
	v.isSet = true
}

func (v NullableTemplateCollection) IsSet() bool {
	return v.isSet
}

func (v *NullableTemplateCollection) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTemplateCollection(val *TemplateCollection) *NullableTemplateCollection {
	return &NullableTemplateCollection{value: val, isSet: true}
}

func (v NullableTemplateCollection) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTemplateCollection) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


