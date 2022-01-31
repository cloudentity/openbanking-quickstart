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

// DiscountEligibilityType DiscountEligibilityType
//
// The type of the specific eligibility constraint for a discount
// Example: BUSINESS
//
// swagger:model DiscountEligibilityType
type DiscountEligibilityType string

const (

	// DiscountEligibilityTypeBUSINESS captures enum value "BUSINESS"
	DiscountEligibilityTypeBUSINESS DiscountEligibilityType = "BUSINESS"

	// DiscountEligibilityTypeEMPLOYMENTSTATUS captures enum value "EMPLOYMENT_STATUS"
	DiscountEligibilityTypeEMPLOYMENTSTATUS DiscountEligibilityType = "EMPLOYMENT_STATUS"

	// DiscountEligibilityTypeINTRODUCTORY captures enum value "INTRODUCTORY"
	DiscountEligibilityTypeINTRODUCTORY DiscountEligibilityType = "INTRODUCTORY"

	// DiscountEligibilityTypeMAXAGE captures enum value "MAX_AGE"
	DiscountEligibilityTypeMAXAGE DiscountEligibilityType = "MAX_AGE"

	// DiscountEligibilityTypeMINAGE captures enum value "MIN_AGE"
	DiscountEligibilityTypeMINAGE DiscountEligibilityType = "MIN_AGE"

	// DiscountEligibilityTypeMININCOME captures enum value "MIN_INCOME"
	DiscountEligibilityTypeMININCOME DiscountEligibilityType = "MIN_INCOME"

	// DiscountEligibilityTypeMINTURNOVER captures enum value "MIN_TURNOVER"
	DiscountEligibilityTypeMINTURNOVER DiscountEligibilityType = "MIN_TURNOVER"

	// DiscountEligibilityTypeNATURALPERSON captures enum value "NATURAL_PERSON"
	DiscountEligibilityTypeNATURALPERSON DiscountEligibilityType = "NATURAL_PERSON"

	// DiscountEligibilityTypeOTHER captures enum value "OTHER"
	DiscountEligibilityTypeOTHER DiscountEligibilityType = "OTHER"

	// DiscountEligibilityTypePENSIONRECIPIENT captures enum value "PENSION_RECIPIENT"
	DiscountEligibilityTypePENSIONRECIPIENT DiscountEligibilityType = "PENSION_RECIPIENT"

	// DiscountEligibilityTypeRESIDENCYSTATUS captures enum value "RESIDENCY_STATUS"
	DiscountEligibilityTypeRESIDENCYSTATUS DiscountEligibilityType = "RESIDENCY_STATUS"

	// DiscountEligibilityTypeSTAFF captures enum value "STAFF"
	DiscountEligibilityTypeSTAFF DiscountEligibilityType = "STAFF"

	// DiscountEligibilityTypeSTUDENT captures enum value "STUDENT"
	DiscountEligibilityTypeSTUDENT DiscountEligibilityType = "STUDENT"
)

// for schema
var discountEligibilityTypeEnum []interface{}

func init() {
	var res []DiscountEligibilityType
	if err := json.Unmarshal([]byte(`["BUSINESS","EMPLOYMENT_STATUS","INTRODUCTORY","MAX_AGE","MIN_AGE","MIN_INCOME","MIN_TURNOVER","NATURAL_PERSON","OTHER","PENSION_RECIPIENT","RESIDENCY_STATUS","STAFF","STUDENT"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		discountEligibilityTypeEnum = append(discountEligibilityTypeEnum, v)
	}
}

func (m DiscountEligibilityType) validateDiscountEligibilityTypeEnum(path, location string, value DiscountEligibilityType) error {
	if err := validate.EnumCase(path, location, value, discountEligibilityTypeEnum, true); err != nil {
		return err
	}
	return nil
}

// Validate validates this discount eligibility type
func (m DiscountEligibilityType) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateDiscountEligibilityTypeEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validates this discount eligibility type based on context it is used
func (m DiscountEligibilityType) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}
