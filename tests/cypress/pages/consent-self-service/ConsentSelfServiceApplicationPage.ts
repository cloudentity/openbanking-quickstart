import {Urls} from '../Urls';

export class ConsentSelfServiceApplicationPage {

  private readonly accountsTabLocator: string = `.accounts-tab`;
  private readonly paymentsTabLocator: string = `.payments-tab`;
  private readonly revokeButtonLocator: string = `#revoke-access-button`
  private readonly purposeHeaderLocator: string = `.purpose-header`
  private readonly consentRowLocator: string = `.consent-row`

  public visit(force: boolean = false): void {
    Urls.visit(Urls.consentSelfServiceUrl, force);
  }

  public expandAccountsTab(): void {
    cy.get(this.accountsTabLocator).click();
  }

  public expandPaymentsTab(): void {
    cy.get(this.paymentsTabLocator).click();
  }

  public expandAccountConsentRow(): void {
    cy.get(this.consentRowLocator).first().click();
    cy.get(this.revokeButtonLocator).should(`contain.text`, `Revoke access`)
  }

  public expandPaymentConsentRow(): void {
    cy.get(this.consentRowLocator).first().click();
    cy.get(this.purposeHeaderLocator).should(`contain.text`, `Purpose for sharing data`)
  }

}
