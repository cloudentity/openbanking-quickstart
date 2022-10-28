// Code generated by go-swagger; DO NOT EDIT.

package domestic_standing_orders

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

// NewGetDomesticStandingOrdersDomesticStandingOrderIDParams creates a new GetDomesticStandingOrdersDomesticStandingOrderIDParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetDomesticStandingOrdersDomesticStandingOrderIDParams() *GetDomesticStandingOrdersDomesticStandingOrderIDParams {
	return &GetDomesticStandingOrdersDomesticStandingOrderIDParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetDomesticStandingOrdersDomesticStandingOrderIDParamsWithTimeout creates a new GetDomesticStandingOrdersDomesticStandingOrderIDParams object
// with the ability to set a timeout on a request.
func NewGetDomesticStandingOrdersDomesticStandingOrderIDParamsWithTimeout(timeout time.Duration) *GetDomesticStandingOrdersDomesticStandingOrderIDParams {
	return &GetDomesticStandingOrdersDomesticStandingOrderIDParams{
		timeout: timeout,
	}
}

// NewGetDomesticStandingOrdersDomesticStandingOrderIDParamsWithContext creates a new GetDomesticStandingOrdersDomesticStandingOrderIDParams object
// with the ability to set a context for a request.
func NewGetDomesticStandingOrdersDomesticStandingOrderIDParamsWithContext(ctx context.Context) *GetDomesticStandingOrdersDomesticStandingOrderIDParams {
	return &GetDomesticStandingOrdersDomesticStandingOrderIDParams{
		Context: ctx,
	}
}

// NewGetDomesticStandingOrdersDomesticStandingOrderIDParamsWithHTTPClient creates a new GetDomesticStandingOrdersDomesticStandingOrderIDParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetDomesticStandingOrdersDomesticStandingOrderIDParamsWithHTTPClient(client *http.Client) *GetDomesticStandingOrdersDomesticStandingOrderIDParams {
	return &GetDomesticStandingOrdersDomesticStandingOrderIDParams{
		HTTPClient: client,
	}
}

/* GetDomesticStandingOrdersDomesticStandingOrderIDParams contains all the parameters to send to the API endpoint
   for the get domestic standing orders domestic standing order Id operation.

   Typically these are written to a http.Request.
*/
type GetDomesticStandingOrdersDomesticStandingOrderIDParams struct {

	/* Authorization.

	   An Authorisation Token as per https://tools.ietf.org/html/rfc6750
	*/
	Authorization string

	/* DomesticStandingOrderID.

	   DomesticStandingOrderId
	*/
	DomesticStandingOrderID string

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

// WithDefaults hydrates default values in the get domestic standing orders domestic standing order Id params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetDomesticStandingOrdersDomesticStandingOrderIDParams) WithDefaults() *GetDomesticStandingOrdersDomesticStandingOrderIDParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get domestic standing orders domestic standing order Id params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetDomesticStandingOrdersDomesticStandingOrderIDParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get domestic standing orders domestic standing order Id params
func (o *GetDomesticStandingOrdersDomesticStandingOrderIDParams) WithTimeout(timeout time.Duration) *GetDomesticStandingOrdersDomesticStandingOrderIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get domestic standing orders domestic standing order Id params
func (o *GetDomesticStandingOrdersDomesticStandingOrderIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get domestic standing orders domestic standing order Id params
func (o *GetDomesticStandingOrdersDomesticStandingOrderIDParams) WithContext(ctx context.Context) *GetDomesticStandingOrdersDomesticStandingOrderIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get domestic standing orders domestic standing order Id params
func (o *GetDomesticStandingOrdersDomesticStandingOrderIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get domestic standing orders domestic standing order Id params
func (o *GetDomesticStandingOrdersDomesticStandingOrderIDParams) WithHTTPClient(client *http.Client) *GetDomesticStandingOrdersDomesticStandingOrderIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get domestic standing orders domestic standing order Id params
func (o *GetDomesticStandingOrdersDomesticStandingOrderIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAuthorization adds the authorization to the get domestic standing orders domestic standing order Id params
func (o *GetDomesticStandingOrdersDomesticStandingOrderIDParams) WithAuthorization(authorization string) *GetDomesticStandingOrdersDomesticStandingOrderIDParams {
	o.SetAuthorization(authorization)
	return o
}

// SetAuthorization adds the authorization to the get domestic standing orders domestic standing order Id params
func (o *GetDomesticStandingOrdersDomesticStandingOrderIDParams) SetAuthorization(authorization string) {
	o.Authorization = authorization
}

// WithDomesticStandingOrderID adds the domesticStandingOrderID to the get domestic standing orders domestic standing order Id params
func (o *GetDomesticStandingOrdersDomesticStandingOrderIDParams) WithDomesticStandingOrderID(domesticStandingOrderID string) *GetDomesticStandingOrdersDomesticStandingOrderIDParams {
	o.SetDomesticStandingOrderID(domesticStandingOrderID)
	return o
}

// SetDomesticStandingOrderID adds the domesticStandingOrderId to the get domestic standing orders domestic standing order Id params
func (o *GetDomesticStandingOrdersDomesticStandingOrderIDParams) SetDomesticStandingOrderID(domesticStandingOrderID string) {
	o.DomesticStandingOrderID = domesticStandingOrderID
}

// WithXCustomerUserAgent adds the xCustomerUserAgent to the get domestic standing orders domestic standing order Id params
func (o *GetDomesticStandingOrdersDomesticStandingOrderIDParams) WithXCustomerUserAgent(xCustomerUserAgent *string) *GetDomesticStandingOrdersDomesticStandingOrderIDParams {
	o.SetXCustomerUserAgent(xCustomerUserAgent)
	return o
}

// SetXCustomerUserAgent adds the xCustomerUserAgent to the get domestic standing orders domestic standing order Id params
func (o *GetDomesticStandingOrdersDomesticStandingOrderIDParams) SetXCustomerUserAgent(xCustomerUserAgent *string) {
	o.XCustomerUserAgent = xCustomerUserAgent
}

// WithXFapiAuthDate adds the xFapiAuthDate to the get domestic standing orders domestic standing order Id params
func (o *GetDomesticStandingOrdersDomesticStandingOrderIDParams) WithXFapiAuthDate(xFapiAuthDate *string) *GetDomesticStandingOrdersDomesticStandingOrderIDParams {
	o.SetXFapiAuthDate(xFapiAuthDate)
	return o
}

// SetXFapiAuthDate adds the xFapiAuthDate to the get domestic standing orders domestic standing order Id params
func (o *GetDomesticStandingOrdersDomesticStandingOrderIDParams) SetXFapiAuthDate(xFapiAuthDate *string) {
	o.XFapiAuthDate = xFapiAuthDate
}

// WithXFapiCustomerIPAddress adds the xFapiCustomerIPAddress to the get domestic standing orders domestic standing order Id params
func (o *GetDomesticStandingOrdersDomesticStandingOrderIDParams) WithXFapiCustomerIPAddress(xFapiCustomerIPAddress *string) *GetDomesticStandingOrdersDomesticStandingOrderIDParams {
	o.SetXFapiCustomerIPAddress(xFapiCustomerIPAddress)
	return o
}

// SetXFapiCustomerIPAddress adds the xFapiCustomerIpAddress to the get domestic standing orders domestic standing order Id params
func (o *GetDomesticStandingOrdersDomesticStandingOrderIDParams) SetXFapiCustomerIPAddress(xFapiCustomerIPAddress *string) {
	o.XFapiCustomerIPAddress = xFapiCustomerIPAddress
}

// WithXFapiInteractionID adds the xFapiInteractionID to the get domestic standing orders domestic standing order Id params
func (o *GetDomesticStandingOrdersDomesticStandingOrderIDParams) WithXFapiInteractionID(xFapiInteractionID *string) *GetDomesticStandingOrdersDomesticStandingOrderIDParams {
	o.SetXFapiInteractionID(xFapiInteractionID)
	return o
}

// SetXFapiInteractionID adds the xFapiInteractionId to the get domestic standing orders domestic standing order Id params
func (o *GetDomesticStandingOrdersDomesticStandingOrderIDParams) SetXFapiInteractionID(xFapiInteractionID *string) {
	o.XFapiInteractionID = xFapiInteractionID
}

// WriteToRequest writes these params to a swagger request
func (o *GetDomesticStandingOrdersDomesticStandingOrderIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// header param Authorization
	if err := r.SetHeaderParam("Authorization", o.Authorization); err != nil {
		return err
	}

	// path param DomesticStandingOrderId
	if err := r.SetPathParam("DomesticStandingOrderId", o.DomesticStandingOrderID); err != nil {
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