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

// DeliveryAddress2 DeliveryAddress2
//
// Address of the payee used to execute the payment
//
// swagger:model DeliveryAddress2
type DeliveryAddress2 struct {

	// City
	// Max Length: 64
	City string `json:"city,omitempty"`

	// country
	Country ISO3166CountryCode3 `json:"country,omitempty"`

	// Address line 1
	// Max Length: 64
	Line1 string `json:"line1,omitempty"`

	// Address line 2
	// Max Length: 64
	Line2 string `json:"line2,omitempty"`

	// Address line 3
	// Max Length: 64
	Line3 string `json:"line3,omitempty"`

	// Postal code
	// Max Length: 16
	PostalCode string `json:"postalCode,omitempty"`

	// State or province or territory. Replaces "state" property
	// Max Length: 64
	Region string `json:"region,omitempty"`

	// State or province. Deprecated, will remove in FDX V6.0
	// Max Length: 64
	State string `json:"state,omitempty"`

	// type
	Type DeliveryAddressType2 `json:"type,omitempty"`
}

// Validate validates this delivery address2
func (m *DeliveryAddress2) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCity(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCountry(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLine1(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLine2(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLine3(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePostalCode(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRegion(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateState(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DeliveryAddress2) validateCity(formats strfmt.Registry) error {
	if swag.IsZero(m.City) { // not required
		return nil
	}

	if err := validate.MaxLength("city", "body", m.City, 64); err != nil {
		return err
	}

	return nil
}

func (m *DeliveryAddress2) validateCountry(formats strfmt.Registry) error {
	if swag.IsZero(m.Country) { // not required
		return nil
	}

	if err := m.Country.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("country")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("country")
		}
		return err
	}

	return nil
}

func (m *DeliveryAddress2) validateLine1(formats strfmt.Registry) error {
	if swag.IsZero(m.Line1) { // not required
		return nil
	}

	if err := validate.MaxLength("line1", "body", m.Line1, 64); err != nil {
		return err
	}

	return nil
}

func (m *DeliveryAddress2) validateLine2(formats strfmt.Registry) error {
	if swag.IsZero(m.Line2) { // not required
		return nil
	}

	if err := validate.MaxLength("line2", "body", m.Line2, 64); err != nil {
		return err
	}

	return nil
}

func (m *DeliveryAddress2) validateLine3(formats strfmt.Registry) error {
	if swag.IsZero(m.Line3) { // not required
		return nil
	}

	if err := validate.MaxLength("line3", "body", m.Line3, 64); err != nil {
		return err
	}

	return nil
}

func (m *DeliveryAddress2) validatePostalCode(formats strfmt.Registry) error {
	if swag.IsZero(m.PostalCode) { // not required
		return nil
	}

	if err := validate.MaxLength("postalCode", "body", m.PostalCode, 16); err != nil {
		return err
	}

	return nil
}

func (m *DeliveryAddress2) validateRegion(formats strfmt.Registry) error {
	if swag.IsZero(m.Region) { // not required
		return nil
	}

	if err := validate.MaxLength("region", "body", m.Region, 64); err != nil {
		return err
	}

	return nil
}

func (m *DeliveryAddress2) validateState(formats strfmt.Registry) error {
	if swag.IsZero(m.State) { // not required
		return nil
	}

	if err := validate.MaxLength("state", "body", m.State, 64); err != nil {
		return err
	}

	return nil
}

func (m *DeliveryAddress2) validateType(formats strfmt.Registry) error {
	if swag.IsZero(m.Type) { // not required
		return nil
	}

	if err := m.Type.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("type")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("type")
		}
		return err
	}

	return nil
}

// ContextValidate validate this delivery address2 based on the context it is used
func (m *DeliveryAddress2) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateCountry(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateType(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DeliveryAddress2) contextValidateCountry(ctx context.Context, formats strfmt.Registry) error {

	if err := m.Country.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("country")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("country")
		}
		return err
	}

	return nil
}

func (m *DeliveryAddress2) contextValidateType(ctx context.Context, formats strfmt.Registry) error {

	if err := m.Type.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("type")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("type")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *DeliveryAddress2) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DeliveryAddress2) UnmarshalBinary(b []byte) error {
	var res DeliveryAddress2
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}