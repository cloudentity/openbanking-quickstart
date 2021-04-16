package main

type Permission struct {
	Cluster    string
	Permission string
	Language   string
}

var DataClusterLanguage = []Permission{
	{
		Cluster:    "Your Account Details",
		Permission: "ReadAccountsBasic",
		Language:   "Any other name by which you refer to this account and/or the currency of the account.",
	},
	{
		Cluster:    "Your Account Details",
		Permission: "ReadAccountsDetail",
		Language:   "Your account name, number and sort-code",
	},
	{
		Cluster:    "Your Account Details",
		Permission: "ReadBalances",
		Language:   "Your account balance",
	},
	{
		Cluster:    "Your Account Details",
		Permission: "ReadPAN",
		Language:   "Your card number",
	},
	{
		Cluster:    "Your Regular Payments",
		Permission: "ReadBeneficiariesBasic",
		Language:   "Payee agreements you have set up",
	},
	{
		Cluster:    "Your Regular Payments",
		Permission: "ReadBeneficiariesDetail",
		Language:   "Details of Payee agreements you have set up",
	},
	{
		Cluster:    "Your Regular Payments",
		Permission: "ReadDirectDebits",
		Language:   "Your Direct Debits",
	},
	{
		Cluster:    "Your Regular Payments",
		Permission: "ReadStandingOrdersBasic",
		Language:   "Your Standing Orders",
	},
	{
		Cluster:    "Your Regular Payments",
		Permission: "ReadStandingOrdersDetail",
		Language:   "Details of your Standing Orders",
	},
	{
		Cluster:    "Your Regular Payments",
		Permission: "ReadScheduledPaymentsBasic",
		Language:   "Recurring and future dated payments",
	},
	{
		Cluster:    "Your Regular Payments",
		Permission: "ReadScheduledPaymentsDetail",
		Language:   "Details of recurring and future dated payments",
	},
	{
		Cluster:    "Your Account Transactions",
		Permission: "ReadTransactionsBasic",
		Language:   "Your transactions",
	},
	{
		Cluster:    "Your Account Transactions",
		Permission: "ReadTransactionsDetail",
		Language:   "Details of your transactions",
	},
	{
		Cluster:    "Your Account Transactions",
		Permission: "ReadTransactionsCredits",
		Language:   "Your incoming transactions",
	},
	{
		Cluster:    "Your Account Transactions",
		Permission: "ReadTransactionsDebits",
		Language:   "Your outgoing transactions",
	},
	{
		Cluster:    "Your Statements",
		Permission: "ReadStatementsBasic",
		Language:   "Information contained in your statement",
	},
	{
		Cluster:    "Your Statements",
		Permission: "ReadStatementsDetail",
		Language:   "Details of information contained in your statement",
	},
	{
		Cluster:    "Your Account Features and Benefits",
		Permission: "ReadProducts",
		Language:   "Product details – fees, charges, interest, benefits/rewards",
	},
	{
		Cluster:    "Your Account Features and Benefits",
		Permission: "ReadOffers",
		Language:   "Offers available on your account",
	},
	{
		Cluster:    "Contact and party details",
		Permission: "ReadParty",
		Language:   "The full legal name(s) of account holder(s). Address(es), telephone number(s) and email address(es)*",
	},
}
