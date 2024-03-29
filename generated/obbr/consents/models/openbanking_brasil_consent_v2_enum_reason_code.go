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

// OpenbankingBrasilConsentV2EnumReasonCode EnumReasonCode
//
// Define o cdigo da razo pela qual o consentimento foi rejeitado.
//
// - CONSENT_EXPIRED  consentimento que ultrapassou o tempo limite para autorizao.
// - CUSTOMER_MANUALLY_REJECTED  cliente efetuou a rejeio do consentimento manualmente atravs de interao nas instituies participantes.
// - CUSTOMER_MANUALLY_REVOKED  cliente efetuou a revogao aps a autorizao do consentimento.
// - CONSENT_MAX_DATE_REACHED  consentimento que ultrapassou o tempo limite de compartilhamento.
// - CONSENT_TECHNICAL_ISSUE  consentimento que foi rejeitado devido a um problema tcnico que impossibilita seu uso pela instituio receptora, por exemplo: falha associada a troca do AuthCode pelo AccessToken, durante o processo de Hybid Flow.
// - INTERNAL_SECURITY_REASON  consentimento que foi rejeitado devido as polticas de segurana aplicada pela instituio transmissora.
// Example: CONSENT_EXPIRED
//
// swagger:model OpenbankingBrasilConsentV2EnumReasonCode
type OpenbankingBrasilConsentV2EnumReasonCode string

func NewOpenbankingBrasilConsentV2EnumReasonCode(value OpenbankingBrasilConsentV2EnumReasonCode) *OpenbankingBrasilConsentV2EnumReasonCode {
	v := value
	return &v
}

const (

	// OpenbankingBrasilConsentV2EnumReasonCodeCONSENTEXPIRED captures enum value "CONSENT_EXPIRED"
	OpenbankingBrasilConsentV2EnumReasonCodeCONSENTEXPIRED OpenbankingBrasilConsentV2EnumReasonCode = "CONSENT_EXPIRED"

	// OpenbankingBrasilConsentV2EnumReasonCodeCUSTOMERMANUALLYREJECTED captures enum value "CUSTOMER_MANUALLY_REJECTED"
	OpenbankingBrasilConsentV2EnumReasonCodeCUSTOMERMANUALLYREJECTED OpenbankingBrasilConsentV2EnumReasonCode = "CUSTOMER_MANUALLY_REJECTED"

	// OpenbankingBrasilConsentV2EnumReasonCodeCUSTOMERMANUALLYREVOKED captures enum value "CUSTOMER_MANUALLY_REVOKED"
	OpenbankingBrasilConsentV2EnumReasonCodeCUSTOMERMANUALLYREVOKED OpenbankingBrasilConsentV2EnumReasonCode = "CUSTOMER_MANUALLY_REVOKED"

	// OpenbankingBrasilConsentV2EnumReasonCodeCONSENTMAXDATEREACHED captures enum value "CONSENT_MAX_DATE_REACHED"
	OpenbankingBrasilConsentV2EnumReasonCodeCONSENTMAXDATEREACHED OpenbankingBrasilConsentV2EnumReasonCode = "CONSENT_MAX_DATE_REACHED"

	// OpenbankingBrasilConsentV2EnumReasonCodeCONSENTTECHNICALISSUE captures enum value "CONSENT_TECHNICAL_ISSUE"
	OpenbankingBrasilConsentV2EnumReasonCodeCONSENTTECHNICALISSUE OpenbankingBrasilConsentV2EnumReasonCode = "CONSENT_TECHNICAL_ISSUE"

	// OpenbankingBrasilConsentV2EnumReasonCodeINTERNALSECURITYREASON captures enum value "INTERNAL_SECURITY_REASON"
	OpenbankingBrasilConsentV2EnumReasonCodeINTERNALSECURITYREASON OpenbankingBrasilConsentV2EnumReasonCode = "INTERNAL_SECURITY_REASON"
)

// for schema
var openbankingBrasilConsentV2EnumReasonCodeEnum []interface{}

func init() {
	var res []OpenbankingBrasilConsentV2EnumReasonCode
	if err := json.Unmarshal([]byte(`["CONSENT_EXPIRED","CUSTOMER_MANUALLY_REJECTED","CUSTOMER_MANUALLY_REVOKED","CONSENT_MAX_DATE_REACHED","CONSENT_TECHNICAL_ISSUE","INTERNAL_SECURITY_REASON"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		openbankingBrasilConsentV2EnumReasonCodeEnum = append(openbankingBrasilConsentV2EnumReasonCodeEnum, v)
	}
}

func (m OpenbankingBrasilConsentV2EnumReasonCode) validateOpenbankingBrasilConsentV2EnumReasonCodeEnum(path, location string, value OpenbankingBrasilConsentV2EnumReasonCode) error {
	if err := validate.EnumCase(path, location, value, openbankingBrasilConsentV2EnumReasonCodeEnum, true); err != nil {
		return err
	}
	return nil
}

// Validate validates this openbanking brasil consent v2 enum reason code
func (m OpenbankingBrasilConsentV2EnumReasonCode) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateOpenbankingBrasilConsentV2EnumReasonCodeEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validates this openbanking brasil consent v2 enum reason code based on context it is used
func (m OpenbankingBrasilConsentV2EnumReasonCode) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}
