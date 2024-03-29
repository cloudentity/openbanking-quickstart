import { AcpLoginPage } from "../../pages/acp/AcpLoginPage";
import { AccountConsentPage } from "../../pages/consent/AccountConsentPage";
import { ConsentSelfServicePage } from "../../pages/consent-self-service/ConsentSelfServicePage";
import { ConsentSelfServiceApplicationPage } from "../../pages/consent-self-service/ConsentSelfServiceApplicationPage";
import { ConsentSelfServiceAccountDetailsPage } from "../../pages/consent-self-service/ConsentSelfServiceAccountDetailsPage";
import { MockDataRecipientNavigationPage } from "../../pages/mock-data-recipient/MockDataRecipientNavigationPage";
import { DiscoverDataHoldersPage } from "../../pages/mock-data-recipient/DiscoverDataHoldersPage";
import { DynamicClientRegistrationPage } from "../../pages/mock-data-recipient/DynamicClientRegistrationPage";
import { PushedAuthorisationRequestPage } from "../../pages/mock-data-recipient/PushedAuthorisationRequestPage";
import { ConsentAndAuthorisationCallbackPage } from "../../pages/mock-data-recipient/ConsentAndAuthorisationCallbackPage";
import { Accounts } from "../../pages/Accounts";

describe(`CDR Consent self service tests`, () => {
  const mockDataRecipientNavigationPage: MockDataRecipientNavigationPage = new MockDataRecipientNavigationPage();
  const discoverDataHoldersPage: DiscoverDataHoldersPage = new DiscoverDataHoldersPage();
  const dynamicClientRegistrationPage: DynamicClientRegistrationPage = new DynamicClientRegistrationPage();
  const pushedAuthorisationRequestPage: PushedAuthorisationRequestPage = new PushedAuthorisationRequestPage();
  const consentAndAuthorisationCallbackPage: ConsentAndAuthorisationCallbackPage = new ConsentAndAuthorisationCallbackPage();
  const acpLoginPage: AcpLoginPage = new AcpLoginPage();
  const accountConsentPage: AccountConsentPage = new AccountConsentPage();
  const consentSelfServicePage: ConsentSelfServicePage = new ConsentSelfServicePage();
  const consentSelfServiceApplicationPage: ConsentSelfServiceApplicationPage = new ConsentSelfServiceApplicationPage();
  const consentSelfServiceAccountDetailsPage: ConsentSelfServiceAccountDetailsPage = new ConsentSelfServiceAccountDetailsPage();


  beforeEach(`Dynamic Client Registration via CDR mock data recipient`, () => {
    mockDataRecipientNavigationPage.visit(true);
    mockDataRecipientNavigationPage.clickDiscoverDataHoldersLink();

    discoverDataHoldersPage.assertThatPageIsDisplayed();
    discoverDataHoldersPage.clickRefreshDataHoldersButton();
    discoverDataHoldersPage.clickResetDataHoldersButton();
    discoverDataHoldersPage.clickRefreshDataHoldersButton();
    discoverDataHoldersPage.assertThatDataHolderBrandsLoaded();

    mockDataRecipientNavigationPage.clickDynamicClientRegistrationLink();

    dynamicClientRegistrationPage.clickMockDataHolderBankingBrand();
    dynamicClientRegistrationPage.assertThatDynamicClientRegistrationFormIsDisplayed();
    dynamicClientRegistrationPage.clickDCRRegisterButton();
    dynamicClientRegistrationPage.assertThatClientRegistered();
  });

  beforeEach(`Authorize via CDR mock data recipient`, () => {
    mockDataRecipientNavigationPage.clickParLink();

    pushedAuthorisationRequestPage.assertThatPageIsDisplayed();
    pushedAuthorisationRequestPage.selectClientRegistration(1);
    pushedAuthorisationRequestPage.setSharingDuration(1000000);
    pushedAuthorisationRequestPage.setScopes([
      'openid', 'profile', 'bank:accounts.basic:read',
      'bank:accounts.detail:read', 'bank:transactions:read',
      'common:customer.basic:read', 'offline_access']);
    pushedAuthorisationRequestPage.clickInitiateParButton();
    pushedAuthorisationRequestPage.assertThatAuthorizationUriIsGenerated();
    pushedAuthorisationRequestPage.clickOnAuthorizationUriLink();

    acpLoginPage.assertThatModalIsDisplayed();
    acpLoginPage.login();

    accountConsentPage.checkAccounts([Accounts.ids.CDR.savings, Accounts.ids.CDR.checking]);
    accountConsentPage.clickAgree();
    accountConsentPage.assertThatPageIsNotVisible();

    consentAndAuthorisationCallbackPage.assertThatPageIsDisplayed();
  });

  beforeEach(`Go to Consent Self Service Page`, () => {
    consentSelfServicePage.visit(true);

    acpLoginPage.assertThatModalIsDisplayed();
    acpLoginPage.login();

    consentSelfServicePage.clickOnApplicationCard();
  });

  it(`Happy path with account consent`, () => {
    consentSelfServiceApplicationPage.expandAccountsTab();
    consentSelfServiceApplicationPage.checkAccount(Accounts.ids.CDR.savings);
    consentSelfServiceApplicationPage.checkAccount(Accounts.ids.CDR.checking);
    consentSelfServiceApplicationPage.expandAccountConsentRow();

    consentSelfServiceAccountDetailsPage.assertThatAccountDetailsAreVisible()
    consentSelfServiceAccountDetailsPage.assertAccount(Accounts.ids.CDR.savings);
    consentSelfServiceAccountDetailsPage.assertAccount(Accounts.ids.CDR.checking);
  });

  it(`Revoke CDR Arrangement`, () => {
    consentSelfServiceApplicationPage.expandAccountsTab();
    consentSelfServiceApplicationPage.assertAuthorisedAccountRowExists(Accounts.ids.CDR.savings);
    consentSelfServiceApplicationPage.assertAuthorisedAccountRowExists(Accounts.ids.CDR.checking);
    consentSelfServiceApplicationPage.expandAccountConsentRow();

    consentSelfServiceAccountDetailsPage.assertThatAccountDetailsAreVisible();
    consentSelfServiceAccountDetailsPage.clickRevokeAccessButton();
    consentSelfServiceAccountDetailsPage.assertThatRevokeAccountDetailsAreVisible();
    consentSelfServiceAccountDetailsPage.confirmRevokeAccessAction();

    consentSelfServicePage.clickOnApplicationCardWithName("MyBudgetHelper");
    consentSelfServiceApplicationPage.assertAuthorisedAccountRowDoesNotExist(Accounts.ids.CDR.savings);
    consentSelfServiceApplicationPage.assertAuthorisedAccountRowDoesNotExist(Accounts.ids.CDR.checking);
  });

  afterEach(`Remove DCR client from CDR mock data recipient`, () => {
    mockDataRecipientNavigationPage.visit(true);
    mockDataRecipientNavigationPage.clickDynamicClientRegistrationLink();

    dynamicClientRegistrationPage.assertThatPageIsDisplayed();
    dynamicClientRegistrationPage.clickDeleteClientButton();
    dynamicClientRegistrationPage.assertThatRegisteredClientWasRemoved();
  });

});
