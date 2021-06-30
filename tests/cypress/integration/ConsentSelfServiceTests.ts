import {TppIntentPage} from '../pages/tpp/TppIntentPage';
import {TppLoginPage} from '../pages/tpp/TppLoginPage';
import {AcpLoginPage} from '../pages/acp/AcpLoginPage';
import {ConsentPage} from '../pages/consent/ConsentPage';
import {ErrorPage} from '../pages/ErrorPage';
import {Credentials} from "../pages/Credentials";
import {ConsentSelfServicePage} from '../pages/consent-self-service/ConsentSelfServicePage';
import {Urls} from "../pages/Urls";
import {MfaPage} from "../pages/mfa/MfaPage";
import {FinancrooLoginPage} from "../pages/financroo/FinancrooLoginPage";
import {FinancrooWelcomePage} from "../pages/financroo/FinancrooWelcomePage";
import {FinancrooAccountsPage} from "../pages/financroo/accounts/FinancrooAccountsPage";
import {FinancrooInvestmentsPage} from "../pages/financroo/investments/FinancrooInvestmentsPage";
import {FinancrooContributePage} from "../pages/financroo/investments/FinancrooContributePage";
import {ConsentSelfServiceApplicationPage} from "../pages/consent-self-service/ConsentSelfServiceApplicationPage";
import {EnvironmentVariables} from "../pages/EnvironmentVariables"

describe(`Consent self service app`, () => {
  const tppIntentPage: TppIntentPage = new TppIntentPage();
  const tppLoginPage: TppLoginPage = new TppLoginPage();
  const acpLoginPage: AcpLoginPage = new AcpLoginPage();
  const consentPage: ConsentPage = new ConsentPage();
  const errorPage: ErrorPage = new ErrorPage();
  const consentSelfServicePage: ConsentSelfServicePage = new ConsentSelfServicePage();
  const consentSelfServiceApplicationPage: ConsentSelfServiceApplicationPage = new ConsentSelfServiceApplicationPage();
  const mfaPage: MfaPage = new MfaPage();
  const financrooLoginPage: FinancrooLoginPage = new FinancrooLoginPage();
  const financrooWelcomePage: FinancrooWelcomePage = new FinancrooWelcomePage();
  const financrooAccountsPage: FinancrooAccountsPage = new FinancrooAccountsPage();
  const financrooInvestmentsPage: FinancrooInvestmentsPage = new FinancrooInvestmentsPage();
  const financrooContributePage: FinancrooContributePage = new FinancrooContributePage();
  const environmentVariables: EnvironmentVariables = new EnvironmentVariables();

  before(() => {
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

    financrooLoginPage.visit()
    financrooAccountsPage.goToInvestmentsTab()
    financrooInvestmentsPage.invest()
    financrooContributePage.contribute(1)
    acpLoginPage.login(Credentials.tppUsername, Credentials.defaultPassword)
    if (environmentVariables.isMfaEnabled()) {
      mfaPage.typePin()
    }
    consentPage.confirm()
  })

  beforeEach(() => {
    consentSelfServicePage.visit(true)
  })

  it(`Happy path with account consent`, () => {
    acpLoginPage.login(Credentials.tppUsername, Credentials.defaultPassword);
    consentSelfServicePage.clickOnApplicationCard()
    consentSelfServiceApplicationPage.expandAccountsTab()
    consentSelfServiceApplicationPage.checkAccount()
    consentSelfServiceApplicationPage.expandAccountConsentRow()
  })

  it(`Happy path with payment consent`, () => {
    acpLoginPage.login(Credentials.tppUsername, Credentials.defaultPassword);
    consentSelfServicePage.clickOnApplicationCard()
    consentSelfServiceApplicationPage.expandPaymentsTab()
    consentSelfServiceApplicationPage.checkAccount()
    consentSelfServiceApplicationPage.expandPaymentConsentRow()
  })

  it(`Cancel ACP login`, () => {
    acpLoginPage.cancel();
    errorPage.assertError("The user rejected the authentication")
  })

})
