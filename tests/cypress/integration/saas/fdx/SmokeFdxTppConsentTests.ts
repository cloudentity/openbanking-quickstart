import { AcpLoginPage } from "../../../pages/acp/AcpLoginPage";
import { AccountConsentPage } from "../../../pages/consent/AccountConsentPage";
import { Accounts } from "../../../pages/Accounts";
import { FdxTppLandingPage } from "../../../pages/fdx-tpp/FdxTppLandingPage";
import { FdxTppIntentRegisteredPage } from "../../../pages/fdx-tpp/FdxTppIntentRegisteredPage";
import { FdxTppAuthenticatedPage } from "../../../pages/fdx-tpp/FdxTppAuthenticatedPage";

describe(`FDX Tpp consent app`, () => {
  const fdxTppLandingPage: FdxTppLandingPage = new FdxTppLandingPage();
  const fdxTppIntentRegisteredPage: FdxTppIntentRegisteredPage =
    new FdxTppIntentRegisteredPage();
  const fdxTppAuthenticatedPage: FdxTppAuthenticatedPage =
    new FdxTppAuthenticatedPage();
  const acpLoginPage: AcpLoginPage = new AcpLoginPage();
  const accountConsentPage: AccountConsentPage = new AccountConsentPage();


  beforeEach(() => {
    fdxTppLandingPage.visit();
  });

  [
    [Accounts.ids.FDX.checkingAcc, Accounts.ids.FDX.savings1, Accounts.ids.FDX.savings2],
    [Accounts.ids.FDX.savings1],
    [Accounts.ids.FDX.savings2],
    [Accounts.ids.FDX.checkingAcc],
  ].forEach((accountsIds) => {
    it(`Happy path with selected accounts: ${accountsIds}`, () => {
      fdxTppLandingPage.assertThatPageIsDisplayed();
      fdxTppLandingPage.clickNext();

      fdxTppIntentRegisteredPage.assertThatPageIsDisplayed();
      fdxTppIntentRegisteredPage.clickLogin();

      acpLoginPage.loginWithMfaOption();

      accountConsentPage.assertPermissions(4);
      accountConsentPage.clickContinue();
      accountConsentPage.checkAccounts(accountsIds);
      accountConsentPage.clickAgree();

      fdxTppAuthenticatedPage.assertThatPageIsDisplayed();
      fdxTppAuthenticatedPage.assertThatTokenResponseFieldIsNotEmpty();
      fdxTppAuthenticatedPage.assertThatAccessTokenFieldIsNotEmpty();
      fdxTppAuthenticatedPage.assertThatConsentResponseFieldContainsAccountsIds(accountsIds);

      fdxTppAuthenticatedPage.clickTryNext();
      fdxTppLandingPage.assertThatPageIsDisplayed();
      cy.wait(3000) // Workaround due to pipeline issues >>> AUT-6292
    });
  });

  it(`Happy path with not selected account`, () => {
    fdxTppLandingPage.assertThatPageIsDisplayed();
    fdxTppLandingPage.clickNext();

    fdxTppIntentRegisteredPage.assertThatPageIsDisplayed();
    fdxTppIntentRegisteredPage.clickLogin();

    acpLoginPage.loginWithMfaOption();

    accountConsentPage.clickContinue();
    accountConsentPage.clickAgree();

    fdxTppAuthenticatedPage.assertThatPageIsDisplayed();
    fdxTppAuthenticatedPage.assertThatConsentResponseFieldNotContainsAccountsIds(
      [Accounts.ids.FDX.checkingAcc, Accounts.ids.FDX.savings1, Accounts.ids.FDX.savings1]
    );
  });

});
