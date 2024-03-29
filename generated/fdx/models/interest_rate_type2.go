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

// InterestRateType2 InterestRateType2
//
// FIXED or VARIABLE
//
// swagger:model InterestRateType2
type InterestRateType2 string

func NewInterestRateType2(value InterestRateType2) *InterestRateType2 {
	v := value
	return &v
}

const (

	// InterestRateType2FIXED captures enum value "FIXED"
	InterestRateType2FIXED InterestRateType2 = "FIXED"

	// InterestRateType2VARIABLE captures enum value "VARIABLE"
	InterestRateType2VARIABLE InterestRateType2 = "VARIABLE"
)

// for schema
var interestRateType2Enum []interface{}

func init() {
	var res []InterestRateType2
	if err := json.Unmarshal([]byte(`["FIXED","VARIABLE"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		interestRateType2Enum = append(interestRateType2Enum, v)
	}
}

func (m InterestRateType2) validateInterestRateType2Enum(path, location string, value InterestRateType2) error {
	if err := validate.EnumCase(path, location, value, interestRateType2Enum, true); err != nil {
		return err
	}
	return nil
}

// Validate validates this interest rate type2
func (m InterestRateType2) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateInterestRateType2Enum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validates this interest rate type2 based on context it is used
func (m InterestRateType2) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}
