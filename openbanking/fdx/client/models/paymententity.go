// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// Paymententity Paymententity
//
// Represents a payment
//
// swagger:model Paymententity
type Paymententity struct {

	// Amount for the payment. Must be positive
	// Required: true
	// Minimum: 0
	Amount *float64 `json:"amount"`

	// When the payment was cancelled
	// Example: 2021-07-15T14:46:41.375Z
	// Format: date-time
	CancelledTimestamp strfmt.DateTime `json:"cancelledTimestamp,omitempty"`

	// Date that the funds are scheduled to be delivered
	// Example: 2021-07-15T00:00:00.000Z
	// Required: true
	// Format: date
	DueDate *strfmt.Date `json:"dueDate"`

	// When the payment failed. Includes when the payment was determined to lack sufficient funds
	// Example: 2021-07-15T14:46:41.375Z
	// Format: date-time
	FailedTimestamp strfmt.DateTime `json:"failedTimestamp,omitempty"`

	// ID of the account used to source funds for payment
	// Required: true
	// Max Length: 256
	FromAccountID *string `json:"fromAccountId"`

	// Links to related payment entities
	Links []*HATEOASLink `json:"links"`

	// User's account identifier with the payee
	MerchantAccountID string `json:"merchantAccountId,omitempty"`

	// Uniquely identifies a payment. Used within the API to reference a payment
	// Required: true
	// Max Length: 256
	PaymentID *string `json:"paymentId"`

	// When the payment was processed
	// Example: 2021-07-15T14:46:41.375Z
	// Format: date-time
	ProcessedTimestamp strfmt.DateTime `json:"processedTimestamp,omitempty"`

	// The recurring payment that spawned this payment. Null if payment is not associated with a recurring payment
	// Max Length: 256
	RecurringPaymentID string `json:"recurringPaymentId,omitempty"`

	// When the payment was scheduled
	// Example: 2021-07-15T14:46:41.375Z
	// Format: date-time
	ScheduledTimestamp strfmt.DateTime `json:"scheduledTimestamp,omitempty"`

	// When the payment execution started
	// Example: 2021-07-15T14:46:41.375Z
	// Format: date-time
	StartedProcessingTimestamp strfmt.DateTime `json:"startedProcessingTimestamp,omitempty"`

	// status
	// Required: true
	Status *PaymentStatus4 `json:"status"`

	// ID of the payee to receive funds for the payment
	// Required: true
	// Max Length: 256
	ToPayeeID *string `json:"toPayeeId"`
}

// Validate validates this paymententity
func (m *Paymententity) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAmount(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCancelledTimestamp(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDueDate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFailedTimestamp(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFromAccountID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLinks(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePaymentID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateProcessedTimestamp(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRecurringPaymentID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateScheduledTimestamp(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStartedProcessingTimestamp(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStatus(formats); err != nil {
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

func (m *Paymententity) validateAmount(formats strfmt.Registry) error {

	if err := validate.Required("amount", "body", m.Amount); err != nil {
		return err
	}

	if err := validate.Minimum("amount", "body", *m.Amount, 0, false); err != nil {
		return err
	}

	return nil
}

func (m *Paymententity) validateCancelledTimestamp(formats strfmt.Registry) error {
	if swag.IsZero(m.CancelledTimestamp) { // not required
		return nil
	}

	if err := validate.FormatOf("cancelledTimestamp", "body", "date-time", m.CancelledTimestamp.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *Paymententity) validateDueDate(formats strfmt.Registry) error {

	if err := validate.Required("dueDate", "body", m.DueDate); err != nil {
		return err
	}

	if err := validate.FormatOf("dueDate", "body", "date", m.DueDate.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *Paymententity) validateFailedTimestamp(formats strfmt.Registry) error {
	if swag.IsZero(m.FailedTimestamp) { // not required
		return nil
	}

	if err := validate.FormatOf("failedTimestamp", "body", "date-time", m.FailedTimestamp.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *Paymententity) validateFromAccountID(formats strfmt.Registry) error {

	if err := validate.Required("fromAccountId", "body", m.FromAccountID); err != nil {
		return err
	}

	if err := validate.MaxLength("fromAccountId", "body", *m.FromAccountID, 256); err != nil {
		return err
	}

	return nil
}

func (m *Paymententity) validateLinks(formats strfmt.Registry) error {
	if swag.IsZero(m.Links) { // not required
		return nil
	}

	for i := 0; i < len(m.Links); i++ {
		if swag.IsZero(m.Links[i]) { // not required
			continue
		}

		if m.Links[i] != nil {
			if err := m.Links[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("links" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("links" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *Paymententity) validatePaymentID(formats strfmt.Registry) error {

	if err := validate.Required("paymentId", "body", m.PaymentID); err != nil {
		return err
	}

	if err := validate.MaxLength("paymentId", "body", *m.PaymentID, 256); err != nil {
		return err
	}

	return nil
}

func (m *Paymententity) validateProcessedTimestamp(formats strfmt.Registry) error {
	if swag.IsZero(m.ProcessedTimestamp) { // not required
		return nil
	}

	if err := validate.FormatOf("processedTimestamp", "body", "date-time", m.ProcessedTimestamp.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *Paymententity) validateRecurringPaymentID(formats strfmt.Registry) error {
	if swag.IsZero(m.RecurringPaymentID) { // not required
		return nil
	}

	if err := validate.MaxLength("recurringPaymentId", "body", m.RecurringPaymentID, 256); err != nil {
		return err
	}

	return nil
}

func (m *Paymententity) validateScheduledTimestamp(formats strfmt.Registry) error {
	if swag.IsZero(m.ScheduledTimestamp) { // not required
		return nil
	}

	if err := validate.FormatOf("scheduledTimestamp", "body", "date-time", m.ScheduledTimestamp.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *Paymententity) validateStartedProcessingTimestamp(formats strfmt.Registry) error {
	if swag.IsZero(m.StartedProcessingTimestamp) { // not required
		return nil
	}

	if err := validate.FormatOf("startedProcessingTimestamp", "body", "date-time", m.StartedProcessingTimestamp.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *Paymententity) validateStatus(formats strfmt.Registry) error {

	if err := validate.Required("status", "body", m.Status); err != nil {
		return err
	}

	if err := validate.Required("status", "body", m.Status); err != nil {
		return err
	}

	if m.Status != nil {
		if err := m.Status.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("status")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("status")
			}
			return err
		}
	}

	return nil
}

func (m *Paymententity) validateToPayeeID(formats strfmt.Registry) error {

	if err := validate.Required("toPayeeId", "body", m.ToPayeeID); err != nil {
		return err
	}

	if err := validate.MaxLength("toPayeeId", "body", *m.ToPayeeID, 256); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this paymententity based on the context it is used
func (m *Paymententity) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateLinks(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateStatus(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Paymententity) contextValidateLinks(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Links); i++ {

		if m.Links[i] != nil {
			if err := m.Links[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("links" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("links" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *Paymententity) contextValidateStatus(ctx context.Context, formats strfmt.Registry) error {

	if m.Status != nil {
		if err := m.Status.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("status")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("status")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Paymententity) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Paymententity) UnmarshalBinary(b []byte) error {
	var res Paymententity
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
