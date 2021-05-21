import {Urls} from '../Urls';

export class ConsentSelfServicePage {

  private readonly applicationCardLocator: string = `.application-card`;

  public visit(force: boolean = false): void {
    Urls.visit(Urls.consentSelfServiceUrl, force);
  }

  public clickOnApplicationCard(): void {
    cy.get(this.applicationCardLocator).click();
  }

}
