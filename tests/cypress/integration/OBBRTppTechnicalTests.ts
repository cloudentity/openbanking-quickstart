import {TppAuthenticatedPage} from '../pages/tpp/TppAuthenticatedPage';
import {TppIntentPage} from '../pages/tpp/TppIntentPage';
import {TppLoginPage} from '../pages/tpp/TppLoginPage';
import {AcpLoginPage} from '../pages/acp/AcpLoginPage';
import {ConsentPage} from '../pages/consent/ConsentPage';
import {ErrorPage} from '../pages/ErrorPage';
import {Credentials} from "../pages/Credentials";
import {Urls} from "../pages/Urls";
import {MfaPage} from "../pages/mfa/MfaPage";
import {EnvironmentVariables} from "../pages/EnvironmentVariables"

describe(`Tpp technical app`, () => {
  const tppAuthenticatedPage: TppAuthenticatedPage = new TppAuthenticatedPage();
  const tppIntentPage: TppIntentPage = new TppIntentPage();
  const tppLoginPage: TppLoginPage = new TppLoginPage();
  const acpLoginPage: AcpLoginPage = new AcpLoginPage();
  const consentPage: ConsentPage = new ConsentPage();
  const errorPage: ErrorPage = new ErrorPage();
  const mfaPage: MfaPage = new MfaPage();
  const environmentVariables: EnvironmentVariables = new EnvironmentVariables();

  const accountsReadPermission: string = `ACCOUNTS_READ`;
  const accountsOverdraftLimitsReadPermission: string = `ACCOUNTS_OVERDRAFT_LIMITS_READ`;
  const resourcesReadPermission: string = `RESOURCES_READ`;


  beforeEach(() => {
    tppLoginPage.visit()
    Urls.clearLocalStorage()
    tppLoginPage.visit();
  });

  [
    [accountsReadPermission, accountsOverdraftLimitsReadPermission, resourcesReadPermission],
    [accountsReadPermission]
    // [] // todo add better error handling in the app
  ].forEach(permissions => {
    it(`Happy path with permissions: ${permissions}`, () => {
      tppLoginPage.checkAccountsReadPermission(permissions.includes(accountsReadPermission))
      tppLoginPage.checkAccountsOverdraftLimitsReadPermission(permissions.includes(accountsOverdraftLimitsReadPermission))
      tppLoginPage.checkResourcesReadPermission(permissions.includes(resourcesReadPermission))
      tppLoginPage.next();
      if (!permissions.includes(accountsReadPermission) || !permissions.includes(accountsOverdraftLimitsReadPermission) || !permissions.includes(resourcesReadPermission)) {
        errorPage.assertError(`failed to register account access consent`)
        return 
      } 
      tppIntentPage.login();
      acpLoginPage.login(Credentials.tppUsername, Credentials.defaultPassword);
      if (environmentVariables.isMfaEnabled()) {
        mfaPage.typePin()
      }
      // TODO: consent page needs work with obbr permissions 
      //consentPage.expandPermissions()
      //consentPage.assertPermissions(permissions.length)
      consentPage.confirm();
      tppAuthenticatedPage.assertSuccess()
    })
  })

  it(`Cancel on ACP login`, () => {
    tppLoginPage.next();
    tppIntentPage.login();
    acpLoginPage.cancel();
    errorPage.assertError(`The user rejected the authentication`)
  })

  it(`Cancel on consent`, () => {
    tppLoginPage.next();
    tppIntentPage.login();
    acpLoginPage.login(`user`, Credentials.defaultPassword);
    if (environmentVariables.isMfaEnabled()) {
      mfaPage.typePin()
    }
    consentPage.cancel()
    errorPage.assertError(`rejected`)
  })

})
