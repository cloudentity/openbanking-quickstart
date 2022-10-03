import { AcpLoginPage } from "../../pages/acp/AcpLoginPage";
import { AccountConsentPage } from "../../pages/consent/AccountConsentPage";
import { Credentials } from "../../pages/Credentials";
import { ConsentSelfServicePage } from "../../pages/consent-self-service/ConsentSelfServicePage";
import { ConsentSelfServiceApplicationPage } from "../../pages/consent-self-service/ConsentSelfServiceApplicationPage";
import { ConsentSelfServiceAccountDetailsPage } from "../../pages/consent-self-service/ConsentSelfServiceAccountDetailsPage";
import { MockDataRecipientNavigationPage } from "../../pages/mock-data-recipient/MockDataRecipientNavigationPage";
import { DiscoverDataHoldersPage } from "../../pages/mock-data-recipient/DiscoverDataHoldersPage";
import { DynamicClientRegistrationPage } from "../../pages/mock-data-recipient/DynamicClientRegistrationPage";
import { ConsentAndAuthorisationPage } from "../../pages/mock-data-recipient/ConsentAndAuthorisationPage";
import { ConsentAndAuthorisationCallbackPage } from "../../pages/mock-data-recipient/ConsentAndAuthorisationCallbackPage";
import { Urls } from "../../pages/Urls";
import { Accounts } from "../../pages/Accounts";

describe(`CDR Consent self service tests`, () => {
  const mockDataRecipientNavigationPage: MockDataRecipientNavigationPage = new MockDataRecipientNavigationPage();
  const discoverDataHoldersPage: DiscoverDataHoldersPage = new DiscoverDataHoldersPage();
  const dynamicClientRegistrationPage: DynamicClientRegistrationPage = new DynamicClientRegistrationPage();
  const consentAndAuthorisationPage: ConsentAndAuthorisationPage = new ConsentAndAuthorisationPage();
  const consentAndAuthorisationCallbackPage: ConsentAndAuthorisationCallbackPage = new ConsentAndAuthorisationCallbackPage();
  const acpLoginPage: AcpLoginPage = new AcpLoginPage();
  const accountConsentPage: AccountConsentPage = new AccountConsentPage();
  const consentSelfServicePage: ConsentSelfServicePage = new ConsentSelfServicePage();
  const consentSelfServiceApplicationPage: ConsentSelfServiceApplicationPage = new ConsentSelfServiceApplicationPage();
  const consentSelfServiceAccountDetailsPage: ConsentSelfServiceAccountDetailsPage = new ConsentSelfServiceAccountDetailsPage();


  before(`Dynamic Client Registration and Authorization via CDR mock data recipient`, () => {
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

    accountConsentPage.checkAccounts([Accounts.ids.CDR.savings, Accounts.ids.CDR.checking]);
    accountConsentPage.clickAgree();
    accountConsentPage.assertThatPageIsNotVisible();

    consentAndAuthorisationCallbackPage.assertThatPageIsDisplayed();
  });

  beforeEach(() => {
    consentSelfServicePage.visit(true);
    Urls.clearLocalStorage();
    consentSelfServicePage.visit(true);

    acpLoginPage.login(Credentials.tppUsername, Credentials.defaultPassword);

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
    consentSelfServiceApplicationPage.assertNumberOfConsents(1);
    consentSelfServiceApplicationPage.expandAccountConsentRow();

    consentSelfServiceAccountDetailsPage.assertThatAccountDetailsAreVisible();
    consentSelfServiceAccountDetailsPage.clickRevokeAccessButton();
    consentSelfServiceAccountDetailsPage.assertThatRevokeAccountDetailsAreVisible();
    consentSelfServiceAccountDetailsPage.confirmRevokeAccessAction();

    consentSelfServicePage.assertThatNoAccountsPageIsDisplayed();
  });

});
