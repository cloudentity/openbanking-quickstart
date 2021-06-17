import React, { Suspense } from "react";
import { Switch } from "react-router";
import { Route } from "react-router-dom";
import { useSilentAuthentication } from "./useSilentAuthentication";
import Progress from "./Progress";
import AccountView from "./AccountView/AccountView";
import AccountViewDetails from "./AccountView/AccountViewDetails";
import ThirdPartyProvidersView from "./ThirdPartyProvidersView/ThirdPartyProvidersView";
import ConsentManagementView from "./ConsentManagementView/ConsentManagementView";

interface PropTypes {
  authorizationServerURL?: string;
  authorizationServerId?: string;
  tenantId?: string;
  clientId?: string;
  scopes: any;
}

export default function AuthenticatedAppBase({
  authorizationServerURL,
  authorizationServerId,
  tenantId,
  clientId,
  scopes,
}: PropTypes) {
  useSilentAuthentication(
    authorizationServerURL,
    authorizationServerId,
    tenantId,
    clientId,
    scopes
  );

  return (
    <div style={{ marginTop: 64 }}>
      <Suspense fallback={<Progress />}>
        <Switch>
          <Route
            path="/accounts/:id/apps/:clientId"
            render={() => (
              <AccountViewDetails
                authorizationServerURL={authorizationServerURL}
                authorizationServerId={authorizationServerId}
                tenantId={tenantId}
              />
            )}
          />
          <Route
            path="/accounts/:id"
            render={() => (
              <AccountView
                authorizationServerURL={authorizationServerURL}
                authorizationServerId={authorizationServerId}
                tenantId={tenantId}
              />
            )}
          />
          <Route
            path="/providers"
            render={() => (
              <ThirdPartyProvidersView
                authorizationServerURL={authorizationServerURL}
                authorizationServerId={authorizationServerId}
                tenantId={tenantId}
              />
            )}
          />
          <Route
            path="/"
            render={() => (
              <ConsentManagementView
                authorizationServerURL={authorizationServerURL}
                authorizationServerId={authorizationServerId}
                tenantId={tenantId}
              />
            )}
          />
        </Switch>
      </Suspense>
    </div>
  );
}
