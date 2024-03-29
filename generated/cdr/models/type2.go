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

// Type2 Type2
//
// The type of the PayID
// Example: ABN
//
// swagger:model Type2
type Type2 string

func NewType2(value Type2) *Type2 {
	v := value
	return &v
}

const (

	// Type2ABN captures enum value "ABN"
	Type2ABN Type2 = "ABN"

	// Type2EMAIL captures enum value "EMAIL"
	Type2EMAIL Type2 = "EMAIL"

	// Type2ORGIDENTIFIER captures enum value "ORG_IDENTIFIER"
	Type2ORGIDENTIFIER Type2 = "ORG_IDENTIFIER"

	// Type2TELEPHONE captures enum value "TELEPHONE"
	Type2TELEPHONE Type2 = "TELEPHONE"
)

// for schema
var type2Enum []interface{}

func init() {
	var res []Type2
	if err := json.Unmarshal([]byte(`["ABN","EMAIL","ORG_IDENTIFIER","TELEPHONE"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		type2Enum = append(type2Enum, v)
	}
}

func (m Type2) validateType2Enum(path, location string, value Type2) error {
	if err := validate.EnumCase(path, location, value, type2Enum, true); err != nil {
		return err
	}
	return nil
}

// Validate validates this type2
func (m Type2) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateType2Enum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validates this type2 based on context it is used
func (m Type2) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}
