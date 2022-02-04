// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// BankingAuthorisedEntity BankingAuthorisedEntity
//
// swagger:model BankingAuthorisedEntity
type BankingAuthorisedEntity struct {

	// Australian Business Number for the authorised entity
	Abn string `json:"abn,omitempty"`

	// Australian Company Number for the authorised entity
	Acn string `json:"acn,omitempty"`

	// Australian Registered Body Number for the authorised entity
	Arbn string `json:"arbn,omitempty"`

	// Description of the authorised entity derived from previously executed direct debits
	Description string `json:"description,omitempty"`

	// Name of the financial institution through which the direct debit will be executed. Is required unless the payment is made via a credit card scheme
	FinancialInstitution string `json:"financialInstitution,omitempty"`
}

// Validate validates this banking authorised entity
func (m *BankingAuthorisedEntity) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this banking authorised entity based on context it is used
func (m *BankingAuthorisedEntity) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *BankingAuthorisedEntity) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *BankingAuthorisedEntity) UnmarshalBinary(b []byte) error {
	var res BankingAuthorisedEntity
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}