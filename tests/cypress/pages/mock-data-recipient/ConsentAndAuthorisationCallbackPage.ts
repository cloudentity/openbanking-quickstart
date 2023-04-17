import { Urls } from "../Urls";

export class ConsentAndAuthorisationCallbackPage {
  private readonly mainHeaderSelector: string = `[role='main'] > h2`;
  private readonly idTokenSelector: string = '#consent-callback-idToken + dd';
  private readonly accessTokenSelector: string = '#consent-callback-accessToken + dd';
  private readonly refreshTokenSelector: string = '#consent-callback-refreshToken + dd';
  private readonly expiresInSelector: string = '#consent-callback-expiresIn + dd';
  private readonly tokenTypeSelector: string = '#consent-callback-tokenType + dd';
  private readonly cdrArrangementIdSelector: string = '#consent-callback-cdrArrangementId + dd';

  private readonly scopes: Array<string> = [
    "openid",
    "profile",
    "bank:accounts.basic:read",
    "bank:accounts.detail:read",
    "bank:transactions:read",
    "common:customer.basic:read",
    "offline_access",
  ];


  public visit(force: boolean = false): void {
    Urls.visit(Cypress.env(`mock_data_recipient_url`) + '/consent/callback', false);
    Urls.clearLocalStorage();
    this.assertThatPageIsDisplayed();
  }

  public assertThatPageIsDisplayed(): void {
    cy.get(this.mainHeaderSelector)
      .should('have.text', 'Consent and Authorisation - Callback');

    cy.get(this.idTokenSelector).then((element) => {
      expect(element.text()).to.match(/[a-zA-Z0-9_-]+\.[a-zA-Z0-9_-]+\.[a-zA-Z0-9_-]+\.[a-zA-Z0-9_-]+\.[a-zA-Z0-9_-]+/);
      // expect(text).to.match(/([a-zA-Z0-9_-]+)\.(?1)\.(?1)\.(?1)\.(?1)/);
    });

    cy.get(this.accessTokenSelector).then((element) => {
      let text: string = element.text();
      expect(element.text()).to.match(/[a-zA-Z0-9_-]+\.[a-zA-Z0-9_-]+\.[a-zA-Z0-9_-]+/);
      // expect(element.text()).to.match(/([a-zA-Z0-9_-]+)\.(?1)\.(?1)/);
    });

    cy.get(this.refreshTokenSelector).then((element) => {
      let text: string = element.text();
      expect(element.text()).to.match(/[a-zA-Z0-9_-]+\.[a-zA-Z0-9_-]+/);
      // expect(element.text()).to.match(/([a-zA-Z0-9_-]+)\.(?1)/);
    });

    cy.get(this.expiresInSelector).then((element) => {
      let text: string = element.text();
      expect(element.text()).to.match(/\d{3}/);
    });

    this.scopes.forEach(element => {
      cy.get('#consent-callback-scope + dd').should("contain.text", element);
    });

    cy.get(this.tokenTypeSelector).should("have.text", "bearer");

    cy.get(this.cdrArrangementIdSelector).then((element) => {
      expect(element.text()).to.match(/[a-zA-Z0-9]{20}/);
    });
  }
}
