import { TppAuthenticatedPage } from "../../pages/tpp/TppAuthenticatedPage";
import { TppIntentPage } from "../../pages/tpp/TppIntentPage";
import { TppLoginPage } from "../../pages/tpp/TppLoginPage";
import { AcpLoginPage } from "../../pages/acp/AcpLoginPage";
import { AccountConsentPage } from "../../pages/consent/AccountConsentPage";
import { ErrorPage } from "../../pages/ErrorPage";

describe(`Tpp technical app`, () => {
  const tppAuthenticatedPage: TppAuthenticatedPage = new TppAuthenticatedPage();
  const tppIntentPage: TppIntentPage = new TppIntentPage();
  const tppLoginPage: TppLoginPage = new TppLoginPage();
  const acpLoginPage: AcpLoginPage = new AcpLoginPage();
  const accountConsentPage: AccountConsentPage = new AccountConsentPage();
  const errorPage: ErrorPage = new ErrorPage();

  const basicPermission: string = `ReadAccountsBasic`;
  const detailPermission: string = `ReadAccountsDetail`;

  beforeEach(() => {
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

        acpLoginPage.assertThatModalIsDisplayed("Open Banking UK");
        acpLoginPage.loginWithMfaOption();

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

    acpLoginPage.assertThatModalIsDisplayed("Open Banking UK");
    acpLoginPage.cancelLogin();
    // UI error page improvements AUT-5845
    errorPage.assertError(`The user rejected the authentication`);
  });

  it(`Cancel on consent`, () => {
    tppLoginPage.next();
    tppIntentPage.login();

    acpLoginPage.assertThatModalIsDisplayed("Open Banking UK");
    acpLoginPage.loginWithMfaOption();
    
    accountConsentPage.clickCancel();
    // UI error page improvements AUT-5845
    errorPage.assertError(`rejected`);
  });
});
