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

// InterestPaymentDue InterestPaymentDue
//
// When loan payments are due to be paid within each period. The investment benefit of earlier payments affect the rate that can be offered
// Example: IN_ADVANCE
//
// swagger:model InterestPaymentDue
type InterestPaymentDue string

func NewInterestPaymentDue(value InterestPaymentDue) *InterestPaymentDue {
	v := value
	return &v
}

const (

	// InterestPaymentDueINADVANCE captures enum value "IN_ADVANCE"
	InterestPaymentDueINADVANCE InterestPaymentDue = "IN_ADVANCE"

	// InterestPaymentDueINARREARS captures enum value "IN_ARREARS"
	InterestPaymentDueINARREARS InterestPaymentDue = "IN_ARREARS"
)

// for schema
var interestPaymentDueEnum []interface{}

func init() {
	var res []InterestPaymentDue
	if err := json.Unmarshal([]byte(`["IN_ADVANCE","IN_ARREARS"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		interestPaymentDueEnum = append(interestPaymentDueEnum, v)
	}
}

func (m InterestPaymentDue) validateInterestPaymentDueEnum(path, location string, value InterestPaymentDue) error {
	if err := validate.EnumCase(path, location, value, interestPaymentDueEnum, true); err != nil {
		return err
	}
	return nil
}

// Validate validates this interest payment due
func (m InterestPaymentDue) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateInterestPaymentDueEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validates this interest payment due based on context it is used
func (m InterestPaymentDue) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}
