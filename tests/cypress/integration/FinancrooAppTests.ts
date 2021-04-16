import {AcpLoginPage} from '../pages/acp/AcpLoginPage';
import {FinancrooLoginPage} from '../pages/financroo/FinancrooLoginPage';
import {ConsentPage} from '../pages/consent/ConsentPage';
import {ErrorPage} from '../pages/ErrorPage';
import {FinancrooWelcomePage} from '../pages/financroo/FinancrooWelcomePage';
import {FinancrooDashboardPage} from '../pages/financroo/FinancrooDashboardPage';
import {Credentials} from "../pages/Credentials";
import {Urls} from "../pages/Urls";

describe(`Financroo app`, () => {
  const acpLoginPage: AcpLoginPage = new AcpLoginPage();
  const consentPage: ConsentPage = new ConsentPage();
  const errorPage: ErrorPage = new ErrorPage();
  const financrooLoginPage: FinancrooLoginPage = new FinancrooLoginPage();
  const financrooWelcomePage: FinancrooWelcomePage = new FinancrooWelcomePage();
  const financrooDashboardPage: FinancrooDashboardPage = new FinancrooDashboardPage();

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
      acpLoginPage.login(Credentials.financrooUsername, Credentials.defaultPassword)
      financrooWelcomePage.connect()
      acpLoginPage.login(Credentials.tppUsername, Credentials.defaultPassword)
      consentPage.checkAccounts(accounts)
      consentPage.assertPermissions([`ReadAccountsDetail`, `ReadAccountsBasic`, `ReadBalances`,
        `ReadTransactionsBasic`, `ReadTransactionsDetail`, `ReadTransactionsCredits`, `ReadTransactionsDebits`])
      consentPage.confirm()
      financrooDashboardPage.assertAccounts(accounts)
    })
  })

  it(`Cancel on ACP login`, () => {
    acpLoginPage.cancel()
    errorPage.assertError(`The user rejected the authentication`)
  })

  it(`Cancel on second ACP login`, () => {
    acpLoginPage.login(Credentials.financrooUsername, Credentials.defaultPassword)
    financrooWelcomePage.connect()
    acpLoginPage.cancel()
    errorPage.assertError(`The user rejected the authentication`)
  })

  it(`Cancel on consent`, () => {
    acpLoginPage.login(Credentials.financrooUsername, Credentials.defaultPassword)
    financrooWelcomePage.connect()
    acpLoginPage.login(Credentials.tppUsername, Credentials.defaultPassword)
    consentPage.cancel()
    errorPage.assertError(`rejected`)
  })

})
