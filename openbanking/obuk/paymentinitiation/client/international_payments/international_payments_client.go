// Code generated by go-swagger; DO NOT EDIT.

package international_payments

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new international payments API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for international payments API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	CreateInternationalPaymentConsents(params *CreateInternationalPaymentConsentsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*CreateInternationalPaymentConsentsCreated, error)

	CreateInternationalPayments(params *CreateInternationalPaymentsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*CreateInternationalPaymentsCreated, error)

	GetInternationalPaymentConsentsConsentID(params *GetInternationalPaymentConsentsConsentIDParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetInternationalPaymentConsentsConsentIDOK, error)

	GetInternationalPaymentConsentsConsentIDFundsConfirmation(params *GetInternationalPaymentConsentsConsentIDFundsConfirmationParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetInternationalPaymentConsentsConsentIDFundsConfirmationOK, error)

	GetInternationalPaymentsInternationalPaymentID(params *GetInternationalPaymentsInternationalPaymentIDParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetInternationalPaymentsInternationalPaymentIDOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  CreateInternationalPaymentConsents creates international payment consents
*/
func (a *Client) CreateInternationalPaymentConsents(params *CreateInternationalPaymentConsentsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*CreateInternationalPaymentConsentsCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateInternationalPaymentConsentsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "CreateInternationalPaymentConsents",
		Method:             "POST",
		PathPattern:        "/international-payment-consents",
		ProducesMediaTypes: []string{"application/jose+jwe", "application/json", "application/json; charset=utf-8"},
		ConsumesMediaTypes: []string{"application/jose+jwe", "application/json", "application/json; charset=utf-8"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &CreateInternationalPaymentConsentsReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*CreateInternationalPaymentConsentsCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for CreateInternationalPaymentConsents: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  CreateInternationalPayments creates international payments
*/
func (a *Client) CreateInternationalPayments(params *CreateInternationalPaymentsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*CreateInternationalPaymentsCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateInternationalPaymentsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "CreateInternationalPayments",
		Method:             "POST",
		PathPattern:        "/international-payments",
		ProducesMediaTypes: []string{"application/jose+jwe", "application/json", "application/json; charset=utf-8"},
		ConsumesMediaTypes: []string{"application/jose+jwe", "application/json", "application/json; charset=utf-8"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &CreateInternationalPaymentsReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*CreateInternationalPaymentsCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for CreateInternationalPayments: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetInternationalPaymentConsentsConsentID gets international payment consents
*/
func (a *Client) GetInternationalPaymentConsentsConsentID(params *GetInternationalPaymentConsentsConsentIDParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetInternationalPaymentConsentsConsentIDOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetInternationalPaymentConsentsConsentIDParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetInternationalPaymentConsentsConsentId",
		Method:             "GET",
		PathPattern:        "/international-payment-consents/{ConsentId}",
		ProducesMediaTypes: []string{"application/jose+jwe", "application/json", "application/json; charset=utf-8"},
		ConsumesMediaTypes: []string{"application/jose+jwe", "application/json", "application/json; charset=utf-8"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetInternationalPaymentConsentsConsentIDReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetInternationalPaymentConsentsConsentIDOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetInternationalPaymentConsentsConsentId: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetInternationalPaymentConsentsConsentIDFundsConfirmation gets international payment consents funds confirmation
*/
func (a *Client) GetInternationalPaymentConsentsConsentIDFundsConfirmation(params *GetInternationalPaymentConsentsConsentIDFundsConfirmationParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetInternationalPaymentConsentsConsentIDFundsConfirmationOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetInternationalPaymentConsentsConsentIDFundsConfirmationParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetInternationalPaymentConsentsConsentIdFundsConfirmation",
		Method:             "GET",
		PathPattern:        "/international-payment-consents/{ConsentId}/funds-confirmation",
		ProducesMediaTypes: []string{"application/jose+jwe", "application/json", "application/json; charset=utf-8"},
		ConsumesMediaTypes: []string{"application/jose+jwe", "application/json", "application/json; charset=utf-8"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetInternationalPaymentConsentsConsentIDFundsConfirmationReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetInternationalPaymentConsentsConsentIDFundsConfirmationOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetInternationalPaymentConsentsConsentIdFundsConfirmation: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetInternationalPaymentsInternationalPaymentID gets international payments
*/
func (a *Client) GetInternationalPaymentsInternationalPaymentID(params *GetInternationalPaymentsInternationalPaymentIDParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetInternationalPaymentsInternationalPaymentIDOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetInternationalPaymentsInternationalPaymentIDParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetInternationalPaymentsInternationalPaymentId",
		Method:             "GET",
		PathPattern:        "/international-payments/{InternationalPaymentId}",
		ProducesMediaTypes: []string{"application/jose+jwe", "application/json", "application/json; charset=utf-8"},
		ConsumesMediaTypes: []string{"application/jose+jwe", "application/json", "application/json; charset=utf-8"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetInternationalPaymentsInternationalPaymentIDReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetInternationalPaymentsInternationalPaymentIDOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetInternationalPaymentsInternationalPaymentId: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
