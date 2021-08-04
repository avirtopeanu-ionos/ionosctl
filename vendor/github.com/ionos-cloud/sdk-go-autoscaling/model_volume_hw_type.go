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
	"fmt"
)

// VolumeHwType Storage Type for this template volume (SSD or HDD).
type VolumeHwType string

// List of VolumeHwType
const (
	HDD VolumeHwType = "HDD"
	SSD VolumeHwType = "SSD"
	SSD_PREMIUM VolumeHwType = "SSD_PREMIUM"
	SSD_STANDARD VolumeHwType = "SSD_STANDARD"
)

func (v *VolumeHwType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := VolumeHwType(value)
	for _, existing := range []VolumeHwType{ "HDD", "SSD", "SSD_PREMIUM", "SSD_STANDARD",   } {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid VolumeHwType", value)
}

// Ptr returns reference to VolumeHwType value
func (v VolumeHwType) Ptr() *VolumeHwType {
	return &v
}

type NullableVolumeHwType struct {
	value *VolumeHwType
	isSet bool
}

func (v NullableVolumeHwType) Get() *VolumeHwType {
	return v.value
}

func (v *NullableVolumeHwType) Set(val *VolumeHwType) {
	v.value = val
	v.isSet = true
}

func (v NullableVolumeHwType) IsSet() bool {
	return v.isSet
}

func (v *NullableVolumeHwType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVolumeHwType(val *VolumeHwType) *NullableVolumeHwType {
	return &NullableVolumeHwType{value: val, isSet: true}
}

func (v NullableVolumeHwType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVolumeHwType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

