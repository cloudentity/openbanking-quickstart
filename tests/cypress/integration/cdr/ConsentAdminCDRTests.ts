import {AcpLoginPage} from '../../pages/acp/AcpLoginPage';
import {ConsentPage} from '../../pages/consent/ConsentPage';
import {Credentials} from "../../pages/Credentials";
import {ConsentAdminPage} from '../../pages/consent-admin/ConsentAdminPage';
import { MockDataRecipientPage } from '../../pages/mock-data-recipient/MockDataRecipientPage';
import {Urls} from "../../pages/Urls";

describe(`Consent admin portal CDR`, () => {
  const acpLoginPage: AcpLoginPage = new AcpLoginPage();
  const consentPage: ConsentPage = new ConsentPage();
  const consentAdminPage: ConsentAdminPage = new ConsentAdminPage();
  const mockDataRecipientPage: MockDataRecipientPage = new MockDataRecipientPage(); 


  before(() => {
    mockDataRecipientPage.visit()
    mockDataRecipientPage.visitDiscoverDataHoldersTab()
    mockDataRecipientPage.clickDataHoldersRefresh()
    mockDataRecipientPage.assertThatDataHolderBrandsLoaded()
    mockDataRecipientPage.visitDynamicClientRegistrationTab() 
    mockDataRecipientPage.clickDCRRegisterButton()
    mockDataRecipientPage.assertThatClientRegistered()
    mockDataRecipientPage.visitConsentAndAuthorisationTab() 
    mockDataRecipientPage.selectClientRegistration(1)
    mockDataRecipientPage.inputSharingDuration(1000000)
    mockDataRecipientPage.clickConstructAuthorisationURI()
    acpLoginPage.login(Credentials.tppUsername, Credentials.defaultPassword)
    consentPage.confirm()
    consentPage.assertThatPageIsNotVisible()
    // verify data recipient mock consent callback page - UI error page improvements AUT-5845
  })

  it(`Happy path with revoking consent from Consent management page`, () => {
    consentAdminPage.visit(true)
    Urls.clearLocalStorage()
    consentAdminPage.visit(true)
    consentAdminPage.login()
  
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
    mockDataRecipientPage.assertThatDataHolderBrandsLoaded()
    mockDataRecipientPage.visitDynamicClientRegistrationTab() 
    mockDataRecipientPage.clickDCRRegisterButton()
    mockDataRecipientPage.assertThatClientRegistered()
    mockDataRecipientPage.visitConsentAndAuthorisationTab() 
    mockDataRecipientPage.selectClientRegistration(1)
    mockDataRecipientPage.inputSharingDuration(1000000)
    mockDataRecipientPage.clickConstructAuthorisationURI()
    acpLoginPage.login(Credentials.tppUsername, Credentials.defaultPassword)
    consentPage.confirm()
    consentPage.assertThatPageIsNotVisible()
    // verify data recipient mock consent callback page - UI error page improvements AUT-5845

    consentAdminPage.visit(true)
    Urls.clearLocalStorage()
    consentAdminPage.visit(true)
    consentAdminPage.login()

    consentAdminPage.assertThatConsentManagementTabIsDisplayed()
    consentAdminPage.revokeClientConsent();
  });
})
