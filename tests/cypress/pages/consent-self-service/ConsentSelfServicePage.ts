import {Urls} from '../Urls';
import {Aliases} from "../Aliases";

export class ConsentSelfServicePage {

  private readonly expandIconSelector: string = `.MuiSvgIcon-root`;
  private readonly confirmButtonLocator: string = `#confirm-button`;

  public visit(): void {
    Urls.forceVisit(Urls.consentSelfServiceUrl);
  }

  public expandTab(): void {
    cy.get(this.expandIconSelector).click();
  }

  public assertConsentIsNotDisplayed(): void {
    cy.get(`@${Aliases.intentId}`).then(intentId => cy.get(`#${intentId}`).should(`not.exist`))
  }

  public revokeConsent(): void {
    cy.get(`@${Aliases.intentId}`).then(intentId => cy.get(`#${intentId} .revoke-button`).click())
    cy.get(this.confirmButtonLocator).click()
  }

}
