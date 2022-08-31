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

  CREDIT_CARDS_ACCOUNTS_BILLS_TRANSACTIONS_READ: {
    Cluster: "TODO",
    Language: "TODO",
  },
  UNARRANGED_ACCOUNTS_OVERDRAFT_READ: {
    Cluster: "TODO",
    Language: "TODO",
  },
  ACCOUNTS_READ: {
    Cluster: "TODO",
    Language: "TODO",
  },
  LOANS_READ: {
    Cluster: "TODO",
    Language: "TODO",
  },
  LOANS_SCHEDULED_INSTALMENTS_READ: {
    Cluster: "TODO",
    Language: "TODO",
  },
  FINANCINGS_SCHEDULED_INSTALMENTS_READ: {
    Cluster: "TODO",
    Language: "TODO",
  },
  UNARRANGED_ACCOUNTS_OVERDRAFT_WARRANTIES_READ: {
    Cluster: "TODO",
    Language: "TODO",
  },
  UNARRANGED_ACCOUNTS_OVERDRAFT_PAYMENTS_READ: {
    Cluster: "TODO",
    Language: "TODO",
  },
  INVOICE_FINANCINGS_WARRANTIES_READ: {
    Cluster: "TODO",
    Language: "TODO",
  },
  CUSTOMERS_PERSONAL_IDENTIFICATIONS_READ: {
    Cluster: "TODO",
    Language: "TODO",
  },
  CREDIT_CARDS_ACCOUNTS_LIMITS_READ: {
    Cluster: "TODO",
    Language: "TODO",
  },
  FINANCINGS_READ: {
    Cluster: "TODO",
    Language: "TODO",
  },
  INVOICE_FINANCINGS_PAYMENTS_READ: {
    Cluster: "TODO",
    Language: "TODO",
  },
  RESOURCES_READ: {
    Cluster: "TODO",
    Language: "TODO",
  },
  ACCOUNTS_BALANCES_READ: {
    Cluster: "TODO",
    Language: "TODO",
  },
  ACCOUNTS_OVERDRAFT_LIMITS_READ: {
    Cluster: "TODO",
    Language: "TODO",
  },
  LOANS_WARRANTIES_READ: {
    Cluster: "TODO",
    Language: "TODO",
  },
  UNARRANGED_ACCOUNTS_OVERDRAFT_SCHEDULED_INSTALMENTS_READ: {
    Cluster: "TODO",
    Language: "TODO",
  },
  CREDIT_CARDS_ACCOUNTS_BILLS_READ: {
    Cluster: "TODO",
    Language: "TODO",
  },
  INVOICE_FINANCINGS_READ: {
    Cluster: "TODO",
    Language: "TODO",
  },
  CUSTOMERS_PERSONAL_ADITTIONALINFO_READ: {
    Cluster: "TODO",
    Language: "TODO",
  },
  FINANCINGS_WARRANTIES_READ: {
    Cluster: "TODO",
    Language: "TODO",
  },
  INVOICE_FINANCINGS_SCHEDULED_INSTALMENTS_READ: {
    Cluster: "TODO",
    Language: "TODO",
  },
  CUSTOMERS_BUSINESS_IDENTIFICATIONS_READ: {
    Cluster: "TODO",
    Language: "TODO",
  },
  CREDIT_CARDS_ACCOUNTS_READ: {
    Cluster: "TODO",
    Language: "TODO",
  },
  CREDIT_CARDS_ACCOUNTS_TRANSACTIONS_READ: {
    Cluster: "TODO",
    Language: "TODO",
  },
  LOANS_PAYMENTS_READ: {
    Cluster: "TODO",
    Language: "TODO",
  },
  CUSTOMERS_BUSINESS_ADITTIONALINFO_READ: {
    Cluster: "TODO",
    Language: "TODO",
  },
  ACCOUNTS_TRANSACTIONS_READ: {
    Cluster: "TODO",
    Language: "TODO",
  },
  FINANCINGS_PAYMENTS_READ: {
    Cluster: "TODO",
    Language: "TODO",
  },
};

export const currencyDict = {
  USD: "$",
  GBP: "£",
  EUR: "€",
};

// TODO:
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
    // ...theme.custom.body2,
    marginBottom: 24,
  },
  subHeader: {
    //  ...theme.custom.caption,
    textTransform: "uppercase" as "uppercase",
    fontWeight: "bold" as "bold",
    color: "#002D4C",
    margin: "16px 0",
  },

  cardTitle: {
    fontWeight: "bold" as "bold",
    marginBottom: 4,
    fontSize: 12,
  },
  cardContent: {
    // ...theme.custom.caption,
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
      //   ...theme.custom.body2,
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
