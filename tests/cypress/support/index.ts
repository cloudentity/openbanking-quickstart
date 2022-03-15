export {};

declare global {
  namespace Cypress {
    /**
     * Use `cy.disableSameSiteCookieRestrictions()` to disable SameSite cookie restrictions
     * and solve problem with invalid csrf token error.
     * Run `cy.disableSameSiteCookieRestrictions()` together with `cy.visit()` 
     * @example
     *    cy.disableSameSiteCookieRestrictions()
     *    cy.visit('http://localhost:3000')
     */
    interface Chainable<Subject> {
      disableSameSiteCookieRestrictions(): void;
    }
  }
}
