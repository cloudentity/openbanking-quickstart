// Code generated by go-swagger; DO NOT EDIT.

package reward_program_information

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

// NewSearchRewardProgramsParams creates a new SearchRewardProgramsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewSearchRewardProgramsParams() *SearchRewardProgramsParams {
	return &SearchRewardProgramsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewSearchRewardProgramsParamsWithTimeout creates a new SearchRewardProgramsParams object
// with the ability to set a timeout on a request.
func NewSearchRewardProgramsParamsWithTimeout(timeout time.Duration) *SearchRewardProgramsParams {
	return &SearchRewardProgramsParams{
		timeout: timeout,
	}
}

// NewSearchRewardProgramsParamsWithContext creates a new SearchRewardProgramsParams object
// with the ability to set a context for a request.
func NewSearchRewardProgramsParamsWithContext(ctx context.Context) *SearchRewardProgramsParams {
	return &SearchRewardProgramsParams{
		Context: ctx,
	}
}

// NewSearchRewardProgramsParamsWithHTTPClient creates a new SearchRewardProgramsParams object
// with the ability to set a custom HTTPClient for a request.
func NewSearchRewardProgramsParamsWithHTTPClient(client *http.Client) *SearchRewardProgramsParams {
	return &SearchRewardProgramsParams{
		HTTPClient: client,
	}
}

/* SearchRewardProgramsParams contains all the parameters to send to the API endpoint
   for the search reward programs operation.

   Typically these are written to a http.Request.
*/
type SearchRewardProgramsParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the search reward programs params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SearchRewardProgramsParams) WithDefaults() *SearchRewardProgramsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the search reward programs params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SearchRewardProgramsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the search reward programs params
func (o *SearchRewardProgramsParams) WithTimeout(timeout time.Duration) *SearchRewardProgramsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the search reward programs params
func (o *SearchRewardProgramsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the search reward programs params
func (o *SearchRewardProgramsParams) WithContext(ctx context.Context) *SearchRewardProgramsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the search reward programs params
func (o *SearchRewardProgramsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the search reward programs params
func (o *SearchRewardProgramsParams) WithHTTPClient(client *http.Client) *SearchRewardProgramsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the search reward programs params
func (o *SearchRewardProgramsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *SearchRewardProgramsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}