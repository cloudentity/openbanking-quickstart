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

// HoldingSubType2 HoldingSubType2
//
// MONEYMARKET, CASH
//
// swagger:model HoldingSubType2
type HoldingSubType2 string

func NewHoldingSubType2(value HoldingSubType2) *HoldingSubType2 {
	v := value
	return &v
}

const (

	// HoldingSubType2CASH captures enum value "CASH"
	HoldingSubType2CASH HoldingSubType2 = "CASH"

	// HoldingSubType2MONEYMARKET captures enum value "MONEYMARKET"
	HoldingSubType2MONEYMARKET HoldingSubType2 = "MONEYMARKET"
)

// for schema
var holdingSubType2Enum []interface{}

func init() {
	var res []HoldingSubType2
	if err := json.Unmarshal([]byte(`["CASH","MONEYMARKET"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		holdingSubType2Enum = append(holdingSubType2Enum, v)
	}
}

func (m HoldingSubType2) validateHoldingSubType2Enum(path, location string, value HoldingSubType2) error {
	if err := validate.EnumCase(path, location, value, holdingSubType2Enum, true); err != nil {
		return err
	}
	return nil
}

// Validate validates this holding sub type2
func (m HoldingSubType2) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateHoldingSubType2Enum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validates this holding sub type2 based on context it is used
func (m HoldingSubType2) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}
