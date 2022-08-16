import "./App.css";
import AuthPage from "./components/AuthPage";
import AuthenticatedAppBase from "./components/AuthenticatedAppBase";
import PrivateRoute from "./components/PrivateRoute";
import Progress from "./components/Progress";
import {
  putExpiresInInStore,
  putIATInInStore,
  putIdTokenInStore,
  putTokenInStore,
} from "./components/auth.utils";
import { theme } from "./theme";
import { ThemeProvider, StyledEngineProvider } from "@mui/material/styles";

import React, { Suspense } from "react";
import { QueryClient, QueryClientProvider } from "react-query";
import { ReactQueryDevtools } from "react-query/devtools";
import { BrowserRouter as Router, Route, Routes } from "react-router-dom";

declare global {
  interface Window {
    featureFlags: any;
    spec: any;
  }
}

window.featureFlags = window.featureFlags || {};
window.spec = window.spec || {};

const queryClient = new QueryClient();

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
    <QueryClientProvider client={queryClient}>
      <StyledEngineProvider injectFirst>
        <ThemeProvider theme={theme}>
          <ReactQueryDevtools />
          <Router>
            <Suspense fallback={<Progress />}>
              <Routes>
                <Route path="/auth" element={<AuthPage loginFn={login} />} />
                <Route path="*" element={<PrivateRoute />}>
                  <Route path="*" element={<AuthenticatedAppBase />} />
                </Route>
              </Routes>
            </Suspense>
          </Router>
        </ThemeProvider>
      </StyledEngineProvider>
    </QueryClientProvider>
  );
}

export default App;
