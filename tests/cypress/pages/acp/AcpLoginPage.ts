import { MfaPage } from "../../pages/mfa/MfaPage";
import { EnvironmentVariables } from "../../pages/EnvironmentVariables";
import { Credentials } from "../../pages/Credentials";

export class AcpLoginPage {
  mfaPage: MfaPage = new MfaPage();
  environmentVariables: EnvironmentVariables = new EnvironmentVariables();

  private readonly signInHeaderLocator: string = `#sign-in h3`;
  private readonly usernameLocator: string = `#text-field-username-input`;
  private readonly passwordLocator: string = `#text-field-password-input`;
  private readonly loginButtonLocator: string = `button[type='submit']`;
  private readonly cancelButtonLocator: string = `#cancel`;

  public login(username: string, password: string): void {
    cy.get(this.usernameLocator).type(username);
    cy.get(this.passwordLocator).type(password);
    cy.get(this.loginButtonLocator).click();
  }

  public confirmLogin(): void {
    cy.get(this.usernameLocator).type(Credentials.defaultUsername);
    cy.get(this.passwordLocator).type(Credentials.defaultPassword);
    cy.get(this.loginButtonLocator).click();

    if (this.environmentVariables.isMfaEnabled()) {
      this.mfaPage.assertThatMfaIsDisplayed();
      this.mfaPage.typePin();
    }
  }

  public cancelLogin(): void {
    cy.get(this.cancelButtonLocator).click();
  }
  
  public assertThatModalIsDisplayed(headerName: string): void {
    cy.get(this.signInHeaderLocator).should("have.text", headerName);
    cy.get(this.usernameLocator).should("be.visible");
    cy.get(this.passwordLocator).should("be.visible");
    cy.get(this.loginButtonLocator).should("be.visible");
    cy.get(this.cancelButtonLocator).should("be.visible");
  }
}
