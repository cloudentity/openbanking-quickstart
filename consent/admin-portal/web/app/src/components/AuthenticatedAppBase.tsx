import React, { Suspense } from "react";
import { Switch } from "react-router";
import { Route } from "react-router-dom";
import Progress from "./Progress";
import AccountView from "./AccountView/AccountView";
import AccountViewDetails from "./AccountView/AccountViewDetails";
import ThirdPartyProvidersView from "./ThirdPartyProvidersView/ThirdPartyProvidersView";
import ConsentManagementView from "./ConsentManagementView/ConsentManagementView";

export default function AuthenticatedAppBase() {
  return (
    <div style={{ marginTop: 64 }}>
      <Suspense fallback={<Progress />}>
        <Switch>
          <Route
            path="/accounts/:id/apps/:clientId"
            render={() => (
              <AccountViewDetails />
            )}
          />
          <Route
            path="/accounts/:id"
            render={() => (
              <AccountView />
            )}
          />
          <Route
            path="/providers"
            render={() => (
              <ThirdPartyProvidersView />
            )}
          />
          <Route
            path="/"
            render={() => (
              <ConsentManagementView />
            )}
          />
        </Switch>
      </Suspense>
    </div>
  );
}
