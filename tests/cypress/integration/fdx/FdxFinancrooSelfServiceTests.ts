import { FinancrooLoginPage } from "../../pages/financroo/FinancrooLoginPage";
import { Accounts } from "../../pages/Accounts";
import { FinancrooWelcomePage } from "../../pages/financroo/FinancrooWelcomePage";
import { AcpLoginPage } from "../../pages/acp/AcpLoginPage";
import { AccountConsentPage } from "../../pages/consent/AccountConsentPage";
import { ConsentSelfServicePage } from "../../pages/consent-self-service/ConsentSelfServicePage";
import { FinancrooModalPage } from "../../pages/financroo/accounts/FinancrooModalPage";
import { FinancrooAccountsPage } from "../../pages/financroo/accounts/FinancrooAccountsPage";
import { ConsentSelfServiceApplicationPage } from "../../pages/consent-self-service/ConsentSelfServiceApplicationPage";
import { ConsentSelfServiceAccountDetailsPage } from "../../pages/consent-self-service/ConsentSelfServiceAccountDetailsPage";

describe(`FDX Financroo Consent self service tests`, () => {
  const financrooLoginPage: FinancrooLoginPage = new FinancrooLoginPage();
  const financrooWelcomePage: FinancrooWelcomePage = new FinancrooWelcomePage();
  const acpLoginPage: AcpLoginPage = new AcpLoginPage();
  const accountConsentPage: AccountConsentPage = new AccountConsentPage();
  const financrooModalPage: FinancrooModalPage = new FinancrooModalPage();
  const financrooAccountsPage: FinancrooAccountsPage = new FinancrooAccountsPage();
  const consentSelfServicePage: ConsentSelfServicePage = new ConsentSelfServicePage();
  const consentSelfServiceApplicationPage: ConsentSelfServiceApplicationPage = new ConsentSelfServiceApplicationPage();
  const consentSelfServiceAccountDetailsPage: ConsentSelfServiceAccountDetailsPage = new ConsentSelfServiceAccountDetailsPage();

  const accountsIDs = [Accounts.ids.FDX.checkingAcc, Accounts.ids.FDX.savings1];


  before(() => {
    financrooLoginPage.visit();
    financrooLoginPage.login();

    financrooWelcomePage.reconnectGoBank();

    acpLoginPage.assertThatModalIsDisplayed();
    acpLoginPage.loginWithMfaOption();

    accountConsentPage.assertPermissions(4);
    accountConsentPage.assertThatAccountsAreNotVisible([
      Accounts.ids.FDX.checkingAcc,
      Accounts.ids.FDX.savings1,
      Accounts.ids.FDX.savings2,
    ]);

    accountConsentPage.clickContinue();
    accountConsentPage.checkAccounts(accountsIDs);
    accountConsentPage.clickAgree();

    financrooModalPage.assertThatModalIsDisplayed();
    financrooModalPage.close();

    financrooAccountsPage.assertThatPageIsDisplayed();
    financrooAccountsPage.assertAccountsSyncedNumber(accountsIDs.length);
    financrooAccountsPage.assertAccountsIds(accountsIDs);
  });

  beforeEach(() => {
    consentSelfServicePage.visit(true);

    acpLoginPage.assertThatModalIsDisplayed();
    acpLoginPage.login();

    consentSelfServicePage.clickOnApplicationCardWithName("Financroo");
  });

  it(`Happy path with account consent`, () => {
    consentSelfServiceApplicationPage.expandAccountsTab();
    consentSelfServiceApplicationPage.checkAccountHasStatus(accountsIDs[0], "Authorised");
    consentSelfServiceApplicationPage.checkAccountHasStatus(accountsIDs[1], "Authorised");
    consentSelfServiceApplicationPage.expandAccountConsentRow();

    consentSelfServiceAccountDetailsPage.assertThatAccountDetailsAreVisible()
    consentSelfServiceAccountDetailsPage.assertAccount(accountsIDs[0]);
    consentSelfServiceAccountDetailsPage.assertAccount(accountsIDs[1]);
  });

  it(`Revoke account consent`, () => {
    consentSelfServiceApplicationPage.expandAccountsTab();
    consentSelfServiceApplicationPage.assertAuthorisedAccountRowExists(accountsIDs[0]);
    consentSelfServiceApplicationPage.assertAuthorisedAccountRowExists(accountsIDs[1]);
    consentSelfServiceApplicationPage.expandAccountConsentRow();

    consentSelfServiceAccountDetailsPage.assertThatAccountDetailsAreVisible();
    consentSelfServiceAccountDetailsPage.clickRevokeAccessButton();
    consentSelfServiceAccountDetailsPage.assertThatRevokeAccountDetailsAreVisible();
    consentSelfServiceAccountDetailsPage.confirmRevokeAccessAction();

    consentSelfServicePage.assertThatNoAccountsPageIsDisplayed();
  });
});
