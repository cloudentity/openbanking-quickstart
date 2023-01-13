export class MfaPage {
  private readonly titleLocator: string = `.heading-2`;
  private readonly subtitleLocator: string = `.heading-4`;
  private readonly textMessageButtonLocator: string = `[onclick^="submit"]`;
  private readonly pinTextFieldsLocator: string = `.pinlogin-field`
  private readonly verifyButtonLocator: string = `#verify-button[data-pin='111111']`


  public assertThatMfaIsDisplayed(): void {
    cy.get(this.titleLocator).contains("Additional verification required");
    cy.get(this.subtitleLocator).contains("To proceed, select where to send your security code");
    cy.get(this.textMessageButtonLocator).should("be.visible");
  }

  public typePin(pin: string = `111111`): void {
    cy.get(this.textMessageButtonLocator).click();

    cy.get(this.subtitleLocator).contains("Enter the security code sent to user");

    cy.get(this.pinTextFieldsLocator).each((pinTextField, index) => {
      cy.wrap(pinTextField).type(pin.charAt(index))
    });
    cy.get(this.verifyButtonLocator).click();
  }

}
