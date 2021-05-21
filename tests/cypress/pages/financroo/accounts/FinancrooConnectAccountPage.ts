export class FinancrooConnectAccountPage {
  private readonly goBankLocator: string = `#gobank`;
  private readonly cancelButtonLocator: string = `#cancel-button`;
  private readonly allowButtonLocator: string = `#allow-button`;

  public connectGoBank(): void {
    cy.get(this.goBankLocator).click()
  }

  public allow(): void {
    cy.get(this.allowButtonLocator).click()
  }

  public cancel(): void {
    cy.get(this.cancelButtonLocator).click()
  }
}
