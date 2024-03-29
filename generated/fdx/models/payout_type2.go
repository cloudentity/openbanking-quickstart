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

// PayoutType2 PayoutType2
//
// Indicates type of payout such as immediate or deferred
//
// swagger:model PayoutType2
type PayoutType2 string

func NewPayoutType2(value PayoutType2) *PayoutType2 {
	v := value
	return &v
}

const (

	// PayoutType2DEFERRED captures enum value "DEFERRED"
	PayoutType2DEFERRED PayoutType2 = "DEFERRED"

	// PayoutType2IMMEDIATE captures enum value "IMMEDIATE"
	PayoutType2IMMEDIATE PayoutType2 = "IMMEDIATE"
)

// for schema
var payoutType2Enum []interface{}

func init() {
	var res []PayoutType2
	if err := json.Unmarshal([]byte(`["DEFERRED","IMMEDIATE"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		payoutType2Enum = append(payoutType2Enum, v)
	}
}

func (m PayoutType2) validatePayoutType2Enum(path, location string, value PayoutType2) error {
	if err := validate.EnumCase(path, location, value, payoutType2Enum, true); err != nil {
		return err
	}
	return nil
}

// Validate validates this payout type2
func (m PayoutType2) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validatePayoutType2Enum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validates this payout type2 based on context it is used
func (m PayoutType2) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}
