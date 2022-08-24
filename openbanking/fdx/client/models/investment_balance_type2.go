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

// InvestmentBalanceType2 InvestmentBalanceType2
//
// AMOUNT, PERCENTAGE
//
// swagger:model InvestmentBalanceType2
type InvestmentBalanceType2 string

func NewInvestmentBalanceType2(value InvestmentBalanceType2) *InvestmentBalanceType2 {
	v := value
	return &v
}

const (

	// InvestmentBalanceType2AMOUNT captures enum value "AMOUNT"
	InvestmentBalanceType2AMOUNT InvestmentBalanceType2 = "AMOUNT"

	// InvestmentBalanceType2PERCENTAGE captures enum value "PERCENTAGE"
	InvestmentBalanceType2PERCENTAGE InvestmentBalanceType2 = "PERCENTAGE"
)

// for schema
var investmentBalanceType2Enum []interface{}

func init() {
	var res []InvestmentBalanceType2
	if err := json.Unmarshal([]byte(`["AMOUNT","PERCENTAGE"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		investmentBalanceType2Enum = append(investmentBalanceType2Enum, v)
	}
}

func (m InvestmentBalanceType2) validateInvestmentBalanceType2Enum(path, location string, value InvestmentBalanceType2) error {
	if err := validate.EnumCase(path, location, value, investmentBalanceType2Enum, true); err != nil {
		return err
	}
	return nil
}

// Validate validates this investment balance type2
func (m InvestmentBalanceType2) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateInvestmentBalanceType2Enum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validates this investment balance type2 based on context it is used
func (m InvestmentBalanceType2) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}
