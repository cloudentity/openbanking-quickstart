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

// DepositRateType DepositRateType
//
// The type of rate (base, bonus, etc). See the next section for an overview of valid values and their meaning
// Example: BONUS
//
// swagger:model DepositRateType
type DepositRateType string

func NewDepositRateType(value DepositRateType) *DepositRateType {
	v := value
	return &v
}

const (

	// DepositRateTypeBONUS captures enum value "BONUS"
	DepositRateTypeBONUS DepositRateType = "BONUS"

	// DepositRateTypeBUNDLEBONUS captures enum value "BUNDLE_BONUS"
	DepositRateTypeBUNDLEBONUS DepositRateType = "BUNDLE_BONUS"

	// DepositRateTypeFIXED captures enum value "FIXED"
	DepositRateTypeFIXED DepositRateType = "FIXED"

	// DepositRateTypeFLOATING captures enum value "FLOATING"
	DepositRateTypeFLOATING DepositRateType = "FLOATING"

	// DepositRateTypeINTRODUCTORY captures enum value "INTRODUCTORY"
	DepositRateTypeINTRODUCTORY DepositRateType = "INTRODUCTORY"

	// DepositRateTypeMARKETLINKED captures enum value "MARKET_LINKED"
	DepositRateTypeMARKETLINKED DepositRateType = "MARKET_LINKED"

	// DepositRateTypeVARIABLE captures enum value "VARIABLE"
	DepositRateTypeVARIABLE DepositRateType = "VARIABLE"
)

// for schema
var depositRateTypeEnum []interface{}

func init() {
	var res []DepositRateType
	if err := json.Unmarshal([]byte(`["BONUS","BUNDLE_BONUS","FIXED","FLOATING","INTRODUCTORY","MARKET_LINKED","VARIABLE"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		depositRateTypeEnum = append(depositRateTypeEnum, v)
	}
}

func (m DepositRateType) validateDepositRateTypeEnum(path, location string, value DepositRateType) error {
	if err := validate.EnumCase(path, location, value, depositRateTypeEnum, true); err != nil {
		return err
	}
	return nil
}

// Validate validates this deposit rate type
func (m DepositRateType) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateDepositRateTypeEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validates this deposit rate type based on context it is used
func (m DepositRateType) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}