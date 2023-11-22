// FIXME
// describe(`Smoke Financroo Consent self service tests`, () => {
//   const financrooLoginPage: FinancrooLoginPage = new FinancrooLoginPage();
//   const financrooWelcomePage: FinancrooWelcomePage = new FinancrooWelcomePage();
//   const acpLoginPage: AcpLoginPage = new AcpLoginPage();
//   const accountConsentPage: AccountConsentPage = new AccountConsentPage();
//   const financrooModalPage: FinancrooModalPage = new FinancrooModalPage();
//   const financrooAccountsPage: FinancrooAccountsPage = new FinancrooAccountsPage();
//   const consentSelfServicePage: ConsentSelfServicePage = new ConsentSelfServicePage();
//   const consentSelfServiceApplicationPage: ConsentSelfServiceApplicationPage = new ConsentSelfServiceApplicationPage();
//   const consentSelfServiceAccountDetailsPage: ConsentSelfServiceAccountDetailsPage = new ConsentSelfServiceAccountDetailsPage();

//   const accountsIDs = [Accounts.ids.CDR.savings, Accounts.ids.CDR.checking];


//   before(() => {
//     financrooLoginPage.visit();
//     financrooLoginPage.login();

//     financrooWelcomePage.reconnectGoBank();

//     acpLoginPage.assertThatModalIsDisplayed();
//     acpLoginPage.loginWithMfaOption();

//     accountConsentPage.checkAccounts(accountsIDs);
//     accountConsentPage.expandPermissions();
//     accountConsentPage.assertPermissionsDetails(
//       "Purpose for sharing data",
//       "To uncover insights that can improve your financial well being."
//     );
//     accountConsentPage.clickAgree();

//     financrooModalPage.assertThatModalIsDisplayed();
//     financrooModalPage.close();

//     financrooAccountsPage.assertThatPageIsDisplayed();
//     financrooAccountsPage.assertAccountsSyncedNumber(accountsIDs.length);
//     financrooAccountsPage.assertAccountsIds(accountsIDs);
//   });

//   beforeEach(() => {
//     consentSelfServicePage.visit(true);

//     acpLoginPage.assertThatModalIsDisplayed();
//     acpLoginPage.login();

//     consentSelfServicePage.clickOnApplicationCard();
//   });

//   it(`Happy path with account consent`, () => {
//     consentSelfServiceApplicationPage.expandAccountsTab();
//     consentSelfServiceApplicationPage.checkAccount(accountsIDs[0]);
//     consentSelfServiceApplicationPage.checkAccount(accountsIDs[1]);
//     consentSelfServiceApplicationPage.expandAccountConsentRow();

//     consentSelfServiceAccountDetailsPage.assertThatAccountDetailsAreVisible()
//     consentSelfServiceAccountDetailsPage.assertAccount(accountsIDs[0]);
//     consentSelfServiceAccountDetailsPage.assertAccount(accountsIDs[1]);
//   });

//   it(`Revoke account consent`, () => {
//     consentSelfServiceApplicationPage.expandAccountsTab();
//     consentSelfServiceApplicationPage.assertAuthorisedAccountRowExists(accountsIDs[0]);
//     consentSelfServiceApplicationPage.assertAuthorisedAccountRowExists(accountsIDs[1]);
//     consentSelfServiceApplicationPage.expandAccountConsentRow();

//     consentSelfServiceAccountDetailsPage.assertThatAccountDetailsAreVisible();
//     consentSelfServiceAccountDetailsPage.clickRevokeAccessButton();
//     consentSelfServiceAccountDetailsPage.assertThatRevokeAccountDetailsAreVisible();
//     consentSelfServiceAccountDetailsPage.confirmRevokeAccessAction();

//     consentSelfServicePage.clickOnApplicationCardWithName("Financroo");
//     consentSelfServiceApplicationPage.assertAuthorisedAccountRowDoesNotExist(accountsIDs[0]);
//     consentSelfServiceApplicationPage.assertAuthorisedAccountRowDoesNotExist(accountsIDs[1]);

//     financrooLoginPage.visit();
//     financrooLoginPage.login();

//     financrooAccountsPage.assertThatAccountsAreDisconnected();
//   });

// });
