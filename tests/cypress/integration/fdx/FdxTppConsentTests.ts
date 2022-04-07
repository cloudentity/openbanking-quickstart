import {TppAuthenticatedPage} from '../../pages/tpp/TppAuthenticatedPage';
import {TppIntentPage} from '../../pages/tpp/TppIntentPage';
import {TppLoginPage} from '../../pages/tpp/TppLoginPage';
import {AcpLoginPage} from '../../pages/acp/AcpLoginPage';
import {ConsentPage} from '../../pages/consent/ConsentPage';
import {ErrorPage} from '../../pages/ErrorPage';
import {Credentials} from "../../pages/Credentials";
import {Urls} from "../../pages/Urls";
import {MfaPage} from "../../pages/mfa/MfaPage";
import {EnvironmentVariables} from "../../pages/EnvironmentVariables"
import { FdxTppLoginPage } from '../../pages/fdx-tpp/FdxTppLoginPage';
import { FdxTppIntentRegisteredPage } from '../../pages/fdx-tpp/FdxTppIntentRegisteredPage';

describe(`Tpp technical app`, () => {
  const fdxTppLoginPage: FdxTppLoginPage = new FdxTppLoginPage();
  const fdxTppIntentRegisteredPage: FdxTppIntentRegisteredPage = new FdxTppIntentRegisteredPage();

  const tppAuthenticatedPage: TppAuthenticatedPage = new TppAuthenticatedPage();
  const tppIntentPage: TppIntentPage = new TppIntentPage();
  const tppLoginPage: TppLoginPage = new TppLoginPage();
  const acpLoginPage: AcpLoginPage = new AcpLoginPage();
  const consentPage: ConsentPage = new ConsentPage();
  const errorPage: ErrorPage = new ErrorPage();
  const mfaPage: MfaPage = new MfaPage();
  const environmentVariables: EnvironmentVariables = new EnvironmentVariables();

  const basicPermission: string = `ReadAccountsBasic`;
  const detailPermission: string = `ReadAccountsDetail`;


  beforeEach(() => {
    fdxTppLoginPage.visit()
    Urls.clearLocalStorage()
    fdxTppLoginPage.visit();
  });

  it(`TEST`, () => {
    fdxTppLoginPage.assertThatPageIsDisplayed();
    fdxTppLoginPage.assertThatAuthorizationDetailsAreDispalyed()
    fdxTppLoginPage.clickNext();

    fdxTppIntentRegisteredPage.assertThatPageIsDisplayed();
    fdxTppIntentRegisteredPage.assertThatRequestUriIsDisplayed();
    fdxTppIntentRegisteredPage.clickLogin();

  })














    // [
    //   [basicPermission, detailPermission],
    //   [basicPermission],
    //   [detailPermission]
      
    // ].forEach(permissions => {
    //   it(`Happy path with permissions: ${permissions}`, () => {
    //     tppLoginPage.checkBasicPermission(permissions.includes(basicPermission))
    //     tppLoginPage.checkDetailPermission(permissions.includes(detailPermission))
    //     tppLoginPage.next();
    //     if (!permissions.includes(basicPermission) && !permissions.includes(detailPermission)) {
    //       errorPage.assertError(`Invalid consent request`)
    //     } else {
    //       tppIntentPage.login();
    //       acpLoginPage.login(Credentials.tppUsername, Credentials.defaultPassword);
    //       if (environmentVariables.isMfaEnabled()) {
    //         mfaPage.typePin()
    //       }
    //       consentPage.expandPermissions()
    //       consentPage.assertPermissions(permissions.length)
    //       consentPage.confirm();
    //       if (!permissions.includes(basicPermission) && permissions.includes(detailPermission)) {
    //         errorPage.assertError(`failed to call bank get accounts`)
    //       } else {
    //         tppAuthenticatedPage.assertSuccess()
    //       }
    //     }
    //   })
    // });

 

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
