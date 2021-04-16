import {base, baseWithCustomBaseUrl} from './api-base';

export const api = {
  fetchBanks: () => base.get({url: `/banks`}),
  fetchAccounts: () => base.get({url: `/accounts`}),
  connectBank: (bankId, body) => base.post({url: `/connect/${bankId}`, body, query: {}}),
  domesticPaymentConsent: (body) => base.post({url: `/domestic-payment-consent`, body, query: {}}),
  disconnectBank: (bankId) => base.delete({url: `/disconnect/${bankId}`, query: {}}),
  fetchTransactions: () => base.get({url: `/transactions`}),
  fetchBalances: () => base.get({url: `/balances`}),
  userinfo: (authorizationServerURL, tenantId, authorizationServerId) => baseWithCustomBaseUrl('/', authorizationServerURL).get({url: `/${tenantId}/${authorizationServerId}/userinfo`}),
};
