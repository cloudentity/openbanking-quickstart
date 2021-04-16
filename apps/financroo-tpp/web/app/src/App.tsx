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

declare global {
  interface Window {
    featureFlags: any;
  }
}

window.featureFlags = window.featureFlags || {};

export type Config = {
  authorizationServerURL: string;
  authorizationServerId: string;
  clientId: string;
  tenantId: string;
};

const queryCache = new QueryCache();

const scopes = [];

const login = (data) => {
  if (data.token) {
    putTokenInStore(data.token);
    data.expires_in && putExpiresInInStore(data.expires_in);
    data.iat && putIATInInStore(data.iat);
    data.idToken && putIdTokenInStore(data.idToken);
  }
};

function App() {
  const [progress, setProgress] = useState(true);
  const [config, setConfig] = useState<Config>();

  useEffect(() => {
    superagent
      .get("/config.json")
      .then(toJson)
      .then((res) => setConfig(res))
      .finally(() => setProgress(false));
  }, []);

  return (
    <>
      <ThemeProvider theme={theme}>
        <StylesProvider injectFirst>
          <ReactQueryCacheProvider queryCache={queryCache}>
            <ReactQueryDevtools />
            {progress && <Progress />}
            {!progress && (
              <Router>
                <Suspense fallback={<Progress />}>
                  <Switch>
                    <Route
                      path="/callback"
                      render={() => (
                        <Callback
                          authorizationServerURL={
                            config?.authorizationServerURL
                          }
                          authorizationServerId={config?.authorizationServerId}
                          tenantId={config?.tenantId}
                          clientId={config?.clientId}
                          login={login}
                        />
                      )}
                    />
                    <Route
                      path="/silent"
                      render={() => (
                        <Callback
                          silent
                          authorizationServerURL={
                            config?.authorizationServerURL
                          }
                          authorizationServerId={config?.authorizationServerId}
                          tenantId={config?.tenantId}
                          clientId={config?.clientId}
                          login={login}
                        />
                      )}
                    />
                    <Route
                      path={"/auth"}
                      render={() => (
                        <AuthPage
                          login={login}
                          authorizationServerURL={
                            config?.authorizationServerURL
                          }
                          authorizationServerId={config?.authorizationServerId}
                          tenantId={config?.tenantId}
                          clientId={config?.clientId}
                          scopes={scopes}
                        />
                      )}
                    />
                    <PrivateRoute
                      path="/"
                      authorizationServerURL={config?.authorizationServerURL}
                      authorizationServerId={config?.authorizationServerId}
                      tenantId={config?.tenantId}
                      login={login}
                      component={() => (
                        <AuthenticatedAppBase
                          authorizationServerURL={
                            config?.authorizationServerURL
                          }
                          authorizationServerId={config?.authorizationServerId}
                          tenantId={config?.tenantId}
                          clientId={config?.clientId}
                          scopes={scopes}
                        />
                      )}
                    />
                    <Route component={() => <Redirect to={"/auth"} />} />
                  </Switch>
                </Suspense>
              </Router>
            )}
          </ReactQueryCacheProvider>
        </StylesProvider>
      </ThemeProvider>
    </>
  );
}

export default App;
