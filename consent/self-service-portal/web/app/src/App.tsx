import React, { Suspense, useEffect, useState } from "react";
import "./App.css";
import { BrowserRouter as Router, Redirect, Route } from "react-router-dom";
import { Switch } from "react-router";

import { StylesProvider, ThemeProvider } from "@material-ui/core/styles";
import superagent from "superagent";
import Progress from "./components/Progress";
import { toJson } from "./api/api-base";
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
import { CommonProvider } from "./services/common";
import { Snacks } from "./components/Snacks";

export type Config = {
  authorizationServerURL: string;
  authorizationServerId: string;
  clientId: string;
  tenantId: string;
};

const scopes = [];

const login = data => {
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
      .then(res => setConfig(res))
      .finally(() => setProgress(false));
  }, []);

  return (
    <>
      <ThemeProvider theme={theme}>
        <StylesProvider injectFirst>
          {progress && <Progress />}
          {!progress && (
            <Router>
              <CommonProvider>
                <Snacks />
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
                      path="/auth"
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
                    <Route component={() => <Redirect to="/auth" />} />
                  </Switch>
                </Suspense>
              </CommonProvider>
            </Router>
          )}
        </StylesProvider>
      </ThemeProvider>
    </>
  );
}

export default App;
