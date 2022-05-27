import { Urls } from '../Urls';

export class DiscoverDataHoldersPage {
  private readonly mainHeaderSelector: string = `[role='main'] > h2`;
  private readonly refreshDataHolderHeaderSelector: string = '#dh-refresh-header';
  private readonly currentDataHolderHeaderSelector: string = '#dh-current-header';
  private readonly refreshDataHolderTitleSelector: string = '#dh-refresh-title';
  private readonly versionFieldSelector = (node: string) => `#dh-refresh-version ${node}`;
  private readonly refreshButtonSelector: string = '#dh-refresh-btnRefresh';
  private readonly refreshMessageSelector: string = '#dh-refresh-message';

  public visit(force: boolean = false): void {
    Urls.visit(Cypress.env(`mock_data_recipient_url`) + '/data-holders', false);
    this.assertThatPageIsDisplayed();
  }

  public assertThatPageIsDisplayed(): void {
    cy.get(this.mainHeaderSelector)
      .should('have.text', 'Discover Data Holders');
    cy.get(this.refreshDataHolderHeaderSelector)
      .should('have.text', 'Refresh Data Holders');
    cy.get(this.currentDataHolderHeaderSelector)
      .should('have.text', 'Current Data Holders');
    cy.get(this.refreshDataHolderTitleSelector)
      .should('have.text', 'Call the Register to retrieve the Data Holder Brands');
    cy.get(this.versionFieldSelector('label'))
      .should('have.text', 'Version');
    cy.get(this.versionFieldSelector('#Version'))
      .should('be.visible');
    cy.get(this.refreshButtonSelector)
      .should('be.visible');
    cy.get(this.refreshMessageSelector)
      .should('have.text', 'Waiting...');
  }

  public clickRefreshDataHoldersButton(): void {
    cy.get(this.refreshButtonSelector).click();
  }

  public assertThatDataHolderBrandsLoaded(): void {
    cy.get(this.refreshMessageSelector, {timeout: 30000,})
      .should('have.text', 'OK - 30 data holder brands loaded.');
  }
}
