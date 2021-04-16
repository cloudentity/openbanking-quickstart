import React, { Suspense } from "react";
import { Switch } from "react-router";
import { Route } from "react-router-dom";
import ApplicationDetailsController from "./applicationDetails/ApplicationDetailsController";
import Dashboard from "./Dashboard";
import { useSilentAuthentication } from "./useSilentAuthentication";

export default function AuthenticatedAppBase({
  authorizationServerURL,
  authorizationServerId,
  tenantId,
  clientId,
  scopes,
  userinfo = {},
}) {
  useSilentAuthentication(
    authorizationServerURL,
    authorizationServerId,
    tenantId,
    clientId,
    scopes
  );

  return (
    <Suspense>
      <Switch>
        <Route
          exact
          path="/"
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
          exact
          path="/app/:id"
          render={() => (
            <ApplicationDetailsController
              authorizationServerURL={authorizationServerURL}
              authorizationServerId={authorizationServerId}
              tenantId={tenantId}
            />
          )}
        />
      </Switch>
    </Suspense>
  );
}
