// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/validate"
)

// Frequency0 Individual Definitions:
// EvryDay - Every day
// EvryWorkgDay - Every working day
// IntrvlWkDay - An interval specified in weeks (01 to 09), and the day within the week (01 to 07)
// WkInMnthDay - A monthly interval, specifying the week of the month (01 to 05) and day within the week (01 to 07)
// IntrvlMnthDay - An interval specified in months (between 01 to 06, 12, 24), specifying the day within the month (-5 to -1, 1 to 31)
// QtrDay - Quarterly (either ENGLISH, SCOTTISH, or RECEIVED)
// ENGLISH = Paid on the 25th March, 24th June, 29th September and 25th December.
// SCOTTISH = Paid on the 2nd February, 15th May, 1st August and 11th November.
// RECEIVED = Paid on the 20th March, 19th June, 24th September and 20th December.
// Individual Patterns:
// EvryDay (ScheduleCode)
// EvryWorkgDay (ScheduleCode)
// IntrvlWkDay:IntervalInWeeks:DayInWeek (ScheduleCode + IntervalInWeeks + DayInWeek)
// WkInMnthDay:WeekInMonth:DayInWeek (ScheduleCode + WeekInMonth + DayInWeek)
// IntrvlMnthDay:IntervalInMonths:DayInMonth (ScheduleCode + IntervalInMonths + DayInMonth)
// QtrDay: + either (ENGLISH, SCOTTISH or RECEIVED) ScheduleCode + QuarterDay
// The regular expression for this element combines five smaller versions for each permitted pattern. To aid legibility - the components are presented individually here:
// EvryDay
// EvryWorkgDay
// IntrvlWkDay:0[1-9]:0[1-7]
// WkInMnthDay:0[1-5]:0[1-7]
// IntrvlMnthDay:(0[1-6]|12|24):(-0[1-5]|0[1-9]|[12][0-9]|3[01])
// QtrDay:(ENGLISH|SCOTTISH|RECEIVED)
// Full Regular Expression:
// ^(EvryDay)$|^(EvryWorkgDay)$|^(IntrvlWkDay:0[1-9]:0[1-7])$|^(WkInMnthDay:0[1-5]:0[1-7])$|^(IntrvlMnthDay:(0[1-6]|12|24):(-0[1-5]|0[1-9]|[12][0-9]|3[01]))$|^(QtrDay:(ENGLISH|SCOTTISH|RECEIVED))$
//
// swagger:model Frequency_0
type Frequency0 string

// Validate validates this frequency 0
func (m Frequency0) Validate(formats strfmt.Registry) error {
	var res []error

	if err := validate.Pattern("", "body", string(m), `^(EvryDay)$|^(EvryWorkgDay)$|^(IntrvlWkDay:0[1-9]:0[1-7])$|^(WkInMnthDay:0[1-5]:0[1-7])$|^(IntrvlMnthDay:(0[1-6]|12|24):(-0[1-5]|0[1-9]|[12][0-9]|3[01]))$|^(QtrDay:(ENGLISH|SCOTTISH|RECEIVED))$`); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validates this frequency 0 based on context it is used
func (m Frequency0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}
