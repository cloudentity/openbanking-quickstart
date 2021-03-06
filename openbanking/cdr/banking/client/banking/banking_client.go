// Code generated by go-swagger; DO NOT EDIT.

package banking

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new banking API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for banking API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	GetAccountDetail(params *GetAccountDetailParams, opts ...ClientOption) (*GetAccountDetailOK, error)

	GetBalance(params *GetBalanceParams, opts ...ClientOption) (*GetBalanceOK, error)

	GetPayeeDetail(params *GetPayeeDetailParams, opts ...ClientOption) (*GetPayeeDetailOK, error)

	GetProductDetail(params *GetProductDetailParams, opts ...ClientOption) (*GetProductDetailOK, error)

	GetTransactionDetail(params *GetTransactionDetailParams, opts ...ClientOption) (*GetTransactionDetailOK, error)

	GetTransactions(params *GetTransactionsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetTransactionsOK, error)

	ListAccounts(params *ListAccountsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ListAccountsOK, error)

	ListBalancesBulk(params *ListBalancesBulkParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ListBalancesBulkOK, error)

	ListBalancesSpecificAccounts(params *ListBalancesSpecificAccountsParams, opts ...ClientOption) (*ListBalancesSpecificAccountsOK, error)

	ListDirectDebits(params *ListDirectDebitsParams, opts ...ClientOption) (*ListDirectDebitsOK, error)

	ListDirectDebitsBulk(params *ListDirectDebitsBulkParams, opts ...ClientOption) (*ListDirectDebitsBulkOK, error)

	ListDirectDebitsSpecificAccounts(params *ListDirectDebitsSpecificAccountsParams, opts ...ClientOption) (*ListDirectDebitsSpecificAccountsOK, error)

	ListPayees(params *ListPayeesParams, opts ...ClientOption) (*ListPayeesOK, error)

	ListProducts(params *ListProductsParams, opts ...ClientOption) (*ListProductsOK, error)

	ListScheduledPayments(params *ListScheduledPaymentsParams, opts ...ClientOption) (*ListScheduledPaymentsOK, error)

	ListScheduledPaymentsBulk(params *ListScheduledPaymentsBulkParams, opts ...ClientOption) (*ListScheduledPaymentsBulkOK, error)

	ListScheduledPaymentsSpecificAccounts(params *ListScheduledPaymentsSpecificAccountsParams, opts ...ClientOption) (*ListScheduledPaymentsSpecificAccountsOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  GetAccountDetail gets account detail

  Obtain detailed information on a single account.

Obsolete versions: [v1](includes/obsolete/get-account-detail-v1.html)
*/
func (a *Client) GetAccountDetail(params *GetAccountDetailParams, opts ...ClientOption) (*GetAccountDetailOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetAccountDetailParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getAccountDetail",
		Method:             "GET",
		PathPattern:        "/banking/accounts/{accountId}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetAccountDetailReader{formats: a.formats},
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
	success, ok := result.(*GetAccountDetailOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getAccountDetail: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetBalance gets balance

  Obtain the balance for a single specified account
*/
func (a *Client) GetBalance(params *GetBalanceParams, opts ...ClientOption) (*GetBalanceOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetBalanceParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getBalance",
		Method:             "GET",
		PathPattern:        "/banking/accounts/{accountId}/balance",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetBalanceReader{formats: a.formats},
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
	success, ok := result.(*GetBalanceOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getBalance: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetPayeeDetail gets payee detail

  Obtain detailed information on a single payee.

Note that the payee sub-structure should be selected to represent the payment destination only rather than any known characteristics of the payment recipient.

Obsolete versions: [v1](includes/obsolete/get-payee-detail-v1.html)
*/
func (a *Client) GetPayeeDetail(params *GetPayeeDetailParams, opts ...ClientOption) (*GetPayeeDetailOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetPayeeDetailParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getPayeeDetail",
		Method:             "GET",
		PathPattern:        "/banking/payees/{payeeId}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetPayeeDetailReader{formats: a.formats},
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
	success, ok := result.(*GetPayeeDetailOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getPayeeDetail: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetProductDetail gets product detail

  Obtain detailed information on a single product offered openly to the market.

Obsolete versions: [v1](includes/obsolete/get-product-detail-v1.html) [v2](includes/obsolete/get-product-detail-v2.html) [v3](includes/obsolete/get-product-detail-v3.html)
*/
func (a *Client) GetProductDetail(params *GetProductDetailParams, opts ...ClientOption) (*GetProductDetailOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetProductDetailParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getProductDetail",
		Method:             "GET",
		PathPattern:        "/banking/products/{productId}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetProductDetailReader{formats: a.formats},
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
	success, ok := result.(*GetProductDetailOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getProductDetail: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetTransactionDetail gets transaction detail

  Obtain detailed information on a transaction for a specific account
*/
func (a *Client) GetTransactionDetail(params *GetTransactionDetailParams, opts ...ClientOption) (*GetTransactionDetailOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetTransactionDetailParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getTransactionDetail",
		Method:             "GET",
		PathPattern:        "/banking/accounts/{accountId}/transactions/{transactionId}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetTransactionDetailReader{formats: a.formats},
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
	success, ok := result.(*GetTransactionDetailOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getTransactionDetail: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetTransactions gets transactions

  Obtain transactions for a specific account.

Some general notes that apply to all end points that retrieve transactions:

- Where multiple transactions are returned, transactions should be ordered according to effective date in descending order
- As the date and time for a transaction can alter depending on status and transaction type two separate date/times are included in the payload. There are still some scenarios where neither of these time stamps is available. For the purpose of filtering and ordering it is expected that the data holder will use the "effective" date/time which will be defined as:
  - Posted date/time if available, then
  - Execution date/time if available, then
  - A reasonable date/time nominated by the data holder using internal data structures
- For transaction amounts it should be assumed that a negative value indicates a reduction of the available balance on the account while a positive value indicates an increase in the available balance on the account
- For aggregated transactions (ie. groups of sub transactions reported as a single entry for the account) only the aggregated information, with as much consistent information accross the subsidiary transactions as possible, is required to be shared
*/
func (a *Client) GetTransactions(params *GetTransactionsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetTransactionsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetTransactionsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getTransactions",
		Method:             "GET",
		PathPattern:        "/banking/accounts/{accountId}/transactions",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetTransactionsReader{formats: a.formats},
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
	success, ok := result.(*GetTransactionsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getTransactions: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  ListAccounts lists accounts

  Obtain a list of accounts
*/
func (a *Client) ListAccounts(params *ListAccountsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ListAccountsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewListAccountsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "listAccounts",
		Method:             "GET",
		PathPattern:        "/banking/accounts",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ListAccountsReader{formats: a.formats},
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
	success, ok := result.(*ListAccountsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for listAccounts: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  ListBalancesBulk lists balances bulk

  Obtain balances for multiple, filtered accounts
*/
func (a *Client) ListBalancesBulk(params *ListBalancesBulkParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ListBalancesBulkOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewListBalancesBulkParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "listBalancesBulk",
		Method:             "GET",
		PathPattern:        "/banking/accounts/balances",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ListBalancesBulkReader{formats: a.formats},
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
	success, ok := result.(*ListBalancesBulkOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for listBalancesBulk: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  ListBalancesSpecificAccounts lists balances specific accounts

  Obtain balances for a specified list of accounts
*/
func (a *Client) ListBalancesSpecificAccounts(params *ListBalancesSpecificAccountsParams, opts ...ClientOption) (*ListBalancesSpecificAccountsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewListBalancesSpecificAccountsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "listBalancesSpecificAccounts",
		Method:             "POST",
		PathPattern:        "/banking/accounts/balances",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ListBalancesSpecificAccountsReader{formats: a.formats},
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
	success, ok := result.(*ListBalancesSpecificAccountsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for listBalancesSpecificAccounts: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  ListDirectDebits lists direct debits

  Obtain direct debit authorisations for a specific account
*/
func (a *Client) ListDirectDebits(params *ListDirectDebitsParams, opts ...ClientOption) (*ListDirectDebitsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewListDirectDebitsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "listDirectDebits",
		Method:             "GET",
		PathPattern:        "/banking/accounts/{accountId}/direct-debits",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ListDirectDebitsReader{formats: a.formats},
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
	success, ok := result.(*ListDirectDebitsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for listDirectDebits: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  ListDirectDebitsBulk lists direct debits bulk

  Obtain direct debit authorisations for multiple, filtered accounts
*/
func (a *Client) ListDirectDebitsBulk(params *ListDirectDebitsBulkParams, opts ...ClientOption) (*ListDirectDebitsBulkOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewListDirectDebitsBulkParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "listDirectDebitsBulk",
		Method:             "GET",
		PathPattern:        "/banking/accounts/direct-debits",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ListDirectDebitsBulkReader{formats: a.formats},
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
	success, ok := result.(*ListDirectDebitsBulkOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for listDirectDebitsBulk: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  ListDirectDebitsSpecificAccounts lists direct debits specific accounts

  Obtain direct debit authorisations for a specified list of accounts
*/
func (a *Client) ListDirectDebitsSpecificAccounts(params *ListDirectDebitsSpecificAccountsParams, opts ...ClientOption) (*ListDirectDebitsSpecificAccountsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewListDirectDebitsSpecificAccountsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "listDirectDebitsSpecificAccounts",
		Method:             "POST",
		PathPattern:        "/banking/accounts/direct-debits",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ListDirectDebitsSpecificAccountsReader{formats: a.formats},
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
	success, ok := result.(*ListDirectDebitsSpecificAccountsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for listDirectDebitsSpecificAccounts: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  ListPayees lists payees

  Obtain a list of pre-registered payees.

Obsolete versions: [v1](includes/obsolete/get-payees-v1.html)
*/
func (a *Client) ListPayees(params *ListPayeesParams, opts ...ClientOption) (*ListPayeesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewListPayeesParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "listPayees",
		Method:             "GET",
		PathPattern:        "/banking/payees",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ListPayeesReader{formats: a.formats},
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
	success, ok := result.(*ListPayeesOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for listPayees: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  ListProducts lists products

  Obtain a list of products that are currently openly offered to the market

Note that the results returned by this end point are expected to be ordered in descending order according to ``lastUpdated``.

### Conventions
In the product reference payloads there are a number of recurring conventions that are explained here, in one place.

#### Arrays Of Features

In the product detail payload there are a number of arrays articulating generic features, constraints, prices, etc. The intent of these arrays is as follows:

- Each element in an array has the same structure so that clients can reliably interpret the payloads
- Each element as a type element that is an enumeration of the specific aspect of a product being described, such as types of fees.
- Each element has a field name [additionalValue](#productfeaturetypedoc).  This is a generic field with contents that will vary based on the type of object being described. The contents of this field for the ADDITIONAL_CARDS feature is the number of cards allowed while the contents of this field for the MAX_LIMIT constraint would be the maximum credit limit allowed for the product.
- An element in these arrays of the same type may appear more than once. For instance, a product may offer two separate loyalty programs that the customer can select from. A fixed term mortgage may have different rates for different term lengths.
- An element in these arrays may contain an additionalInfo and additionalInfoUri field. The additionalInfo field is used to provide displayable text clarifying the purpose of the element in some way when the product is presented to a customer. The additionalInfoUri provides a link to externally hosted information specifically relevant to that feature of the product.
- Depending on the type of data being represented there may be additional specific fields.

#### URIs To More Information

As the complexities and nuances of a financial product can not easily be fully expressed in a data structure without a high degree of complexity it is necessary to provide additional reference information that a potential customer can access so that they are fully informed of the features and implications of the product. The payloads for product reference therefore contain numerous fields that are provided to allow the product holder to describe the product more fully using a web page hosted on their online channels.

These URIs do not need to all link to different pages. If desired, they can all link to a single hosted page and use difference HTML anchors to focus on a specific topic such as eligibility or fees.

#### Linkage To Accounts
From the moment that a customer applies for a product and an account is created the account and the product that spawned it will diverge.  Rates and features of the product may change and a discount may be negotiated for the account.

For this reason, while productCategory is a common field between accounts and products, there is no specific ID that can be used to link an account to a product within the regime.

Similarly, many of the fields and objects in the product payload will appear in the account detail payload but the structures and semantics are not identical as one refers to a product that can potentially be originated and one refers to an account that actual has been instantiated and created along with the associated decisions inherent in that process.

#### Dates
It is expected that data consumers needing this data will call relatively frequently to ensure the data they have is representative of the current offering from a bank.  To minimise the volume and frequency of these calls the ability to set a lastUpdated field with the date and time of the last update to this product is included.  A call for a list of products can then be filtered to only return products that have been updated since the last time that data was obtained using the updated-since query parameter.

In addition, the concept of effective date and time has also been included.  This allows for a product to be marked for obsolescence, or introduction, from a certain time without the need for an update to show that a product has been changed.  The inclusion of these dates also removes the need to represent deleted products in the payload.  Products that are no long offered can be marked not effective for a few weeks before they are then removed from the product set as an option entirely.

Obsolete versions: [v1](includes/obsolete/get-products-v1.html) [v2](includes/obsolete/get-products-v2.html)
*/
func (a *Client) ListProducts(params *ListProductsParams, opts ...ClientOption) (*ListProductsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewListProductsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "listProducts",
		Method:             "GET",
		PathPattern:        "/banking/products",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ListProductsReader{formats: a.formats},
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
	success, ok := result.(*ListProductsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for listProducts: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  ListScheduledPayments lists scheduled payments

  Obtain scheduled, outgoing payments for a specific account
*/
func (a *Client) ListScheduledPayments(params *ListScheduledPaymentsParams, opts ...ClientOption) (*ListScheduledPaymentsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewListScheduledPaymentsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "listScheduledPayments",
		Method:             "GET",
		PathPattern:        "/banking/accounts/{accountId}/payments/scheduled",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ListScheduledPaymentsReader{formats: a.formats},
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
	success, ok := result.(*ListScheduledPaymentsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for listScheduledPayments: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  ListScheduledPaymentsBulk lists scheduled payments bulk

  Obtain scheduled payments for multiple, filtered accounts that are the source of funds for the payments
*/
func (a *Client) ListScheduledPaymentsBulk(params *ListScheduledPaymentsBulkParams, opts ...ClientOption) (*ListScheduledPaymentsBulkOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewListScheduledPaymentsBulkParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "listScheduledPaymentsBulk",
		Method:             "GET",
		PathPattern:        "/banking/payments/scheduled",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ListScheduledPaymentsBulkReader{formats: a.formats},
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
	success, ok := result.(*ListScheduledPaymentsBulkOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for listScheduledPaymentsBulk: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  ListScheduledPaymentsSpecificAccounts lists scheduled payments specific accounts

  Obtain scheduled payments for a specified list of accounts
*/
func (a *Client) ListScheduledPaymentsSpecificAccounts(params *ListScheduledPaymentsSpecificAccountsParams, opts ...ClientOption) (*ListScheduledPaymentsSpecificAccountsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewListScheduledPaymentsSpecificAccountsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "listScheduledPaymentsSpecificAccounts",
		Method:             "POST",
		PathPattern:        "/banking/payments/scheduled",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ListScheduledPaymentsSpecificAccountsReader{formats: a.formats},
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
	success, ok := result.(*ListScheduledPaymentsSpecificAccountsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for listScheduledPaymentsSpecificAccounts: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
