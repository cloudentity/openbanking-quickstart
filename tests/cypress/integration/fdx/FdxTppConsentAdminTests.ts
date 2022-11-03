import { AcpLoginPage } from "../../pages/acp/AcpLoginPage";
import { AccountConsentPage } from "../../pages/consent/AccountConsentPage";
import { Accounts } from "../../pages/Accounts";
import { FdxTppLandingPage } from "../../pages/fdx-tpp/FdxTppLandingPage";
import { FdxTppIntentRegisteredPage } from "../../pages/fdx-tpp/FdxTppIntentRegisteredPage";
import { FdxTppAuthenticatedPage } from "../../pages/fdx-tpp/FdxTppAuthenticatedPage";
import { ConsentAdminPage } from "../../pages/consent-admin/ConsentAdminPage";

import { ErrorPage } from "../../pages/ErrorPage";

describe(`FDX Tpp Consent admin portal tests`, () => {
  const fdxTppLoginPage: FdxTppLandingPage = new FdxTppLandingPage();
  const fdxTppIntentRegisteredPage: FdxTppIntentRegisteredPage = new FdxTppIntentRegisteredPage();
  const fdxTppAuthenticatedPage: FdxTppAuthenticatedPage = new FdxTppAuthenticatedPage();
  const acpLoginPage: AcpLoginPage = new AcpLoginPage();
  const consentPage: AccountConsentPage = new AccountConsentPage();
  const consentAdminPage: ConsentAdminPage = new ConsentAdminPage();

  const errorPage: ErrorPage = new ErrorPage();

  beforeEach(() => {
    fdxTppLoginPage.visit();

    fdxTppLoginPage.assertThatPageIsDisplayed();
    fdxTppLoginPage.clickNext();

    fdxTppIntentRegisteredPage.assertThatPageIsDisplayed();
    fdxTppIntentRegisteredPage.clickLogin();

    acpLoginPage.assertThatModalIsDisplayed("FDX");
    acpLoginPage.confirmLogin();

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

  async function acceptConsentWithIds(
    consentPage: AccountConsentPage,
    fdxTppAuthenticatedPage: FdxTppAuthenticatedPage,
    fdxTppLoginPage: FdxTppLandingPage,
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
