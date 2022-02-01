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

  // obuk 
  const basicPermission: string = `ReadAccountsBasic`;
  const detailPermission: string = `ReadAccountsDetail`;


  // obbr 
  const accountsReadPermission: string = `ACCOUNTS_READ`;
  const accountsOverdraftLimitsReadPermission: string = `ACCOUNTS_OVERDRAFT_LIMITS_READ`;
  const resourcesReadPermission: string = `RESOURCES_READ`;

  beforeEach(() => {
    tppLoginPage.visit()
    Urls.clearLocalStorage()
    tppLoginPage.visit();
  });

  if (environmentVariables.isOBBRSpecification()) {
    [
      // FIXME restore when this fix has been made
      // https://github.com/cloudentity/openbanking-quickstart/pull/108
      // [accountsReadPermission, accountsOverdraftLimitsReadPermission, resourcesReadPermission],
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
    });
  };

  if (environmentVariables.isOBUKSpecification()) {
    [
      [basicPermission, detailPermission],
      [basicPermission],
      [detailPermission]
      // [] // todo add better error handling in the app
    ].forEach(permissions => {
      it(`Happy path with permissions: ${permissions}`, () => {
        tppLoginPage.checkBasicPermission(permissions.includes(basicPermission))
        tppLoginPage.checkDetailPermission(permissions.includes(detailPermission))
        tppLoginPage.next();
        if (!permissions.includes(basicPermission) && !permissions.includes(detailPermission)) {
          errorPage.assertError(`Invalid consent request`)
        } else {
          tppIntentPage.login();
          acpLoginPage.login(Credentials.tppUsername, Credentials.defaultPassword);
          if (environmentVariables.isMfaEnabled()) {
            mfaPage.typePin()
          }
          consentPage.expandPermissions()
          consentPage.assertPermissions(permissions.length)
          consentPage.confirm();
          if (!permissions.includes(basicPermission) && permissions.includes(detailPermission)) {
            errorPage.assertError(`failed to call bank get accounts`)
          } else {
            tppAuthenticatedPage.assertSuccess()
          }
        }
      })
    });
  };

 

  it(`Cancel on ACP login`, () => {
    // FIXME restore when this fix has been made
    // https://github.com/cloudentity/openbanking-quickstart/pull/108
    // tppLoginPage.next();
    // tppIntentPage.login();
    // acpLoginPage.cancel();
    // errorPage.assertError(`The user rejected the authentication`)
  })

  it(`Cancel on consent`, () => {
    // FIXME restore when this fix has been made
    // https://github.com/cloudentity/openbanking-quickstart/pull/108
    // tppLoginPage.next();
    // tppIntentPage.login();
    // acpLoginPage.login(Credentials.tppUsername, Credentials.defaultPassword);
    // if (environmentVariables.isMfaEnabled()) {
    //   mfaPage.typePin()
    // }
    // consentPage.cancel()
    // errorPage.assertError(`rejected`)
  })

})
