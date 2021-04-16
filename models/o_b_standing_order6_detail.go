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

// OBStandingOrder6Detail o b standing order6 detail
//
// swagger:model OBStandingOrder6Detail
type OBStandingOrder6Detail struct {

	// account Id
	// Required: true
	AccountID *AccountID `json:"AccountId"`

	// creditor account
	// Required: true
	CreditorAccount *OBCashAccount51 `json:"CreditorAccount"`

	// creditor agent
	CreditorAgent *OBBranchAndFinancialInstitutionIdentification51 `json:"CreditorAgent,omitempty"`

	// final payment amount
	FinalPaymentAmount *OBActiveOrHistoricCurrencyAndAmount4 `json:"FinalPaymentAmount,omitempty"`

	// final payment date time
	// Format: date-time
	FinalPaymentDateTime FinalPaymentDateTime `json:"FinalPaymentDateTime,omitempty"`

	// first payment amount
	FirstPaymentAmount *OBActiveOrHistoricCurrencyAndAmount2 `json:"FirstPaymentAmount,omitempty"`

	// first payment date time
	// Format: date-time
	FirstPaymentDateTime FirstPaymentDateTime `json:"FirstPaymentDateTime,omitempty"`

	// frequency
	// Required: true
	Frequency *Frequency1 `json:"Frequency"`

	// last payment amount
	LastPaymentAmount *OBActiveOrHistoricCurrencyAndAmount11 `json:"LastPaymentAmount,omitempty"`

	// last payment date time
	// Format: date-time
	LastPaymentDateTime LastPaymentDateTime `json:"LastPaymentDateTime,omitempty"`

	// next payment amount
	NextPaymentAmount *OBActiveOrHistoricCurrencyAndAmount3 `json:"NextPaymentAmount,omitempty"`

	// next payment date time
	// Format: date-time
	NextPaymentDateTime NextPaymentDateTime `json:"NextPaymentDateTime,omitempty"`

	// number of payments
	NumberOfPayments NumberOfPayments `json:"NumberOfPayments,omitempty"`

	// reference
	Reference Reference `json:"Reference,omitempty"`

	// standing order Id
	StandingOrderID StandingOrderID `json:"StandingOrderId,omitempty"`

	// standing order status code
	StandingOrderStatusCode OBExternalStandingOrderStatus1Code `json:"StandingOrderStatusCode,omitempty"`

	// supplementary data
	SupplementaryData OBSupplementaryData1 `json:"SupplementaryData,omitempty"`
}

// Validate validates this o b standing order6 detail
func (m *OBStandingOrder6Detail) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAccountID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCreditorAccount(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCreditorAgent(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFinalPaymentAmount(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFinalPaymentDateTime(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFirstPaymentAmount(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFirstPaymentDateTime(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFrequency(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLastPaymentAmount(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLastPaymentDateTime(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNextPaymentAmount(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNextPaymentDateTime(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNumberOfPayments(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateReference(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStandingOrderID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStandingOrderStatusCode(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *OBStandingOrder6Detail) validateAccountID(formats strfmt.Registry) error {

	if err := validate.Required("AccountId", "body", m.AccountID); err != nil {
		return err
	}

	if err := validate.Required("AccountId", "body", m.AccountID); err != nil {
		return err
	}

	if m.AccountID != nil {
		if err := m.AccountID.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("AccountId")
			}
			return err
		}
	}

	return nil
}

func (m *OBStandingOrder6Detail) validateCreditorAccount(formats strfmt.Registry) error {

	if err := validate.Required("CreditorAccount", "body", m.CreditorAccount); err != nil {
		return err
	}

	if m.CreditorAccount != nil {
		if err := m.CreditorAccount.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("CreditorAccount")
			}
			return err
		}
	}

	return nil
}

func (m *OBStandingOrder6Detail) validateCreditorAgent(formats strfmt.Registry) error {
	if swag.IsZero(m.CreditorAgent) { // not required
		return nil
	}

	if m.CreditorAgent != nil {
		if err := m.CreditorAgent.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("CreditorAgent")
			}
			return err
		}
	}

	return nil
}

func (m *OBStandingOrder6Detail) validateFinalPaymentAmount(formats strfmt.Registry) error {
	if swag.IsZero(m.FinalPaymentAmount) { // not required
		return nil
	}

	if m.FinalPaymentAmount != nil {
		if err := m.FinalPaymentAmount.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("FinalPaymentAmount")
			}
			return err
		}
	}

	return nil
}

func (m *OBStandingOrder6Detail) validateFinalPaymentDateTime(formats strfmt.Registry) error {
	if swag.IsZero(m.FinalPaymentDateTime) { // not required
		return nil
	}

	if err := m.FinalPaymentDateTime.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("FinalPaymentDateTime")
		}
		return err
	}

	return nil
}

func (m *OBStandingOrder6Detail) validateFirstPaymentAmount(formats strfmt.Registry) error {
	if swag.IsZero(m.FirstPaymentAmount) { // not required
		return nil
	}

	if m.FirstPaymentAmount != nil {
		if err := m.FirstPaymentAmount.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("FirstPaymentAmount")
			}
			return err
		}
	}

	return nil
}

func (m *OBStandingOrder6Detail) validateFirstPaymentDateTime(formats strfmt.Registry) error {
	if swag.IsZero(m.FirstPaymentDateTime) { // not required
		return nil
	}

	if err := m.FirstPaymentDateTime.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("FirstPaymentDateTime")
		}
		return err
	}

	return nil
}

func (m *OBStandingOrder6Detail) validateFrequency(formats strfmt.Registry) error {

	if err := validate.Required("Frequency", "body", m.Frequency); err != nil {
		return err
	}

	if err := validate.Required("Frequency", "body", m.Frequency); err != nil {
		return err
	}

	if m.Frequency != nil {
		if err := m.Frequency.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("Frequency")
			}
			return err
		}
	}

	return nil
}

func (m *OBStandingOrder6Detail) validateLastPaymentAmount(formats strfmt.Registry) error {
	if swag.IsZero(m.LastPaymentAmount) { // not required
		return nil
	}

	if m.LastPaymentAmount != nil {
		if err := m.LastPaymentAmount.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("LastPaymentAmount")
			}
			return err
		}
	}

	return nil
}

func (m *OBStandingOrder6Detail) validateLastPaymentDateTime(formats strfmt.Registry) error {
	if swag.IsZero(m.LastPaymentDateTime) { // not required
		return nil
	}

	if err := m.LastPaymentDateTime.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("LastPaymentDateTime")
		}
		return err
	}

	return nil
}

func (m *OBStandingOrder6Detail) validateNextPaymentAmount(formats strfmt.Registry) error {
	if swag.IsZero(m.NextPaymentAmount) { // not required
		return nil
	}

	if m.NextPaymentAmount != nil {
		if err := m.NextPaymentAmount.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("NextPaymentAmount")
			}
			return err
		}
	}

	return nil
}

func (m *OBStandingOrder6Detail) validateNextPaymentDateTime(formats strfmt.Registry) error {
	if swag.IsZero(m.NextPaymentDateTime) { // not required
		return nil
	}

	if err := m.NextPaymentDateTime.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("NextPaymentDateTime")
		}
		return err
	}

	return nil
}

func (m *OBStandingOrder6Detail) validateNumberOfPayments(formats strfmt.Registry) error {
	if swag.IsZero(m.NumberOfPayments) { // not required
		return nil
	}

	if err := m.NumberOfPayments.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("NumberOfPayments")
		}
		return err
	}

	return nil
}

func (m *OBStandingOrder6Detail) validateReference(formats strfmt.Registry) error {
	if swag.IsZero(m.Reference) { // not required
		return nil
	}

	if err := m.Reference.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("Reference")
		}
		return err
	}

	return nil
}

func (m *OBStandingOrder6Detail) validateStandingOrderID(formats strfmt.Registry) error {
	if swag.IsZero(m.StandingOrderID) { // not required
		return nil
	}

	if err := m.StandingOrderID.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("StandingOrderId")
		}
		return err
	}

	return nil
}

func (m *OBStandingOrder6Detail) validateStandingOrderStatusCode(formats strfmt.Registry) error {
	if swag.IsZero(m.StandingOrderStatusCode) { // not required
		return nil
	}

	if err := m.StandingOrderStatusCode.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("StandingOrderStatusCode")
		}
		return err
	}

	return nil
}

// ContextValidate validate this o b standing order6 detail based on the context it is used
func (m *OBStandingOrder6Detail) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateAccountID(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateCreditorAccount(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateCreditorAgent(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateFinalPaymentAmount(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateFinalPaymentDateTime(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateFirstPaymentAmount(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateFirstPaymentDateTime(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateFrequency(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateLastPaymentAmount(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateLastPaymentDateTime(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateNextPaymentAmount(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateNextPaymentDateTime(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateNumberOfPayments(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateReference(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateStandingOrderID(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateStandingOrderStatusCode(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *OBStandingOrder6Detail) contextValidateAccountID(ctx context.Context, formats strfmt.Registry) error {

	if m.AccountID != nil {
		if err := m.AccountID.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("AccountId")
			}
			return err
		}
	}

	return nil
}

func (m *OBStandingOrder6Detail) contextValidateCreditorAccount(ctx context.Context, formats strfmt.Registry) error {

	if m.CreditorAccount != nil {
		if err := m.CreditorAccount.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("CreditorAccount")
			}
			return err
		}
	}

	return nil
}

func (m *OBStandingOrder6Detail) contextValidateCreditorAgent(ctx context.Context, formats strfmt.Registry) error {

	if m.CreditorAgent != nil {
		if err := m.CreditorAgent.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("CreditorAgent")
			}
			return err
		}
	}

	return nil
}

func (m *OBStandingOrder6Detail) contextValidateFinalPaymentAmount(ctx context.Context, formats strfmt.Registry) error {

	if m.FinalPaymentAmount != nil {
		if err := m.FinalPaymentAmount.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("FinalPaymentAmount")
			}
			return err
		}
	}

	return nil
}

func (m *OBStandingOrder6Detail) contextValidateFinalPaymentDateTime(ctx context.Context, formats strfmt.Registry) error {

	if err := m.FinalPaymentDateTime.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("FinalPaymentDateTime")
		}
		return err
	}

	return nil
}

func (m *OBStandingOrder6Detail) contextValidateFirstPaymentAmount(ctx context.Context, formats strfmt.Registry) error {

	if m.FirstPaymentAmount != nil {
		if err := m.FirstPaymentAmount.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("FirstPaymentAmount")
			}
			return err
		}
	}

	return nil
}

func (m *OBStandingOrder6Detail) contextValidateFirstPaymentDateTime(ctx context.Context, formats strfmt.Registry) error {

	if err := m.FirstPaymentDateTime.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("FirstPaymentDateTime")
		}
		return err
	}

	return nil
}

func (m *OBStandingOrder6Detail) contextValidateFrequency(ctx context.Context, formats strfmt.Registry) error {

	if m.Frequency != nil {
		if err := m.Frequency.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("Frequency")
			}
			return err
		}
	}

	return nil
}

func (m *OBStandingOrder6Detail) contextValidateLastPaymentAmount(ctx context.Context, formats strfmt.Registry) error {

	if m.LastPaymentAmount != nil {
		if err := m.LastPaymentAmount.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("LastPaymentAmount")
			}
			return err
		}
	}

	return nil
}

func (m *OBStandingOrder6Detail) contextValidateLastPaymentDateTime(ctx context.Context, formats strfmt.Registry) error {

	if err := m.LastPaymentDateTime.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("LastPaymentDateTime")
		}
		return err
	}

	return nil
}

func (m *OBStandingOrder6Detail) contextValidateNextPaymentAmount(ctx context.Context, formats strfmt.Registry) error {

	if m.NextPaymentAmount != nil {
		if err := m.NextPaymentAmount.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("NextPaymentAmount")
			}
			return err
		}
	}

	return nil
}

func (m *OBStandingOrder6Detail) contextValidateNextPaymentDateTime(ctx context.Context, formats strfmt.Registry) error {

	if err := m.NextPaymentDateTime.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("NextPaymentDateTime")
		}
		return err
	}

	return nil
}

func (m *OBStandingOrder6Detail) contextValidateNumberOfPayments(ctx context.Context, formats strfmt.Registry) error {

	if err := m.NumberOfPayments.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("NumberOfPayments")
		}
		return err
	}

	return nil
}

func (m *OBStandingOrder6Detail) contextValidateReference(ctx context.Context, formats strfmt.Registry) error {

	if err := m.Reference.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("Reference")
		}
		return err
	}

	return nil
}

func (m *OBStandingOrder6Detail) contextValidateStandingOrderID(ctx context.Context, formats strfmt.Registry) error {

	if err := m.StandingOrderID.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("StandingOrderId")
		}
		return err
	}

	return nil
}

func (m *OBStandingOrder6Detail) contextValidateStandingOrderStatusCode(ctx context.Context, formats strfmt.Registry) error {

	if err := m.StandingOrderStatusCode.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("StandingOrderStatusCode")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *OBStandingOrder6Detail) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *OBStandingOrder6Detail) UnmarshalBinary(b []byte) error {
	var res OBStandingOrder6Detail
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
