export class FinancrooInvestmentsPage {
  private readonly investButtonLocator: string = `#invest-button`;

  public invest(): void {
    cy.get(this.investButtonLocator).click()
  }
}
