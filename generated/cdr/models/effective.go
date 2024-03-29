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

// Effective effective
// Example: ALL
//
// swagger:model effective
type Effective string

func NewEffective(value Effective) *Effective {
	v := value
	return &v
}

const (

	// EffectiveALL captures enum value "ALL"
	EffectiveALL Effective = "ALL"

	// EffectiveCURRENT captures enum value "CURRENT"
	EffectiveCURRENT Effective = "CURRENT"

	// EffectiveFUTURE captures enum value "FUTURE"
	EffectiveFUTURE Effective = "FUTURE"
)

// for schema
var effectiveEnum []interface{}

func init() {
	var res []Effective
	if err := json.Unmarshal([]byte(`["ALL","CURRENT","FUTURE"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		effectiveEnum = append(effectiveEnum, v)
	}
}

func (m Effective) validateEffectiveEnum(path, location string, value Effective) error {
	if err := validate.EnumCase(path, location, value, effectiveEnum, true); err != nil {
		return err
	}
	return nil
}

// Validate validates this effective
func (m Effective) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateEffectiveEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validates this effective based on context it is used
func (m Effective) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}
