import {AcpLoginPage} from '../pages/acp/AcpLoginPage';
import {FinancrooLoginPage} from '../pages/financroo/FinancrooLoginPage';
import {ConsentPage} from '../pages/consent/ConsentPage';
import {ErrorPage} from '../pages/ErrorPage';
import {FinancrooWelcomePage} from '../pages/financroo/FinancrooWelcomePage';
import {FinancrooAccountsPage} from '../pages/financroo/FinancrooAccountsPage';
import {Credentials} from "../pages/Credentials";
import {Urls} from "../pages/Urls";
import {MfaPage} from "../pages/mfa/MfaPage";

describe(`Financroo app`, () => {
  const acpLoginPage: AcpLoginPage = new AcpLoginPage();
  const consentPage: ConsentPage = new ConsentPage();
  const errorPage: ErrorPage = new ErrorPage();
  const financrooLoginPage: FinancrooLoginPage = new FinancrooLoginPage();
  const financrooWelcomePage: FinancrooWelcomePage = new FinancrooWelcomePage();
  const financrooAccountsPage: FinancrooAccountsPage = new FinancrooAccountsPage();
  const mfaPage: MfaPage = new MfaPage();

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
      mfaPage.typePin()
      consentPage.checkAccounts(accounts)
      consentPage.expandPermissions()
      consentPage.assertPermissions(7)
      consentPage.confirm()
      financrooAccountsPage.assertAccounts(accounts)
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
    mfaPage.typePin()
    consentPage.cancel()
    errorPage.assertError(`rejected`)
  })

})
