import { AcpLoginPage } from "../../../pages/acp/AcpLoginPage";
import { FinancrooLoginPage } from "../../../pages/financroo/FinancrooLoginPage";
import { AccountConsentPage } from "../../../pages/consent/AccountConsentPage";
import { PaymentConsentPage } from "../../../pages/consent/PaymentConsentPage";
import { FinancrooWelcomePage } from "../../../pages/financroo/FinancrooWelcomePage";
import { FinancrooAccountsPage } from "../../../pages/financroo/accounts/FinancrooAccountsPage";
import { Credentials } from "../../../pages/Credentials";
import { Currencies } from "../../../pages/Currencies";
import { Accounts } from "../../../pages/Accounts";
import { Urls } from "../../../pages/Urls";
import { MfaPage } from "../../../pages/mfa/MfaPage";
import { FinancrooInvestmentsPage } from "../../../pages/financroo/investments/FinancrooInvestmentsPage";
import { FinancrooContributePage } from "../../../pages/financroo/investments/FinancrooContributePage";
import { EnvironmentVariables } from "../../../pages/EnvironmentVariables";
import { FinancrooModalPage } from "../../../pages/financroo/accounts/FinancrooModalPage";

describe(`Smoke Financroo payments app test`, () => {
  const acpLoginPage: AcpLoginPage = new AcpLoginPage();
  const accountConsentPage: AccountConsentPage = new AccountConsentPage();
  const paymentConsentPage: PaymentConsentPage = new PaymentConsentPage();
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
    financrooLoginPage.visit();
    Urls.clearLocalStorage();
    financrooLoginPage.visit();
    financrooLoginPage.login();

    financrooWelcomePage.reconnectGoBank();

    acpLoginPage.login();
    if (environmentVariables.isMfaEnabled()) {
      mfaPage.typePin();
    }

    accountConsentPage.checkAllAccounts();
    accountConsentPage.clickAgree();

    financrooModalPage.assertThatModalIsDisplayed();
  });

  it(`Happy path with confirm consent to add new amount for account ${Accounts.ids.UK.bills}`, () => {
    financrooLoginPage.visit();
    financrooAccountsPage.assertThatPageIsDisplayed();
    financrooAccountsPage.goToInvestmentsTab();
    
    financrooInvestmentsPage.assertThatDashboardIsVisible(Currencies.currency.UK.code);
    financrooInvestmentsPage.clickInvest();

    financrooContributePage.contributeAmount(amount, Currencies.currency.UK.symbol);
    financrooContributePage.contributePaymentMethod(amount, Currencies.currency.UK.symbol, Accounts.ids.UK.bills);
    financrooContributePage.contributeInvestmentSummary(amount, Currencies.currency.UK.symbol, Accounts.ids.UK.bills);

    acpLoginPage.login();
    if (environmentVariables.isMfaEnabled()) {
      mfaPage.typePin();
    }
    
    paymentConsentPage.assertThatConsentPageIsVisible(amount, Currencies.currency.UK.code, Accounts.ids.UK.bills);  
    paymentConsentPage.clickConfirm();

    financrooInvestmentsPage.assertThatTransactionWasCompleted(amount, Currencies.currency.UK.symbol);
    financrooInvestmentsPage.clickBackToPortfolio();
    financrooInvestmentsPage.assertThatDashboardIsVisible(Currencies.currency.UK.code);
  });
});
