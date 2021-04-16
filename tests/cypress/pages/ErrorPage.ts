export class ErrorPage {
  public assertError(error: String): void {
    cy.get("body").should(`contain.text`, error)
  }
}
