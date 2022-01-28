import {Urls} from '../Urls';

export class ConsentSelfServicePage {

  private readonly applicationCardLocator: string = `.application-card`;

  public visit(force: boolean = false): void {
    Urls.visit(Cypress.env('consent_self_service_url'), force);
  }

  public clickOnApplicationCard(): void {
    cy.get(this.applicationCardLocator, { timeout: 30000 }).should('be.visible').click();
  }

}
