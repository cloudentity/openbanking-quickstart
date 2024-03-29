import { FinancrooLoginPage } from "../../pages/financroo/FinancrooLoginPage";
import { FinancrooWelcomePage } from "../../pages/financroo/FinancrooWelcomePage";
import { AcpLoginPage } from "../../pages/acp/AcpLoginPage";
import { AccountConsentPage } from "../../pages/consent/AccountConsentPage";
import { FinancrooModalPage } from "../../pages/financroo/accounts/FinancrooModalPage";
import { FinancrooAccountsPage } from "../../pages/financroo/accounts/FinancrooAccountsPage";
import { Accounts } from "../../pages/Accounts";
import { ConsentAdminPage } from "../../pages/consent-admin/ConsentAdminPage";


describe(`Financroo Consent admin portal tests`, () => {
  const financrooLoginPage: FinancrooLoginPage = new FinancrooLoginPage();
  const financrooWelcomePage: FinancrooWelcomePage = new FinancrooWelcomePage();
  const acpLoginPage: AcpLoginPage = new AcpLoginPage();
  const accountConsentPage: AccountConsentPage = new AccountConsentPage();
  const financrooModalPage: FinancrooModalPage = new FinancrooModalPage();
  const financrooAccountsPage: FinancrooAccountsPage = new FinancrooAccountsPage();
  const consentAdminPage: ConsentAdminPage = new ConsentAdminPage();

  beforeEach(() => {
    financrooLoginPage.visit();
    financrooLoginPage.login();

    financrooWelcomePage.reconnectGoBank();

    acpLoginPage.assertThatModalIsDisplayed();
    acpLoginPage.loginWithMfaOption();
  });

  it(`Happy path with revoking consent from Consent management page`, () => {
    const accountsIDs = [Accounts.ids.CDR.loan, Accounts.ids.CDR.checking];

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

    consentAdminPage.visit(true);
    consentAdminPage.login();

    consentAdminPage.assertThatConsentManagementTabIsDisplayed();
    consentAdminPage.searchAccount(Accounts.ids.CDR.loan);
    consentAdminPage.assertAccountResult(Accounts.ids.CDR.loan);
    consentAdminPage.assertClientAccountWithStatus("Financroo", "Active");
    consentAdminPage.manageAccount("Financroo");
    consentAdminPage.assertConsentsDetails();
    consentAdminPage.revokeClientConsentByAccountName("Financroo");
    consentAdminPage.assertClientAccountWithStatus("Financroo", "Inactive");
  });

  it(`Happy path with revoking consent from Third party providers page`, () => {
    const accountsIDs = [Accounts.ids.CDR.savings];

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

    consentAdminPage.visit(true);
    consentAdminPage.login();

    consentAdminPage.assertThatConsentManagementTabIsDisplayed();
    consentAdminPage.revokeClientConsent();
  });

  afterEach(() => {
    financrooLoginPage.visit();
    financrooLoginPage.login();

    financrooAccountsPage.assertThatAccountsAreDisconnected();
  });
});
