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

// BankingScheduledPaymentRecurrenceOnceOff BankingScheduledPaymentRecurrenceOnceOff
//
// Indicates that the payment is a once off payment on a specific future date. Mandatory if recurrenceUType is set to onceOff
//
// swagger:model BankingScheduledPaymentRecurrenceOnceOff
type BankingScheduledPaymentRecurrenceOnceOff struct {

	// The scheduled date for the once off payment
	// Required: true
	PaymentDate *string `json:"paymentDate"`
}

// Validate validates this banking scheduled payment recurrence once off
func (m *BankingScheduledPaymentRecurrenceOnceOff) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validatePaymentDate(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *BankingScheduledPaymentRecurrenceOnceOff) validatePaymentDate(formats strfmt.Registry) error {

	if err := validate.Required("paymentDate", "body", m.PaymentDate); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this banking scheduled payment recurrence once off based on context it is used
func (m *BankingScheduledPaymentRecurrenceOnceOff) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *BankingScheduledPaymentRecurrenceOnceOff) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *BankingScheduledPaymentRecurrenceOnceOff) UnmarshalBinary(b []byte) error {
	var res BankingScheduledPaymentRecurrenceOnceOff
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}