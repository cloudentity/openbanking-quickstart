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
	"github.com/go-openapi/swag"
)

// NewSearchForRecurringPaymentsParams creates a new SearchForRecurringPaymentsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewSearchForRecurringPaymentsParams() *SearchForRecurringPaymentsParams {
	return &SearchForRecurringPaymentsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewSearchForRecurringPaymentsParamsWithTimeout creates a new SearchForRecurringPaymentsParams object
// with the ability to set a timeout on a request.
func NewSearchForRecurringPaymentsParamsWithTimeout(timeout time.Duration) *SearchForRecurringPaymentsParams {
	return &SearchForRecurringPaymentsParams{
		timeout: timeout,
	}
}

// NewSearchForRecurringPaymentsParamsWithContext creates a new SearchForRecurringPaymentsParams object
// with the ability to set a context for a request.
func NewSearchForRecurringPaymentsParamsWithContext(ctx context.Context) *SearchForRecurringPaymentsParams {
	return &SearchForRecurringPaymentsParams{
		Context: ctx,
	}
}

// NewSearchForRecurringPaymentsParamsWithHTTPClient creates a new SearchForRecurringPaymentsParams object
// with the ability to set a custom HTTPClient for a request.
func NewSearchForRecurringPaymentsParamsWithHTTPClient(client *http.Client) *SearchForRecurringPaymentsParams {
	return &SearchForRecurringPaymentsParams{
		HTTPClient: client,
	}
}

/* SearchForRecurringPaymentsParams contains all the parameters to send to the API endpoint
   for the search for recurring payments operation.

   Typically these are written to a http.Request.
*/
type SearchForRecurringPaymentsParams struct {

	/* Limit.

	   Number of elements that the consumer wishes to receive. Providers should implement reasonable default and maximum values

	   Format: int32
	*/
	Limit *int32

	/* Offset.

	   Opaque cursor used by the provider to send the next set of records
	*/
	Offset *string

	/* UpdatedSince.

	   Return items that have been created or updated since the update id
	*/
	UpdatedSince *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the search for recurring payments params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SearchForRecurringPaymentsParams) WithDefaults() *SearchForRecurringPaymentsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the search for recurring payments params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SearchForRecurringPaymentsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the search for recurring payments params
func (o *SearchForRecurringPaymentsParams) WithTimeout(timeout time.Duration) *SearchForRecurringPaymentsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the search for recurring payments params
func (o *SearchForRecurringPaymentsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the search for recurring payments params
func (o *SearchForRecurringPaymentsParams) WithContext(ctx context.Context) *SearchForRecurringPaymentsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the search for recurring payments params
func (o *SearchForRecurringPaymentsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the search for recurring payments params
func (o *SearchForRecurringPaymentsParams) WithHTTPClient(client *http.Client) *SearchForRecurringPaymentsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the search for recurring payments params
func (o *SearchForRecurringPaymentsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithLimit adds the limit to the search for recurring payments params
func (o *SearchForRecurringPaymentsParams) WithLimit(limit *int32) *SearchForRecurringPaymentsParams {
	o.SetLimit(limit)
	return o
}

// SetLimit adds the limit to the search for recurring payments params
func (o *SearchForRecurringPaymentsParams) SetLimit(limit *int32) {
	o.Limit = limit
}

// WithOffset adds the offset to the search for recurring payments params
func (o *SearchForRecurringPaymentsParams) WithOffset(offset *string) *SearchForRecurringPaymentsParams {
	o.SetOffset(offset)
	return o
}

// SetOffset adds the offset to the search for recurring payments params
func (o *SearchForRecurringPaymentsParams) SetOffset(offset *string) {
	o.Offset = offset
}

// WithUpdatedSince adds the updatedSince to the search for recurring payments params
func (o *SearchForRecurringPaymentsParams) WithUpdatedSince(updatedSince *string) *SearchForRecurringPaymentsParams {
	o.SetUpdatedSince(updatedSince)
	return o
}

// SetUpdatedSince adds the updatedSince to the search for recurring payments params
func (o *SearchForRecurringPaymentsParams) SetUpdatedSince(updatedSince *string) {
	o.UpdatedSince = updatedSince
}

// WriteToRequest writes these params to a swagger request
func (o *SearchForRecurringPaymentsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Limit != nil {

		// query param limit
		var qrLimit int32

		if o.Limit != nil {
			qrLimit = *o.Limit
		}
		qLimit := swag.FormatInt32(qrLimit)
		if qLimit != "" {

			if err := r.SetQueryParam("limit", qLimit); err != nil {
				return err
			}
		}
	}

	if o.Offset != nil {

		// query param offset
		var qrOffset string

		if o.Offset != nil {
			qrOffset = *o.Offset
		}
		qOffset := qrOffset
		if qOffset != "" {

			if err := r.SetQueryParam("offset", qOffset); err != nil {
				return err
			}
		}
	}

	if o.UpdatedSince != nil {

		// query param updatedSince
		var qrUpdatedSince string

		if o.UpdatedSince != nil {
			qrUpdatedSince = *o.UpdatedSince
		}
		qUpdatedSince := qrUpdatedSince
		if qUpdatedSince != "" {

			if err := r.SetQueryParam("updatedSince", qUpdatedSince); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}