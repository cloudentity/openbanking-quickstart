// Code generated by go-swagger; DO NOT EDIT.

package statements

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

// NewGetAccountsAccountIDStatementsParams creates a new GetAccountsAccountIDStatementsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetAccountsAccountIDStatementsParams() *GetAccountsAccountIDStatementsParams {
	return &GetAccountsAccountIDStatementsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetAccountsAccountIDStatementsParamsWithTimeout creates a new GetAccountsAccountIDStatementsParams object
// with the ability to set a timeout on a request.
func NewGetAccountsAccountIDStatementsParamsWithTimeout(timeout time.Duration) *GetAccountsAccountIDStatementsParams {
	return &GetAccountsAccountIDStatementsParams{
		timeout: timeout,
	}
}

// NewGetAccountsAccountIDStatementsParamsWithContext creates a new GetAccountsAccountIDStatementsParams object
// with the ability to set a context for a request.
func NewGetAccountsAccountIDStatementsParamsWithContext(ctx context.Context) *GetAccountsAccountIDStatementsParams {
	return &GetAccountsAccountIDStatementsParams{
		Context: ctx,
	}
}

// NewGetAccountsAccountIDStatementsParamsWithHTTPClient creates a new GetAccountsAccountIDStatementsParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetAccountsAccountIDStatementsParamsWithHTTPClient(client *http.Client) *GetAccountsAccountIDStatementsParams {
	return &GetAccountsAccountIDStatementsParams{
		HTTPClient: client,
	}
}

/* GetAccountsAccountIDStatementsParams contains all the parameters to send to the API endpoint
   for the get accounts account Id statements operation.

   Typically these are written to a http.Request.
*/
type GetAccountsAccountIDStatementsParams struct {

	/* AccountID.

	   AccountId
	*/
	AccountID string

	/* Authorization.

	   An Authorisation Token as per https://tools.ietf.org/html/rfc6750
	*/
	Authorization string

	/* FromStatementDateTime.

	     The UTC ISO 8601 Date Time to filter statements FROM
	NB Time component is optional - set to 00:00:00 for just Date.
	If the Date Time contains a timezone, the ASPSP must ignore the timezone component.

	     Format: date-time
	*/
	FromStatementDateTime *strfmt.DateTime

	/* ToStatementDateTime.

	     The UTC ISO 8601 Date Time to filter statements TO
	NB Time component is optional - set to 00:00:00 for just Date.
	If the Date Time contains a timezone, the ASPSP must ignore the timezone component.

	     Format: date-time
	*/
	ToStatementDateTime *strfmt.DateTime

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

// WithDefaults hydrates default values in the get accounts account Id statements params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetAccountsAccountIDStatementsParams) WithDefaults() *GetAccountsAccountIDStatementsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get accounts account Id statements params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetAccountsAccountIDStatementsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get accounts account Id statements params
func (o *GetAccountsAccountIDStatementsParams) WithTimeout(timeout time.Duration) *GetAccountsAccountIDStatementsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get accounts account Id statements params
func (o *GetAccountsAccountIDStatementsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get accounts account Id statements params
func (o *GetAccountsAccountIDStatementsParams) WithContext(ctx context.Context) *GetAccountsAccountIDStatementsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get accounts account Id statements params
func (o *GetAccountsAccountIDStatementsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get accounts account Id statements params
func (o *GetAccountsAccountIDStatementsParams) WithHTTPClient(client *http.Client) *GetAccountsAccountIDStatementsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get accounts account Id statements params
func (o *GetAccountsAccountIDStatementsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAccountID adds the accountID to the get accounts account Id statements params
func (o *GetAccountsAccountIDStatementsParams) WithAccountID(accountID string) *GetAccountsAccountIDStatementsParams {
	o.SetAccountID(accountID)
	return o
}

// SetAccountID adds the accountId to the get accounts account Id statements params
func (o *GetAccountsAccountIDStatementsParams) SetAccountID(accountID string) {
	o.AccountID = accountID
}

// WithAuthorization adds the authorization to the get accounts account Id statements params
func (o *GetAccountsAccountIDStatementsParams) WithAuthorization(authorization string) *GetAccountsAccountIDStatementsParams {
	o.SetAuthorization(authorization)
	return o
}

// SetAuthorization adds the authorization to the get accounts account Id statements params
func (o *GetAccountsAccountIDStatementsParams) SetAuthorization(authorization string) {
	o.Authorization = authorization
}

// WithFromStatementDateTime adds the fromStatementDateTime to the get accounts account Id statements params
func (o *GetAccountsAccountIDStatementsParams) WithFromStatementDateTime(fromStatementDateTime *strfmt.DateTime) *GetAccountsAccountIDStatementsParams {
	o.SetFromStatementDateTime(fromStatementDateTime)
	return o
}

// SetFromStatementDateTime adds the fromStatementDateTime to the get accounts account Id statements params
func (o *GetAccountsAccountIDStatementsParams) SetFromStatementDateTime(fromStatementDateTime *strfmt.DateTime) {
	o.FromStatementDateTime = fromStatementDateTime
}

// WithToStatementDateTime adds the toStatementDateTime to the get accounts account Id statements params
func (o *GetAccountsAccountIDStatementsParams) WithToStatementDateTime(toStatementDateTime *strfmt.DateTime) *GetAccountsAccountIDStatementsParams {
	o.SetToStatementDateTime(toStatementDateTime)
	return o
}

// SetToStatementDateTime adds the toStatementDateTime to the get accounts account Id statements params
func (o *GetAccountsAccountIDStatementsParams) SetToStatementDateTime(toStatementDateTime *strfmt.DateTime) {
	o.ToStatementDateTime = toStatementDateTime
}

// WithXCustomerUserAgent adds the xCustomerUserAgent to the get accounts account Id statements params
func (o *GetAccountsAccountIDStatementsParams) WithXCustomerUserAgent(xCustomerUserAgent *string) *GetAccountsAccountIDStatementsParams {
	o.SetXCustomerUserAgent(xCustomerUserAgent)
	return o
}

// SetXCustomerUserAgent adds the xCustomerUserAgent to the get accounts account Id statements params
func (o *GetAccountsAccountIDStatementsParams) SetXCustomerUserAgent(xCustomerUserAgent *string) {
	o.XCustomerUserAgent = xCustomerUserAgent
}

// WithXFapiAuthDate adds the xFapiAuthDate to the get accounts account Id statements params
func (o *GetAccountsAccountIDStatementsParams) WithXFapiAuthDate(xFapiAuthDate *string) *GetAccountsAccountIDStatementsParams {
	o.SetXFapiAuthDate(xFapiAuthDate)
	return o
}

// SetXFapiAuthDate adds the xFapiAuthDate to the get accounts account Id statements params
func (o *GetAccountsAccountIDStatementsParams) SetXFapiAuthDate(xFapiAuthDate *string) {
	o.XFapiAuthDate = xFapiAuthDate
}

// WithXFapiCustomerIPAddress adds the xFapiCustomerIPAddress to the get accounts account Id statements params
func (o *GetAccountsAccountIDStatementsParams) WithXFapiCustomerIPAddress(xFapiCustomerIPAddress *string) *GetAccountsAccountIDStatementsParams {
	o.SetXFapiCustomerIPAddress(xFapiCustomerIPAddress)
	return o
}

// SetXFapiCustomerIPAddress adds the xFapiCustomerIpAddress to the get accounts account Id statements params
func (o *GetAccountsAccountIDStatementsParams) SetXFapiCustomerIPAddress(xFapiCustomerIPAddress *string) {
	o.XFapiCustomerIPAddress = xFapiCustomerIPAddress
}

// WithXFapiInteractionID adds the xFapiInteractionID to the get accounts account Id statements params
func (o *GetAccountsAccountIDStatementsParams) WithXFapiInteractionID(xFapiInteractionID *string) *GetAccountsAccountIDStatementsParams {
	o.SetXFapiInteractionID(xFapiInteractionID)
	return o
}

// SetXFapiInteractionID adds the xFapiInteractionId to the get accounts account Id statements params
func (o *GetAccountsAccountIDStatementsParams) SetXFapiInteractionID(xFapiInteractionID *string) {
	o.XFapiInteractionID = xFapiInteractionID
}

// WriteToRequest writes these params to a swagger request
func (o *GetAccountsAccountIDStatementsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param AccountId
	if err := r.SetPathParam("AccountId", o.AccountID); err != nil {
		return err
	}

	// header param Authorization
	if err := r.SetHeaderParam("Authorization", o.Authorization); err != nil {
		return err
	}

	if o.FromStatementDateTime != nil {

		// query param fromStatementDateTime
		var qrFromStatementDateTime strfmt.DateTime

		if o.FromStatementDateTime != nil {
			qrFromStatementDateTime = *o.FromStatementDateTime
		}
		qFromStatementDateTime := qrFromStatementDateTime.String()
		if qFromStatementDateTime != "" {

			if err := r.SetQueryParam("fromStatementDateTime", qFromStatementDateTime); err != nil {
				return err
			}
		}
	}

	if o.ToStatementDateTime != nil {

		// query param toStatementDateTime
		var qrToStatementDateTime strfmt.DateTime

		if o.ToStatementDateTime != nil {
			qrToStatementDateTime = *o.ToStatementDateTime
		}
		qToStatementDateTime := qrToStatementDateTime.String()
		if qToStatementDateTime != "" {

			if err := r.SetQueryParam("toStatementDateTime", qToStatementDateTime); err != nil {
				return err
			}
		}
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