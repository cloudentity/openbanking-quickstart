export class FinancrooContributePage {
  private readonly headerLocator: string = `#contribute-header`;
  private readonly titleLocator: string = `#contribute-title`;
  private readonly amountCurrencyLocator: string = `#contribution-currency`;
  private readonly amountInputLocator: string = `#amount-to-contribute`;
  private readonly totalAmountLocator: string = `#total-amount`;
  private readonly nextButtonLocator: string = `#next-button`;
  private readonly getAccountLocator = (id: string) => `#account-id-${id}`;
  private readonly getAccountInputLocator = (id: string) => `#account-id-${id} input`;


  public contributeAmount(amount: number, currency: string): void {
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

}
