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
	"time"
)

// Metadata struct for Metadata
type Metadata struct {
	// The user who created the resource
	CreatedBy *string `json:"createdBy"`
	// The user who created the resource
	CreatedByUserId *string `json:"createdByUserId"`
	// When the resource was created
	CreatedDate *IonosTime
	// The etag for the resource
	Etag *string `json:"etag"`
	// The last user who modified the resource
	LastModifiedBy *string `json:"lastModifiedBy"`
	// The last user who modified the resource
	LastModifiedByUserId *string `json:"lastModifiedByUserId"`
	// When the resource was last modified
	LastModifiedDate *IonosTime
	State *MetadataState `json:"state"`
}



// GetCreatedBy returns the CreatedBy field value
// If the value is explicit nil, the zero value for string will be returned
func (o *Metadata) GetCreatedBy() *string {
	if o == nil {
		return nil
	}


	return o.CreatedBy

}

// GetCreatedByOk returns a tuple with the CreatedBy field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Metadata) GetCreatedByOk() (*string, bool) {
	if o == nil {
		return nil, false
	}


	return o.CreatedBy, true
}

// SetCreatedBy sets field value
func (o *Metadata) SetCreatedBy(v string) {


	o.CreatedBy = &v

}

// HasCreatedBy returns a boolean if a field has been set.
func (o *Metadata) HasCreatedBy() bool {
	if o != nil && o.CreatedBy != nil {
		return true
	}

	return false
}



// GetCreatedByUserId returns the CreatedByUserId field value
// If the value is explicit nil, the zero value for string will be returned
func (o *Metadata) GetCreatedByUserId() *string {
	if o == nil {
		return nil
	}


	return o.CreatedByUserId

}

// GetCreatedByUserIdOk returns a tuple with the CreatedByUserId field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Metadata) GetCreatedByUserIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}


	return o.CreatedByUserId, true
}

// SetCreatedByUserId sets field value
func (o *Metadata) SetCreatedByUserId(v string) {


	o.CreatedByUserId = &v

}

// HasCreatedByUserId returns a boolean if a field has been set.
func (o *Metadata) HasCreatedByUserId() bool {
	if o != nil && o.CreatedByUserId != nil {
		return true
	}

	return false
}



// GetCreatedDate returns the CreatedDate field value
// If the value is explicit nil, the zero value for time.Time will be returned
func (o *Metadata) GetCreatedDate() *time.Time {
	if o == nil {
		return nil
	}

	if o.CreatedDate == nil {
		return nil
	}
	return &o.CreatedDate.Time


}

// GetCreatedDateOk returns a tuple with the CreatedDate field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Metadata) GetCreatedDateOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}

	if o.CreatedDate == nil {
		return nil, false
	}
	return &o.CreatedDate.Time, true

}

// SetCreatedDate sets field value
func (o *Metadata) SetCreatedDate(v time.Time) {

	o.CreatedDate = &IonosTime{v}


}

// HasCreatedDate returns a boolean if a field has been set.
func (o *Metadata) HasCreatedDate() bool {
	if o != nil && o.CreatedDate != nil {
		return true
	}

	return false
}



// GetEtag returns the Etag field value
// If the value is explicit nil, the zero value for string will be returned
func (o *Metadata) GetEtag() *string {
	if o == nil {
		return nil
	}


	return o.Etag

}

// GetEtagOk returns a tuple with the Etag field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Metadata) GetEtagOk() (*string, bool) {
	if o == nil {
		return nil, false
	}


	return o.Etag, true
}

// SetEtag sets field value
func (o *Metadata) SetEtag(v string) {


	o.Etag = &v

}

// HasEtag returns a boolean if a field has been set.
func (o *Metadata) HasEtag() bool {
	if o != nil && o.Etag != nil {
		return true
	}

	return false
}



// GetLastModifiedBy returns the LastModifiedBy field value
// If the value is explicit nil, the zero value for string will be returned
func (o *Metadata) GetLastModifiedBy() *string {
	if o == nil {
		return nil
	}


	return o.LastModifiedBy

}

// GetLastModifiedByOk returns a tuple with the LastModifiedBy field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Metadata) GetLastModifiedByOk() (*string, bool) {
	if o == nil {
		return nil, false
	}


	return o.LastModifiedBy, true
}

// SetLastModifiedBy sets field value
func (o *Metadata) SetLastModifiedBy(v string) {


	o.LastModifiedBy = &v

}

// HasLastModifiedBy returns a boolean if a field has been set.
func (o *Metadata) HasLastModifiedBy() bool {
	if o != nil && o.LastModifiedBy != nil {
		return true
	}

	return false
}



// GetLastModifiedByUserId returns the LastModifiedByUserId field value
// If the value is explicit nil, the zero value for string will be returned
func (o *Metadata) GetLastModifiedByUserId() *string {
	if o == nil {
		return nil
	}


	return o.LastModifiedByUserId

}

// GetLastModifiedByUserIdOk returns a tuple with the LastModifiedByUserId field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Metadata) GetLastModifiedByUserIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}


	return o.LastModifiedByUserId, true
}

// SetLastModifiedByUserId sets field value
func (o *Metadata) SetLastModifiedByUserId(v string) {


	o.LastModifiedByUserId = &v

}

// HasLastModifiedByUserId returns a boolean if a field has been set.
func (o *Metadata) HasLastModifiedByUserId() bool {
	if o != nil && o.LastModifiedByUserId != nil {
		return true
	}

	return false
}



// GetLastModifiedDate returns the LastModifiedDate field value
// If the value is explicit nil, the zero value for time.Time will be returned
func (o *Metadata) GetLastModifiedDate() *time.Time {
	if o == nil {
		return nil
	}

	if o.LastModifiedDate == nil {
		return nil
	}
	return &o.LastModifiedDate.Time


}

// GetLastModifiedDateOk returns a tuple with the LastModifiedDate field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Metadata) GetLastModifiedDateOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}

	if o.LastModifiedDate == nil {
		return nil, false
	}
	return &o.LastModifiedDate.Time, true

}

// SetLastModifiedDate sets field value
func (o *Metadata) SetLastModifiedDate(v time.Time) {

	o.LastModifiedDate = &IonosTime{v}


}

// HasLastModifiedDate returns a boolean if a field has been set.
func (o *Metadata) HasLastModifiedDate() bool {
	if o != nil && o.LastModifiedDate != nil {
		return true
	}

	return false
}



// GetState returns the State field value
// If the value is explicit nil, the zero value for MetadataState will be returned
func (o *Metadata) GetState() *MetadataState {
	if o == nil {
		return nil
	}


	return o.State

}

// GetStateOk returns a tuple with the State field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Metadata) GetStateOk() (*MetadataState, bool) {
	if o == nil {
		return nil, false
	}


	return o.State, true
}

// SetState sets field value
func (o *Metadata) SetState(v MetadataState) {


	o.State = &v

}

// HasState returns a boolean if a field has been set.
func (o *Metadata) HasState() bool {
	if o != nil && o.State != nil {
		return true
	}

	return false
}


func (o Metadata) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}

	if o.CreatedBy != nil {
		toSerialize["createdBy"] = o.CreatedBy
	}
	

	if o.CreatedByUserId != nil {
		toSerialize["createdByUserId"] = o.CreatedByUserId
	}
	

	if o.CreatedDate != nil {
		toSerialize["createdDate"] = o.CreatedDate
	}
	

	if o.Etag != nil {
		toSerialize["etag"] = o.Etag
	}
	

	if o.LastModifiedBy != nil {
		toSerialize["lastModifiedBy"] = o.LastModifiedBy
	}
	

	if o.LastModifiedByUserId != nil {
		toSerialize["lastModifiedByUserId"] = o.LastModifiedByUserId
	}
	

	if o.LastModifiedDate != nil {
		toSerialize["lastModifiedDate"] = o.LastModifiedDate
	}
	

	if o.State != nil {
		toSerialize["state"] = o.State
	}
	
	return json.Marshal(toSerialize)
}

type NullableMetadata struct {
	value *Metadata
	isSet bool
}

func (v NullableMetadata) Get() *Metadata {
	return v.value
}

func (v *NullableMetadata) Set(val *Metadata) {
	v.value = val
	v.isSet = true
}

func (v NullableMetadata) IsSet() bool {
	return v.isSet
}

func (v *NullableMetadata) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMetadata(val *Metadata) *NullableMetadata {
	return &NullableMetadata{value: val, isSet: true}
}

func (v NullableMetadata) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMetadata) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


