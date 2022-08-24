// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// PageMetadata1 PageMetadata1
//
// Offset IDs for navigating result sets
//
// swagger:model PageMetadata1
type PageMetadata1 struct {

	// Opaque identifier. Does not need to be numeric or have any specific pattern. Implementation specific
	// Example: 2
	NextOffset string `json:"nextOffset,omitempty"`

	// Opaque identifier. Does not need to be numeric or have any specific pattern. Implementation specific
	// Example: 1
	PrevOffset string `json:"prevOffset,omitempty"`

	// Total number of elements
	// Example: 3
	TotalElements int32 `json:"totalElements,omitempty"`
}

// Validate validates this page metadata1
func (m *PageMetadata1) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this page metadata1 based on context it is used
func (m *PageMetadata1) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *PageMetadata1) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PageMetadata1) UnmarshalBinary(b []byte) error {
	var res PageMetadata1
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
