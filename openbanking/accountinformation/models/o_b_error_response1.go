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

// OBErrorResponse1 An array of detail error codes, and messages, and URLs to documentation to help remediation.
//
// swagger:model OBErrorResponse1
type OBErrorResponse1 struct {

	// High level textual error code, to help categorize the errors.
	// Required: true
	// Max Length: 40
	// Min Length: 1
	Code *string `json:"Code"`

	// errors
	// Required: true
	// Min Items: 1
	Errors []*OBError1 `json:"Errors"`

	// A unique reference for the error instance, for audit purposes, in case of unknown/unclassified errors.
	// Max Length: 40
	// Min Length: 1
	ID string `json:"Id,omitempty"`

	// Brief Error message, e.g., 'There is something wrong with the request parameters provided'
	// Required: true
	// Max Length: 500
	// Min Length: 1
	Message *string `json:"Message"`
}

// Validate validates this o b error response1
func (m *OBErrorResponse1) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCode(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateErrors(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMessage(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *OBErrorResponse1) validateCode(formats strfmt.Registry) error {

	if err := validate.Required("Code", "body", m.Code); err != nil {
		return err
	}

	if err := validate.MinLength("Code", "body", *m.Code, 1); err != nil {
		return err
	}

	if err := validate.MaxLength("Code", "body", *m.Code, 40); err != nil {
		return err
	}

	return nil
}

func (m *OBErrorResponse1) validateErrors(formats strfmt.Registry) error {

	if err := validate.Required("Errors", "body", m.Errors); err != nil {
		return err
	}

	iErrorsSize := int64(len(m.Errors))

	if err := validate.MinItems("Errors", "body", iErrorsSize, 1); err != nil {
		return err
	}

	for i := 0; i < len(m.Errors); i++ {
		if swag.IsZero(m.Errors[i]) { // not required
			continue
		}

		if m.Errors[i] != nil {
			if err := m.Errors[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("Errors" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *OBErrorResponse1) validateID(formats strfmt.Registry) error {
	if swag.IsZero(m.ID) { // not required
		return nil
	}

	if err := validate.MinLength("Id", "body", m.ID, 1); err != nil {
		return err
	}

	if err := validate.MaxLength("Id", "body", m.ID, 40); err != nil {
		return err
	}

	return nil
}

func (m *OBErrorResponse1) validateMessage(formats strfmt.Registry) error {

	if err := validate.Required("Message", "body", m.Message); err != nil {
		return err
	}

	if err := validate.MinLength("Message", "body", *m.Message, 1); err != nil {
		return err
	}

	if err := validate.MaxLength("Message", "body", *m.Message, 500); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this o b error response1 based on the context it is used
func (m *OBErrorResponse1) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateErrors(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *OBErrorResponse1) contextValidateErrors(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Errors); i++ {

		if m.Errors[i] != nil {
			if err := m.Errors[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("Errors" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *OBErrorResponse1) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *OBErrorResponse1) UnmarshalBinary(b []byte) error {
	var res OBErrorResponse1
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
