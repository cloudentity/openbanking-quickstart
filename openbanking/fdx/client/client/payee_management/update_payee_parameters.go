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

	"github.com/cloudentity/openbanking-quickstart/openbanking/fdx/client/models"
)

// NewUpdatePayeeParams creates a new UpdatePayeeParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewUpdatePayeeParams() *UpdatePayeeParams {
	return &UpdatePayeeParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewUpdatePayeeParamsWithTimeout creates a new UpdatePayeeParams object
// with the ability to set a timeout on a request.
func NewUpdatePayeeParamsWithTimeout(timeout time.Duration) *UpdatePayeeParams {
	return &UpdatePayeeParams{
		timeout: timeout,
	}
}

// NewUpdatePayeeParamsWithContext creates a new UpdatePayeeParams object
// with the ability to set a context for a request.
func NewUpdatePayeeParamsWithContext(ctx context.Context) *UpdatePayeeParams {
	return &UpdatePayeeParams{
		Context: ctx,
	}
}

// NewUpdatePayeeParamsWithHTTPClient creates a new UpdatePayeeParams object
// with the ability to set a custom HTTPClient for a request.
func NewUpdatePayeeParamsWithHTTPClient(client *http.Client) *UpdatePayeeParams {
	return &UpdatePayeeParams{
		HTTPClient: client,
	}
}

/* UpdatePayeeParams contains all the parameters to send to the API endpoint
   for the update payee operation.

   Typically these are written to a http.Request.
*/
type UpdatePayeeParams struct {

	// ContentType.
	ContentType *string

	// Body.
	Body *models.PayeeForUpdateentity

	/* IdempotencyKey.

	   Used to de-duplicate requests
	*/
	IdempotencyKey string

	/* PayeeID.

	   Payee Identifier. Uniquely identifies a payee
	*/
	PayeeID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the update payee params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdatePayeeParams) WithDefaults() *UpdatePayeeParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the update payee params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdatePayeeParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the update payee params
func (o *UpdatePayeeParams) WithTimeout(timeout time.Duration) *UpdatePayeeParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update payee params
func (o *UpdatePayeeParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update payee params
func (o *UpdatePayeeParams) WithContext(ctx context.Context) *UpdatePayeeParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update payee params
func (o *UpdatePayeeParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update payee params
func (o *UpdatePayeeParams) WithHTTPClient(client *http.Client) *UpdatePayeeParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update payee params
func (o *UpdatePayeeParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithContentType adds the contentType to the update payee params
func (o *UpdatePayeeParams) WithContentType(contentType *string) *UpdatePayeeParams {
	o.SetContentType(contentType)
	return o
}

// SetContentType adds the contentType to the update payee params
func (o *UpdatePayeeParams) SetContentType(contentType *string) {
	o.ContentType = contentType
}

// WithBody adds the body to the update payee params
func (o *UpdatePayeeParams) WithBody(body *models.PayeeForUpdateentity) *UpdatePayeeParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the update payee params
func (o *UpdatePayeeParams) SetBody(body *models.PayeeForUpdateentity) {
	o.Body = body
}

// WithIdempotencyKey adds the idempotencyKey to the update payee params
func (o *UpdatePayeeParams) WithIdempotencyKey(idempotencyKey string) *UpdatePayeeParams {
	o.SetIdempotencyKey(idempotencyKey)
	return o
}

// SetIdempotencyKey adds the idempotencyKey to the update payee params
func (o *UpdatePayeeParams) SetIdempotencyKey(idempotencyKey string) {
	o.IdempotencyKey = idempotencyKey
}

// WithPayeeID adds the payeeID to the update payee params
func (o *UpdatePayeeParams) WithPayeeID(payeeID string) *UpdatePayeeParams {
	o.SetPayeeID(payeeID)
	return o
}

// SetPayeeID adds the payeeId to the update payee params
func (o *UpdatePayeeParams) SetPayeeID(payeeID string) {
	o.PayeeID = payeeID
}

// WriteToRequest writes these params to a swagger request
func (o *UpdatePayeeParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param payeeId
	if err := r.SetPathParam("payeeId", o.PayeeID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
