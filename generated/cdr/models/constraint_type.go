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

// ConstraintType ConstraintType
//
// The type of constraint described.  See the next section for an overview of valid values and their meaning
// Example: MAX_BALANCE
//
// swagger:model ConstraintType
type ConstraintType string

func NewConstraintType(value ConstraintType) *ConstraintType {
	v := value
	return &v
}

const (

	// ConstraintTypeMAXBALANCE captures enum value "MAX_BALANCE"
	ConstraintTypeMAXBALANCE ConstraintType = "MAX_BALANCE"

	// ConstraintTypeMAXLIMIT captures enum value "MAX_LIMIT"
	ConstraintTypeMAXLIMIT ConstraintType = "MAX_LIMIT"

	// ConstraintTypeMINBALANCE captures enum value "MIN_BALANCE"
	ConstraintTypeMINBALANCE ConstraintType = "MIN_BALANCE"

	// ConstraintTypeMINLIMIT captures enum value "MIN_LIMIT"
	ConstraintTypeMINLIMIT ConstraintType = "MIN_LIMIT"

	// ConstraintTypeOPENINGBALANCE captures enum value "OPENING_BALANCE"
	ConstraintTypeOPENINGBALANCE ConstraintType = "OPENING_BALANCE"
)

// for schema
var constraintTypeEnum []interface{}

func init() {
	var res []ConstraintType
	if err := json.Unmarshal([]byte(`["MAX_BALANCE","MAX_LIMIT","MIN_BALANCE","MIN_LIMIT","OPENING_BALANCE"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		constraintTypeEnum = append(constraintTypeEnum, v)
	}
}

func (m ConstraintType) validateConstraintTypeEnum(path, location string, value ConstraintType) error {
	if err := validate.EnumCase(path, location, value, constraintTypeEnum, true); err != nil {
		return err
	}
	return nil
}

// Validate validates this constraint type
func (m ConstraintType) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateConstraintTypeEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validates this constraint type based on context it is used
func (m ConstraintType) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}
