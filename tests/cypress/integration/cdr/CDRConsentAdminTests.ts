import { AcpLoginPage } from "../../pages/acp/AcpLoginPage";
import { AccountConsentPage } from "../../pages/consent/AccountConsentPage";
import { Credentials } from "../../pages/Credentials";
import { ConsentAdminPage } from "../../pages/consent-admin/ConsentAdminPage";
import { MockDataRecipientNavigationPage } from "../../pages/mock-data-recipient/MockDataRecipientNavigationPage";
import { DiscoverDataHoldersPage } from "../../pages/mock-data-recipient/DiscoverDataHoldersPage";
import { DynamicClientRegistrationPage } from "../../pages/mock-data-recipient/DynamicClientRegistrationPage";
import { PushedAuthorisationRequestPage } from "../../pages/mock-data-recipient/PushedAuthorisationRequestPage";
import { ConsentAndAuthorisationCallbackPage } from "../../pages/mock-data-recipient/ConsentAndAuthorisationCallbackPage";
import { Urls } from "../../pages/Urls";
import { Accounts } from "../../pages/Accounts";

describe(`CDR Consent admin portal tests`, () => {
  const mockDataRecipientNavigationPage: MockDataRecipientNavigationPage = new MockDataRecipientNavigationPage();
  const discoverDataHoldersPage: DiscoverDataHoldersPage = new DiscoverDataHoldersPage();
  const dynamicClientRegistrationPage: DynamicClientRegistrationPage = new DynamicClientRegistrationPage();
  const pushedAuthorisationRequestPage: PushedAuthorisationRequestPage = new PushedAuthorisationRequestPage();
  const consentAndAuthorisationCallbackPage: ConsentAndAuthorisationCallbackPage = new ConsentAndAuthorisationCallbackPage();
  const acpLoginPage: AcpLoginPage = new AcpLoginPage();
  const accountConsentPage: AccountConsentPage = new AccountConsentPage();
  const consentAdminPage: ConsentAdminPage = new ConsentAdminPage();

  beforeEach(`Dynamic Client Registration via CDR mock data recipient`, () => {
    mockDataRecipientNavigationPage.visit(true);
    Urls.clearLocalStorage();
    mockDataRecipientNavigationPage.visit(true);
    mockDataRecipientNavigationPage.clickDiscoverDataHoldersLink();

    discoverDataHoldersPage.assertThatPageIsDisplayed();
    discoverDataHoldersPage.clickRefreshDataHoldersButton();
    discoverDataHoldersPage.assertThatDataHolderBrandsLoaded();

    mockDataRecipientNavigationPage.clickDynamicClientRegistrationLink();

    dynamicClientRegistrationPage.assertThatPageIsDisplayed();
    dynamicClientRegistrationPage.assertThatBrandIdIsSelected();
    dynamicClientRegistrationPage.clickDCRRegisterButton();
    dynamicClientRegistrationPage.assertThatClientRegistered();
  });

  beforeEach(`Authorize via CDR mock data recipient`, () => {
    mockDataRecipientNavigationPage.clickParLink();

    pushedAuthorisationRequestPage.assertThatPageIsDisplayed();
    pushedAuthorisationRequestPage.selectClientRegistration(1);
    pushedAuthorisationRequestPage.setSharingDuration(1000000);
    pushedAuthorisationRequestPage.clickInitiateParButton();
    pushedAuthorisationRequestPage.assertThatAuthorizationUriIsGenerated();
    pushedAuthorisationRequestPage.clickOnAuthorizationUriLink();

    acpLoginPage.login();

    accountConsentPage.checkAccounts([Accounts.ids.CDR.savings, Accounts.ids.CDR.loan]);
    accountConsentPage.clickAgree();
    accountConsentPage.assertThatPageIsNotVisible();

    consentAndAuthorisationCallbackPage.assertThatPageIsDisplayed();
  });

  it(`Happy path with revoking consent from Consent management page`, () => {
    consentAdminPage.visit(true);
    consentAdminPage.login();

    consentAdminPage.assertThatConsentManagementTabIsDisplayed();
    consentAdminPage.searchAccount(Accounts.ids.CDR.savings);
    consentAdminPage.assertAccountResult(Accounts.ids.CDR.savings);
    consentAdminPage.assertClientAccountWithStatus("MyBudgetHelper", "Active");
    consentAdminPage.manageAccount("MyBudgetHelper");
    consentAdminPage.assertConsentsDetails();
    consentAdminPage.revokeClientConsentByAccountName("MyBudgetHelper");
    consentAdminPage.assertClientAccountWithStatus("MyBudgetHelper", "Inactive");
  });

  it(`Happy path with revoking consent from Third party providers page`, () => {
    consentAdminPage.visit(true);
    consentAdminPage.login();

    consentAdminPage.assertThatConsentManagementTabIsDisplayed();
    consentAdminPage.revokeClientConsent();
  });

  afterEach(`Remove DCR client from CDR mock data recipient`, () => {
    mockDataRecipientNavigationPage.visit(true);
    Urls.clearLocalStorage();
    mockDataRecipientNavigationPage.visit(true);

    mockDataRecipientNavigationPage.clickDynamicClientRegistrationLink();

    dynamicClientRegistrationPage.assertThatPageIsDisplayed();
    dynamicClientRegistrationPage.clickDeleteClientButton();
    dynamicClientRegistrationPage.assertThatRegisteredClientWasRemoved();
  });

});
