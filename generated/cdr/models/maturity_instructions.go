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

// MaturityInstructions MaturityInstructions
//
// Current instructions on action to be taken at maturity. This includes default actions that may be specified in the terms and conditions for the product e.g. roll-over to the same term and frequency of interest payments
// Example: HOLD_ON_MATURITY
//
// swagger:model MaturityInstructions
type MaturityInstructions string

func NewMaturityInstructions(value MaturityInstructions) *MaturityInstructions {
	v := value
	return &v
}

const (

	// MaturityInstructionsHOLDONMATURITY captures enum value "HOLD_ON_MATURITY"
	MaturityInstructionsHOLDONMATURITY MaturityInstructions = "HOLD_ON_MATURITY"

	// MaturityInstructionsPAIDOUTATMATURITY captures enum value "PAID_OUT_AT_MATURITY"
	MaturityInstructionsPAIDOUTATMATURITY MaturityInstructions = "PAID_OUT_AT_MATURITY"

	// MaturityInstructionsROLLEDOVER captures enum value "ROLLED_OVER"
	MaturityInstructionsROLLEDOVER MaturityInstructions = "ROLLED_OVER"
)

// for schema
var maturityInstructionsEnum []interface{}

func init() {
	var res []MaturityInstructions
	if err := json.Unmarshal([]byte(`["HOLD_ON_MATURITY","PAID_OUT_AT_MATURITY","ROLLED_OVER"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		maturityInstructionsEnum = append(maturityInstructionsEnum, v)
	}
}

func (m MaturityInstructions) validateMaturityInstructionsEnum(path, location string, value MaturityInstructions) error {
	if err := validate.EnumCase(path, location, value, maturityInstructionsEnum, true); err != nil {
		return err
	}
	return nil
}

// Validate validates this maturity instructions
func (m MaturityInstructions) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateMaturityInstructionsEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validates this maturity instructions based on context it is used
func (m MaturityInstructions) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}
