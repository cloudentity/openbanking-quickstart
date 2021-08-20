import {Urls} from '../Urls';

export class ConsentSelfServicePage {

  private readonly applicationCardLocator: string = `.application-card`;

  public visit(force: boolean = false): void {
    Urls.visit(Cypress.env('CONSENT_SELF_SERVICE_URL'), force);
  }

  public clickOnApplicationCard(): void {
    cy.get(this.applicationCardLocator).click();
  }

}
