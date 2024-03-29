// Code generated by go-swagger; DO NOT EDIT.

package recurring_payments

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
)

// NewGetRecurringPaymentParams creates a new GetRecurringPaymentParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetRecurringPaymentParams() *GetRecurringPaymentParams {
	return &GetRecurringPaymentParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetRecurringPaymentParamsWithTimeout creates a new GetRecurringPaymentParams object
// with the ability to set a timeout on a request.
func NewGetRecurringPaymentParamsWithTimeout(timeout time.Duration) *GetRecurringPaymentParams {
	return &GetRecurringPaymentParams{
		timeout: timeout,
	}
}

// NewGetRecurringPaymentParamsWithContext creates a new GetRecurringPaymentParams object
// with the ability to set a context for a request.
func NewGetRecurringPaymentParamsWithContext(ctx context.Context) *GetRecurringPaymentParams {
	return &GetRecurringPaymentParams{
		Context: ctx,
	}
}

// NewGetRecurringPaymentParamsWithHTTPClient creates a new GetRecurringPaymentParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetRecurringPaymentParamsWithHTTPClient(client *http.Client) *GetRecurringPaymentParams {
	return &GetRecurringPaymentParams{
		HTTPClient: client,
	}
}

/* GetRecurringPaymentParams contains all the parameters to send to the API endpoint
   for the get recurring payment operation.

   Typically these are written to a http.Request.
*/
type GetRecurringPaymentParams struct {

	/* RecurringPaymentID.

	   Recurring Payment Identifier. Uniquely identifies a recurring payment
	*/
	RecurringPaymentID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get recurring payment params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetRecurringPaymentParams) WithDefaults() *GetRecurringPaymentParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get recurring payment params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetRecurringPaymentParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get recurring payment params
func (o *GetRecurringPaymentParams) WithTimeout(timeout time.Duration) *GetRecurringPaymentParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get recurring payment params
func (o *GetRecurringPaymentParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get recurring payment params
func (o *GetRecurringPaymentParams) WithContext(ctx context.Context) *GetRecurringPaymentParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get recurring payment params
func (o *GetRecurringPaymentParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get recurring payment params
func (o *GetRecurringPaymentParams) WithHTTPClient(client *http.Client) *GetRecurringPaymentParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get recurring payment params
func (o *GetRecurringPaymentParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithRecurringPaymentID adds the recurringPaymentID to the get recurring payment params
func (o *GetRecurringPaymentParams) WithRecurringPaymentID(recurringPaymentID string) *GetRecurringPaymentParams {
	o.SetRecurringPaymentID(recurringPaymentID)
	return o
}

// SetRecurringPaymentID adds the recurringPaymentId to the get recurring payment params
func (o *GetRecurringPaymentParams) SetRecurringPaymentID(recurringPaymentID string) {
	o.RecurringPaymentID = recurringPaymentID
}

// WriteToRequest writes these params to a swagger request
func (o *GetRecurringPaymentParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param recurringPaymentId
	if err := r.SetPathParam("recurringPaymentId", o.RecurringPaymentID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
