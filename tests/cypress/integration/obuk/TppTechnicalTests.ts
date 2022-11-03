import { TppAuthenticatedPage } from "../../pages/tpp/TppAuthenticatedPage";
import { TppIntentPage } from "../../pages/tpp/TppIntentPage";
import { TppLoginPage } from "../../pages/tpp/TppLoginPage";
import { AcpLoginPage } from "../../pages/acp/AcpLoginPage";
import { AccountConsentPage } from "../../pages/consent/AccountConsentPage";
import { ErrorPage } from "../../pages/ErrorPage";
import { Credentials } from "../../pages/Credentials";
import { Urls } from "../../pages/Urls";
import { MfaPage } from "../../pages/mfa/MfaPage";
import { EnvironmentVariables } from "../../pages/EnvironmentVariables";

describe(`Tpp technical app`, () => {
  const tppAuthenticatedPage: TppAuthenticatedPage = new TppAuthenticatedPage();
  const tppIntentPage: TppIntentPage = new TppIntentPage();
  const tppLoginPage: TppLoginPage = new TppLoginPage();
  const acpLoginPage: AcpLoginPage = new AcpLoginPage();
  const accountConsentPage: AccountConsentPage = new AccountConsentPage();
  const errorPage: ErrorPage = new ErrorPage();
  const mfaPage: MfaPage = new MfaPage();
  const environmentVariables: EnvironmentVariables = new EnvironmentVariables();

  const basicPermission: string = `ReadAccountsBasic`;
  const detailPermission: string = `ReadAccountsDetail`;

  beforeEach(() => {
    tppLoginPage.visit();
    Urls.clearLocalStorage();
    tppLoginPage.visit();
  });

  [
    [basicPermission, detailPermission],
    [basicPermission],
    [detailPermission],
    [], // none permissions selected - UI error page improvements AUT-5845
  ].forEach((permissions) => {
    it(`Happy path with permissions: ${permissions}`, () => {
      tppLoginPage.checkBasicPermission(permissions.includes(basicPermission));
      tppLoginPage.checkDetailPermission(
        permissions.includes(detailPermission)
      );
      tppLoginPage.next();
      if (
        !permissions.includes(basicPermission) &&
        !permissions.includes(detailPermission)
      ) {
        errorPage.assertError(`Invalid consent request`);
      } else {
        tppIntentPage.login();
        acpLoginPage.login(
          Credentials.defaultUsername,
          Credentials.defaultPassword
        );
        if (environmentVariables.isMfaEnabled()) {
          mfaPage.typePin();
        }
        accountConsentPage.expandPermissions();
        accountConsentPage.assertPermissions(permissions.length);
        accountConsentPage.clickAgree();
        if (
          !permissions.includes(basicPermission) &&
          permissions.includes(detailPermission)
        ) {
          // ReadAccountsDetail permission selected - UI error page improvements AUT-5845
          errorPage.assertError(`failed to call bank get accounts`);
        } else {
          tppAuthenticatedPage.assertSuccess();
        }
      }
    });
  });

  it(`Cancel on ACP login`, () => {
    tppLoginPage.next();
    tppIntentPage.login();
    acpLoginPage.cancelLogin();
    // UI error page improvements AUT-5845
    errorPage.assertError(`The user rejected the authentication`);
  });

  it(`Cancel on consent`, () => {
    tppLoginPage.next();
    tppIntentPage.login();
    acpLoginPage.login(Credentials.defaultUsername, Credentials.defaultPassword);
    if (environmentVariables.isMfaEnabled()) {
      mfaPage.typePin();
    }
    accountConsentPage.clickCancel();
    // UI error page improvements AUT-5845
    errorPage.assertError(`rejected`);
  });
});
