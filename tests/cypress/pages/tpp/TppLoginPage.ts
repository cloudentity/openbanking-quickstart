import { Urls } from '../Urls'

export class TppLoginPage {
  private readonly nextButtonLocator: string = `[type="submit"]`;
  private readonly basicPermissionCheckboxLocator: string = `[value="ReadAccountsBasic"]`
  private readonly detailPermissionCheckboxLocator: string = `[value="ReadAccountsDetail"]`
  private readonly accountsReadPermissionCheckboxLocator: string = `[value="ACCOUNTS_READ"]`
  private readonly accountsOverdraftLimitsReadPermissionCheckboxLocator: string = `[value="ACCOUNTS_OVERDRAFT_LIMITS_READ"]`
  private readonly resourcesReadPermissionCheckboxLocator: string = `[value="RESOURCES_READ"]`

  public visit(force: boolean = false): void {
    Urls.visit(Cypress.env('tpp_url'), force);
    Urls.clearLocalStorage();
    Urls.visit(Cypress.env('tpp_url'), force);
  }

  public checkBasicPermission(check: boolean): void {
    const basicCheckbox = cy.get(this.basicPermissionCheckboxLocator);
    if (check) {
      basicCheckbox.check()
    } else {
      basicCheckbox.uncheck()
    }
  }

  public checkDetailPermission(check: boolean): void {
    const detailCheckbox = cy.get(this.detailPermissionCheckboxLocator);
    if (check) {
      detailCheckbox.check()
    } else {
      detailCheckbox.uncheck()
    }
  }

  public checkAccountsReadPermission(check: boolean): void {
    const cb = cy.get(this.accountsReadPermissionCheckboxLocator);
    if (check) {
      cb.check();
    } else {
      cb.uncheck();
    }
  }

  public checkAccountsOverdraftLimitsReadPermission(check: boolean): void {
    const cb = cy.get(this.accountsOverdraftLimitsReadPermissionCheckboxLocator);
    if (check) {
      cb.check();
    } else {
      cb.uncheck();
    }
  }

  public checkResourcesReadPermission(check: boolean): void {
    const cb = cy.get(this.resourcesReadPermissionCheckboxLocator);
    if (check) {
      cb.check();
    } else {
      cb.uncheck();
    }
  }

  public next(): void {
    cy.get(this.nextButtonLocator).click();
  }
}
