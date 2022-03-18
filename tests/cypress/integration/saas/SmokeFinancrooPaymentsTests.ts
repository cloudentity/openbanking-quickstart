import {AcpLoginPage} from '../../pages/acp/AcpLoginPage';
import {FinancrooLoginPage} from '../../pages/financroo/FinancrooLoginPage';
import {ConsentPage} from '../../pages/consent/ConsentPage';
import {FinancrooWelcomePage} from '../../pages/financroo/FinancrooWelcomePage';
import {FinancrooAccountsPage} from '../../pages/financroo/accounts/FinancrooAccountsPage';
import {Credentials} from "../../pages/Credentials";
import {Urls} from "../../pages/Urls";
import {MfaPage} from "../../pages/mfa/MfaPage";
import {FinancrooInvestmentsPage} from "../../pages/financroo/investments/FinancrooInvestmentsPage";
import {FinancrooContributePage} from "../../pages/financroo/investments/FinancrooContributePage";
import {EnvironmentVariables} from "../../pages/EnvironmentVariables"
import { FinancrooModalPage } from '../../pages/financroo/accounts/FinancrooModalPage';

describe(`Smoke Financroo payments app test`, () => {
  const acpLoginPage: AcpLoginPage = new AcpLoginPage();
  const consentPage: ConsentPage = new ConsentPage();
  const financrooLoginPage: FinancrooLoginPage = new FinancrooLoginPage();
  const financrooWelcomePage: FinancrooWelcomePage = new FinancrooWelcomePage();
  const financrooModalPage: FinancrooModalPage = new FinancrooModalPage();
  const financrooAccountsPage: FinancrooAccountsPage = new FinancrooAccountsPage();
  const financrooInvestmentsPage: FinancrooInvestmentsPage = new FinancrooInvestmentsPage();
  const financrooContributePage: FinancrooContributePage = new FinancrooContributePage();
  const mfaPage: MfaPage = new MfaPage();
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
  });

  it(`Happy path with confirm consent to add new amount`, () => {
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
    financrooContributePage.assertAmount(amount)
    financrooContributePage.assertItIsFinished()
  })
})
