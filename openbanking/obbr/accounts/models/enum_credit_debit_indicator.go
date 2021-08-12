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

// EnumCreditDebitIndicator EnumCreditDebitIndicator
//
// Indicador do tipo de lançamento:
// Débito (no extrato) Em um extrato bancário, os débitos, marcados com a letra “D” ao lado do valor registrado, informam as saídas de dinheiro na conta-corrente.
// Crédito (no extrato) Em um extrato bancário, os créditos, marcados com a letra “C” ao lado do valor registrado, informam as entradas de dinheiro na conta-corrente.
//
// swagger:model EnumCreditDebitIndicator
type EnumCreditDebitIndicator string

const (

	// EnumCreditDebitIndicatorCREDITO captures enum value "CREDITO"
	EnumCreditDebitIndicatorCREDITO EnumCreditDebitIndicator = "CREDITO"

	// EnumCreditDebitIndicatorDEBITO captures enum value "DEBITO"
	EnumCreditDebitIndicatorDEBITO EnumCreditDebitIndicator = "DEBITO"
)

// for schema
var enumCreditDebitIndicatorEnum []interface{}

func init() {
	var res []EnumCreditDebitIndicator
	if err := json.Unmarshal([]byte(`["CREDITO","DEBITO"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		enumCreditDebitIndicatorEnum = append(enumCreditDebitIndicatorEnum, v)
	}
}

func (m EnumCreditDebitIndicator) validateEnumCreditDebitIndicatorEnum(path, location string, value EnumCreditDebitIndicator) error {
	if err := validate.EnumCase(path, location, value, enumCreditDebitIndicatorEnum, true); err != nil {
		return err
	}
	return nil
}

// Validate validates this enum credit debit indicator
func (m EnumCreditDebitIndicator) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateEnumCreditDebitIndicatorEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validates this enum credit debit indicator based on context it is used
func (m EnumCreditDebitIndicator) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}
