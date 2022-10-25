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

// ResultType ResultType
//
// Flag to indicate if you want a lightweight array of metadata (AccountDescriptor or Tax or Operations) or full item details (Account or a Tax subclass or Availability details). If set to 'lightweight', should only return the fields associated with the metadata entity.
//
// swagger:model ResultType
type ResultType string

func NewResultType(value ResultType) *ResultType {
	v := value
	return &v
}

const (

	// ResultTypeLightweight captures enum value "lightweight"
	ResultTypeLightweight ResultType = "lightweight"

	// ResultTypeDetails captures enum value "details"
	ResultTypeDetails ResultType = "details"
)

// for schema
var resultTypeEnum []interface{}

func init() {
	var res []ResultType
	if err := json.Unmarshal([]byte(`["lightweight","details"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		resultTypeEnum = append(resultTypeEnum, v)
	}
}

func (m ResultType) validateResultTypeEnum(path, location string, value ResultType) error {
	if err := validate.EnumCase(path, location, value, resultTypeEnum, true); err != nil {
		return err
	}
	return nil
}

// Validate validates this result type
func (m ResultType) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateResultTypeEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validates this result type based on context it is used
func (m ResultType) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}