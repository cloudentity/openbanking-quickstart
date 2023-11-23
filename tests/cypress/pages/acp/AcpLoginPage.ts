import { MfaPage } from "../mfa/MfaPage";
import { EnvironmentVariables } from "../EnvironmentVariables";
import { Credentials } from "../Credentials";

export class AcpLoginPage {
  mfaPage: MfaPage = new MfaPage();
  environmentVariables: EnvironmentVariables = new EnvironmentVariables();

  private readonly usernameLocator: string = `#text-field-username-input`;
  private readonly passwordLocator: string = `#text-field-password-input`;
  private readonly loginButtonLocator: string = `button[type='submit']`;

  public login(): void {
    cy.get(this.usernameLocator).type(Credentials.defaultUsername);
    cy.get(this.passwordLocator).type(Credentials.defaultPassword);
    cy.get(this.loginButtonLocator).click();
  }

  public loginWithMfaOption(): void {
    cy.get(this.usernameLocator).type(Credentials.defaultUsername);
    cy.get(this.passwordLocator).type(Credentials.defaultPassword);
    cy.get(this.loginButtonLocator).click();

    if (this.environmentVariables.isMfaEnabled()) {
      this.mfaPage.assertThatMfaIsDisplayed();
      this.mfaPage.typePin();
    }
  }

  public assertThatModalIsDisplayed(): void {
    cy.get(this.usernameLocator).should("be.visible");
    cy.get(this.passwordLocator).should("be.visible");
    cy.get(this.loginButtonLocator).should("be.visible");
  }
}
