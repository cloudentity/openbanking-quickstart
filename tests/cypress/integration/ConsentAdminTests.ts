import {TppIntentPage} from '../pages/tpp/TppIntentPage';
import {TppLoginPage} from '../pages/tpp/TppLoginPage';
import {AcpLoginPage} from '../pages/acp/AcpLoginPage';
import {ConsentPage} from '../pages/consent/ConsentPage';
import {ErrorPage} from '../pages/ErrorPage';
import {Credentials} from "../pages/Credentials";
import {ConsentAdminPage} from "../pages/consent-admin/ConsentAdminPage";
import {Urls} from "../pages/Urls";

describe(`Consent admin app`, () => {
  const tppIntentPage: TppIntentPage = new TppIntentPage();
  const tppLoginPage: TppLoginPage = new TppLoginPage();
  const acpLoginPage: AcpLoginPage = new AcpLoginPage();
  const consentPage: ConsentPage = new ConsentPage();
  const errorPage: ErrorPage = new ErrorPage();
  const consentAdminPage: ConsentAdminPage = new ConsentAdminPage();

  beforeEach(() => {
    consentAdminPage.visit()
    Urls.clearLocalStorage()
    tppLoginPage.visit();
    tppLoginPage.next();
    tppIntentPage.saveIntentId()
    tppIntentPage.login();
  });

  it(`Happy path with revoking consent`, () => {
    acpLoginPage.login(Credentials.tppUsername, Credentials.defaultPassword);
    consentPage.confirm();
    consentAdminPage.visit();
    acpLoginPage.login(Credentials.consentAdminUsername, Credentials.defaultPassword);
    consentAdminPage.expandTab()
    consentAdminPage.revokeConsent()
    consentAdminPage.expandTab()
    consentAdminPage.assertConsentIsNotDisplayed()
  })

  it(`Happy path with revoking all consents`, () => {
    acpLoginPage.login(Credentials.tppUsername, Credentials.defaultPassword);
    consentPage.confirm();
    consentAdminPage.visit();
    acpLoginPage.login(Credentials.consentAdminUsername, Credentials.defaultPassword);
    consentAdminPage.expandTab()
    consentAdminPage.revokeAllConsents()
    consentAdminPage.expandTab()
    consentAdminPage.assertConsentIsNotDisplayed()
  })

  it(`Cancel first ACP login`, () => {
    acpLoginPage.cancel();
    errorPage.assertError("The user rejected the authentication")
  })

  it(`Cancel second ACP login`, () => {
    acpLoginPage.login(Credentials.tppUsername, Credentials.defaultPassword);
    consentPage.confirm();
    consentAdminPage.visit();
    acpLoginPage.cancel();
    errorPage.assertError("The user rejected the authentication")
  })

})
