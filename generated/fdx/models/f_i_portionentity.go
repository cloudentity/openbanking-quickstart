// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// FIPortionentity FIPortionentity
//
// Financial Institution-specific asset allocation
//
// swagger:model FIPortionentity
type FIPortionentity struct {

	// Financial Institution-specific asset class
	AssetClass string `json:"assetClass,omitempty"`

	// Percentage of asset class that falls under this asset
	Percent float64 `json:"percent,omitempty"`
}

// Validate validates this f i portionentity
func (m *FIPortionentity) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this f i portionentity based on context it is used
func (m *FIPortionentity) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *FIPortionentity) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *FIPortionentity) UnmarshalBinary(b []byte) error {
	var res FIPortionentity
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
