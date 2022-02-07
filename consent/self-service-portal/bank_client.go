package main

type BankClient interface {
	GetInternalAccounts(subject string) (InternalAccounts, error)
}

type InternalAccounts struct {
	Accounts []InternalAccount `json:"accounts"`
}

type InternalAccount struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

func (i *InternalAccounts) GetAccountIDs() []string {
	var accountIDs []string
	for _, account := range i.Accounts {
		accountIDs = append(accountIDs, account.ID)
	}
	return accountIDs
}
