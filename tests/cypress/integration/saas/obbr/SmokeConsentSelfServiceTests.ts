import { AcpLoginPage } from "../../../pages/acp/AcpLoginPage";
import { AccountConsentPage } from "../../../pages/consent/AccountConsentPage";
import { ConsentSelfServicePage } from "../../../pages/consent-self-service/ConsentSelfServicePage";
import { ConsentSelfServiceAccountDetailsPage } from "../../../pages/consent-self-service/ConsentSelfServiceAccountDetailsPage";
import { Accounts } from "../../../pages/Accounts";
import { FinancrooLoginPage } from "../../../pages/financroo/FinancrooLoginPage";
import { FinancrooWelcomePage } from "../../../pages/financroo/FinancrooWelcomePage";
import { FinancrooAccountsPage } from "../../../pages/financroo/accounts/FinancrooAccountsPage";
import { ConsentSelfServiceApplicationPage } from "../../../pages/consent-self-service/ConsentSelfServiceApplicationPage";
import { FinancrooModalPage } from "../../../pages/financroo/accounts/FinancrooModalPage";

describe(`Consent self service app`, () => {
  const acpLoginPage: AcpLoginPage = new AcpLoginPage();
  const accountConsentPage: AccountConsentPage = new AccountConsentPage();
  const consentSelfServicePage: ConsentSelfServicePage = new ConsentSelfServicePage();
  const consentSelfServiceAccountDetailsPage: ConsentSelfServiceAccountDetailsPage = new ConsentSelfServiceAccountDetailsPage();
  const consentSelfServiceApplicationPage: ConsentSelfServiceApplicationPage = new ConsentSelfServiceApplicationPage();
  const financrooLoginPage: FinancrooLoginPage = new FinancrooLoginPage();
  const financrooWelcomePage: FinancrooWelcomePage = new FinancrooWelcomePage();
  const financrooModalPage: FinancrooModalPage = new FinancrooModalPage();
  const financrooAccountsPage: FinancrooAccountsPage = new FinancrooAccountsPage();


  before(() => {
    financrooLoginPage.visit();
    financrooLoginPage.login();

    financrooWelcomePage.reconnectGoBank();

    acpLoginPage.assertThatModalIsDisplayed();
    acpLoginPage.loginWithMfaOption();
    
    accountConsentPage.checkAllAccounts();
    accountConsentPage.clickAgree();

    financrooModalPage.assertThatModalIsDisplayed();

    financrooLoginPage.visit();
    financrooLoginPage.login();

    financrooAccountsPage.assertThatPageIsDisplayed();
  });

  beforeEach(() => {
    consentSelfServicePage.visit(true);

    acpLoginPage.assertThatModalIsDisplayed();
    acpLoginPage.login();

    consentSelfServicePage.clickOnAccountOnlyButton();
    consentSelfServicePage.clickOnApplicationCard();
  });

  it(`Happy path with account consent`, () => {
    consentSelfServiceApplicationPage.expandAccountsTab();
    consentSelfServiceApplicationPage.checkAccount(Accounts.ids.BR.account1);
    consentSelfServiceApplicationPage.expandAccountConsentRow();

    consentSelfServiceAccountDetailsPage.assertThatAccountDetailsAreVisible()
    consentSelfServiceAccountDetailsPage.assertAccount(Accounts.ids.BR.account1);
  });

  it(`Revoke consent`, () => {
    consentSelfServiceApplicationPage.expandAccountsTab();
    consentSelfServiceApplicationPage.assertNumberOfConsents(1);
    consentSelfServiceApplicationPage.expandAccountConsentRow();

    consentSelfServiceAccountDetailsPage.assertThatAccountDetailsAreVisible();
    consentSelfServiceAccountDetailsPage.clickRevokeAccessButton();
    consentSelfServiceAccountDetailsPage.assertThatRevokeAccountDetailsAreVisible();
    consentSelfServiceAccountDetailsPage.confirmRevokeAccessAction();
    
    consentSelfServiceApplicationPage.assertNumberOfConsents(0);
  });
});
