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

// Transaction12 Transaction12
//
// swagger:model Transaction12
type Transaction12 struct {

	// loc transaction
	LocTransaction *LineOfCreditTransactionentity `json:"locTransaction,omitempty"`
}

// Validate validates this transaction12
func (m *Transaction12) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLocTransaction(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Transaction12) validateLocTransaction(formats strfmt.Registry) error {
	if swag.IsZero(m.LocTransaction) { // not required
		return nil
	}

	if m.LocTransaction != nil {
		if err := m.LocTransaction.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("locTransaction")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("locTransaction")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this transaction12 based on the context it is used
func (m *Transaction12) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateLocTransaction(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Transaction12) contextValidateLocTransaction(ctx context.Context, formats strfmt.Registry) error {

	if m.LocTransaction != nil {
		if err := m.LocTransaction.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("locTransaction")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("locTransaction")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Transaction12) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Transaction12) UnmarshalBinary(b []byte) error {
	var res Transaction12
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
