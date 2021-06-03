import {TppIntentPage} from '../pages/tpp/TppIntentPage';
import {TppLoginPage} from '../pages/tpp/TppLoginPage';
import {AcpLoginPage} from '../pages/acp/AcpLoginPage';
import {ConsentPage} from '../pages/consent/ConsentPage';
import {ErrorPage} from '../pages/ErrorPage';
import {Credentials} from "../pages/Credentials";
import {ConsentAdminPage} from "../pages/consent-admin/ConsentAdminPage";
import {Urls} from "../pages/Urls";
import {MfaPage} from "../pages/mfa/MfaPage";

describe(`Consent admin app`, () => {
  const tppIntentPage: TppIntentPage = new TppIntentPage();
  const tppLoginPage: TppLoginPage = new TppLoginPage();
  const acpLoginPage: AcpLoginPage = new AcpLoginPage();
  const consentPage: ConsentPage = new ConsentPage();
  const errorPage: ErrorPage = new ErrorPage();
  const consentAdminPage: ConsentAdminPage = new ConsentAdminPage();
  const mfaPage: MfaPage = new MfaPage();
  const enableMfa = "true" //process.env.ENABLE_MFA

  beforeEach(() => {
    consentAdminPage.visit()
    Urls.clearLocalStorage()
    tppLoginPage.visit(true);
    tppLoginPage.next();
    tppIntentPage.saveIntentId()
    tppIntentPage.login();
  });

  it(`Happy path with revoking consent`, () => {

    // START Logging
    Cypress.log({
      displayName: 'Environment variable from node types >>> ' + process.env.LOG_LEVEL
    })
    console.log(process.env.LOG_LEVEL);

    Cypress.log({
      displayName: "Environment variable 'makefile_mfa' set via Makefile >>> " + Cypress.env('makefile_mfa')
    })
    console.log(Cypress.env('makefile_mfa'))

    Cypress.log({
      displayName: "Environment variable 'test' set in cypress.json file >>> " + Cypress.env('test')
    })
    console.log(Cypress.env('test'))


    Cypress.log({
      displayName: "Environment variable 'dotEnvMfa' set in index.js file via dotenv >>> " + Cypress.env('dotEnvMfa')
    })
    console.log(Cypress.env('dotEnvMfa'))

    Cypress.log({
      displayName: "Environment variable 'LOG_LEVEL' set in index.js file via dotenv >>> " + Cypress.env('logLevel')
    })
    console.log(Cypress.env('logLevel'))
    // END Logging

    acpLoginPage.login(Credentials.tppUsername, Credentials.defaultPassword)
    if (Cypress.env('makefile_mfa')) {
      mfaPage.typePin()
    }
    consentPage.confirm();
    consentAdminPage.visit(true);
    acpLoginPage.login(Credentials.consentAdminUsername, Credentials.defaultPassword);
    consentAdminPage.expandTab()
    consentAdminPage.revokeConsent()
    consentAdminPage.expandTab()
    consentAdminPage.assertConsentIsNotDisplayed()
  })

  // it(`Happy path with revoking all consents`, () => {
  //   acpLoginPage.login(Credentials.tppUsername, Credentials.defaultPassword);
  //   mfaPage.typePin()
  //   consentPage.confirm();
  //   consentAdminPage.visit();
  //   acpLoginPage.login(Credentials.consentAdminUsername, Credentials.defaultPassword);
  //   consentAdminPage.expandTab()
  //   consentAdminPage.revokeAllConsents()
  //   consentAdminPage.expandTab()
  //   consentAdminPage.assertConsentIsNotDisplayed()
  // })
/*
  it(`Cancel first ACP login`, () => {
    acpLoginPage.cancel();
    errorPage.assertError(`The user rejected the authentication`)
  })

  it(`Cancel second ACP login`, () => {
    acpLoginPage.login(Credentials.tppUsername, Credentials.defaultPassword);
    if (enableMfa) {
      mfaPage.typePin()
    }
    consentPage.confirm();
    consentAdminPage.visit();
    acpLoginPage.cancel();
    errorPage.assertError(`The user rejected the authentication`)
  })
*/
})
