package main

type CDREnergyClient struct {
}

var _ BankClient = &CDREnergyClient{}

func NewCDREnergyClient(config Config) *CDREnergyClient {
	c := CDREnergyClient{}
	return &c
}

func (c *CDREnergyClient) GetInternalAccounts(id string) (InternalAccounts, error) {
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

// TODO: mock data holder cdr app doesn't even have this data yet
func (c *CDREnergyClient) GetInternalBalances(id string) (BalanceResponse, error) {
	return BalanceResponse{}, nil
}
