import {AcpLoginPage} from '../../pages/acp/AcpLoginPage';
import {ConsentPage} from '../../pages/consent/ConsentPage';
import {ErrorPage} from '../../pages/ErrorPage';
import {Credentials} from "../../pages/Credentials";
import {ConsentSelfServicePage} from '../../pages/consent-self-service/ConsentSelfServicePage';
import {Urls} from "../../pages/Urls";
import {MfaPage} from "../../pages/mfa/MfaPage";
import {FinancrooLoginPage} from "../../pages/financroo/FinancrooLoginPage";
import {FinancrooWelcomePage} from "../../pages/financroo/FinancrooWelcomePage";
import {FinancrooAccountsPage} from "../../pages/financroo/accounts/FinancrooAccountsPage";
import {FinancrooInvestmentsPage} from "../../pages/financroo/investments/FinancrooInvestmentsPage";
import {FinancrooContributePage} from "../../pages/financroo/investments/FinancrooContributePage";
import {ConsentSelfServiceApplicationPage} from "../../pages/consent-self-service/ConsentSelfServiceApplicationPage";
import {EnvironmentVariables} from "../../pages/EnvironmentVariables"
import { FinancrooModalPage } from '../../pages/financroo/accounts/FinancrooModalPage';


describe(`Consent self service app`, () => {
  const acpLoginPage: AcpLoginPage = new AcpLoginPage();
  const consentPage: ConsentPage = new ConsentPage();
  const errorPage: ErrorPage = new ErrorPage();
  const consentSelfServicePage: ConsentSelfServicePage = new ConsentSelfServicePage();
  const consentSelfServiceApplicationPage: ConsentSelfServiceApplicationPage = new ConsentSelfServiceApplicationPage();
  const mfaPage: MfaPage = new MfaPage();
  const financrooLoginPage: FinancrooLoginPage = new FinancrooLoginPage();
  const financrooWelcomePage: FinancrooWelcomePage = new FinancrooWelcomePage();
  const financrooModalPage: FinancrooModalPage = new FinancrooModalPage();
  const financrooAccountsPage: FinancrooAccountsPage = new FinancrooAccountsPage();
  const financrooInvestmentsPage: FinancrooInvestmentsPage = new FinancrooInvestmentsPage();
  const financrooContributePage: FinancrooContributePage = new FinancrooContributePage();
  const environmentVariables: EnvironmentVariables = new EnvironmentVariables();

  const amount: number = Math.floor(Math.random() * 50) + 1;

  before(() => {
    financrooLoginPage.visit()
    Urls.clearLocalStorage()
    financrooLoginPage.visit()
    financrooLoginPage.login()

    financrooWelcomePage.connectGoBank()
    acpLoginPage.login(Credentials.tppUsername, Credentials.defaultPassword)
    if (environmentVariables.isMfaEnabled()) {
      mfaPage.typePin()
    }
    consentPage.confirm()
    financrooModalPage.assertThatModalIsDisplayed()

    financrooLoginPage.visit()
    financrooAccountsPage.assertThatPageIsDisplayed()
    financrooAccountsPage.goToInvestmentsTab()
    financrooInvestmentsPage.invest()
    financrooContributePage.contribute(amount)
    acpLoginPage.login(Credentials.tppUsername, Credentials.defaultPassword)
    if (environmentVariables.isMfaEnabled()) {
      mfaPage.typePin()
    }
    consentPage.confirm()
  })

  beforeEach(() => {
    consentSelfServicePage.visit()
    Urls.clearLocalStorage()
    consentSelfServicePage.visit()
    acpLoginPage.login(Credentials.tppUsername, Credentials.defaultPassword);
  })

  it(`Happy path with account consent`, () => {
    consentSelfServicePage.clickOnApplicationCard()
    consentSelfServiceApplicationPage.expandAccountsTab()
    consentSelfServiceApplicationPage.checkAccount("22289")
    consentSelfServiceApplicationPage.expandAccountConsentRow()
  })

  it(`Happy path with payment consent`, () => {
    consentSelfServicePage.clickOnApplicationCard()
    consentSelfServiceApplicationPage.expandPaymentsTab()
    consentSelfServiceApplicationPage.checkAccount("22289")
    consentSelfServiceApplicationPage.expandPaymentConsentRow()
    consentSelfServiceApplicationPage.assertAmount(amount)
  })

  it(`Cancel ACP login`, () => {
    acpLoginPage.cancel();
    // UI error page improvements AUT-5845
    errorPage.assertError("The user rejected the authentication")
  })

})
