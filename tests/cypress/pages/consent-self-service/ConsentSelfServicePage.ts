import {Urls} from '../Urls';

export class ConsentSelfServicePage {

  private readonly applicationCardLocator: string = `.application-card`;
  private readonly allTypesLocator: string = `#all-types`;
  private readonly accountOnlyLocator: string = `#account-only`;
  private readonly paymentOnlyLocator: string = `#payment-only`;
  private readonly accountOnlySelectedLocator: string = `[id='account-only'][style*='color: white']`

  public visit(force: boolean = false): void {
    Urls.visit(Cypress.env('consent_self_service_url'), force);
  }

  public clickOnAccountOnlyButton(): void {
    cy.get(this.accountOnlyLocator).should('be.visible').click();
    cy.get(this.accountOnlySelectedLocator).should('be.visible');
  }

  public clickOnApplicationCard(): void {
    cy.get(this.applicationCardLocator, { timeout: 30000 }).should('be.visible').click();
  }

  public assertThatFilterPermissionsButtonsAreDisplayed(): void {
    cy.get(this.allTypesLocator, { timeout: 5000 }).should('be.visible');
    cy.get(this.accountOnlyLocator, { timeout: 5000 }).should('be.visible');
    cy.get(this.paymentOnlyLocator, { timeout: 5000 }).should('be.visible');
  }

  public assertThatApplicationCardIsNotDisplayed(): void {
    cy.get(this.applicationCardLocator, { timeout: 30000 }).should('not.exist');
  }
}
