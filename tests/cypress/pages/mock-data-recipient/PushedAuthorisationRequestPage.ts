import { Urls } from '../Urls';

export class PushedAuthorisationRequestPage {
  private readonly mainHeaderSelector: string = `[role='main'] > h2`;
  private readonly initiateAuthorisationHeaderSelector: string = '#par-header';
  private readonly registrationSelector = (node: string) => `#par-registration ${node}`;
  private readonly cdrArrangementSelector = (node: string) => `#par-CdrArrangement ${node}`;
  private readonly sharingDurationSelector = (node: string) => `#par-sharingDuration ${node}`;
  private readonly scopeSelector = (node: string) => `#par-scope ${node}`;
  private readonly initiateParButtonSelector: string = '#par-initiatePar';
  private readonly consentAuthorizationUriSelector: string = '#par-authorisationUri';
  private readonly requestUriLabelSelector: string = '#par-requestUri';
  private readonly requestUriOutputSelector: string = '#par-requestUri + dd';
  private readonly expiresInLabelSelector: string = '#par-expiresIn';
  private readonly expiresInOutputSelector: string = '#par-expiresIn + dd';

  public visit(force: boolean = false): void {
    Urls.visit(Cypress.env(`mock_data_recipient_url`) + '/par', false);
    this.assertThatPageIsDisplayed();
  }

  public assertThatPageIsDisplayed(): void {
    cy.get(this.mainHeaderSelector)
      .should('have.text', 'Pushed Authorisation Request (PAR)');
    cy.get(this.initiateAuthorisationHeaderSelector)
      .should('have.text', 'Initiate PAR with the selected Data Holder');
    cy.get(this.registrationSelector('label'))
      .should('have.text', 'Registration');
    cy.get(this.registrationSelector('#RegistrationId'))
      .should('be.visible');
    cy.get(this.cdrArrangementSelector('label'))
      .should('have.text', 'CDR Arrangement');
    cy.get(this.cdrArrangementSelector('#CdrArrangementId'))
      .should('be.visible');
    cy.get(this.sharingDurationSelector('label'))
      .should('have.text', 'SharingDuration (in Seconds)');
    cy.get(this.sharingDurationSelector('#SharingDuration'))
      .should('be.visible');
    cy.get(this.scopeSelector('label'))
      .should('have.text', 'Scope');
    cy.get(this.scopeSelector('#Scope'))
      .should('be.visible');
    cy.get(this.initiateParButtonSelector)
      .should('be.visible');
  }

  public selectClientRegistration(index: number = 0): void {
    cy.get(this.registrationSelector('#RegistrationId > option'))
      .eq(index)
      .then((element) => cy.get('#RegistrationId').select(element.val().toString()));
  }

  public setSharingDuration(duration: number): void {
    cy.get(this.sharingDurationSelector('#SharingDuration'))
    .type(duration.toString());
  }

  public setScopes(scopes: string[]): void {
    cy.get(this.scopeSelector('#Scope'))
    .type(scopes.join(' ').toString());
  }

  public clickInitiateParButton(): void {
    cy.get(this.initiateParButtonSelector).click();
  }

  public assertThatAuthorizationUriIsGenerated(): void {
    cy.get(this.consentAuthorizationUriSelector, {timeout: 30000,})
      .should('be.visible');
    cy.get(this.requestUriLabelSelector, {timeout: 30000,})
      .should('have.text', 'request_uri');
    cy.get(this.requestUriOutputSelector, {timeout: 30000,})
      .should('contain.text', 'urn:ietf:params:oauth:request_uri:');
    cy.get(this.expiresInLabelSelector, {timeout: 30000,})
      .should('have.text', 'expires_in');
    cy.get(this.expiresInOutputSelector, {timeout: 30000,})
      .should('have.text', '60');

    cy.get(this.registrationSelector(`#RegistrationId > option`)).eq(0)
      .then(($registrationId) => {
        const clientId = $registrationId.val().toString();
        cy.get(this.consentAuthorizationUriSelector)
          .should('contain.text', clientId);
      });

    cy.get(this.scopeSelector(`#Scope`))
      .should('be.visible')
      .then(($scope) => {
        cy.get(this.consentAuthorizationUriSelector)
          .should('contain.text', $scope.val());
      });

    cy.get(this.requestUriOutputSelector)
      .should('be.visible')
      .then(($requestUri) => {
        cy.get(this.consentAuthorizationUriSelector)
          .should('contain.text', $requestUri.text());
      });
  }

  public clickOnAuthorizationUriLink(): void {
    cy.get(this.consentAuthorizationUriSelector).click();
  }
}
