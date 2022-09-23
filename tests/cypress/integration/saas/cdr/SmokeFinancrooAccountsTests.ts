import {AcpLoginPage} from '../../../pages/acp/AcpLoginPage';
import {FinancrooLoginPage} from '../../../pages/financroo/FinancrooLoginPage';
import {ConsentPage} from '../../../pages/consent/ConsentPage';
import {FinancrooWelcomePage} from '../../../pages/financroo/FinancrooWelcomePage';
import {FinancrooAccountsPage} from '../../../pages/financroo/accounts/FinancrooAccountsPage';
import {Credentials} from "../../../pages/Credentials";
import {Urls} from "../../../pages/Urls";
import {MfaPage} from "../../../pages/mfa/MfaPage";
import {EnvironmentVariables} from "../../../pages/EnvironmentVariables"

describe(`Smoke Financroo app`, () => {
  const acpLoginPage: AcpLoginPage = new AcpLoginPage();
  const consentPage: ConsentPage = new ConsentPage();
  const financrooLoginPage: FinancrooLoginPage = new FinancrooLoginPage();
  const financrooWelcomePage: FinancrooWelcomePage = new FinancrooWelcomePage();
  const financrooAccountsPage: FinancrooAccountsPage = new FinancrooAccountsPage();
  const mfaPage: MfaPage = new MfaPage();
  const environmentVariables: EnvironmentVariables = new EnvironmentVariables();

  const savingsAccount: string = `Savings Account`;
  const loanAccount: string = `Loan Account`; 
  const checkingsAccount: string = `Checkings Account`; 

  beforeEach(() => {
    financrooLoginPage.visit()
    Urls.clearLocalStorage()
    financrooLoginPage.visit()
    financrooLoginPage.login()
  });

  [
    [savingsAccount],
    [savingsAccount, loanAccount],
    [savingsAccount, loanAccount, checkingsAccount],
    []
  ].forEach(accounts => {
    it(`Happy path with accounts: ${accounts}`, () => {
      cy.wait(5000)
      financrooWelcomePage.reconnectGoBank()
      acpLoginPage.login(Credentials.tppUsername, Credentials.defaultPassword)
      if (environmentVariables.isMfaEnabled()) {
        mfaPage.typePin()
      }
      consentPage.checkAccounts(accounts)
      consentPage.confirm()
      financrooAccountsPage.assertAccounts(accounts)
    })
  })

})
