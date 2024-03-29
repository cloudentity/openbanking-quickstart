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

// OpenbankingBrasilConsentV2EnumRejectedBy EnumRejectedBy
//
// Informar usurio responsvel pela rejeio.
// 1. USER usurio
// 2. ASPSP instituio transmissora
// 3. TPP instituio receptora
// Example: USER
//
// swagger:model OpenbankingBrasilConsentV2EnumRejectedBy
type OpenbankingBrasilConsentV2EnumRejectedBy string

func NewOpenbankingBrasilConsentV2EnumRejectedBy(value OpenbankingBrasilConsentV2EnumRejectedBy) *OpenbankingBrasilConsentV2EnumRejectedBy {
	v := value
	return &v
}

const (

	// OpenbankingBrasilConsentV2EnumRejectedByUSER captures enum value "USER"
	OpenbankingBrasilConsentV2EnumRejectedByUSER OpenbankingBrasilConsentV2EnumRejectedBy = "USER"

	// OpenbankingBrasilConsentV2EnumRejectedByASPSP captures enum value "ASPSP"
	OpenbankingBrasilConsentV2EnumRejectedByASPSP OpenbankingBrasilConsentV2EnumRejectedBy = "ASPSP"

	// OpenbankingBrasilConsentV2EnumRejectedByTPP captures enum value "TPP"
	OpenbankingBrasilConsentV2EnumRejectedByTPP OpenbankingBrasilConsentV2EnumRejectedBy = "TPP"
)

// for schema
var openbankingBrasilConsentV2EnumRejectedByEnum []interface{}

func init() {
	var res []OpenbankingBrasilConsentV2EnumRejectedBy
	if err := json.Unmarshal([]byte(`["USER","ASPSP","TPP"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		openbankingBrasilConsentV2EnumRejectedByEnum = append(openbankingBrasilConsentV2EnumRejectedByEnum, v)
	}
}

func (m OpenbankingBrasilConsentV2EnumRejectedBy) validateOpenbankingBrasilConsentV2EnumRejectedByEnum(path, location string, value OpenbankingBrasilConsentV2EnumRejectedBy) error {
	if err := validate.EnumCase(path, location, value, openbankingBrasilConsentV2EnumRejectedByEnum, true); err != nil {
		return err
	}
	return nil
}

// Validate validates this openbanking brasil consent v2 enum rejected by
func (m OpenbankingBrasilConsentV2EnumRejectedBy) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateOpenbankingBrasilConsentV2EnumRejectedByEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validates this openbanking brasil consent v2 enum rejected by based on context it is used
func (m OpenbankingBrasilConsentV2EnumRejectedBy) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}
