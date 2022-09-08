export type Consent = {
  AccountIDs: string[];
  ConsentID: string;
  client_id: string;
  tenant_id: string;
  server_id: string;
  Status: string;
  type: string;
  CreationDateTime: string;
  ExpirationDateTime: string;
  StatusUpdateDateTime: string;
  Permissions: string[] | null;
  DebtorAccountIdentification: string;
  DebtorAccountName: string;
  CreditorAccountIdentification: string;
  CreditorAccountName: string;
  Currency: string;
  Amount: string;
  CompletionDateTime: string;
};

export type ClientConsent = {
  id: string;
  name: string;
  logo_uri: string;
  client_uri: string;
  consents: Consent[];
};

export type ConsentAccount = {
  id: string;
  name: string;
};

export type ConsentsResponse = {
  client_consents: ClientConsent[];
  accounts: {
    accounts: ConsentAccount[];
  };
};
