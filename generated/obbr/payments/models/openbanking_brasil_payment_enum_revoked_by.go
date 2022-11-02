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

// OpenbankingBrasilPaymentEnumRevokedBy EnumRevokedBy
//
// Define qual das partes envolvidas na transao est realizando a revogao. Valores possveis:
// - USER (Revogado pelo usurio)
// - ASPSP (Provedor de servios de pagamento para servios de conta - Detentora de conta)
// - TPP (Instituies Provedoras - iniciadora de pagamentos)
// Example: USER
//
// swagger:model OpenbankingBrasilPaymentEnumRevokedBy
type OpenbankingBrasilPaymentEnumRevokedBy string

func NewOpenbankingBrasilPaymentEnumRevokedBy(value OpenbankingBrasilPaymentEnumRevokedBy) *OpenbankingBrasilPaymentEnumRevokedBy {
	v := value
	return &v
}

const (

	// OpenbankingBrasilPaymentEnumRevokedByUSER captures enum value "USER"
	OpenbankingBrasilPaymentEnumRevokedByUSER OpenbankingBrasilPaymentEnumRevokedBy = "USER"

	// OpenbankingBrasilPaymentEnumRevokedByASPSP captures enum value "ASPSP"
	OpenbankingBrasilPaymentEnumRevokedByASPSP OpenbankingBrasilPaymentEnumRevokedBy = "ASPSP"

	// OpenbankingBrasilPaymentEnumRevokedByTPP captures enum value "TPP"
	OpenbankingBrasilPaymentEnumRevokedByTPP OpenbankingBrasilPaymentEnumRevokedBy = "TPP"
)

// for schema
var openbankingBrasilPaymentEnumRevokedByEnum []interface{}

func init() {
	var res []OpenbankingBrasilPaymentEnumRevokedBy
	if err := json.Unmarshal([]byte(`["USER","ASPSP","TPP"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		openbankingBrasilPaymentEnumRevokedByEnum = append(openbankingBrasilPaymentEnumRevokedByEnum, v)
	}
}

func (m OpenbankingBrasilPaymentEnumRevokedBy) validateOpenbankingBrasilPaymentEnumRevokedByEnum(path, location string, value OpenbankingBrasilPaymentEnumRevokedBy) error {
	if err := validate.EnumCase(path, location, value, openbankingBrasilPaymentEnumRevokedByEnum, true); err != nil {
		return err
	}
	return nil
}

// Validate validates this openbanking brasil payment enum revoked by
func (m OpenbankingBrasilPaymentEnumRevokedBy) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateOpenbankingBrasilPaymentEnumRevokedByEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validates this openbanking brasil payment enum revoked by based on context it is used
func (m OpenbankingBrasilPaymentEnumRevokedBy) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}