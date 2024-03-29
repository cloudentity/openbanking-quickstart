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

// EnumAccountSubType EnumAccountSubType
//
// Subtipo de conta (vide Enum):
// Conta individual - possui um único titular
// Conta conjunta simples - onde as movimentações financeiras só podem serem realizadas mediante autorização de TODOS os correntistas da conta.
// Conta conjunta solidária - é a modalidade cujos titulares podem realizar movimentações de forma isolada, isto é, sem que seja necessária a autorização dos demais titulares
//
// swagger:model EnumAccountSubType
type EnumAccountSubType string

func NewEnumAccountSubType(value EnumAccountSubType) *EnumAccountSubType {
	v := value
	return &v
}

const (

	// EnumAccountSubTypeINDIVIDUAL captures enum value "INDIVIDUAL"
	EnumAccountSubTypeINDIVIDUAL EnumAccountSubType = "INDIVIDUAL"

	// EnumAccountSubTypeCONJUNTASIMPLES captures enum value "CONJUNTA_SIMPLES"
	EnumAccountSubTypeCONJUNTASIMPLES EnumAccountSubType = "CONJUNTA_SIMPLES"

	// EnumAccountSubTypeCONJUNTASOLIDARIA captures enum value "CONJUNTA_SOLIDARIA"
	EnumAccountSubTypeCONJUNTASOLIDARIA EnumAccountSubType = "CONJUNTA_SOLIDARIA"
)

// for schema
var enumAccountSubTypeEnum []interface{}

func init() {
	var res []EnumAccountSubType
	if err := json.Unmarshal([]byte(`["INDIVIDUAL","CONJUNTA_SIMPLES","CONJUNTA_SOLIDARIA"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		enumAccountSubTypeEnum = append(enumAccountSubTypeEnum, v)
	}
}

func (m EnumAccountSubType) validateEnumAccountSubTypeEnum(path, location string, value EnumAccountSubType) error {
	if err := validate.EnumCase(path, location, value, enumAccountSubTypeEnum, true); err != nil {
		return err
	}
	return nil
}

// Validate validates this enum account sub type
func (m EnumAccountSubType) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateEnumAccountSubTypeEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validates this enum account sub type based on context it is used
func (m EnumAccountSubType) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}
