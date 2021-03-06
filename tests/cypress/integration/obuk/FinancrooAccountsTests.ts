import {AcpLoginPage} from '../../pages/acp/AcpLoginPage';
import {FinancrooLoginPage} from '../../pages/financroo/FinancrooLoginPage';
import {ConsentPage} from '../../pages/consent/ConsentPage';
import {ErrorPage} from '../../pages/ErrorPage';
import {FinancrooWelcomePage} from '../../pages/financroo/FinancrooWelcomePage';
import {FinancrooAccountsPage} from '../../pages/financroo/accounts/FinancrooAccountsPage';
import {Credentials} from "../../pages/Credentials";
import {Urls} from "../../pages/Urls";
import {MfaPage} from "../../pages/mfa/MfaPage";
import {EnvironmentVariables} from "../../pages/EnvironmentVariables"

describe(`Financroo app`, () => {
  const acpLoginPage: AcpLoginPage = new AcpLoginPage();
  const consentPage: ConsentPage = new ConsentPage();
  const errorPage: ErrorPage = new ErrorPage();
  const financrooLoginPage: FinancrooLoginPage = new FinancrooLoginPage();
  const financrooWelcomePage: FinancrooWelcomePage = new FinancrooWelcomePage();
  const financrooAccountsPage: FinancrooAccountsPage = new FinancrooAccountsPage();
  const mfaPage: MfaPage = new MfaPage();
  const environmentVariables: EnvironmentVariables = new EnvironmentVariables();

  const billsAccount: string = `Bills`;
  const householdAccount: string = `Household`;

  beforeEach(() => {
    financrooLoginPage.visit()
    Urls.clearLocalStorage()
    financrooLoginPage.visit()
    financrooLoginPage.login()
  });

  [
    [billsAccount, householdAccount],
    [billsAccount],
    [householdAccount],
    []
  ].forEach(accounts => {
    it(`Happy path with accounts: ${accounts}`, () => {
      
      financrooWelcomePage.connectGoBank()
      acpLoginPage.login(Credentials.tppUsername, Credentials.defaultPassword)
      if (environmentVariables.isMfaEnabled()) {
        mfaPage.typePin()
      }
      consentPage.checkAccounts(accounts)
      consentPage.expandPermissions()
      consentPage.assertPermissions(7)
      consentPage.confirm()
      financrooAccountsPage.assertAccounts(accounts)
    })
  })

  it(`Cancel on ACP login`, () => {
    financrooWelcomePage.connectGoBank()
    acpLoginPage.cancel()
    // UI error page improvements AUT-5845
    errorPage.assertError(`The user rejected the authentication`)
  })

  it(`Cancel on consent`, () => {
    financrooWelcomePage.connectGoBank()
    acpLoginPage.login(Credentials.tppUsername, Credentials.defaultPassword)
    if (environmentVariables.isMfaEnabled()) {
      mfaPage.typePin()
    }
    consentPage.cancel()
    // UI error page improvements AUT-5845
    errorPage.assertError(`rejected`)
  })

})
