// Code generated by go-swagger; DO NOT EDIT.

package accounts

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

// NewAccountsGetAccountsParams creates a new AccountsGetAccountsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewAccountsGetAccountsParams() *AccountsGetAccountsParams {
	return &AccountsGetAccountsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewAccountsGetAccountsParamsWithTimeout creates a new AccountsGetAccountsParams object
// with the ability to set a timeout on a request.
func NewAccountsGetAccountsParamsWithTimeout(timeout time.Duration) *AccountsGetAccountsParams {
	return &AccountsGetAccountsParams{
		timeout: timeout,
	}
}

// NewAccountsGetAccountsParamsWithContext creates a new AccountsGetAccountsParams object
// with the ability to set a context for a request.
func NewAccountsGetAccountsParamsWithContext(ctx context.Context) *AccountsGetAccountsParams {
	return &AccountsGetAccountsParams{
		Context: ctx,
	}
}

// NewAccountsGetAccountsParamsWithHTTPClient creates a new AccountsGetAccountsParams object
// with the ability to set a custom HTTPClient for a request.
func NewAccountsGetAccountsParamsWithHTTPClient(client *http.Client) *AccountsGetAccountsParams {
	return &AccountsGetAccountsParams{
		HTTPClient: client,
	}
}

/* AccountsGetAccountsParams contains all the parameters to send to the API endpoint
   for the accounts get accounts operation.

   Typically these are written to a http.Request.
*/
type AccountsGetAccountsParams struct {

	/* Authorization.

	   Cabeçalho HTTP padrão. Permite que as credenciais sejam fornecidas dependendo do tipo de recurso solicitado
	*/
	Authorization string

	/* AccountType.

	   Tipos de contas. Modalidades tradicionais previstas pela Resolução 4.753, não contemplando contas vinculadas, conta de domiciliados no exterior, contas em moedas estrangeiras e conta correspondente moeda eletrônica. Vide Enum.
	*/
	AccountType *string

	/* Page.

	   Número da página que está sendo requisitada (o valor da primeira página é 1).

	   Format: int32
	   Default: 1
	*/
	Page *int32

	/* PageSize.

	   Quantidade total de registros por páginas.

	   Format: int32
	   Default: 25
	*/
	PageSize *int32

	/* XCustomerUserAgent.

	   Indica o user-agent que o usuário utiliza.
	*/
	XCustomerUserAgent *string

	/* XFapiAuthDate.

	   Data em que o usuário logou pela última vez com o receptor. Representada de acordo com a [RFC7231](https://tools.ietf.org/html/rfc7231).Exemplo: Sun, 10 Sep 2017 19:43:31 UTC
	*/
	XFapiAuthDate *string

	/* XFapiCustomerIPAddress.

	   O endereço IP do usuário se estiver atualmente logado com o receptor.
	*/
	XFapiCustomerIPAddress *string

	/* XFapiInteractionID.

	   Um UID [RFC4122](https://tools.ietf.org/html/rfc4122) usado como um ID de correlação. Se fornecido, o transmissor deve "reproduzir" esse valor no cabeçalho de resposta.
	*/
	XFapiInteractionID *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the accounts get accounts params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *AccountsGetAccountsParams) WithDefaults() *AccountsGetAccountsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the accounts get accounts params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *AccountsGetAccountsParams) SetDefaults() {
	var (
		pageDefault = int32(1)

		pageSizeDefault = int32(25)
	)

	val := AccountsGetAccountsParams{
		Page:     &pageDefault,
		PageSize: &pageSizeDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the accounts get accounts params
func (o *AccountsGetAccountsParams) WithTimeout(timeout time.Duration) *AccountsGetAccountsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the accounts get accounts params
func (o *AccountsGetAccountsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the accounts get accounts params
func (o *AccountsGetAccountsParams) WithContext(ctx context.Context) *AccountsGetAccountsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the accounts get accounts params
func (o *AccountsGetAccountsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the accounts get accounts params
func (o *AccountsGetAccountsParams) WithHTTPClient(client *http.Client) *AccountsGetAccountsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the accounts get accounts params
func (o *AccountsGetAccountsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAuthorization adds the authorization to the accounts get accounts params
func (o *AccountsGetAccountsParams) WithAuthorization(authorization string) *AccountsGetAccountsParams {
	o.SetAuthorization(authorization)
	return o
}

// SetAuthorization adds the authorization to the accounts get accounts params
func (o *AccountsGetAccountsParams) SetAuthorization(authorization string) {
	o.Authorization = authorization
}

// WithAccountType adds the accountType to the accounts get accounts params
func (o *AccountsGetAccountsParams) WithAccountType(accountType *string) *AccountsGetAccountsParams {
	o.SetAccountType(accountType)
	return o
}

// SetAccountType adds the accountType to the accounts get accounts params
func (o *AccountsGetAccountsParams) SetAccountType(accountType *string) {
	o.AccountType = accountType
}

// WithPage adds the page to the accounts get accounts params
func (o *AccountsGetAccountsParams) WithPage(page *int32) *AccountsGetAccountsParams {
	o.SetPage(page)
	return o
}

// SetPage adds the page to the accounts get accounts params
func (o *AccountsGetAccountsParams) SetPage(page *int32) {
	o.Page = page
}

// WithPageSize adds the pageSize to the accounts get accounts params
func (o *AccountsGetAccountsParams) WithPageSize(pageSize *int32) *AccountsGetAccountsParams {
	o.SetPageSize(pageSize)
	return o
}

// SetPageSize adds the pageSize to the accounts get accounts params
func (o *AccountsGetAccountsParams) SetPageSize(pageSize *int32) {
	o.PageSize = pageSize
}

// WithXCustomerUserAgent adds the xCustomerUserAgent to the accounts get accounts params
func (o *AccountsGetAccountsParams) WithXCustomerUserAgent(xCustomerUserAgent *string) *AccountsGetAccountsParams {
	o.SetXCustomerUserAgent(xCustomerUserAgent)
	return o
}

// SetXCustomerUserAgent adds the xCustomerUserAgent to the accounts get accounts params
func (o *AccountsGetAccountsParams) SetXCustomerUserAgent(xCustomerUserAgent *string) {
	o.XCustomerUserAgent = xCustomerUserAgent
}

// WithXFapiAuthDate adds the xFapiAuthDate to the accounts get accounts params
func (o *AccountsGetAccountsParams) WithXFapiAuthDate(xFapiAuthDate *string) *AccountsGetAccountsParams {
	o.SetXFapiAuthDate(xFapiAuthDate)
	return o
}

// SetXFapiAuthDate adds the xFapiAuthDate to the accounts get accounts params
func (o *AccountsGetAccountsParams) SetXFapiAuthDate(xFapiAuthDate *string) {
	o.XFapiAuthDate = xFapiAuthDate
}

// WithXFapiCustomerIPAddress adds the xFapiCustomerIPAddress to the accounts get accounts params
func (o *AccountsGetAccountsParams) WithXFapiCustomerIPAddress(xFapiCustomerIPAddress *string) *AccountsGetAccountsParams {
	o.SetXFapiCustomerIPAddress(xFapiCustomerIPAddress)
	return o
}

// SetXFapiCustomerIPAddress adds the xFapiCustomerIpAddress to the accounts get accounts params
func (o *AccountsGetAccountsParams) SetXFapiCustomerIPAddress(xFapiCustomerIPAddress *string) {
	o.XFapiCustomerIPAddress = xFapiCustomerIPAddress
}

// WithXFapiInteractionID adds the xFapiInteractionID to the accounts get accounts params
func (o *AccountsGetAccountsParams) WithXFapiInteractionID(xFapiInteractionID *string) *AccountsGetAccountsParams {
	o.SetXFapiInteractionID(xFapiInteractionID)
	return o
}

// SetXFapiInteractionID adds the xFapiInteractionId to the accounts get accounts params
func (o *AccountsGetAccountsParams) SetXFapiInteractionID(xFapiInteractionID *string) {
	o.XFapiInteractionID = xFapiInteractionID
}

// WriteToRequest writes these params to a swagger request
func (o *AccountsGetAccountsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// header param Authorization
	if err := r.SetHeaderParam("Authorization", o.Authorization); err != nil {
		return err
	}

	if o.AccountType != nil {

		// query param accountType
		var qrAccountType string

		if o.AccountType != nil {
			qrAccountType = *o.AccountType
		}
		qAccountType := qrAccountType
		if qAccountType != "" {

			if err := r.SetQueryParam("accountType", qAccountType); err != nil {
				return err
			}
		}
	}

	if o.Page != nil {

		// query param page
		var qrPage int32

		if o.Page != nil {
			qrPage = *o.Page
		}
		qPage := swag.FormatInt32(qrPage)
		if qPage != "" {

			if err := r.SetQueryParam("page", qPage); err != nil {
				return err
			}
		}
	}

	if o.PageSize != nil {

		// query param page-size
		var qrPageSize int32

		if o.PageSize != nil {
			qrPageSize = *o.PageSize
		}
		qPageSize := swag.FormatInt32(qrPageSize)
		if qPageSize != "" {

			if err := r.SetQueryParam("page-size", qPageSize); err != nil {
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
