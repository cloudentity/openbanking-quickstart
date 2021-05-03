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


  beforeEach(() => {
    financrooLoginPage.visit()
    Urls.clearLocalStorage()
    financrooLoginPage.visit()
    financrooLoginPage.login()
    acpLoginPage.login(Credentials.financrooUsername, Credentials.defaultPassword)
    financrooWelcomePage.connect()
    acpLoginPage.login(Credentials.tppUsername, Credentials.defaultPassword)
    mfaPage.typePin()
    consentPage.confirm()
  });

  it(`Foo1`, () => {
    financrooLoginPage.visit()
    financrooAccountsPage.goToInvestmentsTab()
    financrooInvestmentsPage.invest()
    financrooContributePage.contribute(1)
  })

})
