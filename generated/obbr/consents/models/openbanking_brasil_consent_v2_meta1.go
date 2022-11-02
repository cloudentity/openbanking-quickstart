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

// OpenbankingBrasilConsentV2Meta1 Meta1
//
// Meta informaes referente  API requisitada.
//
// swagger:model OpenbankingBrasilConsentV2Meta1
type OpenbankingBrasilConsentV2Meta1 struct {

	// Data e hora da consulta, conforme especificao RFC-3339, formato UTC.
	// Example: 2021-05-21T08:30:00Z
	// Required: true
	// Format: date-time
	RequestDateTime strfmt.DateTime `json:"requestDateTime"`
}

// Validate validates this openbanking brasil consent v2 meta1
func (m *OpenbankingBrasilConsentV2Meta1) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateRequestDateTime(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *OpenbankingBrasilConsentV2Meta1) validateRequestDateTime(formats strfmt.Registry) error {

	if err := validate.Required("requestDateTime", "body", strfmt.DateTime(m.RequestDateTime)); err != nil {
		return err
	}

	if err := validate.FormatOf("requestDateTime", "body", "date-time", m.RequestDateTime.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this openbanking brasil consent v2 meta1 based on context it is used
func (m *OpenbankingBrasilConsentV2Meta1) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *OpenbankingBrasilConsentV2Meta1) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *OpenbankingBrasilConsentV2Meta1) UnmarshalBinary(b []byte) error {
	var res OpenbankingBrasilConsentV2Meta1
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}