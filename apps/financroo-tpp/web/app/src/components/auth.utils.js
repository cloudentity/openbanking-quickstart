let ACCESS_TOKEN_KEY = `access_token`;
let EXPIRES_IN_KEY = `expires_in`;
let IAT_KEY = `iat`;
let ID_TOKEN_KEY = `id_token`;

export const isTokenInStore = () => !!localStorage.getItem(ACCESS_TOKEN_KEY);
export const getTokenFromStore = () => localStorage.getItem(ACCESS_TOKEN_KEY);
export const putTokenInStore = (token, name = ACCESS_TOKEN_KEY) => localStorage.setItem(name, token);
export const removeTokenFromStore = () => localStorage.removeItem(ACCESS_TOKEN_KEY);

export const getExpiresInFromStore = () => localStorage.getItem(EXPIRES_IN_KEY);
export const putExpiresInInStore = (token, name = EXPIRES_IN_KEY) => localStorage.setItem(name, token);
export const removeExpiresInFromStore = () => localStorage.removeItem(EXPIRES_IN_KEY);

export const getIATInFromStore = () => localStorage.getItem(IAT_KEY);
export const putIATInInStore = (token, name = IAT_KEY) => localStorage.setItem(name, token);
export const removeIATFromStore = () => localStorage.removeItem(IAT_KEY);

export const getIdTokenFromStore = () => localStorage.getItem(ID_TOKEN_KEY);
export const putIdTokenInStore = (token, name = ID_TOKEN_KEY) => localStorage.setItem(name, token);
export const removeIdTokenFromStore = () => localStorage.removeItem(ID_TOKEN_KEY);

export const removeAllAuthDataFromStore = () => {
  removeTokenFromStore();
  removeExpiresInFromStore();
  removeIATFromStore();
  removeIdTokenFromStore();
}
