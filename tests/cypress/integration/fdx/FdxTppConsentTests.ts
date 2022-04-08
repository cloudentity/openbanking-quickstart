import { AcpLoginPage } from "../../pages/acp/AcpLoginPage";
import { ConsentPage } from "../../pages/consent/ConsentPage";
import { Credentials } from "../../pages/Credentials";
import { Urls } from "../../pages/Urls";
import { MfaPage } from "../../pages/mfa/MfaPage";
import { EnvironmentVariables } from "../../pages/EnvironmentVariables";
import { FdxTppLoginPage } from "../../pages/fdx-tpp/FdxTppLoginPage";
import { FdxTppIntentRegisteredPage } from "../../pages/fdx-tpp/FdxTppIntentRegisteredPage";
import { FdxTppAuthenticatedPage } from "../../pages/fdx-tpp/FdxTppAuthenticatedPage";
import {ErrorPage} from '../../pages/ErrorPage';

describe(`FDX Tpp consent app`, () => {
  const fdxTppLoginPage: FdxTppLoginPage = new FdxTppLoginPage();
  const fdxTppIntentRegisteredPage: FdxTppIntentRegisteredPage =
    new FdxTppIntentRegisteredPage();
  const fdxTppAuthenticatedPage: FdxTppAuthenticatedPage =
    new FdxTppAuthenticatedPage();
  const acpLoginPage: AcpLoginPage = new AcpLoginPage();
  const consentPage: ConsentPage = new ConsentPage();
  const mfaPage: MfaPage = new MfaPage();
  const environmentVariables: EnvironmentVariables = new EnvironmentVariables();
  const errorPage: ErrorPage = new ErrorPage();

  const digitalBankingAccountId: string = `96534987`;
  const savingsAccountId: string = `1000001`;
  const savings2AccountId: string = `1000002`;

  beforeEach(() => {
    fdxTppLoginPage.visit();
    Urls.clearLocalStorage();
    fdxTppLoginPage.visit();
  });

  [
    [digitalBankingAccountId, savings2AccountId],
    [savingsAccountId],
    [savings2AccountId],
  ].forEach((accountsIds) => {
    it(`Happy path with selcted accounts: ${accountsIds}`, () => {
      fdxTppLoginPage.assertThatPageIsDisplayed();
      fdxTppLoginPage.assertThatAuthorizationDetailsAreDispalyed();
      fdxTppLoginPage.clickNext();

      fdxTppIntentRegisteredPage.assertThatPageIsDisplayed();
      fdxTppIntentRegisteredPage.assertThatRequestUriFieldsAreNotEmpty();
      fdxTppIntentRegisteredPage.clickLogin();

      acpLoginPage.login(Credentials.tppUsername, Credentials.defaultPassword);
      if (environmentVariables.isMfaEnabled()) {
        mfaPage.typePin();
      }

      consentPage.checkAccounts(accountsIds);
      consentPage.expandPermissions();
      consentPage.assertPermissions(4);
      consentPage.confirm();

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
    fdxTppLoginPage.assertThatAuthorizationDetailsAreDispalyed();
    fdxTppLoginPage.clickNext();

    fdxTppIntentRegisteredPage.assertThatPageIsDisplayed();
    fdxTppIntentRegisteredPage.assertThatRequestUriFieldsAreNotEmpty();
    fdxTppIntentRegisteredPage.clickLogin();

    acpLoginPage.login(Credentials.tppUsername, Credentials.defaultPassword);
    if (environmentVariables.isMfaEnabled()) {
      mfaPage.typePin();
    }

    consentPage.confirm();

    fdxTppAuthenticatedPage.assertThatPageIsDisplayed();
    fdxTppAuthenticatedPage.assertThatConsentResponseFieldNotContainsAccountsIds(
      [digitalBankingAccountId, savingsAccountId, savings2AccountId]
    );
  });

  it('Cancel on consent page', () => {
    fdxTppLoginPage.assertThatPageIsDisplayed();
    fdxTppLoginPage.assertThatAuthorizationDetailsAreDispalyed();
    fdxTppLoginPage.clickNext();

    fdxTppIntentRegisteredPage.assertThatPageIsDisplayed();
    fdxTppIntentRegisteredPage.assertThatRequestUriFieldsAreNotEmpty();
    fdxTppIntentRegisteredPage.clickLogin();

    acpLoginPage.login(Credentials.tppUsername, Credentials.defaultPassword);
    if (environmentVariables.isMfaEnabled()) {
      mfaPage.typePin();
    }

    consentPage.checkAccounts([digitalBankingAccountId]);
    consentPage.cancel();

    errorPage.assertError(`acp returned an error: rejected: `);
  })

  it('Cancel on ACP login', () => {
    fdxTppLoginPage.assertThatPageIsDisplayed();
    fdxTppLoginPage.assertThatAuthorizationDetailsAreDispalyed();
    fdxTppLoginPage.clickNext();

    fdxTppIntentRegisteredPage.assertThatPageIsDisplayed();
    fdxTppIntentRegisteredPage.assertThatRequestUriFieldsAreNotEmpty();
    fdxTppIntentRegisteredPage.clickLogin();

    acpLoginPage.cancel();
    errorPage.assertError(`The user rejected the authentication`);
  })
});
