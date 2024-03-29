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

// PensionSourceentity PensionSourceentity
//
// The source of pension funds
//
// swagger:model PensionSourceentity
type PensionSourceentity struct {

	// Benefit Amount
	Amount float64 `json:"amount,omitempty"`

	// Date benefit was calculated
	// Example: 2021-07-15T00:00:00.000Z
	// Format: date
	AsOfDate strfmt.Date `json:"asOfDate,omitempty"`

	// Name of the source
	DisplayName string `json:"displayName,omitempty"`

	// frequency
	Frequency PaymentFrequency1 `json:"frequency,omitempty"`

	// Form of payment
	PaymentOption string `json:"paymentOption,omitempty"`

	// Assumed retirement date. As of date amount is payable
	// Example: 2021-07-15T00:00:00.000Z
	// Format: date
	StartDate strfmt.Date `json:"startDate,omitempty"`
}

// Validate validates this pension sourceentity
func (m *PensionSourceentity) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAsOfDate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFrequency(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStartDate(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PensionSourceentity) validateAsOfDate(formats strfmt.Registry) error {
	if swag.IsZero(m.AsOfDate) { // not required
		return nil
	}

	if err := validate.FormatOf("asOfDate", "body", "date", m.AsOfDate.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *PensionSourceentity) validateFrequency(formats strfmt.Registry) error {
	if swag.IsZero(m.Frequency) { // not required
		return nil
	}

	if err := m.Frequency.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("frequency")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("frequency")
		}
		return err
	}

	return nil
}

func (m *PensionSourceentity) validateStartDate(formats strfmt.Registry) error {
	if swag.IsZero(m.StartDate) { // not required
		return nil
	}

	if err := validate.FormatOf("startDate", "body", "date", m.StartDate.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this pension sourceentity based on the context it is used
func (m *PensionSourceentity) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateFrequency(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PensionSourceentity) contextValidateFrequency(ctx context.Context, formats strfmt.Registry) error {

	if err := m.Frequency.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("frequency")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("frequency")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *PensionSourceentity) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PensionSourceentity) UnmarshalBinary(b []byte) error {
	var res PensionSourceentity
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
