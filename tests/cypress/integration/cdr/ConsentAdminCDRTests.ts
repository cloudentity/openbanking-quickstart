import {AcpLoginPage} from '../../pages/acp/AcpLoginPage';
import {ConsentPage} from '../../pages/consent/ConsentPage';
import {Credentials} from "../../pages/Credentials";
import {ConsentAdminPage} from '../../pages/consent-admin/ConsentAdminPage';
import { MockDataRecipientPage } from '../../pages/mock-data-recipient/MockDataRecipientPage';
import { ErrorPage } from "../../pages/ErrorPage";

describe(`Consent admin portal CDR`, () => {
  const acpLoginPage: AcpLoginPage = new AcpLoginPage();
  const consentPage: ConsentPage = new ConsentPage();
  const consentAdminPage: ConsentAdminPage = new ConsentAdminPage();
  const mockDataRecipientPage: MockDataRecipientPage = new MockDataRecipientPage(); 
  const errorPage: ErrorPage = new ErrorPage();


  before(() => {
    mockDataRecipientPage.visit()
    mockDataRecipientPage.visitDiscoverDataHoldersTab()
    mockDataRecipientPage.clickDataHoldersRefresh()
    mockDataRecipientPage.visitDynamicClientRegistrationTab() 
    mockDataRecipientPage.clickDCRRegisterButton()
    mockDataRecipientPage.visitConsentAndAuthorisationTab() 
    mockDataRecipientPage.selectClientRegistration(1)
    mockDataRecipientPage.inputSharingDuration(1000000)
    mockDataRecipientPage.clickConstructAuthorisationURI()
    acpLoginPage.login(Credentials.tppUsername, Credentials.defaultPassword)
    consentPage.confirm()
  })

  it(`Happy path with revoking consent from Consent management page`, () => {
    consentAdminPage.visit(true)
    acpLoginPage.login(
      Credentials.consentAdminUsername,
      Credentials.defaultPassword
    );
  
    consentAdminPage.assertThatConsentManagementTabIsDisplayed()
    consentAdminPage.searchAccount("1000001");
    consentAdminPage.assertAccountResult("1000001");
    consentAdminPage.assertClientAccountWithStatus("MyBudgetHelper", "Active");
    consentAdminPage.manageAccount("MyBudgetHelper");
    consentAdminPage.assertConsentsDetails();
    consentAdminPage.revokeClientConsentByAccountName("MyBudgetHelper");
    consentAdminPage.assertClientAccountWithStatus("MyBudgetHelper", "Inactive");
  })

  it(`Happy path with revoking consent from Third party providers page`, () => {
    mockDataRecipientPage.visit()
    mockDataRecipientPage.visitDiscoverDataHoldersTab()
    mockDataRecipientPage.clickDataHoldersRefresh()
    mockDataRecipientPage.visitDynamicClientRegistrationTab() 
    mockDataRecipientPage.clickDCRRegisterButton()
    mockDataRecipientPage.visitConsentAndAuthorisationTab() 
    mockDataRecipientPage.selectClientRegistration(1)
    mockDataRecipientPage.inputSharingDuration(1000000)
    mockDataRecipientPage.clickConstructAuthorisationURI()
    acpLoginPage.login(Credentials.tppUsername, Credentials.defaultPassword)
    consentPage.confirm()

    consentAdminPage.visit(true)

    acpLoginPage.login(
      Credentials.consentAdminUsername,
      Credentials.defaultPassword
    );
    consentAdminPage.assertThatConsentManagementTabIsDisplayed()
    consentAdminPage.revokeClientConsent();
  });


  it(`Cancel first ACP login`, () => {
    consentAdminPage.visit(true)
    acpLoginPage.cancel();
    errorPage.assertError(`The user rejected the authentication`);
  });

  it(`Cancel second ACP login`, () => {
    consentAdminPage.visit(true)
    acpLoginPage.cancel();
    errorPage.assertError(`The user rejected the authentication`);
  });
})