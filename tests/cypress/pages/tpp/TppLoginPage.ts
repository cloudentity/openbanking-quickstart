import {Urls} from '../Urls'

export class TppLoginPage {
  private readonly nextButtonLocator: string = `[type="submit"]`;
  private readonly basicPermissionCheckboxLocator: string = `[value="ReadAccountsBasic"]`
  private readonly detailPermissionCheckboxLocator: string = `[value="ReadAccountsDetail"]`

  public visit(): void {
    Urls.forceVisit(Urls.tppTechnicalUrl);
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

  public next(): void {
    cy.get(this.nextButtonLocator).click();
  }
}
