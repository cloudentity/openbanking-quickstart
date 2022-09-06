import {FinancrooConnectAccountPage} from './accounts/FinancrooConnectAccountPage';

export class FinancrooWelcomePage {
  private readonly connectBankLocator: string = `#connect-bank-button`
  private readonly welcomeTitleLocator: string = `#financroo-welcome-title`
  private readonly welcometiSubtleLocator: string = `#financroo-welcome-subtitle`
  private readonly financrooConnectAccountPage: FinancrooConnectAccountPage = new FinancrooConnectAccountPage();

  public connectGoBank(): void {
    cy.get(this.connectBankLocator).then(ele => {
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
    cy.get(this.connectBankLocator).then(ele => {
      cy.wrap(ele).click()
      if (ele.text().includes(`disconnect`)) {
        this.connectSantanderBank()
      } else if (!ele.text().includes(`reconnect`)) {
        this.financrooConnectAccountPage.clickSantanderBankIcon()
        this.financrooConnectAccountPage.allow()
      }
    })
 }

 public assertThatConnectBankPageIsDisplayed(): void {
    cy.get(this.welcomeTitleLocator).should('be.visible')
    cy.get(this.welcomeTitleLocator).should('contain.text', 'Welcome to Financroo Smart Banking');
    cy.get(this.welcometiSubtleLocator).should('be.visible')
    cy.get(this.welcometiSubtleLocator).should('contain.text', 'Connect your bank, bills and credit cards to uncover insights that can improve your financial well being');
    cy.get(this.connectBankLocator).should('be.visible')
    cy.get(this.connectBankLocator).should('contain.text', 'Connect your bank');
 }
 
}
