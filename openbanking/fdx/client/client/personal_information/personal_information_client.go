// Code generated by go-swagger; DO NOT EDIT.

package personal_information

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new personal information API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for personal information API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	GetAccountContact(params *GetAccountContactParams, opts ...ClientOption) (*GetAccountContactOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  GetAccountContact gets account contact

  Get contact information on the account
*/
func (a *Client) GetAccountContact(params *GetAccountContactParams, opts ...ClientOption) (*GetAccountContactOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetAccountContactParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getAccountContact",
		Method:             "GET",
		PathPattern:        "/accounts/{accountId}/contact",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetAccountContactReader{formats: a.formats},
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
	success, ok := result.(*GetAccountContactOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getAccountContact: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
