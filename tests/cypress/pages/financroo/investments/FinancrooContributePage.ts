export class FinancrooContributePage {
  private readonly headerLocator: string = `#contribute-header`;
  private readonly titleLocator: string = `#contribute-title`;
  private readonly amountCurrencyLocator: string = `#contribution-currency`;
  private readonly amountInputLocator: string = `#amount-to-contribute`;
  private readonly totalAmountLocator: string = `#total-amount`;

  private readonly getAccountLocator = (id: string) => `#account-id-${id}`;
  private readonly getAccountInputLocator = (id: string) => `#account-id-${id} input`;


  private readonly nextButtonLocator: string = `#next-button`;
        private readonly backToPortfolioButtonLocator: string = `#back-to-portfolio`;
  private readonly amountDisplayedLocator: string = `.MuiChip-label`;
  private readonly accountListLocator: string = `#accounts-list`;

          public contribute(amount: number): void {
            cy.get(this.titleLocator).contains("How much would you like to transfer?");
            cy.get(this.amountInputLocator).type(String(amount));
            cy.get(this.nextButtonLocator).click();
            cy.get(this.titleLocator).contains("PAYMENT TOTAL");
            cy.get(this.accountListLocator).contains("Checking account");
            cy.get(this.nextButtonLocator).click();
            cy.get(this.titleLocator).contains("INVESTMENT SUMMARY");
            cy.get(this.nextButtonLocator).click();
          }

  public contributeAmmount(amount: number, currency: string): void {
    cy.get(this.headerLocator).contains("Contributions");
    cy.get(this.titleLocator).contains("How much would you like to transfer?");
    cy.get(this.amountCurrencyLocator).contains(currency);
    cy.get(this.amountInputLocator).type(String(amount));
    cy.get(this.nextButtonLocator).contains("Next").click();
  }

  public contributePaymentMethod(
    amount: number,
    currency: string,
    accountId: string
  ): void {
    cy.get(this.headerLocator).contains("Select the payment method");
    cy.get(this.titleLocator).contains("PAYMENT TOTAL");
    cy.get(this.totalAmountLocator).contains(currency + " " + amount.toFixed(2));
    cy.get(this.getAccountLocator(accountId)).should("contain.text", currency)
    cy.get(this.getAccountInputLocator(accountId)).check();
    cy.get(this.nextButtonLocator).contains("Next").click();
  }

  public contributeInvestmentSummary(
    amount: number,
    currency: string,
    accountId: string
  ): void {
    cy.get(this.headerLocator).contains("Investment summary");
    cy.get(this.titleLocator).contains("INVESTMENT SUMMARY");
    cy.get(this.totalAmountLocator).contains(currency + " " + amount.toFixed(2));
    cy.get(this.getAccountLocator(accountId)).should("contain.text", "**** ***** **** " + accountId);
    cy.get(this.nextButtonLocator).contains("Confirm").click();
  }

  public assertItIsFinished(): void {
    cy.get(this.backToPortfolioButtonLocator).should(
      `contain.text`,
      `Back to portfolio`
    );
  }

  public assertAmount(amount: number, currency: string): void {
    cy.get(this.amountDisplayedLocator).should(
      `contain.text`,
      amount.toFixed(2) + " " + currency
    );
  }
}
