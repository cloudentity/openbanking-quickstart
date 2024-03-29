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

// OBBeneficiary5Detail o b beneficiary5 detail
//
// swagger:model OBBeneficiary5Detail
type OBBeneficiary5Detail struct {

	// account Id
	AccountID AccountID `json:"AccountId,omitempty"`

	// beneficiary Id
	BeneficiaryID BeneficiaryID `json:"BeneficiaryId,omitempty"`

	// beneficiary type
	BeneficiaryType OBBeneficiaryType1Code `json:"BeneficiaryType,omitempty"`

	// creditor account
	// Required: true
	CreditorAccount *OBCashAccount50 `json:"CreditorAccount"`

	// creditor agent
	CreditorAgent *OBBranchAndFinancialInstitutionIdentification60 `json:"CreditorAgent,omitempty"`

	// reference
	Reference Reference `json:"Reference,omitempty"`

	// supplementary data
	SupplementaryData OBSupplementaryData1 `json:"SupplementaryData,omitempty"`
}

// Validate validates this o b beneficiary5 detail
func (m *OBBeneficiary5Detail) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAccountID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateBeneficiaryID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateBeneficiaryType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCreditorAccount(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCreditorAgent(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateReference(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *OBBeneficiary5Detail) validateAccountID(formats strfmt.Registry) error {
	if swag.IsZero(m.AccountID) { // not required
		return nil
	}

	if err := m.AccountID.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("AccountId")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("AccountId")
		}
		return err
	}

	return nil
}

func (m *OBBeneficiary5Detail) validateBeneficiaryID(formats strfmt.Registry) error {
	if swag.IsZero(m.BeneficiaryID) { // not required
		return nil
	}

	if err := m.BeneficiaryID.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("BeneficiaryId")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("BeneficiaryId")
		}
		return err
	}

	return nil
}

func (m *OBBeneficiary5Detail) validateBeneficiaryType(formats strfmt.Registry) error {
	if swag.IsZero(m.BeneficiaryType) { // not required
		return nil
	}

	if err := m.BeneficiaryType.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("BeneficiaryType")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("BeneficiaryType")
		}
		return err
	}

	return nil
}

func (m *OBBeneficiary5Detail) validateCreditorAccount(formats strfmt.Registry) error {

	if err := validate.Required("CreditorAccount", "body", m.CreditorAccount); err != nil {
		return err
	}

	if m.CreditorAccount != nil {
		if err := m.CreditorAccount.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("CreditorAccount")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("CreditorAccount")
			}
			return err
		}
	}

	return nil
}

func (m *OBBeneficiary5Detail) validateCreditorAgent(formats strfmt.Registry) error {
	if swag.IsZero(m.CreditorAgent) { // not required
		return nil
	}

	if m.CreditorAgent != nil {
		if err := m.CreditorAgent.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("CreditorAgent")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("CreditorAgent")
			}
			return err
		}
	}

	return nil
}

func (m *OBBeneficiary5Detail) validateReference(formats strfmt.Registry) error {
	if swag.IsZero(m.Reference) { // not required
		return nil
	}

	if err := m.Reference.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("Reference")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("Reference")
		}
		return err
	}

	return nil
}

// ContextValidate validate this o b beneficiary5 detail based on the context it is used
func (m *OBBeneficiary5Detail) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateAccountID(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateBeneficiaryID(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateBeneficiaryType(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateCreditorAccount(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateCreditorAgent(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateReference(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *OBBeneficiary5Detail) contextValidateAccountID(ctx context.Context, formats strfmt.Registry) error {

	if err := m.AccountID.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("AccountId")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("AccountId")
		}
		return err
	}

	return nil
}

func (m *OBBeneficiary5Detail) contextValidateBeneficiaryID(ctx context.Context, formats strfmt.Registry) error {

	if err := m.BeneficiaryID.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("BeneficiaryId")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("BeneficiaryId")
		}
		return err
	}

	return nil
}

func (m *OBBeneficiary5Detail) contextValidateBeneficiaryType(ctx context.Context, formats strfmt.Registry) error {

	if err := m.BeneficiaryType.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("BeneficiaryType")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("BeneficiaryType")
		}
		return err
	}

	return nil
}

func (m *OBBeneficiary5Detail) contextValidateCreditorAccount(ctx context.Context, formats strfmt.Registry) error {

	if m.CreditorAccount != nil {
		if err := m.CreditorAccount.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("CreditorAccount")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("CreditorAccount")
			}
			return err
		}
	}

	return nil
}

func (m *OBBeneficiary5Detail) contextValidateCreditorAgent(ctx context.Context, formats strfmt.Registry) error {

	if m.CreditorAgent != nil {
		if err := m.CreditorAgent.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("CreditorAgent")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("CreditorAgent")
			}
			return err
		}
	}

	return nil
}

func (m *OBBeneficiary5Detail) contextValidateReference(ctx context.Context, formats strfmt.Registry) error {

	if err := m.Reference.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("Reference")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("Reference")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *OBBeneficiary5Detail) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *OBBeneficiary5Detail) UnmarshalBinary(b []byte) error {
	var res OBBeneficiary5Detail
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
