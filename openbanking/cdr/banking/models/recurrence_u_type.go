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

// RecurrenceUType RecurrenceUType
//
// The type of recurrence used to define the schedule
// Example: eventBased
//
// swagger:model RecurrenceUType
type RecurrenceUType string

const (

	// RecurrenceUTypeEventBased captures enum value "eventBased"
	RecurrenceUTypeEventBased RecurrenceUType = "eventBased"

	// RecurrenceUTypeIntervalSchedule captures enum value "intervalSchedule"
	RecurrenceUTypeIntervalSchedule RecurrenceUType = "intervalSchedule"

	// RecurrenceUTypeLastWeekDay captures enum value "lastWeekDay"
	RecurrenceUTypeLastWeekDay RecurrenceUType = "lastWeekDay"

	// RecurrenceUTypeOnceOff captures enum value "onceOff"
	RecurrenceUTypeOnceOff RecurrenceUType = "onceOff"
)

// for schema
var recurrenceUTypeEnum []interface{}

func init() {
	var res []RecurrenceUType
	if err := json.Unmarshal([]byte(`["eventBased","intervalSchedule","lastWeekDay","onceOff"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		recurrenceUTypeEnum = append(recurrenceUTypeEnum, v)
	}
}

func (m RecurrenceUType) validateRecurrenceUTypeEnum(path, location string, value RecurrenceUType) error {
	if err := validate.EnumCase(path, location, value, recurrenceUTypeEnum, true); err != nil {
		return err
	}
	return nil
}

// Validate validates this recurrence u type
func (m RecurrenceUType) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateRecurrenceUTypeEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validates this recurrence u type based on context it is used
func (m RecurrenceUType) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}
