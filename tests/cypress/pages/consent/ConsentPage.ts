export class ConsentPage {
  private readonly confirmButtonLocator: string = `[value="confirm"]`;
  private readonly cancelButtonLocator: string = `[value="deny"]`;
  private readonly permissionNameLocator: string = `[data-desc-id="account_permissions"] .permission-name`;
  private readonly permissionContentLocator: string = `[data-desc-id="account_permissions"] .caption`;
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

  public assertPermissionsDetails(name: string, content: string): void {
      cy.get(this.permissionNameLocator).should('contain.text', name)
      cy.get(this.permissionContentLocator).should('contain.text', content)
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

  public assertThatPageIsNotVisible(): void {
    cy.get(this.cancelButtonLocator, { timeout: 5000 } ).should('not.exist');
    cy.get(this.confirmButtonLocator, { timeout: 5000 } ).should('not.exist');
  }
}
