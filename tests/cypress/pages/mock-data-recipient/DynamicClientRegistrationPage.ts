import { Urls } from '../Urls';

export class DynamicClientRegistrationPage {
  private readonly mainHeaderSelector: string = `[role='main'] > h2`;
  private readonly createClientRegistrationHeaderSelector: string = '#dcr-create-header';
  private readonly currentRegistrationsHeaderSelector: string = '#dcr-current-header';
  private readonly brandIdSelector = (node: string) => `#dcr-create-dataHolderBrandId ${node}`;
  private readonly clientIdSelector = (node: string) => `#dcr-create-clientId ${node}`;
  private readonly softwareProductIdSelector = (node: string) => `#dcr-create-softwareProductId ${node}`;
  private readonly redirectUrisSelector = (node: string) => `#dcr-create-redirectUris ${node}`;
  private readonly tokenEndpointAuthSigningAlgSelector = (node: string) => `#dcr-create-tokenEndpointAuthSigningAlg ${node}`;
  private readonly tokenEndpointAuthMethodSelector = (node: string) => `#dcr-create-tokenEndpointAuthMethod ${node}`;
  private readonly grantTypesSelector = (node: string) => `#dcr-create-grantTypes ${node}`;
  private readonly responseTypesSelector = (node: string) => `#dcr-create-responseTypes ${node}`;
  private readonly applicationTypeSelector = (node: string) => `#dcr-create-applicationType ${node}`;
  private readonly idTokenSignedResponseAlgSelector = (node: string) => `#dcr-create-idTokenSignedResponseAlg ${node}`;
  private readonly idTokenEncryptedResponseAlgSelector = (node: string) => `#dcr-create-idTokenEncryptedResponseAlg ${node}`;
  private readonly idTokenEncryptedResponseEncSelector = (node: string) => `#dcr-create-idTokenEncryptedResponseEnc ${node}`;
  private readonly requestObjectSigningAlgSelector = (node: string) => `#dcr-create-requestObjectSigningAlg ${node}`;
  private readonly registerButtonSelector: string = '#dcr-create-btnRegister';
  private readonly cancelButtonSelector: string = '#dcr-create-btnCancel';
  private readonly registerMessageSelector: string = '#dcr-create-message';
  private readonly noRegistrationMessage: string = '#dcr-current-noDataMessage';
  private readonly registrationsCountMessageSelector = (node: string) => `#dcr-current-count ${node}`;
  private readonly deleteClientButtonSelector: string = `[id*='dcr-current-btnDelete']`;
  private readonly dhBrandNameFormSelector: string = `#DataHolderBrandId`;


  public visit(force: boolean = false): void {
    Urls.visit(Cypress.env(`mock_data_recipient_url`) + '/dcr', force);
    Urls.clearLocalStorage();
    this.assertThatPageIsDisplayed();
  }

  public assertThatPageIsDisplayed(): void {
    cy.get(this.mainHeaderSelector)
      .should('have.text', 'Dynamic Client Registration');
    cy.get(this.createClientRegistrationHeaderSelector)
      .should('have.text', 'Create Client Registration');
    cy.get(this.currentRegistrationsHeaderSelector)
      .should('have.text', 'Current Registrations');
    cy.get(this.brandIdSelector('#DataHolderBrandId'))
      .should('be.visible');
    cy.get(this.registerMessageSelector)
      .should('contain.text', 'Waiting...');
  }

  public assertThatDynamicClientRegistrationFormIsDisplayed() :void {
    cy.get(this.mainHeaderSelector)
      .should('have.text', 'Dynamic Client Registration');
    cy.get(this.createClientRegistrationHeaderSelector)
      .should('have.text', 'Create Client Registration');
    cy.get(this.currentRegistrationsHeaderSelector)
      .should('have.text', 'Current Registrations');
    cy.get(this.brandIdSelector('#DataHolderBrandId'))
      .should('be.visible');
    cy.get(this.clientIdSelector('label'))
      .should('have.text', 'Client ID');
    cy.get(this.clientIdSelector('#ClientId'))
      .should('be.visible');
    cy.get(this.softwareProductIdSelector('label'))
      .should('have.text', 'Software Product ID');
    cy.get(this.softwareProductIdSelector('#SoftwareProductId'))
      .should('be.visible');
    cy.get(this.redirectUrisSelector('label'))
      .should('have.text', 'Redirect URIs');
    cy.get(this.redirectUrisSelector('#RedirectUris'))
      .should('be.visible');
    cy.get(this.tokenEndpointAuthSigningAlgSelector('label'))
      .should('have.text', 'Token Endpoint Auth Signing Alg');
    cy.get(this.tokenEndpointAuthSigningAlgSelector('#TokenEndpointAuthSigningAlg'))
      .should('be.visible');
    cy.get(this.tokenEndpointAuthMethodSelector('label'))
      .should('have.text','Token Endpoint Auth Method');
    cy.get(this.tokenEndpointAuthMethodSelector('#TokenEndpointAuthMethod'))
      .should('be.visible');
    cy.get(this.grantTypesSelector('label'))
      .should('have.text', 'Grant Types');
    cy.get(this.grantTypesSelector('#GrantTypes'))
      .should('be.visible');
    cy.get(this.responseTypesSelector('label'))
      .should('have.text', 'Response Types');
    cy.get(this.responseTypesSelector('#ResponseTypes'))
      .should('be.visible');
    cy.get(this.applicationTypeSelector('label'))
      .should('have.text','Application Type');
    cy.get(this.applicationTypeSelector('#ApplicationType'))
      .should('be.visible');
    cy.get(this.idTokenSignedResponseAlgSelector('label'))
      .should('have.text', 'Id Token Signed Response Alg');
    cy.get(this.idTokenSignedResponseAlgSelector('#IdTokenSignedResponseAlg'))
      .should('be.visible');
    cy.get(this.idTokenEncryptedResponseAlgSelector('label'))
      .should('have.text', 'Id Token Encrypted Response Alg');
    cy.get(this.idTokenEncryptedResponseAlgSelector('#IdTokenEncryptedResponseAlg'))
      .should('be.visible');
    cy.get(this.idTokenEncryptedResponseEncSelector('label'))
      .should('have.text', 'Id Token Encrypted Response Enc');
    cy.get(this.idTokenEncryptedResponseEncSelector('#IdTokenEncryptedResponseEnc'))
      .should('be.visible');
    cy.get(this.requestObjectSigningAlgSelector('label'))
      .should('have.text', 'Request Object Signing Alg');
    cy.get(this.requestObjectSigningAlgSelector('#RequestObjectSigningAlg'))
      .should('be.visible');
    cy.get(this.registerButtonSelector)
      .should('be.visible');
    cy.get(this.cancelButtonSelector)
      .should('be.visible');
    cy.get(this.registerMessageSelector)
      .should('contain.text', 'Waiting...');
  }

  public clickMockDataHolderBankingBrand(): void {
    cy.get(this.dhBrandNameFormSelector).select("804fc2fb-18a7-4235-9a49-2af393d18bc7");
  }

  public assertThatBrandIdIsSelected(): void {
    cy.get(this.brandIdSelector(`option[selected='selected']`))
      .should('not.have.text', 'Select Data Holder Brand...');
  }

  public clickDCRRegisterButton(): void {
    cy.get(this.registerButtonSelector).click();
  }

  public clickDeleteClientButton(): void {
    cy.get(this.deleteClientButtonSelector).click();
  }

  public assertThatClientRegistered(): void {
    cy.get(this.registerMessageSelector, {timeout: 30000,})
      .should('contain.text', 'Created - Registered');
    cy.get(this.noRegistrationMessage)
      .should('not.exist');
    cy.get(this.registrationsCountMessageSelector('strong'))
      .should('have.text', '1');
  }

  public assertThatRegisteredClientWasRemoved(): void {
    cy.get(this.noRegistrationMessage, {timeout: 30000,})
      .should('have.text', 'No existing registrations found.');
    cy.get(this.registrationsCountMessageSelector('strong'))
      .should('have.text', '0');
  }
}
