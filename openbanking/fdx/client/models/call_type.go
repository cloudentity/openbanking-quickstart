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

// CallType CallType
//
// The call type for a stock option
//
// swagger:model CallType
type CallType string

func NewCallType(value CallType) *CallType {
	v := value
	return &v
}

const (

	// CallTypeCALL captures enum value "CALL"
	CallTypeCALL CallType = "CALL"

	// CallTypeMATURITY captures enum value "MATURITY"
	CallTypeMATURITY CallType = "MATURITY"

	// CallTypePREFUND captures enum value "PREFUND"
	CallTypePREFUND CallType = "PREFUND"

	// CallTypePUT captures enum value "PUT"
	CallTypePUT CallType = "PUT"
)

// for schema
var callTypeEnum []interface{}

func init() {
	var res []CallType
	if err := json.Unmarshal([]byte(`["CALL","MATURITY","PREFUND","PUT"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		callTypeEnum = append(callTypeEnum, v)
	}
}

func (m CallType) validateCallTypeEnum(path, location string, value CallType) error {
	if err := validate.EnumCase(path, location, value, callTypeEnum, true); err != nil {
		return err
	}
	return nil
}

// Validate validates this call type
func (m CallType) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateCallTypeEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validates this call type based on context it is used
func (m CallType) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}
