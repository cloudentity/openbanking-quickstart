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

// BankingProductAdditionalInformationV2AdditionalInformationUris BankingProductAdditionalInformationV2_additionalInformationUris
//
// swagger:model BankingProductAdditionalInformationV2_additionalInformationUris
type BankingProductAdditionalInformationV2AdditionalInformationUris struct {

	// The URI describing the additional information
	// Required: true
	AdditionalInfoURI *string `json:"additionalInfoUri"`

	// Display text providing more information about the document URI
	Description string `json:"description,omitempty"`
}

// Validate validates this banking product additional information v2 additional information uris
func (m *BankingProductAdditionalInformationV2AdditionalInformationUris) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAdditionalInfoURI(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *BankingProductAdditionalInformationV2AdditionalInformationUris) validateAdditionalInfoURI(formats strfmt.Registry) error {

	if err := validate.Required("additionalInfoUri", "body", m.AdditionalInfoURI); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this banking product additional information v2 additional information uris based on context it is used
func (m *BankingProductAdditionalInformationV2AdditionalInformationUris) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *BankingProductAdditionalInformationV2AdditionalInformationUris) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *BankingProductAdditionalInformationV2AdditionalInformationUris) UnmarshalBinary(b []byte) error {
	var res BankingProductAdditionalInformationV2AdditionalInformationUris
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
