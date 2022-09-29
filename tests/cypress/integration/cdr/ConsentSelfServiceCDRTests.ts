import { AcpLoginPage } from "../../pages/acp/AcpLoginPage";
import { AccountConsentPage } from "../../pages/consent/AccountConsentPage";
import { Credentials } from "../../pages/Credentials";
import { ConsentSelfServicePage } from "../../pages/consent-self-service/ConsentSelfServicePage";
import { ConsentSelfServiceApplicationPage } from "../../pages/consent-self-service/ConsentSelfServiceApplicationPage";
import { MockDataRecipientNavigationPage } from "../../pages/mock-data-recipient/MockDataRecipientNavigationPage";
import { DiscoverDataHoldersPage } from "../../pages/mock-data-recipient/DiscoverDataHoldersPage";
import { DynamicClientRegistrationPage } from "../../pages/mock-data-recipient/DynamicClientRegistrationPage";
import { ConsentAndAuthorisationPage } from "../../pages/mock-data-recipient/ConsentAndAuthorisationPage";
import { ConsentAndAuthorisationCallbackPage } from "../../pages/mock-data-recipient/ConsentAndAuthorisationCallbackPage";
import { Urls } from "../../pages/Urls";

describe(`Consent self service app CDR`, () => {
  const mockDataRecipientNavigationPage: MockDataRecipientNavigationPage = new MockDataRecipientNavigationPage();
  const discoverDataHoldersPage: DiscoverDataHoldersPage = new DiscoverDataHoldersPage();
  const dynamicClientRegistrationPage: DynamicClientRegistrationPage = new DynamicClientRegistrationPage();
  const consentAndAuthorisationPage: ConsentAndAuthorisationPage = new ConsentAndAuthorisationPage();
  const consentAndAuthorisationCallbackPage: ConsentAndAuthorisationCallbackPage = new ConsentAndAuthorisationCallbackPage();
  const acpLoginPage: AcpLoginPage = new AcpLoginPage();
  const consentPage: AccountConsentPage = new AccountConsentPage();
  const consentSelfServicePage: ConsentSelfServicePage = new ConsentSelfServicePage();
  const consentSelfServiceApplicationPage: ConsentSelfServiceApplicationPage = new ConsentSelfServiceApplicationPage();

  before(() => {
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

    mockDataRecipientNavigationPage.clickConsentAndAuthorisationLink();
    consentAndAuthorisationPage.assertThatPageIsDisplayed();
    consentAndAuthorisationPage.selectClientRegistration(1);
    consentAndAuthorisationPage.setSharingDuration(1000000);
    consentAndAuthorisationPage.clickConstructAuthorizationUriButton();
    consentAndAuthorisationPage.assertThatAuthorizationUriIsGenerated();
    consentAndAuthorisationPage.clickOnAuthorizationUriLink();

    acpLoginPage.login(Credentials.tppUsername, Credentials.defaultPassword);
    consentPage.checkAccounts(["1000001", "1000002"]);
    consentPage.clickAgree();
    consentPage.assertThatPageIsNotVisible();
    consentAndAuthorisationCallbackPage.assertThatPageIsDisplayed();
  });

  beforeEach(() => {
    consentSelfServicePage.visit(true);
  });

  it(`Happy path with account consent`, () => {
    acpLoginPage.login(Credentials.tppUsername, Credentials.defaultPassword);
    consentSelfServicePage.clickOnApplicationCard();
    consentSelfServiceApplicationPage.expandAccountsTab();
    consentSelfServiceApplicationPage.checkAccount("1000001");
    consentSelfServiceApplicationPage.checkAccount("1000002");
    consentSelfServiceApplicationPage.expandAccountConsentRow();
    consentSelfServiceApplicationPage.assertAccountRevokePopupContainsText(
      "Your Name and occupation"
    );
  });

  it(`Revoke CDR Arrangement`, () => {
    acpLoginPage.login(Credentials.tppUsername, Credentials.defaultPassword);
    consentSelfServicePage.clickOnApplicationCard();
    consentSelfServiceApplicationPage.expandAccountsTab();
    consentSelfServiceApplicationPage.assertNumberOfConsents(1);
    consentSelfServiceApplicationPage.expandAccountConsentRow();
    consentSelfServiceApplicationPage.clickRevokeAccessButton();
    consentSelfServiceApplicationPage.assertNumberOfConsents(0);
  });
});
