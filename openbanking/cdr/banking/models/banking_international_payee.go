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

// BankingInternationalPayee BankingInternationalPayee
//
// swagger:model BankingInternationalPayee
type BankingInternationalPayee struct {

	// bank details
	// Required: true
	BankDetails *BankDetails `json:"bankDetails"`

	// beneficiary details
	// Required: true
	BeneficiaryDetails *BeneficiaryDetails `json:"beneficiaryDetails"`
}

// Validate validates this banking international payee
func (m *BankingInternationalPayee) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateBankDetails(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateBeneficiaryDetails(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *BankingInternationalPayee) validateBankDetails(formats strfmt.Registry) error {

	if err := validate.Required("bankDetails", "body", m.BankDetails); err != nil {
		return err
	}

	if m.BankDetails != nil {
		if err := m.BankDetails.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("bankDetails")
			}
			return err
		}
	}

	return nil
}

func (m *BankingInternationalPayee) validateBeneficiaryDetails(formats strfmt.Registry) error {

	if err := validate.Required("beneficiaryDetails", "body", m.BeneficiaryDetails); err != nil {
		return err
	}

	if m.BeneficiaryDetails != nil {
		if err := m.BeneficiaryDetails.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("beneficiaryDetails")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this banking international payee based on the context it is used
func (m *BankingInternationalPayee) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateBankDetails(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateBeneficiaryDetails(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *BankingInternationalPayee) contextValidateBankDetails(ctx context.Context, formats strfmt.Registry) error {

	if m.BankDetails != nil {
		if err := m.BankDetails.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("bankDetails")
			}
			return err
		}
	}

	return nil
}

func (m *BankingInternationalPayee) contextValidateBeneficiaryDetails(ctx context.Context, formats strfmt.Registry) error {

	if m.BeneficiaryDetails != nil {
		if err := m.BeneficiaryDetails.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("beneficiaryDetails")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *BankingInternationalPayee) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *BankingInternationalPayee) UnmarshalBinary(b []byte) error {
	var res BankingInternationalPayee
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
