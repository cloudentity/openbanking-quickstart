// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// BeneficiaryDetails BeneficiaryDetails
//
// swagger:model BeneficiaryDetails
type BeneficiaryDetails struct {

	// Country where the beneficiary resides. A valid [ISO 3166 Alpha-3](https://www.iso.org/iso-3166-country-codes.html) country code
	// Required: true
	Country *string `json:"country"`

	// Response message for the payment
	Message string `json:"message,omitempty"`

	// Name of the beneficiary
	Name string `json:"name,omitempty"`
}

// Validate validates this beneficiary details
func (m *BeneficiaryDetails) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCountry(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *BeneficiaryDetails) validateCountry(formats strfmt.Registry) error {

	if err := validate.Required("country", "body", m.Country); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this beneficiary details based on context it is used
func (m *BeneficiaryDetails) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *BeneficiaryDetails) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *BeneficiaryDetails) UnmarshalBinary(b []byte) error {
	var res BeneficiaryDetails
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
