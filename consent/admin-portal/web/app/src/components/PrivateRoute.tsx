import React from "react";
import { Navigate, Outlet, useLocation } from "react-router-dom";
import { isTokenInStore } from "./auth.utils";

export default function PrivateRoute() {
  const location = useLocation();
  return isTokenInStore() ? (
    <Outlet />
  ) : (
    <Navigate to="/auth" state={{ from: location }} replace />
  );
}
