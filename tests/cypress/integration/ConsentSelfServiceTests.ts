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
import { MockDataRecipientPage } from '../pages/mock-data-recipient/MockDataRecipientPage';

describe(`Consent self service app CDR`, () => {
  const acpLoginPage: AcpLoginPage = new AcpLoginPage();
  const consentPage: ConsentPage = new ConsentPage();
  const consentSelfServicePage: ConsentSelfServicePage = new ConsentSelfServicePage();
  const consentSelfServiceApplicationPage: ConsentSelfServiceApplicationPage = new ConsentSelfServiceApplicationPage();
  const environmentVariables: EnvironmentVariables = new EnvironmentVariables();
  const mockDataRecipientPage: MockDataRecipientPage = new MockDataRecipientPage(); 

  before(() => {
    mockDataRecipientPage.visit()
    
    // TODO: refresh wait? 
    mockDataRecipientPage.visitDiscoverDataHoldersTab()
    mockDataRecipientPage.clickDataHoldersRefresh()

    mockDataRecipientPage.visitDynamicClientRegistrationTab() 
    mockDataRecipientPage.clickDCRRegisterButton()

    mockDataRecipientPage.visitConsentAndAuthorisationTab() 
    mockDataRecipientPage.selectClientRegistration(1)
    mockDataRecipientPage.inputSharingDuration(1000000)
    mockDataRecipientPage.clickConstructAuthorisationURI(); 
    
    acpLoginPage.login(Credentials.tppUsername, Credentials.defaultPassword);

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
})

describe(`Consent self service app`, () => {
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

  const amount: number = Math.floor(Math.random() * 50) + 1;

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
    consentSelfServiceApplicationPage.assertAmount(amount)
  })

  it(`Cancel ACP login`, () => {
    acpLoginPage.cancel();
    errorPage.assertError("The user rejected the authentication")
  })

})
