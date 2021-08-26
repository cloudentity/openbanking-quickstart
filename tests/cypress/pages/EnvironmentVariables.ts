export class EnvironmentVariables {

    public isMfaEnabled(): boolean {
        return Cypress.env('ENABLE_MFA') === 'true'
    }

    public isOBBRSpecification(): boolean {
        return Cypress.env("SPEC") === 'obbr'
    }

    public isOBUKSpecification(): boolean {
        return Cypress.env("SPEC") === 'obuk'
    }

}
