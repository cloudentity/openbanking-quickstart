package main

type CDREnergyClient struct {
}

var _ BankClient = &CDREnergyClient{}

func NewCDREnergyClient(config Config) *CDREnergyClient {
	c := CDREnergyClient{}
	return &c
}

// TODO pull data from mock data holder
func (c *CDREnergyClient) GetInternalAccounts(id string) (InternalAccounts, error) {
	return InternalAccounts{
		Accounts: []InternalAccount{
			{
				ID:   "1234567890",
				Name: "Test Account 1",
				Balance: Balance{
					AccountID: "1234567890",
					Amount: BalanceAmount{
						Amount:   "100",
						Currency: "USD",
					},
				},
			},
			{
				ID:   "0987654321",
				Name: "Test Account 2",
				Balance: Balance{
					AccountID: "0987654321",
					Amount: BalanceAmount{
						Amount:   "150",
						Currency: "USD",
					},
				},
			},
		},
	}, nil
}

// TODO pull data from mock data holder
func (c *CDREnergyClient) GetInternalBalances(id string) (BalanceResponse, error) {
	return BalanceResponse{}, nil
}
