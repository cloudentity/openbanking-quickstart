// Code generated by go-swagger; DO NOT EDIT.

package personal_information

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

// NewGetAccountContactParams creates a new GetAccountContactParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetAccountContactParams() *GetAccountContactParams {
	return &GetAccountContactParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetAccountContactParamsWithTimeout creates a new GetAccountContactParams object
// with the ability to set a timeout on a request.
func NewGetAccountContactParamsWithTimeout(timeout time.Duration) *GetAccountContactParams {
	return &GetAccountContactParams{
		timeout: timeout,
	}
}

// NewGetAccountContactParamsWithContext creates a new GetAccountContactParams object
// with the ability to set a context for a request.
func NewGetAccountContactParamsWithContext(ctx context.Context) *GetAccountContactParams {
	return &GetAccountContactParams{
		Context: ctx,
	}
}

// NewGetAccountContactParamsWithHTTPClient creates a new GetAccountContactParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetAccountContactParamsWithHTTPClient(client *http.Client) *GetAccountContactParams {
	return &GetAccountContactParams{
		HTTPClient: client,
	}
}

/* GetAccountContactParams contains all the parameters to send to the API endpoint
   for the get account contact operation.

   Typically these are written to a http.Request.
*/
type GetAccountContactParams struct {

	/* AccountID.

	   Account Identifier
	*/
	AccountID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get account contact params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetAccountContactParams) WithDefaults() *GetAccountContactParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get account contact params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetAccountContactParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get account contact params
func (o *GetAccountContactParams) WithTimeout(timeout time.Duration) *GetAccountContactParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get account contact params
func (o *GetAccountContactParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get account contact params
func (o *GetAccountContactParams) WithContext(ctx context.Context) *GetAccountContactParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get account contact params
func (o *GetAccountContactParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get account contact params
func (o *GetAccountContactParams) WithHTTPClient(client *http.Client) *GetAccountContactParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get account contact params
func (o *GetAccountContactParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAccountID adds the accountID to the get account contact params
func (o *GetAccountContactParams) WithAccountID(accountID string) *GetAccountContactParams {
	o.SetAccountID(accountID)
	return o
}

// SetAccountID adds the accountId to the get account contact params
func (o *GetAccountContactParams) SetAccountID(accountID string) {
	o.AccountID = accountID
}

// WriteToRequest writes these params to a swagger request
func (o *GetAccountContactParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param accountId
	if err := r.SetPathParam("accountId", o.AccountID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}