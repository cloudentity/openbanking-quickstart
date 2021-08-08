// Code generated by go-swagger; DO NOT EDIT.

package consents

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

	"github.com/cloudentity/openbanking-quickstart/openbanking/obbr/consents/models"
)

// NewConsentsPostConsentsParams creates a new ConsentsPostConsentsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewConsentsPostConsentsParams() *ConsentsPostConsentsParams {
	return &ConsentsPostConsentsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewConsentsPostConsentsParamsWithTimeout creates a new ConsentsPostConsentsParams object
// with the ability to set a timeout on a request.
func NewConsentsPostConsentsParamsWithTimeout(timeout time.Duration) *ConsentsPostConsentsParams {
	return &ConsentsPostConsentsParams{
		timeout: timeout,
	}
}

// NewConsentsPostConsentsParamsWithContext creates a new ConsentsPostConsentsParams object
// with the ability to set a context for a request.
func NewConsentsPostConsentsParamsWithContext(ctx context.Context) *ConsentsPostConsentsParams {
	return &ConsentsPostConsentsParams{
		Context: ctx,
	}
}

// NewConsentsPostConsentsParamsWithHTTPClient creates a new ConsentsPostConsentsParams object
// with the ability to set a custom HTTPClient for a request.
func NewConsentsPostConsentsParamsWithHTTPClient(client *http.Client) *ConsentsPostConsentsParams {
	return &ConsentsPostConsentsParams{
		HTTPClient: client,
	}
}

/* ConsentsPostConsentsParams contains all the parameters to send to the API endpoint
   for the consents post consents operation.

   Typically these are written to a http.Request.
*/
type ConsentsPostConsentsParams struct {

	/* Authorization.

	   Cabealho HTTP padro. Permite que as credenciais sejam fornecidas dependendo do tipo de recurso solicitado
	*/
	Authorization string

	/* Body.

	   Payload para criao do consentimento.
	*/
	Body *models.OpenbankingBrasilCreateConsent

	/* XCustomerUserAgent.

	   Indica o user-agent que o usurio utiliza.
	*/
	XCustomerUserAgent *string

	/* XFapiAuthDate.

	   Data em que o usurio logou pela ltima vez com o receptor. Representada de acordo com a [RFC7231](https://tools.ietf.org/html/rfc7231).Exemplo: Sun, 10 Sep 2017 19:43:31 UTC
	*/
	XFapiAuthDate *string

	/* XFapiCustomerIPAddress.

	   O endereo IP do usurio se estiver atualmente logado com o receptor.
	*/
	XFapiCustomerIPAddress *string

	/* XFapiInteractionID.

	   Um UID [RFC4122](https://tools.ietf.org/html/rfc4122) usado como um ID de correlao. Se fornecido, o transmissor deve "reproduzir" esse valor no cabealho de resposta.
	*/
	XFapiInteractionID *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the consents post consents params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ConsentsPostConsentsParams) WithDefaults() *ConsentsPostConsentsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the consents post consents params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ConsentsPostConsentsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the consents post consents params
func (o *ConsentsPostConsentsParams) WithTimeout(timeout time.Duration) *ConsentsPostConsentsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the consents post consents params
func (o *ConsentsPostConsentsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the consents post consents params
func (o *ConsentsPostConsentsParams) WithContext(ctx context.Context) *ConsentsPostConsentsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the consents post consents params
func (o *ConsentsPostConsentsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the consents post consents params
func (o *ConsentsPostConsentsParams) WithHTTPClient(client *http.Client) *ConsentsPostConsentsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the consents post consents params
func (o *ConsentsPostConsentsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAuthorization adds the authorization to the consents post consents params
func (o *ConsentsPostConsentsParams) WithAuthorization(authorization string) *ConsentsPostConsentsParams {
	o.SetAuthorization(authorization)
	return o
}

// SetAuthorization adds the authorization to the consents post consents params
func (o *ConsentsPostConsentsParams) SetAuthorization(authorization string) {
	o.Authorization = authorization
}

// WithBody adds the body to the consents post consents params
func (o *ConsentsPostConsentsParams) WithBody(body *models.OpenbankingBrasilCreateConsent) *ConsentsPostConsentsParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the consents post consents params
func (o *ConsentsPostConsentsParams) SetBody(body *models.OpenbankingBrasilCreateConsent) {
	o.Body = body
}

// WithXCustomerUserAgent adds the xCustomerUserAgent to the consents post consents params
func (o *ConsentsPostConsentsParams) WithXCustomerUserAgent(xCustomerUserAgent *string) *ConsentsPostConsentsParams {
	o.SetXCustomerUserAgent(xCustomerUserAgent)
	return o
}

// SetXCustomerUserAgent adds the xCustomerUserAgent to the consents post consents params
func (o *ConsentsPostConsentsParams) SetXCustomerUserAgent(xCustomerUserAgent *string) {
	o.XCustomerUserAgent = xCustomerUserAgent
}

// WithXFapiAuthDate adds the xFapiAuthDate to the consents post consents params
func (o *ConsentsPostConsentsParams) WithXFapiAuthDate(xFapiAuthDate *string) *ConsentsPostConsentsParams {
	o.SetXFapiAuthDate(xFapiAuthDate)
	return o
}

// SetXFapiAuthDate adds the xFapiAuthDate to the consents post consents params
func (o *ConsentsPostConsentsParams) SetXFapiAuthDate(xFapiAuthDate *string) {
	o.XFapiAuthDate = xFapiAuthDate
}

// WithXFapiCustomerIPAddress adds the xFapiCustomerIPAddress to the consents post consents params
func (o *ConsentsPostConsentsParams) WithXFapiCustomerIPAddress(xFapiCustomerIPAddress *string) *ConsentsPostConsentsParams {
	o.SetXFapiCustomerIPAddress(xFapiCustomerIPAddress)
	return o
}

// SetXFapiCustomerIPAddress adds the xFapiCustomerIpAddress to the consents post consents params
func (o *ConsentsPostConsentsParams) SetXFapiCustomerIPAddress(xFapiCustomerIPAddress *string) {
	o.XFapiCustomerIPAddress = xFapiCustomerIPAddress
}

// WithXFapiInteractionID adds the xFapiInteractionID to the consents post consents params
func (o *ConsentsPostConsentsParams) WithXFapiInteractionID(xFapiInteractionID *string) *ConsentsPostConsentsParams {
	o.SetXFapiInteractionID(xFapiInteractionID)
	return o
}

// SetXFapiInteractionID adds the xFapiInteractionId to the consents post consents params
func (o *ConsentsPostConsentsParams) SetXFapiInteractionID(xFapiInteractionID *string) {
	o.XFapiInteractionID = xFapiInteractionID
}

// WriteToRequest writes these params to a swagger request
func (o *ConsentsPostConsentsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// header param Authorization
	if err := r.SetHeaderParam("Authorization", o.Authorization); err != nil {
		return err
	}
	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
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
