// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/validate"
)

// LoanPaymentFrequency LoanPaymentFrequency
//
// The frequency of payments on a loan
//
// swagger:model LoanPaymentFrequency
type LoanPaymentFrequency string

func NewLoanPaymentFrequency(value LoanPaymentFrequency) *LoanPaymentFrequency {
	v := value
	return &v
}

const (

	// LoanPaymentFrequencyANNUALLY captures enum value "ANNUALLY"
	LoanPaymentFrequencyANNUALLY LoanPaymentFrequency = "ANNUALLY"

	// LoanPaymentFrequencyBIMONTHLY captures enum value "BIMONTHLY"
	LoanPaymentFrequencyBIMONTHLY LoanPaymentFrequency = "BIMONTHLY"

	// LoanPaymentFrequencyBIWEEKLY captures enum value "BIWEEKLY"
	LoanPaymentFrequencyBIWEEKLY LoanPaymentFrequency = "BIWEEKLY"

	// LoanPaymentFrequencyFOURWEEKS captures enum value "FOURWEEKS"
	LoanPaymentFrequencyFOURWEEKS LoanPaymentFrequency = "FOURWEEKS"

	// LoanPaymentFrequencyMONTHLY captures enum value "MONTHLY"
	LoanPaymentFrequencyMONTHLY LoanPaymentFrequency = "MONTHLY"

	// LoanPaymentFrequencyOTHER captures enum value "OTHER"
	LoanPaymentFrequencyOTHER LoanPaymentFrequency = "OTHER"

	// LoanPaymentFrequencyQUARTERLY captures enum value "QUARTERLY"
	LoanPaymentFrequencyQUARTERLY LoanPaymentFrequency = "QUARTERLY"

	// LoanPaymentFrequencySEMIANNUALLY captures enum value "SEMIANNUALLY"
	LoanPaymentFrequencySEMIANNUALLY LoanPaymentFrequency = "SEMIANNUALLY"

	// LoanPaymentFrequencyTWICEMONTHLY captures enum value "TWICEMONTHLY"
	LoanPaymentFrequencyTWICEMONTHLY LoanPaymentFrequency = "TWICEMONTHLY"

	// LoanPaymentFrequencyWEEKLY captures enum value "WEEKLY"
	LoanPaymentFrequencyWEEKLY LoanPaymentFrequency = "WEEKLY"
)

// for schema
var loanPaymentFrequencyEnum []interface{}

func init() {
	var res []LoanPaymentFrequency
	if err := json.Unmarshal([]byte(`["ANNUALLY","BIMONTHLY","BIWEEKLY","FOURWEEKS","MONTHLY","OTHER","QUARTERLY","SEMIANNUALLY","TWICEMONTHLY","WEEKLY"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		loanPaymentFrequencyEnum = append(loanPaymentFrequencyEnum, v)
	}
}

func (m LoanPaymentFrequency) validateLoanPaymentFrequencyEnum(path, location string, value LoanPaymentFrequency) error {
	if err := validate.EnumCase(path, location, value, loanPaymentFrequencyEnum, true); err != nil {
		return err
	}
	return nil
}

// Validate validates this loan payment frequency
func (m LoanPaymentFrequency) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateLoanPaymentFrequencyEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validates this loan payment frequency based on context it is used
func (m LoanPaymentFrequency) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}