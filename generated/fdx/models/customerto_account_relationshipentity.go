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

// CustomertoAccountRelationshipentity CustomertoAccountRelationshipentity
//
// Describes an account related to a customer
//
// swagger:model CustomertoAccountRelationshipentity
type CustomertoAccountRelationshipentity struct {

	// Account ID of the related account
	// Max Length: 256
	AccountID string `json:"accountId,omitempty"`

	// Links to the account, or to invoke other APIs
	Links []*HATEOASLink `json:"links"`

	// relationship
	Relationship AccountHolderRelationship3 `json:"relationship,omitempty"`
}

// Validate validates this customerto account relationshipentity
func (m *CustomertoAccountRelationshipentity) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAccountID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLinks(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRelationship(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CustomertoAccountRelationshipentity) validateAccountID(formats strfmt.Registry) error {
	if swag.IsZero(m.AccountID) { // not required
		return nil
	}

	if err := validate.MaxLength("accountId", "body", m.AccountID, 256); err != nil {
		return err
	}

	return nil
}

func (m *CustomertoAccountRelationshipentity) validateLinks(formats strfmt.Registry) error {
	if swag.IsZero(m.Links) { // not required
		return nil
	}

	for i := 0; i < len(m.Links); i++ {
		if swag.IsZero(m.Links[i]) { // not required
			continue
		}

		if m.Links[i] != nil {
			if err := m.Links[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("links" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("links" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *CustomertoAccountRelationshipentity) validateRelationship(formats strfmt.Registry) error {
	if swag.IsZero(m.Relationship) { // not required
		return nil
	}

	if err := m.Relationship.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("relationship")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("relationship")
		}
		return err
	}

	return nil
}

// ContextValidate validate this customerto account relationshipentity based on the context it is used
func (m *CustomertoAccountRelationshipentity) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateLinks(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateRelationship(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CustomertoAccountRelationshipentity) contextValidateLinks(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Links); i++ {

		if m.Links[i] != nil {
			if err := m.Links[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("links" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("links" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *CustomertoAccountRelationshipentity) contextValidateRelationship(ctx context.Context, formats strfmt.Registry) error {

	if err := m.Relationship.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("relationship")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("relationship")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *CustomertoAccountRelationshipentity) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CustomertoAccountRelationshipentity) UnmarshalBinary(b []byte) error {
	var res CustomertoAccountRelationshipentity
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}