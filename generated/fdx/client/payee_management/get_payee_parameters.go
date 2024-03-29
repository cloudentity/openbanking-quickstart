// Code generated by go-swagger; DO NOT EDIT.

package payee_management

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

// NewGetPayeeParams creates a new GetPayeeParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetPayeeParams() *GetPayeeParams {
	return &GetPayeeParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetPayeeParamsWithTimeout creates a new GetPayeeParams object
// with the ability to set a timeout on a request.
func NewGetPayeeParamsWithTimeout(timeout time.Duration) *GetPayeeParams {
	return &GetPayeeParams{
		timeout: timeout,
	}
}

// NewGetPayeeParamsWithContext creates a new GetPayeeParams object
// with the ability to set a context for a request.
func NewGetPayeeParamsWithContext(ctx context.Context) *GetPayeeParams {
	return &GetPayeeParams{
		Context: ctx,
	}
}

// NewGetPayeeParamsWithHTTPClient creates a new GetPayeeParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetPayeeParamsWithHTTPClient(client *http.Client) *GetPayeeParams {
	return &GetPayeeParams{
		HTTPClient: client,
	}
}

/* GetPayeeParams contains all the parameters to send to the API endpoint
   for the get payee operation.

   Typically these are written to a http.Request.
*/
type GetPayeeParams struct {

	/* PayeeID.

	   Payee Identifier. Uniquely identifies a payee
	*/
	PayeeID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get payee params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetPayeeParams) WithDefaults() *GetPayeeParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get payee params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetPayeeParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get payee params
func (o *GetPayeeParams) WithTimeout(timeout time.Duration) *GetPayeeParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get payee params
func (o *GetPayeeParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get payee params
func (o *GetPayeeParams) WithContext(ctx context.Context) *GetPayeeParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get payee params
func (o *GetPayeeParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get payee params
func (o *GetPayeeParams) WithHTTPClient(client *http.Client) *GetPayeeParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get payee params
func (o *GetPayeeParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithPayeeID adds the payeeID to the get payee params
func (o *GetPayeeParams) WithPayeeID(payeeID string) *GetPayeeParams {
	o.SetPayeeID(payeeID)
	return o
}

// SetPayeeID adds the payeeId to the get payee params
func (o *GetPayeeParams) SetPayeeID(payeeID string) {
	o.PayeeID = payeeID
}

// WriteToRequest writes these params to a swagger request
func (o *GetPayeeParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param payeeId
	if err := r.SetPathParam("payeeId", o.PayeeID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
