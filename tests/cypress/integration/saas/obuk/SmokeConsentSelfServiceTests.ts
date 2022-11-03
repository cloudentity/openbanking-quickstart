import { AcpLoginPage } from "../../../pages/acp/AcpLoginPage";
import { AccountConsentPage } from "../../../pages/consent/AccountConsentPage";
import { PaymentConsentPage } from "../../../pages/consent/PaymentConsentPage";import { Credentials } from "../../../pages/Credentials";
import { ConsentSelfServicePage } from "../../../pages/consent-self-service/ConsentSelfServicePage";
import { ConsentSelfServicePaymentDetailsPage } from "../../../pages/consent-self-service/ConsentSelfServicePaymentDetailsPage";
import { ConsentSelfServiceAccountDetailsPage } from "../../../pages/consent-self-service/ConsentSelfServiceAccountDetailsPage";
import { Currencies } from "../../../pages/Currencies";
import { Accounts } from "../../../pages/Accounts";
import { FinancrooLoginPage } from "../../../pages/financroo/FinancrooLoginPage";
import { FinancrooWelcomePage } from "../../../pages/financroo/FinancrooWelcomePage";
import { FinancrooAccountsPage } from "../../../pages/financroo/accounts/FinancrooAccountsPage";
import { FinancrooInvestmentsPage } from "../../../pages/financroo/investments/FinancrooInvestmentsPage";
import { FinancrooContributePage } from "../../../pages/financroo/investments/FinancrooContributePage";
import { ConsentSelfServiceApplicationPage } from "../../../pages/consent-self-service/ConsentSelfServiceApplicationPage";
import { FinancrooModalPage } from "../../../pages/financroo/accounts/FinancrooModalPage";

describe(`Smoke Consent self service app`, () => {
  const acpLoginPage: AcpLoginPage = new AcpLoginPage();
  const accountConsentPage: AccountConsentPage = new AccountConsentPage();
  const paymentConsentPage: PaymentConsentPage = new PaymentConsentPage();
  const consentSelfServicePage: ConsentSelfServicePage = new ConsentSelfServicePage();
  const consentSelfServicePaymentDetailsPage: ConsentSelfServicePaymentDetailsPage = new ConsentSelfServicePaymentDetailsPage();
  const consentSelfServiceAccountDetailsPage: ConsentSelfServiceAccountDetailsPage = new ConsentSelfServiceAccountDetailsPage();
  const consentSelfServiceApplicationPage: ConsentSelfServiceApplicationPage = new ConsentSelfServiceApplicationPage();
  const financrooLoginPage: FinancrooLoginPage = new FinancrooLoginPage();
  const financrooWelcomePage: FinancrooWelcomePage = new FinancrooWelcomePage();
  const financrooModalPage: FinancrooModalPage = new FinancrooModalPage();
  const financrooAccountsPage: FinancrooAccountsPage = new FinancrooAccountsPage();
  const financrooInvestmentsPage: FinancrooInvestmentsPage =  new FinancrooInvestmentsPage();
  const financrooContributePage: FinancrooContributePage = new FinancrooContributePage();

  const amount: number = Math.floor(Math.random() * 50) + 1;

  before(() => {
    financrooLoginPage.visit();
    financrooLoginPage.login();

    financrooWelcomePage.reconnectGoBank();

    acpLoginPage.assertThatModalIsDisplayed("XXX");
    acpLoginPage.loginWithMfaOption();
    
    accountConsentPage.checkAllAccounts();
    accountConsentPage.clickAgree();

    financrooModalPage.assertThatModalIsDisplayed();

    financrooLoginPage.visit();

    financrooAccountsPage.assertThatPageIsDisplayed();
    financrooAccountsPage.goToInvestmentsTab();

    financrooInvestmentsPage.assertThatDashboardIsVisible(Currencies.currency.UK.code);
    financrooInvestmentsPage.clickInvest();

    financrooContributePage.contributeAmount(amount, Currencies.currency.UK.symbol);
    financrooContributePage.contributePaymentMethod(amount, Currencies.currency.UK.symbol, Accounts.ids.UK.bills);
    financrooContributePage.contributeInvestmentSummary(amount, Currencies.currency.UK.symbol, Accounts.ids.UK.bills);

    acpLoginPage.assertThatModalIsDisplayed("XXX");
    acpLoginPage.loginWithMfaOption();

    paymentConsentPage.assertThatConsentPageIsVisible(amount, Currencies.currency.UK.code, Accounts.ids.UK.bills); 
    paymentConsentPage.clickConfirm();

    financrooInvestmentsPage.assertThatTransactionWasCompleted(amount, Currencies.currency.UK.symbol);
  });

  beforeEach(() => {
    consentSelfServicePage.visit(true);
  });

  it(`Happy path with account consent`, () => {
    acpLoginPage.assertThatModalIsDisplayed("XXX");
    acpLoginPage.login();

    consentSelfServicePage.clickOnApplicationCard();

    consentSelfServiceApplicationPage.expandAccountsTab();
    consentSelfServiceApplicationPage.checkAccount(Accounts.ids.UK.bills);
    consentSelfServiceApplicationPage.expandAccountConsentRow();

    consentSelfServiceAccountDetailsPage.assertThatAccountDetailsAreVisible()
    consentSelfServiceAccountDetailsPage.assertAccount(Accounts.ids.UK.bills);
  });

  it(`Revoke consent`, () => {
    acpLoginPage.assertThatModalIsDisplayed("XXX");
    acpLoginPage.login();

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

  it(`Happy path with payment consent`, () => {
    acpLoginPage.assertThatModalIsDisplayed("XXX");
    acpLoginPage.login();

    consentSelfServicePage.clickOnApplicationCard();

    consentSelfServiceApplicationPage.expandPaymentsTab();
    consentSelfServiceApplicationPage.checkAccount(Accounts.ids.UK.bills);
    consentSelfServiceApplicationPage.checkAmount(Currencies.currency.UK.symbol, amount);
    consentSelfServiceApplicationPage.expandPaymentConsentRow();

    consentSelfServicePaymentDetailsPage.assertThatPaymentDetailsAreVisible();
    consentSelfServicePaymentDetailsPage.assertAmount(Currencies.currency.UK.symbol, amount);
    consentSelfServicePaymentDetailsPage.assertAccount(Accounts.ids.UK.bills);
  });
});
