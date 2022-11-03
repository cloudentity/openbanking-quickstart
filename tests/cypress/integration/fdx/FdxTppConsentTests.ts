import { AcpLoginPage } from "../../pages/acp/AcpLoginPage";
import { AccountConsentPage } from "../../pages/consent/AccountConsentPage";
import { Credentials } from "../../pages/Credentials";
import { Urls } from "../../pages/Urls";
import { Accounts } from "../../pages/Accounts";
import { MfaPage } from "../../pages/mfa/MfaPage";
import { EnvironmentVariables } from "../../pages/EnvironmentVariables";
import { FdxTppLandingPage } from "../../pages/fdx-tpp/FdxTppLandingPage";
import { FdxTppIntentRegisteredPage } from "../../pages/fdx-tpp/FdxTppIntentRegisteredPage";
import { FdxTppAuthenticatedPage } from "../../pages/fdx-tpp/FdxTppAuthenticatedPage";
import {ErrorPage} from '../../pages/ErrorPage';

describe(`FDX Tpp consent app`, () => {
  const fdxTppLoginPage: FdxTppLandingPage = new FdxTppLandingPage();
  const fdxTppIntentRegisteredPage: FdxTppIntentRegisteredPage = new FdxTppIntentRegisteredPage();
  const fdxTppAuthenticatedPage: FdxTppAuthenticatedPage = new FdxTppAuthenticatedPage();
  const acpLoginPage: AcpLoginPage = new AcpLoginPage();
  const consentPage: AccountConsentPage = new AccountConsentPage();
  const mfaPage: MfaPage = new MfaPage();
  const environmentVariables: EnvironmentVariables = new EnvironmentVariables();
  const errorPage: ErrorPage = new ErrorPage();


  beforeEach(() => {
    fdxTppLoginPage.visit();
    Urls.clearLocalStorage();
    fdxTppLoginPage.visit();
  });

  [
    [Accounts.ids.FDX.checkingAcc, Accounts.ids.FDX.savings2],
    [Accounts.ids.FDX.savings1],
    [Accounts.ids.FDX.savings2],
  ].forEach((accountsIds) => {
    it(`Happy path with selected accounts: ${accountsIds}`, () => {
      fdxTppLoginPage.assertThatPageIsDisplayed();
      fdxTppLoginPage.clickNext();

      fdxTppIntentRegisteredPage.assertThatPageIsDisplayed();
      fdxTppIntentRegisteredPage.clickLogin();

      acpLoginPage.login(Credentials.defaultUsername, Credentials.defaultPassword);
      if (environmentVariables.isMfaEnabled()) {
        mfaPage.typePin();
      }

      consentPage.assertPermissions(4);
      consentPage.assertThatAccountsAreNotVisible(accountsIds);
      consentPage.clickContinue();
      consentPage.checkAccounts(accountsIds);
      consentPage.clickAgree();

      fdxTppAuthenticatedPage.assertThatPageIsDisplayed();
      fdxTppAuthenticatedPage.assertThatTokenResponseFieldIsNotEmpty();
      fdxTppAuthenticatedPage.assertThatAccessTokenFieldIsNotEmpty();
      fdxTppAuthenticatedPage.assertThatConsentResponseFieldContainsAccountsIds(
        accountsIds
      );

      fdxTppAuthenticatedPage.clickTryNext();
      fdxTppLoginPage.assertThatPageIsDisplayed();
    });
  });

  it(`Happy path with not selected account`, () => {
    fdxTppLoginPage.assertThatPageIsDisplayed();
    fdxTppLoginPage.clickNext();

    fdxTppIntentRegisteredPage.assertThatPageIsDisplayed();
    fdxTppIntentRegisteredPage.clickLogin();

    acpLoginPage.login(Credentials.defaultUsername, Credentials.defaultPassword);
    if (environmentVariables.isMfaEnabled()) {
      mfaPage.typePin();
    }

    consentPage.clickContinue();
    consentPage.clickAgree();

    fdxTppAuthenticatedPage.assertThatPageIsDisplayed();
    fdxTppAuthenticatedPage.assertThatConsentResponseFieldNotContainsAccountsIds(
      [Accounts.ids.FDX.checkingAcc, Accounts.ids.FDX.savings1, Accounts.ids.FDX.savings1]
    );
  });

  it('Cancel on consent page', () => {
    fdxTppLoginPage.assertThatPageIsDisplayed();
    fdxTppLoginPage.clickNext();

    fdxTppIntentRegisteredPage.assertThatPageIsDisplayed();
    fdxTppIntentRegisteredPage.clickLogin();

    acpLoginPage.login(Credentials.defaultUsername, Credentials.defaultPassword);
    if (environmentVariables.isMfaEnabled()) {
      mfaPage.typePin();
    }

    consentPage.assertPermissions(4);
    consentPage.clickCancel();

    // UI error page improvements AUT-5845
    errorPage.assertError(`acp returned an error: rejected: `);
  })

  it('Cancel on ACP login', () => {
    fdxTppLoginPage.assertThatPageIsDisplayed();
    fdxTppLoginPage.clickNext();

    fdxTppIntentRegisteredPage.assertThatPageIsDisplayed();
    fdxTppIntentRegisteredPage.clickLogin();

    acpLoginPage.cancelLogin();
    // UI error page improvements AUT-5845
    errorPage.assertError(`The user rejected the authentication`);
  })
});
