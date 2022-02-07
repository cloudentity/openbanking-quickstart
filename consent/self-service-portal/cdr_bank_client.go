package main

type CDRBankClient struct{}

func (c *CDRBankClient) GetInternalAccounts(subject string) (InternalAccounts, error) {
	if subject == "user" {
		return InternalAccounts{
				Accounts: []InternalAccount{
					{
						ID:   "96534987",
						Name: "Digital banking account",
					},
					{
						ID:   "1000001",
						Name: "Savings",
					},
					{
						ID:   "1000002",
						Name: "Savings 2",
					},
				},
			},
			nil
	}

	return InternalAccounts{
			Accounts: []InternalAccount{
				{
					ID:   "96565987",
					Name: "Credit",
				},
				{
					ID:   "1122334455",
					Name: "Savings",
				},
			},
		},
		nil
}
