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

// UnitType3 UnitType3
//
// Type of unit. One of SHARES, CURRENCY
//
// swagger:model UnitType3
type UnitType3 string

func NewUnitType3(value UnitType3) *UnitType3 {
	v := value
	return &v
}

const (

	// UnitType3CURRENCY captures enum value "CURRENCY"
	UnitType3CURRENCY UnitType3 = "CURRENCY"

	// UnitType3SHARES captures enum value "SHARES"
	UnitType3SHARES UnitType3 = "SHARES"
)

// for schema
var unitType3Enum []interface{}

func init() {
	var res []UnitType3
	if err := json.Unmarshal([]byte(`["CURRENCY","SHARES"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		unitType3Enum = append(unitType3Enum, v)
	}
}

func (m UnitType3) validateUnitType3Enum(path, location string, value UnitType3) error {
	if err := validate.EnumCase(path, location, value, unitType3Enum, true); err != nil {
		return err
	}
	return nil
}

// Validate validates this unit type3
func (m UnitType3) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateUnitType3Enum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validates this unit type3 based on context it is used
func (m UnitType3) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}
