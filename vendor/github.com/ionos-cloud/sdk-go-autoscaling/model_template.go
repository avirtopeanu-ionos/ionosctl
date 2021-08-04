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

// Template struct for Template
type Template struct {
	// Absolute URL to the resource's representation
	Href *string `json:"href,omitempty"`
	// Unique identifier for the resource
	Id *string `json:"id,omitempty"`
	// Type of resource
	Type *string `json:"type,omitempty"`
	// Entities associated with this resource. Contents depend on the type of resource.
	Entities *map[string]interface{} `json:"entities,omitempty"`
	Metadata *Metadata `json:"metadata,omitempty"`
	Properties *TemplateProperties `json:"properties,omitempty"`
}



// GetHref returns the Href field value
// If the value is explicit nil, the zero value for string will be returned
func (o *Template) GetHref() *string {
	if o == nil {
		return nil
	}


	return o.Href

}

// GetHrefOk returns a tuple with the Href field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Template) GetHrefOk() (*string, bool) {
	if o == nil {
		return nil, false
	}


	return o.Href, true
}

// SetHref sets field value
func (o *Template) SetHref(v string) {


	o.Href = &v

}

// HasHref returns a boolean if a field has been set.
func (o *Template) HasHref() bool {
	if o != nil && o.Href != nil {
		return true
	}

	return false
}



// GetId returns the Id field value
// If the value is explicit nil, the zero value for string will be returned
func (o *Template) GetId() *string {
	if o == nil {
		return nil
	}


	return o.Id

}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Template) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}


	return o.Id, true
}

// SetId sets field value
func (o *Template) SetId(v string) {


	o.Id = &v

}

// HasId returns a boolean if a field has been set.
func (o *Template) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}



// GetType returns the Type field value
// If the value is explicit nil, the zero value for string will be returned
func (o *Template) GetType() *string {
	if o == nil {
		return nil
	}


	return o.Type

}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Template) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}


	return o.Type, true
}

// SetType sets field value
func (o *Template) SetType(v string) {


	o.Type = &v

}

// HasType returns a boolean if a field has been set.
func (o *Template) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}



// GetEntities returns the Entities field value
// If the value is explicit nil, the zero value for map[string]interface{} will be returned
func (o *Template) GetEntities() *map[string]interface{} {
	if o == nil {
		return nil
	}


	return o.Entities

}

// GetEntitiesOk returns a tuple with the Entities field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Template) GetEntitiesOk() (*map[string]interface{}, bool) {
	if o == nil {
		return nil, false
	}


	return o.Entities, true
}

// SetEntities sets field value
func (o *Template) SetEntities(v map[string]interface{}) {


	o.Entities = &v

}

// HasEntities returns a boolean if a field has been set.
func (o *Template) HasEntities() bool {
	if o != nil && o.Entities != nil {
		return true
	}

	return false
}



// GetMetadata returns the Metadata field value
// If the value is explicit nil, the zero value for Metadata will be returned
func (o *Template) GetMetadata() *Metadata {
	if o == nil {
		return nil
	}


	return o.Metadata

}

// GetMetadataOk returns a tuple with the Metadata field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Template) GetMetadataOk() (*Metadata, bool) {
	if o == nil {
		return nil, false
	}


	return o.Metadata, true
}

// SetMetadata sets field value
func (o *Template) SetMetadata(v Metadata) {


	o.Metadata = &v

}

// HasMetadata returns a boolean if a field has been set.
func (o *Template) HasMetadata() bool {
	if o != nil && o.Metadata != nil {
		return true
	}

	return false
}



// GetProperties returns the Properties field value
// If the value is explicit nil, the zero value for TemplateProperties will be returned
func (o *Template) GetProperties() *TemplateProperties {
	if o == nil {
		return nil
	}


	return o.Properties

}

// GetPropertiesOk returns a tuple with the Properties field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Template) GetPropertiesOk() (*TemplateProperties, bool) {
	if o == nil {
		return nil, false
	}


	return o.Properties, true
}

// SetProperties sets field value
func (o *Template) SetProperties(v TemplateProperties) {


	o.Properties = &v

}

// HasProperties returns a boolean if a field has been set.
func (o *Template) HasProperties() bool {
	if o != nil && o.Properties != nil {
		return true
	}

	return false
}


func (o Template) MarshalJSON() ([]byte, error) {
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
	

	if o.Entities != nil {
		toSerialize["entities"] = o.Entities
	}
	

	if o.Metadata != nil {
		toSerialize["metadata"] = o.Metadata
	}
	

	if o.Properties != nil {
		toSerialize["properties"] = o.Properties
	}
	
	return json.Marshal(toSerialize)
}

type NullableTemplate struct {
	value *Template
	isSet bool
}

func (v NullableTemplate) Get() *Template {
	return v.value
}

func (v *NullableTemplate) Set(val *Template) {
	v.value = val
	v.isSet = true
}

func (v NullableTemplate) IsSet() bool {
	return v.isSet
}

func (v *NullableTemplate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTemplate(val *Template) *NullableTemplate {
	return &NullableTemplate{value: val, isSet: true}
}

func (v NullableTemplate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTemplate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


