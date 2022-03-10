import './commands'


Cypress.Commands.add("disableSameSiteCookieRestrictions", () => {
  cy.intercept("*", (req) => {
    req.on("response", (res) => {
      if (!res.headers["set-cookie"]) {
        return;
      }

      const disableSameSite = (headerContent: string): string => {
        return headerContent.replace(
          /samesite=(lax|strict)/gi,
          "samesite=none"
        );
      };

      if (Array.isArray(res.headers["set-cookie"])) {
        res.headers["set-cookie"] = (res.headers["set-cookie"] as any).map(
          disableSameSite
        ) as any;
      } else {
        res.headers["set-cookie"] = disableSameSite(res.headers["set-cookie"]);
      }
    });
  });
});
