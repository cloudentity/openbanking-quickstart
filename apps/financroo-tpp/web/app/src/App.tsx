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
import { StylesProvider, ThemeProvider } from "@material-ui/core/styles";
import React, { Suspense } from "react";
import { QueryCache, ReactQueryCacheProvider } from "react-query";
import { ReactQueryDevtools } from "react-query-devtools";
import { Switch } from "react-router";
import { BrowserRouter as Router, Redirect, Route } from "react-router-dom";

declare global {
  interface Window {
    featureFlags: any;
    spec: any;
  }
}

window.featureFlags = window.featureFlags || {};
window.spec = window.spec || {};

const queryCache = new QueryCache();

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
    <ThemeProvider theme={theme}>
      <StylesProvider injectFirst>
        <ReactQueryCacheProvider queryCache={queryCache}>
          <ReactQueryDevtools />
          <Router>
            <Suspense fallback={<Progress />}>
              <Switch>
                <Route
                  path="/auth"
                  render={() => <AuthPage loginFn={login} />}
                />
                <PrivateRoute
                  path="/"
                  component={() => <AuthenticatedAppBase />}
                />
                <Route component={() => <Redirect to="/auth" />} />
              </Switch>
            </Suspense>
          </Router>
        </ReactQueryCacheProvider>
      </StylesProvider>
    </ThemeProvider>
  );
}

export default App;
