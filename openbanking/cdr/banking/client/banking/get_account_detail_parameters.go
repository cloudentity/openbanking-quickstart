// Code generated by go-swagger; DO NOT EDIT.

package banking

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

// NewGetAccountDetailParams creates a new GetAccountDetailParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetAccountDetailParams() *GetAccountDetailParams {
	return &GetAccountDetailParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetAccountDetailParamsWithTimeout creates a new GetAccountDetailParams object
// with the ability to set a timeout on a request.
func NewGetAccountDetailParamsWithTimeout(timeout time.Duration) *GetAccountDetailParams {
	return &GetAccountDetailParams{
		timeout: timeout,
	}
}

// NewGetAccountDetailParamsWithContext creates a new GetAccountDetailParams object
// with the ability to set a context for a request.
func NewGetAccountDetailParamsWithContext(ctx context.Context) *GetAccountDetailParams {
	return &GetAccountDetailParams{
		Context: ctx,
	}
}

// NewGetAccountDetailParamsWithHTTPClient creates a new GetAccountDetailParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetAccountDetailParamsWithHTTPClient(client *http.Client) *GetAccountDetailParams {
	return &GetAccountDetailParams{
		HTTPClient: client,
	}
}

/* GetAccountDetailParams contains all the parameters to send to the API endpoint
   for the get account detail operation.

   Typically these are written to a http.Request.
*/
type GetAccountDetailParams struct {

	/* AccountID.

	   A tokenised identifier for the account which is unique but not shareable
	*/
	AccountID string

	/* XCdsClientHeaders.

	   The customer's original standard http headers [Base64](#common-field-types) encoded, including the original User Agent header, if the customer is currently logged in to the Data Recipient Software Product. Mandatory for customer present calls.  Not required for unattended or unauthenticated calls.
	*/
	XCdsClientHeaders *string

	/* XFapiAuthDate.

	   The time when the customer last logged in to the Data Recipient Software Product. Required for all resource calls (customer present and unattended) if the customer has logged in. Not to be included for unauthenticated calls.
	*/
	XFapiAuthDate *string

	/* XFapiCustomerIPAddress.

	   The customer's original IP address if the customer is currently logged in to the Data Recipient Software Product. The presence of this header indicates that the API is being called in a customer present context. Not to be included for unauthenticated calls.
	*/
	XFapiCustomerIPAddress *string

	/* XFapiInteractionID.

	   An [RFC4122](https://tools.ietf.org/html/rfc4122) UUID used as a correlation id. If provided, the data holder must play back this value in the x-fapi-interaction-id response header. If not provided a [RFC4122] UUID value is required to be provided in the response header to track the interaction.
	*/
	XFapiInteractionID *string

	/* XMinv.

	   Minimum version of the API end point requested by the client. Must be set to a positive integer if provided. The data holder should respond with the highest supported version between [x-min-v](#request-headers) and [x-v](#request-headers). If all versions requested are not supported then the data holder must respond with a 406 Not Acceptable.
	*/
	XMinv *string

	/* Xv.

	   Version of the API end point requested by the client. Must be set to a positive integer. The data holder should respond with the highest supported version between [x-min-v](#request-headers) and [x-v](#request-headers). If the value of [x-min-v](#request-headers) is equal to or higher than the value of [x-v](#request-headers) then the [x-min-v](#request-headers) header should be treated as absent. If all versions requested are not supported then the data holder must respond with a 406 Not Acceptable. See [HTTP Headers](#request-headers)
	*/
	Xv string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get account detail params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetAccountDetailParams) WithDefaults() *GetAccountDetailParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get account detail params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetAccountDetailParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get account detail params
func (o *GetAccountDetailParams) WithTimeout(timeout time.Duration) *GetAccountDetailParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get account detail params
func (o *GetAccountDetailParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get account detail params
func (o *GetAccountDetailParams) WithContext(ctx context.Context) *GetAccountDetailParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get account detail params
func (o *GetAccountDetailParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get account detail params
func (o *GetAccountDetailParams) WithHTTPClient(client *http.Client) *GetAccountDetailParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get account detail params
func (o *GetAccountDetailParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAccountID adds the accountID to the get account detail params
func (o *GetAccountDetailParams) WithAccountID(accountID string) *GetAccountDetailParams {
	o.SetAccountID(accountID)
	return o
}

// SetAccountID adds the accountId to the get account detail params
func (o *GetAccountDetailParams) SetAccountID(accountID string) {
	o.AccountID = accountID
}

// WithXCdsClientHeaders adds the xCdsClientHeaders to the get account detail params
func (o *GetAccountDetailParams) WithXCdsClientHeaders(xCdsClientHeaders *string) *GetAccountDetailParams {
	o.SetXCdsClientHeaders(xCdsClientHeaders)
	return o
}

// SetXCdsClientHeaders adds the xCdsClientHeaders to the get account detail params
func (o *GetAccountDetailParams) SetXCdsClientHeaders(xCdsClientHeaders *string) {
	o.XCdsClientHeaders = xCdsClientHeaders
}

// WithXFapiAuthDate adds the xFapiAuthDate to the get account detail params
func (o *GetAccountDetailParams) WithXFapiAuthDate(xFapiAuthDate *string) *GetAccountDetailParams {
	o.SetXFapiAuthDate(xFapiAuthDate)
	return o
}

// SetXFapiAuthDate adds the xFapiAuthDate to the get account detail params
func (o *GetAccountDetailParams) SetXFapiAuthDate(xFapiAuthDate *string) {
	o.XFapiAuthDate = xFapiAuthDate
}

// WithXFapiCustomerIPAddress adds the xFapiCustomerIPAddress to the get account detail params
func (o *GetAccountDetailParams) WithXFapiCustomerIPAddress(xFapiCustomerIPAddress *string) *GetAccountDetailParams {
	o.SetXFapiCustomerIPAddress(xFapiCustomerIPAddress)
	return o
}

// SetXFapiCustomerIPAddress adds the xFapiCustomerIpAddress to the get account detail params
func (o *GetAccountDetailParams) SetXFapiCustomerIPAddress(xFapiCustomerIPAddress *string) {
	o.XFapiCustomerIPAddress = xFapiCustomerIPAddress
}

// WithXFapiInteractionID adds the xFapiInteractionID to the get account detail params
func (o *GetAccountDetailParams) WithXFapiInteractionID(xFapiInteractionID *string) *GetAccountDetailParams {
	o.SetXFapiInteractionID(xFapiInteractionID)
	return o
}

// SetXFapiInteractionID adds the xFapiInteractionId to the get account detail params
func (o *GetAccountDetailParams) SetXFapiInteractionID(xFapiInteractionID *string) {
	o.XFapiInteractionID = xFapiInteractionID
}

// WithXMinv adds the xMinv to the get account detail params
func (o *GetAccountDetailParams) WithXMinv(xMinv *string) *GetAccountDetailParams {
	o.SetXMinv(xMinv)
	return o
}

// SetXMinv adds the xMinV to the get account detail params
func (o *GetAccountDetailParams) SetXMinv(xMinv *string) {
	o.XMinv = xMinv
}

// WithXv adds the xv to the get account detail params
func (o *GetAccountDetailParams) WithXv(xv string) *GetAccountDetailParams {
	o.SetXv(xv)
	return o
}

// SetXv adds the xV to the get account detail params
func (o *GetAccountDetailParams) SetXv(xv string) {
	o.Xv = xv
}

// WriteToRequest writes these params to a swagger request
func (o *GetAccountDetailParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param accountId
	if err := r.SetPathParam("accountId", o.AccountID); err != nil {
		return err
	}

	if o.XCdsClientHeaders != nil {

		// header param x-cds-client-headers
		if err := r.SetHeaderParam("x-cds-client-headers", *o.XCdsClientHeaders); err != nil {
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

	if o.XMinv != nil {

		// header param x-min-v
		if err := r.SetHeaderParam("x-min-v", *o.XMinv); err != nil {
			return err
		}
	}

	// header param x-v
	if err := r.SetHeaderParam("x-v", o.Xv); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}