import Button from "@mui/material/Button";
import Card from "@mui/material/Card";
import Typography from "@mui/material/Typography";
import ErrorIcon from "@mui/icons-material/ErrorOutline";
import { pathOr } from "ramda";
import React, { useEffect, useState } from "react";
import { Navigate, useNavigate } from "react-router-dom";
import authApi from "./authApi";
import {
  SILENT_AUTH_ERROR_MESSAGE,
  SILENT_AUTH_SUCCESS_MESSAGE,
} from "./AuthPage";
import Progress from "./Progress";
import { LoginData } from "../App";

const getParamFromUrl = param => {
  const match = RegExp(`${param}=([^&]*)`, "g").exec(window.location.href);
  return match && match[1];
};

const capitalizeFirstLetter = (string = "") => {
  return string.charAt(0).toUpperCase() + string.slice(1);
};

const errorCodeToDisplay = error =>
  capitalizeFirstLetter((error || "").replace(/(\+|_)/g, " "));

interface Props {
  authorizationServerURL: string | undefined;
  tenantId: string | undefined;
  authorizationServerId: string | undefined;
  clientId: string | undefined;
  login: (data: LoginData) => void;
  silent?: boolean;
}

export default function Callback({
  authorizationServerURL,
  tenantId,
  authorizationServerId,
  clientId,
  login,
  silent = false,
}: Props) {
  const navigate = useNavigate();

  const [exchangeCompleted, setExchangeCompleted] = useState(false);
  const [error, setError] = useState<{
    error: string;
    errorDescription: string;
    errorHint: string;
  } | null>(null);

  useEffect(() => {
    const PKCE_STATE_KEY = `pkce_state`;
    const PKCE_CODE_VERIFIER_KEY = `pkce_code_verifier`;
    const code = getParamFromUrl("code");
    const errorParam = getParamFromUrl("error");

    const authorizationUri = `${authorizationServerURL}/${tenantId}/${authorizationServerId}/oauth2/authorize`;
    const tokenUri = `${authorizationServerURL}/${tenantId}/${authorizationServerId}/oauth2/token`;

    const config = {
      client_id: authorizationServerId,
      redirect_uri:
        window.location.origin + `/${silent ? "silent" : "callback"}`,
      authorization_endpoint: authorizationUri,
    };

    const body = `grant_type=authorization_code&client_id=${clientId}&redirect_uri=${
      config.redirect_uri
    }&code=${code}&code_verifier=${localStorage.getItem(
      PKCE_CODE_VERIFIER_KEY
    )}`;

    if (code) {
      authApi
        .exchangeCodeForToken(tokenUri, { body })
        .then(res => {
          const currentInSec = new Date().getTime() / 1000;
          login({
            token: res.body.access_token,
            iat: currentInSec,
            expires_in: currentInSec + res.body.expires_in,
            idToken: res.body.id_token,
          });
          localStorage.removeItem(PKCE_STATE_KEY);
          localStorage.removeItem(PKCE_CODE_VERIFIER_KEY);
          setExchangeCompleted(true);
          if (silent) {
            window.parent.postMessage(
              SILENT_AUTH_SUCCESS_MESSAGE,
              window.location.origin
            );
          }
        })
        .catch(err => {
          localStorage.removeItem(PKCE_STATE_KEY);
          localStorage.removeItem(PKCE_CODE_VERIFIER_KEY);
          setError({
            error: errorCodeToDisplay(
              pathOr("", ["response", "body", "error"], err)
            ),
            errorDescription: pathOr(
              "",
              ["response", "body", "error_description"],
              err
            ),
            errorHint: pathOr("", ["response", "body", "error_hint"], err),
          });
          if (silent) {
            window.parent.postMessage(
              SILENT_AUTH_ERROR_MESSAGE,
              window.location.origin
            );
          }
        });
    }

    if (silent && !code && errorParam === "login_required") {
      window.parent.postMessage(
        SILENT_AUTH_ERROR_MESSAGE,
        window.location.origin
      );
      return;
    }

    if (errorParam) {
      const errorDescription = (
        getParamFromUrl("error_description") || ""
      ).replace(/\+/g, " ");
      const error = errorCodeToDisplay(errorParam);
      setError({ error, errorDescription, errorHint: "" });
    }
  }, [
    tenantId,
    authorizationServerId,
    authorizationServerURL,
    clientId,
    login,
    silent,
  ]);

  if (error) {
    return (
      <>
        <Card
          style={{
            width: 500,
            padding: 30,
            margin: "0 auto",
            marginTop: 120,
            display: "flex",
            flexDirection: "column",
            alignItems: "center",
          }}
        >
          <ErrorIcon style={{ fontSize: 120, opacity: 0.1 }} />
          <Typography variant="h5" style={{ marginTop: 24 }}>
            {error.error}
          </Typography>
          <Typography
            variant="subtitle1"
            style={{ opacity: 0.6, marginTop: 12 }}
          >
            {error.errorDescription}
          </Typography>
          <div style={{ textAlign: "left", width: "100%", marginTop: 12 }}>
            {error.errorHint && (
              <Typography variant="caption" style={{ opacity: 0.6 }}>
                Hint: {error.errorHint}
              </Typography>
            )}
          </div>
          <div style={{ textAlign: "right", marginTop: 32 }}>
            <Button
              variant="contained"
              color="primary"
              onClick={() => navigate("/")}
            >
              Try again
            </Button>
          </div>
        </Card>
      </>
    );
  }

  if (exchangeCompleted && !silent) {
    return <Navigate to="/" replace />;
  }

  return <Progress />;
}
