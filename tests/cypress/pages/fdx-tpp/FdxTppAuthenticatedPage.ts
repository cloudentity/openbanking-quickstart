export class FdxTppAuthenticatedPage {
  private readonly authenticatedPageInfoLabelLocator: string = `#fdx-auth-info-text p:nth-child(2)`;
  private readonly tokenResponseLocator: string = `#fdx-auth-token-response`;
  private readonly accessTokenLocator: string = `#fdx-auth-access-token`;
  private readonly consentResponseLocator: string = `#fdx-auth-consent-response`;
  private readonly tryNextButtonLocator: string = `#fdx-auth-try-again-button`;

  private expectedBankDetails: string =
    '{"homeUri":"https://www.midwest.com","logoUri":"https://www.midwest.com/81d88112572c.jpg",' +
    '"name":"MidwestPrimaryBank,NA","registeredEntityIdentifier":"549300ATG070THRDJ595",' +
    '"registeredEntityName":"MidwestPrimaryBank,NA","registryName":"GLEIF"}';

  private expectedAccountDetails = (id: string) =>
    `{"dataClusters":["ACCOUNT_DETAILED","TRANSACTIONS","STATEMENTS"],"id":"${id}","resourceType":"ACCOUNT"}`;

  public assertThatPageIsDisplayed(): void {
    cy.get(this.authenticatedPageInfoLabelLocator).should(
      "have.text",
      "User has been authenticated. Authorization code has been exchanged for an access token using mtls. " +
        "Implicit FDX consent has been created, see `grant_id` parameter"
    );
    cy.get(this.tryNextButtonLocator, { timeout: 3000 }).should("be.visible");
  }

  public assertThatTokenResponseFieldIsNotEmpty(): void {
    cy.get(this.tokenResponseLocator).then((element) => {
      let text: string = element.text();
      cy.log(text);
      expect(text).to.match(/\"access_token\"\: \"[a-zA-Z0-9_.-]+\"/);
      expect(text).to.match(/\"id_token\"\: \"[a-zA-Z0-9_.-]+\"/);
      expect(text).to.contain('"token_type": "bearer"');
      expect(text).to.match(/\"scope\"\: +[a-zA-Z_ "]+openid+[a-zA-Z_ ]*\"/);
      expect(text).to.match(/\"scope\"\: +[a-zA-Z_ "]+READ_CONSENTS+[a-zA-Z_ ]*\"/);
      expect(text).to.match(/\"scope\"\: +[a-zA-Z_ "]+UPDATE_CONSENTS+[a-zA-Z_ ]*\"/);
      expect(text).to.match(/\"expires_in\"\: [0-9]+\,/);
      expect(text).to.match(/\"grant_id\"\: \"[a-zA-Z0-9]+\"/);
    });
  }

  public assertThatAccessTokenFieldIsNotEmpty(): void {
    cy.get(this.accessTokenLocator).then((element) => {
      let text: string = element.text();
      cy.log(text);
      expect(text).to.contain('"acr": "1"');
      expect(text).to.match(/\"aid\"\: \"(?:fdx|[a-z0-9-]+fdx)\"/);
      expect(text).to.match(/\"spiffe\:\/\/[a-z0-9\/.-]+\/(?:fdx|[a-z0-9-]+fdx)\/(?:fdx-profile|[a-z0-9-]+fdx-profile)\"/);
      expect(text).to.match(/\"exp\"\: [0-9]+\,/);
      expect(text).to.match(/\"iat\"\: [0-9]+\,/);
      expect(text).to.match(/\"idp\"\: \"[a-zA-Z0-9]+\"/);
      expect(text).to.match(/\"iss\"\: \"https\:\/\/[a-z0-9\/:.-]+\/(?:fdx|[a-z0-9-]+fdx)\"/);
      expect(text).to.match(/\"jti\"\: \"[a-zA-Z0-9-]+\"/);
      expect(text).to.match(/\"nbf\"\: [0-9]+\,/);
      expect(text).to.contain('"st": "pairwise"');
      expect(text).to.match(/\"sub\"\: \"[a-zA-Z0-9]+\"/);
    });
  }

  public assertThatConsentResponseFieldContainsAccountsIds(
    accountsIDs: string[]
  ): void {
    cy.get(this.consentResponseLocator).then((element) => {
      let text: string = element.text().replace(/\s/g, "");
      cy.log(text);
      expect(text).to.contain(this.expectedBankDetails);
      accountsIDs.forEach((accountID) =>
        expect(text).to.contain(this.expectedAccountDetails(accountID))
      );
    });
  }

  public assertThatConsentResponseFieldNotContainsAccountsIds(
    accountsIDs: string[]
  ): void {
    cy.get(this.consentResponseLocator).then((element) => {
      let text: string = element.text().replace(/\s/g, "");
      cy.log(text);
      accountsIDs.forEach((accountID) =>
        expect(text).to.not.contain(this.expectedAccountDetails(accountID))
      );
    });
  }

  public clickTryNext(): void {
    cy.get(this.tryNextButtonLocator).click();
  }
}
