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

// BankDetails BankDetails
//
// swagger:model BankDetails
type BankDetails struct {

	// Account Targeted for payment
	// Required: true
	AccountNumber *string `json:"accountNumber"`

	// bank address
	BankAddress *BankAddress `json:"bankAddress,omitempty"`

	// Swift bank code.  Aligns with standard [ISO 9362](https://www.iso.org/standard/60390.html)
	BeneficiaryBankBIC string `json:"beneficiaryBankBIC,omitempty"`

	// Number for the Clearing House Interbank Payments System
	ChipNumber string `json:"chipNumber,omitempty"`

	// Country of the recipient institution. A valid [ISO 3166 Alpha-3](https://www.iso.org/iso-3166-country-codes.html) country code
	// Required: true
	Country *string `json:"country"`

	// Number for Fedwire payment (Federal Reserve Wire Network)
	FedWireNumber string `json:"fedWireNumber,omitempty"`

	// The legal entity identifier (LEI) for the beneficiary.  Aligns with [ISO 17442](https://www.iso.org/standard/59771.html)
	LegalEntityIdentifier string `json:"legalEntityIdentifier,omitempty"`

	// International bank routing number
	RoutingNumber string `json:"routingNumber,omitempty"`

	// Sort code used for account identification in some jurisdictions
	SortCode string `json:"sortCode,omitempty"`
}

// Validate validates this bank details
func (m *BankDetails) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAccountNumber(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateBankAddress(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCountry(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *BankDetails) validateAccountNumber(formats strfmt.Registry) error {

	if err := validate.Required("accountNumber", "body", m.AccountNumber); err != nil {
		return err
	}

	return nil
}

func (m *BankDetails) validateBankAddress(formats strfmt.Registry) error {
	if swag.IsZero(m.BankAddress) { // not required
		return nil
	}

	if m.BankAddress != nil {
		if err := m.BankAddress.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("bankAddress")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("bankAddress")
			}
			return err
		}
	}

	return nil
}

func (m *BankDetails) validateCountry(formats strfmt.Registry) error {

	if err := validate.Required("country", "body", m.Country); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this bank details based on the context it is used
func (m *BankDetails) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateBankAddress(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *BankDetails) contextValidateBankAddress(ctx context.Context, formats strfmt.Registry) error {

	if m.BankAddress != nil {
		if err := m.BankAddress.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("bankAddress")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("bankAddress")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *BankDetails) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *BankDetails) UnmarshalBinary(b []byte) error {
	var res BankDetails
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
