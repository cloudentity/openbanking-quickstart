import { Urls } from "../Urls";

export class ConsentSelfServicePage {
  private readonly applicationCardLocator: string = `.application-card`;
  private readonly allTypesLocator: string = `#all-types`;
  private readonly accountOnlyLocator: string = `#account-only`;
  private readonly paymentOnlyLocator: string = `#payment-only`;
  private readonly accountOnlySelectedLocator: string = `[id='account-only'][style*='color: white']`;
  private readonly noAccountTitleLabelLocator: string = `#no-account-title`;
  private readonly noAccountSubtitleLabelLocator: string = `#no-account-subtitle`;

  public visit(force: boolean = false): void {
    Urls.visit(Cypress.env("consent_self_service_url"), force);
  }

  public clickOnAccountOnlyButton(): void {
    cy.get(this.accountOnlyLocator).should("be.visible").click();
    cy.get(this.accountOnlySelectedLocator).should("be.visible");
  }

  public clickOnApplicationCard(): void {
    cy.get(this.applicationCardLocator, { timeout: 30000 })
      .should("be.visible")
      .click();
  }

  public assertThatNoAccountsPageIsDisplayed(): void {
    cy.get(this.noAccountTitleLabelLocator, { timeout: 5000 })
      .should("be.visible")
      .should(`contain.text`, `No connected accounts`);
    cy.get(this.noAccountSubtitleLabelLocator, { timeout: 5000 })
      .should("be.visible")
      .should(
        `contain.text`,
        `You haven't connected any accounts yet to manage access`
      );
  }

  public assertThatApplicationCardIsNotDisplayed(): void {
    cy.get(this.applicationCardLocator, { timeout: 30000 }).should("not.exist");
  }

}
