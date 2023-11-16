export class FinancrooConnectAccountPage {
  private readonly goBankLocator: string = `#gobank`;
  private readonly hyperscaleBankLocator: string = '#hyperscalebank';
  private readonly cancelButtonLocator: string = `#cancel-button`;
  private readonly allowButtonLocator: string = `#allow-button`;

  public clickGoBankIcon(): void {
    cy.get(this.goBankLocator).click()
  }

  public clickHyperscaleBankIcon(): void {
    cy.get(this.hyperscaleBankLocator).click()
  }

  public allow(): void {
    cy.get(this.allowButtonLocator).click()
  }

  public cancel(): void {
    cy.get(this.cancelButtonLocator).click()
  }
}
