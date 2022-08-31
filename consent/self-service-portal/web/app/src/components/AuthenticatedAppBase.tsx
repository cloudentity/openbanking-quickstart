import React, { Suspense } from "react";
import { Route, Routes } from "react-router-dom";
import ApplicationDetailsController from "./applicationDetails/ApplicationDetailsController";
import Dashboard from "./Dashboard";
import Progress from "./Progress";
import { useSilentAuthentication } from "./useSilentAuthentication";

interface Props {
  authorizationServerURL: string | undefined;
  authorizationServerId: string | undefined;
  tenantId: string | undefined;
  clientId: string | undefined;
  scopes: string[];
}

export default function AuthenticatedAppBase({
  authorizationServerURL,
  authorizationServerId,
  tenantId,
  clientId,
  scopes,
}: Props) {
  useSilentAuthentication(
    authorizationServerURL,
    authorizationServerId,
    tenantId,
    clientId,
    scopes
  );

  return (
    <Suspense fallback={<Progress />}>
      <Routes>
        <Route
          path="/"
          element={
            <Dashboard
              authorizationServerURL={authorizationServerURL}
              authorizationServerId={authorizationServerId}
              tenantId={tenantId}
            />
          }
        />
        <Route
          path="/app/:id"
          element={
            <ApplicationDetailsController
              authorizationServerURL={authorizationServerURL}
              authorizationServerId={authorizationServerId}
              tenantId={tenantId}
            />
          }
        />
      </Routes>
    </Suspense>
  );
}
