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
	"github.com/go-openapi/swag"
)

// NewListProductsParams creates a new ListProductsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewListProductsParams() *ListProductsParams {
	return &ListProductsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewListProductsParamsWithTimeout creates a new ListProductsParams object
// with the ability to set a timeout on a request.
func NewListProductsParamsWithTimeout(timeout time.Duration) *ListProductsParams {
	return &ListProductsParams{
		timeout: timeout,
	}
}

// NewListProductsParamsWithContext creates a new ListProductsParams object
// with the ability to set a context for a request.
func NewListProductsParamsWithContext(ctx context.Context) *ListProductsParams {
	return &ListProductsParams{
		Context: ctx,
	}
}

// NewListProductsParamsWithHTTPClient creates a new ListProductsParams object
// with the ability to set a custom HTTPClient for a request.
func NewListProductsParamsWithHTTPClient(client *http.Client) *ListProductsParams {
	return &ListProductsParams{
		HTTPClient: client,
	}
}

/* ListProductsParams contains all the parameters to send to the API endpoint
   for the list products operation.

   Typically these are written to a http.Request.
*/
type ListProductsParams struct {

	/* Brand.

	   Filter results based on a specific brand
	*/
	Brand *string

	/* Effective.

	   Allows for the filtering of products based on whether the current time is within the period of time defined as effective by the effectiveFrom and effectiveTo fields. Valid values are ‘CURRENT’, ‘FUTURE’ and ‘ALL’. If absent defaults to 'CURRENT'

	   Default: "CURRENT"
	*/
	Effective *string

	/* Page.

	   Page of results to request (standard pagination)

	   Format: int32
	   Default: 1
	*/
	Page *int32

	/* PageSize.

	   Page size to request. Default is 25 (standard pagination)

	   Format: int32
	   Default: 25
	*/
	PageSize *int32

	/* ProductCategory.

	   Used to filter results on the productCategory field applicable to accounts. Any one of the valid values for this field can be supplied. If absent then all accounts returned.
	*/
	ProductCategory *string

	/* UpdatedSince.

	   Only include products that have been updated after the specified date and time. If absent defaults to include all products
	*/
	UpdatedSince *string

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

// WithDefaults hydrates default values in the list products params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ListProductsParams) WithDefaults() *ListProductsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the list products params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ListProductsParams) SetDefaults() {
	var (
		effectiveDefault = string("CURRENT")

		pageDefault = int32(1)

		pageSizeDefault = int32(25)
	)

	val := ListProductsParams{
		Effective: &effectiveDefault,
		Page:      &pageDefault,
		PageSize:  &pageSizeDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the list products params
func (o *ListProductsParams) WithTimeout(timeout time.Duration) *ListProductsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the list products params
func (o *ListProductsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the list products params
func (o *ListProductsParams) WithContext(ctx context.Context) *ListProductsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the list products params
func (o *ListProductsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the list products params
func (o *ListProductsParams) WithHTTPClient(client *http.Client) *ListProductsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the list products params
func (o *ListProductsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBrand adds the brand to the list products params
func (o *ListProductsParams) WithBrand(brand *string) *ListProductsParams {
	o.SetBrand(brand)
	return o
}

// SetBrand adds the brand to the list products params
func (o *ListProductsParams) SetBrand(brand *string) {
	o.Brand = brand
}

// WithEffective adds the effective to the list products params
func (o *ListProductsParams) WithEffective(effective *string) *ListProductsParams {
	o.SetEffective(effective)
	return o
}

// SetEffective adds the effective to the list products params
func (o *ListProductsParams) SetEffective(effective *string) {
	o.Effective = effective
}

// WithPage adds the page to the list products params
func (o *ListProductsParams) WithPage(page *int32) *ListProductsParams {
	o.SetPage(page)
	return o
}

// SetPage adds the page to the list products params
func (o *ListProductsParams) SetPage(page *int32) {
	o.Page = page
}

// WithPageSize adds the pageSize to the list products params
func (o *ListProductsParams) WithPageSize(pageSize *int32) *ListProductsParams {
	o.SetPageSize(pageSize)
	return o
}

// SetPageSize adds the pageSize to the list products params
func (o *ListProductsParams) SetPageSize(pageSize *int32) {
	o.PageSize = pageSize
}

// WithProductCategory adds the productCategory to the list products params
func (o *ListProductsParams) WithProductCategory(productCategory *string) *ListProductsParams {
	o.SetProductCategory(productCategory)
	return o
}

// SetProductCategory adds the productCategory to the list products params
func (o *ListProductsParams) SetProductCategory(productCategory *string) {
	o.ProductCategory = productCategory
}

// WithUpdatedSince adds the updatedSince to the list products params
func (o *ListProductsParams) WithUpdatedSince(updatedSince *string) *ListProductsParams {
	o.SetUpdatedSince(updatedSince)
	return o
}

// SetUpdatedSince adds the updatedSince to the list products params
func (o *ListProductsParams) SetUpdatedSince(updatedSince *string) {
	o.UpdatedSince = updatedSince
}

// WithXMinv adds the xMinv to the list products params
func (o *ListProductsParams) WithXMinv(xMinv *string) *ListProductsParams {
	o.SetXMinv(xMinv)
	return o
}

// SetXMinv adds the xMinV to the list products params
func (o *ListProductsParams) SetXMinv(xMinv *string) {
	o.XMinv = xMinv
}

// WithXv adds the xv to the list products params
func (o *ListProductsParams) WithXv(xv string) *ListProductsParams {
	o.SetXv(xv)
	return o
}

// SetXv adds the xV to the list products params
func (o *ListProductsParams) SetXv(xv string) {
	o.Xv = xv
}

// WriteToRequest writes these params to a swagger request
func (o *ListProductsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Brand != nil {

		// query param brand
		var qrBrand string

		if o.Brand != nil {
			qrBrand = *o.Brand
		}
		qBrand := qrBrand
		if qBrand != "" {

			if err := r.SetQueryParam("brand", qBrand); err != nil {
				return err
			}
		}
	}

	if o.Effective != nil {

		// query param effective
		var qrEffective string

		if o.Effective != nil {
			qrEffective = *o.Effective
		}
		qEffective := qrEffective
		if qEffective != "" {

			if err := r.SetQueryParam("effective", qEffective); err != nil {
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

	if o.ProductCategory != nil {

		// query param product-category
		var qrProductCategory string

		if o.ProductCategory != nil {
			qrProductCategory = *o.ProductCategory
		}
		qProductCategory := qrProductCategory
		if qProductCategory != "" {

			if err := r.SetQueryParam("product-category", qProductCategory); err != nil {
				return err
			}
		}
	}

	if o.UpdatedSince != nil {

		// query param updated-since
		var qrUpdatedSince string

		if o.UpdatedSince != nil {
			qrUpdatedSince = *o.UpdatedSince
		}
		qUpdatedSince := qrUpdatedSince
		if qUpdatedSince != "" {

			if err := r.SetQueryParam("updated-since", qUpdatedSince); err != nil {
				return err
			}
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
