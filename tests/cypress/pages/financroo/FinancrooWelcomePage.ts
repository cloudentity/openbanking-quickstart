import {FinancrooConnectAccountPage} from "./accounts/FinancrooConnectAccountPage";

export class FinancrooWelcomePage {
  private readonly financrooConnectAccountPage: FinancrooConnectAccountPage = new FinancrooConnectAccountPage();

  public connect(): void {
    cy.get(`[class*="connect-button"]`).then(ele => {
      cy.wrap(ele).click()
      if (ele.text().includes('disconnect')) {
        this.connect()
      } else if (!ele.text().includes('reconnect')) {
        this.financrooConnectAccountPage.connectGoBank()
        this.financrooConnectAccountPage.allow()
      }
    })
  }
}
