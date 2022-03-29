package main

type FDXClient struct {
}

var _ BankClient = &FDXClient{}

func NewFDXClient(config Config) *FDXClient {
	c := FDXClient{}
	return &c
}

func (c *FDXClient) GetInternalAccounts(id string) (InternalAccounts, error) {
	if id == "user" {
		return InternalAccounts{
			Accounts: []InternalAccount{
				{
					ID:   "96534987",
					Name: "Digital banking account",
					Balance: Balance{
						AccountID: "96534987",
						Amount: BalanceAmount{
							Amount:   "100",
							Currency: "USD",
						},
					},
				},
				{
					ID:   "1000001",
					Name: "Savings",
					Balance: Balance{
						AccountID: "1000001",
						Amount: BalanceAmount{
							Amount:   "150",
							Currency: "USD",
						},
					},
				},
				{
					ID:   "1000002",
					Name: "Savings 2",
					Balance: Balance{
						AccountID: "1000002",
						Amount: BalanceAmount{
							Amount:   "175",
							Currency: "USD",
						},
					},
				},
			},
		}, nil
	}

	return InternalAccounts{
		Accounts: []InternalAccount{
			{
				ID:   "96565987",
				Name: "Credit",
				Balance: Balance{
					AccountID: "96565987",
					Amount: BalanceAmount{
						Amount:   "100",
						Currency: "USD",
					},
				},
			},
			{
				ID:   "1122334455",
				Name: "Savings",
				Balance: Balance{
					AccountID: "1122334455",
					Amount: BalanceAmount{
						Amount:   "150",
						Currency: "USD",
					},
				},
			},
		},
	}, nil
}

// TODO: mock data holder cdr app doesn't even have this data yet
func (c *FDXClient) GetInternalBalances(id string) (BalanceResponse, error) {
	return BalanceResponse{}, nil
}
