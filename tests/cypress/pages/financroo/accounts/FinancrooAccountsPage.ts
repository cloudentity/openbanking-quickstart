
export class FinancrooAccountsPage {
  private readonly accountsLocator: string = `.account-name`;
  private readonly accountsSyncedNumberLocator: string = `#accounts-synced-number`;
  private readonly investmentsTabLocator: string = `#investments-tab`;
  private readonly accountsTabLocator: string = `#accounts-tab`;
  private readonly disconnectAccountsButtonLocator: string = `#access-bank-button`;
  private readonly connectBankLocator: string = "#connect-bank";
  private getAccountLocator = (id: string) => `#account-id-${id}`;

  public assertAccounts(accounts: string[]): void {
    const accountElements = cy.get(this.accountsLocator);
    if (accounts.length) {
      accountElements.invoke(`text`).should(`equal`, accounts.join(``));
    } else {
      accountElements.should(`not.exist`);
    }
  }

  public assertAccountsIds(accountsIds: string[]): void {
    accountsIds.forEach(accountId => {
      cy.get(this.getAccountLocator(accountId)).should(`be.visible`)
    });
  }

  public assertAccountsSyncedNumber(accountsNumber): void {
    cy.get(this.accountsSyncedNumberLocator).should(`have.text`, `${accountsNumber} accounts synced`)

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
    cy.get(this.accountsSyncedNumberLocator).should("be.visible");
  }

  public assertThatAccountsAreDisconnected(): void {
    cy.get(this.accountsTabLocator).should("have.text", "Accounts");
    cy.get(this.disconnectAccountsButtonLocator).should(
      "have.text",
      "reconnect"
    );
    cy.get(this.accountsSyncedNumberLocator).should("be.visible");
    this.assertAccountsSyncedNumber(0);
  }

  public disconnectAccounts(): void {
    cy.get(this.disconnectAccountsButtonLocator).click();
  }

  public addNewBankAccount(): void {
    cy.get(this.connectBankLocator).click()
  }

}
