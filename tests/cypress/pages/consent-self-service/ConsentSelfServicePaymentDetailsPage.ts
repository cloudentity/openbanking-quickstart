export class ConsentSelfServicePaymentDetailsPage {

  private readonly purposeLocator: string = `#payment-purpose`;
  private readonly detailsLocator: string = `#payment-details`;
  private readonly permissionDatesLocator: string = `#payment-permission-dates`;
  private readonly detailsSharedLocator: string = `#payment-details-shared`;


  public assertThatPaymentDetailsAreVisible(): void {
    cy.get(this.purposeLocator).should(`contain.text`, `Purpose for sharing data`);
    cy.get(this.detailsLocator).should(`contain.text`, `TRANSACTION Details`);
    cy.get(this.permissionDatesLocator).should(`contain.text`, `Permission dates`);
    cy.get(this.detailsSharedLocator).should(`contain.text`, `Details being shared`);
  }

  public assertAmount(currency: string, amount: number): void {
    cy.get(this.detailsLocator).should(`contain.text`, currency + " " + amount.toFixed(2));
  }

  public assertAccount(accountId: string): void {
    cy.get(this.detailsLocator).should(`contain.text`, accountId);
  }

}
