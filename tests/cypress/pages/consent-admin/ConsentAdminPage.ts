import { Urls } from "../Urls";
import { Aliases } from "../Aliases";

export class ConsentAdminPage {
  private readonly thirdPartyProvidersTabSelector: string =
    "#third-party-providers-tab";
  private readonly revokeButtonSelector: string = ".MuiButton-label";
  private readonly revokeConfirmDialogSelector: string =
    "#revoke-confirm-dialog";
  private readonly revokeConfirmButtonSelector: string =
    "#revoke-confirm-button";

  public visit(force: boolean = false): void {
    Urls.visit(Urls.consentAdminUrl, force);
  }

  public revokeClientConsent(): void {
    cy.get(this.thirdPartyProvidersTabSelector).click();
    cy.contains(this.revokeButtonSelector, "Revoke").click();
    cy.get(this.revokeConfirmDialogSelector).should("exist");
    cy.get(this.revokeConfirmButtonSelector).click();
    cy.get(this.revokeConfirmDialogSelector).should("not.exist");
  }
}
