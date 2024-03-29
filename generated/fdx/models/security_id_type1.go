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

// SecurityIDType1 SecurityIdType1
//
// CUSIP, ISIN, SEDOL, SICC, VALOR, WKN
//
// swagger:model SecurityIdType1
type SecurityIDType1 string

func NewSecurityIDType1(value SecurityIDType1) *SecurityIDType1 {
	v := value
	return &v
}

const (

	// SecurityIDType1CINS captures enum value "CINS"
	SecurityIDType1CINS SecurityIDType1 = "CINS"

	// SecurityIDType1CMC captures enum value "CMC"
	SecurityIDType1CMC SecurityIDType1 = "CMC"

	// SecurityIDType1CME captures enum value "CME"
	SecurityIDType1CME SecurityIDType1 = "CME"

	// SecurityIDType1CUSIP captures enum value "CUSIP"
	SecurityIDType1CUSIP SecurityIDType1 = "CUSIP"

	// SecurityIDType1ISIN captures enum value "ISIN"
	SecurityIDType1ISIN SecurityIDType1 = "ISIN"

	// SecurityIDType1ITSA captures enum value "ITSA"
	SecurityIDType1ITSA SecurityIDType1 = "ITSA"

	// SecurityIDType1NASDAQ captures enum value "NASDAQ"
	SecurityIDType1NASDAQ SecurityIDType1 = "NASDAQ"

	// SecurityIDType1SEDOL captures enum value "SEDOL"
	SecurityIDType1SEDOL SecurityIDType1 = "SEDOL"

	// SecurityIDType1SICC captures enum value "SICC"
	SecurityIDType1SICC SecurityIDType1 = "SICC"

	// SecurityIDType1VALOR captures enum value "VALOR"
	SecurityIDType1VALOR SecurityIDType1 = "VALOR"

	// SecurityIDType1WKN captures enum value "WKN"
	SecurityIDType1WKN SecurityIDType1 = "WKN"
)

// for schema
var securityIdType1Enum []interface{}

func init() {
	var res []SecurityIDType1
	if err := json.Unmarshal([]byte(`["CINS","CMC","CME","CUSIP","ISIN","ITSA","NASDAQ","SEDOL","SICC","VALOR","WKN"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		securityIdType1Enum = append(securityIdType1Enum, v)
	}
}

func (m SecurityIDType1) validateSecurityIDType1Enum(path, location string, value SecurityIDType1) error {
	if err := validate.EnumCase(path, location, value, securityIdType1Enum, true); err != nil {
		return err
	}
	return nil
}

// Validate validates this security Id type1
func (m SecurityIDType1) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateSecurityIDType1Enum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validates this security Id type1 based on context it is used
func (m SecurityIDType1) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}
