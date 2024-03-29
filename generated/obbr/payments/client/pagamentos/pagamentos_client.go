// Code generated by go-swagger; DO NOT EDIT.

package pagamentos

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new pagamentos API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for pagamentos API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	PaymentsGetConsentsConsentID(params *PaymentsGetConsentsConsentIDParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PaymentsGetConsentsConsentIDOK, error)

	PaymentsGetPixPaymentsPaymentID(params *PaymentsGetPixPaymentsPaymentIDParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PaymentsGetPixPaymentsPaymentIDOK, error)

	PaymentsPatchPixPayments(params *PaymentsPatchPixPaymentsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PaymentsPatchPixPaymentsOK, error)

	PaymentsPostConsents(params *PaymentsPostConsentsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PaymentsPostConsentsCreated, error)

	PaymentsPostPixPayments(params *PaymentsPostPixPaymentsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PaymentsPostPixPaymentsCreated, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  PaymentsGetConsentsConsentID payments get consents consent Id

  Mtodo para consulta do consentimento para a iniciao de pagamento.
*/
func (a *Client) PaymentsGetConsentsConsentID(params *PaymentsGetConsentsConsentIDParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PaymentsGetConsentsConsentIDOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPaymentsGetConsentsConsentIDParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "paymentsGetConsentsConsentId",
		Method:             "GET",
		PathPattern:        "/consents/{consentId}",
		ProducesMediaTypes: []string{"application/json", "application/json; charset=utf-8", "application/jwt"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PaymentsGetConsentsConsentIDReader{formats: a.formats},
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
	success, ok := result.(*PaymentsGetConsentsConsentIDOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*PaymentsGetConsentsConsentIDDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  PaymentsGetPixPaymentsPaymentID payments get pix payments payment Id

  Mtodo para consultar uma iniciao de pagamento.
*/
func (a *Client) PaymentsGetPixPaymentsPaymentID(params *PaymentsGetPixPaymentsPaymentIDParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PaymentsGetPixPaymentsPaymentIDOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPaymentsGetPixPaymentsPaymentIDParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "paymentsGetPixPaymentsPaymentId",
		Method:             "GET",
		PathPattern:        "/pix/payments/{paymentId}",
		ProducesMediaTypes: []string{"application/json", "application/json; charset=utf-8", "application/jwt"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PaymentsGetPixPaymentsPaymentIDReader{formats: a.formats},
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
	success, ok := result.(*PaymentsGetPixPaymentsPaymentIDOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*PaymentsGetPixPaymentsPaymentIDDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  PaymentsPatchPixPayments payments patch pix payments

  Mtodo para revogao do consentimento.
*/
func (a *Client) PaymentsPatchPixPayments(params *PaymentsPatchPixPaymentsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PaymentsPatchPixPaymentsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPaymentsPatchPixPaymentsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "paymentsPatchPixPayments",
		Method:             "PATCH",
		PathPattern:        "/consents/{consentId}",
		ProducesMediaTypes: []string{"application/json", "application/json; charset=utf-8", "application/jwt"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PaymentsPatchPixPaymentsReader{formats: a.formats},
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
	success, ok := result.(*PaymentsPatchPixPaymentsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*PaymentsPatchPixPaymentsDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  PaymentsPostConsents payments post consents

  Mtodo de criao do consentimento para a iniciao de pagamento.
*/
func (a *Client) PaymentsPostConsents(params *PaymentsPostConsentsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PaymentsPostConsentsCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPaymentsPostConsentsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "paymentsPostConsents",
		Method:             "POST",
		PathPattern:        "/consents",
		ProducesMediaTypes: []string{"application/json", "application/json; charset=utf-8", "application/jwt"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PaymentsPostConsentsReader{formats: a.formats},
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
	success, ok := result.(*PaymentsPostConsentsCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*PaymentsPostConsentsDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  PaymentsPostPixPayments payments post pix payments

  Mtodo para criar uma iniciao de pagamento.
*/
func (a *Client) PaymentsPostPixPayments(params *PaymentsPostPixPaymentsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PaymentsPostPixPaymentsCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPaymentsPostPixPaymentsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "paymentsPostPixPayments",
		Method:             "POST",
		PathPattern:        "/pix/payments",
		ProducesMediaTypes: []string{"application/json", "application/json; charset=utf-8", "application/jwt"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PaymentsPostPixPaymentsReader{formats: a.formats},
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
	success, ok := result.(*PaymentsPostPixPaymentsCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*PaymentsPostPixPaymentsDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
