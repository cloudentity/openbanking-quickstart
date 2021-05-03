export class FinancrooAccountsPage {
  private readonly accountsLocator: string = `.account-name`;

  public assertAccounts(accounts: string[]): void {
    const accountElements = cy.get(this.accountsLocator);
    if (accounts.length) {
      accountElements.invoke(`text`).should(`equal`, accounts.join(``))
    } else {
      accountElements.should(`not.exist`);
    }
  }
}
