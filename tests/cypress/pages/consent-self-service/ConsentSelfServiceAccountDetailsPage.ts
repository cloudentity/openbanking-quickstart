export class ConsentSelfServiceAccountDetailsPage {
  private readonly permissionDatesLocator: string = `#account-permission-dates`;
  private readonly infoLocator: string = `#accounts-info`;
  private readonly detailsLocator: string = `#account-details`;
  private readonly revokeButtonLocator: string = `#revoke-access-button`;
  private readonly cancelRevokeButtonLocator: string = `#cancel-revoke-access-button`;
  private readonly revokeInfoLocator: string = `#account-revoke-info`;
  private readonly revokeCheckboxLocator: string = `#account-revoke-checkbox`;

  public assertThatAccountDetailsAreVisible(): void {
    cy.get(this.permissionDatesLocator).should(`contain.text`, `Permission dates`);
    cy.get(this.infoLocator).should(`contain.text`, `Accounts`);
    cy.get(this.detailsLocator).should(`contain.text`, `Details being shared`);
    cy.get(this.revokeButtonLocator).should(`contain.text`, `Revoke access`);
    cy.get(this.cancelRevokeButtonLocator).should(`contain.text`, `Cancel`);
  }

  public assertAccount(accountId: string): void {
    cy.get(this.infoLocator).should(`contain.text`, accountId);
  }

  public clickRevokeAccessButton(): void {
    cy.get(this.revokeButtonLocator).click();
  }

  public assertThatRevokeAccountDetailsAreVisible(): void {
    cy.get(this.revokeInfoLocator).should(
      `contain.text`,
      `Warning: Deleteing this consent will remove access to all accounts to which you have previously granted access.`
    );
    cy.get(this.revokeInfoLocator).should(
      `contain.text`,
      `Are you sure you want to revoke access for all accounts connected with this application?`
    );
  }

  public confirmRevokeAccessAction(): void {
    cy.get(this.revokeCheckboxLocator).check();
    cy.get(this.revokeButtonLocator).click();
  }
}
