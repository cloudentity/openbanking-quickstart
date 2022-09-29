export class AccountConsentPage {
  private readonly confirmButtonLocator: string = `[value="confirm"]`;
  private readonly cancelButtonLocator: string = `[value="deny"]`;
  private readonly permissionNameLocator: string = `[data-desc-id="account_permissions"] .permission-name`;
  private readonly permissionContentLocator: string = `[data-desc-id="account_permissions"] .caption`;
  private readonly permissionRowLocator: string = `[data-desc-id="account_permissions"] li`;
  private readonly expandPermissionsButtonLocator: string = `[data-icon-id="account_permissions"]`
  private readonly accountsIdsLocator: string = `[id^="account-id"]`;
  private readonly getAccountIdLocator = (id: string) => `#account-id-${id}`;

  public checkAccounts(accounts: string[]): void {
    this.uncheckAllAccounts()
    accounts.forEach(account => cy.get(this.getAccountIdLocator(account)).check())
  }

  public uncheckAllAccounts(): void {
    cy.get(this.accountsIdsLocator).uncheck()
  }

  public checkAllAccounts(): void {
    cy.get(this.accountsIdsLocator).check()
  }

  public expandPermissions(): void {
    cy.get(this.expandPermissionsButtonLocator).click();
  }

  public assertPermissionsDetails(name: string, content: string): void {
      cy.get(this.permissionNameLocator).should('contain.text', name)
      cy.get(this.permissionContentLocator).should('contain.text', content)
  }

  public assertPermissions(length: number): void {
    cy.get(this.permissionRowLocator).should('have.length', length)
  }

  public clickContinue(): void {
    this.clickButton(this.confirmButtonLocator, 'Continue');
  }

  public clickAgree(): void {
    this.clickButton(this.confirmButtonLocator, 'I Agree');
  }

  public clickConfirm(): void {
    this.clickButton(this.confirmButtonLocator, 'Confirm');
  }

  public clickCancel(): void {
    this.clickButton(this.cancelButtonLocator, 'Cancel');
  }

  public assertThatPageIsNotVisible(): void {
    cy.get(this.cancelButtonLocator, { timeout: 5000 } ).should('not.exist');
    cy.get(this.confirmButtonLocator, { timeout: 5000 } ).should('not.exist');
  }

  public assertThatAccountsAreNotVisible(accounts: string[]): void {
    accounts.forEach(account => cy.get(this.getAccountIdLocator(account)).should('not.be.visible'))
  }

  private clickButton(locator: string, label: string): void {
    cy.get(locator).should('contain.text', label);
    cy.get(locator).click();
  }
}
