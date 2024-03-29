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

// OrderType OrderType
//
// The type of the order
//
// swagger:model OrderType
type OrderType string

func NewOrderType(value OrderType) *OrderType {
	v := value
	return &v
}

const (

	// OrderTypeBUY captures enum value "BUY"
	OrderTypeBUY OrderType = "BUY"

	// OrderTypeSELL captures enum value "SELL"
	OrderTypeSELL OrderType = "SELL"

	// OrderTypeBUYTOCOVER captures enum value "BUYTOCOVER"
	OrderTypeBUYTOCOVER OrderType = "BUYTOCOVER"

	// OrderTypeBUYTOOPEN captures enum value "BUYTOOPEN"
	OrderTypeBUYTOOPEN OrderType = "BUYTOOPEN"

	// OrderTypeSELLTOCOVER captures enum value "SELLTOCOVER"
	OrderTypeSELLTOCOVER OrderType = "SELLTOCOVER"

	// OrderTypeSELLTOOPEN captures enum value "SELLTOOPEN"
	OrderTypeSELLTOOPEN OrderType = "SELLTOOPEN"

	// OrderTypeSELLSHORT captures enum value "SELLSHORT"
	OrderTypeSELLSHORT OrderType = "SELLSHORT"

	// OrderTypeSELLCLOSE captures enum value "SELLCLOSE"
	OrderTypeSELLCLOSE OrderType = "SELLCLOSE"
)

// for schema
var orderTypeEnum []interface{}

func init() {
	var res []OrderType
	if err := json.Unmarshal([]byte(`["BUY","SELL","BUYTOCOVER","BUYTOOPEN","SELLTOCOVER","SELLTOOPEN","SELLSHORT","SELLCLOSE"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		orderTypeEnum = append(orderTypeEnum, v)
	}
}

func (m OrderType) validateOrderTypeEnum(path, location string, value OrderType) error {
	if err := validate.EnumCase(path, location, value, orderTypeEnum, true); err != nil {
		return err
	}
	return nil
}

// Validate validates this order type
func (m OrderType) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateOrderTypeEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validates this order type based on context it is used
func (m OrderType) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}
