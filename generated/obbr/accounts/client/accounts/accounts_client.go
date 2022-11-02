// Code generated by go-swagger; DO NOT EDIT.

package accounts

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new accounts API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for accounts API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	AccountsGetAccounts(params *AccountsGetAccountsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*AccountsGetAccountsOK, error)

	AccountsGetAccountsAccountID(params *AccountsGetAccountsAccountIDParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*AccountsGetAccountsAccountIDOK, error)

	AccountsGetAccountsAccountIDBalances(params *AccountsGetAccountsAccountIDBalancesParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*AccountsGetAccountsAccountIDBalancesOK, error)

	AccountsGetAccountsAccountIDOverdraftLimits(params *AccountsGetAccountsAccountIDOverdraftLimitsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*AccountsGetAccountsAccountIDOverdraftLimitsOK, error)

	AccountsGetAccountsAccountIDTransactions(params *AccountsGetAccountsAccountIDTransactionsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*AccountsGetAccountsAccountIDTransactionsOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  AccountsGetAccounts accounts get accounts

  Método para obter a lista de contas depósito à vista, poupança e pagamento pré-pagas mantidas pelo cliente na instituição transmissora e para as quais ele tenha fornecido consentimento.
*/
func (a *Client) AccountsGetAccounts(params *AccountsGetAccountsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*AccountsGetAccountsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewAccountsGetAccountsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "accountsGetAccounts",
		Method:             "GET",
		PathPattern:        "/accounts",
		ProducesMediaTypes: []string{"application/json", "application/json; charset=utf-8"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &AccountsGetAccountsReader{formats: a.formats},
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
	success, ok := result.(*AccountsGetAccountsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*AccountsGetAccountsDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  AccountsGetAccountsAccountID accounts get accounts account Id

  Método para obter os dados de identificação da conta de depósito à vista, poupança ou pagamento pré-paga identificada por accountId mantida pelo cliente na instituição transmissora.
*/
func (a *Client) AccountsGetAccountsAccountID(params *AccountsGetAccountsAccountIDParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*AccountsGetAccountsAccountIDOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewAccountsGetAccountsAccountIDParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "accountsGetAccountsAccountId",
		Method:             "GET",
		PathPattern:        "/accounts/{accountId}",
		ProducesMediaTypes: []string{"application/json", "application/json; charset=utf-8"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &AccountsGetAccountsAccountIDReader{formats: a.formats},
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
	success, ok := result.(*AccountsGetAccountsAccountIDOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*AccountsGetAccountsAccountIDDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  AccountsGetAccountsAccountIDBalances accounts get accounts account Id balances

  Método para obter os saldos da conta de depósito à vista, poupança ou pagamento pré-paga identificada por accountId mantida pelo cliente na instituição transmissora.
*/
func (a *Client) AccountsGetAccountsAccountIDBalances(params *AccountsGetAccountsAccountIDBalancesParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*AccountsGetAccountsAccountIDBalancesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewAccountsGetAccountsAccountIDBalancesParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "accountsGetAccountsAccountIdBalances",
		Method:             "GET",
		PathPattern:        "/accounts/{accountId}/balances",
		ProducesMediaTypes: []string{"application/json", "application/json; charset=utf-8"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &AccountsGetAccountsAccountIDBalancesReader{formats: a.formats},
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
	success, ok := result.(*AccountsGetAccountsAccountIDBalancesOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*AccountsGetAccountsAccountIDBalancesDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  AccountsGetAccountsAccountIDOverdraftLimits accounts get accounts account Id overdraft limits

  Método para obter os limites da conta de depósito à vista, poupança ou pagamento pré-paga identificada por accountId mantida pelo cliente na instituição transmissora.
*/
func (a *Client) AccountsGetAccountsAccountIDOverdraftLimits(params *AccountsGetAccountsAccountIDOverdraftLimitsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*AccountsGetAccountsAccountIDOverdraftLimitsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewAccountsGetAccountsAccountIDOverdraftLimitsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "accountsGetAccountsAccountIdOverdraftLimits",
		Method:             "GET",
		PathPattern:        "/accounts/{accountId}/overdraft-limits",
		ProducesMediaTypes: []string{"application/json", "application/json; charset=utf-8"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &AccountsGetAccountsAccountIDOverdraftLimitsReader{formats: a.formats},
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
	success, ok := result.(*AccountsGetAccountsAccountIDOverdraftLimitsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*AccountsGetAccountsAccountIDOverdraftLimitsDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  AccountsGetAccountsAccountIDTransactions accounts get accounts account Id transactions

  Método para obter a lista de transações da conta de depósito à vista, poupança ou pagamento pré-paga identificada por accountId mantida pelo cliente na instituição transmissora.
*/
func (a *Client) AccountsGetAccountsAccountIDTransactions(params *AccountsGetAccountsAccountIDTransactionsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*AccountsGetAccountsAccountIDTransactionsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewAccountsGetAccountsAccountIDTransactionsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "accountsGetAccountsAccountIdTransactions",
		Method:             "GET",
		PathPattern:        "/accounts/{accountId}/transactions",
		ProducesMediaTypes: []string{"application/json", "application/json; charset=utf-8"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &AccountsGetAccountsAccountIDTransactionsReader{formats: a.formats},
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
	success, ok := result.(*AccountsGetAccountsAccountIDTransactionsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*AccountsGetAccountsAccountIDTransactionsDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}