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

// NewGetFilePaymentsFilePaymentIDPaymentDetailsParams creates a new GetFilePaymentsFilePaymentIDPaymentDetailsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetFilePaymentsFilePaymentIDPaymentDetailsParams() *GetFilePaymentsFilePaymentIDPaymentDetailsParams {
	return &GetFilePaymentsFilePaymentIDPaymentDetailsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetFilePaymentsFilePaymentIDPaymentDetailsParamsWithTimeout creates a new GetFilePaymentsFilePaymentIDPaymentDetailsParams object
// with the ability to set a timeout on a request.
func NewGetFilePaymentsFilePaymentIDPaymentDetailsParamsWithTimeout(timeout time.Duration) *GetFilePaymentsFilePaymentIDPaymentDetailsParams {
	return &GetFilePaymentsFilePaymentIDPaymentDetailsParams{
		timeout: timeout,
	}
}

// NewGetFilePaymentsFilePaymentIDPaymentDetailsParamsWithContext creates a new GetFilePaymentsFilePaymentIDPaymentDetailsParams object
// with the ability to set a context for a request.
func NewGetFilePaymentsFilePaymentIDPaymentDetailsParamsWithContext(ctx context.Context) *GetFilePaymentsFilePaymentIDPaymentDetailsParams {
	return &GetFilePaymentsFilePaymentIDPaymentDetailsParams{
		Context: ctx,
	}
}

// NewGetFilePaymentsFilePaymentIDPaymentDetailsParamsWithHTTPClient creates a new GetFilePaymentsFilePaymentIDPaymentDetailsParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetFilePaymentsFilePaymentIDPaymentDetailsParamsWithHTTPClient(client *http.Client) *GetFilePaymentsFilePaymentIDPaymentDetailsParams {
	return &GetFilePaymentsFilePaymentIDPaymentDetailsParams{
		HTTPClient: client,
	}
}

/* GetFilePaymentsFilePaymentIDPaymentDetailsParams contains all the parameters to send to the API endpoint
   for the get file payments file payment Id payment details operation.

   Typically these are written to a http.Request.
*/
type GetFilePaymentsFilePaymentIDPaymentDetailsParams struct {

	/* Authorization.

	   An Authorisation Token as per https://tools.ietf.org/html/rfc6750
	*/
	Authorization string

	/* FilePaymentID.

	   FilePaymentId
	*/
	FilePaymentID string

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

// WithDefaults hydrates default values in the get file payments file payment Id payment details params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetFilePaymentsFilePaymentIDPaymentDetailsParams) WithDefaults() *GetFilePaymentsFilePaymentIDPaymentDetailsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get file payments file payment Id payment details params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetFilePaymentsFilePaymentIDPaymentDetailsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get file payments file payment Id payment details params
func (o *GetFilePaymentsFilePaymentIDPaymentDetailsParams) WithTimeout(timeout time.Duration) *GetFilePaymentsFilePaymentIDPaymentDetailsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get file payments file payment Id payment details params
func (o *GetFilePaymentsFilePaymentIDPaymentDetailsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get file payments file payment Id payment details params
func (o *GetFilePaymentsFilePaymentIDPaymentDetailsParams) WithContext(ctx context.Context) *GetFilePaymentsFilePaymentIDPaymentDetailsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get file payments file payment Id payment details params
func (o *GetFilePaymentsFilePaymentIDPaymentDetailsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get file payments file payment Id payment details params
func (o *GetFilePaymentsFilePaymentIDPaymentDetailsParams) WithHTTPClient(client *http.Client) *GetFilePaymentsFilePaymentIDPaymentDetailsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get file payments file payment Id payment details params
func (o *GetFilePaymentsFilePaymentIDPaymentDetailsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAuthorization adds the authorization to the get file payments file payment Id payment details params
func (o *GetFilePaymentsFilePaymentIDPaymentDetailsParams) WithAuthorization(authorization string) *GetFilePaymentsFilePaymentIDPaymentDetailsParams {
	o.SetAuthorization(authorization)
	return o
}

// SetAuthorization adds the authorization to the get file payments file payment Id payment details params
func (o *GetFilePaymentsFilePaymentIDPaymentDetailsParams) SetAuthorization(authorization string) {
	o.Authorization = authorization
}

// WithFilePaymentID adds the filePaymentID to the get file payments file payment Id payment details params
func (o *GetFilePaymentsFilePaymentIDPaymentDetailsParams) WithFilePaymentID(filePaymentID string) *GetFilePaymentsFilePaymentIDPaymentDetailsParams {
	o.SetFilePaymentID(filePaymentID)
	return o
}

// SetFilePaymentID adds the filePaymentId to the get file payments file payment Id payment details params
func (o *GetFilePaymentsFilePaymentIDPaymentDetailsParams) SetFilePaymentID(filePaymentID string) {
	o.FilePaymentID = filePaymentID
}

// WithXCustomerUserAgent adds the xCustomerUserAgent to the get file payments file payment Id payment details params
func (o *GetFilePaymentsFilePaymentIDPaymentDetailsParams) WithXCustomerUserAgent(xCustomerUserAgent *string) *GetFilePaymentsFilePaymentIDPaymentDetailsParams {
	o.SetXCustomerUserAgent(xCustomerUserAgent)
	return o
}

// SetXCustomerUserAgent adds the xCustomerUserAgent to the get file payments file payment Id payment details params
func (o *GetFilePaymentsFilePaymentIDPaymentDetailsParams) SetXCustomerUserAgent(xCustomerUserAgent *string) {
	o.XCustomerUserAgent = xCustomerUserAgent
}

// WithXFapiAuthDate adds the xFapiAuthDate to the get file payments file payment Id payment details params
func (o *GetFilePaymentsFilePaymentIDPaymentDetailsParams) WithXFapiAuthDate(xFapiAuthDate *string) *GetFilePaymentsFilePaymentIDPaymentDetailsParams {
	o.SetXFapiAuthDate(xFapiAuthDate)
	return o
}

// SetXFapiAuthDate adds the xFapiAuthDate to the get file payments file payment Id payment details params
func (o *GetFilePaymentsFilePaymentIDPaymentDetailsParams) SetXFapiAuthDate(xFapiAuthDate *string) {
	o.XFapiAuthDate = xFapiAuthDate
}

// WithXFapiCustomerIPAddress adds the xFapiCustomerIPAddress to the get file payments file payment Id payment details params
func (o *GetFilePaymentsFilePaymentIDPaymentDetailsParams) WithXFapiCustomerIPAddress(xFapiCustomerIPAddress *string) *GetFilePaymentsFilePaymentIDPaymentDetailsParams {
	o.SetXFapiCustomerIPAddress(xFapiCustomerIPAddress)
	return o
}

// SetXFapiCustomerIPAddress adds the xFapiCustomerIpAddress to the get file payments file payment Id payment details params
func (o *GetFilePaymentsFilePaymentIDPaymentDetailsParams) SetXFapiCustomerIPAddress(xFapiCustomerIPAddress *string) {
	o.XFapiCustomerIPAddress = xFapiCustomerIPAddress
}

// WithXFapiInteractionID adds the xFapiInteractionID to the get file payments file payment Id payment details params
func (o *GetFilePaymentsFilePaymentIDPaymentDetailsParams) WithXFapiInteractionID(xFapiInteractionID *string) *GetFilePaymentsFilePaymentIDPaymentDetailsParams {
	o.SetXFapiInteractionID(xFapiInteractionID)
	return o
}

// SetXFapiInteractionID adds the xFapiInteractionId to the get file payments file payment Id payment details params
func (o *GetFilePaymentsFilePaymentIDPaymentDetailsParams) SetXFapiInteractionID(xFapiInteractionID *string) {
	o.XFapiInteractionID = xFapiInteractionID
}

// WriteToRequest writes these params to a swagger request
func (o *GetFilePaymentsFilePaymentIDPaymentDetailsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// header param Authorization
	if err := r.SetHeaderParam("Authorization", o.Authorization); err != nil {
		return err
	}

	// path param FilePaymentId
	if err := r.SetPathParam("FilePaymentId", o.FilePaymentID); err != nil {
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