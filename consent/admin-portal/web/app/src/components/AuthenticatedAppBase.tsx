import React, { Suspense } from "react";
import { Routes, Route, Navigate } from "react-router-dom";
import Progress from "./Progress";
import AccountView from "./AccountView/AccountView";
import AccountViewDetails from "./AccountView/AccountViewDetails";
import ThirdPartyProvidersView from "./ThirdPartyProvidersView/ThirdPartyProvidersView";
import ConsentManagementView from "./ConsentManagementView/ConsentManagementView";

export default function AuthenticatedAppBase() {
  return (
    <div style={{ marginTop: 64 }}>
      <Suspense fallback={<Progress />}>
        <Routes>
          <Route
            path="/accounts/:id/apps/:clientId"
            element={<AccountViewDetails />}
          />
          <Route path="/accounts/:id" element={<AccountView />} />
          <Route path="/providers" element={<ThirdPartyProvidersView />} />
          <Route path="/" element={<ConsentManagementView />} />
          <Route path="*" element={<Navigate to="/" replace />} />
        </Routes>
      </Suspense>
    </div>
  );
}
