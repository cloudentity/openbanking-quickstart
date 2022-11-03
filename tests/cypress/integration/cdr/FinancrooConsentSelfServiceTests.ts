import { AcpLoginPage } from "../../pages/acp/AcpLoginPage";
import { AccountConsentPage } from "../../pages/consent/AccountConsentPage";
import { ConsentSelfServicePage } from "../../pages/consent-self-service/ConsentSelfServicePage";
import { ConsentSelfServiceApplicationPage } from "../../pages/consent-self-service/ConsentSelfServiceApplicationPage";
import { ConsentSelfServiceAccountDetailsPage } from "../../pages/consent-self-service/ConsentSelfServiceAccountDetailsPage";
import { Accounts } from "../../pages/Accounts";
import { FinancrooLoginPage } from "../../pages/financroo/FinancrooLoginPage";
import { FinancrooWelcomePage } from "../../pages/financroo/FinancrooWelcomePage";
import { FinancrooModalPage } from "../../pages/financroo/accounts/FinancrooModalPage";
import { FinancrooAccountsPage } from "../../pages/financroo/accounts/FinancrooAccountsPage";


describe(`Financroo Consent self service tests`, () => {
  const financrooLoginPage: FinancrooLoginPage = new FinancrooLoginPage();
  const financrooWelcomePage: FinancrooWelcomePage = new FinancrooWelcomePage();
  const acpLoginPage: AcpLoginPage = new AcpLoginPage();
  const accountConsentPage: AccountConsentPage = new AccountConsentPage();
  const financrooModalPage: FinancrooModalPage = new FinancrooModalPage();
  const financrooAccountsPage: FinancrooAccountsPage = new FinancrooAccountsPage();
  const consentSelfServicePage: ConsentSelfServicePage = new ConsentSelfServicePage();
  const consentSelfServiceApplicationPage: ConsentSelfServiceApplicationPage = new ConsentSelfServiceApplicationPage();
  const consentSelfServiceAccountDetailsPage: ConsentSelfServiceAccountDetailsPage = new ConsentSelfServiceAccountDetailsPage();

  const accountsIDs = [Accounts.ids.CDR.savings, Accounts.ids.CDR.checking];


  before(() => {
    financrooLoginPage.visit();
    financrooLoginPage.login();

    financrooWelcomePage.reconnectGoBank();

    acpLoginPage.loginWithMfaOption();

    accountConsentPage.checkAccounts(accountsIDs);
    accountConsentPage.expandPermissions();
    accountConsentPage.assertPermissionsDetails(
      "Purpose for sharing data",
      "To uncover insights that can improve your financial well being."
    );
    accountConsentPage.clickAgree();

    financrooModalPage.assertThatModalIsDisplayed();
    financrooModalPage.close();

    financrooAccountsPage.assertThatPageIsDisplayed();
    financrooAccountsPage.assertAccountsSyncedNumber(accountsIDs.length);
    financrooAccountsPage.assertAccountsIds(accountsIDs);
  });

  beforeEach(() => {
    consentSelfServicePage.visit(true);

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

    consentSelfServicePage.clickOnApplicationCardWithName("Financroo");
    consentSelfServiceApplicationPage.assertAuthorisedAccountRowDoesNotExist(accountsIDs[0]);
    consentSelfServiceApplicationPage.assertAuthorisedAccountRowDoesNotExist(accountsIDs[1]);

    financrooLoginPage.visit();
    financrooLoginPage.login();

    financrooAccountsPage.assertThatAccountsAreDisconnected();
  });

});
