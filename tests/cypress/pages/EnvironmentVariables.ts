export class EnvironmentVariables {

    public getMfaVariable() {
        const newLocal = Cypress.env('isMfaEnabled')
        this.logVariable(newLocal, 'isMfaEnabled')
        return newLocal
    }

    private logVariable(value: any, name: any): void {
        Cypress.log({
            displayName: `Environment variable ${name} >>> ${value}`
        })
        console.log(`${name}: ${value}`)
    }
}
