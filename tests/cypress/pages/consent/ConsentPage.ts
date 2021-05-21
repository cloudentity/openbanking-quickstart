export class ConsentPage {
  private readonly confirmButtonLocator: string = `[value="confirm"]`;
  private readonly cancelButtonLocator: string = `[value="deny"]`;
  private readonly permissionRowLocator: string = `[data-desc-id="account_permissions"] li`;
  private readonly expandPermissionsButtonLocator: string = `[data-icon-id="account_permissions"]`
  private readonly accountIdsLocator: string = `[name="account_ids"]`;

  public checkAccounts(accounts: string[]): void {
    cy.get(this.accountIdsLocator).uncheck()
    accounts.forEach(account => cy.get(`[id*="${account}"]`).check())
  }

  public expandPermissions(): void {
    cy.get(this.expandPermissionsButtonLocator).click();
  }

  public assertPermissions(length: number): void {
    cy.get(this.permissionRowLocator).should('have.length', length)
  }

  public confirm(): void {
    cy.get(this.confirmButtonLocator).click();
  }

  public cancel(): void {
    cy.get(this.cancelButtonLocator).click();
  }
}
