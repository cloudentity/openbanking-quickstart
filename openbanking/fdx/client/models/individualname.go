// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// Individualname Individualname
//
// First name, middle initial, last name, suffix fields
//
// swagger:model Individualname
type Individualname struct {

	// First name
	First string `json:"first,omitempty"`

	// Last name
	Last string `json:"last,omitempty"`

	// Middle initial
	Middle string `json:"middle,omitempty"`

	// Generational or academic suffix
	Suffix string `json:"suffix,omitempty"`
}

// Validate validates this individualname
func (m *Individualname) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this individualname based on context it is used
func (m *Individualname) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *Individualname) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Individualname) UnmarshalBinary(b []byte) error {
	var res Individualname
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
