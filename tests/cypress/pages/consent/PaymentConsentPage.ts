export class PaymentConsentPage {
  private readonly consentTitleLocator: string = `.consent-title`;
  private readonly totalAmountLocator: string = `#total-amount`;
  private readonly getAccountIdLocator = (id: string) => `#account-id-${id}`;
  private readonly confirmButtonLocator: string = `[value="confirm"]`;
  private readonly cancelButtonLocator: string = `[value="deny"]`;


  public assertThatConsentPageIsVisible(
    amount: number,
    currency: string,
    accountId: string
  ): void {
    cy.get(this.consentTitleLocator).contains("Account Confirm payment");
    cy.get(this.totalAmountLocator).contains(amount.toFixed(2) + " " + currency);
    cy.get(this.getAccountIdLocator(accountId)).should("contain.text", "**** **** **** " + accountId);
  }

  public clickConfirm(): void {
    this.clickButton(this.confirmButtonLocator, "Confirm");
  }

  public clickCancel(): void {
    this.clickButton(this.cancelButtonLocator, "Cancel");
  }

  private clickButton(locator: string, label: string): void {
    cy.get(locator).should("contain.text", label);
    cy.get(locator).click();
  }
}
