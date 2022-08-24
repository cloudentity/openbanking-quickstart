// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// HATEOASLink2 HATEOASLink2
//
// Resource URL for retrieving previous dataset
//
// swagger:model HATEOASLink2
type HATEOASLink2 struct {

	// action
	Action Action `json:"action,omitempty"`

	// URL to invoke the action on the resource
	// Example: https://api.fi.com/fdx/v4/accounts/12345
	// Required: true
	Href *string `json:"href"`

	// Relation of this link to its containing entity, as defined by and with many example relation values at [IETF RFC5988](https://datatracker.ietf.org/doc/html/rfc5988)
	Rel string `json:"rel,omitempty"`

	// Content-types that can be used in the Accept header
	Types []ContentTypes `json:"types"`
}

// Validate validates this h a t e o a s link2
func (m *HATEOASLink2) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAction(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateHref(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTypes(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *HATEOASLink2) validateAction(formats strfmt.Registry) error {
	if swag.IsZero(m.Action) { // not required
		return nil
	}

	if err := m.Action.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("action")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("action")
		}
		return err
	}

	return nil
}

func (m *HATEOASLink2) validateHref(formats strfmt.Registry) error {

	if err := validate.Required("href", "body", m.Href); err != nil {
		return err
	}

	return nil
}

func (m *HATEOASLink2) validateTypes(formats strfmt.Registry) error {
	if swag.IsZero(m.Types) { // not required
		return nil
	}

	for i := 0; i < len(m.Types); i++ {

		if err := m.Types[i].Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("types" + "." + strconv.Itoa(i))
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("types" + "." + strconv.Itoa(i))
			}
			return err
		}

	}

	return nil
}

// ContextValidate validate this h a t e o a s link2 based on the context it is used
func (m *HATEOASLink2) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateAction(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateTypes(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *HATEOASLink2) contextValidateAction(ctx context.Context, formats strfmt.Registry) error {

	if err := m.Action.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("action")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("action")
		}
		return err
	}

	return nil
}

func (m *HATEOASLink2) contextValidateTypes(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Types); i++ {

		if err := m.Types[i].ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("types" + "." + strconv.Itoa(i))
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("types" + "." + strconv.Itoa(i))
			}
			return err
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *HATEOASLink2) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *HATEOASLink2) UnmarshalBinary(b []byte) error {
	var res HATEOASLink2
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
