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

// OpenbankingBrasilPaymentConsent PaymentConsent
//
// Objeto contendo dados de pagamento para consentimento.
//
// swagger:model OpenbankingBrasilPaymentConsent
type OpenbankingBrasilPaymentConsent struct {

	// Valor da transao com 2 casas decimais.
	// Example: 100000.12
	// Required: true
	// Max Length: 19
	// Min Length: 4
	// Pattern: ^((\d{1,16}\.\d{2}))$
	Amount string `json:"amount"`

	// Cdigo da moeda nacional segundo modelo ISO-4217, ou seja, 'BRL'.
	// Todos os valores monetrios informados esto representados com a moeda vigente do Brasil.
	// Example: BRL
	// Required: true
	// Max Length: 3
	// Pattern: ^([A-Z]{3})$
	Currency string `json:"currency"`

	// Data do pagamento, conforme especificao RFC-3339.
	// Example: 2021-01-01
	// Required: true
	// Format: date
	Date strfmt.Date `json:"date"`

	// Este campo define o tipo de pagamento que ser iniciado aps a autorizao do consentimento.
	// Example: PIX
	// Required: true
	Type string `json:"type"`
}

// Validate validates this openbanking brasil payment consent
func (m *OpenbankingBrasilPaymentConsent) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAmount(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCurrency(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDate(formats); err != nil {
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

func (m *OpenbankingBrasilPaymentConsent) validateAmount(formats strfmt.Registry) error {

	if err := validate.RequiredString("amount", "body", m.Amount); err != nil {
		return err
	}

	if err := validate.MinLength("amount", "body", m.Amount, 4); err != nil {
		return err
	}

	if err := validate.MaxLength("amount", "body", m.Amount, 19); err != nil {
		return err
	}

	if err := validate.Pattern("amount", "body", m.Amount, `^((\d{1,16}\.\d{2}))$`); err != nil {
		return err
	}

	return nil
}

func (m *OpenbankingBrasilPaymentConsent) validateCurrency(formats strfmt.Registry) error {

	if err := validate.RequiredString("currency", "body", m.Currency); err != nil {
		return err
	}

	if err := validate.MaxLength("currency", "body", m.Currency, 3); err != nil {
		return err
	}

	if err := validate.Pattern("currency", "body", m.Currency, `^([A-Z]{3})$`); err != nil {
		return err
	}

	return nil
}

func (m *OpenbankingBrasilPaymentConsent) validateDate(formats strfmt.Registry) error {

	if err := validate.Required("date", "body", strfmt.Date(m.Date)); err != nil {
		return err
	}

	if err := validate.FormatOf("date", "body", "date", m.Date.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *OpenbankingBrasilPaymentConsent) validateType(formats strfmt.Registry) error {

	if err := validate.RequiredString("type", "body", m.Type); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this openbanking brasil payment consent based on context it is used
func (m *OpenbankingBrasilPaymentConsent) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *OpenbankingBrasilPaymentConsent) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *OpenbankingBrasilPaymentConsent) UnmarshalBinary(b []byte) error {
	var res OpenbankingBrasilPaymentConsent
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
