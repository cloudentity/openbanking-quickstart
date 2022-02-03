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

// Status Status
//
// Status of the transaction whether pending or posted. Note that there is currently no provision in the standards to guarantee the ability to correlate a pending transaction with an associated posted transaction
// Example: PENDING
//
// swagger:model Status
type Status string

const (

	// StatusPENDING captures enum value "PENDING"
	StatusPENDING Status = "PENDING"

	// StatusPOSTED captures enum value "POSTED"
	StatusPOSTED Status = "POSTED"
)

// for schema
var statusEnum []interface{}

func init() {
	var res []Status
	if err := json.Unmarshal([]byte(`["PENDING","POSTED"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		statusEnum = append(statusEnum, v)
	}
}

func (m Status) validateStatusEnum(path, location string, value Status) error {
	if err := validate.EnumCase(path, location, value, statusEnum, true); err != nil {
		return err
	}
	return nil
}

// Validate validates this status
func (m Status) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateStatusEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validates this status based on context it is used
func (m Status) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}
