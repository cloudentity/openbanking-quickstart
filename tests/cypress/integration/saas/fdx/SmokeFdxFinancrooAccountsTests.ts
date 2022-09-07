import {FinancrooLoginPage} from '../../../pages/financroo/FinancrooLoginPage';
import { Urls } from "../../../pages/Urls";
import {FinancrooWelcomePage} from '../../../pages/financroo/FinancrooWelcomePage';
import { AcpLoginPage } from "../../../pages/acp/AcpLoginPage";
import { Credentials } from "../../../pages/Credentials";
import { EnvironmentVariables } from "../../../pages/EnvironmentVariables";
import { MfaPage } from "../../../pages/mfa/MfaPage";
import { ConsentPage } from "../../../pages/consent/ConsentPage";
import {FinancrooModalPage} from '../../../pages/financroo/accounts/FinancrooModalPage';
import {FinancrooAccountsPage} from '../../../pages/financroo/accounts/FinancrooAccountsPage';


describe(`FDX Financroo app`, () => {
  const financrooLoginPage: FinancrooLoginPage = new FinancrooLoginPage();
  const financrooWelcomePage: FinancrooWelcomePage = new FinancrooWelcomePage();
  const acpLoginPage: AcpLoginPage = new AcpLoginPage();
  const environmentVariables: EnvironmentVariables = new EnvironmentVariables();
  const mfaPage: MfaPage = new MfaPage();
  const consentPage: ConsentPage = new ConsentPage();
  const financrooModalPage: FinancrooModalPage = new FinancrooModalPage();
  const financrooAccountsPage: FinancrooAccountsPage = new FinancrooAccountsPage();

  const creditsAccountId: string = `96534987`;
  const savingsAccountId: string = `1000002`;
  const savings2AccountId: string = `1000001`;

  beforeEach(() => {
    financrooLoginPage.visit()
    Urls.clearLocalStorage()
    financrooLoginPage.visit()
    financrooLoginPage.login()
  });

  [
    [creditsAccountId, savingsAccountId, savings2AccountId],
    [creditsAccountId],
    [savingsAccountId],
    [savings2AccountId],
  ].forEach((accountsIds) => {
    it(`Happy path with selected accounts: ${accountsIds}`, () => {
      financrooWelcomePage.reconnectGoBank()

      acpLoginPage.login(Credentials.tppUsername, Credentials.defaultPassword)
      if (environmentVariables.isMfaEnabled()) {
        mfaPage.typePin()
      }
      consentPage.assertPermissions(4);
      consentPage.assertThatAccountsAreNotVisible(accountsIds);
      consentPage.confirm();
      consentPage.checkAccounts(accountsIds)
      consentPage.confirm();

      financrooModalPage.assertThatModalIsDisplayed()
      financrooModalPage.close()

      financrooAccountsPage.assertThatPageIsDisplayed()
      financrooAccountsPage.disconnectAccounts()

      financrooWelcomePage.assertThatConnectBankPageIsDisplayed()
    });
  });

  // it(`Happy path with not selected account`, () => {
  //   TBA
  // });

});
