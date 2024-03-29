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

// EligibilityType EligibilityType
//
// The type of eligibility criteria described.  See the next section for an overview of valid values and their meaning
// Example: BUSINESS
//
// swagger:model EligibilityType
type EligibilityType string

func NewEligibilityType(value EligibilityType) *EligibilityType {
	v := value
	return &v
}

const (

	// EligibilityTypeBUSINESS captures enum value "BUSINESS"
	EligibilityTypeBUSINESS EligibilityType = "BUSINESS"

	// EligibilityTypeEMPLOYMENTSTATUS captures enum value "EMPLOYMENT_STATUS"
	EligibilityTypeEMPLOYMENTSTATUS EligibilityType = "EMPLOYMENT_STATUS"

	// EligibilityTypeMAXAGE captures enum value "MAX_AGE"
	EligibilityTypeMAXAGE EligibilityType = "MAX_AGE"

	// EligibilityTypeMINAGE captures enum value "MIN_AGE"
	EligibilityTypeMINAGE EligibilityType = "MIN_AGE"

	// EligibilityTypeMININCOME captures enum value "MIN_INCOME"
	EligibilityTypeMININCOME EligibilityType = "MIN_INCOME"

	// EligibilityTypeMINTURNOVER captures enum value "MIN_TURNOVER"
	EligibilityTypeMINTURNOVER EligibilityType = "MIN_TURNOVER"

	// EligibilityTypeNATURALPERSON captures enum value "NATURAL_PERSON"
	EligibilityTypeNATURALPERSON EligibilityType = "NATURAL_PERSON"

	// EligibilityTypeOTHER captures enum value "OTHER"
	EligibilityTypeOTHER EligibilityType = "OTHER"

	// EligibilityTypePENSIONRECIPIENT captures enum value "PENSION_RECIPIENT"
	EligibilityTypePENSIONRECIPIENT EligibilityType = "PENSION_RECIPIENT"

	// EligibilityTypeRESIDENCYSTATUS captures enum value "RESIDENCY_STATUS"
	EligibilityTypeRESIDENCYSTATUS EligibilityType = "RESIDENCY_STATUS"

	// EligibilityTypeSTAFF captures enum value "STAFF"
	EligibilityTypeSTAFF EligibilityType = "STAFF"

	// EligibilityTypeSTUDENT captures enum value "STUDENT"
	EligibilityTypeSTUDENT EligibilityType = "STUDENT"
)

// for schema
var eligibilityTypeEnum []interface{}

func init() {
	var res []EligibilityType
	if err := json.Unmarshal([]byte(`["BUSINESS","EMPLOYMENT_STATUS","MAX_AGE","MIN_AGE","MIN_INCOME","MIN_TURNOVER","NATURAL_PERSON","OTHER","PENSION_RECIPIENT","RESIDENCY_STATUS","STAFF","STUDENT"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		eligibilityTypeEnum = append(eligibilityTypeEnum, v)
	}
}

func (m EligibilityType) validateEligibilityTypeEnum(path, location string, value EligibilityType) error {
	if err := validate.EnumCase(path, location, value, eligibilityTypeEnum, true); err != nil {
		return err
	}
	return nil
}

// Validate validates this eligibility type
func (m EligibilityType) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateEligibilityTypeEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validates this eligibility type based on context it is used
func (m EligibilityType) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}
