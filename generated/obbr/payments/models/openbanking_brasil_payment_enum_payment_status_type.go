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

// OpenbankingBrasilPaymentEnumPaymentStatusType EnumPaymentStatusType
//
// Estado atual da iniciao de pagamento. O estado evolui na seguinte ordem:
// 1. PDNG (PENDING) - Iniciao de pagamento ou transao de pagamento est pendente. Checagens adicionais em realizao.
// 2. SASP (SCHEDULE_ACCEPTED_SETTLEMENT_IN_PROCESS) - Indica que o processo de agendamento est em processamento.
// 3. SASC (SCHEDULE_ACCEPTED_SETTLEMENT_COMPLETED) - Indica que o processo de agendamento foi realizado.
// 4. PART (PARTIALLY ACCEPTED) - Aguardando autorizao mltipla alada.
// 5. ACSP (ACCEPTED_SETTLEMENT_IN_PROCESS) - Iniciao de pagamento aceita e processamento do pagamento foi iniciado.
// 6. ACSC (ACCEPTED_SETTLEMENT_COMPLETED_DEBITOR_ACCOUNT) - Dbito realizado na conta do pagador.
// 7. ACCC (ACCEPTED_SETTLEMENT_COMPLETED) - Crdito realizado na instituio de destino.
//
// Em caso insucesso:
// RJCT (REJECTED) - Instruo de pagamento rejeitada.
// Example: PDNG
//
// swagger:model OpenbankingBrasilPaymentEnumPaymentStatusType
type OpenbankingBrasilPaymentEnumPaymentStatusType string

func NewOpenbankingBrasilPaymentEnumPaymentStatusType(value OpenbankingBrasilPaymentEnumPaymentStatusType) *OpenbankingBrasilPaymentEnumPaymentStatusType {
	v := value
	return &v
}

const (

	// OpenbankingBrasilPaymentEnumPaymentStatusTypePDNG captures enum value "PDNG"
	OpenbankingBrasilPaymentEnumPaymentStatusTypePDNG OpenbankingBrasilPaymentEnumPaymentStatusType = "PDNG"

	// OpenbankingBrasilPaymentEnumPaymentStatusTypeSASP captures enum value "SASP"
	OpenbankingBrasilPaymentEnumPaymentStatusTypeSASP OpenbankingBrasilPaymentEnumPaymentStatusType = "SASP"

	// OpenbankingBrasilPaymentEnumPaymentStatusTypeSASC captures enum value "SASC"
	OpenbankingBrasilPaymentEnumPaymentStatusTypeSASC OpenbankingBrasilPaymentEnumPaymentStatusType = "SASC"

	// OpenbankingBrasilPaymentEnumPaymentStatusTypePART captures enum value "PART"
	OpenbankingBrasilPaymentEnumPaymentStatusTypePART OpenbankingBrasilPaymentEnumPaymentStatusType = "PART"

	// OpenbankingBrasilPaymentEnumPaymentStatusTypeACSP captures enum value "ACSP"
	OpenbankingBrasilPaymentEnumPaymentStatusTypeACSP OpenbankingBrasilPaymentEnumPaymentStatusType = "ACSP"

	// OpenbankingBrasilPaymentEnumPaymentStatusTypeACSC captures enum value "ACSC"
	OpenbankingBrasilPaymentEnumPaymentStatusTypeACSC OpenbankingBrasilPaymentEnumPaymentStatusType = "ACSC"

	// OpenbankingBrasilPaymentEnumPaymentStatusTypeACCC captures enum value "ACCC"
	OpenbankingBrasilPaymentEnumPaymentStatusTypeACCC OpenbankingBrasilPaymentEnumPaymentStatusType = "ACCC"

	// OpenbankingBrasilPaymentEnumPaymentStatusTypeRJCT captures enum value "RJCT"
	OpenbankingBrasilPaymentEnumPaymentStatusTypeRJCT OpenbankingBrasilPaymentEnumPaymentStatusType = "RJCT"
)

// for schema
var openbankingBrasilPaymentEnumPaymentStatusTypeEnum []interface{}

func init() {
	var res []OpenbankingBrasilPaymentEnumPaymentStatusType
	if err := json.Unmarshal([]byte(`["PDNG","SASP","SASC","PART","ACSP","ACSC","ACCC","RJCT"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		openbankingBrasilPaymentEnumPaymentStatusTypeEnum = append(openbankingBrasilPaymentEnumPaymentStatusTypeEnum, v)
	}
}

func (m OpenbankingBrasilPaymentEnumPaymentStatusType) validateOpenbankingBrasilPaymentEnumPaymentStatusTypeEnum(path, location string, value OpenbankingBrasilPaymentEnumPaymentStatusType) error {
	if err := validate.EnumCase(path, location, value, openbankingBrasilPaymentEnumPaymentStatusTypeEnum, true); err != nil {
		return err
	}
	return nil
}

// Validate validates this openbanking brasil payment enum payment status type
func (m OpenbankingBrasilPaymentEnumPaymentStatusType) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateOpenbankingBrasilPaymentEnumPaymentStatusTypeEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validates this openbanking brasil payment enum payment status type based on context it is used
func (m OpenbankingBrasilPaymentEnumPaymentStatusType) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}