// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// VolumeUsageData VolumeUsageData VolumeUsageData VolumeUsageData VolumeUsageData VolumeUsageData VolumeUsageData VolumeUsageData VolumeUsageData VolumeUsageData VolumeUsageData Usage details about the volume. This information is used by the
// `GET /system/df` endpoint, and omitted in other endpoints.
//
// swagger:model VolumeUsageData
type VolumeUsageData struct {

	// The number of containers referencing this volume. This field
	// is set to `-1` if the reference-count is not available.
	// Required: true
	RefCount *int64 `json:"RefCount"`

	// Amount of disk space used by the volume (in bytes). This information
	// is only available for volumes created with the `"local"` volume
	// driver. For volumes created with other volume drivers, this field
	// is set to `-1` ("not available")
	// Required: true
	Size *int64 `json:"Size"`
}

// Validate validates this volume usage data
func (m *VolumeUsageData) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateRefCount(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSize(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *VolumeUsageData) validateRefCount(formats strfmt.Registry) error {

	if err := validate.Required("RefCount", "body", m.RefCount); err != nil {
		return err
	}

	return nil
}

func (m *VolumeUsageData) validateSize(formats strfmt.Registry) error {

	if err := validate.Required("Size", "body", m.Size); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this volume usage data based on context it is used
func (m *VolumeUsageData) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *VolumeUsageData) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *VolumeUsageData) UnmarshalBinary(b []byte) error {
	var res VolumeUsageData
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
