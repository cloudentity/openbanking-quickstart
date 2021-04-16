export class TppAuthenticatedPage {
  private readonly statusTextLocator: string = `.aut-demo-info-text`;

  public assertSuccess(): void {
    cy.get(this.statusTextLocator).should(`contain.text`, `Authenticated`)
  }
}
