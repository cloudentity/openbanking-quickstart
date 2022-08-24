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

// PartyRegistry PartyRegistry
//
// The registry containing the party’s registration with name and id
//
// swagger:model PartyRegistry
type PartyRegistry string

func NewPartyRegistry(value PartyRegistry) *PartyRegistry {
	v := value
	return &v
}

const (

	// PartyRegistryFDX captures enum value "FDX"
	PartyRegistryFDX PartyRegistry = "FDX"

	// PartyRegistryGLEIF captures enum value "GLEIF"
	PartyRegistryGLEIF PartyRegistry = "GLEIF"

	// PartyRegistryICANN captures enum value "ICANN"
	PartyRegistryICANN PartyRegistry = "ICANN"

	// PartyRegistryPRIVATE captures enum value "PRIVATE"
	PartyRegistryPRIVATE PartyRegistry = "PRIVATE"
)

// for schema
var partyRegistryEnum []interface{}

func init() {
	var res []PartyRegistry
	if err := json.Unmarshal([]byte(`["FDX","GLEIF","ICANN","PRIVATE"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		partyRegistryEnum = append(partyRegistryEnum, v)
	}
}

func (m PartyRegistry) validatePartyRegistryEnum(path, location string, value PartyRegistry) error {
	if err := validate.EnumCase(path, location, value, partyRegistryEnum, true); err != nil {
		return err
	}
	return nil
}

// Validate validates this party registry
func (m PartyRegistry) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validatePartyRegistryEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validates this party registry based on context it is used
func (m PartyRegistry) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}
