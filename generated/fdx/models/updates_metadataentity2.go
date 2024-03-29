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

// UpdatesMetadataentity2 UpdatesMetadataentity2
//
// Update IDs for retrieving updates since query
//
// swagger:model UpdatesMetadataentity2
type UpdatesMetadataentity2 struct {

	// Opaque identifier. Does not need to be numeric or have any specific pattern. Implementation specific
	// Max Length: 256
	NextUpdateID string `json:"nextUpdateId,omitempty"`
}

// Validate validates this updates metadataentity2
func (m *UpdatesMetadataentity2) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateNextUpdateID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *UpdatesMetadataentity2) validateNextUpdateID(formats strfmt.Registry) error {
	if swag.IsZero(m.NextUpdateID) { // not required
		return nil
	}

	if err := validate.MaxLength("nextUpdateId", "body", m.NextUpdateID, 256); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this updates metadataentity2 based on context it is used
func (m *UpdatesMetadataentity2) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *UpdatesMetadataentity2) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *UpdatesMetadataentity2) UnmarshalBinary(b []byte) error {
	var res UpdatesMetadataentity2
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
