import {Urls} from '../Urls'

export class MockDataRecipientPage {
    public visit(force: boolean = false): void {
        Urls.visit(Cypress.env('mock_data_recipient_url'), force);
    }
    public visitDiscoverDataHoldersTab(): void {
        Urls.visit(Cypress.env(`mock_data_recipient_url`)+'/data-holders', false); 
    }
    public clickDataHoldersRefresh(): void {
        cy.get('.btn.btn-primary').contains('Refresh').click(); 
    }

    public visitDynamicClientRegistrationTab(): void {
        Urls.visit(Cypress.env(`mock_data_recipient_url`)+'/dcr', false); 
    }
    public clickDCRRegisterButton(): void {
        cy.get('.btn.btn-primary').contains('Register').click(); 
    }

    public visitConsentAndAuthorisationTab(): void {
        Urls.visit(Cypress.env(`mock_data_recipient_url`)+'/consent', false); 
    }
    public selectClientRegistration(idx: number = 0): void {
        cy.get(`select>option`).eq(idx).then(element => cy.get(`select`).select(element.val().toString()));
    }
    public inputSharingDuration(duration: number): void {
        cy.get('#SharingDuration').type(duration.toString()); 
    }
    public clickConstructAuthorisationURI(): void {
        cy.get(`.btn.btn-primary`).contains('Construct Authorisation Uri').click();
        cy.get(`.results>a`).click();
    }
}
