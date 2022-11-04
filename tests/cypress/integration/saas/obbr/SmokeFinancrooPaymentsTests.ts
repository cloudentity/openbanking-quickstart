import { AcpLoginPage } from "../../../pages/acp/AcpLoginPage";
import { FinancrooLoginPage } from "../../../pages/financroo/FinancrooLoginPage";
import { AccountConsentPage } from "../../../pages/consent/AccountConsentPage";
import { PaymentConsentPage } from "../../../pages/consent/PaymentConsentPage";
import { ErrorPage } from "../../../pages/ErrorPage";
import { FinancrooWelcomePage } from "../../../pages/financroo/FinancrooWelcomePage";
import { FinancrooAccountsPage } from "../../../pages/financroo/accounts/FinancrooAccountsPage";
import { Currencies } from "../../../pages/Currencies";
import { Accounts } from "../../../pages/Accounts";
import { FinancrooInvestmentsPage } from "../../../pages/financroo/investments/FinancrooInvestmentsPage";
import { FinancrooContributePage } from "../../../pages/financroo/investments/FinancrooContributePage";
import { FinancrooModalPage } from "../../../pages/financroo/accounts/FinancrooModalPage";

describe(`Financroo payments app test`, () => {
  const acpLoginPage: AcpLoginPage = new AcpLoginPage();
  const accountConsentPage: AccountConsentPage = new AccountConsentPage();
  const paymentConsentPage: PaymentConsentPage = new PaymentConsentPage();
  const errorPage: ErrorPage = new ErrorPage();
  const financrooLoginPage: FinancrooLoginPage = new FinancrooLoginPage();
  const financrooWelcomePage: FinancrooWelcomePage = new FinancrooWelcomePage();
  const financrooModalPage: FinancrooModalPage = new FinancrooModalPage();
  const financrooAccountsPage: FinancrooAccountsPage = new FinancrooAccountsPage();
  const financrooInvestmentsPage: FinancrooInvestmentsPage = new FinancrooInvestmentsPage();
  const financrooContributePage: FinancrooContributePage = new FinancrooContributePage();

  
  beforeEach(() => {
    financrooLoginPage.visit();
    financrooLoginPage.login();

    financrooWelcomePage.reconnectGoBank();

    acpLoginPage.assertThatModalIsDisplayed("Open Finance Brazil");
    acpLoginPage.loginWithMfaOption();
    
    accountConsentPage.checkAllAccounts();
    accountConsentPage.clickAgree();

    financrooModalPage.assertThatModalIsDisplayed();
  });

  it(`Happy path with confirm consent to add new amount`, () => {
    const amount: number = Math.floor(Math.random() * 50) + 1;

    financrooLoginPage.visit();
    financrooLoginPage.login();

    financrooAccountsPage.assertThatPageIsDisplayed();
    financrooAccountsPage.goToInvestmentsTab();

    financrooInvestmentsPage.assertThatDashboardIsVisible(Currencies.currency.BR.code);
    financrooInvestmentsPage.clickInvest();

    financrooContributePage.contributeAmount(amount, Currencies.currency.BR.symbol);
    financrooContributePage.contributePaymentMethod(amount, Currencies.currency.BR.symbol, Accounts.ids.BR.account1);
    financrooContributePage.contributeInvestmentSummary(amount, Currencies.currency.BR.symbol, Accounts.ids.BR.account1);

    acpLoginPage.assertThatModalIsDisplayed("Open Finance Brazil");
    acpLoginPage.loginWithMfaOption();

    paymentConsentPage.assertThatConsentPageIsVisible(amount, Currencies.currency.BR.code, Accounts.ids.BR.account1);  
    paymentConsentPage.clickConfirm();

    financrooInvestmentsPage.assertThatTransactionWasCompleted(amount, Currencies.currency.BR.symbol);
    financrooInvestmentsPage.clickBackToPortfolio();
    financrooInvestmentsPage.assertThatDashboardIsVisible(Currencies.currency.BR.code);
  });

  it(`Reject path with decline consent to add new amount`, () => {
    const amount: number = Math.floor(Math.random() * 50) + 1;

    financrooLoginPage.visit();
    financrooLoginPage.login();
    
    financrooAccountsPage.assertThatPageIsDisplayed();
    financrooAccountsPage.goToInvestmentsTab();

    financrooInvestmentsPage.assertThatDashboardIsVisible(Currencies.currency.BR.code);
    financrooInvestmentsPage.clickInvest();

    financrooContributePage.contributeAmount(amount, Currencies.currency.BR.symbol);
    financrooContributePage.contributePaymentMethod(amount, Currencies.currency.BR.symbol, Accounts.ids.BR.account1);
    financrooContributePage.contributeInvestmentSummary(amount, Currencies.currency.BR.symbol, Accounts.ids.BR.account1);

    acpLoginPage.assertThatModalIsDisplayed("Open Finance Brazil");
    acpLoginPage.loginWithMfaOption();
    
    paymentConsentPage.assertThatConsentPageIsVisible(amount, Currencies.currency.BR.code, Accounts.ids.BR.account1); 
    paymentConsentPage.clickCancel();

    // UI error page improvements AUT-5845
    errorPage.assertError(`acp returned an error: access_denied: rejected`);
  });
});
