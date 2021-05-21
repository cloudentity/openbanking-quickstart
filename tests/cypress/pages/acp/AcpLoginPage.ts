export class AcpLoginPage {
  private readonly usernameLocator: string = `#text-field-username-input`;
  private readonly passwordLocator: string = `#text-field-password-input`;
  private readonly loginButtonLocator: string = `.mdc-button__ripple`;
  private readonly cancelButtonLocator: string = `#cancel`;

  public login(username: string, password: string): void {
    cy.get(this.usernameLocator).type(username)
    cy.get(this.passwordLocator).type(password)
    cy.get(this.loginButtonLocator).click();
  }

  public cancel(): void {
    cy.get(this.cancelButtonLocator).click();
  }
}
