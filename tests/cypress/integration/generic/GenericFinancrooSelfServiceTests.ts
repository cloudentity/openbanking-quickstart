import {FinancrooLoginPage} from "../../pages/financroo/FinancrooLoginPage";
import {Accounts} from "../../pages/Accounts";
import {FinancrooWelcomePage} from "../../pages/financroo/FinancrooWelcomePage";
import {AcpLoginPage} from "../../pages/acp/AcpLoginPage";
import {AccountConsentPage} from "../../pages/consent/AccountConsentPage";
import {FinancrooAccountsPage} from "../../pages/financroo/accounts/FinancrooAccountsPage";
import {FinancrooConnectAccountPage} from "../../pages/financroo/accounts/FinancrooConnectAccountPage";

describe(`Generic Financroo Consent self service tests`, () => {
  const financrooLoginPage: FinancrooLoginPage = new FinancrooLoginPage();
  const financrooWelcomePage: FinancrooWelcomePage = new FinancrooWelcomePage();
  const acpLoginPage: AcpLoginPage = new AcpLoginPage();
  const accountConsentPage: AccountConsentPage = new AccountConsentPage();
  const financrooAccountsPage: FinancrooAccountsPage = new FinancrooAccountsPage();
  const financrooConnectAccountPage: FinancrooConnectAccountPage = new FinancrooConnectAccountPage();

  const gobankAccountsIDs = [Accounts.ids.GENERIC_GOBANK.checkings,
    Accounts.ids.GENERIC_GOBANK.loan, Accounts.ids.GENERIC_GOBANK.savings];
  const hyperscaleAccountsIDs = [Accounts.ids.GENERIC_HYPERSCALE.checkings,
    Accounts.ids.GENERIC_HYPERSCALE.loan, Accounts.ids.GENERIC_HYPERSCALE.savings];

  beforeEach(() => {
    financrooLoginPage.visit();
    financrooLoginPage.login();
  });

  it(`Happy path with GO bank`, () => {
    financrooWelcomePage.reconnectGoBank();

    acpLoginPage.assertThatModalIsDisplayed();
    acpLoginPage.loginWithMfaOption();

    accountConsentPage.checkAccounts(gobankAccountsIDs);
    accountConsentPage.clickAgree();

    financrooAccountsPage.assertThatPageIsDisplayed();
    financrooAccountsPage.assertAccountsSyncedNumber(gobankAccountsIDs.length);
    financrooAccountsPage.assertAccountsIds(gobankAccountsIDs);
  });

  it(`Happy path with Hyperscale bank`, () => {
    financrooWelcomePage.reconnectHyperscaleBank();

    acpLoginPage.assertThatModalIsDisplayed();
    acpLoginPage.loginWithMfaOption();

    accountConsentPage.checkAccounts(hyperscaleAccountsIDs);
    accountConsentPage.clickAgree();

    financrooAccountsPage.assertThatPageIsDisplayed();
    financrooAccountsPage.assertAccountsSyncedNumber(hyperscaleAccountsIDs.length);
    financrooAccountsPage.assertAccountsIds(hyperscaleAccountsIDs);
  });

  it(`Happy path with both banks`, () => {
    financrooWelcomePage.reconnectHyperscaleBank();

    acpLoginPage.assertThatModalIsDisplayed();
    acpLoginPage.loginWithMfaOption();

    accountConsentPage.checkAccounts(hyperscaleAccountsIDs);
    accountConsentPage.clickAgree();

    financrooAccountsPage.assertThatPageIsDisplayed();
    financrooAccountsPage.addNewBankAccount();

    financrooConnectAccountPage.clickGoBankIcon()
    financrooConnectAccountPage.allow()

    acpLoginPage.assertThatModalIsDisplayed();
    acpLoginPage.loginWithMfaOption();

    accountConsentPage.checkAccounts(gobankAccountsIDs);
    accountConsentPage.clickAgree();

    financrooAccountsPage.assertThatPageIsDisplayed();
    financrooAccountsPage.assertAccountsIds(hyperscaleAccountsIDs);
    financrooAccountsPage.assertAccountsIds(gobankAccountsIDs);
  });

  it(`Happy path with both banks reversed`, () => {
    financrooWelcomePage.reconnectGoBank();

    acpLoginPage.assertThatModalIsDisplayed();
    acpLoginPage.loginWithMfaOption();

    accountConsentPage.checkAccounts(gobankAccountsIDs);
    accountConsentPage.clickAgree();

    financrooAccountsPage.assertThatPageIsDisplayed();
    financrooAccountsPage.addNewBankAccount();

    financrooConnectAccountPage.clickHyperscaleBankIcon()
    financrooConnectAccountPage.allow()

    acpLoginPage.assertThatModalIsDisplayed();
    acpLoginPage.loginWithMfaOption();

    accountConsentPage.checkAccounts(hyperscaleAccountsIDs);
    accountConsentPage.clickAgree();

    financrooAccountsPage.assertThatPageIsDisplayed();
    financrooAccountsPage.assertAccountsIds(hyperscaleAccountsIDs);
    financrooAccountsPage.assertAccountsIds(gobankAccountsIDs);
  });

});
