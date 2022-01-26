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

// NewGetProductDetailParams creates a new GetProductDetailParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetProductDetailParams() *GetProductDetailParams {
	return &GetProductDetailParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetProductDetailParamsWithTimeout creates a new GetProductDetailParams object
// with the ability to set a timeout on a request.
func NewGetProductDetailParamsWithTimeout(timeout time.Duration) *GetProductDetailParams {
	return &GetProductDetailParams{
		timeout: timeout,
	}
}

// NewGetProductDetailParamsWithContext creates a new GetProductDetailParams object
// with the ability to set a context for a request.
func NewGetProductDetailParamsWithContext(ctx context.Context) *GetProductDetailParams {
	return &GetProductDetailParams{
		Context: ctx,
	}
}

// NewGetProductDetailParamsWithHTTPClient creates a new GetProductDetailParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetProductDetailParamsWithHTTPClient(client *http.Client) *GetProductDetailParams {
	return &GetProductDetailParams{
		HTTPClient: client,
	}
}

/* GetProductDetailParams contains all the parameters to send to the API endpoint
   for the get product detail operation.

   Typically these are written to a http.Request.
*/
type GetProductDetailParams struct {

	/* ProductID.

	   ID of the specific product requested
	*/
	ProductID string

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

// WithDefaults hydrates default values in the get product detail params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetProductDetailParams) WithDefaults() *GetProductDetailParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get product detail params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetProductDetailParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get product detail params
func (o *GetProductDetailParams) WithTimeout(timeout time.Duration) *GetProductDetailParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get product detail params
func (o *GetProductDetailParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get product detail params
func (o *GetProductDetailParams) WithContext(ctx context.Context) *GetProductDetailParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get product detail params
func (o *GetProductDetailParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get product detail params
func (o *GetProductDetailParams) WithHTTPClient(client *http.Client) *GetProductDetailParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get product detail params
func (o *GetProductDetailParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithProductID adds the productID to the get product detail params
func (o *GetProductDetailParams) WithProductID(productID string) *GetProductDetailParams {
	o.SetProductID(productID)
	return o
}

// SetProductID adds the productId to the get product detail params
func (o *GetProductDetailParams) SetProductID(productID string) {
	o.ProductID = productID
}

// WithXMinv adds the xMinv to the get product detail params
func (o *GetProductDetailParams) WithXMinv(xMinv *string) *GetProductDetailParams {
	o.SetXMinv(xMinv)
	return o
}

// SetXMinv adds the xMinV to the get product detail params
func (o *GetProductDetailParams) SetXMinv(xMinv *string) {
	o.XMinv = xMinv
}

// WithXv adds the xv to the get product detail params
func (o *GetProductDetailParams) WithXv(xv string) *GetProductDetailParams {
	o.SetXv(xv)
	return o
}

// SetXv adds the xV to the get product detail params
func (o *GetProductDetailParams) SetXv(xv string) {
	o.Xv = xv
}

// WriteToRequest writes these params to a swagger request
func (o *GetProductDetailParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param productId
	if err := r.SetPathParam("productId", o.ProductID); err != nil {
		return err
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
