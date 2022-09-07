
export class FinancrooAccountsPage {
  private readonly accountsLocator: string = `.account-name`;
  private readonly investmentsTabLocator: string = `#investments-tab`;
  private readonly accountsTabLocator: string = `#accounts-tab`;
  private readonly disconnectAccountsButtonLocator: string = `#access-bank-button`;

  public assertAccounts(accounts: string[]): void {
    const accountElements = cy.get(this.accountsLocator);
    if (accounts.length) {
      accountElements.invoke(`text`).should(`equal`, accounts.join(``));
    } else {
      accountElements.should(`not.exist`);
    }
  }

  public goToInvestmentsTab(): void {
    cy.get(this.investmentsTabLocator).click();
  }

  public assertThatPageIsDisplayed(): void {
    cy.get(this.accountsTabLocator).should("have.text", "Accounts");
    cy.get(this.disconnectAccountsButtonLocator).should(
      "have.text",
      "disconnect"
    );
    cy.get(this.accountsLocator).should("be.visible");
  }

  public disconnectAccounts(): void {
    cy.get(this.disconnectAccountsButtonLocator).click();
  }

}
