import React, { Suspense, useEffect, useState } from "react";
import "./App.css";
import { BrowserRouter, Routes, Route } from "react-router-dom";
import { ThemeProvider, StyledEngineProvider } from "@mui/material/styles";

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

export type LoginData = {
  token: string | null;
  expires_in?: string;
  iat?: number;
  idToken?: string;
};

const scopes: string[] = [];

const login = (data: LoginData) => {
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
      <StyledEngineProvider injectFirst>
        <ThemeProvider theme={theme}>
          {progress && <Progress />}
          {!progress && (
            <BrowserRouter>
              <CommonProvider>
                <Snacks />
                <Suspense fallback={<Progress />}>
                  <Routes>
                    <Route
                      path="/callback"
                      element={
                        <Callback
                          authorizationServerURL={
                            config?.authorizationServerURL
                          }
                          authorizationServerId={config?.authorizationServerId}
                          tenantId={config?.tenantId}
                          clientId={config?.clientId}
                          login={login}
                        />
                      }
                    />
                    <Route
                      path="/silent"
                      element={
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
                      }
                    />
                    <Route
                      path="/auth"
                      element={
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
                      }
                    />
                    <Route
                      path="*"
                      element={
                        <PrivateRoute
                          authorizationServerURL={
                            config?.authorizationServerURL
                          }
                          authorizationServerId={config?.authorizationServerId}
                          tenantId={config?.tenantId}
                        />
                      }
                    >
                      <Route
                        path="*"
                        element={
                          <AuthenticatedAppBase
                            authorizationServerURL={
                              config?.authorizationServerURL
                            }
                            authorizationServerId={
                              config?.authorizationServerId
                            }
                            tenantId={config?.tenantId}
                            clientId={config?.clientId}
                            scopes={scopes}
                          />
                        }
                      />
                    </Route>
                  </Routes>
                </Suspense>
              </CommonProvider>
            </BrowserRouter>
          )}
        </ThemeProvider>
      </StyledEngineProvider>
    </>
  );
}

export default App;
