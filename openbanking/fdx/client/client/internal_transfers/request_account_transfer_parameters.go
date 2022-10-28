// Code generated by go-swagger; DO NOT EDIT.

package internal_transfers

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

	"github.com/cloudentity/openbanking-quickstart/openbanking/fdx/client/models"
)

// NewRequestAccountTransferParams creates a new RequestAccountTransferParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewRequestAccountTransferParams() *RequestAccountTransferParams {
	return &RequestAccountTransferParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewRequestAccountTransferParamsWithTimeout creates a new RequestAccountTransferParams object
// with the ability to set a timeout on a request.
func NewRequestAccountTransferParamsWithTimeout(timeout time.Duration) *RequestAccountTransferParams {
	return &RequestAccountTransferParams{
		timeout: timeout,
	}
}

// NewRequestAccountTransferParamsWithContext creates a new RequestAccountTransferParams object
// with the ability to set a context for a request.
func NewRequestAccountTransferParamsWithContext(ctx context.Context) *RequestAccountTransferParams {
	return &RequestAccountTransferParams{
		Context: ctx,
	}
}

// NewRequestAccountTransferParamsWithHTTPClient creates a new RequestAccountTransferParams object
// with the ability to set a custom HTTPClient for a request.
func NewRequestAccountTransferParamsWithHTTPClient(client *http.Client) *RequestAccountTransferParams {
	return &RequestAccountTransferParams{
		HTTPClient: client,
	}
}

/* RequestAccountTransferParams contains all the parameters to send to the API endpoint
   for the request account transfer operation.

   Typically these are written to a http.Request.
*/
type RequestAccountTransferParams struct {

	// ContentType.
	ContentType *string

	/* Body.

	   Data of the transfer request
	*/
	Body *models.TransferforCreateentity1

	/* IdempotencyKey.

	   Used to de-duplicate requests
	*/
	IdempotencyKey string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the request account transfer params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *RequestAccountTransferParams) WithDefaults() *RequestAccountTransferParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the request account transfer params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *RequestAccountTransferParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the request account transfer params
func (o *RequestAccountTransferParams) WithTimeout(timeout time.Duration) *RequestAccountTransferParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the request account transfer params
func (o *RequestAccountTransferParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the request account transfer params
func (o *RequestAccountTransferParams) WithContext(ctx context.Context) *RequestAccountTransferParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the request account transfer params
func (o *RequestAccountTransferParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the request account transfer params
func (o *RequestAccountTransferParams) WithHTTPClient(client *http.Client) *RequestAccountTransferParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the request account transfer params
func (o *RequestAccountTransferParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithContentType adds the contentType to the request account transfer params
func (o *RequestAccountTransferParams) WithContentType(contentType *string) *RequestAccountTransferParams {
	o.SetContentType(contentType)
	return o
}

// SetContentType adds the contentType to the request account transfer params
func (o *RequestAccountTransferParams) SetContentType(contentType *string) {
	o.ContentType = contentType
}

// WithBody adds the body to the request account transfer params
func (o *RequestAccountTransferParams) WithBody(body *models.TransferforCreateentity1) *RequestAccountTransferParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the request account transfer params
func (o *RequestAccountTransferParams) SetBody(body *models.TransferforCreateentity1) {
	o.Body = body
}

// WithIdempotencyKey adds the idempotencyKey to the request account transfer params
func (o *RequestAccountTransferParams) WithIdempotencyKey(idempotencyKey string) *RequestAccountTransferParams {
	o.SetIdempotencyKey(idempotencyKey)
	return o
}

// SetIdempotencyKey adds the idempotencyKey to the request account transfer params
func (o *RequestAccountTransferParams) SetIdempotencyKey(idempotencyKey string) {
	o.IdempotencyKey = idempotencyKey
}

// WriteToRequest writes these params to a swagger request
func (o *RequestAccountTransferParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.ContentType != nil {

		// header param Content-Type
		if err := r.SetHeaderParam("Content-Type", *o.ContentType); err != nil {
			return err
		}
	}
	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	// header param idempotency-key
	if err := r.SetHeaderParam("idempotency-key", o.IdempotencyKey); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}