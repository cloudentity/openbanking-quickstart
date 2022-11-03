import { FinancrooLoginPage } from "../../../pages/financroo/FinancrooLoginPage";
import { Accounts } from "../../../pages/Accounts";
import { FinancrooWelcomePage } from "../../../pages/financroo/FinancrooWelcomePage";
import { AcpLoginPage } from "../../../pages/acp/AcpLoginPage";
import { AccountConsentPage } from "../../../pages/consent/AccountConsentPage";
import { FinancrooModalPage } from "../../../pages/financroo/accounts/FinancrooModalPage";
import { FinancrooAccountsPage } from "../../../pages/financroo/accounts/FinancrooAccountsPage";

describe(`FDX Financroo app`, () => {
  const financrooLoginPage: FinancrooLoginPage = new FinancrooLoginPage();
  const financrooWelcomePage: FinancrooWelcomePage = new FinancrooWelcomePage();
  const acpLoginPage: AcpLoginPage = new AcpLoginPage();
  const accountConsentPage: AccountConsentPage = new AccountConsentPage();
  const financrooModalPage: FinancrooModalPage = new FinancrooModalPage();
  const financrooAccountsPage: FinancrooAccountsPage = new FinancrooAccountsPage();


  beforeEach(() => {
    financrooLoginPage.visit();
    financrooLoginPage.login();
  });

  [
    [Accounts.ids.FDX.checkingAcc, Accounts.ids.FDX.savings1, Accounts.ids.FDX.savings2],
    [Accounts.ids.FDX.savings1],
    [Accounts.ids.FDX.savings2],
    [Accounts.ids.FDX.checkingAcc],
  ].forEach((accountsIds) => {
    it(`Happy path with selected accounts: ${accountsIds}`, () => {
      financrooWelcomePage.reconnectGoBank();

      acpLoginPage.assertThatModalIsDisplayed("XXX");
      acpLoginPage.loginWithMfaOption();

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
      cy.wait(3000); // Workaround due to pipeline issues >>> AUT-6292
    });
  });

  it(`Happy path with not selected account`, () => {
    financrooWelcomePage.reconnectGoBank();

    acpLoginPage.assertThatModalIsDisplayed("XXX");
    acpLoginPage.loginWithMfaOption();
    
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
});
