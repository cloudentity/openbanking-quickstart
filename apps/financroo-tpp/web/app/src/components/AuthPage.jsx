import React from 'react';
import {Redirect} from 'react-router';

import {getTokenFromStore, isTokenInStore, removeAllAuthDataFromStore} from './auth.utils';
import {generateRandomString, pkceChallengeFromVerifier} from './pkce.utils';
import Register from './Register';

const calcAuthorizationUrl = async (authorizationServerURL, tenantId, authorizationServerId, clientId, scopes = [], silent = false, idTokenHint = "") => {
  const authorizationUri = `${authorizationServerURL}/${tenantId}/${authorizationServerId}/oauth2/authorize`;

  // Create and store a random "state" value
  const state = generateRandomString();
  localStorage.setItem(`pkce_state`, state);

  // Create and store a new PKCE code_verifier (the plaintext random secret)
  const code_verifier = generateRandomString();
  localStorage.setItem(`pkce_code_verifier`, code_verifier);

  // Hash and base64-urlencode the secret to use as the challenge
  const code_challenge = await pkceChallengeFromVerifier(code_verifier);

  return authorizationUri
    + "?response_type=code"
    + "&client_id=" + encodeURIComponent(clientId)
    + "&state=" + encodeURIComponent(state)
    + "&scope=" + encodeURIComponent(scopes.join(' '))
    + "&redirect_uri=" + encodeURIComponent(window.location.origin + `/${silent ? 'silent' : 'callback'}`)
    + "&code_challenge=" + encodeURIComponent(code_challenge)
    + "&code_challenge_method=S256"
    + `${silent ? `&prompt=none&id_token_hint=${idTokenHint}` : ''}`
}

export const authorize = async (authorizationServerURL, tenantId, authorizationServerId, clientId, scopes = []) => {

  // Authorization URL
  window.location.href = await calcAuthorizationUrl(authorizationServerURL, tenantId, authorizationServerId, clientId, scopes)
};

const IFRAME_ID = 'silent-auth-iframe';
export const SILENT_AUTH_SUCCESS_MESSAGE = 'silentAuthSuccess';
export const SILENT_AUTH_ERROR_MESSAGE = 'silentAuthFailure';

export const silentAuthentication = async (authorizationServerURL, tenantId, authorizationServerId, clientId, scopes, idTokenHint) => {
  const iframe = document.createElement("iframe");
  const src = await calcAuthorizationUrl(authorizationServerURL, tenantId, authorizationServerId, clientId, scopes, true, idTokenHint);
  iframe.setAttribute("src", src)
  iframe.setAttribute("id", IFRAME_ID)
  iframe.style.display = 'none';

  const listener = e => {
    if (e.data === SILENT_AUTH_SUCCESS_MESSAGE || e.data === SILENT_AUTH_ERROR_MESSAGE) {
      const iframeToRemove = document.querySelector(`#${IFRAME_ID}`);
      iframeToRemove && document.body.removeChild(iframeToRemove);
      window.removeEventListener('message', listener);
    }
  };

  window.addEventListener("message", listener);

  document.body.appendChild(iframe);
}

export const logout = (authorizationServerURL, tenantId, authorizationServerId) => {
  removeAllAuthDataFromStore();
  window.location.href = `${authorizationServerURL}/${tenantId}/${authorizationServerId}/logout?redirect_to=${window.location.origin}`
};


const AuthPage = ({login, authorizationServerURL, tenantId, authorizationServerId, clientId, scopes}) => {
  const handleLogin = () => {
    authorize(authorizationServerURL, tenantId, authorizationServerId, clientId, scopes);
  }

  if (isTokenInStore()) {
    login({token: getTokenFromStore()});
    return (<Redirect to={'/'}/>)
  }

  return <Register onLogin={handleLogin}/>
};


export default AuthPage;
