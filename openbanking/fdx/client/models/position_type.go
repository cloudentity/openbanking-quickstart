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

// PositionType PositionType
//
// The type of an investment position
//
// swagger:model PositionType
type PositionType string

func NewPositionType(value PositionType) *PositionType {
	v := value
	return &v
}

const (

	// PositionTypeLONG captures enum value "LONG"
	PositionTypeLONG PositionType = "LONG"

	// PositionTypeSHORT captures enum value "SHORT"
	PositionTypeSHORT PositionType = "SHORT"
)

// for schema
var positionTypeEnum []interface{}

func init() {
	var res []PositionType
	if err := json.Unmarshal([]byte(`["LONG","SHORT"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		positionTypeEnum = append(positionTypeEnum, v)
	}
}

func (m PositionType) validatePositionTypeEnum(path, location string, value PositionType) error {
	if err := validate.EnumCase(path, location, value, positionTypeEnum, true); err != nil {
		return err
	}
	return nil
}

// Validate validates this position type
func (m PositionType) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validatePositionTypeEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validates this position type based on context it is used
func (m PositionType) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}
