import {Urls} from '../Urls';
import {Aliases} from "../Aliases";

export class ConsentAdminPage {

  private readonly expandIconSelector: string = `.MuiSvgIcon-root`;

  public visit(): void {
    Urls.forceVisit(Urls.consentAdminUrl);
  }

  public expandTab(): void {
    cy.get(this.expandIconSelector).click();
  }

  public assertConsentIsNotDisplayed(): void {
    cy.get(`@${Aliases.intentId}`).then(intentId => cy.get(`#${intentId}`).should(`not.exist`))
  }

  public revokeConsent(): void {
    cy.get(`@${Aliases.intentId}`).then(intentId => cy.get(`#${intentId} .revoke-button`).click())
  }

  public revokeAllConsents(): void {
    cy.get(`.revoke-all-button`).click()
  }

}
