// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// BankingLoanAccountV2 BankingLoanAccountV2
//
// swagger:model BankingLoanAccountV2
type BankingLoanAccountV2 struct {

	// Date that the loan is due to be repaid in full
	LoanEndDate string `json:"loanEndDate,omitempty"`

	// Maximum amount of funds that can be redrawn. If not present redraw is not available even if the feature exists for the account
	MaxRedraw string `json:"maxRedraw,omitempty"`

	// If absent assumes AUD
	MaxRedrawCurrency string `json:"maxRedrawCurrency,omitempty"`

	// Minimum amount of next instalment
	MinInstalmentAmount string `json:"minInstalmentAmount,omitempty"`

	// If absent assumes AUD
	MinInstalmentCurrency string `json:"minInstalmentCurrency,omitempty"`

	// Minimum redraw amount
	MinRedraw string `json:"minRedraw,omitempty"`

	// If absent assumes AUD
	MinRedrawCurrency string `json:"minRedrawCurrency,omitempty"`

	// Next date that an instalment is required
	NextInstalmentDate string `json:"nextInstalmentDate,omitempty"`

	// Set to true if one or more offset accounts are configured for this loan account
	OffsetAccountEnabled bool `json:"offsetAccountEnabled,omitempty"`

	// The accountIDs of the configured offset accounts attached to this loan. Only offset accounts that can be accessed under the current authorisation should be included. It is expected behaviour that offsetAccountEnabled is set to true but the offsetAccountIds field is absent or empty. This represents a situation where an offset account exists but details can not be accessed under the current authorisation
	OffsetAccountIds []string `json:"offsetAccountIds"`

	// Optional original loan value
	OriginalLoanAmount string `json:"originalLoanAmount,omitempty"`

	// If absent assumes AUD
	OriginalLoanCurrency string `json:"originalLoanCurrency,omitempty"`

	// Optional original start date for the loan
	OriginalStartDate string `json:"originalStartDate,omitempty"`

	// The expected or required repayment frequency. Formatted according to [ISO 8601 Durations](https://en.wikipedia.org/wiki/ISO_8601#Durations) (excludes recurrence syntax)
	RepaymentFrequency string `json:"repaymentFrequency,omitempty"`

	// repayment type
	RepaymentType RepaymentType1 `json:"repaymentType,omitempty"`
}

// Validate validates this banking loan account v2
func (m *BankingLoanAccountV2) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateRepaymentType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *BankingLoanAccountV2) validateRepaymentType(formats strfmt.Registry) error {
	if swag.IsZero(m.RepaymentType) { // not required
		return nil
	}

	if err := m.RepaymentType.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("repaymentType")
		}
		return err
	}

	return nil
}

// ContextValidate validate this banking loan account v2 based on the context it is used
func (m *BankingLoanAccountV2) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateRepaymentType(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *BankingLoanAccountV2) contextValidateRepaymentType(ctx context.Context, formats strfmt.Registry) error {

	if err := m.RepaymentType.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("repaymentType")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *BankingLoanAccountV2) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *BankingLoanAccountV2) UnmarshalBinary(b []byte) error {
	var res BankingLoanAccountV2
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
