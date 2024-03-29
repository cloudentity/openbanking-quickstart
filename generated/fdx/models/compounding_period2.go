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

// CompoundingPeriod2 CompoundingPeriod2
//
// DAILY, WEEKLY, BIWEEKLY, SEMIMONTHLY, MONTHLY, SEMIANNUALLY, ANNUALLY
//
// swagger:model CompoundingPeriod2
type CompoundingPeriod2 string

func NewCompoundingPeriod2(value CompoundingPeriod2) *CompoundingPeriod2 {
	v := value
	return &v
}

const (

	// CompoundingPeriod2ANNUALLY captures enum value "ANNUALLY"
	CompoundingPeriod2ANNUALLY CompoundingPeriod2 = "ANNUALLY"

	// CompoundingPeriod2BIWEEKLY captures enum value "BIWEEKLY"
	CompoundingPeriod2BIWEEKLY CompoundingPeriod2 = "BIWEEKLY"

	// CompoundingPeriod2DAILY captures enum value "DAILY"
	CompoundingPeriod2DAILY CompoundingPeriod2 = "DAILY"

	// CompoundingPeriod2MONTHLY captures enum value "MONTHLY"
	CompoundingPeriod2MONTHLY CompoundingPeriod2 = "MONTHLY"

	// CompoundingPeriod2SEMIANNUALLY captures enum value "SEMIANNUALLY"
	CompoundingPeriod2SEMIANNUALLY CompoundingPeriod2 = "SEMIANNUALLY"

	// CompoundingPeriod2SEMIMONTHLY captures enum value "SEMIMONTHLY"
	CompoundingPeriod2SEMIMONTHLY CompoundingPeriod2 = "SEMIMONTHLY"

	// CompoundingPeriod2WEEKLY captures enum value "WEEKLY"
	CompoundingPeriod2WEEKLY CompoundingPeriod2 = "WEEKLY"
)

// for schema
var compoundingPeriod2Enum []interface{}

func init() {
	var res []CompoundingPeriod2
	if err := json.Unmarshal([]byte(`["ANNUALLY","BIWEEKLY","DAILY","MONTHLY","SEMIANNUALLY","SEMIMONTHLY","WEEKLY"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		compoundingPeriod2Enum = append(compoundingPeriod2Enum, v)
	}
}

func (m CompoundingPeriod2) validateCompoundingPeriod2Enum(path, location string, value CompoundingPeriod2) error {
	if err := validate.EnumCase(path, location, value, compoundingPeriod2Enum, true); err != nil {
		return err
	}
	return nil
}

// Validate validates this compounding period2
func (m CompoundingPeriod2) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateCompoundingPeriod2Enum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validates this compounding period2 based on context it is used
func (m CompoundingPeriod2) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}
