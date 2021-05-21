export class MfaPage {
  private readonly pinTextFieldsLocator: string = `.pinlogin-field`
  private readonly verifyButtonLocator: string = `#verify-button`
  private readonly textMessageButtonLocator: string = `[onclick^="submit"]`;

  public typePin(pin: string = `111111`): void {
    cy.get(this.textMessageButtonLocator).click();
    cy.get(this.pinTextFieldsLocator).each((pinTextField, index) => {
      cy.wrap(pinTextField).type(pin.charAt(index))
    });
    cy.get(this.verifyButtonLocator).click();
  }

}
