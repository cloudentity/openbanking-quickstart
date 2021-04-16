export class Urls {
  public static readonly tppTechnicalUrl: string = `https://localhost:8090`;
  public static readonly financrooUrl: string = `https://localhost:8091`
  public static readonly acpUrl: string = `https://localhost:8443`
  public static readonly consentSelfServiceUrl: string = `https://localhost:8085`
  public static readonly consentAdminUrl: string = `https://localhost:8086`

  public static forceVisit(url: string) {
    cy.window().then(window => window.open(url, '_self'))
  }

  public static clearLocalStorage(): void {
    cy.clearLocalStorage();
  }

}
