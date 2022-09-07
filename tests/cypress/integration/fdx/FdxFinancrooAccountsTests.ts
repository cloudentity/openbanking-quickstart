import {FinancrooLoginPage} from '../../pages/financroo/FinancrooLoginPage';
import { Urls } from "../../pages/Urls";
import {FinancrooWelcomePage} from '../../pages/financroo/FinancrooWelcomePage';
import { AcpLoginPage } from "../../pages/acp/AcpLoginPage";
import { Credentials } from "../../pages/Credentials";
import { EnvironmentVariables } from "../../pages/EnvironmentVariables";
import { MfaPage } from "../../pages/mfa/MfaPage";
import { ConsentPage } from "../../pages/consent/ConsentPage";
import {FinancrooModalPage} from '../../pages/financroo/accounts/FinancrooModalPage';
import {FinancrooAccountsPage} from '../../pages/financroo/accounts/FinancrooAccountsPage';
import {ErrorPage} from '../../pages/ErrorPage';


describe(`FDX Financroo app`, () => {
  const financrooLoginPage: FinancrooLoginPage = new FinancrooLoginPage();
  const financrooWelcomePage: FinancrooWelcomePage = new FinancrooWelcomePage();
  const acpLoginPage: AcpLoginPage = new AcpLoginPage();
  const environmentVariables: EnvironmentVariables = new EnvironmentVariables();
  const mfaPage: MfaPage = new MfaPage();
  const consentPage: ConsentPage = new ConsentPage();
  const financrooModalPage: FinancrooModalPage = new FinancrooModalPage();
  const financrooAccountsPage: FinancrooAccountsPage = new FinancrooAccountsPage();
  const errorPage: ErrorPage = new ErrorPage();

  const digitalBankingAccountId: string = `96534987`;
  const savingsAccountId: string = `1000001`;
  const savings2AccountId: string = `1000002`;

  beforeEach(() => {
    financrooLoginPage.visit()
    Urls.clearLocalStorage()
    financrooLoginPage.visit()
    financrooLoginPage.login()
  });

  [
    [digitalBankingAccountId, savings2AccountId],
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
      consentPage.clickContinue();
      consentPage.clickContinue();
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

  it('Cancel on consent page', () => {
    financrooWelcomePage.reconnectGoBank()

    acpLoginPage.login(Credentials.tppUsername, Credentials.defaultPassword)
    if (environmentVariables.isMfaEnabled()) {
      mfaPage.typePin()
    }
    consentPage.assertPermissions(4);
    consentPage.clickContinue();
    consentPage.cancel()

    // UI error page improvements AUT-5845
    errorPage.assertError(`acp returned an error: rejected: `);
  })

  it('Cancel on ACP login', () => {
    financrooWelcomePage.reconnectGoBank()

    acpLoginPage.cancel();
    // UI error page improvements AUT-5845
    errorPage.assertError(`The user rejected the authentication`);
  })
});
