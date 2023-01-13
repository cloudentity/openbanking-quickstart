import React, { Fragment, useEffect, useState } from "react";
import { Navigate, Outlet, useLocation } from "react-router-dom";
import { api } from "../api/api";
import { isTokenInStore, removeAllAuthDataFromStore } from "./auth.utils";
import Progress from "./Progress";

interface Props {
  authorizationServerURL: string | undefined;
  tenantId: string | undefined;
  authorizationServerId: string | undefined;
}

export default function PrivateRoute({
  authorizationServerURL,
  tenantId,
  authorizationServerId,
}: Props) {
  const [progress, setProgress] = useState(true);
  const location = useLocation();

  useEffect(() => {
    api
      .userinfo(authorizationServerURL, tenantId, authorizationServerId)
      .catch(() => removeAllAuthDataFromStore())
      .finally(() => setProgress(false));
  }, [authorizationServerURL, tenantId, authorizationServerId]);

  return (
    <Fragment>
      {progress && <Progress />}
      {!progress &&
        (isTokenInStore() ? (
          <Outlet />
        ) : (
          <Navigate to="/auth" state={{ from: location }} replace />
        ))}
    </Fragment>
  );
}
