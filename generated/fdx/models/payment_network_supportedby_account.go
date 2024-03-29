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

// PaymentNetworkSupportedbyAccount PaymentNetworkSupportedbyAccount
//
// This provides details required to execute a transaction against the account within the payment network
//
// swagger:model PaymentNetworkSupportedbyAccount
type PaymentNetworkSupportedbyAccount struct {

	// Bank identifier used by the payment network ie. Routing Number
	BankID string `json:"bankId,omitempty"`

	// The number used to identify the account within the payment network. If identifierType is ACCOUNT_NUMBER, this is the account number; if identifierType is TOKENIZED_ACCOUNT_NUMBER, this is a tokenized account number
	Identifier string `json:"identifier,omitempty"`

	// identifier type
	IdentifierType PaymentNetworkIdentifierType2 `json:"identifierType,omitempty"`

	// Can transfer funds to the account using this information
	TransferIn bool `json:"transferIn,omitempty"`

	// Can transfer funds from the account using this information
	TransferOut bool `json:"transferOut,omitempty"`

	// type
	Type PaymentNetworkType2 `json:"type,omitempty"`
}

// Validate validates this payment network supportedby account
func (m *PaymentNetworkSupportedbyAccount) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateIdentifierType(formats); err != nil {
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

func (m *PaymentNetworkSupportedbyAccount) validateIdentifierType(formats strfmt.Registry) error {
	if swag.IsZero(m.IdentifierType) { // not required
		return nil
	}

	if err := m.IdentifierType.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("identifierType")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("identifierType")
		}
		return err
	}

	return nil
}

func (m *PaymentNetworkSupportedbyAccount) validateType(formats strfmt.Registry) error {
	if swag.IsZero(m.Type) { // not required
		return nil
	}

	if err := m.Type.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("type")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("type")
		}
		return err
	}

	return nil
}

// ContextValidate validate this payment network supportedby account based on the context it is used
func (m *PaymentNetworkSupportedbyAccount) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateIdentifierType(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateType(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PaymentNetworkSupportedbyAccount) contextValidateIdentifierType(ctx context.Context, formats strfmt.Registry) error {

	if err := m.IdentifierType.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("identifierType")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("identifierType")
		}
		return err
	}

	return nil
}

func (m *PaymentNetworkSupportedbyAccount) contextValidateType(ctx context.Context, formats strfmt.Registry) error {

	if err := m.Type.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("type")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("type")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *PaymentNetworkSupportedbyAccount) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PaymentNetworkSupportedbyAccount) UnmarshalBinary(b []byte) error {
	var res PaymentNetworkSupportedbyAccount
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
