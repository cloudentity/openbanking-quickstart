export class FdxTppIntentRegisteredPage {
  private readonly intentRegisteredPageInfoLabelLocator: string = `#fdx-int-reg-info-text p:nth-child(2)`;
  private readonly uriJsonLocator: string = `#fdx-int-reg-request-uri`;
  private readonly fullAuthorizeRequestLocator: string = `#fdx-int-reg-full-authorize-request`;
  private readonly loginButtonLocator: string = `#fdx-int-reg-login-button`;

  public assertThatPageIsDisplayed(): void {
    cy.get(this.intentRegisteredPageInfoLabelLocator).should(
      'have.text',
      'PAR request has been sent. Below you can find the response from POST /par endpoint.'
    );
    cy.get(this.loginButtonLocator, { timeout: 3000 }).should('be.visible');
    this.assertThatRequestUriFieldsAreNotEmpty();
  }

  private assertThatRequestUriFieldsAreNotEmpty(): void {
    cy.get(this.uriJsonLocator).then((element) => {
      let text: string = element.text();
      cy.log(text);
      expect(text).to.match(
        /\"request_uri\"\: \"urn\:ietf:params:oauth:request_uri\:[a-zA-Z0-9]+\"/
      );
    });
    cy.get(this.fullAuthorizeRequestLocator).then((element) => {
      let text: string = element.text();
      cy.log(text);
      expect(text).to.match(
        /request_uri\=urn\%3Aietf\%3Aparams\%3Aoauth\%3Arequest_uri\%3A[a-zA-Z0-9]+/
      );
    });
  }

  public clickLogin(): void {
    cy.get(this.loginButtonLocator).click();
  }
}
