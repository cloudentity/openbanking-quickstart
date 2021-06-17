import { Urls } from "../Urls";

export class ConsentAdminPage {
  private readonly thirdPartyProvidersTabSelector: string =
    "#third-party-providers-tab";
  private readonly revokeConfirmDrawerSelector: string = ".MuiDrawer-paper";
  private readonly revokeButtonSelector: string = ".MuiButton-label";
  private readonly clientRevokeAccessAcceptCheckboxSelector: string =
    "#revoke-access-accept-checkbox";
  private readonly revokeConfirmButtonSelector: string =
    "#revoke-access-button";

  public visit(force: boolean = false): void {
    Urls.visit(Urls.consentAdminUrl, force);
  }

  public revokeClientConsent(): void {
    cy.get(this.thirdPartyProvidersTabSelector).click();
    cy.get(this.revokeConfirmDrawerSelector).should("not.exist");
    cy.contains(this.revokeButtonSelector, "Revoke").click();
    cy.get(this.clientRevokeAccessAcceptCheckboxSelector).click();
    cy.get(this.revokeConfirmButtonSelector).click();
    cy.get(this.revokeConfirmDrawerSelector).should("not.exist");
  }
}
