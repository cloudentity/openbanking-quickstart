// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// Data6 Data6
//
// swagger:model Data6
type Data6 struct {

	// The list of authorisations returned
	// Required: true
	DirectDebitAuthorisations []*BankingDirectDebit `json:"directDebitAuthorisations"`
}

// Validate validates this data6
func (m *Data6) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDirectDebitAuthorisations(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Data6) validateDirectDebitAuthorisations(formats strfmt.Registry) error {

	if err := validate.Required("directDebitAuthorisations", "body", m.DirectDebitAuthorisations); err != nil {
		return err
	}

	for i := 0; i < len(m.DirectDebitAuthorisations); i++ {
		if swag.IsZero(m.DirectDebitAuthorisations[i]) { // not required
			continue
		}

		if m.DirectDebitAuthorisations[i] != nil {
			if err := m.DirectDebitAuthorisations[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("directDebitAuthorisations" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("directDebitAuthorisations" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this data6 based on the context it is used
func (m *Data6) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateDirectDebitAuthorisations(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Data6) contextValidateDirectDebitAuthorisations(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.DirectDebitAuthorisations); i++ {

		if m.DirectDebitAuthorisations[i] != nil {
			if err := m.DirectDebitAuthorisations[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("directDebitAuthorisations" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("directDebitAuthorisations" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *Data6) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Data6) UnmarshalBinary(b []byte) error {
	var res Data6
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}