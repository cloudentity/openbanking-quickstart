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

// OpenbankingBrasilConsentV2Data Data
//
// swagger:model OpenbankingBrasilConsentV2Data
type OpenbankingBrasilConsentV2Data struct {

	// business entity
	BusinessEntity *OpenbankingBrasilConsentV2BusinessEntity `json:"businessEntity,omitempty"`

	// Data e hora de expirao da permisso. De preenchimento obrigatrio, reflete a data limite de validade do consentimento. Uma string com data e hora conforme especificao RFC-3339, sempre com a utilizao de timezone UTC(UTC time format).
	// Example: 2021-05-21T08:30:00Z
	// Required: true
	// Format: date-time
	ExpirationDateTime strfmt.DateTime `json:"expirationDateTime"`

	// logged user
	// Required: true
	LoggedUser *OpenbankingBrasilConsentV2LoggedUser `json:"loggedUser"`

	// permissions
	// Example: ["ACCOUNTS_READ","ACCOUNTS_OVERDRAFT_LIMITS_READ","RESOURCES_READ"]
	// Required: true
	// Max Items: 30
	// Min Items: 1
	Permissions []OpenbankingBrasilConsentV2Permission `json:"permissions"`
}

// Validate validates this openbanking brasil consent v2 data
func (m *OpenbankingBrasilConsentV2Data) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateBusinessEntity(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateExpirationDateTime(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLoggedUser(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePermissions(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *OpenbankingBrasilConsentV2Data) validateBusinessEntity(formats strfmt.Registry) error {
	if swag.IsZero(m.BusinessEntity) { // not required
		return nil
	}

	if m.BusinessEntity != nil {
		if err := m.BusinessEntity.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("businessEntity")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("businessEntity")
			}
			return err
		}
	}

	return nil
}

func (m *OpenbankingBrasilConsentV2Data) validateExpirationDateTime(formats strfmt.Registry) error {

	if err := validate.Required("expirationDateTime", "body", strfmt.DateTime(m.ExpirationDateTime)); err != nil {
		return err
	}

	if err := validate.FormatOf("expirationDateTime", "body", "date-time", m.ExpirationDateTime.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *OpenbankingBrasilConsentV2Data) validateLoggedUser(formats strfmt.Registry) error {

	if err := validate.Required("loggedUser", "body", m.LoggedUser); err != nil {
		return err
	}

	if m.LoggedUser != nil {
		if err := m.LoggedUser.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("loggedUser")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("loggedUser")
			}
			return err
		}
	}

	return nil
}

func (m *OpenbankingBrasilConsentV2Data) validatePermissions(formats strfmt.Registry) error {

	if err := validate.Required("permissions", "body", m.Permissions); err != nil {
		return err
	}

	iPermissionsSize := int64(len(m.Permissions))

	if err := validate.MinItems("permissions", "body", iPermissionsSize, 1); err != nil {
		return err
	}

	if err := validate.MaxItems("permissions", "body", iPermissionsSize, 30); err != nil {
		return err
	}

	for i := 0; i < len(m.Permissions); i++ {

		if err := m.Permissions[i].Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("permissions" + "." + strconv.Itoa(i))
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("permissions" + "." + strconv.Itoa(i))
			}
			return err
		}

	}

	return nil
}

// ContextValidate validate this openbanking brasil consent v2 data based on the context it is used
func (m *OpenbankingBrasilConsentV2Data) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateBusinessEntity(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateLoggedUser(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidatePermissions(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *OpenbankingBrasilConsentV2Data) contextValidateBusinessEntity(ctx context.Context, formats strfmt.Registry) error {

	if m.BusinessEntity != nil {
		if err := m.BusinessEntity.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("businessEntity")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("businessEntity")
			}
			return err
		}
	}

	return nil
}

func (m *OpenbankingBrasilConsentV2Data) contextValidateLoggedUser(ctx context.Context, formats strfmt.Registry) error {

	if m.LoggedUser != nil {
		if err := m.LoggedUser.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("loggedUser")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("loggedUser")
			}
			return err
		}
	}

	return nil
}

func (m *OpenbankingBrasilConsentV2Data) contextValidatePermissions(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Permissions); i++ {

		if err := m.Permissions[i].ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("permissions" + "." + strconv.Itoa(i))
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("permissions" + "." + strconv.Itoa(i))
			}
			return err
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *OpenbankingBrasilConsentV2Data) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *OpenbankingBrasilConsentV2Data) UnmarshalBinary(b []byte) error {
	var res OpenbankingBrasilConsentV2Data
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
