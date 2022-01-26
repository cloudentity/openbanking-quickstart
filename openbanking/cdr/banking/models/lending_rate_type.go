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

// LendingRateType LendingRateType
//
// The type of rate (fixed, variable, etc). See the next section for an overview of valid values and their meaning
// Example: BUNDLE_DISCOUNT_FIXED
//
// swagger:model LendingRateType
type LendingRateType string

const (

	// LendingRateTypeBUNDLEDISCOUNTFIXED captures enum value "BUNDLE_DISCOUNT_FIXED"
	LendingRateTypeBUNDLEDISCOUNTFIXED LendingRateType = "BUNDLE_DISCOUNT_FIXED"

	// LendingRateTypeBUNDLEDISCOUNTVARIABLE captures enum value "BUNDLE_DISCOUNT_VARIABLE"
	LendingRateTypeBUNDLEDISCOUNTVARIABLE LendingRateType = "BUNDLE_DISCOUNT_VARIABLE"

	// LendingRateTypeCASHADVANCE captures enum value "CASH_ADVANCE"
	LendingRateTypeCASHADVANCE LendingRateType = "CASH_ADVANCE"

	// LendingRateTypeDISCOUNT captures enum value "DISCOUNT"
	LendingRateTypeDISCOUNT LendingRateType = "DISCOUNT"

	// LendingRateTypeFIXED captures enum value "FIXED"
	LendingRateTypeFIXED LendingRateType = "FIXED"

	// LendingRateTypeFLOATING captures enum value "FLOATING"
	LendingRateTypeFLOATING LendingRateType = "FLOATING"

	// LendingRateTypeINTRODUCTORY captures enum value "INTRODUCTORY"
	LendingRateTypeINTRODUCTORY LendingRateType = "INTRODUCTORY"

	// LendingRateTypeMARKETLINKED captures enum value "MARKET_LINKED"
	LendingRateTypeMARKETLINKED LendingRateType = "MARKET_LINKED"

	// LendingRateTypePENALTY captures enum value "PENALTY"
	LendingRateTypePENALTY LendingRateType = "PENALTY"

	// LendingRateTypePURCHASE captures enum value "PURCHASE"
	LendingRateTypePURCHASE LendingRateType = "PURCHASE"

	// LendingRateTypeVARIABLE captures enum value "VARIABLE"
	LendingRateTypeVARIABLE LendingRateType = "VARIABLE"
)

// for schema
var lendingRateTypeEnum []interface{}

func init() {
	var res []LendingRateType
	if err := json.Unmarshal([]byte(`["BUNDLE_DISCOUNT_FIXED","BUNDLE_DISCOUNT_VARIABLE","CASH_ADVANCE","DISCOUNT","FIXED","FLOATING","INTRODUCTORY","MARKET_LINKED","PENALTY","PURCHASE","VARIABLE"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		lendingRateTypeEnum = append(lendingRateTypeEnum, v)
	}
}

func (m LendingRateType) validateLendingRateTypeEnum(path, location string, value LendingRateType) error {
	if err := validate.EnumCase(path, location, value, lendingRateTypeEnum, true); err != nil {
		return err
	}
	return nil
}

// Validate validates this lending rate type
func (m LendingRateType) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateLendingRateTypeEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validates this lending rate type based on context it is used
func (m LendingRateType) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}
