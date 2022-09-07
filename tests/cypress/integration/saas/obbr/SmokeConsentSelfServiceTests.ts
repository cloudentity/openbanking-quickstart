import {AcpLoginPage} from '../../../pages/acp/AcpLoginPage';
import {ConsentPage} from '../../../pages/consent/ConsentPage';
import {Credentials} from "../../../pages/Credentials";
import {ConsentSelfServicePage} from '../../../pages/consent-self-service/ConsentSelfServicePage';
import {Urls} from "../../../pages/Urls";
import {MfaPage} from "../../../pages/mfa/MfaPage";
import {FinancrooLoginPage} from "../../../pages/financroo/FinancrooLoginPage";
import {FinancrooWelcomePage} from "../../../pages/financroo/FinancrooWelcomePage";
import {FinancrooAccountsPage} from "../../../pages/financroo/accounts/FinancrooAccountsPage";
import {ConsentSelfServiceApplicationPage} from "../../../pages/consent-self-service/ConsentSelfServiceApplicationPage";
import {EnvironmentVariables} from "../../../pages/EnvironmentVariables"
import {FinancrooModalPage} from '../../../pages/financroo/accounts/FinancrooModalPage';


describe(`Consent self service app`, () => {
  const acpLoginPage: AcpLoginPage = new AcpLoginPage();
  const consentPage: ConsentPage = new ConsentPage();
  const consentSelfServicePage: ConsentSelfServicePage = new ConsentSelfServicePage();
  const consentSelfServiceApplicationPage: ConsentSelfServiceApplicationPage = new ConsentSelfServiceApplicationPage();
  const mfaPage: MfaPage = new MfaPage();
  const financrooLoginPage: FinancrooLoginPage = new FinancrooLoginPage();
  const financrooWelcomePage: FinancrooWelcomePage = new FinancrooWelcomePage();
  const financrooModalPage: FinancrooModalPage = new FinancrooModalPage();
  const financrooAccountsPage: FinancrooAccountsPage = new FinancrooAccountsPage();
  const environmentVariables: EnvironmentVariables = new EnvironmentVariables();

  before(() => {
    financrooLoginPage.visit()
    Urls.clearLocalStorage()
    financrooLoginPage.visit()
    financrooLoginPage.login()

    financrooWelcomePage.reconnectGoBank()
    acpLoginPage.login(Credentials.tppUsername, Credentials.defaultPassword)
    if (environmentVariables.isMfaEnabled()) {
      mfaPage.typePin()
    }
    consentPage.confirm()
    financrooModalPage.assertThatModalIsDisplayed()

    financrooLoginPage.visit()
    financrooAccountsPage.assertThatPageIsDisplayed()
  })

  beforeEach(() => {
    consentSelfServicePage.visit(true)
    Urls.clearLocalStorage()
    consentSelfServicePage.visit(true)
    acpLoginPage.login(Credentials.tppUsername, Credentials.defaultPassword);
    consentSelfServicePage.clickOnAccountOnlyButton()
    consentSelfServicePage.clickOnApplicationCard()
  })

  it(`Happy path with account consent`, () => {
    consentSelfServiceApplicationPage.expandAccountsTab()
    consentSelfServiceApplicationPage.checkAccount("94088392")
    consentSelfServiceApplicationPage.expandAccountConsentRow()
  })

  it(`Revoke consent`, () => {
    consentSelfServiceApplicationPage.expandAccountsTab()
    consentSelfServiceApplicationPage.assertNumberOfConsents(1)
    consentSelfServiceApplicationPage.expandAccountConsentRow()
    consentSelfServiceApplicationPage.clickRevokeAccessButton()
    consentSelfServicePage.assertThatNoAccountsPageIsDisplayed()
  })

})
