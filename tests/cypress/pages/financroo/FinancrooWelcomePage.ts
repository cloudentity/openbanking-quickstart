import {FinancrooConnectAccountPage} from './accounts/FinancrooConnectAccountPage';

export class FinancrooWelcomePage {
  private readonly financrooConnectAccountPage: FinancrooConnectAccountPage = new FinancrooConnectAccountPage();

  public connectGoBank(): void {
    cy.get(`[class*="connect-button"]`).then(ele => {
      cy.wrap(ele).click()
      if (ele.text().includes(`disconnect`)) {
        this.connectGoBank()
      } else if (!ele.text().includes(`reconnect`)) {
        this.financrooConnectAccountPage.clickGoBankIcon()
        this.financrooConnectAccountPage.allow()
      }
    })
  }

  public connectSantanderBank(): void {
    cy.get(`[class*="connect-button"]`).then(ele => {
      cy.wrap(ele).click()
      if (ele.text().includes(`disconnect`)) {
        this.connectSantanderBank()
      } else if (!ele.text().includes(`reconnect`)) {
        this.financrooConnectAccountPage.clickSantanderBankIcon()
        this.financrooConnectAccountPage.allow()
      }
    })
 }
}


