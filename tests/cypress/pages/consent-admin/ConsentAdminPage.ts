import { Urls } from "../Urls";

export class ConsentAdminPage {
  private readonly loginButtonSelector: string = "#login-button";
  private readonly thirdPartyProvidersTabSelector: string =
    "#third-party-providers-tab";
  private readonly consentManagementTabSelector: string =
    "#consent-management-tab";
  private readonly searchLabelSelector: string = "#search-content";
  private readonly searchByAccountNumberSelector: string =
    "#outlined-adornment-password";
  private readonly searchButtonSelector: string = "#search-account";
  private readonly searchResultsSelector: string = "#search-results";
  private readonly consentsDetailsSelector: string = "#consents-tabs";
  private readonly revokeConfirmDrawerSelector: string = ".MuiDrawer-paper";
  private readonly revokeButtonSelector: string = "#revoke-access";
  private readonly clientRevokeAccessAcceptCheckboxSelector: string =
    "#revoke-access-accept-checkbox";
  private readonly revokeConfirmButtonSelector: string =
    "#revoke-access-button";
  private getClientSelector = (name: string) => `[id='client-${name}']`;
  private getClientStatusSelector = (name: string) =>
    `[id='client-${name}'] [id*='status-']`;
  private getMangeClientButtonSelector = (name: string) =>
    `[id='client-${name}'] #manage-account`;
  private getRevokeAccessButtonSelector = (name: string) =>
    `[id='client-${name}'] #revoke-access`;

  public visit(force: boolean = false): void {
    Urls.visit(Cypress.env("consent_admin_url"), force);
  }

  public login(): void {
    cy.get(this.loginButtonSelector).click();
  }

  public revokeClientConsent(): void {
    cy.get(this.thirdPartyProvidersTabSelector).click();
    cy.get(this.revokeConfirmDrawerSelector).should("not.exist");
    cy.contains(this.revokeButtonSelector, "Revoke").click();
    cy.get(this.clientRevokeAccessAcceptCheckboxSelector).click();
    cy.get(this.revokeConfirmButtonSelector).click();
    cy.get(this.revokeConfirmDrawerSelector).should("not.exist");
    cy.get(this.revokeButtonSelector).should("not.exist");
  }

  public searchAccount(accountNumber: string): void {
    cy.get(this.consentManagementTabSelector).click();
    cy.get(this.searchByAccountNumberSelector).type(accountNumber);
    cy.get(this.searchButtonSelector).click();
  }

  public assertAccountResult(accountNumber: string): void {
    cy.get(this.searchResultsSelector).should("be.visible");
    cy.get(this.searchResultsSelector).contains(accountNumber);
  }

  public assertClientAccountWithStatus(
    accountName: string,
    status: string
  ): void {
    cy.get(this.getClientSelector(accountName.toLowerCase())).should(
      "be.visible"
    );
    cy.get(
      this.getClientStatusSelector(accountName.toLocaleLowerCase())
    ).should("be.visible");
    cy.get(
      this.getClientStatusSelector(accountName.toLocaleLowerCase())
    ).contains(status);
  }

  public manageAccount(accountName: string): void {
    cy.get(this.getMangeClientButtonSelector(accountName.toLowerCase()))
      .should("be.visible")
      .click();
  }

  public assertConsentsDetails(): void {
    cy.get(this.consentsDetailsSelector).should("be.visible");
    cy.get(this.consentsDetailsSelector).contains("Account access");
    cy.get(this.consentsDetailsSelector).contains("Payment access");
  }

  public revokeClientConsentByAccountName(accountName: string): void {
    cy.get(this.revokeConfirmDrawerSelector).should("not.exist");
    cy.get(this.getRevokeAccessButtonSelector(accountName.toLowerCase()))
      .should("be.visible")
      .click();
    cy.get(this.clientRevokeAccessAcceptCheckboxSelector).click();
    cy.get(this.revokeConfirmButtonSelector).click();
    cy.get(this.revokeConfirmDrawerSelector).should("not.exist");
  }

  public assertThatConsentManagementTabIsDisplayed(): void {
    cy.get(this.searchByAccountNumberSelector, { timeout: 30000 }).should(
      "be.visible"
    );
    cy.get(this.searchButtonSelector).should("be.visible");
    cy.get(this.searchLabelSelector).should(
      "have.text",
      "Search and manage consents on behalf of bank members"
    );
  }
}
