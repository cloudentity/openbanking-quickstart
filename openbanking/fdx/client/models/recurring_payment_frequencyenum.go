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

// RecurringPaymentFrequencyenum RecurringPaymentFrequencyenum
//
// Defines how often a payment is made relative to the starting day.<br/> * `BIWEEKLY`: Every 14 days<br/> * `TWICEMONTHLY`: Every 15 days<br/> * `WEEKLY`: Every 7 days
//
// swagger:model RecurringPaymentFrequencyenum
type RecurringPaymentFrequencyenum string

func NewRecurringPaymentFrequencyenum(value RecurringPaymentFrequencyenum) *RecurringPaymentFrequencyenum {
	v := value
	return &v
}

const (

	// RecurringPaymentFrequencyenumWEEKLY captures enum value "WEEKLY"
	RecurringPaymentFrequencyenumWEEKLY RecurringPaymentFrequencyenum = "WEEKLY"

	// RecurringPaymentFrequencyenumBIWEEKLY captures enum value "BIWEEKLY"
	RecurringPaymentFrequencyenumBIWEEKLY RecurringPaymentFrequencyenum = "BIWEEKLY"

	// RecurringPaymentFrequencyenumTWICEMONTHLY captures enum value "TWICEMONTHLY"
	RecurringPaymentFrequencyenumTWICEMONTHLY RecurringPaymentFrequencyenum = "TWICEMONTHLY"

	// RecurringPaymentFrequencyenumMONTHLY captures enum value "MONTHLY"
	RecurringPaymentFrequencyenumMONTHLY RecurringPaymentFrequencyenum = "MONTHLY"

	// RecurringPaymentFrequencyenumFOURWEEKS captures enum value "FOURWEEKS"
	RecurringPaymentFrequencyenumFOURWEEKS RecurringPaymentFrequencyenum = "FOURWEEKS"

	// RecurringPaymentFrequencyenumBIMONTHLY captures enum value "BIMONTHLY"
	RecurringPaymentFrequencyenumBIMONTHLY RecurringPaymentFrequencyenum = "BIMONTHLY"

	// RecurringPaymentFrequencyenumQUARTERLY captures enum value "QUARTERLY"
	RecurringPaymentFrequencyenumQUARTERLY RecurringPaymentFrequencyenum = "QUARTERLY"

	// RecurringPaymentFrequencyenumSEMIANNUALLY captures enum value "SEMIANNUALLY"
	RecurringPaymentFrequencyenumSEMIANNUALLY RecurringPaymentFrequencyenum = "SEMIANNUALLY"

	// RecurringPaymentFrequencyenumANNUALLY captures enum value "ANNUALLY"
	RecurringPaymentFrequencyenumANNUALLY RecurringPaymentFrequencyenum = "ANNUALLY"
)

// for schema
var recurringPaymentFrequencyenumEnum []interface{}

func init() {
	var res []RecurringPaymentFrequencyenum
	if err := json.Unmarshal([]byte(`["WEEKLY","BIWEEKLY","TWICEMONTHLY","MONTHLY","FOURWEEKS","BIMONTHLY","QUARTERLY","SEMIANNUALLY","ANNUALLY"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		recurringPaymentFrequencyenumEnum = append(recurringPaymentFrequencyenumEnum, v)
	}
}

func (m RecurringPaymentFrequencyenum) validateRecurringPaymentFrequencyenumEnum(path, location string, value RecurringPaymentFrequencyenum) error {
	if err := validate.EnumCase(path, location, value, recurringPaymentFrequencyenumEnum, true); err != nil {
		return err
	}
	return nil
}

// Validate validates this recurring payment frequencyenum
func (m RecurringPaymentFrequencyenum) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateRecurringPaymentFrequencyenumEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validates this recurring payment frequencyenum based on context it is used
func (m RecurringPaymentFrequencyenum) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}
