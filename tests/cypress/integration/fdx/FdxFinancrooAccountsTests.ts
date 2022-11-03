import { FinancrooLoginPage } from "../../pages/financroo/FinancrooLoginPage";
import { Urls } from "../../pages/Urls";
import { Accounts } from "../../pages/Accounts";
import { FinancrooWelcomePage } from "../../pages/financroo/FinancrooWelcomePage";
import { AcpLoginPage } from "../../pages/acp/AcpLoginPage";
import { Credentials } from "../../pages/Credentials";
import { EnvironmentVariables } from "../../pages/EnvironmentVariables";
import { MfaPage } from "../../pages/mfa/MfaPage";
import { AccountConsentPage } from "../../pages/consent/AccountConsentPage";
import { FinancrooModalPage } from "../../pages/financroo/accounts/FinancrooModalPage";
import { FinancrooAccountsPage } from "../../pages/financroo/accounts/FinancrooAccountsPage";
import { ErrorPage } from "../../pages/ErrorPage";

describe(`FDX Financroo app`, () => {
  const financrooLoginPage: FinancrooLoginPage = new FinancrooLoginPage();
  const financrooWelcomePage: FinancrooWelcomePage = new FinancrooWelcomePage();
  const acpLoginPage: AcpLoginPage = new AcpLoginPage();
  const environmentVariables: EnvironmentVariables = new EnvironmentVariables();
  const mfaPage: MfaPage = new MfaPage();
  const accountConsentPage: AccountConsentPage = new AccountConsentPage();
  const financrooModalPage: FinancrooModalPage = new FinancrooModalPage();
  const financrooAccountsPage: FinancrooAccountsPage = new FinancrooAccountsPage();
  const errorPage: ErrorPage = new ErrorPage();


  beforeEach(() => {
    financrooLoginPage.visit();
    Urls.clearLocalStorage();
    financrooLoginPage.visit();
    financrooLoginPage.login();
  });

  [
    [Accounts.ids.FDX.checkingAcc, Accounts.ids.FDX.savings2],
    [Accounts.ids.FDX.savings1],
    [Accounts.ids.FDX.savings2],
  ].forEach((accountsIds) => {
    it(`Happy path with selected accounts: ${accountsIds}`, () => {
      financrooWelcomePage.reconnectGoBank();

      acpLoginPage.login(Credentials.defaultUsername, Credentials.defaultPassword);
      if (environmentVariables.isMfaEnabled()) {
        mfaPage.typePin();
      }
      accountConsentPage.assertPermissions(4);
      accountConsentPage.assertThatAccountsAreNotVisible(accountsIds);
      accountConsentPage.clickContinue();
      accountConsentPage.checkAccounts(accountsIds);
      accountConsentPage.clickAgree();

      financrooModalPage.assertThatModalIsDisplayed();
      financrooModalPage.close();

      financrooAccountsPage.assertThatPageIsDisplayed();
      financrooAccountsPage.assertAccountsSyncedNumber(accountsIds.length);
      financrooAccountsPage.assertAccountsIds(accountsIds);
      financrooAccountsPage.disconnectAccounts();

      financrooWelcomePage.assertThatConnectBankPageIsDisplayed();
    });
  });

  it(`Happy path with not selected account`, () => {
    financrooWelcomePage.reconnectGoBank();

    acpLoginPage.login(Credentials.defaultUsername, Credentials.defaultPassword);
    if (environmentVariables.isMfaEnabled()) {
      mfaPage.typePin();
    }
    accountConsentPage.assertPermissions(4);
    accountConsentPage.clickContinue();
    accountConsentPage.clickAgree();

    financrooModalPage.assertThatModalIsDisplayed();
    financrooModalPage.close();

    financrooAccountsPage.assertThatPageIsDisplayed();
    financrooAccountsPage.assertAccountsSyncedNumber(0);
    financrooAccountsPage.disconnectAccounts();

    financrooWelcomePage.assertThatConnectBankPageIsDisplayed();
  });

  it("Cancel on consent page", () => {
    financrooWelcomePage.reconnectGoBank();

    acpLoginPage.login(Credentials.defaultUsername, Credentials.defaultPassword);
    if (environmentVariables.isMfaEnabled()) {
      mfaPage.typePin();
    }
    accountConsentPage.assertPermissions(4);
    accountConsentPage.clickContinue();
    accountConsentPage.clickCancel();

    // UI error page improvements AUT-5845
    errorPage.assertError(`acp returned an error: rejected: `);
  });

  it("Cancel on ACP login", () => {
    financrooWelcomePage.reconnectGoBank();

    acpLoginPage.cancelLogin();
    // UI error page improvements AUT-5845
    errorPage.assertError(`The user rejected the authentication`);
  });
});
