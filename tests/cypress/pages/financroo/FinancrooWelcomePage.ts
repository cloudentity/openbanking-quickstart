import {FinancrooConnectAccountPage} from "./FinancrooConnectAccountPage";

export class FinancrooWelcomePage {
  private readonly connectButtonLocator: string = `.connect-button`;
  private readonly financrooConnectAccountPage: FinancrooConnectAccountPage = new FinancrooConnectAccountPage();

  public connect(): void {
    cy.get(`[class*="connect-button"]`).then(ele => {
      ele.click()
      if (ele.text().includes('disconnect')) {
        this.connect()
      } else if (!ele.text().includes('reconnect')) {
        this.financrooConnectAccountPage.connectGoBank()
        this.financrooConnectAccountPage.allow()
      }
    })
  }
}
