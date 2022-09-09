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

// SubAccountType SubAccountType
//
// The subtype of an account
//
// swagger:model SubAccountType
type SubAccountType string

func NewSubAccountType(value SubAccountType) *SubAccountType {
	v := value
	return &v
}

const (

	// SubAccountTypeCASH captures enum value "CASH"
	SubAccountTypeCASH SubAccountType = "CASH"

	// SubAccountTypeMARGIN captures enum value "MARGIN"
	SubAccountTypeMARGIN SubAccountType = "MARGIN"

	// SubAccountTypeSHORT captures enum value "SHORT"
	SubAccountTypeSHORT SubAccountType = "SHORT"

	// SubAccountTypeOTHER captures enum value "OTHER"
	SubAccountTypeOTHER SubAccountType = "OTHER"
)

// for schema
var subAccountTypeEnum []interface{}

func init() {
	var res []SubAccountType
	if err := json.Unmarshal([]byte(`["CASH","MARGIN","SHORT","OTHER"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		subAccountTypeEnum = append(subAccountTypeEnum, v)
	}
}

func (m SubAccountType) validateSubAccountTypeEnum(path, location string, value SubAccountType) error {
	if err := validate.EnumCase(path, location, value, subAccountTypeEnum, true); err != nil {
		return err
	}
	return nil
}

// Validate validates this sub account type
func (m SubAccountType) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateSubAccountTypeEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validates this sub account type based on context it is used
func (m SubAccountType) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}
