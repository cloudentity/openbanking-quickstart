export class FinancrooContributePage {
  private readonly title: string =`#title`;
  private readonly nextButtonLocator: string = `#next-button`;
  private readonly backToPortfolioButtonLocator: string = `#back-to-portfolio`;
  private readonly amountInput: string = `input`;
  private readonly amountDisplayed: string = `.MuiChip-label`;
  private readonly accountList: string = `#accounts-list`;

  public contribute(amount: number): void {
    cy.get(this.amountInput).type(String(amount))
    cy.get(this.nextButtonLocator).click()
    cy.get(this.title).contains('PAYMENT TOTAL')
    cy.get(this.accountList).contains('Checking account')
    cy.get(this.nextButtonLocator).click()
    cy.get(this.title).contains('INVESTMENT SUMMARY')
    cy.get(this.nextButtonLocator).click()
  }

  public assertItIsFinished(): void {
    cy.get(this.backToPortfolioButtonLocator).should(`contain.text`, `Back to portfolio`)
  }

  public assertAmount(amount: number): void {
    cy.get(this.amountDisplayed).should(`contain.text`, amount.toString() + " GBP")
  }
}
