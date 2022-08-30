// Code generated by go-swagger; DO NOT EDIT.

package account_information

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new account information API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for account information API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	GetAccount(params *GetAccountParams, opts ...ClientOption) (*GetAccountOK, error)

	SearchForAccounts(params *SearchForAccountsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*SearchForAccountsOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  GetAccount gets account

  Get a specific account
*/
func (a *Client) GetAccount(params *GetAccountParams, opts ...ClientOption) (*GetAccountOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetAccountParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getAccount",
		Method:             "GET",
		PathPattern:        "/accounts/{accountId}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetAccountReader{formats: a.formats},
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
	success, ok := result.(*GetAccountOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getAccount: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  SearchForAccounts searches for accounts

  Query all information for a set of accounts provided in the payload
*/
func (a *Client) SearchForAccounts(params *SearchForAccountsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*SearchForAccountsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewSearchForAccountsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "searchForAccounts",
		Method:             "GET",
		PathPattern:        "/accounts",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &SearchForAccountsReader{formats: a.formats},
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
	success, ok := result.(*SearchForAccountsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for searchForAccounts: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
