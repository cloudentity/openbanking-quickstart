export const permissionNameDescriptionMap = {
    "ReadAccountsBasic": "Ability to read basic account information",
    "ReadAccountsDetail": "Ability to read account identification details",
    "ReadBalances": "Ability to read all balance information",
    "ReadBeneficiariesBasic": "Ability to read basic beneficiary details",
    "ReadBeneficiariesDetail": "Ability to read account identification details for the beneficiary",
    "ReadDirectDebits": "Ability to read all direct debit information",
    "ReadStandingOrdersBasic": "Ability to read basic standing order information",
    "ReadStandingOrdersDetail": "Ability to read account identification details for beneficiary of the standing order",
    "ReadTransactionsBasic": "Ability to read basic transaction information",
    "ReadTransactionsDetail": "Ability to read transaction data elements which may hold silent party details",
    "ReadTransactionsCredits": "Ability to read only credit transactions",
    "ReadTransactionsDebits": "Ability to read only debit transactions",
    "ReadStatementsBasic": "Ability to read basic statement details",
    "ReadStatementsDetail": "Ability to read statement data elements which may leak other information about the account",
    "ReadProducts": "Ability to read all product information relating to the account",
    "ReadOffers": "Ability to read all offer information",
    "ReadParty": "Ability to read party information on the account owner.",
    "ReadPartyPSU": "Ability to read party information on the PSU logged in.",
    "ReadScheduledPaymentsBasic": "Ability to read basic statement details",
    "ReadScheduledPaymentsDetail": "",
    "ReadPAN": "Request to access PAN in the clear across the available endpoints. \n" +
        "\n" +
        "If this permission code is not in the account-access-consent, the AISP will receive a masked PAN.\n" +
        "\n" +
        "While an AISP may request to access PAN in the clear, an ASPSP may still respond with a masked PAN if:\n" +
        "\n" +
        "The ASPSP does not display PAN in the clear in existing online channels\n" +
        "The ASPSP takes a legal view to respond with only the masked PAN",
}

