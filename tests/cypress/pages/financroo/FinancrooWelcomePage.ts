import {FinancrooConnectAccountPage} from './accounts/FinancrooConnectAccountPage';

export class FinancrooWelcomePage {
  private readonly accessBankLocator: string = `#access-bank-button`
  private readonly welcomeTitleLocator: string = `#financroo-welcome-title`
  private readonly welcomeSubtleLocator: string = `#financroo-welcome-subtitle`
  private readonly financrooConnectAccountPage: FinancrooConnectAccountPage = new FinancrooConnectAccountPage();

  public reconnectGoBank(): void {
    cy.get(this.accessBankLocator).then(ele => {
      cy.wrap(ele).click()
      if (ele.text().includes(`disconnect`)) {
        this.reconnectGoBank()
      } else if (ele.text().includes(`Connect your bank`)) {
        this.financrooConnectAccountPage.clickGoBankIcon()
        this.financrooConnectAccountPage.allow()
      }
    })
  }

 public assertThatConnectBankPageIsDisplayed(): void {
    cy.get(this.welcomeTitleLocator).should('be.visible')
    cy.get(this.welcomeTitleLocator).should('contain.text', 'Welcome to Financroo Smart Banking');
    cy.get(this.welcomeSubtleLocator).should('be.visible')
    cy.get(this.welcomeSubtleLocator).should('contain.text', 'Connect your bank, bills and credit cards to uncover insights that can improve your financial well being');
    cy.get(this.accessBankLocator).should('be.visible')
    cy.get(this.accessBankLocator).should('contain.text', 'Connect your bank');
 }
 
}
