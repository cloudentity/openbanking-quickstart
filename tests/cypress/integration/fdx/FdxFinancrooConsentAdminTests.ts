import { FinancrooLoginPage } from "../../pages/financroo/FinancrooLoginPage";
import { Accounts } from "../../pages/Accounts";
import { FinancrooWelcomePage } from "../../pages/financroo/FinancrooWelcomePage";
import { AcpLoginPage } from "../../pages/acp/AcpLoginPage";
import { AccountConsentPage } from "../../pages/consent/AccountConsentPage";
import { ConsentAdminPage } from "../../pages/consent-admin/ConsentAdminPage";
import { FinancrooModalPage } from "../../pages/financroo/accounts/FinancrooModalPage";
import { FinancrooAccountsPage } from "../../pages/financroo/accounts/FinancrooAccountsPage";
import { ErrorPage } from "../../pages/ErrorPage";

describe(`FDX Financroo Consent admin portal tests`, () => {
  const financrooLoginPage: FinancrooLoginPage = new FinancrooLoginPage();
  const financrooWelcomePage: FinancrooWelcomePage = new FinancrooWelcomePage();
  const acpLoginPage: AcpLoginPage = new AcpLoginPage();
  const accountConsentPage: AccountConsentPage = new AccountConsentPage();
  const consentAdminPage: ConsentAdminPage = new ConsentAdminPage();
  const financrooModalPage: FinancrooModalPage = new FinancrooModalPage();
  const financrooAccountsPage: FinancrooAccountsPage = new FinancrooAccountsPage();
  const errorPage: ErrorPage = new ErrorPage();


  beforeEach(() => {
    financrooLoginPage.visit();
    financrooLoginPage.login();

    financrooWelcomePage.reconnectGoBank();

    acpLoginPage.assertThatModalIsDisplayed();
    acpLoginPage.loginWithMfaOption();

    accountConsentPage.assertPermissions(4);
    accountConsentPage.assertThatAccountsAreNotVisible([
      Accounts.ids.FDX.checkingAcc,
      Accounts.ids.FDX.savings1,
      Accounts.ids.FDX.savings2,
    ]);
  });

  it(`Happy path with revoking consent from Consent management page`, () => {
    const accountsIDs = [
      Accounts.ids.FDX.savings1,
      Accounts.ids.FDX.savings2,
    ];

    acceptConsentWithIds(
      accountConsentPage,
      financrooModalPage,
      financrooAccountsPage,
      accountsIDs
    );

    consentAdminPage.visit(true);
    consentAdminPage.login();

    consentAdminPage.assertThatConsentManagementTabIsDisplayed();
    consentAdminPage.searchAccount(accountsIDs[0]);
    consentAdminPage.assertAccountResult(accountsIDs[0]);
    consentAdminPage.assertClientAccountWithStatus("Financroo", "Active");
    consentAdminPage.manageAccount("Financroo");
    consentAdminPage.assertConsentsDetails();
    consentAdminPage.revokeClientConsentByAccountName("Financroo");
    consentAdminPage.assertClientAccountWithStatus("Financroo", "Inactive");
  });

  it(`Happy path with revoking consent from Third party providers page`, () => {
    const accountsIDs = [
      Accounts.ids.FDX.savings2,
      Accounts.ids.FDX.checkingAcc,
    ];

    acceptConsentWithIds(
      accountConsentPage,
      financrooModalPage,
      financrooAccountsPage,
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

    financrooAccountsPage.assertThatPageIsDisplayed();
    financrooAccountsPage.assertAccountsSyncedNumber(0);

    consentAdminPage.visit(true);
    consentAdminPage.login();

    consentAdminPage.assertThatConsentManagementTabIsDisplayed();
    consentAdminPage.revokeClientConsent();
  });

  it("Cancel on consent page", () => {
    accountConsentPage.assertPermissions(4);
    accountConsentPage.clickCancel();

    // UI error page improvements AUT-5845
    errorPage.assertError(`acp returned an error: rejected:`);
  });

  afterEach(() => {
    financrooLoginPage.visit();
    financrooLoginPage.login();

    financrooWelcomePage.assertThatConnectBankPageIsDisplayed()
  });

  async function acceptConsentWithIds(
    accountConsentPage: AccountConsentPage,
    financrooModalPage: FinancrooModalPage,
    financrooAccountsPage: FinancrooAccountsPage,
    accountsIDs: string[]
  ) {
    accountConsentPage.clickContinue();
    accountConsentPage.checkAccounts(accountsIDs);
    accountConsentPage.clickAgree();

    financrooAccountsPage.assertThatPageIsDisplayed();
    financrooAccountsPage.assertAccountsSyncedNumber(accountsIDs.length);
    financrooAccountsPage.assertAccountsIds(accountsIDs);
  }
});
