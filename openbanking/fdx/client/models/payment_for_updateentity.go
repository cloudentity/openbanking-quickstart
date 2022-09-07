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

// PaymentForUpdateentity PaymentForUpdateentity
//
// Payment entity used for creation and update of a payment
//
// swagger:model PaymentForUpdateentity
type PaymentForUpdateentity struct {

	// Amount for the payment. Must be positive
	// Required: true
	// Minimum: 0
	Amount *float64 `json:"amount"`

	// Date that the funds are scheduled to be delivered
	// Example: 2021-07-15T00:00:00.000Z
	// Required: true
	// Format: date
	DueDate *strfmt.Date `json:"dueDate"`

	// ID of the account used to source funds for payment
	// Required: true
	// Max Length: 256
	FromAccountID *string `json:"fromAccountId"`

	// User's account identifier with the payee
	MerchantAccountID string `json:"merchantAccountId,omitempty"`

	// ID of the payee to receive funds for the payment
	// Required: true
	// Max Length: 256
	ToPayeeID *string `json:"toPayeeId"`
}

// Validate validates this payment for updateentity
func (m *PaymentForUpdateentity) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAmount(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDueDate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFromAccountID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateToPayeeID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PaymentForUpdateentity) validateAmount(formats strfmt.Registry) error {

	if err := validate.Required("amount", "body", m.Amount); err != nil {
		return err
	}

	if err := validate.Minimum("amount", "body", *m.Amount, 0, false); err != nil {
		return err
	}

	return nil
}

func (m *PaymentForUpdateentity) validateDueDate(formats strfmt.Registry) error {

	if err := validate.Required("dueDate", "body", m.DueDate); err != nil {
		return err
	}

	if err := validate.FormatOf("dueDate", "body", "date", m.DueDate.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *PaymentForUpdateentity) validateFromAccountID(formats strfmt.Registry) error {

	if err := validate.Required("fromAccountId", "body", m.FromAccountID); err != nil {
		return err
	}

	if err := validate.MaxLength("fromAccountId", "body", *m.FromAccountID, 256); err != nil {
		return err
	}

	return nil
}

func (m *PaymentForUpdateentity) validateToPayeeID(formats strfmt.Registry) error {

	if err := validate.Required("toPayeeId", "body", m.ToPayeeID); err != nil {
		return err
	}

	if err := validate.MaxLength("toPayeeId", "body", *m.ToPayeeID, 256); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this payment for updateentity based on context it is used
func (m *PaymentForUpdateentity) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *PaymentForUpdateentity) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PaymentForUpdateentity) UnmarshalBinary(b []byte) error {
	var res PaymentForUpdateentity
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
