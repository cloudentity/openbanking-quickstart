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

describe(`Foo`, () => {
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

  beforeEach(() => {
    financrooLoginPage.visit()
    Urls.clearLocalStorage()
    financrooLoginPage.visit()
    financrooLoginPage.login()
    acpLoginPage.login(Credentials.financrooUsername, Credentials.defaultPassword)
    financrooWelcomePage.connect()
    acpLoginPage.login(Credentials.tppUsername, Credentials.defaultPassword)
    if (environmentVariables.getMfaVariable()) {
      mfaPage.typePin()
    }
    consentPage.confirm()
  });

  it(`Happy path`, () => {
    financrooLoginPage.visit()
    financrooAccountsPage.goToInvestmentsTab()
    financrooInvestmentsPage.invest()
    financrooContributePage.contribute(1)
    acpLoginPage.login(Credentials.tppUsername, Credentials.defaultPassword)
    if (environmentVariables.getMfaVariable()) {
      mfaPage.typePin()
    }
    consentPage.confirm()
    financrooContributePage.assertItIsFinished()
  })

  it(`Cancel on ACP login`, () => {
    financrooLoginPage.visit()
    financrooAccountsPage.goToInvestmentsTab()
    financrooInvestmentsPage.invest()
    financrooContributePage.contribute(1)
    acpLoginPage.cancel()
    errorPage.assertError(`The user rejected the authentication`)
  })

})
