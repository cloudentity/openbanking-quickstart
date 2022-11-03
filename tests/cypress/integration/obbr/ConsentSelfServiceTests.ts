import { AcpLoginPage } from "../../pages/acp/AcpLoginPage";
import { AccountConsentPage } from "../../pages/consent/AccountConsentPage";
import { ErrorPage } from "../../pages/ErrorPage";
import { ConsentSelfServicePage } from "../../pages/consent-self-service/ConsentSelfServicePage";
import { ConsentSelfServiceAccountDetailsPage } from "../../pages/consent-self-service/ConsentSelfServiceAccountDetailsPage";
import { Urls } from "../../pages/Urls";
import { Accounts } from "../../pages/Accounts";
import { FinancrooLoginPage } from "../../pages/financroo/FinancrooLoginPage";
import { FinancrooWelcomePage } from "../../pages/financroo/FinancrooWelcomePage";
import { FinancrooAccountsPage } from "../../pages/financroo/accounts/FinancrooAccountsPage";
import { ConsentSelfServiceApplicationPage } from "../../pages/consent-self-service/ConsentSelfServiceApplicationPage";
import { FinancrooModalPage } from "../../pages/financroo/accounts/FinancrooModalPage";

describe(`Consent self service app`, () => {
  const acpLoginPage: AcpLoginPage = new AcpLoginPage();
  const accountConsentPage: AccountConsentPage = new AccountConsentPage();
  const errorPage: ErrorPage = new ErrorPage();
  const consentSelfServicePage: ConsentSelfServicePage = new ConsentSelfServicePage();
  const consentSelfServiceAccountDetailsPage: ConsentSelfServiceAccountDetailsPage = new ConsentSelfServiceAccountDetailsPage();
  const consentSelfServiceApplicationPage: ConsentSelfServiceApplicationPage = new ConsentSelfServiceApplicationPage();
  const financrooLoginPage: FinancrooLoginPage = new FinancrooLoginPage();
  const financrooWelcomePage: FinancrooWelcomePage = new FinancrooWelcomePage();
  const financrooModalPage: FinancrooModalPage = new FinancrooModalPage();
  const financrooAccountsPage: FinancrooAccountsPage = new FinancrooAccountsPage();


  before(() => {
    financrooLoginPage.visit();
    Urls.clearLocalStorage();
    financrooLoginPage.visit();
    financrooLoginPage.login();

    financrooWelcomePage.reconnectGoBank();

    acpLoginPage.loginWithMfaOption();
    
    accountConsentPage.checkAllAccounts();
    accountConsentPage.clickAgree();

    financrooModalPage.assertThatModalIsDisplayed();

    financrooLoginPage.visit();

    financrooAccountsPage.assertThatPageIsDisplayed();
  });

  beforeEach(() => {
    consentSelfServicePage.visit(true);
    Urls.clearLocalStorage();
    consentSelfServicePage.visit(true);
  });

  it(`Happy path with account consent`, () => {
    acpLoginPage.login();

    consentSelfServicePage.clickOnAccountOnlyButton();
    consentSelfServicePage.clickOnApplicationCard();

    consentSelfServiceApplicationPage.expandAccountsTab();
    consentSelfServiceApplicationPage.checkAccount(Accounts.ids.BR.account1);
    consentSelfServiceApplicationPage.expandAccountConsentRow();

    consentSelfServiceAccountDetailsPage.assertThatAccountDetailsAreVisible()
    consentSelfServiceAccountDetailsPage.assertAccount(Accounts.ids.BR.account1);
  });

  it(`Revoke consent`, () => {
    acpLoginPage.login();

    consentSelfServicePage.clickOnAccountOnlyButton();
    consentSelfServicePage.clickOnApplicationCard();

    consentSelfServiceApplicationPage.expandAccountsTab();
    consentSelfServiceApplicationPage.assertNumberOfConsents(1);
    consentSelfServiceApplicationPage.expandAccountConsentRow();

    consentSelfServiceAccountDetailsPage.assertThatAccountDetailsAreVisible();
    consentSelfServiceAccountDetailsPage.clickRevokeAccessButton();
    consentSelfServiceAccountDetailsPage.assertThatRevokeAccountDetailsAreVisible();
    consentSelfServiceAccountDetailsPage.confirmRevokeAccessAction();
    
    consentSelfServiceApplicationPage.assertNumberOfConsents(0);
  });

  it(`Cancel ACP login`, () => {
    acpLoginPage.cancelLogin();
    // UI error page improvements AUT-5845
    errorPage.assertError("The user rejected the authentication");
  });
});
