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

// AccountOverdraftLimitsData AccountOverdraftLimitsData
//
// Conjunto de informações da Conta de: depósito à vista
//
// swagger:model AccountOverdraftLimitsData
type AccountOverdraftLimitsData struct {

	// Valor do limite contratado do cheque especial.
	// Example: 99.9999
	// Required: true
	OverdraftContractedLimit *float64 `json:"overdraftContractedLimit"`

	// Moeda referente ao valor do limite contratado do cheque especial, segundo modelo ISO-4217. p.ex. 'BRL'. Pode ser preenchido com “NA” caso a instituição não possua a informação.
	// Example: BRL
	// Required: true
	// Max Length: 3
	// Pattern: ^(\w{3}){1}$
	OverdraftContractedLimitCurrency *string `json:"overdraftContractedLimitCurrency"`

	// Valor utilizado total do limite do cheque especial e o adiantamento a depositante.
	// Example: 10000.9999
	// Required: true
	OverdraftUsedLimit *float64 `json:"overdraftUsedLimit"`

	// Moeda referente ao valor utilizado total do limite do cheque especial e o adiantamento a depositante, segundo modelo ISO-4217. p.ex. 'BRL'. Pode ser preenchido com “NA” caso a instituição não possua a informação.
	// Example: BRL
	// Required: true
	// Max Length: 3
	// Pattern: ^(\w{3}){1}$
	OverdraftUsedLimitCurrency *string `json:"overdraftUsedLimitCurrency"`

	// Valor de operação contratada em caráter emergencial para cobertura de saldo devedor em conta de depósitos à vista e de excesso sobre o limite pactuado de cheque especial.
	// Example: 99.9999
	// Required: true
	UnarrangedOverdraftAmount *float64 `json:"unarrangedOverdraftAmount"`

	// Moeda referente ao valor de operação contratada em caráter emergencial para cobertura de saldo devedor em conta de depósitos à vista e de excesso sobre o limite pactuado de cheque especial, segundo modelo ISO-4217. p.ex. 'BRL'. Pode ser preenchido com “NA” caso a instituição não possua a informação.
	// Example: BRL
	// Required: true
	// Max Length: 3
	// Pattern: ^(\w{3}){1}$
	UnarrangedOverdraftAmountCurrency *string `json:"unarrangedOverdraftAmountCurrency"`
}

// Validate validates this account overdraft limits data
func (m *AccountOverdraftLimitsData) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateOverdraftContractedLimit(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOverdraftContractedLimitCurrency(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOverdraftUsedLimit(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOverdraftUsedLimitCurrency(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUnarrangedOverdraftAmount(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUnarrangedOverdraftAmountCurrency(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AccountOverdraftLimitsData) validateOverdraftContractedLimit(formats strfmt.Registry) error {

	if err := validate.Required("overdraftContractedLimit", "body", m.OverdraftContractedLimit); err != nil {
		return err
	}

	return nil
}

func (m *AccountOverdraftLimitsData) validateOverdraftContractedLimitCurrency(formats strfmt.Registry) error {

	if err := validate.Required("overdraftContractedLimitCurrency", "body", m.OverdraftContractedLimitCurrency); err != nil {
		return err
	}

	if err := validate.MaxLength("overdraftContractedLimitCurrency", "body", *m.OverdraftContractedLimitCurrency, 3); err != nil {
		return err
	}

	if err := validate.Pattern("overdraftContractedLimitCurrency", "body", *m.OverdraftContractedLimitCurrency, `^(\w{3}){1}$`); err != nil {
		return err
	}

	return nil
}

func (m *AccountOverdraftLimitsData) validateOverdraftUsedLimit(formats strfmt.Registry) error {

	if err := validate.Required("overdraftUsedLimit", "body", m.OverdraftUsedLimit); err != nil {
		return err
	}

	return nil
}

func (m *AccountOverdraftLimitsData) validateOverdraftUsedLimitCurrency(formats strfmt.Registry) error {

	if err := validate.Required("overdraftUsedLimitCurrency", "body", m.OverdraftUsedLimitCurrency); err != nil {
		return err
	}

	if err := validate.MaxLength("overdraftUsedLimitCurrency", "body", *m.OverdraftUsedLimitCurrency, 3); err != nil {
		return err
	}

	if err := validate.Pattern("overdraftUsedLimitCurrency", "body", *m.OverdraftUsedLimitCurrency, `^(\w{3}){1}$`); err != nil {
		return err
	}

	return nil
}

func (m *AccountOverdraftLimitsData) validateUnarrangedOverdraftAmount(formats strfmt.Registry) error {

	if err := validate.Required("unarrangedOverdraftAmount", "body", m.UnarrangedOverdraftAmount); err != nil {
		return err
	}

	return nil
}

func (m *AccountOverdraftLimitsData) validateUnarrangedOverdraftAmountCurrency(formats strfmt.Registry) error {

	if err := validate.Required("unarrangedOverdraftAmountCurrency", "body", m.UnarrangedOverdraftAmountCurrency); err != nil {
		return err
	}

	if err := validate.MaxLength("unarrangedOverdraftAmountCurrency", "body", *m.UnarrangedOverdraftAmountCurrency, 3); err != nil {
		return err
	}

	if err := validate.Pattern("unarrangedOverdraftAmountCurrency", "body", *m.UnarrangedOverdraftAmountCurrency, `^(\w{3}){1}$`); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this account overdraft limits data based on context it is used
func (m *AccountOverdraftLimitsData) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *AccountOverdraftLimitsData) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AccountOverdraftLimitsData) UnmarshalBinary(b []byte) error {
	var res AccountOverdraftLimitsData
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
