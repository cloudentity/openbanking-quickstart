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

// AccountStatus AccountStatus
//
// The status of an account
//
// swagger:model AccountStatus
type AccountStatus string

func NewAccountStatus(value AccountStatus) *AccountStatus {
	v := value
	return &v
}

const (

	// AccountStatusCLOSED captures enum value "CLOSED"
	AccountStatusCLOSED AccountStatus = "CLOSED"

	// AccountStatusDELINQUENT captures enum value "DELINQUENT"
	AccountStatusDELINQUENT AccountStatus = "DELINQUENT"

	// AccountStatusNEGATIVECURRENTBALANCE captures enum value "NEGATIVECURRENTBALANCE"
	AccountStatusNEGATIVECURRENTBALANCE AccountStatus = "NEGATIVECURRENTBALANCE"

	// AccountStatusOPEN captures enum value "OPEN"
	AccountStatusOPEN AccountStatus = "OPEN"

	// AccountStatusPAID captures enum value "PAID"
	AccountStatusPAID AccountStatus = "PAID"

	// AccountStatusPENDINGCLOSE captures enum value "PENDINGCLOSE"
	AccountStatusPENDINGCLOSE AccountStatus = "PENDINGCLOSE"

	// AccountStatusPENDINGOPEN captures enum value "PENDINGOPEN"
	AccountStatusPENDINGOPEN AccountStatus = "PENDINGOPEN"

	// AccountStatusRESTRICTED captures enum value "RESTRICTED"
	AccountStatusRESTRICTED AccountStatus = "RESTRICTED"
)

// for schema
var accountStatusEnum []interface{}

func init() {
	var res []AccountStatus
	if err := json.Unmarshal([]byte(`["CLOSED","DELINQUENT","NEGATIVECURRENTBALANCE","OPEN","PAID","PENDINGCLOSE","PENDINGOPEN","RESTRICTED"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		accountStatusEnum = append(accountStatusEnum, v)
	}
}

func (m AccountStatus) validateAccountStatusEnum(path, location string, value AccountStatus) error {
	if err := validate.EnumCase(path, location, value, accountStatusEnum, true); err != nil {
		return err
	}
	return nil
}

// Validate validates this account status
func (m AccountStatus) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateAccountStatusEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validates this account status based on context it is used
func (m AccountStatus) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}