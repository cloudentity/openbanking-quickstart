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

// PaymentNetworkType2 PaymentNetworkType2
//
// Type of payment network
//
// swagger:model PaymentNetworkType2
type PaymentNetworkType2 string

func NewPaymentNetworkType2(value PaymentNetworkType2) *PaymentNetworkType2 {
	v := value
	return &v
}

const (

	// PaymentNetworkType2CAACSS captures enum value "CA_ACSS"
	PaymentNetworkType2CAACSS PaymentNetworkType2 = "CA_ACSS"

	// PaymentNetworkType2CALVTS captures enum value "CA_LVTS"
	PaymentNetworkType2CALVTS PaymentNetworkType2 = "CA_LVTS"

	// PaymentNetworkType2USACH captures enum value "US_ACH"
	PaymentNetworkType2USACH PaymentNetworkType2 = "US_ACH"

	// PaymentNetworkType2USCHIPS captures enum value "US_CHIPS"
	PaymentNetworkType2USCHIPS PaymentNetworkType2 = "US_CHIPS"

	// PaymentNetworkType2USFEDWIRE captures enum value "US_FEDWIRE"
	PaymentNetworkType2USFEDWIRE PaymentNetworkType2 = "US_FEDWIRE"

	// PaymentNetworkType2USRTP captures enum value "US_RTP"
	PaymentNetworkType2USRTP PaymentNetworkType2 = "US_RTP"
)

// for schema
var paymentNetworkType2Enum []interface{}

func init() {
	var res []PaymentNetworkType2
	if err := json.Unmarshal([]byte(`["CA_ACSS","CA_LVTS","US_ACH","US_CHIPS","US_FEDWIRE","US_RTP"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		paymentNetworkType2Enum = append(paymentNetworkType2Enum, v)
	}
}

func (m PaymentNetworkType2) validatePaymentNetworkType2Enum(path, location string, value PaymentNetworkType2) error {
	if err := validate.EnumCase(path, location, value, paymentNetworkType2Enum, true); err != nil {
		return err
	}
	return nil
}

// Validate validates this payment network type2
func (m PaymentNetworkType2) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validatePaymentNetworkType2Enum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validates this payment network type2 based on context it is used
func (m PaymentNetworkType2) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}