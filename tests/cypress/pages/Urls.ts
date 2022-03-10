export class Urls {

  public static visit(url: string, force: boolean) {
    cy.disableSameSiteCookieRestrictions();
    if (force) {
      cy.window().then(window => window.open(url, '_self'))
    } else {
      cy.visit(url)
    }
  }

  public static clearLocalStorage(): void {
    cy.clearLocalStorage();
  }

}
