import {TppIntentPage} from '../pages/tpp/TppIntentPage';
import {TppLoginPage} from '../pages/tpp/TppLoginPage';
import {AcpLoginPage} from '../pages/acp/AcpLoginPage';
import {ConsentPage} from '../pages/consent/ConsentPage';
import {ErrorPage} from '../pages/ErrorPage';
import {Credentials} from "../pages/Credentials";
import {ConsentSelfServicePage} from '../pages/consent-self-service/ConsentSelfServicePage';
import {Urls} from "../pages/Urls";
import {MfaPage} from "../pages/mfa/MfaPage";

describe(`Consent self service app`, () => {
  const tppIntentPage: TppIntentPage = new TppIntentPage();
  const tppLoginPage: TppLoginPage = new TppLoginPage();
  const acpLoginPage: AcpLoginPage = new AcpLoginPage();
  const consentPage: ConsentPage = new ConsentPage();
  const errorPage: ErrorPage = new ErrorPage();
  const consentSelfServicePage: ConsentSelfServicePage = new ConsentSelfServicePage();
  const mfaPage: MfaPage = new MfaPage();

  beforeEach(() => {
    consentSelfServicePage.visit()
    Urls.clearLocalStorage()
    tppLoginPage.visit(true);
    tppLoginPage.next();
    tppIntentPage.saveIntentId()
    tppIntentPage.login();
  });

  it(`Happy path`, () => {
    acpLoginPage.login(Credentials.tppUsername, Credentials.defaultPassword);
    mfaPage.typePin()
    consentPage.confirm();
    consentSelfServicePage.visit(true);
    acpLoginPage.login(Credentials.tppUsername, Credentials.defaultPassword);

    // consentSelfServicePage.expandTab()
    // consentSelfServicePage.revokeConsent()
    // consentSelfServicePage.expandTab()
    // consentSelfServicePage.assertConsentIsNotDisplayed()
  })

  /* it(`Cancel first ACP login`, () => {
     acpLoginPage.cancel();
     errorPage.assertError("The user rejected the authentication")
   })

   it(`Cancel second ACP login`, () => {
     acpLoginPage.login(Credentials.tppUsername, Credentials.defaultPassword);
     consentPage.confirm();
     consentSelfServicePage.visit();
     acpLoginPage.cancel();
     errorPage.assertError("The user rejected the authentication")
   })*/

})
