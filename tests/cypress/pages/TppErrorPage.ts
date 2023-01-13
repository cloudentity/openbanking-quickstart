export class TppErrorPage {
  private readonly errorTitleLocator: string = `#tpp-consent-error`;
  private readonly errorDescriptionLocator: string = `#tpp-consent-error-description`;
  private readonly errorCauseLocator: string = `#tpp-consent-error-cause`;
  private readonly traceIdLocator: string = `#tpp-consent-trace-id`;
  private readonly tryNextButtonLocator: string = `#tpp-consent-error-try-again-button`;


  public assertThatRejectConsentErrorPageIsDisplayed(error: string, description: string, cause: string): void {
    this.assertThatCancelLoginErrorPageIsDisplayed(error, description);
    cy.get(this.errorCauseLocator).contains("cause: " + cause);
  }

  public assertThatCancelLoginErrorPageIsDisplayed(error: string, description: string): void {
    cy.get(this.tryNextButtonLocator, { timeout: 3000 }).should("be.visible");
    cy.get(this.errorTitleLocator).contains(error);
    cy.get(this.errorDescriptionLocator).contains("description: " + description);
    cy.get(this.traceIdLocator).then((element) => {
      let text: string = element.text();
      expect(text).to.match(/ID\: [a-zA-Z0-9]+/);
    });
  }

  public clickTryNext(): void {
    cy.get(this.tryNextButtonLocator).click();
  }
}
