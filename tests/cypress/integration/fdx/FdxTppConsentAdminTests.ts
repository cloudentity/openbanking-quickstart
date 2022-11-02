import { AcpLoginPage } from "../../pages/acp/AcpLoginPage";
import { AccountConsentPage } from "../../pages/consent/AccountConsentPage";
import { Credentials } from "../../pages/Credentials";
import { Urls } from "../../pages/Urls";
import { Accounts } from "../../pages/Accounts";
import { MfaPage } from "../../pages/mfa/MfaPage";
import { EnvironmentVariables } from "../../pages/EnvironmentVariables";
import { FdxTppLoginPage } from "../../pages/fdx-tpp/FdxTppLoginPage";
import { FdxTppIntentRegisteredPage } from "../../pages/fdx-tpp/FdxTppIntentRegisteredPage";
import { FdxTppAuthenticatedPage } from "../../pages/fdx-tpp/FdxTppAuthenticatedPage";
import { ConsentAdminPage } from "../../pages/consent-admin/ConsentAdminPage";

import { ErrorPage } from "../../pages/ErrorPage";

describe(`FDX Tpp Consent admin portal tests`, () => {
  const fdxTppLoginPage: FdxTppLoginPage = new FdxTppLoginPage();
  const fdxTppIntentRegisteredPage: FdxTppIntentRegisteredPage = new FdxTppIntentRegisteredPage();
  const fdxTppAuthenticatedPage: FdxTppAuthenticatedPage = new FdxTppAuthenticatedPage();
  const acpLoginPage: AcpLoginPage = new AcpLoginPage();
  const consentPage: AccountConsentPage = new AccountConsentPage();
  const mfaPage: MfaPage = new MfaPage();
  const environmentVariables: EnvironmentVariables = new EnvironmentVariables();
  const consentAdminPage: ConsentAdminPage = new ConsentAdminPage();

  const errorPage: ErrorPage = new ErrorPage();

  beforeEach(() => {
    fdxTppLoginPage.visit();
    Urls.clearLocalStorage();
    fdxTppLoginPage.visit();

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
    consentPage.assertThatAccountsAreNotVisible([
      Accounts.ids.FDX.checkingAcc,
      Accounts.ids.FDX.savings1,
      Accounts.ids.FDX.savings1,
    ]);
  });

  it(`Happy path with revoking consent from Consent management page`, () => {
    const accountsIDs = [
      Accounts.ids.FDX.checkingAcc,
      Accounts.ids.FDX.savings2,
    ];

    acceptConsentWithIds(
      consentPage,
      fdxTppAuthenticatedPage,
      fdxTppLoginPage,
      accountsIDs
    );

    consentAdminPage.visit(true);
    Urls.clearLocalStorage();
    consentAdminPage.visit(true);
    consentAdminPage.login();

    consentAdminPage.assertThatConsentManagementTabIsDisplayed();
    consentAdminPage.searchAccount(accountsIDs[0]);
    consentAdminPage.assertAccountResult(accountsIDs[0]);
    consentAdminPage.assertClientAccountWithStatus("Developer TPP", "Active");
    consentAdminPage.manageAccount("Developer TPP");
    consentAdminPage.assertConsentsDetails();
    consentAdminPage.revokeClientConsentByAccountName("Developer TPP");
    consentAdminPage.assertClientAccountWithStatus("Developer TPP", "Inactive");
  });

  it(`Happy path with revoking consent from Third party providers page`, () => {
    const accountsIDs = [
      Accounts.ids.FDX.savings1,
      Accounts.ids.FDX.savings2
    ];

    acceptConsentWithIds(
      consentPage,
      fdxTppAuthenticatedPage,
      fdxTppLoginPage,
      accountsIDs
    );

    consentAdminPage.visit(true);
    Urls.clearLocalStorage();
    consentAdminPage.visit(true);
    consentAdminPage.login();

    consentAdminPage.assertThatConsentManagementTabIsDisplayed();
    consentAdminPage.revokeClientConsent();
  });

  it(`Happy path with revoking consent with not selected accounts`, () => {
    consentPage.clickContinue();
    consentPage.clickAgree();

    fdxTppAuthenticatedPage.assertThatPageIsDisplayed();
    fdxTppAuthenticatedPage.assertThatConsentResponseFieldNotContainsAccountsIds(
      [
        Accounts.ids.FDX.checkingAcc,
        Accounts.ids.FDX.savings1,
        Accounts.ids.FDX.savings1,
      ]
    );

    consentAdminPage.visit(true);
    Urls.clearLocalStorage();
    consentAdminPage.visit(true);
    consentAdminPage.login();

    consentAdminPage.assertThatConsentManagementTabIsDisplayed();
    consentAdminPage.revokeClientConsent();
  });

  it("Cancel on consent page", () => {
    consentPage.assertPermissions(4);
    consentPage.clickCancel();

    // UI error page improvements AUT-5845
    errorPage.assertError(`acp returned an error: rejected: `);
  });
/*
  it("Cancel on ACP login", () => {
    fdxTppLoginPage.assertThatPageIsDisplayed();
    fdxTppLoginPage.assertThatAuthorizationDetailsAreDisplayed();
    fdxTppLoginPage.clickNext();

    fdxTppIntentRegisteredPage.assertThatPageIsDisplayed();
    fdxTppIntentRegisteredPage.assertThatRequestUriFieldsAreNotEmpty();
    fdxTppIntentRegisteredPage.clickLogin();

    acpLoginPage.cancel();
    // UI error page improvements AUT-5845
    errorPage.assertError(`The user rejected the authentication`);
  });
*/
  async function acceptConsentWithIds(
    consentPage: AccountConsentPage,
    fdxTppAuthenticatedPage: FdxTppAuthenticatedPage,
    fdxTppLoginPage: FdxTppLoginPage,
    accountsIDs: string[]
  ) {
    consentPage.clickContinue();
    consentPage.checkAccounts(accountsIDs);
    consentPage.clickAgree();

    fdxTppAuthenticatedPage.assertThatPageIsDisplayed();
    fdxTppAuthenticatedPage.assertThatTokenResponseFieldIsNotEmpty();
    fdxTppAuthenticatedPage.assertThatAccessTokenFieldIsNotEmpty();
    fdxTppAuthenticatedPage.assertThatConsentResponseFieldContainsAccountsIds(
      accountsIDs
    );

    fdxTppAuthenticatedPage.clickTryNext();
    fdxTppLoginPage.assertThatPageIsDisplayed();
  }
});
