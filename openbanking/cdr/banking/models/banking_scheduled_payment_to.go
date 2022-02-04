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

// BankingScheduledPaymentTo BankingScheduledPaymentTo
//
// Object containing details of the destination of the payment. Used to specify a variety of payment destination types
//
// swagger:model BankingScheduledPaymentTo
type BankingScheduledPaymentTo struct {

	// Present if toUType is set to accountId. Indicates that the payment is to another account that is accessible under the current consent
	AccountID string `json:"accountId,omitempty"`

	// biller
	Biller *BankingBillerPayee `json:"biller,omitempty"`

	// domestic
	Domestic *BankingDomesticPayee `json:"domestic,omitempty"`

	// international
	International *BankingInternationalPayee `json:"international,omitempty"`

	// The short display name of the payee as provided by the customer unless toUType is set to payeeId. Where a customer has not provided a nickname, a display name derived by the bank for payee should be provided that is consistent with existing digital banking channels
	Nickname string `json:"nickname,omitempty"`

	// Present if toUType is set to payeeId. Indicates that the payment is to registered payee that can be accessed using the payee end point. If the Bank Payees scope has not been consented to then a payeeId should not be provided and the full payee details should be provided instead
	PayeeID string `json:"payeeId,omitempty"`

	// The reference for the transaction, if applicable, that will be provided by the originating institution for the specific payment. If not empty, it overrides the value provided at the BankingScheduledPayment level.
	PayeeReference string `json:"payeeReference,omitempty"`

	// to u type
	// Required: true
	ToUType *ToUType `json:"toUType"`
}

// Validate validates this banking scheduled payment to
func (m *BankingScheduledPaymentTo) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateBiller(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDomestic(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateInternational(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateToUType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *BankingScheduledPaymentTo) validateBiller(formats strfmt.Registry) error {
	if swag.IsZero(m.Biller) { // not required
		return nil
	}

	if m.Biller != nil {
		if err := m.Biller.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("biller")
			}
			return err
		}
	}

	return nil
}

func (m *BankingScheduledPaymentTo) validateDomestic(formats strfmt.Registry) error {
	if swag.IsZero(m.Domestic) { // not required
		return nil
	}

	if m.Domestic != nil {
		if err := m.Domestic.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("domestic")
			}
			return err
		}
	}

	return nil
}

func (m *BankingScheduledPaymentTo) validateInternational(formats strfmt.Registry) error {
	if swag.IsZero(m.International) { // not required
		return nil
	}

	if m.International != nil {
		if err := m.International.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("international")
			}
			return err
		}
	}

	return nil
}

func (m *BankingScheduledPaymentTo) validateToUType(formats strfmt.Registry) error {

	if err := validate.Required("toUType", "body", m.ToUType); err != nil {
		return err
	}

	if err := validate.Required("toUType", "body", m.ToUType); err != nil {
		return err
	}

	if m.ToUType != nil {
		if err := m.ToUType.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("toUType")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this banking scheduled payment to based on the context it is used
func (m *BankingScheduledPaymentTo) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateBiller(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateDomestic(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateInternational(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateToUType(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *BankingScheduledPaymentTo) contextValidateBiller(ctx context.Context, formats strfmt.Registry) error {

	if m.Biller != nil {
		if err := m.Biller.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("biller")
			}
			return err
		}
	}

	return nil
}

func (m *BankingScheduledPaymentTo) contextValidateDomestic(ctx context.Context, formats strfmt.Registry) error {

	if m.Domestic != nil {
		if err := m.Domestic.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("domestic")
			}
			return err
		}
	}

	return nil
}

func (m *BankingScheduledPaymentTo) contextValidateInternational(ctx context.Context, formats strfmt.Registry) error {

	if m.International != nil {
		if err := m.International.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("international")
			}
			return err
		}
	}

	return nil
}

func (m *BankingScheduledPaymentTo) contextValidateToUType(ctx context.Context, formats strfmt.Registry) error {

	if m.ToUType != nil {
		if err := m.ToUType.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("toUType")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *BankingScheduledPaymentTo) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *BankingScheduledPaymentTo) UnmarshalBinary(b []byte) error {
	var res BankingScheduledPaymentTo
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}