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

// BankingScheduledPaymentRecurrenceIntervalSchedule BankingScheduledPaymentRecurrenceIntervalSchedule
//
// Indicates that the schedule of payments is defined by a series of intervals. Mandatory if recurrenceUType is set to intervalSchedule
//
// swagger:model BankingScheduledPaymentRecurrenceIntervalSchedule
type BankingScheduledPaymentRecurrenceIntervalSchedule struct {

	// The limit date after which no more payments should be made using this schedule. If both finalPaymentDate and paymentsRemaining are present then payments will stop according to the most constraining value. If neither field is present the payments will continue indefinitely
	FinalPaymentDate string `json:"finalPaymentDate,omitempty"`

	// An array of interval objects defining the payment schedule.  Each entry in the array is additive, in that it adds payments to the overall payment schedule.  If multiple intervals result in a payment on the same day then only one payment will be made. Must have at least one entry
	// Required: true
	Intervals []*BankingScheduledPaymentInterval `json:"intervals"`

	// non business day treatment
	NonBusinessDayTreatment NonBusinessDayTreatment `json:"nonBusinessDayTreatment,omitempty"`

	// Indicates the number of payments remaining in the schedule. If both finalPaymentDate and paymentsRemaining are present then payments will stop according to the most constraining value, If neither field is present the payments will continue indefinitely
	// Example: 1
	PaymentsRemaining int32 `json:"paymentsRemaining,omitempty"`
}

// Validate validates this banking scheduled payment recurrence interval schedule
func (m *BankingScheduledPaymentRecurrenceIntervalSchedule) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateIntervals(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNonBusinessDayTreatment(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *BankingScheduledPaymentRecurrenceIntervalSchedule) validateIntervals(formats strfmt.Registry) error {

	if err := validate.Required("intervals", "body", m.Intervals); err != nil {
		return err
	}

	for i := 0; i < len(m.Intervals); i++ {
		if swag.IsZero(m.Intervals[i]) { // not required
			continue
		}

		if m.Intervals[i] != nil {
			if err := m.Intervals[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("intervals" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("intervals" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *BankingScheduledPaymentRecurrenceIntervalSchedule) validateNonBusinessDayTreatment(formats strfmt.Registry) error {
	if swag.IsZero(m.NonBusinessDayTreatment) { // not required
		return nil
	}

	if err := m.NonBusinessDayTreatment.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("nonBusinessDayTreatment")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("nonBusinessDayTreatment")
		}
		return err
	}

	return nil
}

// ContextValidate validate this banking scheduled payment recurrence interval schedule based on the context it is used
func (m *BankingScheduledPaymentRecurrenceIntervalSchedule) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateIntervals(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateNonBusinessDayTreatment(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *BankingScheduledPaymentRecurrenceIntervalSchedule) contextValidateIntervals(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Intervals); i++ {

		if m.Intervals[i] != nil {
			if err := m.Intervals[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("intervals" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("intervals" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *BankingScheduledPaymentRecurrenceIntervalSchedule) contextValidateNonBusinessDayTreatment(ctx context.Context, formats strfmt.Registry) error {

	if err := m.NonBusinessDayTreatment.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("nonBusinessDayTreatment")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("nonBusinessDayTreatment")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *BankingScheduledPaymentRecurrenceIntervalSchedule) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *BankingScheduledPaymentRecurrenceIntervalSchedule) UnmarshalBinary(b []byte) error {
	var res BankingScheduledPaymentRecurrenceIntervalSchedule
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
