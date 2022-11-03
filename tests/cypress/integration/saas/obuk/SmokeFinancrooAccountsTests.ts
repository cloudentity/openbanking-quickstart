import { AcpLoginPage } from "../../../pages/acp/AcpLoginPage";
import { FinancrooLoginPage } from "../../../pages/financroo/FinancrooLoginPage";
import { AccountConsentPage } from "../../../pages/consent/AccountConsentPage";
import { FinancrooWelcomePage } from "../../../pages/financroo/FinancrooWelcomePage";
import { FinancrooAccountsPage } from "../../../pages/financroo/accounts/FinancrooAccountsPage";
import { Credentials } from "../../../pages/Credentials";
import { Urls } from "../../../pages/Urls";
import { Accounts } from "../../../pages/Accounts";
import { MfaPage } from "../../../pages/mfa/MfaPage";
import { EnvironmentVariables } from "../../../pages/EnvironmentVariables";
import { FinancrooModalPage } from "../../../pages/financroo/accounts/FinancrooModalPage";


describe(`Smoke Financroo app`, () => {
  const acpLoginPage: AcpLoginPage = new AcpLoginPage();
  const accountConsentPage: AccountConsentPage = new AccountConsentPage();
  const financrooLoginPage: FinancrooLoginPage = new FinancrooLoginPage();
  const financrooWelcomePage: FinancrooWelcomePage = new FinancrooWelcomePage();
  const financrooAccountsPage: FinancrooAccountsPage = new FinancrooAccountsPage();
  const mfaPage: MfaPage = new MfaPage();
  const environmentVariables: EnvironmentVariables = new EnvironmentVariables();
  const financrooModalPage: FinancrooModalPage = new FinancrooModalPage();


  beforeEach(() => {
    financrooLoginPage.visit();
    Urls.clearLocalStorage();
    financrooLoginPage.visit();
    financrooLoginPage.login();
  });

  [
    [Accounts.ids.UK.bills, Accounts.ids.UK.household],
    [Accounts.ids.UK.bills],
    [Accounts.ids.UK.household]
  ].forEach((accountsIds) => {
    it(`Happy path with accounts: ${accountsIds}`, () => {
      financrooWelcomePage.reconnectGoBank();

      acpLoginPage.login();
      if (environmentVariables.isMfaEnabled()) {
        mfaPage.typePin();
      }

      accountConsentPage.checkAccounts(accountsIds);
      accountConsentPage.expandPermissions();
      accountConsentPage.assertPermissions(7);
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

    acpLoginPage.login();
    if (environmentVariables.isMfaEnabled()) {
      mfaPage.typePin();
    }

    accountConsentPage.uncheckAllAccounts();
    accountConsentPage.clickAgree();

    financrooModalPage.assertThatModalIsDisplayed();
    financrooModalPage.close();

    financrooAccountsPage.assertThatPageIsDisplayed();
    financrooAccountsPage.assertAccountsSyncedNumber(0);
    financrooAccountsPage.disconnectAccounts();

    financrooWelcomePage.assertThatConnectBankPageIsDisplayed();
  });
});
