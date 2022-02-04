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

// BankingCreditCardAccount BankingCreditCardAccount
//
// swagger:model BankingCreditCardAccount
type BankingCreditCardAccount struct {

	// The minimum payment amount due for the next card payment
	// Required: true
	MinPaymentAmount *string `json:"minPaymentAmount"`

	// If absent assumes AUD
	PaymentCurrency string `json:"paymentCurrency,omitempty"`

	// The amount due for the next card payment
	// Required: true
	PaymentDueAmount *string `json:"paymentDueAmount"`

	// Date that the next payment for the card is due
	// Required: true
	PaymentDueDate *string `json:"paymentDueDate"`
}

// Validate validates this banking credit card account
func (m *BankingCreditCardAccount) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateMinPaymentAmount(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePaymentDueAmount(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePaymentDueDate(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *BankingCreditCardAccount) validateMinPaymentAmount(formats strfmt.Registry) error {

	if err := validate.Required("minPaymentAmount", "body", m.MinPaymentAmount); err != nil {
		return err
	}

	return nil
}

func (m *BankingCreditCardAccount) validatePaymentDueAmount(formats strfmt.Registry) error {

	if err := validate.Required("paymentDueAmount", "body", m.PaymentDueAmount); err != nil {
		return err
	}

	return nil
}

func (m *BankingCreditCardAccount) validatePaymentDueDate(formats strfmt.Registry) error {

	if err := validate.Required("paymentDueDate", "body", m.PaymentDueDate); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this banking credit card account based on context it is used
func (m *BankingCreditCardAccount) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *BankingCreditCardAccount) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *BankingCreditCardAccount) UnmarshalBinary(b []byte) error {
	var res BankingCreditCardAccount
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}