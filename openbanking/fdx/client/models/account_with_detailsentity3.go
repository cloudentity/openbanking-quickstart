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

// AccountWithDetailsentity3 AccountWithDetailsentity3
//
// swagger:model AccountWithDetailsentity3
type AccountWithDetailsentity3 struct {

	// investment account
	InvestmentAccount *InvestmentAccountentity2 `json:"investmentAccount,omitempty"`
}

// Validate validates this account with detailsentity3
func (m *AccountWithDetailsentity3) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateInvestmentAccount(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AccountWithDetailsentity3) validateInvestmentAccount(formats strfmt.Registry) error {
	if swag.IsZero(m.InvestmentAccount) { // not required
		return nil
	}

	if m.InvestmentAccount != nil {
		if err := m.InvestmentAccount.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("investmentAccount")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("investmentAccount")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this account with detailsentity3 based on the context it is used
func (m *AccountWithDetailsentity3) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateInvestmentAccount(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AccountWithDetailsentity3) contextValidateInvestmentAccount(ctx context.Context, formats strfmt.Registry) error {

	if m.InvestmentAccount != nil {
		if err := m.InvestmentAccount.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("investmentAccount")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("investmentAccount")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *AccountWithDetailsentity3) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AccountWithDetailsentity3) UnmarshalBinary(b []byte) error {
	var res AccountWithDetailsentity3
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
