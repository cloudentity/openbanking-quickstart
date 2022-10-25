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

// SecurityType SecurityType
//
// The type of a security
//
// swagger:model SecurityType
type SecurityType string

func NewSecurityType(value SecurityType) *SecurityType {
	v := value
	return &v
}

const (

	// SecurityTypeBOND captures enum value "BOND"
	SecurityTypeBOND SecurityType = "BOND"

	// SecurityTypeDEBT captures enum value "DEBT"
	SecurityTypeDEBT SecurityType = "DEBT"

	// SecurityTypeDIGITALASSET captures enum value "DIGITALASSET"
	SecurityTypeDIGITALASSET SecurityType = "DIGITALASSET"

	// SecurityTypeMUTUALFUND captures enum value "MUTUALFUND"
	SecurityTypeMUTUALFUND SecurityType = "MUTUALFUND"

	// SecurityTypeOPTION captures enum value "OPTION"
	SecurityTypeOPTION SecurityType = "OPTION"

	// SecurityTypeOTHER captures enum value "OTHER"
	SecurityTypeOTHER SecurityType = "OTHER"

	// SecurityTypeSTOCK captures enum value "STOCK"
	SecurityTypeSTOCK SecurityType = "STOCK"

	// SecurityTypeSWEEP captures enum value "SWEEP"
	SecurityTypeSWEEP SecurityType = "SWEEP"
)

// for schema
var securityTypeEnum []interface{}

func init() {
	var res []SecurityType
	if err := json.Unmarshal([]byte(`["BOND","DEBT","DIGITALASSET","MUTUALFUND","OPTION","OTHER","STOCK","SWEEP"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		securityTypeEnum = append(securityTypeEnum, v)
	}
}

func (m SecurityType) validateSecurityTypeEnum(path, location string, value SecurityType) error {
	if err := validate.EnumCase(path, location, value, securityTypeEnum, true); err != nil {
		return err
	}
	return nil
}

// Validate validates this security type
func (m SecurityType) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateSecurityTypeEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validates this security type based on context it is used
func (m SecurityType) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}