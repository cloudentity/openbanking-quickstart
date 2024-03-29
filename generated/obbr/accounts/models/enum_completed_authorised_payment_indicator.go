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

// EnumCompletedAuthorisedPaymentIndicator EnumCompletedAuthorisedPaymentIndicator
//
// Indicador da transação:
// - Transação efetivada
// - Lançamento futuro
//
// swagger:model EnumCompletedAuthorisedPaymentIndicator
type EnumCompletedAuthorisedPaymentIndicator string

func NewEnumCompletedAuthorisedPaymentIndicator(value EnumCompletedAuthorisedPaymentIndicator) *EnumCompletedAuthorisedPaymentIndicator {
	v := value
	return &v
}

const (

	// EnumCompletedAuthorisedPaymentIndicatorTRANSACAOEFETIVADA captures enum value "TRANSACAO_EFETIVADA"
	EnumCompletedAuthorisedPaymentIndicatorTRANSACAOEFETIVADA EnumCompletedAuthorisedPaymentIndicator = "TRANSACAO_EFETIVADA"

	// EnumCompletedAuthorisedPaymentIndicatorLANCAMENTOFUTURO captures enum value "LANCAMENTO_FUTURO"
	EnumCompletedAuthorisedPaymentIndicatorLANCAMENTOFUTURO EnumCompletedAuthorisedPaymentIndicator = "LANCAMENTO_FUTURO"
)

// for schema
var enumCompletedAuthorisedPaymentIndicatorEnum []interface{}

func init() {
	var res []EnumCompletedAuthorisedPaymentIndicator
	if err := json.Unmarshal([]byte(`["TRANSACAO_EFETIVADA","LANCAMENTO_FUTURO"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		enumCompletedAuthorisedPaymentIndicatorEnum = append(enumCompletedAuthorisedPaymentIndicatorEnum, v)
	}
}

func (m EnumCompletedAuthorisedPaymentIndicator) validateEnumCompletedAuthorisedPaymentIndicatorEnum(path, location string, value EnumCompletedAuthorisedPaymentIndicator) error {
	if err := validate.EnumCase(path, location, value, enumCompletedAuthorisedPaymentIndicatorEnum, true); err != nil {
		return err
	}
	return nil
}

// Validate validates this enum completed authorised payment indicator
func (m EnumCompletedAuthorisedPaymentIndicator) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateEnumCompletedAuthorisedPaymentIndicatorEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validates this enum completed authorised payment indicator based on context it is used
func (m EnumCompletedAuthorisedPaymentIndicator) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}
