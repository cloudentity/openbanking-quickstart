// Code generated by go-swagger; DO NOT EDIT.

package internal_transfers

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

// NewSearchForTransfersParams creates a new SearchForTransfersParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewSearchForTransfersParams() *SearchForTransfersParams {
	return &SearchForTransfersParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewSearchForTransfersParamsWithTimeout creates a new SearchForTransfersParams object
// with the ability to set a timeout on a request.
func NewSearchForTransfersParamsWithTimeout(timeout time.Duration) *SearchForTransfersParams {
	return &SearchForTransfersParams{
		timeout: timeout,
	}
}

// NewSearchForTransfersParamsWithContext creates a new SearchForTransfersParams object
// with the ability to set a context for a request.
func NewSearchForTransfersParamsWithContext(ctx context.Context) *SearchForTransfersParams {
	return &SearchForTransfersParams{
		Context: ctx,
	}
}

// NewSearchForTransfersParamsWithHTTPClient creates a new SearchForTransfersParams object
// with the ability to set a custom HTTPClient for a request.
func NewSearchForTransfersParamsWithHTTPClient(client *http.Client) *SearchForTransfersParams {
	return &SearchForTransfersParams{
		HTTPClient: client,
	}
}

/* SearchForTransfersParams contains all the parameters to send to the API endpoint
   for the search for transfers operation.

   Typically these are written to a http.Request.
*/
type SearchForTransfersParams struct {

	/* SearchEndTransferDate.

	   End time for use in retrieval of transfers by transfer date

	   Format: date
	*/
	SearchEndTransferDate *strfmt.Date

	/* SearchFromAccountIds.

	   Search for transfers by source account
	*/
	SearchFromAccountIds []string

	/* SearchStartTransferDate.

	   Start time for use in retrieval of transfers by transfer date

	   Format: date
	*/
	SearchStartTransferDate *strfmt.Date

	/* SearchStatuses.

	   Search for transfers by source account
	*/
	SearchStatuses []string

	/* SearchToAccountIds.

	   Search for transfers by source account
	*/
	SearchToAccountIds []string

	/* SearchTransferIds.

	   Search for transfers by id
	*/
	SearchTransferIds []string

	/* UpdatedSince.

	   Return items that have been created or updated since the update id
	*/
	UpdatedSince *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the search for transfers params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SearchForTransfersParams) WithDefaults() *SearchForTransfersParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the search for transfers params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SearchForTransfersParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the search for transfers params
func (o *SearchForTransfersParams) WithTimeout(timeout time.Duration) *SearchForTransfersParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the search for transfers params
func (o *SearchForTransfersParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the search for transfers params
func (o *SearchForTransfersParams) WithContext(ctx context.Context) *SearchForTransfersParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the search for transfers params
func (o *SearchForTransfersParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the search for transfers params
func (o *SearchForTransfersParams) WithHTTPClient(client *http.Client) *SearchForTransfersParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the search for transfers params
func (o *SearchForTransfersParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithSearchEndTransferDate adds the searchEndTransferDate to the search for transfers params
func (o *SearchForTransfersParams) WithSearchEndTransferDate(searchEndTransferDate *strfmt.Date) *SearchForTransfersParams {
	o.SetSearchEndTransferDate(searchEndTransferDate)
	return o
}

// SetSearchEndTransferDate adds the searchEndTransferDate to the search for transfers params
func (o *SearchForTransfersParams) SetSearchEndTransferDate(searchEndTransferDate *strfmt.Date) {
	o.SearchEndTransferDate = searchEndTransferDate
}

// WithSearchFromAccountIds adds the searchFromAccountIds to the search for transfers params
func (o *SearchForTransfersParams) WithSearchFromAccountIds(searchFromAccountIds []string) *SearchForTransfersParams {
	o.SetSearchFromAccountIds(searchFromAccountIds)
	return o
}

// SetSearchFromAccountIds adds the searchFromAccountIds to the search for transfers params
func (o *SearchForTransfersParams) SetSearchFromAccountIds(searchFromAccountIds []string) {
	o.SearchFromAccountIds = searchFromAccountIds
}

// WithSearchStartTransferDate adds the searchStartTransferDate to the search for transfers params
func (o *SearchForTransfersParams) WithSearchStartTransferDate(searchStartTransferDate *strfmt.Date) *SearchForTransfersParams {
	o.SetSearchStartTransferDate(searchStartTransferDate)
	return o
}

// SetSearchStartTransferDate adds the searchStartTransferDate to the search for transfers params
func (o *SearchForTransfersParams) SetSearchStartTransferDate(searchStartTransferDate *strfmt.Date) {
	o.SearchStartTransferDate = searchStartTransferDate
}

// WithSearchStatuses adds the searchStatuses to the search for transfers params
func (o *SearchForTransfersParams) WithSearchStatuses(searchStatuses []string) *SearchForTransfersParams {
	o.SetSearchStatuses(searchStatuses)
	return o
}

// SetSearchStatuses adds the searchStatuses to the search for transfers params
func (o *SearchForTransfersParams) SetSearchStatuses(searchStatuses []string) {
	o.SearchStatuses = searchStatuses
}

// WithSearchToAccountIds adds the searchToAccountIds to the search for transfers params
func (o *SearchForTransfersParams) WithSearchToAccountIds(searchToAccountIds []string) *SearchForTransfersParams {
	o.SetSearchToAccountIds(searchToAccountIds)
	return o
}

// SetSearchToAccountIds adds the searchToAccountIds to the search for transfers params
func (o *SearchForTransfersParams) SetSearchToAccountIds(searchToAccountIds []string) {
	o.SearchToAccountIds = searchToAccountIds
}

// WithSearchTransferIds adds the searchTransferIds to the search for transfers params
func (o *SearchForTransfersParams) WithSearchTransferIds(searchTransferIds []string) *SearchForTransfersParams {
	o.SetSearchTransferIds(searchTransferIds)
	return o
}

// SetSearchTransferIds adds the searchTransferIds to the search for transfers params
func (o *SearchForTransfersParams) SetSearchTransferIds(searchTransferIds []string) {
	o.SearchTransferIds = searchTransferIds
}

// WithUpdatedSince adds the updatedSince to the search for transfers params
func (o *SearchForTransfersParams) WithUpdatedSince(updatedSince *string) *SearchForTransfersParams {
	o.SetUpdatedSince(updatedSince)
	return o
}

// SetUpdatedSince adds the updatedSince to the search for transfers params
func (o *SearchForTransfersParams) SetUpdatedSince(updatedSince *string) {
	o.UpdatedSince = updatedSince
}

// WriteToRequest writes these params to a swagger request
func (o *SearchForTransfersParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.SearchEndTransferDate != nil {

		// query param searchEndTransferDate
		var qrSearchEndTransferDate strfmt.Date

		if o.SearchEndTransferDate != nil {
			qrSearchEndTransferDate = *o.SearchEndTransferDate
		}
		qSearchEndTransferDate := qrSearchEndTransferDate.String()
		if qSearchEndTransferDate != "" {

			if err := r.SetQueryParam("searchEndTransferDate", qSearchEndTransferDate); err != nil {
				return err
			}
		}
	}

	if o.SearchFromAccountIds != nil {

		// binding items for searchFromAccountIds
		joinedSearchFromAccountIds := o.bindParamSearchFromAccountIds(reg)

		// query array param searchFromAccountIds
		if err := r.SetQueryParam("searchFromAccountIds", joinedSearchFromAccountIds...); err != nil {
			return err
		}
	}

	if o.SearchStartTransferDate != nil {

		// query param searchStartTransferDate
		var qrSearchStartTransferDate strfmt.Date

		if o.SearchStartTransferDate != nil {
			qrSearchStartTransferDate = *o.SearchStartTransferDate
		}
		qSearchStartTransferDate := qrSearchStartTransferDate.String()
		if qSearchStartTransferDate != "" {

			if err := r.SetQueryParam("searchStartTransferDate", qSearchStartTransferDate); err != nil {
				return err
			}
		}
	}

	if o.SearchStatuses != nil {

		// binding items for searchStatuses
		joinedSearchStatuses := o.bindParamSearchStatuses(reg)

		// query array param searchStatuses
		if err := r.SetQueryParam("searchStatuses", joinedSearchStatuses...); err != nil {
			return err
		}
	}

	if o.SearchToAccountIds != nil {

		// binding items for searchToAccountIds
		joinedSearchToAccountIds := o.bindParamSearchToAccountIds(reg)

		// query array param searchToAccountIds
		if err := r.SetQueryParam("searchToAccountIds", joinedSearchToAccountIds...); err != nil {
			return err
		}
	}

	if o.SearchTransferIds != nil {

		// binding items for searchTransferIds
		joinedSearchTransferIds := o.bindParamSearchTransferIds(reg)

		// query array param searchTransferIds
		if err := r.SetQueryParam("searchTransferIds", joinedSearchTransferIds...); err != nil {
			return err
		}
	}

	if o.UpdatedSince != nil {

		// query param updatedSince
		var qrUpdatedSince string

		if o.UpdatedSince != nil {
			qrUpdatedSince = *o.UpdatedSince
		}
		qUpdatedSince := qrUpdatedSince
		if qUpdatedSince != "" {

			if err := r.SetQueryParam("updatedSince", qUpdatedSince); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindParamSearchForTransfers binds the parameter searchFromAccountIds
func (o *SearchForTransfersParams) bindParamSearchFromAccountIds(formats strfmt.Registry) []string {
	searchFromAccountIdsIR := o.SearchFromAccountIds

	var searchFromAccountIdsIC []string
	for _, searchFromAccountIdsIIR := range searchFromAccountIdsIR { // explode []string

		searchFromAccountIdsIIV := searchFromAccountIdsIIR // string as string
		searchFromAccountIdsIC = append(searchFromAccountIdsIC, searchFromAccountIdsIIV)
	}

	// items.CollectionFormat: ""
	searchFromAccountIdsIS := swag.JoinByFormat(searchFromAccountIdsIC, "")

	return searchFromAccountIdsIS
}

// bindParamSearchForTransfers binds the parameter searchStatuses
func (o *SearchForTransfersParams) bindParamSearchStatuses(formats strfmt.Registry) []string {
	searchStatusesIR := o.SearchStatuses

	var searchStatusesIC []string
	for _, searchStatusesIIR := range searchStatusesIR { // explode []string

		searchStatusesIIV := searchStatusesIIR // string as string
		searchStatusesIC = append(searchStatusesIC, searchStatusesIIV)
	}

	// items.CollectionFormat: ""
	searchStatusesIS := swag.JoinByFormat(searchStatusesIC, "")

	return searchStatusesIS
}

// bindParamSearchForTransfers binds the parameter searchToAccountIds
func (o *SearchForTransfersParams) bindParamSearchToAccountIds(formats strfmt.Registry) []string {
	searchToAccountIdsIR := o.SearchToAccountIds

	var searchToAccountIdsIC []string
	for _, searchToAccountIdsIIR := range searchToAccountIdsIR { // explode []string

		searchToAccountIdsIIV := searchToAccountIdsIIR // string as string
		searchToAccountIdsIC = append(searchToAccountIdsIC, searchToAccountIdsIIV)
	}

	// items.CollectionFormat: ""
	searchToAccountIdsIS := swag.JoinByFormat(searchToAccountIdsIC, "")

	return searchToAccountIdsIS
}

// bindParamSearchForTransfers binds the parameter searchTransferIds
func (o *SearchForTransfersParams) bindParamSearchTransferIds(formats strfmt.Registry) []string {
	searchTransferIdsIR := o.SearchTransferIds

	var searchTransferIdsIC []string
	for _, searchTransferIdsIIR := range searchTransferIdsIR { // explode []string

		searchTransferIdsIIV := searchTransferIdsIIR // string as string
		searchTransferIdsIC = append(searchTransferIdsIC, searchTransferIdsIIV)
	}

	// items.CollectionFormat: ""
	searchTransferIdsIS := swag.JoinByFormat(searchTransferIdsIC, "")

	return searchTransferIdsIS
}