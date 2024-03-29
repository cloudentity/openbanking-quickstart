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

// TransferforCreateentity1 TransferforCreateentity1
//
// Data of the transfer request
//
// swagger:model TransferforCreateentity1
type TransferforCreateentity1 struct {

	// Positive amount of money to be transferred
	Amount float64 `json:"amount,omitempty"`

	// Long-term persistent identity of the source account
	// Max Length: 256
	FromAccountID string `json:"fromAccountId,omitempty"`

	// User-entered reason for transfer
	// Max Length: 255
	Memo string `json:"memo,omitempty"`

	// payment details
	PaymentDetails *PaymentDetailsentity4 `json:"paymentDetails,omitempty"`

	// Long-term persistent identity of the destination account
	// Max Length: 256
	ToAccountID string `json:"toAccountId,omitempty"`

	// Client generated, long-term persistent identity of the transfer action. This ID should be maintained and returned by institution
	// Max Length: 256
	TransferID string `json:"transferId,omitempty"`
}

// Validate validates this transferfor createentity1
func (m *TransferforCreateentity1) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateFromAccountID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMemo(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePaymentDetails(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateToAccountID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTransferID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TransferforCreateentity1) validateFromAccountID(formats strfmt.Registry) error {
	if swag.IsZero(m.FromAccountID) { // not required
		return nil
	}

	if err := validate.MaxLength("fromAccountId", "body", m.FromAccountID, 256); err != nil {
		return err
	}

	return nil
}

func (m *TransferforCreateentity1) validateMemo(formats strfmt.Registry) error {
	if swag.IsZero(m.Memo) { // not required
		return nil
	}

	if err := validate.MaxLength("memo", "body", m.Memo, 255); err != nil {
		return err
	}

	return nil
}

func (m *TransferforCreateentity1) validatePaymentDetails(formats strfmt.Registry) error {
	if swag.IsZero(m.PaymentDetails) { // not required
		return nil
	}

	if m.PaymentDetails != nil {
		if err := m.PaymentDetails.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("paymentDetails")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("paymentDetails")
			}
			return err
		}
	}

	return nil
}

func (m *TransferforCreateentity1) validateToAccountID(formats strfmt.Registry) error {
	if swag.IsZero(m.ToAccountID) { // not required
		return nil
	}

	if err := validate.MaxLength("toAccountId", "body", m.ToAccountID, 256); err != nil {
		return err
	}

	return nil
}

func (m *TransferforCreateentity1) validateTransferID(formats strfmt.Registry) error {
	if swag.IsZero(m.TransferID) { // not required
		return nil
	}

	if err := validate.MaxLength("transferId", "body", m.TransferID, 256); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this transferfor createentity1 based on the context it is used
func (m *TransferforCreateentity1) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidatePaymentDetails(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TransferforCreateentity1) contextValidatePaymentDetails(ctx context.Context, formats strfmt.Registry) error {

	if m.PaymentDetails != nil {
		if err := m.PaymentDetails.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("paymentDetails")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("paymentDetails")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *TransferforCreateentity1) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TransferforCreateentity1) UnmarshalBinary(b []byte) error {
	var res TransferforCreateentity1
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
