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

// OBMinMaxType1Code Min Max type
//
// swagger:model OB_MinMaxType1Code
type OBMinMaxType1Code string

const (

	// OBMinMaxType1CodeFMMN captures enum value "FMMN"
	OBMinMaxType1CodeFMMN OBMinMaxType1Code = "FMMN"

	// OBMinMaxType1CodeFMMX captures enum value "FMMX"
	OBMinMaxType1CodeFMMX OBMinMaxType1Code = "FMMX"
)

// for schema
var oBMinMaxType1CodeEnum []interface{}

func init() {
	var res []OBMinMaxType1Code
	if err := json.Unmarshal([]byte(`["FMMN","FMMX"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		oBMinMaxType1CodeEnum = append(oBMinMaxType1CodeEnum, v)
	}
}

func (m OBMinMaxType1Code) validateOBMinMaxType1CodeEnum(path, location string, value OBMinMaxType1Code) error {
	if err := validate.EnumCase(path, location, value, oBMinMaxType1CodeEnum, true); err != nil {
		return err
	}
	return nil
}

// Validate validates this o b min max type1 code
func (m OBMinMaxType1Code) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateOBMinMaxType1CodeEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validates this o b min max type1 code based on context it is used
func (m OBMinMaxType1Code) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}
