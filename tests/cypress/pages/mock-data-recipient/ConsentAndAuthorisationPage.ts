import { Urls } from '../Urls';

export class ConsentAndAuthorisationPage {
  private readonly mainHeaderSelector: string = `[role='main'] > h2`;
  private readonly initiateAuthorisationHeaderSelector: string = '#consent-header';
  private readonly initiateAuthorisationTitleSelector: string = '#consent-title';
  private readonly registrationSelector = (node: string) => `#consent-registration ${node}`;
  private readonly sharingDurationSelector = (node: string) => `#consent-sharingDuration ${node}`;
  private readonly scopeSelector = (node: string) => `#consent-scope ${node}`;
  private readonly constructAuthorizationUriButtonSelector: string = '#consent-btnConstructAuthorizationUri';
  private readonly consentAuthorizationUriSelector: string = '#consent-authorisationUri';

  public visit(force: boolean = false): void {
    Urls.visit(Cypress.env(`mock_data_recipient_url`) + '/consent', false);
    Urls.clearLocalStorage();
    this.assertThatPageIsDisplayed();
  }

  public assertThatPageIsDisplayed(): void {
    cy.get(this.mainHeaderSelector)
      .should('have.text', 'Consent and Authorisation');
    cy.get(this.initiateAuthorisationHeaderSelector)
      .should('have.text', 'Initiate Authorisation');
    cy.get(this.initiateAuthorisationTitleSelector)
      .should('have.text', 'Initiate the authorisation flow with the selected Data Holder');
    cy.get(this.registrationSelector('label'))
      .should('have.text', 'Registration');
    cy.get(this.registrationSelector('#ClientId'))
      .should('be.visible');
    cy.get(this.sharingDurationSelector('label'))
      .should('have.text', 'Sharing Duration');
    cy.get(this.sharingDurationSelector('#SharingDuration'))
      .should('be.visible');
    cy.get(this.scopeSelector('label'))
      .should('have.text', 'Scope');
    cy.get(this.scopeSelector('#Scope'))
      .should('be.visible');
    cy.get(this.constructAuthorizationUriButtonSelector)
      .should('be.visible');
  }

  public selectClientRegistration(index: number = 0): void {
    cy.get(this.registrationSelector('#ClientId > option'))
      .eq(index)
      .then((element) => cy.get('#ClientId').select(element.val().toString()));
  }

  public setSharingDuration(duration: number): void {
    cy.get(this.sharingDurationSelector('#SharingDuration'))
    .type(duration.toString());
  }

  public clickConstructAuthorizationUriButton(): void {
    cy.get(this.constructAuthorizationUriButtonSelector).click();
  }

  public assertThatAuthorizationUriIsGenerated(): void {
    cy.get(this.consentAuthorizationUriSelector, {timeout: 30000,})
      .should('be.visible');
    
    cy.get(this.registrationSelector(`[selected='selected']`))
      .should('be.visible')
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
  }

  public clickOnAuthorizationUriLink(): void {
    cy.get(this.consentAuthorizationUriSelector).click();
  }
}
