import {Urls} from '../Urls'

export class FinancrooLoginPage {
  private readonly loginButtonLocator: string = `.login-button`;

  public visit(force: boolean = false): void {
    Urls.visit(Cypress.env('financroo_url'), force);
    Urls.clearLocalStorage();
    Urls.visit(Cypress.env("financroo_url"), force);
  }

  public login(): void {
    cy.get(this.loginButtonLocator).click();
  }
}
