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

// PayoutMode PayoutMode
//
// Frequency of annuity payments.<br/> <br/> | Value | Description |<br/> |-----|-----|<br/> | ANNUALLY | Paid Annually |<br/> | BIWEEKLY | Paid Bi-weekly |<br/> | DAILY | Paid Daily |<br/> | MONTHLY | Paid Monthly |<br/> | SEMIANNUALLY | Paid Semi-annually |<br/> | SEMIMONTHLY | Paid Semi-monthly |
//
// swagger:model PayoutMode
type PayoutMode string

func NewPayoutMode(value PayoutMode) *PayoutMode {
	v := value
	return &v
}

const (

	// PayoutModeANNUALLY captures enum value "ANNUALLY"
	PayoutModeANNUALLY PayoutMode = "ANNUALLY"

	// PayoutModeBIWEEKLY captures enum value "BIWEEKLY"
	PayoutModeBIWEEKLY PayoutMode = "BIWEEKLY"

	// PayoutModeDAILY captures enum value "DAILY"
	PayoutModeDAILY PayoutMode = "DAILY"

	// PayoutModeMONTHLY captures enum value "MONTHLY"
	PayoutModeMONTHLY PayoutMode = "MONTHLY"

	// PayoutModeSEMIANNUALLY captures enum value "SEMIANNUALLY"
	PayoutModeSEMIANNUALLY PayoutMode = "SEMIANNUALLY"

	// PayoutModeSEMIMONTHLY captures enum value "SEMIMONTHLY"
	PayoutModeSEMIMONTHLY PayoutMode = "SEMIMONTHLY"

	// PayoutModeWEEKLY captures enum value "WEEKLY"
	PayoutModeWEEKLY PayoutMode = "WEEKLY"
)

// for schema
var payoutModeEnum []interface{}

func init() {
	var res []PayoutMode
	if err := json.Unmarshal([]byte(`["ANNUALLY","BIWEEKLY","DAILY","MONTHLY","SEMIANNUALLY","SEMIMONTHLY","WEEKLY"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		payoutModeEnum = append(payoutModeEnum, v)
	}
}

func (m PayoutMode) validatePayoutModeEnum(path, location string, value PayoutMode) error {
	if err := validate.EnumCase(path, location, value, payoutModeEnum, true); err != nil {
		return err
	}
	return nil
}

// Validate validates this payout mode
func (m PayoutMode) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validatePayoutModeEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validates this payout mode based on context it is used
func (m PayoutMode) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}
