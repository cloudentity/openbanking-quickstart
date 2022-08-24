import React, { Suspense } from "react";
import "./App.css";
import { BrowserRouter, Routes, Route } from "react-router-dom";
import { ThemeProvider, StyledEngineProvider } from "@mui/material/styles";
import Progress from "./components/Progress";
import PrivateRoute from "./components/PrivateRoute";
import AuthPage from "./components/AuthPage";
import AuthenticatedAppBase from "./components/AuthenticatedAppBase";
import {
  putExpiresInInStore,
  putIATInInStore,
  putIdTokenInStore,
  putTokenInStore,
} from "./components/auth.utils";
import { theme } from "./theme";

export type Config = {
  authorizationServerURL: string;
  authorizationServerId: string;
  clientId: string;
  tenantId: string;
};

const login = data => {
  if (data.token) {
    putTokenInStore(data.token);
    data.expires_in && putExpiresInInStore(data.expires_in);
    data.iat && putIATInInStore(data.iat);
    data.idToken && putIdTokenInStore(data.idToken);
  }
};

function App() {
  return (
    <>
      <StyledEngineProvider injectFirst>
        <ThemeProvider theme={theme}>
          <BrowserRouter>
            <Suspense fallback={<Progress />}>
              <Routes>
                <Route path="/auth" element={<AuthPage login={login} />} />
                <Route path="*" element={<PrivateRoute />}>
                  <Route path="*" element={<AuthenticatedAppBase />} />
                </Route>
              </Routes>
            </Suspense>
          </BrowserRouter>
        </ThemeProvider>
      </StyledEngineProvider>
    </>
  );
}

export default App;
