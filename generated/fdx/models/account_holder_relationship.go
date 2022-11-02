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

// AccountHolderRelationship AccountHolderRelationship
//
// Types of relationships between accounts and holders. Suggested values
//
// swagger:model AccountHolderRelationship
type AccountHolderRelationship string

func NewAccountHolderRelationship(value AccountHolderRelationship) *AccountHolderRelationship {
	v := value
	return &v
}

const (

	// AccountHolderRelationshipAUTHORIZEDUSER captures enum value "AUTHORIZED_USER"
	AccountHolderRelationshipAUTHORIZEDUSER AccountHolderRelationship = "AUTHORIZED_USER"

	// AccountHolderRelationshipBUSINESS captures enum value "BUSINESS"
	AccountHolderRelationshipBUSINESS AccountHolderRelationship = "BUSINESS"

	// AccountHolderRelationshipFORBENEFITOF captures enum value "FOR_BENEFIT_OF"
	AccountHolderRelationshipFORBENEFITOF AccountHolderRelationship = "FOR_BENEFIT_OF"

	// AccountHolderRelationshipFORBENEFITOFPRIMARY captures enum value "FOR_BENEFIT_OF_PRIMARY"
	AccountHolderRelationshipFORBENEFITOFPRIMARY AccountHolderRelationship = "FOR_BENEFIT_OF_PRIMARY"

	// AccountHolderRelationshipFORBENEFITOFPRIMARYJOINTRESTRICTED captures enum value "FOR_BENEFIT_OF_PRIMARY_JOINT_RESTRICTED"
	AccountHolderRelationshipFORBENEFITOFPRIMARYJOINTRESTRICTED AccountHolderRelationship = "FOR_BENEFIT_OF_PRIMARY_JOINT_RESTRICTED"

	// AccountHolderRelationshipFORBENEFITOFSECONDARY captures enum value "FOR_BENEFIT_OF_SECONDARY"
	AccountHolderRelationshipFORBENEFITOFSECONDARY AccountHolderRelationship = "FOR_BENEFIT_OF_SECONDARY"

	// AccountHolderRelationshipFORBENEFITOFSECONDARYJOINTRESTRICTED captures enum value "FOR_BENEFIT_OF_SECONDARY_JOINT_RESTRICTED"
	AccountHolderRelationshipFORBENEFITOFSECONDARYJOINTRESTRICTED AccountHolderRelationship = "FOR_BENEFIT_OF_SECONDARY_JOINT_RESTRICTED"

	// AccountHolderRelationshipFORBENEFITOFSOLEOWNERRESTRICTED captures enum value "FOR_BENEFIT_OF_SOLE_OWNER_RESTRICTED"
	AccountHolderRelationshipFORBENEFITOFSOLEOWNERRESTRICTED AccountHolderRelationship = "FOR_BENEFIT_OF_SOLE_OWNER_RESTRICTED"

	// AccountHolderRelationshipPOWEROFATTORNEY captures enum value "POWER_OF_ATTORNEY"
	AccountHolderRelationshipPOWEROFATTORNEY AccountHolderRelationship = "POWER_OF_ATTORNEY"

	// AccountHolderRelationshipPRIMARYJOINTTENANTS captures enum value "PRIMARY_JOINT_TENANTS"
	AccountHolderRelationshipPRIMARYJOINTTENANTS AccountHolderRelationship = "PRIMARY_JOINT_TENANTS"

	// AccountHolderRelationshipPRIMARY captures enum value "PRIMARY"
	AccountHolderRelationshipPRIMARY AccountHolderRelationship = "PRIMARY"

	// AccountHolderRelationshipPRIMARYBORROWER captures enum value "PRIMARY_BORROWER"
	AccountHolderRelationshipPRIMARYBORROWER AccountHolderRelationship = "PRIMARY_BORROWER"

	// AccountHolderRelationshipPRIMARYJOINT captures enum value "PRIMARY_JOINT"
	AccountHolderRelationshipPRIMARYJOINT AccountHolderRelationship = "PRIMARY_JOINT"

	// AccountHolderRelationshipSECONDARY captures enum value "SECONDARY"
	AccountHolderRelationshipSECONDARY AccountHolderRelationship = "SECONDARY"

	// AccountHolderRelationshipSECONDARYJOINTTENANTS captures enum value "SECONDARY_JOINT_TENANTS"
	AccountHolderRelationshipSECONDARYJOINTTENANTS AccountHolderRelationship = "SECONDARY_JOINT_TENANTS"

	// AccountHolderRelationshipSECONDARYBORROWER captures enum value "SECONDARY_BORROWER"
	AccountHolderRelationshipSECONDARYBORROWER AccountHolderRelationship = "SECONDARY_BORROWER"

	// AccountHolderRelationshipSECONDARYJOINT captures enum value "SECONDARY_JOINT"
	AccountHolderRelationshipSECONDARYJOINT AccountHolderRelationship = "SECONDARY_JOINT"

	// AccountHolderRelationshipSOLEOWNER captures enum value "SOLE_OWNER"
	AccountHolderRelationshipSOLEOWNER AccountHolderRelationship = "SOLE_OWNER"

	// AccountHolderRelationshipTRUSTEE captures enum value "TRUSTEE"
	AccountHolderRelationshipTRUSTEE AccountHolderRelationship = "TRUSTEE"

	// AccountHolderRelationshipUNIFORMTRANSFERTOMINOR captures enum value "UNIFORM_TRANSFER_TO_MINOR"
	AccountHolderRelationshipUNIFORMTRANSFERTOMINOR AccountHolderRelationship = "UNIFORM_TRANSFER_TO_MINOR"
)

// for schema
var accountHolderRelationshipEnum []interface{}

func init() {
	var res []AccountHolderRelationship
	if err := json.Unmarshal([]byte(`["AUTHORIZED_USER","BUSINESS","FOR_BENEFIT_OF","FOR_BENEFIT_OF_PRIMARY","FOR_BENEFIT_OF_PRIMARY_JOINT_RESTRICTED","FOR_BENEFIT_OF_SECONDARY","FOR_BENEFIT_OF_SECONDARY_JOINT_RESTRICTED","FOR_BENEFIT_OF_SOLE_OWNER_RESTRICTED","POWER_OF_ATTORNEY","PRIMARY_JOINT_TENANTS","PRIMARY","PRIMARY_BORROWER","PRIMARY_JOINT","SECONDARY","SECONDARY_JOINT_TENANTS","SECONDARY_BORROWER","SECONDARY_JOINT","SOLE_OWNER","TRUSTEE","UNIFORM_TRANSFER_TO_MINOR"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		accountHolderRelationshipEnum = append(accountHolderRelationshipEnum, v)
	}
}

func (m AccountHolderRelationship) validateAccountHolderRelationshipEnum(path, location string, value AccountHolderRelationship) error {
	if err := validate.EnumCase(path, location, value, accountHolderRelationshipEnum, true); err != nil {
		return err
	}
	return nil
}

// Validate validates this account holder relationship
func (m AccountHolderRelationship) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateAccountHolderRelationshipEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validates this account holder relationship based on context it is used
func (m AccountHolderRelationship) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}