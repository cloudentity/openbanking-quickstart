import { AcpLoginPage } from "../../../pages/acp/AcpLoginPage";
import { ConsentPage } from "../../../pages/consent/ConsentPage";
import { Credentials } from "../../../pages/Credentials";
import { Urls } from "../../../pages/Urls";
import { FdxTppLoginPage } from "../../../pages/fdx-tpp/FdxTppLoginPage";
import { FdxTppIntentRegisteredPage } from "../../../pages/fdx-tpp/FdxTppIntentRegisteredPage";
import { FdxTppAuthenticatedPage } from "../../../pages/fdx-tpp/FdxTppAuthenticatedPage";
import { MfaPage } from "../../../pages/mfa/MfaPage";
import { EnvironmentVariables } from "../../../pages/EnvironmentVariables";

describe(`FDX Tpp consent app`, () => {
  const fdxTppLoginPage: FdxTppLoginPage = new FdxTppLoginPage();
  const fdxTppIntentRegisteredPage: FdxTppIntentRegisteredPage =
    new FdxTppIntentRegisteredPage();
  const fdxTppAuthenticatedPage: FdxTppAuthenticatedPage =
    new FdxTppAuthenticatedPage();
  const acpLoginPage: AcpLoginPage = new AcpLoginPage();
  const consentPage: ConsentPage = new ConsentPage();
  const mfaPage: MfaPage = new MfaPage();
  const environmentVariables: EnvironmentVariables = new EnvironmentVariables();

  const creditsAccountId: string = `96565987`;
  const savingsAccountId: string = `1122334455`;

  beforeEach(() => {
    fdxTppLoginPage.visit();
    Urls.clearLocalStorage();
    fdxTppLoginPage.visit();
  });

  [
    [creditsAccountId, savingsAccountId],
    [savingsAccountId],
    [creditsAccountId],
  ].forEach((accountsIds) => {
    it(`Happy path with selected accounts: ${accountsIds}`, () => {
      fdxTppLoginPage.assertThatPageIsDisplayed();
      fdxTppLoginPage.assertThatAuthorizationDetailsAreDisplayed();
      fdxTppLoginPage.clickNext();

      fdxTppIntentRegisteredPage.assertThatPageIsDisplayed();
      fdxTppIntentRegisteredPage.assertThatRequestUriFieldsAreNotEmpty();
      fdxTppIntentRegisteredPage.clickLogin();

      acpLoginPage.login(Credentials.tppUsername, Credentials.defaultPassword);
      if (environmentVariables.isMfaEnabled()) {
        mfaPage.typePin();
      }

      consentPage.assertPermissions(4);
      consentPage.clickContinue();
      consentPage.checkAccounts(accountsIds);
      consentPage.confirm();

      fdxTppAuthenticatedPage.assertThatPageIsDisplayed();
      fdxTppAuthenticatedPage.assertThatTokenResponseFieldIsNotEmpty();
      fdxTppAuthenticatedPage.assertThatAccessTokenFieldIsNotEmpty();
      fdxTppAuthenticatedPage.assertThatConsentResponseFieldContainsAccountsIds(
        accountsIds
      );

      fdxTppAuthenticatedPage.clickTryNext();
      fdxTppLoginPage.assertThatPageIsDisplayed();
      cy.wait(3000) // Workaround due to pipeline issues >>> AUT-6292
    });
  });

  it(`Happy path with not selected account`, () => {
    fdxTppLoginPage.assertThatPageIsDisplayed();
    fdxTppLoginPage.assertThatAuthorizationDetailsAreDisplayed();
    fdxTppLoginPage.clickNext();

    fdxTppIntentRegisteredPage.assertThatPageIsDisplayed();
    fdxTppIntentRegisteredPage.assertThatRequestUriFieldsAreNotEmpty();
    fdxTppIntentRegisteredPage.clickLogin();

    acpLoginPage.login(Credentials.tppUsername, Credentials.defaultPassword);
    if (environmentVariables.isMfaEnabled()) {
      mfaPage.typePin();
    }

    consentPage.clickContinue();
    consentPage.confirm();

    fdxTppAuthenticatedPage.assertThatPageIsDisplayed();
    fdxTppAuthenticatedPage.assertThatConsentResponseFieldNotContainsAccountsIds(
      [creditsAccountId, savingsAccountId]
    );
  });

});
