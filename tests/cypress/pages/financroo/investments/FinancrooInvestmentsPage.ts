export class FinancrooInvestmentsPage {
  private readonly investButtonLocator: string = `#invest-button`;
  private readonly getDashboardLocator = (currency: string) => `#dashboard-${currency}`;
  private readonly transactionCompleteLocator: string = `#transaction-completed-title`;
  private readonly totalAmountLocator: string = `#total-amount`;
  private readonly backToPortfolioButtonLocator: string = `#back-to-portfolio`;

  public assertThatDashboardIsVisible(currency: string): void {
    cy.get(this.getDashboardLocator(currency)).should('be.visible');
  }

  public clickInvest(): void {
    cy.get(this.investButtonLocator).click();
  }

  public assertThatTransactionWasCompleted(amount: number, currency: string): void {
    cy.get(this.transactionCompleteLocator).contains('Transaction completed');
    cy.get(this.totalAmountLocator).contains(amount.toFixed(2) + " " + currency);
    cy.get(this.backToPortfolioButtonLocator).should('be.visible');
  }

  public clickBackToPortfolio(): void {
    cy.get(this.backToPortfolioButtonLocator).click();
  }
}
