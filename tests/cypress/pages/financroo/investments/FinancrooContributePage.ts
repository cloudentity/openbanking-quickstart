export class FinancrooContributePage {
  private readonly nextButtonLocator: string = `.MuiButton-containedPrimary`;
  private readonly amountInput: string = `input`;

  public contribute(amount: number): void {
    cy.get(this.amountInput).type(String(amount))
    cy.get(this.nextButtonLocator).click()
    cy.wait(1000)
    cy.get(this.nextButtonLocator).click()
    cy.wait(1000)
    cy.get(this.nextButtonLocator).click()
  }
}
