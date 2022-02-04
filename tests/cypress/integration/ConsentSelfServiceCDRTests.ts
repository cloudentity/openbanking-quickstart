import {AcpLoginPage} from '../pages/acp/AcpLoginPage';
import {ConsentPage} from '../pages/consent/ConsentPage';
import {Credentials} from "../pages/Credentials";
import {ConsentSelfServicePage} from '../pages/consent-self-service/ConsentSelfServicePage';
import {ConsentSelfServiceApplicationPage} from "../pages/consent-self-service/ConsentSelfServiceApplicationPage";
import { MockDataRecipientPage } from '../pages/mock-data-recipient/MockDataRecipientPage';

describe(`Consent self service app CDR`, () => {
  const acpLoginPage: AcpLoginPage = new AcpLoginPage();
  const consentPage: ConsentPage = new ConsentPage();
  const consentSelfServicePage: ConsentSelfServicePage = new ConsentSelfServicePage();
  const consentSelfServiceApplicationPage: ConsentSelfServiceApplicationPage = new ConsentSelfServiceApplicationPage();
  const mockDataRecipientPage: MockDataRecipientPage = new MockDataRecipientPage(); 

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

  beforeEach(() => {
    consentSelfServicePage.visit(true)
  })

  it(`Happy path with account consent`, () => {
    acpLoginPage.login(Credentials.tppUsername, Credentials.defaultPassword);
    consentSelfServicePage.clickOnApplicationCard()
    consentSelfServiceApplicationPage.expandAccountsTab()
    consentSelfServiceApplicationPage.checkAccount("1000001")
    consentSelfServiceApplicationPage.checkAccount("1000002")
    consentSelfServiceApplicationPage.expandAccountConsentRow()
    consentSelfServiceApplicationPage.assertAccountRevokePopupContainsText('Your Name and occupation'); 
  })

  it(`Revoke CDR Arrangement`, () => {
    acpLoginPage.login(Credentials.tppUsername, Credentials.defaultPassword);
    consentSelfServicePage.clickOnApplicationCard()
    consentSelfServiceApplicationPage.expandAccountsTab()
    consentSelfServiceApplicationPage.assertNumberOfConsents(1)
    consentSelfServiceApplicationPage.expandAccountConsentRow()
    consentSelfServiceApplicationPage.clickRevokeAccessButton() 
    consentSelfServiceApplicationPage.assertNumberOfConsents(0)
  })
})