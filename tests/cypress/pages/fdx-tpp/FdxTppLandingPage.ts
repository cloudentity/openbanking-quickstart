import { Urls } from "../Urls";

export class FdxTppLandingPage {
  private readonly loginPageInfoLabelLocator: string = `#fdx-login-info-text p`;
  private readonly authorizationDetailsContentLocator: string = `#fdx-login-authorization-details`;
  private readonly nextButtonLocator: string = `#fdx-login-next-button`;

  private expectedAuthorizationDetails: string =
    '{"authorization_details":[{"type":"fdx_v1.0","consentRequest":' +
    '{"durationType":"ONE_TIME","lookbackPeriod":60,"resources":' +
    '[{"resourceType":"ACCOUNT","dataClusters":["ACCOUNT_DETAILED","TRANSACTIONS","STATEMENTS"]},' +
    '{"resourceType":"CUSTOMER","dataClusters":["CUSTOMER_CONTACT"]}]}}]}';

  public visit(force: boolean = false): void {
    Urls.visit(Cypress.env("tpp_url"), force);
    Urls.clearLocalStorage();
    Urls.visit(Cypress.env("tpp_url"), force);
  }

  public assertThatPageIsDisplayed(): void {
    cy.get(this.loginPageInfoLabelLocator).should(
      'have.text',
      'In the FDX, the flow starts with the TPP sending authorization_details to the PAR endpoint.'
    );
    cy.get(this.authorizationDetailsContentLocator).then((element) => {
      let text: string = element.text().replace(/\s/g, "");
      expect(text).to.contain(this.expectedAuthorizationDetails);
    });
    cy.get(this.nextButtonLocator, { timeout: 3000 }).should("be.visible");
  }

  public clickNext(): void {
    cy.get(this.nextButtonLocator).click();
  }
}
