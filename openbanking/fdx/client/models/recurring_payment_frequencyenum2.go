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

// RecurringPaymentFrequencyenum2 RecurringPaymentFrequencyenum2
//
// Defines how often the payment repeats
//
// swagger:model RecurringPaymentFrequencyenum2
type RecurringPaymentFrequencyenum2 string

func NewRecurringPaymentFrequencyenum2(value RecurringPaymentFrequencyenum2) *RecurringPaymentFrequencyenum2 {
	v := value
	return &v
}

const (

	// RecurringPaymentFrequencyenum2WEEKLY captures enum value "WEEKLY"
	RecurringPaymentFrequencyenum2WEEKLY RecurringPaymentFrequencyenum2 = "WEEKLY"

	// RecurringPaymentFrequencyenum2BIWEEKLY captures enum value "BIWEEKLY"
	RecurringPaymentFrequencyenum2BIWEEKLY RecurringPaymentFrequencyenum2 = "BIWEEKLY"

	// RecurringPaymentFrequencyenum2TWICEMONTHLY captures enum value "TWICEMONTHLY"
	RecurringPaymentFrequencyenum2TWICEMONTHLY RecurringPaymentFrequencyenum2 = "TWICEMONTHLY"

	// RecurringPaymentFrequencyenum2MONTHLY captures enum value "MONTHLY"
	RecurringPaymentFrequencyenum2MONTHLY RecurringPaymentFrequencyenum2 = "MONTHLY"

	// RecurringPaymentFrequencyenum2FOURWEEKS captures enum value "FOURWEEKS"
	RecurringPaymentFrequencyenum2FOURWEEKS RecurringPaymentFrequencyenum2 = "FOURWEEKS"

	// RecurringPaymentFrequencyenum2BIMONTHLY captures enum value "BIMONTHLY"
	RecurringPaymentFrequencyenum2BIMONTHLY RecurringPaymentFrequencyenum2 = "BIMONTHLY"

	// RecurringPaymentFrequencyenum2QUARTERLY captures enum value "QUARTERLY"
	RecurringPaymentFrequencyenum2QUARTERLY RecurringPaymentFrequencyenum2 = "QUARTERLY"

	// RecurringPaymentFrequencyenum2SEMIANNUALLY captures enum value "SEMIANNUALLY"
	RecurringPaymentFrequencyenum2SEMIANNUALLY RecurringPaymentFrequencyenum2 = "SEMIANNUALLY"

	// RecurringPaymentFrequencyenum2ANNUALLY captures enum value "ANNUALLY"
	RecurringPaymentFrequencyenum2ANNUALLY RecurringPaymentFrequencyenum2 = "ANNUALLY"
)

// for schema
var recurringPaymentFrequencyenum2Enum []interface{}

func init() {
	var res []RecurringPaymentFrequencyenum2
	if err := json.Unmarshal([]byte(`["WEEKLY","BIWEEKLY","TWICEMONTHLY","MONTHLY","FOURWEEKS","BIMONTHLY","QUARTERLY","SEMIANNUALLY","ANNUALLY"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		recurringPaymentFrequencyenum2Enum = append(recurringPaymentFrequencyenum2Enum, v)
	}
}

func (m RecurringPaymentFrequencyenum2) validateRecurringPaymentFrequencyenum2Enum(path, location string, value RecurringPaymentFrequencyenum2) error {
	if err := validate.EnumCase(path, location, value, recurringPaymentFrequencyenum2Enum, true); err != nil {
		return err
	}
	return nil
}

// Validate validates this recurring payment frequencyenum2
func (m RecurringPaymentFrequencyenum2) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateRecurringPaymentFrequencyenum2Enum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validates this recurring payment frequencyenum2 based on context it is used
func (m RecurringPaymentFrequencyenum2) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}