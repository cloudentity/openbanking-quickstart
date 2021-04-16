import superagent from 'superagent';

const authApi = {
  exchangeCodeForToken: ({tokenUri, body}) => superagent
    .post(tokenUri)
    .send(body)
    .then(res => res)
};

export default authApi;
