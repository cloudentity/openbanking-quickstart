// Code generated by go-swagger; DO NOT EDIT.

package pagamentos

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

// NewPaymentsPostPixPaymentsParams creates a new PaymentsPostPixPaymentsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPaymentsPostPixPaymentsParams() *PaymentsPostPixPaymentsParams {
	return &PaymentsPostPixPaymentsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPaymentsPostPixPaymentsParamsWithTimeout creates a new PaymentsPostPixPaymentsParams object
// with the ability to set a timeout on a request.
func NewPaymentsPostPixPaymentsParamsWithTimeout(timeout time.Duration) *PaymentsPostPixPaymentsParams {
	return &PaymentsPostPixPaymentsParams{
		timeout: timeout,
	}
}

// NewPaymentsPostPixPaymentsParamsWithContext creates a new PaymentsPostPixPaymentsParams object
// with the ability to set a context for a request.
func NewPaymentsPostPixPaymentsParamsWithContext(ctx context.Context) *PaymentsPostPixPaymentsParams {
	return &PaymentsPostPixPaymentsParams{
		Context: ctx,
	}
}

// NewPaymentsPostPixPaymentsParamsWithHTTPClient creates a new PaymentsPostPixPaymentsParams object
// with the ability to set a custom HTTPClient for a request.
func NewPaymentsPostPixPaymentsParamsWithHTTPClient(client *http.Client) *PaymentsPostPixPaymentsParams {
	return &PaymentsPostPixPaymentsParams{
		HTTPClient: client,
	}
}

/* PaymentsPostPixPaymentsParams contains all the parameters to send to the API endpoint
   for the payments post pix payments operation.

   Typically these are written to a http.Request.
*/
type PaymentsPostPixPaymentsParams struct {

	/* Authorization.

	   Cabealho HTTP padro. Permite que as credenciais sejam fornecidas dependendo do tipo de recurso solicitado
	*/
	Authorization string

	/* Body.

	   Payload para criao da iniciao do pagamento Pix.
	*/
	Body interface{}

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

	/* XIdempotencyKey.

	   Cabealho HTTP personalizado. Identificador de solicitao exclusivo para suportar a idempotncia.
	*/
	XIdempotencyKey string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the payments post pix payments params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PaymentsPostPixPaymentsParams) WithDefaults() *PaymentsPostPixPaymentsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the payments post pix payments params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PaymentsPostPixPaymentsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the payments post pix payments params
func (o *PaymentsPostPixPaymentsParams) WithTimeout(timeout time.Duration) *PaymentsPostPixPaymentsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the payments post pix payments params
func (o *PaymentsPostPixPaymentsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the payments post pix payments params
func (o *PaymentsPostPixPaymentsParams) WithContext(ctx context.Context) *PaymentsPostPixPaymentsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the payments post pix payments params
func (o *PaymentsPostPixPaymentsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the payments post pix payments params
func (o *PaymentsPostPixPaymentsParams) WithHTTPClient(client *http.Client) *PaymentsPostPixPaymentsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the payments post pix payments params
func (o *PaymentsPostPixPaymentsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAuthorization adds the authorization to the payments post pix payments params
func (o *PaymentsPostPixPaymentsParams) WithAuthorization(authorization string) *PaymentsPostPixPaymentsParams {
	o.SetAuthorization(authorization)
	return o
}

// SetAuthorization adds the authorization to the payments post pix payments params
func (o *PaymentsPostPixPaymentsParams) SetAuthorization(authorization string) {
	o.Authorization = authorization
}

// WithBody adds the body to the payments post pix payments params
func (o *PaymentsPostPixPaymentsParams) WithBody(body interface{}) *PaymentsPostPixPaymentsParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the payments post pix payments params
func (o *PaymentsPostPixPaymentsParams) SetBody(body interface{}) {
	o.Body = body
}

// WithXCustomerUserAgent adds the xCustomerUserAgent to the payments post pix payments params
func (o *PaymentsPostPixPaymentsParams) WithXCustomerUserAgent(xCustomerUserAgent *string) *PaymentsPostPixPaymentsParams {
	o.SetXCustomerUserAgent(xCustomerUserAgent)
	return o
}

// SetXCustomerUserAgent adds the xCustomerUserAgent to the payments post pix payments params
func (o *PaymentsPostPixPaymentsParams) SetXCustomerUserAgent(xCustomerUserAgent *string) {
	o.XCustomerUserAgent = xCustomerUserAgent
}

// WithXFapiAuthDate adds the xFapiAuthDate to the payments post pix payments params
func (o *PaymentsPostPixPaymentsParams) WithXFapiAuthDate(xFapiAuthDate *string) *PaymentsPostPixPaymentsParams {
	o.SetXFapiAuthDate(xFapiAuthDate)
	return o
}

// SetXFapiAuthDate adds the xFapiAuthDate to the payments post pix payments params
func (o *PaymentsPostPixPaymentsParams) SetXFapiAuthDate(xFapiAuthDate *string) {
	o.XFapiAuthDate = xFapiAuthDate
}

// WithXFapiCustomerIPAddress adds the xFapiCustomerIPAddress to the payments post pix payments params
func (o *PaymentsPostPixPaymentsParams) WithXFapiCustomerIPAddress(xFapiCustomerIPAddress *string) *PaymentsPostPixPaymentsParams {
	o.SetXFapiCustomerIPAddress(xFapiCustomerIPAddress)
	return o
}

// SetXFapiCustomerIPAddress adds the xFapiCustomerIpAddress to the payments post pix payments params
func (o *PaymentsPostPixPaymentsParams) SetXFapiCustomerIPAddress(xFapiCustomerIPAddress *string) {
	o.XFapiCustomerIPAddress = xFapiCustomerIPAddress
}

// WithXFapiInteractionID adds the xFapiInteractionID to the payments post pix payments params
func (o *PaymentsPostPixPaymentsParams) WithXFapiInteractionID(xFapiInteractionID *string) *PaymentsPostPixPaymentsParams {
	o.SetXFapiInteractionID(xFapiInteractionID)
	return o
}

// SetXFapiInteractionID adds the xFapiInteractionId to the payments post pix payments params
func (o *PaymentsPostPixPaymentsParams) SetXFapiInteractionID(xFapiInteractionID *string) {
	o.XFapiInteractionID = xFapiInteractionID
}

// WithXIdempotencyKey adds the xIdempotencyKey to the payments post pix payments params
func (o *PaymentsPostPixPaymentsParams) WithXIdempotencyKey(xIdempotencyKey string) *PaymentsPostPixPaymentsParams {
	o.SetXIdempotencyKey(xIdempotencyKey)
	return o
}

// SetXIdempotencyKey adds the xIdempotencyKey to the payments post pix payments params
func (o *PaymentsPostPixPaymentsParams) SetXIdempotencyKey(xIdempotencyKey string) {
	o.XIdempotencyKey = xIdempotencyKey
}

// WriteToRequest writes these params to a swagger request
func (o *PaymentsPostPixPaymentsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// header param x-idempotency-key
	if err := r.SetHeaderParam("x-idempotency-key", o.XIdempotencyKey); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
