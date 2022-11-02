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

// BankingProductLendingRateV2 BankingProductLendingRateV2
//
// swagger:model BankingProductLendingRateV2
type BankingProductLendingRateV2 struct {

	// Display text providing more information on the rate.
	AdditionalInfo string `json:"additionalInfo,omitempty"`

	// Link to a web page with more information on this rate
	AdditionalInfoURI string `json:"additionalInfoUri,omitempty"`

	// Generic field containing additional information relevant to the [lendingRateType](#tocSproductlendingratetypedoc) specified. Whether mandatory or not is dependent on the value of [lendingRateType](#tocSproductlendingratetypedoc)
	AdditionalValue string `json:"additionalValue,omitempty"`

	// The period after which the calculated amount(s) (see calculationFrequency) are 'applied' (i.e. debited or credited) to the account. Formatted according to [ISO 8601 Durations](https://en.wikipedia.org/wiki/ISO_8601#Durations) (excludes recurrence syntax)
	ApplicationFrequency string `json:"applicationFrequency,omitempty"`

	// The period after which the rate is applied to the balance to calculate the amount due for the period. Calculation of the amount is often daily (as balances may change) but accumulated until the total amount is 'applied' to the account (see applicationFrequency). Formatted according to [ISO 8601 Durations](https://en.wikipedia.org/wiki/ISO_8601#Durations) (excludes recurrence syntax)
	CalculationFrequency string `json:"calculationFrequency,omitempty"`

	// A comparison rate equivalent for this rate
	ComparisonRate string `json:"comparisonRate,omitempty"`

	// interest payment due
	InterestPaymentDue InterestPaymentDue `json:"interestPaymentDue,omitempty"`

	// lending rate type
	// Required: true
	LendingRateType *LendingRateType `json:"lendingRateType"`

	// loan purpose
	LoanPurpose LoanPurpose `json:"loanPurpose,omitempty"`

	// The rate to be applied
	// Required: true
	Rate *string `json:"rate"`

	// repayment type
	RepaymentType RepaymentType `json:"repaymentType,omitempty"`

	// Rate tiers applicable for this rate
	Tiers []*BankingProductRateTierV3 `json:"tiers"`
}

// Validate validates this banking product lending rate v2
func (m *BankingProductLendingRateV2) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateInterestPaymentDue(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLendingRateType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLoanPurpose(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRepaymentType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTiers(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *BankingProductLendingRateV2) validateInterestPaymentDue(formats strfmt.Registry) error {
	if swag.IsZero(m.InterestPaymentDue) { // not required
		return nil
	}

	if err := m.InterestPaymentDue.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("interestPaymentDue")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("interestPaymentDue")
		}
		return err
	}

	return nil
}

func (m *BankingProductLendingRateV2) validateLendingRateType(formats strfmt.Registry) error {

	if err := validate.Required("lendingRateType", "body", m.LendingRateType); err != nil {
		return err
	}

	if err := validate.Required("lendingRateType", "body", m.LendingRateType); err != nil {
		return err
	}

	if m.LendingRateType != nil {
		if err := m.LendingRateType.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("lendingRateType")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("lendingRateType")
			}
			return err
		}
	}

	return nil
}

func (m *BankingProductLendingRateV2) validateLoanPurpose(formats strfmt.Registry) error {
	if swag.IsZero(m.LoanPurpose) { // not required
		return nil
	}

	if err := m.LoanPurpose.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("loanPurpose")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("loanPurpose")
		}
		return err
	}

	return nil
}

func (m *BankingProductLendingRateV2) validateRate(formats strfmt.Registry) error {

	if err := validate.Required("rate", "body", m.Rate); err != nil {
		return err
	}

	return nil
}

func (m *BankingProductLendingRateV2) validateRepaymentType(formats strfmt.Registry) error {
	if swag.IsZero(m.RepaymentType) { // not required
		return nil
	}

	if err := m.RepaymentType.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("repaymentType")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("repaymentType")
		}
		return err
	}

	return nil
}

func (m *BankingProductLendingRateV2) validateTiers(formats strfmt.Registry) error {
	if swag.IsZero(m.Tiers) { // not required
		return nil
	}

	for i := 0; i < len(m.Tiers); i++ {
		if swag.IsZero(m.Tiers[i]) { // not required
			continue
		}

		if m.Tiers[i] != nil {
			if err := m.Tiers[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("tiers" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("tiers" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this banking product lending rate v2 based on the context it is used
func (m *BankingProductLendingRateV2) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateInterestPaymentDue(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateLendingRateType(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateLoanPurpose(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateRepaymentType(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateTiers(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *BankingProductLendingRateV2) contextValidateInterestPaymentDue(ctx context.Context, formats strfmt.Registry) error {

	if err := m.InterestPaymentDue.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("interestPaymentDue")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("interestPaymentDue")
		}
		return err
	}

	return nil
}

func (m *BankingProductLendingRateV2) contextValidateLendingRateType(ctx context.Context, formats strfmt.Registry) error {

	if m.LendingRateType != nil {
		if err := m.LendingRateType.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("lendingRateType")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("lendingRateType")
			}
			return err
		}
	}

	return nil
}

func (m *BankingProductLendingRateV2) contextValidateLoanPurpose(ctx context.Context, formats strfmt.Registry) error {

	if err := m.LoanPurpose.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("loanPurpose")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("loanPurpose")
		}
		return err
	}

	return nil
}

func (m *BankingProductLendingRateV2) contextValidateRepaymentType(ctx context.Context, formats strfmt.Registry) error {

	if err := m.RepaymentType.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("repaymentType")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("repaymentType")
		}
		return err
	}

	return nil
}

func (m *BankingProductLendingRateV2) contextValidateTiers(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Tiers); i++ {

		if m.Tiers[i] != nil {
			if err := m.Tiers[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("tiers" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("tiers" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *BankingProductLendingRateV2) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *BankingProductLendingRateV2) UnmarshalBinary(b []byte) error {
	var res BankingProductLendingRateV2
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}