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

// NewGetDomesticScheduledPaymentsDomesticScheduledPaymentIDPaymentDetailsParams creates a new GetDomesticScheduledPaymentsDomesticScheduledPaymentIDPaymentDetailsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetDomesticScheduledPaymentsDomesticScheduledPaymentIDPaymentDetailsParams() *GetDomesticScheduledPaymentsDomesticScheduledPaymentIDPaymentDetailsParams {
	return &GetDomesticScheduledPaymentsDomesticScheduledPaymentIDPaymentDetailsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetDomesticScheduledPaymentsDomesticScheduledPaymentIDPaymentDetailsParamsWithTimeout creates a new GetDomesticScheduledPaymentsDomesticScheduledPaymentIDPaymentDetailsParams object
// with the ability to set a timeout on a request.
func NewGetDomesticScheduledPaymentsDomesticScheduledPaymentIDPaymentDetailsParamsWithTimeout(timeout time.Duration) *GetDomesticScheduledPaymentsDomesticScheduledPaymentIDPaymentDetailsParams {
	return &GetDomesticScheduledPaymentsDomesticScheduledPaymentIDPaymentDetailsParams{
		timeout: timeout,
	}
}

// NewGetDomesticScheduledPaymentsDomesticScheduledPaymentIDPaymentDetailsParamsWithContext creates a new GetDomesticScheduledPaymentsDomesticScheduledPaymentIDPaymentDetailsParams object
// with the ability to set a context for a request.
func NewGetDomesticScheduledPaymentsDomesticScheduledPaymentIDPaymentDetailsParamsWithContext(ctx context.Context) *GetDomesticScheduledPaymentsDomesticScheduledPaymentIDPaymentDetailsParams {
	return &GetDomesticScheduledPaymentsDomesticScheduledPaymentIDPaymentDetailsParams{
		Context: ctx,
	}
}

// NewGetDomesticScheduledPaymentsDomesticScheduledPaymentIDPaymentDetailsParamsWithHTTPClient creates a new GetDomesticScheduledPaymentsDomesticScheduledPaymentIDPaymentDetailsParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetDomesticScheduledPaymentsDomesticScheduledPaymentIDPaymentDetailsParamsWithHTTPClient(client *http.Client) *GetDomesticScheduledPaymentsDomesticScheduledPaymentIDPaymentDetailsParams {
	return &GetDomesticScheduledPaymentsDomesticScheduledPaymentIDPaymentDetailsParams{
		HTTPClient: client,
	}
}

/* GetDomesticScheduledPaymentsDomesticScheduledPaymentIDPaymentDetailsParams contains all the parameters to send to the API endpoint
   for the get domestic scheduled payments domestic scheduled payment Id payment details operation.

   Typically these are written to a http.Request.
*/
type GetDomesticScheduledPaymentsDomesticScheduledPaymentIDPaymentDetailsParams struct {

	/* Authorization.

	   An Authorisation Token as per https://tools.ietf.org/html/rfc6750
	*/
	Authorization string

	/* DomesticScheduledPaymentID.

	   DomesticScheduledPaymentId
	*/
	DomesticScheduledPaymentID string

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

// WithDefaults hydrates default values in the get domestic scheduled payments domestic scheduled payment Id payment details params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetDomesticScheduledPaymentsDomesticScheduledPaymentIDPaymentDetailsParams) WithDefaults() *GetDomesticScheduledPaymentsDomesticScheduledPaymentIDPaymentDetailsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get domestic scheduled payments domestic scheduled payment Id payment details params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetDomesticScheduledPaymentsDomesticScheduledPaymentIDPaymentDetailsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get domestic scheduled payments domestic scheduled payment Id payment details params
func (o *GetDomesticScheduledPaymentsDomesticScheduledPaymentIDPaymentDetailsParams) WithTimeout(timeout time.Duration) *GetDomesticScheduledPaymentsDomesticScheduledPaymentIDPaymentDetailsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get domestic scheduled payments domestic scheduled payment Id payment details params
func (o *GetDomesticScheduledPaymentsDomesticScheduledPaymentIDPaymentDetailsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get domestic scheduled payments domestic scheduled payment Id payment details params
func (o *GetDomesticScheduledPaymentsDomesticScheduledPaymentIDPaymentDetailsParams) WithContext(ctx context.Context) *GetDomesticScheduledPaymentsDomesticScheduledPaymentIDPaymentDetailsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get domestic scheduled payments domestic scheduled payment Id payment details params
func (o *GetDomesticScheduledPaymentsDomesticScheduledPaymentIDPaymentDetailsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get domestic scheduled payments domestic scheduled payment Id payment details params
func (o *GetDomesticScheduledPaymentsDomesticScheduledPaymentIDPaymentDetailsParams) WithHTTPClient(client *http.Client) *GetDomesticScheduledPaymentsDomesticScheduledPaymentIDPaymentDetailsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get domestic scheduled payments domestic scheduled payment Id payment details params
func (o *GetDomesticScheduledPaymentsDomesticScheduledPaymentIDPaymentDetailsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAuthorization adds the authorization to the get domestic scheduled payments domestic scheduled payment Id payment details params
func (o *GetDomesticScheduledPaymentsDomesticScheduledPaymentIDPaymentDetailsParams) WithAuthorization(authorization string) *GetDomesticScheduledPaymentsDomesticScheduledPaymentIDPaymentDetailsParams {
	o.SetAuthorization(authorization)
	return o
}

// SetAuthorization adds the authorization to the get domestic scheduled payments domestic scheduled payment Id payment details params
func (o *GetDomesticScheduledPaymentsDomesticScheduledPaymentIDPaymentDetailsParams) SetAuthorization(authorization string) {
	o.Authorization = authorization
}

// WithDomesticScheduledPaymentID adds the domesticScheduledPaymentID to the get domestic scheduled payments domestic scheduled payment Id payment details params
func (o *GetDomesticScheduledPaymentsDomesticScheduledPaymentIDPaymentDetailsParams) WithDomesticScheduledPaymentID(domesticScheduledPaymentID string) *GetDomesticScheduledPaymentsDomesticScheduledPaymentIDPaymentDetailsParams {
	o.SetDomesticScheduledPaymentID(domesticScheduledPaymentID)
	return o
}

// SetDomesticScheduledPaymentID adds the domesticScheduledPaymentId to the get domestic scheduled payments domestic scheduled payment Id payment details params
func (o *GetDomesticScheduledPaymentsDomesticScheduledPaymentIDPaymentDetailsParams) SetDomesticScheduledPaymentID(domesticScheduledPaymentID string) {
	o.DomesticScheduledPaymentID = domesticScheduledPaymentID
}

// WithXCustomerUserAgent adds the xCustomerUserAgent to the get domestic scheduled payments domestic scheduled payment Id payment details params
func (o *GetDomesticScheduledPaymentsDomesticScheduledPaymentIDPaymentDetailsParams) WithXCustomerUserAgent(xCustomerUserAgent *string) *GetDomesticScheduledPaymentsDomesticScheduledPaymentIDPaymentDetailsParams {
	o.SetXCustomerUserAgent(xCustomerUserAgent)
	return o
}

// SetXCustomerUserAgent adds the xCustomerUserAgent to the get domestic scheduled payments domestic scheduled payment Id payment details params
func (o *GetDomesticScheduledPaymentsDomesticScheduledPaymentIDPaymentDetailsParams) SetXCustomerUserAgent(xCustomerUserAgent *string) {
	o.XCustomerUserAgent = xCustomerUserAgent
}

// WithXFapiAuthDate adds the xFapiAuthDate to the get domestic scheduled payments domestic scheduled payment Id payment details params
func (o *GetDomesticScheduledPaymentsDomesticScheduledPaymentIDPaymentDetailsParams) WithXFapiAuthDate(xFapiAuthDate *string) *GetDomesticScheduledPaymentsDomesticScheduledPaymentIDPaymentDetailsParams {
	o.SetXFapiAuthDate(xFapiAuthDate)
	return o
}

// SetXFapiAuthDate adds the xFapiAuthDate to the get domestic scheduled payments domestic scheduled payment Id payment details params
func (o *GetDomesticScheduledPaymentsDomesticScheduledPaymentIDPaymentDetailsParams) SetXFapiAuthDate(xFapiAuthDate *string) {
	o.XFapiAuthDate = xFapiAuthDate
}

// WithXFapiCustomerIPAddress adds the xFapiCustomerIPAddress to the get domestic scheduled payments domestic scheduled payment Id payment details params
func (o *GetDomesticScheduledPaymentsDomesticScheduledPaymentIDPaymentDetailsParams) WithXFapiCustomerIPAddress(xFapiCustomerIPAddress *string) *GetDomesticScheduledPaymentsDomesticScheduledPaymentIDPaymentDetailsParams {
	o.SetXFapiCustomerIPAddress(xFapiCustomerIPAddress)
	return o
}

// SetXFapiCustomerIPAddress adds the xFapiCustomerIpAddress to the get domestic scheduled payments domestic scheduled payment Id payment details params
func (o *GetDomesticScheduledPaymentsDomesticScheduledPaymentIDPaymentDetailsParams) SetXFapiCustomerIPAddress(xFapiCustomerIPAddress *string) {
	o.XFapiCustomerIPAddress = xFapiCustomerIPAddress
}

// WithXFapiInteractionID adds the xFapiInteractionID to the get domestic scheduled payments domestic scheduled payment Id payment details params
func (o *GetDomesticScheduledPaymentsDomesticScheduledPaymentIDPaymentDetailsParams) WithXFapiInteractionID(xFapiInteractionID *string) *GetDomesticScheduledPaymentsDomesticScheduledPaymentIDPaymentDetailsParams {
	o.SetXFapiInteractionID(xFapiInteractionID)
	return o
}

// SetXFapiInteractionID adds the xFapiInteractionId to the get domestic scheduled payments domestic scheduled payment Id payment details params
func (o *GetDomesticScheduledPaymentsDomesticScheduledPaymentIDPaymentDetailsParams) SetXFapiInteractionID(xFapiInteractionID *string) {
	o.XFapiInteractionID = xFapiInteractionID
}

// WriteToRequest writes these params to a swagger request
func (o *GetDomesticScheduledPaymentsDomesticScheduledPaymentIDPaymentDetailsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// header param Authorization
	if err := r.SetHeaderParam("Authorization", o.Authorization); err != nil {
		return err
	}

	// path param DomesticScheduledPaymentId
	if err := r.SetPathParam("DomesticScheduledPaymentId", o.DomesticScheduledPaymentID); err != nil {
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
