// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/validate"
)

// DepositTransactionType2 DepositTransactionType2
//
// CHECK, WITHDRAWAL, TRANSFER, POSDEBIT, ATMWITHDRAWAL, BILLPAYMENT, FEE, DEPOSIT, ADJUSTMENT, INTEREST, DIVIDEND, DIRECTDEPOSIT, ATMDEPOSIT, POSCREDIT
//
// swagger:model DepositTransactionType2
type DepositTransactionType2 string

func NewDepositTransactionType2(value DepositTransactionType2) *DepositTransactionType2 {
	v := value
	return &v
}

const (

	// DepositTransactionType2ADJUSTMENT captures enum value "ADJUSTMENT"
	DepositTransactionType2ADJUSTMENT DepositTransactionType2 = "ADJUSTMENT"

	// DepositTransactionType2ATMDEPOSIT captures enum value "ATMDEPOSIT"
	DepositTransactionType2ATMDEPOSIT DepositTransactionType2 = "ATMDEPOSIT"

	// DepositTransactionType2ATMWITHDRAWAL captures enum value "ATMWITHDRAWAL"
	DepositTransactionType2ATMWITHDRAWAL DepositTransactionType2 = "ATMWITHDRAWAL"

	// DepositTransactionType2BILLPAYMENT captures enum value "BILLPAYMENT"
	DepositTransactionType2BILLPAYMENT DepositTransactionType2 = "BILLPAYMENT"

	// DepositTransactionType2CHECK captures enum value "CHECK"
	DepositTransactionType2CHECK DepositTransactionType2 = "CHECK"

	// DepositTransactionType2DEPOSIT captures enum value "DEPOSIT"
	DepositTransactionType2DEPOSIT DepositTransactionType2 = "DEPOSIT"

	// DepositTransactionType2DIRECTDEPOSIT captures enum value "DIRECTDEPOSIT"
	DepositTransactionType2DIRECTDEPOSIT DepositTransactionType2 = "DIRECTDEPOSIT"

	// DepositTransactionType2DIVIDEND captures enum value "DIVIDEND"
	DepositTransactionType2DIVIDEND DepositTransactionType2 = "DIVIDEND"

	// DepositTransactionType2FEE captures enum value "FEE"
	DepositTransactionType2FEE DepositTransactionType2 = "FEE"

	// DepositTransactionType2INTEREST captures enum value "INTEREST"
	DepositTransactionType2INTEREST DepositTransactionType2 = "INTEREST"

	// DepositTransactionType2POSCREDIT captures enum value "POSCREDIT"
	DepositTransactionType2POSCREDIT DepositTransactionType2 = "POSCREDIT"

	// DepositTransactionType2POSDEBIT captures enum value "POSDEBIT"
	DepositTransactionType2POSDEBIT DepositTransactionType2 = "POSDEBIT"

	// DepositTransactionType2TRANSFER captures enum value "TRANSFER"
	DepositTransactionType2TRANSFER DepositTransactionType2 = "TRANSFER"

	// DepositTransactionType2WITHDRAWAL captures enum value "WITHDRAWAL"
	DepositTransactionType2WITHDRAWAL DepositTransactionType2 = "WITHDRAWAL"
)

// for schema
var depositTransactionType2Enum []interface{}

func init() {
	var res []DepositTransactionType2
	if err := json.Unmarshal([]byte(`["ADJUSTMENT","ATMDEPOSIT","ATMWITHDRAWAL","BILLPAYMENT","CHECK","DEPOSIT","DIRECTDEPOSIT","DIVIDEND","FEE","INTEREST","POSCREDIT","POSDEBIT","TRANSFER","WITHDRAWAL"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		depositTransactionType2Enum = append(depositTransactionType2Enum, v)
	}
}

func (m DepositTransactionType2) validateDepositTransactionType2Enum(path, location string, value DepositTransactionType2) error {
	if err := validate.EnumCase(path, location, value, depositTransactionType2Enum, true); err != nil {
		return err
	}
	return nil
}

// Validate validates this deposit transaction type2
func (m DepositTransactionType2) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateDepositTransactionType2Enum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validates this deposit transaction type2 based on context it is used
func (m DepositTransactionType2) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}
