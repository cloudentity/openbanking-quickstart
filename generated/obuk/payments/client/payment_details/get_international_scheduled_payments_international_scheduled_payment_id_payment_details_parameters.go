// Code generated by go-swagger; DO NOT EDIT.

package payment_details

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

// NewGetInternationalScheduledPaymentsInternationalScheduledPaymentIDPaymentDetailsParams creates a new GetInternationalScheduledPaymentsInternationalScheduledPaymentIDPaymentDetailsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetInternationalScheduledPaymentsInternationalScheduledPaymentIDPaymentDetailsParams() *GetInternationalScheduledPaymentsInternationalScheduledPaymentIDPaymentDetailsParams {
	return &GetInternationalScheduledPaymentsInternationalScheduledPaymentIDPaymentDetailsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetInternationalScheduledPaymentsInternationalScheduledPaymentIDPaymentDetailsParamsWithTimeout creates a new GetInternationalScheduledPaymentsInternationalScheduledPaymentIDPaymentDetailsParams object
// with the ability to set a timeout on a request.
func NewGetInternationalScheduledPaymentsInternationalScheduledPaymentIDPaymentDetailsParamsWithTimeout(timeout time.Duration) *GetInternationalScheduledPaymentsInternationalScheduledPaymentIDPaymentDetailsParams {
	return &GetInternationalScheduledPaymentsInternationalScheduledPaymentIDPaymentDetailsParams{
		timeout: timeout,
	}
}

// NewGetInternationalScheduledPaymentsInternationalScheduledPaymentIDPaymentDetailsParamsWithContext creates a new GetInternationalScheduledPaymentsInternationalScheduledPaymentIDPaymentDetailsParams object
// with the ability to set a context for a request.
func NewGetInternationalScheduledPaymentsInternationalScheduledPaymentIDPaymentDetailsParamsWithContext(ctx context.Context) *GetInternationalScheduledPaymentsInternationalScheduledPaymentIDPaymentDetailsParams {
	return &GetInternationalScheduledPaymentsInternationalScheduledPaymentIDPaymentDetailsParams{
		Context: ctx,
	}
}

// NewGetInternationalScheduledPaymentsInternationalScheduledPaymentIDPaymentDetailsParamsWithHTTPClient creates a new GetInternationalScheduledPaymentsInternationalScheduledPaymentIDPaymentDetailsParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetInternationalScheduledPaymentsInternationalScheduledPaymentIDPaymentDetailsParamsWithHTTPClient(client *http.Client) *GetInternationalScheduledPaymentsInternationalScheduledPaymentIDPaymentDetailsParams {
	return &GetInternationalScheduledPaymentsInternationalScheduledPaymentIDPaymentDetailsParams{
		HTTPClient: client,
	}
}

/* GetInternationalScheduledPaymentsInternationalScheduledPaymentIDPaymentDetailsParams contains all the parameters to send to the API endpoint
   for the get international scheduled payments international scheduled payment Id payment details operation.

   Typically these are written to a http.Request.
*/
type GetInternationalScheduledPaymentsInternationalScheduledPaymentIDPaymentDetailsParams struct {

	/* Authorization.

	   An Authorisation Token as per https://tools.ietf.org/html/rfc6750
	*/
	Authorization string

	/* InternationalScheduledPaymentID.

	   InternationalScheduledPaymentId
	*/
	InternationalScheduledPaymentID string

	/* XCustomerUserAgent.

	   Indicates the user-agent that the PSU is using.
	*/
	XCustomerUserAgent *string

	/* XFapiAuthDate.

	     The time when the PSU last logged in with the TPP.
	All dates in the HTTP headers are represented as RFC 7231 Full Dates. An example is below:
	Sun, 10 Sep 2017 19:43:31 UTC
	*/
	XFapiAuthDate *string

	/* XFapiCustomerIPAddress.

	   The PSU's IP address if the PSU is currently logged in with the TPP.
	*/
	XFapiCustomerIPAddress *string

	/* XFapiInteractionID.

	   An RFC4122 UID used as a correlation id.
	*/
	XFapiInteractionID *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get international scheduled payments international scheduled payment Id payment details params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetInternationalScheduledPaymentsInternationalScheduledPaymentIDPaymentDetailsParams) WithDefaults() *GetInternationalScheduledPaymentsInternationalScheduledPaymentIDPaymentDetailsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get international scheduled payments international scheduled payment Id payment details params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetInternationalScheduledPaymentsInternationalScheduledPaymentIDPaymentDetailsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get international scheduled payments international scheduled payment Id payment details params
func (o *GetInternationalScheduledPaymentsInternationalScheduledPaymentIDPaymentDetailsParams) WithTimeout(timeout time.Duration) *GetInternationalScheduledPaymentsInternationalScheduledPaymentIDPaymentDetailsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get international scheduled payments international scheduled payment Id payment details params
func (o *GetInternationalScheduledPaymentsInternationalScheduledPaymentIDPaymentDetailsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get international scheduled payments international scheduled payment Id payment details params
func (o *GetInternationalScheduledPaymentsInternationalScheduledPaymentIDPaymentDetailsParams) WithContext(ctx context.Context) *GetInternationalScheduledPaymentsInternationalScheduledPaymentIDPaymentDetailsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get international scheduled payments international scheduled payment Id payment details params
func (o *GetInternationalScheduledPaymentsInternationalScheduledPaymentIDPaymentDetailsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get international scheduled payments international scheduled payment Id payment details params
func (o *GetInternationalScheduledPaymentsInternationalScheduledPaymentIDPaymentDetailsParams) WithHTTPClient(client *http.Client) *GetInternationalScheduledPaymentsInternationalScheduledPaymentIDPaymentDetailsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get international scheduled payments international scheduled payment Id payment details params
func (o *GetInternationalScheduledPaymentsInternationalScheduledPaymentIDPaymentDetailsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAuthorization adds the authorization to the get international scheduled payments international scheduled payment Id payment details params
func (o *GetInternationalScheduledPaymentsInternationalScheduledPaymentIDPaymentDetailsParams) WithAuthorization(authorization string) *GetInternationalScheduledPaymentsInternationalScheduledPaymentIDPaymentDetailsParams {
	o.SetAuthorization(authorization)
	return o
}

// SetAuthorization adds the authorization to the get international scheduled payments international scheduled payment Id payment details params
func (o *GetInternationalScheduledPaymentsInternationalScheduledPaymentIDPaymentDetailsParams) SetAuthorization(authorization string) {
	o.Authorization = authorization
}

// WithInternationalScheduledPaymentID adds the internationalScheduledPaymentID to the get international scheduled payments international scheduled payment Id payment details params
func (o *GetInternationalScheduledPaymentsInternationalScheduledPaymentIDPaymentDetailsParams) WithInternationalScheduledPaymentID(internationalScheduledPaymentID string) *GetInternationalScheduledPaymentsInternationalScheduledPaymentIDPaymentDetailsParams {
	o.SetInternationalScheduledPaymentID(internationalScheduledPaymentID)
	return o
}

// SetInternationalScheduledPaymentID adds the internationalScheduledPaymentId to the get international scheduled payments international scheduled payment Id payment details params
func (o *GetInternationalScheduledPaymentsInternationalScheduledPaymentIDPaymentDetailsParams) SetInternationalScheduledPaymentID(internationalScheduledPaymentID string) {
	o.InternationalScheduledPaymentID = internationalScheduledPaymentID
}

// WithXCustomerUserAgent adds the xCustomerUserAgent to the get international scheduled payments international scheduled payment Id payment details params
func (o *GetInternationalScheduledPaymentsInternationalScheduledPaymentIDPaymentDetailsParams) WithXCustomerUserAgent(xCustomerUserAgent *string) *GetInternationalScheduledPaymentsInternationalScheduledPaymentIDPaymentDetailsParams {
	o.SetXCustomerUserAgent(xCustomerUserAgent)
	return o
}

// SetXCustomerUserAgent adds the xCustomerUserAgent to the get international scheduled payments international scheduled payment Id payment details params
func (o *GetInternationalScheduledPaymentsInternationalScheduledPaymentIDPaymentDetailsParams) SetXCustomerUserAgent(xCustomerUserAgent *string) {
	o.XCustomerUserAgent = xCustomerUserAgent
}

// WithXFapiAuthDate adds the xFapiAuthDate to the get international scheduled payments international scheduled payment Id payment details params
func (o *GetInternationalScheduledPaymentsInternationalScheduledPaymentIDPaymentDetailsParams) WithXFapiAuthDate(xFapiAuthDate *string) *GetInternationalScheduledPaymentsInternationalScheduledPaymentIDPaymentDetailsParams {
	o.SetXFapiAuthDate(xFapiAuthDate)
	return o
}

// SetXFapiAuthDate adds the xFapiAuthDate to the get international scheduled payments international scheduled payment Id payment details params
func (o *GetInternationalScheduledPaymentsInternationalScheduledPaymentIDPaymentDetailsParams) SetXFapiAuthDate(xFapiAuthDate *string) {
	o.XFapiAuthDate = xFapiAuthDate
}

// WithXFapiCustomerIPAddress adds the xFapiCustomerIPAddress to the get international scheduled payments international scheduled payment Id payment details params
func (o *GetInternationalScheduledPaymentsInternationalScheduledPaymentIDPaymentDetailsParams) WithXFapiCustomerIPAddress(xFapiCustomerIPAddress *string) *GetInternationalScheduledPaymentsInternationalScheduledPaymentIDPaymentDetailsParams {
	o.SetXFapiCustomerIPAddress(xFapiCustomerIPAddress)
	return o
}

// SetXFapiCustomerIPAddress adds the xFapiCustomerIpAddress to the get international scheduled payments international scheduled payment Id payment details params
func (o *GetInternationalScheduledPaymentsInternationalScheduledPaymentIDPaymentDetailsParams) SetXFapiCustomerIPAddress(xFapiCustomerIPAddress *string) {
	o.XFapiCustomerIPAddress = xFapiCustomerIPAddress
}

// WithXFapiInteractionID adds the xFapiInteractionID to the get international scheduled payments international scheduled payment Id payment details params
func (o *GetInternationalScheduledPaymentsInternationalScheduledPaymentIDPaymentDetailsParams) WithXFapiInteractionID(xFapiInteractionID *string) *GetInternationalScheduledPaymentsInternationalScheduledPaymentIDPaymentDetailsParams {
	o.SetXFapiInteractionID(xFapiInteractionID)
	return o
}

// SetXFapiInteractionID adds the xFapiInteractionId to the get international scheduled payments international scheduled payment Id payment details params
func (o *GetInternationalScheduledPaymentsInternationalScheduledPaymentIDPaymentDetailsParams) SetXFapiInteractionID(xFapiInteractionID *string) {
	o.XFapiInteractionID = xFapiInteractionID
}

// WriteToRequest writes these params to a swagger request
func (o *GetInternationalScheduledPaymentsInternationalScheduledPaymentIDPaymentDetailsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// header param Authorization
	if err := r.SetHeaderParam("Authorization", o.Authorization); err != nil {
		return err
	}

	// path param InternationalScheduledPaymentId
	if err := r.SetPathParam("InternationalScheduledPaymentId", o.InternationalScheduledPaymentID); err != nil {
		return err
	}

	if o.XCustomerUserAgent != nil {

		// header param x-customer-user-agent
		if err := r.SetHeaderParam("x-customer-user-agent", *o.XCustomerUserAgent); err != nil {
			return err
		}
	}

	if o.XFapiAuthDate != nil {

		// header param x-fapi-auth-date
		if err := r.SetHeaderParam("x-fapi-auth-date", *o.XFapiAuthDate); err != nil {
			return err
		}
	}

	if o.XFapiCustomerIPAddress != nil {

		// header param x-fapi-customer-ip-address
		if err := r.SetHeaderParam("x-fapi-customer-ip-address", *o.XFapiCustomerIPAddress); err != nil {
			return err
		}
	}

	if o.XFapiInteractionID != nil {

		// header param x-fapi-interaction-id
		if err := r.SetHeaderParam("x-fapi-interaction-id", *o.XFapiInteractionID); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}