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

// BankingProductBundle BankingProductBundle
//
// swagger:model BankingProductBundle
type BankingProductBundle struct {

	// Display text providing more information on the bundle
	AdditionalInfo string `json:"additionalInfo,omitempty"`

	// Link to a web page with more information on the bundle criteria and benefits
	AdditionalInfoURI string `json:"additionalInfoUri,omitempty"`

	// Description of the bundle
	// Required: true
	Description *string `json:"description"`

	// Name of the bundle
	// Required: true
	Name *string `json:"name"`

	// Array of product IDs for products included in the bundle that are available via the product end points.  Note that this array is not intended to represent a comprehensive model of the products included in the bundle and some products available for the bundle may not be available via the product reference end points
	ProductIds []string `json:"productIds"`
}

// Validate validates this banking product bundle
func (m *BankingProductBundle) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDescription(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *BankingProductBundle) validateDescription(formats strfmt.Registry) error {

	if err := validate.Required("description", "body", m.Description); err != nil {
		return err
	}

	return nil
}

func (m *BankingProductBundle) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this banking product bundle based on context it is used
func (m *BankingProductBundle) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *BankingProductBundle) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *BankingProductBundle) UnmarshalBinary(b []byte) error {
	var res BankingProductBundle
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
