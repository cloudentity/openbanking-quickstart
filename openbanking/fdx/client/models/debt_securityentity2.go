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

// DebtSecurityentity2 DebtSecurityentity2
//
// A debt security
//
// swagger:model DebtSecurityentity2
type DebtSecurityentity2 struct {

	// Bond maturity date
	// Example: 2021-07-15T00:00:00.000Z
	// Format: date
	BondMaturityDate strfmt.Date `json:"bondMaturityDate,omitempty"`

	// Next call date
	// Example: 2021-07-15T00:00:00.000Z
	// Format: date
	CallDate strfmt.Date `json:"callDate,omitempty"`

	// Bond call price
	CallPrice float64 `json:"callPrice,omitempty"`

	// call type
	CallType CallType2 `json:"callType,omitempty"`

	// Maturity date for next coupon
	// Example: 2021-07-15T00:00:00.000Z
	// Format: date
	CouponDate strfmt.Date `json:"couponDate,omitempty"`

	// coupon mature frequency
	CouponMatureFrequency CouponMatureFrequency2 `json:"couponMatureFrequency,omitempty"`

	// Bond coupon rate for next closest call date
	CouponRate float64 `json:"couponRate,omitempty"`

	// debt class
	DebtClass DebtClass2 `json:"debtClass,omitempty"`

	// debt type
	DebtType DebtType2 `json:"debtType,omitempty"`

	// Par value amount
	ParValue float64 `json:"parValue,omitempty"`

	// Yield to next call
	YieldToCall float64 `json:"yieldToCall,omitempty"`

	// Yield to maturity
	YieldToMaturity float64 `json:"yieldToMaturity,omitempty"`
}

// Validate validates this debt securityentity2
func (m *DebtSecurityentity2) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateBondMaturityDate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCallDate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCallType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCouponDate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCouponMatureFrequency(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDebtClass(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDebtType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DebtSecurityentity2) validateBondMaturityDate(formats strfmt.Registry) error {
	if swag.IsZero(m.BondMaturityDate) { // not required
		return nil
	}

	if err := validate.FormatOf("bondMaturityDate", "body", "date", m.BondMaturityDate.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *DebtSecurityentity2) validateCallDate(formats strfmt.Registry) error {
	if swag.IsZero(m.CallDate) { // not required
		return nil
	}

	if err := validate.FormatOf("callDate", "body", "date", m.CallDate.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *DebtSecurityentity2) validateCallType(formats strfmt.Registry) error {
	if swag.IsZero(m.CallType) { // not required
		return nil
	}

	if err := m.CallType.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("callType")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("callType")
		}
		return err
	}

	return nil
}

func (m *DebtSecurityentity2) validateCouponDate(formats strfmt.Registry) error {
	if swag.IsZero(m.CouponDate) { // not required
		return nil
	}

	if err := validate.FormatOf("couponDate", "body", "date", m.CouponDate.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *DebtSecurityentity2) validateCouponMatureFrequency(formats strfmt.Registry) error {
	if swag.IsZero(m.CouponMatureFrequency) { // not required
		return nil
	}

	if err := m.CouponMatureFrequency.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("couponMatureFrequency")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("couponMatureFrequency")
		}
		return err
	}

	return nil
}

func (m *DebtSecurityentity2) validateDebtClass(formats strfmt.Registry) error {
	if swag.IsZero(m.DebtClass) { // not required
		return nil
	}

	if err := m.DebtClass.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("debtClass")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("debtClass")
		}
		return err
	}

	return nil
}

func (m *DebtSecurityentity2) validateDebtType(formats strfmt.Registry) error {
	if swag.IsZero(m.DebtType) { // not required
		return nil
	}

	if err := m.DebtType.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("debtType")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("debtType")
		}
		return err
	}

	return nil
}

// ContextValidate validate this debt securityentity2 based on the context it is used
func (m *DebtSecurityentity2) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateCallType(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateCouponMatureFrequency(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateDebtClass(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateDebtType(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DebtSecurityentity2) contextValidateCallType(ctx context.Context, formats strfmt.Registry) error {

	if err := m.CallType.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("callType")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("callType")
		}
		return err
	}

	return nil
}

func (m *DebtSecurityentity2) contextValidateCouponMatureFrequency(ctx context.Context, formats strfmt.Registry) error {

	if err := m.CouponMatureFrequency.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("couponMatureFrequency")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("couponMatureFrequency")
		}
		return err
	}

	return nil
}

func (m *DebtSecurityentity2) contextValidateDebtClass(ctx context.Context, formats strfmt.Registry) error {

	if err := m.DebtClass.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("debtClass")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("debtClass")
		}
		return err
	}

	return nil
}

func (m *DebtSecurityentity2) contextValidateDebtType(ctx context.Context, formats strfmt.Registry) error {

	if err := m.DebtType.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("debtType")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("debtType")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *DebtSecurityentity2) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DebtSecurityentity2) UnmarshalBinary(b []byte) error {
	var res DebtSecurityentity2
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
