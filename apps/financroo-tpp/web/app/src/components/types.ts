export type Account = {
  Account: {
    Identification: string;
    Name: string;
    SchemeName: string;
    SecondaryIdentification: string;
  }[];
  AccountId: string;
  AccountSubType: string;
  AccountType: string;
  BankId: string;
  Currency: string;
  MaturityDate: string;
  Nickname: string;
  OpeningDate: string;
  Status: string;
  StatusUpdateDateTime: string;
};

export type BanksResponse = {
  connected_banks: string[];
  expired_banks: string[];
};

export type Balance = {
  AccountId: string;
  Amount: string;
  BankId: string;
  Currency: string;
};

export type Transaction = {
  AccountId: string;
  Amount: {
    Amount: string;
    Currency: string;
  };
  Balance: {
    Amount: {
      Amount: string;
      Currency: string;
    };
    CreditDebitIndicator: string;
    Type: string;
  };
  BankTransactionCode: {
    Code: string;
    SubCode: string;
  };
  BookingDateTime: string;
  CreditDebitIndicator: string;
  ProprietaryBankTransactionCode: {
    Code: string;
    Issuer: string;
  };
  StatementReference: null;
  Status: string;
  TransactionId: string;
  TransactionInformation: string;
  TransactionReference: string;
  ValueDateTime: string;
  BankId: string;
};

export type BalancesResponse = {
  balances: Balance[];
};

export type AccountsResponse = {
  accounts: Account[];
};

export type TransactionsResponse = {
  transactions: Transaction[];
};

export type Filter = {
  accounts?: string[];
  months?: string[];
  categories?: string[];
};
