import {Aliases} from "../Aliases";

export class TppIntentPage {
  private readonly loginButtonLocator: string = `[onclick="onLogin()"]`;
  private readonly intentIdLocator: string = `#intent-id`;

  public login(): void {
    cy.get(this.loginButtonLocator).click();
  }

  public saveIntentId() {
    cy.get(this.intentIdLocator).invoke('text').as(Aliases.intentId)
  }
}
