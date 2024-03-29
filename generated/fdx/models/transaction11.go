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

// Transaction11 Transaction11
//
// swagger:model Transaction11
type Transaction11 struct {

	// loan transaction
	LoanTransaction *LoanTransactionentity1 `json:"loanTransaction,omitempty"`
}

// Validate validates this transaction11
func (m *Transaction11) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLoanTransaction(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Transaction11) validateLoanTransaction(formats strfmt.Registry) error {
	if swag.IsZero(m.LoanTransaction) { // not required
		return nil
	}

	if m.LoanTransaction != nil {
		if err := m.LoanTransaction.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("loanTransaction")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("loanTransaction")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this transaction11 based on the context it is used
func (m *Transaction11) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateLoanTransaction(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Transaction11) contextValidateLoanTransaction(ctx context.Context, formats strfmt.Registry) error {

	if m.LoanTransaction != nil {
		if err := m.LoanTransaction.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("loanTransaction")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("loanTransaction")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Transaction11) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Transaction11) UnmarshalBinary(b []byte) error {
	var res Transaction11
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
