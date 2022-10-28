// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// CustomerNameentity3 CustomerNameentity3
//
// Name of the payee used to execute the payment
//
// swagger:model CustomerNameentity3
type CustomerNameentity3 struct {

	// Company name
	Company string `json:"company,omitempty"`

	// First name
	First string `json:"first,omitempty"`

	// Last name
	Last string `json:"last,omitempty"`

	// Middle initial
	Middle string `json:"middle,omitempty"`

	// Name prefix, e.g. Mr.
	Prefix string `json:"prefix,omitempty"`

	// Generational or academic suffix
	Suffix string `json:"suffix,omitempty"`
}

// Validate validates this customer nameentity3
func (m *CustomerNameentity3) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this customer nameentity3 based on context it is used
func (m *CustomerNameentity3) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *CustomerNameentity3) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CustomerNameentity3) UnmarshalBinary(b []byte) error {
	var res CustomerNameentity3
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}