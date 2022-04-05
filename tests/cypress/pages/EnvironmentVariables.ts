export class EnvironmentVariables {

    public isMfaEnabled(): boolean {
        return Cypress.env('ENABLE_MFA') === 'true'
    }
}
