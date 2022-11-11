export class FdxTppErrorPage {
  private readonly errorTitleLocator: string = `#fdx-consent-error`;
  private readonly errorDescriptionLocator: string = `#fdx-consent-error-description`;
  private readonly errorCauseLocator: string = `#fdx-consent-error-cause`;
  private readonly traceIdLocator: string = `#fdx-consent-trace-id`;
  private readonly tryNextButtonLocator: string = `#fdx-consent-error-try-again-button`;


  public assertThatErrorPageIsDisplayed(error: string, description: string, cause: string): void {
    cy.get(this.tryNextButtonLocator, { timeout: 3000 }).should("be.visible");
    cy.get(this.errorTitleLocator).contains(error);
    cy.get(this.errorDescriptionLocator).contains("description: " + description);
    cy.get(this.errorCauseLocator).contains("cause: " + cause);
    cy.get(this.traceIdLocator).then((element) => {
      let text: string = element.text();
      expect(text).to.match(/ID\: [a-zA-Z0-9]+/);
    });
  }

  public clickTryNext(): void {
    cy.get(this.tryNextButtonLocator).click();
  }
}
