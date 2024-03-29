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

// AccountWithDetailsentity AccountWithDetailsentity
//
// swagger:model AccountWithDetailsentity
type AccountWithDetailsentity struct {

	// deposit account
	DepositAccount *DepositAccountentity2 `json:"depositAccount,omitempty"`
}

// Validate validates this account with detailsentity
func (m *AccountWithDetailsentity) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDepositAccount(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AccountWithDetailsentity) validateDepositAccount(formats strfmt.Registry) error {
	if swag.IsZero(m.DepositAccount) { // not required
		return nil
	}

	if m.DepositAccount != nil {
		if err := m.DepositAccount.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("depositAccount")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("depositAccount")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this account with detailsentity based on the context it is used
func (m *AccountWithDetailsentity) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateDepositAccount(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AccountWithDetailsentity) contextValidateDepositAccount(ctx context.Context, formats strfmt.Registry) error {

	if m.DepositAccount != nil {
		if err := m.DepositAccount.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("depositAccount")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("depositAccount")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *AccountWithDetailsentity) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AccountWithDetailsentity) UnmarshalBinary(b []byte) error {
	var res AccountWithDetailsentity
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
