export class FinancrooModalPage {
    private readonly closeIcon: string = `#close-icon`
    private readonly modalContentLabel: string = `#modal-content`
    private readonly cancelButton: string = `#cancel-button`
    private readonly startInvestingButton: string = `#start-investing-button`

    public assertThatModalIsDisplayed(): void {
        cy.get(this.closeIcon).should('be.visible');
        cy.get(this.modalContentLabel).should('contain.text', 'Your Go Bank account(s) has been successfully connected to Financroo');
        cy.get(this.cancelButton).should('have.text', 'Cancel');
        cy.get(this.startInvestingButton).should('have.text', 'Start investing');
    }

    public close(): void {
        cy.get(this.closeIcon).click();
        this.assertThatModalIsNotDisplayed();
    }

    public cancel(): void {
        cy.get(this.cancelButton).click();
        this.assertThatModalIsNotDisplayed();
    }

    public startIvesting(): void {
        cy.get(this.startInvestingButton).click()
    }

    private assertThatModalIsNotDisplayed(): void {
        cy.get(this.closeIcon).should('not.be.visible');
        cy.get(this.modalContentLabel).should('not.exist');
        cy.get(this.cancelButton).should('be.visible', false);
        cy.get(this.startInvestingButton).should('be.visible', false);
    }
}