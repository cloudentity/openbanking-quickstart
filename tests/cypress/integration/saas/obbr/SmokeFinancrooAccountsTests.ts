import { AcpLoginPage } from "../../../pages/acp/AcpLoginPage";
import { FinancrooLoginPage } from "../../../pages/financroo/FinancrooLoginPage";
import { AccountConsentPage } from "../../../pages/consent/AccountConsentPage";
import { FinancrooWelcomePage } from "../../../pages/financroo/FinancrooWelcomePage";
import { FinancrooAccountsPage } from "../../../pages/financroo/accounts/FinancrooAccountsPage";
import { Accounts } from "../../../pages/Accounts";
import { FinancrooModalPage } from "../../../pages/financroo/accounts/FinancrooModalPage";

describe(`Financroo app`, () => {
  const acpLoginPage: AcpLoginPage = new AcpLoginPage();
  const accountConsentPage: AccountConsentPage = new AccountConsentPage();
  const financrooLoginPage: FinancrooLoginPage = new FinancrooLoginPage();
  const financrooWelcomePage: FinancrooWelcomePage = new FinancrooWelcomePage();
  const financrooAccountsPage: FinancrooAccountsPage = new FinancrooAccountsPage();
  const financrooModalPage: FinancrooModalPage = new FinancrooModalPage();


  beforeEach(() => {
    financrooLoginPage.visit();
    financrooLoginPage.login();
  });

  [
    [Accounts.ids.BR.account1, Accounts.ids.BR.account2],
    [Accounts.ids.BR.account1],
    [Accounts.ids.BR.account2],
  ].forEach((accountsIds) => {
    it(`Happy path with accounts: ${accountsIds}`, () => {
      financrooWelcomePage.reconnectGoBank();

      acpLoginPage.assertThatModalIsDisplayed();
      acpLoginPage.loginWithMfaOption();

      accountConsentPage.checkAccounts(accountsIds);
      accountConsentPage.expandPermissions();
      accountConsentPage.assertPermissionsDetails(
        "Purpose for sharing data",
        "To uncover insights that can improve your financial well being."
      );
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

    acpLoginPage.assertThatModalIsDisplayed();
    acpLoginPage.loginWithMfaOption();

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
