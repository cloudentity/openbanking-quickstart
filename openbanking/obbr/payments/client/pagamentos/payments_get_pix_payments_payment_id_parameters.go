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

// NewPaymentsGetPixPaymentsPaymentIDParams creates a new PaymentsGetPixPaymentsPaymentIDParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPaymentsGetPixPaymentsPaymentIDParams() *PaymentsGetPixPaymentsPaymentIDParams {
	return &PaymentsGetPixPaymentsPaymentIDParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPaymentsGetPixPaymentsPaymentIDParamsWithTimeout creates a new PaymentsGetPixPaymentsPaymentIDParams object
// with the ability to set a timeout on a request.
func NewPaymentsGetPixPaymentsPaymentIDParamsWithTimeout(timeout time.Duration) *PaymentsGetPixPaymentsPaymentIDParams {
	return &PaymentsGetPixPaymentsPaymentIDParams{
		timeout: timeout,
	}
}

// NewPaymentsGetPixPaymentsPaymentIDParamsWithContext creates a new PaymentsGetPixPaymentsPaymentIDParams object
// with the ability to set a context for a request.
func NewPaymentsGetPixPaymentsPaymentIDParamsWithContext(ctx context.Context) *PaymentsGetPixPaymentsPaymentIDParams {
	return &PaymentsGetPixPaymentsPaymentIDParams{
		Context: ctx,
	}
}

// NewPaymentsGetPixPaymentsPaymentIDParamsWithHTTPClient creates a new PaymentsGetPixPaymentsPaymentIDParams object
// with the ability to set a custom HTTPClient for a request.
func NewPaymentsGetPixPaymentsPaymentIDParamsWithHTTPClient(client *http.Client) *PaymentsGetPixPaymentsPaymentIDParams {
	return &PaymentsGetPixPaymentsPaymentIDParams{
		HTTPClient: client,
	}
}

/* PaymentsGetPixPaymentsPaymentIDParams contains all the parameters to send to the API endpoint
   for the payments get pix payments payment Id operation.

   Typically these are written to a http.Request.
*/
type PaymentsGetPixPaymentsPaymentIDParams struct {

	/* Authorization.

	   Cabealho HTTP padro. Permite que as credenciais sejam fornecidas dependendo do tipo de recurso solicitado
	*/
	Authorization string

	/* PaymentID.

	   Identificador da operao de pagamento.
	*/
	PaymentID string

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

// WithDefaults hydrates default values in the payments get pix payments payment Id params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PaymentsGetPixPaymentsPaymentIDParams) WithDefaults() *PaymentsGetPixPaymentsPaymentIDParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the payments get pix payments payment Id params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PaymentsGetPixPaymentsPaymentIDParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the payments get pix payments payment Id params
func (o *PaymentsGetPixPaymentsPaymentIDParams) WithTimeout(timeout time.Duration) *PaymentsGetPixPaymentsPaymentIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the payments get pix payments payment Id params
func (o *PaymentsGetPixPaymentsPaymentIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the payments get pix payments payment Id params
func (o *PaymentsGetPixPaymentsPaymentIDParams) WithContext(ctx context.Context) *PaymentsGetPixPaymentsPaymentIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the payments get pix payments payment Id params
func (o *PaymentsGetPixPaymentsPaymentIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the payments get pix payments payment Id params
func (o *PaymentsGetPixPaymentsPaymentIDParams) WithHTTPClient(client *http.Client) *PaymentsGetPixPaymentsPaymentIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the payments get pix payments payment Id params
func (o *PaymentsGetPixPaymentsPaymentIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAuthorization adds the authorization to the payments get pix payments payment Id params
func (o *PaymentsGetPixPaymentsPaymentIDParams) WithAuthorization(authorization string) *PaymentsGetPixPaymentsPaymentIDParams {
	o.SetAuthorization(authorization)
	return o
}

// SetAuthorization adds the authorization to the payments get pix payments payment Id params
func (o *PaymentsGetPixPaymentsPaymentIDParams) SetAuthorization(authorization string) {
	o.Authorization = authorization
}

// WithPaymentID adds the paymentID to the payments get pix payments payment Id params
func (o *PaymentsGetPixPaymentsPaymentIDParams) WithPaymentID(paymentID string) *PaymentsGetPixPaymentsPaymentIDParams {
	o.SetPaymentID(paymentID)
	return o
}

// SetPaymentID adds the paymentId to the payments get pix payments payment Id params
func (o *PaymentsGetPixPaymentsPaymentIDParams) SetPaymentID(paymentID string) {
	o.PaymentID = paymentID
}

// WithXCustomerUserAgent adds the xCustomerUserAgent to the payments get pix payments payment Id params
func (o *PaymentsGetPixPaymentsPaymentIDParams) WithXCustomerUserAgent(xCustomerUserAgent *string) *PaymentsGetPixPaymentsPaymentIDParams {
	o.SetXCustomerUserAgent(xCustomerUserAgent)
	return o
}

// SetXCustomerUserAgent adds the xCustomerUserAgent to the payments get pix payments payment Id params
func (o *PaymentsGetPixPaymentsPaymentIDParams) SetXCustomerUserAgent(xCustomerUserAgent *string) {
	o.XCustomerUserAgent = xCustomerUserAgent
}

// WithXFapiAuthDate adds the xFapiAuthDate to the payments get pix payments payment Id params
func (o *PaymentsGetPixPaymentsPaymentIDParams) WithXFapiAuthDate(xFapiAuthDate *string) *PaymentsGetPixPaymentsPaymentIDParams {
	o.SetXFapiAuthDate(xFapiAuthDate)
	return o
}

// SetXFapiAuthDate adds the xFapiAuthDate to the payments get pix payments payment Id params
func (o *PaymentsGetPixPaymentsPaymentIDParams) SetXFapiAuthDate(xFapiAuthDate *string) {
	o.XFapiAuthDate = xFapiAuthDate
}

// WithXFapiCustomerIPAddress adds the xFapiCustomerIPAddress to the payments get pix payments payment Id params
func (o *PaymentsGetPixPaymentsPaymentIDParams) WithXFapiCustomerIPAddress(xFapiCustomerIPAddress *string) *PaymentsGetPixPaymentsPaymentIDParams {
	o.SetXFapiCustomerIPAddress(xFapiCustomerIPAddress)
	return o
}

// SetXFapiCustomerIPAddress adds the xFapiCustomerIpAddress to the payments get pix payments payment Id params
func (o *PaymentsGetPixPaymentsPaymentIDParams) SetXFapiCustomerIPAddress(xFapiCustomerIPAddress *string) {
	o.XFapiCustomerIPAddress = xFapiCustomerIPAddress
}

// WithXFapiInteractionID adds the xFapiInteractionID to the payments get pix payments payment Id params
func (o *PaymentsGetPixPaymentsPaymentIDParams) WithXFapiInteractionID(xFapiInteractionID *string) *PaymentsGetPixPaymentsPaymentIDParams {
	o.SetXFapiInteractionID(xFapiInteractionID)
	return o
}

// SetXFapiInteractionID adds the xFapiInteractionId to the payments get pix payments payment Id params
func (o *PaymentsGetPixPaymentsPaymentIDParams) SetXFapiInteractionID(xFapiInteractionID *string) {
	o.XFapiInteractionID = xFapiInteractionID
}

// WriteToRequest writes these params to a swagger request
func (o *PaymentsGetPixPaymentsPaymentIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// header param Authorization
	if err := r.SetHeaderParam("Authorization", o.Authorization); err != nil {
		return err
	}

	// path param paymentId
	if err := r.SetPathParam("paymentId", o.PaymentID); err != nil {
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
