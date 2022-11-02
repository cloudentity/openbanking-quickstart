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

// PayeeForUpdateentity PayeeForUpdateentity
//
// Payee's fields to be updated
//
// swagger:model PayeeForUpdateentity
type PayeeForUpdateentity struct {

	// merchant
	// Required: true
	Merchant *MerchantForUpdateentity1 `json:"merchant"`
}

// Validate validates this payee for updateentity
func (m *PayeeForUpdateentity) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateMerchant(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PayeeForUpdateentity) validateMerchant(formats strfmt.Registry) error {

	if err := validate.Required("merchant", "body", m.Merchant); err != nil {
		return err
	}

	if m.Merchant != nil {
		if err := m.Merchant.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("merchant")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("merchant")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this payee for updateentity based on the context it is used
func (m *PayeeForUpdateentity) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateMerchant(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PayeeForUpdateentity) contextValidateMerchant(ctx context.Context, formats strfmt.Registry) error {

	if m.Merchant != nil {
		if err := m.Merchant.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("merchant")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("merchant")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *PayeeForUpdateentity) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PayeeForUpdateentity) UnmarshalBinary(b []byte) error {
	var res PayeeForUpdateentity
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}