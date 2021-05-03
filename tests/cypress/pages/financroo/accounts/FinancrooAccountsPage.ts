export class FinancrooAccountsPage {
  private readonly accountsLocator: string = `.account-name`;
  private readonly investmentsTabLocator: string = `#investments-tab`;

  public assertAccounts(accounts: string[]): void {
    const accountElements = cy.get(this.accountsLocator);
    if (accounts.length) {
      accountElements.invoke(`text`).should(`equal`, accounts.join(``))
    } else {
      accountElements.should(`not.exist`);
    }
  }

  public goToInvestmentsTab(): void {
    cy.get(this.investmentsTabLocator).click()
  }

}
