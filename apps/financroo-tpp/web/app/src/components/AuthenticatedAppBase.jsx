import React, { Suspense } from "react";
import { Switch } from "react-router";
import { Route } from "react-router-dom";

import Dashboard from "./Dashboard";
import Investments from "./investments/Investments";
import InvestmentsContribute from "./investments/InvestmentsContribute";
import InvestmentsContributeSuccess from "./investments/InvestmentsContributeSuccess";

export default function AuthenticatedAppBase({
  authorizationServerURL,
  authorizationServerId,
  tenantId,
  clientId,
  scopes,
  userinfo = {},
}) {
  return (
    <Suspense>
      <Switch>
        <Route
          exact
          path={"/"}
          render={() => (
            <Dashboard
              authorizationServerURL={authorizationServerURL}
              authorizationServerId={authorizationServerId}
              tenantId={tenantId}
              userinfo={userinfo}
            />
          )}
        />
        <Route
          path={"/investments"}
          exact
          render={() => (
            <Investments
              authorizationServerURL={authorizationServerURL}
              authorizationServerId={authorizationServerId}
              tenantId={tenantId}
            />
          )}
        />
        <Route
          path={"/investments/contribute"}
          exact
          render={() => <InvestmentsContribute />}
        />

        <Route
          path={"/investments/contribute/:id/success"}
          exact
          render={() => <InvestmentsContributeSuccess />}
        />
      </Switch>
    </Suspense>
  );
}
