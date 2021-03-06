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

// OpenbankingBrasilEnumAccountPaymentsType EnumAccountPaymentsType
//
// Tipos de contas usadas para pagamento via Pix.
// Modalidades tradicionais previstas pela Resoluo 4.753, no contemplando contas vinculadas,
// conta de domiciliados no exterior, contas em moedas estrangeiras e conta correspondente moeda eletrnica.
// Segue descrio de cada valor do ENUM para o escopo do Pix.
// CACC - Current - Conta Corrente.
// SLRY - Salary - Conta-Salrio.
// SVGS - Savings - Conta de Poupana.
// TRAN - TransactingAccount - Conta de Pagamento pr-paga.
//
// swagger:model OpenbankingBrasilEnumAccountPaymentsType
type OpenbankingBrasilEnumAccountPaymentsType string

const (

	// OpenbankingBrasilEnumAccountPaymentsTypeCACC captures enum value "CACC"
	OpenbankingBrasilEnumAccountPaymentsTypeCACC OpenbankingBrasilEnumAccountPaymentsType = "CACC"

	// OpenbankingBrasilEnumAccountPaymentsTypeSLRY captures enum value "SLRY"
	OpenbankingBrasilEnumAccountPaymentsTypeSLRY OpenbankingBrasilEnumAccountPaymentsType = "SLRY"

	// OpenbankingBrasilEnumAccountPaymentsTypeSVGS captures enum value "SVGS"
	OpenbankingBrasilEnumAccountPaymentsTypeSVGS OpenbankingBrasilEnumAccountPaymentsType = "SVGS"

	// OpenbankingBrasilEnumAccountPaymentsTypeTRAN captures enum value "TRAN"
	OpenbankingBrasilEnumAccountPaymentsTypeTRAN OpenbankingBrasilEnumAccountPaymentsType = "TRAN"
)

// for schema
var openbankingBrasilEnumAccountPaymentsTypeEnum []interface{}

func init() {
	var res []OpenbankingBrasilEnumAccountPaymentsType
	if err := json.Unmarshal([]byte(`["CACC","SLRY","SVGS","TRAN"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		openbankingBrasilEnumAccountPaymentsTypeEnum = append(openbankingBrasilEnumAccountPaymentsTypeEnum, v)
	}
}

func (m OpenbankingBrasilEnumAccountPaymentsType) validateOpenbankingBrasilEnumAccountPaymentsTypeEnum(path, location string, value OpenbankingBrasilEnumAccountPaymentsType) error {
	if err := validate.EnumCase(path, location, value, openbankingBrasilEnumAccountPaymentsTypeEnum, true); err != nil {
		return err
	}
	return nil
}

// Validate validates this openbanking brasil enum account payments type
func (m OpenbankingBrasilEnumAccountPaymentsType) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateOpenbankingBrasilEnumAccountPaymentsTypeEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validates this openbanking brasil enum account payments type based on context it is used
func (m OpenbankingBrasilEnumAccountPaymentsType) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}
