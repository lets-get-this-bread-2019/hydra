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

// PluginConfigArgs PluginConfigArgs PluginConfigArgs PluginConfigArgs PluginConfigArgs PluginConfigArgs PluginConfigArgs PluginConfigArgs PluginConfigArgs PluginConfigArgs PluginConfigArgs plugin config args
//
// swagger:model PluginConfigArgs
type PluginConfigArgs struct {

	// description
	// Required: true
	Description *string `json:"Description"`

	// name
	// Required: true
	Name *string `json:"Name"`

	// settable
	// Required: true
	Settable []string `json:"Settable"`

	// value
	// Required: true
	Value []string `json:"Value"`
}

// Validate validates this plugin config args
func (m *PluginConfigArgs) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDescription(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSettable(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateValue(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PluginConfigArgs) validateDescription(formats strfmt.Registry) error {

	if err := validate.Required("Description", "body", m.Description); err != nil {
		return err
	}

	return nil
}

func (m *PluginConfigArgs) validateName(formats strfmt.Registry) error {

	if err := validate.Required("Name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

func (m *PluginConfigArgs) validateSettable(formats strfmt.Registry) error {

	if err := validate.Required("Settable", "body", m.Settable); err != nil {
		return err
	}

	return nil
}

func (m *PluginConfigArgs) validateValue(formats strfmt.Registry) error {

	if err := validate.Required("Value", "body", m.Value); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this plugin config args based on context it is used
func (m *PluginConfigArgs) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *PluginConfigArgs) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PluginConfigArgs) UnmarshalBinary(b []byte) error {
	var res PluginConfigArgs
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
