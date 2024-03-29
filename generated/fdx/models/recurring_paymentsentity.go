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

// RecurringPaymentsentity RecurringPaymentsentity
//
// A list of recurring payments
//
// swagger:model RecurringPaymentsentity
type RecurringPaymentsentity struct {

	// links
	Links *SynchronizableArrayLinksentity2 `json:"links,omitempty"`

	// page
	Page *PageMetadata2 `json:"page,omitempty"`

	// Recurring payments retrieved by the operation
	// Required: true
	RecurringPayments []*RecurringPaymententity `json:"recurringPayments"`

	// updates
	Updates *UpdatesMetadataentity2 `json:"updates,omitempty"`
}

// Validate validates this recurring paymentsentity
func (m *RecurringPaymentsentity) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLinks(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePage(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRecurringPayments(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUpdates(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *RecurringPaymentsentity) validateLinks(formats strfmt.Registry) error {
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

func (m *RecurringPaymentsentity) validatePage(formats strfmt.Registry) error {
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

func (m *RecurringPaymentsentity) validateRecurringPayments(formats strfmt.Registry) error {

	if err := validate.Required("recurringPayments", "body", m.RecurringPayments); err != nil {
		return err
	}

	for i := 0; i < len(m.RecurringPayments); i++ {
		if swag.IsZero(m.RecurringPayments[i]) { // not required
			continue
		}

		if m.RecurringPayments[i] != nil {
			if err := m.RecurringPayments[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("recurringPayments" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("recurringPayments" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *RecurringPaymentsentity) validateUpdates(formats strfmt.Registry) error {
	if swag.IsZero(m.Updates) { // not required
		return nil
	}

	if m.Updates != nil {
		if err := m.Updates.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("updates")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("updates")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this recurring paymentsentity based on the context it is used
func (m *RecurringPaymentsentity) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateLinks(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidatePage(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateRecurringPayments(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateUpdates(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *RecurringPaymentsentity) contextValidateLinks(ctx context.Context, formats strfmt.Registry) error {

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

func (m *RecurringPaymentsentity) contextValidatePage(ctx context.Context, formats strfmt.Registry) error {

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

func (m *RecurringPaymentsentity) contextValidateRecurringPayments(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.RecurringPayments); i++ {

		if m.RecurringPayments[i] != nil {
			if err := m.RecurringPayments[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("recurringPayments" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("recurringPayments" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *RecurringPaymentsentity) contextValidateUpdates(ctx context.Context, formats strfmt.Registry) error {

	if m.Updates != nil {
		if err := m.Updates.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("updates")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("updates")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *RecurringPaymentsentity) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *RecurringPaymentsentity) UnmarshalBinary(b []byte) error {
	var res RecurringPaymentsentity
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
