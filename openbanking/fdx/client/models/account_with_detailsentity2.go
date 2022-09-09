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

// AccountWithDetailsentity2 AccountWithDetailsentity2
//
// swagger:model AccountWithDetailsentity2
type AccountWithDetailsentity2 struct {

	// loc account
	LocAccount *LineOfCreditAccountentity `json:"locAccount,omitempty"`
}

// Validate validates this account with detailsentity2
func (m *AccountWithDetailsentity2) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLocAccount(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AccountWithDetailsentity2) validateLocAccount(formats strfmt.Registry) error {
	if swag.IsZero(m.LocAccount) { // not required
		return nil
	}

	if m.LocAccount != nil {
		if err := m.LocAccount.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("locAccount")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("locAccount")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this account with detailsentity2 based on the context it is used
func (m *AccountWithDetailsentity2) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateLocAccount(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AccountWithDetailsentity2) contextValidateLocAccount(ctx context.Context, formats strfmt.Registry) error {

	if m.LocAccount != nil {
		if err := m.LocAccount.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("locAccount")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("locAccount")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *AccountWithDetailsentity2) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AccountWithDetailsentity2) UnmarshalBinary(b []byte) error {
	var res AccountWithDetailsentity2
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
