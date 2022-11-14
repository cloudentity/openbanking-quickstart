import { AcpLoginPage } from "../../pages/acp/AcpLoginPage";
import { AccountConsentPage } from "../../pages/consent/AccountConsentPage";
import { Accounts } from "../../pages/Accounts";
import { FdxTppLandingPage } from "../../pages/fdx-tpp/FdxTppLandingPage";
import { FdxTppIntentRegisteredPage } from "../../pages/fdx-tpp/FdxTppIntentRegisteredPage";
import { FdxTppAuthenticatedPage } from "../../pages/fdx-tpp/FdxTppAuthenticatedPage";
import { ConsentAdminPage } from "../../pages/consent-admin/ConsentAdminPage";
import { TppErrorPage } from "../../pages/TppErrorPage";

describe(`FDX TPP Consent admin portal tests`, () => {
  const fdxTppLandingPage: FdxTppLandingPage = new FdxTppLandingPage();
  const fdxTppIntentRegisteredPage: FdxTppIntentRegisteredPage = new FdxTppIntentRegisteredPage();
  const fdxTppAuthenticatedPage: FdxTppAuthenticatedPage = new FdxTppAuthenticatedPage();
  const acpLoginPage: AcpLoginPage = new AcpLoginPage();
  const accountConsentPage: AccountConsentPage = new AccountConsentPage();
  const consentAdminPage: ConsentAdminPage = new ConsentAdminPage();
  const tppErrorPage: TppErrorPage = new TppErrorPage();

  beforeEach(() => {
    fdxTppLandingPage.visit();

    fdxTppLandingPage.assertThatPageIsDisplayed();
    fdxTppLandingPage.clickNext();

    fdxTppIntentRegisteredPage.assertThatPageIsDisplayed();
    fdxTppIntentRegisteredPage.clickLogin();

    acpLoginPage.assertThatModalIsDisplayed("FDX");
    acpLoginPage.loginWithMfaOption();

    accountConsentPage.assertPermissions(4);
    accountConsentPage.assertThatAccountsAreNotVisible([
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
      accountConsentPage,
      fdxTppAuthenticatedPage,
      fdxTppLandingPage,
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
      accountConsentPage,
      fdxTppAuthenticatedPage,
      fdxTppLandingPage,
      accountsIDs
    );

    consentAdminPage.visit(true);
    consentAdminPage.login();

    consentAdminPage.assertThatConsentManagementTabIsDisplayed();
    consentAdminPage.revokeClientConsent();
  });

  it(`Happy path with revoking consent with not selected accounts`, () => {
    accountConsentPage.clickContinue();
    accountConsentPage.clickAgree();

    fdxTppAuthenticatedPage.assertThatPageIsDisplayed();
    fdxTppAuthenticatedPage.assertThatConsentResponseFieldNotContainsAccountsIds(
      [
        Accounts.ids.FDX.checkingAcc,
        Accounts.ids.FDX.savings1,
        Accounts.ids.FDX.savings2,
      ]
    );

    consentAdminPage.visit(true);
    consentAdminPage.login();

    consentAdminPage.assertThatConsentManagementTabIsDisplayed();
    consentAdminPage.revokeClientConsent();
  });

  it("Cancel on consent page", () => {
    accountConsentPage.assertPermissions(4);
    accountConsentPage.clickCancel();

    tppErrorPage.assertThatRejectConsentErrorPageIsDisplayed(
      `Rejected`,
      `The user rejected the authentication.`,
      `consent_rejected`
    );
  });

  async function acceptConsentWithIds(
    accountConsentPage: AccountConsentPage,
    fdxTppAuthenticatedPage: FdxTppAuthenticatedPage,
    fdxTppLandingPage: FdxTppLandingPage,
    accountsIDs: string[]
  ) {
    accountConsentPage.clickContinue();
    accountConsentPage.checkAccounts(accountsIDs);
    accountConsentPage.clickAgree();

    fdxTppAuthenticatedPage.assertThatPageIsDisplayed();
    fdxTppAuthenticatedPage.assertThatTokenResponseFieldIsNotEmpty();
    fdxTppAuthenticatedPage.assertThatAccessTokenFieldIsNotEmpty();
    fdxTppAuthenticatedPage.assertThatConsentResponseFieldContainsAccountsIds(
      accountsIDs
    );

    fdxTppAuthenticatedPage.clickTryNext();
    fdxTppLandingPage.assertThatPageIsDisplayed();
  }
});
