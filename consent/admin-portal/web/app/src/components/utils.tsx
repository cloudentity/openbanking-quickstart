import React from "react";

import Chip from "./Chip";
import SearchInput from "./SearchInput";
import { theme } from "../theme";

export const drawerStyles = {
  name: {
    fontWeight: "normal" as "normal",
    fontSize: 20,
    lineHeight: "32px",
    margin: "0 16px",
  },
  logo: {
    border: "1.5px solid #F4F4F4",
    borderRadius: 4,
    width: 48,
    height: 48,
    objectFit: "contain" as "contain",
  },
  drawerHeader: {
    height: 72,
    backgroundColor: "#F7FAFF",
    width: "100%",
    display: "flex",
    alignItems: "center",
    paddingLeft: 32,
    fontWeight: 600,
    fontSize: 16,
    lineHeight: "24px",
    color: "#BD271E",
  },
  headerContent: {
    padding: "12px 24px",
    display: "flex" as "flex",
    alignItems: "center" as "center",
    width: "100%",
  },
  purposeHeader: {
    fontWeight: "bold" as "bold",
  },
  purpose: {
    ...theme.custom.body2,
    marginBottom: 24,
  },
  subHeader: {
    ...theme.custom.caption,
    textTransform: "uppercase" as "uppercase",
    fontWeight: "bold" as "bold",
    color: "#002D4C",
    margin: "16px 0",
    "&:first-child": {
      borderBottom: "solid 1px #ECECEC",
      paddingBottom: 12,
    },
  },
  cardTitle: {
    fontWeight: "bold" as "bold",
    marginBottom: 4,
    fontSize: 12,
  },
  cardContent: {
    ...theme.custom.caption,
  },
  card: {
    backgroundColor: "#FCFCFF",
    border: "1px solid #ECECEC",
    boxSizing: "border-box" as "border-box",
    borderRadius: 4,
    padding: 16,
    marginRight: 16,
    marginBottom: 16,
  },
  ulList: {
    marginTop: 0,
    paddingLeft: 16,
    "& > li": {
      ...theme.custom.body2,
    },
  },
  detailsTitle: {
    fontWeight: "bold" as "bold",
    fontSize: 14,
    lineHeight: "24px",
    marginTop: 24,
    marginBottom: 4,
  },
};

const monthNames = [
  "January",
  "February",
  "March",
  "April",
  "May",
  "June",
  "July",
  "August",
  "September",
  "October",
  "November",
  "December",
];

export function getDate(date) {
  try {
    const d = new Date(date);
    const year = d.getFullYear();
    if (year === 1 || isNaN(year)) return "N/A";
    return `${d.getDate()} ${monthNames[d.getMonth()]} ${d.getFullYear()}`;
  } catch (err) {
    console.error(err);
    return "N/A";
  }
}

export const permissionsDict = {
  CommonCustomerBasicRead: {
    Cluster: "Your Name and occupation",
    Language: "Name Occupation",
  },
  ReadAccountsBasic: {
    Cluster: "Your Account Details",
    Language:
      "Any other name by which you refer to this account and/or the currency of the account.",
  },
  ReadAccountsDetail: {
    Cluster: "Your Account Details",
    Language: "Your account name, number and sort-code",
  },
  ReadBalances: {
    Cluster: "Your Account Details",
    Language: "Your account balance",
  },
  ReadPAN: {
    Cluster: "Your Account Details",
    Language: "Your card number",
  },
  ReadBeneficiariesBasic: {
    Cluster: "Your Regular Payments",
    Language: "Payee agreements you have set up",
  },
  ReadBeneficiariesDetail: {
    Cluster: "Your Regular Payments",
    Language: "Details of Payee agreements you have set up",
  },
  ReadDirectDebits: {
    Cluster: "Your Regular Payments",
    Language: "Your Direct Debits",
  },
  ReadStandingOrdersBasic: {
    Cluster: "Your Regular Payments",
    Language: "Your Standing Orders",
  },
  ReadStandingOrdersDetail: {
    Cluster: "Your Regular Payments",
    Language: "Details of your Standing Orders",
  },
  ReadScheduledPaymentsBasic: {
    Cluster: "Your Regular Payments",
    Language: "Recurring and future dated payments",
  },
  ReadScheduledPaymentsDetail: {
    Cluster: "Your Regular Payments",
    Language: "Details of recurring and future dated payments",
  },
  ReadTransactionsBasic: {
    Cluster: "Your Account Transactions",
    Language: "Your transactions",
  },
  ReadTransactionsDetail: {
    Cluster: "Your Account Transactions",
    Language: "Details of your transactions",
  },
  ReadTransactionsCredits: {
    Cluster: "Your Account Transactions",
    Language: "Your incoming transactions",
  },
  ReadTransactionsDebits: {
    Cluster: "Your Account Transactions",
    Language: "Your outgoing transactions",
  },
  ReadStatementsBasic: {
    Cluster: "Your Statements",
    Language: "Information contained in your statement",
  },
  ReadStatementsDetail: {
    Cluster: "Your Statements",
    Language: "Details of information contained in your statement",
  },
  ReadProducts: {
    Cluster: "Your Account Features and Benefits",
    Language: "Product details – fees, charges, interest, benefits/rewards",
  },
  ReadOffers: {
    Cluster: "Your Account Features and Benefits",
    Language: "Offers available on your account",
  },
  ReadParty: {
    Cluster: "Contact and party details",
    Language:
      "The full legal name(s) of account holder(s). Address(es), telephone number(s) and email address(es)*",
  },
};

const availableConstentTypes = [
  "consents",
  "account_access",
  "domestic_payment",
  "cdr_arrangement",
  "domestic_scheduled_payment",
  "domestic_standing_order",
  "file_payment",
  "international_payment",
  "international_scheduled_payment",
  "international_standing_order",
];

export const availableConstentTypesJoined = availableConstentTypes.join(",");

export function getRawConsents(consents) {
  return consents
    .filter((c) => availableConstentTypes.includes(c.consent_type))
    .map((c) => {
      return {
        consent_type: c.consent_type,
        consent: c,
        accounts: c?.account_ids ?? [],
      };
    });
}

export enum ConsentStatus {
  Active = "Active",
  Inactive = "Inactive",
  Authorised = "Authorised",
}

export function getChipForStatus(client?: ClientType) {
  return (
    (client?.mainStatus === ConsentStatus.Active && (
      <Chip type="active" id={`status-active-${client?.client_id}`}>
        Active
      </Chip>
    )) ||
    (client?.mainStatus === ConsentStatus.Authorised && (
      <Chip type="active" id={`status-authorised-${client?.client_id}`}>
        Authorised
      </Chip>
    )) ||
    (client?.mainStatus === ConsentStatus.Inactive && (
      <Chip type="inactive" id={`status-inactive-${client?.client_id}`}>
        Inactive
      </Chip>
    )) ||
    null
  );
}

export type ClientType = {
  client_id: string;
  client_name: string;
  client_uri: string;
  provider_type: string;
  consents: {
    Client: {
      client_uri: string;
      id: string;
      name: string;
    };
    account_ids: string[];
    consent_id: string;
    client_id: string;
    tenant_id: string;
    server_id: string;
    status: string;
    consent_type: string;
    created_at: string;
    expires_at?: string;
    updated_at?: string;
    completed_at?: string | null;
    permissions?: string[];

    currency: string;
    amount: string;
  }[];
  mainStatus?: ConsentStatus;
};

export const handleSearch =
  (searchText: string) =>
  (history: any, accounts?: { [accountId: string]: string[] }) => {
    if (accounts) {
      const foundAccount = accounts[searchText];
      history.push({
        pathname: `/accounts/${searchText}`,
        state: { clientIds: foundAccount || null, accounts },
      });
    }
  };

export const currencyDict = {
  USD: "$",
  GBP: "£",
  EUR: "€",
};

export function getStatus(client: ClientType) {
  const accountConsents = ["account_access", "consents", "cdr_arrangement"];
  const found = client?.consents?.find(
    (consent) =>
      consent &&
      accountConsents.includes(consent.consent_type) &&
      consent.status === "Authorised"
  );
  return found ? ConsentStatus.Active : ConsentStatus.Inactive;
}

export function enrichClientWithStatus(client: ClientType) {
  return {
    ...client,
    mainStatus: getStatus(client),
  };
}

export const searchTabs = (
  onSearch: (searchText: string) => void,
  inputValue?: string
) => [
  {
    key: "account",
    label: "Account",
    content: (
      <div>
        <SearchInput
          placeholder="Search by account number"
          onSearch={onSearch}
          inputValue={inputValue}
        />
      </div>
    ),
  },
];
