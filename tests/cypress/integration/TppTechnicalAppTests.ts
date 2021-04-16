import {TppAuthenticatedPage} from '../pages/tpp/TppAuthenticatedPage';
import {TppIntentPage} from '../pages/tpp/TppIntentPage';
import {TppLoginPage} from '../pages/tpp/TppLoginPage';
import {AcpLoginPage} from '../pages/acp/AcpLoginPage';
import {ConsentPage} from '../pages/consent/ConsentPage';
import {ErrorPage} from '../pages/ErrorPage';
import {Credentials} from "../pages/Credentials";
import {Urls} from "../pages/Urls";

describe(`Tpp technical app`, () => {
  const tppAuthenticatedPage: TppAuthenticatedPage = new TppAuthenticatedPage();
  const tppIntentPage: TppIntentPage = new TppIntentPage();
  const tppLoginPage: TppLoginPage = new TppLoginPage();
  const acpLoginPage: AcpLoginPage = new AcpLoginPage();
  const consentPage: ConsentPage = new ConsentPage();
  const errorPage: ErrorPage = new ErrorPage();

  const basicPermission: string = `ReadAccountsBasic`;
  const detailPermission: string = `ReadAccountsDetail`;

  beforeEach(() => {
    tppLoginPage.visit()
    Urls.clearLocalStorage()
    tppLoginPage.visit();
  });

  [
    [basicPermission, detailPermission],
    [basicPermission],
    [detailPermission],
    []
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
        consentPage.assertPermissions(permissions)
        consentPage.confirm();
        if (!permissions.includes(basicPermission) && permissions.includes(detailPermission)) {
          errorPage.assertError(`failed to call bank get accounts`)
        } else {
          tppAuthenticatedPage.assertSuccess()
        }
      }
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
    consentPage.cancel()
    errorPage.assertError(`rejected`)
  })

})
