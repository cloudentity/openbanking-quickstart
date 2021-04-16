import {useEffect} from 'react';
import {getExpiresInFromStore, getIATInFromStore, getIdTokenFromStore, getTokenFromStore} from './auth.utils';
import throttle from 'lodash.throttle';
import {silentAuthentication} from './AuthPage';

const TIMEOUT_RATIO_FACTOR = 0.75;
const TOKEN_EXPIRATION_RATIO_CHECK_INTERVAL = 5000;

const silentAuthenticationThrottled = throttle(silentAuthentication, 30000, {trailing: false});

export const useSilentAuthentication = (authorizationServerURL, authorizationServerId, tenantId, clientId, scopes) =>
  useEffect(() => {
    const counter = setInterval(async () => {
      const token = getTokenFromStore();
      const current = new Date().getTime() / 1000;
      const expiresIn = parseInt(getExpiresInFromStore());
      const iat = parseInt(getIATInFromStore());
      const idTokenHint = getIdTokenFromStore() || '';

      const lifetimeInSec = expiresIn - iat;
      const validForInSec = expiresIn - current;
      const ratio = (lifetimeInSec - validForInSec) / lifetimeInSec;

      if (ratio > TIMEOUT_RATIO_FACTOR || !token) {
        silentAuthenticationThrottled(authorizationServerURL, tenantId, authorizationServerId, clientId, scopes, idTokenHint);
      }
    }, TOKEN_EXPIRATION_RATIO_CHECK_INTERVAL);

    return () => clearInterval(counter);
  }, [authorizationServerURL, authorizationServerId, tenantId, clientId, scopes]);
