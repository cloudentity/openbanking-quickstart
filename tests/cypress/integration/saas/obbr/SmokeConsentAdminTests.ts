import { TppIntentPage } from "../../../pages/tpp/TppIntentPage";
import { TppLoginPage } from "../../../pages/tpp/TppLoginPage";
import { AcpLoginPage } from "../../../pages/acp/AcpLoginPage";
import { AccountConsentPage } from "../../../pages/consent/AccountConsentPage";
import { Credentials } from "../../../pages/Credentials";
import { ConsentAdminPage } from "../../../pages/consent-admin/ConsentAdminPage";
import { Urls } from "../../../pages/Urls";
import { Accounts } from "../../../pages/Accounts";
import { MfaPage } from "../../../pages/mfa/MfaPage";
import { EnvironmentVariables } from "../../../pages/EnvironmentVariables";

describe(`Consent admin app`, () => {
  const tppIntentPage: TppIntentPage = new TppIntentPage();
  const tppLoginPage: TppLoginPage = new TppLoginPage();
  const acpLoginPage: AcpLoginPage = new AcpLoginPage();
  const accountConsentPage: AccountConsentPage = new AccountConsentPage();
  const consentAdminPage: ConsentAdminPage = new ConsentAdminPage();
  const mfaPage: MfaPage = new MfaPage();
  const environmentVariables: EnvironmentVariables = new EnvironmentVariables();

  beforeEach(() => {
    consentAdminPage.visit();
    Urls.clearLocalStorage();
    tppLoginPage.visit(true);
    tppLoginPage.next();
    tppIntentPage.saveIntentId();
    tppIntentPage.login();
  });

  it(`Happy path with revoking consent from Third party providers page`, () => {
    acpLoginPage.login(Credentials.defaultUsername, Credentials.defaultPassword);
    if (environmentVariables.isMfaEnabled()) {
      mfaPage.typePin();
    }

    accountConsentPage.checkAllAccounts();
    accountConsentPage.clickAgree();

    consentAdminPage.visit(true);
    consentAdminPage.login();

    consentAdminPage.assertThatConsentManagementTabIsDisplayed();
    consentAdminPage.revokeClientConsent();
  });

  it(`Happy path with revoking consent from Consent management page`, () => {
    acpLoginPage.login(Credentials.defaultUsername, Credentials.defaultPassword);
    if (environmentVariables.isMfaEnabled()) {
      mfaPage.typePin();
    }

    accountConsentPage.checkAllAccounts();
    accountConsentPage.clickAgree();

    consentAdminPage.visit();
    consentAdminPage.login();

    consentAdminPage.assertThatConsentManagementTabIsDisplayed()
    consentAdminPage.searchAccount(Accounts.ids.BR.account1);
    consentAdminPage.assertAccountResult(Accounts.ids.BR.account1);
    consentAdminPage.assertClientAccountWithStatus("Developer", "Active");
    consentAdminPage.manageAccount("Developer");
    consentAdminPage.assertConsentsDetails();
    consentAdminPage.revokeClientConsentByAccountName("Developer");
    consentAdminPage.assertClientAccountWithStatus("Developer", "Inactive");
  })

});
