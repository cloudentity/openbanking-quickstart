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

// BankingPayeeV2 BankingPayeeV2
//
// swagger:model BankingPayeeV2
type BankingPayeeV2 struct {

	// The date the payee was created by the customer
	CreationDate string `json:"creationDate,omitempty"`

	// A description of the payee provided by the customer
	Description string `json:"description,omitempty"`

	// The short display name of the payee as provided by the customer. Where a customer has not provided a nickname, a display name derived by the bank for the payee consistent with existing digital banking channels
	// Required: true
	Nickname *string `json:"nickname"`

	// ID of the payee adhering to the rules of ID permanence
	// Required: true
	PayeeID *string `json:"payeeId"`

	// type
	// Required: true
	Type *Type1 `json:"type"`
}

// Validate validates this banking payee v2
func (m *BankingPayeeV2) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateNickname(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePayeeID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *BankingPayeeV2) validateNickname(formats strfmt.Registry) error {

	if err := validate.Required("nickname", "body", m.Nickname); err != nil {
		return err
	}

	return nil
}

func (m *BankingPayeeV2) validatePayeeID(formats strfmt.Registry) error {

	if err := validate.Required("payeeId", "body", m.PayeeID); err != nil {
		return err
	}

	return nil
}

func (m *BankingPayeeV2) validateType(formats strfmt.Registry) error {

	if err := validate.Required("type", "body", m.Type); err != nil {
		return err
	}

	if err := validate.Required("type", "body", m.Type); err != nil {
		return err
	}

	if m.Type != nil {
		if err := m.Type.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("type")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this banking payee v2 based on the context it is used
func (m *BankingPayeeV2) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateType(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *BankingPayeeV2) contextValidateType(ctx context.Context, formats strfmt.Registry) error {

	if m.Type != nil {
		if err := m.Type.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("type")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *BankingPayeeV2) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *BankingPayeeV2) UnmarshalBinary(b []byte) error {
	var res BankingPayeeV2
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}