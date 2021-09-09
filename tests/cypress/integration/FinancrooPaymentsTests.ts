import {AcpLoginPage} from '../pages/acp/AcpLoginPage';
import {FinancrooLoginPage} from '../pages/financroo/FinancrooLoginPage';
import {ConsentPage} from '../pages/consent/ConsentPage';
import {ErrorPage} from '../pages/ErrorPage';
import {FinancrooWelcomePage} from '../pages/financroo/FinancrooWelcomePage';
import {FinancrooAccountsPage} from '../pages/financroo/accounts/FinancrooAccountsPage';
import {Credentials} from "../pages/Credentials";
import {Urls} from "../pages/Urls";
import {MfaPage} from "../pages/mfa/MfaPage";
import {FinancrooInvestmentsPage} from "../pages/financroo/investments/FinancrooInvestmentsPage";
import {FinancrooContributePage} from "../pages/financroo/investments/FinancrooContributePage";
import {EnvironmentVariables} from "../pages/EnvironmentVariables"

describe(`Financroo payments app test`, () => {
  const acpLoginPage: AcpLoginPage = new AcpLoginPage();
  const consentPage: ConsentPage = new ConsentPage();
  const errorPage: ErrorPage = new ErrorPage();
  const financrooLoginPage: FinancrooLoginPage = new FinancrooLoginPage();
  const financrooWelcomePage: FinancrooWelcomePage = new FinancrooWelcomePage();
  const financrooAccountsPage: FinancrooAccountsPage = new FinancrooAccountsPage();
  const financrooInvestmentsPage: FinancrooInvestmentsPage = new FinancrooInvestmentsPage();
  const financrooContributePage: FinancrooContributePage = new FinancrooContributePage();
  const mfaPage: MfaPage = new MfaPage();
  const environmentVariables: EnvironmentVariables = new EnvironmentVariables();

  const amount: number = Math.floor(Math.random() * 50) + 1;

  beforeEach(() => {
    financrooLoginPage.visit()
    Urls.clearLocalStorage()
    financrooLoginPage.visit()
    financrooLoginPage.login()
    acpLoginPage.login(Credentials.financrooUsername, Credentials.defaultPassword)
    financrooWelcomePage.connect()
    acpLoginPage.login(Credentials.tppUsername, Credentials.defaultPassword)
    if (environmentVariables.isMfaEnabled()) {
      mfaPage.typePin()
    }
    consentPage.confirm()
  });

  it(`Happy path with confirm consent to add new amount`, () => {
    financrooLoginPage.visit()
    financrooAccountsPage.goToInvestmentsTab()
    financrooInvestmentsPage.invest()
    financrooContributePage.contribute(amount)
    acpLoginPage.login(Credentials.tppUsername, Credentials.defaultPassword)
    if (environmentVariables.isMfaEnabled()) {
      mfaPage.typePin()
    }
    consentPage.confirm()
    financrooContributePage.assertAmount(amount)
    financrooContributePage.assertItIsFinished()
  })

  it(`Reject path with decline consent to add new amount`, () => {
    financrooLoginPage.visit()
    financrooAccountsPage.goToInvestmentsTab()
    financrooInvestmentsPage.invest()
    financrooContributePage.contribute(amount + 1)
    acpLoginPage.login(Credentials.tppUsername, Credentials.defaultPassword)
    if (environmentVariables.isMfaEnabled()) {
      mfaPage.typePin()
    }
    consentPage.cancel()
    errorPage.assertError(`acp returned an error: rejected:`)
  })

  it(`Cancel on ACP login`, () => {
    financrooLoginPage.visit()
    financrooAccountsPage.goToInvestmentsTab()
    financrooInvestmentsPage.invest()
    financrooContributePage.contribute(amount + 2)
    acpLoginPage.cancel()
    errorPage.assertError(`The user rejected the authentication`)
  })

})
