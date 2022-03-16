import React, { Suspense, useEffect, useState } from "react";
import "./App.css";
import { BrowserRouter as Router, Redirect, Route } from "react-router-dom";
import { Switch } from "react-router";
import { ThemeProvider } from "@material-ui/core";
import { StylesProvider } from "@material-ui/core/styles";
import superagent from "superagent";
import Progress from "./components/Progress";
import { toJson } from "./api/api-base";
import { ReactQueryDevtools } from "react-query-devtools";
import { QueryCache, ReactQueryCacheProvider } from "react-query";
import PrivateRoute from "./components/PrivateRoute";
import AuthPage from "./components/AuthPage";
import Callback from "./components/Callback";
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

const queryCache = new QueryCache();

const login = (data) => {
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
      <ThemeProvider theme={theme}>
        <StylesProvider injectFirst>
          <ReactQueryCacheProvider queryCache={queryCache}>
            <ReactQueryDevtools />
              <Router>
                <Suspense fallback={<Progress />}>
                  <Switch>
                    <Route
                      path={"/auth"}
                      render={() => (
                        <AuthPage login={login} />
                      )}
                    />
                    <PrivateRoute
                      path="/"
                      login={login}
                      component={() => (
                        <AuthenticatedAppBase />
                      )}
                    />
                    <Route component={() => <Redirect to={"/auth"} />} />
                  </Switch>
                </Suspense>
              </Router>
          </ReactQueryCacheProvider>
        </StylesProvider>
      </ThemeProvider>
    </>
  );
}

export default App;
