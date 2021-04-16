import superagent from 'superagent';
// import {unauthorizedRedirect} from './unauthorized-redirect.interceptor';
const getTokenFromStore = () => localStorage.getItem('access_token');

export const toJson = response => response.text ? JSON.parse(response.text) : {};

const normalize = path => path.replace(/\/\/+/g, '/');

const buildUrl = (origin, base, path) => origin + normalize(base + '/' + path);

const accessTokenHeader = () => ['Authorization', `Bearer ${getTokenFromStore()}`];

const http = (request, origin, baseUrl) => ({
  get: ({url, query = null, responseType = null, callback = toJson}) => request
    .get(buildUrl(origin, baseUrl, url))
    // .use(unauthorizedRedirect)
    .query(query)
    .responseType(responseType)
    .set(...accessTokenHeader())
    .then(callback),
  post: ({url, body, query}) => request
    .post(buildUrl(origin, baseUrl, url))
    // .use(unauthorizedRedirect)
    .query(query)
    .send(body)
    .set(...accessTokenHeader())
    .then(toJson),
  put: ({url, body}) => request
    .put(buildUrl(origin, baseUrl, url))
    .send(body)
    // .use(unauthorizedRedirect)
    .set(...accessTokenHeader())
    .then(toJson),
  delete: ({url, query}) => request
    .delete(buildUrl(origin, baseUrl, url))
    // .use(unauthorizedRedirect)
    .query(query)
    .set(...accessTokenHeader())
    .then(toJson),
  postFormData: ({url, formData, query}) => request
    .post(buildUrl(origin, baseUrl, url))
    // .use(unauthorizedRedirect)
    .query(query)
    .set(...accessTokenHeader())
    .send(formData)
    .then(toJson)
});

export default http(superagent, window.location.origin, `/api`);

export const base = http(superagent, window.location.origin, `/api`);

export const baseWithCustomBaseUrl = (baseUrl, origin = window.location.origin) => http(superagent, origin, baseUrl);
