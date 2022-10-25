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

// OpenbankingBrasilConsentV2BusinessEntity BusinessEntity
//
// Titular, pessoa jurdica a quem se referem os dados que so objeto de compartilhamento.
//
// swagger:model OpenbankingBrasilConsentV2BusinessEntity
type OpenbankingBrasilConsentV2BusinessEntity struct {

	// document
	// Required: true
	Document *OpenbankingBrasilConsentV2Document1 `json:"document"`
}

// Validate validates this openbanking brasil consent v2 business entity
func (m *OpenbankingBrasilConsentV2BusinessEntity) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDocument(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *OpenbankingBrasilConsentV2BusinessEntity) validateDocument(formats strfmt.Registry) error {

	if err := validate.Required("document", "body", m.Document); err != nil {
		return err
	}

	if m.Document != nil {
		if err := m.Document.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("document")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("document")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this openbanking brasil consent v2 business entity based on the context it is used
func (m *OpenbankingBrasilConsentV2BusinessEntity) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateDocument(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *OpenbankingBrasilConsentV2BusinessEntity) contextValidateDocument(ctx context.Context, formats strfmt.Registry) error {

	if m.Document != nil {
		if err := m.Document.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("document")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("document")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *OpenbankingBrasilConsentV2BusinessEntity) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *OpenbankingBrasilConsentV2BusinessEntity) UnmarshalBinary(b []byte) error {
	var res OpenbankingBrasilConsentV2BusinessEntity
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
