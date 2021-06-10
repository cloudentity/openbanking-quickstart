import { Urls } from "../Urls";
import { Aliases } from "../Aliases";

export class ConsentAdminPage {
  private readonly clientAvatarSelector: string = ".MuiAvatar-root";
  private readonly clientDrawerSelector: string = ".MuiDrawer-paper";
  private readonly clientRevokeAccessButtonSelector: string =
    "#revoke-access-button";
  private readonly clientRevokeAccessAcceptCheckboxSelector: string =
    "#revoke-access-accept-checkbox";

  public visit(force: boolean = false): void {
    Urls.visit(Urls.consentAdminUrl, force);
  }

  public openClientDrawer(): void {
    cy.get(this.clientAvatarSelector).click();
  }

  public assertAppliationDrawerIsNotDisplayed(): void {
    cy.get(this.clientDrawerSelector).should("not.exist");
  }

  public revokeClientConsent(): void {
    cy.get(this.clientRevokeAccessButtonSelector).click();
    cy.get(this.clientRevokeAccessAcceptCheckboxSelector).click();
    cy.get(this.clientRevokeAccessButtonSelector).click();
  }
}
