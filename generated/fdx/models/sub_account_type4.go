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

// SubAccountType4 SubAccountType4
//
// CASH, MARGIN, SHORT, OTHER
//
// swagger:model SubAccountType4
type SubAccountType4 string

func NewSubAccountType4(value SubAccountType4) *SubAccountType4 {
	v := value
	return &v
}

const (

	// SubAccountType4CASH captures enum value "CASH"
	SubAccountType4CASH SubAccountType4 = "CASH"

	// SubAccountType4MARGIN captures enum value "MARGIN"
	SubAccountType4MARGIN SubAccountType4 = "MARGIN"

	// SubAccountType4SHORT captures enum value "SHORT"
	SubAccountType4SHORT SubAccountType4 = "SHORT"

	// SubAccountType4OTHER captures enum value "OTHER"
	SubAccountType4OTHER SubAccountType4 = "OTHER"
)

// for schema
var subAccountType4Enum []interface{}

func init() {
	var res []SubAccountType4
	if err := json.Unmarshal([]byte(`["CASH","MARGIN","SHORT","OTHER"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		subAccountType4Enum = append(subAccountType4Enum, v)
	}
}

func (m SubAccountType4) validateSubAccountType4Enum(path, location string, value SubAccountType4) error {
	if err := validate.EnumCase(path, location, value, subAccountType4Enum, true); err != nil {
		return err
	}
	return nil
}

// Validate validates this sub account type4
func (m SubAccountType4) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateSubAccountType4Enum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validates this sub account type4 based on context it is used
func (m SubAccountType4) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}
