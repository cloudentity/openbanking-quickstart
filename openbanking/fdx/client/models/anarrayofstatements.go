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
)

// Anarrayofstatements Anarrayofstatements
//
// A paginated array of account statements
//
// swagger:model Anarrayofstatements
type Anarrayofstatements struct {

	// links
	Links *PageMetadataLinks1 `json:"links,omitempty"`

	// page
	Page *PageMetadata1 `json:"page,omitempty"`

	// An array of Statement, each with its HATEOAS link to retrieve the account statement
	Statements []*Statemententity `json:"statements"`
}

// Validate validates this anarrayofstatements
func (m *Anarrayofstatements) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLinks(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePage(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStatements(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Anarrayofstatements) validateLinks(formats strfmt.Registry) error {
	if swag.IsZero(m.Links) { // not required
		return nil
	}

	if m.Links != nil {
		if err := m.Links.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("links")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("links")
			}
			return err
		}
	}

	return nil
}

func (m *Anarrayofstatements) validatePage(formats strfmt.Registry) error {
	if swag.IsZero(m.Page) { // not required
		return nil
	}

	if m.Page != nil {
		if err := m.Page.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("page")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("page")
			}
			return err
		}
	}

	return nil
}

func (m *Anarrayofstatements) validateStatements(formats strfmt.Registry) error {
	if swag.IsZero(m.Statements) { // not required
		return nil
	}

	for i := 0; i < len(m.Statements); i++ {
		if swag.IsZero(m.Statements[i]) { // not required
			continue
		}

		if m.Statements[i] != nil {
			if err := m.Statements[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("statements" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("statements" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this anarrayofstatements based on the context it is used
func (m *Anarrayofstatements) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateLinks(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidatePage(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateStatements(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Anarrayofstatements) contextValidateLinks(ctx context.Context, formats strfmt.Registry) error {

	if m.Links != nil {
		if err := m.Links.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("links")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("links")
			}
			return err
		}
	}

	return nil
}

func (m *Anarrayofstatements) contextValidatePage(ctx context.Context, formats strfmt.Registry) error {

	if m.Page != nil {
		if err := m.Page.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("page")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("page")
			}
			return err
		}
	}

	return nil
}

func (m *Anarrayofstatements) contextValidateStatements(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Statements); i++ {

		if m.Statements[i] != nil {
			if err := m.Statements[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("statements" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("statements" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *Anarrayofstatements) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Anarrayofstatements) UnmarshalBinary(b []byte) error {
	var res Anarrayofstatements
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
