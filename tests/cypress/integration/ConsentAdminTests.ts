import { TppIntentPage } from "../pages/tpp/TppIntentPage";
import { TppLoginPage } from "../pages/tpp/TppLoginPage";
import { AcpLoginPage } from "../pages/acp/AcpLoginPage";
import { ConsentPage } from "../pages/consent/ConsentPage";
import { ErrorPage } from "../pages/ErrorPage";
import { Credentials } from "../pages/Credentials";
import { ConsentAdminPage } from "../pages/consent-admin/ConsentAdminPage";
import { Urls } from "../pages/Urls";
import { MfaPage } from "../pages/mfa/MfaPage";
import { EnvironmentVariables } from "../pages/EnvironmentVariables";

describe(`Consent admin app`, () => {
  const tppIntentPage: TppIntentPage = new TppIntentPage();
  const tppLoginPage: TppLoginPage = new TppLoginPage();
  const acpLoginPage: AcpLoginPage = new AcpLoginPage();
  const consentPage: ConsentPage = new ConsentPage();
  const errorPage: ErrorPage = new ErrorPage();
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
    acpLoginPage.login(Credentials.tppUsername, Credentials.defaultPassword);
    if (environmentVariables.isMfaEnabled()) {
      mfaPage.typePin();
    }
    consentPage.confirm();
    consentAdminPage.visit(true);
    acpLoginPage.login(
      Credentials.consentAdminUsername,
      Credentials.defaultPassword
    );
    // FIXME restore when this fix has been made
    // https://github.com/cloudentity/openbanking-quickstart/pull/108
    // consentAdminPage.assertThatConsentManagementTabIsDisplayed()
    // consentAdminPage.revokeClientConsent();
  });

  it(`Happy path with revoking consent from Consent management page`, () => {
    acpLoginPage.login(Credentials.tppUsername, Credentials.defaultPassword);
    if (environmentVariables.isMfaEnabled()) {
      mfaPage.typePin();
    }
    consentPage.confirm();
    consentAdminPage.visit();
    acpLoginPage.login(
      Credentials.consentAdminUsername,
      Credentials.defaultPassword
    );
    // FIXME restore when this fix has been made
    // https://github.com/cloudentity/openbanking-quickstart/pull/108
    // consentAdminPage.assertThatConsentManagementTabIsDisplayed()
    // consentAdminPage.searchAccount("22289");
    // consentAdminPage.assertAccountResult("22289");
    // consentAdminPage.assertClientAccountWithStatus("Financroo", "Active");
    // consentAdminPage.manageAccount("Financroo");
    // consentAdminPage.assertConsentsDetails();
    // consentAdminPage.revokeClientConsentByAccountName("Financroo");
    // consentAdminPage.assertClientAccountWithStatus("Financroo", "Inactive");
  })

  it(`Cancel first ACP login`, () => {
    acpLoginPage.cancel();
    errorPage.assertError(`The user rejected the authentication`);
  });

  it(`Cancel second ACP login`, () => {
    acpLoginPage.login(Credentials.tppUsername, Credentials.defaultPassword);
    if (environmentVariables.isMfaEnabled()) {
      mfaPage.typePin();
    }
    consentPage.confirm();
    consentAdminPage.visit();
    acpLoginPage.cancel();
    errorPage.assertError(`The user rejected the authentication`);
  });
});
