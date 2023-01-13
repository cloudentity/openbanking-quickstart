export class ConsentSelfServiceApplicationPage {

  private readonly accountsTabLocator: string = `.accounts-tab`;
  private readonly accountsTabSelectedLocator: string = `[class*='accounts-tab'][aria-selected=true]`;
  private readonly paymentsTabLocator: string = `.payments-tab`;
  private readonly paymentsTabSelectedLocator: string = `[class*='payments-tab'][aria-selected=true]`;
  private readonly consentRowLocator: string = `.consent-row`
  private readonly consentRowAccountsLocator: string = `.consent-row > td:nth-child(2)`
  private readonly consentRowAmountLocator: string = `.consent-row > td:nth-child(5)`

  private readonly consentRowAuthorisedAccountLocator: string = `.consent-row:contains('Authorised')`


  public expandAccountsTab(): void {
    cy.get(this.accountsTabLocator).click();
    cy.get(this.accountsTabSelectedLocator).should('be.visible');
  }

  public expandPaymentsTab(): void {
    cy.get(this.paymentsTabLocator).click();
    cy.get(this.paymentsTabSelectedLocator).should('be.visible');
  }

  public checkAccount(accountID: string): void {
    cy.get(this.consentRowAccountsLocator).first().should("contain.text", accountID);
  }

  public checkAccountHasStatus(accountID: string, status: string): void {
    cy.get(this.consentRowLocator)
    .filter(`:contains(${accountID})`)
    .filter(`:contains(${status})`)
    .should("exist");
  }

  public checkAmount(currency: string, amount: number): void {
    cy.get(this.consentRowAmountLocator).first().should("contain.text", currency + " " + amount.toFixed(2));
  }

  public expandAccountConsentRow(): void {
    cy.get(this.consentRowLocator).first().click();
  }

  public expandPaymentConsentRow(): void {
    cy.get(this.consentRowLocator).first().click();
  }

  public assertNumberOfConsents(num: number): void {
    cy.get(this.consentRowAccountsLocator).should('have.length', num);
  }

  public assertAuthorisedAccountRowExists(accountID: string): void {
    cy.get(this.consentRowLocator)
    .filter(`:contains(${accountID})`)
    .filter(`:contains('Authorised')`)
    .should('have.length', 1);
  }

  public assertAuthorisedAccountRowDoesNotExist(accountID: string): void {
    cy.get(this.consentRowLocator)
    .filter(`:contains(${accountID})`)
    .filter(`:contains('Authorised')`)
    .should('have.length', 0);
  }
}
