import { Urls } from '../Urls';

export class MockDataRecipientNavigationPage {
  private readonly mainHeaderSelector: string = '#main-header';
  private readonly navHomeSelector: string = `#nav-sidebar [href='/']`;
  private readonly navDiscoverDataHoldersSelector: string = `#nav-sidebar [href='/data-holders']`;
  private readonly navGetSsaSelector: string = `#nav-sidebar [href='/ssa']`;
  private readonly navDynamicClientRegistrationSelector: string = `#nav-sidebar [href='/dcr']`;
  private readonly navConsentAndAuthorisationSelector: string = `#nav-sidebar [href='/consent']`;
  private readonly navConsentsSelector: string = `#nav-sidebar [href='/consent/consents']`;
  private readonly navParSelector: string = `#nav-sidebar [href='/par']`;
  private readonly navIdTokenHelperSelector: string = `#nav-sidebar [href='/utilities/id-token']`;
  private readonly navPrivateKeyJwtGeneratorSelector: string = `#nav-sidebar [href='/utilities/private-key-jwt']`;

  public visit(force: boolean = false): void {
    Urls.visit(Cypress.env('mock_data_recipient_url'), force);
    Urls.clearLocalStorage();
    Urls.visit(Cypress.env('mock_data_recipient_url'), force);
    this.assertThatNavigationPageIsDisplayed();
  }

  public assertThatNavigationPageIsDisplayed(): void {
    cy.get(this.mainHeaderSelector)
      .should('have.text', 'Mock Data Recipient');
    cy.get(this.navHomeSelector)
      .should('contain.text', 'Home');
    cy.get(this.navDiscoverDataHoldersSelector)
      .should('contain.text', 'Discover Data Holders');
    cy.get(this.navGetSsaSelector)
      .should('contain.text', 'Get SSA');
    cy.get(this.navDynamicClientRegistrationSelector)
      .should('contain.text', 'Dynamic Client Registration');
    cy.get(this.navConsentAndAuthorisationSelector)
      .should('contain.text', 'Consent and Authorisation');
    cy.get(this.navConsentsSelector)
      .should('contain.text', 'Consents');
    cy.get(this.navParSelector)
      .should('contain.text', 'PAR');
    cy.get(this.navIdTokenHelperSelector)
      .should('contain.text', 'ID Token Helper');
    cy.get(this.navPrivateKeyJwtGeneratorSelector)
      .should('contain.text', 'Private Key JWT Generator');
  }

  public clickHomeLink(): void {
    cy.get(this.navHomeSelector).click();
  }

  public clickDiscoverDataHoldersLink(): void {
    cy.get(this.navDiscoverDataHoldersSelector).click();
  }

  public clickGetSsaLink(): void {
    cy.get(this.navGetSsaSelector).click();
  }

  public clickDynamicClientRegistrationLink(): void {
    cy.get(this.navDynamicClientRegistrationSelector).click();
  }

  public clickConsentAndAuthorisationLink(): void {
    cy.get(this.navConsentAndAuthorisationSelector).click();
  }

  public clickConsentsLink(): void {
    cy.get(this.navConsentsSelector).click();
  }

  public clickParLink(): void {
    cy.get(this.navParSelector).click();
  }

  public clickIdTokenHelperLink(): void {
    cy.get(this.navIdTokenHelperSelector).click();
  }

  public clicknavPrivateKeyJwtGeneratorLink(): void {
    cy.get(this.navPrivateKeyJwtGeneratorSelector).click();
  }
}
