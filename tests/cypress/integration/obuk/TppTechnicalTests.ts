import { TppAuthenticatedPage } from "../../pages/tpp/TppAuthenticatedPage";
import { TppIntentPage } from "../../pages/tpp/TppIntentPage";
import { TppLoginPage } from "../../pages/tpp/TppLoginPage";
import { AcpLoginPage } from "../../pages/acp/AcpLoginPage";
import { AccountConsentPage } from "../../pages/consent/AccountConsentPage";
import { TppErrorPage } from "../../pages/TppErrorPage";
import { ErrorPage } from "../../pages/ErrorPage";

describe(`Tpp technical app`, () => {
  const tppAuthenticatedPage: TppAuthenticatedPage = new TppAuthenticatedPage();
  const tppIntentPage: TppIntentPage = new TppIntentPage();
  const tppLoginPage: TppLoginPage = new TppLoginPage();
  const acpLoginPage: AcpLoginPage = new AcpLoginPage();
  const accountConsentPage: AccountConsentPage = new AccountConsentPage();
  const tppErrorPage: TppErrorPage = new TppErrorPage();
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

        acpLoginPage.assertThatModalIsDisplayed();
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

    acpLoginPage.assertThatModalIsDisplayed();
    acpLoginPage.cancelLogin();
    
    tppErrorPage.assertThatCancelLoginErrorPageIsDisplayed(
      `access denied`,
      `The user rejected the authentication`
    );
  });

  it(`Cancel on consent`, () => {
    tppLoginPage.next();
    tppIntentPage.login();

    acpLoginPage.assertThatModalIsDisplayed();
    acpLoginPage.loginWithMfaOption();
    
    accountConsentPage.clickCancel();

    tppErrorPage.assertThatRejectConsentErrorPageIsDisplayed(
      `rejected`,
      `The user rejected the authentication.`,
      `consent_rejected`
    );
  });
});
