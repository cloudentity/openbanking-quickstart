import {AcpLoginPage} from '../../pages/acp/AcpLoginPage';
import {ConsentPage} from '../../pages/consent/ConsentPage';
import {Credentials} from '../../pages/Credentials';
import {ConsentAdminPage} from '../../pages/consent-admin/ConsentAdminPage';
import {MockDataRecipientNavigationPage} from '../../pages/mock-data-recipient/MockDataRecipientNavigationPage'
import {DiscoverDataHoldersPage} from '../../pages/mock-data-recipient/DiscoverDataHoldersPage'
import {DynamicClientRegistrationPage} from '../../pages/mock-data-recipient/DynamicClientRegistrationPage'
import {ConsentAndAuthorisationPage} from '../../pages/mock-data-recipient/ConsentAndAuthorisationPage'
import {ConsentAndAuthorisationCallbackPage} from '../../pages/mock-data-recipient/ConsentAndAuthorisationCallbackPage'
import {Urls} from '../../pages/Urls';

describe(`Consent admin portal CDR`, () => {
  const mockDataRecipientNavigationPage: MockDataRecipientNavigationPage = new MockDataRecipientNavigationPage(); 
  const discoverDataHoldersPage: DiscoverDataHoldersPage = new DiscoverDataHoldersPage();
  const dynamicClientRegistrationPage: DynamicClientRegistrationPage = new DynamicClientRegistrationPage();
  const consentAndAuthorisationPage: ConsentAndAuthorisationPage = new ConsentAndAuthorisationPage();
  const consentAndAuthorisationCallbackPage: ConsentAndAuthorisationCallbackPage = new ConsentAndAuthorisationCallbackPage();
  const acpLoginPage: AcpLoginPage = new AcpLoginPage();
  const consentPage: ConsentPage = new ConsentPage();
  const consentAdminPage: ConsentAdminPage = new ConsentAdminPage();

  before(() => {
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

    mockDataRecipientNavigationPage.clickConsentAndAuthorisationLink();
    consentAndAuthorisationPage.assertThatPageIsDisplayed();
    consentAndAuthorisationPage.selectClientRegistration(1);
    consentAndAuthorisationPage.setSharingDuration(1000000);
    consentAndAuthorisationPage.clickConstructAuthorizationUriButton();
    consentAndAuthorisationPage.assertThatAuthorizationUriIsGenerated();
    consentAndAuthorisationPage.clickOnAuthorizationUriLink();

    acpLoginPage.login(Credentials.tppUsername, Credentials.defaultPassword);
    consentPage.confirm();
    consentPage.assertThatPageIsNotVisible();
    consentAndAuthorisationCallbackPage.assertThatPageIsDisplayed();
  })

  it(`Happy path with revoking consent from Consent management page`, () => {
    consentAdminPage.visit(true);
    Urls.clearLocalStorage();
    consentAdminPage.visit(true);
    consentAdminPage.login();
  
    consentAdminPage.assertThatConsentManagementTabIsDisplayed();
    consentAdminPage.searchAccount("1000001");
    consentAdminPage.assertAccountResult("1000001");
    consentAdminPage.assertClientAccountWithStatus("MyBudgetHelper", "Active");
    consentAdminPage.manageAccount("MyBudgetHelper");
    consentAdminPage.assertConsentsDetails();
    consentAdminPage.revokeClientConsentByAccountName("MyBudgetHelper");
    consentAdminPage.assertClientAccountWithStatus("MyBudgetHelper", "Inactive");
  })

  it(`Happy path with revoking consent from Third party providers page`, () => {
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

    mockDataRecipientNavigationPage.clickConsentAndAuthorisationLink();
    consentAndAuthorisationPage.assertThatPageIsDisplayed();
    consentAndAuthorisationPage.selectClientRegistration(1);
    consentAndAuthorisationPage.setSharingDuration(1000000);
    consentAndAuthorisationPage.clickConstructAuthorizationUriButton();
    consentAndAuthorisationPage.assertThatAuthorizationUriIsGenerated();
    consentAndAuthorisationPage.clickOnAuthorizationUriLink();

    acpLoginPage.login(Credentials.tppUsername, Credentials.defaultPassword)
    consentPage.confirm();
    consentPage.assertThatPageIsNotVisible();
    consentAndAuthorisationCallbackPage.assertThatPageIsDisplayed();

    consentAdminPage.visit(true);
    Urls.clearLocalStorage();
    consentAdminPage.visit(true);
    consentAdminPage.login();

    consentAdminPage.assertThatConsentManagementTabIsDisplayed();
    consentAdminPage.revokeClientConsent();
  });
})
