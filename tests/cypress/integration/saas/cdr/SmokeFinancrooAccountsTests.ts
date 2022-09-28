import { FinancrooLoginPage } from "../../../pages/financroo/FinancrooLoginPage";
import { Urls } from "../../../pages/Urls";
import { FinancrooWelcomePage } from "../../../pages/financroo/FinancrooWelcomePage";
import { AcpLoginPage } from "../../../pages/acp/AcpLoginPage";
import { EnvironmentVariables } from "../../../pages/EnvironmentVariables";
import { MfaPage } from "../../../pages/mfa/MfaPage";
import { ConsentPage } from "../../../pages/consent/ConsentPage";
import { FinancrooModalPage } from "../../../pages/financroo/accounts/FinancrooModalPage";
import { FinancrooAccountsPage } from "../../../pages/financroo/accounts/FinancrooAccountsPage";
import { Credentials } from "../../../pages/Credentials";


describe(`Smoke Financroo app`, () => {
  const financrooLoginPage: FinancrooLoginPage = new FinancrooLoginPage();
  const financrooWelcomePage: FinancrooWelcomePage = new FinancrooWelcomePage();
  const acpLoginPage: AcpLoginPage = new AcpLoginPage();
  const environmentVariables: EnvironmentVariables = new EnvironmentVariables();
  const mfaPage: MfaPage = new MfaPage();
  const consentPage: ConsentPage = new ConsentPage();
  const financrooModalPage: FinancrooModalPage = new FinancrooModalPage();
  const financrooAccountsPage: FinancrooAccountsPage = new FinancrooAccountsPage();

  const savingsAccountId: string = `1000001`;
  const loanAccountId: string = `1000002`;
  const checkingsAccountId: string = `96534987`;

  beforeEach(() => {
    financrooLoginPage.visit();
    Urls.clearLocalStorage();
    financrooLoginPage.visit();
    financrooLoginPage.login();
  });

  [
    [savingsAccountId],
    [savingsAccountId, loanAccountId],
    [savingsAccountId, loanAccountId, checkingsAccountId]
  ].forEach((accountsIds) => {
    it(`Happy path with selected accounts: ${accountsIds}`, () => {
      financrooWelcomePage.reconnectGoBank();

      acpLoginPage.login(Credentials.tppUsername, Credentials.defaultPassword);
      if (environmentVariables.isMfaEnabled()) {
        mfaPage.typePin();
      }

      consentPage.checkAccounts(accountsIds);
      consentPage.clickConfirm();

      financrooModalPage.assertThatModalIsDisplayed()
      financrooModalPage.close()

      financrooAccountsPage.assertThatPageIsDisplayed()
      financrooAccountsPage.assertAccountsSyncedNumber(accountsIds.length)
      financrooAccountsPage.assertAccountsIds(accountsIds)
      financrooAccountsPage.disconnectAccounts()

      financrooWelcomePage.assertThatConnectBankPageIsDisplayed()
      cy.wait(3000) // Workaround due to pipeline issues >>> AUT-6292
    });
  });

  it(`Happy path with not selected account`, () => {
    financrooWelcomePage.reconnectGoBank()

    acpLoginPage.login(Credentials.tppUsername, Credentials.defaultPassword)
    if (environmentVariables.isMfaEnabled()) {
      mfaPage.typePin()
    }

    consentPage.uncheckAllAccounts();
    consentPage.clickConfirm();

    financrooModalPage.assertThatModalIsDisplayed()
    financrooModalPage.close()

    financrooAccountsPage.assertThatPageIsDisplayed()
    financrooAccountsPage.assertAccountsSyncedNumber(0)
    financrooAccountsPage.disconnectAccounts()

    financrooWelcomePage.assertThatConnectBankPageIsDisplayed()
});
});
