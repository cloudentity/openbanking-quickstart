package main

type BankClient interface {
	GetInternalAccounts(subject string) (InternalAccounts, error)
	GetInternalBalances(subject string) (BalanceResponse, error)
}

type InternalAccounts struct {
	Accounts []InternalAccount `json:"accounts"`
}

type InternalAccount struct {
	ID          string  `json:"id"`
	Name        string  `json:"name"`
	Balance     Balance `json:"balance"`
	Preselected bool    `json:"preselected"`
}

type Balance struct {
	AccountID string        `json:"AccountId"`
	Amount    BalanceAmount `json:"Amount"`
}

type BalanceAmount struct {
	Amount   string
	Currency string
}

type BalanceResponse struct {
	Data BalanceData `json:"Data"`
}

type BalanceData struct {
	Balance []Balance `json:"Balance"`
}
