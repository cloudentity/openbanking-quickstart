// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// RegistryReference RegistryReference
//
// Used for registry references. In snake case to match IETF RFC 7591 naming formats
//
// swagger:model RegistryReference
type RegistryReference struct {

	// An ID representing the intermediary that can be looked up from a legal identity registry source
	RegisteredEntityID string `json:"registered_entity_id,omitempty"`

	// The legal company name for the intermediary
	RegisteredEntityName string `json:"registered_entity_name,omitempty"`

	// registry
	Registry Registry2 `json:"registry,omitempty"`
}

// Validate validates this registry reference
func (m *RegistryReference) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateRegistry(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *RegistryReference) validateRegistry(formats strfmt.Registry) error {
	if swag.IsZero(m.Registry) { // not required
		return nil
	}

	if err := m.Registry.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("registry")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("registry")
		}
		return err
	}

	return nil
}

// ContextValidate validate this registry reference based on the context it is used
func (m *RegistryReference) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateRegistry(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *RegistryReference) contextValidateRegistry(ctx context.Context, formats strfmt.Registry) error {

	if err := m.Registry.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("registry")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("registry")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *RegistryReference) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *RegistryReference) UnmarshalBinary(b []byte) error {
	var res RegistryReference
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
