import { AcpLoginPage } from "../../../pages/acp/AcpLoginPage";
import { AccountConsentPage } from "../../../pages/consent/AccountConsentPage";
import { Credentials } from "../../../pages/Credentials";
import { ConsentSelfServicePage } from "../../../pages/consent-self-service/ConsentSelfServicePage";
import { ConsentSelfServiceApplicationPage } from "../../../pages/consent-self-service/ConsentSelfServiceApplicationPage";
import { ConsentSelfServiceAccountDetailsPage } from "../../../pages/consent-self-service/ConsentSelfServiceAccountDetailsPage";
import { Urls } from "../../../pages/Urls";
import { Accounts } from "../../../pages/Accounts";
import { FinancrooLoginPage } from "../../../pages/financroo/FinancrooLoginPage";
import { FinancrooWelcomePage } from "../../../pages/financroo/FinancrooWelcomePage";
import { EnvironmentVariables } from "../../../pages/EnvironmentVariables";
import { MfaPage } from "../../../pages/mfa/MfaPage";
import { FinancrooModalPage } from "../../../pages/financroo/accounts/FinancrooModalPage";
import { FinancrooAccountsPage } from "../../../pages/financroo/accounts/FinancrooAccountsPage";


describe(`Smoke Financroo Consent self service tests`, () => {
  const financrooLoginPage: FinancrooLoginPage = new FinancrooLoginPage();
  const financrooWelcomePage: FinancrooWelcomePage = new FinancrooWelcomePage();
  const acpLoginPage: AcpLoginPage = new AcpLoginPage();
  const environmentVariables: EnvironmentVariables = new EnvironmentVariables();
  const mfaPage: MfaPage = new MfaPage();
  const accountConsentPage: AccountConsentPage = new AccountConsentPage();
  const financrooModalPage: FinancrooModalPage = new FinancrooModalPage();
  const financrooAccountsPage: FinancrooAccountsPage = new FinancrooAccountsPage();
  const consentSelfServicePage: ConsentSelfServicePage = new ConsentSelfServicePage();
  const consentSelfServiceApplicationPage: ConsentSelfServiceApplicationPage = new ConsentSelfServiceApplicationPage();
  const consentSelfServiceAccountDetailsPage: ConsentSelfServiceAccountDetailsPage = new ConsentSelfServiceAccountDetailsPage();

  const accountsIDs = [Accounts.ids.CDR.savings, Accounts.ids.CDR.checking];


  before(() => {
    financrooLoginPage.visit();
    Urls.clearLocalStorage();
    financrooLoginPage.visit();
    financrooLoginPage.login();

    financrooWelcomePage.reconnectGoBank();

    acpLoginPage.login(Credentials.tppUsername, Credentials.defaultPassword);
    if (environmentVariables.isMfaEnabled()) {
      mfaPage.typePin();
    }

    accountConsentPage.checkAccounts(accountsIDs);
    accountConsentPage.expandPermissions();
    accountConsentPage.assertPermissionsDetails(
      "Purpose for sharing data",
      "To uncover insights that can improve your financial well being."
    );
    accountConsentPage.clickAgree();

    financrooModalPage.assertThatModalIsDisplayed();
    financrooModalPage.close();

    financrooAccountsPage.assertThatPageIsDisplayed();
    financrooAccountsPage.assertAccountsSyncedNumber(accountsIDs.length);
    financrooAccountsPage.assertAccountsIds(accountsIDs);
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
    consentSelfServiceApplicationPage.checkAccount(accountsIDs[0]);
    consentSelfServiceApplicationPage.checkAccount(accountsIDs[1]);
    consentSelfServiceApplicationPage.expandAccountConsentRow();

    consentSelfServiceAccountDetailsPage.assertThatAccountDetailsAreVisible()
    consentSelfServiceAccountDetailsPage.assertAccount(accountsIDs[0]);
    consentSelfServiceAccountDetailsPage.assertAccount(accountsIDs[1]);
  });

  it(`Revoke account consent`, () => {
    consentSelfServiceApplicationPage.expandAccountsTab();
    consentSelfServiceApplicationPage.assertNumberOfConsents(1);
    consentSelfServiceApplicationPage.expandAccountConsentRow();

    consentSelfServiceAccountDetailsPage.assertThatAccountDetailsAreVisible();
    consentSelfServiceAccountDetailsPage.clickRevokeAccessButton();
    consentSelfServiceAccountDetailsPage.assertThatRevokeAccountDetailsAreVisible();
    consentSelfServiceAccountDetailsPage.confirmRevokeAccessAction();

    consentSelfServicePage.assertThatNoAccountsPageIsDisplayed();

    financrooLoginPage.visit();
    financrooLoginPage.login();

    financrooAccountsPage.assertThatAccountsAreDisconnected();
  });

});
