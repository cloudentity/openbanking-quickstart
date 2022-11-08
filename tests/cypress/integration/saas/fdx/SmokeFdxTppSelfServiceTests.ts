import { AcpLoginPage } from "../../../pages/acp/AcpLoginPage";
import { AccountConsentPage } from "../../../pages/consent/AccountConsentPage";
import { ConsentSelfServicePage } from "../../../pages/consent-self-service/ConsentSelfServicePage";
import { ConsentSelfServiceApplicationPage } from "../../../pages/consent-self-service/ConsentSelfServiceApplicationPage";
import { ConsentSelfServiceAccountDetailsPage } from "../../../pages/consent-self-service/ConsentSelfServiceAccountDetailsPage";
import { Accounts } from "../../../pages/Accounts";
import { FdxTppLandingPage } from "../../../pages/fdx-tpp/FdxTppLandingPage";
import { FdxTppIntentRegisteredPage } from "../../../pages/fdx-tpp/FdxTppIntentRegisteredPage";
import { FdxTppAuthenticatedPage } from "../../../pages/fdx-tpp/FdxTppAuthenticatedPage";

describe(`FDX TPP Consent admin portal tests`, () => {
  const fdxTppLandingPage: FdxTppLandingPage = new FdxTppLandingPage();
  const fdxTppIntentRegisteredPage: FdxTppIntentRegisteredPage = new FdxTppIntentRegisteredPage();
  const fdxTppAuthenticatedPage: FdxTppAuthenticatedPage = new FdxTppAuthenticatedPage();
  const acpLoginPage: AcpLoginPage = new AcpLoginPage();
  const accountConsentPage: AccountConsentPage = new AccountConsentPage();
  const consentSelfServicePage: ConsentSelfServicePage = new ConsentSelfServicePage();
  const consentSelfServiceApplicationPage: ConsentSelfServiceApplicationPage = new ConsentSelfServiceApplicationPage();
  const consentSelfServiceAccountDetailsPage: ConsentSelfServiceAccountDetailsPage = new ConsentSelfServiceAccountDetailsPage();

  const accountsIDs = [Accounts.ids.FDX.checkingAcc, Accounts.ids.FDX.savings1];

  //  Ignored due to BUG - AUT-7531

  // before(() => {
  //   fdxTppLandingPage.visit();

  //   fdxTppLandingPage.assertThatPageIsDisplayed();
  //   fdxTppLandingPage.clickNext();

  //   fdxTppIntentRegisteredPage.assertThatPageIsDisplayed();
  //   fdxTppIntentRegisteredPage.clickLogin();

  //   acpLoginPage.assertThatModalIsDisplayed("FDX");
  //   acpLoginPage.loginWithMfaOption();

  //   accountConsentPage.assertPermissions(4);
  //   accountConsentPage.assertThatAccountsAreNotVisible([
  //     Accounts.ids.FDX.checkingAcc,
  //     Accounts.ids.FDX.savings1,
  //     Accounts.ids.FDX.savings2,
  //   ]);

  //   accountConsentPage.clickContinue();
  //   accountConsentPage.checkAccounts(accountsIDs);
  //   accountConsentPage.clickAgree();

  //   fdxTppAuthenticatedPage.assertThatPageIsDisplayed();
  //   fdxTppAuthenticatedPage.assertThatTokenResponseFieldIsNotEmpty();
  //   fdxTppAuthenticatedPage.assertThatAccessTokenFieldIsNotEmpty();
  //   fdxTppAuthenticatedPage.assertThatConsentResponseFieldContainsAccountsIds(accountsIDs);

  //   fdxTppAuthenticatedPage.clickTryNext();
  //   fdxTppLandingPage.assertThatPageIsDisplayed();
  // });

  // beforeEach(() => {
  //   consentSelfServicePage.visit(true);

  //   acpLoginPage.assertThatModalIsDisplayed("Bank customers");
  //   acpLoginPage.login();

  //   consentSelfServicePage.clickOnApplicationCardWithName("Developer TPP");
  // });

  it(`Happy path with account consent`, () => {
  //   consentSelfServiceApplicationPage.expandAccountsTab();
  //   consentSelfServiceApplicationPage.checkAccountHasStatus(accountsIDs[0], "Authorised");
  //   consentSelfServiceApplicationPage.checkAccountHasStatus(accountsIDs[1], "Authorised");
  //   consentSelfServiceApplicationPage.expandAccountConsentRow();

  //   consentSelfServiceAccountDetailsPage.assertThatAccountDetailsAreVisible()
  //   consentSelfServiceAccountDetailsPage.assertAccount(accountsIDs[0]);
  //   consentSelfServiceAccountDetailsPage.assertAccount(accountsIDs[1]);
  });

  it(`Revoke account consent`, () => {
  //   consentSelfServiceApplicationPage.expandAccountsTab();
  //   consentSelfServiceApplicationPage.assertAuthorisedAccountRowExists(accountsIDs[0]);
  //   consentSelfServiceApplicationPage.assertAuthorisedAccountRowExists(accountsIDs[1]);
  //   consentSelfServiceApplicationPage.expandAccountConsentRow();

  //   consentSelfServiceAccountDetailsPage.assertThatAccountDetailsAreVisible();
  //   consentSelfServiceAccountDetailsPage.clickRevokeAccessButton();
  //   consentSelfServiceAccountDetailsPage.assertThatRevokeAccountDetailsAreVisible();
  //   consentSelfServiceAccountDetailsPage.confirmRevokeAccessAction();

  //   consentSelfServicePage.assertThatNoAccountsPageIsDisplayed();
  });
});
