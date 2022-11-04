import { TppIntentPage } from "../../pages/tpp/TppIntentPage";
import { TppLoginPage } from "../../pages/tpp/TppLoginPage";
import { AcpLoginPage } from "../../pages/acp/AcpLoginPage";
import { AccountConsentPage } from "../../pages/consent/AccountConsentPage";
import { ErrorPage } from "../../pages/ErrorPage";
import { ConsentAdminPage } from "../../pages/consent-admin/ConsentAdminPage";
import { Accounts } from "../../pages/Accounts";

describe(`Consent admin app`, () => {
  const tppIntentPage: TppIntentPage = new TppIntentPage();
  const tppLoginPage: TppLoginPage = new TppLoginPage();
  const acpLoginPage: AcpLoginPage = new AcpLoginPage();
  const accountConsentPage: AccountConsentPage = new AccountConsentPage();
  const errorPage: ErrorPage = new ErrorPage();
  const consentAdminPage: ConsentAdminPage = new ConsentAdminPage();

  beforeEach(() => {
    tppLoginPage.visit(true);
    tppLoginPage.next();

    tppIntentPage.saveIntentId();
    tppIntentPage.login();
  });

  it(`Happy path with revoking consent from Third party providers page`, () => {
    acpLoginPage.assertThatModalIsDisplayed("Open Banking UK");
    acpLoginPage.loginWithMfaOption();

    accountConsentPage.checkAllAccounts();
    accountConsentPage.clickAgree();

    consentAdminPage.visit(true);
    consentAdminPage.login();

    consentAdminPage.assertThatConsentManagementTabIsDisplayed()
    consentAdminPage.revokeClientConsent();
  });

  it(`Happy path with revoking consent from Consent management page`, () => {
    acpLoginPage.assertThatModalIsDisplayed("Open Banking UK");
    acpLoginPage.loginWithMfaOption();

    accountConsentPage.checkAllAccounts();
    accountConsentPage.clickAgree();
    
    consentAdminPage.visit();
    consentAdminPage.login();

    consentAdminPage.assertThatConsentManagementTabIsDisplayed()
    consentAdminPage.searchAccount(Accounts.ids.UK.bills);
    consentAdminPage.assertAccountResult(Accounts.ids.UK.bills);
    consentAdminPage.assertClientAccountWithStatus("Developer TPP", "Active");
    consentAdminPage.manageAccount("Developer TPP");
    consentAdminPage.assertConsentsDetails();
    consentAdminPage.revokeClientConsentByAccountName("Developer TPP");
    consentAdminPage.assertClientAccountWithStatus("Developer TPP", "Inactive");
  })

  it(`Cancel first ACP login`, () => {
    acpLoginPage.assertThatModalIsDisplayed("Open Banking UK");
    acpLoginPage.cancelLogin();
    // UI error page improvements AUT-5845
    errorPage.assertError(`The user rejected the authentication`);
  });
});
